/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PurchaseMethod 
type PurchaseMethod string

// List of PurchaseMethod
const (
	PURCHASEMETHOD_IN_STORE PurchaseMethod = "InStore"
	PURCHASEMETHOD_PHONE_ORDER PurchaseMethod = "PhoneOrder"
	PURCHASEMETHOD_E_COMMERCE PurchaseMethod = "ECommerce"
)