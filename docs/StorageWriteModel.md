# StorageWriteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Storage name, if not provided, falls back to a default. | [optional] 
**ApplicationPassword** | Pointer to **string** | The application password to use for the Nextcloud storage. Ignored if the provider type is not Nextcloud.  If a string is provided, the password is set and automatic management is enabled for the storage. If null is provided, the password is unset and automatic management is disabled for the storage. | [optional] 
**Links** | Pointer to [**StorageWriteModelLinks**](StorageWriteModelLinks.md) |  | [optional] 

## Methods

### NewStorageWriteModel

`func NewStorageWriteModel() *StorageWriteModel`

NewStorageWriteModel instantiates a new StorageWriteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageWriteModelWithDefaults

`func NewStorageWriteModelWithDefaults() *StorageWriteModel`

NewStorageWriteModelWithDefaults instantiates a new StorageWriteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageWriteModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageWriteModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageWriteModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageWriteModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplicationPassword

`func (o *StorageWriteModel) GetApplicationPassword() string`

GetApplicationPassword returns the ApplicationPassword field if non-nil, zero value otherwise.

### GetApplicationPasswordOk

`func (o *StorageWriteModel) GetApplicationPasswordOk() (*string, bool)`

GetApplicationPasswordOk returns a tuple with the ApplicationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPassword

`func (o *StorageWriteModel) SetApplicationPassword(v string)`

SetApplicationPassword sets ApplicationPassword field to given value.

### HasApplicationPassword

`func (o *StorageWriteModel) HasApplicationPassword() bool`

HasApplicationPassword returns a boolean if a field has been set.

### GetLinks

`func (o *StorageWriteModel) GetLinks() StorageWriteModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageWriteModel) GetLinksOk() (*StorageWriteModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageWriteModel) SetLinks(v StorageWriteModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StorageWriteModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


