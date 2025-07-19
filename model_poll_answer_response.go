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

// checks if the PollAnswerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollAnswerResponse{}

// PollAnswerResponse struct for PollAnswerResponse
type PollAnswerResponse struct {
	AnswerId int32 `json:"answer_id"`
	PollMedia PollMediaResponse `json:"poll_media"`
}

type _PollAnswerResponse PollAnswerResponse

// NewPollAnswerResponse instantiates a new PollAnswerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollAnswerResponse(answerId int32, pollMedia PollMediaResponse) *PollAnswerResponse {
	this := PollAnswerResponse{}
	this.AnswerId = answerId
	this.PollMedia = pollMedia
	return &this
}

// NewPollAnswerResponseWithDefaults instantiates a new PollAnswerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollAnswerResponseWithDefaults() *PollAnswerResponse {
	this := PollAnswerResponse{}
	return &this
}

// GetAnswerId returns the AnswerId field value
func (o *PollAnswerResponse) GetAnswerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnswerId
}

// GetAnswerIdOk returns a tuple with the AnswerId field value
// and a boolean to check if the value has been set.
func (o *PollAnswerResponse) GetAnswerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerId, true
}

// SetAnswerId sets field value
func (o *PollAnswerResponse) SetAnswerId(v int32) {
	o.AnswerId = v
}

// GetPollMedia returns the PollMedia field value
func (o *PollAnswerResponse) GetPollMedia() PollMediaResponse {
	if o == nil {
		var ret PollMediaResponse
		return ret
	}

	return o.PollMedia
}

// GetPollMediaOk returns a tuple with the PollMedia field value
// and a boolean to check if the value has been set.
func (o *PollAnswerResponse) GetPollMediaOk() (*PollMediaResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollMedia, true
}

// SetPollMedia sets field value
func (o *PollAnswerResponse) SetPollMedia(v PollMediaResponse) {
	o.PollMedia = v
}

func (o PollAnswerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollAnswerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answer_id"] = o.AnswerId
	toSerialize["poll_media"] = o.PollMedia
	return toSerialize, nil
}

func (o *PollAnswerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"answer_id",
		"poll_media",
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

	varPollAnswerResponse := _PollAnswerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollAnswerResponse)

	if err != nil {
		return err
	}

	*o = PollAnswerResponse(varPollAnswerResponse)

	return err
}

type NullablePollAnswerResponse struct {
	value *PollAnswerResponse
	isSet bool
}

func (v NullablePollAnswerResponse) Get() *PollAnswerResponse {
	return v.value
}

func (v *NullablePollAnswerResponse) Set(val *PollAnswerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePollAnswerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePollAnswerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollAnswerResponse(val *PollAnswerResponse) *NullablePollAnswerResponse {
	return &NullablePollAnswerResponse{value: val, isSet: true}
}

func (v NullablePollAnswerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollAnswerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


