# InviteGuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Splash** | Pointer to **NullableString** |  | [optional] 
**Banner** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Features** | **[]string** |  | 
**VerificationLevel** | Pointer to **int32** |  | [optional] 
**VanityUrlCode** | Pointer to **NullableString** |  | [optional] 
**NsfwLevel** | Pointer to **NullableInt32** |  | [optional] 
**Nsfw** | Pointer to **NullableBool** |  | [optional] 
**PremiumSubscriptionCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewInviteGuildResponse

`func NewInviteGuildResponse(id string, name string, features []string, ) *InviteGuildResponse`

NewInviteGuildResponse instantiates a new InviteGuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteGuildResponseWithDefaults

`func NewInviteGuildResponseWithDefaults() *InviteGuildResponse`

NewInviteGuildResponseWithDefaults instantiates a new InviteGuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InviteGuildResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InviteGuildResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InviteGuildResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InviteGuildResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InviteGuildResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InviteGuildResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSplash

`func (o *InviteGuildResponse) GetSplash() string`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *InviteGuildResponse) GetSplashOk() (*string, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *InviteGuildResponse) SetSplash(v string)`

SetSplash sets Splash field to given value.

### HasSplash

`func (o *InviteGuildResponse) HasSplash() bool`

HasSplash returns a boolean if a field has been set.

### SetSplashNil

`func (o *InviteGuildResponse) SetSplashNil(b bool)`

 SetSplashNil sets the value for Splash to be an explicit nil

### UnsetSplash
`func (o *InviteGuildResponse) UnsetSplash()`

UnsetSplash ensures that no value is present for Splash, not even an explicit nil
### GetBanner

`func (o *InviteGuildResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *InviteGuildResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *InviteGuildResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *InviteGuildResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *InviteGuildResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *InviteGuildResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetDescription

`func (o *InviteGuildResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InviteGuildResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InviteGuildResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InviteGuildResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InviteGuildResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InviteGuildResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *InviteGuildResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *InviteGuildResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *InviteGuildResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *InviteGuildResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *InviteGuildResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *InviteGuildResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetFeatures

`func (o *InviteGuildResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InviteGuildResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InviteGuildResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetVerificationLevel

`func (o *InviteGuildResponse) GetVerificationLevel() int32`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *InviteGuildResponse) GetVerificationLevelOk() (*int32, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *InviteGuildResponse) SetVerificationLevel(v int32)`

SetVerificationLevel sets VerificationLevel field to given value.

### HasVerificationLevel

`func (o *InviteGuildResponse) HasVerificationLevel() bool`

HasVerificationLevel returns a boolean if a field has been set.

### GetVanityUrlCode

`func (o *InviteGuildResponse) GetVanityUrlCode() string`

GetVanityUrlCode returns the VanityUrlCode field if non-nil, zero value otherwise.

### GetVanityUrlCodeOk

`func (o *InviteGuildResponse) GetVanityUrlCodeOk() (*string, bool)`

GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVanityUrlCode

`func (o *InviteGuildResponse) SetVanityUrlCode(v string)`

SetVanityUrlCode sets VanityUrlCode field to given value.

### HasVanityUrlCode

`func (o *InviteGuildResponse) HasVanityUrlCode() bool`

HasVanityUrlCode returns a boolean if a field has been set.

### SetVanityUrlCodeNil

`func (o *InviteGuildResponse) SetVanityUrlCodeNil(b bool)`

 SetVanityUrlCodeNil sets the value for VanityUrlCode to be an explicit nil

### UnsetVanityUrlCode
`func (o *InviteGuildResponse) UnsetVanityUrlCode()`

UnsetVanityUrlCode ensures that no value is present for VanityUrlCode, not even an explicit nil
### GetNsfwLevel

`func (o *InviteGuildResponse) GetNsfwLevel() int32`

GetNsfwLevel returns the NsfwLevel field if non-nil, zero value otherwise.

### GetNsfwLevelOk

`func (o *InviteGuildResponse) GetNsfwLevelOk() (*int32, bool)`

GetNsfwLevelOk returns a tuple with the NsfwLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwLevel

`func (o *InviteGuildResponse) SetNsfwLevel(v int32)`

SetNsfwLevel sets NsfwLevel field to given value.

### HasNsfwLevel

`func (o *InviteGuildResponse) HasNsfwLevel() bool`

HasNsfwLevel returns a boolean if a field has been set.

### SetNsfwLevelNil

`func (o *InviteGuildResponse) SetNsfwLevelNil(b bool)`

 SetNsfwLevelNil sets the value for NsfwLevel to be an explicit nil

### UnsetNsfwLevel
`func (o *InviteGuildResponse) UnsetNsfwLevel()`

UnsetNsfwLevel ensures that no value is present for NsfwLevel, not even an explicit nil
### GetNsfw

`func (o *InviteGuildResponse) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *InviteGuildResponse) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *InviteGuildResponse) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *InviteGuildResponse) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### SetNsfwNil

`func (o *InviteGuildResponse) SetNsfwNil(b bool)`

 SetNsfwNil sets the value for Nsfw to be an explicit nil

### UnsetNsfw
`func (o *InviteGuildResponse) UnsetNsfw()`

UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
### GetPremiumSubscriptionCount

`func (o *InviteGuildResponse) GetPremiumSubscriptionCount() int32`

GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field if non-nil, zero value otherwise.

### GetPremiumSubscriptionCountOk

`func (o *InviteGuildResponse) GetPremiumSubscriptionCountOk() (*int32, bool)`

GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionCount

`func (o *InviteGuildResponse) SetPremiumSubscriptionCount(v int32)`

SetPremiumSubscriptionCount sets PremiumSubscriptionCount field to given value.

### HasPremiumSubscriptionCount

`func (o *InviteGuildResponse) HasPremiumSubscriptionCount() bool`

HasPremiumSubscriptionCount returns a boolean if a field has been set.

### SetPremiumSubscriptionCountNil

`func (o *InviteGuildResponse) SetPremiumSubscriptionCountNil(b bool)`

 SetPremiumSubscriptionCountNil sets the value for PremiumSubscriptionCount to be an explicit nil

### UnsetPremiumSubscriptionCount
`func (o *InviteGuildResponse) UnsetPremiumSubscriptionCount()`

UnsetPremiumSubscriptionCount ensures that no value is present for PremiumSubscriptionCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


