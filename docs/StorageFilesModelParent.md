# StorageFilesModelParent

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

### NewStorageFilesModelParent

`func NewStorageFilesModelParent(id string, name string, type_ string, location string, links StorageFileModelAllOfLinks, ) *StorageFilesModelParent`

NewStorageFilesModelParent instantiates a new StorageFilesModelParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFilesModelParentWithDefaults

`func NewStorageFilesModelParentWithDefaults() *StorageFilesModelParent`

NewStorageFilesModelParentWithDefaults instantiates a new StorageFilesModelParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageFilesModelParent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageFilesModelParent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageFilesModelParent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StorageFilesModelParent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageFilesModelParent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageFilesModelParent) SetName(v string)`

SetName sets Name field to given value.


### GetMimeType

`func (o *StorageFilesModelParent) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *StorageFilesModelParent) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *StorageFilesModelParent) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *StorageFilesModelParent) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSize

`func (o *StorageFilesModelParent) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFilesModelParent) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFilesModelParent) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFilesModelParent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StorageFilesModelParent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageFilesModelParent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageFilesModelParent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StorageFilesModelParent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *StorageFilesModelParent) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *StorageFilesModelParent) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *StorageFilesModelParent) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *StorageFilesModelParent) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *StorageFilesModelParent) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *StorageFilesModelParent) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *StorageFilesModelParent) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *StorageFilesModelParent) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetLastModifiedByName

`func (o *StorageFilesModelParent) GetLastModifiedByName() string`

GetLastModifiedByName returns the LastModifiedByName field if non-nil, zero value otherwise.

### GetLastModifiedByNameOk

`func (o *StorageFilesModelParent) GetLastModifiedByNameOk() (*string, bool)`

GetLastModifiedByNameOk returns a tuple with the LastModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByName

`func (o *StorageFilesModelParent) SetLastModifiedByName(v string)`

SetLastModifiedByName sets LastModifiedByName field to given value.

### HasLastModifiedByName

`func (o *StorageFilesModelParent) HasLastModifiedByName() bool`

HasLastModifiedByName returns a boolean if a field has been set.

### GetType

`func (o *StorageFilesModelParent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageFilesModelParent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageFilesModelParent) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *StorageFilesModelParent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageFilesModelParent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageFilesModelParent) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetLinks

`func (o *StorageFilesModelParent) GetLinks() StorageFileModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageFilesModelParent) GetLinksOk() (*StorageFileModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageFilesModelParent) SetLinks(v StorageFileModelAllOfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


