# StringSelectOptionForRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 
**Emoji** | Pointer to [**NullableComponentEmojiForRequest**](ComponentEmojiForRequest.md) |  | [optional] 

## Methods

### NewStringSelectOptionForRequest

`func NewStringSelectOptionForRequest(label string, value string, ) *StringSelectOptionForRequest`

NewStringSelectOptionForRequest instantiates a new StringSelectOptionForRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSelectOptionForRequestWithDefaults

`func NewStringSelectOptionForRequestWithDefaults() *StringSelectOptionForRequest`

NewStringSelectOptionForRequestWithDefaults instantiates a new StringSelectOptionForRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *StringSelectOptionForRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StringSelectOptionForRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StringSelectOptionForRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *StringSelectOptionForRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringSelectOptionForRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringSelectOptionForRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *StringSelectOptionForRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringSelectOptionForRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringSelectOptionForRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringSelectOptionForRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StringSelectOptionForRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StringSelectOptionForRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefault

`func (o *StringSelectOptionForRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *StringSelectOptionForRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *StringSelectOptionForRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *StringSelectOptionForRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *StringSelectOptionForRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *StringSelectOptionForRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetEmoji

`func (o *StringSelectOptionForRequest) GetEmoji() ComponentEmojiForRequest`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *StringSelectOptionForRequest) GetEmojiOk() (*ComponentEmojiForRequest, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *StringSelectOptionForRequest) SetEmoji(v ComponentEmojiForRequest)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *StringSelectOptionForRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *StringSelectOptionForRequest) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *StringSelectOptionForRequest) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


