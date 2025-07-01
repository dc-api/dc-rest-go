# CreateGuildInviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAge** | Pointer to **NullableInt32** |  | [optional] 
**Temporary** | Pointer to **NullableBool** |  | [optional] 
**MaxUses** | Pointer to **NullableInt32** |  | [optional] 
**Unique** | Pointer to **NullableBool** |  | [optional] 
**TargetUserId** | Pointer to **string** |  | [optional] 
**TargetApplicationId** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **NullableInt32** |  | [optional] 

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

### SetMaxAgeNil

`func (o *CreateGuildInviteRequest) SetMaxAgeNil(b bool)`

 SetMaxAgeNil sets the value for MaxAge to be an explicit nil

### UnsetMaxAge
`func (o *CreateGuildInviteRequest) UnsetMaxAge()`

UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
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

### SetTemporaryNil

`func (o *CreateGuildInviteRequest) SetTemporaryNil(b bool)`

 SetTemporaryNil sets the value for Temporary to be an explicit nil

### UnsetTemporary
`func (o *CreateGuildInviteRequest) UnsetTemporary()`

UnsetTemporary ensures that no value is present for Temporary, not even an explicit nil
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

### SetMaxUsesNil

`func (o *CreateGuildInviteRequest) SetMaxUsesNil(b bool)`

 SetMaxUsesNil sets the value for MaxUses to be an explicit nil

### UnsetMaxUses
`func (o *CreateGuildInviteRequest) UnsetMaxUses()`

UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
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

### SetUniqueNil

`func (o *CreateGuildInviteRequest) SetUniqueNil(b bool)`

 SetUniqueNil sets the value for Unique to be an explicit nil

### UnsetUnique
`func (o *CreateGuildInviteRequest) UnsetUnique()`

UnsetUnique ensures that no value is present for Unique, not even an explicit nil
### GetTargetUserId

`func (o *CreateGuildInviteRequest) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *CreateGuildInviteRequest) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *CreateGuildInviteRequest) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.

### HasTargetUserId

`func (o *CreateGuildInviteRequest) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.

### GetTargetApplicationId

`func (o *CreateGuildInviteRequest) GetTargetApplicationId() string`

GetTargetApplicationId returns the TargetApplicationId field if non-nil, zero value otherwise.

### GetTargetApplicationIdOk

`func (o *CreateGuildInviteRequest) GetTargetApplicationIdOk() (*string, bool)`

GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplicationId

`func (o *CreateGuildInviteRequest) SetTargetApplicationId(v string)`

SetTargetApplicationId sets TargetApplicationId field to given value.

### HasTargetApplicationId

`func (o *CreateGuildInviteRequest) HasTargetApplicationId() bool`

HasTargetApplicationId returns a boolean if a field has been set.

### GetTargetType

`func (o *CreateGuildInviteRequest) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CreateGuildInviteRequest) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CreateGuildInviteRequest) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CreateGuildInviteRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *CreateGuildInviteRequest) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *CreateGuildInviteRequest) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


