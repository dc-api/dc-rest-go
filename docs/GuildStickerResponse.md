# GuildStickerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Tags** | **string** |  | 
**Type** | **int32** |  | 
**FormatType** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Available** | **bool** |  | 
**GuildId** | **string** |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewGuildStickerResponse

`func NewGuildStickerResponse(id string, name string, tags string, type_ int32, available bool, guildId string, ) *GuildStickerResponse`

NewGuildStickerResponse instantiates a new GuildStickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildStickerResponseWithDefaults

`func NewGuildStickerResponseWithDefaults() *GuildStickerResponse`

NewGuildStickerResponseWithDefaults instantiates a new GuildStickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuildStickerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuildStickerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuildStickerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GuildStickerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildStickerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildStickerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *GuildStickerResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GuildStickerResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GuildStickerResponse) SetTags(v string)`

SetTags sets Tags field to given value.


### GetType

`func (o *GuildStickerResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuildStickerResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuildStickerResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetFormatType

`func (o *GuildStickerResponse) GetFormatType() int32`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *GuildStickerResponse) GetFormatTypeOk() (*int32, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *GuildStickerResponse) SetFormatType(v int32)`

SetFormatType sets FormatType field to given value.

### HasFormatType

`func (o *GuildStickerResponse) HasFormatType() bool`

HasFormatType returns a boolean if a field has been set.

### SetFormatTypeNil

`func (o *GuildStickerResponse) SetFormatTypeNil(b bool)`

 SetFormatTypeNil sets the value for FormatType to be an explicit nil

### UnsetFormatType
`func (o *GuildStickerResponse) UnsetFormatType()`

UnsetFormatType ensures that no value is present for FormatType, not even an explicit nil
### GetDescription

`func (o *GuildStickerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildStickerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildStickerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildStickerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildStickerResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildStickerResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvailable

`func (o *GuildStickerResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GuildStickerResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GuildStickerResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetGuildId

`func (o *GuildStickerResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GuildStickerResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GuildStickerResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetUser

`func (o *GuildStickerResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GuildStickerResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GuildStickerResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *GuildStickerResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GuildStickerResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GuildStickerResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


