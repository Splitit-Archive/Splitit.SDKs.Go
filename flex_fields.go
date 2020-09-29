package splitit

import (
	_context "context"
)

type FlexFields struct {
	Client               *APIClient
	NumberOfInstallments int32
	ActiveInitRequest    InitiateInstallmentPlanRequest
}

type FlexFieldsCaptureSettings struct {
	AutoCapture            bool
	FirstInstallmentAmount float64
	CurrencyCode           string
	FirstChargeDate        *SplititTime
}

func NewFlexFields(client *APIClient) *FlexFields {
	return &FlexFields{
		Client: client,
		ActiveInitRequest: InitiateInstallmentPlanRequest{
			PlanData: &PlanData{},
		},
	}
}

func (ff *FlexFields) AddInstallments(installmentOptions string, defaultNumInstallments int32) {
	ff.ActiveInitRequest.PlanData.NumberOfInstallments = defaultNumInstallments

	if ff.ActiveInitRequest.PaymentWizardData == nil {
		ff.ActiveInitRequest.PaymentWizardData = &PaymentWizardData{}
	}

	ff.ActiveInitRequest.PaymentWizardData.RequestedNumberOfInstallments = installmentOptions
}

func (ff *FlexFields) AddBillingInformation(billingAddress *AddressData, consumerData *ConsumerData) {
	ff.ActiveInitRequest.BillingAddress = billingAddress
	ff.ActiveInitRequest.ConsumerData = consumerData
}

func (ff *FlexFields) GetPublicToken(ctx _context.Context, amount float64, currencyCode string) (string, error) {
	ff.ActiveInitRequest.PlanData.Amount = &MoneyWithCurrencyCode{amount, currencyCode}

	initResponse, _, err := ff.Client.InstallmentPlanApi.InstallmentPlanInitiate(ctx, ff.ActiveInitRequest)

	if err != nil {
		return "", err
	}

	return initResponse.PublicToken, nil
}

func (ff *FlexFields) VerifyPayment(ctx _context.Context, planNumber string, orderAmount float64) (bool, error) {
	req := VerifyPaymentRequest{
		InstallmentPlanNumber: planNumber,
	}

	res, _, err := ff.Client.InstallmentPlanApi.InstallmentPlanVerifyPayment(ctx, req)
	if err != nil {
		return false, err
	}

	return res.IsPaid && res.OriginalAmountPaid == orderAmount, err
}

func (ff *FlexFields) Add3DSecure(redirectUrls *RedirectUrls) {
	ff.ActiveInitRequest.PlanData.Attempt3DSecure = true
	ff.ActiveInitRequest.RedirectUrls = redirectUrls
}

func (ff *FlexFields) AddCaptureSettings(captureSettings *FlexFieldsCaptureSettings) {
	ff.ActiveInitRequest.PlanData.AutoCapture = captureSettings.AutoCapture

	if captureSettings.FirstInstallmentAmount > 0 {
		ff.ActiveInitRequest.PlanData.FirstInstallmentAmount = &MoneyWithCurrencyCode{captureSettings.FirstInstallmentAmount, captureSettings.CurrencyCode}
	}

	if captureSettings.FirstChargeDate != nil {
		ff.ActiveInitRequest.PlanData.FirstChargeDate = captureSettings.FirstChargeDate
	}
}
