package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"82.156.171.8:10004"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("test") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("list = ", partitionList, len(partitionList))

	pc, err := consumer.ConsumePartition("test", int32(partitionList[0]), sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("failed to start consumer for partition %d,err:%v\n", partitionList[0], err)
		return
	}
	for msg := range pc.Messages() {
		//fmt.Println("打印信息")
		fmt.Println("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
	}
}
