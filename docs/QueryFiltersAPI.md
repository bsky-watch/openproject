# \QueryFiltersAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewQueryFilter**](QueryFiltersAPI.md#ViewQueryFilter) | **Get** /api/v3/queries/filters/{id} | View Query Filter



## ViewQueryFilter

> QueryFilterModel ViewQueryFilter(ctx, id).Execute()

View Query Filter



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
	id := "status" // string | QueryFilter identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryFiltersAPI.ViewQueryFilter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryFiltersAPI.ViewQueryFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewQueryFilter`: QueryFilterModel
	fmt.Fprintf(os.Stdout, "Response from `QueryFiltersAPI.ViewQueryFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | QueryFilter identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewQueryFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryFilterModel**](QueryFilterModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

