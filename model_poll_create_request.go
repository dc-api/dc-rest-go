/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the PollCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollCreateRequest{}

// PollCreateRequest struct for PollCreateRequest
type PollCreateRequest struct {
	Question PollMedia `json:"question"`
	Answers []PollAnswerCreateRequest `json:"answers"`
	AllowMultiselect NullableBool `json:"allow_multiselect,omitempty"`
	LayoutType *int32 `json:"layout_type,omitempty"`
	Duration NullableInt32 `json:"duration,omitempty"`
}

type _PollCreateRequest PollCreateRequest

// NewPollCreateRequest instantiates a new PollCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollCreateRequest(question PollMedia, answers []PollAnswerCreateRequest) *PollCreateRequest {
	this := PollCreateRequest{}
	this.Question = question
	this.Answers = answers
	return &this
}

// NewPollCreateRequestWithDefaults instantiates a new PollCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollCreateRequestWithDefaults() *PollCreateRequest {
	this := PollCreateRequest{}
	return &this
}

// GetQuestion returns the Question field value
func (o *PollCreateRequest) GetQuestion() PollMedia {
	if o == nil {
		var ret PollMedia
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *PollCreateRequest) GetQuestionOk() (*PollMedia, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *PollCreateRequest) SetQuestion(v PollMedia) {
	o.Question = v
}

// GetAnswers returns the Answers field value
func (o *PollCreateRequest) GetAnswers() []PollAnswerCreateRequest {
	if o == nil {
		var ret []PollAnswerCreateRequest
		return ret
	}

	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value
// and a boolean to check if the value has been set.
func (o *PollCreateRequest) GetAnswersOk() ([]PollAnswerCreateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Answers, true
}

// SetAnswers sets field value
func (o *PollCreateRequest) SetAnswers(v []PollAnswerCreateRequest) {
	o.Answers = v
}

// GetAllowMultiselect returns the AllowMultiselect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollCreateRequest) GetAllowMultiselect() bool {
	if o == nil || IsNil(o.AllowMultiselect.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowMultiselect.Get()
}

// GetAllowMultiselectOk returns a tuple with the AllowMultiselect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollCreateRequest) GetAllowMultiselectOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMultiselect.Get()) {
		return nil, false
	}
	return o.AllowMultiselect.Get(), o.AllowMultiselect.IsSet()
}

// HasAllowMultiselect returns a boolean if a field has been set.
func (o *PollCreateRequest) HasAllowMultiselect() bool {
	if o != nil && o.AllowMultiselect.IsSet() {
		return true
	}

	return false
}

// SetAllowMultiselect gets a reference to the given NullableBool and assigns it to the AllowMultiselect field.
func (o *PollCreateRequest) SetAllowMultiselect(v bool) {
	o.AllowMultiselect.Set(&v)
}

// SetAllowMultiselectNil sets the value for AllowMultiselect to be an explicit nil
func (o *PollCreateRequest) SetAllowMultiselectNil() {
	o.AllowMultiselect.Set(nil)
}

// UnsetAllowMultiselect ensures that no value is present for AllowMultiselect, not even an explicit nil
func (o *PollCreateRequest) UnsetAllowMultiselect() {
	o.AllowMultiselect.Unset()
}

// GetLayoutType returns the LayoutType field value if set, zero value otherwise.
func (o *PollCreateRequest) GetLayoutType() int32 {
	if o == nil || IsNil(o.LayoutType) {
		var ret int32
		return ret
	}
	return *o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollCreateRequest) GetLayoutTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.LayoutType) {
		return nil, false
	}
	return o.LayoutType, true
}

// HasLayoutType returns a boolean if a field has been set.
func (o *PollCreateRequest) HasLayoutType() bool {
	if o != nil && !IsNil(o.LayoutType) {
		return true
	}

	return false
}

// SetLayoutType gets a reference to the given int32 and assigns it to the LayoutType field.
func (o *PollCreateRequest) SetLayoutType(v int32) {
	o.LayoutType = &v
}


// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollCreateRequest) GetDuration() int32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollCreateRequest) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration.Get()) {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *PollCreateRequest) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *PollCreateRequest) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil
func (o *PollCreateRequest) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *PollCreateRequest) UnsetDuration() {
	o.Duration.Unset()
}

func (o PollCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["question"] = o.Question
	toSerialize["answers"] = o.Answers
	if o.AllowMultiselect.IsSet() {
		toSerialize["allow_multiselect"] = o.AllowMultiselect.Get()
	}
	if !IsNil(o.LayoutType) {
		toSerialize["layout_type"] = o.LayoutType
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	return toSerialize, nil
}

func (o *PollCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"question",
		"answers",
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

	varPollCreateRequest := _PollCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollCreateRequest)

	if err != nil {
		return err
	}

	*o = PollCreateRequest(varPollCreateRequest)

	return err
}

type NullablePollCreateRequest struct {
	value *PollCreateRequest
	isSet bool
}

func (v NullablePollCreateRequest) Get() *PollCreateRequest {
	return v.value
}

func (v *NullablePollCreateRequest) Set(val *PollCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePollCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePollCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollCreateRequest(val *PollCreateRequest) *NullablePollCreateRequest {
	return &NullablePollCreateRequest{value: val, isSet: true}
}

func (v NullablePollCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


