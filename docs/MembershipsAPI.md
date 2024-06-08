# \MembershipsAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMembership**](MembershipsAPI.md#CreateMembership) | **Post** /api/v3/memberships | Create a membership
[**DeleteMembership**](MembershipsAPI.md#DeleteMembership) | **Delete** /api/v3/memberships/{id} | Delete membership
[**FormCreateMembership**](MembershipsAPI.md#FormCreateMembership) | **Post** /api/v3/memberships/form | Form create membership
[**FormUpdateMembership**](MembershipsAPI.md#FormUpdateMembership) | **Post** /api/v3/memberships/{id}/form | Form update membership
[**GetMembership**](MembershipsAPI.md#GetMembership) | **Get** /api/v3/memberships/{id} | Get a membership
[**GetMembershipSchema**](MembershipsAPI.md#GetMembershipSchema) | **Get** /api/v3/memberships/schema | Schema membership
[**GetMembershipsAvailableProjects**](MembershipsAPI.md#GetMembershipsAvailableProjects) | **Get** /api/v3/memberships/available_projects | Available projects for memberships
[**ListMemberships**](MembershipsAPI.md#ListMemberships) | **Get** /api/v3/memberships | List memberships
[**UpdateMembership**](MembershipsAPI.md#UpdateMembership) | **Patch** /api/v3/memberships/{id} | Update membership



## CreateMembership

> MembershipReadModel CreateMembership(ctx).MembershipWriteModel(membershipWriteModel).Execute()

Create a membership



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
	membershipWriteModel := *openapiclient.NewMembershipWriteModel(*openapiclient.NewMembershipWriteModelLinks()) // MembershipWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.CreateMembership(context.Background()).MembershipWriteModel(membershipWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.CreateMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMembership`: MembershipReadModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.CreateMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipWriteModel** | [**MembershipWriteModel**](MembershipWriteModel.md) |  | 

### Return type

[**MembershipReadModel**](MembershipReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMembership

> DeleteMembership(ctx, id).Execute()

Delete membership



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
	id := int64(1) // int64 | Membership id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.DeleteMembership(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.DeleteMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMembershipRequest struct via the builder pattern


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


## FormCreateMembership

> MembershipFormModel FormCreateMembership(ctx).MembershipWriteModel(membershipWriteModel).Execute()

Form create membership



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
	membershipWriteModel := *openapiclient.NewMembershipWriteModel(*openapiclient.NewMembershipWriteModelLinks()) // MembershipWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.FormCreateMembership(context.Background()).MembershipWriteModel(membershipWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.FormCreateMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormCreateMembership`: MembershipFormModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.FormCreateMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFormCreateMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipWriteModel** | [**MembershipWriteModel**](MembershipWriteModel.md) |  | 

### Return type

[**MembershipFormModel**](MembershipFormModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FormUpdateMembership

> MembershipReadModel FormUpdateMembership(ctx, id).MembershipWriteModel(membershipWriteModel).Execute()

Form update membership



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
	id := int64(1) // int64 | Membership id
	membershipWriteModel := *openapiclient.NewMembershipWriteModel(*openapiclient.NewMembershipWriteModelLinks()) // MembershipWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.FormUpdateMembership(context.Background(), id).MembershipWriteModel(membershipWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.FormUpdateMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormUpdateMembership`: MembershipReadModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.FormUpdateMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFormUpdateMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipWriteModel** | [**MembershipWriteModel**](MembershipWriteModel.md) |  | 

### Return type

[**MembershipReadModel**](MembershipReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembership

> MembershipReadModel GetMembership(ctx, id).Execute()

Get a membership



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
	id := int64(1) // int64 | Membership id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetMembership(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembership`: MembershipReadModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MembershipReadModel**](MembershipReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembershipSchema

> MembershipSchemaModel GetMembershipSchema(ctx).Execute()

Schema membership



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
	resp, r, err := apiClient.MembershipsAPI.GetMembershipSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetMembershipSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembershipSchema`: MembershipSchemaModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetMembershipSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembershipSchemaRequest struct via the builder pattern


### Return type

[**MembershipSchemaModel**](MembershipSchemaModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembershipsAvailableProjects

> ProjectCollectionModel GetMembershipsAvailableProjects(ctx).Execute()

Available projects for memberships



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
	resp, r, err := apiClient.MembershipsAPI.GetMembershipsAvailableProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetMembershipsAvailableProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembershipsAvailableProjects`: ProjectCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetMembershipsAvailableProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembershipsAvailableProjectsRequest struct via the builder pattern


### Return type

[**ProjectCollectionModel**](ProjectCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMemberships

> MembershipCollectionModel ListMemberships(ctx).Filters(filters).SortBy(sortBy).Execute()

List memberships



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
	filters := "[{ "name": { "operator": "=", "values": ["A User"] }" }]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + any_name_attribute: filters memberships based on the name of the principal. All possible name variants   (and also email and login) are searched. + blocked: reduces the result set to all memberships that are temporarily blocked or that are not blocked   temporarily. + group: filters memberships based on the name of a group. The group however is not the principal used for   filtering. Rather, the memberships of the group are used as the filter values. + name: filters memberships based on the name of the principal. Note that only the name is used which depends   on a setting in the OpenProject instance. + principal: filters memberships based on the id of the principal. + project: filters memberships based on the id of the project. + role: filters memberships based on the id of any role assigned to the membership. + status: filters memberships based on the status of the principal. + created_at: filters memberships based on the time the membership was created. + updated_at: filters memberships based on the time the membership was updated last. (optional)
	sortBy := "[["id", "asc"]]" // string | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported sorts are:  + id: Sort by primary key + name: Sort by the name of the principal. Note that this depends on the setting for how the name is to be   displayed at least for users. + email: Sort by the email address of the principal. Groups and principal users, which do not have an email,   are sorted last. + status: Sort by the status of the principal. Groups and principal users, which do not have a status, are   sorted together with the active users. + created_at: Sort by membership creation datetime + updated_at: Sort by the time the membership was updated last (optional) (default to "[[\"id\", \"asc\"]]")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.ListMemberships(context.Background()).Filters(filters).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.ListMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMemberships`: MembershipCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.ListMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + any_name_attribute: filters memberships based on the name of the principal. All possible name variants   (and also email and login) are searched. + blocked: reduces the result set to all memberships that are temporarily blocked or that are not blocked   temporarily. + group: filters memberships based on the name of a group. The group however is not the principal used for   filtering. Rather, the memberships of the group are used as the filter values. + name: filters memberships based on the name of the principal. Note that only the name is used which depends   on a setting in the OpenProject instance. + principal: filters memberships based on the id of the principal. + project: filters memberships based on the id of the project. + role: filters memberships based on the id of any role assigned to the membership. + status: filters memberships based on the status of the principal. + created_at: filters memberships based on the time the membership was created. + updated_at: filters memberships based on the time the membership was updated last. | 
 **sortBy** | **string** | JSON specifying sort criteria. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported sorts are:  + id: Sort by primary key + name: Sort by the name of the principal. Note that this depends on the setting for how the name is to be   displayed at least for users. + email: Sort by the email address of the principal. Groups and principal users, which do not have an email,   are sorted last. + status: Sort by the status of the principal. Groups and principal users, which do not have a status, are   sorted together with the active users. + created_at: Sort by membership creation datetime + updated_at: Sort by the time the membership was updated last | [default to &quot;[[\&quot;id\&quot;, \&quot;asc\&quot;]]&quot;]

### Return type

[**MembershipCollectionModel**](MembershipCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMembership

> MembershipReadModel UpdateMembership(ctx, id).MembershipWriteModel(membershipWriteModel).Execute()

Update membership



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
	id := int64(1) // int64 | Membership id
	membershipWriteModel := *openapiclient.NewMembershipWriteModel(*openapiclient.NewMembershipWriteModelLinks()) // MembershipWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.UpdateMembership(context.Background(), id).MembershipWriteModel(membershipWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.UpdateMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMembership`: MembershipReadModel
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.UpdateMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Membership id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipWriteModel** | [**MembershipWriteModel**](MembershipWriteModel.md) |  | 

### Return type

[**MembershipReadModel**](MembershipReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

