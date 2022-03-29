package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "172.16.0.99:9092,172.16.0.99:9093",
		"group.id":          "maxhu2022",
		"auto.offset.reset": "latest", // 默认从最新消费
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topic_hk_quote"}, nil)

	// 测试耗时
	// 1W - 544.412399ms
	// 10W -  1.684926912s
	// 100W - 15.0426431s
	//i := 0
	//start := time.Now()
	//for i < 100000 {
	//	_, _ = c.ReadMessage(1000 * time.Second)
	//	i++
	//}
	//fmt.Println("总耗时：", time.Since(start))

	//
	for {
		//
		m, err := c.ReadMessage(1000 * time.Second)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, m)
		} else {
			//ms := time.Now().UnixMilli() - m.Value
			fmt.Printf("Offset[%s] - Partition[%d] - Value: %s\n", m.TopicPartition.Offset, m.TopicPartition.Partition, m.Value)
		}
		//i++
	}

	defer c.Close()
}
