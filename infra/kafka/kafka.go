package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func NewProducer(vip *viper.Viper) (*kafka.Writer, error) {
	addr := fmt.Sprintf(
		"%s:%d",
		vip.GetString("kafka.host"),
		vip.GetInt("kafka.port"),
	)

	conn, err := kafka.DialContext(context.Background(), "tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	writer := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Topic:        vip.GetString("kafka.topic"),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return writer, nil
}

func NewConsumer(vip *viper.Viper) (*kafka.Reader, error) {
	addr := fmt.Sprintf(
		"%s:%d",
		vip.GetString("kafka.host"),
		vip.GetInt("kafka.port"),
	)

	conn, err := kafka.DialContext(context.Background(), "tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{addr},
		Topic:   vip.GetString("kafka.topic"),
		GroupID: vip.GetString("kafka.group_id"),

		MinBytes: 10e3, // 10 KB
		MaxBytes: 10e6, // 10 MB

		StartOffset: kafka.FirstOffset,
	})

	return reader, nil
}
