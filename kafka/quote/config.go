package quote

var kafkaConfig = map[string]string{
	"bootstrap.servers": "172.16.0.99:9092,172.16.0.99:9093",
	"group.id":          "maxhu2022",
	"auto.offset.reset": "latest",
}
