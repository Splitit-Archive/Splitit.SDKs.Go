/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PlanStrategy 
type PlanStrategy string

// List of PlanStrategy
const (
	PLANSTRATEGY_SECURED_PLAN PlanStrategy = "SecuredPlan"
	PLANSTRATEGY_NON_SECURED_PLAN PlanStrategy = "NonSecuredPlan"
	PLANSTRATEGY_SECURED_PLAN_CAPTURE_EXISTING_SECURITY_AUTH PlanStrategy = "SecuredPlanCaptureExistingSecurityAuth"
	PLANSTRATEGY_SECURED_PLAN3 PlanStrategy = "SecuredPlan3"
	PLANSTRATEGY_SINGLE_PAYMENT PlanStrategy = "SinglePayment"
	PLANSTRATEGY_EXTERNAL_INSTALLMENT_PROVIDER PlanStrategy = "ExternalInstallmentProvider"
	PLANSTRATEGY_SECURED_PLAN3_A PlanStrategy = "SecuredPlan3A"
	PLANSTRATEGY_SECURED_PLAN3_B PlanStrategy = "SecuredPlan3B"
	PLANSTRATEGY_SECURED_PLAN2_A PlanStrategy = "SecuredPlan2A"
)
