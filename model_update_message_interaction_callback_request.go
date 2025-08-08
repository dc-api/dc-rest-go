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

// checks if the UpdateMessageInteractionCallbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessageInteractionCallbackRequest{}

// UpdateMessageInteractionCallbackRequest struct for UpdateMessageInteractionCallbackRequest
type UpdateMessageInteractionCallbackRequest struct {
	Type NullableInt32 `json:"type"`
	Data NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial `json:"data,omitempty"`
}

type _UpdateMessageInteractionCallbackRequest UpdateMessageInteractionCallbackRequest

// NewUpdateMessageInteractionCallbackRequest instantiates a new UpdateMessageInteractionCallbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessageInteractionCallbackRequest(type_ NullableInt32) *UpdateMessageInteractionCallbackRequest {
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
// If the value is explicit nil, the zero value for int32 will be returned
func (o *UpdateMessageInteractionCallbackRequest) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMessageInteractionCallbackRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *UpdateMessageInteractionCallbackRequest) SetType(v int32) {
	o.Type.Set(&v)
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMessageInteractionCallbackRequest) GetData() IncomingWebhookUpdateForInteractionCallbackRequestPartial {
	if o == nil || IsNil(o.Data.Get()) {
		var ret IncomingWebhookUpdateForInteractionCallbackRequestPartial
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMessageInteractionCallbackRequest) GetDataOk() (*IncomingWebhookUpdateForInteractionCallbackRequestPartial, bool) {
	if o == nil || IsNil(o.Data.Get()) {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *UpdateMessageInteractionCallbackRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial and assigns it to the Data field.
func (o *UpdateMessageInteractionCallbackRequest) SetData(v IncomingWebhookUpdateForInteractionCallbackRequestPartial) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil
func (o *UpdateMessageInteractionCallbackRequest) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *UpdateMessageInteractionCallbackRequest) UnsetData() {
	o.Data.Unset()
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
	toSerialize["type"] = o.Type.Get()
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
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


