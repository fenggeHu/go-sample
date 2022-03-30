package quote

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go-sample/websocket/ws"
	"log"
	"strconv"
	"strings"
	"time"
)

func SnapshotConsumer() {
	c, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topic_hk_quote"}, nil)
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
			t, _ := strconv.ParseInt(strArray[3], 10, 64)
			if count%50000 == 1 {
				log.Printf("quoteConsumer->第%d条Kafka数据分区[%d] 生产时间: %d 消费时间差: %d", count, m.TopicPartition.Partition,
					t, time.Now().UnixMilli()-t)
			}
			o, _ := strconv.ParseFloat(strArray[4], 64)
			h, _ := strconv.ParseFloat(strArray[5], 64)
			l, _ := strconv.ParseFloat(strArray[6], 64)
			c, _ := strconv.ParseFloat(strArray[7], 64)
			response := &ws.ResponseMessage{
				Time:   int64(t),
				Symbol: strArray[2],
				Data:   []float64{o, h, l, c},
			}
			ws.Manager.Send(response)
		}
		//time.Sleep(5)
	}

}
