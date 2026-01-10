# TextInputComponentForModalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**CustomId** | **string** |  | 
**Style** | [**TextInputStyleTypes**](TextInputStyleTypes.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 

## Methods

### NewTextInputComponentForModalRequest

`func NewTextInputComponentForModalRequest(type_ MessageComponentTypes, customId string, style TextInputStyleTypes, ) *TextInputComponentForModalRequest`

NewTextInputComponentForModalRequest instantiates a new TextInputComponentForModalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextInputComponentForModalRequestWithDefaults

`func NewTextInputComponentForModalRequestWithDefaults() *TextInputComponentForModalRequest`

NewTextInputComponentForModalRequestWithDefaults instantiates a new TextInputComponentForModalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextInputComponentForModalRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextInputComponentForModalRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextInputComponentForModalRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *TextInputComponentForModalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextInputComponentForModalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextInputComponentForModalRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TextInputComponentForModalRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomId

`func (o *TextInputComponentForModalRequest) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *TextInputComponentForModalRequest) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *TextInputComponentForModalRequest) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetStyle

`func (o *TextInputComponentForModalRequest) GetStyle() TextInputStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *TextInputComponentForModalRequest) GetStyleOk() (*TextInputStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *TextInputComponentForModalRequest) SetStyle(v TextInputStyleTypes)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *TextInputComponentForModalRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TextInputComponentForModalRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TextInputComponentForModalRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TextInputComponentForModalRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *TextInputComponentForModalRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TextInputComponentForModalRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TextInputComponentForModalRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TextInputComponentForModalRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPlaceholder

`func (o *TextInputComponentForModalRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *TextInputComponentForModalRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *TextInputComponentForModalRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *TextInputComponentForModalRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetRequired

`func (o *TextInputComponentForModalRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TextInputComponentForModalRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TextInputComponentForModalRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *TextInputComponentForModalRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMinLength

`func (o *TextInputComponentForModalRequest) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *TextInputComponentForModalRequest) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *TextInputComponentForModalRequest) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *TextInputComponentForModalRequest) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *TextInputComponentForModalRequest) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *TextInputComponentForModalRequest) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *TextInputComponentForModalRequest) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *TextInputComponentForModalRequest) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


