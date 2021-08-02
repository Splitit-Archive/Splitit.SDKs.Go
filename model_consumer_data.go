/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// ConsumerData struct for ConsumerData
type ConsumerData struct {
	Id string `json:"Id,omitempty"` 
	UniqueId string `json:"UniqueId,omitempty"` 
	UserName string `json:"UserName,omitempty"` 
	FullName string `json:"FullName,omitempty"` 
	Email string `json:"Email,omitempty"` 
	PhoneNumber string `json:"PhoneNumber,omitempty"` 
	CultureName string `json:"CultureName,omitempty"` 
	RoleName string `json:"RoleName,omitempty"` 
	IsLocked bool `json:"IsLocked"` 
	IsDataRestricted bool `json:"IsDataRestricted"` 
	IsDataPrivateRestricted bool `json:"IsDataPrivateRestricted"` 
}
