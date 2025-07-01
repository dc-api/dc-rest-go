/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the WidgetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetResponse{}

// WidgetResponse struct for WidgetResponse
type WidgetResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	InstantInvite NullableString `json:"instant_invite,omitempty"`
	Channels []WidgetChannel `json:"channels"`
	Members []WidgetMember `json:"members"`
	PresenceCount int32 `json:"presence_count"`
}

type _WidgetResponse WidgetResponse

// NewWidgetResponse instantiates a new WidgetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetResponse(id string, name string, channels []WidgetChannel, members []WidgetMember, presenceCount int32) *WidgetResponse {
	this := WidgetResponse{}
	this.Id = id
	this.Name = name
	this.Channels = channels
	this.Members = members
	this.PresenceCount = presenceCount
	return &this
}

// NewWidgetResponseWithDefaults instantiates a new WidgetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetResponseWithDefaults() *WidgetResponse {
	this := WidgetResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WidgetResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WidgetResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WidgetResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WidgetResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WidgetResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WidgetResponse) SetName(v string) {
	o.Name = v
}

// GetInstantInvite returns the InstantInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetResponse) GetInstantInvite() string {
	if o == nil || IsNil(o.InstantInvite.Get()) {
		var ret string
		return ret
	}
	return *o.InstantInvite.Get()
}

// GetInstantInviteOk returns a tuple with the InstantInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetResponse) GetInstantInviteOk() (*string, bool) {
	if o == nil || IsNil(o.InstantInvite.Get()) {
		return nil, false
	}
	return o.InstantInvite.Get(), o.InstantInvite.IsSet()
}

// HasInstantInvite returns a boolean if a field has been set.
func (o *WidgetResponse) HasInstantInvite() bool {
	if o != nil && o.InstantInvite.IsSet() {
		return true
	}

	return false
}

// SetInstantInvite gets a reference to the given NullableString and assigns it to the InstantInvite field.
func (o *WidgetResponse) SetInstantInvite(v string) {
	o.InstantInvite.Set(&v)
}

// SetInstantInviteNil sets the value for InstantInvite to be an explicit nil
func (o *WidgetResponse) SetInstantInviteNil() {
	o.InstantInvite.Set(nil)
}

// UnsetInstantInvite ensures that no value is present for InstantInvite, not even an explicit nil
func (o *WidgetResponse) UnsetInstantInvite() {
	o.InstantInvite.Unset()
}

// GetChannels returns the Channels field value
func (o *WidgetResponse) GetChannels() []WidgetChannel {
	if o == nil {
		var ret []WidgetChannel
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *WidgetResponse) GetChannelsOk() ([]WidgetChannel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *WidgetResponse) SetChannels(v []WidgetChannel) {
	o.Channels = v
}

// GetMembers returns the Members field value
func (o *WidgetResponse) GetMembers() []WidgetMember {
	if o == nil {
		var ret []WidgetMember
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *WidgetResponse) GetMembersOk() ([]WidgetMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *WidgetResponse) SetMembers(v []WidgetMember) {
	o.Members = v
}

// GetPresenceCount returns the PresenceCount field value
func (o *WidgetResponse) GetPresenceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PresenceCount
}

// GetPresenceCountOk returns a tuple with the PresenceCount field value
// and a boolean to check if the value has been set.
func (o *WidgetResponse) GetPresenceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresenceCount, true
}

// SetPresenceCount sets field value
func (o *WidgetResponse) SetPresenceCount(v int32) {
	o.PresenceCount = v
}

func (o WidgetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.InstantInvite.IsSet() {
		toSerialize["instant_invite"] = o.InstantInvite.Get()
	}
	toSerialize["channels"] = o.Channels
	toSerialize["members"] = o.Members
	toSerialize["presence_count"] = o.PresenceCount
	return toSerialize, nil
}

func (o *WidgetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"channels",
		"members",
		"presence_count",
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

	varWidgetResponse := _WidgetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWidgetResponse)

	if err != nil {
		return err
	}

	*o = WidgetResponse(varWidgetResponse)

	return err
}

type NullableWidgetResponse struct {
	value *WidgetResponse
	isSet bool
}

func (v NullableWidgetResponse) Get() *WidgetResponse {
	return v.value
}

func (v *NullableWidgetResponse) Set(val *WidgetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetResponse(val *WidgetResponse) *NullableWidgetResponse {
	return &NullableWidgetResponse{value: val, isSet: true}
}

func (v NullableWidgetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


