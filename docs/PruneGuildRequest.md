# PruneGuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **NullableInt32** |  | [optional] 
**ComputePruneCount** | Pointer to **NullableBool** |  | [optional] 
**IncludeRoles** | Pointer to [**NullablePruneGuildRequestIncludeRoles**](PruneGuildRequestIncludeRoles.md) |  | [optional] 

## Methods

### NewPruneGuildRequest

`func NewPruneGuildRequest() *PruneGuildRequest`

NewPruneGuildRequest instantiates a new PruneGuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPruneGuildRequestWithDefaults

`func NewPruneGuildRequestWithDefaults() *PruneGuildRequest`

NewPruneGuildRequestWithDefaults instantiates a new PruneGuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *PruneGuildRequest) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PruneGuildRequest) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PruneGuildRequest) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *PruneGuildRequest) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *PruneGuildRequest) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *PruneGuildRequest) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetComputePruneCount

`func (o *PruneGuildRequest) GetComputePruneCount() bool`

GetComputePruneCount returns the ComputePruneCount field if non-nil, zero value otherwise.

### GetComputePruneCountOk

`func (o *PruneGuildRequest) GetComputePruneCountOk() (*bool, bool)`

GetComputePruneCountOk returns a tuple with the ComputePruneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePruneCount

`func (o *PruneGuildRequest) SetComputePruneCount(v bool)`

SetComputePruneCount sets ComputePruneCount field to given value.

### HasComputePruneCount

`func (o *PruneGuildRequest) HasComputePruneCount() bool`

HasComputePruneCount returns a boolean if a field has been set.

### SetComputePruneCountNil

`func (o *PruneGuildRequest) SetComputePruneCountNil(b bool)`

 SetComputePruneCountNil sets the value for ComputePruneCount to be an explicit nil

### UnsetComputePruneCount
`func (o *PruneGuildRequest) UnsetComputePruneCount()`

UnsetComputePruneCount ensures that no value is present for ComputePruneCount, not even an explicit nil
### GetIncludeRoles

`func (o *PruneGuildRequest) GetIncludeRoles() PruneGuildRequestIncludeRoles`

GetIncludeRoles returns the IncludeRoles field if non-nil, zero value otherwise.

### GetIncludeRolesOk

`func (o *PruneGuildRequest) GetIncludeRolesOk() (*PruneGuildRequestIncludeRoles, bool)`

GetIncludeRolesOk returns a tuple with the IncludeRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRoles

`func (o *PruneGuildRequest) SetIncludeRoles(v PruneGuildRequestIncludeRoles)`

SetIncludeRoles sets IncludeRoles field to given value.

### HasIncludeRoles

`func (o *PruneGuildRequest) HasIncludeRoles() bool`

HasIncludeRoles returns a boolean if a field has been set.

### SetIncludeRolesNil

`func (o *PruneGuildRequest) SetIncludeRolesNil(b bool)`

 SetIncludeRolesNil sets the value for IncludeRoles to be an explicit nil

### UnsetIncludeRoles
`func (o *PruneGuildRequest) UnsetIncludeRoles()`

UnsetIncludeRoles ensures that no value is present for IncludeRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


