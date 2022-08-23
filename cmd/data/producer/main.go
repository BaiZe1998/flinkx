package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"82.156.171.8:10004"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息

	wordList := []string{"apple", "banana", "peach", "watermelon", "orange", "tomato", "potato"}
	rand.Seed(time.Now().UnixNano())

	for {
		idx := rand.Intn(len(wordList))
		//time.Sleep(time.Second * 3)
		msg.Value = sarama.StringEncoder(wordList[idx] + "," + strconv.FormatInt(time.Now().Unix(), 10))
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}

}
