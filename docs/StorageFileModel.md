# StorageFileModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Linked file&#39;s id on the origin | 
**Name** | **string** | Linked file&#39;s name on the origin | 
**MimeType** | Pointer to **string** | MIME type of the linked file.  To link a folder entity, the custom MIME type &#x60;application/x-op-directory&#x60; MUST be provided. Otherwise it defaults back to an unknown MIME type. | [optional] 
**Size** | Pointer to **int64** | file size on origin in bytes | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp of the creation datetime of the file on the origin | [optional] 
**LastModifiedAt** | Pointer to **time.Time** | Timestamp of the datetime of the last modification of the file on the origin | [optional] 
**CreatedByName** | Pointer to **string** | Display name of the author that created the file on the origin | [optional] 
**LastModifiedByName** | Pointer to **string** | Display name of the author that modified the file on the origin last | [optional] 
**Type** | **string** |  | 
**Location** | **string** | Location identification for file in storage | 
**Links** | [**StorageFileModelAllOfLinks**](StorageFileModelAllOfLinks.md) |  | 

## Methods

### NewStorageFileModel

`func NewStorageFileModel(id string, name string, type_ string, location string, links StorageFileModelAllOfLinks, ) *StorageFileModel`

NewStorageFileModel instantiates a new StorageFileModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFileModelWithDefaults

`func NewStorageFileModelWithDefaults() *StorageFileModel`

NewStorageFileModelWithDefaults instantiates a new StorageFileModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageFileModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageFileModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageFileModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StorageFileModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageFileModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageFileModel) SetName(v string)`

SetName sets Name field to given value.


### GetMimeType

`func (o *StorageFileModel) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *StorageFileModel) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *StorageFileModel) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *StorageFileModel) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSize

`func (o *StorageFileModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFileModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFileModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFileModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StorageFileModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageFileModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageFileModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StorageFileModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *StorageFileModel) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *StorageFileModel) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *StorageFileModel) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *StorageFileModel) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *StorageFileModel) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *StorageFileModel) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *StorageFileModel) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *StorageFileModel) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetLastModifiedByName

`func (o *StorageFileModel) GetLastModifiedByName() string`

GetLastModifiedByName returns the LastModifiedByName field if non-nil, zero value otherwise.

### GetLastModifiedByNameOk

`func (o *StorageFileModel) GetLastModifiedByNameOk() (*string, bool)`

GetLastModifiedByNameOk returns a tuple with the LastModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByName

`func (o *StorageFileModel) SetLastModifiedByName(v string)`

SetLastModifiedByName sets LastModifiedByName field to given value.

### HasLastModifiedByName

`func (o *StorageFileModel) HasLastModifiedByName() bool`

HasLastModifiedByName returns a boolean if a field has been set.

### GetType

`func (o *StorageFileModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageFileModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageFileModel) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *StorageFileModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageFileModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageFileModel) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetLinks

`func (o *StorageFileModel) GetLinks() StorageFileModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageFileModel) GetLinksOk() (*StorageFileModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageFileModel) SetLinks(v StorageFileModelAllOfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


