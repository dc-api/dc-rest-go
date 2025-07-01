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
	"bytes"
	"fmt"
)

// checks if the CreateGuildChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildChannelRequest{}

// CreateGuildChannelRequest struct for CreateGuildChannelRequest
type CreateGuildChannelRequest struct {
	Type NullableInt32 `json:"type,omitempty"`
	Name string `json:"name"`
	Position NullableInt32 `json:"position,omitempty"`
	Topic NullableString `json:"topic,omitempty"`
	Bitrate NullableInt32 `json:"bitrate,omitempty"`
	UserLimit NullableInt32 `json:"user_limit,omitempty"`
	Nsfw NullableBool `json:"nsfw,omitempty"`
	RateLimitPerUser NullableInt32 `json:"rate_limit_per_user,omitempty"`
	ParentId *string `json:"parent_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	PermissionOverwrites []ChannelPermissionOverwriteRequest `json:"permission_overwrites,omitempty"`
	RtcRegion NullableString `json:"rtc_region,omitempty"`
	VideoQualityMode *int32 `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration *int32 `json:"default_auto_archive_duration,omitempty"`
	DefaultReactionEmoji NullableUpdateDefaultReactionEmojiRequest `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser NullableInt32 `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder *int32 `json:"default_sort_order,omitempty"`
	DefaultForumLayout *int32 `json:"default_forum_layout,omitempty"`
	DefaultTagSetting *string `json:"default_tag_setting,omitempty"`
	AvailableTags []*CreateOrUpdateThreadTagRequest `json:"available_tags,omitempty"`
}

type _CreateGuildChannelRequest CreateGuildChannelRequest

// NewCreateGuildChannelRequest instantiates a new CreateGuildChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGuildChannelRequest(name string) *CreateGuildChannelRequest {
	this := CreateGuildChannelRequest{}
	this.Name = name
	return &this
}

// NewCreateGuildChannelRequestWithDefaults instantiates a new CreateGuildChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGuildChannelRequestWithDefaults() *CreateGuildChannelRequest {
	this := CreateGuildChannelRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *CreateGuildChannelRequest) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateGuildChannelRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetType() {
	o.Type.Unset()
}

// GetName returns the Name field value
func (o *CreateGuildChannelRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGuildChannelRequest) SetName(v string) {
	o.Name = v
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetPosition() int32 {
	if o == nil || IsNil(o.Position.Get()) {
		var ret int32
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position.Get()) {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableInt32 and assigns it to the Position field.
func (o *CreateGuildChannelRequest) SetPosition(v int32) {
	o.Position.Set(&v)
}

// SetPositionNil sets the value for Position to be an explicit nil
func (o *CreateGuildChannelRequest) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetPosition() {
	o.Position.Unset()
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic.Get()) {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *CreateGuildChannelRequest) SetTopic(v string) {
	o.Topic.Set(&v)
}

// SetTopicNil sets the value for Topic to be an explicit nil
func (o *CreateGuildChannelRequest) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetTopic() {
	o.Topic.Unset()
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate.Get()) {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt32 and assigns it to the Bitrate field.
func (o *CreateGuildChannelRequest) SetBitrate(v int32) {
	o.Bitrate.Set(&v)
}

// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *CreateGuildChannelRequest) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetUserLimit returns the UserLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.UserLimit.Get()
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit.Get()) {
		return nil, false
	}
	return o.UserLimit.Get(), o.UserLimit.IsSet()
}

// HasUserLimit returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasUserLimit() bool {
	if o != nil && o.UserLimit.IsSet() {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given NullableInt32 and assigns it to the UserLimit field.
func (o *CreateGuildChannelRequest) SetUserLimit(v int32) {
	o.UserLimit.Set(&v)
}

// SetUserLimitNil sets the value for UserLimit to be an explicit nil
func (o *CreateGuildChannelRequest) SetUserLimitNil() {
	o.UserLimit.Set(nil)
}

// UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetUserLimit() {
	o.UserLimit.Unset()
}

// GetNsfw returns the Nsfw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw.Get()) {
		var ret bool
		return ret
	}
	return *o.Nsfw.Get()
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw.Get()) {
		return nil, false
	}
	return o.Nsfw.Get(), o.Nsfw.IsSet()
}

// HasNsfw returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasNsfw() bool {
	if o != nil && o.Nsfw.IsSet() {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given NullableBool and assigns it to the Nsfw field.
func (o *CreateGuildChannelRequest) SetNsfw(v bool) {
	o.Nsfw.Set(&v)
}

// SetNsfwNil sets the value for Nsfw to be an explicit nil
func (o *CreateGuildChannelRequest) SetNsfwNil() {
	o.Nsfw.Set(nil)
}

// UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetNsfw() {
	o.Nsfw.Unset()
}

// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser.Get()
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		return nil, false
	}
	return o.RateLimitPerUser.Get(), o.RateLimitPerUser.IsSet()
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasRateLimitPerUser() bool {
	if o != nil && o.RateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the RateLimitPerUser field.
func (o *CreateGuildChannelRequest) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser.Set(&v)
}

// SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil
func (o *CreateGuildChannelRequest) SetRateLimitPerUserNil() {
	o.RateLimitPerUser.Set(nil)
}

// UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetRateLimitPerUser() {
	o.RateLimitPerUser.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CreateGuildChannelRequest) SetParentId(v string) {
	o.ParentId = &v
}


// GetPermissionOverwrites returns the PermissionOverwrites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest {
	if o == nil {
		var ret []ChannelPermissionOverwriteRequest
		return ret
	}
	return o.PermissionOverwrites
}

// GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetPermissionOverwritesOk() ([]ChannelPermissionOverwriteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionOverwrites, true
}

// HasPermissionOverwrites returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasPermissionOverwrites() bool {
	if o != nil && !IsNil(o.PermissionOverwrites) {
		return true
	}

	return false
}

// SetPermissionOverwrites gets a reference to the given []ChannelPermissionOverwriteRequest and assigns it to the PermissionOverwrites field.
func (o *CreateGuildChannelRequest) SetPermissionOverwrites(v []ChannelPermissionOverwriteRequest) {
	o.PermissionOverwrites = v
}


// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion.Get()) {
		var ret string
		return ret
	}
	return *o.RtcRegion.Get()
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion.Get()) {
		return nil, false
	}
	return o.RtcRegion.Get(), o.RtcRegion.IsSet()
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasRtcRegion() bool {
	if o != nil && o.RtcRegion.IsSet() {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given NullableString and assigns it to the RtcRegion field.
func (o *CreateGuildChannelRequest) SetRtcRegion(v string) {
	o.RtcRegion.Set(&v)
}

// SetRtcRegionNil sets the value for RtcRegion to be an explicit nil
func (o *CreateGuildChannelRequest) SetRtcRegionNil() {
	o.RtcRegion.Set(nil)
}

// UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetRtcRegion() {
	o.RtcRegion.Unset()
}

// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetVideoQualityMode() int32 {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret int32
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetVideoQualityModeOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoQualityMode) {
		return nil, false
	}
	return o.VideoQualityMode, true
}

// HasVideoQualityMode returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasVideoQualityMode() bool {
	if o != nil && !IsNil(o.VideoQualityMode) {
		return true
	}

	return false
}

// SetVideoQualityMode gets a reference to the given int32 and assigns it to the VideoQualityMode field.
func (o *CreateGuildChannelRequest) SetVideoQualityMode(v int32) {
	o.VideoQualityMode = &v
}


// GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultAutoArchiveDuration() int32 {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		var ret int32
		return ret
	}
	return *o.DefaultAutoArchiveDuration
}

// GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultAutoArchiveDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		return nil, false
	}
	return o.DefaultAutoArchiveDuration, true
}

// HasDefaultAutoArchiveDuration returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.DefaultAutoArchiveDuration) {
		return true
	}

	return false
}

// SetDefaultAutoArchiveDuration gets a reference to the given int32 and assigns it to the DefaultAutoArchiveDuration field.
func (o *CreateGuildChannelRequest) SetDefaultAutoArchiveDuration(v int32) {
	o.DefaultAutoArchiveDuration = &v
}


// GetDefaultReactionEmoji returns the DefaultReactionEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetDefaultReactionEmoji() UpdateDefaultReactionEmojiRequest {
	if o == nil || IsNil(o.DefaultReactionEmoji.Get()) {
		var ret UpdateDefaultReactionEmojiRequest
		return ret
	}
	return *o.DefaultReactionEmoji.Get()
}

// GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetDefaultReactionEmojiOk() (*UpdateDefaultReactionEmojiRequest, bool) {
	if o == nil || IsNil(o.DefaultReactionEmoji.Get()) {
		return nil, false
	}
	return o.DefaultReactionEmoji.Get(), o.DefaultReactionEmoji.IsSet()
}

// HasDefaultReactionEmoji returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultReactionEmoji() bool {
	if o != nil && o.DefaultReactionEmoji.IsSet() {
		return true
	}

	return false
}

// SetDefaultReactionEmoji gets a reference to the given NullableUpdateDefaultReactionEmojiRequest and assigns it to the DefaultReactionEmoji field.
func (o *CreateGuildChannelRequest) SetDefaultReactionEmoji(v UpdateDefaultReactionEmojiRequest) {
	o.DefaultReactionEmoji.Set(&v)
}

// SetDefaultReactionEmojiNil sets the value for DefaultReactionEmoji to be an explicit nil
func (o *CreateGuildChannelRequest) SetDefaultReactionEmojiNil() {
	o.DefaultReactionEmoji.Set(nil)
}

// UnsetDefaultReactionEmoji ensures that no value is present for DefaultReactionEmoji, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetDefaultReactionEmoji() {
	o.DefaultReactionEmoji.Unset()
}

// GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetDefaultThreadRateLimitPerUser() int32 {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultThreadRateLimitPerUser.Get()
}

// GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetDefaultThreadRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser.Get()) {
		return nil, false
	}
	return o.DefaultThreadRateLimitPerUser.Get(), o.DefaultThreadRateLimitPerUser.IsSet()
}

// HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultThreadRateLimitPerUser() bool {
	if o != nil && o.DefaultThreadRateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetDefaultThreadRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the DefaultThreadRateLimitPerUser field.
func (o *CreateGuildChannelRequest) SetDefaultThreadRateLimitPerUser(v int32) {
	o.DefaultThreadRateLimitPerUser.Set(&v)
}

// SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil
func (o *CreateGuildChannelRequest) SetDefaultThreadRateLimitPerUserNil() {
	o.DefaultThreadRateLimitPerUser.Set(nil)
}

// UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
func (o *CreateGuildChannelRequest) UnsetDefaultThreadRateLimitPerUser() {
	o.DefaultThreadRateLimitPerUser.Unset()
}

// GetDefaultSortOrder returns the DefaultSortOrder field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultSortOrder() int32 {
	if o == nil || IsNil(o.DefaultSortOrder) {
		var ret int32
		return ret
	}
	return *o.DefaultSortOrder
}

// GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultSortOrder) {
		return nil, false
	}
	return o.DefaultSortOrder, true
}

// HasDefaultSortOrder returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultSortOrder() bool {
	if o != nil && !IsNil(o.DefaultSortOrder) {
		return true
	}

	return false
}

// SetDefaultSortOrder gets a reference to the given int32 and assigns it to the DefaultSortOrder field.
func (o *CreateGuildChannelRequest) SetDefaultSortOrder(v int32) {
	o.DefaultSortOrder = &v
}


// GetDefaultForumLayout returns the DefaultForumLayout field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultForumLayout() int32 {
	if o == nil || IsNil(o.DefaultForumLayout) {
		var ret int32
		return ret
	}
	return *o.DefaultForumLayout
}

// GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultForumLayoutOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultForumLayout) {
		return nil, false
	}
	return o.DefaultForumLayout, true
}

// HasDefaultForumLayout returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultForumLayout() bool {
	if o != nil && !IsNil(o.DefaultForumLayout) {
		return true
	}

	return false
}

// SetDefaultForumLayout gets a reference to the given int32 and assigns it to the DefaultForumLayout field.
func (o *CreateGuildChannelRequest) SetDefaultForumLayout(v int32) {
	o.DefaultForumLayout = &v
}


// GetDefaultTagSetting returns the DefaultTagSetting field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultTagSetting() string {
	if o == nil || IsNil(o.DefaultTagSetting) {
		var ret string
		return ret
	}
	return *o.DefaultTagSetting
}

// GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultTagSettingOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTagSetting) {
		return nil, false
	}
	return o.DefaultTagSetting, true
}

// HasDefaultTagSetting returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultTagSetting() bool {
	if o != nil && !IsNil(o.DefaultTagSetting) {
		return true
	}

	return false
}

// SetDefaultTagSetting gets a reference to the given string and assigns it to the DefaultTagSetting field.
func (o *CreateGuildChannelRequest) SetDefaultTagSetting(v string) {
	o.DefaultTagSetting = &v
}


// GetAvailableTags returns the AvailableTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildChannelRequest) GetAvailableTags() []*CreateOrUpdateThreadTagRequest {
	if o == nil {
		var ret []*CreateOrUpdateThreadTagRequest
		return ret
	}
	return o.AvailableTags
}

// GetAvailableTagsOk returns a tuple with the AvailableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildChannelRequest) GetAvailableTagsOk() ([]*CreateOrUpdateThreadTagRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableTags, true
}

// HasAvailableTags returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasAvailableTags() bool {
	if o != nil && !IsNil(o.AvailableTags) {
		return true
	}

	return false
}

// SetAvailableTags gets a reference to the given []*CreateOrUpdateThreadTagRequest and assigns it to the AvailableTags field.
func (o *CreateGuildChannelRequest) SetAvailableTags(v []*CreateOrUpdateThreadTagRequest) {
	o.AvailableTags = v
}


func (o CreateGuildChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGuildChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	toSerialize["name"] = o.Name
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	if o.Bitrate.IsSet() {
		toSerialize["bitrate"] = o.Bitrate.Get()
	}
	if o.UserLimit.IsSet() {
		toSerialize["user_limit"] = o.UserLimit.Get()
	}
	if o.Nsfw.IsSet() {
		toSerialize["nsfw"] = o.Nsfw.Get()
	}
	if o.RateLimitPerUser.IsSet() {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser.Get()
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.PermissionOverwrites != nil {
		toSerialize["permission_overwrites"] = o.PermissionOverwrites
	}
	if o.RtcRegion.IsSet() {
		toSerialize["rtc_region"] = o.RtcRegion.Get()
	}
	if !IsNil(o.VideoQualityMode) {
		toSerialize["video_quality_mode"] = o.VideoQualityMode
	}
	if !IsNil(o.DefaultAutoArchiveDuration) {
		toSerialize["default_auto_archive_duration"] = o.DefaultAutoArchiveDuration
	}
	if o.DefaultReactionEmoji.IsSet() {
		toSerialize["default_reaction_emoji"] = o.DefaultReactionEmoji.Get()
	}
	if o.DefaultThreadRateLimitPerUser.IsSet() {
		toSerialize["default_thread_rate_limit_per_user"] = o.DefaultThreadRateLimitPerUser.Get()
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
	if o.AvailableTags != nil {
		toSerialize["available_tags"] = o.AvailableTags
	}
	return toSerialize, nil
}

func (o *CreateGuildChannelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateGuildChannelRequest := _CreateGuildChannelRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGuildChannelRequest)

	if err != nil {
		return err
	}

	*o = CreateGuildChannelRequest(varCreateGuildChannelRequest)

	return err
}

type NullableCreateGuildChannelRequest struct {
	value *CreateGuildChannelRequest
	isSet bool
}

func (v NullableCreateGuildChannelRequest) Get() *CreateGuildChannelRequest {
	return v.value
}

func (v *NullableCreateGuildChannelRequest) Set(val *CreateGuildChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildChannelRequest(val *CreateGuildChannelRequest) *NullableCreateGuildChannelRequest {
	return &NullableCreateGuildChannelRequest{value: val, isSet: true}
}

func (v NullableCreateGuildChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


