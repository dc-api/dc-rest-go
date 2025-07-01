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
)

// checks if the UpdateGuildOnboardingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGuildOnboardingRequest{}

// UpdateGuildOnboardingRequest struct for UpdateGuildOnboardingRequest
type UpdateGuildOnboardingRequest struct {
	Prompts []UpdateOnboardingPromptRequest `json:"prompts,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	DefaultChannelIds []string `json:"default_channel_ids,omitempty"`
	Mode *int32 `json:"mode,omitempty"`
}

// NewUpdateGuildOnboardingRequest instantiates a new UpdateGuildOnboardingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGuildOnboardingRequest() *UpdateGuildOnboardingRequest {
	this := UpdateGuildOnboardingRequest{}
	return &this
}

// NewUpdateGuildOnboardingRequestWithDefaults instantiates a new UpdateGuildOnboardingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGuildOnboardingRequestWithDefaults() *UpdateGuildOnboardingRequest {
	this := UpdateGuildOnboardingRequest{}
	return &this
}

// GetPrompts returns the Prompts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildOnboardingRequest) GetPrompts() []UpdateOnboardingPromptRequest {
	if o == nil {
		var ret []UpdateOnboardingPromptRequest
		return ret
	}
	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildOnboardingRequest) GetPromptsOk() ([]UpdateOnboardingPromptRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *UpdateGuildOnboardingRequest) HasPrompts() bool {
	if o != nil && !IsNil(o.Prompts) {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []UpdateOnboardingPromptRequest and assigns it to the Prompts field.
func (o *UpdateGuildOnboardingRequest) SetPrompts(v []UpdateOnboardingPromptRequest) {
	o.Prompts = v
}


// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildOnboardingRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildOnboardingRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled.Get()) {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateGuildOnboardingRequest) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *UpdateGuildOnboardingRequest) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *UpdateGuildOnboardingRequest) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *UpdateGuildOnboardingRequest) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetDefaultChannelIds returns the DefaultChannelIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildOnboardingRequest) GetDefaultChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultChannelIds
}

// GetDefaultChannelIdsOk returns a tuple with the DefaultChannelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildOnboardingRequest) GetDefaultChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultChannelIds, true
}

// HasDefaultChannelIds returns a boolean if a field has been set.
func (o *UpdateGuildOnboardingRequest) HasDefaultChannelIds() bool {
	if o != nil && !IsNil(o.DefaultChannelIds) {
		return true
	}

	return false
}

// SetDefaultChannelIds gets a reference to the given []string and assigns it to the DefaultChannelIds field.
func (o *UpdateGuildOnboardingRequest) SetDefaultChannelIds(v []string) {
	o.DefaultChannelIds = v
}


// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UpdateGuildOnboardingRequest) GetMode() int32 {
	if o == nil || IsNil(o.Mode) {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildOnboardingRequest) GetModeOk() (*int32, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UpdateGuildOnboardingRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *UpdateGuildOnboardingRequest) SetMode(v int32) {
	o.Mode = &v
}


func (o UpdateGuildOnboardingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGuildOnboardingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Prompts != nil {
		toSerialize["prompts"] = o.Prompts
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.DefaultChannelIds != nil {
		toSerialize["default_channel_ids"] = o.DefaultChannelIds
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableUpdateGuildOnboardingRequest struct {
	value *UpdateGuildOnboardingRequest
	isSet bool
}

func (v NullableUpdateGuildOnboardingRequest) Get() *UpdateGuildOnboardingRequest {
	return v.value
}

func (v *NullableUpdateGuildOnboardingRequest) Set(val *UpdateGuildOnboardingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildOnboardingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildOnboardingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildOnboardingRequest(val *UpdateGuildOnboardingRequest) *NullableUpdateGuildOnboardingRequest {
	return &NullableUpdateGuildOnboardingRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildOnboardingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildOnboardingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


