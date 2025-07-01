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

// checks if the UserGuildOnboardingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGuildOnboardingResponse{}

// UserGuildOnboardingResponse struct for UserGuildOnboardingResponse
type UserGuildOnboardingResponse struct {
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Prompts []OnboardingPromptResponse `json:"prompts"`
	DefaultChannelIds []string `json:"default_channel_ids"`
	Enabled bool `json:"enabled"`
}

type _UserGuildOnboardingResponse UserGuildOnboardingResponse

// NewUserGuildOnboardingResponse instantiates a new UserGuildOnboardingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGuildOnboardingResponse(guildId string, prompts []OnboardingPromptResponse, defaultChannelIds []string, enabled bool) *UserGuildOnboardingResponse {
	this := UserGuildOnboardingResponse{}
	this.GuildId = guildId
	this.Prompts = prompts
	this.DefaultChannelIds = defaultChannelIds
	this.Enabled = enabled
	return &this
}

// NewUserGuildOnboardingResponseWithDefaults instantiates a new UserGuildOnboardingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGuildOnboardingResponseWithDefaults() *UserGuildOnboardingResponse {
	this := UserGuildOnboardingResponse{}
	return &this
}

// GetGuildId returns the GuildId field value
func (o *UserGuildOnboardingResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *UserGuildOnboardingResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetPrompts returns the Prompts field value
func (o *UserGuildOnboardingResponse) GetPrompts() []OnboardingPromptResponse {
	if o == nil {
		var ret []OnboardingPromptResponse
		return ret
	}

	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetPromptsOk() ([]OnboardingPromptResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompts, true
}

// SetPrompts sets field value
func (o *UserGuildOnboardingResponse) SetPrompts(v []OnboardingPromptResponse) {
	o.Prompts = v
}

// GetDefaultChannelIds returns the DefaultChannelIds field value
func (o *UserGuildOnboardingResponse) GetDefaultChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DefaultChannelIds
}

// GetDefaultChannelIdsOk returns a tuple with the DefaultChannelIds field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetDefaultChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultChannelIds, true
}

// SetDefaultChannelIds sets field value
func (o *UserGuildOnboardingResponse) SetDefaultChannelIds(v []string) {
	o.DefaultChannelIds = v
}

// GetEnabled returns the Enabled field value
func (o *UserGuildOnboardingResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UserGuildOnboardingResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UserGuildOnboardingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGuildOnboardingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["prompts"] = o.Prompts
	toSerialize["default_channel_ids"] = o.DefaultChannelIds
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *UserGuildOnboardingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guild_id",
		"prompts",
		"default_channel_ids",
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

	varUserGuildOnboardingResponse := _UserGuildOnboardingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGuildOnboardingResponse)

	if err != nil {
		return err
	}

	*o = UserGuildOnboardingResponse(varUserGuildOnboardingResponse)

	return err
}

type NullableUserGuildOnboardingResponse struct {
	value *UserGuildOnboardingResponse
	isSet bool
}

func (v NullableUserGuildOnboardingResponse) Get() *UserGuildOnboardingResponse {
	return v.value
}

func (v *NullableUserGuildOnboardingResponse) Set(val *UserGuildOnboardingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGuildOnboardingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGuildOnboardingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGuildOnboardingResponse(val *UserGuildOnboardingResponse) *NullableUserGuildOnboardingResponse {
	return &NullableUserGuildOnboardingResponse{value: val, isSet: true}
}

func (v NullableUserGuildOnboardingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGuildOnboardingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


