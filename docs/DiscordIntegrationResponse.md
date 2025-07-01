# DiscordIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableString** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**NullableAccountResponse**](AccountResponse.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**Id** | **string** |  | 
**Application** | [**IntegrationApplicationResponse**](IntegrationApplicationResponse.md) |  | 
**Scopes** | **[]string** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewDiscordIntegrationResponse

`func NewDiscordIntegrationResponse(type_ NullableString, id string, application IntegrationApplicationResponse, scopes []string, ) *DiscordIntegrationResponse`

NewDiscordIntegrationResponse instantiates a new DiscordIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordIntegrationResponseWithDefaults

`func NewDiscordIntegrationResponseWithDefaults() *DiscordIntegrationResponse`

NewDiscordIntegrationResponseWithDefaults instantiates a new DiscordIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscordIntegrationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscordIntegrationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscordIntegrationResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *DiscordIntegrationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DiscordIntegrationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *DiscordIntegrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscordIntegrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscordIntegrationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscordIntegrationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DiscordIntegrationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DiscordIntegrationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccount

`func (o *DiscordIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DiscordIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DiscordIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DiscordIntegrationResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *DiscordIntegrationResponse) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *DiscordIntegrationResponse) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetEnabled

`func (o *DiscordIntegrationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiscordIntegrationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiscordIntegrationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DiscordIntegrationResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *DiscordIntegrationResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *DiscordIntegrationResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetId

`func (o *DiscordIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscordIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscordIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApplication

`func (o *DiscordIntegrationResponse) GetApplication() IntegrationApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DiscordIntegrationResponse) GetApplicationOk() (*IntegrationApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DiscordIntegrationResponse) SetApplication(v IntegrationApplicationResponse)`

SetApplication sets Application field to given value.


### GetScopes

`func (o *DiscordIntegrationResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DiscordIntegrationResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DiscordIntegrationResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetUser

`func (o *DiscordIntegrationResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DiscordIntegrationResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DiscordIntegrationResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *DiscordIntegrationResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *DiscordIntegrationResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *DiscordIntegrationResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


