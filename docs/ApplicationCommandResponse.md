# ApplicationCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApplicationId** | **string** |  | 
**Version** | **string** |  | 
**DefaultMemberPermissions** | Pointer to **string** |  | [optional] 
**Type** | [**ApplicationCommandType**](ApplicationCommandType.md) |  | 
**Name** | **string** |  | 
**NameLocalized** | Pointer to **string** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalized** | Pointer to **string** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 
**DmPermission** | Pointer to **bool** |  | [optional] 
**Contexts** | Pointer to [**[]InteractionContextType**](InteractionContextType.md) |  | [optional] 
**IntegrationTypes** | Pointer to [**[]ApplicationIntegrationType**](ApplicationIntegrationType.md) |  | [optional] 
**Options** | Pointer to [**[]ApplicationCommandResponseOptionsInner**](ApplicationCommandResponseOptionsInner.md) |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationCommandResponse

`func NewApplicationCommandResponse(id string, applicationId string, version string, type_ ApplicationCommandType, name string, description string, ) *ApplicationCommandResponse`

NewApplicationCommandResponse instantiates a new ApplicationCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandResponseWithDefaults

`func NewApplicationCommandResponseWithDefaults() *ApplicationCommandResponse`

NewApplicationCommandResponseWithDefaults instantiates a new ApplicationCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCommandResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCommandResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCommandResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *ApplicationCommandResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationCommandResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationCommandResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetVersion

`func (o *ApplicationCommandResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationCommandResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationCommandResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDefaultMemberPermissions

`func (o *ApplicationCommandResponse) GetDefaultMemberPermissions() string`

GetDefaultMemberPermissions returns the DefaultMemberPermissions field if non-nil, zero value otherwise.

### GetDefaultMemberPermissionsOk

`func (o *ApplicationCommandResponse) GetDefaultMemberPermissionsOk() (*string, bool)`

GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemberPermissions

`func (o *ApplicationCommandResponse) SetDefaultMemberPermissions(v string)`

SetDefaultMemberPermissions sets DefaultMemberPermissions field to given value.

### HasDefaultMemberPermissions

`func (o *ApplicationCommandResponse) HasDefaultMemberPermissions() bool`

HasDefaultMemberPermissions returns a boolean if a field has been set.

### GetType

`func (o *ApplicationCommandResponse) GetType() ApplicationCommandType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandResponse) GetTypeOk() (*ApplicationCommandType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandResponse) SetType(v ApplicationCommandType)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### GetNameLocalizations

`func (o *ApplicationCommandResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandResponse) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandResponse) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandResponse) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandResponse) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### GetDescriptionLocalizations

`func (o *ApplicationCommandResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetGuildId

`func (o *ApplicationCommandResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ApplicationCommandResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ApplicationCommandResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *ApplicationCommandResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetDmPermission

`func (o *ApplicationCommandResponse) GetDmPermission() bool`

GetDmPermission returns the DmPermission field if non-nil, zero value otherwise.

### GetDmPermissionOk

`func (o *ApplicationCommandResponse) GetDmPermissionOk() (*bool, bool)`

GetDmPermissionOk returns a tuple with the DmPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPermission

`func (o *ApplicationCommandResponse) SetDmPermission(v bool)`

SetDmPermission sets DmPermission field to given value.

### HasDmPermission

`func (o *ApplicationCommandResponse) HasDmPermission() bool`

HasDmPermission returns a boolean if a field has been set.

### GetContexts

`func (o *ApplicationCommandResponse) GetContexts() []InteractionContextType`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *ApplicationCommandResponse) GetContextsOk() (*[]InteractionContextType, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *ApplicationCommandResponse) SetContexts(v []InteractionContextType)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *ApplicationCommandResponse) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetIntegrationTypes

`func (o *ApplicationCommandResponse) GetIntegrationTypes() []ApplicationIntegrationType`

GetIntegrationTypes returns the IntegrationTypes field if non-nil, zero value otherwise.

### GetIntegrationTypesOk

`func (o *ApplicationCommandResponse) GetIntegrationTypesOk() (*[]ApplicationIntegrationType, bool)`

GetIntegrationTypesOk returns a tuple with the IntegrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypes

`func (o *ApplicationCommandResponse) SetIntegrationTypes(v []ApplicationIntegrationType)`

SetIntegrationTypes sets IntegrationTypes field to given value.

### HasIntegrationTypes

`func (o *ApplicationCommandResponse) HasIntegrationTypes() bool`

HasIntegrationTypes returns a boolean if a field has been set.

### GetOptions

`func (o *ApplicationCommandResponse) GetOptions() []ApplicationCommandResponseOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandResponse) GetOptionsOk() (*[]ApplicationCommandResponseOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandResponse) SetOptions(v []ApplicationCommandResponseOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetNsfw

`func (o *ApplicationCommandResponse) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *ApplicationCommandResponse) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *ApplicationCommandResponse) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *ApplicationCommandResponse) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


