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

// checks if the FlagToChannelAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlagToChannelAction{}

// FlagToChannelAction struct for FlagToChannelAction
type FlagToChannelAction struct {
	Type int32 `json:"type"`
	Metadata FlagToChannelActionMetadata `json:"metadata"`
}

type _FlagToChannelAction FlagToChannelAction

// NewFlagToChannelAction instantiates a new FlagToChannelAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagToChannelAction(type_ int32, metadata FlagToChannelActionMetadata) *FlagToChannelAction {
	this := FlagToChannelAction{}
	this.Type = type_
	this.Metadata = metadata
	return &this
}

// NewFlagToChannelActionWithDefaults instantiates a new FlagToChannelAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagToChannelActionWithDefaults() *FlagToChannelAction {
	this := FlagToChannelAction{}
	return &this
}

// GetType returns the Type field value
func (o *FlagToChannelAction) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlagToChannelAction) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlagToChannelAction) SetType(v int32) {
	o.Type = v
}

// GetMetadata returns the Metadata field value
func (o *FlagToChannelAction) GetMetadata() FlagToChannelActionMetadata {
	if o == nil {
		var ret FlagToChannelActionMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *FlagToChannelAction) GetMetadataOk() (*FlagToChannelActionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *FlagToChannelAction) SetMetadata(v FlagToChannelActionMetadata) {
	o.Metadata = v
}

func (o FlagToChannelAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlagToChannelAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *FlagToChannelAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"metadata",
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

	varFlagToChannelAction := _FlagToChannelAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlagToChannelAction)

	if err != nil {
		return err
	}

	*o = FlagToChannelAction(varFlagToChannelAction)

	return err
}

type NullableFlagToChannelAction struct {
	value *FlagToChannelAction
	isSet bool
}

func (v NullableFlagToChannelAction) Get() *FlagToChannelAction {
	return v.value
}

func (v *NullableFlagToChannelAction) Set(val *FlagToChannelAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagToChannelAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagToChannelAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagToChannelAction(val *FlagToChannelAction) *NullableFlagToChannelAction {
	return &NullableFlagToChannelAction{value: val, isSet: true}
}

func (v NullableFlagToChannelAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagToChannelAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


