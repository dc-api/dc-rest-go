# Go API client for dc_rest

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

## Overview

- API version: 10
- Package version: 10
- Build date: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
- Generator version: 7.14.0

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import dc_rest "github.com/dc-api/dc-rest-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `dc_rest.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), dc_rest.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `dc_rest.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), dc_rest.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `dc_rest.ContextOperationServerIndices` and `dc_rest.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), dc_rest.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dc_rest.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://discord.com/api/v10*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**AddGroupDmUser**](docs/DefaultAPI.md#addgroupdmuser) | **Put** /channels/{channel_id}/recipients/{user_id} | 
*DefaultAPI* | [**AddGuildMember**](docs/DefaultAPI.md#addguildmember) | **Put** /guilds/{guild_id}/members/{user_id} | 
*DefaultAPI* | [**AddGuildMemberRole**](docs/DefaultAPI.md#addguildmemberrole) | **Put** /guilds/{guild_id}/members/{user_id}/roles/{role_id} | 
*DefaultAPI* | [**AddLobbyMember**](docs/DefaultAPI.md#addlobbymember) | **Put** /lobbies/{lobby_id}/members/{user_id} | 
*DefaultAPI* | [**AddMyMessageReaction**](docs/DefaultAPI.md#addmymessagereaction) | **Put** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me | 
*DefaultAPI* | [**AddThreadMember**](docs/DefaultAPI.md#addthreadmember) | **Put** /channels/{channel_id}/thread-members/{user_id} | 
*DefaultAPI* | [**ApplicationsGetActivityInstance**](docs/DefaultAPI.md#applicationsgetactivityinstance) | **Get** /applications/{application_id}/activity-instances/{instance_id} | 
*DefaultAPI* | [**BanUserFromGuild**](docs/DefaultAPI.md#banuserfromguild) | **Put** /guilds/{guild_id}/bans/{user_id} | 
*DefaultAPI* | [**BulkBanUsersFromGuild**](docs/DefaultAPI.md#bulkbanusersfromguild) | **Post** /guilds/{guild_id}/bulk-ban | 
*DefaultAPI* | [**BulkDeleteMessages**](docs/DefaultAPI.md#bulkdeletemessages) | **Post** /channels/{channel_id}/messages/bulk-delete | 
*DefaultAPI* | [**BulkSetApplicationCommands**](docs/DefaultAPI.md#bulksetapplicationcommands) | **Put** /applications/{application_id}/commands | 
*DefaultAPI* | [**BulkSetGuildApplicationCommands**](docs/DefaultAPI.md#bulksetguildapplicationcommands) | **Put** /applications/{application_id}/guilds/{guild_id}/commands | 
*DefaultAPI* | [**BulkUpdateGuildChannels**](docs/DefaultAPI.md#bulkupdateguildchannels) | **Patch** /guilds/{guild_id}/channels | 
*DefaultAPI* | [**BulkUpdateGuildRoles**](docs/DefaultAPI.md#bulkupdateguildroles) | **Patch** /guilds/{guild_id}/roles | 
*DefaultAPI* | [**BulkUpdateLobbyMembers**](docs/DefaultAPI.md#bulkupdatelobbymembers) | **Post** /lobbies/{lobby_id}/members/bulk | 
*DefaultAPI* | [**ConsumeEntitlement**](docs/DefaultAPI.md#consumeentitlement) | **Post** /applications/{application_id}/entitlements/{entitlement_id}/consume | 
*DefaultAPI* | [**CreateApplicationCommand**](docs/DefaultAPI.md#createapplicationcommand) | **Post** /applications/{application_id}/commands | 
*DefaultAPI* | [**CreateApplicationEmoji**](docs/DefaultAPI.md#createapplicationemoji) | **Post** /applications/{application_id}/emojis | 
*DefaultAPI* | [**CreateAutoModerationRule**](docs/DefaultAPI.md#createautomoderationrule) | **Post** /guilds/{guild_id}/auto-moderation/rules | 
*DefaultAPI* | [**CreateChannelInvite**](docs/DefaultAPI.md#createchannelinvite) | **Post** /channels/{channel_id}/invites | 
*DefaultAPI* | [**CreateDm**](docs/DefaultAPI.md#createdm) | **Post** /users/@me/channels | 
*DefaultAPI* | [**CreateEntitlement**](docs/DefaultAPI.md#createentitlement) | **Post** /applications/{application_id}/entitlements | 
*DefaultAPI* | [**CreateGuild**](docs/DefaultAPI.md#createguild) | **Post** /guilds | 
*DefaultAPI* | [**CreateGuildApplicationCommand**](docs/DefaultAPI.md#createguildapplicationcommand) | **Post** /applications/{application_id}/guilds/{guild_id}/commands | 
*DefaultAPI* | [**CreateGuildChannel**](docs/DefaultAPI.md#createguildchannel) | **Post** /guilds/{guild_id}/channels | 
*DefaultAPI* | [**CreateGuildEmoji**](docs/DefaultAPI.md#createguildemoji) | **Post** /guilds/{guild_id}/emojis | 
*DefaultAPI* | [**CreateGuildFromTemplate**](docs/DefaultAPI.md#createguildfromtemplate) | **Post** /guilds/templates/{code} | 
*DefaultAPI* | [**CreateGuildRole**](docs/DefaultAPI.md#createguildrole) | **Post** /guilds/{guild_id}/roles | 
*DefaultAPI* | [**CreateGuildScheduledEvent**](docs/DefaultAPI.md#createguildscheduledevent) | **Post** /guilds/{guild_id}/scheduled-events | 
*DefaultAPI* | [**CreateGuildSoundboardSound**](docs/DefaultAPI.md#createguildsoundboardsound) | **Post** /guilds/{guild_id}/soundboard-sounds | 
*DefaultAPI* | [**CreateGuildSticker**](docs/DefaultAPI.md#createguildsticker) | **Post** /guilds/{guild_id}/stickers | 
*DefaultAPI* | [**CreateGuildTemplate**](docs/DefaultAPI.md#createguildtemplate) | **Post** /guilds/{guild_id}/templates | 
*DefaultAPI* | [**CreateInteractionResponse**](docs/DefaultAPI.md#createinteractionresponse) | **Post** /interactions/{interaction_id}/{interaction_token}/callback | 
*DefaultAPI* | [**CreateLobby**](docs/DefaultAPI.md#createlobby) | **Post** /lobbies | 
*DefaultAPI* | [**CreateLobbyMessage**](docs/DefaultAPI.md#createlobbymessage) | **Post** /lobbies/{lobby_id}/messages | 
*DefaultAPI* | [**CreateMessage**](docs/DefaultAPI.md#createmessage) | **Post** /channels/{channel_id}/messages | 
*DefaultAPI* | [**CreateOrJoinLobby**](docs/DefaultAPI.md#createorjoinlobby) | **Put** /lobbies | 
*DefaultAPI* | [**CreatePin**](docs/DefaultAPI.md#createpin) | **Put** /channels/{channel_id}/messages/pins/{message_id} | 
*DefaultAPI* | [**CreateStageInstance**](docs/DefaultAPI.md#createstageinstance) | **Post** /stage-instances | 
*DefaultAPI* | [**CreateThread**](docs/DefaultAPI.md#createthread) | **Post** /channels/{channel_id}/threads | 
*DefaultAPI* | [**CreateThreadFromMessage**](docs/DefaultAPI.md#createthreadfrommessage) | **Post** /channels/{channel_id}/messages/{message_id}/threads | 
*DefaultAPI* | [**CreateWebhook**](docs/DefaultAPI.md#createwebhook) | **Post** /channels/{channel_id}/webhooks | 
*DefaultAPI* | [**CrosspostMessage**](docs/DefaultAPI.md#crosspostmessage) | **Post** /channels/{channel_id}/messages/{message_id}/crosspost | 
*DefaultAPI* | [**DeleteAllMessageReactions**](docs/DefaultAPI.md#deleteallmessagereactions) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions | 
*DefaultAPI* | [**DeleteAllMessageReactionsByEmoji**](docs/DefaultAPI.md#deleteallmessagereactionsbyemoji) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name} | 
*DefaultAPI* | [**DeleteApplicationCommand**](docs/DefaultAPI.md#deleteapplicationcommand) | **Delete** /applications/{application_id}/commands/{command_id} | 
*DefaultAPI* | [**DeleteApplicationEmoji**](docs/DefaultAPI.md#deleteapplicationemoji) | **Delete** /applications/{application_id}/emojis/{emoji_id} | 
*DefaultAPI* | [**DeleteApplicationUserRoleConnection**](docs/DefaultAPI.md#deleteapplicationuserroleconnection) | **Delete** /users/@me/applications/{application_id}/role-connection | 
*DefaultAPI* | [**DeleteAutoModerationRule**](docs/DefaultAPI.md#deleteautomoderationrule) | **Delete** /guilds/{guild_id}/auto-moderation/rules/{rule_id} | 
*DefaultAPI* | [**DeleteChannel**](docs/DefaultAPI.md#deletechannel) | **Delete** /channels/{channel_id} | 
*DefaultAPI* | [**DeleteChannelPermissionOverwrite**](docs/DefaultAPI.md#deletechannelpermissionoverwrite) | **Delete** /channels/{channel_id}/permissions/{overwrite_id} | 
*DefaultAPI* | [**DeleteEntitlement**](docs/DefaultAPI.md#deleteentitlement) | **Delete** /applications/{application_id}/entitlements/{entitlement_id} | 
*DefaultAPI* | [**DeleteGroupDmUser**](docs/DefaultAPI.md#deletegroupdmuser) | **Delete** /channels/{channel_id}/recipients/{user_id} | 
*DefaultAPI* | [**DeleteGuild**](docs/DefaultAPI.md#deleteguild) | **Delete** /guilds/{guild_id} | 
*DefaultAPI* | [**DeleteGuildApplicationCommand**](docs/DefaultAPI.md#deleteguildapplicationcommand) | **Delete** /applications/{application_id}/guilds/{guild_id}/commands/{command_id} | 
*DefaultAPI* | [**DeleteGuildEmoji**](docs/DefaultAPI.md#deleteguildemoji) | **Delete** /guilds/{guild_id}/emojis/{emoji_id} | 
*DefaultAPI* | [**DeleteGuildIntegration**](docs/DefaultAPI.md#deleteguildintegration) | **Delete** /guilds/{guild_id}/integrations/{integration_id} | 
*DefaultAPI* | [**DeleteGuildMember**](docs/DefaultAPI.md#deleteguildmember) | **Delete** /guilds/{guild_id}/members/{user_id} | 
*DefaultAPI* | [**DeleteGuildMemberRole**](docs/DefaultAPI.md#deleteguildmemberrole) | **Delete** /guilds/{guild_id}/members/{user_id}/roles/{role_id} | 
*DefaultAPI* | [**DeleteGuildRole**](docs/DefaultAPI.md#deleteguildrole) | **Delete** /guilds/{guild_id}/roles/{role_id} | 
*DefaultAPI* | [**DeleteGuildScheduledEvent**](docs/DefaultAPI.md#deleteguildscheduledevent) | **Delete** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id} | 
*DefaultAPI* | [**DeleteGuildSoundboardSound**](docs/DefaultAPI.md#deleteguildsoundboardsound) | **Delete** /guilds/{guild_id}/soundboard-sounds/{sound_id} | 
*DefaultAPI* | [**DeleteGuildSticker**](docs/DefaultAPI.md#deleteguildsticker) | **Delete** /guilds/{guild_id}/stickers/{sticker_id} | 
*DefaultAPI* | [**DeleteGuildTemplate**](docs/DefaultAPI.md#deleteguildtemplate) | **Delete** /guilds/{guild_id}/templates/{code} | 
*DefaultAPI* | [**DeleteLobbyMember**](docs/DefaultAPI.md#deletelobbymember) | **Delete** /lobbies/{lobby_id}/members/{user_id} | 
*DefaultAPI* | [**DeleteMessage**](docs/DefaultAPI.md#deletemessage) | **Delete** /channels/{channel_id}/messages/{message_id} | 
*DefaultAPI* | [**DeleteMyMessageReaction**](docs/DefaultAPI.md#deletemymessagereaction) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me | 
*DefaultAPI* | [**DeleteOriginalWebhookMessage**](docs/DefaultAPI.md#deleteoriginalwebhookmessage) | **Delete** /webhooks/{webhook_id}/{webhook_token}/messages/@original | 
*DefaultAPI* | [**DeletePin**](docs/DefaultAPI.md#deletepin) | **Delete** /channels/{channel_id}/messages/pins/{message_id} | 
*DefaultAPI* | [**DeleteStageInstance**](docs/DefaultAPI.md#deletestageinstance) | **Delete** /stage-instances/{channel_id} | 
*DefaultAPI* | [**DeleteThreadMember**](docs/DefaultAPI.md#deletethreadmember) | **Delete** /channels/{channel_id}/thread-members/{user_id} | 
*DefaultAPI* | [**DeleteUserMessageReaction**](docs/DefaultAPI.md#deleteusermessagereaction) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/{user_id} | 
*DefaultAPI* | [**DeleteWebhook**](docs/DefaultAPI.md#deletewebhook) | **Delete** /webhooks/{webhook_id} | 
*DefaultAPI* | [**DeleteWebhookByToken**](docs/DefaultAPI.md#deletewebhookbytoken) | **Delete** /webhooks/{webhook_id}/{webhook_token} | 
*DefaultAPI* | [**DeleteWebhookMessage**](docs/DefaultAPI.md#deletewebhookmessage) | **Delete** /webhooks/{webhook_id}/{webhook_token}/messages/{message_id} | 
*DefaultAPI* | [**DeprecatedCreatePin**](docs/DefaultAPI.md#deprecatedcreatepin) | **Put** /channels/{channel_id}/pins/{message_id} | 
*DefaultAPI* | [**DeprecatedDeletePin**](docs/DefaultAPI.md#deprecateddeletepin) | **Delete** /channels/{channel_id}/pins/{message_id} | 
*DefaultAPI* | [**DeprecatedListPins**](docs/DefaultAPI.md#deprecatedlistpins) | **Get** /channels/{channel_id}/pins | 
*DefaultAPI* | [**EditLobby**](docs/DefaultAPI.md#editlobby) | **Patch** /lobbies/{lobby_id} | 
*DefaultAPI* | [**EditLobbyChannelLink**](docs/DefaultAPI.md#editlobbychannellink) | **Patch** /lobbies/{lobby_id}/channel-linking | 
*DefaultAPI* | [**ExecuteGithubCompatibleWebhook**](docs/DefaultAPI.md#executegithubcompatiblewebhook) | **Post** /webhooks/{webhook_id}/{webhook_token}/github | 
*DefaultAPI* | [**ExecuteSlackCompatibleWebhook**](docs/DefaultAPI.md#executeslackcompatiblewebhook) | **Post** /webhooks/{webhook_id}/{webhook_token}/slack | 
*DefaultAPI* | [**ExecuteWebhook**](docs/DefaultAPI.md#executewebhook) | **Post** /webhooks/{webhook_id}/{webhook_token} | 
*DefaultAPI* | [**FollowChannel**](docs/DefaultAPI.md#followchannel) | **Post** /channels/{channel_id}/followers | 
*DefaultAPI* | [**GetActiveGuildThreads**](docs/DefaultAPI.md#getactiveguildthreads) | **Get** /guilds/{guild_id}/threads/active | 
*DefaultAPI* | [**GetAnswerVoters**](docs/DefaultAPI.md#getanswervoters) | **Get** /channels/{channel_id}/polls/{message_id}/answers/{answer_id} | 
*DefaultAPI* | [**GetApplication**](docs/DefaultAPI.md#getapplication) | **Get** /applications/{application_id} | 
*DefaultAPI* | [**GetApplicationCommand**](docs/DefaultAPI.md#getapplicationcommand) | **Get** /applications/{application_id}/commands/{command_id} | 
*DefaultAPI* | [**GetApplicationEmoji**](docs/DefaultAPI.md#getapplicationemoji) | **Get** /applications/{application_id}/emojis/{emoji_id} | 
*DefaultAPI* | [**GetApplicationRoleConnectionsMetadata**](docs/DefaultAPI.md#getapplicationroleconnectionsmetadata) | **Get** /applications/{application_id}/role-connections/metadata | 
*DefaultAPI* | [**GetApplicationUserRoleConnection**](docs/DefaultAPI.md#getapplicationuserroleconnection) | **Get** /users/@me/applications/{application_id}/role-connection | 
*DefaultAPI* | [**GetAutoModerationRule**](docs/DefaultAPI.md#getautomoderationrule) | **Get** /guilds/{guild_id}/auto-moderation/rules/{rule_id} | 
*DefaultAPI* | [**GetBotGateway**](docs/DefaultAPI.md#getbotgateway) | **Get** /gateway/bot | 
*DefaultAPI* | [**GetChannel**](docs/DefaultAPI.md#getchannel) | **Get** /channels/{channel_id} | 
*DefaultAPI* | [**GetEntitlement**](docs/DefaultAPI.md#getentitlement) | **Get** /applications/{application_id}/entitlements/{entitlement_id} | 
*DefaultAPI* | [**GetEntitlements**](docs/DefaultAPI.md#getentitlements) | **Get** /applications/{application_id}/entitlements | 
*DefaultAPI* | [**GetGateway**](docs/DefaultAPI.md#getgateway) | **Get** /gateway | 
*DefaultAPI* | [**GetGuild**](docs/DefaultAPI.md#getguild) | **Get** /guilds/{guild_id} | 
*DefaultAPI* | [**GetGuildApplicationCommand**](docs/DefaultAPI.md#getguildapplicationcommand) | **Get** /applications/{application_id}/guilds/{guild_id}/commands/{command_id} | 
*DefaultAPI* | [**GetGuildApplicationCommandPermissions**](docs/DefaultAPI.md#getguildapplicationcommandpermissions) | **Get** /applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions | 
*DefaultAPI* | [**GetGuildBan**](docs/DefaultAPI.md#getguildban) | **Get** /guilds/{guild_id}/bans/{user_id} | 
*DefaultAPI* | [**GetGuildEmoji**](docs/DefaultAPI.md#getguildemoji) | **Get** /guilds/{guild_id}/emojis/{emoji_id} | 
*DefaultAPI* | [**GetGuildMember**](docs/DefaultAPI.md#getguildmember) | **Get** /guilds/{guild_id}/members/{user_id} | 
*DefaultAPI* | [**GetGuildNewMemberWelcome**](docs/DefaultAPI.md#getguildnewmemberwelcome) | **Get** /guilds/{guild_id}/new-member-welcome | 
*DefaultAPI* | [**GetGuildPreview**](docs/DefaultAPI.md#getguildpreview) | **Get** /guilds/{guild_id}/preview | 
*DefaultAPI* | [**GetGuildRole**](docs/DefaultAPI.md#getguildrole) | **Get** /guilds/{guild_id}/roles/{role_id} | 
*DefaultAPI* | [**GetGuildScheduledEvent**](docs/DefaultAPI.md#getguildscheduledevent) | **Get** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id} | 
*DefaultAPI* | [**GetGuildSoundboardSound**](docs/DefaultAPI.md#getguildsoundboardsound) | **Get** /guilds/{guild_id}/soundboard-sounds/{sound_id} | 
*DefaultAPI* | [**GetGuildSticker**](docs/DefaultAPI.md#getguildsticker) | **Get** /guilds/{guild_id}/stickers/{sticker_id} | 
*DefaultAPI* | [**GetGuildTemplate**](docs/DefaultAPI.md#getguildtemplate) | **Get** /guilds/templates/{code} | 
*DefaultAPI* | [**GetGuildVanityUrl**](docs/DefaultAPI.md#getguildvanityurl) | **Get** /guilds/{guild_id}/vanity-url | 
*DefaultAPI* | [**GetGuildWebhooks**](docs/DefaultAPI.md#getguildwebhooks) | **Get** /guilds/{guild_id}/webhooks | 
*DefaultAPI* | [**GetGuildWelcomeScreen**](docs/DefaultAPI.md#getguildwelcomescreen) | **Get** /guilds/{guild_id}/welcome-screen | 
*DefaultAPI* | [**GetGuildWidget**](docs/DefaultAPI.md#getguildwidget) | **Get** /guilds/{guild_id}/widget.json | 
*DefaultAPI* | [**GetGuildWidgetPng**](docs/DefaultAPI.md#getguildwidgetpng) | **Get** /guilds/{guild_id}/widget.png | 
*DefaultAPI* | [**GetGuildWidgetSettings**](docs/DefaultAPI.md#getguildwidgetsettings) | **Get** /guilds/{guild_id}/widget | 
*DefaultAPI* | [**GetGuildsOnboarding**](docs/DefaultAPI.md#getguildsonboarding) | **Get** /guilds/{guild_id}/onboarding | 
*DefaultAPI* | [**GetLobby**](docs/DefaultAPI.md#getlobby) | **Get** /lobbies/{lobby_id} | 
*DefaultAPI* | [**GetLobbyMessages**](docs/DefaultAPI.md#getlobbymessages) | **Get** /lobbies/{lobby_id}/messages | 
*DefaultAPI* | [**GetMessage**](docs/DefaultAPI.md#getmessage) | **Get** /channels/{channel_id}/messages/{message_id} | 
*DefaultAPI* | [**GetMyApplication**](docs/DefaultAPI.md#getmyapplication) | **Get** /applications/@me | 
*DefaultAPI* | [**GetMyGuildMember**](docs/DefaultAPI.md#getmyguildmember) | **Get** /users/@me/guilds/{guild_id}/member | 
*DefaultAPI* | [**GetMyOauth2Application**](docs/DefaultAPI.md#getmyoauth2application) | **Get** /oauth2/applications/@me | 
*DefaultAPI* | [**GetMyOauth2Authorization**](docs/DefaultAPI.md#getmyoauth2authorization) | **Get** /oauth2/@me | 
*DefaultAPI* | [**GetMyUser**](docs/DefaultAPI.md#getmyuser) | **Get** /users/@me | 
*DefaultAPI* | [**GetOpenidConnectUserinfo**](docs/DefaultAPI.md#getopenidconnectuserinfo) | **Get** /oauth2/userinfo | 
*DefaultAPI* | [**GetOriginalWebhookMessage**](docs/DefaultAPI.md#getoriginalwebhookmessage) | **Get** /webhooks/{webhook_id}/{webhook_token}/messages/@original | 
*DefaultAPI* | [**GetPublicKeys**](docs/DefaultAPI.md#getpublickeys) | **Get** /oauth2/keys | 
*DefaultAPI* | [**GetSelfVoiceState**](docs/DefaultAPI.md#getselfvoicestate) | **Get** /guilds/{guild_id}/voice-states/@me | 
*DefaultAPI* | [**GetSoundboardDefaultSounds**](docs/DefaultAPI.md#getsoundboarddefaultsounds) | **Get** /soundboard-default-sounds | 
*DefaultAPI* | [**GetStageInstance**](docs/DefaultAPI.md#getstageinstance) | **Get** /stage-instances/{channel_id} | 
*DefaultAPI* | [**GetSticker**](docs/DefaultAPI.md#getsticker) | **Get** /stickers/{sticker_id} | 
*DefaultAPI* | [**GetStickerPack**](docs/DefaultAPI.md#getstickerpack) | **Get** /sticker-packs/{pack_id} | 
*DefaultAPI* | [**GetThreadMember**](docs/DefaultAPI.md#getthreadmember) | **Get** /channels/{channel_id}/thread-members/{user_id} | 
*DefaultAPI* | [**GetUser**](docs/DefaultAPI.md#getuser) | **Get** /users/{user_id} | 
*DefaultAPI* | [**GetVoiceState**](docs/DefaultAPI.md#getvoicestate) | **Get** /guilds/{guild_id}/voice-states/{user_id} | 
*DefaultAPI* | [**GetWebhook**](docs/DefaultAPI.md#getwebhook) | **Get** /webhooks/{webhook_id} | 
*DefaultAPI* | [**GetWebhookByToken**](docs/DefaultAPI.md#getwebhookbytoken) | **Get** /webhooks/{webhook_id}/{webhook_token} | 
*DefaultAPI* | [**GetWebhookMessage**](docs/DefaultAPI.md#getwebhookmessage) | **Get** /webhooks/{webhook_id}/{webhook_token}/messages/{message_id} | 
*DefaultAPI* | [**InviteResolve**](docs/DefaultAPI.md#inviteresolve) | **Get** /invites/{code} | 
*DefaultAPI* | [**InviteRevoke**](docs/DefaultAPI.md#inviterevoke) | **Delete** /invites/{code} | 
*DefaultAPI* | [**JoinThread**](docs/DefaultAPI.md#jointhread) | **Put** /channels/{channel_id}/thread-members/@me | 
*DefaultAPI* | [**LeaveGuild**](docs/DefaultAPI.md#leaveguild) | **Delete** /users/@me/guilds/{guild_id} | 
*DefaultAPI* | [**LeaveLobby**](docs/DefaultAPI.md#leavelobby) | **Delete** /lobbies/{lobby_id}/members/@me | 
*DefaultAPI* | [**LeaveThread**](docs/DefaultAPI.md#leavethread) | **Delete** /channels/{channel_id}/thread-members/@me | 
*DefaultAPI* | [**ListApplicationCommands**](docs/DefaultAPI.md#listapplicationcommands) | **Get** /applications/{application_id}/commands | 
*DefaultAPI* | [**ListApplicationEmojis**](docs/DefaultAPI.md#listapplicationemojis) | **Get** /applications/{application_id}/emojis | 
*DefaultAPI* | [**ListAutoModerationRules**](docs/DefaultAPI.md#listautomoderationrules) | **Get** /guilds/{guild_id}/auto-moderation/rules | 
*DefaultAPI* | [**ListChannelInvites**](docs/DefaultAPI.md#listchannelinvites) | **Get** /channels/{channel_id}/invites | 
*DefaultAPI* | [**ListChannelWebhooks**](docs/DefaultAPI.md#listchannelwebhooks) | **Get** /channels/{channel_id}/webhooks | 
*DefaultAPI* | [**ListGuildApplicationCommandPermissions**](docs/DefaultAPI.md#listguildapplicationcommandpermissions) | **Get** /applications/{application_id}/guilds/{guild_id}/commands/permissions | 
*DefaultAPI* | [**ListGuildApplicationCommands**](docs/DefaultAPI.md#listguildapplicationcommands) | **Get** /applications/{application_id}/guilds/{guild_id}/commands | 
*DefaultAPI* | [**ListGuildAuditLogEntries**](docs/DefaultAPI.md#listguildauditlogentries) | **Get** /guilds/{guild_id}/audit-logs | 
*DefaultAPI* | [**ListGuildBans**](docs/DefaultAPI.md#listguildbans) | **Get** /guilds/{guild_id}/bans | 
*DefaultAPI* | [**ListGuildChannels**](docs/DefaultAPI.md#listguildchannels) | **Get** /guilds/{guild_id}/channels | 
*DefaultAPI* | [**ListGuildEmojis**](docs/DefaultAPI.md#listguildemojis) | **Get** /guilds/{guild_id}/emojis | 
*DefaultAPI* | [**ListGuildIntegrations**](docs/DefaultAPI.md#listguildintegrations) | **Get** /guilds/{guild_id}/integrations | 
*DefaultAPI* | [**ListGuildInvites**](docs/DefaultAPI.md#listguildinvites) | **Get** /guilds/{guild_id}/invites | 
*DefaultAPI* | [**ListGuildMembers**](docs/DefaultAPI.md#listguildmembers) | **Get** /guilds/{guild_id}/members | 
*DefaultAPI* | [**ListGuildRoles**](docs/DefaultAPI.md#listguildroles) | **Get** /guilds/{guild_id}/roles | 
*DefaultAPI* | [**ListGuildScheduledEventUsers**](docs/DefaultAPI.md#listguildscheduledeventusers) | **Get** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users | 
*DefaultAPI* | [**ListGuildScheduledEvents**](docs/DefaultAPI.md#listguildscheduledevents) | **Get** /guilds/{guild_id}/scheduled-events | 
*DefaultAPI* | [**ListGuildSoundboardSounds**](docs/DefaultAPI.md#listguildsoundboardsounds) | **Get** /guilds/{guild_id}/soundboard-sounds | 
*DefaultAPI* | [**ListGuildStickers**](docs/DefaultAPI.md#listguildstickers) | **Get** /guilds/{guild_id}/stickers | 
*DefaultAPI* | [**ListGuildTemplates**](docs/DefaultAPI.md#listguildtemplates) | **Get** /guilds/{guild_id}/templates | 
*DefaultAPI* | [**ListGuildVoiceRegions**](docs/DefaultAPI.md#listguildvoiceregions) | **Get** /guilds/{guild_id}/regions | 
*DefaultAPI* | [**ListMessageReactionsByEmoji**](docs/DefaultAPI.md#listmessagereactionsbyemoji) | **Get** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name} | 
*DefaultAPI* | [**ListMessages**](docs/DefaultAPI.md#listmessages) | **Get** /channels/{channel_id}/messages | 
*DefaultAPI* | [**ListMyConnections**](docs/DefaultAPI.md#listmyconnections) | **Get** /users/@me/connections | 
*DefaultAPI* | [**ListMyGuilds**](docs/DefaultAPI.md#listmyguilds) | **Get** /users/@me/guilds | 
*DefaultAPI* | [**ListMyPrivateArchivedThreads**](docs/DefaultAPI.md#listmyprivatearchivedthreads) | **Get** /channels/{channel_id}/users/@me/threads/archived/private | 
*DefaultAPI* | [**ListPins**](docs/DefaultAPI.md#listpins) | **Get** /channels/{channel_id}/messages/pins | 
*DefaultAPI* | [**ListPrivateArchivedThreads**](docs/DefaultAPI.md#listprivatearchivedthreads) | **Get** /channels/{channel_id}/threads/archived/private | 
*DefaultAPI* | [**ListPublicArchivedThreads**](docs/DefaultAPI.md#listpublicarchivedthreads) | **Get** /channels/{channel_id}/threads/archived/public | 
*DefaultAPI* | [**ListStickerPacks**](docs/DefaultAPI.md#liststickerpacks) | **Get** /sticker-packs | 
*DefaultAPI* | [**ListThreadMembers**](docs/DefaultAPI.md#listthreadmembers) | **Get** /channels/{channel_id}/thread-members | 
*DefaultAPI* | [**ListVoiceRegions**](docs/DefaultAPI.md#listvoiceregions) | **Get** /voice/regions | 
*DefaultAPI* | [**PartnerSdkToken**](docs/DefaultAPI.md#partnersdktoken) | **Post** /partner-sdk/token | 
*DefaultAPI* | [**PartnerSdkUnmergeProvisionalAccount**](docs/DefaultAPI.md#partnersdkunmergeprovisionalaccount) | **Post** /partner-sdk/provisional-accounts/unmerge | 
*DefaultAPI* | [**PollExpire**](docs/DefaultAPI.md#pollexpire) | **Post** /channels/{channel_id}/polls/{message_id}/expire | 
*DefaultAPI* | [**PreviewPruneGuild**](docs/DefaultAPI.md#previewpruneguild) | **Get** /guilds/{guild_id}/prune | 
*DefaultAPI* | [**PruneGuild**](docs/DefaultAPI.md#pruneguild) | **Post** /guilds/{guild_id}/prune | 
*DefaultAPI* | [**PutGuildsOnboarding**](docs/DefaultAPI.md#putguildsonboarding) | **Put** /guilds/{guild_id}/onboarding | 
*DefaultAPI* | [**SearchGuildMembers**](docs/DefaultAPI.md#searchguildmembers) | **Get** /guilds/{guild_id}/members/search | 
*DefaultAPI* | [**SendSoundboardSound**](docs/DefaultAPI.md#sendsoundboardsound) | **Post** /channels/{channel_id}/send-soundboard-sound | 
*DefaultAPI* | [**SetChannelPermissionOverwrite**](docs/DefaultAPI.md#setchannelpermissionoverwrite) | **Put** /channels/{channel_id}/permissions/{overwrite_id} | 
*DefaultAPI* | [**SetGuildApplicationCommandPermissions**](docs/DefaultAPI.md#setguildapplicationcommandpermissions) | **Put** /applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions | 
*DefaultAPI* | [**SetGuildMfaLevel**](docs/DefaultAPI.md#setguildmfalevel) | **Post** /guilds/{guild_id}/mfa | 
*DefaultAPI* | [**SyncGuildTemplate**](docs/DefaultAPI.md#syncguildtemplate) | **Put** /guilds/{guild_id}/templates/{code} | 
*DefaultAPI* | [**ThreadSearch**](docs/DefaultAPI.md#threadsearch) | **Get** /channels/{channel_id}/threads/search | 
*DefaultAPI* | [**TriggerTypingIndicator**](docs/DefaultAPI.md#triggertypingindicator) | **Post** /channels/{channel_id}/typing | 
*DefaultAPI* | [**UnbanUserFromGuild**](docs/DefaultAPI.md#unbanuserfromguild) | **Delete** /guilds/{guild_id}/bans/{user_id} | 
*DefaultAPI* | [**UpdateApplication**](docs/DefaultAPI.md#updateapplication) | **Patch** /applications/{application_id} | 
*DefaultAPI* | [**UpdateApplicationCommand**](docs/DefaultAPI.md#updateapplicationcommand) | **Patch** /applications/{application_id}/commands/{command_id} | 
*DefaultAPI* | [**UpdateApplicationEmoji**](docs/DefaultAPI.md#updateapplicationemoji) | **Patch** /applications/{application_id}/emojis/{emoji_id} | 
*DefaultAPI* | [**UpdateApplicationRoleConnectionsMetadata**](docs/DefaultAPI.md#updateapplicationroleconnectionsmetadata) | **Put** /applications/{application_id}/role-connections/metadata | 
*DefaultAPI* | [**UpdateApplicationUserRoleConnection**](docs/DefaultAPI.md#updateapplicationuserroleconnection) | **Put** /users/@me/applications/{application_id}/role-connection | 
*DefaultAPI* | [**UpdateAutoModerationRule**](docs/DefaultAPI.md#updateautomoderationrule) | **Patch** /guilds/{guild_id}/auto-moderation/rules/{rule_id} | 
*DefaultAPI* | [**UpdateChannel**](docs/DefaultAPI.md#updatechannel) | **Patch** /channels/{channel_id} | 
*DefaultAPI* | [**UpdateGuild**](docs/DefaultAPI.md#updateguild) | **Patch** /guilds/{guild_id} | 
*DefaultAPI* | [**UpdateGuildApplicationCommand**](docs/DefaultAPI.md#updateguildapplicationcommand) | **Patch** /applications/{application_id}/guilds/{guild_id}/commands/{command_id} | 
*DefaultAPI* | [**UpdateGuildEmoji**](docs/DefaultAPI.md#updateguildemoji) | **Patch** /guilds/{guild_id}/emojis/{emoji_id} | 
*DefaultAPI* | [**UpdateGuildMember**](docs/DefaultAPI.md#updateguildmember) | **Patch** /guilds/{guild_id}/members/{user_id} | 
*DefaultAPI* | [**UpdateGuildRole**](docs/DefaultAPI.md#updateguildrole) | **Patch** /guilds/{guild_id}/roles/{role_id} | 
*DefaultAPI* | [**UpdateGuildScheduledEvent**](docs/DefaultAPI.md#updateguildscheduledevent) | **Patch** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id} | 
*DefaultAPI* | [**UpdateGuildSoundboardSound**](docs/DefaultAPI.md#updateguildsoundboardsound) | **Patch** /guilds/{guild_id}/soundboard-sounds/{sound_id} | 
*DefaultAPI* | [**UpdateGuildSticker**](docs/DefaultAPI.md#updateguildsticker) | **Patch** /guilds/{guild_id}/stickers/{sticker_id} | 
*DefaultAPI* | [**UpdateGuildTemplate**](docs/DefaultAPI.md#updateguildtemplate) | **Patch** /guilds/{guild_id}/templates/{code} | 
*DefaultAPI* | [**UpdateGuildWelcomeScreen**](docs/DefaultAPI.md#updateguildwelcomescreen) | **Patch** /guilds/{guild_id}/welcome-screen | 
*DefaultAPI* | [**UpdateGuildWidgetSettings**](docs/DefaultAPI.md#updateguildwidgetsettings) | **Patch** /guilds/{guild_id}/widget | 
*DefaultAPI* | [**UpdateMessage**](docs/DefaultAPI.md#updatemessage) | **Patch** /channels/{channel_id}/messages/{message_id} | 
*DefaultAPI* | [**UpdateMyApplication**](docs/DefaultAPI.md#updatemyapplication) | **Patch** /applications/@me | 
*DefaultAPI* | [**UpdateMyGuildMember**](docs/DefaultAPI.md#updatemyguildmember) | **Patch** /guilds/{guild_id}/members/@me | 
*DefaultAPI* | [**UpdateMyUser**](docs/DefaultAPI.md#updatemyuser) | **Patch** /users/@me | 
*DefaultAPI* | [**UpdateOriginalWebhookMessage**](docs/DefaultAPI.md#updateoriginalwebhookmessage) | **Patch** /webhooks/{webhook_id}/{webhook_token}/messages/@original | 
*DefaultAPI* | [**UpdateSelfVoiceState**](docs/DefaultAPI.md#updateselfvoicestate) | **Patch** /guilds/{guild_id}/voice-states/@me | 
*DefaultAPI* | [**UpdateStageInstance**](docs/DefaultAPI.md#updatestageinstance) | **Patch** /stage-instances/{channel_id} | 
*DefaultAPI* | [**UpdateVoiceState**](docs/DefaultAPI.md#updatevoicestate) | **Patch** /guilds/{guild_id}/voice-states/{user_id} | 
*DefaultAPI* | [**UpdateWebhook**](docs/DefaultAPI.md#updatewebhook) | **Patch** /webhooks/{webhook_id} | 
*DefaultAPI* | [**UpdateWebhookByToken**](docs/DefaultAPI.md#updatewebhookbytoken) | **Patch** /webhooks/{webhook_id}/{webhook_token} | 
*DefaultAPI* | [**UpdateWebhookMessage**](docs/DefaultAPI.md#updatewebhookmessage) | **Patch** /webhooks/{webhook_id}/{webhook_token}/messages/{message_id} | 
*DefaultAPI* | [**UploadApplicationAttachment**](docs/DefaultAPI.md#uploadapplicationattachment) | **Post** /applications/{application_id}/attachment | 


## Documentation For Models

 - [AccountResponse](docs/AccountResponse.md)
 - [ActionRowComponentForMessageRequest](docs/ActionRowComponentForMessageRequest.md)
 - [ActionRowComponentForMessageRequestComponentsInner](docs/ActionRowComponentForMessageRequestComponentsInner.md)
 - [ActionRowComponentForModalRequest](docs/ActionRowComponentForModalRequest.md)
 - [ActionRowComponentResponse](docs/ActionRowComponentResponse.md)
 - [ActionRowComponentResponseComponentsInner](docs/ActionRowComponentResponseComponentsInner.md)
 - [ActivitiesAttachmentResponse](docs/ActivitiesAttachmentResponse.md)
 - [AddGroupDmUser201Response](docs/AddGroupDmUser201Response.md)
 - [AddGroupDmUserRequest](docs/AddGroupDmUserRequest.md)
 - [AddGuildMemberRequest](docs/AddGuildMemberRequest.md)
 - [AddLobbyMemberRequest](docs/AddLobbyMemberRequest.md)
 - [ApplicationCommandAttachmentOption](docs/ApplicationCommandAttachmentOption.md)
 - [ApplicationCommandAttachmentOptionResponse](docs/ApplicationCommandAttachmentOptionResponse.md)
 - [ApplicationCommandAutocompleteCallbackRequest](docs/ApplicationCommandAutocompleteCallbackRequest.md)
 - [ApplicationCommandAutocompleteCallbackRequestData](docs/ApplicationCommandAutocompleteCallbackRequestData.md)
 - [ApplicationCommandBooleanOption](docs/ApplicationCommandBooleanOption.md)
 - [ApplicationCommandBooleanOptionResponse](docs/ApplicationCommandBooleanOptionResponse.md)
 - [ApplicationCommandChannelOption](docs/ApplicationCommandChannelOption.md)
 - [ApplicationCommandChannelOptionResponse](docs/ApplicationCommandChannelOptionResponse.md)
 - [ApplicationCommandCreateRequest](docs/ApplicationCommandCreateRequest.md)
 - [ApplicationCommandCreateRequestOptionsInner](docs/ApplicationCommandCreateRequestOptionsInner.md)
 - [ApplicationCommandIntegerOption](docs/ApplicationCommandIntegerOption.md)
 - [ApplicationCommandIntegerOptionResponse](docs/ApplicationCommandIntegerOptionResponse.md)
 - [ApplicationCommandInteractionMetadataResponse](docs/ApplicationCommandInteractionMetadataResponse.md)
 - [ApplicationCommandMentionableOption](docs/ApplicationCommandMentionableOption.md)
 - [ApplicationCommandMentionableOptionResponse](docs/ApplicationCommandMentionableOptionResponse.md)
 - [ApplicationCommandNumberOption](docs/ApplicationCommandNumberOption.md)
 - [ApplicationCommandNumberOptionResponse](docs/ApplicationCommandNumberOptionResponse.md)
 - [ApplicationCommandOptionIntegerChoice](docs/ApplicationCommandOptionIntegerChoice.md)
 - [ApplicationCommandOptionIntegerChoiceResponse](docs/ApplicationCommandOptionIntegerChoiceResponse.md)
 - [ApplicationCommandOptionNumberChoice](docs/ApplicationCommandOptionNumberChoice.md)
 - [ApplicationCommandOptionNumberChoiceResponse](docs/ApplicationCommandOptionNumberChoiceResponse.md)
 - [ApplicationCommandOptionStringChoice](docs/ApplicationCommandOptionStringChoice.md)
 - [ApplicationCommandOptionStringChoiceResponse](docs/ApplicationCommandOptionStringChoiceResponse.md)
 - [ApplicationCommandPatchRequestPartial](docs/ApplicationCommandPatchRequestPartial.md)
 - [ApplicationCommandPermission](docs/ApplicationCommandPermission.md)
 - [ApplicationCommandResponse](docs/ApplicationCommandResponse.md)
 - [ApplicationCommandResponseOptionsInner](docs/ApplicationCommandResponseOptionsInner.md)
 - [ApplicationCommandRoleOption](docs/ApplicationCommandRoleOption.md)
 - [ApplicationCommandRoleOptionResponse](docs/ApplicationCommandRoleOptionResponse.md)
 - [ApplicationCommandStringOption](docs/ApplicationCommandStringOption.md)
 - [ApplicationCommandStringOptionResponse](docs/ApplicationCommandStringOptionResponse.md)
 - [ApplicationCommandSubcommandGroupOption](docs/ApplicationCommandSubcommandGroupOption.md)
 - [ApplicationCommandSubcommandGroupOptionResponse](docs/ApplicationCommandSubcommandGroupOptionResponse.md)
 - [ApplicationCommandSubcommandOption](docs/ApplicationCommandSubcommandOption.md)
 - [ApplicationCommandSubcommandOptionOptionsInner](docs/ApplicationCommandSubcommandOptionOptionsInner.md)
 - [ApplicationCommandSubcommandOptionResponse](docs/ApplicationCommandSubcommandOptionResponse.md)
 - [ApplicationCommandSubcommandOptionResponseOptionsInner](docs/ApplicationCommandSubcommandOptionResponseOptionsInner.md)
 - [ApplicationCommandUpdateRequest](docs/ApplicationCommandUpdateRequest.md)
 - [ApplicationCommandUserOption](docs/ApplicationCommandUserOption.md)
 - [ApplicationCommandUserOptionResponse](docs/ApplicationCommandUserOptionResponse.md)
 - [ApplicationFormPartial](docs/ApplicationFormPartial.md)
 - [ApplicationFormPartialDescription](docs/ApplicationFormPartialDescription.md)
 - [ApplicationFormPartialIntegrationTypesConfigValue](docs/ApplicationFormPartialIntegrationTypesConfigValue.md)
 - [ApplicationIncomingWebhookResponse](docs/ApplicationIncomingWebhookResponse.md)
 - [ApplicationIntegrationTypeConfiguration](docs/ApplicationIntegrationTypeConfiguration.md)
 - [ApplicationIntegrationTypeConfigurationResponse](docs/ApplicationIntegrationTypeConfigurationResponse.md)
 - [ApplicationOAuth2InstallParams](docs/ApplicationOAuth2InstallParams.md)
 - [ApplicationOAuth2InstallParamsResponse](docs/ApplicationOAuth2InstallParamsResponse.md)
 - [ApplicationResponse](docs/ApplicationResponse.md)
 - [ApplicationRoleConnectionsMetadataItemRequest](docs/ApplicationRoleConnectionsMetadataItemRequest.md)
 - [ApplicationRoleConnectionsMetadataItemResponse](docs/ApplicationRoleConnectionsMetadataItemResponse.md)
 - [ApplicationUserRoleConnectionResponse](docs/ApplicationUserRoleConnectionResponse.md)
 - [AttachmentResponse](docs/AttachmentResponse.md)
 - [AuditLogEntryResponse](docs/AuditLogEntryResponse.md)
 - [AuditLogObjectChangeResponse](docs/AuditLogObjectChangeResponse.md)
 - [BanUserFromGuildRequest](docs/BanUserFromGuildRequest.md)
 - [BaseCreateMessageCreateRequest](docs/BaseCreateMessageCreateRequest.md)
 - [BaseCreateMessageCreateRequestComponentsInner](docs/BaseCreateMessageCreateRequestComponentsInner.md)
 - [BasicApplicationResponse](docs/BasicApplicationResponse.md)
 - [BasicMessageResponse](docs/BasicMessageResponse.md)
 - [BasicMessageResponseComponentsInner](docs/BasicMessageResponseComponentsInner.md)
 - [BasicMessageResponseInteractionMetadata](docs/BasicMessageResponseInteractionMetadata.md)
 - [BasicMessageResponseNonce](docs/BasicMessageResponseNonce.md)
 - [BlockMessageAction](docs/BlockMessageAction.md)
 - [BlockMessageActionMetadata](docs/BlockMessageActionMetadata.md)
 - [BlockMessageActionMetadataResponse](docs/BlockMessageActionMetadataResponse.md)
 - [BlockMessageActionResponse](docs/BlockMessageActionResponse.md)
 - [BotAccountPatchRequest](docs/BotAccountPatchRequest.md)
 - [BulkBanUsersFromGuildRequest](docs/BulkBanUsersFromGuildRequest.md)
 - [BulkBanUsersResponse](docs/BulkBanUsersResponse.md)
 - [BulkDeleteMessagesRequest](docs/BulkDeleteMessagesRequest.md)
 - [BulkLobbyMemberRequest](docs/BulkLobbyMemberRequest.md)
 - [BulkUpdateGuildChannelsRequestInner](docs/BulkUpdateGuildChannelsRequestInner.md)
 - [BulkUpdateGuildRolesRequestInner](docs/BulkUpdateGuildRolesRequestInner.md)
 - [ButtonComponentForMessageRequest](docs/ButtonComponentForMessageRequest.md)
 - [ButtonComponentResponse](docs/ButtonComponentResponse.md)
 - [ChannelFollowerResponse](docs/ChannelFollowerResponse.md)
 - [ChannelFollowerWebhookResponse](docs/ChannelFollowerWebhookResponse.md)
 - [ChannelPermissionOverwriteRequest](docs/ChannelPermissionOverwriteRequest.md)
 - [ChannelPermissionOverwriteResponse](docs/ChannelPermissionOverwriteResponse.md)
 - [ChannelSelectComponentForMessageRequest](docs/ChannelSelectComponentForMessageRequest.md)
 - [ChannelSelectComponentResponse](docs/ChannelSelectComponentResponse.md)
 - [ChannelSelectDefaultValue](docs/ChannelSelectDefaultValue.md)
 - [ChannelSelectDefaultValueResponse](docs/ChannelSelectDefaultValueResponse.md)
 - [CommandPermissionResponse](docs/CommandPermissionResponse.md)
 - [CommandPermissionsResponse](docs/CommandPermissionsResponse.md)
 - [ComponentEmojiForMessageRequest](docs/ComponentEmojiForMessageRequest.md)
 - [ComponentEmojiResponse](docs/ComponentEmojiResponse.md)
 - [ConnectedAccountGuildResponse](docs/ConnectedAccountGuildResponse.md)
 - [ConnectedAccountIntegrationResponse](docs/ConnectedAccountIntegrationResponse.md)
 - [ConnectedAccountResponse](docs/ConnectedAccountResponse.md)
 - [ContainerComponentForMessageRequest](docs/ContainerComponentForMessageRequest.md)
 - [ContainerComponentForMessageRequestComponentsInner](docs/ContainerComponentForMessageRequestComponentsInner.md)
 - [ContainerComponentResponse](docs/ContainerComponentResponse.md)
 - [ContainerComponentResponseComponentsInner](docs/ContainerComponentResponseComponentsInner.md)
 - [CreateApplicationEmojiRequest](docs/CreateApplicationEmojiRequest.md)
 - [CreateAutoModerationRule200Response](docs/CreateAutoModerationRule200Response.md)
 - [CreateAutoModerationRuleRequest](docs/CreateAutoModerationRuleRequest.md)
 - [CreateChannelInviteRequest](docs/CreateChannelInviteRequest.md)
 - [CreateEntitlementRequestData](docs/CreateEntitlementRequestData.md)
 - [CreateForumThreadRequest](docs/CreateForumThreadRequest.md)
 - [CreateGroupDMInviteRequest](docs/CreateGroupDMInviteRequest.md)
 - [CreateGuildChannelRequest](docs/CreateGuildChannelRequest.md)
 - [CreateGuildEmojiRequest](docs/CreateGuildEmojiRequest.md)
 - [CreateGuildFromTemplateRequest](docs/CreateGuildFromTemplateRequest.md)
 - [CreateGuildInviteRequest](docs/CreateGuildInviteRequest.md)
 - [CreateGuildRequestChannelItem](docs/CreateGuildRequestChannelItem.md)
 - [CreateGuildRequestRoleItem](docs/CreateGuildRequestRoleItem.md)
 - [CreateGuildRoleRequest](docs/CreateGuildRoleRequest.md)
 - [CreateGuildScheduledEventRequest](docs/CreateGuildScheduledEventRequest.md)
 - [CreateGuildTemplateRequest](docs/CreateGuildTemplateRequest.md)
 - [CreateInteractionResponseRequest](docs/CreateInteractionResponseRequest.md)
 - [CreateLobbyRequest](docs/CreateLobbyRequest.md)
 - [CreateMessageInteractionCallbackRequest](docs/CreateMessageInteractionCallbackRequest.md)
 - [CreateMessageInteractionCallbackResponse](docs/CreateMessageInteractionCallbackResponse.md)
 - [CreateOrJoinLobbyRequest](docs/CreateOrJoinLobbyRequest.md)
 - [CreateOrUpdateThreadTagRequest](docs/CreateOrUpdateThreadTagRequest.md)
 - [CreatePrivateChannelRequest](docs/CreatePrivateChannelRequest.md)
 - [CreateStageInstanceRequest](docs/CreateStageInstanceRequest.md)
 - [CreateTextThreadWithMessageRequest](docs/CreateTextThreadWithMessageRequest.md)
 - [CreateTextThreadWithoutMessageRequest](docs/CreateTextThreadWithoutMessageRequest.md)
 - [CreateThreadRequest](docs/CreateThreadRequest.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [CreatedThreadResponse](docs/CreatedThreadResponse.md)
 - [DefaultKeywordListTriggerMetadata](docs/DefaultKeywordListTriggerMetadata.md)
 - [DefaultKeywordListTriggerMetadataResponse](docs/DefaultKeywordListTriggerMetadataResponse.md)
 - [DefaultKeywordListUpsertRequest](docs/DefaultKeywordListUpsertRequest.md)
 - [DefaultKeywordListUpsertRequestActionsInner](docs/DefaultKeywordListUpsertRequestActionsInner.md)
 - [DefaultKeywordListUpsertRequestPartial](docs/DefaultKeywordListUpsertRequestPartial.md)
 - [DefaultKeywordRuleResponse](docs/DefaultKeywordRuleResponse.md)
 - [DefaultKeywordRuleResponseActionsInner](docs/DefaultKeywordRuleResponseActionsInner.md)
 - [DefaultReactionEmojiResponse](docs/DefaultReactionEmojiResponse.md)
 - [DiscordIntegrationResponse](docs/DiscordIntegrationResponse.md)
 - [EditLobbyChannelLinkRequest](docs/EditLobbyChannelLinkRequest.md)
 - [EmbeddedActivityInstance](docs/EmbeddedActivityInstance.md)
 - [EmbeddedActivityInstanceLocation](docs/EmbeddedActivityInstanceLocation.md)
 - [EmojiResponse](docs/EmojiResponse.md)
 - [EntitlementResponse](docs/EntitlementResponse.md)
 - [EntityMetadataExternal](docs/EntityMetadataExternal.md)
 - [EntityMetadataExternalResponse](docs/EntityMetadataExternalResponse.md)
 - [Error](docs/Error.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ExecuteWebhookRequest](docs/ExecuteWebhookRequest.md)
 - [ExternalConnectionIntegrationResponse](docs/ExternalConnectionIntegrationResponse.md)
 - [ExternalScheduledEventCreateRequest](docs/ExternalScheduledEventCreateRequest.md)
 - [ExternalScheduledEventPatchRequestPartial](docs/ExternalScheduledEventPatchRequestPartial.md)
 - [ExternalScheduledEventResponse](docs/ExternalScheduledEventResponse.md)
 - [FileComponentForMessageRequest](docs/FileComponentForMessageRequest.md)
 - [FileComponentResponse](docs/FileComponentResponse.md)
 - [FlagToChannelAction](docs/FlagToChannelAction.md)
 - [FlagToChannelActionMetadata](docs/FlagToChannelActionMetadata.md)
 - [FlagToChannelActionMetadataResponse](docs/FlagToChannelActionMetadataResponse.md)
 - [FlagToChannelActionResponse](docs/FlagToChannelActionResponse.md)
 - [FollowChannelRequest](docs/FollowChannelRequest.md)
 - [ForumTagResponse](docs/ForumTagResponse.md)
 - [FriendInviteResponse](docs/FriendInviteResponse.md)
 - [GatewayBotResponse](docs/GatewayBotResponse.md)
 - [GatewayBotSessionStartLimitResponse](docs/GatewayBotSessionStartLimitResponse.md)
 - [GatewayResponse](docs/GatewayResponse.md)
 - [GetChannel200Response](docs/GetChannel200Response.md)
 - [GetEntitlementsSkuIdsParameter](docs/GetEntitlementsSkuIdsParameter.md)
 - [GetSticker200Response](docs/GetSticker200Response.md)
 - [GithubAuthor](docs/GithubAuthor.md)
 - [GithubCheckApp](docs/GithubCheckApp.md)
 - [GithubCheckPullRequest](docs/GithubCheckPullRequest.md)
 - [GithubCheckRun](docs/GithubCheckRun.md)
 - [GithubCheckRunOutput](docs/GithubCheckRunOutput.md)
 - [GithubCheckSuite](docs/GithubCheckSuite.md)
 - [GithubComment](docs/GithubComment.md)
 - [GithubCommit](docs/GithubCommit.md)
 - [GithubDiscussion](docs/GithubDiscussion.md)
 - [GithubIssue](docs/GithubIssue.md)
 - [GithubRelease](docs/GithubRelease.md)
 - [GithubRepository](docs/GithubRepository.md)
 - [GithubReview](docs/GithubReview.md)
 - [GithubUser](docs/GithubUser.md)
 - [GithubWebhook](docs/GithubWebhook.md)
 - [GroupDMInviteResponse](docs/GroupDMInviteResponse.md)
 - [GuildAuditLogResponse](docs/GuildAuditLogResponse.md)
 - [GuildAuditLogResponseIntegrationsInner](docs/GuildAuditLogResponseIntegrationsInner.md)
 - [GuildBanResponse](docs/GuildBanResponse.md)
 - [GuildChannelLocation](docs/GuildChannelLocation.md)
 - [GuildChannelResponse](docs/GuildChannelResponse.md)
 - [GuildCreateRequest](docs/GuildCreateRequest.md)
 - [GuildHomeSettingsResponse](docs/GuildHomeSettingsResponse.md)
 - [GuildIncomingWebhookResponse](docs/GuildIncomingWebhookResponse.md)
 - [GuildInviteResponse](docs/GuildInviteResponse.md)
 - [GuildMFALevelResponse](docs/GuildMFALevelResponse.md)
 - [GuildMemberResponse](docs/GuildMemberResponse.md)
 - [GuildOnboardingResponse](docs/GuildOnboardingResponse.md)
 - [GuildPatchRequestPartial](docs/GuildPatchRequestPartial.md)
 - [GuildPreviewResponse](docs/GuildPreviewResponse.md)
 - [GuildProductPurchaseResponse](docs/GuildProductPurchaseResponse.md)
 - [GuildPruneResponse](docs/GuildPruneResponse.md)
 - [GuildResponse](docs/GuildResponse.md)
 - [GuildRoleColorsResponse](docs/GuildRoleColorsResponse.md)
 - [GuildRoleResponse](docs/GuildRoleResponse.md)
 - [GuildRoleTagsResponse](docs/GuildRoleTagsResponse.md)
 - [GuildStickerResponse](docs/GuildStickerResponse.md)
 - [GuildSubscriptionIntegrationResponse](docs/GuildSubscriptionIntegrationResponse.md)
 - [GuildTemplateChannelResponse](docs/GuildTemplateChannelResponse.md)
 - [GuildTemplateChannelTags](docs/GuildTemplateChannelTags.md)
 - [GuildTemplateResponse](docs/GuildTemplateResponse.md)
 - [GuildTemplateRoleResponse](docs/GuildTemplateRoleResponse.md)
 - [GuildTemplateSnapshotResponse](docs/GuildTemplateSnapshotResponse.md)
 - [GuildWelcomeChannel](docs/GuildWelcomeChannel.md)
 - [GuildWelcomeScreenChannelResponse](docs/GuildWelcomeScreenChannelResponse.md)
 - [GuildWelcomeScreenResponse](docs/GuildWelcomeScreenResponse.md)
 - [GuildWithCountsResponse](docs/GuildWithCountsResponse.md)
 - [IncomingWebhookInteractionRequest](docs/IncomingWebhookInteractionRequest.md)
 - [IncomingWebhookRequestPartial](docs/IncomingWebhookRequestPartial.md)
 - [IncomingWebhookUpdateForInteractionCallbackRequestPartial](docs/IncomingWebhookUpdateForInteractionCallbackRequestPartial.md)
 - [IncomingWebhookUpdateRequestPartial](docs/IncomingWebhookUpdateRequestPartial.md)
 - [InnerErrors](docs/InnerErrors.md)
 - [IntegrationApplicationResponse](docs/IntegrationApplicationResponse.md)
 - [InteractionApplicationCommandAutocompleteCallbackIntegerData](docs/InteractionApplicationCommandAutocompleteCallbackIntegerData.md)
 - [InteractionApplicationCommandAutocompleteCallbackNumberData](docs/InteractionApplicationCommandAutocompleteCallbackNumberData.md)
 - [InteractionApplicationCommandAutocompleteCallbackStringData](docs/InteractionApplicationCommandAutocompleteCallbackStringData.md)
 - [InteractionCallbackResponse](docs/InteractionCallbackResponse.md)
 - [InteractionCallbackResponseResource](docs/InteractionCallbackResponseResource.md)
 - [InteractionResponse](docs/InteractionResponse.md)
 - [InviteApplicationResponse](docs/InviteApplicationResponse.md)
 - [InviteChannelRecipientResponse](docs/InviteChannelRecipientResponse.md)
 - [InviteChannelResponse](docs/InviteChannelResponse.md)
 - [InviteGuildResponse](docs/InviteGuildResponse.md)
 - [KeywordRuleResponse](docs/KeywordRuleResponse.md)
 - [KeywordTriggerMetadata](docs/KeywordTriggerMetadata.md)
 - [KeywordTriggerMetadataResponse](docs/KeywordTriggerMetadataResponse.md)
 - [KeywordUpsertRequest](docs/KeywordUpsertRequest.md)
 - [KeywordUpsertRequestPartial](docs/KeywordUpsertRequestPartial.md)
 - [LaunchActivityInteractionCallbackRequest](docs/LaunchActivityInteractionCallbackRequest.md)
 - [LaunchActivityInteractionCallbackResponse](docs/LaunchActivityInteractionCallbackResponse.md)
 - [ListApplicationEmojisResponse](docs/ListApplicationEmojisResponse.md)
 - [ListAutoModerationRules200ResponseInner](docs/ListAutoModerationRules200ResponseInner.md)
 - [ListChannelInvites200ResponseInner](docs/ListChannelInvites200ResponseInner.md)
 - [ListChannelWebhooks200ResponseInner](docs/ListChannelWebhooks200ResponseInner.md)
 - [ListGuildIntegrations200ResponseInner](docs/ListGuildIntegrations200ResponseInner.md)
 - [ListGuildScheduledEvents200ResponseInner](docs/ListGuildScheduledEvents200ResponseInner.md)
 - [ListGuildSoundboardSoundsResponse](docs/ListGuildSoundboardSoundsResponse.md)
 - [LobbyMemberRequest](docs/LobbyMemberRequest.md)
 - [LobbyMemberResponse](docs/LobbyMemberResponse.md)
 - [LobbyMessageResponse](docs/LobbyMessageResponse.md)
 - [LobbyResponse](docs/LobbyResponse.md)
 - [MLSpamRuleResponse](docs/MLSpamRuleResponse.md)
 - [MLSpamUpsertRequest](docs/MLSpamUpsertRequest.md)
 - [MLSpamUpsertRequestPartial](docs/MLSpamUpsertRequestPartial.md)
 - [MediaGalleryComponentForMessageRequest](docs/MediaGalleryComponentForMessageRequest.md)
 - [MediaGalleryComponentResponse](docs/MediaGalleryComponentResponse.md)
 - [MediaGalleryItemRequest](docs/MediaGalleryItemRequest.md)
 - [MediaGalleryItemResponse](docs/MediaGalleryItemResponse.md)
 - [MentionSpamRuleResponse](docs/MentionSpamRuleResponse.md)
 - [MentionSpamTriggerMetadata](docs/MentionSpamTriggerMetadata.md)
 - [MentionSpamTriggerMetadataResponse](docs/MentionSpamTriggerMetadataResponse.md)
 - [MentionSpamUpsertRequest](docs/MentionSpamUpsertRequest.md)
 - [MentionSpamUpsertRequestPartial](docs/MentionSpamUpsertRequestPartial.md)
 - [MentionableSelectComponentForMessageRequest](docs/MentionableSelectComponentForMessageRequest.md)
 - [MentionableSelectComponentForMessageRequestDefaultValuesInner](docs/MentionableSelectComponentForMessageRequestDefaultValuesInner.md)
 - [MentionableSelectComponentResponse](docs/MentionableSelectComponentResponse.md)
 - [MentionableSelectComponentResponseDefaultValuesInner](docs/MentionableSelectComponentResponseDefaultValuesInner.md)
 - [MessageAllowedMentionsRequest](docs/MessageAllowedMentionsRequest.md)
 - [MessageAttachmentRequest](docs/MessageAttachmentRequest.md)
 - [MessageAttachmentResponse](docs/MessageAttachmentResponse.md)
 - [MessageCallResponse](docs/MessageCallResponse.md)
 - [MessageComponentInteractionMetadataResponse](docs/MessageComponentInteractionMetadataResponse.md)
 - [MessageCreateRequest](docs/MessageCreateRequest.md)
 - [MessageEditRequestPartial](docs/MessageEditRequestPartial.md)
 - [MessageEmbedAuthorResponse](docs/MessageEmbedAuthorResponse.md)
 - [MessageEmbedFieldResponse](docs/MessageEmbedFieldResponse.md)
 - [MessageEmbedFooterResponse](docs/MessageEmbedFooterResponse.md)
 - [MessageEmbedImageResponse](docs/MessageEmbedImageResponse.md)
 - [MessageEmbedProviderResponse](docs/MessageEmbedProviderResponse.md)
 - [MessageEmbedResponse](docs/MessageEmbedResponse.md)
 - [MessageEmbedVideoResponse](docs/MessageEmbedVideoResponse.md)
 - [MessageInteractionResponse](docs/MessageInteractionResponse.md)
 - [MessageMentionChannelResponse](docs/MessageMentionChannelResponse.md)
 - [MessageReactionCountDetailsResponse](docs/MessageReactionCountDetailsResponse.md)
 - [MessageReactionEmojiResponse](docs/MessageReactionEmojiResponse.md)
 - [MessageReactionResponse](docs/MessageReactionResponse.md)
 - [MessageReferenceRequest](docs/MessageReferenceRequest.md)
 - [MessageReferenceResponse](docs/MessageReferenceResponse.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [MessageRoleSubscriptionDataResponse](docs/MessageRoleSubscriptionDataResponse.md)
 - [MessageSnapshotResponse](docs/MessageSnapshotResponse.md)
 - [MessageStickerItemResponse](docs/MessageStickerItemResponse.md)
 - [MinimalContentMessageResponse](docs/MinimalContentMessageResponse.md)
 - [ModalInteractionCallbackRequest](docs/ModalInteractionCallbackRequest.md)
 - [ModalInteractionCallbackRequestData](docs/ModalInteractionCallbackRequestData.md)
 - [ModalSubmitInteractionMetadataResponse](docs/ModalSubmitInteractionMetadataResponse.md)
 - [ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata](docs/ModalSubmitInteractionMetadataResponseTriggeringInteractionMetadata.md)
 - [MyGuildResponse](docs/MyGuildResponse.md)
 - [NewMemberActionResponse](docs/NewMemberActionResponse.md)
 - [OAuth2GetAuthorizationResponse](docs/OAuth2GetAuthorizationResponse.md)
 - [OAuth2GetKeys](docs/OAuth2GetKeys.md)
 - [OAuth2GetOpenIDConnectUserInfoResponse](docs/OAuth2GetOpenIDConnectUserInfoResponse.md)
 - [OAuth2Key](docs/OAuth2Key.md)
 - [OnboardingPromptOptionRequest](docs/OnboardingPromptOptionRequest.md)
 - [OnboardingPromptOptionResponse](docs/OnboardingPromptOptionResponse.md)
 - [OnboardingPromptResponse](docs/OnboardingPromptResponse.md)
 - [PartialDiscordIntegrationResponse](docs/PartialDiscordIntegrationResponse.md)
 - [PartialExternalConnectionIntegrationResponse](docs/PartialExternalConnectionIntegrationResponse.md)
 - [PartialGuildSubscriptionIntegrationResponse](docs/PartialGuildSubscriptionIntegrationResponse.md)
 - [PartnerSdkUnmergeProvisionalAccountRequest](docs/PartnerSdkUnmergeProvisionalAccountRequest.md)
 - [PinnedMessageResponse](docs/PinnedMessageResponse.md)
 - [PinnedMessagesResponse](docs/PinnedMessagesResponse.md)
 - [PollAnswerCreateRequest](docs/PollAnswerCreateRequest.md)
 - [PollAnswerDetailsResponse](docs/PollAnswerDetailsResponse.md)
 - [PollAnswerResponse](docs/PollAnswerResponse.md)
 - [PollCreateRequest](docs/PollCreateRequest.md)
 - [PollEmoji](docs/PollEmoji.md)
 - [PollEmojiCreateRequest](docs/PollEmojiCreateRequest.md)
 - [PollMedia](docs/PollMedia.md)
 - [PollMediaCreateRequest](docs/PollMediaCreateRequest.md)
 - [PollMediaResponse](docs/PollMediaResponse.md)
 - [PollResponse](docs/PollResponse.md)
 - [PollResultsEntryResponse](docs/PollResultsEntryResponse.md)
 - [PollResultsResponse](docs/PollResultsResponse.md)
 - [PongInteractionCallbackRequest](docs/PongInteractionCallbackRequest.md)
 - [PrivateApplicationResponse](docs/PrivateApplicationResponse.md)
 - [PrivateChannelLocation](docs/PrivateChannelLocation.md)
 - [PrivateChannelResponse](docs/PrivateChannelResponse.md)
 - [PrivateGroupChannelResponse](docs/PrivateGroupChannelResponse.md)
 - [PrivateGuildMemberResponse](docs/PrivateGuildMemberResponse.md)
 - [ProvisionalTokenResponse](docs/ProvisionalTokenResponse.md)
 - [PruneGuildRequest](docs/PruneGuildRequest.md)
 - [PruneGuildRequestIncludeRoles](docs/PruneGuildRequestIncludeRoles.md)
 - [PurchaseNotificationResponse](docs/PurchaseNotificationResponse.md)
 - [QuarantineUserAction](docs/QuarantineUserAction.md)
 - [QuarantineUserActionResponse](docs/QuarantineUserActionResponse.md)
 - [ResolvedObjectsResponse](docs/ResolvedObjectsResponse.md)
 - [ResourceChannelResponse](docs/ResourceChannelResponse.md)
 - [RichEmbed](docs/RichEmbed.md)
 - [RichEmbedAuthor](docs/RichEmbedAuthor.md)
 - [RichEmbedField](docs/RichEmbedField.md)
 - [RichEmbedFooter](docs/RichEmbedFooter.md)
 - [RichEmbedImage](docs/RichEmbedImage.md)
 - [RichEmbedProvider](docs/RichEmbedProvider.md)
 - [RichEmbedThumbnail](docs/RichEmbedThumbnail.md)
 - [RichEmbedVideo](docs/RichEmbedVideo.md)
 - [RoleSelectComponentForMessageRequest](docs/RoleSelectComponentForMessageRequest.md)
 - [RoleSelectComponentResponse](docs/RoleSelectComponentResponse.md)
 - [RoleSelectDefaultValue](docs/RoleSelectDefaultValue.md)
 - [RoleSelectDefaultValueResponse](docs/RoleSelectDefaultValueResponse.md)
 - [SDKMessageRequest](docs/SDKMessageRequest.md)
 - [ScheduledEventResponse](docs/ScheduledEventResponse.md)
 - [ScheduledEventUserResponse](docs/ScheduledEventUserResponse.md)
 - [SectionComponentForMessageRequest](docs/SectionComponentForMessageRequest.md)
 - [SectionComponentForMessageRequestAccessory](docs/SectionComponentForMessageRequestAccessory.md)
 - [SectionComponentResponse](docs/SectionComponentResponse.md)
 - [SectionComponentResponseAccessory](docs/SectionComponentResponseAccessory.md)
 - [SeparatorComponentForMessageRequest](docs/SeparatorComponentForMessageRequest.md)
 - [SeparatorComponentResponse](docs/SeparatorComponentResponse.md)
 - [SetChannelPermissionOverwriteRequest](docs/SetChannelPermissionOverwriteRequest.md)
 - [SetGuildApplicationCommandPermissionsRequest](docs/SetGuildApplicationCommandPermissionsRequest.md)
 - [SetGuildMfaLevelRequest](docs/SetGuildMfaLevelRequest.md)
 - [SettingsEmojiResponse](docs/SettingsEmojiResponse.md)
 - [SlackWebhook](docs/SlackWebhook.md)
 - [SoundboardCreateRequest](docs/SoundboardCreateRequest.md)
 - [SoundboardPatchRequestPartial](docs/SoundboardPatchRequestPartial.md)
 - [SoundboardSoundResponse](docs/SoundboardSoundResponse.md)
 - [SoundboardSoundSendRequest](docs/SoundboardSoundSendRequest.md)
 - [SpamLinkRuleResponse](docs/SpamLinkRuleResponse.md)
 - [StageInstanceResponse](docs/StageInstanceResponse.md)
 - [StageScheduledEventCreateRequest](docs/StageScheduledEventCreateRequest.md)
 - [StageScheduledEventPatchRequestPartial](docs/StageScheduledEventPatchRequestPartial.md)
 - [StageScheduledEventResponse](docs/StageScheduledEventResponse.md)
 - [StandardStickerResponse](docs/StandardStickerResponse.md)
 - [StickerPackCollectionResponse](docs/StickerPackCollectionResponse.md)
 - [StickerPackResponse](docs/StickerPackResponse.md)
 - [StringSelectComponentForMessageRequest](docs/StringSelectComponentForMessageRequest.md)
 - [StringSelectComponentResponse](docs/StringSelectComponentResponse.md)
 - [StringSelectOptionForMessageRequest](docs/StringSelectOptionForMessageRequest.md)
 - [StringSelectOptionResponse](docs/StringSelectOptionResponse.md)
 - [TeamMemberResponse](docs/TeamMemberResponse.md)
 - [TeamResponse](docs/TeamResponse.md)
 - [TextDisplayComponentForMessageRequest](docs/TextDisplayComponentForMessageRequest.md)
 - [TextDisplayComponentResponse](docs/TextDisplayComponentResponse.md)
 - [TextInputComponentForModalRequest](docs/TextInputComponentForModalRequest.md)
 - [TextInputComponentResponse](docs/TextInputComponentResponse.md)
 - [ThreadMemberResponse](docs/ThreadMemberResponse.md)
 - [ThreadMetadataResponse](docs/ThreadMetadataResponse.md)
 - [ThreadResponse](docs/ThreadResponse.md)
 - [ThreadSearchResponse](docs/ThreadSearchResponse.md)
 - [ThreadSearchTagParameter](docs/ThreadSearchTagParameter.md)
 - [ThreadsResponse](docs/ThreadsResponse.md)
 - [ThumbnailComponentForMessageRequest](docs/ThumbnailComponentForMessageRequest.md)
 - [ThumbnailComponentResponse](docs/ThumbnailComponentResponse.md)
 - [UnfurledMediaRequest](docs/UnfurledMediaRequest.md)
 - [UnfurledMediaRequestWithAttachmentReferenceRequired](docs/UnfurledMediaRequestWithAttachmentReferenceRequired.md)
 - [UnfurledMediaResponse](docs/UnfurledMediaResponse.md)
 - [UpdateApplicationEmojiRequest](docs/UpdateApplicationEmojiRequest.md)
 - [UpdateApplicationUserRoleConnectionRequest](docs/UpdateApplicationUserRoleConnectionRequest.md)
 - [UpdateAutoModerationRuleRequest](docs/UpdateAutoModerationRuleRequest.md)
 - [UpdateChannelRequest](docs/UpdateChannelRequest.md)
 - [UpdateDMRequestPartial](docs/UpdateDMRequestPartial.md)
 - [UpdateDefaultReactionEmojiRequest](docs/UpdateDefaultReactionEmojiRequest.md)
 - [UpdateGroupDMRequestPartial](docs/UpdateGroupDMRequestPartial.md)
 - [UpdateGuildChannelRequestPartial](docs/UpdateGuildChannelRequestPartial.md)
 - [UpdateGuildEmojiRequest](docs/UpdateGuildEmojiRequest.md)
 - [UpdateGuildMemberRequest](docs/UpdateGuildMemberRequest.md)
 - [UpdateGuildOnboardingRequest](docs/UpdateGuildOnboardingRequest.md)
 - [UpdateGuildScheduledEventRequest](docs/UpdateGuildScheduledEventRequest.md)
 - [UpdateGuildStickerRequest](docs/UpdateGuildStickerRequest.md)
 - [UpdateGuildTemplateRequest](docs/UpdateGuildTemplateRequest.md)
 - [UpdateGuildWidgetSettingsRequest](docs/UpdateGuildWidgetSettingsRequest.md)
 - [UpdateMessageInteractionCallbackRequest](docs/UpdateMessageInteractionCallbackRequest.md)
 - [UpdateMessageInteractionCallbackResponse](docs/UpdateMessageInteractionCallbackResponse.md)
 - [UpdateMyGuildMemberRequest](docs/UpdateMyGuildMemberRequest.md)
 - [UpdateOnboardingPromptRequest](docs/UpdateOnboardingPromptRequest.md)
 - [UpdateSelfVoiceStateRequest](docs/UpdateSelfVoiceStateRequest.md)
 - [UpdateStageInstanceRequest](docs/UpdateStageInstanceRequest.md)
 - [UpdateThreadRequestPartial](docs/UpdateThreadRequestPartial.md)
 - [UpdateThreadTagRequest](docs/UpdateThreadTagRequest.md)
 - [UpdateVoiceStateRequest](docs/UpdateVoiceStateRequest.md)
 - [UpdateWebhookByTokenRequest](docs/UpdateWebhookByTokenRequest.md)
 - [UpdateWebhookRequest](docs/UpdateWebhookRequest.md)
 - [UserAvatarDecorationResponse](docs/UserAvatarDecorationResponse.md)
 - [UserCollectiblesResponse](docs/UserCollectiblesResponse.md)
 - [UserCommunicationDisabledAction](docs/UserCommunicationDisabledAction.md)
 - [UserCommunicationDisabledActionMetadata](docs/UserCommunicationDisabledActionMetadata.md)
 - [UserCommunicationDisabledActionMetadataResponse](docs/UserCommunicationDisabledActionMetadataResponse.md)
 - [UserCommunicationDisabledActionResponse](docs/UserCommunicationDisabledActionResponse.md)
 - [UserGuildOnboardingResponse](docs/UserGuildOnboardingResponse.md)
 - [UserNameplateResponse](docs/UserNameplateResponse.md)
 - [UserPIIResponse](docs/UserPIIResponse.md)
 - [UserPrimaryGuildResponse](docs/UserPrimaryGuildResponse.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserSelectComponentForMessageRequest](docs/UserSelectComponentForMessageRequest.md)
 - [UserSelectComponentResponse](docs/UserSelectComponentResponse.md)
 - [UserSelectDefaultValue](docs/UserSelectDefaultValue.md)
 - [UserSelectDefaultValueResponse](docs/UserSelectDefaultValueResponse.md)
 - [VanityURLErrorResponse](docs/VanityURLErrorResponse.md)
 - [VanityURLResponse](docs/VanityURLResponse.md)
 - [VoiceRegionResponse](docs/VoiceRegionResponse.md)
 - [VoiceScheduledEventCreateRequest](docs/VoiceScheduledEventCreateRequest.md)
 - [VoiceScheduledEventPatchRequestPartial](docs/VoiceScheduledEventPatchRequestPartial.md)
 - [VoiceScheduledEventResponse](docs/VoiceScheduledEventResponse.md)
 - [VoiceStateResponse](docs/VoiceStateResponse.md)
 - [WebhookSlackEmbed](docs/WebhookSlackEmbed.md)
 - [WebhookSlackEmbedField](docs/WebhookSlackEmbedField.md)
 - [WebhookSourceChannelResponse](docs/WebhookSourceChannelResponse.md)
 - [WebhookSourceGuildResponse](docs/WebhookSourceGuildResponse.md)
 - [WelcomeMessageResponse](docs/WelcomeMessageResponse.md)
 - [WelcomeScreenPatchRequestPartial](docs/WelcomeScreenPatchRequestPartial.md)
 - [WidgetActivity](docs/WidgetActivity.md)
 - [WidgetChannel](docs/WidgetChannel.md)
 - [WidgetMember](docs/WidgetMember.md)
 - [WidgetResponse](docs/WidgetResponse.md)
 - [WidgetSettingsResponse](docs/WidgetSettingsResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BotToken

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: BotToken and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		dc_rest.ContextAPIKeys,
		map[string]dc_rest.APIKey{
			"BotToken": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### OAuth2


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://discord.com/api/oauth2/authorize
- **Scopes**: 
 - **activities.invites.write**: allows your app to send activity invites - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
 - **activities.read**: allows your app to fetch data from a user's \"Now Playing/Recently Played\" list - requires Discord approval
 - **activities.write**: allows your app to update a user's activity - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
 - **applications.builds.read**: allows your app to read build data for a user's applications
 - **applications.builds.upload**: allows your app to upload/update builds for a user's applications - requires Discord approval
 - **applications.commands**: allows your app to use commands in a guild
 - **applications.commands.permissions.update**: allows your app to update permissions for its commands in a guild a user has permissions to
 - **applications.entitlements**: allows your app to read entitlements for a user's applications
 - **applications.store.update**: allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications
 - **bot**: for oauth2 bots, this puts the bot in the user's selected guild by default
 - **connections**: allows /users/@me/connections to return linked third-party accounts
 - **dm_channels.read**: allows your app to see information about the user's DMs and group DMs - requires Discord approval
 - **email**: enables /users/@me to return an email
 - **gdm.join**: allows your app to join users to a group dm
 - **guilds**: allows /users/@me/guilds to return basic information about all of a user's guilds
 - **guilds.join**: allows /guilds/{guild.id}/members/{user.id} to be used for joining users to a guild
 - **guilds.members.read**: allows /users/@me/guilds/{guild.id}/member to return a user's member information in a guild
 - **identify**: allows /users/@me without email
 - **messages.read**: for local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates)
 - **openid**: for OpenID Connect, this allows your app to receive user id and basic profile information
 - **relationships.read**: allows your app to know a user's friends and implicit relationships - requires Discord approval
 - **rpc**: for local rpc server access, this allows you to control a user's local Discord client - requires Discord approval
 - **rpc.activities.write**: for local rpc server access, this allows you to update a user's activity - requires Discord approval
 - **rpc.notifications.read**: for local rpc server access, this allows you to receive notifications pushed out to the user - requires Discord approval
 - **rpc.screenshare.read**: for local rpc server access, this allows you to read a user's screenshare status- requires Discord approval
 - **rpc.screenshare.write**: for local rpc server access, this allows you to update a user's screenshare settings- requires Discord approval
 - **rpc.video.read**: for local rpc server access, this allows you to read a user's video status - requires Discord approval
 - **rpc.video.write**: for local rpc server access, this allows you to update a user's video settings - requires Discord approval
 - **rpc.voice.read**: for local rpc server access, this allows you to read a user's voice settings and listen for voice events - requires Discord approval
 - **rpc.voice.write**: for local rpc server access, this allows you to update a user's voice settings - requires Discord approval
 - **voice**: allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval
 - **webhook.incoming**: this generates a webhook that is returned in the oauth token response for authorization code grants

Example

```go
auth := context.WithValue(context.Background(), dc_rest.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, dc_rest.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **activities.invites.write**: allows your app to send activity invites - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
 - **activities.read**: allows your app to fetch data from a user's \"Now Playing/Recently Played\" list - requires Discord approval
 - **activities.write**: allows your app to update a user's activity - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
 - **applications.builds.read**: allows your app to read build data for a user's applications
 - **applications.builds.upload**: allows your app to upload/update builds for a user's applications - requires Discord approval
 - **applications.commands**: allows your app to use commands in a guild
 - **applications.commands.permissions.update**: allows your app to update permissions for its commands in a guild a user has permissions to
 - **applications.commands.update**: allows your app to update its commands using a Bearer token - client credentials grant only
 - **applications.entitlements**: allows your app to read entitlements for a user's applications
 - **applications.store.update**: allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications
 - **bot**: for oauth2 bots, this puts the bot in the user's selected guild by default
 - **connections**: allows /users/@me/connections to return linked third-party accounts
 - **dm_channels.read**: allows your app to see information about the user's DMs and group DMs - requires Discord approval
 - **email**: enables /users/@me to return an email
 - **gdm.join**: allows your app to join users to a group dm
 - **guilds**: allows /users/@me/guilds to return basic information about all of a user's guilds
 - **guilds.join**: allows /guilds/{guild.id}/members/{user.id} to be used for joining users to a guild
 - **guilds.members.read**: allows /users/@me/guilds/{guild.id}/member to return a user's member information in a guild
 - **identify**: allows /users/@me without email
 - **messages.read**: for local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates)
 - **openid**: for OpenID Connect, this allows your app to receive user id and basic profile information
 - **relationships.read**: allows your app to know a user's friends and implicit relationships - requires Discord approval
 - **rpc**: for local rpc server access, this allows you to control a user's local Discord client - requires Discord approval
 - **rpc.activities.write**: for local rpc server access, this allows you to update a user's activity - requires Discord approval
 - **rpc.notifications.read**: for local rpc server access, this allows you to receive notifications pushed out to the user - requires Discord approval
 - **rpc.screenshare.read**: for local rpc server access, this allows you to read a user's screenshare status- requires Discord approval
 - **rpc.screenshare.write**: for local rpc server access, this allows you to update a user's screenshare settings- requires Discord approval
 - **rpc.video.read**: for local rpc server access, this allows you to read a user's video status - requires Discord approval
 - **rpc.video.write**: for local rpc server access, this allows you to update a user's video settings - requires Discord approval
 - **rpc.voice.read**: for local rpc server access, this allows you to read a user's voice settings and listen for voice events - requires Discord approval
 - **rpc.voice.write**: for local rpc server access, this allows you to update a user's voice settings - requires Discord approval
 - **voice**: allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval
 - **webhook.incoming**: this generates a webhook that is returned in the oauth token response for authorization code grants

Example

```go
auth := context.WithValue(context.Background(), dc_rest.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, dc_rest.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### OAuth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://discord.com/api/oauth2/authorize
- **Scopes**: 
 - **activities.invites.write**: allows your app to send activity invites - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
 - **activities.read**: allows your app to fetch data from a user's \"Now Playing/Recently Played\" list - requires Discord approval
 - **activities.write**: allows your app to update a user's activity - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
 - **applications.builds.read**: allows your app to read build data for a user's applications
 - **applications.builds.upload**: allows your app to upload/update builds for a user's applications - requires Discord approval
 - **applications.commands**: allows your app to use commands in a guild
 - **applications.commands.permissions.update**: allows your app to update permissions for its commands in a guild a user has permissions to
 - **applications.entitlements**: allows your app to read entitlements for a user's applications
 - **applications.store.update**: allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications
 - **bot**: for oauth2 bots, this puts the bot in the user's selected guild by default
 - **connections**: allows /users/@me/connections to return linked third-party accounts
 - **dm_channels.read**: allows your app to see information about the user's DMs and group DMs - requires Discord approval
 - **email**: enables /users/@me to return an email
 - **gdm.join**: allows your app to join users to a group dm
 - **guilds**: allows /users/@me/guilds to return basic information about all of a user's guilds
 - **guilds.join**: allows /guilds/{guild.id}/members/{user.id} to be used for joining users to a guild
 - **guilds.members.read**: allows /users/@me/guilds/{guild.id}/member to return a user's member information in a guild
 - **identify**: allows /users/@me without email
 - **messages.read**: for local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates)
 - **openid**: for OpenID Connect, this allows your app to receive user id and basic profile information
 - **relationships.read**: allows your app to know a user's friends and implicit relationships - requires Discord approval
 - **role_connections.write**: allows your app to update a user's connection and metadata for the app
 - **rpc**: for local rpc server access, this allows you to control a user's local Discord client - requires Discord approval
 - **rpc.activities.write**: for local rpc server access, this allows you to update a user's activity - requires Discord approval
 - **rpc.notifications.read**: for local rpc server access, this allows you to receive notifications pushed out to the user - requires Discord approval
 - **rpc.screenshare.read**: for local rpc server access, this allows you to read a user's screenshare status- requires Discord approval
 - **rpc.screenshare.write**: for local rpc server access, this allows you to update a user's screenshare settings- requires Discord approval
 - **rpc.video.read**: for local rpc server access, this allows you to read a user's video status - requires Discord approval
 - **rpc.video.write**: for local rpc server access, this allows you to update a user's video settings - requires Discord approval
 - **rpc.voice.read**: for local rpc server access, this allows you to read a user's voice settings and listen for voice events - requires Discord approval
 - **rpc.voice.write**: for local rpc server access, this allows you to update a user's voice settings - requires Discord approval
 - **voice**: allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval
 - **webhook.incoming**: this generates a webhook that is returned in the oauth token response for authorization code grants

Example

```go
auth := context.WithValue(context.Background(), dc_rest.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, dc_rest.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



