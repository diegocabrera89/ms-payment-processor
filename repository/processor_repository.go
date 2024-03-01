package repository

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/diegocabrera89/ms-payment-core/dynamodbcore"
	"github.com/diegocabrera89/ms-payment-core/helpers"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-processor/constantsmicro"
	"github.com/diegocabrera89/ms-payment-processor/models"
	"os"
)

// ProcessorRepositoryImpl implements the ProcessorRepository interface of the ms-payment-core package.
type ProcessorRepositoryImpl struct {
	CoreRepository dynamodbcore.CoreRepository
}

// NewProcessorRepository create a new ProcessorRepository instance.
func NewProcessorRepository() *ProcessorRepositoryImpl {
	processorTable := os.Getenv(constantsmicro.ProcessorTable)
	region := os.Getenv(constantsmicro.Region)

	coreRepository, _ := dynamodbcore.NewDynamoDBRepository(processorTable, region)

	return &ProcessorRepositoryImpl{
		CoreRepository: coreRepository,
	}
}

// CreateProcessorRepository put item in DynamoDB.
func (r *ProcessorRepositoryImpl) CreateProcessorRepository(ctx context.Context, request events.APIGatewayProxyRequest, processor models.Processor) (models.Processor, error) {
	logs.LogTrackingInfo("CreateProcessorRepository", ctx, request)
	item, errorMarshallItem := helpers.MarshallItem(processor)
	if errorMarshallItem != nil {
		logs.LogTrackingError("CreateProcessorRepository", "MarshallItem", ctx, request, errorMarshallItem)
		return models.Processor{}, errorMarshallItem
	}

	errorPutItemCore := r.CoreRepository.PutItemCore(ctx, request, item)
	if errorPutItemCore != nil {
		return models.Processor{}, errorPutItemCore
	}
	return processor, nil
}
