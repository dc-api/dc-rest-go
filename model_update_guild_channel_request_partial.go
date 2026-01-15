/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
)

// checks if the UpdateGuildChannelRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGuildChannelRequestPartial{}

// UpdateGuildChannelRequestPartial struct for UpdateGuildChannelRequestPartial
type UpdateGuildChannelRequestPartial struct {
	Type *CreateGuildChannelRequestType `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Topic *string `json:"topic,omitempty"`
	Bitrate *int32 `json:"bitrate,omitempty"`
	UserLimit *int32 `json:"user_limit,omitempty"`
	Nsfw *bool `json:"nsfw,omitempty"`
	RateLimitPerUser *int32 `json:"rate_limit_per_user,omitempty"`
	ParentId *GetEntitlementsSkuIdsParameterOneOfInner `json:"parent_id,omitempty"`
	PermissionOverwrites []ChannelPermissionOverwriteRequest `json:"permission_overwrites,omitempty"`
	RtcRegion *string `json:"rtc_region,omitempty"`
	VideoQualityMode *CreateGuildChannelRequestVideoQualityMode `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration *CreateForumThreadRequestAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	DefaultReactionEmoji *CreateGuildChannelRequestDefaultReactionEmoji `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int32 `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder *CreateGuildChannelRequestDefaultSortOrder `json:"default_sort_order,omitempty"`
	DefaultForumLayout *CreateGuildChannelRequestDefaultForumLayout `json:"default_forum_layout,omitempty"`
	DefaultTagSetting *CreateGuildChannelRequestDefaultTagSetting `json:"default_tag_setting,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
	AvailableTags []UpdateThreadTagRequest `json:"available_tags,omitempty"`
}

// NewUpdateGuildChannelRequestPartial instantiates a new UpdateGuildChannelRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGuildChannelRequestPartial() *UpdateGuildChannelRequestPartial {
	this := UpdateGuildChannelRequestPartial{}
	return &this
}

// NewUpdateGuildChannelRequestPartialWithDefaults instantiates a new UpdateGuildChannelRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGuildChannelRequestPartialWithDefaults() *UpdateGuildChannelRequestPartial {
	this := UpdateGuildChannelRequestPartial{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetType() CreateGuildChannelRequestType {
	if o == nil || IsNil(o.Type) {
		var ret CreateGuildChannelRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetTypeOk() (*CreateGuildChannelRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CreateGuildChannelRequestType and assigns it to the Type field.
func (o *UpdateGuildChannelRequestPartial) SetType(v CreateGuildChannelRequestType) {
	o.Type = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGuildChannelRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetPosition returns the Position field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *UpdateGuildChannelRequestPartial) SetPosition(v int32) {
	o.Position = &v
}


// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *UpdateGuildChannelRequestPartial) SetTopic(v string) {
	o.Topic = &v
}


// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate) {
		var ret int32
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate) {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasBitrate() bool {
	if o != nil && !IsNil(o.Bitrate) {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.
func (o *UpdateGuildChannelRequestPartial) SetBitrate(v int32) {
	o.Bitrate = &v
}


// GetUserLimit returns the UserLimit field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit) {
		var ret int32
		return ret
	}
	return *o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit) {
		return nil, false
	}
	return o.UserLimit, true
}

// HasUserLimit returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasUserLimit() bool {
	if o != nil && !IsNil(o.UserLimit) {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.
func (o *UpdateGuildChannelRequestPartial) SetUserLimit(v int32) {
	o.UserLimit = &v
}


// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *UpdateGuildChannelRequestPartial) SetNsfw(v bool) {
	o.Nsfw = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser) {
		return nil, false
	}
	return o.RateLimitPerUser, true
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasRateLimitPerUser() bool {
	if o != nil && !IsNil(o.RateLimitPerUser) {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given int32 and assigns it to the RateLimitPerUser field.
func (o *UpdateGuildChannelRequestPartial) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser = &v
}


// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ParentId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ParentId field.
func (o *UpdateGuildChannelRequestPartial) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ParentId = &v
}


// GetPermissionOverwrites returns the PermissionOverwrites field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest {
	if o == nil || IsNil(o.PermissionOverwrites) {
		var ret []ChannelPermissionOverwriteRequest
		return ret
	}
	return o.PermissionOverwrites
}

// GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetPermissionOverwritesOk() ([]ChannelPermissionOverwriteRequest, bool) {
	if o == nil || IsNil(o.PermissionOverwrites) {
		return nil, false
	}
	return o.PermissionOverwrites, true
}

// HasPermissionOverwrites returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasPermissionOverwrites() bool {
	if o != nil && !IsNil(o.PermissionOverwrites) {
		return true
	}

	return false
}

// SetPermissionOverwrites gets a reference to the given []ChannelPermissionOverwriteRequest and assigns it to the PermissionOverwrites field.
func (o *UpdateGuildChannelRequestPartial) SetPermissionOverwrites(v []ChannelPermissionOverwriteRequest) {
	o.PermissionOverwrites = v
}


// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion) {
		var ret string
		return ret
	}
	return *o.RtcRegion
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion) {
		return nil, false
	}
	return o.RtcRegion, true
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasRtcRegion() bool {
	if o != nil && !IsNil(o.RtcRegion) {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given string and assigns it to the RtcRegion field.
func (o *UpdateGuildChannelRequestPartial) SetRtcRegion(v string) {
	o.RtcRegion = &v
}


// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetVideoQualityMode() CreateGuildChannelRequestVideoQualityMode {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret CreateGuildChannelRequestVideoQualityMode
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetVideoQualityModeOk() (*CreateGuildChannelRequestVideoQualityMode, bool) {
	if o == nil || IsNil(o.VideoQualityMode) {
		return nil, false
	}
	return o.VideoQualityMode, true
}

// HasVideoQualityMode returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasVideoQualityMode() bool {
	if o != nil && !IsNil(o.VideoQualityMode) {
		return true
	}

	return false
}

// SetVideoQualityMode gets a reference to the given CreateGuildChannelRequestVideoQualityMode and assigns it to the VideoQualityMode field.
func (o *UpdateGuildChannelRequestPartial) SetVideoQualityMode(v CreateGuildChannelRequestVideoQualityMode) {
	o.VideoQualityMode = &v
}


// GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetDefaultAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		var ret CreateForumThreadRequestAutoArchiveDuration
		return ret
	}
	return *o.DefaultAutoArchiveDuration
}

// GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetDefaultAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool) {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		return nil, false
	}
	return o.DefaultAutoArchiveDuration, true
}

// HasDefaultAutoArchiveDuration returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasDefaultAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.DefaultAutoArchiveDuration) {
		return true
	}

	return false
}

// SetDefaultAutoArchiveDuration gets a reference to the given CreateForumThreadRequestAutoArchiveDuration and assigns it to the DefaultAutoArchiveDuration field.
func (o *UpdateGuildChannelRequestPartial) SetDefaultAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration) {
	o.DefaultAutoArchiveDuration = &v
}


// GetDefaultReactionEmoji returns the DefaultReactionEmoji field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetDefaultReactionEmoji() CreateGuildChannelRequestDefaultReactionEmoji {
	if o == nil || IsNil(o.DefaultReactionEmoji) {
		var ret CreateGuildChannelRequestDefaultReactionEmoji
		return ret
	}
	return *o.DefaultReactionEmoji
}

// GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetDefaultReactionEmojiOk() (*CreateGuildChannelRequestDefaultReactionEmoji, bool) {
	if o == nil || IsNil(o.DefaultReactionEmoji) {
		return nil, false
	}
	return o.DefaultReactionEmoji, true
}

// HasDefaultReactionEmoji returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasDefaultReactionEmoji() bool {
	if o != nil && !IsNil(o.DefaultReactionEmoji) {
		return true
	}

	return false
}

// SetDefaultReactionEmoji gets a reference to the given CreateGuildChannelRequestDefaultReactionEmoji and assigns it to the DefaultReactionEmoji field.
func (o *UpdateGuildChannelRequestPartial) SetDefaultReactionEmoji(v CreateGuildChannelRequestDefaultReactionEmoji) {
	o.DefaultReactionEmoji = &v
}


// GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetDefaultThreadRateLimitPerUser() int32 {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.DefaultThreadRateLimitPerUser
}

// GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetDefaultThreadRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser) {
		return nil, false
	}
	return o.DefaultThreadRateLimitPerUser, true
}

// HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasDefaultThreadRateLimitPerUser() bool {
	if o != nil && !IsNil(o.DefaultThreadRateLimitPerUser) {
		return true
	}

	return false
}

// SetDefaultThreadRateLimitPerUser gets a reference to the given int32 and assigns it to the DefaultThreadRateLimitPerUser field.
func (o *UpdateGuildChannelRequestPartial) SetDefaultThreadRateLimitPerUser(v int32) {
	o.DefaultThreadRateLimitPerUser = &v
}


// GetDefaultSortOrder returns the DefaultSortOrder field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder {
	if o == nil || IsNil(o.DefaultSortOrder) {
		var ret CreateGuildChannelRequestDefaultSortOrder
		return ret
	}
	return *o.DefaultSortOrder
}

// GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool) {
	if o == nil || IsNil(o.DefaultSortOrder) {
		return nil, false
	}
	return o.DefaultSortOrder, true
}

// HasDefaultSortOrder returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasDefaultSortOrder() bool {
	if o != nil && !IsNil(o.DefaultSortOrder) {
		return true
	}

	return false
}

// SetDefaultSortOrder gets a reference to the given CreateGuildChannelRequestDefaultSortOrder and assigns it to the DefaultSortOrder field.
func (o *UpdateGuildChannelRequestPartial) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder) {
	o.DefaultSortOrder = &v
}


// GetDefaultForumLayout returns the DefaultForumLayout field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout {
	if o == nil || IsNil(o.DefaultForumLayout) {
		var ret CreateGuildChannelRequestDefaultForumLayout
		return ret
	}
	return *o.DefaultForumLayout
}

// GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool) {
	if o == nil || IsNil(o.DefaultForumLayout) {
		return nil, false
	}
	return o.DefaultForumLayout, true
}

// HasDefaultForumLayout returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasDefaultForumLayout() bool {
	if o != nil && !IsNil(o.DefaultForumLayout) {
		return true
	}

	return false
}

// SetDefaultForumLayout gets a reference to the given CreateGuildChannelRequestDefaultForumLayout and assigns it to the DefaultForumLayout field.
func (o *UpdateGuildChannelRequestPartial) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout) {
	o.DefaultForumLayout = &v
}


// GetDefaultTagSetting returns the DefaultTagSetting field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting {
	if o == nil || IsNil(o.DefaultTagSetting) {
		var ret CreateGuildChannelRequestDefaultTagSetting
		return ret
	}
	return *o.DefaultTagSetting
}

// GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool) {
	if o == nil || IsNil(o.DefaultTagSetting) {
		return nil, false
	}
	return o.DefaultTagSetting, true
}

// HasDefaultTagSetting returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasDefaultTagSetting() bool {
	if o != nil && !IsNil(o.DefaultTagSetting) {
		return true
	}

	return false
}

// SetDefaultTagSetting gets a reference to the given CreateGuildChannelRequestDefaultTagSetting and assigns it to the DefaultTagSetting field.
func (o *UpdateGuildChannelRequestPartial) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting) {
	o.DefaultTagSetting = &v
}


// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *UpdateGuildChannelRequestPartial) SetFlags(v int32) {
	o.Flags = &v
}


// GetAvailableTags returns the AvailableTags field value if set, zero value otherwise.
func (o *UpdateGuildChannelRequestPartial) GetAvailableTags() []UpdateThreadTagRequest {
	if o == nil || IsNil(o.AvailableTags) {
		var ret []UpdateThreadTagRequest
		return ret
	}
	return o.AvailableTags
}

// GetAvailableTagsOk returns a tuple with the AvailableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildChannelRequestPartial) GetAvailableTagsOk() ([]UpdateThreadTagRequest, bool) {
	if o == nil || IsNil(o.AvailableTags) {
		return nil, false
	}
	return o.AvailableTags, true
}

// HasAvailableTags returns a boolean if a field has been set.
func (o *UpdateGuildChannelRequestPartial) HasAvailableTags() bool {
	if o != nil && !IsNil(o.AvailableTags) {
		return true
	}

	return false
}

// SetAvailableTags gets a reference to the given []UpdateThreadTagRequest and assigns it to the AvailableTags field.
func (o *UpdateGuildChannelRequestPartial) SetAvailableTags(v []UpdateThreadTagRequest) {
	o.AvailableTags = v
}


func (o UpdateGuildChannelRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGuildChannelRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Bitrate) {
		toSerialize["bitrate"] = o.Bitrate
	}
	if !IsNil(o.UserLimit) {
		toSerialize["user_limit"] = o.UserLimit
	}
	if !IsNil(o.Nsfw) {
		toSerialize["nsfw"] = o.Nsfw
	}
	if !IsNil(o.RateLimitPerUser) {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.PermissionOverwrites) {
		toSerialize["permission_overwrites"] = o.PermissionOverwrites
	}
	if !IsNil(o.RtcRegion) {
		toSerialize["rtc_region"] = o.RtcRegion
	}
	if !IsNil(o.VideoQualityMode) {
		toSerialize["video_quality_mode"] = o.VideoQualityMode
	}
	if !IsNil(o.DefaultAutoArchiveDuration) {
		toSerialize["default_auto_archive_duration"] = o.DefaultAutoArchiveDuration
	}
	if !IsNil(o.DefaultReactionEmoji) {
		toSerialize["default_reaction_emoji"] = o.DefaultReactionEmoji
	}
	if !IsNil(o.DefaultThreadRateLimitPerUser) {
		toSerialize["default_thread_rate_limit_per_user"] = o.DefaultThreadRateLimitPerUser
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
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.AvailableTags) {
		toSerialize["available_tags"] = o.AvailableTags
	}
	return toSerialize, nil
}

type NullableUpdateGuildChannelRequestPartial struct {
	value *UpdateGuildChannelRequestPartial
	isSet bool
}

func (v NullableUpdateGuildChannelRequestPartial) Get() *UpdateGuildChannelRequestPartial {
	return v.value
}

func (v *NullableUpdateGuildChannelRequestPartial) Set(val *UpdateGuildChannelRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildChannelRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildChannelRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildChannelRequestPartial(val *UpdateGuildChannelRequestPartial) *NullableUpdateGuildChannelRequestPartial {
	return &NullableUpdateGuildChannelRequestPartial{value: val, isSet: true}
}

func (v NullableUpdateGuildChannelRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildChannelRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

