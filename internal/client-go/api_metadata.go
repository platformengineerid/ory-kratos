/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.

API version:
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type MetadataApi interface {

	/*
			GetVersion Return Running Software Version.

			This endpoint returns the version of Ory Kratos.

		If the service supports TLS Edge Termination, this endpoint does not require the
		`X-Forwarded-Proto` header to be set.

		Be aware that if you are running multiple nodes of this service, the version will never
		refer to the cluster state, only to a single instance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return MetadataApiGetVersionRequest
	*/
	GetVersion(ctx context.Context) MetadataApiGetVersionRequest

	// GetVersionExecute executes the request
	//  @return GetVersion200Response
	GetVersionExecute(r MetadataApiGetVersionRequest) (*GetVersion200Response, *http.Response, error)

	/*
			IsAlive Check HTTP Server Status

			This endpoint returns a HTTP 200 status code when Ory Kratos is accepting incoming
		HTTP requests. This status does currently not include checks whether the database connection is working.

		If the service supports TLS Edge Termination, this endpoint does not require the
		`X-Forwarded-Proto` header to be set.

		Be aware that if you are running multiple nodes of this service, the health status will never
		refer to the cluster state, only to a single instance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return MetadataApiIsAliveRequest
	*/
	IsAlive(ctx context.Context) MetadataApiIsAliveRequest

	// IsAliveExecute executes the request
	//  @return IsAlive200Response
	IsAliveExecute(r MetadataApiIsAliveRequest) (*IsAlive200Response, *http.Response, error)

	/*
			IsReady Check HTTP Server and Database Status

			This endpoint returns a HTTP 200 status code when Ory Kratos is up running and the environment dependencies (e.g.
		the database) are responsive as well.

		If the service supports TLS Edge Termination, this endpoint does not require the
		`X-Forwarded-Proto` header to be set.

		Be aware that if you are running multiple nodes of Ory Kratos, the health status will never
		refer to the cluster state, only to a single instance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return MetadataApiIsReadyRequest
	*/
	IsReady(ctx context.Context) MetadataApiIsReadyRequest

	// IsReadyExecute executes the request
	//  @return IsAlive200Response
	IsReadyExecute(r MetadataApiIsReadyRequest) (*IsAlive200Response, *http.Response, error)
}

// MetadataApiService MetadataApi service
type MetadataApiService service

type MetadataApiGetVersionRequest struct {
	ctx        context.Context
	ApiService MetadataApi
}

func (r MetadataApiGetVersionRequest) Execute() (*GetVersion200Response, *http.Response, error) {
	return r.ApiService.GetVersionExecute(r)
}

/*
GetVersion Return Running Software Version.

This endpoint returns the version of Ory Kratos.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the version will never
refer to the cluster state, only to a single instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MetadataApiGetVersionRequest
*/
func (a *MetadataApiService) GetVersion(ctx context.Context) MetadataApiGetVersionRequest {
	return MetadataApiGetVersionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetVersion200Response
func (a *MetadataApiService) GetVersionExecute(r MetadataApiGetVersionRequest) (*GetVersion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetVersion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.GetVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/version"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type MetadataApiIsAliveRequest struct {
	ctx        context.Context
	ApiService MetadataApi
}

func (r MetadataApiIsAliveRequest) Execute() (*IsAlive200Response, *http.Response, error) {
	return r.ApiService.IsAliveExecute(r)
}

/*
IsAlive Check HTTP Server Status

This endpoint returns a HTTP 200 status code when Ory Kratos is accepting incoming
HTTP requests. This status does currently not include checks whether the database connection is working.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MetadataApiIsAliveRequest
*/
func (a *MetadataApiService) IsAlive(ctx context.Context) MetadataApiIsAliveRequest {
	return MetadataApiIsAliveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IsAlive200Response
func (a *MetadataApiService) IsAliveExecute(r MetadataApiIsAliveRequest) (*IsAlive200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IsAlive200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.IsAlive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health/alive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
		var v string
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type MetadataApiIsReadyRequest struct {
	ctx        context.Context
	ApiService MetadataApi
}

func (r MetadataApiIsReadyRequest) Execute() (*IsAlive200Response, *http.Response, error) {
	return r.ApiService.IsReadyExecute(r)
}

/*
IsReady Check HTTP Server and Database Status

This endpoint returns a HTTP 200 status code when Ory Kratos is up running and the environment dependencies (e.g.
the database) are responsive as well.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of Ory Kratos, the health status will never
refer to the cluster state, only to a single instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MetadataApiIsReadyRequest
*/
func (a *MetadataApiService) IsReady(ctx context.Context) MetadataApiIsReadyRequest {
	return MetadataApiIsReadyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IsAlive200Response
func (a *MetadataApiService) IsReadyExecute(r MetadataApiIsReadyRequest) (*IsAlive200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IsAlive200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.IsReady")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health/ready"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
		if localVarHTTPResponse.StatusCode == 503 {
			var v IsReady503Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v string
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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
