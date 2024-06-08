# \WorkPackagesAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWatcher**](WorkPackagesAPI.md#AddWatcher) | **Post** /api/v3/work_packages/{id}/watchers | Add watcher
[**AvailableProjectsForWorkPackage**](WorkPackagesAPI.md#AvailableProjectsForWorkPackage) | **Get** /api/v3/work_packages/{id}/available_projects | Available projects for work package
[**AvailableWatchers**](WorkPackagesAPI.md#AvailableWatchers) | **Get** /api/v3/work_packages/{id}/available_watchers | Available watchers
[**CommentWorkPackage**](WorkPackagesAPI.md#CommentWorkPackage) | **Post** /api/v3/work_packages/{id}/activities | Comment work package
[**CreateProjectWorkPackage**](WorkPackagesAPI.md#CreateProjectWorkPackage) | **Post** /api/v3/projects/{id}/work_packages | Create work package in project
[**CreateRelation**](WorkPackagesAPI.md#CreateRelation) | **Post** /api/v3/work_packages/{id}/relations | Create Relation
[**CreateWorkPackage**](WorkPackagesAPI.md#CreateWorkPackage) | **Post** /api/v3/work_packages | Create Work Package
[**CreateWorkPackageFileLink**](WorkPackagesAPI.md#CreateWorkPackageFileLink) | **Post** /api/v3/work_packages/{id}/file_links | Creates file links.
[**DeleteWorkPackage**](WorkPackagesAPI.md#DeleteWorkPackage) | **Delete** /api/v3/work_packages/{id} | Delete Work Package
[**GetProjectWorkPackageCollection**](WorkPackagesAPI.md#GetProjectWorkPackageCollection) | **Get** /api/v3/projects/{id}/work_packages | Get work packages of project
[**ListAvailableRelationCandidates**](WorkPackagesAPI.md#ListAvailableRelationCandidates) | **Get** /api/v3/work_packages/{id}/available_relation_candidates | Available relation candidates
[**ListWatchers**](WorkPackagesAPI.md#ListWatchers) | **Get** /api/v3/work_packages/{id}/watchers | List watchers
[**ListWorkPackageActivities**](WorkPackagesAPI.md#ListWorkPackageActivities) | **Get** /api/v3/work_packages/{id}/activities | List work package activities
[**ListWorkPackageFileLinks**](WorkPackagesAPI.md#ListWorkPackageFileLinks) | **Get** /api/v3/work_packages/{id}/file_links | Gets all file links of a work package
[**ListWorkPackageRelations**](WorkPackagesAPI.md#ListWorkPackageRelations) | **Get** /api/v3/work_packages/{id}/relations | List relations
[**ListWorkPackageSchemas**](WorkPackagesAPI.md#ListWorkPackageSchemas) | **Get** /api/v3/work_packages/schemas | List Work Package Schemas
[**ListWorkPackages**](WorkPackagesAPI.md#ListWorkPackages) | **Get** /api/v3/work_packages | List work packages
[**ProjectAvailableAssignees**](WorkPackagesAPI.md#ProjectAvailableAssignees) | **Get** /api/v3/projects/{id}/available_assignees | Project Available assignees
[**RemoveWatcher**](WorkPackagesAPI.md#RemoveWatcher) | **Delete** /api/v3/work_packages/{id}/watchers/{user_id} | Remove watcher
[**Revisions**](WorkPackagesAPI.md#Revisions) | **Get** /api/v3/work_packages/{id}/revisions | Revisions
[**UpdateWorkPackage**](WorkPackagesAPI.md#UpdateWorkPackage) | **Patch** /api/v3/work_packages/{id} | Update a Work Package
[**ViewWorkPackage**](WorkPackagesAPI.md#ViewWorkPackage) | **Get** /api/v3/work_packages/{id} | View Work Package
[**ViewWorkPackageSchema**](WorkPackagesAPI.md#ViewWorkPackageSchema) | **Get** /api/v3/work_packages/schemas/{identifier} | View Work Package Schema
[**WorkPackageAvailableAssignees**](WorkPackagesAPI.md#WorkPackageAvailableAssignees) | **Get** /api/v3/work_packages/{id}/available_assignees | Work Package Available assignees
[**WorkPackageCreateForm**](WorkPackagesAPI.md#WorkPackageCreateForm) | **Post** /api/v3/work_packages/form | Work Package Create Form
[**WorkPackageCreateFormForProject**](WorkPackagesAPI.md#WorkPackageCreateFormForProject) | **Post** /api/v3/projects/{id}/work_packages/form | Work Package Create Form For Project
[**WorkPackageEditForm**](WorkPackagesAPI.md#WorkPackageEditForm) | **Post** /api/v3/work_packages/{id}/form | Work Package Edit Form



## AddWatcher

> AddWatcher(ctx, id).AddWatcherRequest(addWatcherRequest).Execute()

Add watcher



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
	id := int64(1) // int64 | Work package id
	addWatcherRequest := *openapiclient.NewAddWatcherRequest() // AddWatcherRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.AddWatcher(context.Background(), id).AddWatcherRequest(addWatcherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.AddWatcher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWatcherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addWatcherRequest** | [**AddWatcherRequest**](AddWatcherRequest.md) |  | 

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


## AvailableProjectsForWorkPackage

> map[string]interface{} AvailableProjectsForWorkPackage(ctx, id).Execute()

Available projects for work package



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
	id := int64(1) // int64 | work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.AvailableProjectsForWorkPackage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.AvailableProjectsForWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableProjectsForWorkPackage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.AvailableProjectsForWorkPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableProjectsForWorkPackageRequest struct via the builder pattern


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


## AvailableWatchers

> map[string]interface{} AvailableWatchers(ctx, id).Execute()

Available watchers



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
	id := int64(1) // int64 | work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.AvailableWatchers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.AvailableWatchers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailableWatchers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.AvailableWatchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableWatchersRequest struct via the builder pattern


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


## CommentWorkPackage

> ActivityModel CommentWorkPackage(ctx, id).Notify(notify).CommentWorkPackageRequest(commentWorkPackageRequest).Execute()

Comment work package



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
	id := int64(1) // int64 | Work package id
	notify := false // bool | Indicates whether change notifications (e.g. via E-Mail) should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. (optional) (default to true)
	commentWorkPackageRequest := *openapiclient.NewCommentWorkPackageRequest() // CommentWorkPackageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.CommentWorkPackage(context.Background(), id).Notify(notify).CommentWorkPackageRequest(commentWorkPackageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.CommentWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentWorkPackage`: ActivityModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.CommentWorkPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommentWorkPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notify** | **bool** | Indicates whether change notifications (e.g. via E-Mail) should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. | [default to true]
 **commentWorkPackageRequest** | [**CommentWorkPackageRequest**](CommentWorkPackageRequest.md) |  | 

### Return type

[**ActivityModel**](ActivityModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectWorkPackage

> WorkPackageModel CreateProjectWorkPackage(ctx, id).Notify(notify).WorkPackageModel(workPackageModel).Execute()

Create work package in project



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
	notify := false // bool | Indicates whether change notifications (e.g. via E-Mail) should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. (optional) (default to true)
	workPackageModel := *openapiclient.NewWorkPackageModel() // WorkPackageModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.CreateProjectWorkPackage(context.Background(), id).Notify(notify).WorkPackageModel(workPackageModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.CreateProjectWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectWorkPackage`: WorkPackageModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.CreateProjectWorkPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectWorkPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notify** | **bool** | Indicates whether change notifications (e.g. via E-Mail) should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. | [default to true]
 **workPackageModel** | [**WorkPackageModel**](WorkPackageModel.md) |  | 

### Return type

[**WorkPackageModel**](WorkPackageModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRelation

> CreateRelation(ctx, id).CreateRelationRequest(createRelationRequest).Execute()

Create Relation



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
	id := int64(1) // int64 | Work package id
	createRelationRequest := *openapiclient.NewCreateRelationRequest("Type_example", *openapiclient.NewCreateRelationRequestLinks(*openapiclient.NewLink("Href_example"), *openapiclient.NewLink("Href_example"))) // CreateRelationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.CreateRelation(context.Background(), id).CreateRelationRequest(createRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.CreateRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRelationRequest** | [**CreateRelationRequest**](CreateRelationRequest.md) |  | 

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


## CreateWorkPackage

> WorkPackageModel CreateWorkPackage(ctx).Notify(notify).WorkPackageModel(workPackageModel).Execute()

Create Work Package



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
	notify := false // bool | Indicates whether change notifications (e.g. via E-Mail) should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. (optional) (default to true)
	workPackageModel := *openapiclient.NewWorkPackageModel() // WorkPackageModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.CreateWorkPackage(context.Background()).Notify(notify).WorkPackageModel(workPackageModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.CreateWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkPackage`: WorkPackageModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.CreateWorkPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notify** | **bool** | Indicates whether change notifications (e.g. via E-Mail) should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. | [default to true]
 **workPackageModel** | [**WorkPackageModel**](WorkPackageModel.md) |  | 

### Return type

[**WorkPackageModel**](WorkPackageModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkPackageFileLink

> FileLinkCollectionReadModel CreateWorkPackageFileLink(ctx, id).FileLinkCollectionWriteModel(fileLinkCollectionWriteModel).Execute()

Creates file links.



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
	id := int64(1337) // int64 | Work package id
	fileLinkCollectionWriteModel := *openapiclient.NewFileLinkCollectionWriteModel("Type_example", int64(123), int64(123), *openapiclient.NewCollectionModelLinks(*openapiclient.NewCollectionModelLinksSelf("Href_example")), *openapiclient.NewFileLinkCollectionWriteModelAllOfEmbedded([]openapiclient.FileLinkWriteModel{*openapiclient.NewFileLinkWriteModel(*openapiclient.NewFileLinkOriginDataModel("Id_example", "Name_example"), openapiclient.FileLinkWriteModel__links{FileLinkWriteModelLinksOneOf: openapiclient.NewFileLinkWriteModelLinksOneOf(*openapiclient.NewFileLinkReadModelLinksStorage("Href_example"))})})) // FileLinkCollectionWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.CreateWorkPackageFileLink(context.Background(), id).FileLinkCollectionWriteModel(fileLinkCollectionWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.CreateWorkPackageFileLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkPackageFileLink`: FileLinkCollectionReadModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.CreateWorkPackageFileLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkPackageFileLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkCollectionWriteModel** | [**FileLinkCollectionWriteModel**](FileLinkCollectionWriteModel.md) |  | 

### Return type

[**FileLinkCollectionReadModel**](FileLinkCollectionReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkPackage

> DeleteWorkPackage(ctx, id).Execute()

Delete Work Package



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
	id := int64(1) // int64 | Work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.DeleteWorkPackage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.DeleteWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkPackageRequest struct via the builder pattern


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


## GetProjectWorkPackageCollection

> WorkPackagesModel GetProjectWorkPackageCollection(ctx, id).Offset(offset).PageSize(pageSize).Filters(filters).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Select_(select_).Execute()

Get work packages of project



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
	offset := int64(25) // int64 | Page number inside the requested collection. (optional) (default to 1)
	pageSize := int64(25) // int64 | Number of elements to display per page. (optional)
	filters := "[{ "type_id": { "operator": "=", "values": ['1', '2'] }}]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. If no filter is to be applied, the client should send an empty array (`[]`). (optional) (default to "[{ \"status_id\": { \"operator\": \"o\", \"values\": null }}]")
	sortBy := "[["status", "asc"]]" // string | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. (optional) (default to "[[\"id\", \"asc\"]]")
	groupBy := "status" // string | The column to group by. (optional)
	showSums := true // bool | Indicates whether properties should be summed up if they support it. (optional) (default to false)
	select_ := "total,elements/subject,elements/id,self" // string | Comma separated list of properties to include. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.GetProjectWorkPackageCollection(context.Background(), id).Offset(offset).PageSize(pageSize).Filters(filters).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.GetProjectWorkPackageCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectWorkPackageCollection`: WorkPackagesModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.GetProjectWorkPackageCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectWorkPackageCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | Page number inside the requested collection. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page. | 
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. If no filter is to be applied, the client should send an empty array (&#x60;[]&#x60;). | [default to &quot;[{ \&quot;status_id\&quot;: { \&quot;operator\&quot;: \&quot;o\&quot;, \&quot;values\&quot;: null }}]&quot;]
 **sortBy** | **string** | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]
 **groupBy** | **string** | The column to group by. | 
 **showSums** | **bool** | Indicates whether properties should be summed up if they support it. | [default to false]
 **select_** | **string** | Comma separated list of properties to include. | 

### Return type

[**WorkPackagesModel**](WorkPackagesModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableRelationCandidates

> map[string]interface{} ListAvailableRelationCandidates(ctx, id).PageSize(pageSize).Filters(filters).Query(query).Type_(type_).SortBy(sortBy).Execute()

Available relation candidates



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
	pageSize := int64(25) // int64 | Maximum number of candidates to list (default 10) (optional)
	filters := "[{ "status_id": { "operator": "o", "values": null } }]" // string | JSON specifying filter conditions. Accepts the same filters as the [work packages](https://www.openproject.org/docs/api/endpoints/work-packages/) endpoint. (optional)
	query := ""rollout"" // string | Shortcut for filtering by ID or subject (optional)
	type_ := ""follows"" // string | Type of relation to find candidates for (default \"relates\") (optional)
	sortBy := "[["status", "asc"]]" // string | JSON specifying sort criteria. Accepts the same sort criteria as the [work packages](https://www.openproject.org/docs/api/endpoints/work-packages/) endpoint. (optional) (default to "[[\"id\", \"asc\"]]")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ListAvailableRelationCandidates(context.Background(), id).PageSize(pageSize).Filters(filters).Query(query).Type_(type_).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListAvailableRelationCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableRelationCandidates`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ListAvailableRelationCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableRelationCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | Maximum number of candidates to list (default 10) | 
 **filters** | **string** | JSON specifying filter conditions. Accepts the same filters as the [work packages](https://www.openproject.org/docs/api/endpoints/work-packages/) endpoint. | 
 **query** | **string** | Shortcut for filtering by ID or subject | 
 **type_** | **string** | Type of relation to find candidates for (default \&quot;relates\&quot;) | 
 **sortBy** | **string** | JSON specifying sort criteria. Accepts the same sort criteria as the [work packages](https://www.openproject.org/docs/api/endpoints/work-packages/) endpoint. | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]

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


## ListWatchers

> WatchersModel ListWatchers(ctx, id).Execute()

List watchers



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
	id := int64(1) // int64 | Work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ListWatchers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListWatchers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchers`: WatchersModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ListWatchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWatchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WatchersModel**](WatchersModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkPackageActivities

> WorkPackageActivitiesModel ListWorkPackageActivities(ctx, id).Execute()

List work package activities



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
	id := int64(1) // int64 | Work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ListWorkPackageActivities(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListWorkPackageActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkPackageActivities`: WorkPackageActivitiesModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ListWorkPackageActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkPackageActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkPackageActivitiesModel**](WorkPackageActivitiesModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkPackageFileLinks

> FileLinkCollectionReadModel ListWorkPackageFileLinks(ctx, id).Filters(filters).Execute()

Gets all file links of a work package



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
	id := int64(1337) // int64 | Work package id
	filters := "[{"storage":{"operator":"=","values":["42"]}}]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. The following filters are supported:  - storage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ListWorkPackageFileLinks(context.Background(), id).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListWorkPackageFileLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkPackageFileLinks`: FileLinkCollectionReadModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ListWorkPackageFileLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkPackageFileLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. The following filters are supported:  - storage | 

### Return type

[**FileLinkCollectionReadModel**](FileLinkCollectionReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkPackageRelations

> ListWorkPackageRelations(ctx, id).Execute()

List relations



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
	id := int64(1) // int64 | Work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.ListWorkPackageRelations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListWorkPackageRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkPackageRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkPackageSchemas

> map[string]interface{} ListWorkPackageSchemas(ctx).Filters(filters).Execute()

List Work Package Schemas



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
	filters := "[{ "id": { "operator": "=", "values": ["12-1", "14-2"] } }]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + id: The schema's id  Schema id has the form `project_id-work_package_type_id`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ListWorkPackageSchemas(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListWorkPackageSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkPackageSchemas`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ListWorkPackageSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkPackageSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + id: The schema&#39;s id  Schema id has the form &#x60;project_id-work_package_type_id&#x60;. | 

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


## ListWorkPackages

> WorkPackagesModel ListWorkPackages(ctx).Offset(offset).PageSize(pageSize).Filters(filters).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Select_(select_).Timestamps(timestamps).Execute()

List work packages



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
	filters := "[{ "type_id": { "operator": "=", "values": ["1", "2"] }}]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. If no filter is to be applied, the client should send an empty array (`[]`), otherwise a default filter is applied. A Currently supported filters are (there are additional filters added by modules):  - assigned_to - assignee_or_group - attachment_base - attachment_content - attachment_file_name - author - blocked - blocks - category - comment - created_at - custom_field - dates_interval - description - done_ratio - due_date - duplicated - duplicates - duration - estimated_hours - file_link_origin_id - follows - group - id - includes - linkable_to_storage_id - linkable_to_storage_url - manual_sort - milestone - only_subproject - parent - partof - precedes - principal_base - priority - project - relatable - relates - required - requires - responsible - role - search - start_date - status - storage_id - storage_url - subject - subject_or_id - subproject - type - typeahead - updated_at - version - watcher - work_package (optional) (default to "[{ \"status_id\": { \"operator\": \"o\", \"values\": null }}]")
	sortBy := "[["status", "asc"]]" // string | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. (optional) (default to "[[\"id\", \"asc\"]]")
	groupBy := "status" // string | The column to group by. (optional)
	showSums := true // bool | Indicates whether properties should be summed up if they support it. (optional) (default to false)
	select_ := "total,elements/subject,elements/id,self" // string | Comma separated list of properties to include. (optional)
	timestamps := "2022-01-01T00:00:00Z,PT0S" // string | In order to perform a [baseline comparison](/docs/api/baseline-comparisons), you may provide one or several timestamps in ISO-8601 format as comma-separated list. The timestamps may be absolute or relative, such as ISO8601 dates, ISO8601 durations and the following relative date keywords: \"oneDayAgo@HH:MM+HH:MM\", \"lastWorkingDay@HH:MM+HH:MM\", \"oneWeekAgo@HH:MM+HH:MM\", \"oneMonthAgo@HH:MM+HH:MM\". The first \"HH:MM\" part represents the zero paded hours and minutes. The last \"+HH:MM\" part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\"oneDayAgo@01:00+01:00\", \"oneDayAgo@01:00-01:00\".  Usually, the first timestamp is the baseline date, the last timestamp is the current date. Values older than 1 day are accepted only with valid Enterprise Token available. (optional) (default to "PT0S")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ListWorkPackages(context.Background()).Offset(offset).PageSize(pageSize).Filters(filters).SortBy(sortBy).GroupBy(groupBy).ShowSums(showSums).Select_(select_).Timestamps(timestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ListWorkPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkPackages`: WorkPackagesModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ListWorkPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Page number inside the requested collection. | [default to 1]
 **pageSize** | **int64** | Number of elements to display per page. | 
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. If no filter is to be applied, the client should send an empty array (&#x60;[]&#x60;), otherwise a default filter is applied. A Currently supported filters are (there are additional filters added by modules):  - assigned_to - assignee_or_group - attachment_base - attachment_content - attachment_file_name - author - blocked - blocks - category - comment - created_at - custom_field - dates_interval - description - done_ratio - due_date - duplicated - duplicates - duration - estimated_hours - file_link_origin_id - follows - group - id - includes - linkable_to_storage_id - linkable_to_storage_url - manual_sort - milestone - only_subproject - parent - partof - precedes - principal_base - priority - project - relatable - relates - required - requires - responsible - role - search - start_date - status - storage_id - storage_url - subject - subject_or_id - subproject - type - typeahead - updated_at - version - watcher - work_package | [default to &quot;[{ \&quot;status_id\&quot;: { \&quot;operator\&quot;: \&quot;o\&quot;, \&quot;values\&quot;: null }}]&quot;]
 **sortBy** | **string** | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]
 **groupBy** | **string** | The column to group by. | 
 **showSums** | **bool** | Indicates whether properties should be summed up if they support it. | [default to false]
 **select_** | **string** | Comma separated list of properties to include. | 
 **timestamps** | **string** | In order to perform a [baseline comparison](/docs/api/baseline-comparisons), you may provide one or several timestamps in ISO-8601 format as comma-separated list. The timestamps may be absolute or relative, such as ISO8601 dates, ISO8601 durations and the following relative date keywords: \&quot;oneDayAgo@HH:MM+HH:MM\&quot;, \&quot;lastWorkingDay@HH:MM+HH:MM\&quot;, \&quot;oneWeekAgo@HH:MM+HH:MM\&quot;, \&quot;oneMonthAgo@HH:MM+HH:MM\&quot;. The first \&quot;HH:MM\&quot; part represents the zero paded hours and minutes. The last \&quot;+HH:MM\&quot; part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\&quot;oneDayAgo@01:00+01:00\&quot;, \&quot;oneDayAgo@01:00-01:00\&quot;.  Usually, the first timestamp is the baseline date, the last timestamp is the current date. Values older than 1 day are accepted only with valid Enterprise Token available. | [default to &quot;PT0S&quot;]

### Return type

[**WorkPackagesModel**](WorkPackagesModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectAvailableAssignees

> AvailableAssigneesModel ProjectAvailableAssignees(ctx, id).Execute()

Project Available assignees



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
	resp, r, err := apiClient.WorkPackagesAPI.ProjectAvailableAssignees(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ProjectAvailableAssignees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectAvailableAssignees`: AvailableAssigneesModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ProjectAvailableAssignees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectAvailableAssigneesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailableAssigneesModel**](AvailableAssigneesModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWatcher

> RemoveWatcher(ctx, id, userId).Execute()

Remove watcher



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
	id := int64(1) // int64 | Work package id
	userId := int64(1) // int64 | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.RemoveWatcher(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.RemoveWatcher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 
**userId** | **int64** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWatcherRequest struct via the builder pattern


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


## Revisions

> map[string]interface{} Revisions(ctx, id).Execute()

Revisions



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
	id := int64(1) // int64 | Work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.Revisions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.Revisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Revisions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.Revisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevisionsRequest struct via the builder pattern


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


## UpdateWorkPackage

> WorkPackagePatchModel UpdateWorkPackage(ctx, id).Notify(notify).WorkPackageModel(workPackageModel).Execute()

Update a Work Package



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
	id := int64(42) // int64 | Work package id
	notify := false // bool | Indicates whether change notifications should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. (optional) (default to true)
	workPackageModel := *openapiclient.NewWorkPackageModel() // WorkPackageModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.UpdateWorkPackage(context.Background(), id).Notify(notify).WorkPackageModel(workPackageModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.UpdateWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkPackage`: WorkPackagePatchModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.UpdateWorkPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notify** | **bool** | Indicates whether change notifications should be sent. Note that this controls notifications for all users interested in changes to the work package (e.g. watchers, author and assignee), not just the current user. | [default to true]
 **workPackageModel** | [**WorkPackageModel**](WorkPackageModel.md) |  | 

### Return type

[**WorkPackagePatchModel**](WorkPackagePatchModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewWorkPackage

> WorkPackageModel ViewWorkPackage(ctx, id).Timestamps(timestamps).Execute()

View Work Package



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
	id := int64(1) // int64 | Work package id
	timestamps := "2022-01-01T00:00:00Z,PT0S" // string | In order to perform a [baseline comparison](/docs/api/baseline-comparisons) of the work-package attributes, you may provide one or several timestamps in ISO-8601 format as comma-separated list. The timestamps may be absolute or relative, such as ISO8601 dates, ISO8601 durations and the following relative date keywords: \"oneDayAgo@HH:MM+HH:MM\", \"lastWorkingDay@HH:MM+HH:MM\", \"oneWeekAgo@HH:MM+HH:MM\", \"oneMonthAgo@HH:MM+HH:MM\". The first \"HH:MM\" part represents the zero paded hours and minutes. The last \"+HH:MM\" part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\"oneDayAgo@01:00+01:00\", \"oneDayAgo@01:00-01:00\".  Usually, the first timestamp is the baseline date, the last timestamp is the current date. Values older than 1 day are accepted only with valid Enterprise Token available. (optional) (default to "PT0S")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.ViewWorkPackage(context.Background(), id).Timestamps(timestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ViewWorkPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewWorkPackage`: WorkPackageModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.ViewWorkPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewWorkPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamps** | **string** | In order to perform a [baseline comparison](/docs/api/baseline-comparisons) of the work-package attributes, you may provide one or several timestamps in ISO-8601 format as comma-separated list. The timestamps may be absolute or relative, such as ISO8601 dates, ISO8601 durations and the following relative date keywords: \&quot;oneDayAgo@HH:MM+HH:MM\&quot;, \&quot;lastWorkingDay@HH:MM+HH:MM\&quot;, \&quot;oneWeekAgo@HH:MM+HH:MM\&quot;, \&quot;oneMonthAgo@HH:MM+HH:MM\&quot;. The first \&quot;HH:MM\&quot; part represents the zero paded hours and minutes. The last \&quot;+HH:MM\&quot; part represents the timezone offset from UTC associated with the time, the offset can be positive or negative e.g.\&quot;oneDayAgo@01:00+01:00\&quot;, \&quot;oneDayAgo@01:00-01:00\&quot;.  Usually, the first timestamp is the baseline date, the last timestamp is the current date. Values older than 1 day are accepted only with valid Enterprise Token available. | [default to &quot;PT0S&quot;]

### Return type

[**WorkPackageModel**](WorkPackageModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewWorkPackageSchema

> ViewWorkPackageSchema(ctx, identifier).Execute()

View Work Package Schema



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
	identifier := "12-13" // string | Identifier of the schema

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.ViewWorkPackageSchema(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.ViewWorkPackageSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Identifier of the schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewWorkPackageSchemaRequest struct via the builder pattern


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


## WorkPackageAvailableAssignees

> AvailableAssigneesModel WorkPackageAvailableAssignees(ctx, id).Execute()

Work Package Available assignees



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
	id := int64(1) // int64 | Work package id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkPackagesAPI.WorkPackageAvailableAssignees(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.WorkPackageAvailableAssignees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkPackageAvailableAssignees`: AvailableAssigneesModel
	fmt.Fprintf(os.Stdout, "Response from `WorkPackagesAPI.WorkPackageAvailableAssignees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Work package id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkPackageAvailableAssigneesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailableAssigneesModel**](AvailableAssigneesModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkPackageCreateForm

> WorkPackageCreateForm(ctx).Execute()

Work Package Create Form



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
	r, err := apiClient.WorkPackagesAPI.WorkPackageCreateForm(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.WorkPackageCreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkPackageCreateFormRequest struct via the builder pattern


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


## WorkPackageCreateFormForProject

> WorkPackageCreateFormForProject(ctx, id).Execute()

Work Package Create Form For Project



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
	id := int64(1) // int64 | ID of the project in which the work package will be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.WorkPackageCreateFormForProject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.WorkPackageCreateFormForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the project in which the work package will be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkPackageCreateFormForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## WorkPackageEditForm

> WorkPackageEditForm(ctx, id).WorkPackageModel(workPackageModel).Execute()

Work Package Edit Form



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
	id := int64(1) // int64 | ID of the work package being modified
	workPackageModel := *openapiclient.NewWorkPackageModel() // WorkPackageModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkPackagesAPI.WorkPackageEditForm(context.Background(), id).WorkPackageModel(workPackageModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkPackagesAPI.WorkPackageEditForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the work package being modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkPackageEditFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workPackageModel** | [**WorkPackageModel**](WorkPackageModel.md) |  | 

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

