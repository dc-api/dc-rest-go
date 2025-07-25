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

// checks if the PollResultsEntryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollResultsEntryResponse{}

// PollResultsEntryResponse struct for PollResultsEntryResponse
type PollResultsEntryResponse struct {
	Id int32 `json:"id"`
	Count int32 `json:"count"`
	MeVoted NullableBool `json:"me_voted,omitempty"`
}

type _PollResultsEntryResponse PollResultsEntryResponse

// NewPollResultsEntryResponse instantiates a new PollResultsEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollResultsEntryResponse(id int32, count int32) *PollResultsEntryResponse {
	this := PollResultsEntryResponse{}
	this.Id = id
	this.Count = count
	return &this
}

// NewPollResultsEntryResponseWithDefaults instantiates a new PollResultsEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollResultsEntryResponseWithDefaults() *PollResultsEntryResponse {
	this := PollResultsEntryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PollResultsEntryResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PollResultsEntryResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PollResultsEntryResponse) SetId(v int32) {
	o.Id = v
}

// GetCount returns the Count field value
func (o *PollResultsEntryResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PollResultsEntryResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PollResultsEntryResponse) SetCount(v int32) {
	o.Count = v
}

// GetMeVoted returns the MeVoted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PollResultsEntryResponse) GetMeVoted() bool {
	if o == nil || IsNil(o.MeVoted.Get()) {
		var ret bool
		return ret
	}
	return *o.MeVoted.Get()
}

// GetMeVotedOk returns a tuple with the MeVoted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PollResultsEntryResponse) GetMeVotedOk() (*bool, bool) {
	if o == nil || IsNil(o.MeVoted.Get()) {
		return nil, false
	}
	return o.MeVoted.Get(), o.MeVoted.IsSet()
}

// HasMeVoted returns a boolean if a field has been set.
func (o *PollResultsEntryResponse) HasMeVoted() bool {
	if o != nil && o.MeVoted.IsSet() {
		return true
	}

	return false
}

// SetMeVoted gets a reference to the given NullableBool and assigns it to the MeVoted field.
func (o *PollResultsEntryResponse) SetMeVoted(v bool) {
	o.MeVoted.Set(&v)
}

// SetMeVotedNil sets the value for MeVoted to be an explicit nil
func (o *PollResultsEntryResponse) SetMeVotedNil() {
	o.MeVoted.Set(nil)
}

// UnsetMeVoted ensures that no value is present for MeVoted, not even an explicit nil
func (o *PollResultsEntryResponse) UnsetMeVoted() {
	o.MeVoted.Unset()
}

func (o PollResultsEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollResultsEntryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["count"] = o.Count
	if o.MeVoted.IsSet() {
		toSerialize["me_voted"] = o.MeVoted.Get()
	}
	return toSerialize, nil
}

func (o *PollResultsEntryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"count",
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

	varPollResultsEntryResponse := _PollResultsEntryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollResultsEntryResponse)

	if err != nil {
		return err
	}

	*o = PollResultsEntryResponse(varPollResultsEntryResponse)

	return err
}

type NullablePollResultsEntryResponse struct {
	value *PollResultsEntryResponse
	isSet bool
}

func (v NullablePollResultsEntryResponse) Get() *PollResultsEntryResponse {
	return v.value
}

func (v *NullablePollResultsEntryResponse) Set(val *PollResultsEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePollResultsEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePollResultsEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollResultsEntryResponse(val *PollResultsEntryResponse) *NullablePollResultsEntryResponse {
	return &NullablePollResultsEntryResponse{value: val, isSet: true}
}

func (v NullablePollResultsEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollResultsEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


