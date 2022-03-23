package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	topic := "topic_hk_quote"
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"172.16.0.99:9092", "172.16.0.99:9093"},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(kafka.LastOffset)

	// 测试耗时
	// 1W - 258.488914ms
	// 10W -  1.493010664s
	// 100W - ?
	//i := 0
	//start := time.Now()
	//for i < 1000000 {
	//	_, _ = r.ReadMessage(context.Background())
	//	i++
	//}
	//fmt.Println("总耗时：", time.Since(start))

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Errorf("%s", err)
			continue
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
