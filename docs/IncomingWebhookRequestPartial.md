# IncomingWebhookRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**NullableMessageAllowedMentionsRequest**](MessageAllowedMentionsRequest.md) |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Poll** | Pointer to [**NullablePollCreateRequest**](PollCreateRequest.md) |  | [optional] 
**Tts** | Pointer to **NullableBool** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**ThreadName** | Pointer to **NullableString** |  | [optional] 
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

### SetContentNil

`func (o *IncomingWebhookRequestPartial) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *IncomingWebhookRequestPartial) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
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

### SetEmbedsNil

`func (o *IncomingWebhookRequestPartial) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *IncomingWebhookRequestPartial) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetAllowedMentions

`func (o *IncomingWebhookRequestPartial) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *IncomingWebhookRequestPartial) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *IncomingWebhookRequestPartial) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *IncomingWebhookRequestPartial) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### SetAllowedMentionsNil

`func (o *IncomingWebhookRequestPartial) SetAllowedMentionsNil(b bool)`

 SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil

### UnsetAllowedMentions
`func (o *IncomingWebhookRequestPartial) UnsetAllowedMentions()`

UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
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

### SetComponentsNil

`func (o *IncomingWebhookRequestPartial) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *IncomingWebhookRequestPartial) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
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

### SetAttachmentsNil

`func (o *IncomingWebhookRequestPartial) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *IncomingWebhookRequestPartial) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetPoll

`func (o *IncomingWebhookRequestPartial) GetPoll() PollCreateRequest`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *IncomingWebhookRequestPartial) GetPollOk() (*PollCreateRequest, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *IncomingWebhookRequestPartial) SetPoll(v PollCreateRequest)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *IncomingWebhookRequestPartial) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### SetPollNil

`func (o *IncomingWebhookRequestPartial) SetPollNil(b bool)`

 SetPollNil sets the value for Poll to be an explicit nil

### UnsetPoll
`func (o *IncomingWebhookRequestPartial) UnsetPoll()`

UnsetPoll ensures that no value is present for Poll, not even an explicit nil
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

### SetTtsNil

`func (o *IncomingWebhookRequestPartial) SetTtsNil(b bool)`

 SetTtsNil sets the value for Tts to be an explicit nil

### UnsetTts
`func (o *IncomingWebhookRequestPartial) UnsetTts()`

UnsetTts ensures that no value is present for Tts, not even an explicit nil
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

### SetFlagsNil

`func (o *IncomingWebhookRequestPartial) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *IncomingWebhookRequestPartial) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
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

### SetUsernameNil

`func (o *IncomingWebhookRequestPartial) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *IncomingWebhookRequestPartial) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
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

### SetAvatarUrlNil

`func (o *IncomingWebhookRequestPartial) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *IncomingWebhookRequestPartial) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
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

### SetThreadNameNil

`func (o *IncomingWebhookRequestPartial) SetThreadNameNil(b bool)`

 SetThreadNameNil sets the value for ThreadName to be an explicit nil

### UnsetThreadName
`func (o *IncomingWebhookRequestPartial) UnsetThreadName()`

UnsetThreadName ensures that no value is present for ThreadName, not even an explicit nil
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

### SetAppliedTagsNil

`func (o *IncomingWebhookRequestPartial) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *IncomingWebhookRequestPartial) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


