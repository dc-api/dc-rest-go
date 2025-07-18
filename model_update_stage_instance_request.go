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
)

// checks if the UpdateStageInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStageInstanceRequest{}

// UpdateStageInstanceRequest struct for UpdateStageInstanceRequest
type UpdateStageInstanceRequest struct {
	Topic *string `json:"topic,omitempty"`
	PrivacyLevel *int32 `json:"privacy_level,omitempty"`
}

// NewUpdateStageInstanceRequest instantiates a new UpdateStageInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStageInstanceRequest() *UpdateStageInstanceRequest {
	this := UpdateStageInstanceRequest{}
	return &this
}

// NewUpdateStageInstanceRequestWithDefaults instantiates a new UpdateStageInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStageInstanceRequestWithDefaults() *UpdateStageInstanceRequest {
	this := UpdateStageInstanceRequest{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *UpdateStageInstanceRequest) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStageInstanceRequest) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *UpdateStageInstanceRequest) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *UpdateStageInstanceRequest) SetTopic(v string) {
	o.Topic = &v
}


// GetPrivacyLevel returns the PrivacyLevel field value if set, zero value otherwise.
func (o *UpdateStageInstanceRequest) GetPrivacyLevel() int32 {
	if o == nil || IsNil(o.PrivacyLevel) {
		var ret int32
		return ret
	}
	return *o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStageInstanceRequest) GetPrivacyLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return o.PrivacyLevel, true
}

// HasPrivacyLevel returns a boolean if a field has been set.
func (o *UpdateStageInstanceRequest) HasPrivacyLevel() bool {
	if o != nil && !IsNil(o.PrivacyLevel) {
		return true
	}

	return false
}

// SetPrivacyLevel gets a reference to the given int32 and assigns it to the PrivacyLevel field.
func (o *UpdateStageInstanceRequest) SetPrivacyLevel(v int32) {
	o.PrivacyLevel = &v
}


func (o UpdateStageInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStageInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.PrivacyLevel) {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	return toSerialize, nil
}

type NullableUpdateStageInstanceRequest struct {
	value *UpdateStageInstanceRequest
	isSet bool
}

func (v NullableUpdateStageInstanceRequest) Get() *UpdateStageInstanceRequest {
	return v.value
}

func (v *NullableUpdateStageInstanceRequest) Set(val *UpdateStageInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStageInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStageInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStageInstanceRequest(val *UpdateStageInstanceRequest) *NullableUpdateStageInstanceRequest {
	return &NullableUpdateStageInstanceRequest{value: val, isSet: true}
}

func (v NullableUpdateStageInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStageInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


