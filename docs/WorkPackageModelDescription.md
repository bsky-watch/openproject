# WorkPackageModelDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | Indicates the formatting language of the raw text | [readonly] 
**Raw** | Pointer to **string** | The raw text, as entered by the user | [optional] 
**Html** | Pointer to **string** | The text converted to HTML according to the format | [optional] [readonly] 

## Methods

### NewWorkPackageModelDescription

`func NewWorkPackageModelDescription(format string, ) *WorkPackageModelDescription`

NewWorkPackageModelDescription instantiates a new WorkPackageModelDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkPackageModelDescriptionWithDefaults

`func NewWorkPackageModelDescriptionWithDefaults() *WorkPackageModelDescription`

NewWorkPackageModelDescriptionWithDefaults instantiates a new WorkPackageModelDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *WorkPackageModelDescription) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkPackageModelDescription) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkPackageModelDescription) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetRaw

`func (o *WorkPackageModelDescription) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *WorkPackageModelDescription) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *WorkPackageModelDescription) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *WorkPackageModelDescription) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetHtml

`func (o *WorkPackageModelDescription) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *WorkPackageModelDescription) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *WorkPackageModelDescription) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *WorkPackageModelDescription) HasHtml() bool`

HasHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


