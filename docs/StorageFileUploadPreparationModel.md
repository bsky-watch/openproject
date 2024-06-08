# StorageFileUploadPreparationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int64** | The project identifier, from where a user starts uploading a file. | 
**FileName** | **string** | The file name. | 
**Parent** | **string** | The directory to which the file is to be uploaded. For root directories, the value &#x60;/&#x60; must be provided. | 

## Methods

### NewStorageFileUploadPreparationModel

`func NewStorageFileUploadPreparationModel(projectId int64, fileName string, parent string, ) *StorageFileUploadPreparationModel`

NewStorageFileUploadPreparationModel instantiates a new StorageFileUploadPreparationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFileUploadPreparationModelWithDefaults

`func NewStorageFileUploadPreparationModelWithDefaults() *StorageFileUploadPreparationModel`

NewStorageFileUploadPreparationModelWithDefaults instantiates a new StorageFileUploadPreparationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *StorageFileUploadPreparationModel) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StorageFileUploadPreparationModel) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StorageFileUploadPreparationModel) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetFileName

`func (o *StorageFileUploadPreparationModel) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *StorageFileUploadPreparationModel) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *StorageFileUploadPreparationModel) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetParent

`func (o *StorageFileUploadPreparationModel) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageFileUploadPreparationModel) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageFileUploadPreparationModel) SetParent(v string)`

SetParent sets Parent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


