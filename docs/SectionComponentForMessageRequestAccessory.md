# SectionComponentForMessageRequestAccessory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**CustomId** | Pointer to **string** |  | [optional] 
**Style** | **int32** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to [**ComponentEmojiForMessageRequest**](ComponentEmojiForMessageRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Spoiler** | Pointer to **bool** |  | [optional] 
**Media** | [**UnfurledMediaRequest**](UnfurledMediaRequest.md) |  | 

## Methods

### NewSectionComponentForMessageRequestAccessory

`func NewSectionComponentForMessageRequestAccessory(type_ int32, style int32, media UnfurledMediaRequest, ) *SectionComponentForMessageRequestAccessory`

NewSectionComponentForMessageRequestAccessory instantiates a new SectionComponentForMessageRequestAccessory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionComponentForMessageRequestAccessoryWithDefaults

`func NewSectionComponentForMessageRequestAccessoryWithDefaults() *SectionComponentForMessageRequestAccessory`

NewSectionComponentForMessageRequestAccessoryWithDefaults instantiates a new SectionComponentForMessageRequestAccessory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SectionComponentForMessageRequestAccessory) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectionComponentForMessageRequestAccessory) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectionComponentForMessageRequestAccessory) SetType(v int32)`

SetType sets Type field to given value.


### GetCustomId

`func (o *SectionComponentForMessageRequestAccessory) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *SectionComponentForMessageRequestAccessory) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *SectionComponentForMessageRequestAccessory) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *SectionComponentForMessageRequestAccessory) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetStyle

`func (o *SectionComponentForMessageRequestAccessory) GetStyle() int32`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *SectionComponentForMessageRequestAccessory) GetStyleOk() (*int32, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *SectionComponentForMessageRequestAccessory) SetStyle(v int32)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *SectionComponentForMessageRequestAccessory) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SectionComponentForMessageRequestAccessory) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SectionComponentForMessageRequestAccessory) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SectionComponentForMessageRequestAccessory) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisabled

`func (o *SectionComponentForMessageRequestAccessory) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SectionComponentForMessageRequestAccessory) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SectionComponentForMessageRequestAccessory) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SectionComponentForMessageRequestAccessory) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *SectionComponentForMessageRequestAccessory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SectionComponentForMessageRequestAccessory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SectionComponentForMessageRequestAccessory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SectionComponentForMessageRequestAccessory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSkuId

`func (o *SectionComponentForMessageRequestAccessory) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *SectionComponentForMessageRequestAccessory) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *SectionComponentForMessageRequestAccessory) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *SectionComponentForMessageRequestAccessory) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetEmoji

`func (o *SectionComponentForMessageRequestAccessory) GetEmoji() ComponentEmojiForMessageRequest`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *SectionComponentForMessageRequestAccessory) GetEmojiOk() (*ComponentEmojiForMessageRequest, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *SectionComponentForMessageRequestAccessory) SetEmoji(v ComponentEmojiForMessageRequest)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *SectionComponentForMessageRequestAccessory) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetDescription

`func (o *SectionComponentForMessageRequestAccessory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SectionComponentForMessageRequestAccessory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SectionComponentForMessageRequestAccessory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SectionComponentForMessageRequestAccessory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpoiler

`func (o *SectionComponentForMessageRequestAccessory) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *SectionComponentForMessageRequestAccessory) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *SectionComponentForMessageRequestAccessory) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *SectionComponentForMessageRequestAccessory) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### GetMedia

`func (o *SectionComponentForMessageRequestAccessory) GetMedia() UnfurledMediaRequest`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SectionComponentForMessageRequestAccessory) GetMediaOk() (*UnfurledMediaRequest, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SectionComponentForMessageRequestAccessory) SetMedia(v UnfurledMediaRequest)`

SetMedia sets Media field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


