/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the MLSpamRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLSpamRuleResponse{}

// MLSpamRuleResponse struct for MLSpamRuleResponse
type MLSpamRuleResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	CreatorId string `json:"creator_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	EventType int32 `json:"event_type"`
	Actions []DefaultKeywordRuleResponseActionsInner `json:"actions"`
	TriggerType int32 `json:"trigger_type"`
	Enabled NullableBool `json:"enabled,omitempty"`
	ExemptRoles []string `json:"exempt_roles,omitempty"`
	ExemptChannels []string `json:"exempt_channels,omitempty"`
	TriggerMetadata map[string]interface{} `json:"trigger_metadata"`
}

type _MLSpamRuleResponse MLSpamRuleResponse

// NewMLSpamRuleResponse instantiates a new MLSpamRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLSpamRuleResponse(id string, guildId string, creatorId string, name string, eventType int32, actions []DefaultKeywordRuleResponseActionsInner, triggerType int32, triggerMetadata map[string]interface{}) *MLSpamRuleResponse {
	this := MLSpamRuleResponse{}
	this.Id = id
	this.GuildId = guildId
	this.CreatorId = creatorId
	this.Name = name
	this.EventType = eventType
	this.Actions = actions
	this.TriggerType = triggerType
	this.TriggerMetadata = triggerMetadata
	return &this
}

// NewMLSpamRuleResponseWithDefaults instantiates a new MLSpamRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLSpamRuleResponseWithDefaults() *MLSpamRuleResponse {
	this := MLSpamRuleResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MLSpamRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MLSpamRuleResponse) SetId(v string) {
	o.Id = v
}

// GetGuildId returns the GuildId field value
func (o *MLSpamRuleResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *MLSpamRuleResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetCreatorId returns the CreatorId field value
func (o *MLSpamRuleResponse) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *MLSpamRuleResponse) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetName returns the Name field value
func (o *MLSpamRuleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MLSpamRuleResponse) SetName(v string) {
	o.Name = v
}

// GetEventType returns the EventType field value
func (o *MLSpamRuleResponse) GetEventType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetEventTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MLSpamRuleResponse) SetEventType(v int32) {
	o.EventType = v
}

// GetActions returns the Actions field value
func (o *MLSpamRuleResponse) GetActions() []DefaultKeywordRuleResponseActionsInner {
	if o == nil {
		var ret []DefaultKeywordRuleResponseActionsInner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetActionsOk() ([]DefaultKeywordRuleResponseActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *MLSpamRuleResponse) SetActions(v []DefaultKeywordRuleResponseActionsInner) {
	o.Actions = v
}

// GetTriggerType returns the TriggerType field value
func (o *MLSpamRuleResponse) GetTriggerType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetTriggerTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *MLSpamRuleResponse) SetTriggerType(v int32) {
	o.TriggerType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MLSpamRuleResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MLSpamRuleResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled.Get()) {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *MLSpamRuleResponse) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *MLSpamRuleResponse) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *MLSpamRuleResponse) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *MLSpamRuleResponse) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MLSpamRuleResponse) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MLSpamRuleResponse) GetExemptRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *MLSpamRuleResponse) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *MLSpamRuleResponse) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}


// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MLSpamRuleResponse) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MLSpamRuleResponse) GetExemptChannelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *MLSpamRuleResponse) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *MLSpamRuleResponse) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}


// GetTriggerMetadata returns the TriggerMetadata field value
func (o *MLSpamRuleResponse) GetTriggerMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.TriggerMetadata
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value
// and a boolean to check if the value has been set.
func (o *MLSpamRuleResponse) GetTriggerMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.TriggerMetadata, true
}

// SetTriggerMetadata sets field value
func (o *MLSpamRuleResponse) SetTriggerMetadata(v map[string]interface{}) {
	o.TriggerMetadata = v
}

func (o MLSpamRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLSpamRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["guild_id"] = o.GuildId
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["name"] = o.Name
	toSerialize["event_type"] = o.EventType
	toSerialize["actions"] = o.Actions
	toSerialize["trigger_type"] = o.TriggerType
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.ExemptRoles != nil {
		toSerialize["exempt_roles"] = o.ExemptRoles
	}
	if o.ExemptChannels != nil {
		toSerialize["exempt_channels"] = o.ExemptChannels
	}
	toSerialize["trigger_metadata"] = o.TriggerMetadata
	return toSerialize, nil
}

func (o *MLSpamRuleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"guild_id",
		"creator_id",
		"name",
		"event_type",
		"actions",
		"trigger_type",
		"trigger_metadata",
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

	varMLSpamRuleResponse := _MLSpamRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMLSpamRuleResponse)

	if err != nil {
		return err
	}

	*o = MLSpamRuleResponse(varMLSpamRuleResponse)

	return err
}

type NullableMLSpamRuleResponse struct {
	value *MLSpamRuleResponse
	isSet bool
}

func (v NullableMLSpamRuleResponse) Get() *MLSpamRuleResponse {
	return v.value
}

func (v *NullableMLSpamRuleResponse) Set(val *MLSpamRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMLSpamRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMLSpamRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLSpamRuleResponse(val *MLSpamRuleResponse) *NullableMLSpamRuleResponse {
	return &NullableMLSpamRuleResponse{value: val, isSet: true}
}

func (v NullableMLSpamRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLSpamRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


