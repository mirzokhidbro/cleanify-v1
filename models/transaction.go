package models

type PaymentPurposeId uint8

const (
	PaymentPurposeFromOrder             PaymentPurposeId = 1
	PaymentPurposeSalaryOfEmployee      PaymentPurposeId = 2
	PaymentPurposeMoneyReceiverByWorker PaymentPurposeId = 3
)
