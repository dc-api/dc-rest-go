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

// checks if the AuditLogEntryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogEntryResponse{}

// AuditLogEntryResponse struct for AuditLogEntryResponse
type AuditLogEntryResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ActionType NullableInt32 `json:"action_type"`
	UserId *string `json:"user_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	TargetId *string `json:"target_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Changes []AuditLogObjectChangeResponse `json:"changes,omitempty"`
	Options map[string]string `json:"options,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
}

type _AuditLogEntryResponse AuditLogEntryResponse

// NewAuditLogEntryResponse instantiates a new AuditLogEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogEntryResponse(id string, actionType NullableInt32) *AuditLogEntryResponse {
	this := AuditLogEntryResponse{}
	this.Id = id
	this.ActionType = actionType
	return &this
}

// NewAuditLogEntryResponseWithDefaults instantiates a new AuditLogEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogEntryResponseWithDefaults() *AuditLogEntryResponse {
	this := AuditLogEntryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AuditLogEntryResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntryResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuditLogEntryResponse) SetId(v string) {
	o.Id = v
}

// GetActionType returns the ActionType field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AuditLogEntryResponse) GetActionType() int32 {
	if o == nil || o.ActionType.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ActionType.Get()
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogEntryResponse) GetActionTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionType.Get(), o.ActionType.IsSet()
}

// SetActionType sets field value
func (o *AuditLogEntryResponse) SetActionType(v int32) {
	o.ActionType.Set(&v)
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AuditLogEntryResponse) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntryResponse) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AuditLogEntryResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AuditLogEntryResponse) SetUserId(v string) {
	o.UserId = &v
}


// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AuditLogEntryResponse) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntryResponse) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AuditLogEntryResponse) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AuditLogEntryResponse) SetTargetId(v string) {
	o.TargetId = &v
}


// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLogEntryResponse) GetChanges() []AuditLogObjectChangeResponse {
	if o == nil {
		var ret []AuditLogObjectChangeResponse
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogEntryResponse) GetChangesOk() ([]AuditLogObjectChangeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *AuditLogEntryResponse) HasChanges() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []AuditLogObjectChangeResponse and assigns it to the Changes field.
func (o *AuditLogEntryResponse) SetChanges(v []AuditLogObjectChangeResponse) {
	o.Changes = v
}


// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AuditLogEntryResponse) GetOptions() map[string]string {
	if o == nil || IsNil(o.Options) {
		var ret map[string]string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntryResponse) GetOptionsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]string{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AuditLogEntryResponse) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]string and assigns it to the Options field.
func (o *AuditLogEntryResponse) SetOptions(v map[string]string) {
	o.Options = v
}


// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLogEntryResponse) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogEntryResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason.Get()) {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *AuditLogEntryResponse) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *AuditLogEntryResponse) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *AuditLogEntryResponse) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *AuditLogEntryResponse) UnsetReason() {
	o.Reason.Unset()
}

func (o AuditLogEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogEntryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["action_type"] = o.ActionType.Get()
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	return toSerialize, nil
}

func (o *AuditLogEntryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"action_type",
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

	varAuditLogEntryResponse := _AuditLogEntryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditLogEntryResponse)

	if err != nil {
		return err
	}

	*o = AuditLogEntryResponse(varAuditLogEntryResponse)

	return err
}

type NullableAuditLogEntryResponse struct {
	value *AuditLogEntryResponse
	isSet bool
}

func (v NullableAuditLogEntryResponse) Get() *AuditLogEntryResponse {
	return v.value
}

func (v *NullableAuditLogEntryResponse) Set(val *AuditLogEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogEntryResponse(val *AuditLogEntryResponse) *NullableAuditLogEntryResponse {
	return &NullableAuditLogEntryResponse{value: val, isSet: true}
}

func (v NullableAuditLogEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


