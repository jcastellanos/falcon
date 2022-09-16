package adapters

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/jcastellanos/falcon/core/constants"
	"github.com/jcastellanos/falcon/core/models"
	"github.com/jcastellanos/falcon/core/utils"
	"log"
)

type AmazonConnectNotifierAdapter struct {}

func NewAmazonConnectNotifierAdapter() AmazonConnectNotifierAdapter {
	return AmazonConnectNotifierAdapter {}
}

func (a AmazonConnectNotifierAdapter) Notify(alert models.Alert, guard models.Guard) (bool, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(utils.GetConfig(constants.AWS_ZONE_KEY))},
	)
	if err != nil {
		log.Fatal(err)
	}
	conn := connect.New(sess)
	contactInput := connect.StartOutboundVoiceContactInput {
		ContactFlowId:                aws.String(utils.GetConfig(constants.AWS_CONNECT_CONTACT_FLOW_ID_KEY)),
		DestinationPhoneNumber:       aws.String(guard.Primary.Phone),
		InstanceId:                   aws.String(utils.GetConfig(constants.AWS_CONNECT_INSTANCE_ID_KEY)),
		SourcePhoneNumber:            aws.String(utils.GetConfig(constants.AWS_CONNECT_SOURCE_PHONE_NUMBER_KEY)),
	}
	log.Printf("Calling %s", guard.Primary.Phone)
	_, err = conn.StartOutboundVoiceContact(&contactInput)
	if err != nil {
		log.Fatal(err)
	}
	return true, nil
}