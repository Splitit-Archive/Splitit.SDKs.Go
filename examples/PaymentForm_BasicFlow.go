package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/splitit/splitit.sdks.go"
)

func InitiateInstallmentPlan(transactionValue float64) (err error) {
	ctx := context.TODO()

	apiClient := splitit.NewSandboxAPIClient(
		splitit.APIKey(apiKey),
		splitit.Culture("en-US"),
	)

	_, _, err = apiClient.LoginApi.LoginPost(ctx, splitit.LoginRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("Login call failed: %w", err)
	}

	initPaymentResp, _, err := apiClient.InstallmentPlanApi.InstallmentPlanInitiate(
		ctx,
		splitit.InitiateInstallmentPlanRequest{
			PlanData: &splitit.PlanData{
				AutoCapture: true,
				Amount: &splitit.MoneyWithCurrencyCode{
					Value:        transactionValue,
					CurrencyCode: "USD",
				},
				NumberOfInstallments: 6,
			},
			// Optional data to pre-fill the form
			BillingAddress: &splitit.AddressData{
				AddressLine:  "260 Madison Avenue.",
				AddressLine2: "Apartment 1",
				City:         "New York",
				State:        "NY",
				Country:      "USA",
				Zip:          "10016",
			},
			// Optional data to pre-fill the form
			ConsumerData: &splitit.ConsumerData{
				FullName:    "John Smith",
				Email:       "JohnGo@splitit.com",
				PhoneNumber: "1-844-775-4848",
				CultureName: "en-us",
			},
			// After user successfully interacts with splitit.com they would be
			// redirected to provided Succeeded URL with InstallmentPlanNumber as
			// a parameter in GET request. It is required to continue the flow.
			RedirectUrls: &splitit.RedirectUrls{
				Succeeded: "https://example.com/success",
				Canceled:  "https://example.com/canceled",
				Failed:    "https://example.com/failed",
			},
		},
	)
	if err != nil {
		return fmt.Errorf("InstallmentPlanInitiate call failed: %w", err)
	}

	// TODO: Redirect customer to initPaymentResp.CheckoutUrl
	fmt.Printf("Go to %s\n\n", initPaymentResp.CheckoutUrl)
	return nil
}

func VerifyPayment(installmentPlanNumber string, storedTransactionValue float64) (err error) {
	ctx := context.TODO()

	apiClient := splitit.NewSandboxAPIClient(
		splitit.APIKey(apiKey),
		splitit.Culture("en-US"),
	)

	_, _, err = apiClient.LoginApi.LoginPost(ctx, splitit.LoginRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("Login call failed: %w", err)
	}

	resp, _, err := apiClient.InstallmentPlanApi.InstallmentPlanVerifyPayment(
		ctx,
		splitit.VerifyPaymentRequest{
			InstallmentPlanNumber: installmentPlanNumber,
		},
	)
	if err != nil {
		return fmt.Errorf("InstallmentPlanVerifyPayment call failed: %w", err)
	}
	if !resp.IsPaid {
		return errors.New("transaction wasn't payed")
	}

	// TODO: Finances and floats... in flex_fields they are just compared with 0.02 precision. Do the same?
	if resp.OriginalAmountPaid == storedTransactionValue {
		// transaction was successfull
		return nil
	}

	// Transaction amount was tampered with. Try to refund.

	_, _, err = apiClient.InstallmentPlanApi.InstallmentPlanCancel(
		ctx,
		splitit.CancelInstallmentPlanRequest{
			InstallmentPlanNumber:  installmentPlanNumber,
			RefundUnderCancelation: splitit.REFUNDUNDERCANCELATION_ONLY_IF_A_FULL_REFUND_IS_POSSIBLE,
			CancelationReason:      splitit.INSTALLMENTPLANCANCELATIONREASON_OTHER,
			IsExecutedUnattended:   true,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"transaction value was tampered with, but the cancellation refund failed: %w",
			err,
		)
	}
	return errors.New("transaction value was tampered with, payment was refunded")
}

func main() {
	// TODO: store the transaction value for the verification later
	var storedTransactionValue = 300.00
	err := InitiateInstallmentPlan(storedTransactionValue)
	if err != nil {
		fmt.Printf("Failed to initiate payment: %s\n", err)
		return
	}

	// Simulating redirect to Succeeded URL via console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter InstallmentPlanNumber: ")
	installmentPlanNumber, _ := reader.ReadString('\n')
	installmentPlanNumber = strings.TrimSpace(installmentPlanNumber)

	// Succeeded URL handler
	err = VerifyPayment(installmentPlanNumber, storedTransactionValue)
	if err != nil {
		fmt.Printf("Failed to verify payment: %s\n", err)
		return
	}

	fmt.Printf("Installment plan for %.2f was successfully initiated and verified", storedTransactionValue)
}
