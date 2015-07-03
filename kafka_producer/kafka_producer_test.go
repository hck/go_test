package kafka_producer

import (
	"testing"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
)

func TestInit(t *testing.T) {
  p := Producer{}

  p.Init([]string{"localhost:9092"})
  v, ok := p.producer.(sarama.SyncProducer)

  if !ok {
     t.Errorf("expected InitProducer() to return SyncProducer, but returned %v", v)
  }
}

func TestClose(t *testing.T) {
  p := Producer{}

  sync_producer := mocks.NewSyncProducer(t, nil)
  p.producer = sync_producer

  p.Close()
  sync_producer.ExpectSendMessageAndSucceed()

  msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("message")}
  _, _, err := p.producer.SendMessage(msg)

  if err != nil {
    t.Errorf("extected SyncProducer to be closed, but no error after sending message")
  }
}

func TestSendMessage(t *testing.T) {
		var result bool
		error := errors.New("fail")

    sync_producer := mocks.NewSyncProducer(t, nil)
    p := Producer{producer: sync_producer}

		sync_producer.ExpectSendMessageAndSucceed()
		result = p.SendMessage("test", "test message")
		if !result {
			t.Errorf("expected message to be sent, SendMessage() returned %v", result)
		}

		sync_producer.ExpectSendMessageAndFail(error)
		result = p.SendMessage("test", "test message")
		if result {
			t.Errorf("expected message not to be sent, SendMessage() returned %v", result)
		}
}
