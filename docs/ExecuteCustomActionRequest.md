# ExecuteCustomActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ExecuteCustomActionRequestLinks**](ExecuteCustomActionRequestLinks.md) |  | [optional] 
**LockVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewExecuteCustomActionRequest

`func NewExecuteCustomActionRequest() *ExecuteCustomActionRequest`

NewExecuteCustomActionRequest instantiates a new ExecuteCustomActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteCustomActionRequestWithDefaults

`func NewExecuteCustomActionRequestWithDefaults() *ExecuteCustomActionRequest`

NewExecuteCustomActionRequestWithDefaults instantiates a new ExecuteCustomActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExecuteCustomActionRequest) GetLinks() ExecuteCustomActionRequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExecuteCustomActionRequest) GetLinksOk() (*ExecuteCustomActionRequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExecuteCustomActionRequest) SetLinks(v ExecuteCustomActionRequestLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExecuteCustomActionRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLockVersion

`func (o *ExecuteCustomActionRequest) GetLockVersion() string`

GetLockVersion returns the LockVersion field if non-nil, zero value otherwise.

### GetLockVersionOk

`func (o *ExecuteCustomActionRequest) GetLockVersionOk() (*string, bool)`

GetLockVersionOk returns a tuple with the LockVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockVersion

`func (o *ExecuteCustomActionRequest) SetLockVersion(v string)`

SetLockVersion sets LockVersion field to given value.

### HasLockVersion

`func (o *ExecuteCustomActionRequest) HasLockVersion() bool`

HasLockVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


