# GuildAuditLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogEntries** | [**[]AuditLogEntryResponse**](AuditLogEntryResponse.md) |  | 
**Users** | [**[]UserResponse**](UserResponse.md) |  | 
**Integrations** | [**[]GuildAuditLogResponseIntegrationsInner**](GuildAuditLogResponseIntegrationsInner.md) |  | 
**Webhooks** | [**[]ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md) |  | 
**GuildScheduledEvents** | [**[]ListGuildScheduledEvents200ResponseInner**](ListGuildScheduledEvents200ResponseInner.md) |  | 
**Threads** | [**[]ThreadResponse**](ThreadResponse.md) |  | 
**ApplicationCommands** | [**[]ApplicationCommandResponse**](ApplicationCommandResponse.md) |  | 
**AutoModerationRules** | [**[]ListAutoModerationRules200ResponseInner**](ListAutoModerationRules200ResponseInner.md) |  | 

## Methods

### NewGuildAuditLogResponse

`func NewGuildAuditLogResponse(auditLogEntries []AuditLogEntryResponse, users []UserResponse, integrations []GuildAuditLogResponseIntegrationsInner, webhooks []ListChannelWebhooks200ResponseInner, guildScheduledEvents []ListGuildScheduledEvents200ResponseInner, threads []ThreadResponse, applicationCommands []ApplicationCommandResponse, autoModerationRules []ListAutoModerationRules200ResponseInner, ) *GuildAuditLogResponse`

NewGuildAuditLogResponse instantiates a new GuildAuditLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuildAuditLogResponseWithDefaults

`func NewGuildAuditLogResponseWithDefaults() *GuildAuditLogResponse`

NewGuildAuditLogResponseWithDefaults instantiates a new GuildAuditLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogEntries

`func (o *GuildAuditLogResponse) GetAuditLogEntries() []AuditLogEntryResponse`

GetAuditLogEntries returns the AuditLogEntries field if non-nil, zero value otherwise.

### GetAuditLogEntriesOk

`func (o *GuildAuditLogResponse) GetAuditLogEntriesOk() (*[]AuditLogEntryResponse, bool)`

GetAuditLogEntriesOk returns a tuple with the AuditLogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogEntries

`func (o *GuildAuditLogResponse) SetAuditLogEntries(v []AuditLogEntryResponse)`

SetAuditLogEntries sets AuditLogEntries field to given value.


### GetUsers

`func (o *GuildAuditLogResponse) GetUsers() []UserResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GuildAuditLogResponse) GetUsersOk() (*[]UserResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GuildAuditLogResponse) SetUsers(v []UserResponse)`

SetUsers sets Users field to given value.


### GetIntegrations

`func (o *GuildAuditLogResponse) GetIntegrations() []GuildAuditLogResponseIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *GuildAuditLogResponse) GetIntegrationsOk() (*[]GuildAuditLogResponseIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *GuildAuditLogResponse) SetIntegrations(v []GuildAuditLogResponseIntegrationsInner)`

SetIntegrations sets Integrations field to given value.


### GetWebhooks

`func (o *GuildAuditLogResponse) GetWebhooks() []ListChannelWebhooks200ResponseInner`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *GuildAuditLogResponse) GetWebhooksOk() (*[]ListChannelWebhooks200ResponseInner, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *GuildAuditLogResponse) SetWebhooks(v []ListChannelWebhooks200ResponseInner)`

SetWebhooks sets Webhooks field to given value.


### GetGuildScheduledEvents

`func (o *GuildAuditLogResponse) GetGuildScheduledEvents() []ListGuildScheduledEvents200ResponseInner`

GetGuildScheduledEvents returns the GuildScheduledEvents field if non-nil, zero value otherwise.

### GetGuildScheduledEventsOk

`func (o *GuildAuditLogResponse) GetGuildScheduledEventsOk() (*[]ListGuildScheduledEvents200ResponseInner, bool)`

GetGuildScheduledEventsOk returns a tuple with the GuildScheduledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEvents

`func (o *GuildAuditLogResponse) SetGuildScheduledEvents(v []ListGuildScheduledEvents200ResponseInner)`

SetGuildScheduledEvents sets GuildScheduledEvents field to given value.


### GetThreads

`func (o *GuildAuditLogResponse) GetThreads() []ThreadResponse`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *GuildAuditLogResponse) GetThreadsOk() (*[]ThreadResponse, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *GuildAuditLogResponse) SetThreads(v []ThreadResponse)`

SetThreads sets Threads field to given value.


### GetApplicationCommands

`func (o *GuildAuditLogResponse) GetApplicationCommands() []ApplicationCommandResponse`

GetApplicationCommands returns the ApplicationCommands field if non-nil, zero value otherwise.

### GetApplicationCommandsOk

`func (o *GuildAuditLogResponse) GetApplicationCommandsOk() (*[]ApplicationCommandResponse, bool)`

GetApplicationCommandsOk returns a tuple with the ApplicationCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCommands

`func (o *GuildAuditLogResponse) SetApplicationCommands(v []ApplicationCommandResponse)`

SetApplicationCommands sets ApplicationCommands field to given value.


### GetAutoModerationRules

`func (o *GuildAuditLogResponse) GetAutoModerationRules() []ListAutoModerationRules200ResponseInner`

GetAutoModerationRules returns the AutoModerationRules field if non-nil, zero value otherwise.

### GetAutoModerationRulesOk

`func (o *GuildAuditLogResponse) GetAutoModerationRulesOk() (*[]ListAutoModerationRules200ResponseInner, bool)`

GetAutoModerationRulesOk returns a tuple with the AutoModerationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoModerationRules

`func (o *GuildAuditLogResponse) SetAutoModerationRules(v []ListAutoModerationRules200ResponseInner)`

SetAutoModerationRules sets AutoModerationRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


