# RoleSelectComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **string** |  | [optional] 
**MinValues** | Pointer to **int32** |  | [optional] 
**MaxValues** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DefaultValues** | Pointer to [**[]RoleSelectDefaultValueResponse**](RoleSelectDefaultValueResponse.md) |  | [optional] 

## Methods

### NewRoleSelectComponentResponse

`func NewRoleSelectComponentResponse(type_ MessageComponentTypes, id int32, customId string, ) *RoleSelectComponentResponse`

NewRoleSelectComponentResponse instantiates a new RoleSelectComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleSelectComponentResponseWithDefaults

`func NewRoleSelectComponentResponseWithDefaults() *RoleSelectComponentResponse`

NewRoleSelectComponentResponseWithDefaults instantiates a new RoleSelectComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoleSelectComponentResponse) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleSelectComponentResponse) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleSelectComponentResponse) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *RoleSelectComponentResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleSelectComponentResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleSelectComponentResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *RoleSelectComponentResponse) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *RoleSelectComponentResponse) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *RoleSelectComponentResponse) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *RoleSelectComponentResponse) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *RoleSelectComponentResponse) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *RoleSelectComponentResponse) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *RoleSelectComponentResponse) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMinValues

`func (o *RoleSelectComponentResponse) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *RoleSelectComponentResponse) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *RoleSelectComponentResponse) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *RoleSelectComponentResponse) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### GetMaxValues

`func (o *RoleSelectComponentResponse) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *RoleSelectComponentResponse) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *RoleSelectComponentResponse) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *RoleSelectComponentResponse) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### GetDisabled

`func (o *RoleSelectComponentResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RoleSelectComponentResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RoleSelectComponentResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RoleSelectComponentResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDefaultValues

`func (o *RoleSelectComponentResponse) GetDefaultValues() []RoleSelectDefaultValueResponse`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *RoleSelectComponentResponse) GetDefaultValuesOk() (*[]RoleSelectDefaultValueResponse, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *RoleSelectComponentResponse) SetDefaultValues(v []RoleSelectDefaultValueResponse)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *RoleSelectComponentResponse) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


