# \WorkScheduleAPI

All URIs are relative to *https://qa.openproject-edge.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNonWorkingDay**](WorkScheduleAPI.md#CreateNonWorkingDay) | **Post** /api/v3/days/non_working | Creates a non-working day (NOT IMPLEMENTED)
[**DeleteNonWorkingDay**](WorkScheduleAPI.md#DeleteNonWorkingDay) | **Delete** /api/v3/days/non_working/{date} | Removes a non-working day (NOT IMPLEMENTED)
[**ListDays**](WorkScheduleAPI.md#ListDays) | **Get** /api/v3/days | Lists days
[**ListNonWorkingDays**](WorkScheduleAPI.md#ListNonWorkingDays) | **Get** /api/v3/days/non_working | Lists all non working days
[**ListWeekDays**](WorkScheduleAPI.md#ListWeekDays) | **Get** /api/v3/days/week | Lists week days
[**UpdateNonWorkingDay**](WorkScheduleAPI.md#UpdateNonWorkingDay) | **Patch** /api/v3/days/non_working/{date} | Update a non-working day attributes (NOT IMPLEMENTED)
[**UpdateWeekDay**](WorkScheduleAPI.md#UpdateWeekDay) | **Patch** /api/v3/days/week/{day} | Update a week day attributes (NOT IMPLEMENTED)
[**UpdateWeekDays**](WorkScheduleAPI.md#UpdateWeekDays) | **Patch** /api/v3/days/week | Update week days (NOT IMPLEMENTED)
[**ViewDay**](WorkScheduleAPI.md#ViewDay) | **Get** /api/v3/days/{date} | View day
[**ViewNonWorkingDay**](WorkScheduleAPI.md#ViewNonWorkingDay) | **Get** /api/v3/days/non_working/{date} | View a non-working day
[**ViewWeekDay**](WorkScheduleAPI.md#ViewWeekDay) | **Get** /api/v3/days/week/{day} | View a week day



## CreateNonWorkingDay

> NonWorkingDayModel CreateNonWorkingDay(ctx).NonWorkingDayModel(nonWorkingDayModel).Execute()

Creates a non-working day (NOT IMPLEMENTED)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	nonWorkingDayModel := *openapiclient.NewNonWorkingDayModel("Type_example", time.Now(), "Name_example") // NonWorkingDayModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.CreateNonWorkingDay(context.Background()).NonWorkingDayModel(nonWorkingDayModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.CreateNonWorkingDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNonWorkingDay`: NonWorkingDayModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.CreateNonWorkingDay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNonWorkingDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonWorkingDayModel** | [**NonWorkingDayModel**](NonWorkingDayModel.md) |  | 

### Return type

[**NonWorkingDayModel**](NonWorkingDayModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonWorkingDay

> DeleteNonWorkingDay(ctx, date).Execute()

Removes a non-working day (NOT IMPLEMENTED)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	date := time.Now() // string | The date of the non-working day to view in ISO 8601 format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkScheduleAPI.DeleteNonWorkingDay(context.Background(), date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.DeleteNonWorkingDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | The date of the non-working day to view in ISO 8601 format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonWorkingDayRequest struct via the builder pattern


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


## ListDays

> DayCollectionModel ListDays(ctx).Filters(filters).Execute()

Lists days



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
	filters := "[{ "date": { "operator": "<>d", "values": ["2022-05-02","2022-05-26"] } }, { "working": { "operator": "=", "values": ["f"] } }]" // string | JSON specifying filter conditions.  Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + date: the inclusive date interval to scope days to look up. When   unspecified, default is from the beginning of current month to the end   of following month.    Example: `{ \"date\": { \"operator\": \"<>d\", \"values\": [\"2022-05-02\",\"2022-05-26\"] } }`   would return days between May 5 and May 26 2022, inclusive.  + working: when `true`, returns only the working days. When `false`,   returns only the non-working days (weekend days and non-working days).   When unspecified, returns both working and non-working days.    Example: `{ \"working\": { \"operator\": \"=\", \"values\": [\"t\"] } }`   would exclude non-working days from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.ListDays(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.ListDays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDays`: DayCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.ListDays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions.  Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + date: the inclusive date interval to scope days to look up. When   unspecified, default is from the beginning of current month to the end   of following month.    Example: &#x60;{ \&quot;date\&quot;: { \&quot;operator\&quot;: \&quot;&lt;&gt;d\&quot;, \&quot;values\&quot;: [\&quot;2022-05-02\&quot;,\&quot;2022-05-26\&quot;] } }&#x60;   would return days between May 5 and May 26 2022, inclusive.  + working: when &#x60;true&#x60;, returns only the working days. When &#x60;false&#x60;,   returns only the non-working days (weekend days and non-working days).   When unspecified, returns both working and non-working days.    Example: &#x60;{ \&quot;working\&quot;: { \&quot;operator\&quot;: \&quot;&#x3D;\&quot;, \&quot;values\&quot;: [\&quot;t\&quot;] } }&#x60;   would exclude non-working days from the response. | 

### Return type

[**DayCollectionModel**](DayCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNonWorkingDays

> NonWorkingDayCollectionModel ListNonWorkingDays(ctx).Filters(filters).Execute()

Lists all non working days



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
	filters := "[{ "date": { "operator": "<>d", "values": ["2022-05-02","2022-05-26"] } }]" // string | JSON specifying filter conditions.  Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + date: the inclusive date interval to scope days to look up. When   unspecified, default is from the beginning to the end of current year.    Example: `{ \"date\": { \"operator\": \"<>d\", \"values\": [\"2022-05-02\",\"2022-05-26\"] } }`   would return days between May 5 and May 26 2022, inclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.ListNonWorkingDays(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.ListNonWorkingDays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNonWorkingDays`: NonWorkingDayCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.ListNonWorkingDays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonWorkingDaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | JSON specifying filter conditions.  Accepts the same format as returned by the [queries](https://www.openproject.org/docs/api/endpoints/queries/) endpoint. Currently supported filters are:  + date: the inclusive date interval to scope days to look up. When   unspecified, default is from the beginning to the end of current year.    Example: &#x60;{ \&quot;date\&quot;: { \&quot;operator\&quot;: \&quot;&lt;&gt;d\&quot;, \&quot;values\&quot;: [\&quot;2022-05-02\&quot;,\&quot;2022-05-26\&quot;] } }&#x60;   would return days between May 5 and May 26 2022, inclusive. | 

### Return type

[**NonWorkingDayCollectionModel**](NonWorkingDayCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWeekDays

> WeekDayCollectionModel ListWeekDays(ctx).Execute()

Lists week days



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
	resp, r, err := apiClient.WorkScheduleAPI.ListWeekDays(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.ListWeekDays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWeekDays`: WeekDayCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.ListWeekDays`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWeekDaysRequest struct via the builder pattern


### Return type

[**WeekDayCollectionModel**](WeekDayCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNonWorkingDay

> NonWorkingDayModel UpdateNonWorkingDay(ctx, date).NonWorkingDayModel(nonWorkingDayModel).Execute()

Update a non-working day attributes (NOT IMPLEMENTED)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	date := time.Now() // string | The date of the non-working day to view in ISO 8601 format.
	nonWorkingDayModel := *openapiclient.NewNonWorkingDayModel("Type_example", time.Now(), "Name_example") // NonWorkingDayModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.UpdateNonWorkingDay(context.Background(), date).NonWorkingDayModel(nonWorkingDayModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.UpdateNonWorkingDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNonWorkingDay`: NonWorkingDayModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.UpdateNonWorkingDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | The date of the non-working day to view in ISO 8601 format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNonWorkingDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonWorkingDayModel** | [**NonWorkingDayModel**](NonWorkingDayModel.md) |  | 

### Return type

[**NonWorkingDayModel**](NonWorkingDayModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWeekDay

> WeekDayModel UpdateWeekDay(ctx, day).WeekDayWriteModel(weekDayWriteModel).Execute()

Update a week day attributes (NOT IMPLEMENTED)



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
	day := int64(56) // int64 | The week day from 1 to 7. 1 is Monday. 7 is Sunday.
	weekDayWriteModel := *openapiclient.NewWeekDayWriteModel("Type_example", false) // WeekDayWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.UpdateWeekDay(context.Background(), day).WeekDayWriteModel(weekDayWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.UpdateWeekDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWeekDay`: WeekDayModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.UpdateWeekDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **int64** | The week day from 1 to 7. 1 is Monday. 7 is Sunday. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWeekDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **weekDayWriteModel** | [**WeekDayWriteModel**](WeekDayWriteModel.md) |  | 

### Return type

[**WeekDayModel**](WeekDayModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWeekDays

> WeekDayCollectionModel UpdateWeekDays(ctx).WeekDayCollectionWriteModel(weekDayCollectionWriteModel).Execute()

Update week days (NOT IMPLEMENTED)



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
	weekDayCollectionWriteModel := *openapiclient.NewWeekDayCollectionWriteModel("Type_example", *openapiclient.NewWeekDayCollectionWriteModelEmbedded([]openapiclient.WeekDayCollectionWriteModelEmbeddedElementsInner{*openapiclient.NewWeekDayCollectionWriteModelEmbeddedElementsInner("Type_example", false, *openapiclient.NewWeekDaySelfLinkModel())})) // WeekDayCollectionWriteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.UpdateWeekDays(context.Background()).WeekDayCollectionWriteModel(weekDayCollectionWriteModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.UpdateWeekDays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWeekDays`: WeekDayCollectionModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.UpdateWeekDays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWeekDaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **weekDayCollectionWriteModel** | [**WeekDayCollectionWriteModel**](WeekDayCollectionWriteModel.md) |  | 

### Return type

[**WeekDayCollectionModel**](WeekDayCollectionModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewDay

> DayModel ViewDay(ctx, date).Execute()

View day



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	date := time.Now() // string | The date of the non-working day to view in ISO 8601 format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.ViewDay(context.Background(), date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.ViewDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewDay`: DayModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.ViewDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | The date of the non-working day to view in ISO 8601 format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DayModel**](DayModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewNonWorkingDay

> NonWorkingDayModel ViewNonWorkingDay(ctx, date).Execute()

View a non-working day



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ua-bsky-watch/openproject"
)

func main() {
	date := time.Now() // string | The date of the non-working day to view in ISO 8601 format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.ViewNonWorkingDay(context.Background(), date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.ViewNonWorkingDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewNonWorkingDay`: NonWorkingDayModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.ViewNonWorkingDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | The date of the non-working day to view in ISO 8601 format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewNonWorkingDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonWorkingDayModel**](NonWorkingDayModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewWeekDay

> WeekDayModel ViewWeekDay(ctx, day).Execute()

View a week day



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
	day := int64(56) // int64 | The week day from 1 to 7. 1 is Monday. 7 is Sunday.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkScheduleAPI.ViewWeekDay(context.Background(), day).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkScheduleAPI.ViewWeekDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewWeekDay`: WeekDayModel
	fmt.Fprintf(os.Stdout, "Response from `WorkScheduleAPI.ViewWeekDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **int64** | The week day from 1 to 7. 1 is Monday. 7 is Sunday. | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewWeekDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WeekDayModel**](WeekDayModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

