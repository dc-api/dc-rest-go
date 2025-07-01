# ActionRowComponentForMessageRequestComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**CustomId** | **string** |  | 
**Style** | **int32** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to [**ComponentEmojiForMessageRequest**](ComponentEmojiForMessageRequest.md) |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**MinValues** | Pointer to **int32** |  | [optional] 
**MaxValues** | Pointer to **int32** |  | [optional] 
**DefaultValues** | Pointer to [**[]UserSelectDefaultValue**](UserSelectDefaultValue.md) |  | [optional] 
**ChannelTypes** | Pointer to **[]int32** |  | [optional] 
**Options** | [**[]StringSelectOptionForMessageRequest**](StringSelectOptionForMessageRequest.md) |  | 

## Methods

### NewActionRowComponentForMessageRequestComponentsInner

`func NewActionRowComponentForMessageRequestComponentsInner(type_ int32, customId string, style int32, options []StringSelectOptionForMessageRequest, ) *ActionRowComponentForMessageRequestComponentsInner`

NewActionRowComponentForMessageRequestComponentsInner instantiates a new ActionRowComponentForMessageRequestComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRowComponentForMessageRequestComponentsInnerWithDefaults

`func NewActionRowComponentForMessageRequestComponentsInnerWithDefaults() *ActionRowComponentForMessageRequestComponentsInner`

NewActionRowComponentForMessageRequestComponentsInnerWithDefaults instantiates a new ActionRowComponentForMessageRequestComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetType(v int32)`

SetType sets Type field to given value.


### GetCustomId

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetStyle

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetStyle() int32`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetStyleOk() (*int32, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetStyle(v int32)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisabled

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSkuId

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetEmoji

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetEmoji() ComponentEmojiForMessageRequest`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetEmojiOk() (*ComponentEmojiForMessageRequest, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetEmoji(v ComponentEmojiForMessageRequest)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMinValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### GetMaxValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### GetDefaultValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetDefaultValues() []UserSelectDefaultValue`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetDefaultValuesOk() (*[]UserSelectDefaultValue, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetDefaultValues(v []UserSelectDefaultValue)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetChannelTypes

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetChannelTypes() []int32`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetChannelTypesOk() (*[]int32, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetChannelTypes(v []int32)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### GetOptions

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetOptions() []StringSelectOptionForMessageRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetOptionsOk() (*[]StringSelectOptionForMessageRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetOptions(v []StringSelectOptionForMessageRequest)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


