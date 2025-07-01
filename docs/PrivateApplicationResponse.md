# PrivateApplicationResponse

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
**RedirectUris** | **[]string** |  | 
**InteractionsEndpointUrl** | Pointer to **NullableString** |  | [optional] 
**RoleConnectionsVerificationUrl** | Pointer to **NullableString** |  | [optional] 
**Owner** | [**UserResponse**](UserResponse.md) |  | 
**ApproximateGuildCount** | Pointer to **NullableInt32** |  | [optional] 
**ApproximateUserInstallCount** | **int32** |  | 
**ApproximateUserAuthorizationCount** | **int32** |  | 
**ExplicitContentFilter** | **int32** |  | 
**Team** | Pointer to [**NullableTeamResponse**](TeamResponse.md) |  | [optional] 

## Methods

### NewPrivateApplicationResponse

`func NewPrivateApplicationResponse(id string, name string, description string, verifyKey string, flags int32, redirectUris []*string, owner UserResponse, approximateUserInstallCount int32, approximateUserAuthorizationCount int32, explicitContentFilter int32, ) *PrivateApplicationResponse`

NewPrivateApplicationResponse instantiates a new PrivateApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateApplicationResponseWithDefaults

`func NewPrivateApplicationResponseWithDefaults() *PrivateApplicationResponse`

NewPrivateApplicationResponseWithDefaults instantiates a new PrivateApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateApplicationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *PrivateApplicationResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PrivateApplicationResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PrivateApplicationResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PrivateApplicationResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PrivateApplicationResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PrivateApplicationResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *PrivateApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *PrivateApplicationResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateApplicationResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateApplicationResponse) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateApplicationResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PrivateApplicationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PrivateApplicationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCoverImage

`func (o *PrivateApplicationResponse) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *PrivateApplicationResponse) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *PrivateApplicationResponse) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *PrivateApplicationResponse) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### SetCoverImageNil

`func (o *PrivateApplicationResponse) SetCoverImageNil(b bool)`

 SetCoverImageNil sets the value for CoverImage to be an explicit nil

### UnsetCoverImage
`func (o *PrivateApplicationResponse) UnsetCoverImage()`

UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
### GetPrimarySkuId

`func (o *PrivateApplicationResponse) GetPrimarySkuId() string`

GetPrimarySkuId returns the PrimarySkuId field if non-nil, zero value otherwise.

### GetPrimarySkuIdOk

`func (o *PrivateApplicationResponse) GetPrimarySkuIdOk() (*string, bool)`

GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySkuId

`func (o *PrivateApplicationResponse) SetPrimarySkuId(v string)`

SetPrimarySkuId sets PrimarySkuId field to given value.

### HasPrimarySkuId

`func (o *PrivateApplicationResponse) HasPrimarySkuId() bool`

HasPrimarySkuId returns a boolean if a field has been set.

### GetBot

`func (o *PrivateApplicationResponse) GetBot() UserResponse`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *PrivateApplicationResponse) GetBotOk() (*UserResponse, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *PrivateApplicationResponse) SetBot(v UserResponse)`

SetBot sets Bot field to given value.

### HasBot

`func (o *PrivateApplicationResponse) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *PrivateApplicationResponse) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *PrivateApplicationResponse) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil
### GetSlug

`func (o *PrivateApplicationResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PrivateApplicationResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PrivateApplicationResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PrivateApplicationResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *PrivateApplicationResponse) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *PrivateApplicationResponse) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetGuildId

`func (o *PrivateApplicationResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *PrivateApplicationResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *PrivateApplicationResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *PrivateApplicationResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetRpcOrigins

`func (o *PrivateApplicationResponse) GetRpcOrigins() []*string`

GetRpcOrigins returns the RpcOrigins field if non-nil, zero value otherwise.

### GetRpcOriginsOk

`func (o *PrivateApplicationResponse) GetRpcOriginsOk() (*[]*string, bool)`

GetRpcOriginsOk returns a tuple with the RpcOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcOrigins

`func (o *PrivateApplicationResponse) SetRpcOrigins(v []*string)`

SetRpcOrigins sets RpcOrigins field to given value.

### HasRpcOrigins

`func (o *PrivateApplicationResponse) HasRpcOrigins() bool`

HasRpcOrigins returns a boolean if a field has been set.

### SetRpcOriginsNil

`func (o *PrivateApplicationResponse) SetRpcOriginsNil(b bool)`

 SetRpcOriginsNil sets the value for RpcOrigins to be an explicit nil

### UnsetRpcOrigins
`func (o *PrivateApplicationResponse) UnsetRpcOrigins()`

UnsetRpcOrigins ensures that no value is present for RpcOrigins, not even an explicit nil
### GetBotPublic

`func (o *PrivateApplicationResponse) GetBotPublic() bool`

GetBotPublic returns the BotPublic field if non-nil, zero value otherwise.

### GetBotPublicOk

`func (o *PrivateApplicationResponse) GetBotPublicOk() (*bool, bool)`

GetBotPublicOk returns a tuple with the BotPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotPublic

`func (o *PrivateApplicationResponse) SetBotPublic(v bool)`

SetBotPublic sets BotPublic field to given value.

### HasBotPublic

`func (o *PrivateApplicationResponse) HasBotPublic() bool`

HasBotPublic returns a boolean if a field has been set.

### SetBotPublicNil

`func (o *PrivateApplicationResponse) SetBotPublicNil(b bool)`

 SetBotPublicNil sets the value for BotPublic to be an explicit nil

### UnsetBotPublic
`func (o *PrivateApplicationResponse) UnsetBotPublic()`

UnsetBotPublic ensures that no value is present for BotPublic, not even an explicit nil
### GetBotRequireCodeGrant

`func (o *PrivateApplicationResponse) GetBotRequireCodeGrant() bool`

GetBotRequireCodeGrant returns the BotRequireCodeGrant field if non-nil, zero value otherwise.

### GetBotRequireCodeGrantOk

`func (o *PrivateApplicationResponse) GetBotRequireCodeGrantOk() (*bool, bool)`

GetBotRequireCodeGrantOk returns a tuple with the BotRequireCodeGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotRequireCodeGrant

`func (o *PrivateApplicationResponse) SetBotRequireCodeGrant(v bool)`

SetBotRequireCodeGrant sets BotRequireCodeGrant field to given value.

### HasBotRequireCodeGrant

`func (o *PrivateApplicationResponse) HasBotRequireCodeGrant() bool`

HasBotRequireCodeGrant returns a boolean if a field has been set.

### SetBotRequireCodeGrantNil

`func (o *PrivateApplicationResponse) SetBotRequireCodeGrantNil(b bool)`

 SetBotRequireCodeGrantNil sets the value for BotRequireCodeGrant to be an explicit nil

### UnsetBotRequireCodeGrant
`func (o *PrivateApplicationResponse) UnsetBotRequireCodeGrant()`

UnsetBotRequireCodeGrant ensures that no value is present for BotRequireCodeGrant, not even an explicit nil
### GetTermsOfServiceUrl

`func (o *PrivateApplicationResponse) GetTermsOfServiceUrl() string`

GetTermsOfServiceUrl returns the TermsOfServiceUrl field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlOk

`func (o *PrivateApplicationResponse) GetTermsOfServiceUrlOk() (*string, bool)`

GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrl

`func (o *PrivateApplicationResponse) SetTermsOfServiceUrl(v string)`

SetTermsOfServiceUrl sets TermsOfServiceUrl field to given value.

### HasTermsOfServiceUrl

`func (o *PrivateApplicationResponse) HasTermsOfServiceUrl() bool`

HasTermsOfServiceUrl returns a boolean if a field has been set.

### SetTermsOfServiceUrlNil

`func (o *PrivateApplicationResponse) SetTermsOfServiceUrlNil(b bool)`

 SetTermsOfServiceUrlNil sets the value for TermsOfServiceUrl to be an explicit nil

### UnsetTermsOfServiceUrl
`func (o *PrivateApplicationResponse) UnsetTermsOfServiceUrl()`

UnsetTermsOfServiceUrl ensures that no value is present for TermsOfServiceUrl, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *PrivateApplicationResponse) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *PrivateApplicationResponse) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *PrivateApplicationResponse) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.

### HasPrivacyPolicyUrl

`func (o *PrivateApplicationResponse) HasPrivacyPolicyUrl() bool`

HasPrivacyPolicyUrl returns a boolean if a field has been set.

### SetPrivacyPolicyUrlNil

`func (o *PrivateApplicationResponse) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *PrivateApplicationResponse) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetCustomInstallUrl

`func (o *PrivateApplicationResponse) GetCustomInstallUrl() string`

GetCustomInstallUrl returns the CustomInstallUrl field if non-nil, zero value otherwise.

### GetCustomInstallUrlOk

`func (o *PrivateApplicationResponse) GetCustomInstallUrlOk() (*string, bool)`

GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInstallUrl

`func (o *PrivateApplicationResponse) SetCustomInstallUrl(v string)`

SetCustomInstallUrl sets CustomInstallUrl field to given value.

### HasCustomInstallUrl

`func (o *PrivateApplicationResponse) HasCustomInstallUrl() bool`

HasCustomInstallUrl returns a boolean if a field has been set.

### SetCustomInstallUrlNil

`func (o *PrivateApplicationResponse) SetCustomInstallUrlNil(b bool)`

 SetCustomInstallUrlNil sets the value for CustomInstallUrl to be an explicit nil

### UnsetCustomInstallUrl
`func (o *PrivateApplicationResponse) UnsetCustomInstallUrl()`

UnsetCustomInstallUrl ensures that no value is present for CustomInstallUrl, not even an explicit nil
### GetInstallParams

`func (o *PrivateApplicationResponse) GetInstallParams() ApplicationOAuth2InstallParamsResponse`

GetInstallParams returns the InstallParams field if non-nil, zero value otherwise.

### GetInstallParamsOk

`func (o *PrivateApplicationResponse) GetInstallParamsOk() (*ApplicationOAuth2InstallParamsResponse, bool)`

GetInstallParamsOk returns a tuple with the InstallParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallParams

`func (o *PrivateApplicationResponse) SetInstallParams(v ApplicationOAuth2InstallParamsResponse)`

SetInstallParams sets InstallParams field to given value.

### HasInstallParams

`func (o *PrivateApplicationResponse) HasInstallParams() bool`

HasInstallParams returns a boolean if a field has been set.

### SetInstallParamsNil

`func (o *PrivateApplicationResponse) SetInstallParamsNil(b bool)`

 SetInstallParamsNil sets the value for InstallParams to be an explicit nil

### UnsetInstallParams
`func (o *PrivateApplicationResponse) UnsetInstallParams()`

UnsetInstallParams ensures that no value is present for InstallParams, not even an explicit nil
### GetIntegrationTypesConfig

`func (o *PrivateApplicationResponse) GetIntegrationTypesConfig() map[string]ApplicationIntegrationTypeConfigurationResponse`

GetIntegrationTypesConfig returns the IntegrationTypesConfig field if non-nil, zero value otherwise.

### GetIntegrationTypesConfigOk

`func (o *PrivateApplicationResponse) GetIntegrationTypesConfigOk() (*map[string]ApplicationIntegrationTypeConfigurationResponse, bool)`

GetIntegrationTypesConfigOk returns a tuple with the IntegrationTypesConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypesConfig

`func (o *PrivateApplicationResponse) SetIntegrationTypesConfig(v map[string]ApplicationIntegrationTypeConfigurationResponse)`

SetIntegrationTypesConfig sets IntegrationTypesConfig field to given value.

### HasIntegrationTypesConfig

`func (o *PrivateApplicationResponse) HasIntegrationTypesConfig() bool`

HasIntegrationTypesConfig returns a boolean if a field has been set.

### GetVerifyKey

`func (o *PrivateApplicationResponse) GetVerifyKey() string`

GetVerifyKey returns the VerifyKey field if non-nil, zero value otherwise.

### GetVerifyKeyOk

`func (o *PrivateApplicationResponse) GetVerifyKeyOk() (*string, bool)`

GetVerifyKeyOk returns a tuple with the VerifyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyKey

`func (o *PrivateApplicationResponse) SetVerifyKey(v string)`

SetVerifyKey sets VerifyKey field to given value.


### GetFlags

`func (o *PrivateApplicationResponse) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *PrivateApplicationResponse) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *PrivateApplicationResponse) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetMaxParticipants

`func (o *PrivateApplicationResponse) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *PrivateApplicationResponse) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *PrivateApplicationResponse) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *PrivateApplicationResponse) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *PrivateApplicationResponse) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *PrivateApplicationResponse) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetTags

`func (o *PrivateApplicationResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrivateApplicationResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrivateApplicationResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrivateApplicationResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PrivateApplicationResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PrivateApplicationResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetRedirectUris

`func (o *PrivateApplicationResponse) GetRedirectUris() []*string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *PrivateApplicationResponse) GetRedirectUrisOk() (*[]*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *PrivateApplicationResponse) SetRedirectUris(v []*string)`

SetRedirectUris sets RedirectUris field to given value.


### GetInteractionsEndpointUrl

`func (o *PrivateApplicationResponse) GetInteractionsEndpointUrl() string`

GetInteractionsEndpointUrl returns the InteractionsEndpointUrl field if non-nil, zero value otherwise.

### GetInteractionsEndpointUrlOk

`func (o *PrivateApplicationResponse) GetInteractionsEndpointUrlOk() (*string, bool)`

GetInteractionsEndpointUrlOk returns a tuple with the InteractionsEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionsEndpointUrl

`func (o *PrivateApplicationResponse) SetInteractionsEndpointUrl(v string)`

SetInteractionsEndpointUrl sets InteractionsEndpointUrl field to given value.

### HasInteractionsEndpointUrl

`func (o *PrivateApplicationResponse) HasInteractionsEndpointUrl() bool`

HasInteractionsEndpointUrl returns a boolean if a field has been set.

### SetInteractionsEndpointUrlNil

`func (o *PrivateApplicationResponse) SetInteractionsEndpointUrlNil(b bool)`

 SetInteractionsEndpointUrlNil sets the value for InteractionsEndpointUrl to be an explicit nil

### UnsetInteractionsEndpointUrl
`func (o *PrivateApplicationResponse) UnsetInteractionsEndpointUrl()`

UnsetInteractionsEndpointUrl ensures that no value is present for InteractionsEndpointUrl, not even an explicit nil
### GetRoleConnectionsVerificationUrl

`func (o *PrivateApplicationResponse) GetRoleConnectionsVerificationUrl() string`

GetRoleConnectionsVerificationUrl returns the RoleConnectionsVerificationUrl field if non-nil, zero value otherwise.

### GetRoleConnectionsVerificationUrlOk

`func (o *PrivateApplicationResponse) GetRoleConnectionsVerificationUrlOk() (*string, bool)`

GetRoleConnectionsVerificationUrlOk returns a tuple with the RoleConnectionsVerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleConnectionsVerificationUrl

`func (o *PrivateApplicationResponse) SetRoleConnectionsVerificationUrl(v string)`

SetRoleConnectionsVerificationUrl sets RoleConnectionsVerificationUrl field to given value.

### HasRoleConnectionsVerificationUrl

`func (o *PrivateApplicationResponse) HasRoleConnectionsVerificationUrl() bool`

HasRoleConnectionsVerificationUrl returns a boolean if a field has been set.

### SetRoleConnectionsVerificationUrlNil

`func (o *PrivateApplicationResponse) SetRoleConnectionsVerificationUrlNil(b bool)`

 SetRoleConnectionsVerificationUrlNil sets the value for RoleConnectionsVerificationUrl to be an explicit nil

### UnsetRoleConnectionsVerificationUrl
`func (o *PrivateApplicationResponse) UnsetRoleConnectionsVerificationUrl()`

UnsetRoleConnectionsVerificationUrl ensures that no value is present for RoleConnectionsVerificationUrl, not even an explicit nil
### GetOwner

`func (o *PrivateApplicationResponse) GetOwner() UserResponse`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PrivateApplicationResponse) GetOwnerOk() (*UserResponse, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PrivateApplicationResponse) SetOwner(v UserResponse)`

SetOwner sets Owner field to given value.


### GetApproximateGuildCount

`func (o *PrivateApplicationResponse) GetApproximateGuildCount() int32`

GetApproximateGuildCount returns the ApproximateGuildCount field if non-nil, zero value otherwise.

### GetApproximateGuildCountOk

`func (o *PrivateApplicationResponse) GetApproximateGuildCountOk() (*int32, bool)`

GetApproximateGuildCountOk returns a tuple with the ApproximateGuildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateGuildCount

`func (o *PrivateApplicationResponse) SetApproximateGuildCount(v int32)`

SetApproximateGuildCount sets ApproximateGuildCount field to given value.

### HasApproximateGuildCount

`func (o *PrivateApplicationResponse) HasApproximateGuildCount() bool`

HasApproximateGuildCount returns a boolean if a field has been set.

### SetApproximateGuildCountNil

`func (o *PrivateApplicationResponse) SetApproximateGuildCountNil(b bool)`

 SetApproximateGuildCountNil sets the value for ApproximateGuildCount to be an explicit nil

### UnsetApproximateGuildCount
`func (o *PrivateApplicationResponse) UnsetApproximateGuildCount()`

UnsetApproximateGuildCount ensures that no value is present for ApproximateGuildCount, not even an explicit nil
### GetApproximateUserInstallCount

`func (o *PrivateApplicationResponse) GetApproximateUserInstallCount() int32`

GetApproximateUserInstallCount returns the ApproximateUserInstallCount field if non-nil, zero value otherwise.

### GetApproximateUserInstallCountOk

`func (o *PrivateApplicationResponse) GetApproximateUserInstallCountOk() (*int32, bool)`

GetApproximateUserInstallCountOk returns a tuple with the ApproximateUserInstallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateUserInstallCount

`func (o *PrivateApplicationResponse) SetApproximateUserInstallCount(v int32)`

SetApproximateUserInstallCount sets ApproximateUserInstallCount field to given value.


### GetApproximateUserAuthorizationCount

`func (o *PrivateApplicationResponse) GetApproximateUserAuthorizationCount() int32`

GetApproximateUserAuthorizationCount returns the ApproximateUserAuthorizationCount field if non-nil, zero value otherwise.

### GetApproximateUserAuthorizationCountOk

`func (o *PrivateApplicationResponse) GetApproximateUserAuthorizationCountOk() (*int32, bool)`

GetApproximateUserAuthorizationCountOk returns a tuple with the ApproximateUserAuthorizationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateUserAuthorizationCount

`func (o *PrivateApplicationResponse) SetApproximateUserAuthorizationCount(v int32)`

SetApproximateUserAuthorizationCount sets ApproximateUserAuthorizationCount field to given value.


### GetExplicitContentFilter

`func (o *PrivateApplicationResponse) GetExplicitContentFilter() int32`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *PrivateApplicationResponse) GetExplicitContentFilterOk() (*int32, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *PrivateApplicationResponse) SetExplicitContentFilter(v int32)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.


### GetTeam

`func (o *PrivateApplicationResponse) GetTeam() TeamResponse`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PrivateApplicationResponse) GetTeamOk() (*TeamResponse, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PrivateApplicationResponse) SetTeam(v TeamResponse)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PrivateApplicationResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *PrivateApplicationResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *PrivateApplicationResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


