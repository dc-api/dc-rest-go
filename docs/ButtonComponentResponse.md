# ButtonComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | **int32** |  | 
**CustomId** | Pointer to **NullableString** |  | [optional] 
**Style** | **NullableInt32** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Emoji** | Pointer to [**NullableComponentEmojiResponse**](ComponentEmojiResponse.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 

## Methods

### NewButtonComponentResponse

`func NewButtonComponentResponse(type_ int32, id int32, style NullableInt32, ) *ButtonComponentResponse`

NewButtonComponentResponse instantiates a new ButtonComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonComponentResponseWithDefaults

`func NewButtonComponentResponseWithDefaults() *ButtonComponentResponse`

NewButtonComponentResponseWithDefaults instantiates a new ButtonComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ButtonComponentResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ButtonComponentResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ButtonComponentResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *ButtonComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ButtonComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ButtonComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *ButtonComponentResponse) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ButtonComponentResponse) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ButtonComponentResponse) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *ButtonComponentResponse) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### SetCustomIdNil

`func (o *ButtonComponentResponse) SetCustomIdNil(b bool)`

 SetCustomIdNil sets the value for CustomId to be an explicit nil

### UnsetCustomId
`func (o *ButtonComponentResponse) UnsetCustomId()`

UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
### GetStyle

`func (o *ButtonComponentResponse) GetStyle() int32`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ButtonComponentResponse) GetStyleOk() (*int32, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ButtonComponentResponse) SetStyle(v int32)`

SetStyle sets Style field to given value.


### SetStyleNil

`func (o *ButtonComponentResponse) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *ButtonComponentResponse) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetLabel

`func (o *ButtonComponentResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ButtonComponentResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ButtonComponentResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ButtonComponentResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ButtonComponentResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ButtonComponentResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisabled

`func (o *ButtonComponentResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ButtonComponentResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ButtonComponentResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ButtonComponentResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *ButtonComponentResponse) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *ButtonComponentResponse) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEmoji

`func (o *ButtonComponentResponse) GetEmoji() ComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ButtonComponentResponse) GetEmojiOk() (*ComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ButtonComponentResponse) SetEmoji(v ComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ButtonComponentResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *ButtonComponentResponse) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *ButtonComponentResponse) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetUrl

`func (o *ButtonComponentResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ButtonComponentResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ButtonComponentResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ButtonComponentResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ButtonComponentResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ButtonComponentResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSkuId

`func (o *ButtonComponentResponse) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *ButtonComponentResponse) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *ButtonComponentResponse) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *ButtonComponentResponse) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


