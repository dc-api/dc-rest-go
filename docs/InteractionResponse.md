# InteractionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**ResponseMessageId** | Pointer to **string** |  | [optional] 
**ResponseMessageLoading** | Pointer to **NullableBool** |  | [optional] 
**ResponseMessageEphemeral** | Pointer to **NullableBool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 

## Methods

### NewInteractionResponse

`func NewInteractionResponse(id string, type_ int32, ) *InteractionResponse`

NewInteractionResponse instantiates a new InteractionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionResponseWithDefaults

`func NewInteractionResponseWithDefaults() *InteractionResponse`

NewInteractionResponseWithDefaults instantiates a new InteractionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InteractionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InteractionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InteractionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InteractionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InteractionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InteractionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetResponseMessageId

`func (o *InteractionResponse) GetResponseMessageId() string`

GetResponseMessageId returns the ResponseMessageId field if non-nil, zero value otherwise.

### GetResponseMessageIdOk

`func (o *InteractionResponse) GetResponseMessageIdOk() (*string, bool)`

GetResponseMessageIdOk returns a tuple with the ResponseMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessageId

`func (o *InteractionResponse) SetResponseMessageId(v string)`

SetResponseMessageId sets ResponseMessageId field to given value.

### HasResponseMessageId

`func (o *InteractionResponse) HasResponseMessageId() bool`

HasResponseMessageId returns a boolean if a field has been set.

### GetResponseMessageLoading

`func (o *InteractionResponse) GetResponseMessageLoading() bool`

GetResponseMessageLoading returns the ResponseMessageLoading field if non-nil, zero value otherwise.

### GetResponseMessageLoadingOk

`func (o *InteractionResponse) GetResponseMessageLoadingOk() (*bool, bool)`

GetResponseMessageLoadingOk returns a tuple with the ResponseMessageLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessageLoading

`func (o *InteractionResponse) SetResponseMessageLoading(v bool)`

SetResponseMessageLoading sets ResponseMessageLoading field to given value.

### HasResponseMessageLoading

`func (o *InteractionResponse) HasResponseMessageLoading() bool`

HasResponseMessageLoading returns a boolean if a field has been set.

### SetResponseMessageLoadingNil

`func (o *InteractionResponse) SetResponseMessageLoadingNil(b bool)`

 SetResponseMessageLoadingNil sets the value for ResponseMessageLoading to be an explicit nil

### UnsetResponseMessageLoading
`func (o *InteractionResponse) UnsetResponseMessageLoading()`

UnsetResponseMessageLoading ensures that no value is present for ResponseMessageLoading, not even an explicit nil
### GetResponseMessageEphemeral

`func (o *InteractionResponse) GetResponseMessageEphemeral() bool`

GetResponseMessageEphemeral returns the ResponseMessageEphemeral field if non-nil, zero value otherwise.

### GetResponseMessageEphemeralOk

`func (o *InteractionResponse) GetResponseMessageEphemeralOk() (*bool, bool)`

GetResponseMessageEphemeralOk returns a tuple with the ResponseMessageEphemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessageEphemeral

`func (o *InteractionResponse) SetResponseMessageEphemeral(v bool)`

SetResponseMessageEphemeral sets ResponseMessageEphemeral field to given value.

### HasResponseMessageEphemeral

`func (o *InteractionResponse) HasResponseMessageEphemeral() bool`

HasResponseMessageEphemeral returns a boolean if a field has been set.

### SetResponseMessageEphemeralNil

`func (o *InteractionResponse) SetResponseMessageEphemeralNil(b bool)`

 SetResponseMessageEphemeralNil sets the value for ResponseMessageEphemeral to be an explicit nil

### UnsetResponseMessageEphemeral
`func (o *InteractionResponse) UnsetResponseMessageEphemeral()`

UnsetResponseMessageEphemeral ensures that no value is present for ResponseMessageEphemeral, not even an explicit nil
### GetChannelId

`func (o *InteractionResponse) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *InteractionResponse) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *InteractionResponse) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *InteractionResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetGuildId

`func (o *InteractionResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *InteractionResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *InteractionResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *InteractionResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


