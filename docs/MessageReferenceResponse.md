# MessageReferenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** |  | [optional] 
**ChannelId** | **string** |  | 
**MessageId** | Pointer to **string** |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageReferenceResponse

`func NewMessageReferenceResponse(channelId string, ) *MessageReferenceResponse`

NewMessageReferenceResponse instantiates a new MessageReferenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageReferenceResponseWithDefaults

`func NewMessageReferenceResponseWithDefaults() *MessageReferenceResponse`

NewMessageReferenceResponseWithDefaults instantiates a new MessageReferenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageReferenceResponse) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageReferenceResponse) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageReferenceResponse) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *MessageReferenceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MessageReferenceResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MessageReferenceResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetChannelId

`func (o *MessageReferenceResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MessageReferenceResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MessageReferenceResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetMessageId

`func (o *MessageReferenceResponse) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageReferenceResponse) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageReferenceResponse) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageReferenceResponse) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetGuildId

`func (o *MessageReferenceResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *MessageReferenceResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *MessageReferenceResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *MessageReferenceResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


