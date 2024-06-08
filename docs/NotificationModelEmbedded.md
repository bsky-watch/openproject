# NotificationModelEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to [**UserModel**](UserModel.md) |  | [optional] 
**Project** | [**ProjectModel**](ProjectModel.md) |  | 
**Activity** | Pointer to [**ActivityModel**](ActivityModel.md) |  | [optional] 
**Resource** | [**NotificationModelEmbeddedResource**](NotificationModelEmbeddedResource.md) |  | 

## Methods

### NewNotificationModelEmbedded

`func NewNotificationModelEmbedded(project ProjectModel, resource NotificationModelEmbeddedResource, ) *NotificationModelEmbedded`

NewNotificationModelEmbedded instantiates a new NotificationModelEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelEmbeddedWithDefaults

`func NewNotificationModelEmbeddedWithDefaults() *NotificationModelEmbedded`

NewNotificationModelEmbeddedWithDefaults instantiates a new NotificationModelEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *NotificationModelEmbedded) GetActor() UserModel`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *NotificationModelEmbedded) GetActorOk() (*UserModel, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *NotificationModelEmbedded) SetActor(v UserModel)`

SetActor sets Actor field to given value.

### HasActor

`func (o *NotificationModelEmbedded) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetProject

`func (o *NotificationModelEmbedded) GetProject() ProjectModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NotificationModelEmbedded) GetProjectOk() (*ProjectModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NotificationModelEmbedded) SetProject(v ProjectModel)`

SetProject sets Project field to given value.


### GetActivity

`func (o *NotificationModelEmbedded) GetActivity() ActivityModel`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *NotificationModelEmbedded) GetActivityOk() (*ActivityModel, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *NotificationModelEmbedded) SetActivity(v ActivityModel)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *NotificationModelEmbedded) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetResource

`func (o *NotificationModelEmbedded) GetResource() NotificationModelEmbeddedResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *NotificationModelEmbedded) GetResourceOk() (*NotificationModelEmbeddedResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *NotificationModelEmbedded) SetResource(v NotificationModelEmbeddedResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


