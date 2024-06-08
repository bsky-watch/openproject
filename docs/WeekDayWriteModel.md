# WeekDayWriteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Working** | **bool** | &#x60;true&#x60; for a working day. &#x60;false&#x60; for a weekend day. | 

## Methods

### NewWeekDayWriteModel

`func NewWeekDayWriteModel(type_ string, working bool, ) *WeekDayWriteModel`

NewWeekDayWriteModel instantiates a new WeekDayWriteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeekDayWriteModelWithDefaults

`func NewWeekDayWriteModelWithDefaults() *WeekDayWriteModel`

NewWeekDayWriteModelWithDefaults instantiates a new WeekDayWriteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WeekDayWriteModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WeekDayWriteModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WeekDayWriteModel) SetType(v string)`

SetType sets Type field to given value.


### GetWorking

`func (o *WeekDayWriteModel) GetWorking() bool`

GetWorking returns the Working field if non-nil, zero value otherwise.

### GetWorkingOk

`func (o *WeekDayWriteModel) GetWorkingOk() (*bool, bool)`

GetWorkingOk returns a tuple with the Working field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorking

`func (o *WeekDayWriteModel) SetWorking(v bool)`

SetWorking sets Working field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


