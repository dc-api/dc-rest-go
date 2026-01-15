/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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
)

// checks if the UpdateRoleRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleRequestPartial{}

// UpdateRoleRequestPartial struct for UpdateRoleRequestPartial
type UpdateRoleRequestPartial struct {
	Name *string `json:"name,omitempty"`
	Permissions *int32 `json:"permissions,omitempty"`
	Color *int32 `json:"color,omitempty"`
	Colors *CreateRoleRequestColors `json:"colors,omitempty"`
	Hoist *bool `json:"hoist,omitempty"`
	Mentionable *bool `json:"mentionable,omitempty"`
	Icon *string `json:"icon,omitempty"`
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
}

// NewUpdateRoleRequestPartial instantiates a new UpdateRoleRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleRequestPartial() *UpdateRoleRequestPartial {
	this := UpdateRoleRequestPartial{}
	return &this
}

// NewUpdateRoleRequestPartialWithDefaults instantiates a new UpdateRoleRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleRequestPartialWithDefaults() *UpdateRoleRequestPartial {
	this := UpdateRoleRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRoleRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetPermissions() int32 {
	if o == nil || IsNil(o.Permissions) {
		var ret int32
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetPermissionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given int32 and assigns it to the Permissions field.
func (o *UpdateRoleRequestPartial) SetPermissions(v int32) {
	o.Permissions = &v
}


// GetColor returns the Color field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetColor() int32 {
	if o == nil || IsNil(o.Color) {
		var ret int32
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetColorOk() (*int32, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given int32 and assigns it to the Color field.
func (o *UpdateRoleRequestPartial) SetColor(v int32) {
	o.Color = &v
}


// GetColors returns the Colors field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetColors() CreateRoleRequestColors {
	if o == nil || IsNil(o.Colors) {
		var ret CreateRoleRequestColors
		return ret
	}
	return *o.Colors
}

// GetColorsOk returns a tuple with the Colors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetColorsOk() (*CreateRoleRequestColors, bool) {
	if o == nil || IsNil(o.Colors) {
		return nil, false
	}
	return o.Colors, true
}

// HasColors returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasColors() bool {
	if o != nil && !IsNil(o.Colors) {
		return true
	}

	return false
}

// SetColors gets a reference to the given CreateRoleRequestColors and assigns it to the Colors field.
func (o *UpdateRoleRequestPartial) SetColors(v CreateRoleRequestColors) {
	o.Colors = &v
}


// GetHoist returns the Hoist field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetHoist() bool {
	if o == nil || IsNil(o.Hoist) {
		var ret bool
		return ret
	}
	return *o.Hoist
}

// GetHoistOk returns a tuple with the Hoist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetHoistOk() (*bool, bool) {
	if o == nil || IsNil(o.Hoist) {
		return nil, false
	}
	return o.Hoist, true
}

// HasHoist returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasHoist() bool {
	if o != nil && !IsNil(o.Hoist) {
		return true
	}

	return false
}

// SetHoist gets a reference to the given bool and assigns it to the Hoist field.
func (o *UpdateRoleRequestPartial) SetHoist(v bool) {
	o.Hoist = &v
}


// GetMentionable returns the Mentionable field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetMentionable() bool {
	if o == nil || IsNil(o.Mentionable) {
		var ret bool
		return ret
	}
	return *o.Mentionable
}

// GetMentionableOk returns a tuple with the Mentionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetMentionableOk() (*bool, bool) {
	if o == nil || IsNil(o.Mentionable) {
		return nil, false
	}
	return o.Mentionable, true
}

// HasMentionable returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasMentionable() bool {
	if o != nil && !IsNil(o.Mentionable) {
		return true
	}

	return false
}

// SetMentionable gets a reference to the given bool and assigns it to the Mentionable field.
func (o *UpdateRoleRequestPartial) SetMentionable(v bool) {
	o.Mentionable = &v
}


// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *UpdateRoleRequestPartial) SetIcon(v string) {
	o.Icon = &v
}


// GetUnicodeEmoji returns the UnicodeEmoji field value if set, zero value otherwise.
func (o *UpdateRoleRequestPartial) GetUnicodeEmoji() string {
	if o == nil || IsNil(o.UnicodeEmoji) {
		var ret string
		return ret
	}
	return *o.UnicodeEmoji
}

// GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequestPartial) GetUnicodeEmojiOk() (*string, bool) {
	if o == nil || IsNil(o.UnicodeEmoji) {
		return nil, false
	}
	return o.UnicodeEmoji, true
}

// HasUnicodeEmoji returns a boolean if a field has been set.
func (o *UpdateRoleRequestPartial) HasUnicodeEmoji() bool {
	if o != nil && !IsNil(o.UnicodeEmoji) {
		return true
	}

	return false
}

// SetUnicodeEmoji gets a reference to the given string and assigns it to the UnicodeEmoji field.
func (o *UpdateRoleRequestPartial) SetUnicodeEmoji(v string) {
	o.UnicodeEmoji = &v
}


func (o UpdateRoleRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Colors) {
		toSerialize["colors"] = o.Colors
	}
	if !IsNil(o.Hoist) {
		toSerialize["hoist"] = o.Hoist
	}
	if !IsNil(o.Mentionable) {
		toSerialize["mentionable"] = o.Mentionable
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.UnicodeEmoji) {
		toSerialize["unicode_emoji"] = o.UnicodeEmoji
	}
	return toSerialize, nil
}

type NullableUpdateRoleRequestPartial struct {
	value *UpdateRoleRequestPartial
	isSet bool
}

func (v NullableUpdateRoleRequestPartial) Get() *UpdateRoleRequestPartial {
	return v.value
}

func (v *NullableUpdateRoleRequestPartial) Set(val *UpdateRoleRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleRequestPartial(val *UpdateRoleRequestPartial) *NullableUpdateRoleRequestPartial {
	return &NullableUpdateRoleRequestPartial{value: val, isSet: true}
}

func (v NullableUpdateRoleRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

