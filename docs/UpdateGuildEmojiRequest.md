# UpdateGuildEmojiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateGuildEmojiRequest

`func NewUpdateGuildEmojiRequest() *UpdateGuildEmojiRequest`

NewUpdateGuildEmojiRequest instantiates a new UpdateGuildEmojiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildEmojiRequestWithDefaults

`func NewUpdateGuildEmojiRequestWithDefaults() *UpdateGuildEmojiRequest`

NewUpdateGuildEmojiRequestWithDefaults instantiates a new UpdateGuildEmojiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGuildEmojiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGuildEmojiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGuildEmojiRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGuildEmojiRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateGuildEmojiRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateGuildEmojiRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateGuildEmojiRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateGuildEmojiRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UpdateGuildEmojiRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UpdateGuildEmojiRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


