/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// GetInstallmentsPlanExtendedResponse struct for GetInstallmentsPlanExtendedResponse
type GetInstallmentsPlanExtendedResponse struct {
	PlansList []InstallmentPlan `json:"PlansList,omitempty"` 
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"` 
	PagingResponseHeader *PagingResponseHeader `json:"PagingResponseHeader,omitempty"` 
}
