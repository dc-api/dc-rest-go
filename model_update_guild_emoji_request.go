/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// checks if the UpdateGuildEmojiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGuildEmojiRequest{}

// UpdateGuildEmojiRequest struct for UpdateGuildEmojiRequest
type UpdateGuildEmojiRequest struct {
	Name *string `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

// NewUpdateGuildEmojiRequest instantiates a new UpdateGuildEmojiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGuildEmojiRequest() *UpdateGuildEmojiRequest {
	this := UpdateGuildEmojiRequest{}
	return &this
}

// NewUpdateGuildEmojiRequestWithDefaults instantiates a new UpdateGuildEmojiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGuildEmojiRequestWithDefaults() *UpdateGuildEmojiRequest {
	this := UpdateGuildEmojiRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGuildEmojiRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildEmojiRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGuildEmojiRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGuildEmojiRequest) SetName(v string) {
	o.Name = &v
}


// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildEmojiRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildEmojiRequest) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateGuildEmojiRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UpdateGuildEmojiRequest) SetRoles(v []string) {
	o.Roles = v
}


func (o UpdateGuildEmojiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGuildEmojiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUpdateGuildEmojiRequest struct {
	value *UpdateGuildEmojiRequest
	isSet bool
}

func (v NullableUpdateGuildEmojiRequest) Get() *UpdateGuildEmojiRequest {
	return v.value
}

func (v *NullableUpdateGuildEmojiRequest) Set(val *UpdateGuildEmojiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildEmojiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildEmojiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildEmojiRequest(val *UpdateGuildEmojiRequest) *NullableUpdateGuildEmojiRequest {
	return &NullableUpdateGuildEmojiRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildEmojiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildEmojiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


