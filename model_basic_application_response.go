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
	"bytes"
	"fmt"
)

// checks if the BasicApplicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicApplicationResponse{}

// BasicApplicationResponse struct for BasicApplicationResponse
type BasicApplicationResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon *string `json:"icon,omitempty"`
	Description string `json:"description"`
	Type *ApplicationFormPartialType `json:"type,omitempty"`
	CoverImage *string `json:"cover_image,omitempty"`
	PrimarySkuId *string `json:"primary_sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Bot *UserResponse `json:"bot,omitempty"`
}

type _BasicApplicationResponse BasicApplicationResponse

// NewBasicApplicationResponse instantiates a new BasicApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicApplicationResponse(id string, name string, description string) *BasicApplicationResponse {
	this := BasicApplicationResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewBasicApplicationResponseWithDefaults instantiates a new BasicApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicApplicationResponseWithDefaults() *BasicApplicationResponse {
	this := BasicApplicationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *BasicApplicationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BasicApplicationResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BasicApplicationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BasicApplicationResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *BasicApplicationResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *BasicApplicationResponse) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *BasicApplicationResponse) SetIcon(v string) {
	o.Icon = &v
}


// GetDescription returns the Description field value
func (o *BasicApplicationResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BasicApplicationResponse) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BasicApplicationResponse) GetType() ApplicationFormPartialType {
	if o == nil || IsNil(o.Type) {
		var ret ApplicationFormPartialType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetTypeOk() (*ApplicationFormPartialType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BasicApplicationResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ApplicationFormPartialType and assigns it to the Type field.
func (o *BasicApplicationResponse) SetType(v ApplicationFormPartialType) {
	o.Type = &v
}


// GetCoverImage returns the CoverImage field value if set, zero value otherwise.
func (o *BasicApplicationResponse) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage) {
		var ret string
		return ret
	}
	return *o.CoverImage
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage) {
		return nil, false
	}
	return o.CoverImage, true
}

// HasCoverImage returns a boolean if a field has been set.
func (o *BasicApplicationResponse) HasCoverImage() bool {
	if o != nil && !IsNil(o.CoverImage) {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given string and assigns it to the CoverImage field.
func (o *BasicApplicationResponse) SetCoverImage(v string) {
	o.CoverImage = &v
}


// GetPrimarySkuId returns the PrimarySkuId field value if set, zero value otherwise.
func (o *BasicApplicationResponse) GetPrimarySkuId() string {
	if o == nil || IsNil(o.PrimarySkuId) {
		var ret string
		return ret
	}
	return *o.PrimarySkuId
}

// GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetPrimarySkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimarySkuId) {
		return nil, false
	}
	return o.PrimarySkuId, true
}

// HasPrimarySkuId returns a boolean if a field has been set.
func (o *BasicApplicationResponse) HasPrimarySkuId() bool {
	if o != nil && !IsNil(o.PrimarySkuId) {
		return true
	}

	return false
}

// SetPrimarySkuId gets a reference to the given string and assigns it to the PrimarySkuId field.
func (o *BasicApplicationResponse) SetPrimarySkuId(v string) {
	o.PrimarySkuId = &v
}


// GetBot returns the Bot field value if set, zero value otherwise.
func (o *BasicApplicationResponse) GetBot() UserResponse {
	if o == nil || IsNil(o.Bot) {
		var ret UserResponse
		return ret
	}
	return *o.Bot
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicApplicationResponse) GetBotOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Bot) {
		return nil, false
	}
	return o.Bot, true
}

// HasBot returns a boolean if a field has been set.
func (o *BasicApplicationResponse) HasBot() bool {
	if o != nil && !IsNil(o.Bot) {
		return true
	}

	return false
}

// SetBot gets a reference to the given UserResponse and assigns it to the Bot field.
func (o *BasicApplicationResponse) SetBot(v UserResponse) {
	o.Bot = &v
}


func (o BasicApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicApplicationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CoverImage) {
		toSerialize["cover_image"] = o.CoverImage
	}
	if !IsNil(o.PrimarySkuId) {
		toSerialize["primary_sku_id"] = o.PrimarySkuId
	}
	if !IsNil(o.Bot) {
		toSerialize["bot"] = o.Bot
	}
	return toSerialize, nil
}

func (o *BasicApplicationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBasicApplicationResponse := _BasicApplicationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBasicApplicationResponse)

	if err != nil {
		return err
	}

	*o = BasicApplicationResponse(varBasicApplicationResponse)

	return err
}

type NullableBasicApplicationResponse struct {
	value *BasicApplicationResponse
	isSet bool
}

func (v NullableBasicApplicationResponse) Get() *BasicApplicationResponse {
	return v.value
}

func (v *NullableBasicApplicationResponse) Set(val *BasicApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicApplicationResponse(val *BasicApplicationResponse) *NullableBasicApplicationResponse {
	return &NullableBasicApplicationResponse{value: val, isSet: true}
}

func (v NullableBasicApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

