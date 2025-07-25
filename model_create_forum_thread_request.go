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

// checks if the CreateForumThreadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateForumThreadRequest{}

// CreateForumThreadRequest struct for CreateForumThreadRequest
type CreateForumThreadRequest struct {
	Name string `json:"name"`
	AutoArchiveDuration *int32 `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser NullableInt32 `json:"rate_limit_per_user,omitempty"`
	AppliedTags []string `json:"applied_tags,omitempty"`
	Message BaseCreateMessageCreateRequest `json:"message"`
}

type _CreateForumThreadRequest CreateForumThreadRequest

// NewCreateForumThreadRequest instantiates a new CreateForumThreadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateForumThreadRequest(name string, message BaseCreateMessageCreateRequest) *CreateForumThreadRequest {
	this := CreateForumThreadRequest{}
	this.Name = name
	this.Message = message
	return &this
}

// NewCreateForumThreadRequestWithDefaults instantiates a new CreateForumThreadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateForumThreadRequestWithDefaults() *CreateForumThreadRequest {
	this := CreateForumThreadRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateForumThreadRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateForumThreadRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateForumThreadRequest) SetName(v string) {
	o.Name = v
}

// GetAutoArchiveDuration returns the AutoArchiveDuration field value if set, zero value otherwise.
func (o *CreateForumThreadRequest) GetAutoArchiveDuration() int32 {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		var ret int32
		return ret
	}
	return *o.AutoArchiveDuration
}

// GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateForumThreadRequest) GetAutoArchiveDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoArchiveDuration) {
		return nil, false
	}
	return o.AutoArchiveDuration, true
}

// HasAutoArchiveDuration returns a boolean if a field has been set.
func (o *CreateForumThreadRequest) HasAutoArchiveDuration() bool {
	if o != nil && !IsNil(o.AutoArchiveDuration) {
		return true
	}

	return false
}

// SetAutoArchiveDuration gets a reference to the given int32 and assigns it to the AutoArchiveDuration field.
func (o *CreateForumThreadRequest) SetAutoArchiveDuration(v int32) {
	o.AutoArchiveDuration = &v
}


// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateForumThreadRequest) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser.Get()
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateForumThreadRequest) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		return nil, false
	}
	return o.RateLimitPerUser.Get(), o.RateLimitPerUser.IsSet()
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateForumThreadRequest) HasRateLimitPerUser() bool {
	if o != nil && o.RateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the RateLimitPerUser field.
func (o *CreateForumThreadRequest) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser.Set(&v)
}

// SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil
func (o *CreateForumThreadRequest) SetRateLimitPerUserNil() {
	o.RateLimitPerUser.Set(nil)
}

// UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
func (o *CreateForumThreadRequest) UnsetRateLimitPerUser() {
	o.RateLimitPerUser.Unset()
}

// GetAppliedTags returns the AppliedTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateForumThreadRequest) GetAppliedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AppliedTags
}

// GetAppliedTagsOk returns a tuple with the AppliedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateForumThreadRequest) GetAppliedTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppliedTags, true
}

// HasAppliedTags returns a boolean if a field has been set.
func (o *CreateForumThreadRequest) HasAppliedTags() bool {
	if o != nil && !IsNil(o.AppliedTags) {
		return true
	}

	return false
}

// SetAppliedTags gets a reference to the given []string and assigns it to the AppliedTags field.
func (o *CreateForumThreadRequest) SetAppliedTags(v []string) {
	o.AppliedTags = v
}


// GetMessage returns the Message field value
func (o *CreateForumThreadRequest) GetMessage() BaseCreateMessageCreateRequest {
	if o == nil {
		var ret BaseCreateMessageCreateRequest
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateForumThreadRequest) GetMessageOk() (*BaseCreateMessageCreateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateForumThreadRequest) SetMessage(v BaseCreateMessageCreateRequest) {
	o.Message = v
}

func (o CreateForumThreadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateForumThreadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.AutoArchiveDuration) {
		toSerialize["auto_archive_duration"] = o.AutoArchiveDuration
	}
	if o.RateLimitPerUser.IsSet() {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser.Get()
	}
	if o.AppliedTags != nil {
		toSerialize["applied_tags"] = o.AppliedTags
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *CreateForumThreadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"message",
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

	varCreateForumThreadRequest := _CreateForumThreadRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateForumThreadRequest)

	if err != nil {
		return err
	}

	*o = CreateForumThreadRequest(varCreateForumThreadRequest)

	return err
}

type NullableCreateForumThreadRequest struct {
	value *CreateForumThreadRequest
	isSet bool
}

func (v NullableCreateForumThreadRequest) Get() *CreateForumThreadRequest {
	return v.value
}

func (v *NullableCreateForumThreadRequest) Set(val *CreateForumThreadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateForumThreadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateForumThreadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateForumThreadRequest(val *CreateForumThreadRequest) *NullableCreateForumThreadRequest {
	return &NullableCreateForumThreadRequest{value: val, isSet: true}
}

func (v NullableCreateForumThreadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateForumThreadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


