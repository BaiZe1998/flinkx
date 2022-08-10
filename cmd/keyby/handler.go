package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/rand"
	"sync"
	"time"
	"word-count/cmd/keyby/pack"
	"word-count/config"
	"word-count/kitex_gen/keybydemo"
	"word-count/pkg/errno"
)

// KeybyServiceImpl implements the last service interface defined in the IDL.
type KeybyServiceImpl struct{}

var (
	keybyNum      int
	maxNum        int
	keybyIdx      int
	keybyChanList []chan *keybydemo.CreateKeybyRequest
	keyIdx        int
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// 由于是单进程模型多协程，因此上游的负载均衡策略本质上就是发往下游哪个算子，这里对应就是选择哪个map协程处理上游source的数据
	keyIdx = config.GlobalDAGConfig.GetInt("keyby.key")
	keybyNum = config.GlobalDAGConfig.GetInt("keyby.num")
	maxNum = config.GetConfig().Keyby.Num
	// DAG中map算子数量非法状态
	if keybyNum <= 0 || keybyNum > maxNum {
		klog.Fatal("DAG中指定的keyby数量过小/过大")
	}
	keybyIdx = 0
	wg := sync.WaitGroup{}
	keybyChanList = make([]chan *keybydemo.CreateKeybyRequest, 0)
	for i := 1; i <= keybyNum; i++ {
		wg.Add(1)
		ch := make(chan *keybydemo.CreateKeybyRequest, 1)
		keybyChanList = append(keybyChanList, ch)
		go func(chan *keybydemo.CreateKeybyRequest, *sync.WaitGroup, int) {
			// reduce服务客户端初始化
			//rpc.InitReduceRPC()
			for {
				// 获取map算子发过来的数据，如果没有消息会被阻塞
				msg, _ := <-ch
				// TODO 根据key进行分区，并且根据window的参数配置，创建窗口，进行水位线的定时获取

				klog.Info(fmt.Sprintf("当前处理的是keyby算子%d号，消息内容为%s", keybyIdx, msg))

			}
			wg.Done()
		}(ch, &wg, keyIdx)
	}
	wg.Wait()
}

func HandleKeybyMsg(req *keybydemo.CreateKeybyRequest) {
	// 需要根据keyIdx对应的值进行hash为不同分区，然后需要分发到对应keyby算子
	key := req.Content[keyIdx]
	keybyIdx = int(key[0]) % keybyNum
	keybyChanList[keybyIdx] <- req
}

// CreateKeyby implements the KeybyServiceImpl interface.
func (s *KeybyServiceImpl) CreateKeyby(ctx context.Context, req *keybydemo.CreateKeybyRequest) (resp *keybydemo.CreateKeybyResponse, err error) {
	klog.Info(req.Content, req.TimeStamp, req.Value)

	// TODO 根据key的位置，选择key值，并且根据k进行hash，分发给两个keyby协程算子，并且每个算子维护m个滚动窗口（每个窗口有时间段之分），并且结合水位线，此时在keyby算子上就能实现5分钟的窗口聚合
	// TODO 同时数据不断发送给下游reduce，进行retract的归约（是否也需要按照窗口的呃分钟进行一次？还是说是每个单词来了归一次）

	//HandleKeybyMsg(req)

	resp = new(keybydemo.CreateKeybyResponse)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
