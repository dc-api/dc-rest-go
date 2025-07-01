# BasicMessageResponse

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

## Methods

### NewBasicMessageResponse

`func NewBasicMessageResponse(type_ int32, content string, mentions []UserResponse, mentionRoles []string, attachments []MessageAttachmentResponse, embeds []MessageEmbedResponse, timestamp time.Time, flags int32, components []BasicMessageResponseComponentsInner, id string, channelId string, author UserResponse, pinned bool, mentionEveryone bool, tts bool, ) *BasicMessageResponse`

NewBasicMessageResponse instantiates a new BasicMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicMessageResponseWithDefaults

`func NewBasicMessageResponseWithDefaults() *BasicMessageResponse`

NewBasicMessageResponseWithDefaults instantiates a new BasicMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BasicMessageResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicMessageResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicMessageResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetContent

`func (o *BasicMessageResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BasicMessageResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BasicMessageResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetMentions

`func (o *BasicMessageResponse) GetMentions() []UserResponse`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *BasicMessageResponse) GetMentionsOk() (*[]UserResponse, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *BasicMessageResponse) SetMentions(v []UserResponse)`

SetMentions sets Mentions field to given value.


### GetMentionRoles

`func (o *BasicMessageResponse) GetMentionRoles() []string`

GetMentionRoles returns the MentionRoles field if non-nil, zero value otherwise.

### GetMentionRolesOk

`func (o *BasicMessageResponse) GetMentionRolesOk() (*[]string, bool)`

GetMentionRolesOk returns a tuple with the MentionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionRoles

`func (o *BasicMessageResponse) SetMentionRoles(v []string)`

SetMentionRoles sets MentionRoles field to given value.


### GetAttachments

`func (o *BasicMessageResponse) GetAttachments() []MessageAttachmentResponse`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BasicMessageResponse) GetAttachmentsOk() (*[]MessageAttachmentResponse, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BasicMessageResponse) SetAttachments(v []MessageAttachmentResponse)`

SetAttachments sets Attachments field to given value.


### GetEmbeds

`func (o *BasicMessageResponse) GetEmbeds() []MessageEmbedResponse`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *BasicMessageResponse) GetEmbedsOk() (*[]MessageEmbedResponse, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *BasicMessageResponse) SetEmbeds(v []MessageEmbedResponse)`

SetEmbeds sets Embeds field to given value.


### GetTimestamp

`func (o *BasicMessageResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BasicMessageResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BasicMessageResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEditedTimestamp

`func (o *BasicMessageResponse) GetEditedTimestamp() time.Time`

GetEditedTimestamp returns the EditedTimestamp field if non-nil, zero value otherwise.

### GetEditedTimestampOk

`func (o *BasicMessageResponse) GetEditedTimestampOk() (*time.Time, bool)`

GetEditedTimestampOk returns a tuple with the EditedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTimestamp

`func (o *BasicMessageResponse) SetEditedTimestamp(v time.Time)`

SetEditedTimestamp sets EditedTimestamp field to given value.

### HasEditedTimestamp

`func (o *BasicMessageResponse) HasEditedTimestamp() bool`

HasEditedTimestamp returns a boolean if a field has been set.

### SetEditedTimestampNil

`func (o *BasicMessageResponse) SetEditedTimestampNil(b bool)`

 SetEditedTimestampNil sets the value for EditedTimestamp to be an explicit nil

### UnsetEditedTimestamp
`func (o *BasicMessageResponse) UnsetEditedTimestamp()`

UnsetEditedTimestamp ensures that no value is present for EditedTimestamp, not even an explicit nil
### GetFlags

`func (o *BasicMessageResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BasicMessageResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BasicMessageResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetComponents

`func (o *BasicMessageResponse) GetComponents() []BasicMessageResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BasicMessageResponse) GetComponentsOk() (*[]BasicMessageResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BasicMessageResponse) SetComponents(v []BasicMessageResponseComponentsInner)`

SetComponents sets Components field to given value.


### GetResolved

`func (o *BasicMessageResponse) GetResolved() ResolvedObjectsResponse`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *BasicMessageResponse) GetResolvedOk() (*ResolvedObjectsResponse, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *BasicMessageResponse) SetResolved(v ResolvedObjectsResponse)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *BasicMessageResponse) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *BasicMessageResponse) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *BasicMessageResponse) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetStickers

`func (o *BasicMessageResponse) GetStickers() []GetSticker200Response`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *BasicMessageResponse) GetStickersOk() (*[]GetSticker200Response, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *BasicMessageResponse) SetStickers(v []GetSticker200Response)`

SetStickers sets Stickers field to given value.

### HasStickers

`func (o *BasicMessageResponse) HasStickers() bool`

HasStickers returns a boolean if a field has been set.

### SetStickersNil

`func (o *BasicMessageResponse) SetStickersNil(b bool)`

 SetStickersNil sets the value for Stickers to be an explicit nil

### UnsetStickers
`func (o *BasicMessageResponse) UnsetStickers()`

UnsetStickers ensures that no value is present for Stickers, not even an explicit nil
### GetStickerItems

`func (o *BasicMessageResponse) GetStickerItems() []MessageStickerItemResponse`

GetStickerItems returns the StickerItems field if non-nil, zero value otherwise.

### GetStickerItemsOk

`func (o *BasicMessageResponse) GetStickerItemsOk() (*[]MessageStickerItemResponse, bool)`

GetStickerItemsOk returns a tuple with the StickerItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerItems

`func (o *BasicMessageResponse) SetStickerItems(v []MessageStickerItemResponse)`

SetStickerItems sets StickerItems field to given value.

### HasStickerItems

`func (o *BasicMessageResponse) HasStickerItems() bool`

HasStickerItems returns a boolean if a field has been set.

### SetStickerItemsNil

`func (o *BasicMessageResponse) SetStickerItemsNil(b bool)`

 SetStickerItemsNil sets the value for StickerItems to be an explicit nil

### UnsetStickerItems
`func (o *BasicMessageResponse) UnsetStickerItems()`

UnsetStickerItems ensures that no value is present for StickerItems, not even an explicit nil
### GetId

`func (o *BasicMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicMessageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetChannelId

`func (o *BasicMessageResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *BasicMessageResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *BasicMessageResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetAuthor

`func (o *BasicMessageResponse) GetAuthor() UserResponse`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BasicMessageResponse) GetAuthorOk() (*UserResponse, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BasicMessageResponse) SetAuthor(v UserResponse)`

SetAuthor sets Author field to given value.


### GetPinned

`func (o *BasicMessageResponse) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *BasicMessageResponse) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *BasicMessageResponse) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetMentionEveryone

`func (o *BasicMessageResponse) GetMentionEveryone() bool`

GetMentionEveryone returns the MentionEveryone field if non-nil, zero value otherwise.

### GetMentionEveryoneOk

`func (o *BasicMessageResponse) GetMentionEveryoneOk() (*bool, bool)`

GetMentionEveryoneOk returns a tuple with the MentionEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionEveryone

`func (o *BasicMessageResponse) SetMentionEveryone(v bool)`

SetMentionEveryone sets MentionEveryone field to given value.


### GetTts

`func (o *BasicMessageResponse) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *BasicMessageResponse) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *BasicMessageResponse) SetTts(v bool)`

SetTts sets Tts field to given value.


### GetCall

`func (o *BasicMessageResponse) GetCall() MessageCallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *BasicMessageResponse) GetCallOk() (*MessageCallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *BasicMessageResponse) SetCall(v MessageCallResponse)`

SetCall sets Call field to given value.

### HasCall

`func (o *BasicMessageResponse) HasCall() bool`

HasCall returns a boolean if a field has been set.

### SetCallNil

`func (o *BasicMessageResponse) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *BasicMessageResponse) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetActivity

`func (o *BasicMessageResponse) GetActivity() map[string]interface{}`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *BasicMessageResponse) GetActivityOk() (*map[string]interface{}, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *BasicMessageResponse) SetActivity(v map[string]interface{})`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *BasicMessageResponse) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApplication

`func (o *BasicMessageResponse) GetApplication() BasicApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *BasicMessageResponse) GetApplicationOk() (*BasicApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *BasicMessageResponse) SetApplication(v BasicApplicationResponse)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *BasicMessageResponse) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *BasicMessageResponse) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *BasicMessageResponse) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetApplicationId

`func (o *BasicMessageResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *BasicMessageResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *BasicMessageResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *BasicMessageResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetInteraction

`func (o *BasicMessageResponse) GetInteraction() MessageInteractionResponse`

GetInteraction returns the Interaction field if non-nil, zero value otherwise.

### GetInteractionOk

`func (o *BasicMessageResponse) GetInteractionOk() (*MessageInteractionResponse, bool)`

GetInteractionOk returns a tuple with the Interaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteraction

`func (o *BasicMessageResponse) SetInteraction(v MessageInteractionResponse)`

SetInteraction sets Interaction field to given value.

### HasInteraction

`func (o *BasicMessageResponse) HasInteraction() bool`

HasInteraction returns a boolean if a field has been set.

### SetInteractionNil

`func (o *BasicMessageResponse) SetInteractionNil(b bool)`

 SetInteractionNil sets the value for Interaction to be an explicit nil

### UnsetInteraction
`func (o *BasicMessageResponse) UnsetInteraction()`

UnsetInteraction ensures that no value is present for Interaction, not even an explicit nil
### GetNonce

`func (o *BasicMessageResponse) GetNonce() BasicMessageResponseNonce`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BasicMessageResponse) GetNonceOk() (*BasicMessageResponseNonce, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BasicMessageResponse) SetNonce(v BasicMessageResponseNonce)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BasicMessageResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *BasicMessageResponse) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *BasicMessageResponse) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetWebhookId

`func (o *BasicMessageResponse) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *BasicMessageResponse) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *BasicMessageResponse) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *BasicMessageResponse) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetMessageReference

`func (o *BasicMessageResponse) GetMessageReference() MessageReferenceResponse`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *BasicMessageResponse) GetMessageReferenceOk() (*MessageReferenceResponse, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *BasicMessageResponse) SetMessageReference(v MessageReferenceResponse)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *BasicMessageResponse) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

### SetMessageReferenceNil

`func (o *BasicMessageResponse) SetMessageReferenceNil(b bool)`

 SetMessageReferenceNil sets the value for MessageReference to be an explicit nil

### UnsetMessageReference
`func (o *BasicMessageResponse) UnsetMessageReference()`

UnsetMessageReference ensures that no value is present for MessageReference, not even an explicit nil
### GetThread

`func (o *BasicMessageResponse) GetThread() ThreadResponse`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *BasicMessageResponse) GetThreadOk() (*ThreadResponse, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *BasicMessageResponse) SetThread(v ThreadResponse)`

SetThread sets Thread field to given value.

### HasThread

`func (o *BasicMessageResponse) HasThread() bool`

HasThread returns a boolean if a field has been set.

### SetThreadNil

`func (o *BasicMessageResponse) SetThreadNil(b bool)`

 SetThreadNil sets the value for Thread to be an explicit nil

### UnsetThread
`func (o *BasicMessageResponse) UnsetThread()`

UnsetThread ensures that no value is present for Thread, not even an explicit nil
### GetMentionChannels

`func (o *BasicMessageResponse) GetMentionChannels() []*MessageMentionChannelResponse`

GetMentionChannels returns the MentionChannels field if non-nil, zero value otherwise.

### GetMentionChannelsOk

`func (o *BasicMessageResponse) GetMentionChannelsOk() (*[]*MessageMentionChannelResponse, bool)`

GetMentionChannelsOk returns a tuple with the MentionChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionChannels

`func (o *BasicMessageResponse) SetMentionChannels(v []*MessageMentionChannelResponse)`

SetMentionChannels sets MentionChannels field to given value.

### HasMentionChannels

`func (o *BasicMessageResponse) HasMentionChannels() bool`

HasMentionChannels returns a boolean if a field has been set.

### SetMentionChannelsNil

`func (o *BasicMessageResponse) SetMentionChannelsNil(b bool)`

 SetMentionChannelsNil sets the value for MentionChannels to be an explicit nil

### UnsetMentionChannels
`func (o *BasicMessageResponse) UnsetMentionChannels()`

UnsetMentionChannels ensures that no value is present for MentionChannels, not even an explicit nil
### GetRoleSubscriptionData

`func (o *BasicMessageResponse) GetRoleSubscriptionData() MessageRoleSubscriptionDataResponse`

GetRoleSubscriptionData returns the RoleSubscriptionData field if non-nil, zero value otherwise.

### GetRoleSubscriptionDataOk

`func (o *BasicMessageResponse) GetRoleSubscriptionDataOk() (*MessageRoleSubscriptionDataResponse, bool)`

GetRoleSubscriptionDataOk returns a tuple with the RoleSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSubscriptionData

`func (o *BasicMessageResponse) SetRoleSubscriptionData(v MessageRoleSubscriptionDataResponse)`

SetRoleSubscriptionData sets RoleSubscriptionData field to given value.

### HasRoleSubscriptionData

`func (o *BasicMessageResponse) HasRoleSubscriptionData() bool`

HasRoleSubscriptionData returns a boolean if a field has been set.

### SetRoleSubscriptionDataNil

`func (o *BasicMessageResponse) SetRoleSubscriptionDataNil(b bool)`

 SetRoleSubscriptionDataNil sets the value for RoleSubscriptionData to be an explicit nil

### UnsetRoleSubscriptionData
`func (o *BasicMessageResponse) UnsetRoleSubscriptionData()`

UnsetRoleSubscriptionData ensures that no value is present for RoleSubscriptionData, not even an explicit nil
### GetPurchaseNotification

`func (o *BasicMessageResponse) GetPurchaseNotification() PurchaseNotificationResponse`

GetPurchaseNotification returns the PurchaseNotification field if non-nil, zero value otherwise.

### GetPurchaseNotificationOk

`func (o *BasicMessageResponse) GetPurchaseNotificationOk() (*PurchaseNotificationResponse, bool)`

GetPurchaseNotificationOk returns a tuple with the PurchaseNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseNotification

`func (o *BasicMessageResponse) SetPurchaseNotification(v PurchaseNotificationResponse)`

SetPurchaseNotification sets PurchaseNotification field to given value.

### HasPurchaseNotification

`func (o *BasicMessageResponse) HasPurchaseNotification() bool`

HasPurchaseNotification returns a boolean if a field has been set.

### SetPurchaseNotificationNil

`func (o *BasicMessageResponse) SetPurchaseNotificationNil(b bool)`

 SetPurchaseNotificationNil sets the value for PurchaseNotification to be an explicit nil

### UnsetPurchaseNotification
`func (o *BasicMessageResponse) UnsetPurchaseNotification()`

UnsetPurchaseNotification ensures that no value is present for PurchaseNotification, not even an explicit nil
### GetPosition

`func (o *BasicMessageResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BasicMessageResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BasicMessageResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BasicMessageResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *BasicMessageResponse) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *BasicMessageResponse) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetPoll

`func (o *BasicMessageResponse) GetPoll() PollResponse`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *BasicMessageResponse) GetPollOk() (*PollResponse, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *BasicMessageResponse) SetPoll(v PollResponse)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *BasicMessageResponse) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *BasicMessageResponse) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *BasicMessageResponse) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
### GetInteractionMetadata

`func (o *BasicMessageResponse) GetInteractionMetadata() BasicMessageResponseInteractionMetadata`

GetInteractionMetadata returns the InteractionMetadata field if non-nil, zero value otherwise.

### GetInteractionMetadataOk

`func (o *BasicMessageResponse) GetInteractionMetadataOk() (*BasicMessageResponseInteractionMetadata, bool)`

GetInteractionMetadataOk returns a tuple with the InteractionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionMetadata

`func (o *BasicMessageResponse) SetInteractionMetadata(v BasicMessageResponseInteractionMetadata)`

SetInteractionMetadata sets InteractionMetadata field to given value.

### HasInteractionMetadata

`func (o *BasicMessageResponse) HasInteractionMetadata() bool`

HasInteractionMetadata returns a boolean if a field has been set.

### SetInteractionMetadataNil

`func (o *BasicMessageResponse) SetInteractionMetadataNil(b bool)`

 SetInteractionMetadataNil sets the value for InteractionMetadata to be an explicit nil

### UnsetInteractionMetadata
`func (o *BasicMessageResponse) UnsetInteractionMetadata()`

UnsetInteractionMetadata ensures that no value is present for InteractionMetadata, not even an explicit nil
### GetMessageSnapshots

`func (o *BasicMessageResponse) GetMessageSnapshots() []MessageSnapshotResponse`

GetMessageSnapshots returns the MessageSnapshots field if non-nil, zero value otherwise.

### GetMessageSnapshotsOk

`func (o *BasicMessageResponse) GetMessageSnapshotsOk() (*[]MessageSnapshotResponse, bool)`

GetMessageSnapshotsOk returns a tuple with the MessageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSnapshots

`func (o *BasicMessageResponse) SetMessageSnapshots(v []MessageSnapshotResponse)`

SetMessageSnapshots sets MessageSnapshots field to given value.

### HasMessageSnapshots

`func (o *BasicMessageResponse) HasMessageSnapshots() bool`

HasMessageSnapshots returns a boolean if a field has been set.

### SetMessageSnapshotsNil

`func (o *BasicMessageResponse) SetMessageSnapshotsNil(b bool)`

 SetMessageSnapshotsNil sets the value for MessageSnapshots to be an explicit nil

### UnsetMessageSnapshots
`func (o *BasicMessageResponse) UnsetMessageSnapshots()`

UnsetMessageSnapshots ensures that no value is present for MessageSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


