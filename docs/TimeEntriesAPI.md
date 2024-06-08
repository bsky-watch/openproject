# \TimeEntriesAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableProjectsForTimeEntries**](TimeEntriesAPI.md#AvailableProjectsForTimeEntries) | **Get** /api/v3/time_entries/available_projects | Available projects for time entries
[**CreateTimeEntry**](TimeEntriesAPI.md#CreateTimeEntry) | **Post** /api/v3/time_entries | Create time entry
[**DeleteTimeEntry**](TimeEntriesAPI.md#DeleteTimeEntry) | **Delete** /api/v3/time_entries/{id} | Delete time entry
[**GetTimeEntry**](TimeEntriesAPI.md#GetTimeEntry) | **Get** /api/v3/time_entries/{id} | Get time entry
[**ListTimeEntries**](TimeEntriesAPI.md#ListTimeEntries) | **Get** /api/v3/time_entries | List time entries
[**TimeEntryCreateForm**](TimeEntriesAPI.md#TimeEntryCreateForm) | **Post** /api/v3/time_entries/form | Time entry create form
[**TimeEntryUpdateForm**](TimeEntriesAPI.md#TimeEntryUpdateForm) | **Post** /api/v3/time_entries/{id}/form | Time entry update form
[**UpdateTimeEntry**](TimeEntriesAPI.md#UpdateTimeEntry) | **Patch** /api/v3/time_entries/{id} | update time entry
[**ViewTimeEntrySchema**](TimeEntriesAPI.md#ViewTimeEntrySchema) | **Get** /api/v3/time_entries/schema | View time entry schema



## AvailableProjectsForTimeEntries

> map[string]interface{} AvailableProjectsForTimeEntries(ctx).Execute()

Available projects for time entries



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
	resp, r, err := apiClient.TimeEntriesAPI.AvailableProjectsForTimeEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.AvailableProjectsForTimeEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableProjectsForTimeEntries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TimeEntriesAPI.AvailableProjectsForTimeEntries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableProjectsForTimeEntriesRequest struct via the builder pattern


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


## CreateTimeEntry

> TimeEntryModel CreateTimeEntry(ctx).Execute()

Create time entry



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
	resp, r, err := apiClient.TimeEntriesAPI.CreateTimeEntry(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.CreateTimeEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimeEntry`: TimeEntryModel
	fmt.Fprintf(os.Stdout, "Response from `TimeEntriesAPI.CreateTimeEntry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimeEntryRequest struct via the builder pattern


### Return type

[**TimeEntryModel**](TimeEntryModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTimeEntry

> DeleteTimeEntry(ctx, id).Execute()

Delete time entry



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
	id := int64(1) // int64 | Time entry id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeEntriesAPI.DeleteTimeEntry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.DeleteTimeEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Time entry id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeEntry

> TimeEntryModel GetTimeEntry(ctx, id).Execute()

Get time entry



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
	id := int64(1) // int64 | time entry id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeEntriesAPI.GetTimeEntry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.GetTimeEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeEntry`: TimeEntryModel
	fmt.Fprintf(os.Stdout, "Response from `TimeEntriesAPI.GetTimeEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | time entry id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimeEntryModel**](TimeEntryModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimeEntries

> TimeEntryCollectionModel ListTimeEntries(ctx).Offset(offset).PageSize(pageSize).SortBy(sortBy).Filters(filters).Execute()

List time entries



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
	sortBy := "[["spent_on", "asc"]]" // string | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported sorts are:  + id: Sort by primary key  + hours: Sort by logged hours  + spent_on: Sort by spent on date  + created_at: Sort by time entry creation datetime  + updated_at: Sort by the time the time entry was updated last (optional) (default to "[\"spent_on\", \"asc\"]")
	filters := "[{ "work_package": { "operator": "=", "values": ["1", "2"] } }, { "project": { "operator": "=", "values": ["1"] } }]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + work_package: Filter time entries by work package  + project: Filter time entries by project  + user: Filter time entries by users  + ongoing: Filter for your ongoing timers  + spent_on: Filter time entries by spent on date  + created_at: Filter time entries by creation datetime  + updated_at: Filter time entries by the last time they where updated  + activity: Filter time entries by time entry activity (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeEntriesAPI.ListTimeEntries(context.Background()).Offset(offset).PageSize(pageSize).SortBy(sortBy).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.ListTimeEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTimeEntries`: TimeEntryCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `TimeEntriesAPI.ListTimeEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTimeEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Page number inside the requested collection. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page. | 
 **sortBy** | **string** | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported sorts are:  + id: Sort by primary key  + hours: Sort by logged hours  + spent_on: Sort by spent on date  + created_at: Sort by time entry creation datetime  + updated_at: Sort by the time the time entry was updated last | [default to &quot;[\&quot;spent_on\&quot;, \&quot;asc\&quot;]&quot;]
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + work_package: Filter time entries by work package  + project: Filter time entries by project  + user: Filter time entries by users  + ongoing: Filter for your ongoing timers  + spent_on: Filter time entries by spent on date  + created_at: Filter time entries by creation datetime  + updated_at: Filter time entries by the last time they where updated  + activity: Filter time entries by time entry activity | 

### Return type

[**TimeEntryCollectionModel**](TimeEntryCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimeEntryCreateForm

> TimeEntryCreateForm(ctx).Execute()

Time entry create form



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
	r, err := apiClient.TimeEntriesAPI.TimeEntryCreateForm(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.TimeEntryCreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTimeEntryCreateFormRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimeEntryUpdateForm

> TimeEntryUpdateForm(ctx, id).Body(body).Execute()

Time entry update form



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
	id := int64(1) // int64 | Time entries activity id
	body := int64(56) // int64 | Time entries activity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeEntriesAPI.TimeEntryUpdateForm(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.TimeEntryUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Time entries activity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTimeEntryUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **int64** | Time entries activity id | 

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


## UpdateTimeEntry

> TimeEntryModel UpdateTimeEntry(ctx, id).Execute()

update time entry



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
	id := int64(1) // int64 | Time entry id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeEntriesAPI.UpdateTimeEntry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.UpdateTimeEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTimeEntry`: TimeEntryModel
	fmt.Fprintf(os.Stdout, "Response from `TimeEntriesAPI.UpdateTimeEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Time entry id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTimeEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimeEntryModel**](TimeEntryModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewTimeEntrySchema

> map[string]interface{} ViewTimeEntrySchema(ctx).Execute()

View time entry schema



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
	resp, r, err := apiClient.TimeEntriesAPI.ViewTimeEntrySchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeEntriesAPI.ViewTimeEntrySchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewTimeEntrySchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TimeEntriesAPI.ViewTimeEntrySchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiViewTimeEntrySchemaRequest struct via the builder pattern


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

