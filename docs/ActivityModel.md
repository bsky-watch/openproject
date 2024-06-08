# ActivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Activity id | [optional] [readonly] 
**Version** | Pointer to **int64** | Activity version | [optional] [readonly] 
**Comment** | Pointer to [**ActivityModelComment**](ActivityModelComment.md) |  | [optional] 
**Details** | Pointer to [**[]Formattable**](Formattable.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time of creation | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Time of update | [optional] [readonly] 

## Methods

### NewActivityModel

`func NewActivityModel() *ActivityModel`

NewActivityModel instantiates a new ActivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityModelWithDefaults

`func NewActivityModelWithDefaults() *ActivityModel`

NewActivityModelWithDefaults instantiates a new ActivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *ActivityModel) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActivityModel) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActivityModel) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActivityModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetComment

`func (o *ActivityModel) GetComment() ActivityModelComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ActivityModel) GetCommentOk() (*ActivityModelComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ActivityModel) SetComment(v ActivityModelComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ActivityModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDetails

`func (o *ActivityModel) GetDetails() []Formattable`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ActivityModel) GetDetailsOk() (*[]Formattable, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ActivityModel) SetDetails(v []Formattable)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ActivityModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ActivityModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActivityModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActivityModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActivityModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ActivityModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ActivityModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ActivityModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ActivityModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


