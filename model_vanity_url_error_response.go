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
	"bytes"
	"fmt"
)

// checks if the VanityURLErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VanityURLErrorResponse{}

// VanityURLErrorResponse struct for VanityURLErrorResponse
type VanityURLErrorResponse struct {
	Message string `json:"message"`
	Code int32 `json:"code"`
}

type _VanityURLErrorResponse VanityURLErrorResponse

// NewVanityURLErrorResponse instantiates a new VanityURLErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVanityURLErrorResponse(message string, code int32) *VanityURLErrorResponse {
	this := VanityURLErrorResponse{}
	this.Message = message
	this.Code = code
	return &this
}

// NewVanityURLErrorResponseWithDefaults instantiates a new VanityURLErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVanityURLErrorResponseWithDefaults() *VanityURLErrorResponse {
	this := VanityURLErrorResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *VanityURLErrorResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *VanityURLErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *VanityURLErrorResponse) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value
func (o *VanityURLErrorResponse) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *VanityURLErrorResponse) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *VanityURLErrorResponse) SetCode(v int32) {
	o.Code = v
}

func (o VanityURLErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VanityURLErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *VanityURLErrorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"code",
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

	varVanityURLErrorResponse := _VanityURLErrorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVanityURLErrorResponse)

	if err != nil {
		return err
	}

	*o = VanityURLErrorResponse(varVanityURLErrorResponse)

	return err
}

type NullableVanityURLErrorResponse struct {
	value *VanityURLErrorResponse
	isSet bool
}

func (v NullableVanityURLErrorResponse) Get() *VanityURLErrorResponse {
	return v.value
}

func (v *NullableVanityURLErrorResponse) Set(val *VanityURLErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVanityURLErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVanityURLErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVanityURLErrorResponse(val *VanityURLErrorResponse) *NullableVanityURLErrorResponse {
	return &NullableVanityURLErrorResponse{value: val, isSet: true}
}

func (v NullableVanityURLErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVanityURLErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


