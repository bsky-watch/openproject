# DayModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**Link**](Link.md) |  | 
**NonWorkingReasons** | Pointer to [**[]Link**](Link.md) | A list of resources describing why this day is a non-working day. Linked resources can be &#x60;NonWorkingDay&#x60; and &#x60;WeekDay&#x60; resources. This property is absent for working days. | [optional] 
**WeekDay** | Pointer to [**DayModelLinksWeekDay**](DayModelLinksWeekDay.md) |  | [optional] 

## Methods

### NewDayModelLinks

`func NewDayModelLinks(self Link, ) *DayModelLinks`

NewDayModelLinks instantiates a new DayModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayModelLinksWithDefaults

`func NewDayModelLinksWithDefaults() *DayModelLinks`

NewDayModelLinksWithDefaults instantiates a new DayModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *DayModelLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *DayModelLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *DayModelLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.


### GetNonWorkingReasons

`func (o *DayModelLinks) GetNonWorkingReasons() []Link`

GetNonWorkingReasons returns the NonWorkingReasons field if non-nil, zero value otherwise.

### GetNonWorkingReasonsOk

`func (o *DayModelLinks) GetNonWorkingReasonsOk() (*[]Link, bool)`

GetNonWorkingReasonsOk returns a tuple with the NonWorkingReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonWorkingReasons

`func (o *DayModelLinks) SetNonWorkingReasons(v []Link)`

SetNonWorkingReasons sets NonWorkingReasons field to given value.

### HasNonWorkingReasons

`func (o *DayModelLinks) HasNonWorkingReasons() bool`

HasNonWorkingReasons returns a boolean if a field has been set.

### GetWeekDay

`func (o *DayModelLinks) GetWeekDay() DayModelLinksWeekDay`

GetWeekDay returns the WeekDay field if non-nil, zero value otherwise.

### GetWeekDayOk

`func (o *DayModelLinks) GetWeekDayOk() (*DayModelLinksWeekDay, bool)`

GetWeekDayOk returns a tuple with the WeekDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDay

`func (o *DayModelLinks) SetWeekDay(v DayModelLinksWeekDay)`

SetWeekDay sets WeekDay field to given value.

### HasWeekDay

`func (o *DayModelLinks) HasWeekDay() bool`

HasWeekDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


