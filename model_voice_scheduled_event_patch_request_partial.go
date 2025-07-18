/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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
)

// checks if the VoiceScheduledEventPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceScheduledEventPatchRequestPartial{}

// VoiceScheduledEventPatchRequestPartial struct for VoiceScheduledEventPatchRequestPartial
type VoiceScheduledEventPatchRequestPartial struct {
	Status NullableInt32 `json:"status,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime NullableTime `json:"scheduled_end_time,omitempty"`
	EntityType NullableInt32 `json:"entity_type,omitempty"`
	PrivacyLevel interface{} `json:"privacy_level,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EntityMetadata map[string]interface{} `json:"entity_metadata,omitempty"`
}

// NewVoiceScheduledEventPatchRequestPartial instantiates a new VoiceScheduledEventPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceScheduledEventPatchRequestPartial() *VoiceScheduledEventPatchRequestPartial {
	this := VoiceScheduledEventPatchRequestPartial{}
	return &this
}

// NewVoiceScheduledEventPatchRequestPartialWithDefaults instantiates a new VoiceScheduledEventPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceScheduledEventPatchRequestPartialWithDefaults() *VoiceScheduledEventPatchRequestPartial {
	this := VoiceScheduledEventPatchRequestPartial{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventPatchRequestPartial) GetStatus() int32 {
	if o == nil || IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventPatchRequestPartial) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status.Get()) {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *VoiceScheduledEventPatchRequestPartial) SetStatus(v int32) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) UnsetStatus() {
	o.Status.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VoiceScheduledEventPatchRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventPatchRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VoiceScheduledEventPatchRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VoiceScheduledEventPatchRequestPartial) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) UnsetDescription() {
	o.Description.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventPatchRequestPartial) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventPatchRequestPartial) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image.Get()) {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *VoiceScheduledEventPatchRequestPartial) SetImage(v string) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) UnsetImage() {
	o.Image.Unset()
}

// GetScheduledStartTime returns the ScheduledStartTime field value if set, zero value otherwise.
func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledStartTime() time.Time {
	if o == nil || IsNil(o.ScheduledStartTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledStartTime) {
		return nil, false
	}
	return o.ScheduledStartTime, true
}

// HasScheduledStartTime returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasScheduledStartTime() bool {
	if o != nil && !IsNil(o.ScheduledStartTime) {
		return true
	}

	return false
}

// SetScheduledStartTime gets a reference to the given time.Time and assigns it to the ScheduledStartTime field.
func (o *VoiceScheduledEventPatchRequestPartial) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = &v
}


// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime.Get()
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventPatchRequestPartial) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		return nil, false
	}
	return o.ScheduledEndTime.Get(), o.ScheduledEndTime.IsSet()
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasScheduledEndTime() bool {
	if o != nil && o.ScheduledEndTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given NullableTime and assigns it to the ScheduledEndTime field.
func (o *VoiceScheduledEventPatchRequestPartial) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime.Set(&v)
}

// SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) SetScheduledEndTimeNil() {
	o.ScheduledEndTime.Set(nil)
}

// UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) UnsetScheduledEndTime() {
	o.ScheduledEndTime.Unset()
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventPatchRequestPartial) GetEntityType() int32 {
	if o == nil || IsNil(o.EntityType.Get()) {
		var ret int32
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventPatchRequestPartial) GetEntityTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityType.Get()) {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasEntityType() bool {
	if o != nil && o.EntityType.IsSet() {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableInt32 and assigns it to the EntityType field.
func (o *VoiceScheduledEventPatchRequestPartial) SetEntityType(v int32) {
	o.EntityType.Set(&v)
}

// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *VoiceScheduledEventPatchRequestPartial) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetPrivacyLevel returns the PrivacyLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventPatchRequestPartial) GetPrivacyLevel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventPatchRequestPartial) GetPrivacyLevelOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// HasPrivacyLevel returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasPrivacyLevel() bool {
	if o != nil && !IsNil(o.PrivacyLevel) {
		return true
	}

	return false
}

// SetPrivacyLevel gets a reference to the given interface{} and assigns it to the PrivacyLevel field.
func (o *VoiceScheduledEventPatchRequestPartial) SetPrivacyLevel(v interface{}) {
	o.PrivacyLevel = v
}


// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *VoiceScheduledEventPatchRequestPartial) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventPatchRequestPartial) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *VoiceScheduledEventPatchRequestPartial) SetChannelId(v string) {
	o.ChannelId = &v
}


// GetEntityMetadata returns the EntityMetadata field value if set, zero value otherwise.
func (o *VoiceScheduledEventPatchRequestPartial) GetEntityMetadata() map[string]interface{} {
	if o == nil || IsNil(o.EntityMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventPatchRequestPartial) GetEntityMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EntityMetadata) {
		return map[string]interface{}{}, false
	}
	return o.EntityMetadata, true
}

// HasEntityMetadata returns a boolean if a field has been set.
func (o *VoiceScheduledEventPatchRequestPartial) HasEntityMetadata() bool {
	if o != nil && !IsNil(o.EntityMetadata) {
		return true
	}

	return false
}

// SetEntityMetadata gets a reference to the given map[string]interface{} and assigns it to the EntityMetadata field.
func (o *VoiceScheduledEventPatchRequestPartial) SetEntityMetadata(v map[string]interface{}) {
	o.EntityMetadata = v
}


func (o VoiceScheduledEventPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceScheduledEventPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if !IsNil(o.ScheduledStartTime) {
		toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	}
	if o.ScheduledEndTime.IsSet() {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime.Get()
	}
	if o.EntityType.IsSet() {
		toSerialize["entity_type"] = o.EntityType.Get()
	}
	if o.PrivacyLevel != nil {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.EntityMetadata) {
		toSerialize["entity_metadata"] = o.EntityMetadata
	}
	return toSerialize, nil
}

type NullableVoiceScheduledEventPatchRequestPartial struct {
	value *VoiceScheduledEventPatchRequestPartial
	isSet bool
}

func (v NullableVoiceScheduledEventPatchRequestPartial) Get() *VoiceScheduledEventPatchRequestPartial {
	return v.value
}

func (v *NullableVoiceScheduledEventPatchRequestPartial) Set(val *VoiceScheduledEventPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceScheduledEventPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceScheduledEventPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceScheduledEventPatchRequestPartial(val *VoiceScheduledEventPatchRequestPartial) *NullableVoiceScheduledEventPatchRequestPartial {
	return &NullableVoiceScheduledEventPatchRequestPartial{value: val, isSet: true}
}

func (v NullableVoiceScheduledEventPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceScheduledEventPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


