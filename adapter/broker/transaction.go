package broker

import (
	"context"

	"github.com/MihasBel/test-transactions-rest/model"
	bModel "github.com/MihasBel/test-transactions/broker/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ogen-go/ogen/json"
	"github.com/rs/zerolog/log"
)

// TopicTransaction name of transaction topic
var TopicTransaction = "transaction"

// PlaceTransaction create new transaction nad post to broker
func (b *Broker) PlaceTransaction(_ context.Context, transaction model.Transaction) error {
	t := bModel.Transaction{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		Amount:      transaction.Amount,
		Status:      transaction.Status,
		CreatedAt:   transaction.CreatedAt,
		Description: transaction.Description,
	}
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	delChan := make(chan kafka.Event, 10000)
	err = b.b.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &TopicTransaction, Partition: kafka.PartitionAny},
		Value:          data},
		delChan,
	)
	if err != nil {
		return err
	}
	channelOut := <-delChan
	messageReport := channelOut.(*kafka.Message)

	if messageReport.TopicPartition.Error != nil {
		return messageReport.TopicPartition.Error
	}
	log.Info().Msg("Message delivered")
	close(delChan)
	return nil
}
