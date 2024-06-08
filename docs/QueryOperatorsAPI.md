# \QueryOperatorsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewQueryOperator**](QueryOperatorsAPI.md#ViewQueryOperator) | **Get** /api/v3/queries/operators/{id} | View Query Operator



## ViewQueryOperator

> QueryOperatorModel ViewQueryOperator(ctx, id).Execute()

View Query Operator



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
	id := "!" // string | QueryOperator id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryOperatorsAPI.ViewQueryOperator(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryOperatorsAPI.ViewQueryOperator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewQueryOperator`: QueryOperatorModel
	fmt.Fprintf(os.Stdout, "Response from `QueryOperatorsAPI.ViewQueryOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | QueryOperator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewQueryOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryOperatorModel**](QueryOperatorModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

