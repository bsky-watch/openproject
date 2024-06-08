# VersionModel

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

### NewVersionModel

`func NewVersionModel(name string, status string, sharing string, createdAt time.Time, updatedAt time.Time, ) *VersionModel`

NewVersionModel instantiates a new VersionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionModelWithDefaults

`func NewVersionModelWithDefaults() *VersionModel`

NewVersionModelWithDefaults instantiates a new VersionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VersionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VersionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VersionModel) GetDescription() ActivityModelComment`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionModel) GetDescriptionOk() (*ActivityModelComment, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionModel) SetDescription(v ActivityModelComment)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *VersionModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *VersionModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *VersionModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *VersionModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *VersionModel) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *VersionModel) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *VersionModel) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *VersionModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *VersionModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersionModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersionModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSharing

`func (o *VersionModel) GetSharing() string`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *VersionModel) GetSharingOk() (*string, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *VersionModel) SetSharing(v string)`

SetSharing sets Sharing field to given value.


### GetCreatedAt

`func (o *VersionModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VersionModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VersionModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VersionModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VersionModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VersionModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *VersionModel) GetLinks() VersionModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VersionModel) GetLinksOk() (*VersionModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VersionModel) SetLinks(v VersionModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VersionModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


