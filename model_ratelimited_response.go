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

// checks if the RatelimitedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatelimitedResponse{}

// RatelimitedResponse Ratelimit error object returned by the Discord API
type RatelimitedResponse struct {
	// Discord internal error code. See error code reference
	Code int32 `json:"code"`
	// Human-readable error message
	Message string `json:"message"`
	// The number of seconds to wait before retrying your request
	RetryAfter float32 `json:"retry_after"`
	// Whether you are being ratelimited by the global ratelimit or a per-endpoint ratelimit
	Global bool `json:"global"`
}

type _RatelimitedResponse RatelimitedResponse

// NewRatelimitedResponse instantiates a new RatelimitedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatelimitedResponse(code int32, message string, retryAfter float32, global bool) *RatelimitedResponse {
	this := RatelimitedResponse{}
	this.Code = code
	this.Message = message
	this.RetryAfter = retryAfter
	this.Global = global
	return &this
}

// NewRatelimitedResponseWithDefaults instantiates a new RatelimitedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatelimitedResponseWithDefaults() *RatelimitedResponse {
	this := RatelimitedResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *RatelimitedResponse) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RatelimitedResponse) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RatelimitedResponse) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *RatelimitedResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RatelimitedResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RatelimitedResponse) SetMessage(v string) {
	o.Message = v
}

// GetRetryAfter returns the RetryAfter field value
func (o *RatelimitedResponse) GetRetryAfter() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RetryAfter
}

// GetRetryAfterOk returns a tuple with the RetryAfter field value
// and a boolean to check if the value has been set.
func (o *RatelimitedResponse) GetRetryAfterOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryAfter, true
}

// SetRetryAfter sets field value
func (o *RatelimitedResponse) SetRetryAfter(v float32) {
	o.RetryAfter = v
}

// GetGlobal returns the Global field value
func (o *RatelimitedResponse) GetGlobal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value
// and a boolean to check if the value has been set.
func (o *RatelimitedResponse) GetGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Global, true
}

// SetGlobal sets field value
func (o *RatelimitedResponse) SetGlobal(v bool) {
	o.Global = v
}

func (o RatelimitedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatelimitedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["retry_after"] = o.RetryAfter
	toSerialize["global"] = o.Global
	return toSerialize, nil
}

func (o *RatelimitedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
		"retry_after",
		"global",
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

	varRatelimitedResponse := _RatelimitedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRatelimitedResponse)

	if err != nil {
		return err
	}

	*o = RatelimitedResponse(varRatelimitedResponse)

	return err
}

type NullableRatelimitedResponse struct {
	value *RatelimitedResponse
	isSet bool
}

func (v NullableRatelimitedResponse) Get() *RatelimitedResponse {
	return v.value
}

func (v *NullableRatelimitedResponse) Set(val *RatelimitedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRatelimitedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRatelimitedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatelimitedResponse(val *RatelimitedResponse) *NullableRatelimitedResponse {
	return &NullableRatelimitedResponse{value: val, isSet: true}
}

func (v NullableRatelimitedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatelimitedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


