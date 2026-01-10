# MessageResponseReferencedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageType**](MessageType.md) |  | 
**Content** | **string** |  | 
**Mentions** | [**[]UserResponse**](UserResponse.md) |  | 
**MentionRoles** | **[]string** |  | 
**Attachments** | [**[]MessageAttachmentResponse**](MessageAttachmentResponse.md) |  | 
**Embeds** | [**[]MessageEmbedResponse**](MessageEmbedResponse.md) |  | 
**Timestamp** | **time.Time** |  | 
**EditedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Flags** | **int32** |  | 
**Components** | [**[]BasicMessageResponseComponentsInner**](BasicMessageResponseComponentsInner.md) |  | 
**Stickers** | Pointer to [**[]GetSticker200Response**](GetSticker200Response.md) |  | [optional] 
**StickerItems** | Pointer to [**[]MessageStickerItemResponse**](MessageStickerItemResponse.md) |  | [optional] 
**Id** | **string** |  | 
**ChannelId** | **string** |  | 
**Author** | [**UserResponse**](UserResponse.md) |  | 
**Pinned** | **bool** |  | 
**MentionEveryone** | **bool** |  | 
**Tts** | **bool** |  | 
**Call** | Pointer to [**MessageCallResponse**](MessageCallResponse.md) |  | [optional] 
**Activity** | Pointer to **map[string]interface{}** |  | [optional] 
**Application** | Pointer to [**BasicApplicationResponse**](BasicApplicationResponse.md) |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**Interaction** | Pointer to [**MessageInteractionResponse**](MessageInteractionResponse.md) |  | [optional] 
**Nonce** | Pointer to [**BasicMessageResponseNonce**](BasicMessageResponseNonce.md) |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 
**MessageReference** | Pointer to [**MessageReferenceResponse**](MessageReferenceResponse.md) |  | [optional] 
**Thread** | Pointer to [**ThreadResponse**](ThreadResponse.md) |  | [optional] 
**MentionChannels** | Pointer to [**[]BasicMessageResponseMentionChannelsInner**](BasicMessageResponseMentionChannelsInner.md) |  | [optional] 
**RoleSubscriptionData** | Pointer to [**MessageRoleSubscriptionDataResponse**](MessageRoleSubscriptionDataResponse.md) |  | [optional] 
**PurchaseNotification** | Pointer to [**PurchaseNotificationResponse**](PurchaseNotificationResponse.md) |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Resolved** | Pointer to [**ResolvedObjectsResponse**](ResolvedObjectsResponse.md) |  | [optional] 
**Poll** | Pointer to [**PollResponse**](PollResponse.md) |  | [optional] 
**SharedClientTheme** | Pointer to [**BasicMessageResponseSharedClientTheme**](BasicMessageResponseSharedClientTheme.md) |  | [optional] 
**InteractionMetadata** | Pointer to [**BasicMessageResponseInteractionMetadata**](BasicMessageResponseInteractionMetadata.md) |  | [optional] 
**MessageSnapshots** | Pointer to [**[]MessageSnapshotResponse**](MessageSnapshotResponse.md) |  | [optional] 

## Methods

### NewMessageResponseReferencedMessage

`func NewMessageResponseReferencedMessage(type_ MessageType, content string, mentions []UserResponse, mentionRoles []string, attachments []MessageAttachmentResponse, embeds []MessageEmbedResponse, timestamp time.Time, flags int32, components []BasicMessageResponseComponentsInner, id string, channelId string, author UserResponse, pinned bool, mentionEveryone bool, tts bool, ) *MessageResponseReferencedMessage`

NewMessageResponseReferencedMessage instantiates a new MessageResponseReferencedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResponseReferencedMessageWithDefaults

`func NewMessageResponseReferencedMessageWithDefaults() *MessageResponseReferencedMessage`

NewMessageResponseReferencedMessageWithDefaults instantiates a new MessageResponseReferencedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageResponseReferencedMessage) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageResponseReferencedMessage) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageResponseReferencedMessage) SetType(v MessageType)`

SetType sets Type field to given value.


### GetContent

`func (o *MessageResponseReferencedMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageResponseReferencedMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageResponseReferencedMessage) SetContent(v string)`

SetContent sets Content field to given value.


### GetMentions

`func (o *MessageResponseReferencedMessage) GetMentions() []UserResponse`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *MessageResponseReferencedMessage) GetMentionsOk() (*[]UserResponse, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *MessageResponseReferencedMessage) SetMentions(v []UserResponse)`

SetMentions sets Mentions field to given value.


### GetMentionRoles

`func (o *MessageResponseReferencedMessage) GetMentionRoles() []string`

GetMentionRoles returns the MentionRoles field if non-nil, zero value otherwise.

### GetMentionRolesOk

`func (o *MessageResponseReferencedMessage) GetMentionRolesOk() (*[]string, bool)`

GetMentionRolesOk returns a tuple with the MentionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionRoles

`func (o *MessageResponseReferencedMessage) SetMentionRoles(v []string)`

SetMentionRoles sets MentionRoles field to given value.


### GetAttachments

`func (o *MessageResponseReferencedMessage) GetAttachments() []MessageAttachmentResponse`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageResponseReferencedMessage) GetAttachmentsOk() (*[]MessageAttachmentResponse, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageResponseReferencedMessage) SetAttachments(v []MessageAttachmentResponse)`

SetAttachments sets Attachments field to given value.


### GetEmbeds

`func (o *MessageResponseReferencedMessage) GetEmbeds() []MessageEmbedResponse`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *MessageResponseReferencedMessage) GetEmbedsOk() (*[]MessageEmbedResponse, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *MessageResponseReferencedMessage) SetEmbeds(v []MessageEmbedResponse)`

SetEmbeds sets Embeds field to given value.


### GetTimestamp

`func (o *MessageResponseReferencedMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageResponseReferencedMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageResponseReferencedMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEditedTimestamp

`func (o *MessageResponseReferencedMessage) GetEditedTimestamp() time.Time`

GetEditedTimestamp returns the EditedTimestamp field if non-nil, zero value otherwise.

### GetEditedTimestampOk

`func (o *MessageResponseReferencedMessage) GetEditedTimestampOk() (*time.Time, bool)`

GetEditedTimestampOk returns a tuple with the EditedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTimestamp

`func (o *MessageResponseReferencedMessage) SetEditedTimestamp(v time.Time)`

SetEditedTimestamp sets EditedTimestamp field to given value.

### HasEditedTimestamp

`func (o *MessageResponseReferencedMessage) HasEditedTimestamp() bool`

HasEditedTimestamp returns a boolean if a field has been set.

### GetFlags

`func (o *MessageResponseReferencedMessage) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageResponseReferencedMessage) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageResponseReferencedMessage) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetComponents

`func (o *MessageResponseReferencedMessage) GetComponents() []BasicMessageResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *MessageResponseReferencedMessage) GetComponentsOk() (*[]BasicMessageResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *MessageResponseReferencedMessage) SetComponents(v []BasicMessageResponseComponentsInner)`

SetComponents sets Components field to given value.


### GetStickers

`func (o *MessageResponseReferencedMessage) GetStickers() []GetSticker200Response`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *MessageResponseReferencedMessage) GetStickersOk() (*[]GetSticker200Response, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *MessageResponseReferencedMessage) SetStickers(v []GetSticker200Response)`

SetStickers sets Stickers field to given value.

### HasStickers

`func (o *MessageResponseReferencedMessage) HasStickers() bool`

HasStickers returns a boolean if a field has been set.

### GetStickerItems

`func (o *MessageResponseReferencedMessage) GetStickerItems() []MessageStickerItemResponse`

GetStickerItems returns the StickerItems field if non-nil, zero value otherwise.

### GetStickerItemsOk

`func (o *MessageResponseReferencedMessage) GetStickerItemsOk() (*[]MessageStickerItemResponse, bool)`

GetStickerItemsOk returns a tuple with the StickerItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerItems

`func (o *MessageResponseReferencedMessage) SetStickerItems(v []MessageStickerItemResponse)`

SetStickerItems sets StickerItems field to given value.

### HasStickerItems

`func (o *MessageResponseReferencedMessage) HasStickerItems() bool`

HasStickerItems returns a boolean if a field has been set.

### GetId

`func (o *MessageResponseReferencedMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageResponseReferencedMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageResponseReferencedMessage) SetId(v string)`

SetId sets Id field to given value.


### GetChannelId

`func (o *MessageResponseReferencedMessage) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MessageResponseReferencedMessage) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MessageResponseReferencedMessage) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetAuthor

`func (o *MessageResponseReferencedMessage) GetAuthor() UserResponse`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MessageResponseReferencedMessage) GetAuthorOk() (*UserResponse, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MessageResponseReferencedMessage) SetAuthor(v UserResponse)`

SetAuthor sets Author field to given value.


### GetPinned

`func (o *MessageResponseReferencedMessage) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *MessageResponseReferencedMessage) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *MessageResponseReferencedMessage) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetMentionEveryone

`func (o *MessageResponseReferencedMessage) GetMentionEveryone() bool`

GetMentionEveryone returns the MentionEveryone field if non-nil, zero value otherwise.

### GetMentionEveryoneOk

`func (o *MessageResponseReferencedMessage) GetMentionEveryoneOk() (*bool, bool)`

GetMentionEveryoneOk returns a tuple with the MentionEveryone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionEveryone

`func (o *MessageResponseReferencedMessage) SetMentionEveryone(v bool)`

SetMentionEveryone sets MentionEveryone field to given value.


### GetTts

`func (o *MessageResponseReferencedMessage) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *MessageResponseReferencedMessage) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *MessageResponseReferencedMessage) SetTts(v bool)`

SetTts sets Tts field to given value.


### GetCall

`func (o *MessageResponseReferencedMessage) GetCall() MessageCallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *MessageResponseReferencedMessage) GetCallOk() (*MessageCallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *MessageResponseReferencedMessage) SetCall(v MessageCallResponse)`

SetCall sets Call field to given value.

### HasCall

`func (o *MessageResponseReferencedMessage) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetActivity

`func (o *MessageResponseReferencedMessage) GetActivity() map[string]interface{}`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MessageResponseReferencedMessage) GetActivityOk() (*map[string]interface{}, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *MessageResponseReferencedMessage) SetActivity(v map[string]interface{})`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *MessageResponseReferencedMessage) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApplication

`func (o *MessageResponseReferencedMessage) GetApplication() BasicApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MessageResponseReferencedMessage) GetApplicationOk() (*BasicApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MessageResponseReferencedMessage) SetApplication(v BasicApplicationResponse)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MessageResponseReferencedMessage) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationId

`func (o *MessageResponseReferencedMessage) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MessageResponseReferencedMessage) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MessageResponseReferencedMessage) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *MessageResponseReferencedMessage) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetInteraction

`func (o *MessageResponseReferencedMessage) GetInteraction() MessageInteractionResponse`

GetInteraction returns the Interaction field if non-nil, zero value otherwise.

### GetInteractionOk

`func (o *MessageResponseReferencedMessage) GetInteractionOk() (*MessageInteractionResponse, bool)`

GetInteractionOk returns a tuple with the Interaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteraction

`func (o *MessageResponseReferencedMessage) SetInteraction(v MessageInteractionResponse)`

SetInteraction sets Interaction field to given value.

### HasInteraction

`func (o *MessageResponseReferencedMessage) HasInteraction() bool`

HasInteraction returns a boolean if a field has been set.

### GetNonce

`func (o *MessageResponseReferencedMessage) GetNonce() BasicMessageResponseNonce`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *MessageResponseReferencedMessage) GetNonceOk() (*BasicMessageResponseNonce, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *MessageResponseReferencedMessage) SetNonce(v BasicMessageResponseNonce)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *MessageResponseReferencedMessage) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetWebhookId

`func (o *MessageResponseReferencedMessage) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *MessageResponseReferencedMessage) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *MessageResponseReferencedMessage) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *MessageResponseReferencedMessage) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetMessageReference

`func (o *MessageResponseReferencedMessage) GetMessageReference() MessageReferenceResponse`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *MessageResponseReferencedMessage) GetMessageReferenceOk() (*MessageReferenceResponse, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *MessageResponseReferencedMessage) SetMessageReference(v MessageReferenceResponse)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *MessageResponseReferencedMessage) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

### GetThread

`func (o *MessageResponseReferencedMessage) GetThread() ThreadResponse`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *MessageResponseReferencedMessage) GetThreadOk() (*ThreadResponse, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *MessageResponseReferencedMessage) SetThread(v ThreadResponse)`

SetThread sets Thread field to given value.

### HasThread

`func (o *MessageResponseReferencedMessage) HasThread() bool`

HasThread returns a boolean if a field has been set.

### GetMentionChannels

`func (o *MessageResponseReferencedMessage) GetMentionChannels() []BasicMessageResponseMentionChannelsInner`

GetMentionChannels returns the MentionChannels field if non-nil, zero value otherwise.

### GetMentionChannelsOk

`func (o *MessageResponseReferencedMessage) GetMentionChannelsOk() (*[]BasicMessageResponseMentionChannelsInner, bool)`

GetMentionChannelsOk returns a tuple with the MentionChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionChannels

`func (o *MessageResponseReferencedMessage) SetMentionChannels(v []BasicMessageResponseMentionChannelsInner)`

SetMentionChannels sets MentionChannels field to given value.

### HasMentionChannels

`func (o *MessageResponseReferencedMessage) HasMentionChannels() bool`

HasMentionChannels returns a boolean if a field has been set.

### GetRoleSubscriptionData

`func (o *MessageResponseReferencedMessage) GetRoleSubscriptionData() MessageRoleSubscriptionDataResponse`

GetRoleSubscriptionData returns the RoleSubscriptionData field if non-nil, zero value otherwise.

### GetRoleSubscriptionDataOk

`func (o *MessageResponseReferencedMessage) GetRoleSubscriptionDataOk() (*MessageRoleSubscriptionDataResponse, bool)`

GetRoleSubscriptionDataOk returns a tuple with the RoleSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSubscriptionData

`func (o *MessageResponseReferencedMessage) SetRoleSubscriptionData(v MessageRoleSubscriptionDataResponse)`

SetRoleSubscriptionData sets RoleSubscriptionData field to given value.

### HasRoleSubscriptionData

`func (o *MessageResponseReferencedMessage) HasRoleSubscriptionData() bool`

HasRoleSubscriptionData returns a boolean if a field has been set.

### GetPurchaseNotification

`func (o *MessageResponseReferencedMessage) GetPurchaseNotification() PurchaseNotificationResponse`

GetPurchaseNotification returns the PurchaseNotification field if non-nil, zero value otherwise.

### GetPurchaseNotificationOk

`func (o *MessageResponseReferencedMessage) GetPurchaseNotificationOk() (*PurchaseNotificationResponse, bool)`

GetPurchaseNotificationOk returns a tuple with the PurchaseNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseNotification

`func (o *MessageResponseReferencedMessage) SetPurchaseNotification(v PurchaseNotificationResponse)`

SetPurchaseNotification sets PurchaseNotification field to given value.

### HasPurchaseNotification

`func (o *MessageResponseReferencedMessage) HasPurchaseNotification() bool`

HasPurchaseNotification returns a boolean if a field has been set.

### GetPosition

`func (o *MessageResponseReferencedMessage) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MessageResponseReferencedMessage) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MessageResponseReferencedMessage) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MessageResponseReferencedMessage) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetResolved

`func (o *MessageResponseReferencedMessage) GetResolved() ResolvedObjectsResponse`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MessageResponseReferencedMessage) GetResolvedOk() (*ResolvedObjectsResponse, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MessageResponseReferencedMessage) SetResolved(v ResolvedObjectsResponse)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *MessageResponseReferencedMessage) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetPoll

`func (o *MessageResponseReferencedMessage) GetPoll() PollResponse`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *MessageResponseReferencedMessage) GetPollOk() (*PollResponse, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *MessageResponseReferencedMessage) SetPoll(v PollResponse)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *MessageResponseReferencedMessage) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetSharedClientTheme

`func (o *MessageResponseReferencedMessage) GetSharedClientTheme() BasicMessageResponseSharedClientTheme`

GetSharedClientTheme returns the SharedClientTheme field if non-nil, zero value otherwise.

### GetSharedClientThemeOk

`func (o *MessageResponseReferencedMessage) GetSharedClientThemeOk() (*BasicMessageResponseSharedClientTheme, bool)`

GetSharedClientThemeOk returns a tuple with the SharedClientTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedClientTheme

`func (o *MessageResponseReferencedMessage) SetSharedClientTheme(v BasicMessageResponseSharedClientTheme)`

SetSharedClientTheme sets SharedClientTheme field to given value.

### HasSharedClientTheme

`func (o *MessageResponseReferencedMessage) HasSharedClientTheme() bool`

HasSharedClientTheme returns a boolean if a field has been set.

### GetInteractionMetadata

`func (o *MessageResponseReferencedMessage) GetInteractionMetadata() BasicMessageResponseInteractionMetadata`

GetInteractionMetadata returns the InteractionMetadata field if non-nil, zero value otherwise.

### GetInteractionMetadataOk

`func (o *MessageResponseReferencedMessage) GetInteractionMetadataOk() (*BasicMessageResponseInteractionMetadata, bool)`

GetInteractionMetadataOk returns a tuple with the InteractionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionMetadata

`func (o *MessageResponseReferencedMessage) SetInteractionMetadata(v BasicMessageResponseInteractionMetadata)`

SetInteractionMetadata sets InteractionMetadata field to given value.

### HasInteractionMetadata

`func (o *MessageResponseReferencedMessage) HasInteractionMetadata() bool`

HasInteractionMetadata returns a boolean if a field has been set.

### GetMessageSnapshots

`func (o *MessageResponseReferencedMessage) GetMessageSnapshots() []MessageSnapshotResponse`

GetMessageSnapshots returns the MessageSnapshots field if non-nil, zero value otherwise.

### GetMessageSnapshotsOk

`func (o *MessageResponseReferencedMessage) GetMessageSnapshotsOk() (*[]MessageSnapshotResponse, bool)`

GetMessageSnapshotsOk returns a tuple with the MessageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSnapshots

`func (o *MessageResponseReferencedMessage) SetMessageSnapshots(v []MessageSnapshotResponse)`

SetMessageSnapshots sets MessageSnapshots field to given value.

### HasMessageSnapshots

`func (o *MessageResponseReferencedMessage) HasMessageSnapshots() bool`

HasMessageSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


