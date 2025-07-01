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

// checks if the ChannelPermissionOverwriteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPermissionOverwriteRequest{}

// ChannelPermissionOverwriteRequest struct for ChannelPermissionOverwriteRequest
type ChannelPermissionOverwriteRequest struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type NullableInt32 `json:"type,omitempty"`
	Allow NullableInt32 `json:"allow,omitempty"`
	Deny NullableInt32 `json:"deny,omitempty"`
}

type _ChannelPermissionOverwriteRequest ChannelPermissionOverwriteRequest

// NewChannelPermissionOverwriteRequest instantiates a new ChannelPermissionOverwriteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPermissionOverwriteRequest(id string) *ChannelPermissionOverwriteRequest {
	this := ChannelPermissionOverwriteRequest{}
	this.Id = id
	return &this
}

// NewChannelPermissionOverwriteRequestWithDefaults instantiates a new ChannelPermissionOverwriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPermissionOverwriteRequestWithDefaults() *ChannelPermissionOverwriteRequest {
	this := ChannelPermissionOverwriteRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelPermissionOverwriteRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelPermissionOverwriteRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelPermissionOverwriteRequest) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelPermissionOverwriteRequest) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ChannelPermissionOverwriteRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *ChannelPermissionOverwriteRequest) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *ChannelPermissionOverwriteRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ChannelPermissionOverwriteRequest) UnsetType() {
	o.Type.Unset()
}

// GetAllow returns the Allow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelPermissionOverwriteRequest) GetAllow() int32 {
	if o == nil || IsNil(o.Allow.Get()) {
		var ret int32
		return ret
	}
	return *o.Allow.Get()
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelPermissionOverwriteRequest) GetAllowOk() (*int32, bool) {
	if o == nil || IsNil(o.Allow.Get()) {
		return nil, false
	}
	return o.Allow.Get(), o.Allow.IsSet()
}

// HasAllow returns a boolean if a field has been set.
func (o *ChannelPermissionOverwriteRequest) HasAllow() bool {
	if o != nil && o.Allow.IsSet() {
		return true
	}

	return false
}

// SetAllow gets a reference to the given NullableInt32 and assigns it to the Allow field.
func (o *ChannelPermissionOverwriteRequest) SetAllow(v int32) {
	o.Allow.Set(&v)
}

// SetAllowNil sets the value for Allow to be an explicit nil
func (o *ChannelPermissionOverwriteRequest) SetAllowNil() {
	o.Allow.Set(nil)
}

// UnsetAllow ensures that no value is present for Allow, not even an explicit nil
func (o *ChannelPermissionOverwriteRequest) UnsetAllow() {
	o.Allow.Unset()
}

// GetDeny returns the Deny field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelPermissionOverwriteRequest) GetDeny() int32 {
	if o == nil || IsNil(o.Deny.Get()) {
		var ret int32
		return ret
	}
	return *o.Deny.Get()
}

// GetDenyOk returns a tuple with the Deny field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelPermissionOverwriteRequest) GetDenyOk() (*int32, bool) {
	if o == nil || IsNil(o.Deny.Get()) {
		return nil, false
	}
	return o.Deny.Get(), o.Deny.IsSet()
}

// HasDeny returns a boolean if a field has been set.
func (o *ChannelPermissionOverwriteRequest) HasDeny() bool {
	if o != nil && o.Deny.IsSet() {
		return true
	}

	return false
}

// SetDeny gets a reference to the given NullableInt32 and assigns it to the Deny field.
func (o *ChannelPermissionOverwriteRequest) SetDeny(v int32) {
	o.Deny.Set(&v)
}

// SetDenyNil sets the value for Deny to be an explicit nil
func (o *ChannelPermissionOverwriteRequest) SetDenyNil() {
	o.Deny.Set(nil)
}

// UnsetDeny ensures that no value is present for Deny, not even an explicit nil
func (o *ChannelPermissionOverwriteRequest) UnsetDeny() {
	o.Deny.Unset()
}

func (o ChannelPermissionOverwriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPermissionOverwriteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Allow.IsSet() {
		toSerialize["allow"] = o.Allow.Get()
	}
	if o.Deny.IsSet() {
		toSerialize["deny"] = o.Deny.Get()
	}
	return toSerialize, nil
}

func (o *ChannelPermissionOverwriteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varChannelPermissionOverwriteRequest := _ChannelPermissionOverwriteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelPermissionOverwriteRequest)

	if err != nil {
		return err
	}

	*o = ChannelPermissionOverwriteRequest(varChannelPermissionOverwriteRequest)

	return err
}

type NullableChannelPermissionOverwriteRequest struct {
	value *ChannelPermissionOverwriteRequest
	isSet bool
}

func (v NullableChannelPermissionOverwriteRequest) Get() *ChannelPermissionOverwriteRequest {
	return v.value
}

func (v *NullableChannelPermissionOverwriteRequest) Set(val *ChannelPermissionOverwriteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPermissionOverwriteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPermissionOverwriteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPermissionOverwriteRequest(val *ChannelPermissionOverwriteRequest) *NullableChannelPermissionOverwriteRequest {
	return &NullableChannelPermissionOverwriteRequest{value: val, isSet: true}
}

func (v NullableChannelPermissionOverwriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPermissionOverwriteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


