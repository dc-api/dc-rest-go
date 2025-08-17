# StringSelectComponentForMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**MinValues** | Pointer to **NullableInt32** |  | [optional] 
**MaxValues** | Pointer to **NullableInt32** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Options** | [**[]StringSelectOptionForRequest**](StringSelectOptionForRequest.md) |  | 

## Methods

### NewStringSelectComponentForMessageRequest

`func NewStringSelectComponentForMessageRequest(type_ int32, customId string, options []StringSelectOptionForRequest, ) *StringSelectComponentForMessageRequest`

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

`func (o *StringSelectComponentForMessageRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringSelectComponentForMessageRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringSelectComponentForMessageRequest) SetType(v int32)`

SetType sets Type field to given value.


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

### SetPlaceholderNil

`func (o *StringSelectComponentForMessageRequest) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *StringSelectComponentForMessageRequest) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
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

### SetMinValuesNil

`func (o *StringSelectComponentForMessageRequest) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *StringSelectComponentForMessageRequest) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
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

### SetMaxValuesNil

`func (o *StringSelectComponentForMessageRequest) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *StringSelectComponentForMessageRequest) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
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

### SetDisabledNil

`func (o *StringSelectComponentForMessageRequest) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *StringSelectComponentForMessageRequest) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
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


