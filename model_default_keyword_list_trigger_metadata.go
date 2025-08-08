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
)

// checks if the DefaultKeywordListTriggerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultKeywordListTriggerMetadata{}

// DefaultKeywordListTriggerMetadata struct for DefaultKeywordListTriggerMetadata
type DefaultKeywordListTriggerMetadata struct {
	AllowList []string `json:"allow_list,omitempty"`
	Presets []int32 `json:"presets,omitempty"`
}

// NewDefaultKeywordListTriggerMetadata instantiates a new DefaultKeywordListTriggerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultKeywordListTriggerMetadata() *DefaultKeywordListTriggerMetadata {
	this := DefaultKeywordListTriggerMetadata{}
	return &this
}

// NewDefaultKeywordListTriggerMetadataWithDefaults instantiates a new DefaultKeywordListTriggerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultKeywordListTriggerMetadataWithDefaults() *DefaultKeywordListTriggerMetadata {
	this := DefaultKeywordListTriggerMetadata{}
	return &this
}

// GetAllowList returns the AllowList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListTriggerMetadata) GetAllowList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListTriggerMetadata) GetAllowListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *DefaultKeywordListTriggerMetadata) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *DefaultKeywordListTriggerMetadata) SetAllowList(v []string) {
	o.AllowList = v
}


// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListTriggerMetadata) GetPresets() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListTriggerMetadata) GetPresetsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *DefaultKeywordListTriggerMetadata) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []int32 and assigns it to the Presets field.
func (o *DefaultKeywordListTriggerMetadata) SetPresets(v []int32) {
	o.Presets = v
}


func (o DefaultKeywordListTriggerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultKeywordListTriggerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowList != nil {
		toSerialize["allow_list"] = o.AllowList
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	return toSerialize, nil
}

type NullableDefaultKeywordListTriggerMetadata struct {
	value *DefaultKeywordListTriggerMetadata
	isSet bool
}

func (v NullableDefaultKeywordListTriggerMetadata) Get() *DefaultKeywordListTriggerMetadata {
	return v.value
}

func (v *NullableDefaultKeywordListTriggerMetadata) Set(val *DefaultKeywordListTriggerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListTriggerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListTriggerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListTriggerMetadata(val *DefaultKeywordListTriggerMetadata) *NullableDefaultKeywordListTriggerMetadata {
	return &NullableDefaultKeywordListTriggerMetadata{value: val, isSet: true}
}

func (v NullableDefaultKeywordListTriggerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListTriggerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


