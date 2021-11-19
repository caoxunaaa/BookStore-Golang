package kafka

import (
	"Order/internal/config"
	"fmt"
	"github.com/Shopify/sarama"
)

func NewKafkaProducer(c config.KafkaConf) (sarama.SyncProducer, error) {
	kconfig := sarama.NewConfig()
	kconfig.Producer.RequiredAcks = sarama.WaitForAll //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	kconfig.Producer.Partitioner = sarama.NewHashPartitioner
	kconfig.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = c.Topic
	//目前是Hash分布，当设置了Key值时，消息会被分配到同一个分区(看做有序的队列)
	if c.Key != "" {
		msg.Key = sarama.StringEncoder(c.Key)
	}
	client, err := sarama.NewSyncProducer(c.Host, kconfig)
	if err != nil {
		fmt.Println("producer close err, ", err)
		err = client.Close()
		return nil, err
	}
	return client, nil
}

func NewKafkaConsumer(c config.KafkaConf) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer(c.Host, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err)
		return nil, err
	}
	return consumer, nil
}
