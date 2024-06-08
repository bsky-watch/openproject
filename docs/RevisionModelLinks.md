# RevisionModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**RevisionModelLinksSelf**](RevisionModelLinksSelf.md) |  | 
**Project** | [**RevisionModelLinksProject**](RevisionModelLinksProject.md) |  | 
**Author** | Pointer to [**RevisionModelLinksAuthor**](RevisionModelLinksAuthor.md) |  | [optional] 
**ShowRevision** | [**RevisionModelLinksShowRevision**](RevisionModelLinksShowRevision.md) |  | 

## Methods

### NewRevisionModelLinks

`func NewRevisionModelLinks(self RevisionModelLinksSelf, project RevisionModelLinksProject, showRevision RevisionModelLinksShowRevision, ) *RevisionModelLinks`

NewRevisionModelLinks instantiates a new RevisionModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionModelLinksWithDefaults

`func NewRevisionModelLinksWithDefaults() *RevisionModelLinks`

NewRevisionModelLinksWithDefaults instantiates a new RevisionModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *RevisionModelLinks) GetSelf() RevisionModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RevisionModelLinks) GetSelfOk() (*RevisionModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RevisionModelLinks) SetSelf(v RevisionModelLinksSelf)`

SetSelf sets Self field to given value.


### GetProject

`func (o *RevisionModelLinks) GetProject() RevisionModelLinksProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RevisionModelLinks) GetProjectOk() (*RevisionModelLinksProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RevisionModelLinks) SetProject(v RevisionModelLinksProject)`

SetProject sets Project field to given value.


### GetAuthor

`func (o *RevisionModelLinks) GetAuthor() RevisionModelLinksAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RevisionModelLinks) GetAuthorOk() (*RevisionModelLinksAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RevisionModelLinks) SetAuthor(v RevisionModelLinksAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RevisionModelLinks) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetShowRevision

`func (o *RevisionModelLinks) GetShowRevision() RevisionModelLinksShowRevision`

GetShowRevision returns the ShowRevision field if non-nil, zero value otherwise.

### GetShowRevisionOk

`func (o *RevisionModelLinks) GetShowRevisionOk() (*RevisionModelLinksShowRevision, bool)`

GetShowRevisionOk returns a tuple with the ShowRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRevision

`func (o *RevisionModelLinks) SetShowRevision(v RevisionModelLinksShowRevision)`

SetShowRevision sets ShowRevision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


