# GuildWelcomeScreenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**WelcomeChannels** | [**[]GuildWelcomeScreenChannelResponse**](GuildWelcomeScreenChannelResponse.md) |  | 

## Methods

### NewGuildWelcomeScreenResponse

`func NewGuildWelcomeScreenResponse(welcomeChannels []GuildWelcomeScreenChannelResponse, ) *GuildWelcomeScreenResponse`

NewGuildWelcomeScreenResponse instantiates a new GuildWelcomeScreenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildWelcomeScreenResponseWithDefaults

`func NewGuildWelcomeScreenResponseWithDefaults() *GuildWelcomeScreenResponse`

NewGuildWelcomeScreenResponseWithDefaults instantiates a new GuildWelcomeScreenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GuildWelcomeScreenResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildWelcomeScreenResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildWelcomeScreenResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildWelcomeScreenResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildWelcomeScreenResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildWelcomeScreenResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWelcomeChannels

`func (o *GuildWelcomeScreenResponse) GetWelcomeChannels() []GuildWelcomeScreenChannelResponse`

GetWelcomeChannels returns the WelcomeChannels field if non-nil, zero value otherwise.

### GetWelcomeChannelsOk

`func (o *GuildWelcomeScreenResponse) GetWelcomeChannelsOk() (*[]GuildWelcomeScreenChannelResponse, bool)`

GetWelcomeChannelsOk returns a tuple with the WelcomeChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeChannels

`func (o *GuildWelcomeScreenResponse) SetWelcomeChannels(v []GuildWelcomeScreenChannelResponse)`

SetWelcomeChannels sets WelcomeChannels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


