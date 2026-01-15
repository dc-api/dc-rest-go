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

// checks if the UpdateThreadRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateThreadRequestPartial{}

// UpdateThreadRequestPartial struct for UpdateThreadRequestPartial
type UpdateThreadRequestPartial struct {
	Name *string `json:"name,omitempty"`
	Archived *bool `json:"archived,omitempty"`
	Locked *bool `json:"locked,omitempty"`
	Invitable *bool `json:"invitable,omitempty"`
	AutoArchiveDuration *CreateForumThreadRequestAutoArchiveDuration `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser *int32 `json:"rate_limit_per_user,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
	AppliedTags []string `json:"applied_tags,omitempty"`
	Bitrate *int32 `json:"bitrate,omitempty"`
	UserLimit *int32 `json:"user_limit,omitempty"`
	RtcRegion *string `json:"rtc_region,omitempty"`
	VideoQualityMode *CreateGuildChannelRequestVideoQualityMode `json:"video_quality_mode,omitempty"`
}

// NewUpdateThreadRequestPartial instantiates a new UpdateThreadRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateThreadRequestPartial() *UpdateThreadRequestPartial {
	this := UpdateThreadRequestPartial{}
	return &this
}

// NewUpdateThreadRequestPartialWithDefaults instantiates a new UpdateThreadRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateThreadRequestPartialWithDefaults() *UpdateThreadRequestPartial {
	this := UpdateThreadRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateThreadRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateThreadRequestPartial) SetArchived(v bool) {
	o.Archived = &v
}


// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *UpdateThreadRequestPartial) SetLocked(v bool) {
	o.Locked = &v
}


// GetInvitable returns the Invitable field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetInvitable() bool {
	if o == nil || IsNil(o.Invitable) {
		var ret bool
		return ret
	}
	return *o.Invitable
}

// GetInvitableOk returns a tuple with the Invitable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetInvitableOk() (*bool, bool) {
	if o == nil || IsNil(o.Invitable) {
		return nil, false
	}
	return o.Invitable, true
}

// HasInvitable returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasInvitable() bool {
	if o != nil && !IsNil(o.Invitable) {
		return true
	}

	return false
}

// SetInvitable gets a reference to the given bool and assigns it to the Invitable field.
func (o *UpdateThreadRequestPartial) SetInvitable(v bool) {
	o.Invitable = &v
}


// GetAutoArchiveDuration returns the AutoArchiveDuration field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetAutoArchiveDuration() CreateForumThreadRequestAutoArchiveDuration {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		var ret CreateForumThreadRequestAutoArchiveDuration
		return ret
	}
	return *o.AutoArchiveDuration
}

// GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetAutoArchiveDurationOk() (*CreateForumThreadRequestAutoArchiveDuration, bool) {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		return nil, false
	}
	return o.AutoArchiveDuration, true
}

// HasAutoArchiveDuration returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.AutoArchiveDuration) {
		return true
	}

	return false
}

// SetAutoArchiveDuration gets a reference to the given CreateForumThreadRequestAutoArchiveDuration and assigns it to the AutoArchiveDuration field.
func (o *UpdateThreadRequestPartial) SetAutoArchiveDuration(v CreateForumThreadRequestAutoArchiveDuration) {
	o.AutoArchiveDuration = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser) {
		return nil, false
	}
	return o.RateLimitPerUser, true
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasRateLimitPerUser() bool {
	if o != nil && !IsNil(o.RateLimitPerUser) {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given int32 and assigns it to the RateLimitPerUser field.
func (o *UpdateThreadRequestPartial) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser = &v
}


// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *UpdateThreadRequestPartial) SetFlags(v int32) {
	o.Flags = &v
}


// GetAppliedTags returns the AppliedTags field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetAppliedTags() []string {
	if o == nil || IsNil(o.AppliedTags) {
		var ret []string
		return ret
	}
	return o.AppliedTags
}

// GetAppliedTagsOk returns a tuple with the AppliedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetAppliedTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppliedTags) {
		return nil, false
	}
	return o.AppliedTags, true
}

// HasAppliedTags returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasAppliedTags() bool {
	if o != nil && !IsNil(o.AppliedTags) {
		return true
	}

	return false
}

// SetAppliedTags gets a reference to the given []string and assigns it to the AppliedTags field.
func (o *UpdateThreadRequestPartial) SetAppliedTags(v []string) {
	o.AppliedTags = v
}


// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate) {
		var ret int32
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate) {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasBitrate() bool {
	if o != nil && !IsNil(o.Bitrate) {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.
func (o *UpdateThreadRequestPartial) SetBitrate(v int32) {
	o.Bitrate = &v
}


// GetUserLimit returns the UserLimit field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit) {
		var ret int32
		return ret
	}
	return *o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit) {
		return nil, false
	}
	return o.UserLimit, true
}

// HasUserLimit returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasUserLimit() bool {
	if o != nil && !IsNil(o.UserLimit) {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.
func (o *UpdateThreadRequestPartial) SetUserLimit(v int32) {
	o.UserLimit = &v
}


// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion) {
		var ret string
		return ret
	}
	return *o.RtcRegion
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion) {
		return nil, false
	}
	return o.RtcRegion, true
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasRtcRegion() bool {
	if o != nil && !IsNil(o.RtcRegion) {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given string and assigns it to the RtcRegion field.
func (o *UpdateThreadRequestPartial) SetRtcRegion(v string) {
	o.RtcRegion = &v
}


// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetVideoQualityMode() CreateGuildChannelRequestVideoQualityMode {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret CreateGuildChannelRequestVideoQualityMode
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetVideoQualityModeOk() (*CreateGuildChannelRequestVideoQualityMode, bool) {
	if o == nil || IsNil(o.VideoQualityMode) {
		return nil, false
	}
	return o.VideoQualityMode, true
}

// HasVideoQualityMode returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasVideoQualityMode() bool {
	if o != nil && !IsNil(o.VideoQualityMode) {
		return true
	}

	return false
}

// SetVideoQualityMode gets a reference to the given CreateGuildChannelRequestVideoQualityMode and assigns it to the VideoQualityMode field.
func (o *UpdateThreadRequestPartial) SetVideoQualityMode(v CreateGuildChannelRequestVideoQualityMode) {
	o.VideoQualityMode = &v
}


func (o UpdateThreadRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateThreadRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Invitable) {
		toSerialize["invitable"] = o.Invitable
	}
	if !IsNil(o.AutoArchiveDuration) {
		toSerialize["auto_archive_duration"] = o.AutoArchiveDuration
	}
	if !IsNil(o.RateLimitPerUser) {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.AppliedTags) {
		toSerialize["applied_tags"] = o.AppliedTags
	}
	if !IsNil(o.Bitrate) {
		toSerialize["bitrate"] = o.Bitrate
	}
	if !IsNil(o.UserLimit) {
		toSerialize["user_limit"] = o.UserLimit
	}
	if !IsNil(o.RtcRegion) {
		toSerialize["rtc_region"] = o.RtcRegion
	}
	if !IsNil(o.VideoQualityMode) {
		toSerialize["video_quality_mode"] = o.VideoQualityMode
	}
	return toSerialize, nil
}

type NullableUpdateThreadRequestPartial struct {
	value *UpdateThreadRequestPartial
	isSet bool
}

func (v NullableUpdateThreadRequestPartial) Get() *UpdateThreadRequestPartial {
	return v.value
}

func (v *NullableUpdateThreadRequestPartial) Set(val *UpdateThreadRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateThreadRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateThreadRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateThreadRequestPartial(val *UpdateThreadRequestPartial) *NullableUpdateThreadRequestPartial {
	return &NullableUpdateThreadRequestPartial{value: val, isSet: true}
}

func (v NullableUpdateThreadRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateThreadRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

