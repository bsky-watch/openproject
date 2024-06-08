# WorkPackageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Work package id | [optional] [readonly] 
**LockVersion** | Pointer to **int64** | The version of the item as used for optimistic locking | [optional] [readonly] 
**Subject** | Pointer to **string** | Work package subject | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to [**WorkPackageModelDescription**](WorkPackageModelDescription.md) |  | [optional] 
**ScheduleManually** | Pointer to **bool** | If false (default) schedule automatically. | [optional] 
**Readonly** | Pointer to **bool** | If true, the work package is in a readonly status so with the exception of the status, no other property can be altered. | [optional] 
**StartDate** | Pointer to **string** | Scheduled beginning of a work package | [optional] 
**DueDate** | Pointer to **string** | Scheduled end of a work package | [optional] 
**Date** | Pointer to **string** | Date on which a milestone is achieved | [optional] 
**DerivedStartDate** | Pointer to **string** | Similar to start date but is not set by a client but rather deduced by the work packages&#39; descendants. If manual scheduleManually is active, the two dates can deviate. | [optional] [readonly] 
**DerivedDueDate** | Pointer to **string** | Similar to due date but is not set by a client but rather deduced by the work packages&#39; descendants. If manual scheduleManually is active, the two dates can deviate. | [optional] [readonly] 
**Duration** | Pointer to **string** | **(NOT IMPLEMENTED)** The amount of time in hours the work package needs to be completed. Not available for milestone type of work packages. | [optional] [readonly] 
**EstimatedTime** | Pointer to **string** | Time a work package likely needs to be completed excluding its descendants | [optional] 
**DerivedEstimatedTime** | Pointer to **string** | Time a work package likely needs to be completed including its descendants | [optional] [readonly] 
**IgnoreNonWorkingDays** | Pointer to **bool** | **(NOT IMPLEMENTED)** When scheduling, whether or not to ignore the non working days being defined. A work package with the flag set to true will be allowed to be scheduled to a non working day. | [optional] [readonly] 
**SpentTime** | Pointer to **string** | The time booked for this work package by users working on it  # Conditions  **Permission** view time entries | [optional] [readonly] 
**PercentageDone** | Pointer to **int64** | Amount of total completion for a work package | [optional] 
**DerivedPercentageDone** | Pointer to **int64** | Amount of total completion for a work package derived from itself and its descendant work packages | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Time of creation. Can be writable by admins with the &#x60;apiv3_write_readonly_attributes&#x60; setting enabled. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Time of the most recent change to the work package. | [optional] [readonly] 
**Links** | Pointer to [**WorkPackageModelLinks**](WorkPackageModelLinks.md) |  | [optional] 

## Methods

### NewWorkPackageModel

`func NewWorkPackageModel() *WorkPackageModel`

NewWorkPackageModel instantiates a new WorkPackageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkPackageModelWithDefaults

`func NewWorkPackageModelWithDefaults() *WorkPackageModel`

NewWorkPackageModelWithDefaults instantiates a new WorkPackageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkPackageModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkPackageModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkPackageModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WorkPackageModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLockVersion

`func (o *WorkPackageModel) GetLockVersion() int64`

GetLockVersion returns the LockVersion field if non-nil, zero value otherwise.

### GetLockVersionOk

`func (o *WorkPackageModel) GetLockVersionOk() (*int64, bool)`

GetLockVersionOk returns a tuple with the LockVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockVersion

`func (o *WorkPackageModel) SetLockVersion(v int64)`

SetLockVersion sets LockVersion field to given value.

### HasLockVersion

`func (o *WorkPackageModel) HasLockVersion() bool`

HasLockVersion returns a boolean if a field has been set.

### GetSubject

`func (o *WorkPackageModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WorkPackageModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WorkPackageModel) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WorkPackageModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetType

`func (o *WorkPackageModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkPackageModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkPackageModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkPackageModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *WorkPackageModel) GetDescription() WorkPackageModelDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkPackageModel) GetDescriptionOk() (*WorkPackageModelDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkPackageModel) SetDescription(v WorkPackageModelDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkPackageModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleManually

`func (o *WorkPackageModel) GetScheduleManually() bool`

GetScheduleManually returns the ScheduleManually field if non-nil, zero value otherwise.

### GetScheduleManuallyOk

`func (o *WorkPackageModel) GetScheduleManuallyOk() (*bool, bool)`

GetScheduleManuallyOk returns a tuple with the ScheduleManually field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleManually

`func (o *WorkPackageModel) SetScheduleManually(v bool)`

SetScheduleManually sets ScheduleManually field to given value.

### HasScheduleManually

`func (o *WorkPackageModel) HasScheduleManually() bool`

HasScheduleManually returns a boolean if a field has been set.

### GetReadonly

`func (o *WorkPackageModel) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *WorkPackageModel) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *WorkPackageModel) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *WorkPackageModel) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetStartDate

`func (o *WorkPackageModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WorkPackageModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WorkPackageModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WorkPackageModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDueDate

`func (o *WorkPackageModel) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *WorkPackageModel) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *WorkPackageModel) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *WorkPackageModel) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDate

`func (o *WorkPackageModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WorkPackageModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WorkPackageModel) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *WorkPackageModel) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDerivedStartDate

`func (o *WorkPackageModel) GetDerivedStartDate() string`

GetDerivedStartDate returns the DerivedStartDate field if non-nil, zero value otherwise.

### GetDerivedStartDateOk

`func (o *WorkPackageModel) GetDerivedStartDateOk() (*string, bool)`

GetDerivedStartDateOk returns a tuple with the DerivedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedStartDate

`func (o *WorkPackageModel) SetDerivedStartDate(v string)`

SetDerivedStartDate sets DerivedStartDate field to given value.

### HasDerivedStartDate

`func (o *WorkPackageModel) HasDerivedStartDate() bool`

HasDerivedStartDate returns a boolean if a field has been set.

### GetDerivedDueDate

`func (o *WorkPackageModel) GetDerivedDueDate() string`

GetDerivedDueDate returns the DerivedDueDate field if non-nil, zero value otherwise.

### GetDerivedDueDateOk

`func (o *WorkPackageModel) GetDerivedDueDateOk() (*string, bool)`

GetDerivedDueDateOk returns a tuple with the DerivedDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedDueDate

`func (o *WorkPackageModel) SetDerivedDueDate(v string)`

SetDerivedDueDate sets DerivedDueDate field to given value.

### HasDerivedDueDate

`func (o *WorkPackageModel) HasDerivedDueDate() bool`

HasDerivedDueDate returns a boolean if a field has been set.

### GetDuration

`func (o *WorkPackageModel) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkPackageModel) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkPackageModel) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkPackageModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEstimatedTime

`func (o *WorkPackageModel) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *WorkPackageModel) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *WorkPackageModel) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *WorkPackageModel) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### GetDerivedEstimatedTime

`func (o *WorkPackageModel) GetDerivedEstimatedTime() string`

GetDerivedEstimatedTime returns the DerivedEstimatedTime field if non-nil, zero value otherwise.

### GetDerivedEstimatedTimeOk

`func (o *WorkPackageModel) GetDerivedEstimatedTimeOk() (*string, bool)`

GetDerivedEstimatedTimeOk returns a tuple with the DerivedEstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedEstimatedTime

`func (o *WorkPackageModel) SetDerivedEstimatedTime(v string)`

SetDerivedEstimatedTime sets DerivedEstimatedTime field to given value.

### HasDerivedEstimatedTime

`func (o *WorkPackageModel) HasDerivedEstimatedTime() bool`

HasDerivedEstimatedTime returns a boolean if a field has been set.

### GetIgnoreNonWorkingDays

`func (o *WorkPackageModel) GetIgnoreNonWorkingDays() bool`

GetIgnoreNonWorkingDays returns the IgnoreNonWorkingDays field if non-nil, zero value otherwise.

### GetIgnoreNonWorkingDaysOk

`func (o *WorkPackageModel) GetIgnoreNonWorkingDaysOk() (*bool, bool)`

GetIgnoreNonWorkingDaysOk returns a tuple with the IgnoreNonWorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreNonWorkingDays

`func (o *WorkPackageModel) SetIgnoreNonWorkingDays(v bool)`

SetIgnoreNonWorkingDays sets IgnoreNonWorkingDays field to given value.

### HasIgnoreNonWorkingDays

`func (o *WorkPackageModel) HasIgnoreNonWorkingDays() bool`

HasIgnoreNonWorkingDays returns a boolean if a field has been set.

### GetSpentTime

`func (o *WorkPackageModel) GetSpentTime() string`

GetSpentTime returns the SpentTime field if non-nil, zero value otherwise.

### GetSpentTimeOk

`func (o *WorkPackageModel) GetSpentTimeOk() (*string, bool)`

GetSpentTimeOk returns a tuple with the SpentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentTime

`func (o *WorkPackageModel) SetSpentTime(v string)`

SetSpentTime sets SpentTime field to given value.

### HasSpentTime

`func (o *WorkPackageModel) HasSpentTime() bool`

HasSpentTime returns a boolean if a field has been set.

### GetPercentageDone

`func (o *WorkPackageModel) GetPercentageDone() int64`

GetPercentageDone returns the PercentageDone field if non-nil, zero value otherwise.

### GetPercentageDoneOk

`func (o *WorkPackageModel) GetPercentageDoneOk() (*int64, bool)`

GetPercentageDoneOk returns a tuple with the PercentageDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageDone

`func (o *WorkPackageModel) SetPercentageDone(v int64)`

SetPercentageDone sets PercentageDone field to given value.

### HasPercentageDone

`func (o *WorkPackageModel) HasPercentageDone() bool`

HasPercentageDone returns a boolean if a field has been set.

### GetDerivedPercentageDone

`func (o *WorkPackageModel) GetDerivedPercentageDone() int64`

GetDerivedPercentageDone returns the DerivedPercentageDone field if non-nil, zero value otherwise.

### GetDerivedPercentageDoneOk

`func (o *WorkPackageModel) GetDerivedPercentageDoneOk() (*int64, bool)`

GetDerivedPercentageDoneOk returns a tuple with the DerivedPercentageDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedPercentageDone

`func (o *WorkPackageModel) SetDerivedPercentageDone(v int64)`

SetDerivedPercentageDone sets DerivedPercentageDone field to given value.

### HasDerivedPercentageDone

`func (o *WorkPackageModel) HasDerivedPercentageDone() bool`

HasDerivedPercentageDone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkPackageModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkPackageModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkPackageModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkPackageModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkPackageModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkPackageModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkPackageModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkPackageModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *WorkPackageModel) GetLinks() WorkPackageModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkPackageModel) GetLinksOk() (*WorkPackageModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkPackageModel) SetLinks(v WorkPackageModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WorkPackageModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


