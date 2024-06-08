# TimeEntryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the time entry | [optional] 
**Comment** | Pointer to [**TimeEntryModelComment**](TimeEntryModelComment.md) |  | [optional] 
**SpentOn** | Pointer to **string** | The date the expenditure is booked for | [optional] 
**Hours** | Pointer to **string** | The time quantifying the expenditure | [optional] 
**Ongoing** | Pointer to **bool** | Whether the time entry is actively tracking time | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the time entry was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the time entry was last updated | [optional] 
**Links** | Pointer to [**TimeEntryModelLinks**](TimeEntryModelLinks.md) |  | [optional] 

## Methods

### NewTimeEntryModel

`func NewTimeEntryModel() *TimeEntryModel`

NewTimeEntryModel instantiates a new TimeEntryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryModelWithDefaults

`func NewTimeEntryModelWithDefaults() *TimeEntryModel`

NewTimeEntryModelWithDefaults instantiates a new TimeEntryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeEntryModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeEntryModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeEntryModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TimeEntryModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComment

`func (o *TimeEntryModel) GetComment() TimeEntryModelComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TimeEntryModel) GetCommentOk() (*TimeEntryModelComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TimeEntryModel) SetComment(v TimeEntryModelComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TimeEntryModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSpentOn

`func (o *TimeEntryModel) GetSpentOn() string`

GetSpentOn returns the SpentOn field if non-nil, zero value otherwise.

### GetSpentOnOk

`func (o *TimeEntryModel) GetSpentOnOk() (*string, bool)`

GetSpentOnOk returns a tuple with the SpentOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentOn

`func (o *TimeEntryModel) SetSpentOn(v string)`

SetSpentOn sets SpentOn field to given value.

### HasSpentOn

`func (o *TimeEntryModel) HasSpentOn() bool`

HasSpentOn returns a boolean if a field has been set.

### GetHours

`func (o *TimeEntryModel) GetHours() string`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimeEntryModel) GetHoursOk() (*string, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimeEntryModel) SetHours(v string)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TimeEntryModel) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetOngoing

`func (o *TimeEntryModel) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *TimeEntryModel) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *TimeEntryModel) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *TimeEntryModel) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TimeEntryModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimeEntryModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimeEntryModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TimeEntryModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TimeEntryModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimeEntryModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimeEntryModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TimeEntryModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *TimeEntryModel) GetLinks() TimeEntryModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TimeEntryModel) GetLinksOk() (*TimeEntryModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TimeEntryModel) SetLinks(v TimeEntryModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TimeEntryModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


