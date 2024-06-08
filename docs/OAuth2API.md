# \OAuth2API

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthApplication**](OAuth2API.md#GetOauthApplication) | **Get** /api/v3/oauth_applications/{id} | Get the oauth application.
[**GetOauthClientCredentials**](OAuth2API.md#GetOauthClientCredentials) | **Get** /api/v3/oauth_client_credentials/{id} | Get the oauth client credentials object.



## GetOauthApplication

> OAuthApplicationReadModel GetOauthApplication(ctx, id).Execute()

Get the oauth application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	id := int64(1337) // int64 | OAuth application id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2API.GetOauthApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2API.GetOauthApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthApplication`: OAuthApplicationReadModel
	fmt.Fprintf(os.Stdout, "Response from `OAuth2API.GetOauthApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | OAuth application id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthApplicationReadModel**](OAuthApplicationReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthClientCredentials

> OAuthClientCredentialsReadModel GetOauthClientCredentials(ctx, id).Execute()

Get the oauth client credentials object.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	id := int64(1337) // int64 | OAuth Client Credentials id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2API.GetOauthClientCredentials(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2API.GetOauthClientCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthClientCredentials`: OAuthClientCredentialsReadModel
	fmt.Fprintf(os.Stdout, "Response from `OAuth2API.GetOauthClientCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | OAuth Client Credentials id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthClientCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthClientCredentialsReadModel**](OAuthClientCredentialsReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

