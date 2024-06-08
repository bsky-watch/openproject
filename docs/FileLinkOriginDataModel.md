# FileLinkOriginDataModel

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

## Methods

### NewFileLinkOriginDataModel

`func NewFileLinkOriginDataModel(id string, name string, ) *FileLinkOriginDataModel`

NewFileLinkOriginDataModel instantiates a new FileLinkOriginDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLinkOriginDataModelWithDefaults

`func NewFileLinkOriginDataModelWithDefaults() *FileLinkOriginDataModel`

NewFileLinkOriginDataModelWithDefaults instantiates a new FileLinkOriginDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileLinkOriginDataModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileLinkOriginDataModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileLinkOriginDataModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FileLinkOriginDataModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileLinkOriginDataModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileLinkOriginDataModel) SetName(v string)`

SetName sets Name field to given value.


### GetMimeType

`func (o *FileLinkOriginDataModel) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *FileLinkOriginDataModel) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *FileLinkOriginDataModel) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *FileLinkOriginDataModel) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSize

`func (o *FileLinkOriginDataModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileLinkOriginDataModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileLinkOriginDataModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileLinkOriginDataModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FileLinkOriginDataModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileLinkOriginDataModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileLinkOriginDataModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileLinkOriginDataModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *FileLinkOriginDataModel) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *FileLinkOriginDataModel) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *FileLinkOriginDataModel) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *FileLinkOriginDataModel) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetCreatedByName

`func (o *FileLinkOriginDataModel) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *FileLinkOriginDataModel) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *FileLinkOriginDataModel) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *FileLinkOriginDataModel) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetLastModifiedByName

`func (o *FileLinkOriginDataModel) GetLastModifiedByName() string`

GetLastModifiedByName returns the LastModifiedByName field if non-nil, zero value otherwise.

### GetLastModifiedByNameOk

`func (o *FileLinkOriginDataModel) GetLastModifiedByNameOk() (*string, bool)`

GetLastModifiedByNameOk returns a tuple with the LastModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByName

`func (o *FileLinkOriginDataModel) SetLastModifiedByName(v string)`

SetLastModifiedByName sets LastModifiedByName field to given value.

### HasLastModifiedByName

`func (o *FileLinkOriginDataModel) HasLastModifiedByName() bool`

HasLastModifiedByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


