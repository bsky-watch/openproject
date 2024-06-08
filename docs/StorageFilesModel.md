# StorageFilesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Files** | [**[]StorageFileModel**](StorageFileModel.md) | List of files provided by the selected storage. | 
**Parent** | [**StorageFilesModelParent**](StorageFilesModelParent.md) |  | 
**Ancestors** | [**[]StorageFileModel**](StorageFileModel.md) | List of ancestors of the parent directory. Can be empty, if parent directory was root directory. | 
**Links** | [**StorageFileModelAllOfLinks**](StorageFileModelAllOfLinks.md) |  | 

## Methods

### NewStorageFilesModel

`func NewStorageFilesModel(type_ string, files []StorageFileModel, parent StorageFilesModelParent, ancestors []StorageFileModel, links StorageFileModelAllOfLinks, ) *StorageFilesModel`

NewStorageFilesModel instantiates a new StorageFilesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFilesModelWithDefaults

`func NewStorageFilesModelWithDefaults() *StorageFilesModel`

NewStorageFilesModelWithDefaults instantiates a new StorageFilesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StorageFilesModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageFilesModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageFilesModel) SetType(v string)`

SetType sets Type field to given value.


### GetFiles

`func (o *StorageFilesModel) GetFiles() []StorageFileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *StorageFilesModel) GetFilesOk() (*[]StorageFileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *StorageFilesModel) SetFiles(v []StorageFileModel)`

SetFiles sets Files field to given value.


### GetParent

`func (o *StorageFilesModel) GetParent() StorageFilesModelParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageFilesModel) GetParentOk() (*StorageFilesModelParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageFilesModel) SetParent(v StorageFilesModelParent)`

SetParent sets Parent field to given value.


### GetAncestors

`func (o *StorageFilesModel) GetAncestors() []StorageFileModel`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageFilesModel) GetAncestorsOk() (*[]StorageFileModel, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageFilesModel) SetAncestors(v []StorageFileModel)`

SetAncestors sets Ancestors field to given value.


### GetLinks

`func (o *StorageFilesModel) GetLinks() StorageFileModelAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageFilesModel) GetLinksOk() (*StorageFileModelAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageFilesModel) SetLinks(v StorageFileModelAllOfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


