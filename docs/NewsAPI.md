# \NewsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNews**](NewsAPI.md#ListNews) | **Get** /api/v3/news | List News
[**ViewNews**](NewsAPI.md#ViewNews) | **Get** /api/v3/news/{id} | View news



## ListNews

> map[string]interface{} ListNews(ctx).Offset(offset).PageSize(pageSize).SortBy(sortBy).Filters(filters).Execute()

List News



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
	offset := int64(25) // int64 | Page number inside the requested collection. (optional) (default to 1)
	pageSize := int64(25) // int64 | Number of elements to display per page. (optional)
	sortBy := "[["created_at", "asc"]]" // string | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported sorts are:  + id: Sort by primary key  + created_at: Sort by news creation datetime (optional)
	filters := "[{ "project_id": { "operator": "=", "values": ["1", "2"] } }]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + project_id: Filter news by project (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.ListNews(context.Background()).Offset(offset).PageSize(pageSize).SortBy(sortBy).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.ListNews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNews`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.ListNews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Page number inside the requested collection. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page. | 
 **sortBy** | **string** | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported sorts are:  + id: Sort by primary key  + created_at: Sort by news creation datetime | 
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + project_id: Filter news by project | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewNews

> NewsModel ViewNews(ctx, id).Execute()

View news



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
	id := int64(1) // int64 | news id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.ViewNews(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.ViewNews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewNews`: NewsModel
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.ViewNews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | news id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NewsModel**](NewsModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

