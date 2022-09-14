package adapters

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/jcastellanos/falcon/core/models"
	"log"
)

type AmazonConnectNotifierAdapter struct {}

func NewAmazonConnectNotifierAdapter() AmazonConnectNotifierAdapter {
	return AmazonConnectNotifierAdapter {}
}

func (a AmazonConnectNotifierAdapter) Notify(alert models.Alert, guard models.Guard) (bool, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Fatal(err)
	}
	conn := connect.New(sess)
	contactInput := connect.StartOutboundVoiceContactInput{
		ContactFlowId:                aws.String(""),
		DestinationPhoneNumber:       aws.String(""),
		InstanceId:                   aws.String(""),
		SourcePhoneNumber:            aws.String(""),
	}
	contactOuput, err := conn.StartOutboundVoiceContact(&contactInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contactOuput.GoString())
	return true, nil
}
