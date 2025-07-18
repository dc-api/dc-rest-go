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

// checks if the UpdateApplicationEmojiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApplicationEmojiRequest{}

// UpdateApplicationEmojiRequest struct for UpdateApplicationEmojiRequest
type UpdateApplicationEmojiRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewUpdateApplicationEmojiRequest instantiates a new UpdateApplicationEmojiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApplicationEmojiRequest() *UpdateApplicationEmojiRequest {
	this := UpdateApplicationEmojiRequest{}
	return &this
}

// NewUpdateApplicationEmojiRequestWithDefaults instantiates a new UpdateApplicationEmojiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApplicationEmojiRequestWithDefaults() *UpdateApplicationEmojiRequest {
	this := UpdateApplicationEmojiRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateApplicationEmojiRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationEmojiRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateApplicationEmojiRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateApplicationEmojiRequest) SetName(v string) {
	o.Name = &v
}


func (o UpdateApplicationEmojiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApplicationEmojiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateApplicationEmojiRequest struct {
	value *UpdateApplicationEmojiRequest
	isSet bool
}

func (v NullableUpdateApplicationEmojiRequest) Get() *UpdateApplicationEmojiRequest {
	return v.value
}

func (v *NullableUpdateApplicationEmojiRequest) Set(val *UpdateApplicationEmojiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationEmojiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationEmojiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationEmojiRequest(val *UpdateApplicationEmojiRequest) *NullableUpdateApplicationEmojiRequest {
	return &NullableUpdateApplicationEmojiRequest{value: val, isSet: true}
}

func (v NullableUpdateApplicationEmojiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationEmojiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


