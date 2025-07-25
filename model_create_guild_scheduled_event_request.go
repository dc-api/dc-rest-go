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
	"fmt"
	"gopkg.in/validator.v2"
)

// CreateGuildScheduledEventRequest - struct for CreateGuildScheduledEventRequest
type CreateGuildScheduledEventRequest struct {
	ExternalScheduledEventCreateRequest *ExternalScheduledEventCreateRequest
	StageScheduledEventCreateRequest *StageScheduledEventCreateRequest
	VoiceScheduledEventCreateRequest *VoiceScheduledEventCreateRequest
}

// ExternalScheduledEventCreateRequestAsCreateGuildScheduledEventRequest is a convenience function that returns ExternalScheduledEventCreateRequest wrapped in CreateGuildScheduledEventRequest
func ExternalScheduledEventCreateRequestAsCreateGuildScheduledEventRequest(v *ExternalScheduledEventCreateRequest) CreateGuildScheduledEventRequest {
	return CreateGuildScheduledEventRequest{
		ExternalScheduledEventCreateRequest: v,
	}
}

// StageScheduledEventCreateRequestAsCreateGuildScheduledEventRequest is a convenience function that returns StageScheduledEventCreateRequest wrapped in CreateGuildScheduledEventRequest
func StageScheduledEventCreateRequestAsCreateGuildScheduledEventRequest(v *StageScheduledEventCreateRequest) CreateGuildScheduledEventRequest {
	return CreateGuildScheduledEventRequest{
		StageScheduledEventCreateRequest: v,
	}
}

// VoiceScheduledEventCreateRequestAsCreateGuildScheduledEventRequest is a convenience function that returns VoiceScheduledEventCreateRequest wrapped in CreateGuildScheduledEventRequest
func VoiceScheduledEventCreateRequestAsCreateGuildScheduledEventRequest(v *VoiceScheduledEventCreateRequest) CreateGuildScheduledEventRequest {
	return CreateGuildScheduledEventRequest{
		VoiceScheduledEventCreateRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateGuildScheduledEventRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExternalScheduledEventCreateRequest
	err = newStrictDecoder(data).Decode(&dst.ExternalScheduledEventCreateRequest)
	if err == nil {
		jsonExternalScheduledEventCreateRequest, _ := json.Marshal(dst.ExternalScheduledEventCreateRequest)
		if string(jsonExternalScheduledEventCreateRequest) == "{}" { // empty struct
			dst.ExternalScheduledEventCreateRequest = nil
		} else {
			if err = validator.Validate(dst.ExternalScheduledEventCreateRequest); err != nil {
				dst.ExternalScheduledEventCreateRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ExternalScheduledEventCreateRequest = nil
	}

	// try to unmarshal data into StageScheduledEventCreateRequest
	err = newStrictDecoder(data).Decode(&dst.StageScheduledEventCreateRequest)
	if err == nil {
		jsonStageScheduledEventCreateRequest, _ := json.Marshal(dst.StageScheduledEventCreateRequest)
		if string(jsonStageScheduledEventCreateRequest) == "{}" { // empty struct
			dst.StageScheduledEventCreateRequest = nil
		} else {
			if err = validator.Validate(dst.StageScheduledEventCreateRequest); err != nil {
				dst.StageScheduledEventCreateRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.StageScheduledEventCreateRequest = nil
	}

	// try to unmarshal data into VoiceScheduledEventCreateRequest
	err = newStrictDecoder(data).Decode(&dst.VoiceScheduledEventCreateRequest)
	if err == nil {
		jsonVoiceScheduledEventCreateRequest, _ := json.Marshal(dst.VoiceScheduledEventCreateRequest)
		if string(jsonVoiceScheduledEventCreateRequest) == "{}" { // empty struct
			dst.VoiceScheduledEventCreateRequest = nil
		} else {
			if err = validator.Validate(dst.VoiceScheduledEventCreateRequest); err != nil {
				dst.VoiceScheduledEventCreateRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.VoiceScheduledEventCreateRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ExternalScheduledEventCreateRequest = nil
		dst.StageScheduledEventCreateRequest = nil
		dst.VoiceScheduledEventCreateRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateGuildScheduledEventRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateGuildScheduledEventRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateGuildScheduledEventRequest) MarshalJSON() ([]byte, error) {
	if src.ExternalScheduledEventCreateRequest != nil {
		return json.Marshal(&src.ExternalScheduledEventCreateRequest)
	}

	if src.StageScheduledEventCreateRequest != nil {
		return json.Marshal(&src.StageScheduledEventCreateRequest)
	}

	if src.VoiceScheduledEventCreateRequest != nil {
		return json.Marshal(&src.VoiceScheduledEventCreateRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateGuildScheduledEventRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ExternalScheduledEventCreateRequest != nil {
		return obj.ExternalScheduledEventCreateRequest
	}

	if obj.StageScheduledEventCreateRequest != nil {
		return obj.StageScheduledEventCreateRequest
	}

	if obj.VoiceScheduledEventCreateRequest != nil {
		return obj.VoiceScheduledEventCreateRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateGuildScheduledEventRequest) GetActualInstanceValue() (interface{}) {
	if obj.ExternalScheduledEventCreateRequest != nil {
		return *obj.ExternalScheduledEventCreateRequest
	}

	if obj.StageScheduledEventCreateRequest != nil {
		return *obj.StageScheduledEventCreateRequest
	}

	if obj.VoiceScheduledEventCreateRequest != nil {
		return *obj.VoiceScheduledEventCreateRequest
	}

	// all schemas are nil
	return nil
}

type NullableCreateGuildScheduledEventRequest struct {
	value *CreateGuildScheduledEventRequest
	isSet bool
}

func (v NullableCreateGuildScheduledEventRequest) Get() *CreateGuildScheduledEventRequest {
	return v.value
}

func (v *NullableCreateGuildScheduledEventRequest) Set(val *CreateGuildScheduledEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildScheduledEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildScheduledEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildScheduledEventRequest(val *CreateGuildScheduledEventRequest) *NullableCreateGuildScheduledEventRequest {
	return &NullableCreateGuildScheduledEventRequest{value: val, isSet: true}
}

func (v NullableCreateGuildScheduledEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildScheduledEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


