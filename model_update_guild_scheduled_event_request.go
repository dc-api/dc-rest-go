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
	"fmt"
)


// UpdateGuildScheduledEventRequest struct for UpdateGuildScheduledEventRequest
type UpdateGuildScheduledEventRequest struct {
	ExternalScheduledEventPatchRequestPartial *ExternalScheduledEventPatchRequestPartial
	StageScheduledEventPatchRequestPartial *StageScheduledEventPatchRequestPartial
	VoiceScheduledEventPatchRequestPartial *VoiceScheduledEventPatchRequestPartial
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpdateGuildScheduledEventRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ExternalScheduledEventPatchRequestPartial
	err = json.Unmarshal(data, &dst.ExternalScheduledEventPatchRequestPartial);
	if err == nil {
		jsonExternalScheduledEventPatchRequestPartial, _ := json.Marshal(dst.ExternalScheduledEventPatchRequestPartial)
		if string(jsonExternalScheduledEventPatchRequestPartial) == "{}" { // empty struct
			dst.ExternalScheduledEventPatchRequestPartial = nil
		} else {
			return nil // data stored in dst.ExternalScheduledEventPatchRequestPartial, return on the first match
		}
	} else {
		dst.ExternalScheduledEventPatchRequestPartial = nil
	}

	// try to unmarshal JSON data into StageScheduledEventPatchRequestPartial
	err = json.Unmarshal(data, &dst.StageScheduledEventPatchRequestPartial);
	if err == nil {
		jsonStageScheduledEventPatchRequestPartial, _ := json.Marshal(dst.StageScheduledEventPatchRequestPartial)
		if string(jsonStageScheduledEventPatchRequestPartial) == "{}" { // empty struct
			dst.StageScheduledEventPatchRequestPartial = nil
		} else {
			return nil // data stored in dst.StageScheduledEventPatchRequestPartial, return on the first match
		}
	} else {
		dst.StageScheduledEventPatchRequestPartial = nil
	}

	// try to unmarshal JSON data into VoiceScheduledEventPatchRequestPartial
	err = json.Unmarshal(data, &dst.VoiceScheduledEventPatchRequestPartial);
	if err == nil {
		jsonVoiceScheduledEventPatchRequestPartial, _ := json.Marshal(dst.VoiceScheduledEventPatchRequestPartial)
		if string(jsonVoiceScheduledEventPatchRequestPartial) == "{}" { // empty struct
			dst.VoiceScheduledEventPatchRequestPartial = nil
		} else {
			return nil // data stored in dst.VoiceScheduledEventPatchRequestPartial, return on the first match
		}
	} else {
		dst.VoiceScheduledEventPatchRequestPartial = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UpdateGuildScheduledEventRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateGuildScheduledEventRequest) MarshalJSON() ([]byte, error) {
	if src.ExternalScheduledEventPatchRequestPartial != nil {
		return json.Marshal(&src.ExternalScheduledEventPatchRequestPartial)
	}

	if src.StageScheduledEventPatchRequestPartial != nil {
		return json.Marshal(&src.StageScheduledEventPatchRequestPartial)
	}

	if src.VoiceScheduledEventPatchRequestPartial != nil {
		return json.Marshal(&src.VoiceScheduledEventPatchRequestPartial)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUpdateGuildScheduledEventRequest struct {
	value *UpdateGuildScheduledEventRequest
	isSet bool
}

func (v NullableUpdateGuildScheduledEventRequest) Get() *UpdateGuildScheduledEventRequest {
	return v.value
}

func (v *NullableUpdateGuildScheduledEventRequest) Set(val *UpdateGuildScheduledEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildScheduledEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildScheduledEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildScheduledEventRequest(val *UpdateGuildScheduledEventRequest) *NullableUpdateGuildScheduledEventRequest {
	return &NullableUpdateGuildScheduledEventRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildScheduledEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildScheduledEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


