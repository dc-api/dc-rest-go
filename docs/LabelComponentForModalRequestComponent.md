# LabelComponentForModalRequestComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **string** |  | [optional] 
**MinValues** | Pointer to **int32** |  | [optional] 
**MaxValues** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DefaultValues** | Pointer to [**[]UserSelectDefaultValue**](UserSelectDefaultValue.md) |  | [optional] 
**ChannelTypes** | Pointer to [**[]ChannelTypes**](ChannelTypes.md) |  | [optional] 
**Options** | [**[]StringSelectOptionForRequest**](StringSelectOptionForRequest.md) |  | 
**Style** | [**TextInputStyleTypes**](TextInputStyleTypes.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 

## Methods

### NewLabelComponentForModalRequestComponent

`func NewLabelComponentForModalRequestComponent(type_ MessageComponentTypes, customId string, options []StringSelectOptionForRequest, style TextInputStyleTypes, ) *LabelComponentForModalRequestComponent`

NewLabelComponentForModalRequestComponent instantiates a new LabelComponentForModalRequestComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelComponentForModalRequestComponentWithDefaults

`func NewLabelComponentForModalRequestComponentWithDefaults() *LabelComponentForModalRequestComponent`

NewLabelComponentForModalRequestComponentWithDefaults instantiates a new LabelComponentForModalRequestComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LabelComponentForModalRequestComponent) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelComponentForModalRequestComponent) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelComponentForModalRequestComponent) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *LabelComponentForModalRequestComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelComponentForModalRequestComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelComponentForModalRequestComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LabelComponentForModalRequestComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomId

`func (o *LabelComponentForModalRequestComponent) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *LabelComponentForModalRequestComponent) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *LabelComponentForModalRequestComponent) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *LabelComponentForModalRequestComponent) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *LabelComponentForModalRequestComponent) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *LabelComponentForModalRequestComponent) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *LabelComponentForModalRequestComponent) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMinValues

`func (o *LabelComponentForModalRequestComponent) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *LabelComponentForModalRequestComponent) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *LabelComponentForModalRequestComponent) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *LabelComponentForModalRequestComponent) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### GetMaxValues

`func (o *LabelComponentForModalRequestComponent) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *LabelComponentForModalRequestComponent) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *LabelComponentForModalRequestComponent) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *LabelComponentForModalRequestComponent) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### GetDisabled

`func (o *LabelComponentForModalRequestComponent) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *LabelComponentForModalRequestComponent) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *LabelComponentForModalRequestComponent) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *LabelComponentForModalRequestComponent) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRequired

`func (o *LabelComponentForModalRequestComponent) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *LabelComponentForModalRequestComponent) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *LabelComponentForModalRequestComponent) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *LabelComponentForModalRequestComponent) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValues

`func (o *LabelComponentForModalRequestComponent) GetDefaultValues() []UserSelectDefaultValue`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *LabelComponentForModalRequestComponent) GetDefaultValuesOk() (*[]UserSelectDefaultValue, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *LabelComponentForModalRequestComponent) SetDefaultValues(v []UserSelectDefaultValue)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *LabelComponentForModalRequestComponent) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetChannelTypes

`func (o *LabelComponentForModalRequestComponent) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *LabelComponentForModalRequestComponent) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *LabelComponentForModalRequestComponent) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *LabelComponentForModalRequestComponent) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### GetOptions

`func (o *LabelComponentForModalRequestComponent) GetOptions() []StringSelectOptionForRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *LabelComponentForModalRequestComponent) GetOptionsOk() (*[]StringSelectOptionForRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *LabelComponentForModalRequestComponent) SetOptions(v []StringSelectOptionForRequest)`

SetOptions sets Options field to given value.


### GetStyle

`func (o *LabelComponentForModalRequestComponent) GetStyle() TextInputStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *LabelComponentForModalRequestComponent) GetStyleOk() (*TextInputStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *LabelComponentForModalRequestComponent) SetStyle(v TextInputStyleTypes)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *LabelComponentForModalRequestComponent) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabelComponentForModalRequestComponent) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabelComponentForModalRequestComponent) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LabelComponentForModalRequestComponent) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *LabelComponentForModalRequestComponent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelComponentForModalRequestComponent) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelComponentForModalRequestComponent) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LabelComponentForModalRequestComponent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMinLength

`func (o *LabelComponentForModalRequestComponent) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *LabelComponentForModalRequestComponent) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *LabelComponentForModalRequestComponent) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *LabelComponentForModalRequestComponent) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *LabelComponentForModalRequestComponent) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *LabelComponentForModalRequestComponent) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *LabelComponentForModalRequestComponent) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *LabelComponentForModalRequestComponent) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


