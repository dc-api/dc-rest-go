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
)

// checks if the BulkUpdateGuildChannelsRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateGuildChannelsRequestInner{}

// BulkUpdateGuildChannelsRequestInner struct for BulkUpdateGuildChannelsRequestInner
type BulkUpdateGuildChannelsRequestInner struct {
	Id *GetEntitlementsSkuIdsParameterOneOfInner `json:"id,omitempty"`
	Position *int32 `json:"position,omitempty"`
	ParentId *GetEntitlementsSkuIdsParameterOneOfInner `json:"parent_id,omitempty"`
	LockPermissions *bool `json:"lock_permissions,omitempty"`
}

// NewBulkUpdateGuildChannelsRequestInner instantiates a new BulkUpdateGuildChannelsRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateGuildChannelsRequestInner() *BulkUpdateGuildChannelsRequestInner {
	this := BulkUpdateGuildChannelsRequestInner{}
	return &this
}

// NewBulkUpdateGuildChannelsRequestInnerWithDefaults instantiates a new BulkUpdateGuildChannelsRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateGuildChannelsRequestInnerWithDefaults() *BulkUpdateGuildChannelsRequestInner {
	this := BulkUpdateGuildChannelsRequestInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkUpdateGuildChannelsRequestInner) GetId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.Id) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateGuildChannelsRequestInner) GetIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkUpdateGuildChannelsRequestInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the Id field.
func (o *BulkUpdateGuildChannelsRequestInner) SetId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.Id = &v
}


// GetPosition returns the Position field value if set, zero value otherwise.
func (o *BulkUpdateGuildChannelsRequestInner) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateGuildChannelsRequestInner) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *BulkUpdateGuildChannelsRequestInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *BulkUpdateGuildChannelsRequestInner) SetPosition(v int32) {
	o.Position = &v
}


// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BulkUpdateGuildChannelsRequestInner) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.ParentId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateGuildChannelsRequestInner) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BulkUpdateGuildChannelsRequestInner) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the ParentId field.
func (o *BulkUpdateGuildChannelsRequestInner) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.ParentId = &v
}


// GetLockPermissions returns the LockPermissions field value if set, zero value otherwise.
func (o *BulkUpdateGuildChannelsRequestInner) GetLockPermissions() bool {
	if o == nil || IsNil(o.LockPermissions) {
		var ret bool
		return ret
	}
	return *o.LockPermissions
}

// GetLockPermissionsOk returns a tuple with the LockPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateGuildChannelsRequestInner) GetLockPermissionsOk() (*bool, bool) {
	if o == nil || IsNil(o.LockPermissions) {
		return nil, false
	}
	return o.LockPermissions, true
}

// HasLockPermissions returns a boolean if a field has been set.
func (o *BulkUpdateGuildChannelsRequestInner) HasLockPermissions() bool {
	if o != nil && !IsNil(o.LockPermissions) {
		return true
	}

	return false
}

// SetLockPermissions gets a reference to the given bool and assigns it to the LockPermissions field.
func (o *BulkUpdateGuildChannelsRequestInner) SetLockPermissions(v bool) {
	o.LockPermissions = &v
}


func (o BulkUpdateGuildChannelsRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpdateGuildChannelsRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.LockPermissions) {
		toSerialize["lock_permissions"] = o.LockPermissions
	}
	return toSerialize, nil
}

type NullableBulkUpdateGuildChannelsRequestInner struct {
	value *BulkUpdateGuildChannelsRequestInner
	isSet bool
}

func (v NullableBulkUpdateGuildChannelsRequestInner) Get() *BulkUpdateGuildChannelsRequestInner {
	return v.value
}

func (v *NullableBulkUpdateGuildChannelsRequestInner) Set(val *BulkUpdateGuildChannelsRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateGuildChannelsRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateGuildChannelsRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateGuildChannelsRequestInner(val *BulkUpdateGuildChannelsRequestInner) *NullableBulkUpdateGuildChannelsRequestInner {
	return &NullableBulkUpdateGuildChannelsRequestInner{value: val, isSet: true}
}

func (v NullableBulkUpdateGuildChannelsRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateGuildChannelsRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

