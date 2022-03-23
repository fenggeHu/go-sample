module go-sample

go 1.17

require (
	github.com/confluentinc/confluent-kafka-go v1.8.2
	github.com/segmentio/kafka-go v0.4.30	// 额外引入了外部依赖
)

require (
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
)
