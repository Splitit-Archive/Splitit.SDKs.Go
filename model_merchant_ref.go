/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// MerchantRef struct for MerchantRef
type MerchantRef struct {
	Id int64 `json:"Id"`
	Code string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
	Name string `json:"Name,omitempty"`
}