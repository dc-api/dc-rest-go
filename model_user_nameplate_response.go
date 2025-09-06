/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-09-06T05:32:22.317073992Z[Etc/UTC]
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

// checks if the UserNameplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserNameplateResponse{}

// UserNameplateResponse struct for UserNameplateResponse
type UserNameplateResponse struct {
	SkuId *string `json:"sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Asset string `json:"asset"`
	Label string `json:"label"`
	Palette string `json:"palette"`
}

type _UserNameplateResponse UserNameplateResponse

// NewUserNameplateResponse instantiates a new UserNameplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNameplateResponse(asset string, label string, palette string) *UserNameplateResponse {
	this := UserNameplateResponse{}
	this.Asset = asset
	this.Label = label
	this.Palette = palette
	return &this
}

// NewUserNameplateResponseWithDefaults instantiates a new UserNameplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNameplateResponseWithDefaults() *UserNameplateResponse {
	this := UserNameplateResponse{}
	return &this
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *UserNameplateResponse) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNameplateResponse) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *UserNameplateResponse) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *UserNameplateResponse) SetSkuId(v string) {
	o.SkuId = &v
}


// GetAsset returns the Asset field value
func (o *UserNameplateResponse) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *UserNameplateResponse) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *UserNameplateResponse) SetAsset(v string) {
	o.Asset = v
}

// GetLabel returns the Label field value
func (o *UserNameplateResponse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *UserNameplateResponse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *UserNameplateResponse) SetLabel(v string) {
	o.Label = v
}

// GetPalette returns the Palette field value
func (o *UserNameplateResponse) GetPalette() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *UserNameplateResponse) GetPaletteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value
func (o *UserNameplateResponse) SetPalette(v string) {
	o.Palette = v
}

func (o UserNameplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserNameplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkuId) {
		toSerialize["sku_id"] = o.SkuId
	}
	toSerialize["asset"] = o.Asset
	toSerialize["label"] = o.Label
	toSerialize["palette"] = o.Palette
	return toSerialize, nil
}

func (o *UserNameplateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset",
		"label",
		"palette",
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

	varUserNameplateResponse := _UserNameplateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserNameplateResponse)

	if err != nil {
		return err
	}

	*o = UserNameplateResponse(varUserNameplateResponse)

	return err
}

type NullableUserNameplateResponse struct {
	value *UserNameplateResponse
	isSet bool
}

func (v NullableUserNameplateResponse) Get() *UserNameplateResponse {
	return v.value
}

func (v *NullableUserNameplateResponse) Set(val *UserNameplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNameplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNameplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNameplateResponse(val *UserNameplateResponse) *NullableUserNameplateResponse {
	return &NullableUserNameplateResponse{value: val, isSet: true}
}

func (v NullableUserNameplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNameplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


