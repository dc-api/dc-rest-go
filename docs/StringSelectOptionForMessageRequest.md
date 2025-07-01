# StringSelectOptionForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 
**Emoji** | Pointer to [**NullableComponentEmojiForMessageRequest**](ComponentEmojiForMessageRequest.md) |  | [optional] 

## Methods

### NewStringSelectOptionForMessageRequest

`func NewStringSelectOptionForMessageRequest(label string, value string, ) *StringSelectOptionForMessageRequest`

NewStringSelectOptionForMessageRequest instantiates a new StringSelectOptionForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSelectOptionForMessageRequestWithDefaults

`func NewStringSelectOptionForMessageRequestWithDefaults() *StringSelectOptionForMessageRequest`

NewStringSelectOptionForMessageRequestWithDefaults instantiates a new StringSelectOptionForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *StringSelectOptionForMessageRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StringSelectOptionForMessageRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StringSelectOptionForMessageRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *StringSelectOptionForMessageRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringSelectOptionForMessageRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringSelectOptionForMessageRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *StringSelectOptionForMessageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringSelectOptionForMessageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringSelectOptionForMessageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringSelectOptionForMessageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StringSelectOptionForMessageRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StringSelectOptionForMessageRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefault

`func (o *StringSelectOptionForMessageRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *StringSelectOptionForMessageRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *StringSelectOptionForMessageRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *StringSelectOptionForMessageRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *StringSelectOptionForMessageRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *StringSelectOptionForMessageRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetEmoji

`func (o *StringSelectOptionForMessageRequest) GetEmoji() ComponentEmojiForMessageRequest`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *StringSelectOptionForMessageRequest) GetEmojiOk() (*ComponentEmojiForMessageRequest, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *StringSelectOptionForMessageRequest) SetEmoji(v ComponentEmojiForMessageRequest)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *StringSelectOptionForMessageRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *StringSelectOptionForMessageRequest) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *StringSelectOptionForMessageRequest) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


