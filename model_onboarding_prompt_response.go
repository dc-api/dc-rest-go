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

// checks if the OnboardingPromptResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingPromptResponse{}

// OnboardingPromptResponse struct for OnboardingPromptResponse
type OnboardingPromptResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Title string `json:"title"`
	Options []OnboardingPromptOptionResponse `json:"options"`
	SingleSelect bool `json:"single_select"`
	Required bool `json:"required"`
	InOnboarding bool `json:"in_onboarding"`
	Type int32 `json:"type"`
}

type _OnboardingPromptResponse OnboardingPromptResponse

// NewOnboardingPromptResponse instantiates a new OnboardingPromptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingPromptResponse(id string, title string, options []OnboardingPromptOptionResponse, singleSelect bool, required bool, inOnboarding bool, type_ int32) *OnboardingPromptResponse {
	this := OnboardingPromptResponse{}
	this.Id = id
	this.Title = title
	this.Options = options
	this.SingleSelect = singleSelect
	this.Required = required
	this.InOnboarding = inOnboarding
	this.Type = type_
	return &this
}

// NewOnboardingPromptResponseWithDefaults instantiates a new OnboardingPromptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingPromptResponseWithDefaults() *OnboardingPromptResponse {
	this := OnboardingPromptResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OnboardingPromptResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OnboardingPromptResponse) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *OnboardingPromptResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OnboardingPromptResponse) SetTitle(v string) {
	o.Title = v
}

// GetOptions returns the Options field value
func (o *OnboardingPromptResponse) GetOptions() []OnboardingPromptOptionResponse {
	if o == nil {
		var ret []OnboardingPromptOptionResponse
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetOptionsOk() ([]OnboardingPromptOptionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *OnboardingPromptResponse) SetOptions(v []OnboardingPromptOptionResponse) {
	o.Options = v
}

// GetSingleSelect returns the SingleSelect field value
func (o *OnboardingPromptResponse) GetSingleSelect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SingleSelect
}

// GetSingleSelectOk returns a tuple with the SingleSelect field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetSingleSelectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleSelect, true
}

// SetSingleSelect sets field value
func (o *OnboardingPromptResponse) SetSingleSelect(v bool) {
	o.SingleSelect = v
}

// GetRequired returns the Required field value
func (o *OnboardingPromptResponse) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *OnboardingPromptResponse) SetRequired(v bool) {
	o.Required = v
}

// GetInOnboarding returns the InOnboarding field value
func (o *OnboardingPromptResponse) GetInOnboarding() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InOnboarding
}

// GetInOnboardingOk returns a tuple with the InOnboarding field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetInOnboardingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InOnboarding, true
}

// SetInOnboarding sets field value
func (o *OnboardingPromptResponse) SetInOnboarding(v bool) {
	o.InOnboarding = v
}

// GetType returns the Type field value
func (o *OnboardingPromptResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OnboardingPromptResponse) SetType(v int32) {
	o.Type = v
}

func (o OnboardingPromptResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingPromptResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["options"] = o.Options
	toSerialize["single_select"] = o.SingleSelect
	toSerialize["required"] = o.Required
	toSerialize["in_onboarding"] = o.InOnboarding
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *OnboardingPromptResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"options",
		"single_select",
		"required",
		"in_onboarding",
		"type",
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

	varOnboardingPromptResponse := _OnboardingPromptResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnboardingPromptResponse)

	if err != nil {
		return err
	}

	*o = OnboardingPromptResponse(varOnboardingPromptResponse)

	return err
}

type NullableOnboardingPromptResponse struct {
	value *OnboardingPromptResponse
	isSet bool
}

func (v NullableOnboardingPromptResponse) Get() *OnboardingPromptResponse {
	return v.value
}

func (v *NullableOnboardingPromptResponse) Set(val *OnboardingPromptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingPromptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingPromptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingPromptResponse(val *OnboardingPromptResponse) *NullableOnboardingPromptResponse {
	return &NullableOnboardingPromptResponse{value: val, isSet: true}
}

func (v NullableOnboardingPromptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingPromptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


