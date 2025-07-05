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

// checks if the PruneGuildRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PruneGuildRequest{}

// PruneGuildRequest struct for PruneGuildRequest
type PruneGuildRequest struct {
	Days NullableInt32 `json:"days,omitempty"`
	ComputePruneCount NullableBool `json:"compute_prune_count,omitempty"`
	IncludeRoles NullablePruneGuildRequestIncludeRoles `json:"include_roles,omitempty"`
}

// NewPruneGuildRequest instantiates a new PruneGuildRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPruneGuildRequest() *PruneGuildRequest {
	this := PruneGuildRequest{}
	return &this
}

// NewPruneGuildRequestWithDefaults instantiates a new PruneGuildRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPruneGuildRequestWithDefaults() *PruneGuildRequest {
	this := PruneGuildRequest{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PruneGuildRequest) GetDays() int32 {
	if o == nil || IsNil(o.Days.Get()) {
		var ret int32
		return ret
	}
	return *o.Days.Get()
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PruneGuildRequest) GetDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Days.Get()) {
		return nil, false
	}
	return o.Days.Get(), o.Days.IsSet()
}

// HasDays returns a boolean if a field has been set.
func (o *PruneGuildRequest) HasDays() bool {
	if o != nil && o.Days.IsSet() {
		return true
	}

	return false
}

// SetDays gets a reference to the given NullableInt32 and assigns it to the Days field.
func (o *PruneGuildRequest) SetDays(v int32) {
	o.Days.Set(&v)
}

// SetDaysNil sets the value for Days to be an explicit nil
func (o *PruneGuildRequest) SetDaysNil() {
	o.Days.Set(nil)
}

// UnsetDays ensures that no value is present for Days, not even an explicit nil
func (o *PruneGuildRequest) UnsetDays() {
	o.Days.Unset()
}

// GetComputePruneCount returns the ComputePruneCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PruneGuildRequest) GetComputePruneCount() bool {
	if o == nil || IsNil(o.ComputePruneCount.Get()) {
		var ret bool
		return ret
	}
	return *o.ComputePruneCount.Get()
}

// GetComputePruneCountOk returns a tuple with the ComputePruneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PruneGuildRequest) GetComputePruneCountOk() (*bool, bool) {
	if o == nil || IsNil(o.ComputePruneCount.Get()) {
		return nil, false
	}
	return o.ComputePruneCount.Get(), o.ComputePruneCount.IsSet()
}

// HasComputePruneCount returns a boolean if a field has been set.
func (o *PruneGuildRequest) HasComputePruneCount() bool {
	if o != nil && o.ComputePruneCount.IsSet() {
		return true
	}

	return false
}

// SetComputePruneCount gets a reference to the given NullableBool and assigns it to the ComputePruneCount field.
func (o *PruneGuildRequest) SetComputePruneCount(v bool) {
	o.ComputePruneCount.Set(&v)
}

// SetComputePruneCountNil sets the value for ComputePruneCount to be an explicit nil
func (o *PruneGuildRequest) SetComputePruneCountNil() {
	o.ComputePruneCount.Set(nil)
}

// UnsetComputePruneCount ensures that no value is present for ComputePruneCount, not even an explicit nil
func (o *PruneGuildRequest) UnsetComputePruneCount() {
	o.ComputePruneCount.Unset()
}

// GetIncludeRoles returns the IncludeRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PruneGuildRequest) GetIncludeRoles() PruneGuildRequestIncludeRoles {
	if o == nil || IsNil(o.IncludeRoles.Get()) {
		var ret PruneGuildRequestIncludeRoles
		return ret
	}
	return *o.IncludeRoles.Get()
}

// GetIncludeRolesOk returns a tuple with the IncludeRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PruneGuildRequest) GetIncludeRolesOk() (*PruneGuildRequestIncludeRoles, bool) {
	if o == nil || IsNil(o.IncludeRoles.Get()) {
		return nil, false
	}
	return o.IncludeRoles.Get(), o.IncludeRoles.IsSet()
}

// HasIncludeRoles returns a boolean if a field has been set.
func (o *PruneGuildRequest) HasIncludeRoles() bool {
	if o != nil && o.IncludeRoles.IsSet() {
		return true
	}

	return false
}

// SetIncludeRoles gets a reference to the given NullablePruneGuildRequestIncludeRoles and assigns it to the IncludeRoles field.
func (o *PruneGuildRequest) SetIncludeRoles(v PruneGuildRequestIncludeRoles) {
	o.IncludeRoles.Set(&v)
}

// SetIncludeRolesNil sets the value for IncludeRoles to be an explicit nil
func (o *PruneGuildRequest) SetIncludeRolesNil() {
	o.IncludeRoles.Set(nil)
}

// UnsetIncludeRoles ensures that no value is present for IncludeRoles, not even an explicit nil
func (o *PruneGuildRequest) UnsetIncludeRoles() {
	o.IncludeRoles.Unset()
}

func (o PruneGuildRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PruneGuildRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Days.IsSet() {
		toSerialize["days"] = o.Days.Get()
	}
	if o.ComputePruneCount.IsSet() {
		toSerialize["compute_prune_count"] = o.ComputePruneCount.Get()
	}
	if o.IncludeRoles.IsSet() {
		toSerialize["include_roles"] = o.IncludeRoles.Get()
	}
	return toSerialize, nil
}

type NullablePruneGuildRequest struct {
	value *PruneGuildRequest
	isSet bool
}

func (v NullablePruneGuildRequest) Get() *PruneGuildRequest {
	return v.value
}

func (v *NullablePruneGuildRequest) Set(val *PruneGuildRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePruneGuildRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePruneGuildRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePruneGuildRequest(val *PruneGuildRequest) *NullablePruneGuildRequest {
	return &NullablePruneGuildRequest{value: val, isSet: true}
}

func (v NullablePruneGuildRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePruneGuildRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


