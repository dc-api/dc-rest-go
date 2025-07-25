/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the InviteGuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteGuildResponse{}

// InviteGuildResponse struct for InviteGuildResponse
type InviteGuildResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Splash NullableString `json:"splash,omitempty"`
	Banner NullableString `json:"banner,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Features []string `json:"features"`
	VerificationLevel *int32 `json:"verification_level,omitempty"`
	VanityUrlCode NullableString `json:"vanity_url_code,omitempty"`
	NsfwLevel NullableInt32 `json:"nsfw_level,omitempty"`
	Nsfw NullableBool `json:"nsfw,omitempty"`
	PremiumSubscriptionCount NullableInt32 `json:"premium_subscription_count,omitempty"`
}

type _InviteGuildResponse InviteGuildResponse

// NewInviteGuildResponse instantiates a new InviteGuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteGuildResponse(id string, name string, features []string) *InviteGuildResponse {
	this := InviteGuildResponse{}
	this.Id = id
	this.Name = name
	this.Features = features
	return &this
}

// NewInviteGuildResponseWithDefaults instantiates a new InviteGuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteGuildResponseWithDefaults() *InviteGuildResponse {
	this := InviteGuildResponse{}
	return &this
}

// GetId returns the Id field value
func (o *InviteGuildResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InviteGuildResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InviteGuildResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InviteGuildResponse) SetName(v string) {
	o.Name = v
}

// GetSplash returns the Splash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetSplash() string {
	if o == nil || IsNil(o.Splash.Get()) {
		var ret string
		return ret
	}
	return *o.Splash.Get()
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash.Get()) {
		return nil, false
	}
	return o.Splash.Get(), o.Splash.IsSet()
}

// HasSplash returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasSplash() bool {
	if o != nil && o.Splash.IsSet() {
		return true
	}

	return false
}

// SetSplash gets a reference to the given NullableString and assigns it to the Splash field.
func (o *InviteGuildResponse) SetSplash(v string) {
	o.Splash.Set(&v)
}

// SetSplashNil sets the value for Splash to be an explicit nil
func (o *InviteGuildResponse) SetSplashNil() {
	o.Splash.Set(nil)
}

// UnsetSplash ensures that no value is present for Splash, not even an explicit nil
func (o *InviteGuildResponse) UnsetSplash() {
	o.Splash.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *InviteGuildResponse) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *InviteGuildResponse) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *InviteGuildResponse) UnsetBanner() {
	o.Banner.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InviteGuildResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InviteGuildResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InviteGuildResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *InviteGuildResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *InviteGuildResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *InviteGuildResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetFeatures returns the Features field value
func (o *InviteGuildResponse) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *InviteGuildResponse) SetFeatures(v []string) {
	o.Features = v
}

// GetVerificationLevel returns the VerificationLevel field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetVerificationLevel() int32 {
	if o == nil || IsNil(o.VerificationLevel) {
		var ret int32
		return ret
	}
	return *o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetVerificationLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.VerificationLevel) {
		return nil, false
	}
	return o.VerificationLevel, true
}

// HasVerificationLevel returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasVerificationLevel() bool {
	if o != nil && !IsNil(o.VerificationLevel) {
		return true
	}

	return false
}

// SetVerificationLevel gets a reference to the given int32 and assigns it to the VerificationLevel field.
func (o *InviteGuildResponse) SetVerificationLevel(v int32) {
	o.VerificationLevel = &v
}


// GetVanityUrlCode returns the VanityUrlCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetVanityUrlCode() string {
	if o == nil || IsNil(o.VanityUrlCode.Get()) {
		var ret string
		return ret
	}
	return *o.VanityUrlCode.Get()
}

// GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetVanityUrlCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VanityUrlCode.Get()) {
		return nil, false
	}
	return o.VanityUrlCode.Get(), o.VanityUrlCode.IsSet()
}

// HasVanityUrlCode returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasVanityUrlCode() bool {
	if o != nil && o.VanityUrlCode.IsSet() {
		return true
	}

	return false
}

// SetVanityUrlCode gets a reference to the given NullableString and assigns it to the VanityUrlCode field.
func (o *InviteGuildResponse) SetVanityUrlCode(v string) {
	o.VanityUrlCode.Set(&v)
}

// SetVanityUrlCodeNil sets the value for VanityUrlCode to be an explicit nil
func (o *InviteGuildResponse) SetVanityUrlCodeNil() {
	o.VanityUrlCode.Set(nil)
}

// UnsetVanityUrlCode ensures that no value is present for VanityUrlCode, not even an explicit nil
func (o *InviteGuildResponse) UnsetVanityUrlCode() {
	o.VanityUrlCode.Unset()
}

// GetNsfwLevel returns the NsfwLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetNsfwLevel() int32 {
	if o == nil || IsNil(o.NsfwLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.NsfwLevel.Get()
}

// GetNsfwLevelOk returns a tuple with the NsfwLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetNsfwLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.NsfwLevel.Get()) {
		return nil, false
	}
	return o.NsfwLevel.Get(), o.NsfwLevel.IsSet()
}

// HasNsfwLevel returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasNsfwLevel() bool {
	if o != nil && o.NsfwLevel.IsSet() {
		return true
	}

	return false
}

// SetNsfwLevel gets a reference to the given NullableInt32 and assigns it to the NsfwLevel field.
func (o *InviteGuildResponse) SetNsfwLevel(v int32) {
	o.NsfwLevel.Set(&v)
}

// SetNsfwLevelNil sets the value for NsfwLevel to be an explicit nil
func (o *InviteGuildResponse) SetNsfwLevelNil() {
	o.NsfwLevel.Set(nil)
}

// UnsetNsfwLevel ensures that no value is present for NsfwLevel, not even an explicit nil
func (o *InviteGuildResponse) UnsetNsfwLevel() {
	o.NsfwLevel.Unset()
}

// GetNsfw returns the Nsfw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw.Get()) {
		var ret bool
		return ret
	}
	return *o.Nsfw.Get()
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw.Get()) {
		return nil, false
	}
	return o.Nsfw.Get(), o.Nsfw.IsSet()
}

// HasNsfw returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasNsfw() bool {
	if o != nil && o.Nsfw.IsSet() {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given NullableBool and assigns it to the Nsfw field.
func (o *InviteGuildResponse) SetNsfw(v bool) {
	o.Nsfw.Set(&v)
}

// SetNsfwNil sets the value for Nsfw to be an explicit nil
func (o *InviteGuildResponse) SetNsfwNil() {
	o.Nsfw.Set(nil)
}

// UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
func (o *InviteGuildResponse) UnsetNsfw() {
	o.Nsfw.Unset()
}

// GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteGuildResponse) GetPremiumSubscriptionCount() int32 {
	if o == nil || IsNil(o.PremiumSubscriptionCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PremiumSubscriptionCount.Get()
}

// GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteGuildResponse) GetPremiumSubscriptionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PremiumSubscriptionCount.Get()) {
		return nil, false
	}
	return o.PremiumSubscriptionCount.Get(), o.PremiumSubscriptionCount.IsSet()
}

// HasPremiumSubscriptionCount returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasPremiumSubscriptionCount() bool {
	if o != nil && o.PremiumSubscriptionCount.IsSet() {
		return true
	}

	return false
}

// SetPremiumSubscriptionCount gets a reference to the given NullableInt32 and assigns it to the PremiumSubscriptionCount field.
func (o *InviteGuildResponse) SetPremiumSubscriptionCount(v int32) {
	o.PremiumSubscriptionCount.Set(&v)
}

// SetPremiumSubscriptionCountNil sets the value for PremiumSubscriptionCount to be an explicit nil
func (o *InviteGuildResponse) SetPremiumSubscriptionCountNil() {
	o.PremiumSubscriptionCount.Set(nil)
}

// UnsetPremiumSubscriptionCount ensures that no value is present for PremiumSubscriptionCount, not even an explicit nil
func (o *InviteGuildResponse) UnsetPremiumSubscriptionCount() {
	o.PremiumSubscriptionCount.Unset()
}

func (o InviteGuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteGuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Splash.IsSet() {
		toSerialize["splash"] = o.Splash.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["features"] = o.Features
	if !IsNil(o.VerificationLevel) {
		toSerialize["verification_level"] = o.VerificationLevel
	}
	if o.VanityUrlCode.IsSet() {
		toSerialize["vanity_url_code"] = o.VanityUrlCode.Get()
	}
	if o.NsfwLevel.IsSet() {
		toSerialize["nsfw_level"] = o.NsfwLevel.Get()
	}
	if o.Nsfw.IsSet() {
		toSerialize["nsfw"] = o.Nsfw.Get()
	}
	if o.PremiumSubscriptionCount.IsSet() {
		toSerialize["premium_subscription_count"] = o.PremiumSubscriptionCount.Get()
	}
	return toSerialize, nil
}

func (o *InviteGuildResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"features",
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

	varInviteGuildResponse := _InviteGuildResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInviteGuildResponse)

	if err != nil {
		return err
	}

	*o = InviteGuildResponse(varInviteGuildResponse)

	return err
}

type NullableInviteGuildResponse struct {
	value *InviteGuildResponse
	isSet bool
}

func (v NullableInviteGuildResponse) Get() *InviteGuildResponse {
	return v.value
}

func (v *NullableInviteGuildResponse) Set(val *InviteGuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteGuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteGuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteGuildResponse(val *InviteGuildResponse) *NullableInviteGuildResponse {
	return &NullableInviteGuildResponse{value: val, isSet: true}
}

func (v NullableInviteGuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteGuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


