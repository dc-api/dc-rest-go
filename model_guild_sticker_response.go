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

// checks if the GuildStickerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildStickerResponse{}

// GuildStickerResponse struct for GuildStickerResponse
type GuildStickerResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Tags string `json:"tags"`
	Type int32 `json:"type"`
	FormatType NullableInt32 `json:"format_type,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Available bool `json:"available"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	User NullableUserResponse `json:"user,omitempty"`
}

type _GuildStickerResponse GuildStickerResponse

// NewGuildStickerResponse instantiates a new GuildStickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildStickerResponse(id string, name string, tags string, type_ int32, available bool, guildId string) *GuildStickerResponse {
	this := GuildStickerResponse{}
	this.Id = id
	this.Name = name
	this.Tags = tags
	this.Type = type_
	this.Available = available
	this.GuildId = guildId
	return &this
}

// NewGuildStickerResponseWithDefaults instantiates a new GuildStickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildStickerResponseWithDefaults() *GuildStickerResponse {
	this := GuildStickerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GuildStickerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildStickerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildStickerResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GuildStickerResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildStickerResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildStickerResponse) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value
func (o *GuildStickerResponse) GetTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *GuildStickerResponse) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *GuildStickerResponse) SetTags(v string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *GuildStickerResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GuildStickerResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GuildStickerResponse) SetType(v int32) {
	o.Type = v
}

// GetFormatType returns the FormatType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildStickerResponse) GetFormatType() int32 {
	if o == nil || IsNil(o.FormatType.Get()) {
		var ret int32
		return ret
	}
	return *o.FormatType.Get()
}

// GetFormatTypeOk returns a tuple with the FormatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildStickerResponse) GetFormatTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.FormatType.Get()) {
		return nil, false
	}
	return o.FormatType.Get(), o.FormatType.IsSet()
}

// HasFormatType returns a boolean if a field has been set.
func (o *GuildStickerResponse) HasFormatType() bool {
	if o != nil && o.FormatType.IsSet() {
		return true
	}

	return false
}

// SetFormatType gets a reference to the given NullableInt32 and assigns it to the FormatType field.
func (o *GuildStickerResponse) SetFormatType(v int32) {
	o.FormatType.Set(&v)
}

// SetFormatTypeNil sets the value for FormatType to be an explicit nil
func (o *GuildStickerResponse) SetFormatTypeNil() {
	o.FormatType.Set(nil)
}

// UnsetFormatType ensures that no value is present for FormatType, not even an explicit nil
func (o *GuildStickerResponse) UnsetFormatType() {
	o.FormatType.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildStickerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildStickerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildStickerResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildStickerResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildStickerResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildStickerResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetAvailable returns the Available field value
func (o *GuildStickerResponse) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *GuildStickerResponse) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *GuildStickerResponse) SetAvailable(v bool) {
	o.Available = v
}

// GetGuildId returns the GuildId field value
func (o *GuildStickerResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *GuildStickerResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *GuildStickerResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildStickerResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildStickerResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *GuildStickerResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *GuildStickerResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *GuildStickerResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *GuildStickerResponse) UnsetUser() {
	o.User.Unset()
}

func (o GuildStickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildStickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	if o.FormatType.IsSet() {
		toSerialize["format_type"] = o.FormatType.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["available"] = o.Available
	toSerialize["guild_id"] = o.GuildId
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *GuildStickerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"tags",
		"type",
		"available",
		"guild_id",
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

	varGuildStickerResponse := _GuildStickerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildStickerResponse)

	if err != nil {
		return err
	}

	*o = GuildStickerResponse(varGuildStickerResponse)

	return err
}

type NullableGuildStickerResponse struct {
	value *GuildStickerResponse
	isSet bool
}

func (v NullableGuildStickerResponse) Get() *GuildStickerResponse {
	return v.value
}

func (v *NullableGuildStickerResponse) Set(val *GuildStickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildStickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildStickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildStickerResponse(val *GuildStickerResponse) *NullableGuildStickerResponse {
	return &NullableGuildStickerResponse{value: val, isSet: true}
}

func (v NullableGuildStickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildStickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


