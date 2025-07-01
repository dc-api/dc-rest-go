# GuildAuditLogResponseIntegrationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **NullableString** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**AccountResponse**](AccountResponse.md) |  | [optional] 
**ApplicationId** | **string** |  | 

## Methods

### NewGuildAuditLogResponseIntegrationsInner

`func NewGuildAuditLogResponseIntegrationsInner(id string, type_ NullableString, applicationId string, ) *GuildAuditLogResponseIntegrationsInner`

NewGuildAuditLogResponseIntegrationsInner instantiates a new GuildAuditLogResponseIntegrationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildAuditLogResponseIntegrationsInnerWithDefaults

`func NewGuildAuditLogResponseIntegrationsInnerWithDefaults() *GuildAuditLogResponseIntegrationsInner`

NewGuildAuditLogResponseIntegrationsInnerWithDefaults instantiates a new GuildAuditLogResponseIntegrationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildAuditLogResponseIntegrationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildAuditLogResponseIntegrationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildAuditLogResponseIntegrationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GuildAuditLogResponseIntegrationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildAuditLogResponseIntegrationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildAuditLogResponseIntegrationsInner) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GuildAuditLogResponseIntegrationsInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GuildAuditLogResponseIntegrationsInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *GuildAuditLogResponseIntegrationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildAuditLogResponseIntegrationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildAuditLogResponseIntegrationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuildAuditLogResponseIntegrationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *GuildAuditLogResponseIntegrationsInner) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GuildAuditLogResponseIntegrationsInner) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GuildAuditLogResponseIntegrationsInner) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GuildAuditLogResponseIntegrationsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApplicationId

`func (o *GuildAuditLogResponseIntegrationsInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GuildAuditLogResponseIntegrationsInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GuildAuditLogResponseIntegrationsInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


