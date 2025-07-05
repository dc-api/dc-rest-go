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
)

// checks if the RichEmbedImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RichEmbedImage{}

// RichEmbedImage struct for RichEmbedImage
type RichEmbedImage struct {
	Url NullableString `json:"url,omitempty"`
	Width NullableInt32 `json:"width,omitempty"`
	Height NullableInt32 `json:"height,omitempty"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	PlaceholderVersion NullableInt32 `json:"placeholder_version,omitempty"`
	IsAnimated NullableBool `json:"is_animated,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewRichEmbedImage instantiates a new RichEmbedImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRichEmbedImage() *RichEmbedImage {
	this := RichEmbedImage{}
	return &this
}

// NewRichEmbedImageWithDefaults instantiates a new RichEmbedImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRichEmbedImageWithDefaults() *RichEmbedImage {
	this := RichEmbedImage{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url.Get()) {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RichEmbedImage) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RichEmbedImage) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *RichEmbedImage) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RichEmbedImage) UnsetUrl() {
	o.Url.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width.Get()) {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *RichEmbedImage) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *RichEmbedImage) SetWidth(v int32) {
	o.Width.Set(&v)
}

// SetWidthNil sets the value for Width to be an explicit nil
func (o *RichEmbedImage) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *RichEmbedImage) UnsetWidth() {
	o.Width.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height.Get()) {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *RichEmbedImage) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *RichEmbedImage) SetHeight(v int32) {
	o.Height.Set(&v)
}

// SetHeightNil sets the value for Height to be an explicit nil
func (o *RichEmbedImage) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *RichEmbedImage) UnsetHeight() {
	o.Height.Unset()
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder.Get()) {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *RichEmbedImage) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *RichEmbedImage) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}

// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *RichEmbedImage) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *RichEmbedImage) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetPlaceholderVersion returns the PlaceholderVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetPlaceholderVersion() int32 {
	if o == nil || IsNil(o.PlaceholderVersion.Get()) {
		var ret int32
		return ret
	}
	return *o.PlaceholderVersion.Get()
}

// GetPlaceholderVersionOk returns a tuple with the PlaceholderVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetPlaceholderVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.PlaceholderVersion.Get()) {
		return nil, false
	}
	return o.PlaceholderVersion.Get(), o.PlaceholderVersion.IsSet()
}

// HasPlaceholderVersion returns a boolean if a field has been set.
func (o *RichEmbedImage) HasPlaceholderVersion() bool {
	if o != nil && o.PlaceholderVersion.IsSet() {
		return true
	}

	return false
}

// SetPlaceholderVersion gets a reference to the given NullableInt32 and assigns it to the PlaceholderVersion field.
func (o *RichEmbedImage) SetPlaceholderVersion(v int32) {
	o.PlaceholderVersion.Set(&v)
}

// SetPlaceholderVersionNil sets the value for PlaceholderVersion to be an explicit nil
func (o *RichEmbedImage) SetPlaceholderVersionNil() {
	o.PlaceholderVersion.Set(nil)
}

// UnsetPlaceholderVersion ensures that no value is present for PlaceholderVersion, not even an explicit nil
func (o *RichEmbedImage) UnsetPlaceholderVersion() {
	o.PlaceholderVersion.Unset()
}

// GetIsAnimated returns the IsAnimated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetIsAnimated() bool {
	if o == nil || IsNil(o.IsAnimated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAnimated.Get()
}

// GetIsAnimatedOk returns a tuple with the IsAnimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetIsAnimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnimated.Get()) {
		return nil, false
	}
	return o.IsAnimated.Get(), o.IsAnimated.IsSet()
}

// HasIsAnimated returns a boolean if a field has been set.
func (o *RichEmbedImage) HasIsAnimated() bool {
	if o != nil && o.IsAnimated.IsSet() {
		return true
	}

	return false
}

// SetIsAnimated gets a reference to the given NullableBool and assigns it to the IsAnimated field.
func (o *RichEmbedImage) SetIsAnimated(v bool) {
	o.IsAnimated.Set(&v)
}

// SetIsAnimatedNil sets the value for IsAnimated to be an explicit nil
func (o *RichEmbedImage) SetIsAnimatedNil() {
	o.IsAnimated.Set(nil)
}

// UnsetIsAnimated ensures that no value is present for IsAnimated, not even an explicit nil
func (o *RichEmbedImage) UnsetIsAnimated() {
	o.IsAnimated.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedImage) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedImage) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RichEmbedImage) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RichEmbedImage) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RichEmbedImage) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RichEmbedImage) UnsetDescription() {
	o.Description.Unset()
}

func (o RichEmbedImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RichEmbedImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.PlaceholderVersion.IsSet() {
		toSerialize["placeholder_version"] = o.PlaceholderVersion.Get()
	}
	if o.IsAnimated.IsSet() {
		toSerialize["is_animated"] = o.IsAnimated.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableRichEmbedImage struct {
	value *RichEmbedImage
	isSet bool
}

func (v NullableRichEmbedImage) Get() *RichEmbedImage {
	return v.value
}

func (v *NullableRichEmbedImage) Set(val *RichEmbedImage) {
	v.value = val
	v.isSet = true
}

func (v NullableRichEmbedImage) IsSet() bool {
	return v.isSet
}

func (v *NullableRichEmbedImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRichEmbedImage(val *RichEmbedImage) *NullableRichEmbedImage {
	return &NullableRichEmbedImage{value: val, isSet: true}
}

func (v NullableRichEmbedImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRichEmbedImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


