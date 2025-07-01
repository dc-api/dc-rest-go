# ApplicationFormPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**ApplicationFormPartialDescription**](ApplicationFormPartialDescription.md) |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**CoverImage** | Pointer to **NullableString** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**InteractionsEndpointUrl** | Pointer to **NullableString** |  | [optional] 
**ExplicitContentFilter** | Pointer to **int32** |  | [optional] 
**MaxParticipants** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomInstallUrl** | Pointer to **NullableString** |  | [optional] 
**InstallParams** | Pointer to [**NullableApplicationOAuth2InstallParams**](ApplicationOAuth2InstallParams.md) |  | [optional] 
**RoleConnectionsVerificationUrl** | Pointer to **NullableString** |  | [optional] 
**IntegrationTypesConfig** | Pointer to [**map[string]ApplicationFormPartialIntegrationTypesConfigValue**](ApplicationFormPartialIntegrationTypesConfigValue.md) |  | [optional] 

## Methods

### NewApplicationFormPartial

`func NewApplicationFormPartial() *ApplicationFormPartial`

NewApplicationFormPartial instantiates a new ApplicationFormPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFormPartialWithDefaults

`func NewApplicationFormPartialWithDefaults() *ApplicationFormPartial`

NewApplicationFormPartialWithDefaults instantiates a new ApplicationFormPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApplicationFormPartial) GetDescription() ApplicationFormPartialDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationFormPartial) GetDescriptionOk() (*ApplicationFormPartialDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationFormPartial) SetDescription(v ApplicationFormPartialDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationFormPartial) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *ApplicationFormPartial) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ApplicationFormPartial) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ApplicationFormPartial) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ApplicationFormPartial) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ApplicationFormPartial) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ApplicationFormPartial) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetCoverImage

`func (o *ApplicationFormPartial) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *ApplicationFormPartial) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *ApplicationFormPartial) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *ApplicationFormPartial) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### SetCoverImageNil

`func (o *ApplicationFormPartial) SetCoverImageNil(b bool)`

 SetCoverImageNil sets the value for CoverImage to be an explicit nil

### UnsetCoverImage
`func (o *ApplicationFormPartial) UnsetCoverImage()`

UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
### GetTeamId

`func (o *ApplicationFormPartial) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ApplicationFormPartial) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ApplicationFormPartial) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ApplicationFormPartial) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetFlags

`func (o *ApplicationFormPartial) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApplicationFormPartial) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApplicationFormPartial) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ApplicationFormPartial) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *ApplicationFormPartial) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *ApplicationFormPartial) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetInteractionsEndpointUrl

`func (o *ApplicationFormPartial) GetInteractionsEndpointUrl() string`

GetInteractionsEndpointUrl returns the InteractionsEndpointUrl field if non-nil, zero value otherwise.

### GetInteractionsEndpointUrlOk

`func (o *ApplicationFormPartial) GetInteractionsEndpointUrlOk() (*string, bool)`

GetInteractionsEndpointUrlOk returns a tuple with the InteractionsEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionsEndpointUrl

`func (o *ApplicationFormPartial) SetInteractionsEndpointUrl(v string)`

SetInteractionsEndpointUrl sets InteractionsEndpointUrl field to given value.

### HasInteractionsEndpointUrl

`func (o *ApplicationFormPartial) HasInteractionsEndpointUrl() bool`

HasInteractionsEndpointUrl returns a boolean if a field has been set.

### SetInteractionsEndpointUrlNil

`func (o *ApplicationFormPartial) SetInteractionsEndpointUrlNil(b bool)`

 SetInteractionsEndpointUrlNil sets the value for InteractionsEndpointUrl to be an explicit nil

### UnsetInteractionsEndpointUrl
`func (o *ApplicationFormPartial) UnsetInteractionsEndpointUrl()`

UnsetInteractionsEndpointUrl ensures that no value is present for InteractionsEndpointUrl, not even an explicit nil
### GetExplicitContentFilter

`func (o *ApplicationFormPartial) GetExplicitContentFilter() int32`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *ApplicationFormPartial) GetExplicitContentFilterOk() (*int32, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *ApplicationFormPartial) SetExplicitContentFilter(v int32)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.

### HasExplicitContentFilter

`func (o *ApplicationFormPartial) HasExplicitContentFilter() bool`

HasExplicitContentFilter returns a boolean if a field has been set.

### GetMaxParticipants

`func (o *ApplicationFormPartial) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *ApplicationFormPartial) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *ApplicationFormPartial) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *ApplicationFormPartial) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *ApplicationFormPartial) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *ApplicationFormPartial) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetType

`func (o *ApplicationFormPartial) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationFormPartial) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationFormPartial) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationFormPartial) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApplicationFormPartial) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApplicationFormPartial) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTags

`func (o *ApplicationFormPartial) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationFormPartial) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationFormPartial) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationFormPartial) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationFormPartial) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationFormPartial) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCustomInstallUrl

`func (o *ApplicationFormPartial) GetCustomInstallUrl() string`

GetCustomInstallUrl returns the CustomInstallUrl field if non-nil, zero value otherwise.

### GetCustomInstallUrlOk

`func (o *ApplicationFormPartial) GetCustomInstallUrlOk() (*string, bool)`

GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInstallUrl

`func (o *ApplicationFormPartial) SetCustomInstallUrl(v string)`

SetCustomInstallUrl sets CustomInstallUrl field to given value.

### HasCustomInstallUrl

`func (o *ApplicationFormPartial) HasCustomInstallUrl() bool`

HasCustomInstallUrl returns a boolean if a field has been set.

### SetCustomInstallUrlNil

`func (o *ApplicationFormPartial) SetCustomInstallUrlNil(b bool)`

 SetCustomInstallUrlNil sets the value for CustomInstallUrl to be an explicit nil

### UnsetCustomInstallUrl
`func (o *ApplicationFormPartial) UnsetCustomInstallUrl()`

UnsetCustomInstallUrl ensures that no value is present for CustomInstallUrl, not even an explicit nil
### GetInstallParams

`func (o *ApplicationFormPartial) GetInstallParams() ApplicationOAuth2InstallParams`

GetInstallParams returns the InstallParams field if non-nil, zero value otherwise.

### GetInstallParamsOk

`func (o *ApplicationFormPartial) GetInstallParamsOk() (*ApplicationOAuth2InstallParams, bool)`

GetInstallParamsOk returns a tuple with the InstallParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallParams

`func (o *ApplicationFormPartial) SetInstallParams(v ApplicationOAuth2InstallParams)`

SetInstallParams sets InstallParams field to given value.

### HasInstallParams

`func (o *ApplicationFormPartial) HasInstallParams() bool`

HasInstallParams returns a boolean if a field has been set.

### SetInstallParamsNil

`func (o *ApplicationFormPartial) SetInstallParamsNil(b bool)`

 SetInstallParamsNil sets the value for InstallParams to be an explicit nil

### UnsetInstallParams
`func (o *ApplicationFormPartial) UnsetInstallParams()`

UnsetInstallParams ensures that no value is present for InstallParams, not even an explicit nil
### GetRoleConnectionsVerificationUrl

`func (o *ApplicationFormPartial) GetRoleConnectionsVerificationUrl() string`

GetRoleConnectionsVerificationUrl returns the RoleConnectionsVerificationUrl field if non-nil, zero value otherwise.

### GetRoleConnectionsVerificationUrlOk

`func (o *ApplicationFormPartial) GetRoleConnectionsVerificationUrlOk() (*string, bool)`

GetRoleConnectionsVerificationUrlOk returns a tuple with the RoleConnectionsVerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleConnectionsVerificationUrl

`func (o *ApplicationFormPartial) SetRoleConnectionsVerificationUrl(v string)`

SetRoleConnectionsVerificationUrl sets RoleConnectionsVerificationUrl field to given value.

### HasRoleConnectionsVerificationUrl

`func (o *ApplicationFormPartial) HasRoleConnectionsVerificationUrl() bool`

HasRoleConnectionsVerificationUrl returns a boolean if a field has been set.

### SetRoleConnectionsVerificationUrlNil

`func (o *ApplicationFormPartial) SetRoleConnectionsVerificationUrlNil(b bool)`

 SetRoleConnectionsVerificationUrlNil sets the value for RoleConnectionsVerificationUrl to be an explicit nil

### UnsetRoleConnectionsVerificationUrl
`func (o *ApplicationFormPartial) UnsetRoleConnectionsVerificationUrl()`

UnsetRoleConnectionsVerificationUrl ensures that no value is present for RoleConnectionsVerificationUrl, not even an explicit nil
### GetIntegrationTypesConfig

`func (o *ApplicationFormPartial) GetIntegrationTypesConfig() map[string]ApplicationFormPartialIntegrationTypesConfigValue`

GetIntegrationTypesConfig returns the IntegrationTypesConfig field if non-nil, zero value otherwise.

### GetIntegrationTypesConfigOk

`func (o *ApplicationFormPartial) GetIntegrationTypesConfigOk() (*map[string]ApplicationFormPartialIntegrationTypesConfigValue, bool)`

GetIntegrationTypesConfigOk returns a tuple with the IntegrationTypesConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypesConfig

`func (o *ApplicationFormPartial) SetIntegrationTypesConfig(v map[string]ApplicationFormPartialIntegrationTypesConfigValue)`

SetIntegrationTypesConfig sets IntegrationTypesConfig field to given value.

### HasIntegrationTypesConfig

`func (o *ApplicationFormPartial) HasIntegrationTypesConfig() bool`

HasIntegrationTypesConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


