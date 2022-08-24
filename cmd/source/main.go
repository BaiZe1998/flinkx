package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
	"word-count/cmd/api/rpc"
	"word-count/config"
	"word-count/kitex_gen/mapdemo"
)

func main() {
	enginConfig := config.GetConfig()

	maxNum := enginConfig.Source.Num

LOOP:
	sourcelist := config.GlobalDAGConfig.GetStringSlice("source.data-from")

	sourceNum := config.GlobalDAGConfig.GetInt("source.num")

	// DAG中source算子数量非法状态
	if sourceNum <= 0 || sourceNum > maxNum {
		klog.Fatal("DAG中指定的source数量过小/过大")
	}
	// 阻塞
	for config.GlobalDAGConfig.GetString("mode") == "off" {
	}
	var wg sync.WaitGroup
	// 启动source算子从kafka数据源列表获取数据发往map算子程序
	for _, dataSource := range sourcelist {
		wg.Add(1)
		// 这里开启多协程去往下游map算子发送msg（这里相同类型的算子是单进程多协程模型的，在算子内部开启多协程处理可能多个上游发送过来的数据）
		go func(string, *sync.WaitGroup) {

			consumer, err := sarama.NewConsumer([]string{dataSource}, nil)
			if err != nil {
				fmt.Printf("fail to start consumer, err:%v\n", err)
				return
			}
			partitionList, err := consumer.Partitions("test") // 根据topic取到所有的分区
			if err != nil {
				fmt.Printf("fail to get list of partition:err%v\n", err)
				return
			}
			//fmt.Println("list = ", partitionList, len(partitionList))

			pc, err := consumer.ConsumePartition("test", int32(partitionList[0]), sarama.OffsetNewest)
			if err != nil {
				fmt.Printf("failed to start consumer for partition %d,err:%v\n", partitionList[0], err)
				return
			}

			// 请求map服务客户端初始化
			rpc.InitMapRPC()

			for msg := range pc.Messages() {
				if config.GlobalDAGConfig.GetString("mode") == "off" {
					break
				}

				// 发送msg到map算子
				fmt.Println("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))

				// 如果下游没有在nacos的列表中则会报错
				err := rpc.CreateMap(context.Background(), &mapdemo.CreateMapRequest{
					string(msg.Value),
				})
				if err != nil {
					klog.Error("call map service failed, ", err)
					// TODO 上游服务熔断 & 下游拉起服务（下游没有可以使用的服务）
				}
			}
			wg.Done()
		}(dataSource, &wg)
	}
	wg.Wait()
	goto LOOP
}
