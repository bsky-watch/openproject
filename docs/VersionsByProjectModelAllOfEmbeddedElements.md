# VersionsByProjectModelAllOfEmbeddedElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Version id | [optional] [readonly] 
**Name** | **string** | Version name | 
**Description** | Pointer to [**ActivityModelComment**](ActivityModelComment.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Status** | **string** | The current status of the version | 
**Sharing** | **string** | The current status of the version | 
**CreatedAt** | **time.Time** | Time of creation | [readonly] 
**UpdatedAt** | **time.Time** | Time of the most recent change to the version | [readonly] 
**Links** | Pointer to [**VersionModelLinks**](VersionModelLinks.md) |  | [optional] 

## Methods

### NewVersionsByProjectModelAllOfEmbeddedElements

`func NewVersionsByProjectModelAllOfEmbeddedElements(name string, status string, sharing string, createdAt time.Time, updatedAt time.Time, ) *VersionsByProjectModelAllOfEmbeddedElements`

NewVersionsByProjectModelAllOfEmbeddedElements instantiates a new VersionsByProjectModelAllOfEmbeddedElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsByProjectModelAllOfEmbeddedElementsWithDefaults

`func NewVersionsByProjectModelAllOfEmbeddedElementsWithDefaults() *VersionsByProjectModelAllOfEmbeddedElements`

NewVersionsByProjectModelAllOfEmbeddedElementsWithDefaults instantiates a new VersionsByProjectModelAllOfEmbeddedElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VersionsByProjectModelAllOfEmbeddedElements) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetDescription() ActivityModelComment`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetDescriptionOk() (*ActivityModelComment, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetDescription(v ActivityModelComment)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionsByProjectModelAllOfEmbeddedElements) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *VersionsByProjectModelAllOfEmbeddedElements) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *VersionsByProjectModelAllOfEmbeddedElements) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSharing

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetSharing() string`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetSharingOk() (*string, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetSharing(v string)`

SetSharing sets Sharing field to given value.


### GetCreatedAt

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetLinks() VersionModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VersionsByProjectModelAllOfEmbeddedElements) GetLinksOk() (*VersionModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VersionsByProjectModelAllOfEmbeddedElements) SetLinks(v VersionModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VersionsByProjectModelAllOfEmbeddedElements) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


