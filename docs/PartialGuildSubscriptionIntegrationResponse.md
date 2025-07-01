# PartialGuildSubscriptionIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **NullableString** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**NullableAccountResponse**](AccountResponse.md) |  | [optional] 

## Methods

### NewPartialGuildSubscriptionIntegrationResponse

`func NewPartialGuildSubscriptionIntegrationResponse(id string, type_ NullableString, ) *PartialGuildSubscriptionIntegrationResponse`

NewPartialGuildSubscriptionIntegrationResponse instantiates a new PartialGuildSubscriptionIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialGuildSubscriptionIntegrationResponseWithDefaults

`func NewPartialGuildSubscriptionIntegrationResponseWithDefaults() *PartialGuildSubscriptionIntegrationResponse`

NewPartialGuildSubscriptionIntegrationResponseWithDefaults instantiates a new PartialGuildSubscriptionIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartialGuildSubscriptionIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialGuildSubscriptionIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PartialGuildSubscriptionIntegrationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartialGuildSubscriptionIntegrationResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PartialGuildSubscriptionIntegrationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PartialGuildSubscriptionIntegrationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *PartialGuildSubscriptionIntegrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartialGuildSubscriptionIntegrationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartialGuildSubscriptionIntegrationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PartialGuildSubscriptionIntegrationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PartialGuildSubscriptionIntegrationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccount

`func (o *PartialGuildSubscriptionIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartialGuildSubscriptionIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartialGuildSubscriptionIntegrationResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PartialGuildSubscriptionIntegrationResponse) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PartialGuildSubscriptionIntegrationResponse) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


