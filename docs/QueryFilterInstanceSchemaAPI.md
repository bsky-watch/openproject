# \QueryFilterInstanceSchemaAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQueryFilterInstanceSchemas**](QueryFilterInstanceSchemaAPI.md#ListQueryFilterInstanceSchemas) | **Get** /api/v3/queries/filter_instance_schemas | List Query Filter Instance Schemas
[**ListQueryFilterInstanceSchemasForProject**](QueryFilterInstanceSchemaAPI.md#ListQueryFilterInstanceSchemasForProject) | **Get** /api/v3/projects/{id}/queries/filter_instance_schemas | List Query Filter Instance Schemas for Project
[**ViewQueryFilterInstanceSchema**](QueryFilterInstanceSchemaAPI.md#ViewQueryFilterInstanceSchema) | **Get** /api/v3/queries/filter_instance_schemas/{id} | View Query Filter Instance Schema



## ListQueryFilterInstanceSchemas

> map[string]interface{} ListQueryFilterInstanceSchemas(ctx).Execute()

List Query Filter Instance Schemas



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
	resp, r, err := apiClient.QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueryFilterInstanceSchemas`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryFilterInstanceSchemasRequest struct via the builder pattern


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


## ListQueryFilterInstanceSchemasForProject

> map[string]interface{} ListQueryFilterInstanceSchemasForProject(ctx, id).Execute()

List Query Filter Instance Schemas for Project



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
	id := int64(1) // int64 | Project id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemasForProject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemasForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueryFilterInstanceSchemasForProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueryFilterInstanceSchemaAPI.ListQueryFilterInstanceSchemasForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueryFilterInstanceSchemasForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ViewQueryFilterInstanceSchema

> QueryFilterInstanceSchemaModel ViewQueryFilterInstanceSchema(ctx, id).Execute()

View Query Filter Instance Schema



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
	id := "author" // string | QueryFilterInstanceSchema identifier. The identifier is the filter identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueryFilterInstanceSchemaAPI.ViewQueryFilterInstanceSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueryFilterInstanceSchemaAPI.ViewQueryFilterInstanceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewQueryFilterInstanceSchema`: QueryFilterInstanceSchemaModel
	fmt.Fprintf(os.Stdout, "Response from `QueryFilterInstanceSchemaAPI.ViewQueryFilterInstanceSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | QueryFilterInstanceSchema identifier. The identifier is the filter identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewQueryFilterInstanceSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryFilterInstanceSchemaModel**](QueryFilterInstanceSchemaModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

