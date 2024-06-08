# CustomOptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The identifier | [optional] [readonly] 
**Value** | Pointer to **string** | The value defined for this custom option | [optional] [readonly] 
**Links** | Pointer to [**CustomOptionModelLinks**](CustomOptionModelLinks.md) |  | [optional] 

## Methods

### NewCustomOptionModel

`func NewCustomOptionModel() *CustomOptionModel`

NewCustomOptionModel instantiates a new CustomOptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomOptionModelWithDefaults

`func NewCustomOptionModelWithDefaults() *CustomOptionModel`

NewCustomOptionModelWithDefaults instantiates a new CustomOptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomOptionModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomOptionModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomOptionModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CustomOptionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *CustomOptionModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomOptionModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomOptionModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomOptionModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLinks

`func (o *CustomOptionModel) GetLinks() CustomOptionModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomOptionModel) GetLinksOk() (*CustomOptionModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomOptionModel) SetLinks(v CustomOptionModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomOptionModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


