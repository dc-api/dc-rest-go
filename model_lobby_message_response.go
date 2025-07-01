/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the LobbyMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LobbyMessageResponse{}

// LobbyMessageResponse struct for LobbyMessageResponse
type LobbyMessageResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type int32 `json:"type"`
	Content string `json:"content"`
	LobbyId string `json:"lobby_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Author UserResponse `json:"author"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Flags int32 `json:"flags"`
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _LobbyMessageResponse LobbyMessageResponse

// NewLobbyMessageResponse instantiates a new LobbyMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLobbyMessageResponse(id string, type_ int32, content string, lobbyId string, channelId string, author UserResponse, flags int32) *LobbyMessageResponse {
	this := LobbyMessageResponse{}
	this.Id = id
	this.Type = type_
	this.Content = content
	this.LobbyId = lobbyId
	this.ChannelId = channelId
	this.Author = author
	this.Flags = flags
	return &this
}

// NewLobbyMessageResponseWithDefaults instantiates a new LobbyMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLobbyMessageResponseWithDefaults() *LobbyMessageResponse {
	this := LobbyMessageResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LobbyMessageResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LobbyMessageResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *LobbyMessageResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LobbyMessageResponse) SetType(v int32) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *LobbyMessageResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *LobbyMessageResponse) SetContent(v string) {
	o.Content = v
}

// GetLobbyId returns the LobbyId field value
func (o *LobbyMessageResponse) GetLobbyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LobbyId
}

// GetLobbyIdOk returns a tuple with the LobbyId field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetLobbyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LobbyId, true
}

// SetLobbyId sets field value
func (o *LobbyMessageResponse) SetLobbyId(v string) {
	o.LobbyId = v
}

// GetChannelId returns the ChannelId field value
func (o *LobbyMessageResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *LobbyMessageResponse) SetChannelId(v string) {
	o.ChannelId = v
}

// GetAuthor returns the Author field value
func (o *LobbyMessageResponse) GetAuthor() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetAuthorOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *LobbyMessageResponse) SetAuthor(v UserResponse) {
	o.Author = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LobbyMessageResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetMetadataOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]string{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LobbyMessageResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *LobbyMessageResponse) SetMetadata(v map[string]string) {
	o.Metadata = v
}


// GetFlags returns the Flags field value
func (o *LobbyMessageResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *LobbyMessageResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *LobbyMessageResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LobbyMessageResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *LobbyMessageResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *LobbyMessageResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}


func (o LobbyMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LobbyMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	toSerialize["lobby_id"] = o.LobbyId
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["author"] = o.Author
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["flags"] = o.Flags
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	return toSerialize, nil
}

func (o *LobbyMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"content",
		"lobby_id",
		"channel_id",
		"author",
		"flags",
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

	varLobbyMessageResponse := _LobbyMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLobbyMessageResponse)

	if err != nil {
		return err
	}

	*o = LobbyMessageResponse(varLobbyMessageResponse)

	return err
}

type NullableLobbyMessageResponse struct {
	value *LobbyMessageResponse
	isSet bool
}

func (v NullableLobbyMessageResponse) Get() *LobbyMessageResponse {
	return v.value
}

func (v *NullableLobbyMessageResponse) Set(val *LobbyMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLobbyMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLobbyMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLobbyMessageResponse(val *LobbyMessageResponse) *NullableLobbyMessageResponse {
	return &NullableLobbyMessageResponse{value: val, isSet: true}
}

func (v NullableLobbyMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLobbyMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


