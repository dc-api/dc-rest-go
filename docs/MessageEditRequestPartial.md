# MessageEditRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**AllowedMentions** | Pointer to [**NullableMessageAllowedMentionsRequest**](MessageAllowedMentionsRequest.md) |  | [optional] 
**StickerIds** | Pointer to **[]string** |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 

## Methods

### NewMessageEditRequestPartial

`func NewMessageEditRequestPartial() *MessageEditRequestPartial`

NewMessageEditRequestPartial instantiates a new MessageEditRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageEditRequestPartialWithDefaults

`func NewMessageEditRequestPartialWithDefaults() *MessageEditRequestPartial`

NewMessageEditRequestPartialWithDefaults instantiates a new MessageEditRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MessageEditRequestPartial) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageEditRequestPartial) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageEditRequestPartial) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MessageEditRequestPartial) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MessageEditRequestPartial) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MessageEditRequestPartial) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEmbeds

`func (o *MessageEditRequestPartial) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *MessageEditRequestPartial) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *MessageEditRequestPartial) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *MessageEditRequestPartial) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### SetEmbedsNil

`func (o *MessageEditRequestPartial) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *MessageEditRequestPartial) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetFlags

`func (o *MessageEditRequestPartial) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageEditRequestPartial) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageEditRequestPartial) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageEditRequestPartial) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *MessageEditRequestPartial) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *MessageEditRequestPartial) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAllowedMentions

`func (o *MessageEditRequestPartial) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *MessageEditRequestPartial) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *MessageEditRequestPartial) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *MessageEditRequestPartial) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### SetAllowedMentionsNil

`func (o *MessageEditRequestPartial) SetAllowedMentionsNil(b bool)`

 SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil

### UnsetAllowedMentions
`func (o *MessageEditRequestPartial) UnsetAllowedMentions()`

UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
### GetStickerIds

`func (o *MessageEditRequestPartial) GetStickerIds() []string`

GetStickerIds returns the StickerIds field if non-nil, zero value otherwise.

### GetStickerIdsOk

`func (o *MessageEditRequestPartial) GetStickerIdsOk() (*[]string, bool)`

GetStickerIdsOk returns a tuple with the StickerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerIds

`func (o *MessageEditRequestPartial) SetStickerIds(v []string)`

SetStickerIds sets StickerIds field to given value.

### HasStickerIds

`func (o *MessageEditRequestPartial) HasStickerIds() bool`

HasStickerIds returns a boolean if a field has been set.

### SetStickerIdsNil

`func (o *MessageEditRequestPartial) SetStickerIdsNil(b bool)`

 SetStickerIdsNil sets the value for StickerIds to be an explicit nil

### UnsetStickerIds
`func (o *MessageEditRequestPartial) UnsetStickerIds()`

UnsetStickerIds ensures that no value is present for StickerIds, not even an explicit nil
### GetComponents

`func (o *MessageEditRequestPartial) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *MessageEditRequestPartial) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *MessageEditRequestPartial) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *MessageEditRequestPartial) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *MessageEditRequestPartial) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *MessageEditRequestPartial) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetAttachments

`func (o *MessageEditRequestPartial) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageEditRequestPartial) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageEditRequestPartial) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MessageEditRequestPartial) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *MessageEditRequestPartial) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *MessageEditRequestPartial) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


