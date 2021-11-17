package kafka_plus

import (
	"strings"
)

type Consumer struct {
	Topic     []string `json:"topic"`     // 使用哪个 topic
	GroupID   string   `json:"group_id"`  // 消费groupId
	Goroutine int      `json:"goroutine"` // 启动几个goroutine
}

type KafKa struct {
	Host     []string `json:"host"` // host 端口 类似  localhost:9092
	Consumer Consumer `json:"consumer"`
}

// NewKafka 初始化kafka
func NewKafka(host []string, topics []string, groupId string, limit int) *KafKa {
	
	return &KafKa{
		Host: host,
		Consumer: Consumer{
			Topic:     topics,
			GroupID:   groupId,
			Goroutine: limit,
		},
	}
}

// FormatHost 将 host 输出
func (k *KafKa) FormatHost() string {
	
	return strings.Join(k.Host, ",")
	
}
