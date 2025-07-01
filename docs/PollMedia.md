# PollMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **NullableString** |  | [optional] 
**Emoji** | Pointer to [**NullablePollEmoji**](PollEmoji.md) |  | [optional] 

## Methods

### NewPollMedia

`func NewPollMedia() *PollMedia`

NewPollMedia instantiates a new PollMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollMediaWithDefaults

`func NewPollMediaWithDefaults() *PollMedia`

NewPollMediaWithDefaults instantiates a new PollMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PollMedia) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PollMedia) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PollMedia) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PollMedia) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *PollMedia) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *PollMedia) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetEmoji

`func (o *PollMedia) GetEmoji() PollEmoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *PollMedia) GetEmojiOk() (*PollEmoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *PollMedia) SetEmoji(v PollEmoji)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *PollMedia) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *PollMedia) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *PollMedia) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


