# TypesByProjectModelAllOfEmbeddedElements

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

### NewTypesByProjectModelAllOfEmbeddedElements

`func NewTypesByProjectModelAllOfEmbeddedElements() *TypesByProjectModelAllOfEmbeddedElements`

NewTypesByProjectModelAllOfEmbeddedElements instantiates a new TypesByProjectModelAllOfEmbeddedElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesByProjectModelAllOfEmbeddedElementsWithDefaults

`func NewTypesByProjectModelAllOfEmbeddedElementsWithDefaults() *TypesByProjectModelAllOfEmbeddedElements`

NewTypesByProjectModelAllOfEmbeddedElementsWithDefaults instantiates a new TypesByProjectModelAllOfEmbeddedElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPosition

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsDefault

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsMilestone

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsMilestone() bool`

GetIsMilestone returns the IsMilestone field if non-nil, zero value otherwise.

### GetIsMilestoneOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetIsMilestoneOk() (*bool, bool)`

GetIsMilestoneOk returns a tuple with the IsMilestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMilestone

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetIsMilestone(v bool)`

SetIsMilestone sets IsMilestone field to given value.

### HasIsMilestone

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasIsMilestone() bool`

HasIsMilestone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetLinks() TypeModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TypesByProjectModelAllOfEmbeddedElements) GetLinksOk() (*TypeModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TypesByProjectModelAllOfEmbeddedElements) SetLinks(v TypeModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TypesByProjectModelAllOfEmbeddedElements) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


