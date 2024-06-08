# RelationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Relation ID | [optional] [readonly] 
**Name** | Pointer to **string** | The internationalized name of this kind of relation | [optional] 
**Type** | Pointer to **string** | Which kind of relation (blocks, precedes, etc.) | [optional] 
**ReverseType** | Pointer to **string** | The kind of relation from the other WP&#39;s perspective | [optional] [readonly] 
**Description** | Pointer to **string** | Short text further describing the relation | [optional] 
**Lag** | Pointer to **int64** | The lag in days between closing of &#x60;from&#x60; and start of &#x60;to&#x60; | [optional] 
**Links** | Pointer to [**RelationModelLinks**](RelationModelLinks.md) |  | [optional] 

## Methods

### NewRelationModel

`func NewRelationModel() *RelationModel`

NewRelationModel instantiates a new RelationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationModelWithDefaults

`func NewRelationModelWithDefaults() *RelationModel`

NewRelationModelWithDefaults instantiates a new RelationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RelationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RelationModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RelationModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReverseType

`func (o *RelationModel) GetReverseType() string`

GetReverseType returns the ReverseType field if non-nil, zero value otherwise.

### GetReverseTypeOk

`func (o *RelationModel) GetReverseTypeOk() (*string, bool)`

GetReverseTypeOk returns a tuple with the ReverseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseType

`func (o *RelationModel) SetReverseType(v string)`

SetReverseType sets ReverseType field to given value.

### HasReverseType

`func (o *RelationModel) HasReverseType() bool`

HasReverseType returns a boolean if a field has been set.

### GetDescription

`func (o *RelationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelationModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelationModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLag

`func (o *RelationModel) GetLag() int64`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *RelationModel) GetLagOk() (*int64, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *RelationModel) SetLag(v int64)`

SetLag sets Lag field to given value.

### HasLag

`func (o *RelationModel) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetLinks

`func (o *RelationModel) GetLinks() RelationModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RelationModel) GetLinksOk() (*RelationModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RelationModel) SetLinks(v RelationModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RelationModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


