# ApplicationCommandPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Options** | Pointer to [**[]ApplicationCommandCreateRequestOptionsInner**](ApplicationCommandCreateRequestOptionsInner.md) |  | [optional] 
**DefaultMemberPermissions** | Pointer to **int32** |  | [optional] 
**DmPermission** | Pointer to **bool** |  | [optional] 
**Contexts** | Pointer to [**[]InteractionContextType**](InteractionContextType.md) |  | [optional] 
**IntegrationTypes** | Pointer to [**[]ApplicationIntegrationType**](ApplicationIntegrationType.md) |  | [optional] 
**Handler** | Pointer to [**ApplicationCommandCreateRequestHandler**](ApplicationCommandCreateRequestHandler.md) |  | [optional] 

## Methods

### NewApplicationCommandPatchRequestPartial

`func NewApplicationCommandPatchRequestPartial() *ApplicationCommandPatchRequestPartial`

NewApplicationCommandPatchRequestPartial instantiates a new ApplicationCommandPatchRequestPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandPatchRequestPartialWithDefaults

`func NewApplicationCommandPatchRequestPartialWithDefaults() *ApplicationCommandPatchRequestPartial`

NewApplicationCommandPatchRequestPartialWithDefaults instantiates a new ApplicationCommandPatchRequestPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandPatchRequestPartial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandPatchRequestPartial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandPatchRequestPartial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationCommandPatchRequestPartial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameLocalizations

`func (o *ApplicationCommandPatchRequestPartial) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandPatchRequestPartial) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandPatchRequestPartial) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandPatchRequestPartial) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandPatchRequestPartial) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandPatchRequestPartial) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandPatchRequestPartial) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationCommandPatchRequestPartial) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionLocalizations

`func (o *ApplicationCommandPatchRequestPartial) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandPatchRequestPartial) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandPatchRequestPartial) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandPatchRequestPartial) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetOptions

`func (o *ApplicationCommandPatchRequestPartial) GetOptions() []ApplicationCommandCreateRequestOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandPatchRequestPartial) GetOptionsOk() (*[]ApplicationCommandCreateRequestOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandPatchRequestPartial) SetOptions(v []ApplicationCommandCreateRequestOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandPatchRequestPartial) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDefaultMemberPermissions

`func (o *ApplicationCommandPatchRequestPartial) GetDefaultMemberPermissions() int32`

GetDefaultMemberPermissions returns the DefaultMemberPermissions field if non-nil, zero value otherwise.

### GetDefaultMemberPermissionsOk

`func (o *ApplicationCommandPatchRequestPartial) GetDefaultMemberPermissionsOk() (*int32, bool)`

GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemberPermissions

`func (o *ApplicationCommandPatchRequestPartial) SetDefaultMemberPermissions(v int32)`

SetDefaultMemberPermissions sets DefaultMemberPermissions field to given value.

### HasDefaultMemberPermissions

`func (o *ApplicationCommandPatchRequestPartial) HasDefaultMemberPermissions() bool`

HasDefaultMemberPermissions returns a boolean if a field has been set.

### GetDmPermission

`func (o *ApplicationCommandPatchRequestPartial) GetDmPermission() bool`

GetDmPermission returns the DmPermission field if non-nil, zero value otherwise.

### GetDmPermissionOk

`func (o *ApplicationCommandPatchRequestPartial) GetDmPermissionOk() (*bool, bool)`

GetDmPermissionOk returns a tuple with the DmPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPermission

`func (o *ApplicationCommandPatchRequestPartial) SetDmPermission(v bool)`

SetDmPermission sets DmPermission field to given value.

### HasDmPermission

`func (o *ApplicationCommandPatchRequestPartial) HasDmPermission() bool`

HasDmPermission returns a boolean if a field has been set.

### GetContexts

`func (o *ApplicationCommandPatchRequestPartial) GetContexts() []InteractionContextType`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *ApplicationCommandPatchRequestPartial) GetContextsOk() (*[]InteractionContextType, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *ApplicationCommandPatchRequestPartial) SetContexts(v []InteractionContextType)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *ApplicationCommandPatchRequestPartial) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetIntegrationTypes

`func (o *ApplicationCommandPatchRequestPartial) GetIntegrationTypes() []ApplicationIntegrationType`

GetIntegrationTypes returns the IntegrationTypes field if non-nil, zero value otherwise.

### GetIntegrationTypesOk

`func (o *ApplicationCommandPatchRequestPartial) GetIntegrationTypesOk() (*[]ApplicationIntegrationType, bool)`

GetIntegrationTypesOk returns a tuple with the IntegrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypes

`func (o *ApplicationCommandPatchRequestPartial) SetIntegrationTypes(v []ApplicationIntegrationType)`

SetIntegrationTypes sets IntegrationTypes field to given value.

### HasIntegrationTypes

`func (o *ApplicationCommandPatchRequestPartial) HasIntegrationTypes() bool`

HasIntegrationTypes returns a boolean if a field has been set.

### GetHandler

`func (o *ApplicationCommandPatchRequestPartial) GetHandler() ApplicationCommandCreateRequestHandler`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *ApplicationCommandPatchRequestPartial) GetHandlerOk() (*ApplicationCommandCreateRequestHandler, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *ApplicationCommandPatchRequestPartial) SetHandler(v ApplicationCommandCreateRequestHandler)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *ApplicationCommandPatchRequestPartial) HasHandler() bool`

HasHandler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


