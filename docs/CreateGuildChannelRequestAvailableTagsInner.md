# CreateGuildChannelRequestAvailableTagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EmojiId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**EmojiName** | Pointer to **string** |  | [optional] 
**Moderated** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateGuildChannelRequestAvailableTagsInner

`func NewCreateGuildChannelRequestAvailableTagsInner(name string, ) *CreateGuildChannelRequestAvailableTagsInner`

NewCreateGuildChannelRequestAvailableTagsInner instantiates a new CreateGuildChannelRequestAvailableTagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildChannelRequestAvailableTagsInnerWithDefaults

`func NewCreateGuildChannelRequestAvailableTagsInnerWithDefaults() *CreateGuildChannelRequestAvailableTagsInner`

NewCreateGuildChannelRequestAvailableTagsInnerWithDefaults instantiates a new CreateGuildChannelRequestAvailableTagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildChannelRequestAvailableTagsInner) SetName(v string)`

SetName sets Name field to given value.


### GetEmojiId

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetEmojiId() GetEntitlementsSkuIdsParameterOneOfInner`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetEmojiIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *CreateGuildChannelRequestAvailableTagsInner) SetEmojiId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *CreateGuildChannelRequestAvailableTagsInner) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *CreateGuildChannelRequestAvailableTagsInner) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *CreateGuildChannelRequestAvailableTagsInner) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetModerated

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetModerated() bool`

GetModerated returns the Moderated field if non-nil, zero value otherwise.

### GetModeratedOk

`func (o *CreateGuildChannelRequestAvailableTagsInner) GetModeratedOk() (*bool, bool)`

GetModeratedOk returns a tuple with the Moderated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerated

`func (o *CreateGuildChannelRequestAvailableTagsInner) SetModerated(v bool)`

SetModerated sets Moderated field to given value.

### HasModerated

`func (o *CreateGuildChannelRequestAvailableTagsInner) HasModerated() bool`

HasModerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


