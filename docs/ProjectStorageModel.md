# ProjectStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | The project storage&#39;s id | 
**ProjectFolderMode** | **string** |  | 
**CreatedAt** | **time.Time** | Time of creation | 
**UpdatedAt** | **time.Time** | Time of the most recent change to the project storage | 
**Links** | Pointer to [**ProjectStorageModelLinks**](ProjectStorageModelLinks.md) |  | [optional] 

## Methods

### NewProjectStorageModel

`func NewProjectStorageModel(type_ string, id int64, projectFolderMode string, createdAt time.Time, updatedAt time.Time, ) *ProjectStorageModel`

NewProjectStorageModel instantiates a new ProjectStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectStorageModelWithDefaults

`func NewProjectStorageModelWithDefaults() *ProjectStorageModel`

NewProjectStorageModelWithDefaults instantiates a new ProjectStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProjectStorageModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectStorageModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectStorageModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ProjectStorageModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectStorageModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectStorageModel) SetId(v int64)`

SetId sets Id field to given value.


### GetProjectFolderMode

`func (o *ProjectStorageModel) GetProjectFolderMode() string`

GetProjectFolderMode returns the ProjectFolderMode field if non-nil, zero value otherwise.

### GetProjectFolderModeOk

`func (o *ProjectStorageModel) GetProjectFolderModeOk() (*string, bool)`

GetProjectFolderModeOk returns a tuple with the ProjectFolderMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectFolderMode

`func (o *ProjectStorageModel) SetProjectFolderMode(v string)`

SetProjectFolderMode sets ProjectFolderMode field to given value.


### GetCreatedAt

`func (o *ProjectStorageModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectStorageModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectStorageModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectStorageModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectStorageModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectStorageModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *ProjectStorageModel) GetLinks() ProjectStorageModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectStorageModel) GetLinksOk() (*ProjectStorageModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectStorageModel) SetLinks(v ProjectStorageModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectStorageModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


