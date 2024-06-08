# MembershipWriteModelMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationMessage** | Pointer to [**MembershipWriteModelMetaNotificationMessage**](MembershipWriteModelMetaNotificationMessage.md) |  | [optional] 
**SendNotification** | Pointer to **bool** | Set to false, if no notification should get sent. | [optional] [default to true]

## Methods

### NewMembershipWriteModelMeta

`func NewMembershipWriteModelMeta() *MembershipWriteModelMeta`

NewMembershipWriteModelMeta instantiates a new MembershipWriteModelMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipWriteModelMetaWithDefaults

`func NewMembershipWriteModelMetaWithDefaults() *MembershipWriteModelMeta`

NewMembershipWriteModelMetaWithDefaults instantiates a new MembershipWriteModelMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationMessage

`func (o *MembershipWriteModelMeta) GetNotificationMessage() MembershipWriteModelMetaNotificationMessage`

GetNotificationMessage returns the NotificationMessage field if non-nil, zero value otherwise.

### GetNotificationMessageOk

`func (o *MembershipWriteModelMeta) GetNotificationMessageOk() (*MembershipWriteModelMetaNotificationMessage, bool)`

GetNotificationMessageOk returns a tuple with the NotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessage

`func (o *MembershipWriteModelMeta) SetNotificationMessage(v MembershipWriteModelMetaNotificationMessage)`

SetNotificationMessage sets NotificationMessage field to given value.

### HasNotificationMessage

`func (o *MembershipWriteModelMeta) HasNotificationMessage() bool`

HasNotificationMessage returns a boolean if a field has been set.

### GetSendNotification

`func (o *MembershipWriteModelMeta) GetSendNotification() bool`

GetSendNotification returns the SendNotification field if non-nil, zero value otherwise.

### GetSendNotificationOk

`func (o *MembershipWriteModelMeta) GetSendNotificationOk() (*bool, bool)`

GetSendNotificationOk returns a tuple with the SendNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotification

`func (o *MembershipWriteModelMeta) SetSendNotification(v bool)`

SetSendNotification sets SendNotification field to given value.

### HasSendNotification

`func (o *MembershipWriteModelMeta) HasSendNotification() bool`

HasSendNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


