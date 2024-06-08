# \QueriesAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableProjectsForQuery**](QueriesAPI.md#AvailableProjectsForQuery) | **Get** /api/v3/queries/available_projects | Available projects for query
[**CreateQuery**](QueriesAPI.md#CreateQuery) | **Post** /api/v3/queries | Create query
[**DeleteQuery**](QueriesAPI.md#DeleteQuery) | **Delete** /api/v3/queries/{id} | Delete query
[**EditQuery**](QueriesAPI.md#EditQuery) | **Patch** /api/v3/queries/{id} | Edit Query
[**ListQueries**](QueriesAPI.md#ListQueries) | **Get** /api/v3/queries | List queries
[**QueryCreateForm**](QueriesAPI.md#QueryCreateForm) | **Post** /api/v3/queries/form | Query Create Form
[**QueryUpdateForm**](QueriesAPI.md#QueryUpdateForm) | **Post** /api/v3/queries/{id}/form | Query Update Form
[**StarQuery**](QueriesAPI.md#StarQuery) | **Patch** /api/v3/queries/{id}/star | Star query
[**UnstarQuery**](QueriesAPI.md#UnstarQuery) | **Patch** /api/v3/queries/{id}/unstar | Unstar query
[**ViewDefaultQuery**](QueriesAPI.md#ViewDefaultQuery) | **Get** /api/v3/queries/default | View default query
[**ViewDefaultQueryForProject**](QueriesAPI.md#ViewDefaultQueryForProject) | **Get** /api/v3/projects/{id}/queries/default | View default query for project
[**ViewQuery**](QueriesAPI.md#ViewQuery) | **Get** /api/v3/queries/{id} | View query
[**ViewSchemaForGlobalQueries**](QueriesAPI.md#ViewSchemaForGlobalQueries) | **Get** /api/v3/queries/schema | View schema for global queries
[**ViewSchemaForProjectQueries**](QueriesAPI.md#ViewSchemaForProjectQueries) | **Get** /api/v3/projects/{id}/queries/schema | View schema for project queries



## AvailableProjectsForQuery

> map[string]interface{} AvailableProjectsForQuery(ctx).Execute()

Available projects for query



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
	resp, r, err := apiClient.QueriesAPI.AvailableProjectsForQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.AvailableProjectsForQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableProjectsForQuery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.AvailableProjectsForQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableProjectsForQueryRequest struct via the builder pattern


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


## CreateQuery

> QueryModel CreateQuery(ctx).QueryCreateForm(queryCreateForm).Execute()

Create query



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
	queryCreateForm := *openapiclient.NewQueryCreateForm() // QueryCreateForm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.CreateQuery(context.Background()).QueryCreateForm(queryCreateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.CreateQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuery`: QueryModel
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.CreateQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryCreateForm** | [**QueryCreateForm**](QueryCreateForm.md) |  | 

### Return type

[**QueryModel**](QueryModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuery

> DeleteQuery(ctx, id).Execute()

Delete query



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
	id := int64(1) // int64 | Query id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QueriesAPI.DeleteQuery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.DeleteQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Query id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditQuery

> QueryModel EditQuery(ctx, id).QueryUpdateForm(queryUpdateForm).Execute()

Edit Query



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
	id := int64(1) // int64 | Query id
	queryUpdateForm := *openapiclient.NewQueryUpdateForm() // QueryUpdateForm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.EditQuery(context.Background(), id).QueryUpdateForm(queryUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.EditQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditQuery`: QueryModel
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.EditQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Query id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryUpdateForm** | [**QueryUpdateForm**](QueryUpdateForm.md) |  | 

### Return type

[**QueryModel**](QueryModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueries

> map[string]interface{} ListQueries(ctx).Filters(filters).Execute()

List queries



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
	filters := "[{ "project_id": { "operator": "!*", "values": null }" }]" // string | JSON specifying filter conditions. Currently supported filters are:  + project: filters queries by the project they are assigned to. If the project filter is passed with the `!*` (not any) operator, global queries are returned.  + id: filters queries based on their id  + updated_at: filters queries based on the last time they where updated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.ListQueries(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.ListQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.ListQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions. Currently supported filters are:  + project: filters queries by the project they are assigned to. If the project filter is passed with the &#x60;!*&#x60; (not any) operator, global queries are returned.  + id: filters queries based on their id  + updated_at: filters queries based on the last time they where updated | 

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


## QueryCreateForm

> QueryCreateForm(ctx).QueryCreateForm(queryCreateForm).Execute()

Query Create Form



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
	queryCreateForm := *openapiclient.NewQueryCreateForm() // QueryCreateForm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QueriesAPI.QueryCreateForm(context.Background()).QueryCreateForm(queryCreateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.QueryCreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCreateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryCreateForm** | [**QueryCreateForm**](QueryCreateForm.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryUpdateForm

> QueryUpdateForm(ctx, id).QueryUpdateForm(queryUpdateForm).Execute()

Query Update Form



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
	id := int64(1) // int64 | Query id
	queryUpdateForm := *openapiclient.NewQueryUpdateForm() // QueryUpdateForm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QueriesAPI.QueryUpdateForm(context.Background(), id).QueryUpdateForm(queryUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.QueryUpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Query id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryUpdateForm** | [**QueryUpdateForm**](QueryUpdateForm.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StarQuery

> map[string]interface{} StarQuery(ctx, id).Execute()

Star query



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
	id := int64(1) // int64 | Query id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.StarQuery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.StarQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StarQuery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.StarQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Query id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStarQueryRequest struct via the builder pattern


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


## UnstarQuery

> map[string]interface{} UnstarQuery(ctx, id).Execute()

Unstar query



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
	id := int64(1) // int64 | Query id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.UnstarQuery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.UnstarQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnstarQuery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.UnstarQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Query id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnstarQueryRequest struct via the builder pattern


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


## ViewDefaultQuery

> map[string]interface{} ViewDefaultQuery(ctx).Filters(filters).Offset(offset).PageSize(pageSize).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Timestamps(timestamps).TimelineVisible(timelineVisible).TimelineZoomLevel(timelineZoomLevel).ShowHierarchies(showHierarchies).Execute()

View default query



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
	filters := "[{ "assignee": { "operator": "=", "values": ["1", "5"] }" }]" // string | JSON specifying filter conditions. The filters provided as parameters are not applied to the query but are instead used to override the query's persisted filters. All filters also accepted by the work packages endpoint are accepted. If no filter is to be applied, the client should send an empty array (`[]`). (optional) (default to "[{ \"status_id\": { \"operator\": \"o\", \"values\": null }}]")
	offset := int64(25) // int64 | Page number inside the queries' result collection of work packages. (optional) (default to 1)
	pageSize := int64(25) // int64 | Number of elements to display per page for the queries' result collection of work packages. (optional)
	sortBy := "[["status", "asc"]]" // string | JSON specifying sort criteria. The sort criteria is applied to the query's result collection of work packages overriding the query's persisted sort criteria. (optional) (default to "[[\"id\", \"asc\"]]")
	groupBy := "status" // string | The column to group by. The grouping criteria is applied to the to the query's result collection of work packages overriding the query's persisted group criteria. (optional)
	showSums := true // bool | Indicates whether properties should be summed up if they support it. The showSums parameter is applied to the to the query's result collection of work packages overriding the query's persisted sums property. (optional) (default to false)
	timestamps := "2023-01-01,P-1Y,PT0S,lastWorkingDay@12:00" // string | Indicates the timestamps to filter by when showing changed attributes on work packages. Values can be either ISO8601 dates, ISO8601 durations and the following relative date keywords: \"oneDayAgo@HH:MM+HH:MM\", \"lastWorkingDay@HH:MM+HH:MM\", \"oneWeekAgo@HH:MM+HH:MM\", \"oneMonthAgo@HH:MM+HH:MM\". The first \"HH:MM\" part represents the zero paded hours and minutes. The last \"+HH:MM\" part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\"oneDayAgo@01:00+01:00\", \"oneDayAgo@01:00-01:00\". Values older than 1 day are accepted only with valid Enterprise Token available.  (optional) (default to "PT0S")
	timelineVisible := true // bool | Indicates whether the timeline should be shown. (optional) (default to false)
	timelineZoomLevel := "days" // string | Indicates in what zoom level the timeline should be shown. Valid values are  `days`, `weeks`, `months`, `quarters`, and `years`. (optional) (default to "days")
	showHierarchies := true // bool | Indicates whether the hierarchy mode should be enabled. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.ViewDefaultQuery(context.Background()).Filters(filters).Offset(offset).PageSize(pageSize).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Timestamps(timestamps).TimelineVisible(timelineVisible).TimelineZoomLevel(timelineZoomLevel).ShowHierarchies(showHierarchies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.ViewDefaultQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewDefaultQuery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.ViewDefaultQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewDefaultQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions. The filters provided as parameters are not applied to the query but are instead used to override the query&#39;s persisted filters. All filters also accepted by the work packages endpoint are accepted. If no filter is to be applied, the client should send an empty array (&#x60;[]&#x60;). | [default to &quot;[{ \&quot;status_id\&quot;: { \&quot;operator\&quot;: \&quot;o\&quot;, \&quot;values\&quot;: null }}]&quot;]
 **offset** | **int64** | Page number inside the queries&#39; result collection of work packages. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page for the queries&#39; result collection of work packages. | 
 **sortBy** | **string** | JSON specifying sort criteria. The sort criteria is applied to the query&#39;s result collection of work packages overriding the query&#39;s persisted sort criteria. | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]
 **groupBy** | **string** | The column to group by. The grouping criteria is applied to the to the query&#39;s result collection of work packages overriding the query&#39;s persisted group criteria. | 
 **showSums** | **bool** | Indicates whether properties should be summed up if they support it. The showSums parameter is applied to the to the query&#39;s result collection of work packages overriding the query&#39;s persisted sums property. | [default to false]
 **timestamps** | **string** | Indicates the timestamps to filter by when showing changed attributes on work packages. Values can be either ISO8601 dates, ISO8601 durations and the following relative date keywords: \&quot;oneDayAgo@HH:MM+HH:MM\&quot;, \&quot;lastWorkingDay@HH:MM+HH:MM\&quot;, \&quot;oneWeekAgo@HH:MM+HH:MM\&quot;, \&quot;oneMonthAgo@HH:MM+HH:MM\&quot;. The first \&quot;HH:MM\&quot; part represents the zero paded hours and minutes. The last \&quot;+HH:MM\&quot; part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\&quot;oneDayAgo@01:00+01:00\&quot;, \&quot;oneDayAgo@01:00-01:00\&quot;. Values older than 1 day are accepted only with valid Enterprise Token available.  | [default to &quot;PT0S&quot;]
 **timelineVisible** | **bool** | Indicates whether the timeline should be shown. | [default to false]
 **timelineZoomLevel** | **string** | Indicates in what zoom level the timeline should be shown. Valid values are  &#x60;days&#x60;, &#x60;weeks&#x60;, &#x60;months&#x60;, &#x60;quarters&#x60;, and &#x60;years&#x60;. | [default to &quot;days&quot;]
 **showHierarchies** | **bool** | Indicates whether the hierarchy mode should be enabled. | [default to true]

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


## ViewDefaultQueryForProject

> map[string]interface{} ViewDefaultQueryForProject(ctx, id).Filters(filters).Offset(offset).PageSize(pageSize).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Timestamps(timestamps).TimelineVisible(timelineVisible).ShowHierarchies(showHierarchies).Execute()

View default query for project



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
	id := int64(1) // int64 | Id of the project the default query is requested for
	filters := "[{ "assignee": { "operator": "=", "values": ["1", "5"] }" }]" // string | JSON specifying filter conditions. The filters provided as parameters are not applied to the query but are instead used to override the query's persisted filters. All filters also accepted by the work packages endpoint are accepted. If no filter is to be applied, the client should send an empty array (`[]`). (optional) (default to "[{ \"status_id\": { \"operator\": \"o\", \"values\": null }}]")
	offset := int64(25) // int64 | Page number inside the queries' result collection of work packages. (optional) (default to 1)
	pageSize := int64(25) // int64 | Number of elements to display per page for the queries' result collection of work packages. (optional)
	sortBy := "[["status", "asc"]]" // string | JSON specifying sort criteria. The sort criteria is applied to the query's result collection of work packages overriding the query's persisted sort criteria. (optional) (default to "[[\"id\", \"asc\"]]")
	groupBy := "status" // string | The column to group by. The grouping criteria is applied to the to the query's result collection of work packages overriding the query's persisted group criteria. (optional)
	showSums := true // bool | Indicates whether properties should be summed up if they support it. The showSums parameter is applied to the to the query's result collection of work packages overriding the query's persisted sums property. (optional) (default to false)
	timestamps := "2023-01-01,P-1Y,PT0S,lastWorkingDay@12:00" // string | Indicates the timestamps to filter by when showing changed attributes on work packages. Values can be either ISO8601 dates, ISO8601 durations and the following relative date keywords: \"oneDayAgo@HH:MM+HH:MM\", \"lastWorkingDay@HH:MM+HH:MM\", \"oneWeekAgo@HH:MM+HH:MM\", \"oneMonthAgo@HH:MM+HH:MM\". The first \"HH:MM\" part represents the zero paded hours and minutes. The last \"+HH:MM\" part represents the timezone offset from UTC associated with the time. Values older than 1 day are accepted only with valid Enterprise Token available.  (optional) (default to "PT0S")
	timelineVisible := true // bool | Indicates whether the timeline should be shown. (optional) (default to false)
	showHierarchies := true // bool | Indicates whether the hierarchy mode should be enabled. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.ViewDefaultQueryForProject(context.Background(), id).Filters(filters).Offset(offset).PageSize(pageSize).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Timestamps(timestamps).TimelineVisible(timelineVisible).ShowHierarchies(showHierarchies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.ViewDefaultQueryForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewDefaultQueryForProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.ViewDefaultQueryForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the project the default query is requested for | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewDefaultQueryForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | JSON specifying filter conditions. The filters provided as parameters are not applied to the query but are instead used to override the query&#39;s persisted filters. All filters also accepted by the work packages endpoint are accepted. If no filter is to be applied, the client should send an empty array (&#x60;[]&#x60;). | [default to &quot;[{ \&quot;status_id\&quot;: { \&quot;operator\&quot;: \&quot;o\&quot;, \&quot;values\&quot;: null }}]&quot;]
 **offset** | **int64** | Page number inside the queries&#39; result collection of work packages. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page for the queries&#39; result collection of work packages. | 
 **sortBy** | **string** | JSON specifying sort criteria. The sort criteria is applied to the query&#39;s result collection of work packages overriding the query&#39;s persisted sort criteria. | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]
 **groupBy** | **string** | The column to group by. The grouping criteria is applied to the to the query&#39;s result collection of work packages overriding the query&#39;s persisted group criteria. | 
 **showSums** | **bool** | Indicates whether properties should be summed up if they support it. The showSums parameter is applied to the to the query&#39;s result collection of work packages overriding the query&#39;s persisted sums property. | [default to false]
 **timestamps** | **string** | Indicates the timestamps to filter by when showing changed attributes on work packages. Values can be either ISO8601 dates, ISO8601 durations and the following relative date keywords: \&quot;oneDayAgo@HH:MM+HH:MM\&quot;, \&quot;lastWorkingDay@HH:MM+HH:MM\&quot;, \&quot;oneWeekAgo@HH:MM+HH:MM\&quot;, \&quot;oneMonthAgo@HH:MM+HH:MM\&quot;. The first \&quot;HH:MM\&quot; part represents the zero paded hours and minutes. The last \&quot;+HH:MM\&quot; part represents the timezone offset from UTC associated with the time. Values older than 1 day are accepted only with valid Enterprise Token available.  | [default to &quot;PT0S&quot;]
 **timelineVisible** | **bool** | Indicates whether the timeline should be shown. | [default to false]
 **showHierarchies** | **bool** | Indicates whether the hierarchy mode should be enabled. | [default to true]

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


## ViewQuery

> QueryModel ViewQuery(ctx, id).Filters(filters).Offset(offset).PageSize(pageSize).Columns(columns).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Timestamps(timestamps).TimelineVisible(timelineVisible).TimelineLabels(timelineLabels).HighlightingMode(highlightingMode).HighlightedAttributes(highlightedAttributes).ShowHierarchies(showHierarchies).Execute()

View query



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
	id := int64(1) // int64 | Query id
	filters := "[{ "assignee": { "operator": "=", "values": ["1", "5"] }" }]" // string | JSON specifying filter conditions. The filters provided as parameters are not applied to the query but are instead used to override the query's persisted filters. All filters also accepted by the work packages endpoint are accepted. If no filter is to be applied, the client should send an empty array (`[]`). (optional) (default to "[{ \"status_id\": { \"operator\": \"o\", \"values\": null }}]")
	offset := int64(25) // int64 | Page number inside the queries' result collection of work packages. (optional) (default to 1)
	pageSize := int64(25) // int64 | Number of elements to display per page for the queries' result collection of work packages. (optional)
	columns := "[]" // string | Selected columns for the table view. (optional) (default to "['type', 'priority']")
	sortBy := "[["status", "asc"]]" // string | JSON specifying sort criteria. The sort criteria is applied to the query's result collection of work packages overriding the query's persisted sort criteria. (optional) (default to "[[\"id\", \"asc\"]]")
	groupBy := "status" // string | The column to group by. The grouping criteria is applied to the to the query's result collection of work packages overriding the query's persisted group criteria. (optional)
	showSums := true // bool | Indicates whether properties should be summed up if they support it. The showSums parameter is applied to the to the query's result collection of work packages overriding the query's persisted sums property. (optional) (default to false)
	timestamps := "2023-01-01,P-1Y,PT0S,lastWorkingDay@12:00" // string | Indicates the timestamps to filter by when showing changed attributes on work packages. Values can be either ISO8601 dates, ISO8601 durations and the following relative date keywords: \"oneDayAgo@HH:MM+HH:MM\", \"lastWorkingDay@HH:MM+HH:MM\", \"oneWeekAgo@HH:MM+HH:MM\", \"oneMonthAgo@HH:MM+HH:MM\". The first \"HH:MM\" part represents the zero paded hours and minutes. The last \"+HH:MM\" part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\"oneDayAgo@01:00+01:00\", \"oneDayAgo@01:00-01:00\". Values older than 1 day are accepted only with valid Enterprise Token available.  (optional) (default to "PT0S")
	timelineVisible := true // bool | Indicates whether the timeline should be shown. (optional) (default to false)
	timelineLabels := "{}" // string | Overridden labels in the timeline view (optional) (default to "{}")
	highlightingMode := "inline" // string | Highlighting mode for the table view. (optional) (default to "inline")
	highlightedAttributes := "[]" // string | Highlighted attributes mode for the table view when `highlightingMode` is `inline`. When set to `[]` all highlightable attributes will be returned as `highlightedAttributes`. (optional) (default to "['type', 'priority']")
	showHierarchies := true // bool | Indicates whether the hierarchy mode should be enabled. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueriesAPI.ViewQuery(context.Background(), id).Filters(filters).Offset(offset).PageSize(pageSize).Columns(columns).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Timestamps(timestamps).TimelineVisible(timelineVisible).TimelineLabels(timelineLabels).HighlightingMode(highlightingMode).HighlightedAttributes(highlightedAttributes).ShowHierarchies(showHierarchies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.ViewQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewQuery`: QueryModel
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.ViewQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Query id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | JSON specifying filter conditions. The filters provided as parameters are not applied to the query but are instead used to override the query&#39;s persisted filters. All filters also accepted by the work packages endpoint are accepted. If no filter is to be applied, the client should send an empty array (&#x60;[]&#x60;). | [default to &quot;[{ \&quot;status_id\&quot;: { \&quot;operator\&quot;: \&quot;o\&quot;, \&quot;values\&quot;: null }}]&quot;]
 **offset** | **int64** | Page number inside the queries&#39; result collection of work packages. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page for the queries&#39; result collection of work packages. | 
 **columns** | **string** | Selected columns for the table view. | [default to &quot;[&#39;type&#39;, &#39;priority&#39;]&quot;]
 **sortBy** | **string** | JSON specifying sort criteria. The sort criteria is applied to the query&#39;s result collection of work packages overriding the query&#39;s persisted sort criteria. | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]
 **groupBy** | **string** | The column to group by. The grouping criteria is applied to the to the query&#39;s result collection of work packages overriding the query&#39;s persisted group criteria. | 
 **showSums** | **bool** | Indicates whether properties should be summed up if they support it. The showSums parameter is applied to the to the query&#39;s result collection of work packages overriding the query&#39;s persisted sums property. | [default to false]
 **timestamps** | **string** | Indicates the timestamps to filter by when showing changed attributes on work packages. Values can be either ISO8601 dates, ISO8601 durations and the following relative date keywords: \&quot;oneDayAgo@HH:MM+HH:MM\&quot;, \&quot;lastWorkingDay@HH:MM+HH:MM\&quot;, \&quot;oneWeekAgo@HH:MM+HH:MM\&quot;, \&quot;oneMonthAgo@HH:MM+HH:MM\&quot;. The first \&quot;HH:MM\&quot; part represents the zero paded hours and minutes. The last \&quot;+HH:MM\&quot; part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\&quot;oneDayAgo@01:00+01:00\&quot;, \&quot;oneDayAgo@01:00-01:00\&quot;. Values older than 1 day are accepted only with valid Enterprise Token available.  | [default to &quot;PT0S&quot;]
 **timelineVisible** | **bool** | Indicates whether the timeline should be shown. | [default to false]
 **timelineLabels** | **string** | Overridden labels in the timeline view | [default to &quot;{}&quot;]
 **highlightingMode** | **string** | Highlighting mode for the table view. | [default to &quot;inline&quot;]
 **highlightedAttributes** | **string** | Highlighted attributes mode for the table view when &#x60;highlightingMode&#x60; is &#x60;inline&#x60;. When set to &#x60;[]&#x60; all highlightable attributes will be returned as &#x60;highlightedAttributes&#x60;. | [default to &quot;[&#39;type&#39;, &#39;priority&#39;]&quot;]
 **showHierarchies** | **bool** | Indicates whether the hierarchy mode should be enabled. | [default to true]

### Return type

[**QueryModel**](QueryModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewSchemaForGlobalQueries

> map[string]interface{} ViewSchemaForGlobalQueries(ctx).Execute()

View schema for global queries



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
	resp, r, err := apiClient.QueriesAPI.ViewSchemaForGlobalQueries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.ViewSchemaForGlobalQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewSchemaForGlobalQueries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.ViewSchemaForGlobalQueries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiViewSchemaForGlobalQueriesRequest struct via the builder pattern


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


## ViewSchemaForProjectQueries

> map[string]interface{} ViewSchemaForProjectQueries(ctx, id).Execute()

View schema for project queries



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
	resp, r, err := apiClient.QueriesAPI.ViewSchemaForProjectQueries(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueriesAPI.ViewSchemaForProjectQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewSchemaForProjectQueries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QueriesAPI.ViewSchemaForProjectQueries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewSchemaForProjectQueriesRequest struct via the builder pattern


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

