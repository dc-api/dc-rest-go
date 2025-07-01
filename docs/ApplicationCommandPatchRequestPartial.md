# ApplicationCommandPatchRequestPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Options** | Pointer to [**[]ApplicationCommandCreateRequestOptionsInner**](ApplicationCommandCreateRequestOptionsInner.md) |  | [optional] 
**DefaultMemberPermissions** | Pointer to **NullableInt32** |  | [optional] 
**DmPermission** | Pointer to **NullableBool** |  | [optional] 
**Contexts** | Pointer to **[]int32** |  | [optional] 
**IntegrationTypes** | Pointer to **[]int32** |  | [optional] 
**Handler** | Pointer to **int32** |  | [optional] 

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

### SetDescriptionNil

`func (o *ApplicationCommandPatchRequestPartial) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationCommandPatchRequestPartial) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetOptionsNil

`func (o *ApplicationCommandPatchRequestPartial) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApplicationCommandPatchRequestPartial) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
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

### SetDefaultMemberPermissionsNil

`func (o *ApplicationCommandPatchRequestPartial) SetDefaultMemberPermissionsNil(b bool)`

 SetDefaultMemberPermissionsNil sets the value for DefaultMemberPermissions to be an explicit nil

### UnsetDefaultMemberPermissions
`func (o *ApplicationCommandPatchRequestPartial) UnsetDefaultMemberPermissions()`

UnsetDefaultMemberPermissions ensures that no value is present for DefaultMemberPermissions, not even an explicit nil
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

### SetDmPermissionNil

`func (o *ApplicationCommandPatchRequestPartial) SetDmPermissionNil(b bool)`

 SetDmPermissionNil sets the value for DmPermission to be an explicit nil

### UnsetDmPermission
`func (o *ApplicationCommandPatchRequestPartial) UnsetDmPermission()`

UnsetDmPermission ensures that no value is present for DmPermission, not even an explicit nil
### GetContexts

`func (o *ApplicationCommandPatchRequestPartial) GetContexts() []int32`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *ApplicationCommandPatchRequestPartial) GetContextsOk() (*[]int32, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *ApplicationCommandPatchRequestPartial) SetContexts(v []int32)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *ApplicationCommandPatchRequestPartial) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### SetContextsNil

`func (o *ApplicationCommandPatchRequestPartial) SetContextsNil(b bool)`

 SetContextsNil sets the value for Contexts to be an explicit nil

### UnsetContexts
`func (o *ApplicationCommandPatchRequestPartial) UnsetContexts()`

UnsetContexts ensures that no value is present for Contexts, not even an explicit nil
### GetIntegrationTypes

`func (o *ApplicationCommandPatchRequestPartial) GetIntegrationTypes() []int32`

GetIntegrationTypes returns the IntegrationTypes field if non-nil, zero value otherwise.

### GetIntegrationTypesOk

`func (o *ApplicationCommandPatchRequestPartial) GetIntegrationTypesOk() (*[]int32, bool)`

GetIntegrationTypesOk returns a tuple with the IntegrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypes

`func (o *ApplicationCommandPatchRequestPartial) SetIntegrationTypes(v []int32)`

SetIntegrationTypes sets IntegrationTypes field to given value.

### HasIntegrationTypes

`func (o *ApplicationCommandPatchRequestPartial) HasIntegrationTypes() bool`

HasIntegrationTypes returns a boolean if a field has been set.

### SetIntegrationTypesNil

`func (o *ApplicationCommandPatchRequestPartial) SetIntegrationTypesNil(b bool)`

 SetIntegrationTypesNil sets the value for IntegrationTypes to be an explicit nil

### UnsetIntegrationTypes
`func (o *ApplicationCommandPatchRequestPartial) UnsetIntegrationTypes()`

UnsetIntegrationTypes ensures that no value is present for IntegrationTypes, not even an explicit nil
### GetHandler

`func (o *ApplicationCommandPatchRequestPartial) GetHandler() int32`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *ApplicationCommandPatchRequestPartial) GetHandlerOk() (*int32, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *ApplicationCommandPatchRequestPartial) SetHandler(v int32)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *ApplicationCommandPatchRequestPartial) HasHandler() bool`

HasHandler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


