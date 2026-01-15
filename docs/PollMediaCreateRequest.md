# PollMediaCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to [**PollMediaCreateRequestEmoji**](PollMediaCreateRequestEmoji.md) |  | [optional] 

## Methods

### NewPollMediaCreateRequest

`func NewPollMediaCreateRequest() *PollMediaCreateRequest`

NewPollMediaCreateRequest instantiates a new PollMediaCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollMediaCreateRequestWithDefaults

`func NewPollMediaCreateRequestWithDefaults() *PollMediaCreateRequest`

NewPollMediaCreateRequestWithDefaults instantiates a new PollMediaCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PollMediaCreateRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PollMediaCreateRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PollMediaCreateRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PollMediaCreateRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEmoji

`func (o *PollMediaCreateRequest) GetEmoji() PollMediaCreateRequestEmoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *PollMediaCreateRequest) GetEmojiOk() (*PollMediaCreateRequestEmoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *PollMediaCreateRequest) SetEmoji(v PollMediaCreateRequestEmoji)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *PollMediaCreateRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


