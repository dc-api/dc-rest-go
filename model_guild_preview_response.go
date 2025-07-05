/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// checks if the GuildPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildPreviewResponse{}

// GuildPreviewResponse struct for GuildPreviewResponse
type GuildPreviewResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
	Description NullableString `json:"description,omitempty"`
	HomeHeader NullableString `json:"home_header,omitempty"`
	Splash NullableString `json:"splash,omitempty"`
	DiscoverySplash NullableString `json:"discovery_splash,omitempty"`
	Features []string `json:"features"`
	ApproximateMemberCount int32 `json:"approximate_member_count"`
	ApproximatePresenceCount int32 `json:"approximate_presence_count"`
	Emojis []EmojiResponse `json:"emojis"`
	Stickers []GuildStickerResponse `json:"stickers"`
}

type _GuildPreviewResponse GuildPreviewResponse

// NewGuildPreviewResponse instantiates a new GuildPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildPreviewResponse(id string, name string, features []string, approximateMemberCount int32, approximatePresenceCount int32, emojis []EmojiResponse, stickers []GuildStickerResponse) *GuildPreviewResponse {
	this := GuildPreviewResponse{}
	this.Id = id
	this.Name = name
	this.Features = features
	this.ApproximateMemberCount = approximateMemberCount
	this.ApproximatePresenceCount = approximatePresenceCount
	this.Emojis = emojis
	this.Stickers = stickers
	return &this
}

// NewGuildPreviewResponseWithDefaults instantiates a new GuildPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildPreviewResponseWithDefaults() *GuildPreviewResponse {
	this := GuildPreviewResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildPreviewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildPreviewResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GuildPreviewResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildPreviewResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPreviewResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPreviewResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildPreviewResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GuildPreviewResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *GuildPreviewResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GuildPreviewResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPreviewResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPreviewResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildPreviewResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildPreviewResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildPreviewResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildPreviewResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetHomeHeader returns the HomeHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPreviewResponse) GetHomeHeader() string {
	if o == nil || IsNil(o.HomeHeader.Get()) {
		var ret string
		return ret
	}
	return *o.HomeHeader.Get()
}

// GetHomeHeaderOk returns a tuple with the HomeHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPreviewResponse) GetHomeHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HomeHeader.Get()) {
		return nil, false
	}
	return o.HomeHeader.Get(), o.HomeHeader.IsSet()
}

// HasHomeHeader returns a boolean if a field has been set.
func (o *GuildPreviewResponse) HasHomeHeader() bool {
	if o != nil && o.HomeHeader.IsSet() {
		return true
	}

	return false
}

// SetHomeHeader gets a reference to the given NullableString and assigns it to the HomeHeader field.
func (o *GuildPreviewResponse) SetHomeHeader(v string) {
	o.HomeHeader.Set(&v)
}

// SetHomeHeaderNil sets the value for HomeHeader to be an explicit nil
func (o *GuildPreviewResponse) SetHomeHeaderNil() {
	o.HomeHeader.Set(nil)
}

// UnsetHomeHeader ensures that no value is present for HomeHeader, not even an explicit nil
func (o *GuildPreviewResponse) UnsetHomeHeader() {
	o.HomeHeader.Unset()
}

// GetSplash returns the Splash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPreviewResponse) GetSplash() string {
	if o == nil || IsNil(o.Splash.Get()) {
		var ret string
		return ret
	}
	return *o.Splash.Get()
}

// GetSplashOk returns a tuple with the Splash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPreviewResponse) GetSplashOk() (*string, bool) {
	if o == nil || IsNil(o.Splash.Get()) {
		return nil, false
	}
	return o.Splash.Get(), o.Splash.IsSet()
}

// HasSplash returns a boolean if a field has been set.
func (o *GuildPreviewResponse) HasSplash() bool {
	if o != nil && o.Splash.IsSet() {
		return true
	}

	return false
}

// SetSplash gets a reference to the given NullableString and assigns it to the Splash field.
func (o *GuildPreviewResponse) SetSplash(v string) {
	o.Splash.Set(&v)
}

// SetSplashNil sets the value for Splash to be an explicit nil
func (o *GuildPreviewResponse) SetSplashNil() {
	o.Splash.Set(nil)
}

// UnsetSplash ensures that no value is present for Splash, not even an explicit nil
func (o *GuildPreviewResponse) UnsetSplash() {
	o.Splash.Unset()
}

// GetDiscoverySplash returns the DiscoverySplash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildPreviewResponse) GetDiscoverySplash() string {
	if o == nil || IsNil(o.DiscoverySplash.Get()) {
		var ret string
		return ret
	}
	return *o.DiscoverySplash.Get()
}

// GetDiscoverySplashOk returns a tuple with the DiscoverySplash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildPreviewResponse) GetDiscoverySplashOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoverySplash.Get()) {
		return nil, false
	}
	return o.DiscoverySplash.Get(), o.DiscoverySplash.IsSet()
}

// HasDiscoverySplash returns a boolean if a field has been set.
func (o *GuildPreviewResponse) HasDiscoverySplash() bool {
	if o != nil && o.DiscoverySplash.IsSet() {
		return true
	}

	return false
}

// SetDiscoverySplash gets a reference to the given NullableString and assigns it to the DiscoverySplash field.
func (o *GuildPreviewResponse) SetDiscoverySplash(v string) {
	o.DiscoverySplash.Set(&v)
}

// SetDiscoverySplashNil sets the value for DiscoverySplash to be an explicit nil
func (o *GuildPreviewResponse) SetDiscoverySplashNil() {
	o.DiscoverySplash.Set(nil)
}

// UnsetDiscoverySplash ensures that no value is present for DiscoverySplash, not even an explicit nil
func (o *GuildPreviewResponse) UnsetDiscoverySplash() {
	o.DiscoverySplash.Unset()
}

// GetFeatures returns the Features field value
func (o *GuildPreviewResponse) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *GuildPreviewResponse) SetFeatures(v []string) {
	o.Features = v
}

// GetApproximateMemberCount returns the ApproximateMemberCount field value
func (o *GuildPreviewResponse) GetApproximateMemberCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproximateMemberCount
}

// GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetApproximateMemberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproximateMemberCount, true
}

// SetApproximateMemberCount sets field value
func (o *GuildPreviewResponse) SetApproximateMemberCount(v int32) {
	o.ApproximateMemberCount = v
}

// GetApproximatePresenceCount returns the ApproximatePresenceCount field value
func (o *GuildPreviewResponse) GetApproximatePresenceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproximatePresenceCount
}

// GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetApproximatePresenceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproximatePresenceCount, true
}

// SetApproximatePresenceCount sets field value
func (o *GuildPreviewResponse) SetApproximatePresenceCount(v int32) {
	o.ApproximatePresenceCount = v
}

// GetEmojis returns the Emojis field value
func (o *GuildPreviewResponse) GetEmojis() []EmojiResponse {
	if o == nil {
		var ret []EmojiResponse
		return ret
	}

	return o.Emojis
}

// GetEmojisOk returns a tuple with the Emojis field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetEmojisOk() ([]EmojiResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emojis, true
}

// SetEmojis sets field value
func (o *GuildPreviewResponse) SetEmojis(v []EmojiResponse) {
	o.Emojis = v
}

// GetStickers returns the Stickers field value
func (o *GuildPreviewResponse) GetStickers() []GuildStickerResponse {
	if o == nil {
		var ret []GuildStickerResponse
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *GuildPreviewResponse) GetStickersOk() ([]GuildStickerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *GuildPreviewResponse) SetStickers(v []GuildStickerResponse) {
	o.Stickers = v
}

func (o GuildPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.HomeHeader.IsSet() {
		toSerialize["home_header"] = o.HomeHeader.Get()
	}
	if o.Splash.IsSet() {
		toSerialize["splash"] = o.Splash.Get()
	}
	if o.DiscoverySplash.IsSet() {
		toSerialize["discovery_splash"] = o.DiscoverySplash.Get()
	}
	toSerialize["features"] = o.Features
	toSerialize["approximate_member_count"] = o.ApproximateMemberCount
	toSerialize["approximate_presence_count"] = o.ApproximatePresenceCount
	toSerialize["emojis"] = o.Emojis
	toSerialize["stickers"] = o.Stickers
	return toSerialize, nil
}

func (o *GuildPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"features",
		"approximate_member_count",
		"approximate_presence_count",
		"emojis",
		"stickers",
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

	varGuildPreviewResponse := _GuildPreviewResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildPreviewResponse)

	if err != nil {
		return err
	}

	*o = GuildPreviewResponse(varGuildPreviewResponse)

	return err
}

type NullableGuildPreviewResponse struct {
	value *GuildPreviewResponse
	isSet bool
}

func (v NullableGuildPreviewResponse) Get() *GuildPreviewResponse {
	return v.value
}

func (v *NullableGuildPreviewResponse) Set(val *GuildPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildPreviewResponse(val *GuildPreviewResponse) *NullableGuildPreviewResponse {
	return &NullableGuildPreviewResponse{value: val, isSet: true}
}

func (v NullableGuildPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


