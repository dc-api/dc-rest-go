# BaseCreateMessageCreateRequest

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

## Methods

### NewBaseCreateMessageCreateRequest

`func NewBaseCreateMessageCreateRequest() *BaseCreateMessageCreateRequest`

NewBaseCreateMessageCreateRequest instantiates a new BaseCreateMessageCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCreateMessageCreateRequestWithDefaults

`func NewBaseCreateMessageCreateRequestWithDefaults() *BaseCreateMessageCreateRequest`

NewBaseCreateMessageCreateRequestWithDefaults instantiates a new BaseCreateMessageCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *BaseCreateMessageCreateRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BaseCreateMessageCreateRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BaseCreateMessageCreateRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BaseCreateMessageCreateRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *BaseCreateMessageCreateRequest) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *BaseCreateMessageCreateRequest) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *BaseCreateMessageCreateRequest) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *BaseCreateMessageCreateRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAllowedMentions

`func (o *BaseCreateMessageCreateRequest) GetAllowedMentions() BaseCreateMessageCreateRequestAllowedMentions`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *BaseCreateMessageCreateRequest) GetAllowedMentionsOk() (*BaseCreateMessageCreateRequestAllowedMentions, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *BaseCreateMessageCreateRequest) SetAllowedMentions(v BaseCreateMessageCreateRequestAllowedMentions)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *BaseCreateMessageCreateRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetStickerIds

`func (o *BaseCreateMessageCreateRequest) GetStickerIds() []string`

GetStickerIds returns the StickerIds field if non-nil, zero value otherwise.

### GetStickerIdsOk

`func (o *BaseCreateMessageCreateRequest) GetStickerIdsOk() (*[]string, bool)`

GetStickerIdsOk returns a tuple with the StickerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerIds

`func (o *BaseCreateMessageCreateRequest) SetStickerIds(v []string)`

SetStickerIds sets StickerIds field to given value.

### HasStickerIds

`func (o *BaseCreateMessageCreateRequest) HasStickerIds() bool`

HasStickerIds returns a boolean if a field has been set.

### GetComponents

`func (o *BaseCreateMessageCreateRequest) GetComponents() []BaseCreateMessageCreateRequestComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BaseCreateMessageCreateRequest) GetComponentsOk() (*[]BaseCreateMessageCreateRequestComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BaseCreateMessageCreateRequest) SetComponents(v []BaseCreateMessageCreateRequestComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BaseCreateMessageCreateRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetFlags

`func (o *BaseCreateMessageCreateRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *BaseCreateMessageCreateRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *BaseCreateMessageCreateRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *BaseCreateMessageCreateRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetAttachments

`func (o *BaseCreateMessageCreateRequest) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BaseCreateMessageCreateRequest) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BaseCreateMessageCreateRequest) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BaseCreateMessageCreateRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPoll

`func (o *BaseCreateMessageCreateRequest) GetPoll() BaseCreateMessageCreateRequestPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *BaseCreateMessageCreateRequest) GetPollOk() (*BaseCreateMessageCreateRequestPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *BaseCreateMessageCreateRequest) SetPoll(v BaseCreateMessageCreateRequestPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *BaseCreateMessageCreateRequest) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetSharedClientTheme

`func (o *BaseCreateMessageCreateRequest) GetSharedClientTheme() BaseCreateMessageCreateRequestSharedClientTheme`

GetSharedClientTheme returns the SharedClientTheme field if non-nil, zero value otherwise.

### GetSharedClientThemeOk

`func (o *BaseCreateMessageCreateRequest) GetSharedClientThemeOk() (*BaseCreateMessageCreateRequestSharedClientTheme, bool)`

GetSharedClientThemeOk returns a tuple with the SharedClientTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedClientTheme

`func (o *BaseCreateMessageCreateRequest) SetSharedClientTheme(v BaseCreateMessageCreateRequestSharedClientTheme)`

SetSharedClientTheme sets SharedClientTheme field to given value.

### HasSharedClientTheme

`func (o *BaseCreateMessageCreateRequest) HasSharedClientTheme() bool`

HasSharedClientTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


