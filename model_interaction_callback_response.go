/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the InteractionCallbackResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractionCallbackResponse{}

// InteractionCallbackResponse struct for InteractionCallbackResponse
type InteractionCallbackResponse struct {
	Interaction InteractionResponse `json:"interaction"`
	Resource NullableInteractionCallbackResponseResource `json:"resource,omitempty"`
}

type _InteractionCallbackResponse InteractionCallbackResponse

// NewInteractionCallbackResponse instantiates a new InteractionCallbackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractionCallbackResponse(interaction InteractionResponse) *InteractionCallbackResponse {
	this := InteractionCallbackResponse{}
	this.Interaction = interaction
	return &this
}

// NewInteractionCallbackResponseWithDefaults instantiates a new InteractionCallbackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractionCallbackResponseWithDefaults() *InteractionCallbackResponse {
	this := InteractionCallbackResponse{}
	return &this
}

// GetInteraction returns the Interaction field value
func (o *InteractionCallbackResponse) GetInteraction() InteractionResponse {
	if o == nil {
		var ret InteractionResponse
		return ret
	}

	return o.Interaction
}

// GetInteractionOk returns a tuple with the Interaction field value
// and a boolean to check if the value has been set.
func (o *InteractionCallbackResponse) GetInteractionOk() (*InteractionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interaction, true
}

// SetInteraction sets field value
func (o *InteractionCallbackResponse) SetInteraction(v InteractionResponse) {
	o.Interaction = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractionCallbackResponse) GetResource() InteractionCallbackResponseResource {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret InteractionCallbackResponseResource
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractionCallbackResponse) GetResourceOk() (*InteractionCallbackResponseResource, bool) {
	if o == nil || IsNil(o.Resource.Get()) {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *InteractionCallbackResponse) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableInteractionCallbackResponseResource and assigns it to the Resource field.
func (o *InteractionCallbackResponse) SetResource(v InteractionCallbackResponseResource) {
	o.Resource.Set(&v)
}

// SetResourceNil sets the value for Resource to be an explicit nil
func (o *InteractionCallbackResponse) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *InteractionCallbackResponse) UnsetResource() {
	o.Resource.Unset()
}

func (o InteractionCallbackResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractionCallbackResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interaction"] = o.Interaction
	if o.Resource.IsSet() {
		toSerialize["resource"] = o.Resource.Get()
	}
	return toSerialize, nil
}

func (o *InteractionCallbackResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interaction",
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

	varInteractionCallbackResponse := _InteractionCallbackResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInteractionCallbackResponse)

	if err != nil {
		return err
	}

	*o = InteractionCallbackResponse(varInteractionCallbackResponse)

	return err
}

type NullableInteractionCallbackResponse struct {
	value *InteractionCallbackResponse
	isSet bool
}

func (v NullableInteractionCallbackResponse) Get() *InteractionCallbackResponse {
	return v.value
}

func (v *NullableInteractionCallbackResponse) Set(val *InteractionCallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionCallbackResponse(val *InteractionCallbackResponse) *NullableInteractionCallbackResponse {
	return &NullableInteractionCallbackResponse{value: val, isSet: true}
}

func (v NullableInteractionCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


