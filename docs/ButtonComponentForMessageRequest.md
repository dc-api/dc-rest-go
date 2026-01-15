# ButtonComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**CustomId** | Pointer to **string** |  | [optional] 
**Style** | [**ButtonStyleTypes**](ButtonStyleTypes.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Emoji** | Pointer to [**ButtonComponentForMessageRequestEmoji**](ButtonComponentForMessageRequestEmoji.md) |  | [optional] 

## Methods

### NewButtonComponentForMessageRequest

`func NewButtonComponentForMessageRequest(type_ MessageComponentTypes, style ButtonStyleTypes, ) *ButtonComponentForMessageRequest`

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

`func (o *ButtonComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ButtonComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ButtonComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *ButtonComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ButtonComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ButtonComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ButtonComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

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

### GetStyle

`func (o *ButtonComponentForMessageRequest) GetStyle() ButtonStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ButtonComponentForMessageRequest) GetStyleOk() (*ButtonStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ButtonComponentForMessageRequest) SetStyle(v ButtonStyleTypes)`

SetStyle sets Style field to given value.


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

### GetSkuId

`func (o *ButtonComponentForMessageRequest) GetSkuId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *ButtonComponentForMessageRequest) GetSkuIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *ButtonComponentForMessageRequest) SetSkuId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *ButtonComponentForMessageRequest) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetEmoji

`func (o *ButtonComponentForMessageRequest) GetEmoji() ButtonComponentForMessageRequestEmoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ButtonComponentForMessageRequest) GetEmojiOk() (*ButtonComponentForMessageRequestEmoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ButtonComponentForMessageRequest) SetEmoji(v ButtonComponentForMessageRequestEmoji)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ButtonComponentForMessageRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


