# GridWidgetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | The grid widget&#39;s unique identifier. Can be null, if a new widget is created within a grid. | 
**Identifier** | **string** | An alternative, human legible, and unique identifier. | 
**StartRow** | **int64** | The index of the starting row of the widget. The row is inclusive. | 
**EndRow** | **int64** | The index of the ending row of the widget. The row is exclusive. | 
**StartColumn** | **int64** | The index of the starting column of the widget. The column is inclusive. | 
**EndColumn** | **int64** | The index of the ending column of the widget. The column is exclusive. | 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGridWidgetModel

`func NewGridWidgetModel(type_ string, id int64, identifier string, startRow int64, endRow int64, startColumn int64, endColumn int64, ) *GridWidgetModel`

NewGridWidgetModel instantiates a new GridWidgetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridWidgetModelWithDefaults

`func NewGridWidgetModelWithDefaults() *GridWidgetModel`

NewGridWidgetModelWithDefaults instantiates a new GridWidgetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GridWidgetModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridWidgetModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridWidgetModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GridWidgetModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GridWidgetModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GridWidgetModel) SetId(v int64)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *GridWidgetModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GridWidgetModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GridWidgetModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetStartRow

`func (o *GridWidgetModel) GetStartRow() int64`

GetStartRow returns the StartRow field if non-nil, zero value otherwise.

### GetStartRowOk

`func (o *GridWidgetModel) GetStartRowOk() (*int64, bool)`

GetStartRowOk returns a tuple with the StartRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRow

`func (o *GridWidgetModel) SetStartRow(v int64)`

SetStartRow sets StartRow field to given value.


### GetEndRow

`func (o *GridWidgetModel) GetEndRow() int64`

GetEndRow returns the EndRow field if non-nil, zero value otherwise.

### GetEndRowOk

`func (o *GridWidgetModel) GetEndRowOk() (*int64, bool)`

GetEndRowOk returns a tuple with the EndRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRow

`func (o *GridWidgetModel) SetEndRow(v int64)`

SetEndRow sets EndRow field to given value.


### GetStartColumn

`func (o *GridWidgetModel) GetStartColumn() int64`

GetStartColumn returns the StartColumn field if non-nil, zero value otherwise.

### GetStartColumnOk

`func (o *GridWidgetModel) GetStartColumnOk() (*int64, bool)`

GetStartColumnOk returns a tuple with the StartColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartColumn

`func (o *GridWidgetModel) SetStartColumn(v int64)`

SetStartColumn sets StartColumn field to given value.


### GetEndColumn

`func (o *GridWidgetModel) GetEndColumn() int64`

GetEndColumn returns the EndColumn field if non-nil, zero value otherwise.

### GetEndColumnOk

`func (o *GridWidgetModel) GetEndColumnOk() (*int64, bool)`

GetEndColumnOk returns a tuple with the EndColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndColumn

`func (o *GridWidgetModel) SetEndColumn(v int64)`

SetEndColumn sets EndColumn field to given value.


### GetOptions

`func (o *GridWidgetModel) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GridWidgetModel) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GridWidgetModel) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GridWidgetModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


