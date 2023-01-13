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

type ShortUrlService struct {
	DynamoDb *dynamodb.DynamoDB
}

type IShortUrlService interface {
	AddItem(models.ShortUrlItem) error
	GetItem(string, *models.ShortUrlItem) (error)
}

func (s *ShortUrlService) AddItem(newItem models.ShortUrlItem) error {
	av, err := dynamodbattribute.MarshalMap(newItem)
	if err != nil {
			return err
	}

	input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(os.Getenv("SHORTURL_DB")),
	}

	_, err = s.DynamoDb.PutItem(input)
	if err != nil {
			return err
	}

	return nil;
}

func (s *ShortUrlService) GetItem(urlId string, item *models.ShortUrlItem) (error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("SHORTURL_DB")),
		Key: map[string]*dynamodb.AttributeValue{
			"UrlId": {
				S: aws.String(urlId),
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

func CreateShortUrlService() IShortUrlService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	}))

	service := &ShortUrlService{
		DynamoDb: dynamodb.New(sess),
	}

	return service
}