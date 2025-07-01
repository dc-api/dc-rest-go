# ConnectedAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Type** | **NullableString** |  | 
**FriendSync** | **bool** |  | 
**Integrations** | Pointer to [**[]ConnectedAccountIntegrationResponse**](ConnectedAccountIntegrationResponse.md) |  | [optional] 
**ShowActivity** | **bool** |  | 
**TwoWayLink** | **bool** |  | 
**Verified** | **bool** |  | 
**Visibility** | **NullableInt32** |  | 
**Revoked** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewConnectedAccountResponse

`func NewConnectedAccountResponse(id string, type_ NullableString, friendSync bool, showActivity bool, twoWayLink bool, verified bool, visibility NullableInt32, ) *ConnectedAccountResponse`

NewConnectedAccountResponse instantiates a new ConnectedAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedAccountResponseWithDefaults

`func NewConnectedAccountResponseWithDefaults() *ConnectedAccountResponse`

NewConnectedAccountResponseWithDefaults instantiates a new ConnectedAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectedAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectedAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectedAccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectedAccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectedAccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectedAccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectedAccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConnectedAccountResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectedAccountResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *ConnectedAccountResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectedAccountResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectedAccountResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ConnectedAccountResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ConnectedAccountResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFriendSync

`func (o *ConnectedAccountResponse) GetFriendSync() bool`

GetFriendSync returns the FriendSync field if non-nil, zero value otherwise.

### GetFriendSyncOk

`func (o *ConnectedAccountResponse) GetFriendSyncOk() (*bool, bool)`

GetFriendSyncOk returns a tuple with the FriendSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendSync

`func (o *ConnectedAccountResponse) SetFriendSync(v bool)`

SetFriendSync sets FriendSync field to given value.


### GetIntegrations

`func (o *ConnectedAccountResponse) GetIntegrations() []ConnectedAccountIntegrationResponse`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ConnectedAccountResponse) GetIntegrationsOk() (*[]ConnectedAccountIntegrationResponse, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ConnectedAccountResponse) SetIntegrations(v []ConnectedAccountIntegrationResponse)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ConnectedAccountResponse) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *ConnectedAccountResponse) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *ConnectedAccountResponse) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetShowActivity

`func (o *ConnectedAccountResponse) GetShowActivity() bool`

GetShowActivity returns the ShowActivity field if non-nil, zero value otherwise.

### GetShowActivityOk

`func (o *ConnectedAccountResponse) GetShowActivityOk() (*bool, bool)`

GetShowActivityOk returns a tuple with the ShowActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowActivity

`func (o *ConnectedAccountResponse) SetShowActivity(v bool)`

SetShowActivity sets ShowActivity field to given value.


### GetTwoWayLink

`func (o *ConnectedAccountResponse) GetTwoWayLink() bool`

GetTwoWayLink returns the TwoWayLink field if non-nil, zero value otherwise.

### GetTwoWayLinkOk

`func (o *ConnectedAccountResponse) GetTwoWayLinkOk() (*bool, bool)`

GetTwoWayLinkOk returns a tuple with the TwoWayLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoWayLink

`func (o *ConnectedAccountResponse) SetTwoWayLink(v bool)`

SetTwoWayLink sets TwoWayLink field to given value.


### GetVerified

`func (o *ConnectedAccountResponse) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ConnectedAccountResponse) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ConnectedAccountResponse) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVisibility

`func (o *ConnectedAccountResponse) GetVisibility() int32`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ConnectedAccountResponse) GetVisibilityOk() (*int32, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ConnectedAccountResponse) SetVisibility(v int32)`

SetVisibility sets Visibility field to given value.


### SetVisibilityNil

`func (o *ConnectedAccountResponse) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *ConnectedAccountResponse) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetRevoked

`func (o *ConnectedAccountResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ConnectedAccountResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ConnectedAccountResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ConnectedAccountResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *ConnectedAccountResponse) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *ConnectedAccountResponse) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


