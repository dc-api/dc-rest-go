# StringSelectComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | Pointer to **int32** |  | [optional] 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **string** |  | [optional] 
**MinValues** | Pointer to **int32** |  | [optional] 
**MaxValues** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Options** | [**[]StringSelectOptionForRequest**](StringSelectOptionForRequest.md) |  | 

## Methods

### NewStringSelectComponentForMessageRequest

`func NewStringSelectComponentForMessageRequest(type_ MessageComponentTypes, customId string, options []StringSelectOptionForRequest, ) *StringSelectComponentForMessageRequest`

NewStringSelectComponentForMessageRequest instantiates a new StringSelectComponentForMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSelectComponentForMessageRequestWithDefaults

`func NewStringSelectComponentForMessageRequestWithDefaults() *StringSelectComponentForMessageRequest`

NewStringSelectComponentForMessageRequestWithDefaults instantiates a new StringSelectComponentForMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StringSelectComponentForMessageRequest) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringSelectComponentForMessageRequest) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringSelectComponentForMessageRequest) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *StringSelectComponentForMessageRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StringSelectComponentForMessageRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StringSelectComponentForMessageRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StringSelectComponentForMessageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomId

`func (o *StringSelectComponentForMessageRequest) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *StringSelectComponentForMessageRequest) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *StringSelectComponentForMessageRequest) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *StringSelectComponentForMessageRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *StringSelectComponentForMessageRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *StringSelectComponentForMessageRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *StringSelectComponentForMessageRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMinValues

`func (o *StringSelectComponentForMessageRequest) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *StringSelectComponentForMessageRequest) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *StringSelectComponentForMessageRequest) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *StringSelectComponentForMessageRequest) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### GetMaxValues

`func (o *StringSelectComponentForMessageRequest) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *StringSelectComponentForMessageRequest) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *StringSelectComponentForMessageRequest) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *StringSelectComponentForMessageRequest) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### GetDisabled

`func (o *StringSelectComponentForMessageRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *StringSelectComponentForMessageRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *StringSelectComponentForMessageRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *StringSelectComponentForMessageRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetRequired

`func (o *StringSelectComponentForMessageRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StringSelectComponentForMessageRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StringSelectComponentForMessageRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *StringSelectComponentForMessageRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOptions

`func (o *StringSelectComponentForMessageRequest) GetOptions() []StringSelectOptionForRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StringSelectComponentForMessageRequest) GetOptionsOk() (*[]StringSelectOptionForRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StringSelectComponentForMessageRequest) SetOptions(v []StringSelectOptionForRequest)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


