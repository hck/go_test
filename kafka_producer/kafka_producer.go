package kafka_producer

import (
	"github.com/Shopify/sarama"
	"log"
)

type Producer struct {
  producer sarama.SyncProducer
}

func (p *Producer) Init(addr []string) {
  producer, err := sarama.NewSyncProducer(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
  p.producer = producer
}

func (p *Producer) SendMessage(topic string, message string) bool {
  msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}

  _, _, err := p.producer.SendMessage(msg)

  if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
    return false
	} else {
    return true
	}
}

func (p *Producer) Close() {
  err := p.producer.Close()
  if err != nil {
    log.Fatalln(err)
  }
}
