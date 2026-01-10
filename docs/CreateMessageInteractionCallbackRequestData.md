# CreateMessageInteractionCallbackRequestData

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

### NewCreateMessageInteractionCallbackRequestData

`func NewCreateMessageInteractionCallbackRequestData() *CreateMessageInteractionCallbackRequestData`

NewCreateMessageInteractionCallbackRequestData instantiates a new CreateMessageInteractionCallbackRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageInteractionCallbackRequestDataWithDefaults

`func NewCreateMessageInteractionCallbackRequestDataWithDefaults() *CreateMessageInteractionCallbackRequestData`

NewCreateMessageInteractionCallbackRequestDataWithDefaults instantiates a new CreateMessageInteractionCallbackRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateMessageInteractionCallbackRequestData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateMessageInteractionCallbackRequestData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateMessageInteractionCallbackRequestData) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateMessageInteractionCallbackRequestData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *CreateMessageInteractionCallbackRequestData) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CreateMessageInteractionCallbackRequestData) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CreateMessageInteractionCallbackRequestData) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *CreateMessageInteractionCallbackRequestData) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAllowedMentions

`func (o *CreateMessageInteractionCallbackRequestData) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *CreateMessageInteractionCallbackRequestData) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *CreateMessageInteractionCallbackRequestData) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *CreateMessageInteractionCallbackRequestData) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetComponents

`func (o *CreateMessageInteractionCallbackRequestData) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CreateMessageInteractionCallbackRequestData) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CreateMessageInteractionCallbackRequestData) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CreateMessageInteractionCallbackRequestData) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetAttachments

`func (o *CreateMessageInteractionCallbackRequestData) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateMessageInteractionCallbackRequestData) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateMessageInteractionCallbackRequestData) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateMessageInteractionCallbackRequestData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPoll

`func (o *CreateMessageInteractionCallbackRequestData) GetPoll() BaseCreateMessageCreateRequestPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *CreateMessageInteractionCallbackRequestData) GetPollOk() (*BaseCreateMessageCreateRequestPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *CreateMessageInteractionCallbackRequestData) SetPoll(v BaseCreateMessageCreateRequestPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *CreateMessageInteractionCallbackRequestData) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetTts

`func (o *CreateMessageInteractionCallbackRequestData) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *CreateMessageInteractionCallbackRequestData) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *CreateMessageInteractionCallbackRequestData) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *CreateMessageInteractionCallbackRequestData) HasTts() bool`

HasTts returns a boolean if a field has been set.

### GetFlags

`func (o *CreateMessageInteractionCallbackRequestData) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *CreateMessageInteractionCallbackRequestData) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *CreateMessageInteractionCallbackRequestData) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *CreateMessageInteractionCallbackRequestData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


