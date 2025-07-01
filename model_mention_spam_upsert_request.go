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

// checks if the MentionSpamUpsertRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionSpamUpsertRequest{}

// MentionSpamUpsertRequest struct for MentionSpamUpsertRequest
type MentionSpamUpsertRequest struct {
	Name string `json:"name"`
	EventType int32 `json:"event_type"`
	Actions []DefaultKeywordListUpsertRequestActionsInner `json:"actions,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	ExemptRoles []string `json:"exempt_roles,omitempty"`
	ExemptChannels []string `json:"exempt_channels,omitempty"`
	TriggerType int32 `json:"trigger_type"`
	TriggerMetadata NullableMentionSpamTriggerMetadata `json:"trigger_metadata,omitempty"`
}

type _MentionSpamUpsertRequest MentionSpamUpsertRequest

// NewMentionSpamUpsertRequest instantiates a new MentionSpamUpsertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionSpamUpsertRequest(name string, eventType int32, triggerType int32) *MentionSpamUpsertRequest {
	this := MentionSpamUpsertRequest{}
	this.Name = name
	this.EventType = eventType
	this.TriggerType = triggerType
	return &this
}

// NewMentionSpamUpsertRequestWithDefaults instantiates a new MentionSpamUpsertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionSpamUpsertRequestWithDefaults() *MentionSpamUpsertRequest {
	this := MentionSpamUpsertRequest{}
	return &this
}

// GetName returns the Name field value
func (o *MentionSpamUpsertRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MentionSpamUpsertRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MentionSpamUpsertRequest) SetName(v string) {
	o.Name = v
}

// GetEventType returns the EventType field value
func (o *MentionSpamUpsertRequest) GetEventType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MentionSpamUpsertRequest) GetEventTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MentionSpamUpsertRequest) SetEventType(v int32) {
	o.EventType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamUpsertRequest) GetActions() []DefaultKeywordListUpsertRequestActionsInner {
	if o == nil {
		var ret []DefaultKeywordListUpsertRequestActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamUpsertRequest) GetActionsOk() ([]DefaultKeywordListUpsertRequestActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *MentionSpamUpsertRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []DefaultKeywordListUpsertRequestActionsInner and assigns it to the Actions field.
func (o *MentionSpamUpsertRequest) SetActions(v []DefaultKeywordListUpsertRequestActionsInner) {
	o.Actions = v
}


// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamUpsertRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamUpsertRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled.Get()) {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *MentionSpamUpsertRequest) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *MentionSpamUpsertRequest) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *MentionSpamUpsertRequest) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *MentionSpamUpsertRequest) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamUpsertRequest) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamUpsertRequest) GetExemptRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *MentionSpamUpsertRequest) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *MentionSpamUpsertRequest) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}


// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamUpsertRequest) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamUpsertRequest) GetExemptChannelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *MentionSpamUpsertRequest) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *MentionSpamUpsertRequest) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}


// GetTriggerType returns the TriggerType field value
func (o *MentionSpamUpsertRequest) GetTriggerType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *MentionSpamUpsertRequest) GetTriggerTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *MentionSpamUpsertRequest) SetTriggerType(v int32) {
	o.TriggerType = v
}

// GetTriggerMetadata returns the TriggerMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamUpsertRequest) GetTriggerMetadata() MentionSpamTriggerMetadata {
	if o == nil || IsNil(o.TriggerMetadata.Get()) {
		var ret MentionSpamTriggerMetadata
		return ret
	}
	return *o.TriggerMetadata.Get()
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamUpsertRequest) GetTriggerMetadataOk() (*MentionSpamTriggerMetadata, bool) {
	if o == nil || IsNil(o.TriggerMetadata.Get()) {
		return nil, false
	}
	return o.TriggerMetadata.Get(), o.TriggerMetadata.IsSet()
}

// HasTriggerMetadata returns a boolean if a field has been set.
func (o *MentionSpamUpsertRequest) HasTriggerMetadata() bool {
	if o != nil && o.TriggerMetadata.IsSet() {
		return true
	}

	return false
}

// SetTriggerMetadata gets a reference to the given NullableMentionSpamTriggerMetadata and assigns it to the TriggerMetadata field.
func (o *MentionSpamUpsertRequest) SetTriggerMetadata(v MentionSpamTriggerMetadata) {
	o.TriggerMetadata.Set(&v)
}

// SetTriggerMetadataNil sets the value for TriggerMetadata to be an explicit nil
func (o *MentionSpamUpsertRequest) SetTriggerMetadataNil() {
	o.TriggerMetadata.Set(nil)
}

// UnsetTriggerMetadata ensures that no value is present for TriggerMetadata, not even an explicit nil
func (o *MentionSpamUpsertRequest) UnsetTriggerMetadata() {
	o.TriggerMetadata.Unset()
}

func (o MentionSpamUpsertRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionSpamUpsertRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["event_type"] = o.EventType
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
	toSerialize["trigger_type"] = o.TriggerType
	if o.TriggerMetadata.IsSet() {
		toSerialize["trigger_metadata"] = o.TriggerMetadata.Get()
	}
	return toSerialize, nil
}

func (o *MentionSpamUpsertRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"event_type",
		"trigger_type",
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

	varMentionSpamUpsertRequest := _MentionSpamUpsertRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMentionSpamUpsertRequest)

	if err != nil {
		return err
	}

	*o = MentionSpamUpsertRequest(varMentionSpamUpsertRequest)

	return err
}

type NullableMentionSpamUpsertRequest struct {
	value *MentionSpamUpsertRequest
	isSet bool
}

func (v NullableMentionSpamUpsertRequest) Get() *MentionSpamUpsertRequest {
	return v.value
}

func (v *NullableMentionSpamUpsertRequest) Set(val *MentionSpamUpsertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionSpamUpsertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionSpamUpsertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionSpamUpsertRequest(val *MentionSpamUpsertRequest) *NullableMentionSpamUpsertRequest {
	return &NullableMentionSpamUpsertRequest{value: val, isSet: true}
}

func (v NullableMentionSpamUpsertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionSpamUpsertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


