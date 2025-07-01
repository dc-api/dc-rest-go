# GuildCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Region** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**VerificationLevel** | Pointer to **int32** |  | [optional] 
**DefaultMessageNotifications** | Pointer to **int32** |  | [optional] 
**ExplicitContentFilter** | Pointer to **int32** |  | [optional] 
**PreferredLocale** | Pointer to **string** |  | [optional] 
**AfkTimeout** | Pointer to **NullableInt32** |  | [optional] 
**Roles** | Pointer to [**[]CreateGuildRequestRoleItem**](CreateGuildRequestRoleItem.md) |  | [optional] 
**Channels** | Pointer to [**[]CreateGuildRequestChannelItem**](CreateGuildRequestChannelItem.md) |  | [optional] 
**AfkChannelId** | Pointer to **string** |  | [optional] 
**SystemChannelId** | Pointer to **string** |  | [optional] 
**SystemChannelFlags** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGuildCreateRequest

`func NewGuildCreateRequest(name string, ) *GuildCreateRequest`

NewGuildCreateRequest instantiates a new GuildCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildCreateRequestWithDefaults

`func NewGuildCreateRequestWithDefaults() *GuildCreateRequest`

NewGuildCreateRequestWithDefaults instantiates a new GuildCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GuildCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuildCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuildCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *GuildCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *GuildCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GuildCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GuildCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GuildCreateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *GuildCreateRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *GuildCreateRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetIcon

`func (o *GuildCreateRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GuildCreateRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GuildCreateRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GuildCreateRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *GuildCreateRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *GuildCreateRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetVerificationLevel

`func (o *GuildCreateRequest) GetVerificationLevel() int32`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildCreateRequest) GetVerificationLevelOk() (*int32, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildCreateRequest) SetVerificationLevel(v int32)`

SetVerificationLevel sets VerificationLevel field to given value.

### HasVerificationLevel

`func (o *GuildCreateRequest) HasVerificationLevel() bool`

HasVerificationLevel returns a boolean if a field has been set.

### GetDefaultMessageNotifications

`func (o *GuildCreateRequest) GetDefaultMessageNotifications() int32`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildCreateRequest) GetDefaultMessageNotificationsOk() (*int32, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildCreateRequest) SetDefaultMessageNotifications(v int32)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.

### HasDefaultMessageNotifications

`func (o *GuildCreateRequest) HasDefaultMessageNotifications() bool`

HasDefaultMessageNotifications returns a boolean if a field has been set.

### GetExplicitContentFilter

`func (o *GuildCreateRequest) GetExplicitContentFilter() int32`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildCreateRequest) GetExplicitContentFilterOk() (*int32, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildCreateRequest) SetExplicitContentFilter(v int32)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.

### HasExplicitContentFilter

`func (o *GuildCreateRequest) HasExplicitContentFilter() bool`

HasExplicitContentFilter returns a boolean if a field has been set.

### GetPreferredLocale

`func (o *GuildCreateRequest) GetPreferredLocale() string`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildCreateRequest) GetPreferredLocaleOk() (*string, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildCreateRequest) SetPreferredLocale(v string)`

SetPreferredLocale sets PreferredLocale field to given value.

### HasPreferredLocale

`func (o *GuildCreateRequest) HasPreferredLocale() bool`

HasPreferredLocale returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildCreateRequest) GetAfkTimeout() int32`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildCreateRequest) GetAfkTimeoutOk() (*int32, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildCreateRequest) SetAfkTimeout(v int32)`

SetAfkTimeout sets AfkTimeout field to given value.

### HasAfkTimeout

`func (o *GuildCreateRequest) HasAfkTimeout() bool`

HasAfkTimeout returns a boolean if a field has been set.

### SetAfkTimeoutNil

`func (o *GuildCreateRequest) SetAfkTimeoutNil(b bool)`

 SetAfkTimeoutNil sets the value for AfkTimeout to be an explicit nil

### UnsetAfkTimeout
`func (o *GuildCreateRequest) UnsetAfkTimeout()`

UnsetAfkTimeout ensures that no value is present for AfkTimeout, not even an explicit nil
### GetRoles

`func (o *GuildCreateRequest) GetRoles() []CreateGuildRequestRoleItem`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GuildCreateRequest) GetRolesOk() (*[]CreateGuildRequestRoleItem, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GuildCreateRequest) SetRoles(v []CreateGuildRequestRoleItem)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GuildCreateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *GuildCreateRequest) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *GuildCreateRequest) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetChannels

`func (o *GuildCreateRequest) GetChannels() []CreateGuildRequestChannelItem`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *GuildCreateRequest) GetChannelsOk() (*[]CreateGuildRequestChannelItem, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *GuildCreateRequest) SetChannels(v []CreateGuildRequestChannelItem)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *GuildCreateRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *GuildCreateRequest) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *GuildCreateRequest) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetAfkChannelId

`func (o *GuildCreateRequest) GetAfkChannelId() string`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildCreateRequest) GetAfkChannelIdOk() (*string, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildCreateRequest) SetAfkChannelId(v string)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildCreateRequest) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetSystemChannelId

`func (o *GuildCreateRequest) GetSystemChannelId() string`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildCreateRequest) GetSystemChannelIdOk() (*string, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildCreateRequest) SetSystemChannelId(v string)`

SetSystemChannelId sets SystemChannelId field to given value.

### HasSystemChannelId

`func (o *GuildCreateRequest) HasSystemChannelId() bool`

HasSystemChannelId returns a boolean if a field has been set.

### GetSystemChannelFlags

`func (o *GuildCreateRequest) GetSystemChannelFlags() int32`

GetSystemChannelFlags returns the SystemChannelFlags field if non-nil, zero value otherwise.

### GetSystemChannelFlagsOk

`func (o *GuildCreateRequest) GetSystemChannelFlagsOk() (*int32, bool)`

GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelFlags

`func (o *GuildCreateRequest) SetSystemChannelFlags(v int32)`

SetSystemChannelFlags sets SystemChannelFlags field to given value.

### HasSystemChannelFlags

`func (o *GuildCreateRequest) HasSystemChannelFlags() bool`

HasSystemChannelFlags returns a boolean if a field has been set.

### SetSystemChannelFlagsNil

`func (o *GuildCreateRequest) SetSystemChannelFlagsNil(b bool)`

 SetSystemChannelFlagsNil sets the value for SystemChannelFlags to be an explicit nil

### UnsetSystemChannelFlags
`func (o *GuildCreateRequest) UnsetSystemChannelFlags()`

UnsetSystemChannelFlags ensures that no value is present for SystemChannelFlags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


