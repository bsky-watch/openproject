# StorageReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Storage id | 
**Type** | **string** |  | 
**Name** | **string** | Storage name | 
**HasApplicationPassword** | Pointer to **bool** | Whether the storage has the application password to use for the Nextcloud storage.  Ignored if the provider type is not Nextcloud | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the storage | [optional] 
**Embedded** | Pointer to [**StorageReadModelEmbedded**](StorageReadModelEmbedded.md) |  | [optional] 
**Links** | [**StorageReadModelLinks**](StorageReadModelLinks.md) |  | 

## Methods

### NewStorageReadModel

`func NewStorageReadModel(id int64, type_ string, name string, links StorageReadModelLinks, ) *StorageReadModel`

NewStorageReadModel instantiates a new StorageReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageReadModelWithDefaults

`func NewStorageReadModelWithDefaults() *StorageReadModel`

NewStorageReadModelWithDefaults instantiates a new StorageReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageReadModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageReadModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageReadModel) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *StorageReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *StorageReadModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageReadModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageReadModel) SetName(v string)`

SetName sets Name field to given value.


### GetHasApplicationPassword

`func (o *StorageReadModel) GetHasApplicationPassword() bool`

GetHasApplicationPassword returns the HasApplicationPassword field if non-nil, zero value otherwise.

### GetHasApplicationPasswordOk

`func (o *StorageReadModel) GetHasApplicationPasswordOk() (*bool, bool)`

GetHasApplicationPasswordOk returns a tuple with the HasApplicationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApplicationPassword

`func (o *StorageReadModel) SetHasApplicationPassword(v bool)`

SetHasApplicationPassword sets HasApplicationPassword field to given value.

### HasHasApplicationPassword

`func (o *StorageReadModel) HasHasApplicationPassword() bool`

HasHasApplicationPassword returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StorageReadModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageReadModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageReadModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StorageReadModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StorageReadModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StorageReadModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StorageReadModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StorageReadModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEmbedded

`func (o *StorageReadModel) GetEmbedded() StorageReadModelEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *StorageReadModel) GetEmbeddedOk() (*StorageReadModelEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *StorageReadModel) SetEmbedded(v StorageReadModelEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *StorageReadModel) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *StorageReadModel) GetLinks() StorageReadModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageReadModel) GetLinksOk() (*StorageReadModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageReadModel) SetLinks(v StorageReadModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


