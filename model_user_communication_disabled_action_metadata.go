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

// checks if the UserCommunicationDisabledActionMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCommunicationDisabledActionMetadata{}

// UserCommunicationDisabledActionMetadata struct for UserCommunicationDisabledActionMetadata
type UserCommunicationDisabledActionMetadata struct {
	DurationSeconds NullableInt32 `json:"duration_seconds,omitempty"`
}

// NewUserCommunicationDisabledActionMetadata instantiates a new UserCommunicationDisabledActionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCommunicationDisabledActionMetadata() *UserCommunicationDisabledActionMetadata {
	this := UserCommunicationDisabledActionMetadata{}
	return &this
}

// NewUserCommunicationDisabledActionMetadataWithDefaults instantiates a new UserCommunicationDisabledActionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCommunicationDisabledActionMetadataWithDefaults() *UserCommunicationDisabledActionMetadata {
	this := UserCommunicationDisabledActionMetadata{}
	return &this
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCommunicationDisabledActionMetadata) GetDurationSeconds() int32 {
	if o == nil || IsNil(o.DurationSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCommunicationDisabledActionMetadata) GetDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationSeconds.Get()) {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *UserCommunicationDisabledActionMetadata) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableInt32 and assigns it to the DurationSeconds field.
func (o *UserCommunicationDisabledActionMetadata) SetDurationSeconds(v int32) {
	o.DurationSeconds.Set(&v)
}

// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *UserCommunicationDisabledActionMetadata) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *UserCommunicationDisabledActionMetadata) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

func (o UserCommunicationDisabledActionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCommunicationDisabledActionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DurationSeconds.IsSet() {
		toSerialize["duration_seconds"] = o.DurationSeconds.Get()
	}
	return toSerialize, nil
}

type NullableUserCommunicationDisabledActionMetadata struct {
	value *UserCommunicationDisabledActionMetadata
	isSet bool
}

func (v NullableUserCommunicationDisabledActionMetadata) Get() *UserCommunicationDisabledActionMetadata {
	return v.value
}

func (v *NullableUserCommunicationDisabledActionMetadata) Set(val *UserCommunicationDisabledActionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCommunicationDisabledActionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCommunicationDisabledActionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCommunicationDisabledActionMetadata(val *UserCommunicationDisabledActionMetadata) *NullableUserCommunicationDisabledActionMetadata {
	return &NullableUserCommunicationDisabledActionMetadata{value: val, isSet: true}
}

func (v NullableUserCommunicationDisabledActionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCommunicationDisabledActionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


