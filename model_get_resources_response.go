/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// GetResourcesResponse struct for GetResourcesResponse
type GetResourcesResponse struct {
	Logos map[string]string `json:"Logos,omitempty"`
	TouchPointColors map[string]string `json:"TouchPointColors,omitempty"`
	ResourcesGroupedByCategories *map[string]map[string]string `json:"ResourcesGroupedByCategories,omitempty"`
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"`
}