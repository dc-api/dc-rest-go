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
	"bytes"
	"fmt"
)

// checks if the SeparatorComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeparatorComponentForMessageRequest{}

// SeparatorComponentForMessageRequest struct for SeparatorComponentForMessageRequest
type SeparatorComponentForMessageRequest struct {
	Type int32 `json:"type"`
	Spacing *int32 `json:"spacing,omitempty"`
	Divider NullableBool `json:"divider,omitempty"`
}

type _SeparatorComponentForMessageRequest SeparatorComponentForMessageRequest

// NewSeparatorComponentForMessageRequest instantiates a new SeparatorComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeparatorComponentForMessageRequest(type_ int32) *SeparatorComponentForMessageRequest {
	this := SeparatorComponentForMessageRequest{}
	this.Type = type_
	return &this
}

// NewSeparatorComponentForMessageRequestWithDefaults instantiates a new SeparatorComponentForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeparatorComponentForMessageRequestWithDefaults() *SeparatorComponentForMessageRequest {
	this := SeparatorComponentForMessageRequest{}
	return &this
}

// GetType returns the Type field value
func (o *SeparatorComponentForMessageRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeparatorComponentForMessageRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SeparatorComponentForMessageRequest) SetType(v int32) {
	o.Type = v
}

// GetSpacing returns the Spacing field value if set, zero value otherwise.
func (o *SeparatorComponentForMessageRequest) GetSpacing() int32 {
	if o == nil || IsNil(o.Spacing) {
		var ret int32
		return ret
	}
	return *o.Spacing
}

// GetSpacingOk returns a tuple with the Spacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeparatorComponentForMessageRequest) GetSpacingOk() (*int32, bool) {
	if o == nil || IsNil(o.Spacing) {
		return nil, false
	}
	return o.Spacing, true
}

// HasSpacing returns a boolean if a field has been set.
func (o *SeparatorComponentForMessageRequest) HasSpacing() bool {
	if o != nil && !IsNil(o.Spacing) {
		return true
	}

	return false
}

// SetSpacing gets a reference to the given int32 and assigns it to the Spacing field.
func (o *SeparatorComponentForMessageRequest) SetSpacing(v int32) {
	o.Spacing = &v
}


// GetDivider returns the Divider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeparatorComponentForMessageRequest) GetDivider() bool {
	if o == nil || IsNil(o.Divider.Get()) {
		var ret bool
		return ret
	}
	return *o.Divider.Get()
}

// GetDividerOk returns a tuple with the Divider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeparatorComponentForMessageRequest) GetDividerOk() (*bool, bool) {
	if o == nil || IsNil(o.Divider.Get()) {
		return nil, false
	}
	return o.Divider.Get(), o.Divider.IsSet()
}

// HasDivider returns a boolean if a field has been set.
func (o *SeparatorComponentForMessageRequest) HasDivider() bool {
	if o != nil && o.Divider.IsSet() {
		return true
	}

	return false
}

// SetDivider gets a reference to the given NullableBool and assigns it to the Divider field.
func (o *SeparatorComponentForMessageRequest) SetDivider(v bool) {
	o.Divider.Set(&v)
}

// SetDividerNil sets the value for Divider to be an explicit nil
func (o *SeparatorComponentForMessageRequest) SetDividerNil() {
	o.Divider.Set(nil)
}

// UnsetDivider ensures that no value is present for Divider, not even an explicit nil
func (o *SeparatorComponentForMessageRequest) UnsetDivider() {
	o.Divider.Unset()
}

func (o SeparatorComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeparatorComponentForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Spacing) {
		toSerialize["spacing"] = o.Spacing
	}
	if o.Divider.IsSet() {
		toSerialize["divider"] = o.Divider.Get()
	}
	return toSerialize, nil
}

func (o *SeparatorComponentForMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSeparatorComponentForMessageRequest := _SeparatorComponentForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSeparatorComponentForMessageRequest)

	if err != nil {
		return err
	}

	*o = SeparatorComponentForMessageRequest(varSeparatorComponentForMessageRequest)

	return err
}

type NullableSeparatorComponentForMessageRequest struct {
	value *SeparatorComponentForMessageRequest
	isSet bool
}

func (v NullableSeparatorComponentForMessageRequest) Get() *SeparatorComponentForMessageRequest {
	return v.value
}

func (v *NullableSeparatorComponentForMessageRequest) Set(val *SeparatorComponentForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSeparatorComponentForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSeparatorComponentForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeparatorComponentForMessageRequest(val *SeparatorComponentForMessageRequest) *NullableSeparatorComponentForMessageRequest {
	return &NullableSeparatorComponentForMessageRequest{value: val, isSet: true}
}

func (v NullableSeparatorComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeparatorComponentForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


