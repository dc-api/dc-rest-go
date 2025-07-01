# SoundboardSoundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SoundId** | **string** |  | 
**Volume** | **float64** |  | 
**EmojiId** | Pointer to **string** |  | [optional] 
**EmojiName** | Pointer to **NullableString** |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 
**Available** | **bool** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewSoundboardSoundResponse

`func NewSoundboardSoundResponse(name string, soundId string, volume float64, available bool, ) *SoundboardSoundResponse`

NewSoundboardSoundResponse instantiates a new SoundboardSoundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundboardSoundResponseWithDefaults

`func NewSoundboardSoundResponseWithDefaults() *SoundboardSoundResponse`

NewSoundboardSoundResponseWithDefaults instantiates a new SoundboardSoundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SoundboardSoundResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoundboardSoundResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoundboardSoundResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSoundId

`func (o *SoundboardSoundResponse) GetSoundId() string`

GetSoundId returns the SoundId field if non-nil, zero value otherwise.

### GetSoundIdOk

`func (o *SoundboardSoundResponse) GetSoundIdOk() (*string, bool)`

GetSoundIdOk returns a tuple with the SoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundId

`func (o *SoundboardSoundResponse) SetSoundId(v string)`

SetSoundId sets SoundId field to given value.


### GetVolume

`func (o *SoundboardSoundResponse) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SoundboardSoundResponse) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SoundboardSoundResponse) SetVolume(v float64)`

SetVolume sets Volume field to given value.


### GetEmojiId

`func (o *SoundboardSoundResponse) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *SoundboardSoundResponse) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *SoundboardSoundResponse) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *SoundboardSoundResponse) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *SoundboardSoundResponse) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *SoundboardSoundResponse) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *SoundboardSoundResponse) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *SoundboardSoundResponse) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *SoundboardSoundResponse) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *SoundboardSoundResponse) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
### GetGuildId

`func (o *SoundboardSoundResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *SoundboardSoundResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *SoundboardSoundResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *SoundboardSoundResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetAvailable

`func (o *SoundboardSoundResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *SoundboardSoundResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *SoundboardSoundResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetUser

`func (o *SoundboardSoundResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SoundboardSoundResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SoundboardSoundResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *SoundboardSoundResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SoundboardSoundResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SoundboardSoundResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


