# TimeEntryActivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | Time entry id | 
**Name** | **string** | The human readable name chosen for this activity | 
**Position** | **int64** | The rank the activity has in a list of activities | 
**Default** | **bool** | Flag to signal whether this activity is the default activity | 
**Embedded** | [**TimeEntryActivityModelEmbedded**](TimeEntryActivityModelEmbedded.md) |  | 
**Links** | [**TimeEntryActivityModelLinks**](TimeEntryActivityModelLinks.md) |  | 

## Methods

### NewTimeEntryActivityModel

`func NewTimeEntryActivityModel(type_ string, id int64, name string, position int64, default_ bool, embedded TimeEntryActivityModelEmbedded, links TimeEntryActivityModelLinks, ) *TimeEntryActivityModel`

NewTimeEntryActivityModel instantiates a new TimeEntryActivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryActivityModelWithDefaults

`func NewTimeEntryActivityModelWithDefaults() *TimeEntryActivityModel`

NewTimeEntryActivityModelWithDefaults instantiates a new TimeEntryActivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeEntryActivityModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeEntryActivityModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeEntryActivityModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TimeEntryActivityModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeEntryActivityModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeEntryActivityModel) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *TimeEntryActivityModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeEntryActivityModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeEntryActivityModel) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *TimeEntryActivityModel) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TimeEntryActivityModel) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TimeEntryActivityModel) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetDefault

`func (o *TimeEntryActivityModel) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TimeEntryActivityModel) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TimeEntryActivityModel) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEmbedded

`func (o *TimeEntryActivityModel) GetEmbedded() TimeEntryActivityModelEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *TimeEntryActivityModel) GetEmbeddedOk() (*TimeEntryActivityModelEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *TimeEntryActivityModel) SetEmbedded(v TimeEntryActivityModelEmbedded)`

SetEmbedded sets Embedded field to given value.


### GetLinks

`func (o *TimeEntryActivityModel) GetLinks() TimeEntryActivityModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TimeEntryActivityModel) GetLinksOk() (*TimeEntryActivityModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TimeEntryActivityModel) SetLinks(v TimeEntryActivityModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


