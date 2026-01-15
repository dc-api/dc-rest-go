# ChannelFollowerWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**GuildId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**WebhookTypes**](WebhookTypes.md) |  | 
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**SourceGuild** | Pointer to [**WebhookSourceGuildResponse**](WebhookSourceGuildResponse.md) |  | [optional] 
**SourceChannel** | Pointer to [**WebhookSourceChannelResponse**](WebhookSourceChannelResponse.md) |  | [optional] 

## Methods

### NewChannelFollowerWebhookResponse

`func NewChannelFollowerWebhookResponse(id string, name string, type_ WebhookTypes, ) *ChannelFollowerWebhookResponse`

NewChannelFollowerWebhookResponse instantiates a new ChannelFollowerWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelFollowerWebhookResponseWithDefaults

`func NewChannelFollowerWebhookResponseWithDefaults() *ChannelFollowerWebhookResponse`

NewChannelFollowerWebhookResponseWithDefaults instantiates a new ChannelFollowerWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ChannelFollowerWebhookResponse) GetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ChannelFollowerWebhookResponse) GetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ChannelFollowerWebhookResponse) SetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ChannelFollowerWebhookResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAvatar

`func (o *ChannelFollowerWebhookResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ChannelFollowerWebhookResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ChannelFollowerWebhookResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ChannelFollowerWebhookResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetChannelId

`func (o *ChannelFollowerWebhookResponse) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChannelFollowerWebhookResponse) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChannelFollowerWebhookResponse) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ChannelFollowerWebhookResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetGuildId

`func (o *ChannelFollowerWebhookResponse) GetGuildId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ChannelFollowerWebhookResponse) GetGuildIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ChannelFollowerWebhookResponse) SetGuildId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *ChannelFollowerWebhookResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetId

`func (o *ChannelFollowerWebhookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelFollowerWebhookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelFollowerWebhookResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ChannelFollowerWebhookResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelFollowerWebhookResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelFollowerWebhookResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ChannelFollowerWebhookResponse) GetType() WebhookTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelFollowerWebhookResponse) GetTypeOk() (*WebhookTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelFollowerWebhookResponse) SetType(v WebhookTypes)`

SetType sets Type field to given value.


### GetUser

`func (o *ChannelFollowerWebhookResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChannelFollowerWebhookResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChannelFollowerWebhookResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ChannelFollowerWebhookResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSourceGuild

`func (o *ChannelFollowerWebhookResponse) GetSourceGuild() WebhookSourceGuildResponse`

GetSourceGuild returns the SourceGuild field if non-nil, zero value otherwise.

### GetSourceGuildOk

`func (o *ChannelFollowerWebhookResponse) GetSourceGuildOk() (*WebhookSourceGuildResponse, bool)`

GetSourceGuildOk returns a tuple with the SourceGuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuild

`func (o *ChannelFollowerWebhookResponse) SetSourceGuild(v WebhookSourceGuildResponse)`

SetSourceGuild sets SourceGuild field to given value.

### HasSourceGuild

`func (o *ChannelFollowerWebhookResponse) HasSourceGuild() bool`

HasSourceGuild returns a boolean if a field has been set.

### GetSourceChannel

`func (o *ChannelFollowerWebhookResponse) GetSourceChannel() WebhookSourceChannelResponse`

GetSourceChannel returns the SourceChannel field if non-nil, zero value otherwise.

### GetSourceChannelOk

`func (o *ChannelFollowerWebhookResponse) GetSourceChannelOk() (*WebhookSourceChannelResponse, bool)`

GetSourceChannelOk returns a tuple with the SourceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChannel

`func (o *ChannelFollowerWebhookResponse) SetSourceChannel(v WebhookSourceChannelResponse)`

SetSourceChannel sets SourceChannel field to given value.

### HasSourceChannel

`func (o *ChannelFollowerWebhookResponse) HasSourceChannel() bool`

HasSourceChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


