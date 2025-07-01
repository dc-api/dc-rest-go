# GuildPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**HomeHeader** | Pointer to **NullableString** |  | [optional] 
**Splash** | Pointer to **NullableString** |  | [optional] 
**DiscoverySplash** | Pointer to **NullableString** |  | [optional] 
**Features** | **[]string** |  | 
**ApproximateMemberCount** | **int32** |  | 
**ApproximatePresenceCount** | **int32** |  | 
**Emojis** | [**[]EmojiResponse**](EmojiResponse.md) |  | 
**Stickers** | [**[]GuildStickerResponse**](GuildStickerResponse.md) |  | 

## Methods

### NewGuildPreviewResponse

`func NewGuildPreviewResponse(id string, name string, features []string, approximateMemberCount int32, approximatePresenceCount int32, emojis []EmojiResponse, stickers []GuildStickerResponse, ) *GuildPreviewResponse`

NewGuildPreviewResponse instantiates a new GuildPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildPreviewResponseWithDefaults

`func NewGuildPreviewResponseWithDefaults() *GuildPreviewResponse`

NewGuildPreviewResponseWithDefaults instantiates a new GuildPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildPreviewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildPreviewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildPreviewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GuildPreviewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildPreviewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildPreviewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *GuildPreviewResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildPreviewResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildPreviewResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildPreviewResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *GuildPreviewResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *GuildPreviewResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *GuildPreviewResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildPreviewResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildPreviewResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildPreviewResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildPreviewResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildPreviewResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHomeHeader

`func (o *GuildPreviewResponse) GetHomeHeader() string`

GetHomeHeader returns the HomeHeader field if non-nil, zero value otherwise.

### GetHomeHeaderOk

`func (o *GuildPreviewResponse) GetHomeHeaderOk() (*string, bool)`

GetHomeHeaderOk returns a tuple with the HomeHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeHeader

`func (o *GuildPreviewResponse) SetHomeHeader(v string)`

SetHomeHeader sets HomeHeader field to given value.

### HasHomeHeader

`func (o *GuildPreviewResponse) HasHomeHeader() bool`

HasHomeHeader returns a boolean if a field has been set.

### SetHomeHeaderNil

`func (o *GuildPreviewResponse) SetHomeHeaderNil(b bool)`

 SetHomeHeaderNil sets the value for HomeHeader to be an explicit nil

### UnsetHomeHeader
`func (o *GuildPreviewResponse) UnsetHomeHeader()`

UnsetHomeHeader ensures that no value is present for HomeHeader, not even an explicit nil
### GetSplash

`func (o *GuildPreviewResponse) GetSplash() string`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *GuildPreviewResponse) GetSplashOk() (*string, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *GuildPreviewResponse) SetSplash(v string)`

SetSplash sets Splash field to given value.

### HasSplash

`func (o *GuildPreviewResponse) HasSplash() bool`

HasSplash returns a boolean if a field has been set.

### SetSplashNil

`func (o *GuildPreviewResponse) SetSplashNil(b bool)`

 SetSplashNil sets the value for Splash to be an explicit nil

### UnsetSplash
`func (o *GuildPreviewResponse) UnsetSplash()`

UnsetSplash ensures that no value is present for Splash, not even an explicit nil
### GetDiscoverySplash

`func (o *GuildPreviewResponse) GetDiscoverySplash() string`

GetDiscoverySplash returns the DiscoverySplash field if non-nil, zero value otherwise.

### GetDiscoverySplashOk

`func (o *GuildPreviewResponse) GetDiscoverySplashOk() (*string, bool)`

GetDiscoverySplashOk returns a tuple with the DiscoverySplash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySplash

`func (o *GuildPreviewResponse) SetDiscoverySplash(v string)`

SetDiscoverySplash sets DiscoverySplash field to given value.

### HasDiscoverySplash

`func (o *GuildPreviewResponse) HasDiscoverySplash() bool`

HasDiscoverySplash returns a boolean if a field has been set.

### SetDiscoverySplashNil

`func (o *GuildPreviewResponse) SetDiscoverySplashNil(b bool)`

 SetDiscoverySplashNil sets the value for DiscoverySplash to be an explicit nil

### UnsetDiscoverySplash
`func (o *GuildPreviewResponse) UnsetDiscoverySplash()`

UnsetDiscoverySplash ensures that no value is present for DiscoverySplash, not even an explicit nil
### GetFeatures

`func (o *GuildPreviewResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GuildPreviewResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GuildPreviewResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetApproximateMemberCount

`func (o *GuildPreviewResponse) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *GuildPreviewResponse) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *GuildPreviewResponse) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.


### GetApproximatePresenceCount

`func (o *GuildPreviewResponse) GetApproximatePresenceCount() int32`

GetApproximatePresenceCount returns the ApproximatePresenceCount field if non-nil, zero value otherwise.

### GetApproximatePresenceCountOk

`func (o *GuildPreviewResponse) GetApproximatePresenceCountOk() (*int32, bool)`

GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatePresenceCount

`func (o *GuildPreviewResponse) SetApproximatePresenceCount(v int32)`

SetApproximatePresenceCount sets ApproximatePresenceCount field to given value.


### GetEmojis

`func (o *GuildPreviewResponse) GetEmojis() []EmojiResponse`

GetEmojis returns the Emojis field if non-nil, zero value otherwise.

### GetEmojisOk

`func (o *GuildPreviewResponse) GetEmojisOk() (*[]EmojiResponse, bool)`

GetEmojisOk returns a tuple with the Emojis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojis

`func (o *GuildPreviewResponse) SetEmojis(v []EmojiResponse)`

SetEmojis sets Emojis field to given value.


### GetStickers

`func (o *GuildPreviewResponse) GetStickers() []GuildStickerResponse`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *GuildPreviewResponse) GetStickersOk() (*[]GuildStickerResponse, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *GuildPreviewResponse) SetStickers(v []GuildStickerResponse)`

SetStickers sets Stickers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


