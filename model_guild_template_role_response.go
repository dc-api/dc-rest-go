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

// checks if the GuildTemplateRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateRoleResponse{}

// GuildTemplateRoleResponse struct for GuildTemplateRoleResponse
type GuildTemplateRoleResponse struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Permissions string `json:"permissions"`
	Color int32 `json:"color"`
	Hoist bool `json:"hoist"`
	Mentionable bool `json:"mentionable"`
	Icon NullableString `json:"icon,omitempty"`
	UnicodeEmoji NullableString `json:"unicode_emoji,omitempty"`
}

type _GuildTemplateRoleResponse GuildTemplateRoleResponse

// NewGuildTemplateRoleResponse instantiates a new GuildTemplateRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateRoleResponse(id int32, name string, permissions string, color int32, hoist bool, mentionable bool) *GuildTemplateRoleResponse {
	this := GuildTemplateRoleResponse{}
	this.Id = id
	this.Name = name
	this.Permissions = permissions
	this.Color = color
	this.Hoist = hoist
	this.Mentionable = mentionable
	return &this
}

// NewGuildTemplateRoleResponseWithDefaults instantiates a new GuildTemplateRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildTemplateRoleResponseWithDefaults() *GuildTemplateRoleResponse {
	this := GuildTemplateRoleResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildTemplateRoleResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildTemplateRoleResponse) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GuildTemplateRoleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildTemplateRoleResponse) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value
func (o *GuildTemplateRoleResponse) GetPermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleResponse) GetPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *GuildTemplateRoleResponse) SetPermissions(v string) {
	o.Permissions = v
}

// GetColor returns the Color field value
func (o *GuildTemplateRoleResponse) GetColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleResponse) GetColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *GuildTemplateRoleResponse) SetColor(v int32) {
	o.Color = v
}

// GetHoist returns the Hoist field value
func (o *GuildTemplateRoleResponse) GetHoist() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hoist
}

// GetHoistOk returns a tuple with the Hoist field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleResponse) GetHoistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hoist, true
}

// SetHoist sets field value
func (o *GuildTemplateRoleResponse) SetHoist(v bool) {
	o.Hoist = v
}

// GetMentionable returns the Mentionable field value
func (o *GuildTemplateRoleResponse) GetMentionable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mentionable
}

// GetMentionableOk returns a tuple with the Mentionable field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateRoleResponse) GetMentionableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mentionable, true
}

// SetMentionable sets field value
func (o *GuildTemplateRoleResponse) SetMentionable(v bool) {
	o.Mentionable = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateRoleResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateRoleResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildTemplateRoleResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GuildTemplateRoleResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *GuildTemplateRoleResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GuildTemplateRoleResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetUnicodeEmoji returns the UnicodeEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateRoleResponse) GetUnicodeEmoji() string {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		var ret string
		return ret
	}
	return *o.UnicodeEmoji.Get()
}

// GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateRoleResponse) GetUnicodeEmojiOk() (*string, bool) {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		return nil, false
	}
	return o.UnicodeEmoji.Get(), o.UnicodeEmoji.IsSet()
}

// HasUnicodeEmoji returns a boolean if a field has been set.
func (o *GuildTemplateRoleResponse) HasUnicodeEmoji() bool {
	if o != nil && o.UnicodeEmoji.IsSet() {
		return true
	}

	return false
}

// SetUnicodeEmoji gets a reference to the given NullableString and assigns it to the UnicodeEmoji field.
func (o *GuildTemplateRoleResponse) SetUnicodeEmoji(v string) {
	o.UnicodeEmoji.Set(&v)
}

// SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil
func (o *GuildTemplateRoleResponse) SetUnicodeEmojiNil() {
	o.UnicodeEmoji.Set(nil)
}

// UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil
func (o *GuildTemplateRoleResponse) UnsetUnicodeEmoji() {
	o.UnicodeEmoji.Unset()
}

func (o GuildTemplateRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildTemplateRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["permissions"] = o.Permissions
	toSerialize["color"] = o.Color
	toSerialize["hoist"] = o.Hoist
	toSerialize["mentionable"] = o.Mentionable
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.UnicodeEmoji.IsSet() {
		toSerialize["unicode_emoji"] = o.UnicodeEmoji.Get()
	}
	return toSerialize, nil
}

func (o *GuildTemplateRoleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"permissions",
		"color",
		"hoist",
		"mentionable",
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

	varGuildTemplateRoleResponse := _GuildTemplateRoleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildTemplateRoleResponse)

	if err != nil {
		return err
	}

	*o = GuildTemplateRoleResponse(varGuildTemplateRoleResponse)

	return err
}

type NullableGuildTemplateRoleResponse struct {
	value *GuildTemplateRoleResponse
	isSet bool
}

func (v NullableGuildTemplateRoleResponse) Get() *GuildTemplateRoleResponse {
	return v.value
}

func (v *NullableGuildTemplateRoleResponse) Set(val *GuildTemplateRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildTemplateRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildTemplateRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildTemplateRoleResponse(val *GuildTemplateRoleResponse) *NullableGuildTemplateRoleResponse {
	return &NullableGuildTemplateRoleResponse{value: val, isSet: true}
}

func (v NullableGuildTemplateRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildTemplateRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


