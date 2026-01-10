# ChannelSelectComponentForModalRequest

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
**DefaultValues** | Pointer to [**[]ChannelSelectDefaultValue**](ChannelSelectDefaultValue.md) |  | [optional] 
**ChannelTypes** | Pointer to [**[]ChannelTypes**](ChannelTypes.md) |  | [optional] 

## Methods

### NewChannelSelectComponentForModalRequest

`func NewChannelSelectComponentForModalRequest(type_ MessageComponentTypes, customId string, ) *ChannelSelectComponentForModalRequest`

NewChannelSelectComponentForModalRequest instantiates a new ChannelSelectComponentForModalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSelectComponentForModalRequestWithDefaults

`func NewChannelSelectComponentForModalRequestWithDefaults() *ChannelSelectComponentForModalRequest`

NewChannelSelectComponentForModalRequestWithDefaults instantiates a new ChannelSelectComponentForModalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChannelSelectComponentForModalRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelSelectComponentForModalRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelSelectComponentForModalRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ChannelSelectComponentForModalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelSelectComponentForModalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelSelectComponentForModalRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelSelectComponentForModalRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomId

`func (o *ChannelSelectComponentForModalRequest) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ChannelSelectComponentForModalRequest) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ChannelSelectComponentForModalRequest) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *ChannelSelectComponentForModalRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ChannelSelectComponentForModalRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ChannelSelectComponentForModalRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ChannelSelectComponentForModalRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMinValues

`func (o *ChannelSelectComponentForModalRequest) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *ChannelSelectComponentForModalRequest) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *ChannelSelectComponentForModalRequest) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *ChannelSelectComponentForModalRequest) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### GetMaxValues

`func (o *ChannelSelectComponentForModalRequest) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *ChannelSelectComponentForModalRequest) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *ChannelSelectComponentForModalRequest) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *ChannelSelectComponentForModalRequest) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### GetDisabled

`func (o *ChannelSelectComponentForModalRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ChannelSelectComponentForModalRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ChannelSelectComponentForModalRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ChannelSelectComponentForModalRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRequired

`func (o *ChannelSelectComponentForModalRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ChannelSelectComponentForModalRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ChannelSelectComponentForModalRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ChannelSelectComponentForModalRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValues

`func (o *ChannelSelectComponentForModalRequest) GetDefaultValues() []ChannelSelectDefaultValue`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ChannelSelectComponentForModalRequest) GetDefaultValuesOk() (*[]ChannelSelectDefaultValue, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ChannelSelectComponentForModalRequest) SetDefaultValues(v []ChannelSelectDefaultValue)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ChannelSelectComponentForModalRequest) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetChannelTypes

`func (o *ChannelSelectComponentForModalRequest) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ChannelSelectComponentForModalRequest) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ChannelSelectComponentForModalRequest) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ChannelSelectComponentForModalRequest) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


