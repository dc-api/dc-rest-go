# ActionRowComponentResponseComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**CustomId** | **string** |  | 
**Style** | **int32** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Emoji** | Pointer to [**ComponentEmojiResponse**](ComponentEmojiResponse.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**MinValues** | Pointer to **int32** |  | [optional] 
**MaxValues** | Pointer to **int32** |  | [optional] 
**ChannelTypes** | Pointer to **[]int32** |  | [optional] 
**DefaultValues** | Pointer to [**[]UserSelectDefaultValueResponse**](UserSelectDefaultValueResponse.md) |  | [optional] 
**Options** | [**[]StringSelectOptionResponse**](StringSelectOptionResponse.md) |  | 
**Value** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 

## Methods

### NewActionRowComponentResponseComponentsInner

`func NewActionRowComponentResponseComponentsInner(type_ int32, id int32, customId string, style int32, options []StringSelectOptionResponse, ) *ActionRowComponentResponseComponentsInner`

NewActionRowComponentResponseComponentsInner instantiates a new ActionRowComponentResponseComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRowComponentResponseComponentsInnerWithDefaults

`func NewActionRowComponentResponseComponentsInnerWithDefaults() *ActionRowComponentResponseComponentsInner`

NewActionRowComponentResponseComponentsInnerWithDefaults instantiates a new ActionRowComponentResponseComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRowComponentResponseComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRowComponentResponseComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRowComponentResponseComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *ActionRowComponentResponseComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRowComponentResponseComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRowComponentResponseComponentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *ActionRowComponentResponseComponentsInner) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ActionRowComponentResponseComponentsInner) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ActionRowComponentResponseComponentsInner) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetStyle

`func (o *ActionRowComponentResponseComponentsInner) GetStyle() int32`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ActionRowComponentResponseComponentsInner) GetStyleOk() (*int32, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ActionRowComponentResponseComponentsInner) SetStyle(v int32)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *ActionRowComponentResponseComponentsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ActionRowComponentResponseComponentsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ActionRowComponentResponseComponentsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ActionRowComponentResponseComponentsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisabled

`func (o *ActionRowComponentResponseComponentsInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ActionRowComponentResponseComponentsInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ActionRowComponentResponseComponentsInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ActionRowComponentResponseComponentsInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmoji

`func (o *ActionRowComponentResponseComponentsInner) GetEmoji() ComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ActionRowComponentResponseComponentsInner) GetEmojiOk() (*ComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ActionRowComponentResponseComponentsInner) SetEmoji(v ComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ActionRowComponentResponseComponentsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetUrl

`func (o *ActionRowComponentResponseComponentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActionRowComponentResponseComponentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActionRowComponentResponseComponentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ActionRowComponentResponseComponentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSkuId

`func (o *ActionRowComponentResponseComponentsInner) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *ActionRowComponentResponseComponentsInner) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *ActionRowComponentResponseComponentsInner) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *ActionRowComponentResponseComponentsInner) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ActionRowComponentResponseComponentsInner) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ActionRowComponentResponseComponentsInner) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ActionRowComponentResponseComponentsInner) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ActionRowComponentResponseComponentsInner) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMinValues

`func (o *ActionRowComponentResponseComponentsInner) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *ActionRowComponentResponseComponentsInner) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *ActionRowComponentResponseComponentsInner) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *ActionRowComponentResponseComponentsInner) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### GetMaxValues

`func (o *ActionRowComponentResponseComponentsInner) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *ActionRowComponentResponseComponentsInner) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *ActionRowComponentResponseComponentsInner) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *ActionRowComponentResponseComponentsInner) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### GetChannelTypes

`func (o *ActionRowComponentResponseComponentsInner) GetChannelTypes() []int32`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ActionRowComponentResponseComponentsInner) GetChannelTypesOk() (*[]int32, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ActionRowComponentResponseComponentsInner) SetChannelTypes(v []int32)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ActionRowComponentResponseComponentsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### GetDefaultValues

`func (o *ActionRowComponentResponseComponentsInner) GetDefaultValues() []UserSelectDefaultValueResponse`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ActionRowComponentResponseComponentsInner) GetDefaultValuesOk() (*[]UserSelectDefaultValueResponse, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ActionRowComponentResponseComponentsInner) SetDefaultValues(v []UserSelectDefaultValueResponse)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ActionRowComponentResponseComponentsInner) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetOptions

`func (o *ActionRowComponentResponseComponentsInner) GetOptions() []StringSelectOptionResponse`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ActionRowComponentResponseComponentsInner) GetOptionsOk() (*[]StringSelectOptionResponse, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ActionRowComponentResponseComponentsInner) SetOptions(v []StringSelectOptionResponse)`

SetOptions sets Options field to given value.


### GetValue

`func (o *ActionRowComponentResponseComponentsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActionRowComponentResponseComponentsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActionRowComponentResponseComponentsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ActionRowComponentResponseComponentsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRequired

`func (o *ActionRowComponentResponseComponentsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ActionRowComponentResponseComponentsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ActionRowComponentResponseComponentsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ActionRowComponentResponseComponentsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMinLength

`func (o *ActionRowComponentResponseComponentsInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ActionRowComponentResponseComponentsInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ActionRowComponentResponseComponentsInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ActionRowComponentResponseComponentsInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *ActionRowComponentResponseComponentsInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ActionRowComponentResponseComponentsInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ActionRowComponentResponseComponentsInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ActionRowComponentResponseComponentsInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


