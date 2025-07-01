# SoundboardSoundSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoundId** | **string** |  | 
**SourceGuildId** | Pointer to **string** |  | [optional] 

## Methods

### NewSoundboardSoundSendRequest

`func NewSoundboardSoundSendRequest(soundId string, ) *SoundboardSoundSendRequest`

NewSoundboardSoundSendRequest instantiates a new SoundboardSoundSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundboardSoundSendRequestWithDefaults

`func NewSoundboardSoundSendRequestWithDefaults() *SoundboardSoundSendRequest`

NewSoundboardSoundSendRequestWithDefaults instantiates a new SoundboardSoundSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoundId

`func (o *SoundboardSoundSendRequest) GetSoundId() string`

GetSoundId returns the SoundId field if non-nil, zero value otherwise.

### GetSoundIdOk

`func (o *SoundboardSoundSendRequest) GetSoundIdOk() (*string, bool)`

GetSoundIdOk returns a tuple with the SoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundId

`func (o *SoundboardSoundSendRequest) SetSoundId(v string)`

SetSoundId sets SoundId field to given value.


### GetSourceGuildId

`func (o *SoundboardSoundSendRequest) GetSourceGuildId() string`

GetSourceGuildId returns the SourceGuildId field if non-nil, zero value otherwise.

### GetSourceGuildIdOk

`func (o *SoundboardSoundSendRequest) GetSourceGuildIdOk() (*string, bool)`

GetSourceGuildIdOk returns a tuple with the SourceGuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuildId

`func (o *SoundboardSoundSendRequest) SetSourceGuildId(v string)`

SetSourceGuildId sets SourceGuildId field to given value.

### HasSourceGuildId

`func (o *SoundboardSoundSendRequest) HasSourceGuildId() bool`

HasSourceGuildId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


