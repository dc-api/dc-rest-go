# ApplicationResponse

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
**Slug** | Pointer to **NullableString** |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 
**RpcOrigins** | Pointer to **[]string** |  | [optional] 
**BotPublic** | Pointer to **NullableBool** |  | [optional] 
**BotRequireCodeGrant** | Pointer to **NullableBool** |  | [optional] 
**TermsOfServiceUrl** | Pointer to **NullableString** |  | [optional] 
**PrivacyPolicyUrl** | Pointer to **NullableString** |  | [optional] 
**CustomInstallUrl** | Pointer to **NullableString** |  | [optional] 
**InstallParams** | Pointer to [**NullableApplicationOAuth2InstallParamsResponse**](ApplicationOAuth2InstallParamsResponse.md) |  | [optional] 
**IntegrationTypesConfig** | Pointer to [**map[string]ApplicationIntegrationTypeConfigurationResponse**](ApplicationIntegrationTypeConfigurationResponse.md) |  | [optional] 
**VerifyKey** | **string** |  | 
**Flags** | **int32** |  | 
**MaxParticipants** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApplicationResponse

`func NewApplicationResponse(id string, name string, description string, verifyKey string, flags int32, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *ApplicationResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ApplicationResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ApplicationResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ApplicationResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ApplicationResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ApplicationResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *ApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ApplicationResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResponse) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApplicationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApplicationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCoverImage

`func (o *ApplicationResponse) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *ApplicationResponse) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *ApplicationResponse) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *ApplicationResponse) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### SetCoverImageNil

`func (o *ApplicationResponse) SetCoverImageNil(b bool)`

 SetCoverImageNil sets the value for CoverImage to be an explicit nil

### UnsetCoverImage
`func (o *ApplicationResponse) UnsetCoverImage()`

UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
### GetPrimarySkuId

`func (o *ApplicationResponse) GetPrimarySkuId() string`

GetPrimarySkuId returns the PrimarySkuId field if non-nil, zero value otherwise.

### GetPrimarySkuIdOk

`func (o *ApplicationResponse) GetPrimarySkuIdOk() (*string, bool)`

GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySkuId

`func (o *ApplicationResponse) SetPrimarySkuId(v string)`

SetPrimarySkuId sets PrimarySkuId field to given value.

### HasPrimarySkuId

`func (o *ApplicationResponse) HasPrimarySkuId() bool`

HasPrimarySkuId returns a boolean if a field has been set.

### GetBot

`func (o *ApplicationResponse) GetBot() UserResponse`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *ApplicationResponse) GetBotOk() (*UserResponse, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *ApplicationResponse) SetBot(v UserResponse)`

SetBot sets Bot field to given value.

### HasBot

`func (o *ApplicationResponse) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *ApplicationResponse) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *ApplicationResponse) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil
### GetSlug

`func (o *ApplicationResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ApplicationResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ApplicationResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ApplicationResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *ApplicationResponse) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *ApplicationResponse) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetGuildId

`func (o *ApplicationResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ApplicationResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ApplicationResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *ApplicationResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetRpcOrigins

`func (o *ApplicationResponse) GetRpcOrigins() []*string`

GetRpcOrigins returns the RpcOrigins field if non-nil, zero value otherwise.

### GetRpcOriginsOk

`func (o *ApplicationResponse) GetRpcOriginsOk() (*[]*string, bool)`

GetRpcOriginsOk returns a tuple with the RpcOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcOrigins

`func (o *ApplicationResponse) SetRpcOrigins(v []*string)`

SetRpcOrigins sets RpcOrigins field to given value.

### HasRpcOrigins

`func (o *ApplicationResponse) HasRpcOrigins() bool`

HasRpcOrigins returns a boolean if a field has been set.

### SetRpcOriginsNil

`func (o *ApplicationResponse) SetRpcOriginsNil(b bool)`

 SetRpcOriginsNil sets the value for RpcOrigins to be an explicit nil

### UnsetRpcOrigins
`func (o *ApplicationResponse) UnsetRpcOrigins()`

UnsetRpcOrigins ensures that no value is present for RpcOrigins, not even an explicit nil
### GetBotPublic

`func (o *ApplicationResponse) GetBotPublic() bool`

GetBotPublic returns the BotPublic field if non-nil, zero value otherwise.

### GetBotPublicOk

`func (o *ApplicationResponse) GetBotPublicOk() (*bool, bool)`

GetBotPublicOk returns a tuple with the BotPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotPublic

`func (o *ApplicationResponse) SetBotPublic(v bool)`

SetBotPublic sets BotPublic field to given value.

### HasBotPublic

`func (o *ApplicationResponse) HasBotPublic() bool`

HasBotPublic returns a boolean if a field has been set.

### SetBotPublicNil

`func (o *ApplicationResponse) SetBotPublicNil(b bool)`

 SetBotPublicNil sets the value for BotPublic to be an explicit nil

### UnsetBotPublic
`func (o *ApplicationResponse) UnsetBotPublic()`

UnsetBotPublic ensures that no value is present for BotPublic, not even an explicit nil
### GetBotRequireCodeGrant

`func (o *ApplicationResponse) GetBotRequireCodeGrant() bool`

GetBotRequireCodeGrant returns the BotRequireCodeGrant field if non-nil, zero value otherwise.

### GetBotRequireCodeGrantOk

`func (o *ApplicationResponse) GetBotRequireCodeGrantOk() (*bool, bool)`

GetBotRequireCodeGrantOk returns a tuple with the BotRequireCodeGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotRequireCodeGrant

`func (o *ApplicationResponse) SetBotRequireCodeGrant(v bool)`

SetBotRequireCodeGrant sets BotRequireCodeGrant field to given value.

### HasBotRequireCodeGrant

`func (o *ApplicationResponse) HasBotRequireCodeGrant() bool`

HasBotRequireCodeGrant returns a boolean if a field has been set.

### SetBotRequireCodeGrantNil

`func (o *ApplicationResponse) SetBotRequireCodeGrantNil(b bool)`

 SetBotRequireCodeGrantNil sets the value for BotRequireCodeGrant to be an explicit nil

### UnsetBotRequireCodeGrant
`func (o *ApplicationResponse) UnsetBotRequireCodeGrant()`

UnsetBotRequireCodeGrant ensures that no value is present for BotRequireCodeGrant, not even an explicit nil
### GetTermsOfServiceUrl

`func (o *ApplicationResponse) GetTermsOfServiceUrl() string`

GetTermsOfServiceUrl returns the TermsOfServiceUrl field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlOk

`func (o *ApplicationResponse) GetTermsOfServiceUrlOk() (*string, bool)`

GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrl

`func (o *ApplicationResponse) SetTermsOfServiceUrl(v string)`

SetTermsOfServiceUrl sets TermsOfServiceUrl field to given value.

### HasTermsOfServiceUrl

`func (o *ApplicationResponse) HasTermsOfServiceUrl() bool`

HasTermsOfServiceUrl returns a boolean if a field has been set.

### SetTermsOfServiceUrlNil

`func (o *ApplicationResponse) SetTermsOfServiceUrlNil(b bool)`

 SetTermsOfServiceUrlNil sets the value for TermsOfServiceUrl to be an explicit nil

### UnsetTermsOfServiceUrl
`func (o *ApplicationResponse) UnsetTermsOfServiceUrl()`

UnsetTermsOfServiceUrl ensures that no value is present for TermsOfServiceUrl, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *ApplicationResponse) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *ApplicationResponse) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *ApplicationResponse) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.

### HasPrivacyPolicyUrl

`func (o *ApplicationResponse) HasPrivacyPolicyUrl() bool`

HasPrivacyPolicyUrl returns a boolean if a field has been set.

### SetPrivacyPolicyUrlNil

`func (o *ApplicationResponse) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *ApplicationResponse) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetCustomInstallUrl

`func (o *ApplicationResponse) GetCustomInstallUrl() string`

GetCustomInstallUrl returns the CustomInstallUrl field if non-nil, zero value otherwise.

### GetCustomInstallUrlOk

`func (o *ApplicationResponse) GetCustomInstallUrlOk() (*string, bool)`

GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInstallUrl

`func (o *ApplicationResponse) SetCustomInstallUrl(v string)`

SetCustomInstallUrl sets CustomInstallUrl field to given value.

### HasCustomInstallUrl

`func (o *ApplicationResponse) HasCustomInstallUrl() bool`

HasCustomInstallUrl returns a boolean if a field has been set.

### SetCustomInstallUrlNil

`func (o *ApplicationResponse) SetCustomInstallUrlNil(b bool)`

 SetCustomInstallUrlNil sets the value for CustomInstallUrl to be an explicit nil

### UnsetCustomInstallUrl
`func (o *ApplicationResponse) UnsetCustomInstallUrl()`

UnsetCustomInstallUrl ensures that no value is present for CustomInstallUrl, not even an explicit nil
### GetInstallParams

`func (o *ApplicationResponse) GetInstallParams() ApplicationOAuth2InstallParamsResponse`

GetInstallParams returns the InstallParams field if non-nil, zero value otherwise.

### GetInstallParamsOk

`func (o *ApplicationResponse) GetInstallParamsOk() (*ApplicationOAuth2InstallParamsResponse, bool)`

GetInstallParamsOk returns a tuple with the InstallParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallParams

`func (o *ApplicationResponse) SetInstallParams(v ApplicationOAuth2InstallParamsResponse)`

SetInstallParams sets InstallParams field to given value.

### HasInstallParams

`func (o *ApplicationResponse) HasInstallParams() bool`

HasInstallParams returns a boolean if a field has been set.

### SetInstallParamsNil

`func (o *ApplicationResponse) SetInstallParamsNil(b bool)`

 SetInstallParamsNil sets the value for InstallParams to be an explicit nil

### UnsetInstallParams
`func (o *ApplicationResponse) UnsetInstallParams()`

UnsetInstallParams ensures that no value is present for InstallParams, not even an explicit nil
### GetIntegrationTypesConfig

`func (o *ApplicationResponse) GetIntegrationTypesConfig() map[string]ApplicationIntegrationTypeConfigurationResponse`

GetIntegrationTypesConfig returns the IntegrationTypesConfig field if non-nil, zero value otherwise.

### GetIntegrationTypesConfigOk

`func (o *ApplicationResponse) GetIntegrationTypesConfigOk() (*map[string]ApplicationIntegrationTypeConfigurationResponse, bool)`

GetIntegrationTypesConfigOk returns a tuple with the IntegrationTypesConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypesConfig

`func (o *ApplicationResponse) SetIntegrationTypesConfig(v map[string]ApplicationIntegrationTypeConfigurationResponse)`

SetIntegrationTypesConfig sets IntegrationTypesConfig field to given value.

### HasIntegrationTypesConfig

`func (o *ApplicationResponse) HasIntegrationTypesConfig() bool`

HasIntegrationTypesConfig returns a boolean if a field has been set.

### GetVerifyKey

`func (o *ApplicationResponse) GetVerifyKey() string`

GetVerifyKey returns the VerifyKey field if non-nil, zero value otherwise.

### GetVerifyKeyOk

`func (o *ApplicationResponse) GetVerifyKeyOk() (*string, bool)`

GetVerifyKeyOk returns a tuple with the VerifyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyKey

`func (o *ApplicationResponse) SetVerifyKey(v string)`

SetVerifyKey sets VerifyKey field to given value.


### GetFlags

`func (o *ApplicationResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApplicationResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApplicationResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetMaxParticipants

`func (o *ApplicationResponse) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *ApplicationResponse) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *ApplicationResponse) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *ApplicationResponse) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *ApplicationResponse) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *ApplicationResponse) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetTags

`func (o *ApplicationResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


