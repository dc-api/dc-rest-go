/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Discord HTTP API (Preview) service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package dc_rest

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the GuildChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildChannelResponse{}

// GuildChannelResponse struct for GuildChannelResponse
type GuildChannelResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	LastMessageId *string `json:"last_message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Flags int32 `json:"flags"`
	LastPinTimestamp NullableTime `json:"last_pin_timestamp,omitempty"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	ParentId *string `json:"parent_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	RateLimitPerUser NullableInt32 `json:"rate_limit_per_user,omitempty"`
	Bitrate NullableInt32 `json:"bitrate,omitempty"`
	UserLimit NullableInt32 `json:"user_limit,omitempty"`
	RtcRegion NullableString `json:"rtc_region,omitempty"`
	VideoQualityMode *int32 `json:"video_quality_mode,omitempty"`
	Permissions NullableString `json:"permissions,omitempty"`
	Topic NullableString `json:"topic,omitempty"`
	DefaultAutoArchiveDuration *int32 `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser NullableInt32 `json:"default_thread_rate_limit_per_user,omitempty"`
	Position int32 `json:"position"`
	PermissionOverwrites []ChannelPermissionOverwriteResponse `json:"permission_overwrites,omitempty"`
	Nsfw NullableBool `json:"nsfw,omitempty"`
	AvailableTags []ForumTagResponse `json:"available_tags,omitempty"`
	DefaultReactionEmoji NullableDefaultReactionEmojiResponse `json:"default_reaction_emoji,omitempty"`
	DefaultSortOrder *int32 `json:"default_sort_order,omitempty"`
	DefaultForumLayout *int32 `json:"default_forum_layout,omitempty"`
	DefaultTagSetting *string `json:"default_tag_setting,omitempty"`
	HdStreamingUntil NullableTime `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerId *string `json:"hd_streaming_buyer_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _GuildChannelResponse GuildChannelResponse

// NewGuildChannelResponse instantiates a new GuildChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildChannelResponse(id string, type_ int32, flags int32, guildId string, name string, position int32) *GuildChannelResponse {
	this := GuildChannelResponse{}
	this.Id = id
	this.Type = type_
	this.Flags = flags
	this.GuildId = guildId
	this.Name = name
	this.Position = position
	return &this
}

// NewGuildChannelResponseWithDefaults instantiates a new GuildChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildChannelResponseWithDefaults() *GuildChannelResponse {
	this := GuildChannelResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildChannelResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildChannelResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *GuildChannelResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GuildChannelResponse) SetType(v int32) {
	o.Type = v
}

// GetLastMessageId returns the LastMessageId field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetLastMessageId() string {
	if o == nil || IsNil(o.LastMessageId) {
		var ret string
		return ret
	}
	return *o.LastMessageId
}

// GetLastMessageIdOk returns a tuple with the LastMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetLastMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastMessageId) {
		return nil, false
	}
	return o.LastMessageId, true
}

// HasLastMessageId returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasLastMessageId() bool {
	if o != nil && !IsNil(o.LastMessageId) {
		return true
	}

	return false
}

// SetLastMessageId gets a reference to the given string and assigns it to the LastMessageId field.
func (o *GuildChannelResponse) SetLastMessageId(v string) {
	o.LastMessageId = &v
}


// GetFlags returns the Flags field value
func (o *GuildChannelResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *GuildChannelResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetLastPinTimestamp returns the LastPinTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetLastPinTimestamp() time.Time {
	if o == nil || IsNil(o.LastPinTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPinTimestamp.Get()
}

// GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetLastPinTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPinTimestamp.Get()) {
		return nil, false
	}
	return o.LastPinTimestamp.Get(), o.LastPinTimestamp.IsSet()
}

// HasLastPinTimestamp returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasLastPinTimestamp() bool {
	if o != nil && o.LastPinTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLastPinTimestamp gets a reference to the given NullableTime and assigns it to the LastPinTimestamp field.
func (o *GuildChannelResponse) SetLastPinTimestamp(v time.Time) {
	o.LastPinTimestamp.Set(&v)
}

// SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil
func (o *GuildChannelResponse) SetLastPinTimestampNil() {
	o.LastPinTimestamp.Set(nil)
}

// UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
func (o *GuildChannelResponse) UnsetLastPinTimestamp() {
	o.LastPinTimestamp.Unset()
}

// GetGuildId returns the GuildId field value
func (o *GuildChannelResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *GuildChannelResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetName returns the Name field value
func (o *GuildChannelResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildChannelResponse) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *GuildChannelResponse) SetParentId(v string) {
	o.ParentId = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser.Get()
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		return nil, false
	}
	return o.RateLimitPerUser.Get(), o.RateLimitPerUser.IsSet()
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasRateLimitPerUser() bool {
	if o != nil && o.RateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the RateLimitPerUser field.
func (o *GuildChannelResponse) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser.Set(&v)
}

// SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil
func (o *GuildChannelResponse) SetRateLimitPerUserNil() {
	o.RateLimitPerUser.Set(nil)
}

// UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
func (o *GuildChannelResponse) UnsetRateLimitPerUser() {
	o.RateLimitPerUser.Unset()
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate.Get()) {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt32 and assigns it to the Bitrate field.
func (o *GuildChannelResponse) SetBitrate(v int32) {
	o.Bitrate.Set(&v)
}

// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *GuildChannelResponse) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *GuildChannelResponse) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetUserLimit returns the UserLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.UserLimit.Get()
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit.Get()) {
		return nil, false
	}
	return o.UserLimit.Get(), o.UserLimit.IsSet()
}

// HasUserLimit returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasUserLimit() bool {
	if o != nil && o.UserLimit.IsSet() {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given NullableInt32 and assigns it to the UserLimit field.
func (o *GuildChannelResponse) SetUserLimit(v int32) {
	o.UserLimit.Set(&v)
}

// SetUserLimitNil sets the value for UserLimit to be an explicit nil
func (o *GuildChannelResponse) SetUserLimitNil() {
	o.UserLimit.Set(nil)
}

// UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
func (o *GuildChannelResponse) UnsetUserLimit() {
	o.UserLimit.Unset()
}

// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion.Get()) {
		var ret string
		return ret
	}
	return *o.RtcRegion.Get()
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion.Get()) {
		return nil, false
	}
	return o.RtcRegion.Get(), o.RtcRegion.IsSet()
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasRtcRegion() bool {
	if o != nil && o.RtcRegion.IsSet() {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given NullableString and assigns it to the RtcRegion field.
func (o *GuildChannelResponse) SetRtcRegion(v string) {
	o.RtcRegion.Set(&v)
}

// SetRtcRegionNil sets the value for RtcRegion to be an explicit nil
func (o *GuildChannelResponse) SetRtcRegionNil() {
	o.RtcRegion.Set(nil)
}

// UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
func (o *GuildChannelResponse) UnsetRtcRegion() {
	o.RtcRegion.Unset()
}

// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetVideoQualityMode() int32 {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret int32
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetVideoQualityModeOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoQualityMode) {
		return nil, false
	}
	return o.VideoQualityMode, true
}

// HasVideoQualityMode returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasVideoQualityMode() bool {
	if o != nil && !IsNil(o.VideoQualityMode) {
		return true
	}

	return false
}

// SetVideoQualityMode gets a reference to the given int32 and assigns it to the VideoQualityMode field.
func (o *GuildChannelResponse) SetVideoQualityMode(v int32) {
	o.VideoQualityMode = &v
}


// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetPermissions() string {
	if o == nil || IsNil(o.Permissions.Get()) {
		var ret string
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetPermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.Permissions.Get()) {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasPermissions() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableString and assigns it to the Permissions field.
func (o *GuildChannelResponse) SetPermissions(v string) {
	o.Permissions.Set(&v)
}

// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *GuildChannelResponse) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *GuildChannelResponse) UnsetPermissions() {
	o.Permissions.Unset()
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic.Get()) {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *GuildChannelResponse) SetTopic(v string) {
	o.Topic.Set(&v)
}

// SetTopicNil sets the value for Topic to be an explicit nil
func (o *GuildChannelResponse) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *GuildChannelResponse) UnsetTopic() {
	o.Topic.Unset()
}

// GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetDefaultAutoArchiveDuration() int32 {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		var ret int32
		return ret
	}
	return *o.DefaultAutoArchiveDuration
}

// GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetDefaultAutoArchiveDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		return nil, false
	}
	return o.DefaultAutoArchiveDuration, true
}

// HasDefaultAutoArchiveDuration returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasDefaultAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.DefaultAutoArchiveDuration) {
		return true
	}

	return false
}

// SetDefaultAutoArchiveDuration gets a reference to the given int32 and assigns it to the DefaultAutoArchiveDuration field.
func (o *GuildChannelResponse) SetDefaultAutoArchiveDuration(v int32) {
	o.DefaultAutoArchiveDuration = &v
}


// GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetDefaultThreadRateLimitPerUser() int32 {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultThreadRateLimitPerUser.Get()
}

// GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetDefaultThreadRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser.Get()) {
		return nil, false
	}
	return o.DefaultThreadRateLimitPerUser.Get(), o.DefaultThreadRateLimitPerUser.IsSet()
}

// HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasDefaultThreadRateLimitPerUser() bool {
	if o != nil && o.DefaultThreadRateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetDefaultThreadRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the DefaultThreadRateLimitPerUser field.
func (o *GuildChannelResponse) SetDefaultThreadRateLimitPerUser(v int32) {
	o.DefaultThreadRateLimitPerUser.Set(&v)
}

// SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil
func (o *GuildChannelResponse) SetDefaultThreadRateLimitPerUserNil() {
	o.DefaultThreadRateLimitPerUser.Set(nil)
}

// UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
func (o *GuildChannelResponse) UnsetDefaultThreadRateLimitPerUser() {
	o.DefaultThreadRateLimitPerUser.Unset()
}

// GetPosition returns the Position field value
func (o *GuildChannelResponse) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *GuildChannelResponse) SetPosition(v int32) {
	o.Position = v
}

// GetPermissionOverwrites returns the PermissionOverwrites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetPermissionOverwrites() []ChannelPermissionOverwriteResponse {
	if o == nil {
		var ret []ChannelPermissionOverwriteResponse
		return ret
	}
	return o.PermissionOverwrites
}

// GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetPermissionOverwritesOk() ([]ChannelPermissionOverwriteResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionOverwrites, true
}

// HasPermissionOverwrites returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasPermissionOverwrites() bool {
	if o != nil && !IsNil(o.PermissionOverwrites) {
		return true
	}

	return false
}

// SetPermissionOverwrites gets a reference to the given []ChannelPermissionOverwriteResponse and assigns it to the PermissionOverwrites field.
func (o *GuildChannelResponse) SetPermissionOverwrites(v []ChannelPermissionOverwriteResponse) {
	o.PermissionOverwrites = v
}


// GetNsfw returns the Nsfw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw.Get()) {
		var ret bool
		return ret
	}
	return *o.Nsfw.Get()
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw.Get()) {
		return nil, false
	}
	return o.Nsfw.Get(), o.Nsfw.IsSet()
}

// HasNsfw returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasNsfw() bool {
	if o != nil && o.Nsfw.IsSet() {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given NullableBool and assigns it to the Nsfw field.
func (o *GuildChannelResponse) SetNsfw(v bool) {
	o.Nsfw.Set(&v)
}

// SetNsfwNil sets the value for Nsfw to be an explicit nil
func (o *GuildChannelResponse) SetNsfwNil() {
	o.Nsfw.Set(nil)
}

// UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
func (o *GuildChannelResponse) UnsetNsfw() {
	o.Nsfw.Unset()
}

// GetAvailableTags returns the AvailableTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetAvailableTags() []ForumTagResponse {
	if o == nil {
		var ret []ForumTagResponse
		return ret
	}
	return o.AvailableTags
}

// GetAvailableTagsOk returns a tuple with the AvailableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetAvailableTagsOk() ([]ForumTagResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableTags, true
}

// HasAvailableTags returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasAvailableTags() bool {
	if o != nil && !IsNil(o.AvailableTags) {
		return true
	}

	return false
}

// SetAvailableTags gets a reference to the given []ForumTagResponse and assigns it to the AvailableTags field.
func (o *GuildChannelResponse) SetAvailableTags(v []ForumTagResponse) {
	o.AvailableTags = v
}


// GetDefaultReactionEmoji returns the DefaultReactionEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetDefaultReactionEmoji() DefaultReactionEmojiResponse {
	if o == nil || IsNil(o.DefaultReactionEmoji.Get()) {
		var ret DefaultReactionEmojiResponse
		return ret
	}
	return *o.DefaultReactionEmoji.Get()
}

// GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetDefaultReactionEmojiOk() (*DefaultReactionEmojiResponse, bool) {
	if o == nil || IsNil(o.DefaultReactionEmoji.Get()) {
		return nil, false
	}
	return o.DefaultReactionEmoji.Get(), o.DefaultReactionEmoji.IsSet()
}

// HasDefaultReactionEmoji returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasDefaultReactionEmoji() bool {
	if o != nil && o.DefaultReactionEmoji.IsSet() {
		return true
	}

	return false
}

// SetDefaultReactionEmoji gets a reference to the given NullableDefaultReactionEmojiResponse and assigns it to the DefaultReactionEmoji field.
func (o *GuildChannelResponse) SetDefaultReactionEmoji(v DefaultReactionEmojiResponse) {
	o.DefaultReactionEmoji.Set(&v)
}

// SetDefaultReactionEmojiNil sets the value for DefaultReactionEmoji to be an explicit nil
func (o *GuildChannelResponse) SetDefaultReactionEmojiNil() {
	o.DefaultReactionEmoji.Set(nil)
}

// UnsetDefaultReactionEmoji ensures that no value is present for DefaultReactionEmoji, not even an explicit nil
func (o *GuildChannelResponse) UnsetDefaultReactionEmoji() {
	o.DefaultReactionEmoji.Unset()
}

// GetDefaultSortOrder returns the DefaultSortOrder field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetDefaultSortOrder() int32 {
	if o == nil || IsNil(o.DefaultSortOrder) {
		var ret int32
		return ret
	}
	return *o.DefaultSortOrder
}

// GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetDefaultSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultSortOrder) {
		return nil, false
	}
	return o.DefaultSortOrder, true
}

// HasDefaultSortOrder returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasDefaultSortOrder() bool {
	if o != nil && !IsNil(o.DefaultSortOrder) {
		return true
	}

	return false
}

// SetDefaultSortOrder gets a reference to the given int32 and assigns it to the DefaultSortOrder field.
func (o *GuildChannelResponse) SetDefaultSortOrder(v int32) {
	o.DefaultSortOrder = &v
}


// GetDefaultForumLayout returns the DefaultForumLayout field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetDefaultForumLayout() int32 {
	if o == nil || IsNil(o.DefaultForumLayout) {
		var ret int32
		return ret
	}
	return *o.DefaultForumLayout
}

// GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetDefaultForumLayoutOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultForumLayout) {
		return nil, false
	}
	return o.DefaultForumLayout, true
}

// HasDefaultForumLayout returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasDefaultForumLayout() bool {
	if o != nil && !IsNil(o.DefaultForumLayout) {
		return true
	}

	return false
}

// SetDefaultForumLayout gets a reference to the given int32 and assigns it to the DefaultForumLayout field.
func (o *GuildChannelResponse) SetDefaultForumLayout(v int32) {
	o.DefaultForumLayout = &v
}


// GetDefaultTagSetting returns the DefaultTagSetting field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetDefaultTagSetting() string {
	if o == nil || IsNil(o.DefaultTagSetting) {
		var ret string
		return ret
	}
	return *o.DefaultTagSetting
}

// GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetDefaultTagSettingOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTagSetting) {
		return nil, false
	}
	return o.DefaultTagSetting, true
}

// HasDefaultTagSetting returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasDefaultTagSetting() bool {
	if o != nil && !IsNil(o.DefaultTagSetting) {
		return true
	}

	return false
}

// SetDefaultTagSetting gets a reference to the given string and assigns it to the DefaultTagSetting field.
func (o *GuildChannelResponse) SetDefaultTagSetting(v string) {
	o.DefaultTagSetting = &v
}


// GetHdStreamingUntil returns the HdStreamingUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildChannelResponse) GetHdStreamingUntil() time.Time {
	if o == nil || IsNil(o.HdStreamingUntil.Get()) {
		var ret time.Time
		return ret
	}
	return *o.HdStreamingUntil.Get()
}

// GetHdStreamingUntilOk returns a tuple with the HdStreamingUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildChannelResponse) GetHdStreamingUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.HdStreamingUntil.Get()) {
		return nil, false
	}
	return o.HdStreamingUntil.Get(), o.HdStreamingUntil.IsSet()
}

// HasHdStreamingUntil returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasHdStreamingUntil() bool {
	if o != nil && o.HdStreamingUntil.IsSet() {
		return true
	}

	return false
}

// SetHdStreamingUntil gets a reference to the given NullableTime and assigns it to the HdStreamingUntil field.
func (o *GuildChannelResponse) SetHdStreamingUntil(v time.Time) {
	o.HdStreamingUntil.Set(&v)
}

// SetHdStreamingUntilNil sets the value for HdStreamingUntil to be an explicit nil
func (o *GuildChannelResponse) SetHdStreamingUntilNil() {
	o.HdStreamingUntil.Set(nil)
}

// UnsetHdStreamingUntil ensures that no value is present for HdStreamingUntil, not even an explicit nil
func (o *GuildChannelResponse) UnsetHdStreamingUntil() {
	o.HdStreamingUntil.Unset()
}

// GetHdStreamingBuyerId returns the HdStreamingBuyerId field value if set, zero value otherwise.
func (o *GuildChannelResponse) GetHdStreamingBuyerId() string {
	if o == nil || IsNil(o.HdStreamingBuyerId) {
		var ret string
		return ret
	}
	return *o.HdStreamingBuyerId
}

// GetHdStreamingBuyerIdOk returns a tuple with the HdStreamingBuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildChannelResponse) GetHdStreamingBuyerIdOk() (*string, bool) {
	if o == nil || IsNil(o.HdStreamingBuyerId) {
		return nil, false
	}
	return o.HdStreamingBuyerId, true
}

// HasHdStreamingBuyerId returns a boolean if a field has been set.
func (o *GuildChannelResponse) HasHdStreamingBuyerId() bool {
	if o != nil && !IsNil(o.HdStreamingBuyerId) {
		return true
	}

	return false
}

// SetHdStreamingBuyerId gets a reference to the given string and assigns it to the HdStreamingBuyerId field.
func (o *GuildChannelResponse) SetHdStreamingBuyerId(v string) {
	o.HdStreamingBuyerId = &v
}


func (o GuildChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.LastMessageId) {
		toSerialize["last_message_id"] = o.LastMessageId
	}
	toSerialize["flags"] = o.Flags
	if o.LastPinTimestamp.IsSet() {
		toSerialize["last_pin_timestamp"] = o.LastPinTimestamp.Get()
	}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.RateLimitPerUser.IsSet() {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser.Get()
	}
	if o.Bitrate.IsSet() {
		toSerialize["bitrate"] = o.Bitrate.Get()
	}
	if o.UserLimit.IsSet() {
		toSerialize["user_limit"] = o.UserLimit.Get()
	}
	if o.RtcRegion.IsSet() {
		toSerialize["rtc_region"] = o.RtcRegion.Get()
	}
	if !IsNil(o.VideoQualityMode) {
		toSerialize["video_quality_mode"] = o.VideoQualityMode
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	if !IsNil(o.DefaultAutoArchiveDuration) {
		toSerialize["default_auto_archive_duration"] = o.DefaultAutoArchiveDuration
	}
	if o.DefaultThreadRateLimitPerUser.IsSet() {
		toSerialize["default_thread_rate_limit_per_user"] = o.DefaultThreadRateLimitPerUser.Get()
	}
	toSerialize["position"] = o.Position
	if o.PermissionOverwrites != nil {
		toSerialize["permission_overwrites"] = o.PermissionOverwrites
	}
	if o.Nsfw.IsSet() {
		toSerialize["nsfw"] = o.Nsfw.Get()
	}
	if o.AvailableTags != nil {
		toSerialize["available_tags"] = o.AvailableTags
	}
	if o.DefaultReactionEmoji.IsSet() {
		toSerialize["default_reaction_emoji"] = o.DefaultReactionEmoji.Get()
	}
	if !IsNil(o.DefaultSortOrder) {
		toSerialize["default_sort_order"] = o.DefaultSortOrder
	}
	if !IsNil(o.DefaultForumLayout) {
		toSerialize["default_forum_layout"] = o.DefaultForumLayout
	}
	if !IsNil(o.DefaultTagSetting) {
		toSerialize["default_tag_setting"] = o.DefaultTagSetting
	}
	if o.HdStreamingUntil.IsSet() {
		toSerialize["hd_streaming_until"] = o.HdStreamingUntil.Get()
	}
	if !IsNil(o.HdStreamingBuyerId) {
		toSerialize["hd_streaming_buyer_id"] = o.HdStreamingBuyerId
	}
	return toSerialize, nil
}

func (o *GuildChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"flags",
		"guild_id",
		"name",
		"position",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGuildChannelResponse := _GuildChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildChannelResponse)

	if err != nil {
		return err
	}

	*o = GuildChannelResponse(varGuildChannelResponse)

	return err
}

type NullableGuildChannelResponse struct {
	value *GuildChannelResponse
	isSet bool
}

func (v NullableGuildChannelResponse) Get() *GuildChannelResponse {
	return v.value
}

func (v *NullableGuildChannelResponse) Set(val *GuildChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildChannelResponse(val *GuildChannelResponse) *NullableGuildChannelResponse {
	return &NullableGuildChannelResponse{value: val, isSet: true}
}

func (v NullableGuildChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


