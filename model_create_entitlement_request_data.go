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

// checks if the CreateEntitlementRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEntitlementRequestData{}

// CreateEntitlementRequestData struct for CreateEntitlementRequestData
type CreateEntitlementRequestData struct {
	SkuId string `json:"sku_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	OwnerId string `json:"owner_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	OwnerType int32 `json:"owner_type"`
}

type _CreateEntitlementRequestData CreateEntitlementRequestData

// NewCreateEntitlementRequestData instantiates a new CreateEntitlementRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEntitlementRequestData(skuId string, ownerId string, ownerType int32) *CreateEntitlementRequestData {
	this := CreateEntitlementRequestData{}
	this.SkuId = skuId
	this.OwnerId = ownerId
	this.OwnerType = ownerType
	return &this
}

// NewCreateEntitlementRequestDataWithDefaults instantiates a new CreateEntitlementRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEntitlementRequestDataWithDefaults() *CreateEntitlementRequestData {
	this := CreateEntitlementRequestData{}
	return &this
}

// GetSkuId returns the SkuId field value
func (o *CreateEntitlementRequestData) GetSkuId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value
// and a boolean to check if the value has been set.
func (o *CreateEntitlementRequestData) GetSkuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuId, true
}

// SetSkuId sets field value
func (o *CreateEntitlementRequestData) SetSkuId(v string) {
	o.SkuId = v
}

// GetOwnerId returns the OwnerId field value
func (o *CreateEntitlementRequestData) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *CreateEntitlementRequestData) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *CreateEntitlementRequestData) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerType returns the OwnerType field value
func (o *CreateEntitlementRequestData) GetOwnerType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *CreateEntitlementRequestData) GetOwnerTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value
func (o *CreateEntitlementRequestData) SetOwnerType(v int32) {
	o.OwnerType = v
}

func (o CreateEntitlementRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEntitlementRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku_id"] = o.SkuId
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_type"] = o.OwnerType
	return toSerialize, nil
}

func (o *CreateEntitlementRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku_id",
		"owner_id",
		"owner_type",
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

	varCreateEntitlementRequestData := _CreateEntitlementRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEntitlementRequestData)

	if err != nil {
		return err
	}

	*o = CreateEntitlementRequestData(varCreateEntitlementRequestData)

	return err
}

type NullableCreateEntitlementRequestData struct {
	value *CreateEntitlementRequestData
	isSet bool
}

func (v NullableCreateEntitlementRequestData) Get() *CreateEntitlementRequestData {
	return v.value
}

func (v *NullableCreateEntitlementRequestData) Set(val *CreateEntitlementRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEntitlementRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEntitlementRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEntitlementRequestData(val *CreateEntitlementRequestData) *NullableCreateEntitlementRequestData {
	return &NullableCreateEntitlementRequestData{value: val, isSet: true}
}

func (v NullableCreateEntitlementRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEntitlementRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


