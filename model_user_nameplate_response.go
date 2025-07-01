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

// checks if the UserNameplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserNameplateResponse{}

// UserNameplateResponse struct for UserNameplateResponse
type UserNameplateResponse struct {
	SkuId *string `json:"sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Asset NullableString `json:"asset,omitempty"`
	Label NullableString `json:"label,omitempty"`
	Palette *string `json:"palette,omitempty"`
}

// NewUserNameplateResponse instantiates a new UserNameplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNameplateResponse() *UserNameplateResponse {
	this := UserNameplateResponse{}
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


// GetAsset returns the Asset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserNameplateResponse) GetAsset() string {
	if o == nil || IsNil(o.Asset.Get()) {
		var ret string
		return ret
	}
	return *o.Asset.Get()
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserNameplateResponse) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset.Get()) {
		return nil, false
	}
	return o.Asset.Get(), o.Asset.IsSet()
}

// HasAsset returns a boolean if a field has been set.
func (o *UserNameplateResponse) HasAsset() bool {
	if o != nil && o.Asset.IsSet() {
		return true
	}

	return false
}

// SetAsset gets a reference to the given NullableString and assigns it to the Asset field.
func (o *UserNameplateResponse) SetAsset(v string) {
	o.Asset.Set(&v)
}

// SetAssetNil sets the value for Asset to be an explicit nil
func (o *UserNameplateResponse) SetAssetNil() {
	o.Asset.Set(nil)
}

// UnsetAsset ensures that no value is present for Asset, not even an explicit nil
func (o *UserNameplateResponse) UnsetAsset() {
	o.Asset.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserNameplateResponse) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserNameplateResponse) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label.Get()) {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *UserNameplateResponse) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *UserNameplateResponse) SetLabel(v string) {
	o.Label.Set(&v)
}

// SetLabelNil sets the value for Label to be an explicit nil
func (o *UserNameplateResponse) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *UserNameplateResponse) UnsetLabel() {
	o.Label.Unset()
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *UserNameplateResponse) GetPalette() string {
	if o == nil || IsNil(o.Palette) {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNameplateResponse) GetPaletteOk() (*string, bool) {
	if o == nil || IsNil(o.Palette) {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *UserNameplateResponse) HasPalette() bool {
	if o != nil && !IsNil(o.Palette) {
		return true
	}

	return false
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *UserNameplateResponse) SetPalette(v string) {
	o.Palette = &v
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
	if o.Asset.IsSet() {
		toSerialize["asset"] = o.Asset.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if !IsNil(o.Palette) {
		toSerialize["palette"] = o.Palette
	}
	return toSerialize, nil
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


