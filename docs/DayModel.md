# DayModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Date** | **string** | Date of the day. | 
**Name** | **string** | Descriptive name for the day. | 
**Working** | **bool** | &#x60;true&#x60; for a working day, &#x60;false&#x60; otherwise. | 
**Links** | Pointer to [**DayModelLinks**](DayModelLinks.md) |  | [optional] 

## Methods

### NewDayModel

`func NewDayModel(type_ string, date string, name string, working bool, ) *DayModel`

NewDayModel instantiates a new DayModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayModelWithDefaults

`func NewDayModelWithDefaults() *DayModel`

NewDayModelWithDefaults instantiates a new DayModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DayModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DayModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DayModel) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *DayModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DayModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DayModel) SetDate(v string)`

SetDate sets Date field to given value.


### GetName

`func (o *DayModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DayModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DayModel) SetName(v string)`

SetName sets Name field to given value.


### GetWorking

`func (o *DayModel) GetWorking() bool`

GetWorking returns the Working field if non-nil, zero value otherwise.

### GetWorkingOk

`func (o *DayModel) GetWorkingOk() (*bool, bool)`

GetWorkingOk returns a tuple with the Working field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorking

`func (o *DayModel) SetWorking(v bool)`

SetWorking sets Working field to given value.


### GetLinks

`func (o *DayModel) GetLinks() DayModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DayModel) GetLinksOk() (*DayModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DayModel) SetLinks(v DayModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DayModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


