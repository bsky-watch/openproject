# DocumentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Document&#39;s id | [optional] [readonly] 
**Title** | Pointer to **string** | The title chosen for the collection of documents | [optional] [readonly] 
**Description** | Pointer to **string** | A text describing the documents | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The time the document was created at | [optional] [readonly] 
**Links** | Pointer to [**DocumentModelLinks**](DocumentModelLinks.md) |  | [optional] 

## Methods

### NewDocumentModel

`func NewDocumentModel() *DocumentModel`

NewDocumentModel instantiates a new DocumentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentModelWithDefaults

`func NewDocumentModelWithDefaults() *DocumentModel`

NewDocumentModelWithDefaults instantiates a new DocumentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *DocumentModel) GetLinks() DocumentModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DocumentModel) GetLinksOk() (*DocumentModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DocumentModel) SetLinks(v DocumentModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DocumentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


