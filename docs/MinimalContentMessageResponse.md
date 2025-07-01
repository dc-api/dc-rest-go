# MinimalContentMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Content** | **string** |  | 
**Mentions** | [**[]UserResponse**](UserResponse.md) |  | 
**MentionRoles** | **[]string** |  | 
**Attachments** | [**[]MessageAttachmentResponse**](MessageAttachmentResponse.md) |  | 
**Embeds** | [**[]MessageEmbedResponse**](MessageEmbedResponse.md) |  | 
**Timestamp** | **time.Time** |  | 
**EditedTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Flags** | **int32** |  | 
**Components** | [**[]BasicMessageResponseComponentsInner**](BasicMessageResponseComponentsInner.md) |  | 
**Resolved** | Pointer to [**NullableResolvedObjectsResponse**](ResolvedObjectsResponse.md) |  | [optional] 
**Stickers** | Pointer to [**[]GetSticker200Response**](GetSticker200Response.md) |  | [optional] 
**StickerItems** | Pointer to [**[]MessageStickerItemResponse**](MessageStickerItemResponse.md) |  | [optional] 

## Methods

### NewMinimalContentMessageResponse

`func NewMinimalContentMessageResponse(type_ int32, content string, mentions []UserResponse, mentionRoles []string, attachments []MessageAttachmentResponse, embeds []MessageEmbedResponse, timestamp time.Time, flags int32, components []BasicMessageResponseComponentsInner, ) *MinimalContentMessageResponse`

NewMinimalContentMessageResponse instantiates a new MinimalContentMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalContentMessageResponseWithDefaults

`func NewMinimalContentMessageResponseWithDefaults() *MinimalContentMessageResponse`

NewMinimalContentMessageResponseWithDefaults instantiates a new MinimalContentMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MinimalContentMessageResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MinimalContentMessageResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MinimalContentMessageResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetContent

`func (o *MinimalContentMessageResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MinimalContentMessageResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MinimalContentMessageResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetMentions

`func (o *MinimalContentMessageResponse) GetMentions() []UserResponse`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *MinimalContentMessageResponse) GetMentionsOk() (*[]UserResponse, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *MinimalContentMessageResponse) SetMentions(v []UserResponse)`

SetMentions sets Mentions field to given value.


### GetMentionRoles

`func (o *MinimalContentMessageResponse) GetMentionRoles() []string`

GetMentionRoles returns the MentionRoles field if non-nil, zero value otherwise.

### GetMentionRolesOk

`func (o *MinimalContentMessageResponse) GetMentionRolesOk() (*[]string, bool)`

GetMentionRolesOk returns a tuple with the MentionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionRoles

`func (o *MinimalContentMessageResponse) SetMentionRoles(v []string)`

SetMentionRoles sets MentionRoles field to given value.


### GetAttachments

`func (o *MinimalContentMessageResponse) GetAttachments() []MessageAttachmentResponse`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MinimalContentMessageResponse) GetAttachmentsOk() (*[]MessageAttachmentResponse, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MinimalContentMessageResponse) SetAttachments(v []MessageAttachmentResponse)`

SetAttachments sets Attachments field to given value.


### GetEmbeds

`func (o *MinimalContentMessageResponse) GetEmbeds() []MessageEmbedResponse`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *MinimalContentMessageResponse) GetEmbedsOk() (*[]MessageEmbedResponse, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *MinimalContentMessageResponse) SetEmbeds(v []MessageEmbedResponse)`

SetEmbeds sets Embeds field to given value.


### GetTimestamp

`func (o *MinimalContentMessageResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MinimalContentMessageResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MinimalContentMessageResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEditedTimestamp

`func (o *MinimalContentMessageResponse) GetEditedTimestamp() time.Time`

GetEditedTimestamp returns the EditedTimestamp field if non-nil, zero value otherwise.

### GetEditedTimestampOk

`func (o *MinimalContentMessageResponse) GetEditedTimestampOk() (*time.Time, bool)`

GetEditedTimestampOk returns a tuple with the EditedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTimestamp

`func (o *MinimalContentMessageResponse) SetEditedTimestamp(v time.Time)`

SetEditedTimestamp sets EditedTimestamp field to given value.

### HasEditedTimestamp

`func (o *MinimalContentMessageResponse) HasEditedTimestamp() bool`

HasEditedTimestamp returns a boolean if a field has been set.

### SetEditedTimestampNil

`func (o *MinimalContentMessageResponse) SetEditedTimestampNil(b bool)`

 SetEditedTimestampNil sets the value for EditedTimestamp to be an explicit nil

### UnsetEditedTimestamp
`func (o *MinimalContentMessageResponse) UnsetEditedTimestamp()`

UnsetEditedTimestamp ensures that no value is present for EditedTimestamp, not even an explicit nil
### GetFlags

`func (o *MinimalContentMessageResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MinimalContentMessageResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MinimalContentMessageResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetComponents

`func (o *MinimalContentMessageResponse) GetComponents() []BasicMessageResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *MinimalContentMessageResponse) GetComponentsOk() (*[]BasicMessageResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *MinimalContentMessageResponse) SetComponents(v []BasicMessageResponseComponentsInner)`

SetComponents sets Components field to given value.


### GetResolved

`func (o *MinimalContentMessageResponse) GetResolved() ResolvedObjectsResponse`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MinimalContentMessageResponse) GetResolvedOk() (*ResolvedObjectsResponse, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MinimalContentMessageResponse) SetResolved(v ResolvedObjectsResponse)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *MinimalContentMessageResponse) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *MinimalContentMessageResponse) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *MinimalContentMessageResponse) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetStickers

`func (o *MinimalContentMessageResponse) GetStickers() []GetSticker200Response`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *MinimalContentMessageResponse) GetStickersOk() (*[]GetSticker200Response, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *MinimalContentMessageResponse) SetStickers(v []GetSticker200Response)`

SetStickers sets Stickers field to given value.

### HasStickers

`func (o *MinimalContentMessageResponse) HasStickers() bool`

HasStickers returns a boolean if a field has been set.

### SetStickersNil

`func (o *MinimalContentMessageResponse) SetStickersNil(b bool)`

 SetStickersNil sets the value for Stickers to be an explicit nil

### UnsetStickers
`func (o *MinimalContentMessageResponse) UnsetStickers()`

UnsetStickers ensures that no value is present for Stickers, not even an explicit nil
### GetStickerItems

`func (o *MinimalContentMessageResponse) GetStickerItems() []MessageStickerItemResponse`

GetStickerItems returns the StickerItems field if non-nil, zero value otherwise.

### GetStickerItemsOk

`func (o *MinimalContentMessageResponse) GetStickerItemsOk() (*[]MessageStickerItemResponse, bool)`

GetStickerItemsOk returns a tuple with the StickerItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerItems

`func (o *MinimalContentMessageResponse) SetStickerItems(v []MessageStickerItemResponse)`

SetStickerItems sets StickerItems field to given value.

### HasStickerItems

`func (o *MinimalContentMessageResponse) HasStickerItems() bool`

HasStickerItems returns a boolean if a field has been set.

### SetStickerItemsNil

`func (o *MinimalContentMessageResponse) SetStickerItemsNil(b bool)`

 SetStickerItemsNil sets the value for StickerItems to be an explicit nil

### UnsetStickerItems
`func (o *MinimalContentMessageResponse) UnsetStickerItems()`

UnsetStickerItems ensures that no value is present for StickerItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


