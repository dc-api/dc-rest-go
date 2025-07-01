# SectionComponentResponseAccessory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**CustomId** | Pointer to **string** |  | [optional] 
**Style** | **int32** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Emoji** | Pointer to [**ComponentEmojiResponse**](ComponentEmojiResponse.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**Media** | [**UnfurledMediaResponse**](UnfurledMediaResponse.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Spoiler** | **bool** |  | 

## Methods

### NewSectionComponentResponseAccessory

`func NewSectionComponentResponseAccessory(type_ int32, id int32, style int32, media UnfurledMediaResponse, spoiler bool, ) *SectionComponentResponseAccessory`

NewSectionComponentResponseAccessory instantiates a new SectionComponentResponseAccessory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionComponentResponseAccessoryWithDefaults

`func NewSectionComponentResponseAccessoryWithDefaults() *SectionComponentResponseAccessory`

NewSectionComponentResponseAccessoryWithDefaults instantiates a new SectionComponentResponseAccessory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SectionComponentResponseAccessory) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SectionComponentResponseAccessory) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SectionComponentResponseAccessory) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *SectionComponentResponseAccessory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectionComponentResponseAccessory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectionComponentResponseAccessory) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *SectionComponentResponseAccessory) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *SectionComponentResponseAccessory) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *SectionComponentResponseAccessory) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *SectionComponentResponseAccessory) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetStyle

`func (o *SectionComponentResponseAccessory) GetStyle() int32`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *SectionComponentResponseAccessory) GetStyleOk() (*int32, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *SectionComponentResponseAccessory) SetStyle(v int32)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *SectionComponentResponseAccessory) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SectionComponentResponseAccessory) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SectionComponentResponseAccessory) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SectionComponentResponseAccessory) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDisabled

`func (o *SectionComponentResponseAccessory) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SectionComponentResponseAccessory) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SectionComponentResponseAccessory) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SectionComponentResponseAccessory) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmoji

`func (o *SectionComponentResponseAccessory) GetEmoji() ComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *SectionComponentResponseAccessory) GetEmojiOk() (*ComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *SectionComponentResponseAccessory) SetEmoji(v ComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *SectionComponentResponseAccessory) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetUrl

`func (o *SectionComponentResponseAccessory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SectionComponentResponseAccessory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SectionComponentResponseAccessory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SectionComponentResponseAccessory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSkuId

`func (o *SectionComponentResponseAccessory) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *SectionComponentResponseAccessory) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *SectionComponentResponseAccessory) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *SectionComponentResponseAccessory) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetMedia

`func (o *SectionComponentResponseAccessory) GetMedia() UnfurledMediaResponse`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SectionComponentResponseAccessory) GetMediaOk() (*UnfurledMediaResponse, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SectionComponentResponseAccessory) SetMedia(v UnfurledMediaResponse)`

SetMedia sets Media field to given value.


### GetDescription

`func (o *SectionComponentResponseAccessory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SectionComponentResponseAccessory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SectionComponentResponseAccessory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SectionComponentResponseAccessory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpoiler

`func (o *SectionComponentResponseAccessory) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *SectionComponentResponseAccessory) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *SectionComponentResponseAccessory) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


