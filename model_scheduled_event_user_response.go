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

// checks if the ScheduledEventUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledEventUserResponse{}

// ScheduledEventUserResponse struct for ScheduledEventUserResponse
type ScheduledEventUserResponse struct {
	GuildScheduledEventId string `json:"guild_scheduled_event_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	UserId string `json:"user_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	User NullableUserResponse `json:"user,omitempty"`
	Member NullableGuildMemberResponse `json:"member,omitempty"`
}

type _ScheduledEventUserResponse ScheduledEventUserResponse

// NewScheduledEventUserResponse instantiates a new ScheduledEventUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledEventUserResponse(guildScheduledEventId string, userId string) *ScheduledEventUserResponse {
	this := ScheduledEventUserResponse{}
	this.GuildScheduledEventId = guildScheduledEventId
	this.UserId = userId
	return &this
}

// NewScheduledEventUserResponseWithDefaults instantiates a new ScheduledEventUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledEventUserResponseWithDefaults() *ScheduledEventUserResponse {
	this := ScheduledEventUserResponse{}
	return &this
}

// GetGuildScheduledEventId returns the GuildScheduledEventId field value
func (o *ScheduledEventUserResponse) GetGuildScheduledEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildScheduledEventId
}

// GetGuildScheduledEventIdOk returns a tuple with the GuildScheduledEventId field value
// and a boolean to check if the value has been set.
func (o *ScheduledEventUserResponse) GetGuildScheduledEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildScheduledEventId, true
}

// SetGuildScheduledEventId sets field value
func (o *ScheduledEventUserResponse) SetGuildScheduledEventId(v string) {
	o.GuildScheduledEventId = v
}

// GetUserId returns the UserId field value
func (o *ScheduledEventUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ScheduledEventUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ScheduledEventUserResponse) SetUserId(v string) {
	o.UserId = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledEventUserResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledEventUserResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ScheduledEventUserResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *ScheduledEventUserResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *ScheduledEventUserResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ScheduledEventUserResponse) UnsetUser() {
	o.User.Unset()
}

// GetMember returns the Member field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledEventUserResponse) GetMember() GuildMemberResponse {
	if o == nil || IsNil(o.Member.Get()) {
		var ret GuildMemberResponse
		return ret
	}
	return *o.Member.Get()
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledEventUserResponse) GetMemberOk() (*GuildMemberResponse, bool) {
	if o == nil || IsNil(o.Member.Get()) {
		return nil, false
	}
	return o.Member.Get(), o.Member.IsSet()
}

// HasMember returns a boolean if a field has been set.
func (o *ScheduledEventUserResponse) HasMember() bool {
	if o != nil && o.Member.IsSet() {
		return true
	}

	return false
}

// SetMember gets a reference to the given NullableGuildMemberResponse and assigns it to the Member field.
func (o *ScheduledEventUserResponse) SetMember(v GuildMemberResponse) {
	o.Member.Set(&v)
}

// SetMemberNil sets the value for Member to be an explicit nil
func (o *ScheduledEventUserResponse) SetMemberNil() {
	o.Member.Set(nil)
}

// UnsetMember ensures that no value is present for Member, not even an explicit nil
func (o *ScheduledEventUserResponse) UnsetMember() {
	o.Member.Unset()
}

func (o ScheduledEventUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledEventUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guild_scheduled_event_id"] = o.GuildScheduledEventId
	toSerialize["user_id"] = o.UserId
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Member.IsSet() {
		toSerialize["member"] = o.Member.Get()
	}
	return toSerialize, nil
}

func (o *ScheduledEventUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guild_scheduled_event_id",
		"user_id",
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

	varScheduledEventUserResponse := _ScheduledEventUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScheduledEventUserResponse)

	if err != nil {
		return err
	}

	*o = ScheduledEventUserResponse(varScheduledEventUserResponse)

	return err
}

type NullableScheduledEventUserResponse struct {
	value *ScheduledEventUserResponse
	isSet bool
}

func (v NullableScheduledEventUserResponse) Get() *ScheduledEventUserResponse {
	return v.value
}

func (v *NullableScheduledEventUserResponse) Set(val *ScheduledEventUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledEventUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledEventUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledEventUserResponse(val *ScheduledEventUserResponse) *NullableScheduledEventUserResponse {
	return &NullableScheduledEventUserResponse{value: val, isSet: true}
}

func (v NullableScheduledEventUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledEventUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


