# MessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Content** | **string** |  | 
**Mentions** | [**[]UserResponse**](UserResponse.md) |  | 
**MentionRoles** | **[]string** |  | 
**Attachments** | [**[]MessageAttachmentResponse**](MessageAttachmentResponse.md) |  | 
**Embeds** | [**[]MessageEmbedResponse**](MessageEmbedResponse.md) |  | 
**Timestamp** | **time.Time** |  | 
**EditedTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Flags** | **int32** |  | 
**Components** | [**[]BasicMessageResponseComponentsInner**](BasicMessageResponseComponentsInner.md) |  | 
**Resolved** | Pointer to [**NullableResolvedObjectsResponse**](ResolvedObjectsResponse.md) |  | [optional] 
**Stickers** | Pointer to [**[]GetSticker200Response**](GetSticker200Response.md) |  | [optional] 
**StickerItems** | Pointer to [**[]MessageStickerItemResponse**](MessageStickerItemResponse.md) |  | [optional] 
**Id** | **string** |  | 
**ChannelId** | **string** |  | 
**Author** | [**UserResponse**](UserResponse.md) |  | 
**Pinned** | **bool** |  | 
**MentionEveryone** | **bool** |  | 
**Tts** | **bool** |  | 
**Call** | Pointer to [**NullableMessageCallResponse**](MessageCallResponse.md) |  | [optional] 
**Activity** | Pointer to **map[string]interface{}** |  | [optional] 
**Application** | Pointer to [**NullableBasicApplicationResponse**](BasicApplicationResponse.md) |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**Interaction** | Pointer to [**NullableMessageInteractionResponse**](MessageInteractionResponse.md) |  | [optional] 
**Nonce** | Pointer to [**NullableBasicMessageResponseNonce**](BasicMessageResponseNonce.md) |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**MessageReference** | Pointer to [**NullableMessageReferenceResponse**](MessageReferenceResponse.md) |  | [optional] 
**Thread** | Pointer to [**NullableThreadResponse**](ThreadResponse.md) |  | [optional] 
**MentionChannels** | Pointer to [**[]MessageMentionChannelResponse**](MessageMentionChannelResponse.md) |  | [optional] 
**RoleSubscriptionData** | Pointer to [**NullableMessageRoleSubscriptionDataResponse**](MessageRoleSubscriptionDataResponse.md) |  | [optional] 
**PurchaseNotification** | Pointer to [**NullablePurchaseNotificationResponse**](PurchaseNotificationResponse.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**Poll** | Pointer to [**NullablePollResponse**](PollResponse.md) |  | [optional] 
**InteractionMetadata** | Pointer to [**NullableBasicMessageResponseInteractionMetadata**](BasicMessageResponseInteractionMetadata.md) |  | [optional] 
**MessageSnapshots** | Pointer to [**[]MessageSnapshotResponse**](MessageSnapshotResponse.md) |  | [optional] 
**Reactions** | Pointer to [**[]MessageReactionResponse**](MessageReactionResponse.md) |  | [optional] 
**ReferencedMessage** | Pointer to [**NullableBasicMessageResponse**](BasicMessageResponse.md) |  | [optional] 

## Methods

### NewMessageResponse

`func NewMessageResponse(type_ int32, content string, mentions []UserResponse, mentionRoles []string, attachments []MessageAttachmentResponse, embeds []MessageEmbedResponse, timestamp time.Time, flags int32, components []BasicMessageResponseComponentsInner, id string, channelId string, author UserResponse, pinned bool, mentionEveryone bool, tts bool, ) *MessageResponse`

NewMessageResponse instantiates a new MessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResponseWithDefaults

`func NewMessageResponseWithDefaults() *MessageResponse`

NewMessageResponseWithDefaults instantiates a new MessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetContent

`func (o *MessageResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetMentions

`func (o *MessageResponse) GetMentions() []UserResponse`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *MessageResponse) GetMentionsOk() (*[]UserResponse, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *MessageResponse) SetMentions(v []UserResponse)`

SetMentions sets Mentions field to given value.


### GetMentionRoles

`func (o *MessageResponse) GetMentionRoles() []string`

GetMentionRoles returns the MentionRoles field if non-nil, zero value otherwise.

### GetMentionRolesOk

`func (o *MessageResponse) GetMentionRolesOk() (*[]string, bool)`

GetMentionRolesOk returns a tuple with the MentionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionRoles

`func (o *MessageResponse) SetMentionRoles(v []string)`

SetMentionRoles sets MentionRoles field to given value.


### GetAttachments

`func (o *MessageResponse) GetAttachments() []MessageAttachmentResponse`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageResponse) GetAttachmentsOk() (*[]MessageAttachmentResponse, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageResponse) SetAttachments(v []MessageAttachmentResponse)`

SetAttachments sets Attachments field to given value.


### GetEmbeds

`func (o *MessageResponse) GetEmbeds() []MessageEmbedResponse`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *MessageResponse) GetEmbedsOk() (*[]MessageEmbedResponse, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *MessageResponse) SetEmbeds(v []MessageEmbedResponse)`

SetEmbeds sets Embeds field to given value.


### GetTimestamp

`func (o *MessageResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEditedTimestamp

`func (o *MessageResponse) GetEditedTimestamp() time.Time`

GetEditedTimestamp returns the EditedTimestamp field if non-nil, zero value otherwise.

### GetEditedTimestampOk

`func (o *MessageResponse) GetEditedTimestampOk() (*time.Time, bool)`

GetEditedTimestampOk returns a tuple with the EditedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTimestamp

`func (o *MessageResponse) SetEditedTimestamp(v time.Time)`

SetEditedTimestamp sets EditedTimestamp field to given value.

### HasEditedTimestamp

`func (o *MessageResponse) HasEditedTimestamp() bool`

HasEditedTimestamp returns a boolean if a field has been set.

### SetEditedTimestampNil

`func (o *MessageResponse) SetEditedTimestampNil(b bool)`

 SetEditedTimestampNil sets the value for EditedTimestamp to be an explicit nil

### UnsetEditedTimestamp
`func (o *MessageResponse) UnsetEditedTimestamp()`

UnsetEditedTimestamp ensures that no value is present for EditedTimestamp, not even an explicit nil
### GetFlags

`func (o *MessageResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetComponents

`func (o *MessageResponse) GetComponents() []BasicMessageResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *MessageResponse) GetComponentsOk() (*[]BasicMessageResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *MessageResponse) SetComponents(v []BasicMessageResponseComponentsInner)`

SetComponents sets Components field to given value.


### GetResolved

`func (o *MessageResponse) GetResolved() ResolvedObjectsResponse`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MessageResponse) GetResolvedOk() (*ResolvedObjectsResponse, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MessageResponse) SetResolved(v ResolvedObjectsResponse)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *MessageResponse) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *MessageResponse) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *MessageResponse) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetStickers

`func (o *MessageResponse) GetStickers() []GetSticker200Response`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *MessageResponse) GetStickersOk() (*[]GetSticker200Response, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *MessageResponse) SetStickers(v []GetSticker200Response)`

SetStickers sets Stickers field to given value.

### HasStickers

`func (o *MessageResponse) HasStickers() bool`

HasStickers returns a boolean if a field has been set.

### SetStickersNil

`func (o *MessageResponse) SetStickersNil(b bool)`

 SetStickersNil sets the value for Stickers to be an explicit nil

### UnsetStickers
`func (o *MessageResponse) UnsetStickers()`

UnsetStickers ensures that no value is present for Stickers, not even an explicit nil
### GetStickerItems

`func (o *MessageResponse) GetStickerItems() []MessageStickerItemResponse`

GetStickerItems returns the StickerItems field if non-nil, zero value otherwise.

### GetStickerItemsOk

`func (o *MessageResponse) GetStickerItemsOk() (*[]MessageStickerItemResponse, bool)`

GetStickerItemsOk returns a tuple with the StickerItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerItems

`func (o *MessageResponse) SetStickerItems(v []MessageStickerItemResponse)`

SetStickerItems sets StickerItems field to given value.

### HasStickerItems

`func (o *MessageResponse) HasStickerItems() bool`

HasStickerItems returns a boolean if a field has been set.

### SetStickerItemsNil

`func (o *MessageResponse) SetStickerItemsNil(b bool)`

 SetStickerItemsNil sets the value for StickerItems to be an explicit nil

### UnsetStickerItems
`func (o *MessageResponse) UnsetStickerItems()`

UnsetStickerItems ensures that no value is present for StickerItems, not even an explicit nil
### GetId

`func (o *MessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetChannelId

`func (o *MessageResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MessageResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MessageResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetAuthor

`func (o *MessageResponse) GetAuthor() UserResponse`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MessageResponse) GetAuthorOk() (*UserResponse, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MessageResponse) SetAuthor(v UserResponse)`

SetAuthor sets Author field to given value.


### GetPinned

`func (o *MessageResponse) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *MessageResponse) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *MessageResponse) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetMentionEveryone

`func (o *MessageResponse) GetMentionEveryone() bool`

GetMentionEveryone returns the MentionEveryone field if non-nil, zero value otherwise.

### GetMentionEveryoneOk

`func (o *MessageResponse) GetMentionEveryoneOk() (*bool, bool)`

GetMentionEveryoneOk returns a tuple with the MentionEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionEveryone

`func (o *MessageResponse) SetMentionEveryone(v bool)`

SetMentionEveryone sets MentionEveryone field to given value.


### GetTts

`func (o *MessageResponse) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *MessageResponse) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *MessageResponse) SetTts(v bool)`

SetTts sets Tts field to given value.


### GetCall

`func (o *MessageResponse) GetCall() MessageCallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *MessageResponse) GetCallOk() (*MessageCallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *MessageResponse) SetCall(v MessageCallResponse)`

SetCall sets Call field to given value.

### HasCall

`func (o *MessageResponse) HasCall() bool`

HasCall returns a boolean if a field has been set.

### SetCallNil

`func (o *MessageResponse) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *MessageResponse) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetActivity

`func (o *MessageResponse) GetActivity() map[string]interface{}`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MessageResponse) GetActivityOk() (*map[string]interface{}, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *MessageResponse) SetActivity(v map[string]interface{})`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *MessageResponse) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApplication

`func (o *MessageResponse) GetApplication() BasicApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MessageResponse) GetApplicationOk() (*BasicApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MessageResponse) SetApplication(v BasicApplicationResponse)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MessageResponse) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *MessageResponse) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *MessageResponse) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetApplicationId

`func (o *MessageResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MessageResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MessageResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *MessageResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetInteraction

`func (o *MessageResponse) GetInteraction() MessageInteractionResponse`

GetInteraction returns the Interaction field if non-nil, zero value otherwise.

### GetInteractionOk

`func (o *MessageResponse) GetInteractionOk() (*MessageInteractionResponse, bool)`

GetInteractionOk returns a tuple with the Interaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteraction

`func (o *MessageResponse) SetInteraction(v MessageInteractionResponse)`

SetInteraction sets Interaction field to given value.

### HasInteraction

`func (o *MessageResponse) HasInteraction() bool`

HasInteraction returns a boolean if a field has been set.

### SetInteractionNil

`func (o *MessageResponse) SetInteractionNil(b bool)`

 SetInteractionNil sets the value for Interaction to be an explicit nil

### UnsetInteraction
`func (o *MessageResponse) UnsetInteraction()`

UnsetInteraction ensures that no value is present for Interaction, not even an explicit nil
### GetNonce

`func (o *MessageResponse) GetNonce() BasicMessageResponseNonce`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *MessageResponse) GetNonceOk() (*BasicMessageResponseNonce, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *MessageResponse) SetNonce(v BasicMessageResponseNonce)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *MessageResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *MessageResponse) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *MessageResponse) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetWebhookId

`func (o *MessageResponse) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *MessageResponse) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *MessageResponse) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *MessageResponse) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetMessageReference

`func (o *MessageResponse) GetMessageReference() MessageReferenceResponse`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *MessageResponse) GetMessageReferenceOk() (*MessageReferenceResponse, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *MessageResponse) SetMessageReference(v MessageReferenceResponse)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *MessageResponse) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

### SetMessageReferenceNil

`func (o *MessageResponse) SetMessageReferenceNil(b bool)`

 SetMessageReferenceNil sets the value for MessageReference to be an explicit nil

### UnsetMessageReference
`func (o *MessageResponse) UnsetMessageReference()`

UnsetMessageReference ensures that no value is present for MessageReference, not even an explicit nil
### GetThread

`func (o *MessageResponse) GetThread() ThreadResponse`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *MessageResponse) GetThreadOk() (*ThreadResponse, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *MessageResponse) SetThread(v ThreadResponse)`

SetThread sets Thread field to given value.

### HasThread

`func (o *MessageResponse) HasThread() bool`

HasThread returns a boolean if a field has been set.

### SetThreadNil

`func (o *MessageResponse) SetThreadNil(b bool)`

 SetThreadNil sets the value for Thread to be an explicit nil

### UnsetThread
`func (o *MessageResponse) UnsetThread()`

UnsetThread ensures that no value is present for Thread, not even an explicit nil
### GetMentionChannels

`func (o *MessageResponse) GetMentionChannels() []*MessageMentionChannelResponse`

GetMentionChannels returns the MentionChannels field if non-nil, zero value otherwise.

### GetMentionChannelsOk

`func (o *MessageResponse) GetMentionChannelsOk() (*[]*MessageMentionChannelResponse, bool)`

GetMentionChannelsOk returns a tuple with the MentionChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionChannels

`func (o *MessageResponse) SetMentionChannels(v []*MessageMentionChannelResponse)`

SetMentionChannels sets MentionChannels field to given value.

### HasMentionChannels

`func (o *MessageResponse) HasMentionChannels() bool`

HasMentionChannels returns a boolean if a field has been set.

### SetMentionChannelsNil

`func (o *MessageResponse) SetMentionChannelsNil(b bool)`

 SetMentionChannelsNil sets the value for MentionChannels to be an explicit nil

### UnsetMentionChannels
`func (o *MessageResponse) UnsetMentionChannels()`

UnsetMentionChannels ensures that no value is present for MentionChannels, not even an explicit nil
### GetRoleSubscriptionData

`func (o *MessageResponse) GetRoleSubscriptionData() MessageRoleSubscriptionDataResponse`

GetRoleSubscriptionData returns the RoleSubscriptionData field if non-nil, zero value otherwise.

### GetRoleSubscriptionDataOk

`func (o *MessageResponse) GetRoleSubscriptionDataOk() (*MessageRoleSubscriptionDataResponse, bool)`

GetRoleSubscriptionDataOk returns a tuple with the RoleSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSubscriptionData

`func (o *MessageResponse) SetRoleSubscriptionData(v MessageRoleSubscriptionDataResponse)`

SetRoleSubscriptionData sets RoleSubscriptionData field to given value.

### HasRoleSubscriptionData

`func (o *MessageResponse) HasRoleSubscriptionData() bool`

HasRoleSubscriptionData returns a boolean if a field has been set.

### SetRoleSubscriptionDataNil

`func (o *MessageResponse) SetRoleSubscriptionDataNil(b bool)`

 SetRoleSubscriptionDataNil sets the value for RoleSubscriptionData to be an explicit nil

### UnsetRoleSubscriptionData
`func (o *MessageResponse) UnsetRoleSubscriptionData()`

UnsetRoleSubscriptionData ensures that no value is present for RoleSubscriptionData, not even an explicit nil
### GetPurchaseNotification

`func (o *MessageResponse) GetPurchaseNotification() PurchaseNotificationResponse`

GetPurchaseNotification returns the PurchaseNotification field if non-nil, zero value otherwise.

### GetPurchaseNotificationOk

`func (o *MessageResponse) GetPurchaseNotificationOk() (*PurchaseNotificationResponse, bool)`

GetPurchaseNotificationOk returns a tuple with the PurchaseNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseNotification

`func (o *MessageResponse) SetPurchaseNotification(v PurchaseNotificationResponse)`

SetPurchaseNotification sets PurchaseNotification field to given value.

### HasPurchaseNotification

`func (o *MessageResponse) HasPurchaseNotification() bool`

HasPurchaseNotification returns a boolean if a field has been set.

### SetPurchaseNotificationNil

`func (o *MessageResponse) SetPurchaseNotificationNil(b bool)`

 SetPurchaseNotificationNil sets the value for PurchaseNotification to be an explicit nil

### UnsetPurchaseNotification
`func (o *MessageResponse) UnsetPurchaseNotification()`

UnsetPurchaseNotification ensures that no value is present for PurchaseNotification, not even an explicit nil
### GetPosition

`func (o *MessageResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MessageResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MessageResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MessageResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *MessageResponse) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *MessageResponse) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetPoll

`func (o *MessageResponse) GetPoll() PollResponse`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *MessageResponse) GetPollOk() (*PollResponse, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *MessageResponse) SetPoll(v PollResponse)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *MessageResponse) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *MessageResponse) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *MessageResponse) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
### GetInteractionMetadata

`func (o *MessageResponse) GetInteractionMetadata() BasicMessageResponseInteractionMetadata`

GetInteractionMetadata returns the InteractionMetadata field if non-nil, zero value otherwise.

### GetInteractionMetadataOk

`func (o *MessageResponse) GetInteractionMetadataOk() (*BasicMessageResponseInteractionMetadata, bool)`

GetInteractionMetadataOk returns a tuple with the InteractionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionMetadata

`func (o *MessageResponse) SetInteractionMetadata(v BasicMessageResponseInteractionMetadata)`

SetInteractionMetadata sets InteractionMetadata field to given value.

### HasInteractionMetadata

`func (o *MessageResponse) HasInteractionMetadata() bool`

HasInteractionMetadata returns a boolean if a field has been set.

### SetInteractionMetadataNil

`func (o *MessageResponse) SetInteractionMetadataNil(b bool)`

 SetInteractionMetadataNil sets the value for InteractionMetadata to be an explicit nil

### UnsetInteractionMetadata
`func (o *MessageResponse) UnsetInteractionMetadata()`

UnsetInteractionMetadata ensures that no value is present for InteractionMetadata, not even an explicit nil
### GetMessageSnapshots

`func (o *MessageResponse) GetMessageSnapshots() []MessageSnapshotResponse`

GetMessageSnapshots returns the MessageSnapshots field if non-nil, zero value otherwise.

### GetMessageSnapshotsOk

`func (o *MessageResponse) GetMessageSnapshotsOk() (*[]MessageSnapshotResponse, bool)`

GetMessageSnapshotsOk returns a tuple with the MessageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSnapshots

`func (o *MessageResponse) SetMessageSnapshots(v []MessageSnapshotResponse)`

SetMessageSnapshots sets MessageSnapshots field to given value.

### HasMessageSnapshots

`func (o *MessageResponse) HasMessageSnapshots() bool`

HasMessageSnapshots returns a boolean if a field has been set.

### SetMessageSnapshotsNil

`func (o *MessageResponse) SetMessageSnapshotsNil(b bool)`

 SetMessageSnapshotsNil sets the value for MessageSnapshots to be an explicit nil

### UnsetMessageSnapshots
`func (o *MessageResponse) UnsetMessageSnapshots()`

UnsetMessageSnapshots ensures that no value is present for MessageSnapshots, not even an explicit nil
### GetReactions

`func (o *MessageResponse) GetReactions() []MessageReactionResponse`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *MessageResponse) GetReactionsOk() (*[]MessageReactionResponse, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *MessageResponse) SetReactions(v []MessageReactionResponse)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *MessageResponse) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### SetReactionsNil

`func (o *MessageResponse) SetReactionsNil(b bool)`

 SetReactionsNil sets the value for Reactions to be an explicit nil

### UnsetReactions
`func (o *MessageResponse) UnsetReactions()`

UnsetReactions ensures that no value is present for Reactions, not even an explicit nil
### GetReferencedMessage

`func (o *MessageResponse) GetReferencedMessage() BasicMessageResponse`

GetReferencedMessage returns the ReferencedMessage field if non-nil, zero value otherwise.

### GetReferencedMessageOk

`func (o *MessageResponse) GetReferencedMessageOk() (*BasicMessageResponse, bool)`

GetReferencedMessageOk returns a tuple with the ReferencedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedMessage

`func (o *MessageResponse) SetReferencedMessage(v BasicMessageResponse)`

SetReferencedMessage sets ReferencedMessage field to given value.

### HasReferencedMessage

`func (o *MessageResponse) HasReferencedMessage() bool`

HasReferencedMessage returns a boolean if a field has been set.

### SetReferencedMessageNil

`func (o *MessageResponse) SetReferencedMessageNil(b bool)`

 SetReferencedMessageNil sets the value for ReferencedMessage to be an explicit nil

### UnsetReferencedMessage
`func (o *MessageResponse) UnsetReferencedMessage()`

UnsetReferencedMessage ensures that no value is present for ReferencedMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


