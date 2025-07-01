# StringSelectOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Emoji** | Pointer to [**NullableComponentEmojiResponse**](ComponentEmojiResponse.md) |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewStringSelectOptionResponse

`func NewStringSelectOptionResponse(label string, value string, ) *StringSelectOptionResponse`

NewStringSelectOptionResponse instantiates a new StringSelectOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSelectOptionResponseWithDefaults

`func NewStringSelectOptionResponseWithDefaults() *StringSelectOptionResponse`

NewStringSelectOptionResponseWithDefaults instantiates a new StringSelectOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *StringSelectOptionResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StringSelectOptionResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StringSelectOptionResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *StringSelectOptionResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringSelectOptionResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringSelectOptionResponse) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *StringSelectOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringSelectOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringSelectOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringSelectOptionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StringSelectOptionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StringSelectOptionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmoji

`func (o *StringSelectOptionResponse) GetEmoji() ComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *StringSelectOptionResponse) GetEmojiOk() (*ComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *StringSelectOptionResponse) SetEmoji(v ComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *StringSelectOptionResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *StringSelectOptionResponse) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *StringSelectOptionResponse) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetDefault

`func (o *StringSelectOptionResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *StringSelectOptionResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *StringSelectOptionResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *StringSelectOptionResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *StringSelectOptionResponse) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *StringSelectOptionResponse) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


