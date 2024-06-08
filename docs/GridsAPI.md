# \GridsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGrid**](GridsAPI.md#CreateGrid) | **Post** /api/v3/grids | Create a grid
[**GetGrid**](GridsAPI.md#GetGrid) | **Get** /api/v3/grids/{id} | Get a grid
[**GridCreateForm**](GridsAPI.md#GridCreateForm) | **Post** /api/v3/grids/form | Grid Create Form
[**GridUpdateForm**](GridsAPI.md#GridUpdateForm) | **Post** /api/v3/grids/{id}/form | Grid Update Form
[**ListGrids**](GridsAPI.md#ListGrids) | **Get** /api/v3/grids | List grids
[**UpdateGrid**](GridsAPI.md#UpdateGrid) | **Patch** /api/v3/grids/{id} | Update a grid



## CreateGrid

> GridReadModel CreateGrid(ctx).GridWriteModel(gridWriteModel).Execute()

Create a grid



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
	gridWriteModel := *openapiclient.NewGridWriteModel() // GridWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridsAPI.CreateGrid(context.Background()).GridWriteModel(gridWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridsAPI.CreateGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGrid`: GridReadModel
	fmt.Fprintf(os.Stdout, "Response from `GridsAPI.CreateGrid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gridWriteModel** | [**GridWriteModel**](GridWriteModel.md) |  | 

### Return type

[**GridReadModel**](GridReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGrid

> GridReadModel GetGrid(ctx, id).Execute()

Get a grid



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
	id := int64(42) // int64 | Grid id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridsAPI.GetGrid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridsAPI.GetGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGrid`: GridReadModel
	fmt.Fprintf(os.Stdout, "Response from `GridsAPI.GetGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Grid id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GridReadModel**](GridReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridCreateForm

> GridCreateForm(ctx).Execute()

Grid Create Form



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GridsAPI.GridCreateForm(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridsAPI.GridCreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGridCreateFormRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridUpdateForm

> map[string]interface{} GridUpdateForm(ctx, id).Execute()

Grid Update Form



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
	id := int64(1) // int64 | ID of the grid being modified

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridsAPI.GridUpdateForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridsAPI.GridUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridUpdateForm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GridsAPI.GridUpdateForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the grid being modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrids

> GridCollectionModel ListGrids(ctx).Offset(offset).PageSize(pageSize).Filters(filters).Execute()

List grids



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
	pageSize := int64(25) // int64 | Number of elements to display per page. (optional) (default to 30)
	filters := "[{ "page": { "operator": "=", "values": ["/my/page"] } }]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  - page: Filter grid by work package (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridsAPI.ListGrids(context.Background()).Offset(offset).PageSize(pageSize).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridsAPI.ListGrids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrids`: GridCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `GridsAPI.ListGrids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGridsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Page number inside the requested collection. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page. | [default to 30]
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  - page: Filter grid by work package | 

### Return type

[**GridCollectionModel**](GridCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGrid

> GridReadModel UpdateGrid(ctx).GridWriteModel(gridWriteModel).Execute()

Update a grid



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
	gridWriteModel := *openapiclient.NewGridWriteModel() // GridWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridsAPI.UpdateGrid(context.Background()).GridWriteModel(gridWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridsAPI.UpdateGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGrid`: GridReadModel
	fmt.Fprintf(os.Stdout, "Response from `GridsAPI.UpdateGrid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gridWriteModel** | [**GridWriteModel**](GridWriteModel.md) |  | 

### Return type

[**GridReadModel**](GridReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

