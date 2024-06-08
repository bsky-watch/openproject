# NotificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** | Notification id | [optional] 
**Reason** | Pointer to **string** | The reason for the notification | [optional] 
**ReadIAN** | Pointer to **bool** | Whether the notification is marked as read | [optional] 
**Details** | Pointer to [**[]NotificationModelDetailsInner**](NotificationModelDetailsInner.md) | A list of objects including detailed information about the notification. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the notification was created at | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the notification was last updated | [optional] 
**Embedded** | Pointer to [**NotificationModelEmbedded**](NotificationModelEmbedded.md) |  | [optional] 
**Links** | Pointer to [**NotificationModelLinks**](NotificationModelLinks.md) |  | [optional] 

## Methods

### NewNotificationModel

`func NewNotificationModel() *NotificationModel`

NewNotificationModel instantiates a new NotificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelWithDefaults

`func NewNotificationModelWithDefaults() *NotificationModel`

NewNotificationModelWithDefaults instantiates a new NotificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *NotificationModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *NotificationModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotificationModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotificationModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NotificationModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReadIAN

`func (o *NotificationModel) GetReadIAN() bool`

GetReadIAN returns the ReadIAN field if non-nil, zero value otherwise.

### GetReadIANOk

`func (o *NotificationModel) GetReadIANOk() (*bool, bool)`

GetReadIANOk returns a tuple with the ReadIAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIAN

`func (o *NotificationModel) SetReadIAN(v bool)`

SetReadIAN sets ReadIAN field to given value.

### HasReadIAN

`func (o *NotificationModel) HasReadIAN() bool`

HasReadIAN returns a boolean if a field has been set.

### GetDetails

`func (o *NotificationModel) GetDetails() []NotificationModelDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NotificationModel) GetDetailsOk() (*[]NotificationModelDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NotificationModel) SetDetails(v []NotificationModelDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NotificationModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEmbedded

`func (o *NotificationModel) GetEmbedded() NotificationModelEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *NotificationModel) GetEmbeddedOk() (*NotificationModelEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *NotificationModel) SetEmbedded(v NotificationModelEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *NotificationModel) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *NotificationModel) GetLinks() NotificationModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationModel) GetLinksOk() (*NotificationModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationModel) SetLinks(v NotificationModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


