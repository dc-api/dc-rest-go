# MessageCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**NullableMessageAllowedMentionsRequest**](MessageAllowedMentionsRequest.md) |  | [optional] 
**StickerIds** | Pointer to **[]string** |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Poll** | Pointer to [**NullablePollCreateRequest**](PollCreateRequest.md) |  | [optional] 
**ConfettiPotion** | Pointer to **map[string]interface{}** |  | [optional] 
**MessageReference** | Pointer to [**NullableMessageReferenceRequest**](MessageReferenceRequest.md) |  | [optional] 
**Nonce** | Pointer to [**NullableBasicMessageResponseNonce**](BasicMessageResponseNonce.md) |  | [optional] 
**EnforceNonce** | Pointer to **NullableBool** |  | [optional] 
**Tts** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMessageCreateRequest

`func NewMessageCreateRequest() *MessageCreateRequest`

NewMessageCreateRequest instantiates a new MessageCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCreateRequestWithDefaults

`func NewMessageCreateRequestWithDefaults() *MessageCreateRequest`

NewMessageCreateRequestWithDefaults instantiates a new MessageCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MessageCreateRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageCreateRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageCreateRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MessageCreateRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MessageCreateRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MessageCreateRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEmbeds

`func (o *MessageCreateRequest) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *MessageCreateRequest) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *MessageCreateRequest) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *MessageCreateRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### SetEmbedsNil

`func (o *MessageCreateRequest) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *MessageCreateRequest) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetAllowedMentions

`func (o *MessageCreateRequest) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *MessageCreateRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *MessageCreateRequest) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *MessageCreateRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### SetAllowedMentionsNil

`func (o *MessageCreateRequest) SetAllowedMentionsNil(b bool)`

 SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil

### UnsetAllowedMentions
`func (o *MessageCreateRequest) UnsetAllowedMentions()`

UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
### GetStickerIds

`func (o *MessageCreateRequest) GetStickerIds() []string`

GetStickerIds returns the StickerIds field if non-nil, zero value otherwise.

### GetStickerIdsOk

`func (o *MessageCreateRequest) GetStickerIdsOk() (*[]string, bool)`

GetStickerIdsOk returns a tuple with the StickerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerIds

`func (o *MessageCreateRequest) SetStickerIds(v []string)`

SetStickerIds sets StickerIds field to given value.

### HasStickerIds

`func (o *MessageCreateRequest) HasStickerIds() bool`

HasStickerIds returns a boolean if a field has been set.

### SetStickerIdsNil

`func (o *MessageCreateRequest) SetStickerIdsNil(b bool)`

 SetStickerIdsNil sets the value for StickerIds to be an explicit nil

### UnsetStickerIds
`func (o *MessageCreateRequest) UnsetStickerIds()`

UnsetStickerIds ensures that no value is present for StickerIds, not even an explicit nil
### GetComponents

`func (o *MessageCreateRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *MessageCreateRequest) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *MessageCreateRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *MessageCreateRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *MessageCreateRequest) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *MessageCreateRequest) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetFlags

`func (o *MessageCreateRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageCreateRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageCreateRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageCreateRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *MessageCreateRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *MessageCreateRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAttachments

`func (o *MessageCreateRequest) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageCreateRequest) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageCreateRequest) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MessageCreateRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *MessageCreateRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *MessageCreateRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetPoll

`func (o *MessageCreateRequest) GetPoll() PollCreateRequest`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *MessageCreateRequest) GetPollOk() (*PollCreateRequest, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *MessageCreateRequest) SetPoll(v PollCreateRequest)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *MessageCreateRequest) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *MessageCreateRequest) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *MessageCreateRequest) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
### GetConfettiPotion

`func (o *MessageCreateRequest) GetConfettiPotion() map[string]interface{}`

GetConfettiPotion returns the ConfettiPotion field if non-nil, zero value otherwise.

### GetConfettiPotionOk

`func (o *MessageCreateRequest) GetConfettiPotionOk() (*map[string]interface{}, bool)`

GetConfettiPotionOk returns a tuple with the ConfettiPotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfettiPotion

`func (o *MessageCreateRequest) SetConfettiPotion(v map[string]interface{})`

SetConfettiPotion sets ConfettiPotion field to given value.

### HasConfettiPotion

`func (o *MessageCreateRequest) HasConfettiPotion() bool`

HasConfettiPotion returns a boolean if a field has been set.

### GetMessageReference

`func (o *MessageCreateRequest) GetMessageReference() MessageReferenceRequest`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *MessageCreateRequest) GetMessageReferenceOk() (*MessageReferenceRequest, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *MessageCreateRequest) SetMessageReference(v MessageReferenceRequest)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *MessageCreateRequest) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

### SetMessageReferenceNil

`func (o *MessageCreateRequest) SetMessageReferenceNil(b bool)`

 SetMessageReferenceNil sets the value for MessageReference to be an explicit nil

### UnsetMessageReference
`func (o *MessageCreateRequest) UnsetMessageReference()`

UnsetMessageReference ensures that no value is present for MessageReference, not even an explicit nil
### GetNonce

`func (o *MessageCreateRequest) GetNonce() BasicMessageResponseNonce`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *MessageCreateRequest) GetNonceOk() (*BasicMessageResponseNonce, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *MessageCreateRequest) SetNonce(v BasicMessageResponseNonce)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *MessageCreateRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *MessageCreateRequest) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *MessageCreateRequest) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetEnforceNonce

`func (o *MessageCreateRequest) GetEnforceNonce() bool`

GetEnforceNonce returns the EnforceNonce field if non-nil, zero value otherwise.

### GetEnforceNonceOk

`func (o *MessageCreateRequest) GetEnforceNonceOk() (*bool, bool)`

GetEnforceNonceOk returns a tuple with the EnforceNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceNonce

`func (o *MessageCreateRequest) SetEnforceNonce(v bool)`

SetEnforceNonce sets EnforceNonce field to given value.

### HasEnforceNonce

`func (o *MessageCreateRequest) HasEnforceNonce() bool`

HasEnforceNonce returns a boolean if a field has been set.

### SetEnforceNonceNil

`func (o *MessageCreateRequest) SetEnforceNonceNil(b bool)`

 SetEnforceNonceNil sets the value for EnforceNonce to be an explicit nil

### UnsetEnforceNonce
`func (o *MessageCreateRequest) UnsetEnforceNonce()`

UnsetEnforceNonce ensures that no value is present for EnforceNonce, not even an explicit nil
### GetTts

`func (o *MessageCreateRequest) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *MessageCreateRequest) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *MessageCreateRequest) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *MessageCreateRequest) HasTts() bool`

HasTts returns a boolean if a field has been set.

### SetTtsNil

`func (o *MessageCreateRequest) SetTtsNil(b bool)`

 SetTtsNil sets the value for Tts to be an explicit nil

### UnsetTts
`func (o *MessageCreateRequest) UnsetTts()`

UnsetTts ensures that no value is present for Tts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


