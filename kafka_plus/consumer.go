package kafka_plus

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hewenyu/toolspackage/global_variable"
)

func StartKafka(k Consumer, host string) {
	
	var run = true
	
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration
		// property.
		"broker.address.family": "v4",
		// Consumer group ID
		"group.id": k.GroupID,
		// Use the cooperative incremental rebalance protocol.
		"partition.assignment.strategy": "cooperative-sticky",
		// Start reading from the first message of each assigned
		// partition if there are no previously committed offsets
		// for this group.
		"auto.offset.reset": "earliest",
		
		"max.poll.interval.ms": "6000",
		"session.timeout.ms":   "6000",
	})
	
	if err != nil {
		global_variable.CLOG.Fatal(fmt.Sprintf("Failed to create consumer: %s", err))
	}
	
	err = consumer.SubscribeTopics(k.Topic, rebalancedCallback)
	
	if err != nil {
		global_variable.CLOG.Fatal(fmt.Sprintf("Failed to create consumer: %s", err))
	}
	
	for run {
		ev := consumer.Poll(100)
		if ev == nil {
			continue
		}
		
		switch e := ev.(type) {
		case *kafka.Message:
			result := string(e.Value)
			global_variable.CLOG.Info(result)
		
		case kafka.Error:
			
			global_variable.CLOG.Error(fmt.Sprintf("%% Error: %v: %v", e.Code(), e))
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		default:
			global_variable.CLOG.Info(fmt.Sprintf("Ignored %v", e))
		}
		
	}
	
	global_variable.CLOG.Info("Closing consumer")
	
	err = consumer.Close()
	if err != nil {
		return
	}
	
}
