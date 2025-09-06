# MentionableSelectComponentForModalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Id** | Pointer to **NullableInt32** |  | [optional] 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**MinValues** | Pointer to **NullableInt32** |  | [optional] 
**MaxValues** | Pointer to **NullableInt32** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**DefaultValues** | Pointer to [**[]MentionableSelectComponentForMessageRequestDefaultValuesInner**](MentionableSelectComponentForMessageRequestDefaultValuesInner.md) |  | [optional] 

## Methods

### NewMentionableSelectComponentForModalRequest

`func NewMentionableSelectComponentForModalRequest(type_ int32, customId string, ) *MentionableSelectComponentForModalRequest`

NewMentionableSelectComponentForModalRequest instantiates a new MentionableSelectComponentForModalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionableSelectComponentForModalRequestWithDefaults

`func NewMentionableSelectComponentForModalRequestWithDefaults() *MentionableSelectComponentForModalRequest`

NewMentionableSelectComponentForModalRequestWithDefaults instantiates a new MentionableSelectComponentForModalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MentionableSelectComponentForModalRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MentionableSelectComponentForModalRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MentionableSelectComponentForModalRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetId

`func (o *MentionableSelectComponentForModalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MentionableSelectComponentForModalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MentionableSelectComponentForModalRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MentionableSelectComponentForModalRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MentionableSelectComponentForModalRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MentionableSelectComponentForModalRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCustomId

`func (o *MentionableSelectComponentForModalRequest) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MentionableSelectComponentForModalRequest) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MentionableSelectComponentForModalRequest) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *MentionableSelectComponentForModalRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *MentionableSelectComponentForModalRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *MentionableSelectComponentForModalRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *MentionableSelectComponentForModalRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *MentionableSelectComponentForModalRequest) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *MentionableSelectComponentForModalRequest) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetMinValues

`func (o *MentionableSelectComponentForModalRequest) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *MentionableSelectComponentForModalRequest) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *MentionableSelectComponentForModalRequest) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *MentionableSelectComponentForModalRequest) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### SetMinValuesNil

`func (o *MentionableSelectComponentForModalRequest) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *MentionableSelectComponentForModalRequest) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
### GetMaxValues

`func (o *MentionableSelectComponentForModalRequest) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *MentionableSelectComponentForModalRequest) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *MentionableSelectComponentForModalRequest) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *MentionableSelectComponentForModalRequest) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### SetMaxValuesNil

`func (o *MentionableSelectComponentForModalRequest) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *MentionableSelectComponentForModalRequest) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
### GetDisabled

`func (o *MentionableSelectComponentForModalRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MentionableSelectComponentForModalRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MentionableSelectComponentForModalRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MentionableSelectComponentForModalRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *MentionableSelectComponentForModalRequest) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *MentionableSelectComponentForModalRequest) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetRequired

`func (o *MentionableSelectComponentForModalRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MentionableSelectComponentForModalRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *MentionableSelectComponentForModalRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *MentionableSelectComponentForModalRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *MentionableSelectComponentForModalRequest) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *MentionableSelectComponentForModalRequest) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetDefaultValues

`func (o *MentionableSelectComponentForModalRequest) GetDefaultValues() []MentionableSelectComponentForMessageRequestDefaultValuesInner`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *MentionableSelectComponentForModalRequest) GetDefaultValuesOk() (*[]MentionableSelectComponentForMessageRequestDefaultValuesInner, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *MentionableSelectComponentForModalRequest) SetDefaultValues(v []MentionableSelectComponentForMessageRequestDefaultValuesInner)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *MentionableSelectComponentForModalRequest) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### SetDefaultValuesNil

`func (o *MentionableSelectComponentForModalRequest) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *MentionableSelectComponentForModalRequest) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


