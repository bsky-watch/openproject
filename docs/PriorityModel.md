# PriorityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Priority id | [optional] [readonly] 
**Name** | Pointer to **string** | Priority name | [optional] [readonly] 
**Position** | Pointer to **int64** | Sort index of the priority | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Indicates whether this is the default value | [optional] [readonly] 
**IsActive** | Pointer to **bool** | Indicates whether the priority is available | [optional] 
**Links** | Pointer to [**PriorityModelLinks**](PriorityModelLinks.md) |  | [optional] 

## Methods

### NewPriorityModel

`func NewPriorityModel() *PriorityModel`

NewPriorityModel instantiates a new PriorityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityModelWithDefaults

`func NewPriorityModelWithDefaults() *PriorityModel`

NewPriorityModelWithDefaults instantiates a new PriorityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriorityModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriorityModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriorityModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PriorityModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PriorityModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriorityModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriorityModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriorityModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *PriorityModel) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PriorityModel) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PriorityModel) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PriorityModel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsDefault

`func (o *PriorityModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PriorityModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PriorityModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PriorityModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsActive

`func (o *PriorityModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PriorityModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PriorityModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PriorityModel) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLinks

`func (o *PriorityModel) GetLinks() PriorityModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PriorityModel) GetLinksOk() (*PriorityModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PriorityModel) SetLinks(v PriorityModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PriorityModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


