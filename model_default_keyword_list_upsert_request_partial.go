/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// checks if the DefaultKeywordListUpsertRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultKeywordListUpsertRequestPartial{}

// DefaultKeywordListUpsertRequestPartial struct for DefaultKeywordListUpsertRequestPartial
type DefaultKeywordListUpsertRequestPartial struct {
	Name *string `json:"name,omitempty"`
	EventType *int32 `json:"event_type,omitempty"`
	Actions []DefaultKeywordListUpsertRequestActionsInner `json:"actions,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	ExemptRoles []string `json:"exempt_roles,omitempty"`
	ExemptChannels []string `json:"exempt_channels,omitempty"`
	TriggerType *int32 `json:"trigger_type,omitempty"`
	TriggerMetadata *DefaultKeywordListTriggerMetadata `json:"trigger_metadata,omitempty"`
}

// NewDefaultKeywordListUpsertRequestPartial instantiates a new DefaultKeywordListUpsertRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultKeywordListUpsertRequestPartial() *DefaultKeywordListUpsertRequestPartial {
	this := DefaultKeywordListUpsertRequestPartial{}
	return &this
}

// NewDefaultKeywordListUpsertRequestPartialWithDefaults instantiates a new DefaultKeywordListUpsertRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultKeywordListUpsertRequestPartialWithDefaults() *DefaultKeywordListUpsertRequestPartial {
	this := DefaultKeywordListUpsertRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DefaultKeywordListUpsertRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DefaultKeywordListUpsertRequestPartial) SetName(v string) {
	o.Name = &v
}


// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DefaultKeywordListUpsertRequestPartial) GetEventType() int32 {
	if o == nil || IsNil(o.EventType) {
		var ret int32
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequestPartial) GetEventTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given int32 and assigns it to the EventType field.
func (o *DefaultKeywordListUpsertRequestPartial) SetEventType(v int32) {
	o.EventType = &v
}


// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequestPartial) GetActions() []DefaultKeywordListUpsertRequestActionsInner {
	if o == nil {
		var ret []DefaultKeywordListUpsertRequestActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequestPartial) GetActionsOk() ([]DefaultKeywordListUpsertRequestActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []DefaultKeywordListUpsertRequestActionsInner and assigns it to the Actions field.
func (o *DefaultKeywordListUpsertRequestPartial) SetActions(v []DefaultKeywordListUpsertRequestActionsInner) {
	o.Actions = v
}


// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequestPartial) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequestPartial) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled.Get()) {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DefaultKeywordListUpsertRequestPartial) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DefaultKeywordListUpsertRequestPartial) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DefaultKeywordListUpsertRequestPartial) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequestPartial) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequestPartial) GetExemptRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *DefaultKeywordListUpsertRequestPartial) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}


// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequestPartial) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequestPartial) GetExemptChannelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *DefaultKeywordListUpsertRequestPartial) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}


// GetTriggerType returns the TriggerType field value if set, zero value otherwise.
func (o *DefaultKeywordListUpsertRequestPartial) GetTriggerType() int32 {
	if o == nil || IsNil(o.TriggerType) {
		var ret int32
		return ret
	}
	return *o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequestPartial) GetTriggerTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TriggerType) {
		return nil, false
	}
	return o.TriggerType, true
}

// HasTriggerType returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasTriggerType() bool {
	if o != nil && !IsNil(o.TriggerType) {
		return true
	}

	return false
}

// SetTriggerType gets a reference to the given int32 and assigns it to the TriggerType field.
func (o *DefaultKeywordListUpsertRequestPartial) SetTriggerType(v int32) {
	o.TriggerType = &v
}


// GetTriggerMetadata returns the TriggerMetadata field value if set, zero value otherwise.
func (o *DefaultKeywordListUpsertRequestPartial) GetTriggerMetadata() DefaultKeywordListTriggerMetadata {
	if o == nil || IsNil(o.TriggerMetadata) {
		var ret DefaultKeywordListTriggerMetadata
		return ret
	}
	return *o.TriggerMetadata
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequestPartial) GetTriggerMetadataOk() (*DefaultKeywordListTriggerMetadata, bool) {
	if o == nil || IsNil(o.TriggerMetadata) {
		return nil, false
	}
	return o.TriggerMetadata, true
}

// HasTriggerMetadata returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequestPartial) HasTriggerMetadata() bool {
	if o != nil && !IsNil(o.TriggerMetadata) {
		return true
	}

	return false
}

// SetTriggerMetadata gets a reference to the given DefaultKeywordListTriggerMetadata and assigns it to the TriggerMetadata field.
func (o *DefaultKeywordListUpsertRequestPartial) SetTriggerMetadata(v DefaultKeywordListTriggerMetadata) {
	o.TriggerMetadata = &v
}


func (o DefaultKeywordListUpsertRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultKeywordListUpsertRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.ExemptRoles != nil {
		toSerialize["exempt_roles"] = o.ExemptRoles
	}
	if o.ExemptChannels != nil {
		toSerialize["exempt_channels"] = o.ExemptChannels
	}
	if !IsNil(o.TriggerType) {
		toSerialize["trigger_type"] = o.TriggerType
	}
	if !IsNil(o.TriggerMetadata) {
		toSerialize["trigger_metadata"] = o.TriggerMetadata
	}
	return toSerialize, nil
}

type NullableDefaultKeywordListUpsertRequestPartial struct {
	value *DefaultKeywordListUpsertRequestPartial
	isSet bool
}

func (v NullableDefaultKeywordListUpsertRequestPartial) Get() *DefaultKeywordListUpsertRequestPartial {
	return v.value
}

func (v *NullableDefaultKeywordListUpsertRequestPartial) Set(val *DefaultKeywordListUpsertRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListUpsertRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListUpsertRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListUpsertRequestPartial(val *DefaultKeywordListUpsertRequestPartial) *NullableDefaultKeywordListUpsertRequestPartial {
	return &NullableDefaultKeywordListUpsertRequestPartial{value: val, isSet: true}
}

func (v NullableDefaultKeywordListUpsertRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListUpsertRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


