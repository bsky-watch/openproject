# \DefaultAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewValuesSchema**](DefaultAPI.md#ViewValuesSchema) | **Get** /api/v3/values/schema/{id} | View Values schema



## ViewValuesSchema

> SchemaModel ViewValuesSchema(ctx, id).Execute()

View Values schema



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
	id := "startDate" // string | The identifier of the value. This is typically the value of the `property` property of the `Values` resource. It should be in lower camelcase format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ViewValuesSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ViewValuesSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewValuesSchema`: SchemaModel
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ViewValuesSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the value. This is typically the value of the &#x60;property&#x60; property of the &#x60;Values&#x60; resource. It should be in lower camelcase format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewValuesSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SchemaModel**](SchemaModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

