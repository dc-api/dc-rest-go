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

// checks if the PartialGuildSubscriptionIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialGuildSubscriptionIntegrationResponse{}

// PartialGuildSubscriptionIntegrationResponse struct for PartialGuildSubscriptionIntegrationResponse
type PartialGuildSubscriptionIntegrationResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type NullableString `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Account NullableAccountResponse `json:"account,omitempty"`
}

type _PartialGuildSubscriptionIntegrationResponse PartialGuildSubscriptionIntegrationResponse

// NewPartialGuildSubscriptionIntegrationResponse instantiates a new PartialGuildSubscriptionIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialGuildSubscriptionIntegrationResponse(id string, type_ NullableString) *PartialGuildSubscriptionIntegrationResponse {
	this := PartialGuildSubscriptionIntegrationResponse{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewPartialGuildSubscriptionIntegrationResponseWithDefaults instantiates a new PartialGuildSubscriptionIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialGuildSubscriptionIntegrationResponseWithDefaults() *PartialGuildSubscriptionIntegrationResponse {
	this := PartialGuildSubscriptionIntegrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PartialGuildSubscriptionIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartialGuildSubscriptionIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PartialGuildSubscriptionIntegrationResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGuildSubscriptionIntegrationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *PartialGuildSubscriptionIntegrationResponse) SetType(v string) {
	o.Type.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGuildSubscriptionIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGuildSubscriptionIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PartialGuildSubscriptionIntegrationResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) UnsetName() {
	o.Name.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGuildSubscriptionIntegrationResponse) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account.Get()) {
		var ret AccountResponse
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGuildSubscriptionIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil || IsNil(o.Account.Get()) {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableAccountResponse and assigns it to the Account field.
func (o *PartialGuildSubscriptionIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) UnsetAccount() {
	o.Account.Unset()
}

func (o PartialGuildSubscriptionIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialGuildSubscriptionIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	return toSerialize, nil
}

func (o *PartialGuildSubscriptionIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPartialGuildSubscriptionIntegrationResponse := _PartialGuildSubscriptionIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPartialGuildSubscriptionIntegrationResponse)

	if err != nil {
		return err
	}

	*o = PartialGuildSubscriptionIntegrationResponse(varPartialGuildSubscriptionIntegrationResponse)

	return err
}

type NullablePartialGuildSubscriptionIntegrationResponse struct {
	value *PartialGuildSubscriptionIntegrationResponse
	isSet bool
}

func (v NullablePartialGuildSubscriptionIntegrationResponse) Get() *PartialGuildSubscriptionIntegrationResponse {
	return v.value
}

func (v *NullablePartialGuildSubscriptionIntegrationResponse) Set(val *PartialGuildSubscriptionIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialGuildSubscriptionIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialGuildSubscriptionIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialGuildSubscriptionIntegrationResponse(val *PartialGuildSubscriptionIntegrationResponse) *NullablePartialGuildSubscriptionIntegrationResponse {
	return &NullablePartialGuildSubscriptionIntegrationResponse{value: val, isSet: true}
}

func (v NullablePartialGuildSubscriptionIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialGuildSubscriptionIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


