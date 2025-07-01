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
	"bytes"
	"fmt"
)

// checks if the ThumbnailComponentForMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThumbnailComponentForMessageRequest{}

// ThumbnailComponentForMessageRequest struct for ThumbnailComponentForMessageRequest
type ThumbnailComponentForMessageRequest struct {
	Type int32 `json:"type"`
	Description NullableString `json:"description,omitempty"`
	Spoiler NullableBool `json:"spoiler,omitempty"`
	Media UnfurledMediaRequest `json:"media"`
}

type _ThumbnailComponentForMessageRequest ThumbnailComponentForMessageRequest

// NewThumbnailComponentForMessageRequest instantiates a new ThumbnailComponentForMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbnailComponentForMessageRequest(type_ int32, media UnfurledMediaRequest) *ThumbnailComponentForMessageRequest {
	this := ThumbnailComponentForMessageRequest{}
	this.Type = type_
	this.Media = media
	return &this
}

// NewThumbnailComponentForMessageRequestWithDefaults instantiates a new ThumbnailComponentForMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbnailComponentForMessageRequestWithDefaults() *ThumbnailComponentForMessageRequest {
	this := ThumbnailComponentForMessageRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ThumbnailComponentForMessageRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ThumbnailComponentForMessageRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ThumbnailComponentForMessageRequest) SetType(v int32) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailComponentForMessageRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailComponentForMessageRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ThumbnailComponentForMessageRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ThumbnailComponentForMessageRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ThumbnailComponentForMessageRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ThumbnailComponentForMessageRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetSpoiler returns the Spoiler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThumbnailComponentForMessageRequest) GetSpoiler() bool {
	if o == nil || IsNil(o.Spoiler.Get()) {
		var ret bool
		return ret
	}
	return *o.Spoiler.Get()
}

// GetSpoilerOk returns a tuple with the Spoiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThumbnailComponentForMessageRequest) GetSpoilerOk() (*bool, bool) {
	if o == nil || IsNil(o.Spoiler.Get()) {
		return nil, false
	}
	return o.Spoiler.Get(), o.Spoiler.IsSet()
}

// HasSpoiler returns a boolean if a field has been set.
func (o *ThumbnailComponentForMessageRequest) HasSpoiler() bool {
	if o != nil && o.Spoiler.IsSet() {
		return true
	}

	return false
}

// SetSpoiler gets a reference to the given NullableBool and assigns it to the Spoiler field.
func (o *ThumbnailComponentForMessageRequest) SetSpoiler(v bool) {
	o.Spoiler.Set(&v)
}

// SetSpoilerNil sets the value for Spoiler to be an explicit nil
func (o *ThumbnailComponentForMessageRequest) SetSpoilerNil() {
	o.Spoiler.Set(nil)
}

// UnsetSpoiler ensures that no value is present for Spoiler, not even an explicit nil
func (o *ThumbnailComponentForMessageRequest) UnsetSpoiler() {
	o.Spoiler.Unset()
}

// GetMedia returns the Media field value
func (o *ThumbnailComponentForMessageRequest) GetMedia() UnfurledMediaRequest {
	if o == nil {
		var ret UnfurledMediaRequest
		return ret
	}

	return o.Media
}

// GetMediaOk returns a tuple with the Media field value
// and a boolean to check if the value has been set.
func (o *ThumbnailComponentForMessageRequest) GetMediaOk() (*UnfurledMediaRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Media, true
}

// SetMedia sets field value
func (o *ThumbnailComponentForMessageRequest) SetMedia(v UnfurledMediaRequest) {
	o.Media = v
}

func (o ThumbnailComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThumbnailComponentForMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Spoiler.IsSet() {
		toSerialize["spoiler"] = o.Spoiler.Get()
	}
	toSerialize["media"] = o.Media
	return toSerialize, nil
}

func (o *ThumbnailComponentForMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"media",
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

	varThumbnailComponentForMessageRequest := _ThumbnailComponentForMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThumbnailComponentForMessageRequest)

	if err != nil {
		return err
	}

	*o = ThumbnailComponentForMessageRequest(varThumbnailComponentForMessageRequest)

	return err
}

type NullableThumbnailComponentForMessageRequest struct {
	value *ThumbnailComponentForMessageRequest
	isSet bool
}

func (v NullableThumbnailComponentForMessageRequest) Get() *ThumbnailComponentForMessageRequest {
	return v.value
}

func (v *NullableThumbnailComponentForMessageRequest) Set(val *ThumbnailComponentForMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbnailComponentForMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbnailComponentForMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbnailComponentForMessageRequest(val *ThumbnailComponentForMessageRequest) *NullableThumbnailComponentForMessageRequest {
	return &NullableThumbnailComponentForMessageRequest{value: val, isSet: true}
}

func (v NullableThumbnailComponentForMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbnailComponentForMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


