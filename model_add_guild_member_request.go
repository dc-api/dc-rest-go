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

// checks if the AddGuildMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddGuildMemberRequest{}

// AddGuildMemberRequest struct for AddGuildMemberRequest
type AddGuildMemberRequest struct {
	Nick NullableString `json:"nick,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Mute NullableBool `json:"mute,omitempty"`
	Deaf NullableBool `json:"deaf,omitempty"`
	AccessToken string `json:"access_token"`
	Flags NullableInt32 `json:"flags,omitempty"`
}

type _AddGuildMemberRequest AddGuildMemberRequest

// NewAddGuildMemberRequest instantiates a new AddGuildMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGuildMemberRequest(accessToken string) *AddGuildMemberRequest {
	this := AddGuildMemberRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewAddGuildMemberRequestWithDefaults instantiates a new AddGuildMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGuildMemberRequestWithDefaults() *AddGuildMemberRequest {
	this := AddGuildMemberRequest{}
	return &this
}

// GetNick returns the Nick field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddGuildMemberRequest) GetNick() string {
	if o == nil || IsNil(o.Nick.Get()) {
		var ret string
		return ret
	}
	return *o.Nick.Get()
}

// GetNickOk returns a tuple with the Nick field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddGuildMemberRequest) GetNickOk() (*string, bool) {
	if o == nil || IsNil(o.Nick.Get()) {
		return nil, false
	}
	return o.Nick.Get(), o.Nick.IsSet()
}

// HasNick returns a boolean if a field has been set.
func (o *AddGuildMemberRequest) HasNick() bool {
	if o != nil && o.Nick.IsSet() {
		return true
	}

	return false
}

// SetNick gets a reference to the given NullableString and assigns it to the Nick field.
func (o *AddGuildMemberRequest) SetNick(v string) {
	o.Nick.Set(&v)
}

// SetNickNil sets the value for Nick to be an explicit nil
func (o *AddGuildMemberRequest) SetNickNil() {
	o.Nick.Set(nil)
}

// UnsetNick ensures that no value is present for Nick, not even an explicit nil
func (o *AddGuildMemberRequest) UnsetNick() {
	o.Nick.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddGuildMemberRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddGuildMemberRequest) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AddGuildMemberRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *AddGuildMemberRequest) SetRoles(v []string) {
	o.Roles = v
}


// GetMute returns the Mute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddGuildMemberRequest) GetMute() bool {
	if o == nil || IsNil(o.Mute.Get()) {
		var ret bool
		return ret
	}
	return *o.Mute.Get()
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddGuildMemberRequest) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute.Get()) {
		return nil, false
	}
	return o.Mute.Get(), o.Mute.IsSet()
}

// HasMute returns a boolean if a field has been set.
func (o *AddGuildMemberRequest) HasMute() bool {
	if o != nil && o.Mute.IsSet() {
		return true
	}

	return false
}

// SetMute gets a reference to the given NullableBool and assigns it to the Mute field.
func (o *AddGuildMemberRequest) SetMute(v bool) {
	o.Mute.Set(&v)
}

// SetMuteNil sets the value for Mute to be an explicit nil
func (o *AddGuildMemberRequest) SetMuteNil() {
	o.Mute.Set(nil)
}

// UnsetMute ensures that no value is present for Mute, not even an explicit nil
func (o *AddGuildMemberRequest) UnsetMute() {
	o.Mute.Unset()
}

// GetDeaf returns the Deaf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddGuildMemberRequest) GetDeaf() bool {
	if o == nil || IsNil(o.Deaf.Get()) {
		var ret bool
		return ret
	}
	return *o.Deaf.Get()
}

// GetDeafOk returns a tuple with the Deaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddGuildMemberRequest) GetDeafOk() (*bool, bool) {
	if o == nil || IsNil(o.Deaf.Get()) {
		return nil, false
	}
	return o.Deaf.Get(), o.Deaf.IsSet()
}

// HasDeaf returns a boolean if a field has been set.
func (o *AddGuildMemberRequest) HasDeaf() bool {
	if o != nil && o.Deaf.IsSet() {
		return true
	}

	return false
}

// SetDeaf gets a reference to the given NullableBool and assigns it to the Deaf field.
func (o *AddGuildMemberRequest) SetDeaf(v bool) {
	o.Deaf.Set(&v)
}

// SetDeafNil sets the value for Deaf to be an explicit nil
func (o *AddGuildMemberRequest) SetDeafNil() {
	o.Deaf.Set(nil)
}

// UnsetDeaf ensures that no value is present for Deaf, not even an explicit nil
func (o *AddGuildMemberRequest) UnsetDeaf() {
	o.Deaf.Unset()
}

// GetAccessToken returns the AccessToken field value
func (o *AddGuildMemberRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AddGuildMemberRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AddGuildMemberRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddGuildMemberRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddGuildMemberRequest) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *AddGuildMemberRequest) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *AddGuildMemberRequest) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *AddGuildMemberRequest) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *AddGuildMemberRequest) UnsetFlags() {
	o.Flags.Unset()
}

func (o AddGuildMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddGuildMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Nick.IsSet() {
		toSerialize["nick"] = o.Nick.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Mute.IsSet() {
		toSerialize["mute"] = o.Mute.Get()
	}
	if o.Deaf.IsSet() {
		toSerialize["deaf"] = o.Deaf.Get()
	}
	toSerialize["access_token"] = o.AccessToken
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	return toSerialize, nil
}

func (o *AddGuildMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
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

	varAddGuildMemberRequest := _AddGuildMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddGuildMemberRequest)

	if err != nil {
		return err
	}

	*o = AddGuildMemberRequest(varAddGuildMemberRequest)

	return err
}

type NullableAddGuildMemberRequest struct {
	value *AddGuildMemberRequest
	isSet bool
}

func (v NullableAddGuildMemberRequest) Get() *AddGuildMemberRequest {
	return v.value
}

func (v *NullableAddGuildMemberRequest) Set(val *AddGuildMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGuildMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGuildMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGuildMemberRequest(val *AddGuildMemberRequest) *NullableAddGuildMemberRequest {
	return &NullableAddGuildMemberRequest{value: val, isSet: true}
}

func (v NullableAddGuildMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGuildMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


