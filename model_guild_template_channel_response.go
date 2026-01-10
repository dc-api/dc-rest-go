/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
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

// checks if the GuildTemplateChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateChannelResponse{}

// GuildTemplateChannelResponse struct for GuildTemplateChannelResponse
type GuildTemplateChannelResponse struct {
	Id *int32 `json:"id,omitempty"`
	Type ChannelTypes `json:"type"`
	Name *string `json:"name,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Topic *string `json:"topic,omitempty"`
	Bitrate int32 `json:"bitrate"`
	UserLimit int32 `json:"user_limit"`
	Nsfw bool `json:"nsfw"`
	RateLimitPerUser int32 `json:"rate_limit_per_user"`
	ParentId *GetEntitlementsSkuIdsParameterOneOfInner `json:"parent_id,omitempty"`
	DefaultAutoArchiveDuration *CreateForumThreadRequestAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	PermissionOverwrites []GuildTemplateChannelResponsePermissionOverwritesInner `json:"permission_overwrites"`
	AvailableTags []GuildTemplateChannelTags `json:"available_tags,omitempty"`
	Template string `json:"template"`
	DefaultReactionEmoji *GuildChannelResponseDefaultReactionEmoji `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int32 `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder *CreateGuildChannelRequestDefaultSortOrder `json:"default_sort_order,omitempty"`
	DefaultForumLayout *CreateGuildChannelRequestDefaultForumLayout `json:"default_forum_layout,omitempty"`
	DefaultTagSetting *CreateGuildChannelRequestDefaultTagSetting `json:"default_tag_setting,omitempty"`
	IconEmoji *GuildTemplateChannelResponseIconEmoji `json:"icon_emoji,omitempty"`
	ThemeColor *int32 `json:"theme_color,omitempty"`
}

type _GuildTemplateChannelResponse GuildTemplateChannelResponse

// NewGuildTemplateChannelResponse instantiates a new GuildTemplateChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateChannelResponse(type_ ChannelTypes, bitrate int32, userLimit int32, nsfw bool, rateLimitPerUser int32, permissionOverwrites []GuildTemplateChannelResponsePermissionOverwritesInner, template string) *GuildTemplateChannelResponse {
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GuildTemplateChannelResponse) SetId(v int32) {
	o.Id = &v
}


// GetType returns the Type field value
func (o *GuildTemplateChannelResponse) GetType() ChannelTypes {
	if o == nil {
		var ret ChannelTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetTypeOk() (*ChannelTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GuildTemplateChannelResponse) SetType(v ChannelTypes) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GuildTemplateChannelResponse) SetName(v string) {
	o.Name = &v
}


// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *GuildTemplateChannelResponse) SetPosition(v int32) {
	o.Position = &v
}


// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *GuildTemplateChannelResponse) SetTopic(v string) {
	o.Topic = &v
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
func (o *GuildTemplateChannelResponse) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ParentId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetParentId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ParentId field.
func (o *GuildTemplateChannelResponse) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ParentId = &v
}


// GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration {
	if o == nil || IsNil(o.DefaultAutoArchiveDuration) {
		var ret CreateForumThreadRequestAutoArchiveDuration
		return ret
	}
	return *o.DefaultAutoArchiveDuration
}

// GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool) {
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

// SetDefaultAutoArchiveDuration gets a reference to the given CreateForumThreadRequestAutoArchiveDuration and assigns it to the DefaultAutoArchiveDuration field.
func (o *GuildTemplateChannelResponse) SetDefaultAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration) {
	o.DefaultAutoArchiveDuration = &v
}


// GetPermissionOverwrites returns the PermissionOverwrites field value
func (o *GuildTemplateChannelResponse) GetPermissionOverwrites() []GuildTemplateChannelResponsePermissionOverwritesInner {
	if o == nil {
		var ret []GuildTemplateChannelResponsePermissionOverwritesInner
		return ret
	}

	return o.PermissionOverwrites
}

// GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetPermissionOverwritesOk() ([]GuildTemplateChannelResponsePermissionOverwritesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionOverwrites, true
}

// SetPermissionOverwrites sets field value
func (o *GuildTemplateChannelResponse) SetPermissionOverwrites(v []GuildTemplateChannelResponsePermissionOverwritesInner) {
	o.PermissionOverwrites = v
}

// GetAvailableTags returns the AvailableTags field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetAvailableTags() []GuildTemplateChannelTags {
	if o == nil || IsNil(o.AvailableTags) {
		var ret []GuildTemplateChannelTags
		return ret
	}
	return o.AvailableTags
}

// GetAvailableTagsOk returns a tuple with the AvailableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetAvailableTagsOk() ([]GuildTemplateChannelTags, bool) {
	if o == nil || IsNil(o.AvailableTags) {
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

// GetDefaultReactionEmoji returns the DefaultReactionEmoji field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultReactionEmoji() GuildChannelResponseDefaultReactionEmoji {
	if o == nil || IsNil(o.DefaultReactionEmoji) {
		var ret GuildChannelResponseDefaultReactionEmoji
		return ret
	}
	return *o.DefaultReactionEmoji
}

// GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultReactionEmojiOk() (*GuildChannelResponseDefaultReactionEmoji, bool) {
	if o == nil || IsNil(o.DefaultReactionEmoji) {
		return nil, false
	}
	return o.DefaultReactionEmoji, true
}

// HasDefaultReactionEmoji returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultReactionEmoji() bool {
	if o != nil && !IsNil(o.DefaultReactionEmoji) {
		return true
	}

	return false
}

// SetDefaultReactionEmoji gets a reference to the given GuildChannelResponseDefaultReactionEmoji and assigns it to the DefaultReactionEmoji field.
func (o *GuildTemplateChannelResponse) SetDefaultReactionEmoji(v GuildChannelResponseDefaultReactionEmoji) {
	o.DefaultReactionEmoji = &v
}


// GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultThreadRateLimitPerUser() int32 {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.DefaultThreadRateLimitPerUser
}

// GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultThreadRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultThreadRateLimitPerUser) {
		return nil, false
	}
	return o.DefaultThreadRateLimitPerUser, true
}

// HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasDefaultThreadRateLimitPerUser() bool {
	if o != nil && !IsNil(o.DefaultThreadRateLimitPerUser) {
		return true
	}

	return false
}

// SetDefaultThreadRateLimitPerUser gets a reference to the given int32 and assigns it to the DefaultThreadRateLimitPerUser field.
func (o *GuildTemplateChannelResponse) SetDefaultThreadRateLimitPerUser(v int32) {
	o.DefaultThreadRateLimitPerUser = &v
}


// GetDefaultSortOrder returns the DefaultSortOrder field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultSortOrder() CreateGuildChannelRequestDefaultSortOrder {
	if o == nil || IsNil(o.DefaultSortOrder) {
		var ret CreateGuildChannelRequestDefaultSortOrder
		return ret
	}
	return *o.DefaultSortOrder
}

// GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultSortOrderOk() (*CreateGuildChannelRequestDefaultSortOrder, bool) {
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

// SetDefaultSortOrder gets a reference to the given CreateGuildChannelRequestDefaultSortOrder and assigns it to the DefaultSortOrder field.
func (o *GuildTemplateChannelResponse) SetDefaultSortOrder(v CreateGuildChannelRequestDefaultSortOrder) {
	o.DefaultSortOrder = &v
}


// GetDefaultForumLayout returns the DefaultForumLayout field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultForumLayout() CreateGuildChannelRequestDefaultForumLayout {
	if o == nil || IsNil(o.DefaultForumLayout) {
		var ret CreateGuildChannelRequestDefaultForumLayout
		return ret
	}
	return *o.DefaultForumLayout
}

// GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultForumLayoutOk() (*CreateGuildChannelRequestDefaultForumLayout, bool) {
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

// SetDefaultForumLayout gets a reference to the given CreateGuildChannelRequestDefaultForumLayout and assigns it to the DefaultForumLayout field.
func (o *GuildTemplateChannelResponse) SetDefaultForumLayout(v CreateGuildChannelRequestDefaultForumLayout) {
	o.DefaultForumLayout = &v
}


// GetDefaultTagSetting returns the DefaultTagSetting field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetDefaultTagSetting() CreateGuildChannelRequestDefaultTagSetting {
	if o == nil || IsNil(o.DefaultTagSetting) {
		var ret CreateGuildChannelRequestDefaultTagSetting
		return ret
	}
	return *o.DefaultTagSetting
}

// GetDefaultTagSettingOk returns a tuple with the DefaultTagSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetDefaultTagSettingOk() (*CreateGuildChannelRequestDefaultTagSetting, bool) {
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

// SetDefaultTagSetting gets a reference to the given CreateGuildChannelRequestDefaultTagSetting and assigns it to the DefaultTagSetting field.
func (o *GuildTemplateChannelResponse) SetDefaultTagSetting(v CreateGuildChannelRequestDefaultTagSetting) {
	o.DefaultTagSetting = &v
}


// GetIconEmoji returns the IconEmoji field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetIconEmoji() GuildTemplateChannelResponseIconEmoji {
	if o == nil || IsNil(o.IconEmoji) {
		var ret GuildTemplateChannelResponseIconEmoji
		return ret
	}
	return *o.IconEmoji
}

// GetIconEmojiOk returns a tuple with the IconEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetIconEmojiOk() (*GuildTemplateChannelResponseIconEmoji, bool) {
	if o == nil || IsNil(o.IconEmoji) {
		return nil, false
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

// SetIconEmoji gets a reference to the given GuildTemplateChannelResponseIconEmoji and assigns it to the IconEmoji field.
func (o *GuildTemplateChannelResponse) SetIconEmoji(v GuildTemplateChannelResponseIconEmoji) {
	o.IconEmoji = &v
}


// GetThemeColor returns the ThemeColor field value if set, zero value otherwise.
func (o *GuildTemplateChannelResponse) GetThemeColor() int32 {
	if o == nil || IsNil(o.ThemeColor) {
		var ret int32
		return ret
	}
	return *o.ThemeColor
}

// GetThemeColorOk returns a tuple with the ThemeColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateChannelResponse) GetThemeColorOk() (*int32, bool) {
	if o == nil || IsNil(o.ThemeColor) {
		return nil, false
	}
	return o.ThemeColor, true
}

// HasThemeColor returns a boolean if a field has been set.
func (o *GuildTemplateChannelResponse) HasThemeColor() bool {
	if o != nil && !IsNil(o.ThemeColor) {
		return true
	}

	return false
}

// SetThemeColor gets a reference to the given int32 and assigns it to the ThemeColor field.
func (o *GuildTemplateChannelResponse) SetThemeColor(v int32) {
	o.ThemeColor = &v
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
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
	if !IsNil(o.AvailableTags) {
		toSerialize["available_tags"] = o.AvailableTags
	}
	toSerialize["template"] = o.Template
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
	if !IsNil(o.IconEmoji) {
		toSerialize["icon_emoji"] = o.IconEmoji
	}
	if !IsNil(o.ThemeColor) {
		toSerialize["theme_color"] = o.ThemeColor
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

