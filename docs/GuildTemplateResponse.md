# GuildTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**UsageCount** | **int32** |  | 
**CreatorId** | **string** |  | 
**Creator** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**SourceGuildId** | **string** |  | 
**SerializedSourceGuild** | [**GuildTemplateSnapshotResponse**](GuildTemplateSnapshotResponse.md) |  | 
**IsDirty** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewGuildTemplateResponse

`func NewGuildTemplateResponse(code string, name string, usageCount int32, creatorId string, createdAt time.Time, updatedAt time.Time, sourceGuildId string, serializedSourceGuild GuildTemplateSnapshotResponse, ) *GuildTemplateResponse`

NewGuildTemplateResponse instantiates a new GuildTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateResponseWithDefaults

`func NewGuildTemplateResponseWithDefaults() *GuildTemplateResponse`

NewGuildTemplateResponseWithDefaults instantiates a new GuildTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GuildTemplateResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GuildTemplateResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GuildTemplateResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *GuildTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GuildTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildTemplateResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildTemplateResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUsageCount

`func (o *GuildTemplateResponse) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *GuildTemplateResponse) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *GuildTemplateResponse) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.


### GetCreatorId

`func (o *GuildTemplateResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *GuildTemplateResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *GuildTemplateResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetCreator

`func (o *GuildTemplateResponse) GetCreator() UserResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GuildTemplateResponse) GetCreatorOk() (*UserResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GuildTemplateResponse) SetCreator(v UserResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GuildTemplateResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *GuildTemplateResponse) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *GuildTemplateResponse) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetCreatedAt

`func (o *GuildTemplateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GuildTemplateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GuildTemplateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GuildTemplateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GuildTemplateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GuildTemplateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSourceGuildId

`func (o *GuildTemplateResponse) GetSourceGuildId() string`

GetSourceGuildId returns the SourceGuildId field if non-nil, zero value otherwise.

### GetSourceGuildIdOk

`func (o *GuildTemplateResponse) GetSourceGuildIdOk() (*string, bool)`

GetSourceGuildIdOk returns a tuple with the SourceGuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuildId

`func (o *GuildTemplateResponse) SetSourceGuildId(v string)`

SetSourceGuildId sets SourceGuildId field to given value.


### GetSerializedSourceGuild

`func (o *GuildTemplateResponse) GetSerializedSourceGuild() GuildTemplateSnapshotResponse`

GetSerializedSourceGuild returns the SerializedSourceGuild field if non-nil, zero value otherwise.

### GetSerializedSourceGuildOk

`func (o *GuildTemplateResponse) GetSerializedSourceGuildOk() (*GuildTemplateSnapshotResponse, bool)`

GetSerializedSourceGuildOk returns a tuple with the SerializedSourceGuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSourceGuild

`func (o *GuildTemplateResponse) SetSerializedSourceGuild(v GuildTemplateSnapshotResponse)`

SetSerializedSourceGuild sets SerializedSourceGuild field to given value.


### GetIsDirty

`func (o *GuildTemplateResponse) GetIsDirty() bool`

GetIsDirty returns the IsDirty field if non-nil, zero value otherwise.

### GetIsDirtyOk

`func (o *GuildTemplateResponse) GetIsDirtyOk() (*bool, bool)`

GetIsDirtyOk returns a tuple with the IsDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirty

`func (o *GuildTemplateResponse) SetIsDirty(v bool)`

SetIsDirty sets IsDirty field to given value.

### HasIsDirty

`func (o *GuildTemplateResponse) HasIsDirty() bool`

HasIsDirty returns a boolean if a field has been set.

### SetIsDirtyNil

`func (o *GuildTemplateResponse) SetIsDirtyNil(b bool)`

 SetIsDirtyNil sets the value for IsDirty to be an explicit nil

### UnsetIsDirty
`func (o *GuildTemplateResponse) UnsetIsDirty()`

UnsetIsDirty ensures that no value is present for IsDirty, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


