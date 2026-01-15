# GuildTemplateChannelTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**EmojiId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**EmojiName** | Pointer to **string** |  | [optional] 
**Moderated** | Pointer to **bool** |  | [optional] 

## Methods

### NewGuildTemplateChannelTags

`func NewGuildTemplateChannelTags(name string, ) *GuildTemplateChannelTags`

NewGuildTemplateChannelTags instantiates a new GuildTemplateChannelTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateChannelTagsWithDefaults

`func NewGuildTemplateChannelTagsWithDefaults() *GuildTemplateChannelTags`

NewGuildTemplateChannelTagsWithDefaults instantiates a new GuildTemplateChannelTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildTemplateChannelTags) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildTemplateChannelTags) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildTemplateChannelTags) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GuildTemplateChannelTags) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GuildTemplateChannelTags) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildTemplateChannelTags) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildTemplateChannelTags) SetName(v string)`

SetName sets Name field to given value.


### GetEmojiId

`func (o *GuildTemplateChannelTags) GetEmojiId() GetEntitlementsSkuIdsParameterOneOfInner`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *GuildTemplateChannelTags) GetEmojiIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *GuildTemplateChannelTags) SetEmojiId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *GuildTemplateChannelTags) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *GuildTemplateChannelTags) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *GuildTemplateChannelTags) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *GuildTemplateChannelTags) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *GuildTemplateChannelTags) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetModerated

`func (o *GuildTemplateChannelTags) GetModerated() bool`

GetModerated returns the Moderated field if non-nil, zero value otherwise.

### GetModeratedOk

`func (o *GuildTemplateChannelTags) GetModeratedOk() (*bool, bool)`

GetModeratedOk returns a tuple with the Moderated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerated

`func (o *GuildTemplateChannelTags) SetModerated(v bool)`

SetModerated sets Moderated field to given value.

### HasModerated

`func (o *GuildTemplateChannelTags) HasModerated() bool`

HasModerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


