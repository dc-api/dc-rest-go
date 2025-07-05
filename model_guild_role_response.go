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

// checks if the GuildRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildRoleResponse{}

// GuildRoleResponse struct for GuildRoleResponse
type GuildRoleResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Permissions string `json:"permissions"`
	Position int32 `json:"position"`
	Color int32 `json:"color"`
	Colors NullableGuildRoleColorsResponse `json:"colors,omitempty"`
	Hoist bool `json:"hoist"`
	Managed bool `json:"managed"`
	Mentionable bool `json:"mentionable"`
	Icon NullableString `json:"icon,omitempty"`
	UnicodeEmoji NullableString `json:"unicode_emoji,omitempty"`
	Tags NullableGuildRoleTagsResponse `json:"tags,omitempty"`
	Flags int32 `json:"flags"`
}

type _GuildRoleResponse GuildRoleResponse

// NewGuildRoleResponse instantiates a new GuildRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildRoleResponse(id string, name string, permissions string, position int32, color int32, hoist bool, managed bool, mentionable bool, flags int32) *GuildRoleResponse {
	this := GuildRoleResponse{}
	this.Id = id
	this.Name = name
	this.Permissions = permissions
	this.Position = position
	this.Color = color
	this.Hoist = hoist
	this.Managed = managed
	this.Mentionable = mentionable
	this.Flags = flags
	return &this
}

// NewGuildRoleResponseWithDefaults instantiates a new GuildRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildRoleResponseWithDefaults() *GuildRoleResponse {
	this := GuildRoleResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildRoleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildRoleResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GuildRoleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildRoleResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildRoleResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildRoleResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildRoleResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildRoleResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetPermissions returns the Permissions field value
func (o *GuildRoleResponse) GetPermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *GuildRoleResponse) SetPermissions(v string) {
	o.Permissions = v
}

// GetPosition returns the Position field value
func (o *GuildRoleResponse) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *GuildRoleResponse) SetPosition(v int32) {
	o.Position = v
}

// GetColor returns the Color field value
func (o *GuildRoleResponse) GetColor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *GuildRoleResponse) SetColor(v int32) {
	o.Color = v
}

// GetColors returns the Colors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleResponse) GetColors() GuildRoleColorsResponse {
	if o == nil || IsNil(o.Colors.Get()) {
		var ret GuildRoleColorsResponse
		return ret
	}
	return *o.Colors.Get()
}

// GetColorsOk returns a tuple with the Colors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleResponse) GetColorsOk() (*GuildRoleColorsResponse, bool) {
	if o == nil || IsNil(o.Colors.Get()) {
		return nil, false
	}
	return o.Colors.Get(), o.Colors.IsSet()
}

// HasColors returns a boolean if a field has been set.
func (o *GuildRoleResponse) HasColors() bool {
	if o != nil && o.Colors.IsSet() {
		return true
	}

	return false
}

// SetColors gets a reference to the given NullableGuildRoleColorsResponse and assigns it to the Colors field.
func (o *GuildRoleResponse) SetColors(v GuildRoleColorsResponse) {
	o.Colors.Set(&v)
}

// SetColorsNil sets the value for Colors to be an explicit nil
func (o *GuildRoleResponse) SetColorsNil() {
	o.Colors.Set(nil)
}

// UnsetColors ensures that no value is present for Colors, not even an explicit nil
func (o *GuildRoleResponse) UnsetColors() {
	o.Colors.Unset()
}

// GetHoist returns the Hoist field value
func (o *GuildRoleResponse) GetHoist() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hoist
}

// GetHoistOk returns a tuple with the Hoist field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetHoistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hoist, true
}

// SetHoist sets field value
func (o *GuildRoleResponse) SetHoist(v bool) {
	o.Hoist = v
}

// GetManaged returns the Managed field value
func (o *GuildRoleResponse) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *GuildRoleResponse) SetManaged(v bool) {
	o.Managed = v
}

// GetMentionable returns the Mentionable field value
func (o *GuildRoleResponse) GetMentionable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mentionable
}

// GetMentionableOk returns a tuple with the Mentionable field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetMentionableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mentionable, true
}

// SetMentionable sets field value
func (o *GuildRoleResponse) SetMentionable(v bool) {
	o.Mentionable = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GuildRoleResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GuildRoleResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *GuildRoleResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GuildRoleResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetUnicodeEmoji returns the UnicodeEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleResponse) GetUnicodeEmoji() string {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		var ret string
		return ret
	}
	return *o.UnicodeEmoji.Get()
}

// GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleResponse) GetUnicodeEmojiOk() (*string, bool) {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		return nil, false
	}
	return o.UnicodeEmoji.Get(), o.UnicodeEmoji.IsSet()
}

// HasUnicodeEmoji returns a boolean if a field has been set.
func (o *GuildRoleResponse) HasUnicodeEmoji() bool {
	if o != nil && o.UnicodeEmoji.IsSet() {
		return true
	}

	return false
}

// SetUnicodeEmoji gets a reference to the given NullableString and assigns it to the UnicodeEmoji field.
func (o *GuildRoleResponse) SetUnicodeEmoji(v string) {
	o.UnicodeEmoji.Set(&v)
}

// SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil
func (o *GuildRoleResponse) SetUnicodeEmojiNil() {
	o.UnicodeEmoji.Set(nil)
}

// UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil
func (o *GuildRoleResponse) UnsetUnicodeEmoji() {
	o.UnicodeEmoji.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildRoleResponse) GetTags() GuildRoleTagsResponse {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret GuildRoleTagsResponse
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildRoleResponse) GetTagsOk() (*GuildRoleTagsResponse, bool) {
	if o == nil || IsNil(o.Tags.Get()) {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *GuildRoleResponse) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableGuildRoleTagsResponse and assigns it to the Tags field.
func (o *GuildRoleResponse) SetTags(v GuildRoleTagsResponse) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *GuildRoleResponse) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *GuildRoleResponse) UnsetTags() {
	o.Tags.Unset()
}

// GetFlags returns the Flags field value
func (o *GuildRoleResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *GuildRoleResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *GuildRoleResponse) SetFlags(v int32) {
	o.Flags = v
}

func (o GuildRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["permissions"] = o.Permissions
	toSerialize["position"] = o.Position
	toSerialize["color"] = o.Color
	if o.Colors.IsSet() {
		toSerialize["colors"] = o.Colors.Get()
	}
	toSerialize["hoist"] = o.Hoist
	toSerialize["managed"] = o.Managed
	toSerialize["mentionable"] = o.Mentionable
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.UnicodeEmoji.IsSet() {
		toSerialize["unicode_emoji"] = o.UnicodeEmoji.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	toSerialize["flags"] = o.Flags
	return toSerialize, nil
}

func (o *GuildRoleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"permissions",
		"position",
		"color",
		"hoist",
		"managed",
		"mentionable",
		"flags",
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

	varGuildRoleResponse := _GuildRoleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildRoleResponse)

	if err != nil {
		return err
	}

	*o = GuildRoleResponse(varGuildRoleResponse)

	return err
}

type NullableGuildRoleResponse struct {
	value *GuildRoleResponse
	isSet bool
}

func (v NullableGuildRoleResponse) Get() *GuildRoleResponse {
	return v.value
}

func (v *NullableGuildRoleResponse) Set(val *GuildRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildRoleResponse(val *GuildRoleResponse) *NullableGuildRoleResponse {
	return &NullableGuildRoleResponse{value: val, isSet: true}
}

func (v NullableGuildRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


