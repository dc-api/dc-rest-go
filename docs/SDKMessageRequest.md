# SDKMessageRequest

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

### NewSDKMessageRequest

`func NewSDKMessageRequest() *SDKMessageRequest`

NewSDKMessageRequest instantiates a new SDKMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDKMessageRequestWithDefaults

`func NewSDKMessageRequestWithDefaults() *SDKMessageRequest`

NewSDKMessageRequestWithDefaults instantiates a new SDKMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SDKMessageRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SDKMessageRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SDKMessageRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SDKMessageRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SDKMessageRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SDKMessageRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEmbeds

`func (o *SDKMessageRequest) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *SDKMessageRequest) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *SDKMessageRequest) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *SDKMessageRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### SetEmbedsNil

`func (o *SDKMessageRequest) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *SDKMessageRequest) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetAllowedMentions

`func (o *SDKMessageRequest) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *SDKMessageRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *SDKMessageRequest) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *SDKMessageRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### SetAllowedMentionsNil

`func (o *SDKMessageRequest) SetAllowedMentionsNil(b bool)`

 SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil

### UnsetAllowedMentions
`func (o *SDKMessageRequest) UnsetAllowedMentions()`

UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
### GetStickerIds

`func (o *SDKMessageRequest) GetStickerIds() []string`

GetStickerIds returns the StickerIds field if non-nil, zero value otherwise.

### GetStickerIdsOk

`func (o *SDKMessageRequest) GetStickerIdsOk() (*[]string, bool)`

GetStickerIdsOk returns a tuple with the StickerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerIds

`func (o *SDKMessageRequest) SetStickerIds(v []string)`

SetStickerIds sets StickerIds field to given value.

### HasStickerIds

`func (o *SDKMessageRequest) HasStickerIds() bool`

HasStickerIds returns a boolean if a field has been set.

### SetStickerIdsNil

`func (o *SDKMessageRequest) SetStickerIdsNil(b bool)`

 SetStickerIdsNil sets the value for StickerIds to be an explicit nil

### UnsetStickerIds
`func (o *SDKMessageRequest) UnsetStickerIds()`

UnsetStickerIds ensures that no value is present for StickerIds, not even an explicit nil
### GetComponents

`func (o *SDKMessageRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SDKMessageRequest) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SDKMessageRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SDKMessageRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *SDKMessageRequest) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *SDKMessageRequest) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetFlags

`func (o *SDKMessageRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SDKMessageRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SDKMessageRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SDKMessageRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *SDKMessageRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *SDKMessageRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAttachments

`func (o *SDKMessageRequest) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SDKMessageRequest) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SDKMessageRequest) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SDKMessageRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *SDKMessageRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *SDKMessageRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetPoll

`func (o *SDKMessageRequest) GetPoll() PollCreateRequest`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *SDKMessageRequest) GetPollOk() (*PollCreateRequest, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *SDKMessageRequest) SetPoll(v PollCreateRequest)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *SDKMessageRequest) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *SDKMessageRequest) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *SDKMessageRequest) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
### GetConfettiPotion

`func (o *SDKMessageRequest) GetConfettiPotion() map[string]interface{}`

GetConfettiPotion returns the ConfettiPotion field if non-nil, zero value otherwise.

### GetConfettiPotionOk

`func (o *SDKMessageRequest) GetConfettiPotionOk() (*map[string]interface{}, bool)`

GetConfettiPotionOk returns a tuple with the ConfettiPotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfettiPotion

`func (o *SDKMessageRequest) SetConfettiPotion(v map[string]interface{})`

SetConfettiPotion sets ConfettiPotion field to given value.

### HasConfettiPotion

`func (o *SDKMessageRequest) HasConfettiPotion() bool`

HasConfettiPotion returns a boolean if a field has been set.

### GetMessageReference

`func (o *SDKMessageRequest) GetMessageReference() MessageReferenceRequest`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *SDKMessageRequest) GetMessageReferenceOk() (*MessageReferenceRequest, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *SDKMessageRequest) SetMessageReference(v MessageReferenceRequest)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *SDKMessageRequest) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

### SetMessageReferenceNil

`func (o *SDKMessageRequest) SetMessageReferenceNil(b bool)`

 SetMessageReferenceNil sets the value for MessageReference to be an explicit nil

### UnsetMessageReference
`func (o *SDKMessageRequest) UnsetMessageReference()`

UnsetMessageReference ensures that no value is present for MessageReference, not even an explicit nil
### GetNonce

`func (o *SDKMessageRequest) GetNonce() BasicMessageResponseNonce`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SDKMessageRequest) GetNonceOk() (*BasicMessageResponseNonce, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SDKMessageRequest) SetNonce(v BasicMessageResponseNonce)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SDKMessageRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *SDKMessageRequest) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *SDKMessageRequest) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetEnforceNonce

`func (o *SDKMessageRequest) GetEnforceNonce() bool`

GetEnforceNonce returns the EnforceNonce field if non-nil, zero value otherwise.

### GetEnforceNonceOk

`func (o *SDKMessageRequest) GetEnforceNonceOk() (*bool, bool)`

GetEnforceNonceOk returns a tuple with the EnforceNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceNonce

`func (o *SDKMessageRequest) SetEnforceNonce(v bool)`

SetEnforceNonce sets EnforceNonce field to given value.

### HasEnforceNonce

`func (o *SDKMessageRequest) HasEnforceNonce() bool`

HasEnforceNonce returns a boolean if a field has been set.

### SetEnforceNonceNil

`func (o *SDKMessageRequest) SetEnforceNonceNil(b bool)`

 SetEnforceNonceNil sets the value for EnforceNonce to be an explicit nil

### UnsetEnforceNonce
`func (o *SDKMessageRequest) UnsetEnforceNonce()`

UnsetEnforceNonce ensures that no value is present for EnforceNonce, not even an explicit nil
### GetTts

`func (o *SDKMessageRequest) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *SDKMessageRequest) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *SDKMessageRequest) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *SDKMessageRequest) HasTts() bool`

HasTts returns a boolean if a field has been set.

### SetTtsNil

`func (o *SDKMessageRequest) SetTtsNil(b bool)`

 SetTtsNil sets the value for Tts to be an explicit nil

### UnsetTts
`func (o *SDKMessageRequest) UnsetTts()`

UnsetTts ensures that no value is present for Tts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


