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

// checks if the ChannelPermissionOverwriteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPermissionOverwriteResponse{}

// ChannelPermissionOverwriteResponse struct for ChannelPermissionOverwriteResponse
type ChannelPermissionOverwriteResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type NullableInt32 `json:"type"`
	Allow string `json:"allow"`
	Deny string `json:"deny"`
}

type _ChannelPermissionOverwriteResponse ChannelPermissionOverwriteResponse

// NewChannelPermissionOverwriteResponse instantiates a new ChannelPermissionOverwriteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPermissionOverwriteResponse(id string, type_ NullableInt32, allow string, deny string) *ChannelPermissionOverwriteResponse {
	this := ChannelPermissionOverwriteResponse{}
	this.Id = id
	this.Type = type_
	this.Allow = allow
	this.Deny = deny
	return &this
}

// NewChannelPermissionOverwriteResponseWithDefaults instantiates a new ChannelPermissionOverwriteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPermissionOverwriteResponseWithDefaults() *ChannelPermissionOverwriteResponse {
	this := ChannelPermissionOverwriteResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelPermissionOverwriteResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelPermissionOverwriteResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ChannelPermissionOverwriteResponse) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelPermissionOverwriteResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ChannelPermissionOverwriteResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// GetAllow returns the Allow field value
func (o *ChannelPermissionOverwriteResponse) GetAllow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetAllowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *ChannelPermissionOverwriteResponse) SetAllow(v string) {
	o.Allow = v
}

// GetDeny returns the Deny field value
func (o *ChannelPermissionOverwriteResponse) GetDeny() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetDenyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deny, true
}

// SetDeny sets field value
func (o *ChannelPermissionOverwriteResponse) SetDeny(v string) {
	o.Deny = v
}

func (o ChannelPermissionOverwriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPermissionOverwriteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type.Get()
	toSerialize["allow"] = o.Allow
	toSerialize["deny"] = o.Deny
	return toSerialize, nil
}

func (o *ChannelPermissionOverwriteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"allow",
		"deny",
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

	varChannelPermissionOverwriteResponse := _ChannelPermissionOverwriteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelPermissionOverwriteResponse)

	if err != nil {
		return err
	}

	*o = ChannelPermissionOverwriteResponse(varChannelPermissionOverwriteResponse)

	return err
}

type NullableChannelPermissionOverwriteResponse struct {
	value *ChannelPermissionOverwriteResponse
	isSet bool
}

func (v NullableChannelPermissionOverwriteResponse) Get() *ChannelPermissionOverwriteResponse {
	return v.value
}

func (v *NullableChannelPermissionOverwriteResponse) Set(val *ChannelPermissionOverwriteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPermissionOverwriteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPermissionOverwriteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPermissionOverwriteResponse(val *ChannelPermissionOverwriteResponse) *NullableChannelPermissionOverwriteResponse {
	return &NullableChannelPermissionOverwriteResponse{value: val, isSet: true}
}

func (v NullableChannelPermissionOverwriteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPermissionOverwriteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


