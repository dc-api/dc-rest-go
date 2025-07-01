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

// checks if the GuildWelcomeScreenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildWelcomeScreenResponse{}

// GuildWelcomeScreenResponse struct for GuildWelcomeScreenResponse
type GuildWelcomeScreenResponse struct {
	Description NullableString `json:"description,omitempty"`
	WelcomeChannels []GuildWelcomeScreenChannelResponse `json:"welcome_channels"`
}

type _GuildWelcomeScreenResponse GuildWelcomeScreenResponse

// NewGuildWelcomeScreenResponse instantiates a new GuildWelcomeScreenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildWelcomeScreenResponse(welcomeChannels []GuildWelcomeScreenChannelResponse) *GuildWelcomeScreenResponse {
	this := GuildWelcomeScreenResponse{}
	this.WelcomeChannels = welcomeChannels
	return &this
}

// NewGuildWelcomeScreenResponseWithDefaults instantiates a new GuildWelcomeScreenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildWelcomeScreenResponseWithDefaults() *GuildWelcomeScreenResponse {
	this := GuildWelcomeScreenResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildWelcomeScreenResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildWelcomeScreenResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildWelcomeScreenResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildWelcomeScreenResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildWelcomeScreenResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildWelcomeScreenResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetWelcomeChannels returns the WelcomeChannels field value
func (o *GuildWelcomeScreenResponse) GetWelcomeChannels() []GuildWelcomeScreenChannelResponse {
	if o == nil {
		var ret []GuildWelcomeScreenChannelResponse
		return ret
	}

	return o.WelcomeChannels
}

// GetWelcomeChannelsOk returns a tuple with the WelcomeChannels field value
// and a boolean to check if the value has been set.
func (o *GuildWelcomeScreenResponse) GetWelcomeChannelsOk() ([]GuildWelcomeScreenChannelResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.WelcomeChannels, true
}

// SetWelcomeChannels sets field value
func (o *GuildWelcomeScreenResponse) SetWelcomeChannels(v []GuildWelcomeScreenChannelResponse) {
	o.WelcomeChannels = v
}

func (o GuildWelcomeScreenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildWelcomeScreenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["welcome_channels"] = o.WelcomeChannels
	return toSerialize, nil
}

func (o *GuildWelcomeScreenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"welcome_channels",
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

	varGuildWelcomeScreenResponse := _GuildWelcomeScreenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildWelcomeScreenResponse)

	if err != nil {
		return err
	}

	*o = GuildWelcomeScreenResponse(varGuildWelcomeScreenResponse)

	return err
}

type NullableGuildWelcomeScreenResponse struct {
	value *GuildWelcomeScreenResponse
	isSet bool
}

func (v NullableGuildWelcomeScreenResponse) Get() *GuildWelcomeScreenResponse {
	return v.value
}

func (v *NullableGuildWelcomeScreenResponse) Set(val *GuildWelcomeScreenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildWelcomeScreenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildWelcomeScreenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildWelcomeScreenResponse(val *GuildWelcomeScreenResponse) *NullableGuildWelcomeScreenResponse {
	return &NullableGuildWelcomeScreenResponse{value: val, isSet: true}
}

func (v NullableGuildWelcomeScreenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildWelcomeScreenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


