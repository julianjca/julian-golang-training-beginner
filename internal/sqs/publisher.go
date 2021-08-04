package sqs

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Publisher struct {
	QueueUrl *string
	SQS      *sqs.SQS
}

func NewPublisher(s *session.Session, queue string) (*Publisher, error) {
	var p Publisher
	sq := sqs.New(s)

	res, err := sq.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queue,
	})
	if err != nil {
		return &p, err
	}

	p.SQS = sq
	p.QueueUrl = res.QueueUrl

	return &p, nil
}

func (p Publisher) Publish(msg interface{}) error {
	messageBody, _ := json.Marshal(msg)

	_, err := p.SQS.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    p.QueueUrl,
	})

	return err
}
