package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/rand"
	"regexp"
	"sync"
	"time"
	"word-count/cmd/api/rpc"
	"word-count/cmd/map/pack"
	"word-count/config"
	"word-count/kitex_gen/keybydemo"
	"word-count/kitex_gen/mapdemo"
	"word-count/pkg/errno"
)

// MapServiceImpl implements the last service interface defined in the IDL.
type MapServiceImpl struct{}

var (
	mapNum      int
	maxNum      int
	mapIdx      int
	mapChanList []chan string
	keyValue    int64
	baizeRegexp *regexp.Regexp
	lb          string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// 由于是单进程模型多协程，因此上游的负载均衡策略本质上就是发往下游哪个算子，这里对应就是选择哪个map协程处理上游source的数据
	lb = fmt.Sprintf(config.GetConfig().Map.Lb)
	baizeRegexp = regexp.MustCompile(config.GlobalDAGConfig.GetString("map.regexp"))
	keyValue = config.GlobalDAGConfig.GetInt64("map.value")
	mapNum = config.GlobalDAGConfig.GetInt("map.num")
	maxNum = config.GetConfig().Map.Num
	// DAG中map算子数量非法状态
	if mapNum <= 0 || mapNum > maxNum {
		klog.Fatal("DAG中指定的map数量过小/过大")
	}
	mapIdx = 0
	wg := sync.WaitGroup{}
	mapChanList = make([]chan string, 0)
	for i := 1; i <= mapNum; i++ {
		wg.Add(1)
		ch := make(chan string, 1)
		mapChanList = append(mapChanList, ch)
		go func(chan string, *regexp.Regexp, int64) {
			// keyby服务客户端初始化
			rpc.InitKeybyRPC()
			for {
				// 获取source算子发过来的数据，如果没有消息会被阻塞
				msg, _ := <-ch
				// 使用DAG中描述的规则解析DataStream数据（由于历史原因source的数据数据中content是string，事实上应该封装好DataStream结构，这里从map算子往下游通信开始封装）
				params := baizeRegexp.FindStringSubmatch(msg)

				klog.Info(fmt.Sprintf("当前处理的是map算子%d号，消息内容为%s", mapIdx, msg))

				content := make([]string, 0)

				klog.Info(params)

				for i := 1; i < len(params)-1; i++ {
					content = append(content, params[i])
				}
				// 消费数据，调用keyby算子的client方法，传递DataStream结构
				err := rpc.CreateKeyby(context.Background(), &keybydemo.CreateKeybyRequest{
					content,
					keyValue,
					params[len(params)-1],
				})
				if err != nil {
					klog.Fatal("call keyby service failed")
				}
			}

		}(ch, baizeRegexp, keyValue)
	}
}

func HandleSourceMsg(req *mapdemo.CreateMapRequest) {
	//klog.Info(req.Content)
	// 负载均衡规则
	switch lb {
	case "round":
		mapIdx++
		mapIdx %= mapNum
		mapChanList[mapIdx] <- req.Content
	case "random":
		mapIdx = rand.Intn(mapNum)
		mapChanList[mapIdx] <- req.Content
	}
}

// CreateMap implements the MapServiceImpl interface.
func (s *MapServiceImpl) CreateMap(ctx context.Context, req *mapdemo.CreateMapRequest) (resp *mapdemo.CreateMapResponse, err error) {

	HandleSourceMsg(req)

	resp = new(mapdemo.CreateMapResponse)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
