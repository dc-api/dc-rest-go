# \DefaultAPI

All URIs are relative to *https://discord.com/api/v10*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupDmUser**](DefaultAPI.md#AddGroupDmUser) | **Put** /channels/{channel_id}/recipients/{user_id} | 
[**AddGuildMember**](DefaultAPI.md#AddGuildMember) | **Put** /guilds/{guild_id}/members/{user_id} | 
[**AddGuildMemberRole**](DefaultAPI.md#AddGuildMemberRole) | **Put** /guilds/{guild_id}/members/{user_id}/roles/{role_id} | 
[**AddLobbyMember**](DefaultAPI.md#AddLobbyMember) | **Put** /lobbies/{lobby_id}/members/{user_id} | 
[**AddMyMessageReaction**](DefaultAPI.md#AddMyMessageReaction) | **Put** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me | 
[**AddThreadMember**](DefaultAPI.md#AddThreadMember) | **Put** /channels/{channel_id}/thread-members/{user_id} | 
[**ApplicationsGetActivityInstance**](DefaultAPI.md#ApplicationsGetActivityInstance) | **Get** /applications/{application_id}/activity-instances/{instance_id} | 
[**BanUserFromGuild**](DefaultAPI.md#BanUserFromGuild) | **Put** /guilds/{guild_id}/bans/{user_id} | 
[**BotPartnerSdkToken**](DefaultAPI.md#BotPartnerSdkToken) | **Post** /partner-sdk/token/bot | 
[**BulkBanUsersFromGuild**](DefaultAPI.md#BulkBanUsersFromGuild) | **Post** /guilds/{guild_id}/bulk-ban | 
[**BulkDeleteMessages**](DefaultAPI.md#BulkDeleteMessages) | **Post** /channels/{channel_id}/messages/bulk-delete | 
[**BulkSetApplicationCommands**](DefaultAPI.md#BulkSetApplicationCommands) | **Put** /applications/{application_id}/commands | 
[**BulkSetGuildApplicationCommands**](DefaultAPI.md#BulkSetGuildApplicationCommands) | **Put** /applications/{application_id}/guilds/{guild_id}/commands | 
[**BulkUpdateGuildChannels**](DefaultAPI.md#BulkUpdateGuildChannels) | **Patch** /guilds/{guild_id}/channels | 
[**BulkUpdateGuildRoles**](DefaultAPI.md#BulkUpdateGuildRoles) | **Patch** /guilds/{guild_id}/roles | 
[**BulkUpdateLobbyMembers**](DefaultAPI.md#BulkUpdateLobbyMembers) | **Post** /lobbies/{lobby_id}/members/bulk | 
[**ConsumeEntitlement**](DefaultAPI.md#ConsumeEntitlement) | **Post** /applications/{application_id}/entitlements/{entitlement_id}/consume | 
[**CreateApplicationCommand**](DefaultAPI.md#CreateApplicationCommand) | **Post** /applications/{application_id}/commands | 
[**CreateApplicationEmoji**](DefaultAPI.md#CreateApplicationEmoji) | **Post** /applications/{application_id}/emojis | 
[**CreateAutoModerationRule**](DefaultAPI.md#CreateAutoModerationRule) | **Post** /guilds/{guild_id}/auto-moderation/rules | 
[**CreateChannelInvite**](DefaultAPI.md#CreateChannelInvite) | **Post** /channels/{channel_id}/invites | 
[**CreateDm**](DefaultAPI.md#CreateDm) | **Post** /users/@me/channels | 
[**CreateEntitlement**](DefaultAPI.md#CreateEntitlement) | **Post** /applications/{application_id}/entitlements | 
[**CreateGuild**](DefaultAPI.md#CreateGuild) | **Post** /guilds | 
[**CreateGuildApplicationCommand**](DefaultAPI.md#CreateGuildApplicationCommand) | **Post** /applications/{application_id}/guilds/{guild_id}/commands | 
[**CreateGuildChannel**](DefaultAPI.md#CreateGuildChannel) | **Post** /guilds/{guild_id}/channels | 
[**CreateGuildEmoji**](DefaultAPI.md#CreateGuildEmoji) | **Post** /guilds/{guild_id}/emojis | 
[**CreateGuildFromTemplate**](DefaultAPI.md#CreateGuildFromTemplate) | **Post** /guilds/templates/{code} | 
[**CreateGuildRole**](DefaultAPI.md#CreateGuildRole) | **Post** /guilds/{guild_id}/roles | 
[**CreateGuildScheduledEvent**](DefaultAPI.md#CreateGuildScheduledEvent) | **Post** /guilds/{guild_id}/scheduled-events | 
[**CreateGuildSoundboardSound**](DefaultAPI.md#CreateGuildSoundboardSound) | **Post** /guilds/{guild_id}/soundboard-sounds | 
[**CreateGuildSticker**](DefaultAPI.md#CreateGuildSticker) | **Post** /guilds/{guild_id}/stickers | 
[**CreateGuildTemplate**](DefaultAPI.md#CreateGuildTemplate) | **Post** /guilds/{guild_id}/templates | 
[**CreateInteractionResponse**](DefaultAPI.md#CreateInteractionResponse) | **Post** /interactions/{interaction_id}/{interaction_token}/callback | 
[**CreateLinkedLobbyGuildInviteForSelf**](DefaultAPI.md#CreateLinkedLobbyGuildInviteForSelf) | **Post** /lobbies/{lobby_id}/members/@me/invites | 
[**CreateLinkedLobbyGuildInviteForUser**](DefaultAPI.md#CreateLinkedLobbyGuildInviteForUser) | **Post** /lobbies/{lobby_id}/members/{user_id}/invites | 
[**CreateLobby**](DefaultAPI.md#CreateLobby) | **Post** /lobbies | 
[**CreateLobbyMessage**](DefaultAPI.md#CreateLobbyMessage) | **Post** /lobbies/{lobby_id}/messages | 
[**CreateMessage**](DefaultAPI.md#CreateMessage) | **Post** /channels/{channel_id}/messages | 
[**CreateOrJoinLobby**](DefaultAPI.md#CreateOrJoinLobby) | **Put** /lobbies | 
[**CreatePin**](DefaultAPI.md#CreatePin) | **Put** /channels/{channel_id}/messages/pins/{message_id} | 
[**CreateStageInstance**](DefaultAPI.md#CreateStageInstance) | **Post** /stage-instances | 
[**CreateThread**](DefaultAPI.md#CreateThread) | **Post** /channels/{channel_id}/threads | 
[**CreateThreadFromMessage**](DefaultAPI.md#CreateThreadFromMessage) | **Post** /channels/{channel_id}/messages/{message_id}/threads | 
[**CreateWebhook**](DefaultAPI.md#CreateWebhook) | **Post** /channels/{channel_id}/webhooks | 
[**CrosspostMessage**](DefaultAPI.md#CrosspostMessage) | **Post** /channels/{channel_id}/messages/{message_id}/crosspost | 
[**DeleteAllMessageReactions**](DefaultAPI.md#DeleteAllMessageReactions) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions | 
[**DeleteAllMessageReactionsByEmoji**](DefaultAPI.md#DeleteAllMessageReactionsByEmoji) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name} | 
[**DeleteApplicationCommand**](DefaultAPI.md#DeleteApplicationCommand) | **Delete** /applications/{application_id}/commands/{command_id} | 
[**DeleteApplicationEmoji**](DefaultAPI.md#DeleteApplicationEmoji) | **Delete** /applications/{application_id}/emojis/{emoji_id} | 
[**DeleteApplicationUserRoleConnection**](DefaultAPI.md#DeleteApplicationUserRoleConnection) | **Delete** /users/@me/applications/{application_id}/role-connection | 
[**DeleteAutoModerationRule**](DefaultAPI.md#DeleteAutoModerationRule) | **Delete** /guilds/{guild_id}/auto-moderation/rules/{rule_id} | 
[**DeleteChannel**](DefaultAPI.md#DeleteChannel) | **Delete** /channels/{channel_id} | 
[**DeleteChannelPermissionOverwrite**](DefaultAPI.md#DeleteChannelPermissionOverwrite) | **Delete** /channels/{channel_id}/permissions/{overwrite_id} | 
[**DeleteEntitlement**](DefaultAPI.md#DeleteEntitlement) | **Delete** /applications/{application_id}/entitlements/{entitlement_id} | 
[**DeleteGroupDmUser**](DefaultAPI.md#DeleteGroupDmUser) | **Delete** /channels/{channel_id}/recipients/{user_id} | 
[**DeleteGuild**](DefaultAPI.md#DeleteGuild) | **Delete** /guilds/{guild_id} | 
[**DeleteGuildApplicationCommand**](DefaultAPI.md#DeleteGuildApplicationCommand) | **Delete** /applications/{application_id}/guilds/{guild_id}/commands/{command_id} | 
[**DeleteGuildEmoji**](DefaultAPI.md#DeleteGuildEmoji) | **Delete** /guilds/{guild_id}/emojis/{emoji_id} | 
[**DeleteGuildIntegration**](DefaultAPI.md#DeleteGuildIntegration) | **Delete** /guilds/{guild_id}/integrations/{integration_id} | 
[**DeleteGuildMember**](DefaultAPI.md#DeleteGuildMember) | **Delete** /guilds/{guild_id}/members/{user_id} | 
[**DeleteGuildMemberRole**](DefaultAPI.md#DeleteGuildMemberRole) | **Delete** /guilds/{guild_id}/members/{user_id}/roles/{role_id} | 
[**DeleteGuildRole**](DefaultAPI.md#DeleteGuildRole) | **Delete** /guilds/{guild_id}/roles/{role_id} | 
[**DeleteGuildScheduledEvent**](DefaultAPI.md#DeleteGuildScheduledEvent) | **Delete** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id} | 
[**DeleteGuildSoundboardSound**](DefaultAPI.md#DeleteGuildSoundboardSound) | **Delete** /guilds/{guild_id}/soundboard-sounds/{sound_id} | 
[**DeleteGuildSticker**](DefaultAPI.md#DeleteGuildSticker) | **Delete** /guilds/{guild_id}/stickers/{sticker_id} | 
[**DeleteGuildTemplate**](DefaultAPI.md#DeleteGuildTemplate) | **Delete** /guilds/{guild_id}/templates/{code} | 
[**DeleteLobbyMember**](DefaultAPI.md#DeleteLobbyMember) | **Delete** /lobbies/{lobby_id}/members/{user_id} | 
[**DeleteMessage**](DefaultAPI.md#DeleteMessage) | **Delete** /channels/{channel_id}/messages/{message_id} | 
[**DeleteMyMessageReaction**](DefaultAPI.md#DeleteMyMessageReaction) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me | 
[**DeleteOriginalWebhookMessage**](DefaultAPI.md#DeleteOriginalWebhookMessage) | **Delete** /webhooks/{webhook_id}/{webhook_token}/messages/@original | 
[**DeletePin**](DefaultAPI.md#DeletePin) | **Delete** /channels/{channel_id}/messages/pins/{message_id} | 
[**DeleteStageInstance**](DefaultAPI.md#DeleteStageInstance) | **Delete** /stage-instances/{channel_id} | 
[**DeleteThreadMember**](DefaultAPI.md#DeleteThreadMember) | **Delete** /channels/{channel_id}/thread-members/{user_id} | 
[**DeleteUserMessageReaction**](DefaultAPI.md#DeleteUserMessageReaction) | **Delete** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/{user_id} | 
[**DeleteWebhook**](DefaultAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | 
[**DeleteWebhookByToken**](DefaultAPI.md#DeleteWebhookByToken) | **Delete** /webhooks/{webhook_id}/{webhook_token} | 
[**DeleteWebhookMessage**](DefaultAPI.md#DeleteWebhookMessage) | **Delete** /webhooks/{webhook_id}/{webhook_token}/messages/{message_id} | 
[**DeprecatedCreatePin**](DefaultAPI.md#DeprecatedCreatePin) | **Put** /channels/{channel_id}/pins/{message_id} | 
[**DeprecatedDeletePin**](DefaultAPI.md#DeprecatedDeletePin) | **Delete** /channels/{channel_id}/pins/{message_id} | 
[**DeprecatedListPins**](DefaultAPI.md#DeprecatedListPins) | **Get** /channels/{channel_id}/pins | 
[**EditLobby**](DefaultAPI.md#EditLobby) | **Patch** /lobbies/{lobby_id} | 
[**EditLobbyChannelLink**](DefaultAPI.md#EditLobbyChannelLink) | **Patch** /lobbies/{lobby_id}/channel-linking | 
[**ExecuteGithubCompatibleWebhook**](DefaultAPI.md#ExecuteGithubCompatibleWebhook) | **Post** /webhooks/{webhook_id}/{webhook_token}/github | 
[**ExecuteSlackCompatibleWebhook**](DefaultAPI.md#ExecuteSlackCompatibleWebhook) | **Post** /webhooks/{webhook_id}/{webhook_token}/slack | 
[**ExecuteWebhook**](DefaultAPI.md#ExecuteWebhook) | **Post** /webhooks/{webhook_id}/{webhook_token} | 
[**FollowChannel**](DefaultAPI.md#FollowChannel) | **Post** /channels/{channel_id}/followers | 
[**GetActiveGuildThreads**](DefaultAPI.md#GetActiveGuildThreads) | **Get** /guilds/{guild_id}/threads/active | 
[**GetAnswerVoters**](DefaultAPI.md#GetAnswerVoters) | **Get** /channels/{channel_id}/polls/{message_id}/answers/{answer_id} | 
[**GetApplication**](DefaultAPI.md#GetApplication) | **Get** /applications/{application_id} | 
[**GetApplicationCommand**](DefaultAPI.md#GetApplicationCommand) | **Get** /applications/{application_id}/commands/{command_id} | 
[**GetApplicationEmoji**](DefaultAPI.md#GetApplicationEmoji) | **Get** /applications/{application_id}/emojis/{emoji_id} | 
[**GetApplicationRoleConnectionsMetadata**](DefaultAPI.md#GetApplicationRoleConnectionsMetadata) | **Get** /applications/{application_id}/role-connections/metadata | 
[**GetApplicationUserRoleConnection**](DefaultAPI.md#GetApplicationUserRoleConnection) | **Get** /users/@me/applications/{application_id}/role-connection | 
[**GetAutoModerationRule**](DefaultAPI.md#GetAutoModerationRule) | **Get** /guilds/{guild_id}/auto-moderation/rules/{rule_id} | 
[**GetBotGateway**](DefaultAPI.md#GetBotGateway) | **Get** /gateway/bot | 
[**GetChannel**](DefaultAPI.md#GetChannel) | **Get** /channels/{channel_id} | 
[**GetEntitlement**](DefaultAPI.md#GetEntitlement) | **Get** /applications/{application_id}/entitlements/{entitlement_id} | 
[**GetEntitlements**](DefaultAPI.md#GetEntitlements) | **Get** /applications/{application_id}/entitlements | 
[**GetGateway**](DefaultAPI.md#GetGateway) | **Get** /gateway | 
[**GetGuild**](DefaultAPI.md#GetGuild) | **Get** /guilds/{guild_id} | 
[**GetGuildApplicationCommand**](DefaultAPI.md#GetGuildApplicationCommand) | **Get** /applications/{application_id}/guilds/{guild_id}/commands/{command_id} | 
[**GetGuildApplicationCommandPermissions**](DefaultAPI.md#GetGuildApplicationCommandPermissions) | **Get** /applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions | 
[**GetGuildBan**](DefaultAPI.md#GetGuildBan) | **Get** /guilds/{guild_id}/bans/{user_id} | 
[**GetGuildEmoji**](DefaultAPI.md#GetGuildEmoji) | **Get** /guilds/{guild_id}/emojis/{emoji_id} | 
[**GetGuildMember**](DefaultAPI.md#GetGuildMember) | **Get** /guilds/{guild_id}/members/{user_id} | 
[**GetGuildNewMemberWelcome**](DefaultAPI.md#GetGuildNewMemberWelcome) | **Get** /guilds/{guild_id}/new-member-welcome | 
[**GetGuildPreview**](DefaultAPI.md#GetGuildPreview) | **Get** /guilds/{guild_id}/preview | 
[**GetGuildRole**](DefaultAPI.md#GetGuildRole) | **Get** /guilds/{guild_id}/roles/{role_id} | 
[**GetGuildScheduledEvent**](DefaultAPI.md#GetGuildScheduledEvent) | **Get** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id} | 
[**GetGuildSoundboardSound**](DefaultAPI.md#GetGuildSoundboardSound) | **Get** /guilds/{guild_id}/soundboard-sounds/{sound_id} | 
[**GetGuildSticker**](DefaultAPI.md#GetGuildSticker) | **Get** /guilds/{guild_id}/stickers/{sticker_id} | 
[**GetGuildTemplate**](DefaultAPI.md#GetGuildTemplate) | **Get** /guilds/templates/{code} | 
[**GetGuildVanityUrl**](DefaultAPI.md#GetGuildVanityUrl) | **Get** /guilds/{guild_id}/vanity-url | 
[**GetGuildWebhooks**](DefaultAPI.md#GetGuildWebhooks) | **Get** /guilds/{guild_id}/webhooks | 
[**GetGuildWelcomeScreen**](DefaultAPI.md#GetGuildWelcomeScreen) | **Get** /guilds/{guild_id}/welcome-screen | 
[**GetGuildWidget**](DefaultAPI.md#GetGuildWidget) | **Get** /guilds/{guild_id}/widget.json | 
[**GetGuildWidgetPng**](DefaultAPI.md#GetGuildWidgetPng) | **Get** /guilds/{guild_id}/widget.png | 
[**GetGuildWidgetSettings**](DefaultAPI.md#GetGuildWidgetSettings) | **Get** /guilds/{guild_id}/widget | 
[**GetGuildsOnboarding**](DefaultAPI.md#GetGuildsOnboarding) | **Get** /guilds/{guild_id}/onboarding | 
[**GetLobby**](DefaultAPI.md#GetLobby) | **Get** /lobbies/{lobby_id} | 
[**GetLobbyMessages**](DefaultAPI.md#GetLobbyMessages) | **Get** /lobbies/{lobby_id}/messages | 
[**GetMessage**](DefaultAPI.md#GetMessage) | **Get** /channels/{channel_id}/messages/{message_id} | 
[**GetMyApplication**](DefaultAPI.md#GetMyApplication) | **Get** /applications/@me | 
[**GetMyGuildMember**](DefaultAPI.md#GetMyGuildMember) | **Get** /users/@me/guilds/{guild_id}/member | 
[**GetMyOauth2Application**](DefaultAPI.md#GetMyOauth2Application) | **Get** /oauth2/applications/@me | 
[**GetMyOauth2Authorization**](DefaultAPI.md#GetMyOauth2Authorization) | **Get** /oauth2/@me | 
[**GetMyUser**](DefaultAPI.md#GetMyUser) | **Get** /users/@me | 
[**GetOpenidConnectUserinfo**](DefaultAPI.md#GetOpenidConnectUserinfo) | **Get** /oauth2/userinfo | 
[**GetOriginalWebhookMessage**](DefaultAPI.md#GetOriginalWebhookMessage) | **Get** /webhooks/{webhook_id}/{webhook_token}/messages/@original | 
[**GetPublicKeys**](DefaultAPI.md#GetPublicKeys) | **Get** /oauth2/keys | 
[**GetSelfVoiceState**](DefaultAPI.md#GetSelfVoiceState) | **Get** /guilds/{guild_id}/voice-states/@me | 
[**GetSoundboardDefaultSounds**](DefaultAPI.md#GetSoundboardDefaultSounds) | **Get** /soundboard-default-sounds | 
[**GetStageInstance**](DefaultAPI.md#GetStageInstance) | **Get** /stage-instances/{channel_id} | 
[**GetSticker**](DefaultAPI.md#GetSticker) | **Get** /stickers/{sticker_id} | 
[**GetStickerPack**](DefaultAPI.md#GetStickerPack) | **Get** /sticker-packs/{pack_id} | 
[**GetThreadMember**](DefaultAPI.md#GetThreadMember) | **Get** /channels/{channel_id}/thread-members/{user_id} | 
[**GetUser**](DefaultAPI.md#GetUser) | **Get** /users/{user_id} | 
[**GetVoiceState**](DefaultAPI.md#GetVoiceState) | **Get** /guilds/{guild_id}/voice-states/{user_id} | 
[**GetWebhook**](DefaultAPI.md#GetWebhook) | **Get** /webhooks/{webhook_id} | 
[**GetWebhookByToken**](DefaultAPI.md#GetWebhookByToken) | **Get** /webhooks/{webhook_id}/{webhook_token} | 
[**GetWebhookMessage**](DefaultAPI.md#GetWebhookMessage) | **Get** /webhooks/{webhook_id}/{webhook_token}/messages/{message_id} | 
[**InviteResolve**](DefaultAPI.md#InviteResolve) | **Get** /invites/{code} | 
[**InviteRevoke**](DefaultAPI.md#InviteRevoke) | **Delete** /invites/{code} | 
[**JoinThread**](DefaultAPI.md#JoinThread) | **Put** /channels/{channel_id}/thread-members/@me | 
[**LeaveGuild**](DefaultAPI.md#LeaveGuild) | **Delete** /users/@me/guilds/{guild_id} | 
[**LeaveLobby**](DefaultAPI.md#LeaveLobby) | **Delete** /lobbies/{lobby_id}/members/@me | 
[**LeaveThread**](DefaultAPI.md#LeaveThread) | **Delete** /channels/{channel_id}/thread-members/@me | 
[**ListApplicationCommands**](DefaultAPI.md#ListApplicationCommands) | **Get** /applications/{application_id}/commands | 
[**ListApplicationEmojis**](DefaultAPI.md#ListApplicationEmojis) | **Get** /applications/{application_id}/emojis | 
[**ListAutoModerationRules**](DefaultAPI.md#ListAutoModerationRules) | **Get** /guilds/{guild_id}/auto-moderation/rules | 
[**ListChannelInvites**](DefaultAPI.md#ListChannelInvites) | **Get** /channels/{channel_id}/invites | 
[**ListChannelWebhooks**](DefaultAPI.md#ListChannelWebhooks) | **Get** /channels/{channel_id}/webhooks | 
[**ListGuildApplicationCommandPermissions**](DefaultAPI.md#ListGuildApplicationCommandPermissions) | **Get** /applications/{application_id}/guilds/{guild_id}/commands/permissions | 
[**ListGuildApplicationCommands**](DefaultAPI.md#ListGuildApplicationCommands) | **Get** /applications/{application_id}/guilds/{guild_id}/commands | 
[**ListGuildAuditLogEntries**](DefaultAPI.md#ListGuildAuditLogEntries) | **Get** /guilds/{guild_id}/audit-logs | 
[**ListGuildBans**](DefaultAPI.md#ListGuildBans) | **Get** /guilds/{guild_id}/bans | 
[**ListGuildChannels**](DefaultAPI.md#ListGuildChannels) | **Get** /guilds/{guild_id}/channels | 
[**ListGuildEmojis**](DefaultAPI.md#ListGuildEmojis) | **Get** /guilds/{guild_id}/emojis | 
[**ListGuildIntegrations**](DefaultAPI.md#ListGuildIntegrations) | **Get** /guilds/{guild_id}/integrations | 
[**ListGuildInvites**](DefaultAPI.md#ListGuildInvites) | **Get** /guilds/{guild_id}/invites | 
[**ListGuildMembers**](DefaultAPI.md#ListGuildMembers) | **Get** /guilds/{guild_id}/members | 
[**ListGuildRoles**](DefaultAPI.md#ListGuildRoles) | **Get** /guilds/{guild_id}/roles | 
[**ListGuildScheduledEventUsers**](DefaultAPI.md#ListGuildScheduledEventUsers) | **Get** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users | 
[**ListGuildScheduledEvents**](DefaultAPI.md#ListGuildScheduledEvents) | **Get** /guilds/{guild_id}/scheduled-events | 
[**ListGuildSoundboardSounds**](DefaultAPI.md#ListGuildSoundboardSounds) | **Get** /guilds/{guild_id}/soundboard-sounds | 
[**ListGuildStickers**](DefaultAPI.md#ListGuildStickers) | **Get** /guilds/{guild_id}/stickers | 
[**ListGuildTemplates**](DefaultAPI.md#ListGuildTemplates) | **Get** /guilds/{guild_id}/templates | 
[**ListGuildVoiceRegions**](DefaultAPI.md#ListGuildVoiceRegions) | **Get** /guilds/{guild_id}/regions | 
[**ListMessageReactionsByEmoji**](DefaultAPI.md#ListMessageReactionsByEmoji) | **Get** /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name} | 
[**ListMessages**](DefaultAPI.md#ListMessages) | **Get** /channels/{channel_id}/messages | 
[**ListMyConnections**](DefaultAPI.md#ListMyConnections) | **Get** /users/@me/connections | 
[**ListMyGuilds**](DefaultAPI.md#ListMyGuilds) | **Get** /users/@me/guilds | 
[**ListMyPrivateArchivedThreads**](DefaultAPI.md#ListMyPrivateArchivedThreads) | **Get** /channels/{channel_id}/users/@me/threads/archived/private | 
[**ListPins**](DefaultAPI.md#ListPins) | **Get** /channels/{channel_id}/messages/pins | 
[**ListPrivateArchivedThreads**](DefaultAPI.md#ListPrivateArchivedThreads) | **Get** /channels/{channel_id}/threads/archived/private | 
[**ListPublicArchivedThreads**](DefaultAPI.md#ListPublicArchivedThreads) | **Get** /channels/{channel_id}/threads/archived/public | 
[**ListStickerPacks**](DefaultAPI.md#ListStickerPacks) | **Get** /sticker-packs | 
[**ListThreadMembers**](DefaultAPI.md#ListThreadMembers) | **Get** /channels/{channel_id}/thread-members | 
[**ListVoiceRegions**](DefaultAPI.md#ListVoiceRegions) | **Get** /voice/regions | 
[**PartnerSdkToken**](DefaultAPI.md#PartnerSdkToken) | **Post** /partner-sdk/token | 
[**PartnerSdkUnmergeProvisionalAccount**](DefaultAPI.md#PartnerSdkUnmergeProvisionalAccount) | **Post** /partner-sdk/provisional-accounts/unmerge | 
[**PollExpire**](DefaultAPI.md#PollExpire) | **Post** /channels/{channel_id}/polls/{message_id}/expire | 
[**PreviewPruneGuild**](DefaultAPI.md#PreviewPruneGuild) | **Get** /guilds/{guild_id}/prune | 
[**PruneGuild**](DefaultAPI.md#PruneGuild) | **Post** /guilds/{guild_id}/prune | 
[**PutGuildsOnboarding**](DefaultAPI.md#PutGuildsOnboarding) | **Put** /guilds/{guild_id}/onboarding | 
[**SearchGuildMembers**](DefaultAPI.md#SearchGuildMembers) | **Get** /guilds/{guild_id}/members/search | 
[**SendSoundboardSound**](DefaultAPI.md#SendSoundboardSound) | **Post** /channels/{channel_id}/send-soundboard-sound | 
[**SetChannelPermissionOverwrite**](DefaultAPI.md#SetChannelPermissionOverwrite) | **Put** /channels/{channel_id}/permissions/{overwrite_id} | 
[**SetGuildApplicationCommandPermissions**](DefaultAPI.md#SetGuildApplicationCommandPermissions) | **Put** /applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions | 
[**SetGuildMfaLevel**](DefaultAPI.md#SetGuildMfaLevel) | **Post** /guilds/{guild_id}/mfa | 
[**SyncGuildTemplate**](DefaultAPI.md#SyncGuildTemplate) | **Put** /guilds/{guild_id}/templates/{code} | 
[**ThreadSearch**](DefaultAPI.md#ThreadSearch) | **Get** /channels/{channel_id}/threads/search | 
[**TriggerTypingIndicator**](DefaultAPI.md#TriggerTypingIndicator) | **Post** /channels/{channel_id}/typing | 
[**UnbanUserFromGuild**](DefaultAPI.md#UnbanUserFromGuild) | **Delete** /guilds/{guild_id}/bans/{user_id} | 
[**UpdateApplication**](DefaultAPI.md#UpdateApplication) | **Patch** /applications/{application_id} | 
[**UpdateApplicationCommand**](DefaultAPI.md#UpdateApplicationCommand) | **Patch** /applications/{application_id}/commands/{command_id} | 
[**UpdateApplicationEmoji**](DefaultAPI.md#UpdateApplicationEmoji) | **Patch** /applications/{application_id}/emojis/{emoji_id} | 
[**UpdateApplicationRoleConnectionsMetadata**](DefaultAPI.md#UpdateApplicationRoleConnectionsMetadata) | **Put** /applications/{application_id}/role-connections/metadata | 
[**UpdateApplicationUserRoleConnection**](DefaultAPI.md#UpdateApplicationUserRoleConnection) | **Put** /users/@me/applications/{application_id}/role-connection | 
[**UpdateAutoModerationRule**](DefaultAPI.md#UpdateAutoModerationRule) | **Patch** /guilds/{guild_id}/auto-moderation/rules/{rule_id} | 
[**UpdateChannel**](DefaultAPI.md#UpdateChannel) | **Patch** /channels/{channel_id} | 
[**UpdateGuild**](DefaultAPI.md#UpdateGuild) | **Patch** /guilds/{guild_id} | 
[**UpdateGuildApplicationCommand**](DefaultAPI.md#UpdateGuildApplicationCommand) | **Patch** /applications/{application_id}/guilds/{guild_id}/commands/{command_id} | 
[**UpdateGuildEmoji**](DefaultAPI.md#UpdateGuildEmoji) | **Patch** /guilds/{guild_id}/emojis/{emoji_id} | 
[**UpdateGuildMember**](DefaultAPI.md#UpdateGuildMember) | **Patch** /guilds/{guild_id}/members/{user_id} | 
[**UpdateGuildRole**](DefaultAPI.md#UpdateGuildRole) | **Patch** /guilds/{guild_id}/roles/{role_id} | 
[**UpdateGuildScheduledEvent**](DefaultAPI.md#UpdateGuildScheduledEvent) | **Patch** /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id} | 
[**UpdateGuildSoundboardSound**](DefaultAPI.md#UpdateGuildSoundboardSound) | **Patch** /guilds/{guild_id}/soundboard-sounds/{sound_id} | 
[**UpdateGuildSticker**](DefaultAPI.md#UpdateGuildSticker) | **Patch** /guilds/{guild_id}/stickers/{sticker_id} | 
[**UpdateGuildTemplate**](DefaultAPI.md#UpdateGuildTemplate) | **Patch** /guilds/{guild_id}/templates/{code} | 
[**UpdateGuildWelcomeScreen**](DefaultAPI.md#UpdateGuildWelcomeScreen) | **Patch** /guilds/{guild_id}/welcome-screen | 
[**UpdateGuildWidgetSettings**](DefaultAPI.md#UpdateGuildWidgetSettings) | **Patch** /guilds/{guild_id}/widget | 
[**UpdateMessage**](DefaultAPI.md#UpdateMessage) | **Patch** /channels/{channel_id}/messages/{message_id} | 
[**UpdateMyApplication**](DefaultAPI.md#UpdateMyApplication) | **Patch** /applications/@me | 
[**UpdateMyGuildMember**](DefaultAPI.md#UpdateMyGuildMember) | **Patch** /guilds/{guild_id}/members/@me | 
[**UpdateMyUser**](DefaultAPI.md#UpdateMyUser) | **Patch** /users/@me | 
[**UpdateOriginalWebhookMessage**](DefaultAPI.md#UpdateOriginalWebhookMessage) | **Patch** /webhooks/{webhook_id}/{webhook_token}/messages/@original | 
[**UpdateSelfVoiceState**](DefaultAPI.md#UpdateSelfVoiceState) | **Patch** /guilds/{guild_id}/voice-states/@me | 
[**UpdateStageInstance**](DefaultAPI.md#UpdateStageInstance) | **Patch** /stage-instances/{channel_id} | 
[**UpdateVoiceState**](DefaultAPI.md#UpdateVoiceState) | **Patch** /guilds/{guild_id}/voice-states/{user_id} | 
[**UpdateWebhook**](DefaultAPI.md#UpdateWebhook) | **Patch** /webhooks/{webhook_id} | 
[**UpdateWebhookByToken**](DefaultAPI.md#UpdateWebhookByToken) | **Patch** /webhooks/{webhook_id}/{webhook_token} | 
[**UpdateWebhookMessage**](DefaultAPI.md#UpdateWebhookMessage) | **Patch** /webhooks/{webhook_id}/{webhook_token}/messages/{message_id} | 
[**UploadApplicationAttachment**](DefaultAPI.md#UploadApplicationAttachment) | **Post** /applications/{application_id}/attachment | 



## AddGroupDmUser

> AddGroupDmUser201Response AddGroupDmUser(ctx, channelId, userId).AddGroupDmUserRequest(addGroupDmUserRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	userId := "userId_example" // string | 
	addGroupDmUserRequest := *openapiclient.NewAddGroupDmUserRequest() // AddGroupDmUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddGroupDmUser(context.Background(), channelId, userId).AddGroupDmUserRequest(addGroupDmUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddGroupDmUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupDmUser`: AddGroupDmUser201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddGroupDmUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupDmUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addGroupDmUserRequest** | [**AddGroupDmUserRequest**](AddGroupDmUserRequest.md) |  | 

### Return type

[**AddGroupDmUser201Response**](AddGroupDmUser201Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGuildMember

> GuildMemberResponse AddGuildMember(ctx, guildId, userId).BotAddGuildMemberRequest(botAddGuildMemberRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	botAddGuildMemberRequest := *openapiclient.NewBotAddGuildMemberRequest("AccessToken_example") // BotAddGuildMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddGuildMember(context.Background(), guildId, userId).BotAddGuildMemberRequest(botAddGuildMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddGuildMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGuildMember`: GuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddGuildMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGuildMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **botAddGuildMemberRequest** | [**BotAddGuildMemberRequest**](BotAddGuildMemberRequest.md) |  | 

### Return type

[**GuildMemberResponse**](GuildMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGuildMemberRole

> AddGuildMemberRole(ctx, guildId, userId, roleId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AddGuildMemberRole(context.Background(), guildId, userId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddGuildMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGuildMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLobbyMember

> LobbyMemberResponse AddLobbyMember(ctx, lobbyId, userId).AddLobbyMemberRequest(addLobbyMemberRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	userId := "userId_example" // string | 
	addLobbyMemberRequest := *openapiclient.NewAddLobbyMemberRequest() // AddLobbyMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddLobbyMember(context.Background(), lobbyId, userId).AddLobbyMemberRequest(addLobbyMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddLobbyMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLobbyMember`: LobbyMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddLobbyMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLobbyMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addLobbyMemberRequest** | [**AddLobbyMemberRequest**](AddLobbyMemberRequest.md) |  | 

### Return type

[**LobbyMemberResponse**](LobbyMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMyMessageReaction

> AddMyMessageReaction(ctx, channelId, messageId, emojiName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	emojiName := "emojiName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AddMyMessageReaction(context.Background(), channelId, messageId, emojiName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddMyMessageReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 
**emojiName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMyMessageReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddThreadMember

> AddThreadMember(ctx, channelId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AddThreadMember(context.Background(), channelId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddThreadMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddThreadMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsGetActivityInstance

> EmbeddedActivityInstance ApplicationsGetActivityInstance(ctx, applicationId, instanceId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	instanceId := "instanceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApplicationsGetActivityInstance(context.Background(), applicationId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApplicationsGetActivityInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplicationsGetActivityInstance`: EmbeddedActivityInstance
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApplicationsGetActivityInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsGetActivityInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmbeddedActivityInstance**](EmbeddedActivityInstance.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BanUserFromGuild

> BanUserFromGuild(ctx, guildId, userId).BanUserFromGuildRequest(banUserFromGuildRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	banUserFromGuildRequest := *openapiclient.NewBanUserFromGuildRequest() // BanUserFromGuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.BanUserFromGuild(context.Background(), guildId, userId).BanUserFromGuildRequest(banUserFromGuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BanUserFromGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBanUserFromGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **banUserFromGuildRequest** | [**BanUserFromGuildRequest**](BanUserFromGuildRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotPartnerSdkToken

> ProvisionalTokenResponse BotPartnerSdkToken(ctx).BotPartnerSdkTokenRequest(botPartnerSdkTokenRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	botPartnerSdkTokenRequest := *openapiclient.NewBotPartnerSdkTokenRequest("ExternalUserId_example") // BotPartnerSdkTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BotPartnerSdkToken(context.Background()).BotPartnerSdkTokenRequest(botPartnerSdkTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BotPartnerSdkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotPartnerSdkToken`: ProvisionalTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BotPartnerSdkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotPartnerSdkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botPartnerSdkTokenRequest** | [**BotPartnerSdkTokenRequest**](BotPartnerSdkTokenRequest.md) |  | 

### Return type

[**ProvisionalTokenResponse**](ProvisionalTokenResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkBanUsersFromGuild

> BulkBanUsersResponse BulkBanUsersFromGuild(ctx, guildId).BulkBanUsersFromGuildRequest(bulkBanUsersFromGuildRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	bulkBanUsersFromGuildRequest := *openapiclient.NewBulkBanUsersFromGuildRequest([]string{"UserIds_example"}) // BulkBanUsersFromGuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BulkBanUsersFromGuild(context.Background(), guildId).BulkBanUsersFromGuildRequest(bulkBanUsersFromGuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkBanUsersFromGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkBanUsersFromGuild`: BulkBanUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkBanUsersFromGuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkBanUsersFromGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkBanUsersFromGuildRequest** | [**BulkBanUsersFromGuildRequest**](BulkBanUsersFromGuildRequest.md) |  | 

### Return type

[**BulkBanUsersResponse**](BulkBanUsersResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteMessages

> BulkDeleteMessages(ctx, channelId).BulkDeleteMessagesRequest(bulkDeleteMessagesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	bulkDeleteMessagesRequest := *openapiclient.NewBulkDeleteMessagesRequest([]string{"Messages_example"}) // BulkDeleteMessagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.BulkDeleteMessages(context.Background(), channelId).BulkDeleteMessagesRequest(bulkDeleteMessagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkDeleteMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkDeleteMessagesRequest** | [**BulkDeleteMessagesRequest**](BulkDeleteMessagesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSetApplicationCommands

> []ApplicationCommandResponse BulkSetApplicationCommands(ctx, applicationId).ApplicationCommandUpdateRequest(applicationCommandUpdateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	applicationCommandUpdateRequest := []openapiclient.ApplicationCommandUpdateRequest{*openapiclient.NewApplicationCommandUpdateRequest("Name_example")} // []ApplicationCommandUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BulkSetApplicationCommands(context.Background(), applicationId).ApplicationCommandUpdateRequest(applicationCommandUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkSetApplicationCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkSetApplicationCommands`: []ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkSetApplicationCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetApplicationCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationCommandUpdateRequest** | [**[]ApplicationCommandUpdateRequest**](ApplicationCommandUpdateRequest.md) |  | 

### Return type

[**[]ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSetGuildApplicationCommands

> []ApplicationCommandResponse BulkSetGuildApplicationCommands(ctx, applicationId, guildId).ApplicationCommandUpdateRequest(applicationCommandUpdateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	applicationCommandUpdateRequest := []openapiclient.ApplicationCommandUpdateRequest{*openapiclient.NewApplicationCommandUpdateRequest("Name_example")} // []ApplicationCommandUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BulkSetGuildApplicationCommands(context.Background(), applicationId, guildId).ApplicationCommandUpdateRequest(applicationCommandUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkSetGuildApplicationCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkSetGuildApplicationCommands`: []ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkSetGuildApplicationCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetGuildApplicationCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationCommandUpdateRequest** | [**[]ApplicationCommandUpdateRequest**](ApplicationCommandUpdateRequest.md) |  | 

### Return type

[**[]ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateGuildChannels

> BulkUpdateGuildChannels(ctx, guildId).BulkUpdateGuildChannelsRequestInner(bulkUpdateGuildChannelsRequestInner).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	bulkUpdateGuildChannelsRequestInner := []openapiclient.BulkUpdateGuildChannelsRequestInner{*openapiclient.NewBulkUpdateGuildChannelsRequestInner()} // []BulkUpdateGuildChannelsRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.BulkUpdateGuildChannels(context.Background(), guildId).BulkUpdateGuildChannelsRequestInner(bulkUpdateGuildChannelsRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpdateGuildChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateGuildChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateGuildChannelsRequestInner** | [**[]BulkUpdateGuildChannelsRequestInner**](BulkUpdateGuildChannelsRequestInner.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateGuildRoles

> []GuildRoleResponse BulkUpdateGuildRoles(ctx, guildId).BulkUpdateGuildRolesRequestInner(bulkUpdateGuildRolesRequestInner).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	bulkUpdateGuildRolesRequestInner := []openapiclient.BulkUpdateGuildRolesRequestInner{*openapiclient.NewBulkUpdateGuildRolesRequestInner()} // []BulkUpdateGuildRolesRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BulkUpdateGuildRoles(context.Background(), guildId).BulkUpdateGuildRolesRequestInner(bulkUpdateGuildRolesRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpdateGuildRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateGuildRoles`: []GuildRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkUpdateGuildRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateGuildRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateGuildRolesRequestInner** | [**[]BulkUpdateGuildRolesRequestInner**](BulkUpdateGuildRolesRequestInner.md) |  | 

### Return type

[**[]GuildRoleResponse**](GuildRoleResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateLobbyMembers

> []LobbyMemberResponse BulkUpdateLobbyMembers(ctx, lobbyId).BulkLobbyMemberRequest(bulkLobbyMemberRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	bulkLobbyMemberRequest := []openapiclient.BulkLobbyMemberRequest{*openapiclient.NewBulkLobbyMemberRequest("Id_example")} // []BulkLobbyMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BulkUpdateLobbyMembers(context.Background(), lobbyId).BulkLobbyMemberRequest(bulkLobbyMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpdateLobbyMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateLobbyMembers`: []LobbyMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkUpdateLobbyMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateLobbyMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkLobbyMemberRequest** | [**[]BulkLobbyMemberRequest**](BulkLobbyMemberRequest.md) |  | 

### Return type

[**[]LobbyMemberResponse**](LobbyMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumeEntitlement

> ConsumeEntitlement(ctx, applicationId, entitlementId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	entitlementId := "entitlementId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ConsumeEntitlement(context.Background(), applicationId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConsumeEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumeEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationCommand

> ApplicationCommandResponse CreateApplicationCommand(ctx, applicationId).ApplicationCommandCreateRequest(applicationCommandCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	applicationCommandCreateRequest := *openapiclient.NewApplicationCommandCreateRequest("Name_example") // ApplicationCommandCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateApplicationCommand(context.Background(), applicationId).ApplicationCommandCreateRequest(applicationCommandCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationCommand`: ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateApplicationCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationCommandCreateRequest** | [**ApplicationCommandCreateRequest**](ApplicationCommandCreateRequest.md) |  | 

### Return type

[**ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationEmoji

> EmojiResponse CreateApplicationEmoji(ctx, applicationId).CreateApplicationEmojiRequest(createApplicationEmojiRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	createApplicationEmojiRequest := *openapiclient.NewCreateApplicationEmojiRequest("Name_example", "Image_example") // CreateApplicationEmojiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateApplicationEmoji(context.Background(), applicationId).CreateApplicationEmojiRequest(createApplicationEmojiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateApplicationEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationEmoji`: EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateApplicationEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApplicationEmojiRequest** | [**CreateApplicationEmojiRequest**](CreateApplicationEmojiRequest.md) |  | 

### Return type

[**EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoModerationRule

> CreateAutoModerationRule200Response CreateAutoModerationRule(ctx, guildId).CreateAutoModerationRuleRequest(createAutoModerationRuleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	createAutoModerationRuleRequest := openapiclient.create_auto_moderation_rule_request{DefaultKeywordListUpsertRequest: openapiclient.NewDefaultKeywordListUpsertRequest("Name_example", int32(123), int32(123), *openapiclient.NewDefaultKeywordListTriggerMetadata())} // CreateAutoModerationRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateAutoModerationRule(context.Background(), guildId).CreateAutoModerationRuleRequest(createAutoModerationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateAutoModerationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoModerationRule`: CreateAutoModerationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateAutoModerationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoModerationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAutoModerationRuleRequest** | [**CreateAutoModerationRuleRequest**](CreateAutoModerationRuleRequest.md) |  | 

### Return type

[**CreateAutoModerationRule200Response**](CreateAutoModerationRule200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelInvite

> ListChannelInvites200ResponseInner CreateChannelInvite(ctx, channelId).CreateChannelInviteRequest(createChannelInviteRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	createChannelInviteRequest := *openapiclient.NewCreateChannelInviteRequest() // CreateChannelInviteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateChannelInvite(context.Background(), channelId).CreateChannelInviteRequest(createChannelInviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateChannelInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannelInvite`: ListChannelInvites200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateChannelInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createChannelInviteRequest** | [**CreateChannelInviteRequest**](CreateChannelInviteRequest.md) |  | 

### Return type

[**ListChannelInvites200ResponseInner**](ListChannelInvites200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDm

> AddGroupDmUser201Response CreateDm(ctx).CreatePrivateChannelRequest(createPrivateChannelRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	createPrivateChannelRequest := *openapiclient.NewCreatePrivateChannelRequest() // CreatePrivateChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateDm(context.Background()).CreatePrivateChannelRequest(createPrivateChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateDm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDm`: AddGroupDmUser201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateDm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPrivateChannelRequest** | [**CreatePrivateChannelRequest**](CreatePrivateChannelRequest.md) |  | 

### Return type

[**AddGroupDmUser201Response**](AddGroupDmUser201Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntitlement

> EntitlementResponse CreateEntitlement(ctx, applicationId).CreateEntitlementRequestData(createEntitlementRequestData).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	createEntitlementRequestData := *openapiclient.NewCreateEntitlementRequestData("SkuId_example", "OwnerId_example", int32(123)) // CreateEntitlementRequestData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateEntitlement(context.Background(), applicationId).CreateEntitlementRequestData(createEntitlementRequestData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntitlement`: EntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEntitlementRequestData** | [**CreateEntitlementRequestData**](CreateEntitlementRequestData.md) |  | 

### Return type

[**EntitlementResponse**](EntitlementResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuild

> GuildResponse CreateGuild(ctx).GuildCreateRequest(guildCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildCreateRequest := *openapiclient.NewGuildCreateRequest("Name_example") // GuildCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuild(context.Background()).GuildCreateRequest(guildCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuild`: GuildResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guildCreateRequest** | [**GuildCreateRequest**](GuildCreateRequest.md) |  | 

### Return type

[**GuildResponse**](GuildResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildApplicationCommand

> ApplicationCommandResponse CreateGuildApplicationCommand(ctx, applicationId, guildId).ApplicationCommandCreateRequest(applicationCommandCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	applicationCommandCreateRequest := *openapiclient.NewApplicationCommandCreateRequest("Name_example") // ApplicationCommandCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildApplicationCommand(context.Background(), applicationId, guildId).ApplicationCommandCreateRequest(applicationCommandCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildApplicationCommand`: ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildApplicationCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationCommandCreateRequest** | [**ApplicationCommandCreateRequest**](ApplicationCommandCreateRequest.md) |  | 

### Return type

[**ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildChannel

> GuildChannelResponse CreateGuildChannel(ctx, guildId).CreateGuildChannelRequest(createGuildChannelRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	createGuildChannelRequest := *openapiclient.NewCreateGuildChannelRequest("Name_example") // CreateGuildChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildChannel(context.Background(), guildId).CreateGuildChannelRequest(createGuildChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildChannel`: GuildChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGuildChannelRequest** | [**CreateGuildChannelRequest**](CreateGuildChannelRequest.md) |  | 

### Return type

[**GuildChannelResponse**](GuildChannelResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildEmoji

> EmojiResponse CreateGuildEmoji(ctx, guildId).CreateGuildEmojiRequest(createGuildEmojiRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	createGuildEmojiRequest := *openapiclient.NewCreateGuildEmojiRequest("Name_example", "Image_example") // CreateGuildEmojiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildEmoji(context.Background(), guildId).CreateGuildEmojiRequest(createGuildEmojiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildEmoji`: EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGuildEmojiRequest** | [**CreateGuildEmojiRequest**](CreateGuildEmojiRequest.md) |  | 

### Return type

[**EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildFromTemplate

> GuildResponse CreateGuildFromTemplate(ctx, code).CreateGuildFromTemplateRequest(createGuildFromTemplateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	code := "code_example" // string | 
	createGuildFromTemplateRequest := *openapiclient.NewCreateGuildFromTemplateRequest("Name_example") // CreateGuildFromTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildFromTemplate(context.Background(), code).CreateGuildFromTemplateRequest(createGuildFromTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildFromTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildFromTemplate`: GuildResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildFromTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGuildFromTemplateRequest** | [**CreateGuildFromTemplateRequest**](CreateGuildFromTemplateRequest.md) |  | 

### Return type

[**GuildResponse**](GuildResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildRole

> GuildRoleResponse CreateGuildRole(ctx, guildId).CreateGuildRoleRequest(createGuildRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	createGuildRoleRequest := *openapiclient.NewCreateGuildRoleRequest() // CreateGuildRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildRole(context.Background(), guildId).CreateGuildRoleRequest(createGuildRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildRole`: GuildRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGuildRoleRequest** | [**CreateGuildRoleRequest**](CreateGuildRoleRequest.md) |  | 

### Return type

[**GuildRoleResponse**](GuildRoleResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildScheduledEvent

> ListGuildScheduledEvents200ResponseInner CreateGuildScheduledEvent(ctx, guildId).CreateGuildScheduledEventRequest(createGuildScheduledEventRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	createGuildScheduledEventRequest := openapiclient.create_guild_scheduled_event_request{ExternalScheduledEventCreateRequest: openapiclient.NewExternalScheduledEventCreateRequest("Name_example", time.Now(), interface{}(123), NullableInt32(123), *openapiclient.NewEntityMetadataExternal("Location_example"))} // CreateGuildScheduledEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildScheduledEvent(context.Background(), guildId).CreateGuildScheduledEventRequest(createGuildScheduledEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildScheduledEvent`: ListGuildScheduledEvents200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGuildScheduledEventRequest** | [**CreateGuildScheduledEventRequest**](CreateGuildScheduledEventRequest.md) |  | 

### Return type

[**ListGuildScheduledEvents200ResponseInner**](ListGuildScheduledEvents200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildSoundboardSound

> SoundboardSoundResponse CreateGuildSoundboardSound(ctx, guildId).SoundboardCreateRequest(soundboardCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	soundboardCreateRequest := *openapiclient.NewSoundboardCreateRequest("Name_example", "Sound_example") // SoundboardCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildSoundboardSound(context.Background(), guildId).SoundboardCreateRequest(soundboardCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildSoundboardSound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildSoundboardSound`: SoundboardSoundResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildSoundboardSound`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildSoundboardSoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **soundboardCreateRequest** | [**SoundboardCreateRequest**](SoundboardCreateRequest.md) |  | 

### Return type

[**SoundboardSoundResponse**](SoundboardSoundResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildSticker

> GuildStickerResponse CreateGuildSticker(ctx, guildId).Name(name).Tags(tags).File(file).Description(description).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	name := "name_example" // string | 
	tags := "tags_example" // string | 
	file := "file_example" // string | 
	description := "description_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildSticker(context.Background(), guildId).Name(name).Tags(tags).File(file).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildSticker`: GuildStickerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildSticker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **tags** | **string** |  | 
 **file** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**GuildStickerResponse**](GuildStickerResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuildTemplate

> GuildTemplateResponse CreateGuildTemplate(ctx, guildId).CreateGuildTemplateRequest(createGuildTemplateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	createGuildTemplateRequest := *openapiclient.NewCreateGuildTemplateRequest("Name_example") // CreateGuildTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateGuildTemplate(context.Background(), guildId).CreateGuildTemplateRequest(createGuildTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateGuildTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuildTemplate`: GuildTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateGuildTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuildTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGuildTemplateRequest** | [**CreateGuildTemplateRequest**](CreateGuildTemplateRequest.md) |  | 

### Return type

[**GuildTemplateResponse**](GuildTemplateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInteractionResponse

> InteractionCallbackResponse CreateInteractionResponse(ctx, interactionId, interactionToken).CreateInteractionResponseRequest(createInteractionResponseRequest).WithResponse(withResponse).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	interactionId := "interactionId_example" // string | 
	interactionToken := "interactionToken_example" // string | 
	createInteractionResponseRequest := *openapiclient.NewCreateInteractionResponseRequest(NullableInt32(123), *openapiclient.NewIncomingWebhookUpdateForInteractionCallbackRequestPartial()) // CreateInteractionResponseRequest | 
	withResponse := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateInteractionResponse(context.Background(), interactionId, interactionToken).CreateInteractionResponseRequest(createInteractionResponseRequest).WithResponse(withResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateInteractionResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInteractionResponse`: InteractionCallbackResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateInteractionResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interactionId** | **string** |  | 
**interactionToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInteractionResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createInteractionResponseRequest** | [**CreateInteractionResponseRequest**](CreateInteractionResponseRequest.md) |  | 
 **withResponse** | **bool** |  | 

### Return type

[**InteractionCallbackResponse**](InteractionCallbackResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkedLobbyGuildInviteForSelf

> LobbyGuildInviteResponse CreateLinkedLobbyGuildInviteForSelf(ctx, lobbyId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateLinkedLobbyGuildInviteForSelf(context.Background(), lobbyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLinkedLobbyGuildInviteForSelf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLinkedLobbyGuildInviteForSelf`: LobbyGuildInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLinkedLobbyGuildInviteForSelf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkedLobbyGuildInviteForSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LobbyGuildInviteResponse**](LobbyGuildInviteResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkedLobbyGuildInviteForUser

> LobbyGuildInviteResponse CreateLinkedLobbyGuildInviteForUser(ctx, lobbyId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateLinkedLobbyGuildInviteForUser(context.Background(), lobbyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLinkedLobbyGuildInviteForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLinkedLobbyGuildInviteForUser`: LobbyGuildInviteResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLinkedLobbyGuildInviteForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkedLobbyGuildInviteForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LobbyGuildInviteResponse**](LobbyGuildInviteResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLobby

> LobbyResponse CreateLobby(ctx).CreateLobbyRequest(createLobbyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	createLobbyRequest := *openapiclient.NewCreateLobbyRequest() // CreateLobbyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateLobby(context.Background()).CreateLobbyRequest(createLobbyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLobby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLobby`: LobbyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLobby`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLobbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLobbyRequest** | [**CreateLobbyRequest**](CreateLobbyRequest.md) |  | 

### Return type

[**LobbyResponse**](LobbyResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLobbyMessage

> LobbyMessageResponse CreateLobbyMessage(ctx, lobbyId).SDKMessageRequest(sDKMessageRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	sDKMessageRequest := *openapiclient.NewSDKMessageRequest() // SDKMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateLobbyMessage(context.Background(), lobbyId).SDKMessageRequest(sDKMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLobbyMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLobbyMessage`: LobbyMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLobbyMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLobbyMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sDKMessageRequest** | [**SDKMessageRequest**](SDKMessageRequest.md) |  | 

### Return type

[**LobbyMessageResponse**](LobbyMessageResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> MessageResponse CreateMessage(ctx, channelId).MessageCreateRequest(messageCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageCreateRequest := *openapiclient.NewMessageCreateRequest() // MessageCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateMessage(context.Background(), channelId).MessageCreateRequest(messageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageCreateRequest** | [**MessageCreateRequest**](MessageCreateRequest.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrJoinLobby

> LobbyResponse CreateOrJoinLobby(ctx).CreateOrJoinLobbyRequest(createOrJoinLobbyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	createOrJoinLobbyRequest := *openapiclient.NewCreateOrJoinLobbyRequest("Secret_example") // CreateOrJoinLobbyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateOrJoinLobby(context.Background()).CreateOrJoinLobbyRequest(createOrJoinLobbyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateOrJoinLobby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrJoinLobby`: LobbyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateOrJoinLobby`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrJoinLobbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrJoinLobbyRequest** | [**CreateOrJoinLobbyRequest**](CreateOrJoinLobbyRequest.md) |  | 

### Return type

[**LobbyResponse**](LobbyResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePin

> CreatePin(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CreatePin(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStageInstance

> StageInstanceResponse CreateStageInstance(ctx).CreateStageInstanceRequest(createStageInstanceRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	createStageInstanceRequest := *openapiclient.NewCreateStageInstanceRequest("Topic_example", "ChannelId_example") // CreateStageInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateStageInstance(context.Background()).CreateStageInstanceRequest(createStageInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateStageInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStageInstance`: StageInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateStageInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStageInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStageInstanceRequest** | [**CreateStageInstanceRequest**](CreateStageInstanceRequest.md) |  | 

### Return type

[**StageInstanceResponse**](StageInstanceResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThread

> CreatedThreadResponse CreateThread(ctx, channelId).CreateThreadRequest(createThreadRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	createThreadRequest := *openapiclient.NewCreateThreadRequest("Name_example", *openapiclient.NewBaseCreateMessageCreateRequest()) // CreateThreadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateThread(context.Background(), channelId).CreateThreadRequest(createThreadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThread`: CreatedThreadResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateThread`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createThreadRequest** | [**CreateThreadRequest**](CreateThreadRequest.md) |  | 

### Return type

[**CreatedThreadResponse**](CreatedThreadResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThreadFromMessage

> ThreadResponse CreateThreadFromMessage(ctx, channelId, messageId).CreateTextThreadWithMessageRequest(createTextThreadWithMessageRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	createTextThreadWithMessageRequest := *openapiclient.NewCreateTextThreadWithMessageRequest("Name_example") // CreateTextThreadWithMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateThreadFromMessage(context.Background(), channelId, messageId).CreateTextThreadWithMessageRequest(createTextThreadWithMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateThreadFromMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThreadFromMessage`: ThreadResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateThreadFromMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateThreadFromMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTextThreadWithMessageRequest** | [**CreateTextThreadWithMessageRequest**](CreateTextThreadWithMessageRequest.md) |  | 

### Return type

[**ThreadResponse**](ThreadResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> GuildIncomingWebhookResponse CreateWebhook(ctx, channelId).CreateWebhookRequest(createWebhookRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	createWebhookRequest := *openapiclient.NewCreateWebhookRequest("Name_example") // CreateWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWebhook(context.Background(), channelId).CreateWebhookRequest(createWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: GuildIncomingWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWebhookRequest** | [**CreateWebhookRequest**](CreateWebhookRequest.md) |  | 

### Return type

[**GuildIncomingWebhookResponse**](GuildIncomingWebhookResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrosspostMessage

> MessageResponse CrosspostMessage(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CrosspostMessage(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CrosspostMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrosspostMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CrosspostMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrosspostMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllMessageReactions

> DeleteAllMessageReactions(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteAllMessageReactions(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteAllMessageReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllMessageReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllMessageReactionsByEmoji

> DeleteAllMessageReactionsByEmoji(ctx, channelId, messageId, emojiName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	emojiName := "emojiName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteAllMessageReactionsByEmoji(context.Background(), channelId, messageId, emojiName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteAllMessageReactionsByEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 
**emojiName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllMessageReactionsByEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationCommand

> DeleteApplicationCommand(ctx, applicationId, commandId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	commandId := "commandId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApplicationCommand(context.Background(), applicationId, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationEmoji

> DeleteApplicationEmoji(ctx, applicationId, emojiId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	emojiId := "emojiId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApplicationEmoji(context.Background(), applicationId, emojiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApplicationEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationUserRoleConnection

> DeleteApplicationUserRoleConnection(ctx, applicationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApplicationUserRoleConnection(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApplicationUserRoleConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationUserRoleConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoModerationRule

> DeleteAutoModerationRule(ctx, guildId, ruleId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteAutoModerationRule(context.Background(), guildId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteAutoModerationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoModerationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> GetChannel200Response DeleteChannel(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteChannel(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChannel`: GetChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetChannel200Response**](GetChannel200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannelPermissionOverwrite

> DeleteChannelPermissionOverwrite(ctx, channelId, overwriteId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	overwriteId := "overwriteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteChannelPermissionOverwrite(context.Background(), channelId, overwriteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteChannelPermissionOverwrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**overwriteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelPermissionOverwriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntitlement

> DeleteEntitlement(ctx, applicationId, entitlementId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	entitlementId := "entitlementId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteEntitlement(context.Background(), applicationId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupDmUser

> DeleteGroupDmUser(ctx, channelId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGroupDmUser(context.Background(), channelId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGroupDmUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupDmUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuild

> DeleteGuild(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuild(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildApplicationCommand

> DeleteGuildApplicationCommand(ctx, applicationId, guildId, commandId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	commandId := "commandId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildApplicationCommand(context.Background(), applicationId, guildId, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildEmoji

> DeleteGuildEmoji(ctx, guildId, emojiId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	emojiId := "emojiId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildEmoji(context.Background(), guildId, emojiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildIntegration

> DeleteGuildIntegration(ctx, guildId, integrationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	integrationId := "integrationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildIntegration(context.Background(), guildId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildMember

> DeleteGuildMember(ctx, guildId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildMember(context.Background(), guildId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildMemberRole

> DeleteGuildMemberRole(ctx, guildId, userId, roleId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildMemberRole(context.Background(), guildId, userId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildRole

> DeleteGuildRole(ctx, guildId, roleId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildRole(context.Background(), guildId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildScheduledEvent

> DeleteGuildScheduledEvent(ctx, guildId, guildScheduledEventId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	guildScheduledEventId := "guildScheduledEventId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildScheduledEvent(context.Background(), guildId, guildScheduledEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**guildScheduledEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildSoundboardSound

> DeleteGuildSoundboardSound(ctx, guildId, soundId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	soundId := "soundId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildSoundboardSound(context.Background(), guildId, soundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildSoundboardSound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**soundId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildSoundboardSoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildSticker

> DeleteGuildSticker(ctx, guildId, stickerId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	stickerId := "stickerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteGuildSticker(context.Background(), guildId, stickerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**stickerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuildTemplate

> GuildTemplateResponse DeleteGuildTemplate(ctx, guildId, code).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteGuildTemplate(context.Background(), guildId, code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteGuildTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGuildTemplate`: GuildTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteGuildTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuildTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuildTemplateResponse**](GuildTemplateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLobbyMember

> DeleteLobbyMember(ctx, lobbyId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteLobbyMember(context.Background(), lobbyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteLobbyMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLobbyMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessage

> DeleteMessage(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteMessage(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMyMessageReaction

> DeleteMyMessageReaction(ctx, channelId, messageId, emojiName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	emojiName := "emojiName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteMyMessageReaction(context.Background(), channelId, messageId, emojiName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMyMessageReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 
**emojiName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMyMessageReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOriginalWebhookMessage

> DeleteOriginalWebhookMessage(ctx, webhookId, webhookToken).ThreadId(threadId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	threadId := "threadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteOriginalWebhookMessage(context.Background(), webhookId, webhookToken).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteOriginalWebhookMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOriginalWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **threadId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePin

> DeletePin(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePin(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStageInstance

> DeleteStageInstance(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteStageInstance(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteStageInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStageInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThreadMember

> DeleteThreadMember(ctx, channelId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteThreadMember(context.Background(), channelId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteThreadMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThreadMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserMessageReaction

> DeleteUserMessageReaction(ctx, channelId, messageId, emojiName, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	emojiName := "emojiName_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteUserMessageReaction(context.Background(), channelId, messageId, emojiName, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteUserMessageReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 
**emojiName** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserMessageReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, webhookId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookByToken

> DeleteWebhookByToken(ctx, webhookId, webhookToken).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWebhookByToken(context.Background(), webhookId, webhookToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhookByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookMessage

> DeleteWebhookMessage(ctx, webhookId, webhookToken, messageId).ThreadId(threadId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	messageId := "messageId_example" // string | 
	threadId := "threadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWebhookMessage(context.Background(), webhookId, webhookToken, messageId).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhookMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **threadId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedCreatePin

> DeprecatedCreatePin(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeprecatedCreatePin(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeprecatedCreatePin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedCreatePinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedDeletePin

> DeprecatedDeletePin(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeprecatedDeletePin(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeprecatedDeletePin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedDeletePinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedListPins

> []MessageResponse DeprecatedListPins(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeprecatedListPins(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeprecatedListPins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeprecatedListPins`: []MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeprecatedListPins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedListPinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLobby

> LobbyResponse EditLobby(ctx, lobbyId).CreateLobbyRequest(createLobbyRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	createLobbyRequest := *openapiclient.NewCreateLobbyRequest() // CreateLobbyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.EditLobby(context.Background(), lobbyId).CreateLobbyRequest(createLobbyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EditLobby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditLobby`: LobbyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EditLobby`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLobbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLobbyRequest** | [**CreateLobbyRequest**](CreateLobbyRequest.md) |  | 

### Return type

[**LobbyResponse**](LobbyResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLobbyChannelLink

> LobbyResponse EditLobbyChannelLink(ctx, lobbyId).EditLobbyChannelLinkRequest(editLobbyChannelLinkRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	editLobbyChannelLinkRequest := *openapiclient.NewEditLobbyChannelLinkRequest() // EditLobbyChannelLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.EditLobbyChannelLink(context.Background(), lobbyId).EditLobbyChannelLinkRequest(editLobbyChannelLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EditLobbyChannelLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditLobbyChannelLink`: LobbyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EditLobbyChannelLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLobbyChannelLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editLobbyChannelLinkRequest** | [**EditLobbyChannelLinkRequest**](EditLobbyChannelLinkRequest.md) |  | 

### Return type

[**LobbyResponse**](LobbyResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteGithubCompatibleWebhook

> ExecuteGithubCompatibleWebhook(ctx, webhookId, webhookToken).GithubWebhook(githubWebhook).Wait(wait).ThreadId(threadId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	githubWebhook := *openapiclient.NewGithubWebhook(*openapiclient.NewGithubUser(int32(123), "Login_example", "HtmlUrl_example", "AvatarUrl_example")) // GithubWebhook | 
	wait := true // bool |  (optional)
	threadId := "threadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ExecuteGithubCompatibleWebhook(context.Background(), webhookId, webhookToken).GithubWebhook(githubWebhook).Wait(wait).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExecuteGithubCompatibleWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteGithubCompatibleWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **githubWebhook** | [**GithubWebhook**](GithubWebhook.md) |  | 
 **wait** | **bool** |  | 
 **threadId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteSlackCompatibleWebhook

> string ExecuteSlackCompatibleWebhook(ctx, webhookId, webhookToken).SlackWebhook(slackWebhook).Wait(wait).ThreadId(threadId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	slackWebhook := *openapiclient.NewSlackWebhook() // SlackWebhook | 
	wait := true // bool |  (optional)
	threadId := "threadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExecuteSlackCompatibleWebhook(context.Background(), webhookId, webhookToken).SlackWebhook(slackWebhook).Wait(wait).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExecuteSlackCompatibleWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteSlackCompatibleWebhook`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExecuteSlackCompatibleWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteSlackCompatibleWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **slackWebhook** | [**SlackWebhook**](SlackWebhook.md) |  | 
 **wait** | **bool** |  | 
 **threadId** | **string** |  | 

### Return type

**string**

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWebhook

> MessageResponse ExecuteWebhook(ctx, webhookId, webhookToken).ExecuteWebhookRequest(executeWebhookRequest).Wait(wait).ThreadId(threadId).WithComponents(withComponents).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	executeWebhookRequest := *openapiclient.NewExecuteWebhookRequest() // ExecuteWebhookRequest | 
	wait := true // bool |  (optional)
	threadId := "threadId_example" // string |  (optional)
	withComponents := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExecuteWebhook(context.Background(), webhookId, webhookToken).ExecuteWebhookRequest(executeWebhookRequest).Wait(wait).ThreadId(threadId).WithComponents(withComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExecuteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteWebhook`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExecuteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **executeWebhookRequest** | [**ExecuteWebhookRequest**](ExecuteWebhookRequest.md) |  | 
 **wait** | **bool** |  | 
 **threadId** | **string** |  | 
 **withComponents** | **bool** |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowChannel

> ChannelFollowerResponse FollowChannel(ctx, channelId).FollowChannelRequest(followChannelRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	followChannelRequest := *openapiclient.NewFollowChannelRequest("WebhookChannelId_example") // FollowChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FollowChannel(context.Background(), channelId).FollowChannelRequest(followChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FollowChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowChannel`: ChannelFollowerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FollowChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFollowChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **followChannelRequest** | [**FollowChannelRequest**](FollowChannelRequest.md) |  | 

### Return type

[**ChannelFollowerResponse**](ChannelFollowerResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveGuildThreads

> ThreadsResponse GetActiveGuildThreads(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetActiveGuildThreads(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetActiveGuildThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveGuildThreads`: ThreadsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetActiveGuildThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGuildThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThreadsResponse**](ThreadsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnswerVoters

> PollAnswerDetailsResponse GetAnswerVoters(ctx, channelId, messageId, answerId).After(after).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	answerId := int32(56) // int32 | 
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAnswerVoters(context.Background(), channelId, messageId, answerId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAnswerVoters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnswerVoters`: PollAnswerDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAnswerVoters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 
**answerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnswerVotersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**PollAnswerDetailsResponse**](PollAnswerDetailsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> PrivateApplicationResponse GetApplication(ctx, applicationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: PrivateApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateApplicationResponse**](PrivateApplicationResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCommand

> ApplicationCommandResponse GetApplicationCommand(ctx, applicationId, commandId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	commandId := "commandId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplicationCommand(context.Background(), applicationId, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationCommand`: ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplicationCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEmoji

> EmojiResponse GetApplicationEmoji(ctx, applicationId, emojiId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	emojiId := "emojiId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplicationEmoji(context.Background(), applicationId, emojiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplicationEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationEmoji`: EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplicationEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationRoleConnectionsMetadata

> []ApplicationRoleConnectionsMetadataItemResponse GetApplicationRoleConnectionsMetadata(ctx, applicationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplicationRoleConnectionsMetadata(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplicationRoleConnectionsMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationRoleConnectionsMetadata`: []ApplicationRoleConnectionsMetadataItemResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplicationRoleConnectionsMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRoleConnectionsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationRoleConnectionsMetadataItemResponse**](ApplicationRoleConnectionsMetadataItemResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationUserRoleConnection

> ApplicationUserRoleConnectionResponse GetApplicationUserRoleConnection(ctx, applicationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApplicationUserRoleConnection(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApplicationUserRoleConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationUserRoleConnection`: ApplicationUserRoleConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApplicationUserRoleConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationUserRoleConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationUserRoleConnectionResponse**](ApplicationUserRoleConnectionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoModerationRule

> CreateAutoModerationRule200Response GetAutoModerationRule(ctx, guildId, ruleId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAutoModerationRule(context.Background(), guildId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAutoModerationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoModerationRule`: CreateAutoModerationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAutoModerationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoModerationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateAutoModerationRule200Response**](CreateAutoModerationRule200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBotGateway

> GatewayBotResponse GetBotGateway(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBotGateway(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBotGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotGateway`: GatewayBotResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBotGateway`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotGatewayRequest struct via the builder pattern


### Return type

[**GatewayBotResponse**](GatewayBotResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannel

> GetChannel200Response GetChannel(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetChannel(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannel`: GetChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetChannel200Response**](GetChannel200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlement

> EntitlementResponse GetEntitlement(ctx, applicationId, entitlementId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	entitlementId := "entitlementId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEntitlement(context.Background(), applicationId, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlement`: EntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntitlementResponse**](EntitlementResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlements

> []*EntitlementResponse GetEntitlements(ctx, applicationId).SkuIds(skuIds).UserId(userId).GuildId(guildId).Before(before).After(after).Limit(limit).ExcludeEnded(excludeEnded).ExcludeDeleted(excludeDeleted).OnlyActive(onlyActive).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	skuIds := openapiclient.get_entitlements_sku_ids_parameter{ArrayOfString: new([]string)} // GetEntitlementsSkuIdsParameter | 
	applicationId := "applicationId_example" // string | 
	userId := "userId_example" // string |  (optional)
	guildId := "guildId_example" // string |  (optional)
	before := "before_example" // string |  (optional)
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	excludeEnded := true // bool |  (optional)
	excludeDeleted := true // bool |  (optional)
	onlyActive := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEntitlements(context.Background(), applicationId).SkuIds(skuIds).UserId(userId).GuildId(guildId).Before(before).After(after).Limit(limit).ExcludeEnded(excludeEnded).ExcludeDeleted(excludeDeleted).OnlyActive(onlyActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlements`: []*EntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuIds** | [**GetEntitlementsSkuIdsParameter**](GetEntitlementsSkuIdsParameter.md) |  | 

 **userId** | **string** |  | 
 **guildId** | **string** |  | 
 **before** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **excludeEnded** | **bool** |  | 
 **excludeDeleted** | **bool** |  | 
 **onlyActive** | **bool** |  | 

### Return type

[**[]*EntitlementResponse**](EntitlementResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGateway

> GatewayResponse GetGateway(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGateway(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGateway`: GatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGateway`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayRequest struct via the builder pattern


### Return type

[**GatewayResponse**](GatewayResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuild

> GuildWithCountsResponse GetGuild(ctx, guildId).WithCounts(withCounts).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	withCounts := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuild(context.Background(), guildId).WithCounts(withCounts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuild`: GuildWithCountsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withCounts** | **bool** |  | 

### Return type

[**GuildWithCountsResponse**](GuildWithCountsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildApplicationCommand

> ApplicationCommandResponse GetGuildApplicationCommand(ctx, applicationId, guildId, commandId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	commandId := "commandId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildApplicationCommand(context.Background(), applicationId, guildId, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildApplicationCommand`: ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildApplicationCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildApplicationCommandPermissions

> CommandPermissionsResponse GetGuildApplicationCommandPermissions(ctx, applicationId, guildId, commandId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	commandId := "commandId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildApplicationCommandPermissions(context.Background(), applicationId, guildId, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildApplicationCommandPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildApplicationCommandPermissions`: CommandPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildApplicationCommandPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildApplicationCommandPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CommandPermissionsResponse**](CommandPermissionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildBan

> GuildBanResponse GetGuildBan(ctx, guildId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildBan(context.Background(), guildId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildBan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildBan`: GuildBanResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildBan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildBanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuildBanResponse**](GuildBanResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildEmoji

> EmojiResponse GetGuildEmoji(ctx, guildId, emojiId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	emojiId := "emojiId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildEmoji(context.Background(), guildId, emojiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildEmoji`: EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildMember

> GuildMemberResponse GetGuildMember(ctx, guildId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildMember(context.Background(), guildId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildMember`: GuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuildMemberResponse**](GuildMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildNewMemberWelcome

> GuildHomeSettingsResponse GetGuildNewMemberWelcome(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildNewMemberWelcome(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildNewMemberWelcome``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildNewMemberWelcome`: GuildHomeSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildNewMemberWelcome`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildNewMemberWelcomeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GuildHomeSettingsResponse**](GuildHomeSettingsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildPreview

> GuildPreviewResponse GetGuildPreview(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildPreview(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildPreview`: GuildPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GuildPreviewResponse**](GuildPreviewResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildRole

> GuildRoleResponse GetGuildRole(ctx, guildId, roleId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildRole(context.Background(), guildId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildRole`: GuildRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuildRoleResponse**](GuildRoleResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildScheduledEvent

> ListGuildScheduledEvents200ResponseInner GetGuildScheduledEvent(ctx, guildId, guildScheduledEventId).WithUserCount(withUserCount).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	guildScheduledEventId := "guildScheduledEventId_example" // string | 
	withUserCount := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildScheduledEvent(context.Background(), guildId, guildScheduledEventId).WithUserCount(withUserCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildScheduledEvent`: ListGuildScheduledEvents200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**guildScheduledEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withUserCount** | **bool** |  | 

### Return type

[**ListGuildScheduledEvents200ResponseInner**](ListGuildScheduledEvents200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildSoundboardSound

> SoundboardSoundResponse GetGuildSoundboardSound(ctx, guildId, soundId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	soundId := "soundId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildSoundboardSound(context.Background(), guildId, soundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildSoundboardSound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildSoundboardSound`: SoundboardSoundResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildSoundboardSound`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**soundId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildSoundboardSoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SoundboardSoundResponse**](SoundboardSoundResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildSticker

> GuildStickerResponse GetGuildSticker(ctx, guildId, stickerId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	stickerId := "stickerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildSticker(context.Background(), guildId, stickerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildSticker`: GuildStickerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildSticker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**stickerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuildStickerResponse**](GuildStickerResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildTemplate

> GuildTemplateResponse GetGuildTemplate(ctx, code).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildTemplate(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildTemplate`: GuildTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GuildTemplateResponse**](GuildTemplateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildVanityUrl

> VanityURLResponse GetGuildVanityUrl(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildVanityUrl(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildVanityUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildVanityUrl`: VanityURLResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildVanityUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildVanityUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VanityURLResponse**](VanityURLResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildWebhooks

> []ListChannelWebhooks200ResponseInner GetGuildWebhooks(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildWebhooks(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildWebhooks`: []ListChannelWebhooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildWelcomeScreen

> GuildWelcomeScreenResponse GetGuildWelcomeScreen(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildWelcomeScreen(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildWelcomeScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildWelcomeScreen`: GuildWelcomeScreenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildWelcomeScreen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildWelcomeScreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GuildWelcomeScreenResponse**](GuildWelcomeScreenResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildWidget

> WidgetResponse GetGuildWidget(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildWidget(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildWidget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildWidget`: WidgetResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WidgetResponse**](WidgetResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildWidgetPng

> string GetGuildWidgetPng(ctx, guildId).Style(style).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	style := "style_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildWidgetPng(context.Background(), guildId).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildWidgetPng``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildWidgetPng`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildWidgetPng`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildWidgetPngRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **style** | **string** |  | 

### Return type

**string**

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildWidgetSettings

> WidgetSettingsResponse GetGuildWidgetSettings(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildWidgetSettings(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildWidgetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildWidgetSettings`: WidgetSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildWidgetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildWidgetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WidgetSettingsResponse**](WidgetSettingsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuildsOnboarding

> UserGuildOnboardingResponse GetGuildsOnboarding(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetGuildsOnboarding(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetGuildsOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuildsOnboarding`: UserGuildOnboardingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetGuildsOnboarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuildsOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGuildOnboardingResponse**](UserGuildOnboardingResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLobby

> LobbyResponse GetLobby(ctx, lobbyId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetLobby(context.Background(), lobbyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLobby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLobby`: LobbyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLobby`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLobbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LobbyResponse**](LobbyResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLobbyMessages

> []LobbyMessageResponse GetLobbyMessages(ctx, lobbyId).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetLobbyMessages(context.Background(), lobbyId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLobbyMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLobbyMessages`: []LobbyMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLobbyMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLobbyMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 

### Return type

[**[]LobbyMessageResponse**](LobbyMessageResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessage

> MessageResponse GetMessage(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMessage(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyApplication

> PrivateApplicationResponse GetMyApplication(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMyApplication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMyApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyApplication`: PrivateApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMyApplication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyApplicationRequest struct via the builder pattern


### Return type

[**PrivateApplicationResponse**](PrivateApplicationResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyGuildMember

> PrivateGuildMemberResponse GetMyGuildMember(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMyGuildMember(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMyGuildMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyGuildMember`: PrivateGuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMyGuildMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyGuildMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateGuildMemberResponse**](PrivateGuildMemberResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyOauth2Application

> PrivateApplicationResponse GetMyOauth2Application(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMyOauth2Application(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMyOauth2Application``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyOauth2Application`: PrivateApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMyOauth2Application`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyOauth2ApplicationRequest struct via the builder pattern


### Return type

[**PrivateApplicationResponse**](PrivateApplicationResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyOauth2Authorization

> OAuth2GetAuthorizationResponse GetMyOauth2Authorization(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMyOauth2Authorization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMyOauth2Authorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyOauth2Authorization`: OAuth2GetAuthorizationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMyOauth2Authorization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyOauth2AuthorizationRequest struct via the builder pattern


### Return type

[**OAuth2GetAuthorizationResponse**](OAuth2GetAuthorizationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyUser

> UserPIIResponse GetMyUser(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMyUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUser`: UserPIIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMyUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUserRequest struct via the builder pattern


### Return type

[**UserPIIResponse**](UserPIIResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenidConnectUserinfo

> OAuth2GetOpenIDConnectUserInfoResponse GetOpenidConnectUserinfo(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOpenidConnectUserinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOpenidConnectUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenidConnectUserinfo`: OAuth2GetOpenIDConnectUserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOpenidConnectUserinfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenidConnectUserinfoRequest struct via the builder pattern


### Return type

[**OAuth2GetOpenIDConnectUserInfoResponse**](OAuth2GetOpenIDConnectUserInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOriginalWebhookMessage

> MessageResponse GetOriginalWebhookMessage(ctx, webhookId, webhookToken).ThreadId(threadId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	threadId := "threadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOriginalWebhookMessage(context.Background(), webhookId, webhookToken).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOriginalWebhookMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOriginalWebhookMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOriginalWebhookMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginalWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **threadId** | **string** |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicKeys

> OAuth2GetKeys GetPublicKeys(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPublicKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPublicKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicKeys`: OAuth2GetKeys
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPublicKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicKeysRequest struct via the builder pattern


### Return type

[**OAuth2GetKeys**](OAuth2GetKeys.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfVoiceState

> VoiceStateResponse GetSelfVoiceState(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSelfVoiceState(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSelfVoiceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelfVoiceState`: VoiceStateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSelfVoiceState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfVoiceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VoiceStateResponse**](VoiceStateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoundboardDefaultSounds

> []SoundboardSoundResponse GetSoundboardDefaultSounds(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSoundboardDefaultSounds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSoundboardDefaultSounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSoundboardDefaultSounds`: []SoundboardSoundResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSoundboardDefaultSounds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoundboardDefaultSoundsRequest struct via the builder pattern


### Return type

[**[]SoundboardSoundResponse**](SoundboardSoundResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStageInstance

> StageInstanceResponse GetStageInstance(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetStageInstance(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetStageInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStageInstance`: StageInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetStageInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStageInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StageInstanceResponse**](StageInstanceResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSticker

> GetSticker200Response GetSticker(ctx, stickerId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	stickerId := "stickerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSticker(context.Background(), stickerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSticker`: GetSticker200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSticker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stickerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSticker200Response**](GetSticker200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStickerPack

> StickerPackResponse GetStickerPack(ctx, packId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	packId := "packId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetStickerPack(context.Background(), packId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetStickerPack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStickerPack`: StickerPackResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetStickerPack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStickerPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StickerPackResponse**](StickerPackResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreadMember

> ThreadMemberResponse GetThreadMember(ctx, channelId, userId).WithMember(withMember).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	userId := "userId_example" // string | 
	withMember := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetThreadMember(context.Background(), channelId, userId).WithMember(withMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetThreadMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreadMember`: ThreadMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetThreadMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreadMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withMember** | **bool** |  | 

### Return type

[**ThreadMemberResponse**](ThreadMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserResponse GetUser(ctx, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoiceState

> VoiceStateResponse GetVoiceState(ctx, guildId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetVoiceState(context.Background(), guildId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetVoiceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoiceState`: VoiceStateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetVoiceState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoiceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VoiceStateResponse**](VoiceStateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> ListChannelWebhooks200ResponseInner GetWebhook(ctx, webhookId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: ListChannelWebhooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookByToken

> ListChannelWebhooks200ResponseInner GetWebhookByToken(ctx, webhookId, webhookToken).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetWebhookByToken(context.Background(), webhookId, webhookToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhookByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookByToken`: ListChannelWebhooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhookByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookMessage

> MessageResponse GetWebhookMessage(ctx, webhookId, webhookToken, messageId).ThreadId(threadId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	messageId := "messageId_example" // string | 
	threadId := "threadId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetWebhookMessage(context.Background(), webhookId, webhookToken, messageId).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWebhookMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWebhookMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **threadId** | **string** |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteResolve

> ListChannelInvites200ResponseInner InviteResolve(ctx, code).WithCounts(withCounts).GuildScheduledEventId(guildScheduledEventId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	code := "code_example" // string | 
	withCounts := true // bool |  (optional)
	guildScheduledEventId := "guildScheduledEventId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InviteResolve(context.Background(), code).WithCounts(withCounts).GuildScheduledEventId(guildScheduledEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InviteResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteResolve`: ListChannelInvites200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InviteResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withCounts** | **bool** |  | 
 **guildScheduledEventId** | **string** |  | 

### Return type

[**ListChannelInvites200ResponseInner**](ListChannelInvites200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteRevoke

> ListChannelInvites200ResponseInner InviteRevoke(ctx, code).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InviteRevoke(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InviteRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteRevoke`: ListChannelInvites200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InviteRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListChannelInvites200ResponseInner**](ListChannelInvites200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinThread

> JoinThread(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.JoinThread(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.JoinThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveGuild

> LeaveGuild(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.LeaveGuild(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LeaveGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveLobby

> LeaveLobby(ctx, lobbyId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	lobbyId := "lobbyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.LeaveLobby(context.Background(), lobbyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LeaveLobby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lobbyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveLobbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveThread

> LeaveThread(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.LeaveThread(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LeaveThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationCommands

> []ApplicationCommandResponse ListApplicationCommands(ctx, applicationId).WithLocalizations(withLocalizations).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	withLocalizations := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListApplicationCommands(context.Background(), applicationId).WithLocalizations(withLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListApplicationCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationCommands`: []ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListApplicationCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withLocalizations** | **bool** |  | 

### Return type

[**[]ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationEmojis

> ListApplicationEmojisResponse ListApplicationEmojis(ctx, applicationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListApplicationEmojis(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListApplicationEmojis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationEmojis`: ListApplicationEmojisResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListApplicationEmojis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationEmojisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListApplicationEmojisResponse**](ListApplicationEmojisResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutoModerationRules

> []ListAutoModerationRules200ResponseInner ListAutoModerationRules(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListAutoModerationRules(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListAutoModerationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoModerationRules`: []ListAutoModerationRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListAutoModerationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoModerationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListAutoModerationRules200ResponseInner**](ListAutoModerationRules200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelInvites

> []ListChannelInvites200ResponseInner ListChannelInvites(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListChannelInvites(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListChannelInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannelInvites`: []ListChannelInvites200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListChannelInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListChannelInvites200ResponseInner**](ListChannelInvites200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelWebhooks

> []ListChannelWebhooks200ResponseInner ListChannelWebhooks(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListChannelWebhooks(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListChannelWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannelWebhooks`: []ListChannelWebhooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListChannelWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildApplicationCommandPermissions

> []CommandPermissionsResponse ListGuildApplicationCommandPermissions(ctx, applicationId, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildApplicationCommandPermissions(context.Background(), applicationId, guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildApplicationCommandPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildApplicationCommandPermissions`: []CommandPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildApplicationCommandPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildApplicationCommandPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CommandPermissionsResponse**](CommandPermissionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildApplicationCommands

> []ApplicationCommandResponse ListGuildApplicationCommands(ctx, applicationId, guildId).WithLocalizations(withLocalizations).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	withLocalizations := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildApplicationCommands(context.Background(), applicationId, guildId).WithLocalizations(withLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildApplicationCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildApplicationCommands`: []ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildApplicationCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildApplicationCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withLocalizations** | **bool** |  | 

### Return type

[**[]ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildAuditLogEntries

> GuildAuditLogResponse ListGuildAuditLogEntries(ctx, guildId).UserId(userId).TargetId(targetId).ActionType(actionType).Before(before).After(after).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string |  (optional)
	targetId := "targetId_example" // string |  (optional)
	actionType := int32(56) // int32 |  (optional)
	before := "before_example" // string |  (optional)
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildAuditLogEntries(context.Background(), guildId).UserId(userId).TargetId(targetId).ActionType(actionType).Before(before).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildAuditLogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildAuditLogEntries`: GuildAuditLogResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildAuditLogEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildAuditLogEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 
 **targetId** | **string** |  | 
 **actionType** | **int32** |  | 
 **before** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**GuildAuditLogResponse**](GuildAuditLogResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildBans

> []GuildBanResponse ListGuildBans(ctx, guildId).Limit(limit).Before(before).After(after).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	limit := int32(56) // int32 |  (optional)
	before := "before_example" // string |  (optional)
	after := "after_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildBans(context.Background(), guildId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildBans`: []GuildBanResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildBans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **before** | **string** |  | 
 **after** | **string** |  | 

### Return type

[**[]GuildBanResponse**](GuildBanResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildChannels

> []GetChannel200Response ListGuildChannels(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildChannels(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildChannels`: []GetChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetChannel200Response**](GetChannel200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildEmojis

> []EmojiResponse ListGuildEmojis(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildEmojis(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildEmojis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildEmojis`: []EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildEmojis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildEmojisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildIntegrations

> []ListGuildIntegrations200ResponseInner ListGuildIntegrations(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildIntegrations(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildIntegrations`: []ListGuildIntegrations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListGuildIntegrations200ResponseInner**](ListGuildIntegrations200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildInvites

> []ListChannelInvites200ResponseInner ListGuildInvites(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildInvites(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildInvites`: []ListChannelInvites200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListChannelInvites200ResponseInner**](ListChannelInvites200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildMembers

> []GuildMemberResponse ListGuildMembers(ctx, guildId).Limit(limit).After(after).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	limit := int32(56) // int32 |  (optional)
	after := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildMembers(context.Background(), guildId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildMembers`: []GuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **after** | **int32** |  | 

### Return type

[**[]GuildMemberResponse**](GuildMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildRoles

> []GuildRoleResponse ListGuildRoles(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildRoles(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildRoles`: []GuildRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GuildRoleResponse**](GuildRoleResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildScheduledEventUsers

> []ScheduledEventUserResponse ListGuildScheduledEventUsers(ctx, guildId, guildScheduledEventId).WithMember(withMember).Limit(limit).Before(before).After(after).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	guildScheduledEventId := "guildScheduledEventId_example" // string | 
	withMember := true // bool |  (optional)
	limit := int32(56) // int32 |  (optional)
	before := "before_example" // string |  (optional)
	after := "after_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildScheduledEventUsers(context.Background(), guildId, guildScheduledEventId).WithMember(withMember).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildScheduledEventUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildScheduledEventUsers`: []ScheduledEventUserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildScheduledEventUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**guildScheduledEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildScheduledEventUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withMember** | **bool** |  | 
 **limit** | **int32** |  | 
 **before** | **string** |  | 
 **after** | **string** |  | 

### Return type

[**[]ScheduledEventUserResponse**](ScheduledEventUserResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildScheduledEvents

> []ListGuildScheduledEvents200ResponseInner ListGuildScheduledEvents(ctx, guildId).WithUserCount(withUserCount).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	withUserCount := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildScheduledEvents(context.Background(), guildId).WithUserCount(withUserCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildScheduledEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildScheduledEvents`: []ListGuildScheduledEvents200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildScheduledEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildScheduledEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withUserCount** | **bool** |  | 

### Return type

[**[]ListGuildScheduledEvents200ResponseInner**](ListGuildScheduledEvents200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildSoundboardSounds

> ListGuildSoundboardSoundsResponse ListGuildSoundboardSounds(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildSoundboardSounds(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildSoundboardSounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildSoundboardSounds`: ListGuildSoundboardSoundsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildSoundboardSounds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildSoundboardSoundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListGuildSoundboardSoundsResponse**](ListGuildSoundboardSoundsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildStickers

> []GuildStickerResponse ListGuildStickers(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildStickers(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildStickers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildStickers`: []GuildStickerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildStickers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildStickersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GuildStickerResponse**](GuildStickerResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildTemplates

> []GuildTemplateResponse ListGuildTemplates(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildTemplates(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildTemplates`: []GuildTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GuildTemplateResponse**](GuildTemplateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGuildVoiceRegions

> []VoiceRegionResponse ListGuildVoiceRegions(ctx, guildId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListGuildVoiceRegions(context.Background(), guildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListGuildVoiceRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGuildVoiceRegions`: []VoiceRegionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListGuildVoiceRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGuildVoiceRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]VoiceRegionResponse**](VoiceRegionResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessageReactionsByEmoji

> []UserResponse ListMessageReactionsByEmoji(ctx, channelId, messageId, emojiName).After(after).Limit(limit).Type_(type_).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	emojiName := "emojiName_example" // string | 
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	type_ := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListMessageReactionsByEmoji(context.Background(), channelId, messageId, emojiName).After(after).Limit(limit).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListMessageReactionsByEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessageReactionsByEmoji`: []UserResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListMessageReactionsByEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 
**emojiName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMessageReactionsByEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **type_** | **int32** |  | 

### Return type

[**[]UserResponse**](UserResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessages

> []MessageResponse ListMessages(ctx, channelId).Around(around).Before(before).After(after).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	around := "around_example" // string |  (optional)
	before := "before_example" // string |  (optional)
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListMessages(context.Background(), channelId).Around(around).Before(before).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessages`: []MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **around** | **string** |  | 
 **before** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyConnections

> []ConnectedAccountResponse ListMyConnections(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListMyConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListMyConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyConnections`: []ConnectedAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListMyConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMyConnectionsRequest struct via the builder pattern


### Return type

[**[]ConnectedAccountResponse**](ConnectedAccountResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyGuilds

> []MyGuildResponse ListMyGuilds(ctx).Before(before).After(after).Limit(limit).WithCounts(withCounts).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	before := "before_example" // string |  (optional)
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	withCounts := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListMyGuilds(context.Background()).Before(before).After(after).Limit(limit).WithCounts(withCounts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListMyGuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyGuilds`: []MyGuildResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListMyGuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyGuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 
 **withCounts** | **bool** |  | 

### Return type

[**[]MyGuildResponse**](MyGuildResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyPrivateArchivedThreads

> ThreadsResponse ListMyPrivateArchivedThreads(ctx, channelId).Before(before).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListMyPrivateArchivedThreads(context.Background(), channelId).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListMyPrivateArchivedThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyPrivateArchivedThreads`: ThreadsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListMyPrivateArchivedThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMyPrivateArchivedThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**ThreadsResponse**](ThreadsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPins

> PinnedMessagesResponse ListPins(ctx, channelId).Before(before).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	before := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPins(context.Background(), channelId).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPins`: PinnedMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** |  | 
 **limit** | **int32** |  | 

### Return type

[**PinnedMessagesResponse**](PinnedMessagesResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateArchivedThreads

> ThreadsResponse ListPrivateArchivedThreads(ctx, channelId).Before(before).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	before := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPrivateArchivedThreads(context.Background(), channelId).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPrivateArchivedThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateArchivedThreads`: ThreadsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPrivateArchivedThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateArchivedThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** |  | 
 **limit** | **int32** |  | 

### Return type

[**ThreadsResponse**](ThreadsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicArchivedThreads

> ThreadsResponse ListPublicArchivedThreads(ctx, channelId).Before(before).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	before := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPublicArchivedThreads(context.Background(), channelId).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPublicArchivedThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicArchivedThreads`: ThreadsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPublicArchivedThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPublicArchivedThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** |  | 
 **limit** | **int32** |  | 

### Return type

[**ThreadsResponse**](ThreadsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStickerPacks

> StickerPackCollectionResponse ListStickerPacks(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListStickerPacks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListStickerPacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStickerPacks`: StickerPackCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListStickerPacks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStickerPacksRequest struct via the builder pattern


### Return type

[**StickerPackCollectionResponse**](StickerPackCollectionResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListThreadMembers

> []ThreadMemberResponse ListThreadMembers(ctx, channelId).WithMember(withMember).Limit(limit).After(after).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	withMember := true // bool |  (optional)
	limit := int32(56) // int32 |  (optional)
	after := "after_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListThreadMembers(context.Background(), channelId).WithMember(withMember).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListThreadMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListThreadMembers`: []ThreadMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListThreadMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListThreadMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withMember** | **bool** |  | 
 **limit** | **int32** |  | 
 **after** | **string** |  | 

### Return type

[**[]ThreadMemberResponse**](ThreadMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVoiceRegions

> []VoiceRegionResponse ListVoiceRegions(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListVoiceRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListVoiceRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVoiceRegions`: []VoiceRegionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListVoiceRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVoiceRegionsRequest struct via the builder pattern


### Return type

[**[]VoiceRegionResponse**](VoiceRegionResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerSdkToken

> ProvisionalTokenResponse PartnerSdkToken(ctx).PartnerSdkUnmergeProvisionalAccountRequest(partnerSdkUnmergeProvisionalAccountRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	partnerSdkUnmergeProvisionalAccountRequest := *openapiclient.NewPartnerSdkUnmergeProvisionalAccountRequest("ClientId_example", "ExternalAuthToken_example", "ExternalAuthType_example") // PartnerSdkUnmergeProvisionalAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PartnerSdkToken(context.Background()).PartnerSdkUnmergeProvisionalAccountRequest(partnerSdkUnmergeProvisionalAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PartnerSdkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartnerSdkToken`: ProvisionalTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PartnerSdkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerSdkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnerSdkUnmergeProvisionalAccountRequest** | [**PartnerSdkUnmergeProvisionalAccountRequest**](PartnerSdkUnmergeProvisionalAccountRequest.md) |  | 

### Return type

[**ProvisionalTokenResponse**](ProvisionalTokenResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartnerSdkUnmergeProvisionalAccount

> PartnerSdkUnmergeProvisionalAccount(ctx).PartnerSdkUnmergeProvisionalAccountRequest(partnerSdkUnmergeProvisionalAccountRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	partnerSdkUnmergeProvisionalAccountRequest := *openapiclient.NewPartnerSdkUnmergeProvisionalAccountRequest("ClientId_example", "ExternalAuthToken_example", "ExternalAuthType_example") // PartnerSdkUnmergeProvisionalAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PartnerSdkUnmergeProvisionalAccount(context.Background()).PartnerSdkUnmergeProvisionalAccountRequest(partnerSdkUnmergeProvisionalAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PartnerSdkUnmergeProvisionalAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartnerSdkUnmergeProvisionalAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnerSdkUnmergeProvisionalAccountRequest** | [**PartnerSdkUnmergeProvisionalAccountRequest**](PartnerSdkUnmergeProvisionalAccountRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PollExpire

> MessageResponse PollExpire(ctx, channelId, messageId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PollExpire(context.Background(), channelId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PollExpire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PollExpire`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PollExpire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPollExpireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewPruneGuild

> GuildPruneResponse PreviewPruneGuild(ctx, guildId).Days(days).IncludeRoles(includeRoles).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	days := int32(56) // int32 |  (optional)
	includeRoles := openapiclient.get_entitlements_sku_ids_parameter{ArrayOfString: new([]string)} // GetEntitlementsSkuIdsParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PreviewPruneGuild(context.Background(), guildId).Days(days).IncludeRoles(includeRoles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PreviewPruneGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewPruneGuild`: GuildPruneResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PreviewPruneGuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPruneGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** |  | 
 **includeRoles** | [**GetEntitlementsSkuIdsParameter**](GetEntitlementsSkuIdsParameter.md) |  | 

### Return type

[**GuildPruneResponse**](GuildPruneResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PruneGuild

> GuildPruneResponse PruneGuild(ctx, guildId).PruneGuildRequest(pruneGuildRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	pruneGuildRequest := *openapiclient.NewPruneGuildRequest() // PruneGuildRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PruneGuild(context.Background(), guildId).PruneGuildRequest(pruneGuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PruneGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PruneGuild`: GuildPruneResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PruneGuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPruneGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pruneGuildRequest** | [**PruneGuildRequest**](PruneGuildRequest.md) |  | 

### Return type

[**GuildPruneResponse**](GuildPruneResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutGuildsOnboarding

> GuildOnboardingResponse PutGuildsOnboarding(ctx, guildId).UpdateGuildOnboardingRequest(updateGuildOnboardingRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	updateGuildOnboardingRequest := *openapiclient.NewUpdateGuildOnboardingRequest() // UpdateGuildOnboardingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutGuildsOnboarding(context.Background(), guildId).UpdateGuildOnboardingRequest(updateGuildOnboardingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutGuildsOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutGuildsOnboarding`: GuildOnboardingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutGuildsOnboarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutGuildsOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGuildOnboardingRequest** | [**UpdateGuildOnboardingRequest**](UpdateGuildOnboardingRequest.md) |  | 

### Return type

[**GuildOnboardingResponse**](GuildOnboardingResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGuildMembers

> []GuildMemberResponse SearchGuildMembers(ctx, guildId).Limit(limit).Query(query).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	limit := int32(56) // int32 | 
	query := "query_example" // string | 
	guildId := "guildId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SearchGuildMembers(context.Background(), guildId).Limit(limit).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SearchGuildMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGuildMembers`: []GuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SearchGuildMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGuildMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **query** | **string** |  | 


### Return type

[**[]GuildMemberResponse**](GuildMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSoundboardSound

> SendSoundboardSound(ctx, channelId).SoundboardSoundSendRequest(soundboardSoundSendRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	soundboardSoundSendRequest := *openapiclient.NewSoundboardSoundSendRequest("SoundId_example") // SoundboardSoundSendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SendSoundboardSound(context.Background(), channelId).SoundboardSoundSendRequest(soundboardSoundSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SendSoundboardSound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSoundboardSoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **soundboardSoundSendRequest** | [**SoundboardSoundSendRequest**](SoundboardSoundSendRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetChannelPermissionOverwrite

> SetChannelPermissionOverwrite(ctx, channelId, overwriteId).SetChannelPermissionOverwriteRequest(setChannelPermissionOverwriteRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	overwriteId := "overwriteId_example" // string | 
	setChannelPermissionOverwriteRequest := *openapiclient.NewSetChannelPermissionOverwriteRequest() // SetChannelPermissionOverwriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SetChannelPermissionOverwrite(context.Background(), channelId, overwriteId).SetChannelPermissionOverwriteRequest(setChannelPermissionOverwriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetChannelPermissionOverwrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**overwriteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetChannelPermissionOverwriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setChannelPermissionOverwriteRequest** | [**SetChannelPermissionOverwriteRequest**](SetChannelPermissionOverwriteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGuildApplicationCommandPermissions

> CommandPermissionsResponse SetGuildApplicationCommandPermissions(ctx, applicationId, guildId, commandId).SetGuildApplicationCommandPermissionsRequest(setGuildApplicationCommandPermissionsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	commandId := "commandId_example" // string | 
	setGuildApplicationCommandPermissionsRequest := *openapiclient.NewSetGuildApplicationCommandPermissionsRequest() // SetGuildApplicationCommandPermissionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SetGuildApplicationCommandPermissions(context.Background(), applicationId, guildId, commandId).SetGuildApplicationCommandPermissionsRequest(setGuildApplicationCommandPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetGuildApplicationCommandPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGuildApplicationCommandPermissions`: CommandPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SetGuildApplicationCommandPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGuildApplicationCommandPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setGuildApplicationCommandPermissionsRequest** | [**SetGuildApplicationCommandPermissionsRequest**](SetGuildApplicationCommandPermissionsRequest.md) |  | 

### Return type

[**CommandPermissionsResponse**](CommandPermissionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGuildMfaLevel

> GuildMFALevelResponse SetGuildMfaLevel(ctx, guildId).SetGuildMfaLevelRequest(setGuildMfaLevelRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	setGuildMfaLevelRequest := *openapiclient.NewSetGuildMfaLevelRequest(int32(123)) // SetGuildMfaLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SetGuildMfaLevel(context.Background(), guildId).SetGuildMfaLevelRequest(setGuildMfaLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetGuildMfaLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGuildMfaLevel`: GuildMFALevelResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SetGuildMfaLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGuildMfaLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setGuildMfaLevelRequest** | [**SetGuildMfaLevelRequest**](SetGuildMfaLevelRequest.md) |  | 

### Return type

[**GuildMFALevelResponse**](GuildMFALevelResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGuildTemplate

> GuildTemplateResponse SyncGuildTemplate(ctx, guildId, code).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SyncGuildTemplate(context.Background(), guildId, code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncGuildTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncGuildTemplate`: GuildTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SyncGuildTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncGuildTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuildTemplateResponse**](GuildTemplateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreadSearch

> ThreadSearchResponse ThreadSearch(ctx, channelId).Name(name).Slop(slop).MinId(minId).MaxId(maxId).Tag(tag).TagSetting(tagSetting).Archived(archived).SortBy(sortBy).SortOrder(sortOrder).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	name := "name_example" // string |  (optional)
	slop := int32(56) // int32 |  (optional)
	minId := "minId_example" // string |  (optional)
	maxId := "maxId_example" // string |  (optional)
	tag := openapiclient.thread_search_tag_parameter{ArrayOfString: new([]string)} // ThreadSearchTagParameter |  (optional)
	tagSetting := "tagSetting_example" // string |  (optional)
	archived := true // bool |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortOrder := "sortOrder_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ThreadSearch(context.Background(), channelId).Name(name).Slop(slop).MinId(minId).MaxId(maxId).Tag(tag).TagSetting(tagSetting).Archived(archived).SortBy(sortBy).SortOrder(sortOrder).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ThreadSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreadSearch`: ThreadSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ThreadSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThreadSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **slop** | **int32** |  | 
 **minId** | **string** |  | 
 **maxId** | **string** |  | 
 **tag** | [**ThreadSearchTagParameter**](ThreadSearchTagParameter.md) |  | 
 **tagSetting** | **string** |  | 
 **archived** | **bool** |  | 
 **sortBy** | **string** |  | 
 **sortOrder** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**ThreadSearchResponse**](ThreadSearchResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerTypingIndicator

> map[string]interface{} TriggerTypingIndicator(ctx, channelId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TriggerTypingIndicator(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TriggerTypingIndicator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerTypingIndicator`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TriggerTypingIndicator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerTypingIndicatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbanUserFromGuild

> UnbanUserFromGuild(ctx, guildId, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UnbanUserFromGuild(context.Background(), guildId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UnbanUserFromGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbanUserFromGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> PrivateApplicationResponse UpdateApplication(ctx, applicationId).ApplicationFormPartial(applicationFormPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	applicationFormPartial := *openapiclient.NewApplicationFormPartial() // ApplicationFormPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplication(context.Background(), applicationId).ApplicationFormPartial(applicationFormPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplication`: PrivateApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationFormPartial** | [**ApplicationFormPartial**](ApplicationFormPartial.md) |  | 

### Return type

[**PrivateApplicationResponse**](PrivateApplicationResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationCommand

> ApplicationCommandResponse UpdateApplicationCommand(ctx, applicationId, commandId).ApplicationCommandPatchRequestPartial(applicationCommandPatchRequestPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	commandId := "commandId_example" // string | 
	applicationCommandPatchRequestPartial := *openapiclient.NewApplicationCommandPatchRequestPartial() // ApplicationCommandPatchRequestPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplicationCommand(context.Background(), applicationId, commandId).ApplicationCommandPatchRequestPartial(applicationCommandPatchRequestPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationCommand`: ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplicationCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationCommandPatchRequestPartial** | [**ApplicationCommandPatchRequestPartial**](ApplicationCommandPatchRequestPartial.md) |  | 

### Return type

[**ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationEmoji

> EmojiResponse UpdateApplicationEmoji(ctx, applicationId, emojiId).UpdateApplicationEmojiRequest(updateApplicationEmojiRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	emojiId := "emojiId_example" // string | 
	updateApplicationEmojiRequest := *openapiclient.NewUpdateApplicationEmojiRequest() // UpdateApplicationEmojiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplicationEmoji(context.Background(), applicationId, emojiId).UpdateApplicationEmojiRequest(updateApplicationEmojiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplicationEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationEmoji`: EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplicationEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateApplicationEmojiRequest** | [**UpdateApplicationEmojiRequest**](UpdateApplicationEmojiRequest.md) |  | 

### Return type

[**EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationRoleConnectionsMetadata

> []ApplicationRoleConnectionsMetadataItemResponse UpdateApplicationRoleConnectionsMetadata(ctx, applicationId).ApplicationRoleConnectionsMetadataItemRequest(applicationRoleConnectionsMetadataItemRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	applicationRoleConnectionsMetadataItemRequest := []openapiclient.ApplicationRoleConnectionsMetadataItemRequest{*openapiclient.NewApplicationRoleConnectionsMetadataItemRequest(int32(123), "Key_example", "Name_example", "Description_example")} // []ApplicationRoleConnectionsMetadataItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplicationRoleConnectionsMetadata(context.Background(), applicationId).ApplicationRoleConnectionsMetadataItemRequest(applicationRoleConnectionsMetadataItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplicationRoleConnectionsMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationRoleConnectionsMetadata`: []ApplicationRoleConnectionsMetadataItemResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplicationRoleConnectionsMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRoleConnectionsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationRoleConnectionsMetadataItemRequest** | [**[]ApplicationRoleConnectionsMetadataItemRequest**](ApplicationRoleConnectionsMetadataItemRequest.md) |  | 

### Return type

[**[]ApplicationRoleConnectionsMetadataItemResponse**](ApplicationRoleConnectionsMetadataItemResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationUserRoleConnection

> ApplicationUserRoleConnectionResponse UpdateApplicationUserRoleConnection(ctx, applicationId).UpdateApplicationUserRoleConnectionRequest(updateApplicationUserRoleConnectionRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	updateApplicationUserRoleConnectionRequest := *openapiclient.NewUpdateApplicationUserRoleConnectionRequest() // UpdateApplicationUserRoleConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplicationUserRoleConnection(context.Background(), applicationId).UpdateApplicationUserRoleConnectionRequest(updateApplicationUserRoleConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplicationUserRoleConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationUserRoleConnection`: ApplicationUserRoleConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplicationUserRoleConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationUserRoleConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApplicationUserRoleConnectionRequest** | [**UpdateApplicationUserRoleConnectionRequest**](UpdateApplicationUserRoleConnectionRequest.md) |  | 

### Return type

[**ApplicationUserRoleConnectionResponse**](ApplicationUserRoleConnectionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoModerationRule

> CreateAutoModerationRule200Response UpdateAutoModerationRule(ctx, guildId, ruleId).UpdateAutoModerationRuleRequest(updateAutoModerationRuleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	ruleId := "ruleId_example" // string | 
	updateAutoModerationRuleRequest := *openapiclient.NewUpdateAutoModerationRuleRequest() // UpdateAutoModerationRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateAutoModerationRule(context.Background(), guildId, ruleId).UpdateAutoModerationRuleRequest(updateAutoModerationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateAutoModerationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoModerationRule`: CreateAutoModerationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateAutoModerationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoModerationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAutoModerationRuleRequest** | [**UpdateAutoModerationRuleRequest**](UpdateAutoModerationRuleRequest.md) |  | 

### Return type

[**CreateAutoModerationRule200Response**](CreateAutoModerationRule200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> GetChannel200Response UpdateChannel(ctx, channelId).UpdateChannelRequest(updateChannelRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	updateChannelRequest := *openapiclient.NewUpdateChannelRequest() // UpdateChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateChannel(context.Background(), channelId).UpdateChannelRequest(updateChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannel`: GetChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChannelRequest** | [**UpdateChannelRequest**](UpdateChannelRequest.md) |  | 

### Return type

[**GetChannel200Response**](GetChannel200Response.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuild

> GuildResponse UpdateGuild(ctx, guildId).GuildPatchRequestPartial(guildPatchRequestPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	guildPatchRequestPartial := *openapiclient.NewGuildPatchRequestPartial() // GuildPatchRequestPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuild(context.Background(), guildId).GuildPatchRequestPartial(guildPatchRequestPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuild`: GuildResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guildPatchRequestPartial** | [**GuildPatchRequestPartial**](GuildPatchRequestPartial.md) |  | 

### Return type

[**GuildResponse**](GuildResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildApplicationCommand

> ApplicationCommandResponse UpdateGuildApplicationCommand(ctx, applicationId, guildId, commandId).ApplicationCommandPatchRequestPartial(applicationCommandPatchRequestPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	guildId := "guildId_example" // string | 
	commandId := "commandId_example" // string | 
	applicationCommandPatchRequestPartial := *openapiclient.NewApplicationCommandPatchRequestPartial() // ApplicationCommandPatchRequestPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildApplicationCommand(context.Background(), applicationId, guildId, commandId).ApplicationCommandPatchRequestPartial(applicationCommandPatchRequestPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildApplicationCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildApplicationCommand`: ApplicationCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildApplicationCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**guildId** | **string** |  | 
**commandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildApplicationCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationCommandPatchRequestPartial** | [**ApplicationCommandPatchRequestPartial**](ApplicationCommandPatchRequestPartial.md) |  | 

### Return type

[**ApplicationCommandResponse**](ApplicationCommandResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildEmoji

> EmojiResponse UpdateGuildEmoji(ctx, guildId, emojiId).UpdateGuildEmojiRequest(updateGuildEmojiRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	emojiId := "emojiId_example" // string | 
	updateGuildEmojiRequest := *openapiclient.NewUpdateGuildEmojiRequest() // UpdateGuildEmojiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildEmoji(context.Background(), guildId, emojiId).UpdateGuildEmojiRequest(updateGuildEmojiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildEmoji`: EmojiResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**emojiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGuildEmojiRequest** | [**UpdateGuildEmojiRequest**](UpdateGuildEmojiRequest.md) |  | 

### Return type

[**EmojiResponse**](EmojiResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildMember

> GuildMemberResponse UpdateGuildMember(ctx, guildId, userId).UpdateGuildMemberRequest(updateGuildMemberRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	updateGuildMemberRequest := *openapiclient.NewUpdateGuildMemberRequest() // UpdateGuildMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildMember(context.Background(), guildId, userId).UpdateGuildMemberRequest(updateGuildMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildMember`: GuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGuildMemberRequest** | [**UpdateGuildMemberRequest**](UpdateGuildMemberRequest.md) |  | 

### Return type

[**GuildMemberResponse**](GuildMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildRole

> GuildRoleResponse UpdateGuildRole(ctx, guildId, roleId).CreateGuildRoleRequest(createGuildRoleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	roleId := "roleId_example" // string | 
	createGuildRoleRequest := *openapiclient.NewCreateGuildRoleRequest() // CreateGuildRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildRole(context.Background(), guildId, roleId).CreateGuildRoleRequest(createGuildRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildRole`: GuildRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createGuildRoleRequest** | [**CreateGuildRoleRequest**](CreateGuildRoleRequest.md) |  | 

### Return type

[**GuildRoleResponse**](GuildRoleResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildScheduledEvent

> ListGuildScheduledEvents200ResponseInner UpdateGuildScheduledEvent(ctx, guildId, guildScheduledEventId).UpdateGuildScheduledEventRequest(updateGuildScheduledEventRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	guildScheduledEventId := "guildScheduledEventId_example" // string | 
	updateGuildScheduledEventRequest := *openapiclient.NewUpdateGuildScheduledEventRequest() // UpdateGuildScheduledEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildScheduledEvent(context.Background(), guildId, guildScheduledEventId).UpdateGuildScheduledEventRequest(updateGuildScheduledEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildScheduledEvent`: ListGuildScheduledEvents200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**guildScheduledEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGuildScheduledEventRequest** | [**UpdateGuildScheduledEventRequest**](UpdateGuildScheduledEventRequest.md) |  | 

### Return type

[**ListGuildScheduledEvents200ResponseInner**](ListGuildScheduledEvents200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildSoundboardSound

> SoundboardSoundResponse UpdateGuildSoundboardSound(ctx, guildId, soundId).SoundboardPatchRequestPartial(soundboardPatchRequestPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	soundId := "soundId_example" // string | 
	soundboardPatchRequestPartial := *openapiclient.NewSoundboardPatchRequestPartial() // SoundboardPatchRequestPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildSoundboardSound(context.Background(), guildId, soundId).SoundboardPatchRequestPartial(soundboardPatchRequestPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildSoundboardSound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildSoundboardSound`: SoundboardSoundResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildSoundboardSound`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**soundId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildSoundboardSoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **soundboardPatchRequestPartial** | [**SoundboardPatchRequestPartial**](SoundboardPatchRequestPartial.md) |  | 

### Return type

[**SoundboardSoundResponse**](SoundboardSoundResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildSticker

> GuildStickerResponse UpdateGuildSticker(ctx, guildId, stickerId).UpdateGuildStickerRequest(updateGuildStickerRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	stickerId := "stickerId_example" // string | 
	updateGuildStickerRequest := *openapiclient.NewUpdateGuildStickerRequest() // UpdateGuildStickerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildSticker(context.Background(), guildId, stickerId).UpdateGuildStickerRequest(updateGuildStickerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildSticker`: GuildStickerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildSticker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**stickerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGuildStickerRequest** | [**UpdateGuildStickerRequest**](UpdateGuildStickerRequest.md) |  | 

### Return type

[**GuildStickerResponse**](GuildStickerResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildTemplate

> GuildTemplateResponse UpdateGuildTemplate(ctx, guildId, code).UpdateGuildTemplateRequest(updateGuildTemplateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	code := "code_example" // string | 
	updateGuildTemplateRequest := *openapiclient.NewUpdateGuildTemplateRequest() // UpdateGuildTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildTemplate(context.Background(), guildId, code).UpdateGuildTemplateRequest(updateGuildTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildTemplate`: GuildTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGuildTemplateRequest** | [**UpdateGuildTemplateRequest**](UpdateGuildTemplateRequest.md) |  | 

### Return type

[**GuildTemplateResponse**](GuildTemplateResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildWelcomeScreen

> GuildWelcomeScreenResponse UpdateGuildWelcomeScreen(ctx, guildId).WelcomeScreenPatchRequestPartial(welcomeScreenPatchRequestPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	welcomeScreenPatchRequestPartial := *openapiclient.NewWelcomeScreenPatchRequestPartial() // WelcomeScreenPatchRequestPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildWelcomeScreen(context.Background(), guildId).WelcomeScreenPatchRequestPartial(welcomeScreenPatchRequestPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildWelcomeScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildWelcomeScreen`: GuildWelcomeScreenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildWelcomeScreen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildWelcomeScreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **welcomeScreenPatchRequestPartial** | [**WelcomeScreenPatchRequestPartial**](WelcomeScreenPatchRequestPartial.md) |  | 

### Return type

[**GuildWelcomeScreenResponse**](GuildWelcomeScreenResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuildWidgetSettings

> WidgetSettingsResponse UpdateGuildWidgetSettings(ctx, guildId).UpdateGuildWidgetSettingsRequest(updateGuildWidgetSettingsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	updateGuildWidgetSettingsRequest := *openapiclient.NewUpdateGuildWidgetSettingsRequest() // UpdateGuildWidgetSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateGuildWidgetSettings(context.Background(), guildId).UpdateGuildWidgetSettingsRequest(updateGuildWidgetSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateGuildWidgetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGuildWidgetSettings`: WidgetSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateGuildWidgetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuildWidgetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGuildWidgetSettingsRequest** | [**UpdateGuildWidgetSettingsRequest**](UpdateGuildWidgetSettingsRequest.md) |  | 

### Return type

[**WidgetSettingsResponse**](WidgetSettingsResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> MessageResponse UpdateMessage(ctx, channelId, messageId).MessageEditRequestPartial(messageEditRequestPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	messageEditRequestPartial := *openapiclient.NewMessageEditRequestPartial() // MessageEditRequestPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateMessage(context.Background(), channelId, messageId).MessageEditRequestPartial(messageEditRequestPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messageEditRequestPartial** | [**MessageEditRequestPartial**](MessageEditRequestPartial.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyApplication

> PrivateApplicationResponse UpdateMyApplication(ctx).ApplicationFormPartial(applicationFormPartial).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationFormPartial := *openapiclient.NewApplicationFormPartial() // ApplicationFormPartial | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateMyApplication(context.Background()).ApplicationFormPartial(applicationFormPartial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMyApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMyApplication`: PrivateApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMyApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationFormPartial** | [**ApplicationFormPartial**](ApplicationFormPartial.md) |  | 

### Return type

[**PrivateApplicationResponse**](PrivateApplicationResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyGuildMember

> PrivateGuildMemberResponse UpdateMyGuildMember(ctx, guildId).UpdateMyGuildMemberRequest(updateMyGuildMemberRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	updateMyGuildMemberRequest := *openapiclient.NewUpdateMyGuildMemberRequest() // UpdateMyGuildMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateMyGuildMember(context.Background(), guildId).UpdateMyGuildMemberRequest(updateMyGuildMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMyGuildMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMyGuildMember`: PrivateGuildMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMyGuildMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyGuildMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMyGuildMemberRequest** | [**UpdateMyGuildMemberRequest**](UpdateMyGuildMemberRequest.md) |  | 

### Return type

[**PrivateGuildMemberResponse**](PrivateGuildMemberResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyUser

> UserPIIResponse UpdateMyUser(ctx).BotAccountPatchRequest(botAccountPatchRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	botAccountPatchRequest := *openapiclient.NewBotAccountPatchRequest("Username_example") // BotAccountPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateMyUser(context.Background()).BotAccountPatchRequest(botAccountPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMyUser`: UserPIIResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMyUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botAccountPatchRequest** | [**BotAccountPatchRequest**](BotAccountPatchRequest.md) |  | 

### Return type

[**UserPIIResponse**](UserPIIResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOriginalWebhookMessage

> MessageResponse UpdateOriginalWebhookMessage(ctx, webhookId, webhookToken).IncomingWebhookUpdateRequestPartial(incomingWebhookUpdateRequestPartial).ThreadId(threadId).WithComponents(withComponents).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	incomingWebhookUpdateRequestPartial := *openapiclient.NewIncomingWebhookUpdateRequestPartial() // IncomingWebhookUpdateRequestPartial | 
	threadId := "threadId_example" // string |  (optional)
	withComponents := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateOriginalWebhookMessage(context.Background(), webhookId, webhookToken).IncomingWebhookUpdateRequestPartial(incomingWebhookUpdateRequestPartial).ThreadId(threadId).WithComponents(withComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateOriginalWebhookMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOriginalWebhookMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateOriginalWebhookMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOriginalWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **incomingWebhookUpdateRequestPartial** | [**IncomingWebhookUpdateRequestPartial**](IncomingWebhookUpdateRequestPartial.md) |  | 
 **threadId** | **string** |  | 
 **withComponents** | **bool** |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSelfVoiceState

> UpdateSelfVoiceState(ctx, guildId).UpdateSelfVoiceStateRequest(updateSelfVoiceStateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	updateSelfVoiceStateRequest := *openapiclient.NewUpdateSelfVoiceStateRequest() // UpdateSelfVoiceStateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateSelfVoiceState(context.Background(), guildId).UpdateSelfVoiceStateRequest(updateSelfVoiceStateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSelfVoiceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSelfVoiceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSelfVoiceStateRequest** | [**UpdateSelfVoiceStateRequest**](UpdateSelfVoiceStateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStageInstance

> StageInstanceResponse UpdateStageInstance(ctx, channelId).UpdateStageInstanceRequest(updateStageInstanceRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	channelId := "channelId_example" // string | 
	updateStageInstanceRequest := *openapiclient.NewUpdateStageInstanceRequest() // UpdateStageInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateStageInstance(context.Background(), channelId).UpdateStageInstanceRequest(updateStageInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateStageInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStageInstance`: StageInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateStageInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStageInstanceRequest** | [**UpdateStageInstanceRequest**](UpdateStageInstanceRequest.md) |  | 

### Return type

[**StageInstanceResponse**](StageInstanceResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVoiceState

> UpdateVoiceState(ctx, guildId, userId).UpdateVoiceStateRequest(updateVoiceStateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	updateVoiceStateRequest := *openapiclient.NewUpdateVoiceStateRequest() // UpdateVoiceStateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateVoiceState(context.Background(), guildId, userId).UpdateVoiceStateRequest(updateVoiceStateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateVoiceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoiceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVoiceStateRequest** | [**UpdateVoiceStateRequest**](UpdateVoiceStateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> ListChannelWebhooks200ResponseInner UpdateWebhook(ctx, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	updateWebhookRequest := *openapiclient.NewUpdateWebhookRequest() // UpdateWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateWebhook(context.Background(), webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: ListChannelWebhooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookRequest** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) |  | 

### Return type

[**ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookByToken

> ListChannelWebhooks200ResponseInner UpdateWebhookByToken(ctx, webhookId, webhookToken).UpdateWebhookByTokenRequest(updateWebhookByTokenRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	updateWebhookByTokenRequest := *openapiclient.NewUpdateWebhookByTokenRequest() // UpdateWebhookByTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateWebhookByToken(context.Background(), webhookId, webhookToken).UpdateWebhookByTokenRequest(updateWebhookByTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhookByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookByToken`: ListChannelWebhooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhookByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWebhookByTokenRequest** | [**UpdateWebhookByTokenRequest**](UpdateWebhookByTokenRequest.md) |  | 

### Return type

[**ListChannelWebhooks200ResponseInner**](ListChannelWebhooks200ResponseInner.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookMessage

> MessageResponse UpdateWebhookMessage(ctx, webhookId, webhookToken, messageId).IncomingWebhookUpdateRequestPartial(incomingWebhookUpdateRequestPartial).ThreadId(threadId).WithComponents(withComponents).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhookToken := "webhookToken_example" // string | 
	messageId := "messageId_example" // string | 
	incomingWebhookUpdateRequestPartial := *openapiclient.NewIncomingWebhookUpdateRequestPartial() // IncomingWebhookUpdateRequestPartial | 
	threadId := "threadId_example" // string |  (optional)
	withComponents := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateWebhookMessage(context.Background(), webhookId, webhookToken, messageId).IncomingWebhookUpdateRequestPartial(incomingWebhookUpdateRequestPartial).ThreadId(threadId).WithComponents(withComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhookMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhookMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 
**webhookToken** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **incomingWebhookUpdateRequestPartial** | [**IncomingWebhookUpdateRequestPartial**](IncomingWebhookUpdateRequestPartial.md) |  | 
 **threadId** | **string** |  | 
 **withComponents** | **bool** |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadApplicationAttachment

> ActivitiesAttachmentResponse UploadApplicationAttachment(ctx, applicationId).File(file).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dc-api/dc-rest-go"
)

func main() {
	applicationId := "applicationId_example" // string | 
	file := "file_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UploadApplicationAttachment(context.Background(), applicationId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadApplicationAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadApplicationAttachment`: ActivitiesAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UploadApplicationAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadApplicationAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **string** |  | 

### Return type

[**ActivitiesAttachmentResponse**](ActivitiesAttachmentResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2), [BotToken](../README.md#BotToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

