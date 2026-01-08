/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the CreateGuildChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildChannelRequest{}

// CreateGuildChannelRequest struct for CreateGuildChannelRequest
type CreateGuildChannelRequest struct {
	Type *CreateGuildChannelRequestType `json:"type,omitempty"`
	Name string `json:"name"`
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
	AvailableTags []CreateGuildChannelRequestAvailableTagsInner `json:"available_tags,omitempty"`
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

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetType() CreateGuildChannelRequestType {
	if o == nil || IsNil(o.Type) {
		var ret CreateGuildChannelRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetTypeOk() (*CreateGuildChannelRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CreateGuildChannelRequestType and assigns it to the Type field.
func (o *CreateGuildChannelRequest) SetType(v CreateGuildChannelRequestType) {
	o.Type = &v
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

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *CreateGuildChannelRequest) SetPosition(v int32) {
	o.Position = &v
}


// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *CreateGuildChannelRequest) SetTopic(v string) {
	o.Topic = &v
}


// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate) {
		var ret int32
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate) {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasBitrate() bool {
	if o != nil && !IsNil(o.Bitrate) {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.
func (o *CreateGuildChannelRequest) SetBitrate(v int32) {
	o.Bitrate = &v
}


// GetUserLimit returns the UserLimit field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit) {
		var ret int32
		return ret
	}
	return *o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit) {
		return nil, false
	}
	return o.UserLimit, true
}

// HasUserLimit returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasUserLimit() bool {
	if o != nil && !IsNil(o.UserLimit) {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.
func (o *CreateGuildChannelRequest) SetUserLimit(v int32) {
	o.UserLimit = &v
}


// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *CreateGuildChannelRequest) SetNsfw(v bool) {
	o.Nsfw = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser) {
		return nil, false
	}
	return o.RateLimitPerUser, true
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasRateLimitPerUser() bool {
	if o != nil && !IsNil(o.RateLimitPerUser) {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given int32 and assigns it to the RateLimitPerUser field.
func (o *CreateGuildChannelRequest) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser = &v
}


// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ParentId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetParentId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ParentId field.
func (o *CreateGuildChannelRequest) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ParentId = &v
}


// GetPermissionOverwrites returns the PermissionOverwrites field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest {
	if o == nil || IsNil(o.PermissionOverwrites) {
		var ret []ChannelPermissionOverwriteRequest
		return ret
	}
	return o.PermissionOverwrites
}

// GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetPermissionOverwritesOk() ([]ChannelPermissionOverwriteRequest, bool) {
	if o == nil || IsNil(o.PermissionOverwrites) {
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


// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion) {
		var ret string
		return ret
	}
	return *o.RtcRegion
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion) {
		return nil, false
	}
	return o.RtcRegion, true
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasRtcRegion() bool {
	if o != nil && !IsNil(o.RtcRegion) {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given string and assigns it to the RtcRegion field.
func (o *CreateGuildChannelRequest) SetRtcRegion(v string) {
	o.RtcRegion = &v
}


// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetVideoQualityMode() CreateGuildChannelRequestVideoQualityMode {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret CreateGuildChannelRequestVideoQualityMode
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetVideoQualityModeOk() (*CreateGuildChannelRequestVideoQualityMode, bool) {
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

// SetVideoQualityMode gets a reference to the given CreateGuildChannelRequestVideoQualityMode and assigns it to the VideoQualityMode field.
func (o *CreateGuildChannelRequest) SetVideoQualityMode(v CreateGuildChannelRequestVideoQualityMode) {
	o.VideoQualityMode = &v
}


// GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		var ret CreateForumThreadRequestAutoArchiveDuration
		return ret
	}
	return *o.DefaultAutoArchiveDuration
}

// GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool) {
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

// SetDefaultAutoArchiveDuration gets a reference to the given CreateForumThreadRequestAutoArchiveDuration and assigns it to the DefaultAutoArchiveDuration field.
func (o *CreateGuildChannelRequest) SetDefaultAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration) {
	o.DefaultAutoArchiveDuration = &v
}


// GetDefaultReactionEmoji returns the DefaultReactionEmoji field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultReactionEmoji() CreateGuildChannelRequestDefaultReactionEmoji {
	if o == nil || IsNil(o.DefaultReactionEmoji) {
		var ret CreateGuildChannelRequestDefaultReactionEmoji
		return ret
	}
	return *o.DefaultReactionEmoji
}

// GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultReactionEmojiOk() (*CreateGuildChannelRequestDefaultReactionEmoji, bool) {
	if o == nil || IsNil(o.DefaultReactionEmoji) {
		return nil, false
	}
	return o.DefaultReactionEmoji, true
}

// HasDefaultReactionEmoji returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultReactionEmoji() bool {
	if o != nil && !IsNil(o.DefaultReactionEmoji) {
		return true
	}

	return false
}

// SetDefaultReactionEmoji gets a reference to the given CreateGuildChannelRequestDefaultReactionEmoji and assigns it to the DefaultReactionEmoji field.
func (o *CreateGuildChannelRequest) SetDefaultReactionEmoji(v CreateGuildChannelRequestDefaultReactionEmoji) {
	o.DefaultReactionEmoji = &v
}


// GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultThreadRateLimitPerUser() int32 {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.DefaultThreadRateLimitPerUser
}

// GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultThreadRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser) {
		return nil, false
	}
	return o.DefaultThreadRateLimitPerUser, true
}

// HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateGuildChannelRequest) HasDefaultThreadRateLimitPerUser() bool {
	if o != nil && !IsNil(o.DefaultThreadRateLimitPerUser) {
		return true
	}

	return false
}

// SetDefaultThreadRateLimitPerUser gets a reference to the given int32 and assigns it to the DefaultThreadRateLimitPerUser field.
func (o *CreateGuildChannelRequest) SetDefaultThreadRateLimitPerUser(v int32) {
	o.DefaultThreadRateLimitPerUser = &v
}


// GetDefaultSortOrder returns the DefaultSortOrder field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder {
	if o == nil || IsNil(o.DefaultSortOrder) {
		var ret CreateGuildChannelRequestDefaultSortOrder
		return ret
	}
	return *o.DefaultSortOrder
}

// GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool) {
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

// SetDefaultSortOrder gets a reference to the given CreateGuildChannelRequestDefaultSortOrder and assigns it to the DefaultSortOrder field.
func (o *CreateGuildChannelRequest) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder) {
	o.DefaultSortOrder = &v
}


// GetDefaultForumLayout returns the DefaultForumLayout field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout {
	if o == nil || IsNil(o.DefaultForumLayout) {
		var ret CreateGuildChannelRequestDefaultForumLayout
		return ret
	}
	return *o.DefaultForumLayout
}

// GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool) {
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

// SetDefaultForumLayout gets a reference to the given CreateGuildChannelRequestDefaultForumLayout and assigns it to the DefaultForumLayout field.
func (o *CreateGuildChannelRequest) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout) {
	o.DefaultForumLayout = &v
}


// GetDefaultTagSetting returns the DefaultTagSetting field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting {
	if o == nil || IsNil(o.DefaultTagSetting) {
		var ret CreateGuildChannelRequestDefaultTagSetting
		return ret
	}
	return *o.DefaultTagSetting
}

// GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool) {
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

// SetDefaultTagSetting gets a reference to the given CreateGuildChannelRequestDefaultTagSetting and assigns it to the DefaultTagSetting field.
func (o *CreateGuildChannelRequest) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting) {
	o.DefaultTagSetting = &v
}


// GetAvailableTags returns the AvailableTags field value if set, zero value otherwise.
func (o *CreateGuildChannelRequest) GetAvailableTags() []CreateGuildChannelRequestAvailableTagsInner {
	if o == nil || IsNil(o.AvailableTags) {
		var ret []CreateGuildChannelRequestAvailableTagsInner
		return ret
	}
	return o.AvailableTags
}

// GetAvailableTagsOk returns a tuple with the AvailableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGuildChannelRequest) GetAvailableTagsOk() ([]CreateGuildChannelRequestAvailableTagsInner, bool) {
	if o == nil || IsNil(o.AvailableTags) {
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

// SetAvailableTags gets a reference to the given []CreateGuildChannelRequestAvailableTagsInner and assigns it to the AvailableTags field.
func (o *CreateGuildChannelRequest) SetAvailableTags(v []CreateGuildChannelRequestAvailableTagsInner) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
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
	if !IsNil(o.AvailableTags) {
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

