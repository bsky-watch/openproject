# TypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Type id | [optional] [readonly] 
**Name** | Pointer to **string** | Type name | [optional] [readonly] 
**Color** | Pointer to **string** | The color used to represent this type | [optional] [readonly] 
**Position** | Pointer to **int64** | Sort index of the type | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Is this type active by default in new projects? | [optional] [readonly] 
**IsMilestone** | Pointer to **bool** | Do work packages of this type represent a milestone? | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the user | [optional] 
**Links** | Pointer to [**TypeModelLinks**](TypeModelLinks.md) |  | [optional] 

## Methods

### NewTypeModel

`func NewTypeModel() *TypeModel`

NewTypeModel instantiates a new TypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeModelWithDefaults

`func NewTypeModelWithDefaults() *TypeModel`

NewTypeModelWithDefaults instantiates a new TypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypeModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypeModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypeModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TypeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *TypeModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TypeModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TypeModel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TypeModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPosition

`func (o *TypeModel) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TypeModel) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TypeModel) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TypeModel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsDefault

`func (o *TypeModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TypeModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TypeModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TypeModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsMilestone

`func (o *TypeModel) GetIsMilestone() bool`

GetIsMilestone returns the IsMilestone field if non-nil, zero value otherwise.

### GetIsMilestoneOk

`func (o *TypeModel) GetIsMilestoneOk() (*bool, bool)`

GetIsMilestoneOk returns a tuple with the IsMilestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMilestone

`func (o *TypeModel) SetIsMilestone(v bool)`

SetIsMilestone sets IsMilestone field to given value.

### HasIsMilestone

`func (o *TypeModel) HasIsMilestone() bool`

HasIsMilestone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TypeModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypeModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypeModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TypeModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TypeModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypeModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypeModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TypeModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *TypeModel) GetLinks() TypeModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TypeModel) GetLinksOk() (*TypeModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TypeModel) SetLinks(v TypeModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TypeModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


