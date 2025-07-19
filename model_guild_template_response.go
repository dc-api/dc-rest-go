/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the GuildTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildTemplateResponse{}

// GuildTemplateResponse struct for GuildTemplateResponse
type GuildTemplateResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	UsageCount int32 `json:"usage_count"`
	CreatorId string `json:"creator_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Creator NullableUserResponse `json:"creator,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SourceGuildId string `json:"source_guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SerializedSourceGuild GuildTemplateSnapshotResponse `json:"serialized_source_guild"`
	IsDirty NullableBool `json:"is_dirty,omitempty"`
}

type _GuildTemplateResponse GuildTemplateResponse

// NewGuildTemplateResponse instantiates a new GuildTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildTemplateResponse(code string, name string, usageCount int32, creatorId string, createdAt time.Time, updatedAt time.Time, sourceGuildId string, serializedSourceGuild GuildTemplateSnapshotResponse) *GuildTemplateResponse {
	this := GuildTemplateResponse{}
	this.Code = code
	this.Name = name
	this.UsageCount = usageCount
	this.CreatorId = creatorId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.SourceGuildId = sourceGuildId
	this.SerializedSourceGuild = serializedSourceGuild
	return &this
}

// NewGuildTemplateResponseWithDefaults instantiates a new GuildTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildTemplateResponseWithDefaults() *GuildTemplateResponse {
	this := GuildTemplateResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *GuildTemplateResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GuildTemplateResponse) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *GuildTemplateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GuildTemplateResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuildTemplateResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuildTemplateResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuildTemplateResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuildTemplateResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetUsageCount returns the UsageCount field value
func (o *GuildTemplateResponse) GetUsageCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetUsageCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageCount, true
}

// SetUsageCount sets field value
func (o *GuildTemplateResponse) SetUsageCount(v int32) {
	o.UsageCount = v
}

// GetCreatorId returns the CreatorId field value
func (o *GuildTemplateResponse) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *GuildTemplateResponse) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateResponse) GetCreator() UserResponse {
	if o == nil || IsNil(o.Creator.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateResponse) GetCreatorOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Creator.Get()) {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *GuildTemplateResponse) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given NullableUserResponse and assigns it to the Creator field.
func (o *GuildTemplateResponse) SetCreator(v UserResponse) {
	o.Creator.Set(&v)
}

// SetCreatorNil sets the value for Creator to be an explicit nil
func (o *GuildTemplateResponse) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil
func (o *GuildTemplateResponse) UnsetCreator() {
	o.Creator.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *GuildTemplateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GuildTemplateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GuildTemplateResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GuildTemplateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSourceGuildId returns the SourceGuildId field value
func (o *GuildTemplateResponse) GetSourceGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceGuildId
}

// GetSourceGuildIdOk returns a tuple with the SourceGuildId field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetSourceGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceGuildId, true
}

// SetSourceGuildId sets field value
func (o *GuildTemplateResponse) SetSourceGuildId(v string) {
	o.SourceGuildId = v
}

// GetSerializedSourceGuild returns the SerializedSourceGuild field value
func (o *GuildTemplateResponse) GetSerializedSourceGuild() GuildTemplateSnapshotResponse {
	if o == nil {
		var ret GuildTemplateSnapshotResponse
		return ret
	}

	return o.SerializedSourceGuild
}

// GetSerializedSourceGuildOk returns a tuple with the SerializedSourceGuild field value
// and a boolean to check if the value has been set.
func (o *GuildTemplateResponse) GetSerializedSourceGuildOk() (*GuildTemplateSnapshotResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerializedSourceGuild, true
}

// SetSerializedSourceGuild sets field value
func (o *GuildTemplateResponse) SetSerializedSourceGuild(v GuildTemplateSnapshotResponse) {
	o.SerializedSourceGuild = v
}

// GetIsDirty returns the IsDirty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildTemplateResponse) GetIsDirty() bool {
	if o == nil || IsNil(o.IsDirty.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDirty.Get()
}

// GetIsDirtyOk returns a tuple with the IsDirty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildTemplateResponse) GetIsDirtyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDirty.Get()) {
		return nil, false
	}
	return o.IsDirty.Get(), o.IsDirty.IsSet()
}

// HasIsDirty returns a boolean if a field has been set.
func (o *GuildTemplateResponse) HasIsDirty() bool {
	if o != nil && o.IsDirty.IsSet() {
		return true
	}

	return false
}

// SetIsDirty gets a reference to the given NullableBool and assigns it to the IsDirty field.
func (o *GuildTemplateResponse) SetIsDirty(v bool) {
	o.IsDirty.Set(&v)
}

// SetIsDirtyNil sets the value for IsDirty to be an explicit nil
func (o *GuildTemplateResponse) SetIsDirtyNil() {
	o.IsDirty.Set(nil)
}

// UnsetIsDirty ensures that no value is present for IsDirty, not even an explicit nil
func (o *GuildTemplateResponse) UnsetIsDirty() {
	o.IsDirty.Unset()
}

func (o GuildTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["usage_count"] = o.UsageCount
	toSerialize["creator_id"] = o.CreatorId
	if o.Creator.IsSet() {
		toSerialize["creator"] = o.Creator.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["source_guild_id"] = o.SourceGuildId
	toSerialize["serialized_source_guild"] = o.SerializedSourceGuild
	if o.IsDirty.IsSet() {
		toSerialize["is_dirty"] = o.IsDirty.Get()
	}
	return toSerialize, nil
}

func (o *GuildTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
		"usage_count",
		"creator_id",
		"created_at",
		"updated_at",
		"source_guild_id",
		"serialized_source_guild",
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

	varGuildTemplateResponse := _GuildTemplateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildTemplateResponse)

	if err != nil {
		return err
	}

	*o = GuildTemplateResponse(varGuildTemplateResponse)

	return err
}

type NullableGuildTemplateResponse struct {
	value *GuildTemplateResponse
	isSet bool
}

func (v NullableGuildTemplateResponse) Get() *GuildTemplateResponse {
	return v.value
}

func (v *NullableGuildTemplateResponse) Set(val *GuildTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildTemplateResponse(val *GuildTemplateResponse) *NullableGuildTemplateResponse {
	return &NullableGuildTemplateResponse{value: val, isSet: true}
}

func (v NullableGuildTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


