# ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** | Projects&#39; id | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** | Indicates whether the project is currently active or already archived | [optional] 
**StatusExplanation** | Pointer to [**ProjectModelStatusExplanation**](ProjectModelStatusExplanation.md) |  | [optional] 
**Public** | Pointer to **bool** | Indicates whether the project is accessible for everybody | [optional] 
**Description** | Pointer to [**Formattable**](Formattable.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time of creation. Can be writable by admins with the &#x60;apiv3_write_readonly_attributes&#x60; setting enabled. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the project | [optional] 
**Links** | Pointer to [**ProjectModelLinks**](ProjectModelLinks.md) |  | [optional] 

## Methods

### NewListAvailableParentProjectCandidatesModelAllOfEmbeddedElements

`func NewListAvailableParentProjectCandidatesModelAllOfEmbeddedElements() *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements`

NewListAvailableParentProjectCandidatesModelAllOfEmbeddedElements instantiates a new ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAvailableParentProjectCandidatesModelAllOfEmbeddedElementsWithDefaults

`func NewListAvailableParentProjectCandidatesModelAllOfEmbeddedElementsWithDefaults() *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements`

NewListAvailableParentProjectCandidatesModelAllOfEmbeddedElementsWithDefaults instantiates a new ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatusExplanation

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetStatusExplanation() ProjectModelStatusExplanation`

GetStatusExplanation returns the StatusExplanation field if non-nil, zero value otherwise.

### GetStatusExplanationOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetStatusExplanationOk() (*ProjectModelStatusExplanation, bool)`

GetStatusExplanationOk returns a tuple with the StatusExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusExplanation

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetStatusExplanation(v ProjectModelStatusExplanation)`

SetStatusExplanation sets StatusExplanation field to given value.

### HasStatusExplanation

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasStatusExplanation() bool`

HasStatusExplanation returns a boolean if a field has been set.

### GetPublic

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetDescription

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetDescription() Formattable`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetDescriptionOk() (*Formattable, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetDescription(v Formattable)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetLinks() ProjectModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) GetLinksOk() (*ProjectModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) SetLinks(v ProjectModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListAvailableParentProjectCandidatesModelAllOfEmbeddedElements) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


