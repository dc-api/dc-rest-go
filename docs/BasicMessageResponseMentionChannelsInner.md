# BasicMessageResponseMentionChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**GuildId** | **string** |  | 

## Methods

### NewBasicMessageResponseMentionChannelsInner

`func NewBasicMessageResponseMentionChannelsInner(id string, name string, type_ ChannelTypes, guildId string, ) *BasicMessageResponseMentionChannelsInner`

NewBasicMessageResponseMentionChannelsInner instantiates a new BasicMessageResponseMentionChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicMessageResponseMentionChannelsInnerWithDefaults

`func NewBasicMessageResponseMentionChannelsInnerWithDefaults() *BasicMessageResponseMentionChannelsInner`

NewBasicMessageResponseMentionChannelsInnerWithDefaults instantiates a new BasicMessageResponseMentionChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicMessageResponseMentionChannelsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicMessageResponseMentionChannelsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicMessageResponseMentionChannelsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BasicMessageResponseMentionChannelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicMessageResponseMentionChannelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicMessageResponseMentionChannelsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BasicMessageResponseMentionChannelsInner) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicMessageResponseMentionChannelsInner) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicMessageResponseMentionChannelsInner) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetGuildId

`func (o *BasicMessageResponseMentionChannelsInner) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *BasicMessageResponseMentionChannelsInner) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *BasicMessageResponseMentionChannelsInner) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


