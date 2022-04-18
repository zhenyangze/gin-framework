package helpers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type KafkaMessage struct {
	*sarama.ConsumerMessage
}

func KafkaConsumer(brokerList []string, topic string, f func(msg *KafkaMessage)) {
	var wg sync.WaitGroup
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

type KafkaGroupMessage struct {
	name     string
	callback func(msg *KafkaMessage)
}

func (KafkaGroupMessage) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (KafkaGroupMessage) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (msg KafkaGroupMessage) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message topic:%q partition:%d offset:%d", message.Topic, message.Partition, message.Offset)
		// 调用自自义方法
		msg.callback(&KafkaMessage{
			message,
		})
		// 标记完成
		sess.MarkMessage(message, "")
	}
	return nil
}

func consumerByGroup(wg *sync.WaitGroup, group *sarama.ConsumerGroup, groupId string, topics []string, f func(msg *KafkaMessage)) {
	defer wg.Done()
	ctx := context.Background()
	for {
		handler := KafkaGroupMessage{callback: f, name: groupId}
		err := (*group).Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

func KafkaGroupConsumer(brokerList []string, groupId string, topic string, f func(msg *KafkaMessage)) {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	//config.Version = kfversion
	config.Version = sarama.V0_10_2_0
	config.Consumer.Return.Errors = false

	// Start with a client
	client, err := sarama.NewClient(brokerList, config)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	group, err := sarama.NewConsumerGroupFromClient(groupId, client)
	if err != nil {
		log.Println(err)
	}
	defer group.Close()

	topics := []string{topic}
	wg.Add(1)
	go consumerByGroup(&wg, &group, groupId, topics, f)
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
