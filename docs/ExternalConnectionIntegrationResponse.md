# ExternalConnectionIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableString** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**NullableAccountResponse**](AccountResponse.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**Id** | **string** |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 
**Revoked** | Pointer to **NullableBool** |  | [optional] 
**ExpireBehavior** | Pointer to **int32** |  | [optional] 
**ExpireGracePeriod** | Pointer to **int32** |  | [optional] 
**SubscriberCount** | Pointer to **NullableInt32** |  | [optional] 
**SyncedAt** | Pointer to **NullableTime** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**Syncing** | Pointer to **NullableBool** |  | [optional] 
**EnableEmoticons** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewExternalConnectionIntegrationResponse

`func NewExternalConnectionIntegrationResponse(type_ NullableString, id string, user UserResponse, ) *ExternalConnectionIntegrationResponse`

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

`func (o *ExternalConnectionIntegrationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalConnectionIntegrationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalConnectionIntegrationResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalConnectionIntegrationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalConnectionIntegrationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
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

### SetNameNil

`func (o *ExternalConnectionIntegrationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExternalConnectionIntegrationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### HasAccount

`func (o *ExternalConnectionIntegrationResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ExternalConnectionIntegrationResponse) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ExternalConnectionIntegrationResponse) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
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

### HasEnabled

`func (o *ExternalConnectionIntegrationResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ExternalConnectionIntegrationResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ExternalConnectionIntegrationResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
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

### SetRevokedNil

`func (o *ExternalConnectionIntegrationResponse) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *ExternalConnectionIntegrationResponse) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
### GetExpireBehavior

`func (o *ExternalConnectionIntegrationResponse) GetExpireBehavior() int32`

GetExpireBehavior returns the ExpireBehavior field if non-nil, zero value otherwise.

### GetExpireBehaviorOk

`func (o *ExternalConnectionIntegrationResponse) GetExpireBehaviorOk() (*int32, bool)`

GetExpireBehaviorOk returns a tuple with the ExpireBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireBehavior

`func (o *ExternalConnectionIntegrationResponse) SetExpireBehavior(v int32)`

SetExpireBehavior sets ExpireBehavior field to given value.

### HasExpireBehavior

`func (o *ExternalConnectionIntegrationResponse) HasExpireBehavior() bool`

HasExpireBehavior returns a boolean if a field has been set.

### GetExpireGracePeriod

`func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriod() int32`

GetExpireGracePeriod returns the ExpireGracePeriod field if non-nil, zero value otherwise.

### GetExpireGracePeriodOk

`func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriodOk() (*int32, bool)`

GetExpireGracePeriodOk returns a tuple with the ExpireGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireGracePeriod

`func (o *ExternalConnectionIntegrationResponse) SetExpireGracePeriod(v int32)`

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

### SetSubscriberCountNil

`func (o *ExternalConnectionIntegrationResponse) SetSubscriberCountNil(b bool)`

 SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil

### UnsetSubscriberCount
`func (o *ExternalConnectionIntegrationResponse) UnsetSubscriberCount()`

UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
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

### SetSyncedAtNil

`func (o *ExternalConnectionIntegrationResponse) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *ExternalConnectionIntegrationResponse) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
### GetRoleId

`func (o *ExternalConnectionIntegrationResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ExternalConnectionIntegrationResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ExternalConnectionIntegrationResponse) SetRoleId(v string)`

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

### SetSyncingNil

`func (o *ExternalConnectionIntegrationResponse) SetSyncingNil(b bool)`

 SetSyncingNil sets the value for Syncing to be an explicit nil

### UnsetSyncing
`func (o *ExternalConnectionIntegrationResponse) UnsetSyncing()`

UnsetSyncing ensures that no value is present for Syncing, not even an explicit nil
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

### SetEnableEmoticonsNil

`func (o *ExternalConnectionIntegrationResponse) SetEnableEmoticonsNil(b bool)`

 SetEnableEmoticonsNil sets the value for EnableEmoticons to be an explicit nil

### UnsetEnableEmoticons
`func (o *ExternalConnectionIntegrationResponse) UnsetEnableEmoticons()`

UnsetEnableEmoticons ensures that no value is present for EnableEmoticons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


