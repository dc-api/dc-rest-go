/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the UserCommunicationDisabledActionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCommunicationDisabledActionMetadataResponse{}

// UserCommunicationDisabledActionMetadataResponse struct for UserCommunicationDisabledActionMetadataResponse
type UserCommunicationDisabledActionMetadataResponse struct {
	DurationSeconds int32 `json:"duration_seconds"`
}

type _UserCommunicationDisabledActionMetadataResponse UserCommunicationDisabledActionMetadataResponse

// NewUserCommunicationDisabledActionMetadataResponse instantiates a new UserCommunicationDisabledActionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCommunicationDisabledActionMetadataResponse(durationSeconds int32) *UserCommunicationDisabledActionMetadataResponse {
	this := UserCommunicationDisabledActionMetadataResponse{}
	this.DurationSeconds = durationSeconds
	return &this
}

// NewUserCommunicationDisabledActionMetadataResponseWithDefaults instantiates a new UserCommunicationDisabledActionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCommunicationDisabledActionMetadataResponseWithDefaults() *UserCommunicationDisabledActionMetadataResponse {
	this := UserCommunicationDisabledActionMetadataResponse{}
	return &this
}

// GetDurationSeconds returns the DurationSeconds field value
func (o *UserCommunicationDisabledActionMetadataResponse) GetDurationSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *UserCommunicationDisabledActionMetadataResponse) GetDurationSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value
func (o *UserCommunicationDisabledActionMetadataResponse) SetDurationSeconds(v int32) {
	o.DurationSeconds = v
}

func (o UserCommunicationDisabledActionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCommunicationDisabledActionMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration_seconds"] = o.DurationSeconds
	return toSerialize, nil
}

func (o *UserCommunicationDisabledActionMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration_seconds",
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

	varUserCommunicationDisabledActionMetadataResponse := _UserCommunicationDisabledActionMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCommunicationDisabledActionMetadataResponse)

	if err != nil {
		return err
	}

	*o = UserCommunicationDisabledActionMetadataResponse(varUserCommunicationDisabledActionMetadataResponse)

	return err
}

type NullableUserCommunicationDisabledActionMetadataResponse struct {
	value *UserCommunicationDisabledActionMetadataResponse
	isSet bool
}

func (v NullableUserCommunicationDisabledActionMetadataResponse) Get() *UserCommunicationDisabledActionMetadataResponse {
	return v.value
}

func (v *NullableUserCommunicationDisabledActionMetadataResponse) Set(val *UserCommunicationDisabledActionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCommunicationDisabledActionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCommunicationDisabledActionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCommunicationDisabledActionMetadataResponse(val *UserCommunicationDisabledActionMetadataResponse) *NullableUserCommunicationDisabledActionMetadataResponse {
	return &NullableUserCommunicationDisabledActionMetadataResponse{value: val, isSet: true}
}

func (v NullableUserCommunicationDisabledActionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCommunicationDisabledActionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


