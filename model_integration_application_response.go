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

// checks if the IntegrationApplicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationApplicationResponse{}

// IntegrationApplicationResponse struct for IntegrationApplicationResponse
type IntegrationApplicationResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
	Description string `json:"description"`
	Type NullableInt32 `json:"type,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	PrimarySkuId *string `json:"primary_sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Bot NullableUserResponse `json:"bot,omitempty"`
}

type _IntegrationApplicationResponse IntegrationApplicationResponse

// NewIntegrationApplicationResponse instantiates a new IntegrationApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationApplicationResponse(id string, name string, description string) *IntegrationApplicationResponse {
	this := IntegrationApplicationResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewIntegrationApplicationResponseWithDefaults instantiates a new IntegrationApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationApplicationResponseWithDefaults() *IntegrationApplicationResponse {
	this := IntegrationApplicationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *IntegrationApplicationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegrationApplicationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegrationApplicationResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *IntegrationApplicationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationApplicationResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationApplicationResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationApplicationResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *IntegrationApplicationResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *IntegrationApplicationResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *IntegrationApplicationResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *IntegrationApplicationResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value
func (o *IntegrationApplicationResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IntegrationApplicationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *IntegrationApplicationResponse) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationApplicationResponse) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationApplicationResponse) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationApplicationResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *IntegrationApplicationResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *IntegrationApplicationResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *IntegrationApplicationResponse) UnsetType() {
	o.Type.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationApplicationResponse) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationApplicationResponse) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage.Get()) {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *IntegrationApplicationResponse) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *IntegrationApplicationResponse) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}

// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *IntegrationApplicationResponse) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *IntegrationApplicationResponse) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetPrimarySkuId returns the PrimarySkuId field value if set, zero value otherwise.
func (o *IntegrationApplicationResponse) GetPrimarySkuId() string {
	if o == nil || IsNil(o.PrimarySkuId) {
		var ret string
		return ret
	}
	return *o.PrimarySkuId
}

// GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationApplicationResponse) GetPrimarySkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimarySkuId) {
		return nil, false
	}
	return o.PrimarySkuId, true
}

// HasPrimarySkuId returns a boolean if a field has been set.
func (o *IntegrationApplicationResponse) HasPrimarySkuId() bool {
	if o != nil && !IsNil(o.PrimarySkuId) {
		return true
	}

	return false
}

// SetPrimarySkuId gets a reference to the given string and assigns it to the PrimarySkuId field.
func (o *IntegrationApplicationResponse) SetPrimarySkuId(v string) {
	o.PrimarySkuId = &v
}


// GetBot returns the Bot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationApplicationResponse) GetBot() UserResponse {
	if o == nil || IsNil(o.Bot.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Bot.Get()
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationApplicationResponse) GetBotOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Bot.Get()) {
		return nil, false
	}
	return o.Bot.Get(), o.Bot.IsSet()
}

// HasBot returns a boolean if a field has been set.
func (o *IntegrationApplicationResponse) HasBot() bool {
	if o != nil && o.Bot.IsSet() {
		return true
	}

	return false
}

// SetBot gets a reference to the given NullableUserResponse and assigns it to the Bot field.
func (o *IntegrationApplicationResponse) SetBot(v UserResponse) {
	o.Bot.Set(&v)
}

// SetBotNil sets the value for Bot to be an explicit nil
func (o *IntegrationApplicationResponse) SetBotNil() {
	o.Bot.Set(nil)
}

// UnsetBot ensures that no value is present for Bot, not even an explicit nil
func (o *IntegrationApplicationResponse) UnsetBot() {
	o.Bot.Unset()
}

func (o IntegrationApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationApplicationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["description"] = o.Description
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if !IsNil(o.PrimarySkuId) {
		toSerialize["primary_sku_id"] = o.PrimarySkuId
	}
	if o.Bot.IsSet() {
		toSerialize["bot"] = o.Bot.Get()
	}
	return toSerialize, nil
}

func (o *IntegrationApplicationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
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

	varIntegrationApplicationResponse := _IntegrationApplicationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIntegrationApplicationResponse)

	if err != nil {
		return err
	}

	*o = IntegrationApplicationResponse(varIntegrationApplicationResponse)

	return err
}

type NullableIntegrationApplicationResponse struct {
	value *IntegrationApplicationResponse
	isSet bool
}

func (v NullableIntegrationApplicationResponse) Get() *IntegrationApplicationResponse {
	return v.value
}

func (v *NullableIntegrationApplicationResponse) Set(val *IntegrationApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationApplicationResponse(val *IntegrationApplicationResponse) *NullableIntegrationApplicationResponse {
	return &NullableIntegrationApplicationResponse{value: val, isSet: true}
}

func (v NullableIntegrationApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


