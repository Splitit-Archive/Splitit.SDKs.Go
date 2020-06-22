package splitit

import (
	_context "context"
	"fmt"
	"math"
)

type FlexFields struct {
	ActiveContext        _context.Context
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

func FlexFieldsAuthenticate(ctx _context.Context, client *APIClient, username string, pass string) (*FlexFields, error) {
	loginResponse, _, err := client.LoginApi.LoginPost(ctx, LoginRequest{
		UserName: username,
		Password: pass,
	})

	if err != nil {
		return nil, err
	}

	ff := &FlexFields{
		ActiveContext: _context.WithValue(ctx, "splitit.context", SplititRequestContext{
			SessionId: loginResponse.SessionId,
		}),
		Client: client,
		ActiveInitRequest: InitiateInstallmentPlanRequest{
			PlanData: &PlanData{},
		},
	}

	return ff, nil
}

func (ff *FlexFields) AddCulture(culture string) {
	reqContext, _ := ff.ActiveContext.Value("splitit.context").(SplititRequestContext)
	reqContext.Culture = culture

	ff.ActiveContext = _context.WithValue(ff.ActiveContext, "splitit.context", reqContext)
}

func (ff *FlexFields) AddInstallments(installmentOptions string, defaultNumInstallments int32) {
	ff.ActiveInitRequest.PlanData.NumberOfInstallments = defaultNumInstallments

	fmt.Println(ff.ActiveInitRequest.PlanData.NumberOfInstallments)

	if ff.ActiveInitRequest.PaymentWizardData == nil {
		ff.ActiveInitRequest.PaymentWizardData = &PaymentWizardData{}
	}

	ff.ActiveInitRequest.PaymentWizardData.RequestedNumberOfInstallments = installmentOptions
}

func (ff *FlexFields) AddBillingInformation(billingAddress *AddressData, consumerData *ConsumerData) {
	ff.ActiveInitRequest.BillingAddress = billingAddress
	ff.ActiveInitRequest.ConsumerData = consumerData
}

func (ff *FlexFields) GetPublicToken(amount float64, currencyCode string) (string, error) {
	ff.ActiveInitRequest.PlanData.Amount = &MoneyWithCurrencyCode{
		Value:        amount,
		CurrencyCode: currencyCode,
	}

	initResponse, _, err := ff.Client.InstallmentPlanApi.InstallmentPlanInitiate(ff.ActiveContext, ff.ActiveInitRequest)

	if err != nil {
		return "", err
	}

	return initResponse.PublicToken, nil
}

func (ff *FlexFields) VerifyPayment(planNumber string, orderAmount float64) (bool, error) {
	req := VerifyPaymentRequest{
		InstallmentPlanNumber: planNumber,
	}

	res, _, err := ff.Client.InstallmentPlanApi.InstallmentPlanVerifyPayment(ff.ActiveContext, req)
	if err != nil {
		return false, err
	}

	return res.IsPaid && math.Abs(res.OriginalAmountPaid-orderAmount) < 0.02, err
}

func (ff *FlexFields) Add3DSecure(redirectUrls *RedirectUrls) {
	ff.ActiveInitRequest.PlanData.Attempt3DSecure = true
	ff.ActiveInitRequest.RedirectUrls = redirectUrls
}

func (ff *FlexFields) AddCaptureSettings(captureSettings *FlexFieldsCaptureSettings) {
	ff.ActiveInitRequest.PlanData.AutoCapture = captureSettings.AutoCapture

	if captureSettings.FirstInstallmentAmount > 0 {
		ff.ActiveInitRequest.PlanData.FirstInstallmentAmount = &MoneyWithCurrencyCode{
			Value:        captureSettings.FirstInstallmentAmount,
			CurrencyCode: captureSettings.CurrencyCode,
		}
	}

	if captureSettings.FirstChargeDate != nil {
		ff.ActiveInitRequest.PlanData.FirstChargeDate = captureSettings.FirstChargeDate
	}
}
