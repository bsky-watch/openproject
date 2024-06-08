# WeekDayModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Day** | **int64** | The week day from 1 to 7. 1 is Monday. 7 is Sunday. | [readonly] 
**Name** | **string** | The week day name. | 
**Working** | **bool** | &#x60;true&#x60; for a working week day, &#x60;false&#x60; otherwise. | 
**Links** | Pointer to [**WeekDaySelfLinkModel**](WeekDaySelfLinkModel.md) |  | [optional] 

## Methods

### NewWeekDayModel

`func NewWeekDayModel(type_ string, day int64, name string, working bool, ) *WeekDayModel`

NewWeekDayModel instantiates a new WeekDayModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeekDayModelWithDefaults

`func NewWeekDayModelWithDefaults() *WeekDayModel`

NewWeekDayModelWithDefaults instantiates a new WeekDayModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WeekDayModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WeekDayModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WeekDayModel) SetType(v string)`

SetType sets Type field to given value.


### GetDay

`func (o *WeekDayModel) GetDay() int64`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *WeekDayModel) GetDayOk() (*int64, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *WeekDayModel) SetDay(v int64)`

SetDay sets Day field to given value.


### GetName

`func (o *WeekDayModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WeekDayModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WeekDayModel) SetName(v string)`

SetName sets Name field to given value.


### GetWorking

`func (o *WeekDayModel) GetWorking() bool`

GetWorking returns the Working field if non-nil, zero value otherwise.

### GetWorkingOk

`func (o *WeekDayModel) GetWorkingOk() (*bool, bool)`

GetWorkingOk returns a tuple with the Working field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorking

`func (o *WeekDayModel) SetWorking(v bool)`

SetWorking sets Working field to given value.


### GetLinks

`func (o *WeekDayModel) GetLinks() WeekDaySelfLinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WeekDayModel) GetLinksOk() (*WeekDaySelfLinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WeekDayModel) SetLinks(v WeekDaySelfLinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WeekDayModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


