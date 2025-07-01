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

// checks if the DefaultKeywordListTriggerMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultKeywordListTriggerMetadataResponse{}

// DefaultKeywordListTriggerMetadataResponse struct for DefaultKeywordListTriggerMetadataResponse
type DefaultKeywordListTriggerMetadataResponse struct {
	AllowList []string `json:"allow_list"`
	Presets []int32 `json:"presets"`
}

type _DefaultKeywordListTriggerMetadataResponse DefaultKeywordListTriggerMetadataResponse

// NewDefaultKeywordListTriggerMetadataResponse instantiates a new DefaultKeywordListTriggerMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultKeywordListTriggerMetadataResponse(allowList []string, presets []int32) *DefaultKeywordListTriggerMetadataResponse {
	this := DefaultKeywordListTriggerMetadataResponse{}
	this.AllowList = allowList
	this.Presets = presets
	return &this
}

// NewDefaultKeywordListTriggerMetadataResponseWithDefaults instantiates a new DefaultKeywordListTriggerMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultKeywordListTriggerMetadataResponseWithDefaults() *DefaultKeywordListTriggerMetadataResponse {
	this := DefaultKeywordListTriggerMetadataResponse{}
	return &this
}

// GetAllowList returns the AllowList field value
func (o *DefaultKeywordListTriggerMetadataResponse) GetAllowList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListTriggerMetadataResponse) GetAllowListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowList, true
}

// SetAllowList sets field value
func (o *DefaultKeywordListTriggerMetadataResponse) SetAllowList(v []string) {
	o.AllowList = v
}

// GetPresets returns the Presets field value
func (o *DefaultKeywordListTriggerMetadataResponse) GetPresets() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListTriggerMetadataResponse) GetPresetsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presets, true
}

// SetPresets sets field value
func (o *DefaultKeywordListTriggerMetadataResponse) SetPresets(v []int32) {
	o.Presets = v
}

func (o DefaultKeywordListTriggerMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultKeywordListTriggerMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allow_list"] = o.AllowList
	toSerialize["presets"] = o.Presets
	return toSerialize, nil
}

func (o *DefaultKeywordListTriggerMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allow_list",
		"presets",
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

	varDefaultKeywordListTriggerMetadataResponse := _DefaultKeywordListTriggerMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDefaultKeywordListTriggerMetadataResponse)

	if err != nil {
		return err
	}

	*o = DefaultKeywordListTriggerMetadataResponse(varDefaultKeywordListTriggerMetadataResponse)

	return err
}

type NullableDefaultKeywordListTriggerMetadataResponse struct {
	value *DefaultKeywordListTriggerMetadataResponse
	isSet bool
}

func (v NullableDefaultKeywordListTriggerMetadataResponse) Get() *DefaultKeywordListTriggerMetadataResponse {
	return v.value
}

func (v *NullableDefaultKeywordListTriggerMetadataResponse) Set(val *DefaultKeywordListTriggerMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListTriggerMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListTriggerMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListTriggerMetadataResponse(val *DefaultKeywordListTriggerMetadataResponse) *NullableDefaultKeywordListTriggerMetadataResponse {
	return &NullableDefaultKeywordListTriggerMetadataResponse{value: val, isSet: true}
}

func (v NullableDefaultKeywordListTriggerMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListTriggerMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


