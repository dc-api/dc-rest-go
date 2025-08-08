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

// checks if the GuildTemplateSnapshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateSnapshotResponse{}

// GuildTemplateSnapshotResponse struct for GuildTemplateSnapshotResponse
type GuildTemplateSnapshotResponse struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Region NullableString `json:"region,omitempty"`
	VerificationLevel int32 `json:"verification_level"`
	DefaultMessageNotifications int32 `json:"default_message_notifications"`
	ExplicitContentFilter int32 `json:"explicit_content_filter"`
	PreferredLocale string `json:"preferred_locale"`
	AfkChannelId *string `json:"afk_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	AfkTimeout NullableInt32 `json:"afk_timeout"`
	SystemChannelId *string `json:"system_channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SystemChannelFlags int32 `json:"system_channel_flags"`
	Roles []GuildTemplateRoleResponse `json:"roles"`
	Channels []GuildTemplateChannelResponse `json:"channels"`
}

type _GuildTemplateSnapshotResponse GuildTemplateSnapshotResponse

// NewGuildTemplateSnapshotResponse instantiates a new GuildTemplateSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateSnapshotResponse(name string, verificationLevel int32, defaultMessageNotifications int32, explicitContentFilter int32, preferredLocale string, afkTimeout NullableInt32, systemChannelFlags int32, roles []GuildTemplateRoleResponse, channels []GuildTemplateChannelResponse) *GuildTemplateSnapshotResponse {
	this := GuildTemplateSnapshotResponse{}
	this.Name = name
	this.VerificationLevel = verificationLevel
	this.DefaultMessageNotifications = defaultMessageNotifications
	this.ExplicitContentFilter = explicitContentFilter
	this.PreferredLocale = preferredLocale
	this.AfkTimeout = afkTimeout
	this.SystemChannelFlags = systemChannelFlags
	this.Roles = roles
	this.Channels = channels
	return &this
}

// NewGuildTemplateSnapshotResponseWithDefaults instantiates a new GuildTemplateSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildTemplateSnapshotResponseWithDefaults() *GuildTemplateSnapshotResponse {
	this := GuildTemplateSnapshotResponse{}
	return &this
}

// GetName returns the Name field value
func (o *GuildTemplateSnapshotResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildTemplateSnapshotResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateSnapshotResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateSnapshotResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildTemplateSnapshotResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildTemplateSnapshotResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildTemplateSnapshotResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildTemplateSnapshotResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateSnapshotResponse) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateSnapshotResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region.Get()) {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *GuildTemplateSnapshotResponse) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *GuildTemplateSnapshotResponse) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *GuildTemplateSnapshotResponse) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *GuildTemplateSnapshotResponse) UnsetRegion() {
	o.Region.Unset()
}

// GetVerificationLevel returns the VerificationLevel field value
func (o *GuildTemplateSnapshotResponse) GetVerificationLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VerificationLevel
}

// GetVerificationLevelOk returns a tuple with the VerificationLevel field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetVerificationLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationLevel, true
}

// SetVerificationLevel sets field value
func (o *GuildTemplateSnapshotResponse) SetVerificationLevel(v int32) {
	o.VerificationLevel = v
}

// GetDefaultMessageNotifications returns the DefaultMessageNotifications field value
func (o *GuildTemplateSnapshotResponse) GetDefaultMessageNotifications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultMessageNotifications
}

// GetDefaultMessageNotificationsOk returns a tuple with the DefaultMessageNotifications field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetDefaultMessageNotificationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMessageNotifications, true
}

// SetDefaultMessageNotifications sets field value
func (o *GuildTemplateSnapshotResponse) SetDefaultMessageNotifications(v int32) {
	o.DefaultMessageNotifications = v
}

// GetExplicitContentFilter returns the ExplicitContentFilter field value
func (o *GuildTemplateSnapshotResponse) GetExplicitContentFilter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetExplicitContentFilterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplicitContentFilter, true
}

// SetExplicitContentFilter sets field value
func (o *GuildTemplateSnapshotResponse) SetExplicitContentFilter(v int32) {
	o.ExplicitContentFilter = v
}

// GetPreferredLocale returns the PreferredLocale field value
func (o *GuildTemplateSnapshotResponse) GetPreferredLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetPreferredLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreferredLocale, true
}

// SetPreferredLocale sets field value
func (o *GuildTemplateSnapshotResponse) SetPreferredLocale(v string) {
	o.PreferredLocale = v
}

// GetAfkChannelId returns the AfkChannelId field value if set, zero value otherwise.
func (o *GuildTemplateSnapshotResponse) GetAfkChannelId() string {
	if o == nil || IsNil(o.AfkChannelId) {
		var ret string
		return ret
	}
	return *o.AfkChannelId
}

// GetAfkChannelIdOk returns a tuple with the AfkChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetAfkChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfkChannelId) {
		return nil, false
	}
	return o.AfkChannelId, true
}

// HasAfkChannelId returns a boolean if a field has been set.
func (o *GuildTemplateSnapshotResponse) HasAfkChannelId() bool {
	if o != nil && !IsNil(o.AfkChannelId) {
		return true
	}

	return false
}

// SetAfkChannelId gets a reference to the given string and assigns it to the AfkChannelId field.
func (o *GuildTemplateSnapshotResponse) SetAfkChannelId(v string) {
	o.AfkChannelId = &v
}


// GetAfkTimeout returns the AfkTimeout field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GuildTemplateSnapshotResponse) GetAfkTimeout() int32 {
	if o == nil || o.AfkTimeout.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AfkTimeout.Get()
}

// GetAfkTimeoutOk returns a tuple with the AfkTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateSnapshotResponse) GetAfkTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfkTimeout.Get(), o.AfkTimeout.IsSet()
}

// SetAfkTimeout sets field value
func (o *GuildTemplateSnapshotResponse) SetAfkTimeout(v int32) {
	o.AfkTimeout.Set(&v)
}

// GetSystemChannelId returns the SystemChannelId field value if set, zero value otherwise.
func (o *GuildTemplateSnapshotResponse) GetSystemChannelId() string {
	if o == nil || IsNil(o.SystemChannelId) {
		var ret string
		return ret
	}
	return *o.SystemChannelId
}

// GetSystemChannelIdOk returns a tuple with the SystemChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetSystemChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemChannelId) {
		return nil, false
	}
	return o.SystemChannelId, true
}

// HasSystemChannelId returns a boolean if a field has been set.
func (o *GuildTemplateSnapshotResponse) HasSystemChannelId() bool {
	if o != nil && !IsNil(o.SystemChannelId) {
		return true
	}

	return false
}

// SetSystemChannelId gets a reference to the given string and assigns it to the SystemChannelId field.
func (o *GuildTemplateSnapshotResponse) SetSystemChannelId(v string) {
	o.SystemChannelId = &v
}


// GetSystemChannelFlags returns the SystemChannelFlags field value
func (o *GuildTemplateSnapshotResponse) GetSystemChannelFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SystemChannelFlags
}

// GetSystemChannelFlagsOk returns a tuple with the SystemChannelFlags field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetSystemChannelFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemChannelFlags, true
}

// SetSystemChannelFlags sets field value
func (o *GuildTemplateSnapshotResponse) SetSystemChannelFlags(v int32) {
	o.SystemChannelFlags = v
}

// GetRoles returns the Roles field value
func (o *GuildTemplateSnapshotResponse) GetRoles() []GuildTemplateRoleResponse {
	if o == nil {
		var ret []GuildTemplateRoleResponse
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetRolesOk() ([]GuildTemplateRoleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *GuildTemplateSnapshotResponse) SetRoles(v []GuildTemplateRoleResponse) {
	o.Roles = v
}

// GetChannels returns the Channels field value
func (o *GuildTemplateSnapshotResponse) GetChannels() []GuildTemplateChannelResponse {
	if o == nil {
		var ret []GuildTemplateChannelResponse
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateSnapshotResponse) GetChannelsOk() ([]GuildTemplateChannelResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *GuildTemplateSnapshotResponse) SetChannels(v []GuildTemplateChannelResponse) {
	o.Channels = v
}

func (o GuildTemplateSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildTemplateSnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	toSerialize["verification_level"] = o.VerificationLevel
	toSerialize["default_message_notifications"] = o.DefaultMessageNotifications
	toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	toSerialize["preferred_locale"] = o.PreferredLocale
	if !IsNil(o.AfkChannelId) {
		toSerialize["afk_channel_id"] = o.AfkChannelId
	}
	toSerialize["afk_timeout"] = o.AfkTimeout.Get()
	if !IsNil(o.SystemChannelId) {
		toSerialize["system_channel_id"] = o.SystemChannelId
	}
	toSerialize["system_channel_flags"] = o.SystemChannelFlags
	toSerialize["roles"] = o.Roles
	toSerialize["channels"] = o.Channels
	return toSerialize, nil
}

func (o *GuildTemplateSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"verification_level",
		"default_message_notifications",
		"explicit_content_filter",
		"preferred_locale",
		"afk_timeout",
		"system_channel_flags",
		"roles",
		"channels",
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

	varGuildTemplateSnapshotResponse := _GuildTemplateSnapshotResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildTemplateSnapshotResponse)

	if err != nil {
		return err
	}

	*o = GuildTemplateSnapshotResponse(varGuildTemplateSnapshotResponse)

	return err
}

type NullableGuildTemplateSnapshotResponse struct {
	value *GuildTemplateSnapshotResponse
	isSet bool
}

func (v NullableGuildTemplateSnapshotResponse) Get() *GuildTemplateSnapshotResponse {
	return v.value
}

func (v *NullableGuildTemplateSnapshotResponse) Set(val *GuildTemplateSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildTemplateSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildTemplateSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildTemplateSnapshotResponse(val *GuildTemplateSnapshotResponse) *NullableGuildTemplateSnapshotResponse {
	return &NullableGuildTemplateSnapshotResponse{value: val, isSet: true}
}

func (v NullableGuildTemplateSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildTemplateSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


