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

// checks if the GatewayBotSessionStartLimitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayBotSessionStartLimitResponse{}

// GatewayBotSessionStartLimitResponse struct for GatewayBotSessionStartLimitResponse
type GatewayBotSessionStartLimitResponse struct {
	MaxConcurrency int32 `json:"max_concurrency"`
	Remaining int32 `json:"remaining"`
	ResetAfter int32 `json:"reset_after"`
	Total int32 `json:"total"`
}

type _GatewayBotSessionStartLimitResponse GatewayBotSessionStartLimitResponse

// NewGatewayBotSessionStartLimitResponse instantiates a new GatewayBotSessionStartLimitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayBotSessionStartLimitResponse(maxConcurrency int32, remaining int32, resetAfter int32, total int32) *GatewayBotSessionStartLimitResponse {
	this := GatewayBotSessionStartLimitResponse{}
	this.MaxConcurrency = maxConcurrency
	this.Remaining = remaining
	this.ResetAfter = resetAfter
	this.Total = total
	return &this
}

// NewGatewayBotSessionStartLimitResponseWithDefaults instantiates a new GatewayBotSessionStartLimitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayBotSessionStartLimitResponseWithDefaults() *GatewayBotSessionStartLimitResponse {
	this := GatewayBotSessionStartLimitResponse{}
	return &this
}

// GetMaxConcurrency returns the MaxConcurrency field value
func (o *GatewayBotSessionStartLimitResponse) GetMaxConcurrency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxConcurrency
}

// GetMaxConcurrencyOk returns a tuple with the MaxConcurrency field value
// and a boolean to check if the value has been set.
func (o *GatewayBotSessionStartLimitResponse) GetMaxConcurrencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConcurrency, true
}

// SetMaxConcurrency sets field value
func (o *GatewayBotSessionStartLimitResponse) SetMaxConcurrency(v int32) {
	o.MaxConcurrency = v
}

// GetRemaining returns the Remaining field value
func (o *GatewayBotSessionStartLimitResponse) GetRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *GatewayBotSessionStartLimitResponse) GetRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *GatewayBotSessionStartLimitResponse) SetRemaining(v int32) {
	o.Remaining = v
}

// GetResetAfter returns the ResetAfter field value
func (o *GatewayBotSessionStartLimitResponse) GetResetAfter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResetAfter
}

// GetResetAfterOk returns a tuple with the ResetAfter field value
// and a boolean to check if the value has been set.
func (o *GatewayBotSessionStartLimitResponse) GetResetAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetAfter, true
}

// SetResetAfter sets field value
func (o *GatewayBotSessionStartLimitResponse) SetResetAfter(v int32) {
	o.ResetAfter = v
}

// GetTotal returns the Total field value
func (o *GatewayBotSessionStartLimitResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GatewayBotSessionStartLimitResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GatewayBotSessionStartLimitResponse) SetTotal(v int32) {
	o.Total = v
}

func (o GatewayBotSessionStartLimitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayBotSessionStartLimitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_concurrency"] = o.MaxConcurrency
	toSerialize["remaining"] = o.Remaining
	toSerialize["reset_after"] = o.ResetAfter
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *GatewayBotSessionStartLimitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_concurrency",
		"remaining",
		"reset_after",
		"total",
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

	varGatewayBotSessionStartLimitResponse := _GatewayBotSessionStartLimitResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGatewayBotSessionStartLimitResponse)

	if err != nil {
		return err
	}

	*o = GatewayBotSessionStartLimitResponse(varGatewayBotSessionStartLimitResponse)

	return err
}

type NullableGatewayBotSessionStartLimitResponse struct {
	value *GatewayBotSessionStartLimitResponse
	isSet bool
}

func (v NullableGatewayBotSessionStartLimitResponse) Get() *GatewayBotSessionStartLimitResponse {
	return v.value
}

func (v *NullableGatewayBotSessionStartLimitResponse) Set(val *GatewayBotSessionStartLimitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayBotSessionStartLimitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayBotSessionStartLimitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayBotSessionStartLimitResponse(val *GatewayBotSessionStartLimitResponse) *NullableGatewayBotSessionStartLimitResponse {
	return &NullableGatewayBotSessionStartLimitResponse{value: val, isSet: true}
}

func (v NullableGatewayBotSessionStartLimitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayBotSessionStartLimitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


