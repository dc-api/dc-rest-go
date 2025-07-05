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
	"bytes"
	"fmt"
)

// checks if the VoiceScheduledEventCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceScheduledEventCreateRequest{}

// VoiceScheduledEventCreateRequest struct for VoiceScheduledEventCreateRequest
type VoiceScheduledEventCreateRequest struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ScheduledStartTime time.Time `json:"scheduled_start_time"`
	ScheduledEndTime NullableTime `json:"scheduled_end_time,omitempty"`
	PrivacyLevel interface{} `json:"privacy_level"`
	EntityType NullableInt32 `json:"entity_type"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EntityMetadata map[string]interface{} `json:"entity_metadata,omitempty"`
}

type _VoiceScheduledEventCreateRequest VoiceScheduledEventCreateRequest

// NewVoiceScheduledEventCreateRequest instantiates a new VoiceScheduledEventCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceScheduledEventCreateRequest(name string, scheduledStartTime time.Time, privacyLevel interface{}, entityType NullableInt32) *VoiceScheduledEventCreateRequest {
	this := VoiceScheduledEventCreateRequest{}
	this.Name = name
	this.ScheduledStartTime = scheduledStartTime
	this.PrivacyLevel = privacyLevel
	this.EntityType = entityType
	return &this
}

// NewVoiceScheduledEventCreateRequestWithDefaults instantiates a new VoiceScheduledEventCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceScheduledEventCreateRequestWithDefaults() *VoiceScheduledEventCreateRequest {
	this := VoiceScheduledEventCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *VoiceScheduledEventCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VoiceScheduledEventCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VoiceScheduledEventCreateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VoiceScheduledEventCreateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VoiceScheduledEventCreateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VoiceScheduledEventCreateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventCreateRequest) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventCreateRequest) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image.Get()) {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *VoiceScheduledEventCreateRequest) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *VoiceScheduledEventCreateRequest) SetImage(v string) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *VoiceScheduledEventCreateRequest) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *VoiceScheduledEventCreateRequest) UnsetImage() {
	o.Image.Unset()
}

// GetScheduledStartTime returns the ScheduledStartTime field value
func (o *VoiceScheduledEventCreateRequest) GetScheduledStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventCreateRequest) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledStartTime, true
}

// SetScheduledStartTime sets field value
func (o *VoiceScheduledEventCreateRequest) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VoiceScheduledEventCreateRequest) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime.Get()
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventCreateRequest) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		return nil, false
	}
	return o.ScheduledEndTime.Get(), o.ScheduledEndTime.IsSet()
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *VoiceScheduledEventCreateRequest) HasScheduledEndTime() bool {
	if o != nil && o.ScheduledEndTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given NullableTime and assigns it to the ScheduledEndTime field.
func (o *VoiceScheduledEventCreateRequest) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime.Set(&v)
}

// SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil
func (o *VoiceScheduledEventCreateRequest) SetScheduledEndTimeNil() {
	o.ScheduledEndTime.Set(nil)
}

// UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
func (o *VoiceScheduledEventCreateRequest) UnsetScheduledEndTime() {
	o.ScheduledEndTime.Unset()
}

// GetPrivacyLevel returns the PrivacyLevel field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *VoiceScheduledEventCreateRequest) GetPrivacyLevel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventCreateRequest) GetPrivacyLevelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *VoiceScheduledEventCreateRequest) SetPrivacyLevel(v interface{}) {
	o.PrivacyLevel = v
}

// GetEntityType returns the EntityType field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *VoiceScheduledEventCreateRequest) GetEntityType() int32 {
	if o == nil || o.EntityType.Get() == nil {
		var ret int32
		return ret
	}

	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VoiceScheduledEventCreateRequest) GetEntityTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// SetEntityType sets field value
func (o *VoiceScheduledEventCreateRequest) SetEntityType(v int32) {
	o.EntityType.Set(&v)
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *VoiceScheduledEventCreateRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventCreateRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *VoiceScheduledEventCreateRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *VoiceScheduledEventCreateRequest) SetChannelId(v string) {
	o.ChannelId = &v
}


// GetEntityMetadata returns the EntityMetadata field value if set, zero value otherwise.
func (o *VoiceScheduledEventCreateRequest) GetEntityMetadata() map[string]interface{} {
	if o == nil || IsNil(o.EntityMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceScheduledEventCreateRequest) GetEntityMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EntityMetadata) {
		return map[string]interface{}{}, false
	}
	return o.EntityMetadata, true
}

// HasEntityMetadata returns a boolean if a field has been set.
func (o *VoiceScheduledEventCreateRequest) HasEntityMetadata() bool {
	if o != nil && !IsNil(o.EntityMetadata) {
		return true
	}

	return false
}

// SetEntityMetadata gets a reference to the given map[string]interface{} and assigns it to the EntityMetadata field.
func (o *VoiceScheduledEventCreateRequest) SetEntityMetadata(v map[string]interface{}) {
	o.EntityMetadata = v
}


func (o VoiceScheduledEventCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceScheduledEventCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	if o.ScheduledEndTime.IsSet() {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime.Get()
	}
	if o.PrivacyLevel != nil {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	toSerialize["entity_type"] = o.EntityType.Get()
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.EntityMetadata) {
		toSerialize["entity_metadata"] = o.EntityMetadata
	}
	return toSerialize, nil
}

func (o *VoiceScheduledEventCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scheduled_start_time",
		"privacy_level",
		"entity_type",
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

	varVoiceScheduledEventCreateRequest := _VoiceScheduledEventCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVoiceScheduledEventCreateRequest)

	if err != nil {
		return err
	}

	*o = VoiceScheduledEventCreateRequest(varVoiceScheduledEventCreateRequest)

	return err
}

type NullableVoiceScheduledEventCreateRequest struct {
	value *VoiceScheduledEventCreateRequest
	isSet bool
}

func (v NullableVoiceScheduledEventCreateRequest) Get() *VoiceScheduledEventCreateRequest {
	return v.value
}

func (v *NullableVoiceScheduledEventCreateRequest) Set(val *VoiceScheduledEventCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceScheduledEventCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceScheduledEventCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceScheduledEventCreateRequest(val *VoiceScheduledEventCreateRequest) *NullableVoiceScheduledEventCreateRequest {
	return &NullableVoiceScheduledEventCreateRequest{value: val, isSet: true}
}

func (v NullableVoiceScheduledEventCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceScheduledEventCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


