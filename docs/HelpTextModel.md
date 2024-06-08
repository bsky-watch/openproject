# HelpTextModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** |  | 
**Attribute** | **string** | The attribute the help text is assigned to. | 
**Scope** | **string** |  | 
**HelpText** | [**Formattable**](Formattable.md) |  | 
**Links** | [**HelpTextModelLinks**](HelpTextModelLinks.md) |  | 

## Methods

### NewHelpTextModel

`func NewHelpTextModel(type_ string, id int64, attribute string, scope string, helpText Formattable, links HelpTextModelLinks, ) *HelpTextModel`

NewHelpTextModel instantiates a new HelpTextModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpTextModelWithDefaults

`func NewHelpTextModelWithDefaults() *HelpTextModel`

NewHelpTextModelWithDefaults instantiates a new HelpTextModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HelpTextModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HelpTextModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HelpTextModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *HelpTextModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelpTextModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelpTextModel) SetId(v int64)`

SetId sets Id field to given value.


### GetAttribute

`func (o *HelpTextModel) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *HelpTextModel) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *HelpTextModel) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetScope

`func (o *HelpTextModel) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HelpTextModel) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HelpTextModel) SetScope(v string)`

SetScope sets Scope field to given value.


### GetHelpText

`func (o *HelpTextModel) GetHelpText() Formattable`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *HelpTextModel) GetHelpTextOk() (*Formattable, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *HelpTextModel) SetHelpText(v Formattable)`

SetHelpText sets HelpText field to given value.


### GetLinks

`func (o *HelpTextModel) GetLinks() HelpTextModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HelpTextModel) GetLinksOk() (*HelpTextModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HelpTextModel) SetLinks(v HelpTextModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


