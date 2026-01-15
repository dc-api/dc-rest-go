# ApplicationIncomingWebhookResponse

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

## Methods

### NewApplicationIncomingWebhookResponse

`func NewApplicationIncomingWebhookResponse(id string, name string, type_ WebhookTypes, ) *ApplicationIncomingWebhookResponse`

NewApplicationIncomingWebhookResponse instantiates a new ApplicationIncomingWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationIncomingWebhookResponseWithDefaults

`func NewApplicationIncomingWebhookResponseWithDefaults() *ApplicationIncomingWebhookResponse`

NewApplicationIncomingWebhookResponseWithDefaults instantiates a new ApplicationIncomingWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationIncomingWebhookResponse) GetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationIncomingWebhookResponse) GetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationIncomingWebhookResponse) SetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationIncomingWebhookResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAvatar

`func (o *ApplicationIncomingWebhookResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ApplicationIncomingWebhookResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ApplicationIncomingWebhookResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ApplicationIncomingWebhookResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetChannelId

`func (o *ApplicationIncomingWebhookResponse) GetChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ApplicationIncomingWebhookResponse) GetChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ApplicationIncomingWebhookResponse) SetChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ApplicationIncomingWebhookResponse) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetGuildId

`func (o *ApplicationIncomingWebhookResponse) GetGuildId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ApplicationIncomingWebhookResponse) GetGuildIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ApplicationIncomingWebhookResponse) SetGuildId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *ApplicationIncomingWebhookResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetId

`func (o *ApplicationIncomingWebhookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationIncomingWebhookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationIncomingWebhookResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationIncomingWebhookResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationIncomingWebhookResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationIncomingWebhookResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApplicationIncomingWebhookResponse) GetType() WebhookTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationIncomingWebhookResponse) GetTypeOk() (*WebhookTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationIncomingWebhookResponse) SetType(v WebhookTypes)`

SetType sets Type field to given value.


### GetUser

`func (o *ApplicationIncomingWebhookResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApplicationIncomingWebhookResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApplicationIncomingWebhookResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ApplicationIncomingWebhookResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


