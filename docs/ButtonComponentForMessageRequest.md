# ButtonComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**CustomId** | Pointer to **NullableString** |  | [optional] 
**Style** | **NullableInt32** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to [**NullableComponentEmojiForMessageRequest**](ComponentEmojiForMessageRequest.md) |  | [optional] 

## Methods

### NewButtonComponentForMessageRequest

`func NewButtonComponentForMessageRequest(type_ int32, style NullableInt32, ) *ButtonComponentForMessageRequest`

NewButtonComponentForMessageRequest instantiates a new ButtonComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonComponentForMessageRequestWithDefaults

`func NewButtonComponentForMessageRequestWithDefaults() *ButtonComponentForMessageRequest`

NewButtonComponentForMessageRequestWithDefaults instantiates a new ButtonComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ButtonComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ButtonComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ButtonComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetCustomId

`func (o *ButtonComponentForMessageRequest) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ButtonComponentForMessageRequest) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ButtonComponentForMessageRequest) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *ButtonComponentForMessageRequest) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### SetCustomIdNil

`func (o *ButtonComponentForMessageRequest) SetCustomIdNil(b bool)`

 SetCustomIdNil sets the value for CustomId to be an explicit nil

### UnsetCustomId
`func (o *ButtonComponentForMessageRequest) UnsetCustomId()`

UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
### GetStyle

`func (o *ButtonComponentForMessageRequest) GetStyle() int32`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ButtonComponentForMessageRequest) GetStyleOk() (*int32, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ButtonComponentForMessageRequest) SetStyle(v int32)`

SetStyle sets Style field to given value.


### SetStyleNil

`func (o *ButtonComponentForMessageRequest) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *ButtonComponentForMessageRequest) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetLabel

`func (o *ButtonComponentForMessageRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ButtonComponentForMessageRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ButtonComponentForMessageRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ButtonComponentForMessageRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ButtonComponentForMessageRequest) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ButtonComponentForMessageRequest) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisabled

`func (o *ButtonComponentForMessageRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ButtonComponentForMessageRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ButtonComponentForMessageRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ButtonComponentForMessageRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *ButtonComponentForMessageRequest) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *ButtonComponentForMessageRequest) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetUrl

`func (o *ButtonComponentForMessageRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ButtonComponentForMessageRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ButtonComponentForMessageRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ButtonComponentForMessageRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ButtonComponentForMessageRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ButtonComponentForMessageRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSkuId

`func (o *ButtonComponentForMessageRequest) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *ButtonComponentForMessageRequest) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *ButtonComponentForMessageRequest) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *ButtonComponentForMessageRequest) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetEmoji

`func (o *ButtonComponentForMessageRequest) GetEmoji() ComponentEmojiForMessageRequest`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ButtonComponentForMessageRequest) GetEmojiOk() (*ComponentEmojiForMessageRequest, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ButtonComponentForMessageRequest) SetEmoji(v ComponentEmojiForMessageRequest)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ButtonComponentForMessageRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *ButtonComponentForMessageRequest) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *ButtonComponentForMessageRequest) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


