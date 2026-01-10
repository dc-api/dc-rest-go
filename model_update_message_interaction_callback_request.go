/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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

// checks if the UpdateMessageInteractionCallbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessageInteractionCallbackRequest{}

// UpdateMessageInteractionCallbackRequest struct for UpdateMessageInteractionCallbackRequest
type UpdateMessageInteractionCallbackRequest struct {
	Type InteractionCallbackTypes `json:"type"`
	Data *UpdateMessageInteractionCallbackRequestData `json:"data,omitempty"`
}

type _UpdateMessageInteractionCallbackRequest UpdateMessageInteractionCallbackRequest

// NewUpdateMessageInteractionCallbackRequest instantiates a new UpdateMessageInteractionCallbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessageInteractionCallbackRequest(type_ InteractionCallbackTypes) *UpdateMessageInteractionCallbackRequest {
	this := UpdateMessageInteractionCallbackRequest{}
	this.Type = type_
	return &this
}

// NewUpdateMessageInteractionCallbackRequestWithDefaults instantiates a new UpdateMessageInteractionCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMessageInteractionCallbackRequestWithDefaults() *UpdateMessageInteractionCallbackRequest {
	this := UpdateMessageInteractionCallbackRequest{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateMessageInteractionCallbackRequest) GetType() InteractionCallbackTypes {
	if o == nil {
		var ret InteractionCallbackTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateMessageInteractionCallbackRequest) GetTypeOk() (*InteractionCallbackTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateMessageInteractionCallbackRequest) SetType(v InteractionCallbackTypes) {
	o.Type = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateMessageInteractionCallbackRequest) GetData() UpdateMessageInteractionCallbackRequestData {
	if o == nil || IsNil(o.Data) {
		var ret UpdateMessageInteractionCallbackRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMessageInteractionCallbackRequest) GetDataOk() (*UpdateMessageInteractionCallbackRequestData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateMessageInteractionCallbackRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateMessageInteractionCallbackRequestData and assigns it to the Data field.
func (o *UpdateMessageInteractionCallbackRequest) SetData(v UpdateMessageInteractionCallbackRequestData) {
	o.Data = &v
}


func (o UpdateMessageInteractionCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMessageInteractionCallbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

func (o *UpdateMessageInteractionCallbackRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateMessageInteractionCallbackRequest := _UpdateMessageInteractionCallbackRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMessageInteractionCallbackRequest)

	if err != nil {
		return err
	}

	*o = UpdateMessageInteractionCallbackRequest(varUpdateMessageInteractionCallbackRequest)

	return err
}

type NullableUpdateMessageInteractionCallbackRequest struct {
	value *UpdateMessageInteractionCallbackRequest
	isSet bool
}

func (v NullableUpdateMessageInteractionCallbackRequest) Get() *UpdateMessageInteractionCallbackRequest {
	return v.value
}

func (v *NullableUpdateMessageInteractionCallbackRequest) Set(val *UpdateMessageInteractionCallbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMessageInteractionCallbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMessageInteractionCallbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMessageInteractionCallbackRequest(val *UpdateMessageInteractionCallbackRequest) *NullableUpdateMessageInteractionCallbackRequest {
	return &NullableUpdateMessageInteractionCallbackRequest{value: val, isSet: true}
}

func (v NullableUpdateMessageInteractionCallbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMessageInteractionCallbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

