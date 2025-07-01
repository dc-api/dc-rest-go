# EmojiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**Roles** | **[]string** |  | 
**RequireColons** | **bool** |  | 
**Managed** | **bool** |  | 
**Animated** | **bool** |  | 
**Available** | **bool** |  | 

## Methods

### NewEmojiResponse

`func NewEmojiResponse(id string, name string, roles []string, requireColons bool, managed bool, animated bool, available bool, ) *EmojiResponse`

NewEmojiResponse instantiates a new EmojiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiResponseWithDefaults

`func NewEmojiResponseWithDefaults() *EmojiResponse`

NewEmojiResponseWithDefaults instantiates a new EmojiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmojiResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmojiResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmojiResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EmojiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmojiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmojiResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *EmojiResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EmojiResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EmojiResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *EmojiResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *EmojiResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *EmojiResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRoles

`func (o *EmojiResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *EmojiResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *EmojiResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetRequireColons

`func (o *EmojiResponse) GetRequireColons() bool`

GetRequireColons returns the RequireColons field if non-nil, zero value otherwise.

### GetRequireColonsOk

`func (o *EmojiResponse) GetRequireColonsOk() (*bool, bool)`

GetRequireColonsOk returns a tuple with the RequireColons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireColons

`func (o *EmojiResponse) SetRequireColons(v bool)`

SetRequireColons sets RequireColons field to given value.


### GetManaged

`func (o *EmojiResponse) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *EmojiResponse) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *EmojiResponse) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetAnimated

`func (o *EmojiResponse) GetAnimated() bool`

GetAnimated returns the Animated field if non-nil, zero value otherwise.

### GetAnimatedOk

`func (o *EmojiResponse) GetAnimatedOk() (*bool, bool)`

GetAnimatedOk returns a tuple with the Animated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimated

`func (o *EmojiResponse) SetAnimated(v bool)`

SetAnimated sets Animated field to given value.


### GetAvailable

`func (o *EmojiResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EmojiResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EmojiResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


