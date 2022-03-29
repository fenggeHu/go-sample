package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

func consumer_start() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "172.16.0.99:9092,172.16.0.99:9093",
		"group.id":          "maxhu2022" + uuid.NewString(),
		"auto.offset.reset": "latest", // 默认从最新消费
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topic_hk_quote"}, nil)
	defer c.Close()

	var count = 0
	for {
		//
		m, err := c.ReadMessage(60 * time.Second)
		count++
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, m)
		} else {
			//ms := time.Now().UnixMilli() - m.Value
			//fmt.Printf("Offset[%s] - Partition[%d] - Value: %s\n", m.TopicPartition.Offset, m.TopicPartition.Partition, m.Value)
			strArray := strings.Split(string(m.Value), ",")
			t, _ := strconv.ParseInt(strArray[3], 10, 64)
			if count%50000 == 1 {
				log.Printf("Kafka数据分区[%d] 生产时间: %d 消费时间差: %d", m.TopicPartition.Partition,
					t, (time.Now().UnixMilli()-t)/1000)
			}
			o, _ := strconv.ParseFloat(strArray[4], 64)
			h, _ := strconv.ParseFloat(strArray[5], 64)
			l, _ := strconv.ParseFloat(strArray[6], 64)
			c, _ := strconv.ParseFloat(strArray[7], 64)
			response := &ResponseMessage{
				Time:   int64(t),
				Symbol: strArray[2],
				Data:   []float64{o, h, l, c},
			}
			WsManager.send(response)
		}
		//time.Sleep(5)
	}

}
