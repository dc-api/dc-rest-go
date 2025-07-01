# BaseCreateMessageCreateRequest

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

## Methods

### NewBaseCreateMessageCreateRequest

`func NewBaseCreateMessageCreateRequest() *BaseCreateMessageCreateRequest`

NewBaseCreateMessageCreateRequest instantiates a new BaseCreateMessageCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCreateMessageCreateRequestWithDefaults

`func NewBaseCreateMessageCreateRequestWithDefaults() *BaseCreateMessageCreateRequest`

NewBaseCreateMessageCreateRequestWithDefaults instantiates a new BaseCreateMessageCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *BaseCreateMessageCreateRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BaseCreateMessageCreateRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BaseCreateMessageCreateRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BaseCreateMessageCreateRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *BaseCreateMessageCreateRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *BaseCreateMessageCreateRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEmbeds

`func (o *BaseCreateMessageCreateRequest) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *BaseCreateMessageCreateRequest) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *BaseCreateMessageCreateRequest) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *BaseCreateMessageCreateRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### SetEmbedsNil

`func (o *BaseCreateMessageCreateRequest) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *BaseCreateMessageCreateRequest) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetAllowedMentions

`func (o *BaseCreateMessageCreateRequest) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *BaseCreateMessageCreateRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *BaseCreateMessageCreateRequest) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *BaseCreateMessageCreateRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### SetAllowedMentionsNil

`func (o *BaseCreateMessageCreateRequest) SetAllowedMentionsNil(b bool)`

 SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil

### UnsetAllowedMentions
`func (o *BaseCreateMessageCreateRequest) UnsetAllowedMentions()`

UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
### GetStickerIds

`func (o *BaseCreateMessageCreateRequest) GetStickerIds() []string`

GetStickerIds returns the StickerIds field if non-nil, zero value otherwise.

### GetStickerIdsOk

`func (o *BaseCreateMessageCreateRequest) GetStickerIdsOk() (*[]string, bool)`

GetStickerIdsOk returns a tuple with the StickerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerIds

`func (o *BaseCreateMessageCreateRequest) SetStickerIds(v []string)`

SetStickerIds sets StickerIds field to given value.

### HasStickerIds

`func (o *BaseCreateMessageCreateRequest) HasStickerIds() bool`

HasStickerIds returns a boolean if a field has been set.

### SetStickerIdsNil

`func (o *BaseCreateMessageCreateRequest) SetStickerIdsNil(b bool)`

 SetStickerIdsNil sets the value for StickerIds to be an explicit nil

### UnsetStickerIds
`func (o *BaseCreateMessageCreateRequest) UnsetStickerIds()`

UnsetStickerIds ensures that no value is present for StickerIds, not even an explicit nil
### GetComponents

`func (o *BaseCreateMessageCreateRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BaseCreateMessageCreateRequest) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BaseCreateMessageCreateRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BaseCreateMessageCreateRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *BaseCreateMessageCreateRequest) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *BaseCreateMessageCreateRequest) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetFlags

`func (o *BaseCreateMessageCreateRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BaseCreateMessageCreateRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BaseCreateMessageCreateRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *BaseCreateMessageCreateRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *BaseCreateMessageCreateRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *BaseCreateMessageCreateRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAttachments

`func (o *BaseCreateMessageCreateRequest) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BaseCreateMessageCreateRequest) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BaseCreateMessageCreateRequest) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BaseCreateMessageCreateRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *BaseCreateMessageCreateRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *BaseCreateMessageCreateRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetPoll

`func (o *BaseCreateMessageCreateRequest) GetPoll() PollCreateRequest`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *BaseCreateMessageCreateRequest) GetPollOk() (*PollCreateRequest, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *BaseCreateMessageCreateRequest) SetPoll(v PollCreateRequest)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *BaseCreateMessageCreateRequest) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *BaseCreateMessageCreateRequest) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *BaseCreateMessageCreateRequest) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
### GetConfettiPotion

`func (o *BaseCreateMessageCreateRequest) GetConfettiPotion() map[string]interface{}`

GetConfettiPotion returns the ConfettiPotion field if non-nil, zero value otherwise.

### GetConfettiPotionOk

`func (o *BaseCreateMessageCreateRequest) GetConfettiPotionOk() (*map[string]interface{}, bool)`

GetConfettiPotionOk returns a tuple with the ConfettiPotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfettiPotion

`func (o *BaseCreateMessageCreateRequest) SetConfettiPotion(v map[string]interface{})`

SetConfettiPotion sets ConfettiPotion field to given value.

### HasConfettiPotion

`func (o *BaseCreateMessageCreateRequest) HasConfettiPotion() bool`

HasConfettiPotion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


