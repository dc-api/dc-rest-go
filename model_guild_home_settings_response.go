/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the GuildHomeSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildHomeSettingsResponse{}

// GuildHomeSettingsResponse struct for GuildHomeSettingsResponse
type GuildHomeSettingsResponse struct {
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Enabled bool `json:"enabled"`
	WelcomeMessage NullableWelcomeMessageResponse `json:"welcome_message,omitempty"`
	NewMemberActions []*NewMemberActionResponse `json:"new_member_actions,omitempty"`
	ResourceChannels []*ResourceChannelResponse `json:"resource_channels,omitempty"`
}

type _GuildHomeSettingsResponse GuildHomeSettingsResponse

// NewGuildHomeSettingsResponse instantiates a new GuildHomeSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildHomeSettingsResponse(guildId string, enabled bool) *GuildHomeSettingsResponse {
	this := GuildHomeSettingsResponse{}
	this.GuildId = guildId
	this.Enabled = enabled
	return &this
}

// NewGuildHomeSettingsResponseWithDefaults instantiates a new GuildHomeSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildHomeSettingsResponseWithDefaults() *GuildHomeSettingsResponse {
	this := GuildHomeSettingsResponse{}
	return &this
}

// GetGuildId returns the GuildId field value
func (o *GuildHomeSettingsResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *GuildHomeSettingsResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *GuildHomeSettingsResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetEnabled returns the Enabled field value
func (o *GuildHomeSettingsResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GuildHomeSettingsResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GuildHomeSettingsResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildHomeSettingsResponse) GetWelcomeMessage() WelcomeMessageResponse {
	if o == nil || IsNil(o.WelcomeMessage.Get()) {
		var ret WelcomeMessageResponse
		return ret
	}
	return *o.WelcomeMessage.Get()
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildHomeSettingsResponse) GetWelcomeMessageOk() (*WelcomeMessageResponse, bool) {
	if o == nil || IsNil(o.WelcomeMessage.Get()) {
		return nil, false
	}
	return o.WelcomeMessage.Get(), o.WelcomeMessage.IsSet()
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *GuildHomeSettingsResponse) HasWelcomeMessage() bool {
	if o != nil && o.WelcomeMessage.IsSet() {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given NullableWelcomeMessageResponse and assigns it to the WelcomeMessage field.
func (o *GuildHomeSettingsResponse) SetWelcomeMessage(v WelcomeMessageResponse) {
	o.WelcomeMessage.Set(&v)
}

// SetWelcomeMessageNil sets the value for WelcomeMessage to be an explicit nil
func (o *GuildHomeSettingsResponse) SetWelcomeMessageNil() {
	o.WelcomeMessage.Set(nil)
}

// UnsetWelcomeMessage ensures that no value is present for WelcomeMessage, not even an explicit nil
func (o *GuildHomeSettingsResponse) UnsetWelcomeMessage() {
	o.WelcomeMessage.Unset()
}

// GetNewMemberActions returns the NewMemberActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildHomeSettingsResponse) GetNewMemberActions() []*NewMemberActionResponse {
	if o == nil {
		var ret []*NewMemberActionResponse
		return ret
	}
	return o.NewMemberActions
}

// GetNewMemberActionsOk returns a tuple with the NewMemberActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildHomeSettingsResponse) GetNewMemberActionsOk() ([]*NewMemberActionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewMemberActions, true
}

// HasNewMemberActions returns a boolean if a field has been set.
func (o *GuildHomeSettingsResponse) HasNewMemberActions() bool {
	if o != nil && !IsNil(o.NewMemberActions) {
		return true
	}

	return false
}

// SetNewMemberActions gets a reference to the given []*NewMemberActionResponse and assigns it to the NewMemberActions field.
func (o *GuildHomeSettingsResponse) SetNewMemberActions(v []*NewMemberActionResponse) {
	o.NewMemberActions = v
}


// GetResourceChannels returns the ResourceChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildHomeSettingsResponse) GetResourceChannels() []*ResourceChannelResponse {
	if o == nil {
		var ret []*ResourceChannelResponse
		return ret
	}
	return o.ResourceChannels
}

// GetResourceChannelsOk returns a tuple with the ResourceChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildHomeSettingsResponse) GetResourceChannelsOk() ([]*ResourceChannelResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceChannels, true
}

// HasResourceChannels returns a boolean if a field has been set.
func (o *GuildHomeSettingsResponse) HasResourceChannels() bool {
	if o != nil && !IsNil(o.ResourceChannels) {
		return true
	}

	return false
}

// SetResourceChannels gets a reference to the given []*ResourceChannelResponse and assigns it to the ResourceChannels field.
func (o *GuildHomeSettingsResponse) SetResourceChannels(v []*ResourceChannelResponse) {
	o.ResourceChannels = v
}


func (o GuildHomeSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildHomeSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["enabled"] = o.Enabled
	if o.WelcomeMessage.IsSet() {
		toSerialize["welcome_message"] = o.WelcomeMessage.Get()
	}
	if o.NewMemberActions != nil {
		toSerialize["new_member_actions"] = o.NewMemberActions
	}
	if o.ResourceChannels != nil {
		toSerialize["resource_channels"] = o.ResourceChannels
	}
	return toSerialize, nil
}

func (o *GuildHomeSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guild_id",
		"enabled",
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

	varGuildHomeSettingsResponse := _GuildHomeSettingsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildHomeSettingsResponse)

	if err != nil {
		return err
	}

	*o = GuildHomeSettingsResponse(varGuildHomeSettingsResponse)

	return err
}

type NullableGuildHomeSettingsResponse struct {
	value *GuildHomeSettingsResponse
	isSet bool
}

func (v NullableGuildHomeSettingsResponse) Get() *GuildHomeSettingsResponse {
	return v.value
}

func (v *NullableGuildHomeSettingsResponse) Set(val *GuildHomeSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildHomeSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildHomeSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildHomeSettingsResponse(val *GuildHomeSettingsResponse) *NullableGuildHomeSettingsResponse {
	return &NullableGuildHomeSettingsResponse{value: val, isSet: true}
}

func (v NullableGuildHomeSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildHomeSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


