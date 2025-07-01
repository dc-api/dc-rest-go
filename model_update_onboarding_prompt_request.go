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

// checks if the UpdateOnboardingPromptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOnboardingPromptRequest{}

// UpdateOnboardingPromptRequest struct for UpdateOnboardingPromptRequest
type UpdateOnboardingPromptRequest struct {
	Title string `json:"title"`
	Options []OnboardingPromptOptionRequest `json:"options"`
	SingleSelect NullableBool `json:"single_select,omitempty"`
	Required NullableBool `json:"required,omitempty"`
	InOnboarding NullableBool `json:"in_onboarding,omitempty"`
	Type *int32 `json:"type,omitempty"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _UpdateOnboardingPromptRequest UpdateOnboardingPromptRequest

// NewUpdateOnboardingPromptRequest instantiates a new UpdateOnboardingPromptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOnboardingPromptRequest(title string, options []OnboardingPromptOptionRequest, id string) *UpdateOnboardingPromptRequest {
	this := UpdateOnboardingPromptRequest{}
	this.Title = title
	this.Options = options
	this.Id = id
	return &this
}

// NewUpdateOnboardingPromptRequestWithDefaults instantiates a new UpdateOnboardingPromptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOnboardingPromptRequestWithDefaults() *UpdateOnboardingPromptRequest {
	this := UpdateOnboardingPromptRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *UpdateOnboardingPromptRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UpdateOnboardingPromptRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UpdateOnboardingPromptRequest) SetTitle(v string) {
	o.Title = v
}

// GetOptions returns the Options field value
func (o *UpdateOnboardingPromptRequest) GetOptions() []OnboardingPromptOptionRequest {
	if o == nil {
		var ret []OnboardingPromptOptionRequest
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UpdateOnboardingPromptRequest) GetOptionsOk() ([]OnboardingPromptOptionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UpdateOnboardingPromptRequest) SetOptions(v []OnboardingPromptOptionRequest) {
	o.Options = v
}

// GetSingleSelect returns the SingleSelect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOnboardingPromptRequest) GetSingleSelect() bool {
	if o == nil || IsNil(o.SingleSelect.Get()) {
		var ret bool
		return ret
	}
	return *o.SingleSelect.Get()
}

// GetSingleSelectOk returns a tuple with the SingleSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOnboardingPromptRequest) GetSingleSelectOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleSelect.Get()) {
		return nil, false
	}
	return o.SingleSelect.Get(), o.SingleSelect.IsSet()
}

// HasSingleSelect returns a boolean if a field has been set.
func (o *UpdateOnboardingPromptRequest) HasSingleSelect() bool {
	if o != nil && o.SingleSelect.IsSet() {
		return true
	}

	return false
}

// SetSingleSelect gets a reference to the given NullableBool and assigns it to the SingleSelect field.
func (o *UpdateOnboardingPromptRequest) SetSingleSelect(v bool) {
	o.SingleSelect.Set(&v)
}

// SetSingleSelectNil sets the value for SingleSelect to be an explicit nil
func (o *UpdateOnboardingPromptRequest) SetSingleSelectNil() {
	o.SingleSelect.Set(nil)
}

// UnsetSingleSelect ensures that no value is present for SingleSelect, not even an explicit nil
func (o *UpdateOnboardingPromptRequest) UnsetSingleSelect() {
	o.SingleSelect.Unset()
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOnboardingPromptRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOnboardingPromptRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required.Get()) {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *UpdateOnboardingPromptRequest) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *UpdateOnboardingPromptRequest) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *UpdateOnboardingPromptRequest) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *UpdateOnboardingPromptRequest) UnsetRequired() {
	o.Required.Unset()
}

// GetInOnboarding returns the InOnboarding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOnboardingPromptRequest) GetInOnboarding() bool {
	if o == nil || IsNil(o.InOnboarding.Get()) {
		var ret bool
		return ret
	}
	return *o.InOnboarding.Get()
}

// GetInOnboardingOk returns a tuple with the InOnboarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOnboardingPromptRequest) GetInOnboardingOk() (*bool, bool) {
	if o == nil || IsNil(o.InOnboarding.Get()) {
		return nil, false
	}
	return o.InOnboarding.Get(), o.InOnboarding.IsSet()
}

// HasInOnboarding returns a boolean if a field has been set.
func (o *UpdateOnboardingPromptRequest) HasInOnboarding() bool {
	if o != nil && o.InOnboarding.IsSet() {
		return true
	}

	return false
}

// SetInOnboarding gets a reference to the given NullableBool and assigns it to the InOnboarding field.
func (o *UpdateOnboardingPromptRequest) SetInOnboarding(v bool) {
	o.InOnboarding.Set(&v)
}

// SetInOnboardingNil sets the value for InOnboarding to be an explicit nil
func (o *UpdateOnboardingPromptRequest) SetInOnboardingNil() {
	o.InOnboarding.Set(nil)
}

// UnsetInOnboarding ensures that no value is present for InOnboarding, not even an explicit nil
func (o *UpdateOnboardingPromptRequest) UnsetInOnboarding() {
	o.InOnboarding.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateOnboardingPromptRequest) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOnboardingPromptRequest) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateOnboardingPromptRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UpdateOnboardingPromptRequest) SetType(v int32) {
	o.Type = &v
}


// GetId returns the Id field value
func (o *UpdateOnboardingPromptRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateOnboardingPromptRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateOnboardingPromptRequest) SetId(v string) {
	o.Id = v
}

func (o UpdateOnboardingPromptRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOnboardingPromptRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["options"] = o.Options
	if o.SingleSelect.IsSet() {
		toSerialize["single_select"] = o.SingleSelect.Get()
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if o.InOnboarding.IsSet() {
		toSerialize["in_onboarding"] = o.InOnboarding.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UpdateOnboardingPromptRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"options",
		"id",
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

	varUpdateOnboardingPromptRequest := _UpdateOnboardingPromptRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOnboardingPromptRequest)

	if err != nil {
		return err
	}

	*o = UpdateOnboardingPromptRequest(varUpdateOnboardingPromptRequest)

	return err
}

type NullableUpdateOnboardingPromptRequest struct {
	value *UpdateOnboardingPromptRequest
	isSet bool
}

func (v NullableUpdateOnboardingPromptRequest) Get() *UpdateOnboardingPromptRequest {
	return v.value
}

func (v *NullableUpdateOnboardingPromptRequest) Set(val *UpdateOnboardingPromptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOnboardingPromptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOnboardingPromptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOnboardingPromptRequest(val *UpdateOnboardingPromptRequest) *NullableUpdateOnboardingPromptRequest {
	return &NullableUpdateOnboardingPromptRequest{value: val, isSet: true}
}

func (v NullableUpdateOnboardingPromptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOnboardingPromptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


