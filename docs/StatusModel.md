# StatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** | Status id | [optional] [readonly] 
**Name** | Pointer to **string** | Status name | [optional] [readonly] 
**Position** | Pointer to **int64** | Sort index of the status | [optional] [readonly] 
**IsDefault** | Pointer to **bool** |  | [optional] [readonly] 
**IsClosed** | Pointer to **bool** | are tickets of this status considered closed? | [optional] [readonly] 
**IsReadonly** | Pointer to **bool** | are tickets of this status read only? | [optional] [readonly] 
**DefaultDoneRatio** | Pointer to **int64** | The percentageDone being applied when changing to this status | [optional] 
**Links** | Pointer to [**StatusModelLinks**](StatusModelLinks.md) |  | [optional] 

## Methods

### NewStatusModel

`func NewStatusModel() *StatusModel`

NewStatusModel instantiates a new StatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusModelWithDefaults

`func NewStatusModelWithDefaults() *StatusModel`

NewStatusModelWithDefaults instantiates a new StatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StatusModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatusModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatusModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatusModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *StatusModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatusModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatusModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StatusModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StatusModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatusModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *StatusModel) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StatusModel) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StatusModel) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *StatusModel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsDefault

`func (o *StatusModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *StatusModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *StatusModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *StatusModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsClosed

`func (o *StatusModel) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *StatusModel) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *StatusModel) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *StatusModel) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsReadonly

`func (o *StatusModel) GetIsReadonly() bool`

GetIsReadonly returns the IsReadonly field if non-nil, zero value otherwise.

### GetIsReadonlyOk

`func (o *StatusModel) GetIsReadonlyOk() (*bool, bool)`

GetIsReadonlyOk returns a tuple with the IsReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonly

`func (o *StatusModel) SetIsReadonly(v bool)`

SetIsReadonly sets IsReadonly field to given value.

### HasIsReadonly

`func (o *StatusModel) HasIsReadonly() bool`

HasIsReadonly returns a boolean if a field has been set.

### GetDefaultDoneRatio

`func (o *StatusModel) GetDefaultDoneRatio() int64`

GetDefaultDoneRatio returns the DefaultDoneRatio field if non-nil, zero value otherwise.

### GetDefaultDoneRatioOk

`func (o *StatusModel) GetDefaultDoneRatioOk() (*int64, bool)`

GetDefaultDoneRatioOk returns a tuple with the DefaultDoneRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDoneRatio

`func (o *StatusModel) SetDefaultDoneRatio(v int64)`

SetDefaultDoneRatio sets DefaultDoneRatio field to given value.

### HasDefaultDoneRatio

`func (o *StatusModel) HasDefaultDoneRatio() bool`

HasDefaultDoneRatio returns a boolean if a field has been set.

### GetLinks

`func (o *StatusModel) GetLinks() StatusModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StatusModel) GetLinksOk() (*StatusModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StatusModel) SetLinks(v StatusModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StatusModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


