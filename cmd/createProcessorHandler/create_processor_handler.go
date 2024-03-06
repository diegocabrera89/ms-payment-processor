package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-core/middleware/metadata"
	"github.com/diegocabrera89/ms-payment-core/response"
	"github.com/diegocabrera89/ms-payment-processor/constantsmicro"
	"github.com/diegocabrera89/ms-payment-processor/service"
	"net/http"
)

// ProcessorHandler handles HTTP requests related to the Processor entity.
type ProcessorHandler struct {
	processorService *service.ProcessorService
}

// NewProcessorHandler create a new NewProcessorHandler instance.
func NewProcessorHandler() *ProcessorHandler {
	return &ProcessorHandler{
		processorService: service.NewProcessorService(),
	}
}

// CreateProcessorHandler handler for createProcessorHandler new processor.
func (h *ProcessorHandler) CreateProcessorHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("CreateProcessorHandler", ctx, request)
	createProcessorHandler, errorProcessorHandler := h.processorService.CreateProcessorService(ctx, request)
	if errorProcessorHandler != nil {
		logs.LogTrackingError("CreateProcessorHandler", "CreateProcessorService", ctx, request, errorProcessorHandler)
		return response.ErrorResponse(http.StatusInternalServerError, constantsmicro.ErrorCreatingProcessor)
	}
	return createProcessorHandler, nil
}

func main() {
	// Create an instance of PetHandler in the main function.
	processorHandler := NewProcessorHandler()

	// Wrap the handler function with logging middleware.
	handlerWithLogging := metadata.MiddlewareMetadata(processorHandler.CreateProcessorHandler)

	// Start the Lambda handler with the handler function wrapped in the middleware.
	lambda.Start(handlerWithLogging)
}
