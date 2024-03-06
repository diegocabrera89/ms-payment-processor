package repository

import (
	"context"
	"fmt"
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

// GetProcessorByMerchantIDRepository put item in DynamoDB.
func (r *ProcessorRepositoryImpl) GetProcessorByMerchantIDRepository(ctx context.Context, request events.APIGatewayProxyRequest, fieldValueFilterByID string) (models.Processor, error) {
	logs.LogTrackingInfo("GetProcessorByMerchantIDRepository", ctx, request)
	getProcessorByMerchantID, errorGetPetById := r.CoreRepository.GetItemByFieldCore(ctx, request, constantsmicro.MerchantID, fieldValueFilterByID, constantsmicro.MerchantIDIndex, "", "")
	if errorGetPetById != nil {
		logs.LogTrackingError("GetProcessorByMerchantIDRepository", "GetItemByFieldCore", ctx, request, errorGetPetById)
		return models.Processor{}, errorGetPetById
	}
	logs.LogTrackingInfoData("GetProcessorByMerchantIDRepository", getProcessorByMerchantID, ctx, request)

	// Verificar si hay al menos un elemento en la respuesta
	if getProcessorByMerchantID.Count == 0 {
		return models.Processor{}, fmt.Errorf("No items found")
	}

	// Crear un slice de instancias de models.Merchant
	merchants := make([]models.Processor, len(getProcessorByMerchantID.Items))
	logs.LogTrackingInfoData("GetProcessorByMerchantIDRepository merchants", merchants, ctx, request)

	// Deserializar los maps en las instancias de models.Merchant
	for i, item := range getProcessorByMerchantID.Items {
		var m models.Processor
		err := helpers.UnmarshalMapToType(item, &m)
		if err != nil {
			logs.LogTrackingError("GetProcessorByMerchantIDRepository", "UnmarshalMapToType", ctx, request, err)
			return models.Processor{}, err
		}
		merchants[i] = m
		logs.LogTrackingInfoData("GetProcessorByMerchantIDRepository processors[i]", merchants[i], ctx, request)
	}

	logs.LogTrackingInfoData("GetProcessorByMerchantIDRepository processors", merchants, ctx, request)
	// Si solo esperas un elemento, devuélvelo
	if len(merchants) == 1 {
		return merchants[0], nil
	}

	return models.Processor{}, fmt.Errorf("More than one item found")
}
