# IncomingWebhookUpdateRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**NullableMessageAllowedMentionsRequest**](MessageAllowedMentionsRequest.md) |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Poll** | Pointer to [**NullablePollCreateRequest**](PollCreateRequest.md) |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 

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

### SetContentNil

`func (o *IncomingWebhookUpdateRequestPartial) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *IncomingWebhookUpdateRequestPartial) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
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

### SetEmbedsNil

`func (o *IncomingWebhookUpdateRequestPartial) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *IncomingWebhookUpdateRequestPartial) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetAllowedMentions

`func (o *IncomingWebhookUpdateRequestPartial) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *IncomingWebhookUpdateRequestPartial) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *IncomingWebhookUpdateRequestPartial) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *IncomingWebhookUpdateRequestPartial) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### SetAllowedMentionsNil

`func (o *IncomingWebhookUpdateRequestPartial) SetAllowedMentionsNil(b bool)`

 SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil

### UnsetAllowedMentions
`func (o *IncomingWebhookUpdateRequestPartial) UnsetAllowedMentions()`

UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
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

### SetComponentsNil

`func (o *IncomingWebhookUpdateRequestPartial) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *IncomingWebhookUpdateRequestPartial) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
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

### SetAttachmentsNil

`func (o *IncomingWebhookUpdateRequestPartial) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *IncomingWebhookUpdateRequestPartial) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetPoll

`func (o *IncomingWebhookUpdateRequestPartial) GetPoll() PollCreateRequest`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *IncomingWebhookUpdateRequestPartial) GetPollOk() (*PollCreateRequest, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *IncomingWebhookUpdateRequestPartial) SetPoll(v PollCreateRequest)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *IncomingWebhookUpdateRequestPartial) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *IncomingWebhookUpdateRequestPartial) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *IncomingWebhookUpdateRequestPartial) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
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

### SetFlagsNil

`func (o *IncomingWebhookUpdateRequestPartial) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *IncomingWebhookUpdateRequestPartial) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


