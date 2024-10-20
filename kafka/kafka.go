// pkg/messaging/kafka/producer.go
package kafka

import (
	"encoding/json"
	"helloWorld/models"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Producer) SendOrderBookUpdate(update *models.OrderBookUpdate) error {
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
		Key:   sarama.StringEncoder(update.Symbol),
	}

	_, _, err = p.producer.SendMessage(msg)
	return err
}

// pkg/messaging/kafka/consumer.go
type Consumer struct {
	consumer sarama.Consumer
	topic    string
}

func NewConsumer(brokers []string, topic string) (*Consumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: consumer,
		topic:    topic,
	}, nil
}

func (c *Consumer) Consume(handler func(*models.OrderBookUpdate)) error {
	partitionConsumer, err := c.consumer.ConsumePartition(
		c.topic,
		0,
		sarama.OffsetNewest,
	)
	if err != nil {
		return err
	}

	for msg := range partitionConsumer.Messages() {
		var update models.OrderBookUpdate
		if err := json.Unmarshal(msg.Value, &update); err != nil {
			continue
		}
		handler(&update)
	}

	return nil
}
