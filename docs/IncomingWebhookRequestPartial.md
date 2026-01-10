# IncomingWebhookRequestPartial

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
**Username** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**ThreadName** | Pointer to **string** |  | [optional] 
**AppliedTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIncomingWebhookRequestPartial

`func NewIncomingWebhookRequestPartial() *IncomingWebhookRequestPartial`

NewIncomingWebhookRequestPartial instantiates a new IncomingWebhookRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingWebhookRequestPartialWithDefaults

`func NewIncomingWebhookRequestPartialWithDefaults() *IncomingWebhookRequestPartial`

NewIncomingWebhookRequestPartialWithDefaults instantiates a new IncomingWebhookRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *IncomingWebhookRequestPartial) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncomingWebhookRequestPartial) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncomingWebhookRequestPartial) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IncomingWebhookRequestPartial) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *IncomingWebhookRequestPartial) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *IncomingWebhookRequestPartial) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *IncomingWebhookRequestPartial) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *IncomingWebhookRequestPartial) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAllowedMentions

`func (o *IncomingWebhookRequestPartial) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *IncomingWebhookRequestPartial) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *IncomingWebhookRequestPartial) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *IncomingWebhookRequestPartial) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetComponents

`func (o *IncomingWebhookRequestPartial) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IncomingWebhookRequestPartial) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IncomingWebhookRequestPartial) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *IncomingWebhookRequestPartial) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetAttachments

`func (o *IncomingWebhookRequestPartial) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IncomingWebhookRequestPartial) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IncomingWebhookRequestPartial) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IncomingWebhookRequestPartial) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPoll

`func (o *IncomingWebhookRequestPartial) GetPoll() BaseCreateMessageCreateRequestPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *IncomingWebhookRequestPartial) GetPollOk() (*BaseCreateMessageCreateRequestPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *IncomingWebhookRequestPartial) SetPoll(v BaseCreateMessageCreateRequestPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *IncomingWebhookRequestPartial) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetTts

`func (o *IncomingWebhookRequestPartial) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *IncomingWebhookRequestPartial) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *IncomingWebhookRequestPartial) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *IncomingWebhookRequestPartial) HasTts() bool`

HasTts returns a boolean if a field has been set.

### GetFlags

`func (o *IncomingWebhookRequestPartial) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *IncomingWebhookRequestPartial) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *IncomingWebhookRequestPartial) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *IncomingWebhookRequestPartial) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetUsername

`func (o *IncomingWebhookRequestPartial) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IncomingWebhookRequestPartial) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IncomingWebhookRequestPartial) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IncomingWebhookRequestPartial) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *IncomingWebhookRequestPartial) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *IncomingWebhookRequestPartial) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *IncomingWebhookRequestPartial) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *IncomingWebhookRequestPartial) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetThreadName

`func (o *IncomingWebhookRequestPartial) GetThreadName() string`

GetThreadName returns the ThreadName field if non-nil, zero value otherwise.

### GetThreadNameOk

`func (o *IncomingWebhookRequestPartial) GetThreadNameOk() (*string, bool)`

GetThreadNameOk returns a tuple with the ThreadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadName

`func (o *IncomingWebhookRequestPartial) SetThreadName(v string)`

SetThreadName sets ThreadName field to given value.

### HasThreadName

`func (o *IncomingWebhookRequestPartial) HasThreadName() bool`

HasThreadName returns a boolean if a field has been set.

### GetAppliedTags

`func (o *IncomingWebhookRequestPartial) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *IncomingWebhookRequestPartial) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *IncomingWebhookRequestPartial) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *IncomingWebhookRequestPartial) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


