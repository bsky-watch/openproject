# NotificationModelDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The key of the key - value pair represented by the Values::Property | 
**Value** | **string** | The value of the key - value pair represented by the Values::Property | 
**Links** | [**ValuesPropertyModelLinks**](ValuesPropertyModelLinks.md) |  | 

## Methods

### NewNotificationModelDetailsInner

`func NewNotificationModelDetailsInner(type_ string, value string, links ValuesPropertyModelLinks, ) *NotificationModelDetailsInner`

NewNotificationModelDetailsInner instantiates a new NotificationModelDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelDetailsInnerWithDefaults

`func NewNotificationModelDetailsInnerWithDefaults() *NotificationModelDetailsInner`

NewNotificationModelDetailsInnerWithDefaults instantiates a new NotificationModelDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationModelDetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationModelDetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationModelDetailsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *NotificationModelDetailsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotificationModelDetailsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotificationModelDetailsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetLinks

`func (o *NotificationModelDetailsInner) GetLinks() ValuesPropertyModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationModelDetailsInner) GetLinksOk() (*ValuesPropertyModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationModelDetailsInner) SetLinks(v ValuesPropertyModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


