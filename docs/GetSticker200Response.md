# GetSticker200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Tags** | **string** |  | 
**Type** | **int32** |  | 
**FormatType** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Available** | **bool** |  | 
**GuildId** | **string** |  | 
**User** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**PackId** | **string** |  | 
**SortValue** | **int32** |  | 

## Methods

### NewGetSticker200Response

`func NewGetSticker200Response(id string, name string, tags string, type_ int32, available bool, guildId string, packId string, sortValue int32, ) *GetSticker200Response`

NewGetSticker200Response instantiates a new GetSticker200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSticker200ResponseWithDefaults

`func NewGetSticker200ResponseWithDefaults() *GetSticker200Response`

NewGetSticker200ResponseWithDefaults instantiates a new GetSticker200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSticker200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSticker200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSticker200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetSticker200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSticker200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSticker200Response) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *GetSticker200Response) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetSticker200Response) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetSticker200Response) SetTags(v string)`

SetTags sets Tags field to given value.


### GetType

`func (o *GetSticker200Response) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSticker200Response) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSticker200Response) SetType(v int32)`

SetType sets Type field to given value.


### GetFormatType

`func (o *GetSticker200Response) GetFormatType() int32`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *GetSticker200Response) GetFormatTypeOk() (*int32, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *GetSticker200Response) SetFormatType(v int32)`

SetFormatType sets FormatType field to given value.

### HasFormatType

`func (o *GetSticker200Response) HasFormatType() bool`

HasFormatType returns a boolean if a field has been set.

### GetDescription

`func (o *GetSticker200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSticker200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSticker200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSticker200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvailable

`func (o *GetSticker200Response) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetSticker200Response) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetSticker200Response) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetGuildId

`func (o *GetSticker200Response) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GetSticker200Response) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GetSticker200Response) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetUser

`func (o *GetSticker200Response) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSticker200Response) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSticker200Response) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *GetSticker200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPackId

`func (o *GetSticker200Response) GetPackId() string`

GetPackId returns the PackId field if non-nil, zero value otherwise.

### GetPackIdOk

`func (o *GetSticker200Response) GetPackIdOk() (*string, bool)`

GetPackIdOk returns a tuple with the PackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackId

`func (o *GetSticker200Response) SetPackId(v string)`

SetPackId sets PackId field to given value.


### GetSortValue

`func (o *GetSticker200Response) GetSortValue() int32`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *GetSticker200Response) GetSortValueOk() (*int32, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *GetSticker200Response) SetSortValue(v int32)`

SetSortValue sets SortValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


