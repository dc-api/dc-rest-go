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

// checks if the ThreadSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadSearchResponse{}

// ThreadSearchResponse struct for ThreadSearchResponse
type ThreadSearchResponse struct {
	Threads []ThreadResponse `json:"threads"`
	Members []ThreadMemberResponse `json:"members"`
	HasMore NullableBool `json:"has_more,omitempty"`
	FirstMessages []MessageResponse `json:"first_messages,omitempty"`
	TotalResults NullableInt32 `json:"total_results,omitempty"`
}

type _ThreadSearchResponse ThreadSearchResponse

// NewThreadSearchResponse instantiates a new ThreadSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadSearchResponse(threads []ThreadResponse, members []ThreadMemberResponse) *ThreadSearchResponse {
	this := ThreadSearchResponse{}
	this.Threads = threads
	this.Members = members
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

// GetHasMore returns the HasMore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadSearchResponse) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore.Get()) {
		var ret bool
		return ret
	}
	return *o.HasMore.Get()
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadSearchResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore.Get()) {
		return nil, false
	}
	return o.HasMore.Get(), o.HasMore.IsSet()
}

// HasHasMore returns a boolean if a field has been set.
func (o *ThreadSearchResponse) HasHasMore() bool {
	if o != nil && o.HasMore.IsSet() {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given NullableBool and assigns it to the HasMore field.
func (o *ThreadSearchResponse) SetHasMore(v bool) {
	o.HasMore.Set(&v)
}

// SetHasMoreNil sets the value for HasMore to be an explicit nil
func (o *ThreadSearchResponse) SetHasMoreNil() {
	o.HasMore.Set(nil)
}

// UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
func (o *ThreadSearchResponse) UnsetHasMore() {
	o.HasMore.Unset()
}

// GetFirstMessages returns the FirstMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadSearchResponse) GetFirstMessages() []MessageResponse {
	if o == nil {
		var ret []MessageResponse
		return ret
	}
	return o.FirstMessages
}

// GetFirstMessagesOk returns a tuple with the FirstMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadSearchResponse) GetFirstMessagesOk() ([]MessageResponse, bool) {
	if o == nil {
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


// GetTotalResults returns the TotalResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadSearchResponse) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalResults.Get()
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadSearchResponse) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults.Get()) {
		return nil, false
	}
	return o.TotalResults.Get(), o.TotalResults.IsSet()
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ThreadSearchResponse) HasTotalResults() bool {
	if o != nil && o.TotalResults.IsSet() {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given NullableInt32 and assigns it to the TotalResults field.
func (o *ThreadSearchResponse) SetTotalResults(v int32) {
	o.TotalResults.Set(&v)
}

// SetTotalResultsNil sets the value for TotalResults to be an explicit nil
func (o *ThreadSearchResponse) SetTotalResultsNil() {
	o.TotalResults.Set(nil)
}

// UnsetTotalResults ensures that no value is present for TotalResults, not even an explicit nil
func (o *ThreadSearchResponse) UnsetTotalResults() {
	o.TotalResults.Unset()
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
	if o.HasMore.IsSet() {
		toSerialize["has_more"] = o.HasMore.Get()
	}
	if o.FirstMessages != nil {
		toSerialize["first_messages"] = o.FirstMessages
	}
	if o.TotalResults.IsSet() {
		toSerialize["total_results"] = o.TotalResults.Get()
	}
	return toSerialize, nil
}

func (o *ThreadSearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threads",
		"members",
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


