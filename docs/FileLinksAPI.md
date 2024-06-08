# \FileLinksAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorage**](FileLinksAPI.md#CreateStorage) | **Post** /api/v3/storages | Creates a storage.
[**CreateStorageOauthCredentials**](FileLinksAPI.md#CreateStorageOauthCredentials) | **Post** /api/v3/storages/{id}/oauth_client_credentials | Creates an oauth client credentials object for a storage.
[**DeleteFileLink**](FileLinksAPI.md#DeleteFileLink) | **Delete** /api/v3/file_links/{id} | Removes a file link.
[**DeleteStorage**](FileLinksAPI.md#DeleteStorage) | **Delete** /api/v3/storages/{id} | Delete a storage
[**DownloadFileLink**](FileLinksAPI.md#DownloadFileLink) | **Get** /api/v3/file_links/{id}/download | Creates a download uri of the linked file.
[**GetProjectStorage**](FileLinksAPI.md#GetProjectStorage) | **Get** /api/v3/project_storages/{id} | Gets a project storage
[**GetStorage**](FileLinksAPI.md#GetStorage) | **Get** /api/v3/storages/{id} | Get a storage
[**GetStorageFiles**](FileLinksAPI.md#GetStorageFiles) | **Get** /api/v3/storages/{id}/files | Gets files of a storage.
[**ListProjectStorages**](FileLinksAPI.md#ListProjectStorages) | **Get** /api/v3/project_storages | Gets a list of project storages
[**ListStorages**](FileLinksAPI.md#ListStorages) | **Get** /api/v3/storages | Get Storages
[**OpenFileLink**](FileLinksAPI.md#OpenFileLink) | **Get** /api/v3/file_links/{id}/open | Creates an opening uri of the linked file.
[**OpenProjectStorage**](FileLinksAPI.md#OpenProjectStorage) | **Get** /api/v3/project_storages/{id}/open | Open the project storage
[**OpenStorage**](FileLinksAPI.md#OpenStorage) | **Get** /api/v3/storages/{id}/open | Open the storage
[**PrepareStorageFileUpload**](FileLinksAPI.md#PrepareStorageFileUpload) | **Post** /api/v3/storages/{id}/files/prepare_upload | Preparation of a direct upload of a file to the given storage.
[**UpdateStorage**](FileLinksAPI.md#UpdateStorage) | **Patch** /api/v3/storages/{id} | Update a storage
[**ViewFileLink**](FileLinksAPI.md#ViewFileLink) | **Get** /api/v3/file_links/{id} | Gets a file link.



## CreateStorage

> StorageReadModel CreateStorage(ctx).StorageWriteModel(storageWriteModel).Execute()

Creates a storage.



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
	storageWriteModel := *openapiclient.NewStorageWriteModel() // StorageWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.CreateStorage(context.Background()).StorageWriteModel(storageWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.CreateStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorage`: StorageReadModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.CreateStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageWriteModel** | [**StorageWriteModel**](StorageWriteModel.md) |  | 

### Return type

[**StorageReadModel**](StorageReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageOauthCredentials

> StorageReadModel CreateStorageOauthCredentials(ctx, id).OAuthClientCredentialsWriteModel(oAuthClientCredentialsWriteModel).Execute()

Creates an oauth client credentials object for a storage.



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
	id := int64(1337) // int64 | Storage id
	oAuthClientCredentialsWriteModel := *openapiclient.NewOAuthClientCredentialsWriteModel("ClientId_example", "ClientSecret_example") // OAuthClientCredentialsWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.CreateStorageOauthCredentials(context.Background(), id).OAuthClientCredentialsWriteModel(oAuthClientCredentialsWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.CreateStorageOauthCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorageOauthCredentials`: StorageReadModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.CreateStorageOauthCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageOauthCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuthClientCredentialsWriteModel** | [**OAuthClientCredentialsWriteModel**](OAuthClientCredentialsWriteModel.md) |  | 

### Return type

[**StorageReadModel**](StorageReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileLink

> DeleteFileLink(ctx, id).Execute()

Removes a file link.



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
	id := int64(42) // int64 | File link id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileLinksAPI.DeleteFileLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.DeleteFileLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | File link id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileLinkRequest struct via the builder pattern


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


## DeleteStorage

> DeleteStorage(ctx, id).Execute()

Delete a storage



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
	id := int64(1337) // int64 | Storage id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileLinksAPI.DeleteStorage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.DeleteStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageRequest struct via the builder pattern


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


## DownloadFileLink

> DownloadFileLink(ctx, id).Execute()

Creates a download uri of the linked file.



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
	id := int64(42) // int64 | File link id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileLinksAPI.DownloadFileLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.DownloadFileLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | File link id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileLinkRequest struct via the builder pattern


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


## GetProjectStorage

> ProjectStorageModel GetProjectStorage(ctx, id).Execute()

Gets a project storage



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
	id := int64(1337) // int64 | Project storage id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.GetProjectStorage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.GetProjectStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectStorage`: ProjectStorageModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.GetProjectStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectStorageModel**](ProjectStorageModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorage

> StorageReadModel GetStorage(ctx, id).Execute()

Get a storage



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
	id := int64(1337) // int64 | Storage id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.GetStorage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.GetStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorage`: StorageReadModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.GetStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageReadModel**](StorageReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageFiles

> StorageFilesModel GetStorageFiles(ctx, id).Parent(parent).Execute()

Gets files of a storage.



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
	id := int64(1337) // int64 | Storage id
	parent := "/my/data" // string | Parent file identification (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.GetStorageFiles(context.Background(), id).Parent(parent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.GetStorageFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageFiles`: StorageFilesModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.GetStorageFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parent** | **string** | Parent file identification | 

### Return type

[**StorageFilesModel**](StorageFilesModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectStorages

> ProjectStorageCollectionModel ListProjectStorages(ctx).Filters(filters).Execute()

Gets a list of project storages



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
	filters := "[{ "project_id": { "operator": "=", "values": ["42"] }}, { "storage_id": { "operator": "=", "values": ["1337"] }}]" // string | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  - project_id - storage_id - storage_url (optional) (default to "[]")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.ListProjectStorages(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.ListProjectStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectStorages`: ProjectStorageCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.ListProjectStorages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProjectStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions. Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  - project_id - storage_id - storage_url | [default to &quot;[]&quot;]

### Return type

[**ProjectStorageCollectionModel**](ProjectStorageCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorages

> StorageCollectionModel ListStorages(ctx).Execute()

Get Storages



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
	resp, r, err := apiClient.FileLinksAPI.ListStorages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.ListStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStorages`: StorageCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.ListStorages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStoragesRequest struct via the builder pattern


### Return type

[**StorageCollectionModel**](StorageCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenFileLink

> OpenFileLink(ctx, id).Location(location).Execute()

Creates an opening uri of the linked file.



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
	id := int64(42) // int64 | File link id
	location := true // bool | Boolean flag indicating, if the file should be opened directly or rather the directory location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileLinksAPI.OpenFileLink(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.OpenFileLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | File link id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenFileLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | **bool** | Boolean flag indicating, if the file should be opened directly or rather the directory location. | 

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


## OpenProjectStorage

> OpenProjectStorage(ctx, id).Execute()

Open the project storage



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
	id := int64(1337) // int64 | Project storage id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileLinksAPI.OpenProjectStorage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.OpenProjectStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Project storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenProjectStorageRequest struct via the builder pattern


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


## OpenStorage

> OpenStorage(ctx, id).Execute()

Open the storage



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
	id := int64(1337) // int64 | Storage id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileLinksAPI.OpenStorage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.OpenStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenStorageRequest struct via the builder pattern


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


## PrepareStorageFileUpload

> StorageFileUploadLinkModel PrepareStorageFileUpload(ctx, id).StorageFileUploadPreparationModel(storageFileUploadPreparationModel).Execute()

Preparation of a direct upload of a file to the given storage.



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
	id := int64(1337) // int64 | Storage id
	storageFileUploadPreparationModel := *openapiclient.NewStorageFileUploadPreparationModel(int64(123), "FileName_example", "Parent_example") // StorageFileUploadPreparationModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.PrepareStorageFileUpload(context.Background(), id).StorageFileUploadPreparationModel(storageFileUploadPreparationModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.PrepareStorageFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrepareStorageFileUpload`: StorageFileUploadLinkModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.PrepareStorageFileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrepareStorageFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageFileUploadPreparationModel** | [**StorageFileUploadPreparationModel**](StorageFileUploadPreparationModel.md) |  | 

### Return type

[**StorageFileUploadLinkModel**](StorageFileUploadLinkModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorage

> StorageReadModel UpdateStorage(ctx, id).StorageWriteModel(storageWriteModel).Execute()

Update a storage



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
	id := int64(1337) // int64 | Storage id
	storageWriteModel := *openapiclient.NewStorageWriteModel() // StorageWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.UpdateStorage(context.Background(), id).StorageWriteModel(storageWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.UpdateStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStorage`: StorageReadModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.UpdateStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Storage id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageWriteModel** | [**StorageWriteModel**](StorageWriteModel.md) |  | 

### Return type

[**StorageReadModel**](StorageReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewFileLink

> FileLinkReadModel ViewFileLink(ctx, id).Execute()

Gets a file link.



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
	id := int64(42) // int64 | File link id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileLinksAPI.ViewFileLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileLinksAPI.ViewFileLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewFileLink`: FileLinkReadModel
	fmt.Fprintf(os.Stdout, "Response from `FileLinksAPI.ViewFileLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | File link id | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewFileLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileLinkReadModel**](FileLinkReadModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

