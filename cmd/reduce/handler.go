package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/rand"
	"strconv"
	"time"
	"word-count/cmd/api/rpc"
	"word-count/cmd/reduce/pack"
	"word-count/cmd/sink/kitex_gen/sinkdemo"
	"word-count/config"
	"word-count/kitex_gen/reducedemo"
	"word-count/pkg/errno"
)

// ReduceServiceImpl implements the last service interface defined in the IDL.
type ReduceServiceImpl struct{}

var (
	reduceNum      int
	maxNum         int
	reduceIdx      int
	reduceChanList []chan *reducedemo.CreateReduceRequest
	lb             string
	tables         []string
	windowAssign   string
)

func init() {
	// 由于是单进程模型多协程，因此上游的负载均衡策略本质上就是发往下游哪个算子，这里对应就是选择哪个map协程处理上游source的数据
	lb = fmt.Sprintf(config.GetConfig().Map.Lb)
	reduceNum = config.GlobalDAGConfig.GetInt("reduce.num")
	tables = config.GlobalDAGConfig.GetStringSlice("window.table")
	windowAssign = config.GlobalDAGConfig.GetString("window.assign")
	maxNum = config.GetConfig().Reduce.Num
	// DAG中map算子数量非法状态
	if reduceNum <= 0 || reduceNum > maxNum {
		klog.Fatal("DAG中指定的map数量过小/过大")
	}
	reduceIdx = 0
	reduceChanList = make([]chan *reducedemo.CreateReduceRequest, 0)
	for i := 1; i <= reduceNum; i++ {
		ch := make(chan *reducedemo.CreateReduceRequest, 1)
		reduceChanList = append(reduceChanList, ch)
		go func(chan *reducedemo.CreateReduceRequest) {
			// keyby服务客户端初始化
			rpc.InitSinkRPC()

			tableMap1 := make(map[string]int64)
			tableMap2 := make(map[string]int64)
			var preTime *time.Time

			for {
				// 获取source算子发过来的数据，如果没有消息会被阻塞
				msg, _ := <-ch
				// 使用DAG中描述的规则解析DataStream数据（由于历史原因source的数据数据中content是string，事实上应该封装好DataStream结构，这里从map算子往下游通信开始封装）

				timeNowStamp, _ := strconv.ParseInt(msg.TimeStamp, 10, 64)
				timeNow := time.Unix(timeNowStamp, 0)

			LOOP:
				if preTime == nil {
					preTime = &timeNow
				} else {
					timeGap, _ := time.ParseDuration("+" + windowAssign)
					nextTime := preTime.Add(timeGap)
					if timeNow.After(nextTime) {
						preTime = &timeNow
						klog.Info(fmt.Sprintf("每隔%v的时间段统计一次数据，表%v的结果为%v，表%v的结果为%v", windowAssign, tables[0], tableMap1, tables[1], tableMap2))

						// 消费数据，调用keyby算子的client方法，传递DataStream结构
						content := []map[string]int64{tableMap1, tableMap2}
						err := rpc.CreateSink(context.Background(), &sinkdemo.CreateSinkRequest{
							content,
							msg.TimeStamp,
						})
						if err != nil {
							klog.Fatal("call keyby service failed")
						}

						tableMap1 = make(map[string]int64)
						tableMap2 = make(map[string]int64)
						goto LOOP
					}
				}

				for i := range msg.Content {
					var temp interface{} = msg.Content[i]
					tumple := temp.(*reducedemo.Tuple)
					if tumple.Table == tables[0] {
						tableMap1[tumple.Key] += tumple.Value
					} else {
						tableMap2[tumple.Key] += tumple.Value
					}
				}

				klog.Info(fmt.Sprintf("当前处理的是reduce算子%d号，消息内容为%s", reduceIdx, msg))

				// TODO 消费数据，调用sink算子的client方法，传递DataStream结构
				klog.Info(fmt.Sprintf("获得统计数据表 %v 的内容为 %v，数据表 %v 的内容为 %v ", tables[0], tableMap1, tables[1], tableMap2))
			}

		}(ch)
	}
}

func HandleKeybyMsg(req *reducedemo.CreateReduceRequest) {
	//klog.Info(req.Content)
	// 负载均衡规则
	switch lb {
	case "round":
		reduceIdx++
		reduceIdx %= reduceNum
		reduceChanList[reduceIdx] <- req
	case "random":
		reduceIdx = rand.Intn(reduceNum)
		reduceChanList[reduceIdx] <- req
	}
}

// CreateReduce implements the ReduceServiceImpl interface.
func (s *ReduceServiceImpl) CreateReduce(ctx context.Context, req *reducedemo.CreateReduceRequest) (resp *reducedemo.CreateReduceResponse, err error) {

	HandleKeybyMsg(req)

	resp = new(reducedemo.CreateReduceResponse)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
