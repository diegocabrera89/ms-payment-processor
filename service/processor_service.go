package service

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/diegocabrera89/ms-payment-core/constantscore"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-core/response"
	"github.com/diegocabrera89/ms-payment-processor/models"
	"github.com/diegocabrera89/ms-payment-processor/repository"
	"github.com/diegocabrera89/ms-payment-processor/utils"
	"net/http"
)

// ProcessorService represents the service for the ProcessorService entity.
type ProcessorService struct {
	processorRepo *repository.ProcessorRepositoryImpl
}

// NewProcessorService create a new ProcessorService instance.
func NewProcessorService() *ProcessorService {
	return &ProcessorService{
		processorRepo: repository.NewProcessorRepository(),
	}
}

// CreateProcessorService handles the creation of a new processor.
func (r *ProcessorService) CreateProcessorService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("CreateProcessorService", ctx, request)
	var processor models.Processor
	err := json.Unmarshal([]byte(request.Body), &processor)
	if err != nil {
		logs.LogTrackingError("CreateProcessorService", "JSON Unmarshal", ctx, request, err)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.InvalidRequestBody)
	}

	utils.BuildCreateProcessor(&processor)

	createProcessor, errorProcessorRepository := r.processorRepo.CreateProcessorRepository(ctx, request, processor)
	if errorProcessorRepository != nil {
		logs.LogTrackingError("CreateProcessorService", "CreateProcessorRepository", ctx, request, errorProcessorRepository)
		return response.ErrorResponse(http.StatusInternalServerError, constantscore.ErrorCreatingItem)
	}

	responseBody, err := json.Marshal(createProcessor)
	if err != nil {
		logs.LogTrackingError("CreateProcessorService", "JSON Marshal", ctx, request, err)
		return response.ErrorResponse(http.StatusInternalServerError, constantscore.InvalidResponseBody)
	}
	return response.SuccessResponse(http.StatusCreated, responseBody, constantscore.ItemCreatedSuccessfully)
}
