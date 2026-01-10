# IncomingWebhookUpdateRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**BaseCreateMessageCreateRequestAllowedMentions**](BaseCreateMessageCreateRequestAllowedMentions.md) |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Poll** | Pointer to [**BaseCreateMessageCreateRequestPoll**](BaseCreateMessageCreateRequestPoll.md) |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 

## Methods

### NewIncomingWebhookUpdateRequestPartial

`func NewIncomingWebhookUpdateRequestPartial() *IncomingWebhookUpdateRequestPartial`

NewIncomingWebhookUpdateRequestPartial instantiates a new IncomingWebhookUpdateRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingWebhookUpdateRequestPartialWithDefaults

`func NewIncomingWebhookUpdateRequestPartialWithDefaults() *IncomingWebhookUpdateRequestPartial`

NewIncomingWebhookUpdateRequestPartialWithDefaults instantiates a new IncomingWebhookUpdateRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *IncomingWebhookUpdateRequestPartial) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncomingWebhookUpdateRequestPartial) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncomingWebhookUpdateRequestPartial) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IncomingWebhookUpdateRequestPartial) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *IncomingWebhookUpdateRequestPartial) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *IncomingWebhookUpdateRequestPartial) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *IncomingWebhookUpdateRequestPartial) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *IncomingWebhookUpdateRequestPartial) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAllowedMentions

`func (o *IncomingWebhookUpdateRequestPartial) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *IncomingWebhookUpdateRequestPartial) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *IncomingWebhookUpdateRequestPartial) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *IncomingWebhookUpdateRequestPartial) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetComponents

`func (o *IncomingWebhookUpdateRequestPartial) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IncomingWebhookUpdateRequestPartial) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IncomingWebhookUpdateRequestPartial) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *IncomingWebhookUpdateRequestPartial) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetAttachments

`func (o *IncomingWebhookUpdateRequestPartial) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IncomingWebhookUpdateRequestPartial) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IncomingWebhookUpdateRequestPartial) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IncomingWebhookUpdateRequestPartial) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPoll

`func (o *IncomingWebhookUpdateRequestPartial) GetPoll() BaseCreateMessageCreateRequestPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *IncomingWebhookUpdateRequestPartial) GetPollOk() (*BaseCreateMessageCreateRequestPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *IncomingWebhookUpdateRequestPartial) SetPoll(v BaseCreateMessageCreateRequestPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *IncomingWebhookUpdateRequestPartial) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetFlags

`func (o *IncomingWebhookUpdateRequestPartial) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *IncomingWebhookUpdateRequestPartial) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *IncomingWebhookUpdateRequestPartial) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *IncomingWebhookUpdateRequestPartial) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


