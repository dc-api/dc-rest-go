# BasicApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**CoverImage** | Pointer to **NullableString** |  | [optional] 
**PrimarySkuId** | Pointer to **string** |  | [optional] 
**Bot** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewBasicApplicationResponse

`func NewBasicApplicationResponse(id string, name string, description string, ) *BasicApplicationResponse`

NewBasicApplicationResponse instantiates a new BasicApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicApplicationResponseWithDefaults

`func NewBasicApplicationResponseWithDefaults() *BasicApplicationResponse`

NewBasicApplicationResponseWithDefaults instantiates a new BasicApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BasicApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicApplicationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *BasicApplicationResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BasicApplicationResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BasicApplicationResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *BasicApplicationResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *BasicApplicationResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *BasicApplicationResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *BasicApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BasicApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BasicApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *BasicApplicationResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicApplicationResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicApplicationResponse) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *BasicApplicationResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BasicApplicationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BasicApplicationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCoverImage

`func (o *BasicApplicationResponse) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *BasicApplicationResponse) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *BasicApplicationResponse) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *BasicApplicationResponse) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### SetCoverImageNil

`func (o *BasicApplicationResponse) SetCoverImageNil(b bool)`

 SetCoverImageNil sets the value for CoverImage to be an explicit nil

### UnsetCoverImage
`func (o *BasicApplicationResponse) UnsetCoverImage()`

UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
### GetPrimarySkuId

`func (o *BasicApplicationResponse) GetPrimarySkuId() string`

GetPrimarySkuId returns the PrimarySkuId field if non-nil, zero value otherwise.

### GetPrimarySkuIdOk

`func (o *BasicApplicationResponse) GetPrimarySkuIdOk() (*string, bool)`

GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySkuId

`func (o *BasicApplicationResponse) SetPrimarySkuId(v string)`

SetPrimarySkuId sets PrimarySkuId field to given value.

### HasPrimarySkuId

`func (o *BasicApplicationResponse) HasPrimarySkuId() bool`

HasPrimarySkuId returns a boolean if a field has been set.

### GetBot

`func (o *BasicApplicationResponse) GetBot() UserResponse`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *BasicApplicationResponse) GetBotOk() (*UserResponse, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *BasicApplicationResponse) SetBot(v UserResponse)`

SetBot sets Bot field to given value.

### HasBot

`func (o *BasicApplicationResponse) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *BasicApplicationResponse) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *BasicApplicationResponse) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


