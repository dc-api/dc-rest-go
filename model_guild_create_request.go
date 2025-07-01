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

// checks if the GuildCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildCreateRequest{}

// GuildCreateRequest struct for GuildCreateRequest
type GuildCreateRequest struct {
	Description NullableString `json:"description,omitempty"`
	Name string `json:"name"`
	Region NullableString `json:"region,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	VerificationLevel *int32 `json:"verification_level,omitempty"`
	DefaultMessageNotifications *int32 `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter *int32 `json:"explicit_content_filter,omitempty"`
	PreferredLocale *string `json:"preferred_locale,omitempty"`
	AfkTimeout NullableInt32 `json:"afk_timeout,omitempty"`
	Roles []CreateGuildRequestRoleItem `json:"roles,omitempty"`
	Channels []CreateGuildRequestChannelItem `json:"channels,omitempty"`
	AfkChannelId *string `json:"afk_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SystemChannelId *string `json:"system_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SystemChannelFlags NullableInt32 `json:"system_channel_flags,omitempty"`
}

type _GuildCreateRequest GuildCreateRequest

// NewGuildCreateRequest instantiates a new GuildCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildCreateRequest(name string) *GuildCreateRequest {
	this := GuildCreateRequest{}
	this.Name = name
	return &this
}

// NewGuildCreateRequestWithDefaults instantiates a new GuildCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildCreateRequestWithDefaults() *GuildCreateRequest {
	this := GuildCreateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildCreateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildCreateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildCreateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *GuildCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildCreateRequest) SetName(v string) {
	o.Name = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region.Get()) {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *GuildCreateRequest) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *GuildCreateRequest) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *GuildCreateRequest) UnsetRegion() {
	o.Region.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GuildCreateRequest) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *GuildCreateRequest) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GuildCreateRequest) UnsetIcon() {
	o.Icon.Unset()
}

// GetVerificationLevel returns the VerificationLevel field value if set, zero value otherwise.
func (o *GuildCreateRequest) GetVerificationLevel() int32 {
	if o == nil || IsNil(o.VerificationLevel) {
		var ret int32
		return ret
	}
	return *o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetVerificationLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.VerificationLevel) {
		return nil, false
	}
	return o.VerificationLevel, true
}

// HasVerificationLevel returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasVerificationLevel() bool {
	if o != nil && !IsNil(o.VerificationLevel) {
		return true
	}

	return false
}

// SetVerificationLevel gets a reference to the given int32 and assigns it to the VerificationLevel field.
func (o *GuildCreateRequest) SetVerificationLevel(v int32) {
	o.VerificationLevel = &v
}


// GetDefaultMessageNotifications returns the DefaultMessageNotifications field value if set, zero value otherwise.
func (o *GuildCreateRequest) GetDefaultMessageNotifications() int32 {
	if o == nil || IsNil(o.DefaultMessageNotifications) {
		var ret int32
		return ret
	}
	return *o.DefaultMessageNotifications
}

// GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetDefaultMessageNotificationsOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultMessageNotifications) {
		return nil, false
	}
	return o.DefaultMessageNotifications, true
}

// HasDefaultMessageNotifications returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasDefaultMessageNotifications() bool {
	if o != nil && !IsNil(o.DefaultMessageNotifications) {
		return true
	}

	return false
}

// SetDefaultMessageNotifications gets a reference to the given int32 and assigns it to the DefaultMessageNotifications field.
func (o *GuildCreateRequest) SetDefaultMessageNotifications(v int32) {
	o.DefaultMessageNotifications = &v
}


// GetExplicitContentFilter returns the ExplicitContentFilter field value if set, zero value otherwise.
func (o *GuildCreateRequest) GetExplicitContentFilter() int32 {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		var ret int32
		return ret
	}
	return *o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetExplicitContentFilterOk() (*int32, bool) {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		return nil, false
	}
	return o.ExplicitContentFilter, true
}

// HasExplicitContentFilter returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasExplicitContentFilter() bool {
	if o != nil && !IsNil(o.ExplicitContentFilter) {
		return true
	}

	return false
}

// SetExplicitContentFilter gets a reference to the given int32 and assigns it to the ExplicitContentFilter field.
func (o *GuildCreateRequest) SetExplicitContentFilter(v int32) {
	o.ExplicitContentFilter = &v
}


// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *GuildCreateRequest) GetPreferredLocale() string {
	if o == nil || IsNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredLocale) {
		return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasPreferredLocale() bool {
	if o != nil && !IsNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *GuildCreateRequest) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}


// GetAfkTimeout returns the AfkTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetAfkTimeout() int32 {
	if o == nil || IsNil(o.AfkTimeout.Get()) {
		var ret int32
		return ret
	}
	return *o.AfkTimeout.Get()
}

// GetAfkTimeoutOk returns a tuple with the AfkTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetAfkTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.AfkTimeout.Get()) {
		return nil, false
	}
	return o.AfkTimeout.Get(), o.AfkTimeout.IsSet()
}

// HasAfkTimeout returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasAfkTimeout() bool {
	if o != nil && o.AfkTimeout.IsSet() {
		return true
	}

	return false
}

// SetAfkTimeout gets a reference to the given NullableInt32 and assigns it to the AfkTimeout field.
func (o *GuildCreateRequest) SetAfkTimeout(v int32) {
	o.AfkTimeout.Set(&v)
}

// SetAfkTimeoutNil sets the value for AfkTimeout to be an explicit nil
func (o *GuildCreateRequest) SetAfkTimeoutNil() {
	o.AfkTimeout.Set(nil)
}

// UnsetAfkTimeout ensures that no value is present for AfkTimeout, not even an explicit nil
func (o *GuildCreateRequest) UnsetAfkTimeout() {
	o.AfkTimeout.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetRoles() []CreateGuildRequestRoleItem {
	if o == nil {
		var ret []CreateGuildRequestRoleItem
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetRolesOk() ([]CreateGuildRequestRoleItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CreateGuildRequestRoleItem and assigns it to the Roles field.
func (o *GuildCreateRequest) SetRoles(v []CreateGuildRequestRoleItem) {
	o.Roles = v
}


// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetChannels() []CreateGuildRequestChannelItem {
	if o == nil {
		var ret []CreateGuildRequestChannelItem
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetChannelsOk() ([]CreateGuildRequestChannelItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []CreateGuildRequestChannelItem and assigns it to the Channels field.
func (o *GuildCreateRequest) SetChannels(v []CreateGuildRequestChannelItem) {
	o.Channels = v
}


// GetAfkChannelId returns the AfkChannelId field value if set, zero value otherwise.
func (o *GuildCreateRequest) GetAfkChannelId() string {
	if o == nil || IsNil(o.AfkChannelId) {
		var ret string
		return ret
	}
	return *o.AfkChannelId
}

// GetAfkChannelIdOk returns a tuple with the AfkChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetAfkChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfkChannelId) {
		return nil, false
	}
	return o.AfkChannelId, true
}

// HasAfkChannelId returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasAfkChannelId() bool {
	if o != nil && !IsNil(o.AfkChannelId) {
		return true
	}

	return false
}

// SetAfkChannelId gets a reference to the given string and assigns it to the AfkChannelId field.
func (o *GuildCreateRequest) SetAfkChannelId(v string) {
	o.AfkChannelId = &v
}


// GetSystemChannelId returns the SystemChannelId field value if set, zero value otherwise.
func (o *GuildCreateRequest) GetSystemChannelId() string {
	if o == nil || IsNil(o.SystemChannelId) {
		var ret string
		return ret
	}
	return *o.SystemChannelId
}

// GetSystemChannelIdOk returns a tuple with the SystemChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildCreateRequest) GetSystemChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemChannelId) {
		return nil, false
	}
	return o.SystemChannelId, true
}

// HasSystemChannelId returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasSystemChannelId() bool {
	if o != nil && !IsNil(o.SystemChannelId) {
		return true
	}

	return false
}

// SetSystemChannelId gets a reference to the given string and assigns it to the SystemChannelId field.
func (o *GuildCreateRequest) SetSystemChannelId(v string) {
	o.SystemChannelId = &v
}


// GetSystemChannelFlags returns the SystemChannelFlags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildCreateRequest) GetSystemChannelFlags() int32 {
	if o == nil || IsNil(o.SystemChannelFlags.Get()) {
		var ret int32
		return ret
	}
	return *o.SystemChannelFlags.Get()
}

// GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildCreateRequest) GetSystemChannelFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemChannelFlags.Get()) {
		return nil, false
	}
	return o.SystemChannelFlags.Get(), o.SystemChannelFlags.IsSet()
}

// HasSystemChannelFlags returns a boolean if a field has been set.
func (o *GuildCreateRequest) HasSystemChannelFlags() bool {
	if o != nil && o.SystemChannelFlags.IsSet() {
		return true
	}

	return false
}

// SetSystemChannelFlags gets a reference to the given NullableInt32 and assigns it to the SystemChannelFlags field.
func (o *GuildCreateRequest) SetSystemChannelFlags(v int32) {
	o.SystemChannelFlags.Set(&v)
}

// SetSystemChannelFlagsNil sets the value for SystemChannelFlags to be an explicit nil
func (o *GuildCreateRequest) SetSystemChannelFlagsNil() {
	o.SystemChannelFlags.Set(nil)
}

// UnsetSystemChannelFlags ensures that no value is present for SystemChannelFlags, not even an explicit nil
func (o *GuildCreateRequest) UnsetSystemChannelFlags() {
	o.SystemChannelFlags.Unset()
}

func (o GuildCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if !IsNil(o.VerificationLevel) {
		toSerialize["verification_level"] = o.VerificationLevel
	}
	if !IsNil(o.DefaultMessageNotifications) {
		toSerialize["default_message_notifications"] = o.DefaultMessageNotifications
	}
	if !IsNil(o.ExplicitContentFilter) {
		toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	}
	if !IsNil(o.PreferredLocale) {
		toSerialize["preferred_locale"] = o.PreferredLocale
	}
	if o.AfkTimeout.IsSet() {
		toSerialize["afk_timeout"] = o.AfkTimeout.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.AfkChannelId) {
		toSerialize["afk_channel_id"] = o.AfkChannelId
	}
	if !IsNil(o.SystemChannelId) {
		toSerialize["system_channel_id"] = o.SystemChannelId
	}
	if o.SystemChannelFlags.IsSet() {
		toSerialize["system_channel_flags"] = o.SystemChannelFlags.Get()
	}
	return toSerialize, nil
}

func (o *GuildCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGuildCreateRequest := _GuildCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildCreateRequest)

	if err != nil {
		return err
	}

	*o = GuildCreateRequest(varGuildCreateRequest)

	return err
}

type NullableGuildCreateRequest struct {
	value *GuildCreateRequest
	isSet bool
}

func (v NullableGuildCreateRequest) Get() *GuildCreateRequest {
	return v.value
}

func (v *NullableGuildCreateRequest) Set(val *GuildCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildCreateRequest(val *GuildCreateRequest) *NullableGuildCreateRequest {
	return &NullableGuildCreateRequest{value: val, isSet: true}
}

func (v NullableGuildCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


