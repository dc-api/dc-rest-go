/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-10-14T14:27:02.905060650Z[Etc/UTC]
 *    * - **Generator Version**: 7.16.0
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

// checks if the ThreadsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadsResponse{}

// ThreadsResponse struct for ThreadsResponse
type ThreadsResponse struct {
	Threads []ThreadResponse `json:"threads"`
	Members []ThreadMemberResponse `json:"members"`
	HasMore bool `json:"has_more"`
	FirstMessages []MessageResponse `json:"first_messages,omitempty"`
}

type _ThreadsResponse ThreadsResponse

// NewThreadsResponse instantiates a new ThreadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadsResponse(threads []ThreadResponse, members []ThreadMemberResponse, hasMore bool) *ThreadsResponse {
	this := ThreadsResponse{}
	this.Threads = threads
	this.Members = members
	this.HasMore = hasMore
	return &this
}

// NewThreadsResponseWithDefaults instantiates a new ThreadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadsResponseWithDefaults() *ThreadsResponse {
	this := ThreadsResponse{}
	return &this
}

// GetThreads returns the Threads field value
func (o *ThreadsResponse) GetThreads() []ThreadResponse {
	if o == nil {
		var ret []ThreadResponse
		return ret
	}

	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value
// and a boolean to check if the value has been set.
func (o *ThreadsResponse) GetThreadsOk() ([]ThreadResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads, true
}

// SetThreads sets field value
func (o *ThreadsResponse) SetThreads(v []ThreadResponse) {
	o.Threads = v
}

// GetMembers returns the Members field value
func (o *ThreadsResponse) GetMembers() []ThreadMemberResponse {
	if o == nil {
		var ret []ThreadMemberResponse
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ThreadsResponse) GetMembersOk() ([]ThreadMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *ThreadsResponse) SetMembers(v []ThreadMemberResponse) {
	o.Members = v
}

// GetHasMore returns the HasMore field value
func (o *ThreadsResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ThreadsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ThreadsResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetFirstMessages returns the FirstMessages field value if set, zero value otherwise.
func (o *ThreadsResponse) GetFirstMessages() []MessageResponse {
	if o == nil || IsNil(o.FirstMessages) {
		var ret []MessageResponse
		return ret
	}
	return o.FirstMessages
}

// GetFirstMessagesOk returns a tuple with the FirstMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadsResponse) GetFirstMessagesOk() ([]MessageResponse, bool) {
	if o == nil || IsNil(o.FirstMessages) {
		return nil, false
	}
	return o.FirstMessages, true
}

// HasFirstMessages returns a boolean if a field has been set.
func (o *ThreadsResponse) HasFirstMessages() bool {
	if o != nil && !IsNil(o.FirstMessages) {
		return true
	}

	return false
}

// SetFirstMessages gets a reference to the given []MessageResponse and assigns it to the FirstMessages field.
func (o *ThreadsResponse) SetFirstMessages(v []MessageResponse) {
	o.FirstMessages = v
}


func (o ThreadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threads"] = o.Threads
	toSerialize["members"] = o.Members
	toSerialize["has_more"] = o.HasMore
	if !IsNil(o.FirstMessages) {
		toSerialize["first_messages"] = o.FirstMessages
	}
	return toSerialize, nil
}

func (o *ThreadsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threads",
		"members",
		"has_more",
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

	varThreadsResponse := _ThreadsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThreadsResponse)

	if err != nil {
		return err
	}

	*o = ThreadsResponse(varThreadsResponse)

	return err
}

type NullableThreadsResponse struct {
	value *ThreadsResponse
	isSet bool
}

func (v NullableThreadsResponse) Get() *ThreadsResponse {
	return v.value
}

func (v *NullableThreadsResponse) Set(val *ThreadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadsResponse(val *ThreadsResponse) *NullableThreadsResponse {
	return &NullableThreadsResponse{value: val, isSet: true}
}

func (v NullableThreadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

