/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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

// checks if the SeparatorComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeparatorComponentResponse{}

// SeparatorComponentResponse struct for SeparatorComponentResponse
type SeparatorComponentResponse struct {
	Type int32 `json:"type"`
	Id int32 `json:"id"`
	Spacing int32 `json:"spacing"`
	Divider bool `json:"divider"`
}

type _SeparatorComponentResponse SeparatorComponentResponse

// NewSeparatorComponentResponse instantiates a new SeparatorComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeparatorComponentResponse(type_ int32, id int32, spacing int32, divider bool) *SeparatorComponentResponse {
	this := SeparatorComponentResponse{}
	this.Type = type_
	this.Id = id
	this.Spacing = spacing
	this.Divider = divider
	return &this
}

// NewSeparatorComponentResponseWithDefaults instantiates a new SeparatorComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeparatorComponentResponseWithDefaults() *SeparatorComponentResponse {
	this := SeparatorComponentResponse{}
	return &this
}

// GetType returns the Type field value
func (o *SeparatorComponentResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeparatorComponentResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SeparatorComponentResponse) SetType(v int32) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SeparatorComponentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SeparatorComponentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SeparatorComponentResponse) SetId(v int32) {
	o.Id = v
}

// GetSpacing returns the Spacing field value
func (o *SeparatorComponentResponse) GetSpacing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spacing
}

// GetSpacingOk returns a tuple with the Spacing field value
// and a boolean to check if the value has been set.
func (o *SeparatorComponentResponse) GetSpacingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spacing, true
}

// SetSpacing sets field value
func (o *SeparatorComponentResponse) SetSpacing(v int32) {
	o.Spacing = v
}

// GetDivider returns the Divider field value
func (o *SeparatorComponentResponse) GetDivider() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Divider
}

// GetDividerOk returns a tuple with the Divider field value
// and a boolean to check if the value has been set.
func (o *SeparatorComponentResponse) GetDividerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Divider, true
}

// SetDivider sets field value
func (o *SeparatorComponentResponse) SetDivider(v bool) {
	o.Divider = v
}

func (o SeparatorComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeparatorComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["spacing"] = o.Spacing
	toSerialize["divider"] = o.Divider
	return toSerialize, nil
}

func (o *SeparatorComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"spacing",
		"divider",
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

	varSeparatorComponentResponse := _SeparatorComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSeparatorComponentResponse)

	if err != nil {
		return err
	}

	*o = SeparatorComponentResponse(varSeparatorComponentResponse)

	return err
}

type NullableSeparatorComponentResponse struct {
	value *SeparatorComponentResponse
	isSet bool
}

func (v NullableSeparatorComponentResponse) Get() *SeparatorComponentResponse {
	return v.value
}

func (v *NullableSeparatorComponentResponse) Set(val *SeparatorComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSeparatorComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSeparatorComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeparatorComponentResponse(val *SeparatorComponentResponse) *NullableSeparatorComponentResponse {
	return &NullableSeparatorComponentResponse{value: val, isSet: true}
}

func (v NullableSeparatorComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeparatorComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


