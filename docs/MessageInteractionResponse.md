# MessageInteractionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**Name** | **string** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**NameLocalized** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMessageInteractionResponse

`func NewMessageInteractionResponse(id string, type_ int32, name string, ) *MessageInteractionResponse`

NewMessageInteractionResponse instantiates a new MessageInteractionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageInteractionResponseWithDefaults

`func NewMessageInteractionResponseWithDefaults() *MessageInteractionResponse`

NewMessageInteractionResponseWithDefaults instantiates a new MessageInteractionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageInteractionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageInteractionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageInteractionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MessageInteractionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageInteractionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageInteractionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *MessageInteractionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageInteractionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageInteractionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *MessageInteractionResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MessageInteractionResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MessageInteractionResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *MessageInteractionResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MessageInteractionResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MessageInteractionResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetNameLocalized

`func (o *MessageInteractionResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *MessageInteractionResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *MessageInteractionResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *MessageInteractionResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *MessageInteractionResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *MessageInteractionResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


