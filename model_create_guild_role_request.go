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
)

// checks if the CreateGuildRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildRoleRequest{}

// CreateGuildRoleRequest struct for CreateGuildRoleRequest
type CreateGuildRoleRequest struct {
	Name NullableString `json:"name,omitempty"`
	Permissions NullableInt32 `json:"permissions,omitempty"`
	Color NullableInt32 `json:"color,omitempty"`
	Hoist NullableBool `json:"hoist,omitempty"`
	Mentionable NullableBool `json:"mentionable,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	UnicodeEmoji NullableString `json:"unicode_emoji,omitempty"`
}

// NewCreateGuildRoleRequest instantiates a new CreateGuildRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGuildRoleRequest() *CreateGuildRoleRequest {
	this := CreateGuildRoleRequest{}
	return &this
}

// NewCreateGuildRoleRequestWithDefaults instantiates a new CreateGuildRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGuildRoleRequestWithDefaults() *CreateGuildRoleRequest {
	this := CreateGuildRoleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name.Get()) {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateGuildRoleRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateGuildRoleRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetName() {
	o.Name.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetPermissions() int32 {
	if o == nil || IsNil(o.Permissions.Get()) {
		var ret int32
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetPermissionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Permissions.Get()) {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasPermissions() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableInt32 and assigns it to the Permissions field.
func (o *CreateGuildRoleRequest) SetPermissions(v int32) {
	o.Permissions.Set(&v)
}

// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *CreateGuildRoleRequest) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetPermissions() {
	o.Permissions.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetColor() int32 {
	if o == nil || IsNil(o.Color.Get()) {
		var ret int32
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetColorOk() (*int32, bool) {
	if o == nil || IsNil(o.Color.Get()) {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableInt32 and assigns it to the Color field.
func (o *CreateGuildRoleRequest) SetColor(v int32) {
	o.Color.Set(&v)
}

// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateGuildRoleRequest) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetColor() {
	o.Color.Unset()
}

// GetHoist returns the Hoist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetHoist() bool {
	if o == nil || IsNil(o.Hoist.Get()) {
		var ret bool
		return ret
	}
	return *o.Hoist.Get()
}

// GetHoistOk returns a tuple with the Hoist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetHoistOk() (*bool, bool) {
	if o == nil || IsNil(o.Hoist.Get()) {
		return nil, false
	}
	return o.Hoist.Get(), o.Hoist.IsSet()
}

// HasHoist returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasHoist() bool {
	if o != nil && o.Hoist.IsSet() {
		return true
	}

	return false
}

// SetHoist gets a reference to the given NullableBool and assigns it to the Hoist field.
func (o *CreateGuildRoleRequest) SetHoist(v bool) {
	o.Hoist.Set(&v)
}

// SetHoistNil sets the value for Hoist to be an explicit nil
func (o *CreateGuildRoleRequest) SetHoistNil() {
	o.Hoist.Set(nil)
}

// UnsetHoist ensures that no value is present for Hoist, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetHoist() {
	o.Hoist.Unset()
}

// GetMentionable returns the Mentionable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetMentionable() bool {
	if o == nil || IsNil(o.Mentionable.Get()) {
		var ret bool
		return ret
	}
	return *o.Mentionable.Get()
}

// GetMentionableOk returns a tuple with the Mentionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetMentionableOk() (*bool, bool) {
	if o == nil || IsNil(o.Mentionable.Get()) {
		return nil, false
	}
	return o.Mentionable.Get(), o.Mentionable.IsSet()
}

// HasMentionable returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasMentionable() bool {
	if o != nil && o.Mentionable.IsSet() {
		return true
	}

	return false
}

// SetMentionable gets a reference to the given NullableBool and assigns it to the Mentionable field.
func (o *CreateGuildRoleRequest) SetMentionable(v bool) {
	o.Mentionable.Set(&v)
}

// SetMentionableNil sets the value for Mentionable to be an explicit nil
func (o *CreateGuildRoleRequest) SetMentionableNil() {
	o.Mentionable.Set(nil)
}

// UnsetMentionable ensures that no value is present for Mentionable, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetMentionable() {
	o.Mentionable.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *CreateGuildRoleRequest) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *CreateGuildRoleRequest) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetIcon() {
	o.Icon.Unset()
}

// GetUnicodeEmoji returns the UnicodeEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildRoleRequest) GetUnicodeEmoji() string {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		var ret string
		return ret
	}
	return *o.UnicodeEmoji.Get()
}

// GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildRoleRequest) GetUnicodeEmojiOk() (*string, bool) {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		return nil, false
	}
	return o.UnicodeEmoji.Get(), o.UnicodeEmoji.IsSet()
}

// HasUnicodeEmoji returns a boolean if a field has been set.
func (o *CreateGuildRoleRequest) HasUnicodeEmoji() bool {
	if o != nil && o.UnicodeEmoji.IsSet() {
		return true
	}

	return false
}

// SetUnicodeEmoji gets a reference to the given NullableString and assigns it to the UnicodeEmoji field.
func (o *CreateGuildRoleRequest) SetUnicodeEmoji(v string) {
	o.UnicodeEmoji.Set(&v)
}

// SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil
func (o *CreateGuildRoleRequest) SetUnicodeEmojiNil() {
	o.UnicodeEmoji.Set(nil)
}

// UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil
func (o *CreateGuildRoleRequest) UnsetUnicodeEmoji() {
	o.UnicodeEmoji.Unset()
}

func (o CreateGuildRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGuildRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Hoist.IsSet() {
		toSerialize["hoist"] = o.Hoist.Get()
	}
	if o.Mentionable.IsSet() {
		toSerialize["mentionable"] = o.Mentionable.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.UnicodeEmoji.IsSet() {
		toSerialize["unicode_emoji"] = o.UnicodeEmoji.Get()
	}
	return toSerialize, nil
}

type NullableCreateGuildRoleRequest struct {
	value *CreateGuildRoleRequest
	isSet bool
}

func (v NullableCreateGuildRoleRequest) Get() *CreateGuildRoleRequest {
	return v.value
}

func (v *NullableCreateGuildRoleRequest) Set(val *CreateGuildRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildRoleRequest(val *CreateGuildRoleRequest) *NullableCreateGuildRoleRequest {
	return &NullableCreateGuildRoleRequest{value: val, isSet: true}
}

func (v NullableCreateGuildRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


