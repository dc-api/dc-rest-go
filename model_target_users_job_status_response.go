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
	"time"
	"bytes"
	"fmt"
)

// checks if the TargetUsersJobStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetUsersJobStatusResponse{}

// TargetUsersJobStatusResponse struct for TargetUsersJobStatusResponse
type TargetUsersJobStatusResponse struct {
	// The status of the job processing the target users.
	Status int32 `json:"status"`
	// The total number of users in the provided list.
	TotalUsers int64 `json:"total_users"`
	// The number of users processed so far.
	ProcessedUsers int64 `json:"processed_users"`
	// The timestamp when the job was created.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// The timestamp when the job was successfully completed.
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	// The error message if the job failed.
	ErrorMessage NullableString `json:"error_message,omitempty"`
}

type _TargetUsersJobStatusResponse TargetUsersJobStatusResponse

// NewTargetUsersJobStatusResponse instantiates a new TargetUsersJobStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetUsersJobStatusResponse(status int32, totalUsers int64, processedUsers int64) *TargetUsersJobStatusResponse {
	this := TargetUsersJobStatusResponse{}
	this.Status = status
	this.TotalUsers = totalUsers
	this.ProcessedUsers = processedUsers
	return &this
}

// NewTargetUsersJobStatusResponseWithDefaults instantiates a new TargetUsersJobStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetUsersJobStatusResponseWithDefaults() *TargetUsersJobStatusResponse {
	this := TargetUsersJobStatusResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *TargetUsersJobStatusResponse) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TargetUsersJobStatusResponse) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TargetUsersJobStatusResponse) SetStatus(v int32) {
	o.Status = v
}

// GetTotalUsers returns the TotalUsers field value
func (o *TargetUsersJobStatusResponse) GetTotalUsers() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value
// and a boolean to check if the value has been set.
func (o *TargetUsersJobStatusResponse) GetTotalUsersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalUsers, true
}

// SetTotalUsers sets field value
func (o *TargetUsersJobStatusResponse) SetTotalUsers(v int64) {
	o.TotalUsers = v
}

// GetProcessedUsers returns the ProcessedUsers field value
func (o *TargetUsersJobStatusResponse) GetProcessedUsers() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProcessedUsers
}

// GetProcessedUsersOk returns a tuple with the ProcessedUsers field value
// and a boolean to check if the value has been set.
func (o *TargetUsersJobStatusResponse) GetProcessedUsersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessedUsers, true
}

// SetProcessedUsers sets field value
func (o *TargetUsersJobStatusResponse) SetProcessedUsers(v int64) {
	o.ProcessedUsers = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetUsersJobStatusResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetUsersJobStatusResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TargetUsersJobStatusResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *TargetUsersJobStatusResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TargetUsersJobStatusResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TargetUsersJobStatusResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetUsersJobStatusResponse) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetUsersJobStatusResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *TargetUsersJobStatusResponse) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *TargetUsersJobStatusResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *TargetUsersJobStatusResponse) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *TargetUsersJobStatusResponse) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetUsersJobStatusResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetUsersJobStatusResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TargetUsersJobStatusResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *TargetUsersJobStatusResponse) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *TargetUsersJobStatusResponse) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *TargetUsersJobStatusResponse) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o TargetUsersJobStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetUsersJobStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["total_users"] = o.TotalUsers
	toSerialize["processed_users"] = o.ProcessedUsers
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	return toSerialize, nil
}

func (o *TargetUsersJobStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"total_users",
		"processed_users",
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

	varTargetUsersJobStatusResponse := _TargetUsersJobStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetUsersJobStatusResponse)

	if err != nil {
		return err
	}

	*o = TargetUsersJobStatusResponse(varTargetUsersJobStatusResponse)

	return err
}

type NullableTargetUsersJobStatusResponse struct {
	value *TargetUsersJobStatusResponse
	isSet bool
}

func (v NullableTargetUsersJobStatusResponse) Get() *TargetUsersJobStatusResponse {
	return v.value
}

func (v *NullableTargetUsersJobStatusResponse) Set(val *TargetUsersJobStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetUsersJobStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetUsersJobStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetUsersJobStatusResponse(val *TargetUsersJobStatusResponse) *NullableTargetUsersJobStatusResponse {
	return &NullableTargetUsersJobStatusResponse{value: val, isSet: true}
}

func (v NullableTargetUsersJobStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetUsersJobStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

