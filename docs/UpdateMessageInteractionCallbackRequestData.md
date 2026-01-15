# UpdateMessageInteractionCallbackRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**BaseCreateMessageCreateRequestAllowedMentions**](BaseCreateMessageCreateRequestAllowedMentions.md) |  | [optional] 
**Components** | Pointer to [**[]BaseCreateMessageCreateRequestComponentsInner**](BaseCreateMessageCreateRequestComponentsInner.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateMessageInteractionCallbackRequestData

`func NewUpdateMessageInteractionCallbackRequestData() *UpdateMessageInteractionCallbackRequestData`

NewUpdateMessageInteractionCallbackRequestData instantiates a new UpdateMessageInteractionCallbackRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageInteractionCallbackRequestDataWithDefaults

`func NewUpdateMessageInteractionCallbackRequestDataWithDefaults() *UpdateMessageInteractionCallbackRequestData`

NewUpdateMessageInteractionCallbackRequestDataWithDefaults instantiates a new UpdateMessageInteractionCallbackRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UpdateMessageInteractionCallbackRequestData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateMessageInteractionCallbackRequestData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateMessageInteractionCallbackRequestData) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateMessageInteractionCallbackRequestData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *UpdateMessageInteractionCallbackRequestData) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *UpdateMessageInteractionCallbackRequestData) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *UpdateMessageInteractionCallbackRequestData) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *UpdateMessageInteractionCallbackRequestData) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAllowedMentions

`func (o *UpdateMessageInteractionCallbackRequestData) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *UpdateMessageInteractionCallbackRequestData) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *UpdateMessageInteractionCallbackRequestData) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *UpdateMessageInteractionCallbackRequestData) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetComponents

`func (o *UpdateMessageInteractionCallbackRequestData) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *UpdateMessageInteractionCallbackRequestData) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *UpdateMessageInteractionCallbackRequestData) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *UpdateMessageInteractionCallbackRequestData) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetAttachments

`func (o *UpdateMessageInteractionCallbackRequestData) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UpdateMessageInteractionCallbackRequestData) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UpdateMessageInteractionCallbackRequestData) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *UpdateMessageInteractionCallbackRequestData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetFlags

`func (o *UpdateMessageInteractionCallbackRequestData) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateMessageInteractionCallbackRequestData) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateMessageInteractionCallbackRequestData) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UpdateMessageInteractionCallbackRequestData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


