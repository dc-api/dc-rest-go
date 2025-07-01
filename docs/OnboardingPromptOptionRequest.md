# OnboardingPromptOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EmojiId** | Pointer to **string** |  | [optional] 
**EmojiName** | Pointer to **NullableString** |  | [optional] 
**EmojiAnimated** | Pointer to **NullableBool** |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**ChannelIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOnboardingPromptOptionRequest

`func NewOnboardingPromptOptionRequest(title string, ) *OnboardingPromptOptionRequest`

NewOnboardingPromptOptionRequest instantiates a new OnboardingPromptOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingPromptOptionRequestWithDefaults

`func NewOnboardingPromptOptionRequestWithDefaults() *OnboardingPromptOptionRequest`

NewOnboardingPromptOptionRequestWithDefaults instantiates a new OnboardingPromptOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnboardingPromptOptionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingPromptOptionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingPromptOptionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnboardingPromptOptionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *OnboardingPromptOptionRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OnboardingPromptOptionRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OnboardingPromptOptionRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OnboardingPromptOptionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OnboardingPromptOptionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OnboardingPromptOptionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OnboardingPromptOptionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OnboardingPromptOptionRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OnboardingPromptOptionRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmojiId

`func (o *OnboardingPromptOptionRequest) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *OnboardingPromptOptionRequest) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *OnboardingPromptOptionRequest) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *OnboardingPromptOptionRequest) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *OnboardingPromptOptionRequest) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *OnboardingPromptOptionRequest) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *OnboardingPromptOptionRequest) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *OnboardingPromptOptionRequest) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *OnboardingPromptOptionRequest) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *OnboardingPromptOptionRequest) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
### GetEmojiAnimated

`func (o *OnboardingPromptOptionRequest) GetEmojiAnimated() bool`

GetEmojiAnimated returns the EmojiAnimated field if non-nil, zero value otherwise.

### GetEmojiAnimatedOk

`func (o *OnboardingPromptOptionRequest) GetEmojiAnimatedOk() (*bool, bool)`

GetEmojiAnimatedOk returns a tuple with the EmojiAnimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiAnimated

`func (o *OnboardingPromptOptionRequest) SetEmojiAnimated(v bool)`

SetEmojiAnimated sets EmojiAnimated field to given value.

### HasEmojiAnimated

`func (o *OnboardingPromptOptionRequest) HasEmojiAnimated() bool`

HasEmojiAnimated returns a boolean if a field has been set.

### SetEmojiAnimatedNil

`func (o *OnboardingPromptOptionRequest) SetEmojiAnimatedNil(b bool)`

 SetEmojiAnimatedNil sets the value for EmojiAnimated to be an explicit nil

### UnsetEmojiAnimated
`func (o *OnboardingPromptOptionRequest) UnsetEmojiAnimated()`

UnsetEmojiAnimated ensures that no value is present for EmojiAnimated, not even an explicit nil
### GetRoleIds

`func (o *OnboardingPromptOptionRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *OnboardingPromptOptionRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *OnboardingPromptOptionRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *OnboardingPromptOptionRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### SetRoleIdsNil

`func (o *OnboardingPromptOptionRequest) SetRoleIdsNil(b bool)`

 SetRoleIdsNil sets the value for RoleIds to be an explicit nil

### UnsetRoleIds
`func (o *OnboardingPromptOptionRequest) UnsetRoleIds()`

UnsetRoleIds ensures that no value is present for RoleIds, not even an explicit nil
### GetChannelIds

`func (o *OnboardingPromptOptionRequest) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *OnboardingPromptOptionRequest) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *OnboardingPromptOptionRequest) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *OnboardingPromptOptionRequest) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### SetChannelIdsNil

`func (o *OnboardingPromptOptionRequest) SetChannelIdsNil(b bool)`

 SetChannelIdsNil sets the value for ChannelIds to be an explicit nil

### UnsetChannelIds
`func (o *OnboardingPromptOptionRequest) UnsetChannelIds()`

UnsetChannelIds ensures that no value is present for ChannelIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


