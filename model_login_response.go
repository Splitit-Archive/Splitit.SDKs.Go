/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// LoginResponse struct for LoginResponse
type LoginResponse struct {
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"` 
	SessionId string `json:"SessionId,omitempty"` 
	Result int32 `json:"Result"` 
	ResponseStatus *ResponseStatus `json:"ResponseStatus,omitempty"` 
}
