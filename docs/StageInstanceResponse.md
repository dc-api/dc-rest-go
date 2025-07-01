# StageInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildId** | **string** |  | 
**ChannelId** | **string** |  | 
**Topic** | **string** |  | 
**PrivacyLevel** | **int32** |  | 
**Id** | **string** |  | 
**DiscoverableDisabled** | Pointer to **NullableBool** |  | [optional] 
**GuildScheduledEventId** | Pointer to **string** |  | [optional] 

## Methods

### NewStageInstanceResponse

`func NewStageInstanceResponse(guildId string, channelId string, topic string, privacyLevel int32, id string, ) *StageInstanceResponse`

NewStageInstanceResponse instantiates a new StageInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageInstanceResponseWithDefaults

`func NewStageInstanceResponseWithDefaults() *StageInstanceResponse`

NewStageInstanceResponseWithDefaults instantiates a new StageInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildId

`func (o *StageInstanceResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *StageInstanceResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *StageInstanceResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetChannelId

`func (o *StageInstanceResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *StageInstanceResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *StageInstanceResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetTopic

`func (o *StageInstanceResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *StageInstanceResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *StageInstanceResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetPrivacyLevel

`func (o *StageInstanceResponse) GetPrivacyLevel() int32`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *StageInstanceResponse) GetPrivacyLevelOk() (*int32, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *StageInstanceResponse) SetPrivacyLevel(v int32)`

SetPrivacyLevel sets PrivacyLevel field to given value.


### GetId

`func (o *StageInstanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StageInstanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StageInstanceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDiscoverableDisabled

`func (o *StageInstanceResponse) GetDiscoverableDisabled() bool`

GetDiscoverableDisabled returns the DiscoverableDisabled field if non-nil, zero value otherwise.

### GetDiscoverableDisabledOk

`func (o *StageInstanceResponse) GetDiscoverableDisabledOk() (*bool, bool)`

GetDiscoverableDisabledOk returns a tuple with the DiscoverableDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverableDisabled

`func (o *StageInstanceResponse) SetDiscoverableDisabled(v bool)`

SetDiscoverableDisabled sets DiscoverableDisabled field to given value.

### HasDiscoverableDisabled

`func (o *StageInstanceResponse) HasDiscoverableDisabled() bool`

HasDiscoverableDisabled returns a boolean if a field has been set.

### SetDiscoverableDisabledNil

`func (o *StageInstanceResponse) SetDiscoverableDisabledNil(b bool)`

 SetDiscoverableDisabledNil sets the value for DiscoverableDisabled to be an explicit nil

### UnsetDiscoverableDisabled
`func (o *StageInstanceResponse) UnsetDiscoverableDisabled()`

UnsetDiscoverableDisabled ensures that no value is present for DiscoverableDisabled, not even an explicit nil
### GetGuildScheduledEventId

`func (o *StageInstanceResponse) GetGuildScheduledEventId() string`

GetGuildScheduledEventId returns the GuildScheduledEventId field if non-nil, zero value otherwise.

### GetGuildScheduledEventIdOk

`func (o *StageInstanceResponse) GetGuildScheduledEventIdOk() (*string, bool)`

GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEventId

`func (o *StageInstanceResponse) SetGuildScheduledEventId(v string)`

SetGuildScheduledEventId sets GuildScheduledEventId field to given value.

### HasGuildScheduledEventId

`func (o *StageInstanceResponse) HasGuildScheduledEventId() bool`

HasGuildScheduledEventId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


