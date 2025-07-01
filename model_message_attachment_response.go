/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the MessageAttachmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttachmentResponse{}

// MessageAttachmentResponse struct for MessageAttachmentResponse
type MessageAttachmentResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Filename string `json:"filename"`
	Size int32 `json:"size"`
	Url string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Width NullableInt32 `json:"width,omitempty"`
	Height NullableInt32 `json:"height,omitempty"`
	DurationSecs NullableFloat64 `json:"duration_secs,omitempty"`
	Waveform NullableString `json:"waveform,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ContentType NullableString `json:"content_type,omitempty"`
	Ephemeral NullableBool `json:"ephemeral,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Application NullableApplicationResponse `json:"application,omitempty"`
	ClipCreatedAt NullableTime `json:"clip_created_at,omitempty"`
	ClipParticipants []UserResponse `json:"clip_participants,omitempty"`
}

type _MessageAttachmentResponse MessageAttachmentResponse

// NewMessageAttachmentResponse instantiates a new MessageAttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttachmentResponse(id string, filename string, size int32, url string, proxyUrl string) *MessageAttachmentResponse {
	this := MessageAttachmentResponse{}
	this.Id = id
	this.Filename = filename
	this.Size = size
	this.Url = url
	this.ProxyUrl = proxyUrl
	return &this
}

// NewMessageAttachmentResponseWithDefaults instantiates a new MessageAttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttachmentResponseWithDefaults() *MessageAttachmentResponse {
	this := MessageAttachmentResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MessageAttachmentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageAttachmentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageAttachmentResponse) SetId(v string) {
	o.Id = v
}

// GetFilename returns the Filename field value
func (o *MessageAttachmentResponse) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *MessageAttachmentResponse) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *MessageAttachmentResponse) SetFilename(v string) {
	o.Filename = v
}

// GetSize returns the Size field value
func (o *MessageAttachmentResponse) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *MessageAttachmentResponse) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *MessageAttachmentResponse) SetSize(v int32) {
	o.Size = v
}

// GetUrl returns the Url field value
func (o *MessageAttachmentResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MessageAttachmentResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MessageAttachmentResponse) SetUrl(v string) {
	o.Url = v
}

// GetProxyUrl returns the ProxyUrl field value
func (o *MessageAttachmentResponse) GetProxyUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProxyUrl
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value
// and a boolean to check if the value has been set.
func (o *MessageAttachmentResponse) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProxyUrl, true
}

// SetProxyUrl sets field value
func (o *MessageAttachmentResponse) SetProxyUrl(v string) {
	o.ProxyUrl = v
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width.Get()) {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *MessageAttachmentResponse) SetWidth(v int32) {
	o.Width.Set(&v)
}

// SetWidthNil sets the value for Width to be an explicit nil
func (o *MessageAttachmentResponse) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetWidth() {
	o.Width.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height.Get()) {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *MessageAttachmentResponse) SetHeight(v int32) {
	o.Height.Set(&v)
}

// SetHeightNil sets the value for Height to be an explicit nil
func (o *MessageAttachmentResponse) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetHeight() {
	o.Height.Unset()
}

// GetDurationSecs returns the DurationSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetDurationSecs() float64 {
	if o == nil || IsNil(o.DurationSecs.Get()) {
		var ret float64
		return ret
	}
	return *o.DurationSecs.Get()
}

// GetDurationSecsOk returns a tuple with the DurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetDurationSecsOk() (*float64, bool) {
	if o == nil || IsNil(o.DurationSecs.Get()) {
		return nil, false
	}
	return o.DurationSecs.Get(), o.DurationSecs.IsSet()
}

// HasDurationSecs returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasDurationSecs() bool {
	if o != nil && o.DurationSecs.IsSet() {
		return true
	}

	return false
}

// SetDurationSecs gets a reference to the given NullableFloat64 and assigns it to the DurationSecs field.
func (o *MessageAttachmentResponse) SetDurationSecs(v float64) {
	o.DurationSecs.Set(&v)
}

// SetDurationSecsNil sets the value for DurationSecs to be an explicit nil
func (o *MessageAttachmentResponse) SetDurationSecsNil() {
	o.DurationSecs.Set(nil)
}

// UnsetDurationSecs ensures that no value is present for DurationSecs, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetDurationSecs() {
	o.DurationSecs.Unset()
}

// GetWaveform returns the Waveform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetWaveform() string {
	if o == nil || IsNil(o.Waveform.Get()) {
		var ret string
		return ret
	}
	return *o.Waveform.Get()
}

// GetWaveformOk returns a tuple with the Waveform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetWaveformOk() (*string, bool) {
	if o == nil || IsNil(o.Waveform.Get()) {
		return nil, false
	}
	return o.Waveform.Get(), o.Waveform.IsSet()
}

// HasWaveform returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasWaveform() bool {
	if o != nil && o.Waveform.IsSet() {
		return true
	}

	return false
}

// SetWaveform gets a reference to the given NullableString and assigns it to the Waveform field.
func (o *MessageAttachmentResponse) SetWaveform(v string) {
	o.Waveform.Set(&v)
}

// SetWaveformNil sets the value for Waveform to be an explicit nil
func (o *MessageAttachmentResponse) SetWaveformNil() {
	o.Waveform.Set(nil)
}

// UnsetWaveform ensures that no value is present for Waveform, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetWaveform() {
	o.Waveform.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description.Get()) {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MessageAttachmentResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MessageAttachmentResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType.Get()) {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *MessageAttachmentResponse) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MessageAttachmentResponse) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetContentType() {
	o.ContentType.Unset()
}

// GetEphemeral returns the Ephemeral field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetEphemeral() bool {
	if o == nil || IsNil(o.Ephemeral.Get()) {
		var ret bool
		return ret
	}
	return *o.Ephemeral.Get()
}

// GetEphemeralOk returns a tuple with the Ephemeral field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetEphemeralOk() (*bool, bool) {
	if o == nil || IsNil(o.Ephemeral.Get()) {
		return nil, false
	}
	return o.Ephemeral.Get(), o.Ephemeral.IsSet()
}

// HasEphemeral returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasEphemeral() bool {
	if o != nil && o.Ephemeral.IsSet() {
		return true
	}

	return false
}

// SetEphemeral gets a reference to the given NullableBool and assigns it to the Ephemeral field.
func (o *MessageAttachmentResponse) SetEphemeral(v bool) {
	o.Ephemeral.Set(&v)
}

// SetEphemeralNil sets the value for Ephemeral to be an explicit nil
func (o *MessageAttachmentResponse) SetEphemeralNil() {
	o.Ephemeral.Set(nil)
}

// UnsetEphemeral ensures that no value is present for Ephemeral, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetEphemeral() {
	o.Ephemeral.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title.Get()) {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MessageAttachmentResponse) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *MessageAttachmentResponse) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetTitle() {
	o.Title.Unset()
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetApplication() ApplicationResponse {
	if o == nil || IsNil(o.Application.Get()) {
		var ret ApplicationResponse
		return ret
	}
	return *o.Application.Get()
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetApplicationOk() (*ApplicationResponse, bool) {
	if o == nil || IsNil(o.Application.Get()) {
		return nil, false
	}
	return o.Application.Get(), o.Application.IsSet()
}

// HasApplication returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasApplication() bool {
	if o != nil && o.Application.IsSet() {
		return true
	}

	return false
}

// SetApplication gets a reference to the given NullableApplicationResponse and assigns it to the Application field.
func (o *MessageAttachmentResponse) SetApplication(v ApplicationResponse) {
	o.Application.Set(&v)
}

// SetApplicationNil sets the value for Application to be an explicit nil
func (o *MessageAttachmentResponse) SetApplicationNil() {
	o.Application.Set(nil)
}

// UnsetApplication ensures that no value is present for Application, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetApplication() {
	o.Application.Unset()
}

// GetClipCreatedAt returns the ClipCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetClipCreatedAt() time.Time {
	if o == nil || IsNil(o.ClipCreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ClipCreatedAt.Get()
}

// GetClipCreatedAtOk returns a tuple with the ClipCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetClipCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClipCreatedAt.Get()) {
		return nil, false
	}
	return o.ClipCreatedAt.Get(), o.ClipCreatedAt.IsSet()
}

// HasClipCreatedAt returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasClipCreatedAt() bool {
	if o != nil && o.ClipCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetClipCreatedAt gets a reference to the given NullableTime and assigns it to the ClipCreatedAt field.
func (o *MessageAttachmentResponse) SetClipCreatedAt(v time.Time) {
	o.ClipCreatedAt.Set(&v)
}

// SetClipCreatedAtNil sets the value for ClipCreatedAt to be an explicit nil
func (o *MessageAttachmentResponse) SetClipCreatedAtNil() {
	o.ClipCreatedAt.Set(nil)
}

// UnsetClipCreatedAt ensures that no value is present for ClipCreatedAt, not even an explicit nil
func (o *MessageAttachmentResponse) UnsetClipCreatedAt() {
	o.ClipCreatedAt.Unset()
}

// GetClipParticipants returns the ClipParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttachmentResponse) GetClipParticipants() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}
	return o.ClipParticipants
}

// GetClipParticipantsOk returns a tuple with the ClipParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttachmentResponse) GetClipParticipantsOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClipParticipants, true
}

// HasClipParticipants returns a boolean if a field has been set.
func (o *MessageAttachmentResponse) HasClipParticipants() bool {
	if o != nil && !IsNil(o.ClipParticipants) {
		return true
	}

	return false
}

// SetClipParticipants gets a reference to the given []UserResponse and assigns it to the ClipParticipants field.
func (o *MessageAttachmentResponse) SetClipParticipants(v []UserResponse) {
	o.ClipParticipants = v
}


func (o MessageAttachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttachmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["filename"] = o.Filename
	toSerialize["size"] = o.Size
	toSerialize["url"] = o.Url
	toSerialize["proxy_url"] = o.ProxyUrl
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.DurationSecs.IsSet() {
		toSerialize["duration_secs"] = o.DurationSecs.Get()
	}
	if o.Waveform.IsSet() {
		toSerialize["waveform"] = o.Waveform.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.Ephemeral.IsSet() {
		toSerialize["ephemeral"] = o.Ephemeral.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Application.IsSet() {
		toSerialize["application"] = o.Application.Get()
	}
	if o.ClipCreatedAt.IsSet() {
		toSerialize["clip_created_at"] = o.ClipCreatedAt.Get()
	}
	if o.ClipParticipants != nil {
		toSerialize["clip_participants"] = o.ClipParticipants
	}
	return toSerialize, nil
}

func (o *MessageAttachmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"filename",
		"size",
		"url",
		"proxy_url",
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

	varMessageAttachmentResponse := _MessageAttachmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAttachmentResponse)

	if err != nil {
		return err
	}

	*o = MessageAttachmentResponse(varMessageAttachmentResponse)

	return err
}

type NullableMessageAttachmentResponse struct {
	value *MessageAttachmentResponse
	isSet bool
}

func (v NullableMessageAttachmentResponse) Get() *MessageAttachmentResponse {
	return v.value
}

func (v *NullableMessageAttachmentResponse) Set(val *MessageAttachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttachmentResponse(val *MessageAttachmentResponse) *NullableMessageAttachmentResponse {
	return &NullableMessageAttachmentResponse{value: val, isSet: true}
}

func (v NullableMessageAttachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


