# RelationModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Update** | Pointer to [**RelationModelLinksUpdate**](RelationModelLinksUpdate.md) |  | [optional] 
**UpdateImmediately** | Pointer to [**RelationModelLinksUpdateImmediately**](RelationModelLinksUpdateImmediately.md) |  | [optional] 
**Delete** | Pointer to [**RelationModelLinksDelete**](RelationModelLinksDelete.md) |  | [optional] 
**Self** | [**RelationModelLinksSelf**](RelationModelLinksSelf.md) |  | 
**Schema** | [**RelationModelLinksSchema**](RelationModelLinksSchema.md) |  | 
**From** | [**RelationModelLinksFrom**](RelationModelLinksFrom.md) |  | 
**To** | [**RelationModelLinksTo**](RelationModelLinksTo.md) |  | 

## Methods

### NewRelationModelLinks

`func NewRelationModelLinks(self RelationModelLinksSelf, schema RelationModelLinksSchema, from RelationModelLinksFrom, to RelationModelLinksTo, ) *RelationModelLinks`

NewRelationModelLinks instantiates a new RelationModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationModelLinksWithDefaults

`func NewRelationModelLinksWithDefaults() *RelationModelLinks`

NewRelationModelLinksWithDefaults instantiates a new RelationModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdate

`func (o *RelationModelLinks) GetUpdate() RelationModelLinksUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *RelationModelLinks) GetUpdateOk() (*RelationModelLinksUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *RelationModelLinks) SetUpdate(v RelationModelLinksUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *RelationModelLinks) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdateImmediately

`func (o *RelationModelLinks) GetUpdateImmediately() RelationModelLinksUpdateImmediately`

GetUpdateImmediately returns the UpdateImmediately field if non-nil, zero value otherwise.

### GetUpdateImmediatelyOk

`func (o *RelationModelLinks) GetUpdateImmediatelyOk() (*RelationModelLinksUpdateImmediately, bool)`

GetUpdateImmediatelyOk returns a tuple with the UpdateImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateImmediately

`func (o *RelationModelLinks) SetUpdateImmediately(v RelationModelLinksUpdateImmediately)`

SetUpdateImmediately sets UpdateImmediately field to given value.

### HasUpdateImmediately

`func (o *RelationModelLinks) HasUpdateImmediately() bool`

HasUpdateImmediately returns a boolean if a field has been set.

### GetDelete

`func (o *RelationModelLinks) GetDelete() RelationModelLinksDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *RelationModelLinks) GetDeleteOk() (*RelationModelLinksDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *RelationModelLinks) SetDelete(v RelationModelLinksDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *RelationModelLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetSelf

`func (o *RelationModelLinks) GetSelf() RelationModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RelationModelLinks) GetSelfOk() (*RelationModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RelationModelLinks) SetSelf(v RelationModelLinksSelf)`

SetSelf sets Self field to given value.


### GetSchema

`func (o *RelationModelLinks) GetSchema() RelationModelLinksSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RelationModelLinks) GetSchemaOk() (*RelationModelLinksSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RelationModelLinks) SetSchema(v RelationModelLinksSchema)`

SetSchema sets Schema field to given value.


### GetFrom

`func (o *RelationModelLinks) GetFrom() RelationModelLinksFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RelationModelLinks) GetFromOk() (*RelationModelLinksFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RelationModelLinks) SetFrom(v RelationModelLinksFrom)`

SetFrom sets From field to given value.


### GetTo

`func (o *RelationModelLinks) GetTo() RelationModelLinksTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RelationModelLinks) GetToOk() (*RelationModelLinksTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RelationModelLinks) SetTo(v RelationModelLinksTo)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


