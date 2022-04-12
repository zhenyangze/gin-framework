package helpers

import (
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

type KafkaMessage struct {
	*sarama.ConsumerMessage
}

func KafkaConsumer(brokerList []string, topic string, f func(msg *KafkaMessage)) {
	//创建消费者
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	//获取主题分区
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	//遍历分区
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		wg.Add(1) //+1
		go func(sarama.PartitionConsumer) {
			defer wg.Done() //-1
			for msg := range pc.Messages() {
				log.Printf("[kafka]Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				f(&KafkaMessage{msg})
			}
		}(pc)
	}
	wg.Wait()
}

func KafkaProduct(brokerList []string, topic string, message string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)
	client, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return err
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}
