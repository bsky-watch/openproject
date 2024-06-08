# \FormsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowOrValidateForm**](FormsAPI.md#ShowOrValidateForm) | **Post** /api/v3/example/form | show or validate form



## ShowOrValidateForm

> map[string]interface{} ShowOrValidateForm(ctx).ShowOrValidateFormRequest(showOrValidateFormRequest).Execute()

show or validate form



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
	showOrValidateFormRequest := *openapiclient.NewShowOrValidateFormRequest() // ShowOrValidateFormRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.ShowOrValidateForm(context.Background()).ShowOrValidateFormRequest(showOrValidateFormRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.ShowOrValidateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowOrValidateForm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.ShowOrValidateForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowOrValidateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showOrValidateFormRequest** | [**ShowOrValidateFormRequest**](ShowOrValidateFormRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

