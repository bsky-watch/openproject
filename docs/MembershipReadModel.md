# MembershipReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | The membership&#39;s id | 
**CreatedAt** | **time.Time** | The time the membership was created. | 
**UpdatedAt** | **time.Time** | The time the membership was last updated. | 
**Embedded** | Pointer to [**MembershipReadModelEmbedded**](MembershipReadModelEmbedded.md) |  | [optional] 
**Links** | [**MembershipReadModelLinks**](MembershipReadModelLinks.md) |  | 

## Methods

### NewMembershipReadModel

`func NewMembershipReadModel(type_ string, id int64, createdAt time.Time, updatedAt time.Time, links MembershipReadModelLinks, ) *MembershipReadModel`

NewMembershipReadModel instantiates a new MembershipReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipReadModelWithDefaults

`func NewMembershipReadModelWithDefaults() *MembershipReadModel`

NewMembershipReadModelWithDefaults instantiates a new MembershipReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MembershipReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MembershipReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MembershipReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MembershipReadModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MembershipReadModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MembershipReadModel) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *MembershipReadModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MembershipReadModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MembershipReadModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MembershipReadModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MembershipReadModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MembershipReadModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmbedded

`func (o *MembershipReadModel) GetEmbedded() MembershipReadModelEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *MembershipReadModel) GetEmbeddedOk() (*MembershipReadModelEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *MembershipReadModel) SetEmbedded(v MembershipReadModelEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *MembershipReadModel) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *MembershipReadModel) GetLinks() MembershipReadModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipReadModel) GetLinksOk() (*MembershipReadModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipReadModel) SetLinks(v MembershipReadModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


