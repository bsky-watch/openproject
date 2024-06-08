# GridReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int64** | Grid&#39;s id | 
**RowCount** | **int64** | The number of rows the grid has | 
**ColumnCount** | **int64** | The number of columns the grid has | 
**Widgets** | [**[]GridWidgetModel**](GridWidgetModel.md) | The set of &#x60;GridWidget&#x60;s selected for the grid.  # Conditions  - The widgets must not overlap. | 
**CreatedAt** | Pointer to **time.Time** | The time the grid was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the grid was last updated. | [optional] 
**Links** | [**GridReadModelLinks**](GridReadModelLinks.md) |  | 

## Methods

### NewGridReadModel

`func NewGridReadModel(type_ string, id int64, rowCount int64, columnCount int64, widgets []GridWidgetModel, links GridReadModelLinks, ) *GridReadModel`

NewGridReadModel instantiates a new GridReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridReadModelWithDefaults

`func NewGridReadModelWithDefaults() *GridReadModel`

NewGridReadModelWithDefaults instantiates a new GridReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GridReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GridReadModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GridReadModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GridReadModel) SetId(v int64)`

SetId sets Id field to given value.


### GetRowCount

`func (o *GridReadModel) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *GridReadModel) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *GridReadModel) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.


### GetColumnCount

`func (o *GridReadModel) GetColumnCount() int64`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *GridReadModel) GetColumnCountOk() (*int64, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *GridReadModel) SetColumnCount(v int64)`

SetColumnCount sets ColumnCount field to given value.


### GetWidgets

`func (o *GridReadModel) GetWidgets() []GridWidgetModel`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *GridReadModel) GetWidgetsOk() (*[]GridWidgetModel, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *GridReadModel) SetWidgets(v []GridWidgetModel)`

SetWidgets sets Widgets field to given value.


### GetCreatedAt

`func (o *GridReadModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GridReadModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GridReadModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GridReadModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GridReadModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GridReadModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GridReadModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GridReadModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *GridReadModel) GetLinks() GridReadModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GridReadModel) GetLinksOk() (*GridReadModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GridReadModel) SetLinks(v GridReadModelLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


