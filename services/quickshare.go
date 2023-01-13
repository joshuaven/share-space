package services

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/joshuaven/share-space/models"
)

type QuickShareService struct {
	DynamoDb *dynamodb.DynamoDB
}

type IQuickShareService interface {
	AddItem(models.QSItem) error
	GetItem(string, *models.QSItem) error
}

func (s *QuickShareService) AddItem(newItem models.QSItem) error {
	av, err := dynamodbattribute.MarshalMap(newItem)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(os.Getenv("FILE_DB")),
	}

	_, err = s.DynamoDb.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (s *QuickShareService) GetItem(fileId string, item *models.QSItem) (error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("FILE_DB")),
		Key: map[string]*dynamodb.AttributeValue{
			"FileId": {
				S: aws.String(fileId),	
			},
		},
	}

	res, err := s.DynamoDb.GetItem(query)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalMap(res.Item, item)
	if err != nil {
		return err
	}

	return nil
}

func CreateQuickShareService() IQuickShareService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	}))

	service := &QuickShareService{
		DynamoDb: dynamodb.New(sess),
	}

	return service
}