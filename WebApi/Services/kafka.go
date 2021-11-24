package Services

import (
	"github.com/Shopify/sarama"
)

type KafkaHandler interface {
	Push(c *Config, mes string) (partition int32, offset int64, err error)
}

type KafkaContext struct {
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
}

func NewKafka(c *Config) *KafkaContext {
	return &KafkaContext{
		Producer: NewKafkaProducer(c),
		Consumer: NewKafkaConsumer(c),
	}
}

func (m *KafkaContext) Push(topic, key, mes string) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(mes),
	}
	return m.Producer.SendMessage(msg)
}

//Sync
func NewKafkaProducer(c *Config) sarama.SyncProducer {
	kConfig := sarama.NewConfig()
	kConfig.Producer.RequiredAcks = sarama.WaitForAll //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	kConfig.Producer.Partitioner = sarama.NewHashPartitioner
	kConfig.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer(c.Kafka.Host, kConfig)
	if err != nil {
		panic(err)
	}
	return client
}

func NewKafkaConsumer(c *Config) sarama.Consumer {
	consumer, err := sarama.NewConsumer(c.Kafka.Host, nil)
	if err != nil {
		panic("Failed to start consumer: " + err.Error())
	}
	return consumer
}