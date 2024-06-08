# GroupWriteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new group name. | [optional] 
**Links** | Pointer to [**GroupWriteModelLinks**](GroupWriteModelLinks.md) |  | [optional] 

## Methods

### NewGroupWriteModel

`func NewGroupWriteModel() *GroupWriteModel`

NewGroupWriteModel instantiates a new GroupWriteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWriteModelWithDefaults

`func NewGroupWriteModelWithDefaults() *GroupWriteModel`

NewGroupWriteModelWithDefaults instantiates a new GroupWriteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupWriteModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWriteModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWriteModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupWriteModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLinks

`func (o *GroupWriteModel) GetLinks() GroupWriteModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupWriteModel) GetLinksOk() (*GroupWriteModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupWriteModel) SetLinks(v GroupWriteModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupWriteModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


