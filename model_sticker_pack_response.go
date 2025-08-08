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

// checks if the StickerPackResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StickerPackResponse{}

// StickerPackResponse struct for StickerPackResponse
type StickerPackResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	SkuId string `json:"sku_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Stickers []StandardStickerResponse `json:"stickers"`
	CoverStickerId *string `json:"cover_sticker_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	BannerAssetId *string `json:"banner_asset_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _StickerPackResponse StickerPackResponse

// NewStickerPackResponse instantiates a new StickerPackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStickerPackResponse(id string, skuId string, name string, stickers []StandardStickerResponse) *StickerPackResponse {
	this := StickerPackResponse{}
	this.Id = id
	this.SkuId = skuId
	this.Name = name
	this.Stickers = stickers
	return &this
}

// NewStickerPackResponseWithDefaults instantiates a new StickerPackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStickerPackResponseWithDefaults() *StickerPackResponse {
	this := StickerPackResponse{}
	return &this
}

// GetId returns the Id field value
func (o *StickerPackResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StickerPackResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StickerPackResponse) SetId(v string) {
	o.Id = v
}

// GetSkuId returns the SkuId field value
func (o *StickerPackResponse) GetSkuId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value
// and a boolean to check if the value has been set.
func (o *StickerPackResponse) GetSkuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuId, true
}

// SetSkuId sets field value
func (o *StickerPackResponse) SetSkuId(v string) {
	o.SkuId = v
}

// GetName returns the Name field value
func (o *StickerPackResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StickerPackResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StickerPackResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StickerPackResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StickerPackResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StickerPackResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StickerPackResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StickerPackResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StickerPackResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetStickers returns the Stickers field value
func (o *StickerPackResponse) GetStickers() []StandardStickerResponse {
	if o == nil {
		var ret []StandardStickerResponse
		return ret
	}

	return o.Stickers
}

// GetStickersOk returns a tuple with the Stickers field value
// and a boolean to check if the value has been set.
func (o *StickerPackResponse) GetStickersOk() ([]StandardStickerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stickers, true
}

// SetStickers sets field value
func (o *StickerPackResponse) SetStickers(v []StandardStickerResponse) {
	o.Stickers = v
}

// GetCoverStickerId returns the CoverStickerId field value if set, zero value otherwise.
func (o *StickerPackResponse) GetCoverStickerId() string {
	if o == nil || IsNil(o.CoverStickerId) {
		var ret string
		return ret
	}
	return *o.CoverStickerId
}

// GetCoverStickerIdOk returns a tuple with the CoverStickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StickerPackResponse) GetCoverStickerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoverStickerId) {
		return nil, false
	}
	return o.CoverStickerId, true
}

// HasCoverStickerId returns a boolean if a field has been set.
func (o *StickerPackResponse) HasCoverStickerId() bool {
	if o != nil && !IsNil(o.CoverStickerId) {
		return true
	}

	return false
}

// SetCoverStickerId gets a reference to the given string and assigns it to the CoverStickerId field.
func (o *StickerPackResponse) SetCoverStickerId(v string) {
	o.CoverStickerId = &v
}


// GetBannerAssetId returns the BannerAssetId field value if set, zero value otherwise.
func (o *StickerPackResponse) GetBannerAssetId() string {
	if o == nil || IsNil(o.BannerAssetId) {
		var ret string
		return ret
	}
	return *o.BannerAssetId
}

// GetBannerAssetIdOk returns a tuple with the BannerAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StickerPackResponse) GetBannerAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.BannerAssetId) {
		return nil, false
	}
	return o.BannerAssetId, true
}

// HasBannerAssetId returns a boolean if a field has been set.
func (o *StickerPackResponse) HasBannerAssetId() bool {
	if o != nil && !IsNil(o.BannerAssetId) {
		return true
	}

	return false
}

// SetBannerAssetId gets a reference to the given string and assigns it to the BannerAssetId field.
func (o *StickerPackResponse) SetBannerAssetId(v string) {
	o.BannerAssetId = &v
}


func (o StickerPackResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StickerPackResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["sku_id"] = o.SkuId
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["stickers"] = o.Stickers
	if !IsNil(o.CoverStickerId) {
		toSerialize["cover_sticker_id"] = o.CoverStickerId
	}
	if !IsNil(o.BannerAssetId) {
		toSerialize["banner_asset_id"] = o.BannerAssetId
	}
	return toSerialize, nil
}

func (o *StickerPackResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"sku_id",
		"name",
		"stickers",
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

	varStickerPackResponse := _StickerPackResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStickerPackResponse)

	if err != nil {
		return err
	}

	*o = StickerPackResponse(varStickerPackResponse)

	return err
}

type NullableStickerPackResponse struct {
	value *StickerPackResponse
	isSet bool
}

func (v NullableStickerPackResponse) Get() *StickerPackResponse {
	return v.value
}

func (v *NullableStickerPackResponse) Set(val *StickerPackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStickerPackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStickerPackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStickerPackResponse(val *StickerPackResponse) *NullableStickerPackResponse {
	return &NullableStickerPackResponse{value: val, isSet: true}
}

func (v NullableStickerPackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStickerPackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


