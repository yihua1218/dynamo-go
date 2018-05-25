package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// ListTables list DynamoDB tables
func (db DB) ListTables() {
	input := &dynamodb.ListTablesInput{}
	req := db.Client.ListTablesRequest(input)
	res, err := req.Send()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// DescribeTable get information about the table
func (db DB) DescribeTable(table string) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(table),
	}
	req := db.Client.DescribeTableRequest(input)
	res, err := req.Send()

	if err != nil {
		fmt.Println(err)
	} else {
		db.save(res, fmt.Sprint("cache/", table, ".json"))
	}
}
