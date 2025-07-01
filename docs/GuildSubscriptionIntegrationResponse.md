# GuildSubscriptionIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableString** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**NullableAccountResponse**](AccountResponse.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewGuildSubscriptionIntegrationResponse

`func NewGuildSubscriptionIntegrationResponse(type_ NullableString, id string, ) *GuildSubscriptionIntegrationResponse`

NewGuildSubscriptionIntegrationResponse instantiates a new GuildSubscriptionIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildSubscriptionIntegrationResponseWithDefaults

`func NewGuildSubscriptionIntegrationResponseWithDefaults() *GuildSubscriptionIntegrationResponse`

NewGuildSubscriptionIntegrationResponseWithDefaults instantiates a new GuildSubscriptionIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GuildSubscriptionIntegrationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildSubscriptionIntegrationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildSubscriptionIntegrationResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GuildSubscriptionIntegrationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GuildSubscriptionIntegrationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *GuildSubscriptionIntegrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildSubscriptionIntegrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildSubscriptionIntegrationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuildSubscriptionIntegrationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GuildSubscriptionIntegrationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GuildSubscriptionIntegrationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccount

`func (o *GuildSubscriptionIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GuildSubscriptionIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GuildSubscriptionIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GuildSubscriptionIntegrationResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GuildSubscriptionIntegrationResponse) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GuildSubscriptionIntegrationResponse) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetEnabled

`func (o *GuildSubscriptionIntegrationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuildSubscriptionIntegrationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuildSubscriptionIntegrationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GuildSubscriptionIntegrationResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *GuildSubscriptionIntegrationResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *GuildSubscriptionIntegrationResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetId

`func (o *GuildSubscriptionIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildSubscriptionIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildSubscriptionIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


