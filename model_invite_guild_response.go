/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
	Splash *string `json:"splash,omitempty"`
	Banner *string `json:"banner,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Features []GuildFeatures `json:"features"`
	VerificationLevel *GuildPatchRequestPartialVerificationLevel `json:"verification_level,omitempty"`
	VanityUrlCode *string `json:"vanity_url_code,omitempty"`
	NsfwLevel *InviteGuildResponseNsfwLevel `json:"nsfw_level,omitempty"`
	Nsfw *bool `json:"nsfw,omitempty"`
	PremiumSubscriptionCount int32 `json:"premium_subscription_count"`
}

type _InviteGuildResponse InviteGuildResponse

// NewInviteGuildResponse instantiates a new InviteGuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteGuildResponse(id string, name string, features []GuildFeatures, premiumSubscriptionCount int32) *InviteGuildResponse {
	this := InviteGuildResponse{}
	this.Id = id
	this.Name = name
	this.Features = features
	this.PremiumSubscriptionCount = premiumSubscriptionCount
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

// GetSplash returns the Splash field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetSplash() string {
	if o == nil || IsNil(o.Splash) {
		var ret string
		return ret
	}
	return *o.Splash
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash) {
		return nil, false
	}
	return o.Splash, true
}

// HasSplash returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasSplash() bool {
	if o != nil && !IsNil(o.Splash) {
		return true
	}

	return false
}

// SetSplash gets a reference to the given string and assigns it to the Splash field.
func (o *InviteGuildResponse) SetSplash(v string) {
	o.Splash = &v
}


// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *InviteGuildResponse) SetBanner(v string) {
	o.Banner = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InviteGuildResponse) SetDescription(v string) {
	o.Description = &v
}


// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *InviteGuildResponse) SetIcon(v string) {
	o.Icon = &v
}


// GetFeatures returns the Features field value
func (o *InviteGuildResponse) GetFeatures() []GuildFeatures {
	if o == nil {
		var ret []GuildFeatures
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetFeaturesOk() ([]GuildFeatures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *InviteGuildResponse) SetFeatures(v []GuildFeatures) {
	o.Features = v
}

// GetVerificationLevel returns the VerificationLevel field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetVerificationLevel() GuildPatchRequestPartialVerificationLevel {
	if o == nil || IsNil(o.VerificationLevel) {
		var ret GuildPatchRequestPartialVerificationLevel
		return ret
	}
	return *o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetVerificationLevelOk() (*GuildPatchRequestPartialVerificationLevel, bool) {
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

// SetVerificationLevel gets a reference to the given GuildPatchRequestPartialVerificationLevel and assigns it to the VerificationLevel field.
func (o *InviteGuildResponse) SetVerificationLevel(v GuildPatchRequestPartialVerificationLevel) {
	o.VerificationLevel = &v
}


// GetVanityUrlCode returns the VanityUrlCode field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetVanityUrlCode() string {
	if o == nil || IsNil(o.VanityUrlCode) {
		var ret string
		return ret
	}
	return *o.VanityUrlCode
}

// GetVanityUrlCodeOk returns a tuple with the VanityUrlCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetVanityUrlCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VanityUrlCode) {
		return nil, false
	}
	return o.VanityUrlCode, true
}

// HasVanityUrlCode returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasVanityUrlCode() bool {
	if o != nil && !IsNil(o.VanityUrlCode) {
		return true
	}

	return false
}

// SetVanityUrlCode gets a reference to the given string and assigns it to the VanityUrlCode field.
func (o *InviteGuildResponse) SetVanityUrlCode(v string) {
	o.VanityUrlCode = &v
}


// GetNsfwLevel returns the NsfwLevel field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetNsfwLevel() InviteGuildResponseNsfwLevel {
	if o == nil || IsNil(o.NsfwLevel) {
		var ret InviteGuildResponseNsfwLevel
		return ret
	}
	return *o.NsfwLevel
}

// GetNsfwLevelOk returns a tuple with the NsfwLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetNsfwLevelOk() (*InviteGuildResponseNsfwLevel, bool) {
	if o == nil || IsNil(o.NsfwLevel) {
		return nil, false
	}
	return o.NsfwLevel, true
}

// HasNsfwLevel returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasNsfwLevel() bool {
	if o != nil && !IsNil(o.NsfwLevel) {
		return true
	}

	return false
}

// SetNsfwLevel gets a reference to the given InviteGuildResponseNsfwLevel and assigns it to the NsfwLevel field.
func (o *InviteGuildResponse) SetNsfwLevel(v InviteGuildResponseNsfwLevel) {
	o.NsfwLevel = &v
}


// GetNsfw returns the Nsfw field value if set, zero value otherwise.
func (o *InviteGuildResponse) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw) {
		var ret bool
		return ret
	}
	return *o.Nsfw
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw) {
		return nil, false
	}
	return o.Nsfw, true
}

// HasNsfw returns a boolean if a field has been set.
func (o *InviteGuildResponse) HasNsfw() bool {
	if o != nil && !IsNil(o.Nsfw) {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given bool and assigns it to the Nsfw field.
func (o *InviteGuildResponse) SetNsfw(v bool) {
	o.Nsfw = &v
}


// GetPremiumSubscriptionCount returns the PremiumSubscriptionCount field value
func (o *InviteGuildResponse) GetPremiumSubscriptionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PremiumSubscriptionCount
}

// GetPremiumSubscriptionCountOk returns a tuple with the PremiumSubscriptionCount field value
// and a boolean to check if the value has been set.
func (o *InviteGuildResponse) GetPremiumSubscriptionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PremiumSubscriptionCount, true
}

// SetPremiumSubscriptionCount sets field value
func (o *InviteGuildResponse) SetPremiumSubscriptionCount(v int32) {
	o.PremiumSubscriptionCount = v
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
	if !IsNil(o.Splash) {
		toSerialize["splash"] = o.Splash
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["features"] = o.Features
	if !IsNil(o.VerificationLevel) {
		toSerialize["verification_level"] = o.VerificationLevel
	}
	if !IsNil(o.VanityUrlCode) {
		toSerialize["vanity_url_code"] = o.VanityUrlCode
	}
	if !IsNil(o.NsfwLevel) {
		toSerialize["nsfw_level"] = o.NsfwLevel
	}
	if !IsNil(o.Nsfw) {
		toSerialize["nsfw"] = o.Nsfw
	}
	toSerialize["premium_subscription_count"] = o.PremiumSubscriptionCount
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
		"premium_subscription_count",
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

