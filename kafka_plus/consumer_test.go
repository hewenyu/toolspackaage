package kafka_plus

import "testing"

func TestConsumer(t *testing.T) {
	
	var (
		host    = []string{"kafka1:9092"}
		topic   = []string{"test"}
		groupId = "test"
	)
	config := NewKafka(host, topic, groupId, 3)
	
	for _, v := range config.Host {
		go StartKafka(config.Consumer, v)
	}
}
