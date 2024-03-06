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
	BankID         string `json:"bankID" dynamodbav:"bankID"`
	AccountID      string `json:"accountID" dynamodbav:"accountID"`
	CreatedAt      int64  `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt" dynamodbav:"updatedAt"`
	Status         string `json:"status" dynamodbav:"status"`
}
