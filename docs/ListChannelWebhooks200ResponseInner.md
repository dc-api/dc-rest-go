# ListChannelWebhooks200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | **int32** |  | 
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**SourceGuild** | Pointer to [**WebhookSourceGuildResponse**](WebhookSourceGuildResponse.md) |  | [optional] 
**SourceChannel** | Pointer to [**WebhookSourceChannelResponse**](WebhookSourceChannelResponse.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewListChannelWebhooks200ResponseInner

`func NewListChannelWebhooks200ResponseInner(id string, name string, type_ int32, ) *ListChannelWebhooks200ResponseInner`

NewListChannelWebhooks200ResponseInner instantiates a new ListChannelWebhooks200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChannelWebhooks200ResponseInnerWithDefaults

`func NewListChannelWebhooks200ResponseInnerWithDefaults() *ListChannelWebhooks200ResponseInner`

NewListChannelWebhooks200ResponseInnerWithDefaults instantiates a new ListChannelWebhooks200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ListChannelWebhooks200ResponseInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ListChannelWebhooks200ResponseInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ListChannelWebhooks200ResponseInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ListChannelWebhooks200ResponseInner) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAvatar

`func (o *ListChannelWebhooks200ResponseInner) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ListChannelWebhooks200ResponseInner) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ListChannelWebhooks200ResponseInner) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ListChannelWebhooks200ResponseInner) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetChannelId

`func (o *ListChannelWebhooks200ResponseInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ListChannelWebhooks200ResponseInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ListChannelWebhooks200ResponseInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ListChannelWebhooks200ResponseInner) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetGuildId

`func (o *ListChannelWebhooks200ResponseInner) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ListChannelWebhooks200ResponseInner) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ListChannelWebhooks200ResponseInner) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *ListChannelWebhooks200ResponseInner) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetId

`func (o *ListChannelWebhooks200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListChannelWebhooks200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListChannelWebhooks200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListChannelWebhooks200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListChannelWebhooks200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListChannelWebhooks200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ListChannelWebhooks200ResponseInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListChannelWebhooks200ResponseInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListChannelWebhooks200ResponseInner) SetType(v int32)`

SetType sets Type field to given value.


### GetUser

`func (o *ListChannelWebhooks200ResponseInner) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListChannelWebhooks200ResponseInner) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListChannelWebhooks200ResponseInner) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ListChannelWebhooks200ResponseInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSourceGuild

`func (o *ListChannelWebhooks200ResponseInner) GetSourceGuild() WebhookSourceGuildResponse`

GetSourceGuild returns the SourceGuild field if non-nil, zero value otherwise.

### GetSourceGuildOk

`func (o *ListChannelWebhooks200ResponseInner) GetSourceGuildOk() (*WebhookSourceGuildResponse, bool)`

GetSourceGuildOk returns a tuple with the SourceGuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuild

`func (o *ListChannelWebhooks200ResponseInner) SetSourceGuild(v WebhookSourceGuildResponse)`

SetSourceGuild sets SourceGuild field to given value.

### HasSourceGuild

`func (o *ListChannelWebhooks200ResponseInner) HasSourceGuild() bool`

HasSourceGuild returns a boolean if a field has been set.

### GetSourceChannel

`func (o *ListChannelWebhooks200ResponseInner) GetSourceChannel() WebhookSourceChannelResponse`

GetSourceChannel returns the SourceChannel field if non-nil, zero value otherwise.

### GetSourceChannelOk

`func (o *ListChannelWebhooks200ResponseInner) GetSourceChannelOk() (*WebhookSourceChannelResponse, bool)`

GetSourceChannelOk returns a tuple with the SourceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChannel

`func (o *ListChannelWebhooks200ResponseInner) SetSourceChannel(v WebhookSourceChannelResponse)`

SetSourceChannel sets SourceChannel field to given value.

### HasSourceChannel

`func (o *ListChannelWebhooks200ResponseInner) HasSourceChannel() bool`

HasSourceChannel returns a boolean if a field has been set.

### GetToken

`func (o *ListChannelWebhooks200ResponseInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListChannelWebhooks200ResponseInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListChannelWebhooks200ResponseInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListChannelWebhooks200ResponseInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUrl

`func (o *ListChannelWebhooks200ResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListChannelWebhooks200ResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListChannelWebhooks200ResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListChannelWebhooks200ResponseInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


