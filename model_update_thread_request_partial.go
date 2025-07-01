/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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
)

// checks if the UpdateThreadRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateThreadRequestPartial{}

// UpdateThreadRequestPartial struct for UpdateThreadRequestPartial
type UpdateThreadRequestPartial struct {
	Name NullableString `json:"name,omitempty"`
	Archived NullableBool `json:"archived,omitempty"`
	Locked NullableBool `json:"locked,omitempty"`
	Invitable NullableBool `json:"invitable,omitempty"`
	AutoArchiveDuration *int32 `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser NullableInt32 `json:"rate_limit_per_user,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
	AppliedTags []string `json:"applied_tags,omitempty"`
	Bitrate NullableInt32 `json:"bitrate,omitempty"`
	UserLimit NullableInt32 `json:"user_limit,omitempty"`
	RtcRegion NullableString `json:"rtc_region,omitempty"`
	VideoQualityMode *int32 `json:"video_quality_mode,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateThreadRequestPartial) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateThreadRequestPartial) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetName() {
	o.Name.Unset()
}

// GetArchived returns the Archived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetArchived() bool {
	if o == nil || IsNil(o.Archived.Get()) {
		var ret bool
		return ret
	}
	return *o.Archived.Get()
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived.Get()) {
		return nil, false
	}
	return o.Archived.Get(), o.Archived.IsSet()
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasArchived() bool {
	if o != nil && o.Archived.IsSet() {
		return true
	}

	return false
}

// SetArchived gets a reference to the given NullableBool and assigns it to the Archived field.
func (o *UpdateThreadRequestPartial) SetArchived(v bool) {
	o.Archived.Set(&v)
}

// SetArchivedNil sets the value for Archived to be an explicit nil
func (o *UpdateThreadRequestPartial) SetArchivedNil() {
	o.Archived.Set(nil)
}

// UnsetArchived ensures that no value is present for Archived, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetArchived() {
	o.Archived.Unset()
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetLocked() bool {
	if o == nil || IsNil(o.Locked.Get()) {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked.Get()) {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasLocked() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *UpdateThreadRequestPartial) SetLocked(v bool) {
	o.Locked.Set(&v)
}

// SetLockedNil sets the value for Locked to be an explicit nil
func (o *UpdateThreadRequestPartial) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetLocked() {
	o.Locked.Unset()
}

// GetInvitable returns the Invitable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetInvitable() bool {
	if o == nil || IsNil(o.Invitable.Get()) {
		var ret bool
		return ret
	}
	return *o.Invitable.Get()
}

// GetInvitableOk returns a tuple with the Invitable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetInvitableOk() (*bool, bool) {
	if o == nil || IsNil(o.Invitable.Get()) {
		return nil, false
	}
	return o.Invitable.Get(), o.Invitable.IsSet()
}

// HasInvitable returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasInvitable() bool {
	if o != nil && o.Invitable.IsSet() {
		return true
	}

	return false
}

// SetInvitable gets a reference to the given NullableBool and assigns it to the Invitable field.
func (o *UpdateThreadRequestPartial) SetInvitable(v bool) {
	o.Invitable.Set(&v)
}

// SetInvitableNil sets the value for Invitable to be an explicit nil
func (o *UpdateThreadRequestPartial) SetInvitableNil() {
	o.Invitable.Set(nil)
}

// UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetInvitable() {
	o.Invitable.Unset()
}

// GetAutoArchiveDuration returns the AutoArchiveDuration field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetAutoArchiveDuration() int32 {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		var ret int32
		return ret
	}
	return *o.AutoArchiveDuration
}

// GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetAutoArchiveDurationOk() (*int32, bool) {
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

// SetAutoArchiveDuration gets a reference to the given int32 and assigns it to the AutoArchiveDuration field.
func (o *UpdateThreadRequestPartial) SetAutoArchiveDuration(v int32) {
	o.AutoArchiveDuration = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser.Get()
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		return nil, false
	}
	return o.RateLimitPerUser.Get(), o.RateLimitPerUser.IsSet()
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasRateLimitPerUser() bool {
	if o != nil && o.RateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the RateLimitPerUser field.
func (o *UpdateThreadRequestPartial) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser.Set(&v)
}

// SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil
func (o *UpdateThreadRequestPartial) SetRateLimitPerUserNil() {
	o.RateLimitPerUser.Set(nil)
}

// UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetRateLimitPerUser() {
	o.RateLimitPerUser.Unset()
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *UpdateThreadRequestPartial) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *UpdateThreadRequestPartial) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetFlags() {
	o.Flags.Unset()
}

// GetAppliedTags returns the AppliedTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetAppliedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AppliedTags
}

// GetAppliedTagsOk returns a tuple with the AppliedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetAppliedTagsOk() ([]string, bool) {
	if o == nil {
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


// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.Bitrate.Get()) {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt32 and assigns it to the Bitrate field.
func (o *UpdateThreadRequestPartial) SetBitrate(v int32) {
	o.Bitrate.Set(&v)
}

// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *UpdateThreadRequestPartial) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetUserLimit returns the UserLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.UserLimit.Get()
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit.Get()) {
		return nil, false
	}
	return o.UserLimit.Get(), o.UserLimit.IsSet()
}

// HasUserLimit returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasUserLimit() bool {
	if o != nil && o.UserLimit.IsSet() {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given NullableInt32 and assigns it to the UserLimit field.
func (o *UpdateThreadRequestPartial) SetUserLimit(v int32) {
	o.UserLimit.Set(&v)
}

// SetUserLimitNil sets the value for UserLimit to be an explicit nil
func (o *UpdateThreadRequestPartial) SetUserLimitNil() {
	o.UserLimit.Set(nil)
}

// UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetUserLimit() {
	o.UserLimit.Unset()
}

// GetRtcRegion returns the RtcRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateThreadRequestPartial) GetRtcRegion() string {
	if o == nil || IsNil(o.RtcRegion.Get()) {
		var ret string
		return ret
	}
	return *o.RtcRegion.Get()
}

// GetRtcRegionOk returns a tuple with the RtcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateThreadRequestPartial) GetRtcRegionOk() (*string, bool) {
	if o == nil || IsNil(o.RtcRegion.Get()) {
		return nil, false
	}
	return o.RtcRegion.Get(), o.RtcRegion.IsSet()
}

// HasRtcRegion returns a boolean if a field has been set.
func (o *UpdateThreadRequestPartial) HasRtcRegion() bool {
	if o != nil && o.RtcRegion.IsSet() {
		return true
	}

	return false
}

// SetRtcRegion gets a reference to the given NullableString and assigns it to the RtcRegion field.
func (o *UpdateThreadRequestPartial) SetRtcRegion(v string) {
	o.RtcRegion.Set(&v)
}

// SetRtcRegionNil sets the value for RtcRegion to be an explicit nil
func (o *UpdateThreadRequestPartial) SetRtcRegionNil() {
	o.RtcRegion.Set(nil)
}

// UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
func (o *UpdateThreadRequestPartial) UnsetRtcRegion() {
	o.RtcRegion.Unset()
}

// GetVideoQualityMode returns the VideoQualityMode field value if set, zero value otherwise.
func (o *UpdateThreadRequestPartial) GetVideoQualityMode() int32 {
	if o == nil || IsNil(o.VideoQualityMode) {
		var ret int32
		return ret
	}
	return *o.VideoQualityMode
}

// GetVideoQualityModeOk returns a tuple with the VideoQualityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThreadRequestPartial) GetVideoQualityModeOk() (*int32, bool) {
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

// SetVideoQualityMode gets a reference to the given int32 and assigns it to the VideoQualityMode field.
func (o *UpdateThreadRequestPartial) SetVideoQualityMode(v int32) {
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Archived.IsSet() {
		toSerialize["archived"] = o.Archived.Get()
	}
	if o.Locked.IsSet() {
		toSerialize["locked"] = o.Locked.Get()
	}
	if o.Invitable.IsSet() {
		toSerialize["invitable"] = o.Invitable.Get()
	}
	if !IsNil(o.AutoArchiveDuration) {
		toSerialize["auto_archive_duration"] = o.AutoArchiveDuration
	}
	if o.RateLimitPerUser.IsSet() {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser.Get()
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	if o.AppliedTags != nil {
		toSerialize["applied_tags"] = o.AppliedTags
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


