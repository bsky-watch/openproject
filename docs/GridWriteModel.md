# GridWriteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowCount** | Pointer to **int64** | The number of rows the grid has | [optional] 
**ColumnCount** | Pointer to **int64** | The number of columns the grid has | [optional] 
**Widgets** | Pointer to [**[]GridWidgetModel**](GridWidgetModel.md) | The set of &#x60;GridWidget&#x60;s selected for the grid.  # Conditions  - The widgets must not overlap. | [optional] 
**Links** | Pointer to [**GridWriteModelLinks**](GridWriteModelLinks.md) |  | [optional] 

## Methods

### NewGridWriteModel

`func NewGridWriteModel() *GridWriteModel`

NewGridWriteModel instantiates a new GridWriteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridWriteModelWithDefaults

`func NewGridWriteModelWithDefaults() *GridWriteModel`

NewGridWriteModelWithDefaults instantiates a new GridWriteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowCount

`func (o *GridWriteModel) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *GridWriteModel) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *GridWriteModel) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *GridWriteModel) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetColumnCount

`func (o *GridWriteModel) GetColumnCount() int64`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *GridWriteModel) GetColumnCountOk() (*int64, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *GridWriteModel) SetColumnCount(v int64)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *GridWriteModel) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetWidgets

`func (o *GridWriteModel) GetWidgets() []GridWidgetModel`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *GridWriteModel) GetWidgetsOk() (*[]GridWidgetModel, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *GridWriteModel) SetWidgets(v []GridWidgetModel)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *GridWriteModel) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### GetLinks

`func (o *GridWriteModel) GetLinks() GridWriteModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GridWriteModel) GetLinksOk() (*GridWriteModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GridWriteModel) SetLinks(v GridWriteModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GridWriteModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


