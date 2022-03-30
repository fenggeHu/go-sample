package quote

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

var TOPIC = []string{"topic_hk_trade"}

// TradeConsumer tick消费时间测试
func TradeConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConfig["bootstrap.servers"],
		"group.id":          kafkaConfig["group.id"] + uuid.NewString(),
		"auto.offset.reset": kafkaConfig["auto.offset.reset"], // 默认从最新消费
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(TOPIC, nil)
	//c.Pause()
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
			t, _ := strconv.ParseInt(strArray[4], 10, 64)
			if count%50000 == 1 {
				log.Printf("tradeConsumer->第%d条Kafka数据分区[%d] 生产时间: %d 消费时间差: %d", count, m.TopicPartition.Partition,
					t, time.Now().UnixMilli()-t)
			}
		}
		//time.Sleep(5)
	}

}
