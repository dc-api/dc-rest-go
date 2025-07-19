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

// checks if the OnboardingPromptOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingPromptOptionRequest{}

// OnboardingPromptOptionRequest struct for OnboardingPromptOptionRequest
type OnboardingPromptOptionRequest struct {
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Title string `json:"title"`
	Description NullableString `json:"description,omitempty"`
	EmojiId *string `json:"emoji_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EmojiName NullableString `json:"emoji_name,omitempty"`
	EmojiAnimated NullableBool `json:"emoji_animated,omitempty"`
	RoleIds []string `json:"role_ids,omitempty"`
	ChannelIds []string `json:"channel_ids,omitempty"`
}

type _OnboardingPromptOptionRequest OnboardingPromptOptionRequest

// NewOnboardingPromptOptionRequest instantiates a new OnboardingPromptOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingPromptOptionRequest(title string) *OnboardingPromptOptionRequest {
	this := OnboardingPromptOptionRequest{}
	this.Title = title
	return &this
}

// NewOnboardingPromptOptionRequestWithDefaults instantiates a new OnboardingPromptOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingPromptOptionRequestWithDefaults() *OnboardingPromptOptionRequest {
	this := OnboardingPromptOptionRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OnboardingPromptOptionRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OnboardingPromptOptionRequest) SetId(v string) {
	o.Id = &v
}


// GetTitle returns the Title field value
func (o *OnboardingPromptOptionRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OnboardingPromptOptionRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnboardingPromptOptionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnboardingPromptOptionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OnboardingPromptOptionRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OnboardingPromptOptionRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OnboardingPromptOptionRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetEmojiId returns the EmojiId field value if set, zero value otherwise.
func (o *OnboardingPromptOptionRequest) GetEmojiId() string {
	if o == nil || IsNil(o.EmojiId) {
		var ret string
		return ret
	}
	return *o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionRequest) GetEmojiIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiId) {
		return nil, false
	}
	return o.EmojiId, true
}

// HasEmojiId returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasEmojiId() bool {
	if o != nil && !IsNil(o.EmojiId) {
		return true
	}

	return false
}

// SetEmojiId gets a reference to the given string and assigns it to the EmojiId field.
func (o *OnboardingPromptOptionRequest) SetEmojiId(v string) {
	o.EmojiId = &v
}


// GetEmojiName returns the EmojiName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnboardingPromptOptionRequest) GetEmojiName() string {
	if o == nil || IsNil(o.EmojiName.Get()) {
		var ret string
		return ret
	}
	return *o.EmojiName.Get()
}

// GetEmojiNameOk returns a tuple with the EmojiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnboardingPromptOptionRequest) GetEmojiNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmojiName.Get()) {
		return nil, false
	}
	return o.EmojiName.Get(), o.EmojiName.IsSet()
}

// HasEmojiName returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasEmojiName() bool {
	if o != nil && o.EmojiName.IsSet() {
		return true
	}

	return false
}

// SetEmojiName gets a reference to the given NullableString and assigns it to the EmojiName field.
func (o *OnboardingPromptOptionRequest) SetEmojiName(v string) {
	o.EmojiName.Set(&v)
}

// SetEmojiNameNil sets the value for EmojiName to be an explicit nil
func (o *OnboardingPromptOptionRequest) SetEmojiNameNil() {
	o.EmojiName.Set(nil)
}

// UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
func (o *OnboardingPromptOptionRequest) UnsetEmojiName() {
	o.EmojiName.Unset()
}

// GetEmojiAnimated returns the EmojiAnimated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnboardingPromptOptionRequest) GetEmojiAnimated() bool {
	if o == nil || IsNil(o.EmojiAnimated.Get()) {
		var ret bool
		return ret
	}
	return *o.EmojiAnimated.Get()
}

// GetEmojiAnimatedOk returns a tuple with the EmojiAnimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnboardingPromptOptionRequest) GetEmojiAnimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmojiAnimated.Get()) {
		return nil, false
	}
	return o.EmojiAnimated.Get(), o.EmojiAnimated.IsSet()
}

// HasEmojiAnimated returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasEmojiAnimated() bool {
	if o != nil && o.EmojiAnimated.IsSet() {
		return true
	}

	return false
}

// SetEmojiAnimated gets a reference to the given NullableBool and assigns it to the EmojiAnimated field.
func (o *OnboardingPromptOptionRequest) SetEmojiAnimated(v bool) {
	o.EmojiAnimated.Set(&v)
}

// SetEmojiAnimatedNil sets the value for EmojiAnimated to be an explicit nil
func (o *OnboardingPromptOptionRequest) SetEmojiAnimatedNil() {
	o.EmojiAnimated.Set(nil)
}

// UnsetEmojiAnimated ensures that no value is present for EmojiAnimated, not even an explicit nil
func (o *OnboardingPromptOptionRequest) UnsetEmojiAnimated() {
	o.EmojiAnimated.Unset()
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnboardingPromptOptionRequest) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnboardingPromptOptionRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *OnboardingPromptOptionRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}


// GetChannelIds returns the ChannelIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnboardingPromptOptionRequest) GetChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ChannelIds
}

// GetChannelIdsOk returns a tuple with the ChannelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnboardingPromptOptionRequest) GetChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelIds, true
}

// HasChannelIds returns a boolean if a field has been set.
func (o *OnboardingPromptOptionRequest) HasChannelIds() bool {
	if o != nil && !IsNil(o.ChannelIds) {
		return true
	}

	return false
}

// SetChannelIds gets a reference to the given []string and assigns it to the ChannelIds field.
func (o *OnboardingPromptOptionRequest) SetChannelIds(v []string) {
	o.ChannelIds = v
}


func (o OnboardingPromptOptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingPromptOptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["title"] = o.Title
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.EmojiId) {
		toSerialize["emoji_id"] = o.EmojiId
	}
	if o.EmojiName.IsSet() {
		toSerialize["emoji_name"] = o.EmojiName.Get()
	}
	if o.EmojiAnimated.IsSet() {
		toSerialize["emoji_animated"] = o.EmojiAnimated.Get()
	}
	if o.RoleIds != nil {
		toSerialize["role_ids"] = o.RoleIds
	}
	if o.ChannelIds != nil {
		toSerialize["channel_ids"] = o.ChannelIds
	}
	return toSerialize, nil
}

func (o *OnboardingPromptOptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varOnboardingPromptOptionRequest := _OnboardingPromptOptionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnboardingPromptOptionRequest)

	if err != nil {
		return err
	}

	*o = OnboardingPromptOptionRequest(varOnboardingPromptOptionRequest)

	return err
}

type NullableOnboardingPromptOptionRequest struct {
	value *OnboardingPromptOptionRequest
	isSet bool
}

func (v NullableOnboardingPromptOptionRequest) Get() *OnboardingPromptOptionRequest {
	return v.value
}

func (v *NullableOnboardingPromptOptionRequest) Set(val *OnboardingPromptOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingPromptOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingPromptOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingPromptOptionRequest(val *OnboardingPromptOptionRequest) *NullableOnboardingPromptOptionRequest {
	return &NullableOnboardingPromptOptionRequest{value: val, isSet: true}
}

func (v NullableOnboardingPromptOptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingPromptOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


