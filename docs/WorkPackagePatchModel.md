# WorkPackagePatchModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockVersion** | **int64** | The version of the item as used for optimistic locking | 
**Subject** | Pointer to **string** | Work package subject | [optional] 
**Description** | Pointer to [**WorkPackageModelDescription**](WorkPackageModelDescription.md) |  | [optional] 
**ScheduleManually** | Pointer to **bool** | If false (default) schedule automatically. | [optional] 
**StartDate** | Pointer to **string** | Scheduled beginning of a work package | [optional] 
**DueDate** | Pointer to **string** | Scheduled end of a work package | [optional] 
**Date** | Pointer to **string** | Date on which a milestone is achieved | [optional] 
**EstimatedTime** | Pointer to **string** | Time a work package likely needs to be completed excluding its descendants | [optional] 
**IgnoreNonWorkingDays** | Pointer to **bool** | **(NOT IMPLEMENTED)** When scheduling, whether or not to ignore the non working days being defined. A work package with the flag set to true will be allowed to be scheduled to a non working day. | [optional] [readonly] 
**SpentTime** | Pointer to **string** | The time booked for this work package by users working on it  # Conditions  **Permission** view time entries | [optional] [readonly] 
**PercentageDone** | Pointer to **int64** | Amount of total completion for a work package | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the work package | [optional] [readonly] 
**Links** | Pointer to [**WorkPackagePatchModelLinks**](WorkPackagePatchModelLinks.md) |  | [optional] 

## Methods

### NewWorkPackagePatchModel

`func NewWorkPackagePatchModel(lockVersion int64, ) *WorkPackagePatchModel`

NewWorkPackagePatchModel instantiates a new WorkPackagePatchModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkPackagePatchModelWithDefaults

`func NewWorkPackagePatchModelWithDefaults() *WorkPackagePatchModel`

NewWorkPackagePatchModelWithDefaults instantiates a new WorkPackagePatchModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockVersion

`func (o *WorkPackagePatchModel) GetLockVersion() int64`

GetLockVersion returns the LockVersion field if non-nil, zero value otherwise.

### GetLockVersionOk

`func (o *WorkPackagePatchModel) GetLockVersionOk() (*int64, bool)`

GetLockVersionOk returns a tuple with the LockVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockVersion

`func (o *WorkPackagePatchModel) SetLockVersion(v int64)`

SetLockVersion sets LockVersion field to given value.


### GetSubject

`func (o *WorkPackagePatchModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WorkPackagePatchModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WorkPackagePatchModel) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WorkPackagePatchModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDescription

`func (o *WorkPackagePatchModel) GetDescription() WorkPackageModelDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkPackagePatchModel) GetDescriptionOk() (*WorkPackageModelDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkPackagePatchModel) SetDescription(v WorkPackageModelDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkPackagePatchModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleManually

`func (o *WorkPackagePatchModel) GetScheduleManually() bool`

GetScheduleManually returns the ScheduleManually field if non-nil, zero value otherwise.

### GetScheduleManuallyOk

`func (o *WorkPackagePatchModel) GetScheduleManuallyOk() (*bool, bool)`

GetScheduleManuallyOk returns a tuple with the ScheduleManually field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleManually

`func (o *WorkPackagePatchModel) SetScheduleManually(v bool)`

SetScheduleManually sets ScheduleManually field to given value.

### HasScheduleManually

`func (o *WorkPackagePatchModel) HasScheduleManually() bool`

HasScheduleManually returns a boolean if a field has been set.

### GetStartDate

`func (o *WorkPackagePatchModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WorkPackagePatchModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WorkPackagePatchModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WorkPackagePatchModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *WorkPackagePatchModel) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *WorkPackagePatchModel) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *WorkPackagePatchModel) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *WorkPackagePatchModel) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDate

`func (o *WorkPackagePatchModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WorkPackagePatchModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WorkPackagePatchModel) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *WorkPackagePatchModel) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEstimatedTime

`func (o *WorkPackagePatchModel) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *WorkPackagePatchModel) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *WorkPackagePatchModel) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *WorkPackagePatchModel) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### GetIgnoreNonWorkingDays

`func (o *WorkPackagePatchModel) GetIgnoreNonWorkingDays() bool`

GetIgnoreNonWorkingDays returns the IgnoreNonWorkingDays field if non-nil, zero value otherwise.

### GetIgnoreNonWorkingDaysOk

`func (o *WorkPackagePatchModel) GetIgnoreNonWorkingDaysOk() (*bool, bool)`

GetIgnoreNonWorkingDaysOk returns a tuple with the IgnoreNonWorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreNonWorkingDays

`func (o *WorkPackagePatchModel) SetIgnoreNonWorkingDays(v bool)`

SetIgnoreNonWorkingDays sets IgnoreNonWorkingDays field to given value.

### HasIgnoreNonWorkingDays

`func (o *WorkPackagePatchModel) HasIgnoreNonWorkingDays() bool`

HasIgnoreNonWorkingDays returns a boolean if a field has been set.

### GetSpentTime

`func (o *WorkPackagePatchModel) GetSpentTime() string`

GetSpentTime returns the SpentTime field if non-nil, zero value otherwise.

### GetSpentTimeOk

`func (o *WorkPackagePatchModel) GetSpentTimeOk() (*string, bool)`

GetSpentTimeOk returns a tuple with the SpentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentTime

`func (o *WorkPackagePatchModel) SetSpentTime(v string)`

SetSpentTime sets SpentTime field to given value.

### HasSpentTime

`func (o *WorkPackagePatchModel) HasSpentTime() bool`

HasSpentTime returns a boolean if a field has been set.

### GetPercentageDone

`func (o *WorkPackagePatchModel) GetPercentageDone() int64`

GetPercentageDone returns the PercentageDone field if non-nil, zero value otherwise.

### GetPercentageDoneOk

`func (o *WorkPackagePatchModel) GetPercentageDoneOk() (*int64, bool)`

GetPercentageDoneOk returns a tuple with the PercentageDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageDone

`func (o *WorkPackagePatchModel) SetPercentageDone(v int64)`

SetPercentageDone sets PercentageDone field to given value.

### HasPercentageDone

`func (o *WorkPackagePatchModel) HasPercentageDone() bool`

HasPercentageDone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkPackagePatchModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkPackagePatchModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkPackagePatchModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkPackagePatchModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkPackagePatchModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkPackagePatchModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkPackagePatchModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkPackagePatchModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *WorkPackagePatchModel) GetLinks() WorkPackagePatchModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkPackagePatchModel) GetLinksOk() (*WorkPackagePatchModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkPackagePatchModel) SetLinks(v WorkPackagePatchModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WorkPackagePatchModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


