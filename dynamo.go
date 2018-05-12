package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// TableDef is the definition of a DynamoDB table
type TableDef struct {
	Name      string
	TableName string
	Key       string
	Indexes   []string
}

// Conf is the configuration of this DynamoDB access instance
type Conf struct {
	Tables []TableDef
	Region string
}

var (
	conf = Conf{
		Tables: []TableDef{},
		Region: "us-west-2",
	}
)

func init() {
	fmt.Println(conf)

	cfg, err := external.LoadDefaultAWSConfig()

	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	cfg.Region = conf.Region

	svc := dynamodb.New(cfg)

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":id": {
				S: aws.String("03AFB154"),
			},
		},
		KeyConditionExpression: aws.String("id = :id"),
		TableName:              aws.String("g3_devices"),
	}

	req := svc.QueryRequest(input)

	res, err := req.Send()
	fmt.Println(res)
}
