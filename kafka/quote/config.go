package quote

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

var config = &kafka.ConfigMap{
	"bootstrap.servers": "172.16.0.99:9092,172.16.0.99:9093",
	"group.id":          "maxhu2022" + uuid.NewString(),
	"auto.offset.reset": "latest", // 默认从最新消费
}
