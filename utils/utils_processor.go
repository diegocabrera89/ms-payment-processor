package utils

import (
	"github.com/diegocabrera89/ms-payment-processor/constantsmicro"
	"github.com/diegocabrera89/ms-payment-processor/models"
	"github.com/google/uuid"
	"time"
)

// BuildCreateProcessor build processor object.
func BuildCreateProcessor(processor *models.Processor) {
	processor.ProcessorID = uuid.New().String()   // Generate a unique ID for the client
	processor.CreatedAt = time.Now().UTC().Unix() //Date in UTC
	processor.Status = constantsmicro.StatusProcessorEnable
}
