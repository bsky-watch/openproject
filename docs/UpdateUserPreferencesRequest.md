# UpdateUserPreferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoHidePopups** | Pointer to **bool** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUserPreferencesRequest

`func NewUpdateUserPreferencesRequest() *UpdateUserPreferencesRequest`

NewUpdateUserPreferencesRequest instantiates a new UpdateUserPreferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPreferencesRequestWithDefaults

`func NewUpdateUserPreferencesRequestWithDefaults() *UpdateUserPreferencesRequest`

NewUpdateUserPreferencesRequestWithDefaults instantiates a new UpdateUserPreferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoHidePopups

`func (o *UpdateUserPreferencesRequest) GetAutoHidePopups() bool`

GetAutoHidePopups returns the AutoHidePopups field if non-nil, zero value otherwise.

### GetAutoHidePopupsOk

`func (o *UpdateUserPreferencesRequest) GetAutoHidePopupsOk() (*bool, bool)`

GetAutoHidePopupsOk returns a tuple with the AutoHidePopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoHidePopups

`func (o *UpdateUserPreferencesRequest) SetAutoHidePopups(v bool)`

SetAutoHidePopups sets AutoHidePopups field to given value.

### HasAutoHidePopups

`func (o *UpdateUserPreferencesRequest) HasAutoHidePopups() bool`

HasAutoHidePopups returns a boolean if a field has been set.

### GetTimeZone

`func (o *UpdateUserPreferencesRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UpdateUserPreferencesRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UpdateUserPreferencesRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UpdateUserPreferencesRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


