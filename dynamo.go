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

// DB is the configuration of this DynamoDB access instance
type DB struct {
	Tables []TableDef
	Region string
	Config aws.Config
	Client *dynamodb.DynamoDB
	Error  error
}

var (
	db = DB{
		Tables: []TableDef{},
		Region: "us-west-2",
	}
)

func init() {
	fmt.Println(db)

	db.Config, db.Error = external.LoadDefaultAWSConfig()

	if db.Error != nil {
		panic("unable to load SDK config, " + db.Error.Error())
	}
}

// New carete new DynamoDB access instance
func New(region string) DB {
	db.Region = region

	db.Client = dynamodb.New(db.Config)

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":id": {
				S: aws.String("03AFB154"),
			},
		},
		KeyConditionExpression: aws.String("id = :id"),
		TableName:              aws.String("g3_devices"),
	}

	req := db.Client.QueryRequest(input)

	res, err := req.Send()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	return db
}
