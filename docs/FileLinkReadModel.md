# FileLinkReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | File link id | 
**Type** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the file link | [optional] 
**OriginData** | [**FileLinkOriginDataModel**](FileLinkOriginDataModel.md) |  | 
**Embedded** | Pointer to [**FileLinkReadModelEmbedded**](FileLinkReadModelEmbedded.md) |  | [optional] 
**Links** | [**FileLinkReadModelLinks**](FileLinkReadModelLinks.md) |  | 

## Methods

### NewFileLinkReadModel

`func NewFileLinkReadModel(id int64, type_ string, originData FileLinkOriginDataModel, links FileLinkReadModelLinks, ) *FileLinkReadModel`

NewFileLinkReadModel instantiates a new FileLinkReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLinkReadModelWithDefaults

`func NewFileLinkReadModelWithDefaults() *FileLinkReadModel`

NewFileLinkReadModelWithDefaults instantiates a new FileLinkReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileLinkReadModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileLinkReadModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileLinkReadModel) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *FileLinkReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileLinkReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileLinkReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *FileLinkReadModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileLinkReadModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileLinkReadModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileLinkReadModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileLinkReadModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileLinkReadModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileLinkReadModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FileLinkReadModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOriginData

`func (o *FileLinkReadModel) GetOriginData() FileLinkOriginDataModel`

GetOriginData returns the OriginData field if non-nil, zero value otherwise.

### GetOriginDataOk

`func (o *FileLinkReadModel) GetOriginDataOk() (*FileLinkOriginDataModel, bool)`

GetOriginDataOk returns a tuple with the OriginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginData

`func (o *FileLinkReadModel) SetOriginData(v FileLinkOriginDataModel)`

SetOriginData sets OriginData field to given value.


### GetEmbedded

`func (o *FileLinkReadModel) GetEmbedded() FileLinkReadModelEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *FileLinkReadModel) GetEmbeddedOk() (*FileLinkReadModelEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *FileLinkReadModel) SetEmbedded(v FileLinkReadModelEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *FileLinkReadModel) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *FileLinkReadModel) GetLinks() FileLinkReadModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FileLinkReadModel) GetLinksOk() (*FileLinkReadModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FileLinkReadModel) SetLinks(v FileLinkReadModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


