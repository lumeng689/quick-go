package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
	"time"

	cluster "github.com/bsm/sarama-cluster"
)

var (
	kafkaBrokers = []string{"127.0.0.1:9092"}
	kafkaTopics  = []string{"test_topic_1"}
	groupId      = "test_group_1"
)

func getKafkaConsumer() *cluster.Consumer {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	config.Group.Return.Notifications = true

	kafkaConsumer, err := cluster.NewConsumer(kafkaBrokers, groupId, kafkaTopics, config)

	if err != nil {
		fmt.Printf("error: %v \n", err)
		panic(fmt.Sprintf("create consumer error. kafka info -> {brokers:%v, group: %v}", kafkaBrokers, groupId))
	}

	return kafkaConsumer
}

func getKafkaProducer() sarama.SyncProducer {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 仅需要kafka的leader确认就算成功
	config.Producer.Retry.Max = 5
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Microsecond

	producer, err := sarama.NewSyncProducer(kafkaBrokers, config)

	if err != nil {
		fmt.Printf("error: %v \n", err)
		panic(fmt.Sprintf("create producer error. kafka info -> {brokers:%v, group: %v}", kafkaBrokers, groupId))
	}

	return producer
}

func TestKafkaProducer001(t *testing.T) {
	producer := getKafkaProducer()

	msg := &sarama.ProducerMessage{
		Topic: "test_topic_1",
		Value: sarama.StringEncoder("hello1"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg error:%s \n", err)
	}

	fmt.Printf("partition:%v, offset: %v \n", partition, offset)
}

func TestKafkaConsumer001(t *testing.T) {
	consumer := getKafkaConsumer()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Printf("kafka msg: %s \n", msg.Value)
				consumer.MarkOffset(msg, "")
			} else {
				fmt.Printf("kafka get message failed!")
			}
		case err, ok := <-consumer.Errors():
			if ok {
				fmt.Printf("consumer error: %v", err)
			}
		case ntf, ok := <-consumer.Notifications():
			if ok {
				fmt.Printf("consumer notification: %v", ntf)
			}
		}
	}
}
