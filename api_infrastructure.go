/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// InfrastructureApiService InfrastructureApi service
type InfrastructureApiService service

/*
InfrastructureGetResources Method for InfrastructureGetResources
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param request
@return GetResourcesResponse
*/
func (a *InfrastructureApiService) InfrastructureGetResources(ctx _context.Context, request GetResourcesRequest) (GetResourcesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetResourcesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/Infrastructure/GetResources"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &request
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GetResourcesResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarReturnValue.ResponseHeader.Succeeded != true {
		if len(*localVarReturnValue.ResponseHeader.Errors) > 0 {
			newErr := GenericOpenAPIError{
				body:  localVarBody,
				error: (*localVarReturnValue.ResponseHeader.Errors)[0].Message,
				model: localVarReturnValue.ResponseHeader,
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: "Error ocurred while contacting Splitit API.",
			model: localVarReturnValue.ResponseHeader,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// InfrastructureGetResources2Opts Optional parameters for the method 'InfrastructureGetResources2'
type InfrastructureGetResources2Opts struct {
    ApiKey optional.String
    SessionId optional.String
    MerchantCode optional.String
    CultureName optional.String
    TouchPointCode optional.String
    SystemTextCategories optional.Interface
}

/*
InfrastructureGetResources2 Method for InfrastructureGetResources2
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InfrastructureGetResources2Opts - Optional Parameters:
 * @param "ApiKey" (optional.String) - 
 * @param "SessionId" (optional.String) - 
 * @param "MerchantCode" (optional.String) - 
 * @param "CultureName" (optional.String) - 
 * @param "TouchPointCode" (optional.String) - 
 * @param "SystemTextCategories" (optional.Interface of []SystemTextCategory) - 
@return GetResourcesResponse
*/
func (a *InfrastructureApiService) InfrastructureGetResources2(ctx _context.Context, localVarOptionals *InfrastructureGetResources2Opts) (GetResourcesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetResourcesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/Infrastructure/GetResources"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ApiKey.IsSet() {
		localVarQueryParams.Add("apiKey", parameterToString(localVarOptionals.ApiKey.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SessionId.IsSet() {
		localVarQueryParams.Add("sessionId", parameterToString(localVarOptionals.SessionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MerchantCode.IsSet() {
		localVarQueryParams.Add("merchantCode", parameterToString(localVarOptionals.MerchantCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CultureName.IsSet() {
		localVarQueryParams.Add("cultureName", parameterToString(localVarOptionals.CultureName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TouchPointCode.IsSet() {
		localVarQueryParams.Add("touchPointCode", parameterToString(localVarOptionals.TouchPointCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SystemTextCategories.IsSet() {
		t:=localVarOptionals.SystemTextCategories.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("systemTextCategories", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("systemTextCategories", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GetResourcesResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarReturnValue.ResponseHeader.Succeeded != true {
		if len(*localVarReturnValue.ResponseHeader.Errors) > 0 {
			newErr := GenericOpenAPIError{
				body:  localVarBody,
				error: (*localVarReturnValue.ResponseHeader.Errors)[0].Message,
				model: localVarReturnValue.ResponseHeader,
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: "Error ocurred while contacting Splitit API.",
			model: localVarReturnValue.ResponseHeader,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}