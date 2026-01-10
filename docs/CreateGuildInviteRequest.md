# CreateGuildInviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAge** | Pointer to **int32** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**TargetUserId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**TargetApplicationId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**TargetType** | Pointer to [**CreateGuildInviteRequestTargetType**](CreateGuildInviteRequestTargetType.md) |  | [optional] 

## Methods

### NewCreateGuildInviteRequest

`func NewCreateGuildInviteRequest() *CreateGuildInviteRequest`

NewCreateGuildInviteRequest instantiates a new CreateGuildInviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildInviteRequestWithDefaults

`func NewCreateGuildInviteRequestWithDefaults() *CreateGuildInviteRequest`

NewCreateGuildInviteRequestWithDefaults instantiates a new CreateGuildInviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAge

`func (o *CreateGuildInviteRequest) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CreateGuildInviteRequest) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CreateGuildInviteRequest) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CreateGuildInviteRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetTemporary

`func (o *CreateGuildInviteRequest) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CreateGuildInviteRequest) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CreateGuildInviteRequest) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *CreateGuildInviteRequest) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetMaxUses

`func (o *CreateGuildInviteRequest) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *CreateGuildInviteRequest) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *CreateGuildInviteRequest) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *CreateGuildInviteRequest) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetUnique

`func (o *CreateGuildInviteRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *CreateGuildInviteRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *CreateGuildInviteRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *CreateGuildInviteRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetTargetUserId

`func (o *CreateGuildInviteRequest) GetTargetUserId() GetEntitlementsSkuIdsParameterOneOfInner`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *CreateGuildInviteRequest) GetTargetUserIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *CreateGuildInviteRequest) SetTargetUserId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetTargetUserId sets TargetUserId field to given value.

### HasTargetUserId

`func (o *CreateGuildInviteRequest) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.

### GetTargetApplicationId

`func (o *CreateGuildInviteRequest) GetTargetApplicationId() GetEntitlementsSkuIdsParameterOneOfInner`

GetTargetApplicationId returns the TargetApplicationId field if non-nil, zero value otherwise.

### GetTargetApplicationIdOk

`func (o *CreateGuildInviteRequest) GetTargetApplicationIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplicationId

`func (o *CreateGuildInviteRequest) SetTargetApplicationId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetTargetApplicationId sets TargetApplicationId field to given value.

### HasTargetApplicationId

`func (o *CreateGuildInviteRequest) HasTargetApplicationId() bool`

HasTargetApplicationId returns a boolean if a field has been set.

### GetTargetType

`func (o *CreateGuildInviteRequest) GetTargetType() CreateGuildInviteRequestTargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CreateGuildInviteRequest) GetTargetTypeOk() (*CreateGuildInviteRequestTargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CreateGuildInviteRequest) SetTargetType(v CreateGuildInviteRequestTargetType)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CreateGuildInviteRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


