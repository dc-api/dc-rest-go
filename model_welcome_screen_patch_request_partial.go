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
)

// checks if the WelcomeScreenPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WelcomeScreenPatchRequestPartial{}

// WelcomeScreenPatchRequestPartial struct for WelcomeScreenPatchRequestPartial
type WelcomeScreenPatchRequestPartial struct {
	Description NullableString `json:"description,omitempty"`
	WelcomeChannels []GuildWelcomeChannel `json:"welcome_channels,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewWelcomeScreenPatchRequestPartial instantiates a new WelcomeScreenPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWelcomeScreenPatchRequestPartial() *WelcomeScreenPatchRequestPartial {
	this := WelcomeScreenPatchRequestPartial{}
	return &this
}

// NewWelcomeScreenPatchRequestPartialWithDefaults instantiates a new WelcomeScreenPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWelcomeScreenPatchRequestPartialWithDefaults() *WelcomeScreenPatchRequestPartial {
	this := WelcomeScreenPatchRequestPartial{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WelcomeScreenPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WelcomeScreenPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *WelcomeScreenPatchRequestPartial) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *WelcomeScreenPatchRequestPartial) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *WelcomeScreenPatchRequestPartial) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *WelcomeScreenPatchRequestPartial) UnsetDescription() {
	o.Description.Unset()
}

// GetWelcomeChannels returns the WelcomeChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WelcomeScreenPatchRequestPartial) GetWelcomeChannels() []GuildWelcomeChannel {
	if o == nil {
		var ret []GuildWelcomeChannel
		return ret
	}
	return o.WelcomeChannels
}

// GetWelcomeChannelsOk returns a tuple with the WelcomeChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WelcomeScreenPatchRequestPartial) GetWelcomeChannelsOk() ([]GuildWelcomeChannel, bool) {
	if o == nil {
		return nil, false
	}
	return o.WelcomeChannels, true
}

// HasWelcomeChannels returns a boolean if a field has been set.
func (o *WelcomeScreenPatchRequestPartial) HasWelcomeChannels() bool {
	if o != nil && !IsNil(o.WelcomeChannels) {
		return true
	}

	return false
}

// SetWelcomeChannels gets a reference to the given []GuildWelcomeChannel and assigns it to the WelcomeChannels field.
func (o *WelcomeScreenPatchRequestPartial) SetWelcomeChannels(v []GuildWelcomeChannel) {
	o.WelcomeChannels = v
}


// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WelcomeScreenPatchRequestPartial) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WelcomeScreenPatchRequestPartial) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled.Get()) {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *WelcomeScreenPatchRequestPartial) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *WelcomeScreenPatchRequestPartial) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *WelcomeScreenPatchRequestPartial) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *WelcomeScreenPatchRequestPartial) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o WelcomeScreenPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WelcomeScreenPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.WelcomeChannels != nil {
		toSerialize["welcome_channels"] = o.WelcomeChannels
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableWelcomeScreenPatchRequestPartial struct {
	value *WelcomeScreenPatchRequestPartial
	isSet bool
}

func (v NullableWelcomeScreenPatchRequestPartial) Get() *WelcomeScreenPatchRequestPartial {
	return v.value
}

func (v *NullableWelcomeScreenPatchRequestPartial) Set(val *WelcomeScreenPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableWelcomeScreenPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableWelcomeScreenPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWelcomeScreenPatchRequestPartial(val *WelcomeScreenPatchRequestPartial) *NullableWelcomeScreenPatchRequestPartial {
	return &NullableWelcomeScreenPatchRequestPartial{value: val, isSet: true}
}

func (v NullableWelcomeScreenPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWelcomeScreenPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


