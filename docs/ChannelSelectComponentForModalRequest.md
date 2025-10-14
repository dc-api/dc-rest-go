# ChannelSelectComponentForModalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | Pointer to **NullableInt32** |  | [optional] 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**MinValues** | Pointer to **NullableInt32** |  | [optional] 
**MaxValues** | Pointer to **NullableInt32** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**DefaultValues** | Pointer to [**[]ChannelSelectDefaultValue**](ChannelSelectDefaultValue.md) |  | [optional] 
**ChannelTypes** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewChannelSelectComponentForModalRequest

`func NewChannelSelectComponentForModalRequest(type_ int32, customId string, ) *ChannelSelectComponentForModalRequest`

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

`func (o *ChannelSelectComponentForModalRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelSelectComponentForModalRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelSelectComponentForModalRequest) SetType(v int32)`

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

### SetIdNil

`func (o *ChannelSelectComponentForModalRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ChannelSelectComponentForModalRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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

### SetPlaceholderNil

`func (o *ChannelSelectComponentForModalRequest) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *ChannelSelectComponentForModalRequest) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
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

### SetMinValuesNil

`func (o *ChannelSelectComponentForModalRequest) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *ChannelSelectComponentForModalRequest) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
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

### SetMaxValuesNil

`func (o *ChannelSelectComponentForModalRequest) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *ChannelSelectComponentForModalRequest) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
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

### SetDisabledNil

`func (o *ChannelSelectComponentForModalRequest) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *ChannelSelectComponentForModalRequest) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
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

### SetRequiredNil

`func (o *ChannelSelectComponentForModalRequest) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ChannelSelectComponentForModalRequest) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
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

### SetDefaultValuesNil

`func (o *ChannelSelectComponentForModalRequest) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *ChannelSelectComponentForModalRequest) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil
### GetChannelTypes

`func (o *ChannelSelectComponentForModalRequest) GetChannelTypes() []int32`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ChannelSelectComponentForModalRequest) GetChannelTypesOk() (*[]int32, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ChannelSelectComponentForModalRequest) SetChannelTypes(v []int32)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ChannelSelectComponentForModalRequest) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### SetChannelTypesNil

`func (o *ChannelSelectComponentForModalRequest) SetChannelTypesNil(b bool)`

 SetChannelTypesNil sets the value for ChannelTypes to be an explicit nil

### UnsetChannelTypes
`func (o *ChannelSelectComponentForModalRequest) UnsetChannelTypes()`

UnsetChannelTypes ensures that no value is present for ChannelTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


