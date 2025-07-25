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

// checks if the VanityURLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VanityURLResponse{}

// VanityURLResponse struct for VanityURLResponse
type VanityURLResponse struct {
	Code NullableString `json:"code,omitempty"`
	Uses int32 `json:"uses"`
	Error NullableVanityURLErrorResponse `json:"error,omitempty"`
}

type _VanityURLResponse VanityURLResponse

// NewVanityURLResponse instantiates a new VanityURLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVanityURLResponse(uses int32) *VanityURLResponse {
	this := VanityURLResponse{}
	this.Uses = uses
	return &this
}

// NewVanityURLResponseWithDefaults instantiates a new VanityURLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVanityURLResponseWithDefaults() *VanityURLResponse {
	this := VanityURLResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VanityURLResponse) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VanityURLResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code.Get()) {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *VanityURLResponse) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *VanityURLResponse) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *VanityURLResponse) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *VanityURLResponse) UnsetCode() {
	o.Code.Unset()
}

// GetUses returns the Uses field value
func (o *VanityURLResponse) GetUses() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uses
}

// GetUsesOk returns a tuple with the Uses field value
// and a boolean to check if the value has been set.
func (o *VanityURLResponse) GetUsesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uses, true
}

// SetUses sets field value
func (o *VanityURLResponse) SetUses(v int32) {
	o.Uses = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VanityURLResponse) GetError() VanityURLErrorResponse {
	if o == nil || IsNil(o.Error.Get()) {
		var ret VanityURLErrorResponse
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VanityURLResponse) GetErrorOk() (*VanityURLErrorResponse, bool) {
	if o == nil || IsNil(o.Error.Get()) {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *VanityURLResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableVanityURLErrorResponse and assigns it to the Error field.
func (o *VanityURLResponse) SetError(v VanityURLErrorResponse) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *VanityURLResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *VanityURLResponse) UnsetError() {
	o.Error.Unset()
}

func (o VanityURLResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VanityURLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	toSerialize["uses"] = o.Uses
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

func (o *VanityURLResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uses",
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

	varVanityURLResponse := _VanityURLResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVanityURLResponse)

	if err != nil {
		return err
	}

	*o = VanityURLResponse(varVanityURLResponse)

	return err
}

type NullableVanityURLResponse struct {
	value *VanityURLResponse
	isSet bool
}

func (v NullableVanityURLResponse) Get() *VanityURLResponse {
	return v.value
}

func (v *NullableVanityURLResponse) Set(val *VanityURLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVanityURLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVanityURLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVanityURLResponse(val *VanityURLResponse) *NullableVanityURLResponse {
	return &NullableVanityURLResponse{value: val, isSet: true}
}

func (v NullableVanityURLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVanityURLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


