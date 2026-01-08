# IncomingWebhookInteractionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**BaseCreateMessageCreateRequestAllowedMentions**](BaseCreateMessageCreateRequestAllowedMentions.md) |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Poll** | Pointer to [**BaseCreateMessageCreateRequestPoll**](BaseCreateMessageCreateRequestPoll.md) |  | [optional] 
**Tts** | Pointer to **bool** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 

## Methods

### NewIncomingWebhookInteractionRequest

`func NewIncomingWebhookInteractionRequest() *IncomingWebhookInteractionRequest`

NewIncomingWebhookInteractionRequest instantiates a new IncomingWebhookInteractionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingWebhookInteractionRequestWithDefaults

`func NewIncomingWebhookInteractionRequestWithDefaults() *IncomingWebhookInteractionRequest`

NewIncomingWebhookInteractionRequestWithDefaults instantiates a new IncomingWebhookInteractionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *IncomingWebhookInteractionRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncomingWebhookInteractionRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncomingWebhookInteractionRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IncomingWebhookInteractionRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *IncomingWebhookInteractionRequest) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *IncomingWebhookInteractionRequest) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *IncomingWebhookInteractionRequest) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *IncomingWebhookInteractionRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAllowedMentions

`func (o *IncomingWebhookInteractionRequest) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *IncomingWebhookInteractionRequest) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *IncomingWebhookInteractionRequest) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *IncomingWebhookInteractionRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetComponents

`func (o *IncomingWebhookInteractionRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IncomingWebhookInteractionRequest) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IncomingWebhookInteractionRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *IncomingWebhookInteractionRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetAttachments

`func (o *IncomingWebhookInteractionRequest) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IncomingWebhookInteractionRequest) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IncomingWebhookInteractionRequest) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IncomingWebhookInteractionRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPoll

`func (o *IncomingWebhookInteractionRequest) GetPoll() BaseCreateMessageCreateRequestPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *IncomingWebhookInteractionRequest) GetPollOk() (*BaseCreateMessageCreateRequestPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *IncomingWebhookInteractionRequest) SetPoll(v BaseCreateMessageCreateRequestPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *IncomingWebhookInteractionRequest) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetTts

`func (o *IncomingWebhookInteractionRequest) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *IncomingWebhookInteractionRequest) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *IncomingWebhookInteractionRequest) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *IncomingWebhookInteractionRequest) HasTts() bool`

HasTts returns a boolean if a field has been set.

### GetFlags

`func (o *IncomingWebhookInteractionRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *IncomingWebhookInteractionRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *IncomingWebhookInteractionRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *IncomingWebhookInteractionRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


