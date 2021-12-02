package Services

import (
	"github.com/Shopify/sarama"
	"time"
)

type KafkaContext struct {
	Producer sarama.SyncProducer
	Consumer sarama.ConsumerGroup
}

func NewKafka(c *Config) *KafkaContext {
	return &KafkaContext{
		Producer: NewKafkaProducer(c),
		Consumer: NewKafkaConsumer(c),
	}
}

//Sync Hash
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

//需要保持标记逻辑在插入mysql代码之后即可确保不会出现丢消息
func NewKafkaConsumer(c *Config) sarama.ConsumerGroup {
	kConfig := sarama.NewConfig()
	kConfig.Consumer.Offsets.AutoCommit.Enable = true              // 自动提交
	kConfig.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 间隔1s
	kConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	kConfig.Consumer.Offsets.Retry.Max = 3
	group, err := sarama.NewConsumerGroup(c.Kafka.Host, c.Kafka.OrderGroup, kConfig)
	if err != nil {
		panic("Failed to start consumer: " + err.Error())
	}
	return group
}
