/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
)

// checks if the WebhookSlackEmbedField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSlackEmbedField{}

// WebhookSlackEmbedField struct for WebhookSlackEmbedField
type WebhookSlackEmbedField struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Inline *bool `json:"inline,omitempty"`
}

// NewWebhookSlackEmbedField instantiates a new WebhookSlackEmbedField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSlackEmbedField() *WebhookSlackEmbedField {
	this := WebhookSlackEmbedField{}
	return &this
}

// NewWebhookSlackEmbedFieldWithDefaults instantiates a new WebhookSlackEmbedField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSlackEmbedFieldWithDefaults() *WebhookSlackEmbedField {
	this := WebhookSlackEmbedField{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookSlackEmbedField) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbedField) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookSlackEmbedField) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookSlackEmbedField) SetName(v string) {
	o.Name = &v
}


// GetValue returns the Value field value if set, zero value otherwise.
func (o *WebhookSlackEmbedField) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbedField) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WebhookSlackEmbedField) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WebhookSlackEmbedField) SetValue(v string) {
	o.Value = &v
}


// GetInline returns the Inline field value if set, zero value otherwise.
func (o *WebhookSlackEmbedField) GetInline() bool {
	if o == nil || IsNil(o.Inline) {
		var ret bool
		return ret
	}
	return *o.Inline
}

// GetInlineOk returns a tuple with the Inline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSlackEmbedField) GetInlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Inline) {
		return nil, false
	}
	return o.Inline, true
}

// HasInline returns a boolean if a field has been set.
func (o *WebhookSlackEmbedField) HasInline() bool {
	if o != nil && !IsNil(o.Inline) {
		return true
	}

	return false
}

// SetInline gets a reference to the given bool and assigns it to the Inline field.
func (o *WebhookSlackEmbedField) SetInline(v bool) {
	o.Inline = &v
}


func (o WebhookSlackEmbedField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSlackEmbedField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Inline) {
		toSerialize["inline"] = o.Inline
	}
	return toSerialize, nil
}

type NullableWebhookSlackEmbedField struct {
	value *WebhookSlackEmbedField
	isSet bool
}

func (v NullableWebhookSlackEmbedField) Get() *WebhookSlackEmbedField {
	return v.value
}

func (v *NullableWebhookSlackEmbedField) Set(val *WebhookSlackEmbedField) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSlackEmbedField) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSlackEmbedField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSlackEmbedField(val *WebhookSlackEmbedField) *NullableWebhookSlackEmbedField {
	return &NullableWebhookSlackEmbedField{value: val, isSet: true}
}

func (v NullableWebhookSlackEmbedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSlackEmbedField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

