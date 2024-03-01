package models

// Processor structure to define Processor fields.
type Processor struct {
	ProcessorID    string `json:"processorID" dynamodbav:"processorID"`
	Name           string `json:"name" dynamodbav:"name"`
	Model          string `json:"model" dynamodbav:"model"`
	Nick           string `json:"nick" dynamodbav:"nick"`
	MCC            string `json:"mcc" dynamodbav:"mcc"`
	TerminalID     string `json:"terminalID" dynamodbav:"terminalID"`
	MerchantID     string `json:"merchantID" dynamodbav:"merchantID"`
	SoftDescriptor string `json:"softDescriptor" dynamodbav:"softDescriptor"`
	CreatedAt      string `json:"createdAt" dynamodbav:"createdAt"`
	Status         string `json:"status" dynamodbav:"status"`
}
