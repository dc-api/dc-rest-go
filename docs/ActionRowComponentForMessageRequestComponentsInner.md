# ActionRowComponentForMessageRequestComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**CustomId** | **string** |  | 
**Style** | [**ButtonStyleTypes**](ButtonStyleTypes.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Emoji** | Pointer to [**ButtonComponentForMessageRequestEmoji**](ButtonComponentForMessageRequestEmoji.md) |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**MinValues** | Pointer to **int32** |  | [optional] 
**MaxValues** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DefaultValues** | Pointer to [**[]UserSelectDefaultValue**](UserSelectDefaultValue.md) |  | [optional] 
**ChannelTypes** | Pointer to [**[]ChannelTypes**](ChannelTypes.md) |  | [optional] 
**Options** | [**[]StringSelectOptionForRequest**](StringSelectOptionForRequest.md) |  | 

## Methods

### NewActionRowComponentForMessageRequestComponentsInner

`func NewActionRowComponentForMessageRequestComponentsInner(type_ MessageComponentTypes, customId string, style ButtonStyleTypes, options []StringSelectOptionForRequest, ) *ActionRowComponentForMessageRequestComponentsInner`

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

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

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

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetStyle() ButtonStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetStyleOk() (*ButtonStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetStyle(v ButtonStyleTypes)`

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

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetSkuId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetSkuIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetSkuId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetEmoji

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetEmoji() ButtonComponentForMessageRequestEmoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetEmojiOk() (*ButtonComponentForMessageRequestEmoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetEmoji(v ButtonComponentForMessageRequestEmoji)`

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

### GetRequired

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

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

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ActionRowComponentForMessageRequestComponentsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### GetOptions

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetOptions() []StringSelectOptionForRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ActionRowComponentForMessageRequestComponentsInner) GetOptionsOk() (*[]StringSelectOptionForRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ActionRowComponentForMessageRequestComponentsInner) SetOptions(v []StringSelectOptionForRequest)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


