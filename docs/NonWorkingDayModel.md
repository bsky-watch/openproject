# NonWorkingDayModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Date** | **string** | Date of the non-working day. | 
**Name** | **string** | Descriptive name for the non-working day. | 
**Links** | Pointer to [**NonWorkingDayModelLinks**](NonWorkingDayModelLinks.md) |  | [optional] 

## Methods

### NewNonWorkingDayModel

`func NewNonWorkingDayModel(type_ string, date string, name string, ) *NonWorkingDayModel`

NewNonWorkingDayModel instantiates a new NonWorkingDayModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonWorkingDayModelWithDefaults

`func NewNonWorkingDayModelWithDefaults() *NonWorkingDayModel`

NewNonWorkingDayModelWithDefaults instantiates a new NonWorkingDayModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NonWorkingDayModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NonWorkingDayModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NonWorkingDayModel) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *NonWorkingDayModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NonWorkingDayModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NonWorkingDayModel) SetDate(v string)`

SetDate sets Date field to given value.


### GetName

`func (o *NonWorkingDayModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NonWorkingDayModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NonWorkingDayModel) SetName(v string)`

SetName sets Name field to given value.


### GetLinks

`func (o *NonWorkingDayModel) GetLinks() NonWorkingDayModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NonWorkingDayModel) GetLinksOk() (*NonWorkingDayModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NonWorkingDayModel) SetLinks(v NonWorkingDayModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NonWorkingDayModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


