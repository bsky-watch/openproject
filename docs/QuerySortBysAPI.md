# \QuerySortBysAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewQuerySortBy**](QuerySortBysAPI.md#ViewQuerySortBy) | **Get** /api/v3/queries/sort_bys/{id} | View Query Sort By



## ViewQuerySortBy

> QuerySortByModel ViewQuerySortBy(ctx, id).Execute()

View Query Sort By



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
	id := "status-asc" // string | QuerySortBy identifier. The identifier is a combination of the column identifier and the direction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuerySortBysAPI.ViewQuerySortBy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuerySortBysAPI.ViewQuerySortBy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewQuerySortBy`: QuerySortByModel
	fmt.Fprintf(os.Stdout, "Response from `QuerySortBysAPI.ViewQuerySortBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | QuerySortBy identifier. The identifier is a combination of the column identifier and the direction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewQuerySortByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuerySortByModel**](QuerySortByModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

