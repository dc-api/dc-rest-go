# MessageReferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuildId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**ChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**MessageId** | **string** |  | 
**FailIfNotExists** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**MessageReferenceRequestType**](MessageReferenceRequestType.md) |  | [optional] 

## Methods

### NewMessageReferenceRequest

`func NewMessageReferenceRequest(messageId string, ) *MessageReferenceRequest`

NewMessageReferenceRequest instantiates a new MessageReferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageReferenceRequestWithDefaults

`func NewMessageReferenceRequestWithDefaults() *MessageReferenceRequest`

NewMessageReferenceRequestWithDefaults instantiates a new MessageReferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuildId

`func (o *MessageReferenceRequest) GetGuildId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *MessageReferenceRequest) GetGuildIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *MessageReferenceRequest) SetGuildId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *MessageReferenceRequest) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetChannelId

`func (o *MessageReferenceRequest) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MessageReferenceRequest) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MessageReferenceRequest) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MessageReferenceRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetMessageId

`func (o *MessageReferenceRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageReferenceRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageReferenceRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetFailIfNotExists

`func (o *MessageReferenceRequest) GetFailIfNotExists() bool`

GetFailIfNotExists returns the FailIfNotExists field if non-nil, zero value otherwise.

### GetFailIfNotExistsOk

`func (o *MessageReferenceRequest) GetFailIfNotExistsOk() (*bool, bool)`

GetFailIfNotExistsOk returns a tuple with the FailIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailIfNotExists

`func (o *MessageReferenceRequest) SetFailIfNotExists(v bool)`

SetFailIfNotExists sets FailIfNotExists field to given value.

### HasFailIfNotExists

`func (o *MessageReferenceRequest) HasFailIfNotExists() bool`

HasFailIfNotExists returns a boolean if a field has been set.

### GetType

`func (o *MessageReferenceRequest) GetType() MessageReferenceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageReferenceRequest) GetTypeOk() (*MessageReferenceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageReferenceRequest) SetType(v MessageReferenceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *MessageReferenceRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


