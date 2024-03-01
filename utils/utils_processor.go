package utils

import (
	"github.com/diegocabrera89/ms-payment-processor/constantsmicro"
	"github.com/diegocabrera89/ms-payment-processor/models"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// BuildCreateProcessor build processor object.
func BuildCreateProcessor(processor *models.Processor) {
	processor.ProcessorID = uuid.New().String()                          // Generate a unique ID for the client
	processor.CreatedAt = strconv.FormatInt(time.Now().UTC().Unix(), 10) // Date in UTC
	processor.Status = constantsmicro.StatusProcessorActive
}
