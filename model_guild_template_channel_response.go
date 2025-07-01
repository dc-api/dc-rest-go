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

// checks if the GuildTemplateChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateChannelResponse{}

// GuildTemplateChannelResponse struct for GuildTemplateChannelResponse
type GuildTemplateChannelResponse struct {
	Id NullableInt32 `json:"id,omitempty"`
	Type int32 `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Position NullableInt32 `json:"position,omitempty"`
	Topic NullableString `json:"topic,omitempty"`
	Bitrate int32 `json:"bitrate"`
	UserLimit int32 `json:"user_limit"`
	Nsfw bool `json:"nsfw"`
	RateLimitPerUser int32 `json:"rate_limit_per_user"`
	ParentId *string `json:"parent_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	DefaultAutoArchiveDuration *int32 `json:"default_auto_archive_duration,omitempty"`
	PermissionOverwrites []*ChannelPermissionOverwriteResponse `json:"permission_overwrites"`
	AvailableTags []GuildTemplateChannelTags `json:"available_tags,omitempty"`
	Template string `json:"template"`
	DefaultReactionEmoji NullableDefaultReactionEmojiResponse `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser NullableInt32 `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder *int32 `json:"default_sort_order,omitempty"`
	DefaultForumLayout *int32 `json:"default_forum_layout,omitempty"`
	DefaultTagSetting *string `json:"default_tag_setting,omitempty"`
	IconEmoji map[string]interface{} `json:"icon_emoji,omitempty"`
	ThemeColor NullableInt32 `json:"theme_color,omitempty"`
}

type _GuildTemplateChannelResponse GuildTemplateChannelResponse

// NewGuildTemplateChannelResponse instantiates a new GuildTemplateChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateChannelResponse(type_ int32, bitrate int32, userLimit int32, nsfw bool, rateLimitPerUser int32, permissionOverwrites []*ChannelPermissionOverwriteResponse, template string) *GuildTemplateChannelResponse {
	this := GuildTemplateChannelResponse{}
	this.Type = type_
	this.Bitrate = bitrate
	this.UserLimit = userLimit
	this.Nsfw = nsfw
	this.RateLimitPerUser = rateLimitPerUser
	this.PermissionOverwrites = permissionOverwrites
	this.Template = template
	return &this
}

// NewGuildTemplateChannelResponseWithDefaults instantiates a new GuildTemplateChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildTemplateChannelResponseWithDefaults() *GuildTemplateChannelResponse {
	this := GuildTemplateChannelResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id.Get()) {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GuildTemplateChannelResponse) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *GuildTemplateChannelResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value
func (o *GuildTemplateChannelResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GuildTemplateChannelResponse) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GuildTemplateChannelResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *GuildTemplateChannelResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetName() {
	o.Name.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetPosition() int32 {
	if o == nil || IsNil(o.Position.Get()) {
		var ret int32
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position.Get()) {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableInt32 and assigns it to the Position field.
func (o *GuildTemplateChannelResponse) SetPosition(v int32) {
	o.Position.Set(&v)
}

// SetPositionNil sets the value for Position to be an explicit nil
func (o *GuildTemplateChannelResponse) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetPosition() {
	o.Position.Unset()
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic.Get()) {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *GuildTemplateChannelResponse) SetTopic(v string) {
	o.Topic.Set(&v)
}

// SetTopicNil sets the value for Topic to be an explicit nil
func (o *GuildTemplateChannelResponse) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetTopic() {
	o.Topic.Unset()
}

// GetBitrate returns the Bitrate field value
func (o *GuildTemplateChannelResponse) GetBitrate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bitrate, true
}

// SetBitrate sets field value
func (o *GuildTemplateChannelResponse) SetBitrate(v int32) {
	o.Bitrate = v
}

// GetUserLimit returns the UserLimit field value
func (o *GuildTemplateChannelResponse) GetUserLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetUserLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserLimit, true
}

// SetUserLimit sets field value
func (o *GuildTemplateChannelResponse) SetUserLimit(v int32) {
	o.UserLimit = v
}

// GetNsfw returns the Nsfw field value
func (o *GuildTemplateChannelResponse) GetNsfw() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetNsfwOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nsfw, true
}

// SetNsfw sets field value
func (o *GuildTemplateChannelResponse) SetNsfw(v bool) {
	o.Nsfw = v
}

// GetRateLimitPerUser returns the RateLimitPerUser field value
func (o *GuildTemplateChannelResponse) GetRateLimitPerUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RateLimitPerUser
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimitPerUser, true
}

// SetRateLimitPerUser sets field value
func (o *GuildTemplateChannelResponse) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *GuildTemplateChannelResponse) SetParentId(v string) {
	o.ParentId = &v
}


// GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultAutoArchiveDuration() int32 {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		var ret int32
		return ret
	}
	return *o.DefaultAutoArchiveDuration
}

// GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultAutoArchiveDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		return nil, false
	}
	return o.DefaultAutoArchiveDuration, true
}

// HasDefaultAutoArchiveDuration returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.DefaultAutoArchiveDuration) {
		return true
	}

	return false
}

// SetDefaultAutoArchiveDuration gets a reference to the given int32 and assigns it to the DefaultAutoArchiveDuration field.
func (o *GuildTemplateChannelResponse) SetDefaultAutoArchiveDuration(v int32) {
	o.DefaultAutoArchiveDuration = &v
}


// GetPermissionOverwrites returns the PermissionOverwrites field value
func (o *GuildTemplateChannelResponse) GetPermissionOverwrites() []*ChannelPermissionOverwriteResponse {
	if o == nil {
		var ret []*ChannelPermissionOverwriteResponse
		return ret
	}

	return o.PermissionOverwrites
}

// GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetPermissionOverwritesOk() ([]*ChannelPermissionOverwriteResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionOverwrites, true
}

// SetPermissionOverwrites sets field value
func (o *GuildTemplateChannelResponse) SetPermissionOverwrites(v []*ChannelPermissionOverwriteResponse) {
	o.PermissionOverwrites = v
}

// GetAvailableTags returns the AvailableTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetAvailableTags() []GuildTemplateChannelTags {
	if o == nil {
		var ret []GuildTemplateChannelTags
		return ret
	}
	return o.AvailableTags
}

// GetAvailableTagsOk returns a tuple with the AvailableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetAvailableTagsOk() ([]GuildTemplateChannelTags, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableTags, true
}

// HasAvailableTags returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasAvailableTags() bool {
	if o != nil && !IsNil(o.AvailableTags) {
		return true
	}

	return false
}

// SetAvailableTags gets a reference to the given []GuildTemplateChannelTags and assigns it to the AvailableTags field.
func (o *GuildTemplateChannelResponse) SetAvailableTags(v []GuildTemplateChannelTags) {
	o.AvailableTags = v
}


// GetTemplate returns the Template field value
func (o *GuildTemplateChannelResponse) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *GuildTemplateChannelResponse) SetTemplate(v string) {
	o.Template = v
}

// GetDefaultReactionEmoji returns the DefaultReactionEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetDefaultReactionEmoji() DefaultReactionEmojiResponse {
	if o == nil || IsNil(o.DefaultReactionEmoji.Get()) {
		var ret DefaultReactionEmojiResponse
		return ret
	}
	return *o.DefaultReactionEmoji.Get()
}

// GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetDefaultReactionEmojiOk() (*DefaultReactionEmojiResponse, bool) {
	if o == nil || IsNil(o.DefaultReactionEmoji.Get()) {
		return nil, false
	}
	return o.DefaultReactionEmoji.Get(), o.DefaultReactionEmoji.IsSet()
}

// HasDefaultReactionEmoji returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultReactionEmoji() bool {
	if o != nil && o.DefaultReactionEmoji.IsSet() {
		return true
	}

	return false
}

// SetDefaultReactionEmoji gets a reference to the given NullableDefaultReactionEmojiResponse and assigns it to the DefaultReactionEmoji field.
func (o *GuildTemplateChannelResponse) SetDefaultReactionEmoji(v DefaultReactionEmojiResponse) {
	o.DefaultReactionEmoji.Set(&v)
}

// SetDefaultReactionEmojiNil sets the value for DefaultReactionEmoji to be an explicit nil
func (o *GuildTemplateChannelResponse) SetDefaultReactionEmojiNil() {
	o.DefaultReactionEmoji.Set(nil)
}

// UnsetDefaultReactionEmoji ensures that no value is present for DefaultReactionEmoji, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetDefaultReactionEmoji() {
	o.DefaultReactionEmoji.Unset()
}

// GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetDefaultThreadRateLimitPerUser() int32 {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultThreadRateLimitPerUser.Get()
}

// GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetDefaultThreadRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser.Get()) {
		return nil, false
	}
	return o.DefaultThreadRateLimitPerUser.Get(), o.DefaultThreadRateLimitPerUser.IsSet()
}

// HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultThreadRateLimitPerUser() bool {
	if o != nil && o.DefaultThreadRateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetDefaultThreadRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the DefaultThreadRateLimitPerUser field.
func (o *GuildTemplateChannelResponse) SetDefaultThreadRateLimitPerUser(v int32) {
	o.DefaultThreadRateLimitPerUser.Set(&v)
}

// SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil
func (o *GuildTemplateChannelResponse) SetDefaultThreadRateLimitPerUserNil() {
	o.DefaultThreadRateLimitPerUser.Set(nil)
}

// UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetDefaultThreadRateLimitPerUser() {
	o.DefaultThreadRateLimitPerUser.Unset()
}

// GetDefaultSortOrder returns the DefaultSortOrder field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultSortOrder() int32 {
	if o == nil || IsNil(o.DefaultSortOrder) {
		var ret int32
		return ret
	}
	return *o.DefaultSortOrder
}

// GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultSortOrder) {
		return nil, false
	}
	return o.DefaultSortOrder, true
}

// HasDefaultSortOrder returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultSortOrder() bool {
	if o != nil && !IsNil(o.DefaultSortOrder) {
		return true
	}

	return false
}

// SetDefaultSortOrder gets a reference to the given int32 and assigns it to the DefaultSortOrder field.
func (o *GuildTemplateChannelResponse) SetDefaultSortOrder(v int32) {
	o.DefaultSortOrder = &v
}


// GetDefaultForumLayout returns the DefaultForumLayout field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultForumLayout() int32 {
	if o == nil || IsNil(o.DefaultForumLayout) {
		var ret int32
		return ret
	}
	return *o.DefaultForumLayout
}

// GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultForumLayoutOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultForumLayout) {
		return nil, false
	}
	return o.DefaultForumLayout, true
}

// HasDefaultForumLayout returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultForumLayout() bool {
	if o != nil && !IsNil(o.DefaultForumLayout) {
		return true
	}

	return false
}

// SetDefaultForumLayout gets a reference to the given int32 and assigns it to the DefaultForumLayout field.
func (o *GuildTemplateChannelResponse) SetDefaultForumLayout(v int32) {
	o.DefaultForumLayout = &v
}


// GetDefaultTagSetting returns the DefaultTagSetting field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultTagSetting() string {
	if o == nil || IsNil(o.DefaultTagSetting) {
		var ret string
		return ret
	}
	return *o.DefaultTagSetting
}

// GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultTagSettingOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTagSetting) {
		return nil, false
	}
	return o.DefaultTagSetting, true
}

// HasDefaultTagSetting returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultTagSetting() bool {
	if o != nil && !IsNil(o.DefaultTagSetting) {
		return true
	}

	return false
}

// SetDefaultTagSetting gets a reference to the given string and assigns it to the DefaultTagSetting field.
func (o *GuildTemplateChannelResponse) SetDefaultTagSetting(v string) {
	o.DefaultTagSetting = &v
}


// GetIconEmoji returns the IconEmoji field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetIconEmoji() map[string]interface{} {
	if o == nil || IsNil(o.IconEmoji) {
		var ret map[string]interface{}
		return ret
	}
	return o.IconEmoji
}

// GetIconEmojiOk returns a tuple with the IconEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetIconEmojiOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IconEmoji) {
		return map[string]interface{}{}, false
	}
	return o.IconEmoji, true
}

// HasIconEmoji returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasIconEmoji() bool {
	if o != nil && !IsNil(o.IconEmoji) {
		return true
	}

	return false
}

// SetIconEmoji gets a reference to the given map[string]interface{} and assigns it to the IconEmoji field.
func (o *GuildTemplateChannelResponse) SetIconEmoji(v map[string]interface{}) {
	o.IconEmoji = v
}


// GetThemeColor returns the ThemeColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateChannelResponse) GetThemeColor() int32 {
	if o == nil || IsNil(o.ThemeColor.Get()) {
		var ret int32
		return ret
	}
	return *o.ThemeColor.Get()
}

// GetThemeColorOk returns a tuple with the ThemeColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateChannelResponse) GetThemeColorOk() (*int32, bool) {
	if o == nil || IsNil(o.ThemeColor.Get()) {
		return nil, false
	}
	return o.ThemeColor.Get(), o.ThemeColor.IsSet()
}

// HasThemeColor returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasThemeColor() bool {
	if o != nil && o.ThemeColor.IsSet() {
		return true
	}

	return false
}

// SetThemeColor gets a reference to the given NullableInt32 and assigns it to the ThemeColor field.
func (o *GuildTemplateChannelResponse) SetThemeColor(v int32) {
	o.ThemeColor.Set(&v)
}

// SetThemeColorNil sets the value for ThemeColor to be an explicit nil
func (o *GuildTemplateChannelResponse) SetThemeColorNil() {
	o.ThemeColor.Set(nil)
}

// UnsetThemeColor ensures that no value is present for ThemeColor, not even an explicit nil
func (o *GuildTemplateChannelResponse) UnsetThemeColor() {
	o.ThemeColor.Unset()
}

func (o GuildTemplateChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildTemplateChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["type"] = o.Type
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	toSerialize["bitrate"] = o.Bitrate
	toSerialize["user_limit"] = o.UserLimit
	toSerialize["nsfw"] = o.Nsfw
	toSerialize["rate_limit_per_user"] = o.RateLimitPerUser
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.DefaultAutoArchiveDuration) {
		toSerialize["default_auto_archive_duration"] = o.DefaultAutoArchiveDuration
	}
	toSerialize["permission_overwrites"] = o.PermissionOverwrites
	if o.AvailableTags != nil {
		toSerialize["available_tags"] = o.AvailableTags
	}
	toSerialize["template"] = o.Template
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
	if !IsNil(o.IconEmoji) {
		toSerialize["icon_emoji"] = o.IconEmoji
	}
	if o.ThemeColor.IsSet() {
		toSerialize["theme_color"] = o.ThemeColor.Get()
	}
	return toSerialize, nil
}

func (o *GuildTemplateChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"bitrate",
		"user_limit",
		"nsfw",
		"rate_limit_per_user",
		"permission_overwrites",
		"template",
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

	varGuildTemplateChannelResponse := _GuildTemplateChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildTemplateChannelResponse)

	if err != nil {
		return err
	}

	*o = GuildTemplateChannelResponse(varGuildTemplateChannelResponse)

	return err
}

type NullableGuildTemplateChannelResponse struct {
	value *GuildTemplateChannelResponse
	isSet bool
}

func (v NullableGuildTemplateChannelResponse) Get() *GuildTemplateChannelResponse {
	return v.value
}

func (v *NullableGuildTemplateChannelResponse) Set(val *GuildTemplateChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildTemplateChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildTemplateChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildTemplateChannelResponse(val *GuildTemplateChannelResponse) *NullableGuildTemplateChannelResponse {
	return &NullableGuildTemplateChannelResponse{value: val, isSet: true}
}

func (v NullableGuildTemplateChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildTemplateChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


