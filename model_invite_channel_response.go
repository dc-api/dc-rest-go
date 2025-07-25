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

// checks if the InviteChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteChannelResponse{}

// InviteChannelResponse struct for InviteChannelResponse
type InviteChannelResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Recipients []InviteChannelRecipientResponse `json:"recipients,omitempty"`
}

type _InviteChannelResponse InviteChannelResponse

// NewInviteChannelResponse instantiates a new InviteChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteChannelResponse(id string, type_ int32) *InviteChannelResponse {
	this := InviteChannelResponse{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewInviteChannelResponseWithDefaults instantiates a new InviteChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteChannelResponseWithDefaults() *InviteChannelResponse {
	this := InviteChannelResponse{}
	return &this
}

// GetId returns the Id field value
func (o *InviteChannelResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InviteChannelResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InviteChannelResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *InviteChannelResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InviteChannelResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InviteChannelResponse) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteChannelResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteChannelResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InviteChannelResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InviteChannelResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *InviteChannelResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InviteChannelResponse) UnsetName() {
	o.Name.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteChannelResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteChannelResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *InviteChannelResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *InviteChannelResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *InviteChannelResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *InviteChannelResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetRecipients returns the Recipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteChannelResponse) GetRecipients() []InviteChannelRecipientResponse {
	if o == nil {
		var ret []InviteChannelRecipientResponse
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteChannelResponse) GetRecipientsOk() ([]InviteChannelRecipientResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InviteChannelResponse) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []InviteChannelRecipientResponse and assigns it to the Recipients field.
func (o *InviteChannelResponse) SetRecipients(v []InviteChannelRecipientResponse) {
	o.Recipients = v
}


func (o InviteChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}
	return toSerialize, nil
}

func (o *InviteChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varInviteChannelResponse := _InviteChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInviteChannelResponse)

	if err != nil {
		return err
	}

	*o = InviteChannelResponse(varInviteChannelResponse)

	return err
}

type NullableInviteChannelResponse struct {
	value *InviteChannelResponse
	isSet bool
}

func (v NullableInviteChannelResponse) Get() *InviteChannelResponse {
	return v.value
}

func (v *NullableInviteChannelResponse) Set(val *InviteChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteChannelResponse(val *InviteChannelResponse) *NullableInviteChannelResponse {
	return &NullableInviteChannelResponse{value: val, isSet: true}
}

func (v NullableInviteChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


