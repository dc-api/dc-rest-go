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

// checks if the GuildBanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildBanResponse{}

// GuildBanResponse struct for GuildBanResponse
type GuildBanResponse struct {
	User UserResponse `json:"user"`
	Reason NullableString `json:"reason,omitempty"`
}

type _GuildBanResponse GuildBanResponse

// NewGuildBanResponse instantiates a new GuildBanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildBanResponse(user UserResponse) *GuildBanResponse {
	this := GuildBanResponse{}
	this.User = user
	return &this
}

// NewGuildBanResponseWithDefaults instantiates a new GuildBanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildBanResponseWithDefaults() *GuildBanResponse {
	this := GuildBanResponse{}
	return &this
}

// GetUser returns the User field value
func (o *GuildBanResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GuildBanResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GuildBanResponse) SetUser(v UserResponse) {
	o.User = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildBanResponse) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildBanResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason.Get()) {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *GuildBanResponse) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *GuildBanResponse) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *GuildBanResponse) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *GuildBanResponse) UnsetReason() {
	o.Reason.Unset()
}

func (o GuildBanResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildBanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	return toSerialize, nil
}

func (o *GuildBanResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varGuildBanResponse := _GuildBanResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildBanResponse)

	if err != nil {
		return err
	}

	*o = GuildBanResponse(varGuildBanResponse)

	return err
}

type NullableGuildBanResponse struct {
	value *GuildBanResponse
	isSet bool
}

func (v NullableGuildBanResponse) Get() *GuildBanResponse {
	return v.value
}

func (v *NullableGuildBanResponse) Set(val *GuildBanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildBanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildBanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildBanResponse(val *GuildBanResponse) *NullableGuildBanResponse {
	return &NullableGuildBanResponse{value: val, isSet: true}
}

func (v NullableGuildBanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildBanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


