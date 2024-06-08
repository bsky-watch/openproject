# RevisionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Revision&#39;s id, assigned by OpenProject | [optional] [readonly] 
**Identifier** | **string** | The raw SCM identifier of the revision (e.g. full SHA hash) | [readonly] 
**FormattedIdentifier** | **string** | The SCM identifier of the revision, formatted (e.g. shortened unambiguous SHA hash). May be identical to identifier in many cases | [readonly] 
**AuthorName** | **string** | The name of the author that committed this revision. Note that this name is retrieved from the repository and does not identify a user in OpenProject. | [readonly] 
**Message** | [**RevisionModelMessage**](RevisionModelMessage.md) |  | 
**CreatedAt** | **time.Time** | The time this revision was committed to the repository | 
**Links** | Pointer to [**RevisionModelLinks**](RevisionModelLinks.md) |  | [optional] 

## Methods

### NewRevisionModel

`func NewRevisionModel(identifier string, formattedIdentifier string, authorName string, message RevisionModelMessage, createdAt time.Time, ) *RevisionModel`

NewRevisionModel instantiates a new RevisionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionModelWithDefaults

`func NewRevisionModelWithDefaults() *RevisionModel`

NewRevisionModelWithDefaults instantiates a new RevisionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RevisionModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RevisionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *RevisionModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RevisionModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RevisionModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetFormattedIdentifier

`func (o *RevisionModel) GetFormattedIdentifier() string`

GetFormattedIdentifier returns the FormattedIdentifier field if non-nil, zero value otherwise.

### GetFormattedIdentifierOk

`func (o *RevisionModel) GetFormattedIdentifierOk() (*string, bool)`

GetFormattedIdentifierOk returns a tuple with the FormattedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedIdentifier

`func (o *RevisionModel) SetFormattedIdentifier(v string)`

SetFormattedIdentifier sets FormattedIdentifier field to given value.


### GetAuthorName

`func (o *RevisionModel) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *RevisionModel) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *RevisionModel) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetMessage

`func (o *RevisionModel) GetMessage() RevisionModelMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RevisionModel) GetMessageOk() (*RevisionModelMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RevisionModel) SetMessage(v RevisionModelMessage)`

SetMessage sets Message field to given value.


### GetCreatedAt

`func (o *RevisionModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RevisionModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RevisionModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLinks

`func (o *RevisionModel) GetLinks() RevisionModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RevisionModel) GetLinksOk() (*RevisionModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RevisionModel) SetLinks(v RevisionModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RevisionModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


