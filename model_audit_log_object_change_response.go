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
)

// checks if the AuditLogObjectChangeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogObjectChangeResponse{}

// AuditLogObjectChangeResponse struct for AuditLogObjectChangeResponse
type AuditLogObjectChangeResponse struct {
	Key NullableString `json:"key,omitempty"`
	NewValue interface{} `json:"new_value,omitempty"`
	OldValue interface{} `json:"old_value,omitempty"`
}

// NewAuditLogObjectChangeResponse instantiates a new AuditLogObjectChangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogObjectChangeResponse() *AuditLogObjectChangeResponse {
	this := AuditLogObjectChangeResponse{}
	return &this
}

// NewAuditLogObjectChangeResponseWithDefaults instantiates a new AuditLogObjectChangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogObjectChangeResponseWithDefaults() *AuditLogObjectChangeResponse {
	this := AuditLogObjectChangeResponse{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLogObjectChangeResponse) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogObjectChangeResponse) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key.Get()) {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *AuditLogObjectChangeResponse) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *AuditLogObjectChangeResponse) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *AuditLogObjectChangeResponse) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *AuditLogObjectChangeResponse) UnsetKey() {
	o.Key.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLogObjectChangeResponse) GetNewValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogObjectChangeResponse) GetNewValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *AuditLogObjectChangeResponse) HasNewValue() bool {
	if o != nil && !IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given interface{} and assigns it to the NewValue field.
func (o *AuditLogObjectChangeResponse) SetNewValue(v interface{}) {
	o.NewValue = v
}


// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLogObjectChangeResponse) GetOldValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogObjectChangeResponse) GetOldValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *AuditLogObjectChangeResponse) HasOldValue() bool {
	if o != nil && !IsNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given interface{} and assigns it to the OldValue field.
func (o *AuditLogObjectChangeResponse) SetOldValue(v interface{}) {
	o.OldValue = v
}


func (o AuditLogObjectChangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogObjectChangeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.NewValue != nil {
		toSerialize["new_value"] = o.NewValue
	}
	if o.OldValue != nil {
		toSerialize["old_value"] = o.OldValue
	}
	return toSerialize, nil
}

type NullableAuditLogObjectChangeResponse struct {
	value *AuditLogObjectChangeResponse
	isSet bool
}

func (v NullableAuditLogObjectChangeResponse) Get() *AuditLogObjectChangeResponse {
	return v.value
}

func (v *NullableAuditLogObjectChangeResponse) Set(val *AuditLogObjectChangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogObjectChangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogObjectChangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogObjectChangeResponse(val *AuditLogObjectChangeResponse) *NullableAuditLogObjectChangeResponse {
	return &NullableAuditLogObjectChangeResponse{value: val, isSet: true}
}

func (v NullableAuditLogObjectChangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogObjectChangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


