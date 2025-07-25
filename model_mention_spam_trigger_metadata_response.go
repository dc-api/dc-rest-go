/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the MentionSpamTriggerMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionSpamTriggerMetadataResponse{}

// MentionSpamTriggerMetadataResponse struct for MentionSpamTriggerMetadataResponse
type MentionSpamTriggerMetadataResponse struct {
	MentionTotalLimit int32 `json:"mention_total_limit"`
	MentionRaidProtectionEnabled NullableBool `json:"mention_raid_protection_enabled,omitempty"`
}

type _MentionSpamTriggerMetadataResponse MentionSpamTriggerMetadataResponse

// NewMentionSpamTriggerMetadataResponse instantiates a new MentionSpamTriggerMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionSpamTriggerMetadataResponse(mentionTotalLimit int32) *MentionSpamTriggerMetadataResponse {
	this := MentionSpamTriggerMetadataResponse{}
	this.MentionTotalLimit = mentionTotalLimit
	return &this
}

// NewMentionSpamTriggerMetadataResponseWithDefaults instantiates a new MentionSpamTriggerMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionSpamTriggerMetadataResponseWithDefaults() *MentionSpamTriggerMetadataResponse {
	this := MentionSpamTriggerMetadataResponse{}
	return &this
}

// GetMentionTotalLimit returns the MentionTotalLimit field value
func (o *MentionSpamTriggerMetadataResponse) GetMentionTotalLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MentionTotalLimit
}

// GetMentionTotalLimitOk returns a tuple with the MentionTotalLimit field value
// and a boolean to check if the value has been set.
func (o *MentionSpamTriggerMetadataResponse) GetMentionTotalLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MentionTotalLimit, true
}

// SetMentionTotalLimit sets field value
func (o *MentionSpamTriggerMetadataResponse) SetMentionTotalLimit(v int32) {
	o.MentionTotalLimit = v
}

// GetMentionRaidProtectionEnabled returns the MentionRaidProtectionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamTriggerMetadataResponse) GetMentionRaidProtectionEnabled() bool {
	if o == nil || IsNil(o.MentionRaidProtectionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.MentionRaidProtectionEnabled.Get()
}

// GetMentionRaidProtectionEnabledOk returns a tuple with the MentionRaidProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamTriggerMetadataResponse) GetMentionRaidProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MentionRaidProtectionEnabled.Get()) {
		return nil, false
	}
	return o.MentionRaidProtectionEnabled.Get(), o.MentionRaidProtectionEnabled.IsSet()
}

// HasMentionRaidProtectionEnabled returns a boolean if a field has been set.
func (o *MentionSpamTriggerMetadataResponse) HasMentionRaidProtectionEnabled() bool {
	if o != nil && o.MentionRaidProtectionEnabled.IsSet() {
		return true
	}

	return false
}

// SetMentionRaidProtectionEnabled gets a reference to the given NullableBool and assigns it to the MentionRaidProtectionEnabled field.
func (o *MentionSpamTriggerMetadataResponse) SetMentionRaidProtectionEnabled(v bool) {
	o.MentionRaidProtectionEnabled.Set(&v)
}

// SetMentionRaidProtectionEnabledNil sets the value for MentionRaidProtectionEnabled to be an explicit nil
func (o *MentionSpamTriggerMetadataResponse) SetMentionRaidProtectionEnabledNil() {
	o.MentionRaidProtectionEnabled.Set(nil)
}

// UnsetMentionRaidProtectionEnabled ensures that no value is present for MentionRaidProtectionEnabled, not even an explicit nil
func (o *MentionSpamTriggerMetadataResponse) UnsetMentionRaidProtectionEnabled() {
	o.MentionRaidProtectionEnabled.Unset()
}

func (o MentionSpamTriggerMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionSpamTriggerMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mention_total_limit"] = o.MentionTotalLimit
	if o.MentionRaidProtectionEnabled.IsSet() {
		toSerialize["mention_raid_protection_enabled"] = o.MentionRaidProtectionEnabled.Get()
	}
	return toSerialize, nil
}

func (o *MentionSpamTriggerMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mention_total_limit",
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

	varMentionSpamTriggerMetadataResponse := _MentionSpamTriggerMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMentionSpamTriggerMetadataResponse)

	if err != nil {
		return err
	}

	*o = MentionSpamTriggerMetadataResponse(varMentionSpamTriggerMetadataResponse)

	return err
}

type NullableMentionSpamTriggerMetadataResponse struct {
	value *MentionSpamTriggerMetadataResponse
	isSet bool
}

func (v NullableMentionSpamTriggerMetadataResponse) Get() *MentionSpamTriggerMetadataResponse {
	return v.value
}

func (v *NullableMentionSpamTriggerMetadataResponse) Set(val *MentionSpamTriggerMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionSpamTriggerMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionSpamTriggerMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionSpamTriggerMetadataResponse(val *MentionSpamTriggerMetadataResponse) *NullableMentionSpamTriggerMetadataResponse {
	return &NullableMentionSpamTriggerMetadataResponse{value: val, isSet: true}
}

func (v NullableMentionSpamTriggerMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionSpamTriggerMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


