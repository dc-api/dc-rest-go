/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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

// checks if the BotAddGuildMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BotAddGuildMemberRequest{}

// BotAddGuildMemberRequest struct for BotAddGuildMemberRequest
type BotAddGuildMemberRequest struct {
	Nick *string `json:"nick,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Mute *bool `json:"mute,omitempty"`
	Deaf *bool `json:"deaf,omitempty"`
	AccessToken string `json:"access_token"`
	Flags *int32 `json:"flags,omitempty"`
}

type _BotAddGuildMemberRequest BotAddGuildMemberRequest

// NewBotAddGuildMemberRequest instantiates a new BotAddGuildMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotAddGuildMemberRequest(accessToken string) *BotAddGuildMemberRequest {
	this := BotAddGuildMemberRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewBotAddGuildMemberRequestWithDefaults instantiates a new BotAddGuildMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotAddGuildMemberRequestWithDefaults() *BotAddGuildMemberRequest {
	this := BotAddGuildMemberRequest{}
	return &this
}

// GetNick returns the Nick field value if set, zero value otherwise.
func (o *BotAddGuildMemberRequest) GetNick() string {
	if o == nil || IsNil(o.Nick) {
		var ret string
		return ret
	}
	return *o.Nick
}

// GetNickOk returns a tuple with the Nick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotAddGuildMemberRequest) GetNickOk() (*string, bool) {
	if o == nil || IsNil(o.Nick) {
		return nil, false
	}
	return o.Nick, true
}

// HasNick returns a boolean if a field has been set.
func (o *BotAddGuildMemberRequest) HasNick() bool {
	if o != nil && !IsNil(o.Nick) {
		return true
	}

	return false
}

// SetNick gets a reference to the given string and assigns it to the Nick field.
func (o *BotAddGuildMemberRequest) SetNick(v string) {
	o.Nick = &v
}


// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *BotAddGuildMemberRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotAddGuildMemberRequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *BotAddGuildMemberRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *BotAddGuildMemberRequest) SetRoles(v []string) {
	o.Roles = v
}


// GetMute returns the Mute field value if set, zero value otherwise.
func (o *BotAddGuildMemberRequest) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotAddGuildMemberRequest) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *BotAddGuildMemberRequest) HasMute() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *BotAddGuildMemberRequest) SetMute(v bool) {
	o.Mute = &v
}


// GetDeaf returns the Deaf field value if set, zero value otherwise.
func (o *BotAddGuildMemberRequest) GetDeaf() bool {
	if o == nil || IsNil(o.Deaf) {
		var ret bool
		return ret
	}
	return *o.Deaf
}

// GetDeafOk returns a tuple with the Deaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotAddGuildMemberRequest) GetDeafOk() (*bool, bool) {
	if o == nil || IsNil(o.Deaf) {
		return nil, false
	}
	return o.Deaf, true
}

// HasDeaf returns a boolean if a field has been set.
func (o *BotAddGuildMemberRequest) HasDeaf() bool {
	if o != nil && !IsNil(o.Deaf) {
		return true
	}

	return false
}

// SetDeaf gets a reference to the given bool and assigns it to the Deaf field.
func (o *BotAddGuildMemberRequest) SetDeaf(v bool) {
	o.Deaf = &v
}


// GetAccessToken returns the AccessToken field value
func (o *BotAddGuildMemberRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *BotAddGuildMemberRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *BotAddGuildMemberRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *BotAddGuildMemberRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotAddGuildMemberRequest) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *BotAddGuildMemberRequest) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *BotAddGuildMemberRequest) SetFlags(v int32) {
	o.Flags = &v
}


func (o BotAddGuildMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BotAddGuildMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nick) {
		toSerialize["nick"] = o.Nick
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Mute) {
		toSerialize["mute"] = o.Mute
	}
	if !IsNil(o.Deaf) {
		toSerialize["deaf"] = o.Deaf
	}
	toSerialize["access_token"] = o.AccessToken
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *BotAddGuildMemberRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBotAddGuildMemberRequest := _BotAddGuildMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBotAddGuildMemberRequest)

	if err != nil {
		return err
	}

	*o = BotAddGuildMemberRequest(varBotAddGuildMemberRequest)

	return err
}

type NullableBotAddGuildMemberRequest struct {
	value *BotAddGuildMemberRequest
	isSet bool
}

func (v NullableBotAddGuildMemberRequest) Get() *BotAddGuildMemberRequest {
	return v.value
}

func (v *NullableBotAddGuildMemberRequest) Set(val *BotAddGuildMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBotAddGuildMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBotAddGuildMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotAddGuildMemberRequest(val *BotAddGuildMemberRequest) *NullableBotAddGuildMemberRequest {
	return &NullableBotAddGuildMemberRequest{value: val, isSet: true}
}

func (v NullableBotAddGuildMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotAddGuildMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

