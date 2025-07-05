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

// checks if the ConnectedAccountIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedAccountIntegrationResponse{}

// ConnectedAccountIntegrationResponse struct for ConnectedAccountIntegrationResponse
type ConnectedAccountIntegrationResponse struct {
	Id string `json:"id"`
	Type NullableString `json:"type"`
	Account AccountResponse `json:"account"`
	Guild ConnectedAccountGuildResponse `json:"guild"`
}

type _ConnectedAccountIntegrationResponse ConnectedAccountIntegrationResponse

// NewConnectedAccountIntegrationResponse instantiates a new ConnectedAccountIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAccountIntegrationResponse(id string, type_ NullableString, account AccountResponse, guild ConnectedAccountGuildResponse) *ConnectedAccountIntegrationResponse {
	this := ConnectedAccountIntegrationResponse{}
	this.Id = id
	this.Type = type_
	this.Account = account
	this.Guild = guild
	return &this
}

// NewConnectedAccountIntegrationResponseWithDefaults instantiates a new ConnectedAccountIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAccountIntegrationResponseWithDefaults() *ConnectedAccountIntegrationResponse {
	this := ConnectedAccountIntegrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectedAccountIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectedAccountIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConnectedAccountIntegrationResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedAccountIntegrationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ConnectedAccountIntegrationResponse) SetType(v string) {
	o.Type.Set(&v)
}

// GetAccount returns the Account field value
func (o *ConnectedAccountIntegrationResponse) GetAccount() AccountResponse {
	if o == nil {
		var ret AccountResponse
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ConnectedAccountIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account = v
}

// GetGuild returns the Guild field value
func (o *ConnectedAccountIntegrationResponse) GetGuild() ConnectedAccountGuildResponse {
	if o == nil {
		var ret ConnectedAccountGuildResponse
		return ret
	}

	return o.Guild
}

// GetGuildOk returns a tuple with the Guild field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountIntegrationResponse) GetGuildOk() (*ConnectedAccountGuildResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guild, true
}

// SetGuild sets field value
func (o *ConnectedAccountIntegrationResponse) SetGuild(v ConnectedAccountGuildResponse) {
	o.Guild = v
}

func (o ConnectedAccountIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedAccountIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type.Get()
	toSerialize["account"] = o.Account
	toSerialize["guild"] = o.Guild
	return toSerialize, nil
}

func (o *ConnectedAccountIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"account",
		"guild",
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

	varConnectedAccountIntegrationResponse := _ConnectedAccountIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectedAccountIntegrationResponse)

	if err != nil {
		return err
	}

	*o = ConnectedAccountIntegrationResponse(varConnectedAccountIntegrationResponse)

	return err
}

type NullableConnectedAccountIntegrationResponse struct {
	value *ConnectedAccountIntegrationResponse
	isSet bool
}

func (v NullableConnectedAccountIntegrationResponse) Get() *ConnectedAccountIntegrationResponse {
	return v.value
}

func (v *NullableConnectedAccountIntegrationResponse) Set(val *ConnectedAccountIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAccountIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAccountIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAccountIntegrationResponse(val *ConnectedAccountIntegrationResponse) *NullableConnectedAccountIntegrationResponse {
	return &NullableConnectedAccountIntegrationResponse{value: val, isSet: true}
}

func (v NullableConnectedAccountIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAccountIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


