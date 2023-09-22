package aws

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
)

const (
	STANDARD = iota
	FIFO
)

type AwsContext struct {
	config aws.Config
}

func AwsConfig(profileName string) *AwsContext {
	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profileName),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &AwsContext{config: config}
}

func (ac *AwsContext) SendMessage(fc, qu string) {
	sc := sqs.NewFromConfig(ac.config)
	qt := getQueueType(qu)

	msg := &sqs.SendMessageInput{
		QueueUrl:    &qu,
		MessageBody: &fc,
	}

	if qt == FIFO {
		guid := uuid.New().String()
		msg.MessageDeduplicationId = &guid
        msg.MessageGroupId = &guid
	}

	_, err := sc.SendMessage(context.TODO(), msg)
	if err != nil {
		log.Fatal(err)
	}
}

func getQueueType(qu string) int {
	if strings.HasSuffix(qu, ".fifo") {
		return FIFO
	}

	return STANDARD
}
