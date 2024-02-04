package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type userModelAuth struct {
	Username        string `dynamodbav:"username"`
	AuthorizerToken string `dynamodbav:"authorizerToken"`
	Role            string `dynamodbav:"role"`
}

var dynamoDbObj *dynamodb.Client
var TableName string

func isAuthorized(u string, token string) (userModelAuth, bool) {
	var user userModelAuth
	resp, err := dynamoDbObj.GetItem(context.TODO(), &dynamodb.GetItemInput{Key: GetDynamoKeys(u), TableName: aws.String(TableName)})
	if err != nil {
		fmt.Printf("Cannot Get Item from Table %s: %v\n", TableName, err)
		return userModelAuth{}, false

	} else {
		err = attributevalue.UnmarshalMap(resp.Item, &user)
		if err != nil {
			fmt.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}

	}
	if user.AuthorizerToken == token {
		return user, true
	}
	return userModelAuth{}, false
}
func GetDynamoKeys(username string) map[string]types.AttributeValue {
	encoded, err := attributevalue.Marshal(username)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{"username": encoded}
}
func isLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}
func handleRequest(ctx context.Context, apiGatewayRequest events.APIGatewayProxyRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	if user, isAuth := isAuthorized(apiGatewayRequest.Headers["username"], apiGatewayRequest.Headers["token"]); isAuth {
		m := make(map[string]interface{})
		m["role"] = user.Role
		m["complete"] = user
		fmt.Println(user)
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: true, Context: m}, nil
	}

	return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, nil

}
func main() {
	TableName = "simpleAuthorizer"
	cfg, _ := config.LoadDefaultConfig(context.TODO())

	dynamoDbObj = dynamodb.NewFromConfig(cfg)

	if isLambda() {
		lambda.Start(handleRequest)
	}
}
