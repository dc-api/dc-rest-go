/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-28T06:15:04.313783603Z[Etc/UTC]
 *    * - **Generator Version**: 7.19.0
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

// checks if the ThreadSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadSearchResponse{}

// ThreadSearchResponse struct for ThreadSearchResponse
type ThreadSearchResponse struct {
	Threads []ThreadResponse `json:"threads"`
	Members []ThreadMemberResponse `json:"members"`
	HasMore bool `json:"has_more"`
	FirstMessages []MessageResponse `json:"first_messages,omitempty"`
	TotalResults int32 `json:"total_results"`
}

type _ThreadSearchResponse ThreadSearchResponse

// NewThreadSearchResponse instantiates a new ThreadSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadSearchResponse(threads []ThreadResponse, members []ThreadMemberResponse, hasMore bool, totalResults int32) *ThreadSearchResponse {
	this := ThreadSearchResponse{}
	this.Threads = threads
	this.Members = members
	this.HasMore = hasMore
	this.TotalResults = totalResults
	return &this
}

// NewThreadSearchResponseWithDefaults instantiates a new ThreadSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadSearchResponseWithDefaults() *ThreadSearchResponse {
	this := ThreadSearchResponse{}
	return &this
}

// GetThreads returns the Threads field value
func (o *ThreadSearchResponse) GetThreads() []ThreadResponse {
	if o == nil {
		var ret []ThreadResponse
		return ret
	}

	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value
// and a boolean to check if the value has been set.
func (o *ThreadSearchResponse) GetThreadsOk() ([]ThreadResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads, true
}

// SetThreads sets field value
func (o *ThreadSearchResponse) SetThreads(v []ThreadResponse) {
	o.Threads = v
}

// GetMembers returns the Members field value
func (o *ThreadSearchResponse) GetMembers() []ThreadMemberResponse {
	if o == nil {
		var ret []ThreadMemberResponse
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ThreadSearchResponse) GetMembersOk() ([]ThreadMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *ThreadSearchResponse) SetMembers(v []ThreadMemberResponse) {
	o.Members = v
}

// GetHasMore returns the HasMore field value
func (o *ThreadSearchResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ThreadSearchResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ThreadSearchResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetFirstMessages returns the FirstMessages field value if set, zero value otherwise.
func (o *ThreadSearchResponse) GetFirstMessages() []MessageResponse {
	if o == nil || IsNil(o.FirstMessages) {
		var ret []MessageResponse
		return ret
	}
	return o.FirstMessages
}

// GetFirstMessagesOk returns a tuple with the FirstMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadSearchResponse) GetFirstMessagesOk() ([]MessageResponse, bool) {
	if o == nil || IsNil(o.FirstMessages) {
		return nil, false
	}
	return o.FirstMessages, true
}

// HasFirstMessages returns a boolean if a field has been set.
func (o *ThreadSearchResponse) HasFirstMessages() bool {
	if o != nil && !IsNil(o.FirstMessages) {
		return true
	}

	return false
}

// SetFirstMessages gets a reference to the given []MessageResponse and assigns it to the FirstMessages field.
func (o *ThreadSearchResponse) SetFirstMessages(v []MessageResponse) {
	o.FirstMessages = v
}


// GetTotalResults returns the TotalResults field value
func (o *ThreadSearchResponse) GetTotalResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value
// and a boolean to check if the value has been set.
func (o *ThreadSearchResponse) GetTotalResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalResults, true
}

// SetTotalResults sets field value
func (o *ThreadSearchResponse) SetTotalResults(v int32) {
	o.TotalResults = v
}

func (o ThreadSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threads"] = o.Threads
	toSerialize["members"] = o.Members
	toSerialize["has_more"] = o.HasMore
	if !IsNil(o.FirstMessages) {
		toSerialize["first_messages"] = o.FirstMessages
	}
	toSerialize["total_results"] = o.TotalResults
	return toSerialize, nil
}

func (o *ThreadSearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threads",
		"members",
		"has_more",
		"total_results",
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

	varThreadSearchResponse := _ThreadSearchResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThreadSearchResponse)

	if err != nil {
		return err
	}

	*o = ThreadSearchResponse(varThreadSearchResponse)

	return err
}

type NullableThreadSearchResponse struct {
	value *ThreadSearchResponse
	isSet bool
}

func (v NullableThreadSearchResponse) Get() *ThreadSearchResponse {
	return v.value
}

func (v *NullableThreadSearchResponse) Set(val *ThreadSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadSearchResponse(val *ThreadSearchResponse) *NullableThreadSearchResponse {
	return &NullableThreadSearchResponse{value: val, isSet: true}
}

func (v NullableThreadSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

