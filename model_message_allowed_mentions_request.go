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
)

// checks if the MessageAllowedMentionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAllowedMentionsRequest{}

// MessageAllowedMentionsRequest struct for MessageAllowedMentionsRequest
type MessageAllowedMentionsRequest struct {
	Parse []string `json:"parse,omitempty"`
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
	RepliedUser NullableBool `json:"replied_user,omitempty"`
}

// NewMessageAllowedMentionsRequest instantiates a new MessageAllowedMentionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAllowedMentionsRequest() *MessageAllowedMentionsRequest {
	this := MessageAllowedMentionsRequest{}
	return &this
}

// NewMessageAllowedMentionsRequestWithDefaults instantiates a new MessageAllowedMentionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAllowedMentionsRequestWithDefaults() *MessageAllowedMentionsRequest {
	this := MessageAllowedMentionsRequest{}
	return &this
}

// GetParse returns the Parse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAllowedMentionsRequest) GetParse() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Parse
}

// GetParseOk returns a tuple with the Parse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAllowedMentionsRequest) GetParseOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parse, true
}

// HasParse returns a boolean if a field has been set.
func (o *MessageAllowedMentionsRequest) HasParse() bool {
	if o != nil && !IsNil(o.Parse) {
		return true
	}

	return false
}

// SetParse gets a reference to the given []string and assigns it to the Parse field.
func (o *MessageAllowedMentionsRequest) SetParse(v []string) {
	o.Parse = v
}


// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAllowedMentionsRequest) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAllowedMentionsRequest) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MessageAllowedMentionsRequest) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *MessageAllowedMentionsRequest) SetUsers(v []string) {
	o.Users = v
}


// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAllowedMentionsRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAllowedMentionsRequest) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *MessageAllowedMentionsRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *MessageAllowedMentionsRequest) SetRoles(v []string) {
	o.Roles = v
}


// GetRepliedUser returns the RepliedUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAllowedMentionsRequest) GetRepliedUser() bool {
	if o == nil || IsNil(o.RepliedUser.Get()) {
		var ret bool
		return ret
	}
	return *o.RepliedUser.Get()
}

// GetRepliedUserOk returns a tuple with the RepliedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAllowedMentionsRequest) GetRepliedUserOk() (*bool, bool) {
	if o == nil || IsNil(o.RepliedUser.Get()) {
		return nil, false
	}
	return o.RepliedUser.Get(), o.RepliedUser.IsSet()
}

// HasRepliedUser returns a boolean if a field has been set.
func (o *MessageAllowedMentionsRequest) HasRepliedUser() bool {
	if o != nil && o.RepliedUser.IsSet() {
		return true
	}

	return false
}

// SetRepliedUser gets a reference to the given NullableBool and assigns it to the RepliedUser field.
func (o *MessageAllowedMentionsRequest) SetRepliedUser(v bool) {
	o.RepliedUser.Set(&v)
}

// SetRepliedUserNil sets the value for RepliedUser to be an explicit nil
func (o *MessageAllowedMentionsRequest) SetRepliedUserNil() {
	o.RepliedUser.Set(nil)
}

// UnsetRepliedUser ensures that no value is present for RepliedUser, not even an explicit nil
func (o *MessageAllowedMentionsRequest) UnsetRepliedUser() {
	o.RepliedUser.Unset()
}

func (o MessageAllowedMentionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAllowedMentionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Parse != nil {
		toSerialize["parse"] = o.Parse
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.RepliedUser.IsSet() {
		toSerialize["replied_user"] = o.RepliedUser.Get()
	}
	return toSerialize, nil
}

type NullableMessageAllowedMentionsRequest struct {
	value *MessageAllowedMentionsRequest
	isSet bool
}

func (v NullableMessageAllowedMentionsRequest) Get() *MessageAllowedMentionsRequest {
	return v.value
}

func (v *NullableMessageAllowedMentionsRequest) Set(val *MessageAllowedMentionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAllowedMentionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAllowedMentionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAllowedMentionsRequest(val *MessageAllowedMentionsRequest) *NullableMessageAllowedMentionsRequest {
	return &NullableMessageAllowedMentionsRequest{value: val, isSet: true}
}

func (v NullableMessageAllowedMentionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAllowedMentionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


