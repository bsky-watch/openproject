# \CustomActionsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteCustomAction**](CustomActionsAPI.md#ExecuteCustomAction) | **Post** /api/v3/custom_actions/{id}/execute | Execute custom action
[**GetCustomAction**](CustomActionsAPI.md#GetCustomAction) | **Get** /api/v3/custom_actions/{id} | Get a custom action



## ExecuteCustomAction

> ExecuteCustomAction(ctx, id).ExecuteCustomActionRequest(executeCustomActionRequest).Execute()

Execute custom action



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
	id := int64(1) // int64 | The id of the custom action to execute
	executeCustomActionRequest := *openapiclient.NewExecuteCustomActionRequest() // ExecuteCustomActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomActionsAPI.ExecuteCustomAction(context.Background(), id).ExecuteCustomActionRequest(executeCustomActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomActionsAPI.ExecuteCustomAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of the custom action to execute | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteCustomActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **executeCustomActionRequest** | [**ExecuteCustomActionRequest**](ExecuteCustomActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomAction

> CustomActionModel GetCustomAction(ctx, id).Execute()

Get a custom action



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
	id := int64(42) // int64 | The id of the custom action to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomActionsAPI.GetCustomAction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomActionsAPI.GetCustomAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomAction`: CustomActionModel
	fmt.Fprintf(os.Stdout, "Response from `CustomActionsAPI.GetCustomAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of the custom action to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomActionModel**](CustomActionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

