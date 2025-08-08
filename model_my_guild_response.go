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
	"bytes"
	"fmt"
)

// checks if the MyGuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyGuildResponse{}

// MyGuildResponse struct for MyGuildResponse
type MyGuildResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
	Banner NullableString `json:"banner,omitempty"`
	Owner bool `json:"owner"`
	Permissions string `json:"permissions"`
	Features []string `json:"features"`
	ApproximateMemberCount NullableInt32 `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount NullableInt32 `json:"approximate_presence_count,omitempty"`
}

type _MyGuildResponse MyGuildResponse

// NewMyGuildResponse instantiates a new MyGuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyGuildResponse(id string, name string, owner bool, permissions string, features []string) *MyGuildResponse {
	this := MyGuildResponse{}
	this.Id = id
	this.Name = name
	this.Owner = owner
	this.Permissions = permissions
	this.Features = features
	return &this
}

// NewMyGuildResponseWithDefaults instantiates a new MyGuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyGuildResponseWithDefaults() *MyGuildResponse {
	this := MyGuildResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MyGuildResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MyGuildResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MyGuildResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MyGuildResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MyGuildResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MyGuildResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MyGuildResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MyGuildResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *MyGuildResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *MyGuildResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *MyGuildResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *MyGuildResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MyGuildResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MyGuildResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner.Get()) {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *MyGuildResponse) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *MyGuildResponse) SetBanner(v string) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil
func (o *MyGuildResponse) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *MyGuildResponse) UnsetBanner() {
	o.Banner.Unset()
}

// GetOwner returns the Owner field value
func (o *MyGuildResponse) GetOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *MyGuildResponse) GetOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *MyGuildResponse) SetOwner(v bool) {
	o.Owner = v
}

// GetPermissions returns the Permissions field value
func (o *MyGuildResponse) GetPermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *MyGuildResponse) GetPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *MyGuildResponse) SetPermissions(v string) {
	o.Permissions = v
}

// GetFeatures returns the Features field value
func (o *MyGuildResponse) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *MyGuildResponse) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *MyGuildResponse) SetFeatures(v []string) {
	o.Features = v
}

// GetApproximateMemberCount returns the ApproximateMemberCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MyGuildResponse) GetApproximateMemberCount() int32 {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximateMemberCount.Get()
}

// GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MyGuildResponse) GetApproximateMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximateMemberCount.Get()) {
		return nil, false
	}
	return o.ApproximateMemberCount.Get(), o.ApproximateMemberCount.IsSet()
}

// HasApproximateMemberCount returns a boolean if a field has been set.
func (o *MyGuildResponse) HasApproximateMemberCount() bool {
	if o != nil && o.ApproximateMemberCount.IsSet() {
		return true
	}

	return false
}

// SetApproximateMemberCount gets a reference to the given NullableInt32 and assigns it to the ApproximateMemberCount field.
func (o *MyGuildResponse) SetApproximateMemberCount(v int32) {
	o.ApproximateMemberCount.Set(&v)
}

// SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil
func (o *MyGuildResponse) SetApproximateMemberCountNil() {
	o.ApproximateMemberCount.Set(nil)
}

// UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
func (o *MyGuildResponse) UnsetApproximateMemberCount() {
	o.ApproximateMemberCount.Unset()
}

// GetApproximatePresenceCount returns the ApproximatePresenceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MyGuildResponse) GetApproximatePresenceCount() int32 {
	if o == nil || IsNil(o.ApproximatePresenceCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximatePresenceCount.Get()
}

// GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MyGuildResponse) GetApproximatePresenceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximatePresenceCount.Get()) {
		return nil, false
	}
	return o.ApproximatePresenceCount.Get(), o.ApproximatePresenceCount.IsSet()
}

// HasApproximatePresenceCount returns a boolean if a field has been set.
func (o *MyGuildResponse) HasApproximatePresenceCount() bool {
	if o != nil && o.ApproximatePresenceCount.IsSet() {
		return true
	}

	return false
}

// SetApproximatePresenceCount gets a reference to the given NullableInt32 and assigns it to the ApproximatePresenceCount field.
func (o *MyGuildResponse) SetApproximatePresenceCount(v int32) {
	o.ApproximatePresenceCount.Set(&v)
}

// SetApproximatePresenceCountNil sets the value for ApproximatePresenceCount to be an explicit nil
func (o *MyGuildResponse) SetApproximatePresenceCountNil() {
	o.ApproximatePresenceCount.Set(nil)
}

// UnsetApproximatePresenceCount ensures that no value is present for ApproximatePresenceCount, not even an explicit nil
func (o *MyGuildResponse) UnsetApproximatePresenceCount() {
	o.ApproximatePresenceCount.Unset()
}

func (o MyGuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyGuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	toSerialize["owner"] = o.Owner
	toSerialize["permissions"] = o.Permissions
	toSerialize["features"] = o.Features
	if o.ApproximateMemberCount.IsSet() {
		toSerialize["approximate_member_count"] = o.ApproximateMemberCount.Get()
	}
	if o.ApproximatePresenceCount.IsSet() {
		toSerialize["approximate_presence_count"] = o.ApproximatePresenceCount.Get()
	}
	return toSerialize, nil
}

func (o *MyGuildResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"owner",
		"permissions",
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

	varMyGuildResponse := _MyGuildResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMyGuildResponse)

	if err != nil {
		return err
	}

	*o = MyGuildResponse(varMyGuildResponse)

	return err
}

type NullableMyGuildResponse struct {
	value *MyGuildResponse
	isSet bool
}

func (v NullableMyGuildResponse) Get() *MyGuildResponse {
	return v.value
}

func (v *NullableMyGuildResponse) Set(val *MyGuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMyGuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMyGuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyGuildResponse(val *MyGuildResponse) *NullableMyGuildResponse {
	return &NullableMyGuildResponse{value: val, isSet: true}
}

func (v NullableMyGuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyGuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


