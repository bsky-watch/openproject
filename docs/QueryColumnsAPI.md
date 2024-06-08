# \QueryColumnsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewQueryColumn**](QueryColumnsAPI.md#ViewQueryColumn) | **Get** /api/v3/queries/columns/{id} | View Query Column



## ViewQueryColumn

> QueryColumnModel ViewQueryColumn(ctx, id).Execute()

View Query Column



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
	id := "priority" // string | QueryColumn id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryColumnsAPI.ViewQueryColumn(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryColumnsAPI.ViewQueryColumn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewQueryColumn`: QueryColumnModel
	fmt.Fprintf(os.Stdout, "Response from `QueryColumnsAPI.ViewQueryColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | QueryColumn id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewQueryColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryColumnModel**](QueryColumnModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

