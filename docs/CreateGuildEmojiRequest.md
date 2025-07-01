# CreateGuildEmojiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Image** | **string** |  | 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGuildEmojiRequest

`func NewCreateGuildEmojiRequest(name string, image string, ) *CreateGuildEmojiRequest`

NewCreateGuildEmojiRequest instantiates a new CreateGuildEmojiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildEmojiRequestWithDefaults

`func NewCreateGuildEmojiRequestWithDefaults() *CreateGuildEmojiRequest`

NewCreateGuildEmojiRequestWithDefaults instantiates a new CreateGuildEmojiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGuildEmojiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildEmojiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildEmojiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *CreateGuildEmojiRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateGuildEmojiRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateGuildEmojiRequest) SetImage(v string)`

SetImage sets Image field to given value.


### GetRoles

`func (o *CreateGuildEmojiRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateGuildEmojiRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateGuildEmojiRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateGuildEmojiRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CreateGuildEmojiRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CreateGuildEmojiRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


