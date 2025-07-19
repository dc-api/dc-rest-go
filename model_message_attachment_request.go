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

// checks if the MessageAttachmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttachmentRequest{}

// MessageAttachmentRequest struct for MessageAttachmentRequest
type MessageAttachmentRequest struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Filename NullableString `json:"filename,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DurationSecs NullableFloat64 `json:"duration_secs,omitempty"`
	Waveform NullableString `json:"waveform,omitempty"`
	Title NullableString `json:"title,omitempty"`
	IsRemix NullableBool `json:"is_remix,omitempty"`
}

type _MessageAttachmentRequest MessageAttachmentRequest

// NewMessageAttachmentRequest instantiates a new MessageAttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttachmentRequest(id string) *MessageAttachmentRequest {
	this := MessageAttachmentRequest{}
	this.Id = id
	return &this
}

// NewMessageAttachmentRequestWithDefaults instantiates a new MessageAttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttachmentRequestWithDefaults() *MessageAttachmentRequest {
	this := MessageAttachmentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *MessageAttachmentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageAttachmentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageAttachmentRequest) SetId(v string) {
	o.Id = v
}

// GetFilename returns the Filename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentRequest) GetFilename() string {
	if o == nil || IsNil(o.Filename.Get()) {
		var ret string
		return ret
	}
	return *o.Filename.Get()
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentRequest) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename.Get()) {
		return nil, false
	}
	return o.Filename.Get(), o.Filename.IsSet()
}

// HasFilename returns a boolean if a field has been set.
func (o *MessageAttachmentRequest) HasFilename() bool {
	if o != nil && o.Filename.IsSet() {
		return true
	}

	return false
}

// SetFilename gets a reference to the given NullableString and assigns it to the Filename field.
func (o *MessageAttachmentRequest) SetFilename(v string) {
	o.Filename.Set(&v)
}

// SetFilenameNil sets the value for Filename to be an explicit nil
func (o *MessageAttachmentRequest) SetFilenameNil() {
	o.Filename.Set(nil)
}

// UnsetFilename ensures that no value is present for Filename, not even an explicit nil
func (o *MessageAttachmentRequest) UnsetFilename() {
	o.Filename.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageAttachmentRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MessageAttachmentRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MessageAttachmentRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MessageAttachmentRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetDurationSecs returns the DurationSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentRequest) GetDurationSecs() float64 {
	if o == nil || IsNil(o.DurationSecs.Get()) {
		var ret float64
		return ret
	}
	return *o.DurationSecs.Get()
}

// GetDurationSecsOk returns a tuple with the DurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentRequest) GetDurationSecsOk() (*float64, bool) {
	if o == nil || IsNil(o.DurationSecs.Get()) {
		return nil, false
	}
	return o.DurationSecs.Get(), o.DurationSecs.IsSet()
}

// HasDurationSecs returns a boolean if a field has been set.
func (o *MessageAttachmentRequest) HasDurationSecs() bool {
	if o != nil && o.DurationSecs.IsSet() {
		return true
	}

	return false
}

// SetDurationSecs gets a reference to the given NullableFloat64 and assigns it to the DurationSecs field.
func (o *MessageAttachmentRequest) SetDurationSecs(v float64) {
	o.DurationSecs.Set(&v)
}

// SetDurationSecsNil sets the value for DurationSecs to be an explicit nil
func (o *MessageAttachmentRequest) SetDurationSecsNil() {
	o.DurationSecs.Set(nil)
}

// UnsetDurationSecs ensures that no value is present for DurationSecs, not even an explicit nil
func (o *MessageAttachmentRequest) UnsetDurationSecs() {
	o.DurationSecs.Unset()
}

// GetWaveform returns the Waveform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentRequest) GetWaveform() string {
	if o == nil || IsNil(o.Waveform.Get()) {
		var ret string
		return ret
	}
	return *o.Waveform.Get()
}

// GetWaveformOk returns a tuple with the Waveform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentRequest) GetWaveformOk() (*string, bool) {
	if o == nil || IsNil(o.Waveform.Get()) {
		return nil, false
	}
	return o.Waveform.Get(), o.Waveform.IsSet()
}

// HasWaveform returns a boolean if a field has been set.
func (o *MessageAttachmentRequest) HasWaveform() bool {
	if o != nil && o.Waveform.IsSet() {
		return true
	}

	return false
}

// SetWaveform gets a reference to the given NullableString and assigns it to the Waveform field.
func (o *MessageAttachmentRequest) SetWaveform(v string) {
	o.Waveform.Set(&v)
}

// SetWaveformNil sets the value for Waveform to be an explicit nil
func (o *MessageAttachmentRequest) SetWaveformNil() {
	o.Waveform.Set(nil)
}

// UnsetWaveform ensures that no value is present for Waveform, not even an explicit nil
func (o *MessageAttachmentRequest) UnsetWaveform() {
	o.Waveform.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title.Get()) {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MessageAttachmentRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MessageAttachmentRequest) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *MessageAttachmentRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MessageAttachmentRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetIsRemix returns the IsRemix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentRequest) GetIsRemix() bool {
	if o == nil || IsNil(o.IsRemix.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRemix.Get()
}

// GetIsRemixOk returns a tuple with the IsRemix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentRequest) GetIsRemixOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRemix.Get()) {
		return nil, false
	}
	return o.IsRemix.Get(), o.IsRemix.IsSet()
}

// HasIsRemix returns a boolean if a field has been set.
func (o *MessageAttachmentRequest) HasIsRemix() bool {
	if o != nil && o.IsRemix.IsSet() {
		return true
	}

	return false
}

// SetIsRemix gets a reference to the given NullableBool and assigns it to the IsRemix field.
func (o *MessageAttachmentRequest) SetIsRemix(v bool) {
	o.IsRemix.Set(&v)
}

// SetIsRemixNil sets the value for IsRemix to be an explicit nil
func (o *MessageAttachmentRequest) SetIsRemixNil() {
	o.IsRemix.Set(nil)
}

// UnsetIsRemix ensures that no value is present for IsRemix, not even an explicit nil
func (o *MessageAttachmentRequest) UnsetIsRemix() {
	o.IsRemix.Unset()
}

func (o MessageAttachmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttachmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Filename.IsSet() {
		toSerialize["filename"] = o.Filename.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DurationSecs.IsSet() {
		toSerialize["duration_secs"] = o.DurationSecs.Get()
	}
	if o.Waveform.IsSet() {
		toSerialize["waveform"] = o.Waveform.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.IsRemix.IsSet() {
		toSerialize["is_remix"] = o.IsRemix.Get()
	}
	return toSerialize, nil
}

func (o *MessageAttachmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varMessageAttachmentRequest := _MessageAttachmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAttachmentRequest)

	if err != nil {
		return err
	}

	*o = MessageAttachmentRequest(varMessageAttachmentRequest)

	return err
}

type NullableMessageAttachmentRequest struct {
	value *MessageAttachmentRequest
	isSet bool
}

func (v NullableMessageAttachmentRequest) Get() *MessageAttachmentRequest {
	return v.value
}

func (v *NullableMessageAttachmentRequest) Set(val *MessageAttachmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttachmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttachmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttachmentRequest(val *MessageAttachmentRequest) *NullableMessageAttachmentRequest {
	return &NullableMessageAttachmentRequest{value: val, isSet: true}
}

func (v NullableMessageAttachmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttachmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


