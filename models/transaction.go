package models

type PaymentPurposeId uint8

const (
	PaymentPurposeFromOrder                     PaymentPurposeId = 1
	PaymentPurposeGiveSalaryToWorker            PaymentPurposeId = 2
	PaymentPurposeDebtCollectionFromTheEmployee PaymentPurposeId = 3
	PaymentPurposeEmployeeLoan                  PaymentPurposeId = 4
)
