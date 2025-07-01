# CreateStageInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** |  | 
**ChannelId** | **string** |  | 
**PrivacyLevel** | Pointer to **int32** |  | [optional] 
**GuildScheduledEventId** | Pointer to **string** |  | [optional] 
**SendStartNotification** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateStageInstanceRequest

`func NewCreateStageInstanceRequest(topic string, channelId string, ) *CreateStageInstanceRequest`

NewCreateStageInstanceRequest instantiates a new CreateStageInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStageInstanceRequestWithDefaults

`func NewCreateStageInstanceRequestWithDefaults() *CreateStageInstanceRequest`

NewCreateStageInstanceRequestWithDefaults instantiates a new CreateStageInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *CreateStageInstanceRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *CreateStageInstanceRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *CreateStageInstanceRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetChannelId

`func (o *CreateStageInstanceRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateStageInstanceRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateStageInstanceRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetPrivacyLevel

`func (o *CreateStageInstanceRequest) GetPrivacyLevel() int32`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *CreateStageInstanceRequest) GetPrivacyLevelOk() (*int32, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *CreateStageInstanceRequest) SetPrivacyLevel(v int32)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *CreateStageInstanceRequest) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### GetGuildScheduledEventId

`func (o *CreateStageInstanceRequest) GetGuildScheduledEventId() string`

GetGuildScheduledEventId returns the GuildScheduledEventId field if non-nil, zero value otherwise.

### GetGuildScheduledEventIdOk

`func (o *CreateStageInstanceRequest) GetGuildScheduledEventIdOk() (*string, bool)`

GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEventId

`func (o *CreateStageInstanceRequest) SetGuildScheduledEventId(v string)`

SetGuildScheduledEventId sets GuildScheduledEventId field to given value.

### HasGuildScheduledEventId

`func (o *CreateStageInstanceRequest) HasGuildScheduledEventId() bool`

HasGuildScheduledEventId returns a boolean if a field has been set.

### GetSendStartNotification

`func (o *CreateStageInstanceRequest) GetSendStartNotification() bool`

GetSendStartNotification returns the SendStartNotification field if non-nil, zero value otherwise.

### GetSendStartNotificationOk

`func (o *CreateStageInstanceRequest) GetSendStartNotificationOk() (*bool, bool)`

GetSendStartNotificationOk returns a tuple with the SendStartNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStartNotification

`func (o *CreateStageInstanceRequest) SetSendStartNotification(v bool)`

SetSendStartNotification sets SendStartNotification field to given value.

### HasSendStartNotification

`func (o *CreateStageInstanceRequest) HasSendStartNotification() bool`

HasSendStartNotification returns a boolean if a field has been set.

### SetSendStartNotificationNil

`func (o *CreateStageInstanceRequest) SetSendStartNotificationNil(b bool)`

 SetSendStartNotificationNil sets the value for SendStartNotification to be an explicit nil

### UnsetSendStartNotification
`func (o *CreateStageInstanceRequest) UnsetSendStartNotification()`

UnsetSendStartNotification ensures that no value is present for SendStartNotification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


