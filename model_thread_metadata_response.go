/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the ThreadMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadMetadataResponse{}

// ThreadMetadataResponse struct for ThreadMetadataResponse
type ThreadMetadataResponse struct {
	Archived bool `json:"archived"`
	ArchiveTimestamp NullableTime `json:"archive_timestamp,omitempty"`
	AutoArchiveDuration int32 `json:"auto_archive_duration"`
	Locked bool `json:"locked"`
	CreateTimestamp NullableTime `json:"create_timestamp,omitempty"`
	Invitable NullableBool `json:"invitable,omitempty"`
}

type _ThreadMetadataResponse ThreadMetadataResponse

// NewThreadMetadataResponse instantiates a new ThreadMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadMetadataResponse(archived bool, autoArchiveDuration int32, locked bool) *ThreadMetadataResponse {
	this := ThreadMetadataResponse{}
	this.Archived = archived
	this.AutoArchiveDuration = autoArchiveDuration
	this.Locked = locked
	return &this
}

// NewThreadMetadataResponseWithDefaults instantiates a new ThreadMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadMetadataResponseWithDefaults() *ThreadMetadataResponse {
	this := ThreadMetadataResponse{}
	return &this
}

// GetArchived returns the Archived field value
func (o *ThreadMetadataResponse) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *ThreadMetadataResponse) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *ThreadMetadataResponse) SetArchived(v bool) {
	o.Archived = v
}

// GetArchiveTimestamp returns the ArchiveTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadMetadataResponse) GetArchiveTimestamp() time.Time {
	if o == nil || IsNil(o.ArchiveTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ArchiveTimestamp.Get()
}

// GetArchiveTimestampOk returns a tuple with the ArchiveTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadMetadataResponse) GetArchiveTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArchiveTimestamp.Get()) {
		return nil, false
	}
	return o.ArchiveTimestamp.Get(), o.ArchiveTimestamp.IsSet()
}

// HasArchiveTimestamp returns a boolean if a field has been set.
func (o *ThreadMetadataResponse) HasArchiveTimestamp() bool {
	if o != nil && o.ArchiveTimestamp.IsSet() {
		return true
	}

	return false
}

// SetArchiveTimestamp gets a reference to the given NullableTime and assigns it to the ArchiveTimestamp field.
func (o *ThreadMetadataResponse) SetArchiveTimestamp(v time.Time) {
	o.ArchiveTimestamp.Set(&v)
}

// SetArchiveTimestampNil sets the value for ArchiveTimestamp to be an explicit nil
func (o *ThreadMetadataResponse) SetArchiveTimestampNil() {
	o.ArchiveTimestamp.Set(nil)
}

// UnsetArchiveTimestamp ensures that no value is present for ArchiveTimestamp, not even an explicit nil
func (o *ThreadMetadataResponse) UnsetArchiveTimestamp() {
	o.ArchiveTimestamp.Unset()
}

// GetAutoArchiveDuration returns the AutoArchiveDuration field value
func (o *ThreadMetadataResponse) GetAutoArchiveDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AutoArchiveDuration
}

// GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field value
// and a boolean to check if the value has been set.
func (o *ThreadMetadataResponse) GetAutoArchiveDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoArchiveDuration, true
}

// SetAutoArchiveDuration sets field value
func (o *ThreadMetadataResponse) SetAutoArchiveDuration(v int32) {
	o.AutoArchiveDuration = v
}

// GetLocked returns the Locked field value
func (o *ThreadMetadataResponse) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *ThreadMetadataResponse) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *ThreadMetadataResponse) SetLocked(v bool) {
	o.Locked = v
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadMetadataResponse) GetCreateTimestamp() time.Time {
	if o == nil || IsNil(o.CreateTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreateTimestamp.Get()
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadMetadataResponse) GetCreateTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTimestamp.Get()) {
		return nil, false
	}
	return o.CreateTimestamp.Get(), o.CreateTimestamp.IsSet()
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *ThreadMetadataResponse) HasCreateTimestamp() bool {
	if o != nil && o.CreateTimestamp.IsSet() {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given NullableTime and assigns it to the CreateTimestamp field.
func (o *ThreadMetadataResponse) SetCreateTimestamp(v time.Time) {
	o.CreateTimestamp.Set(&v)
}

// SetCreateTimestampNil sets the value for CreateTimestamp to be an explicit nil
func (o *ThreadMetadataResponse) SetCreateTimestampNil() {
	o.CreateTimestamp.Set(nil)
}

// UnsetCreateTimestamp ensures that no value is present for CreateTimestamp, not even an explicit nil
func (o *ThreadMetadataResponse) UnsetCreateTimestamp() {
	o.CreateTimestamp.Unset()
}

// GetInvitable returns the Invitable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadMetadataResponse) GetInvitable() bool {
	if o == nil || IsNil(o.Invitable.Get()) {
		var ret bool
		return ret
	}
	return *o.Invitable.Get()
}

// GetInvitableOk returns a tuple with the Invitable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadMetadataResponse) GetInvitableOk() (*bool, bool) {
	if o == nil || IsNil(o.Invitable.Get()) {
		return nil, false
	}
	return o.Invitable.Get(), o.Invitable.IsSet()
}

// HasInvitable returns a boolean if a field has been set.
func (o *ThreadMetadataResponse) HasInvitable() bool {
	if o != nil && o.Invitable.IsSet() {
		return true
	}

	return false
}

// SetInvitable gets a reference to the given NullableBool and assigns it to the Invitable field.
func (o *ThreadMetadataResponse) SetInvitable(v bool) {
	o.Invitable.Set(&v)
}

// SetInvitableNil sets the value for Invitable to be an explicit nil
func (o *ThreadMetadataResponse) SetInvitableNil() {
	o.Invitable.Set(nil)
}

// UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil
func (o *ThreadMetadataResponse) UnsetInvitable() {
	o.Invitable.Unset()
}

func (o ThreadMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archived"] = o.Archived
	if o.ArchiveTimestamp.IsSet() {
		toSerialize["archive_timestamp"] = o.ArchiveTimestamp.Get()
	}
	toSerialize["auto_archive_duration"] = o.AutoArchiveDuration
	toSerialize["locked"] = o.Locked
	if o.CreateTimestamp.IsSet() {
		toSerialize["create_timestamp"] = o.CreateTimestamp.Get()
	}
	if o.Invitable.IsSet() {
		toSerialize["invitable"] = o.Invitable.Get()
	}
	return toSerialize, nil
}

func (o *ThreadMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archived",
		"auto_archive_duration",
		"locked",
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

	varThreadMetadataResponse := _ThreadMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThreadMetadataResponse)

	if err != nil {
		return err
	}

	*o = ThreadMetadataResponse(varThreadMetadataResponse)

	return err
}

type NullableThreadMetadataResponse struct {
	value *ThreadMetadataResponse
	isSet bool
}

func (v NullableThreadMetadataResponse) Get() *ThreadMetadataResponse {
	return v.value
}

func (v *NullableThreadMetadataResponse) Set(val *ThreadMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadMetadataResponse(val *ThreadMetadataResponse) *NullableThreadMetadataResponse {
	return &NullableThreadMetadataResponse{value: val, isSet: true}
}

func (v NullableThreadMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


