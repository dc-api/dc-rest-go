# SDKMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**BaseCreateMessageCreateRequestAllowedMentions**](BaseCreateMessageCreateRequestAllowedMentions.md) |  | [optional] 
**StickerIds** | Pointer to **[]string** |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Poll** | Pointer to [**BaseCreateMessageCreateRequestPoll**](BaseCreateMessageCreateRequestPoll.md) |  | [optional] 
**SharedClientTheme** | Pointer to [**BaseCreateMessageCreateRequestSharedClientTheme**](BaseCreateMessageCreateRequestSharedClientTheme.md) |  | [optional] 
**MessageReference** | Pointer to [**MessageCreateRequestMessageReference**](MessageCreateRequestMessageReference.md) |  | [optional] 
**Nonce** | Pointer to [**BasicMessageResponseNonce**](BasicMessageResponseNonce.md) |  | [optional] 
**EnforceNonce** | Pointer to **bool** |  | [optional] 
**Tts** | Pointer to **bool** |  | [optional] 

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

### GetAllowedMentions

`func (o *SDKMessageRequest) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *SDKMessageRequest) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *SDKMessageRequest) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *SDKMessageRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

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

### GetPoll

`func (o *SDKMessageRequest) GetPoll() BaseCreateMessageCreateRequestPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *SDKMessageRequest) GetPollOk() (*BaseCreateMessageCreateRequestPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *SDKMessageRequest) SetPoll(v BaseCreateMessageCreateRequestPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *SDKMessageRequest) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetSharedClientTheme

`func (o *SDKMessageRequest) GetSharedClientTheme() BaseCreateMessageCreateRequestSharedClientTheme`

GetSharedClientTheme returns the SharedClientTheme field if non-nil, zero value otherwise.

### GetSharedClientThemeOk

`func (o *SDKMessageRequest) GetSharedClientThemeOk() (*BaseCreateMessageCreateRequestSharedClientTheme, bool)`

GetSharedClientThemeOk returns a tuple with the SharedClientTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedClientTheme

`func (o *SDKMessageRequest) SetSharedClientTheme(v BaseCreateMessageCreateRequestSharedClientTheme)`

SetSharedClientTheme sets SharedClientTheme field to given value.

### HasSharedClientTheme

`func (o *SDKMessageRequest) HasSharedClientTheme() bool`

HasSharedClientTheme returns a boolean if a field has been set.

### GetMessageReference

`func (o *SDKMessageRequest) GetMessageReference() MessageCreateRequestMessageReference`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *SDKMessageRequest) GetMessageReferenceOk() (*MessageCreateRequestMessageReference, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *SDKMessageRequest) SetMessageReference(v MessageCreateRequestMessageReference)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *SDKMessageRequest) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


