# PollMediaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to [**MessageReactionEmojiResponse**](MessageReactionEmojiResponse.md) |  | [optional] 

## Methods

### NewPollMediaResponse

`func NewPollMediaResponse() *PollMediaResponse`

NewPollMediaResponse instantiates a new PollMediaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollMediaResponseWithDefaults

`func NewPollMediaResponseWithDefaults() *PollMediaResponse`

NewPollMediaResponseWithDefaults instantiates a new PollMediaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PollMediaResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PollMediaResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PollMediaResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PollMediaResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEmoji

`func (o *PollMediaResponse) GetEmoji() MessageReactionEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *PollMediaResponse) GetEmojiOk() (*MessageReactionEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *PollMediaResponse) SetEmoji(v MessageReactionEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *PollMediaResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


