/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SwaggerJsonAPIService SwaggerJsonAPI service
type SwaggerJsonAPIService service

type ApiSwaggerJsonRetrieveRequest struct {
	ctx context.Context
	ApiService *SwaggerJsonAPIService
	lang *SwaggerRetrieveLangParameter
	depth *int32
}

func (r ApiSwaggerJsonRetrieveRequest) Lang(lang SwaggerRetrieveLangParameter) ApiSwaggerJsonRetrieveRequest {
	r.lang = &lang
	return r
}

// Serializer Depth
func (r ApiSwaggerJsonRetrieveRequest) Depth(depth int32) ApiSwaggerJsonRetrieveRequest {
	r.depth = &depth
	return r
}

func (r ApiSwaggerJsonRetrieveRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SwaggerJsonRetrieveExecute(r)
}

/*
SwaggerJsonRetrieve Method for SwaggerJsonRetrieve

OpenApi3 schema for this API. Format can be selected via content negotiation.

- YAML: application/vnd.oai.openapi
- JSON: application/vnd.oai.openapi+json

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSwaggerJsonRetrieveRequest
*/
func (a *SwaggerJsonAPIService) SwaggerJsonRetrieve(ctx context.Context) ApiSwaggerJsonRetrieveRequest {
	return ApiSwaggerJsonRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SwaggerJsonAPIService) SwaggerJsonRetrieveExecute(r ApiSwaggerJsonRetrieveRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SwaggerJsonAPIService.SwaggerJsonRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/swagger.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.lang != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lang", r.lang, "form", "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "form", "")
	} else {
		var defaultValue int32 = 1
		r.depth = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.oai.openapi+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
