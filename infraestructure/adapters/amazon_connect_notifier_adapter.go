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
		ContactFlowId:                aws.String("22cacb61-b4c9-4f1c-9bfa-5e924a026bb6"),
		DestinationPhoneNumber:       aws.String("+576015357113"),
		InstanceId:                   aws.String("ebfc66bd-0344-4878-955e-ed8580637fcb"),
		SourcePhoneNumber:            aws.String("+576015143170"),
	}
	contactOuput, err := conn.StartOutboundVoiceContact(&contactInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contactOuput.GoString())
	return true, nil
}
