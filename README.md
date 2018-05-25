# DynamoDB for Go

This is a package for easy access dynamodb.

## Create DynamoDB Instance

``` go
db := dynamo.New("us-east-2")
```

## Table Operation

1. ListTables()
2. DescribeTable(tableName)

## Record Operation

1. Scan()
2. Query(tableName, primaryKey)
3. Put()
4. Update()

## Reference

1. [AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/api/)
2. [AWS SDK for the Go programming language.](https://github.com/aws/aws-sdk-go-v2)
3. [Expressive DynamoDB library for Go](https://github.com/guregu/dynamo)
