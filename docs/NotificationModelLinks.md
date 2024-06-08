# NotificationModelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**NotificationModelLinksSelf**](NotificationModelLinksSelf.md) |  | 
**ReadIAN** | Pointer to [**NotificationModelLinksReadIAN**](NotificationModelLinksReadIAN.md) |  | [optional] 
**UnreadIAN** | Pointer to [**NotificationModelLinksUnreadIAN**](NotificationModelLinksUnreadIAN.md) |  | [optional] 
**Project** | [**NotificationModelLinksProject**](NotificationModelLinksProject.md) |  | 
**Actor** | [**NotificationModelLinksActor**](NotificationModelLinksActor.md) |  | 
**Resource** | [**NotificationModelLinksResource**](NotificationModelLinksResource.md) |  | 
**Activity** | [**NotificationModelLinksActivity**](NotificationModelLinksActivity.md) |  | 
**Details** | [**[]NotificationModelLinksDetailsInner**](NotificationModelLinksDetailsInner.md) |  | 

## Methods

### NewNotificationModelLinks

`func NewNotificationModelLinks(self NotificationModelLinksSelf, project NotificationModelLinksProject, actor NotificationModelLinksActor, resource NotificationModelLinksResource, activity NotificationModelLinksActivity, details []NotificationModelLinksDetailsInner, ) *NotificationModelLinks`

NewNotificationModelLinks instantiates a new NotificationModelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelLinksWithDefaults

`func NewNotificationModelLinksWithDefaults() *NotificationModelLinks`

NewNotificationModelLinksWithDefaults instantiates a new NotificationModelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NotificationModelLinks) GetSelf() NotificationModelLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NotificationModelLinks) GetSelfOk() (*NotificationModelLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NotificationModelLinks) SetSelf(v NotificationModelLinksSelf)`

SetSelf sets Self field to given value.


### GetReadIAN

`func (o *NotificationModelLinks) GetReadIAN() NotificationModelLinksReadIAN`

GetReadIAN returns the ReadIAN field if non-nil, zero value otherwise.

### GetReadIANOk

`func (o *NotificationModelLinks) GetReadIANOk() (*NotificationModelLinksReadIAN, bool)`

GetReadIANOk returns a tuple with the ReadIAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIAN

`func (o *NotificationModelLinks) SetReadIAN(v NotificationModelLinksReadIAN)`

SetReadIAN sets ReadIAN field to given value.

### HasReadIAN

`func (o *NotificationModelLinks) HasReadIAN() bool`

HasReadIAN returns a boolean if a field has been set.

### GetUnreadIAN

`func (o *NotificationModelLinks) GetUnreadIAN() NotificationModelLinksUnreadIAN`

GetUnreadIAN returns the UnreadIAN field if non-nil, zero value otherwise.

### GetUnreadIANOk

`func (o *NotificationModelLinks) GetUnreadIANOk() (*NotificationModelLinksUnreadIAN, bool)`

GetUnreadIANOk returns a tuple with the UnreadIAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadIAN

`func (o *NotificationModelLinks) SetUnreadIAN(v NotificationModelLinksUnreadIAN)`

SetUnreadIAN sets UnreadIAN field to given value.

### HasUnreadIAN

`func (o *NotificationModelLinks) HasUnreadIAN() bool`

HasUnreadIAN returns a boolean if a field has been set.

### GetProject

`func (o *NotificationModelLinks) GetProject() NotificationModelLinksProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NotificationModelLinks) GetProjectOk() (*NotificationModelLinksProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NotificationModelLinks) SetProject(v NotificationModelLinksProject)`

SetProject sets Project field to given value.


### GetActor

`func (o *NotificationModelLinks) GetActor() NotificationModelLinksActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *NotificationModelLinks) GetActorOk() (*NotificationModelLinksActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *NotificationModelLinks) SetActor(v NotificationModelLinksActor)`

SetActor sets Actor field to given value.


### GetResource

`func (o *NotificationModelLinks) GetResource() NotificationModelLinksResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *NotificationModelLinks) GetResourceOk() (*NotificationModelLinksResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *NotificationModelLinks) SetResource(v NotificationModelLinksResource)`

SetResource sets Resource field to given value.


### GetActivity

`func (o *NotificationModelLinks) GetActivity() NotificationModelLinksActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *NotificationModelLinks) GetActivityOk() (*NotificationModelLinksActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *NotificationModelLinks) SetActivity(v NotificationModelLinksActivity)`

SetActivity sets Activity field to given value.


### GetDetails

`func (o *NotificationModelLinks) GetDetails() []NotificationModelLinksDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NotificationModelLinks) GetDetailsOk() (*[]NotificationModelLinksDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NotificationModelLinks) SetDetails(v []NotificationModelLinksDetailsInner)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


