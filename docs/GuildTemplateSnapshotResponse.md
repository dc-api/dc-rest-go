# GuildTemplateSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**VerificationLevel** | [**VerificationLevels**](VerificationLevels.md) |  | 
**DefaultMessageNotifications** | [**UserNotificationSettings**](UserNotificationSettings.md) |  | 
**ExplicitContentFilter** | [**GuildExplicitContentFilterTypes**](GuildExplicitContentFilterTypes.md) |  | 
**PreferredLocale** | [**AvailableLocalesEnum**](AvailableLocalesEnum.md) |  | 
**AfkChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**AfkTimeout** | [**AfkTimeouts**](AfkTimeouts.md) |  | 
**SystemChannelId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**SystemChannelFlags** | **int32** |  | 
**Roles** | [**[]GuildTemplateRoleResponse**](GuildTemplateRoleResponse.md) |  | 
**Channels** | [**[]GuildTemplateChannelResponse**](GuildTemplateChannelResponse.md) |  | 

## Methods

### NewGuildTemplateSnapshotResponse

`func NewGuildTemplateSnapshotResponse(name string, verificationLevel VerificationLevels, defaultMessageNotifications UserNotificationSettings, explicitContentFilter GuildExplicitContentFilterTypes, preferredLocale AvailableLocalesEnum, afkTimeout AfkTimeouts, systemChannelFlags int32, roles []GuildTemplateRoleResponse, channels []GuildTemplateChannelResponse, ) *GuildTemplateSnapshotResponse`

NewGuildTemplateSnapshotResponse instantiates a new GuildTemplateSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildTemplateSnapshotResponseWithDefaults

`func NewGuildTemplateSnapshotResponseWithDefaults() *GuildTemplateSnapshotResponse`

NewGuildTemplateSnapshotResponseWithDefaults instantiates a new GuildTemplateSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GuildTemplateSnapshotResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuildTemplateSnapshotResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuildTemplateSnapshotResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GuildTemplateSnapshotResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuildTemplateSnapshotResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuildTemplateSnapshotResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuildTemplateSnapshotResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRegion

`func (o *GuildTemplateSnapshotResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GuildTemplateSnapshotResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GuildTemplateSnapshotResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GuildTemplateSnapshotResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVerificationLevel

`func (o *GuildTemplateSnapshotResponse) GetVerificationLevel() VerificationLevels`

GetVerificationLevel returns the VerificationLevel field if non-nil, zero value otherwise.

### GetVerificationLevelOk

`func (o *GuildTemplateSnapshotResponse) GetVerificationLevelOk() (*VerificationLevels, bool)`

GetVerificationLevelOk returns a tuple with the VerificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLevel

`func (o *GuildTemplateSnapshotResponse) SetVerificationLevel(v VerificationLevels)`

SetVerificationLevel sets VerificationLevel field to given value.


### GetDefaultMessageNotifications

`func (o *GuildTemplateSnapshotResponse) GetDefaultMessageNotifications() UserNotificationSettings`

GetDefaultMessageNotifications returns the DefaultMessageNotifications field if non-nil, zero value otherwise.

### GetDefaultMessageNotificationsOk

`func (o *GuildTemplateSnapshotResponse) GetDefaultMessageNotificationsOk() (*UserNotificationSettings, bool)`

GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessageNotifications

`func (o *GuildTemplateSnapshotResponse) SetDefaultMessageNotifications(v UserNotificationSettings)`

SetDefaultMessageNotifications sets DefaultMessageNotifications field to given value.


### GetExplicitContentFilter

`func (o *GuildTemplateSnapshotResponse) GetExplicitContentFilter() GuildExplicitContentFilterTypes`

GetExplicitContentFilter returns the ExplicitContentFilter field if non-nil, zero value otherwise.

### GetExplicitContentFilterOk

`func (o *GuildTemplateSnapshotResponse) GetExplicitContentFilterOk() (*GuildExplicitContentFilterTypes, bool)`

GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContentFilter

`func (o *GuildTemplateSnapshotResponse) SetExplicitContentFilter(v GuildExplicitContentFilterTypes)`

SetExplicitContentFilter sets ExplicitContentFilter field to given value.


### GetPreferredLocale

`func (o *GuildTemplateSnapshotResponse) GetPreferredLocale() AvailableLocalesEnum`

GetPreferredLocale returns the PreferredLocale field if non-nil, zero value otherwise.

### GetPreferredLocaleOk

`func (o *GuildTemplateSnapshotResponse) GetPreferredLocaleOk() (*AvailableLocalesEnum, bool)`

GetPreferredLocaleOk returns a tuple with the PreferredLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocale

`func (o *GuildTemplateSnapshotResponse) SetPreferredLocale(v AvailableLocalesEnum)`

SetPreferredLocale sets PreferredLocale field to given value.


### GetAfkChannelId

`func (o *GuildTemplateSnapshotResponse) GetAfkChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetAfkChannelId returns the AfkChannelId field if non-nil, zero value otherwise.

### GetAfkChannelIdOk

`func (o *GuildTemplateSnapshotResponse) GetAfkChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetAfkChannelIdOk returns a tuple with the AfkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkChannelId

`func (o *GuildTemplateSnapshotResponse) SetAfkChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetAfkChannelId sets AfkChannelId field to given value.

### HasAfkChannelId

`func (o *GuildTemplateSnapshotResponse) HasAfkChannelId() bool`

HasAfkChannelId returns a boolean if a field has been set.

### GetAfkTimeout

`func (o *GuildTemplateSnapshotResponse) GetAfkTimeout() AfkTimeouts`

GetAfkTimeout returns the AfkTimeout field if non-nil, zero value otherwise.

### GetAfkTimeoutOk

`func (o *GuildTemplateSnapshotResponse) GetAfkTimeoutOk() (*AfkTimeouts, bool)`

GetAfkTimeoutOk returns a tuple with the AfkTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkTimeout

`func (o *GuildTemplateSnapshotResponse) SetAfkTimeout(v AfkTimeouts)`

SetAfkTimeout sets AfkTimeout field to given value.


### GetSystemChannelId

`func (o *GuildTemplateSnapshotResponse) GetSystemChannelId() GetEntitlementsSkuIdsParameterOneOfInner`

GetSystemChannelId returns the SystemChannelId field if non-nil, zero value otherwise.

### GetSystemChannelIdOk

`func (o *GuildTemplateSnapshotResponse) GetSystemChannelIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetSystemChannelIdOk returns a tuple with the SystemChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelId

`func (o *GuildTemplateSnapshotResponse) SetSystemChannelId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetSystemChannelId sets SystemChannelId field to given value.

### HasSystemChannelId

`func (o *GuildTemplateSnapshotResponse) HasSystemChannelId() bool`

HasSystemChannelId returns a boolean if a field has been set.

### GetSystemChannelFlags

`func (o *GuildTemplateSnapshotResponse) GetSystemChannelFlags() int32`

GetSystemChannelFlags returns the SystemChannelFlags field if non-nil, zero value otherwise.

### GetSystemChannelFlagsOk

`func (o *GuildTemplateSnapshotResponse) GetSystemChannelFlagsOk() (*int32, bool)`

GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemChannelFlags

`func (o *GuildTemplateSnapshotResponse) SetSystemChannelFlags(v int32)`

SetSystemChannelFlags sets SystemChannelFlags field to given value.


### GetRoles

`func (o *GuildTemplateSnapshotResponse) GetRoles() []GuildTemplateRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GuildTemplateSnapshotResponse) GetRolesOk() (*[]GuildTemplateRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GuildTemplateSnapshotResponse) SetRoles(v []GuildTemplateRoleResponse)`

SetRoles sets Roles field to given value.


### GetChannels

`func (o *GuildTemplateSnapshotResponse) GetChannels() []GuildTemplateChannelResponse`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *GuildTemplateSnapshotResponse) GetChannelsOk() (*[]GuildTemplateChannelResponse, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *GuildTemplateSnapshotResponse) SetChannels(v []GuildTemplateChannelResponse)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


