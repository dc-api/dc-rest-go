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
	"bytes"
	"fmt"
)

// checks if the CreateTextThreadWithoutMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTextThreadWithoutMessageRequest{}

// CreateTextThreadWithoutMessageRequest struct for CreateTextThreadWithoutMessageRequest
type CreateTextThreadWithoutMessageRequest struct {
	Name string `json:"name"`
	AutoArchiveDuration *int32 `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser NullableInt32 `json:"rate_limit_per_user,omitempty"`
	Type NullableInt32 `json:"type,omitempty"`
	Invitable NullableBool `json:"invitable,omitempty"`
}

type _CreateTextThreadWithoutMessageRequest CreateTextThreadWithoutMessageRequest

// NewCreateTextThreadWithoutMessageRequest instantiates a new CreateTextThreadWithoutMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTextThreadWithoutMessageRequest(name string) *CreateTextThreadWithoutMessageRequest {
	this := CreateTextThreadWithoutMessageRequest{}
	this.Name = name
	return &this
}

// NewCreateTextThreadWithoutMessageRequestWithDefaults instantiates a new CreateTextThreadWithoutMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTextThreadWithoutMessageRequestWithDefaults() *CreateTextThreadWithoutMessageRequest {
	this := CreateTextThreadWithoutMessageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTextThreadWithoutMessageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTextThreadWithoutMessageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTextThreadWithoutMessageRequest) SetName(v string) {
	o.Name = v
}

// GetAutoArchiveDuration returns the AutoArchiveDuration field value if set, zero value otherwise.
func (o *CreateTextThreadWithoutMessageRequest) GetAutoArchiveDuration() int32 {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		var ret int32
		return ret
	}
	return *o.AutoArchiveDuration
}

// GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTextThreadWithoutMessageRequest) GetAutoArchiveDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		return nil, false
	}
	return o.AutoArchiveDuration, true
}

// HasAutoArchiveDuration returns a boolean if a field has been set.
func (o *CreateTextThreadWithoutMessageRequest) HasAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.AutoArchiveDuration) {
		return true
	}

	return false
}

// SetAutoArchiveDuration gets a reference to the given int32 and assigns it to the AutoArchiveDuration field.
func (o *CreateTextThreadWithoutMessageRequest) SetAutoArchiveDuration(v int32) {
	o.AutoArchiveDuration = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTextThreadWithoutMessageRequest) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser.Get()
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTextThreadWithoutMessageRequest) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		return nil, false
	}
	return o.RateLimitPerUser.Get(), o.RateLimitPerUser.IsSet()
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateTextThreadWithoutMessageRequest) HasRateLimitPerUser() bool {
	if o != nil && o.RateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the RateLimitPerUser field.
func (o *CreateTextThreadWithoutMessageRequest) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser.Set(&v)
}

// SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil
func (o *CreateTextThreadWithoutMessageRequest) SetRateLimitPerUserNil() {
	o.RateLimitPerUser.Set(nil)
}

// UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
func (o *CreateTextThreadWithoutMessageRequest) UnsetRateLimitPerUser() {
	o.RateLimitPerUser.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTextThreadWithoutMessageRequest) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTextThreadWithoutMessageRequest) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CreateTextThreadWithoutMessageRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *CreateTextThreadWithoutMessageRequest) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateTextThreadWithoutMessageRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateTextThreadWithoutMessageRequest) UnsetType() {
	o.Type.Unset()
}

// GetInvitable returns the Invitable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTextThreadWithoutMessageRequest) GetInvitable() bool {
	if o == nil || IsNil(o.Invitable.Get()) {
		var ret bool
		return ret
	}
	return *o.Invitable.Get()
}

// GetInvitableOk returns a tuple with the Invitable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTextThreadWithoutMessageRequest) GetInvitableOk() (*bool, bool) {
	if o == nil || IsNil(o.Invitable.Get()) {
		return nil, false
	}
	return o.Invitable.Get(), o.Invitable.IsSet()
}

// HasInvitable returns a boolean if a field has been set.
func (o *CreateTextThreadWithoutMessageRequest) HasInvitable() bool {
	if o != nil && o.Invitable.IsSet() {
		return true
	}

	return false
}

// SetInvitable gets a reference to the given NullableBool and assigns it to the Invitable field.
func (o *CreateTextThreadWithoutMessageRequest) SetInvitable(v bool) {
	o.Invitable.Set(&v)
}

// SetInvitableNil sets the value for Invitable to be an explicit nil
func (o *CreateTextThreadWithoutMessageRequest) SetInvitableNil() {
	o.Invitable.Set(nil)
}

// UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil
func (o *CreateTextThreadWithoutMessageRequest) UnsetInvitable() {
	o.Invitable.Unset()
}

func (o CreateTextThreadWithoutMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTextThreadWithoutMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.AutoArchiveDuration) {
		toSerialize["auto_archive_duration"] = o.AutoArchiveDuration
	}
	if o.RateLimitPerUser.IsSet() {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Invitable.IsSet() {
		toSerialize["invitable"] = o.Invitable.Get()
	}
	return toSerialize, nil
}

func (o *CreateTextThreadWithoutMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateTextThreadWithoutMessageRequest := _CreateTextThreadWithoutMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTextThreadWithoutMessageRequest)

	if err != nil {
		return err
	}

	*o = CreateTextThreadWithoutMessageRequest(varCreateTextThreadWithoutMessageRequest)

	return err
}

type NullableCreateTextThreadWithoutMessageRequest struct {
	value *CreateTextThreadWithoutMessageRequest
	isSet bool
}

func (v NullableCreateTextThreadWithoutMessageRequest) Get() *CreateTextThreadWithoutMessageRequest {
	return v.value
}

func (v *NullableCreateTextThreadWithoutMessageRequest) Set(val *CreateTextThreadWithoutMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTextThreadWithoutMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTextThreadWithoutMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTextThreadWithoutMessageRequest(val *CreateTextThreadWithoutMessageRequest) *NullableCreateTextThreadWithoutMessageRequest {
	return &NullableCreateTextThreadWithoutMessageRequest{value: val, isSet: true}
}

func (v NullableCreateTextThreadWithoutMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTextThreadWithoutMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


