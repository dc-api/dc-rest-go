# ExternalConnectionIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IntegrationTypes**](IntegrationTypes.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Account** | [**AccountResponse**](AccountResponse.md) |  | 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 
**Revoked** | Pointer to **bool** |  | [optional] 
**ExpireBehavior** | Pointer to [**IntegrationExpireBehaviorTypes**](IntegrationExpireBehaviorTypes.md) |  | [optional] 
**ExpireGracePeriod** | Pointer to [**IntegrationExpireGracePeriodTypes**](IntegrationExpireGracePeriodTypes.md) |  | [optional] 
**SubscriberCount** | Pointer to **int32** |  | [optional] 
**SyncedAt** | Pointer to **time.Time** |  | [optional] 
**RoleId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Syncing** | Pointer to **bool** |  | [optional] 
**EnableEmoticons** | Pointer to **bool** |  | [optional] 

## Methods

### NewExternalConnectionIntegrationResponse

`func NewExternalConnectionIntegrationResponse(type_ IntegrationTypes, account AccountResponse, enabled bool, id string, user UserResponse, ) *ExternalConnectionIntegrationResponse`

NewExternalConnectionIntegrationResponse instantiates a new ExternalConnectionIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalConnectionIntegrationResponseWithDefaults

`func NewExternalConnectionIntegrationResponseWithDefaults() *ExternalConnectionIntegrationResponse`

NewExternalConnectionIntegrationResponseWithDefaults instantiates a new ExternalConnectionIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalConnectionIntegrationResponse) GetType() IntegrationTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalConnectionIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalConnectionIntegrationResponse) SetType(v IntegrationTypes)`

SetType sets Type field to given value.


### GetName

`func (o *ExternalConnectionIntegrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalConnectionIntegrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalConnectionIntegrationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalConnectionIntegrationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *ExternalConnectionIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExternalConnectionIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExternalConnectionIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.


### GetEnabled

`func (o *ExternalConnectionIntegrationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExternalConnectionIntegrationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExternalConnectionIntegrationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *ExternalConnectionIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalConnectionIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalConnectionIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *ExternalConnectionIntegrationResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExternalConnectionIntegrationResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExternalConnectionIntegrationResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.


### GetRevoked

`func (o *ExternalConnectionIntegrationResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ExternalConnectionIntegrationResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ExternalConnectionIntegrationResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ExternalConnectionIntegrationResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetExpireBehavior

`func (o *ExternalConnectionIntegrationResponse) GetExpireBehavior() IntegrationExpireBehaviorTypes`

GetExpireBehavior returns the ExpireBehavior field if non-nil, zero value otherwise.

### GetExpireBehaviorOk

`func (o *ExternalConnectionIntegrationResponse) GetExpireBehaviorOk() (*IntegrationExpireBehaviorTypes, bool)`

GetExpireBehaviorOk returns a tuple with the ExpireBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireBehavior

`func (o *ExternalConnectionIntegrationResponse) SetExpireBehavior(v IntegrationExpireBehaviorTypes)`

SetExpireBehavior sets ExpireBehavior field to given value.

### HasExpireBehavior

`func (o *ExternalConnectionIntegrationResponse) HasExpireBehavior() bool`

HasExpireBehavior returns a boolean if a field has been set.

### GetExpireGracePeriod

`func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriod() IntegrationExpireGracePeriodTypes`

GetExpireGracePeriod returns the ExpireGracePeriod field if non-nil, zero value otherwise.

### GetExpireGracePeriodOk

`func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriodOk() (*IntegrationExpireGracePeriodTypes, bool)`

GetExpireGracePeriodOk returns a tuple with the ExpireGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireGracePeriod

`func (o *ExternalConnectionIntegrationResponse) SetExpireGracePeriod(v IntegrationExpireGracePeriodTypes)`

SetExpireGracePeriod sets ExpireGracePeriod field to given value.

### HasExpireGracePeriod

`func (o *ExternalConnectionIntegrationResponse) HasExpireGracePeriod() bool`

HasExpireGracePeriod returns a boolean if a field has been set.

### GetSubscriberCount

`func (o *ExternalConnectionIntegrationResponse) GetSubscriberCount() int32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *ExternalConnectionIntegrationResponse) GetSubscriberCountOk() (*int32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *ExternalConnectionIntegrationResponse) SetSubscriberCount(v int32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *ExternalConnectionIntegrationResponse) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### GetSyncedAt

`func (o *ExternalConnectionIntegrationResponse) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *ExternalConnectionIntegrationResponse) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *ExternalConnectionIntegrationResponse) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *ExternalConnectionIntegrationResponse) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### GetRoleId

`func (o *ExternalConnectionIntegrationResponse) GetRoleId() GetEntitlementsSkuIdsParameterOneOfInner`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ExternalConnectionIntegrationResponse) GetRoleIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ExternalConnectionIntegrationResponse) SetRoleId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ExternalConnectionIntegrationResponse) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSyncing

`func (o *ExternalConnectionIntegrationResponse) GetSyncing() bool`

GetSyncing returns the Syncing field if non-nil, zero value otherwise.

### GetSyncingOk

`func (o *ExternalConnectionIntegrationResponse) GetSyncingOk() (*bool, bool)`

GetSyncingOk returns a tuple with the Syncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncing

`func (o *ExternalConnectionIntegrationResponse) SetSyncing(v bool)`

SetSyncing sets Syncing field to given value.

### HasSyncing

`func (o *ExternalConnectionIntegrationResponse) HasSyncing() bool`

HasSyncing returns a boolean if a field has been set.

### GetEnableEmoticons

`func (o *ExternalConnectionIntegrationResponse) GetEnableEmoticons() bool`

GetEnableEmoticons returns the EnableEmoticons field if non-nil, zero value otherwise.

### GetEnableEmoticonsOk

`func (o *ExternalConnectionIntegrationResponse) GetEnableEmoticonsOk() (*bool, bool)`

GetEnableEmoticonsOk returns a tuple with the EnableEmoticons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmoticons

`func (o *ExternalConnectionIntegrationResponse) SetEnableEmoticons(v bool)`

SetEnableEmoticons sets EnableEmoticons field to given value.

### HasEnableEmoticons

`func (o *ExternalConnectionIntegrationResponse) HasEnableEmoticons() bool`

HasEnableEmoticons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


