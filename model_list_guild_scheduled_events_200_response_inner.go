/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// ListGuildScheduledEvents200ResponseInner - struct for ListGuildScheduledEvents200ResponseInner
type ListGuildScheduledEvents200ResponseInner struct {
	ExternalScheduledEventResponse *ExternalScheduledEventResponse
	StageScheduledEventResponse *StageScheduledEventResponse
	VoiceScheduledEventResponse *VoiceScheduledEventResponse
}

// ExternalScheduledEventResponseAsListGuildScheduledEvents200ResponseInner is a convenience function that returns ExternalScheduledEventResponse wrapped in ListGuildScheduledEvents200ResponseInner
func ExternalScheduledEventResponseAsListGuildScheduledEvents200ResponseInner(v *ExternalScheduledEventResponse) ListGuildScheduledEvents200ResponseInner {
	return ListGuildScheduledEvents200ResponseInner{
		ExternalScheduledEventResponse: v,
	}
}

// StageScheduledEventResponseAsListGuildScheduledEvents200ResponseInner is a convenience function that returns StageScheduledEventResponse wrapped in ListGuildScheduledEvents200ResponseInner
func StageScheduledEventResponseAsListGuildScheduledEvents200ResponseInner(v *StageScheduledEventResponse) ListGuildScheduledEvents200ResponseInner {
	return ListGuildScheduledEvents200ResponseInner{
		StageScheduledEventResponse: v,
	}
}

// VoiceScheduledEventResponseAsListGuildScheduledEvents200ResponseInner is a convenience function that returns VoiceScheduledEventResponse wrapped in ListGuildScheduledEvents200ResponseInner
func VoiceScheduledEventResponseAsListGuildScheduledEvents200ResponseInner(v *VoiceScheduledEventResponse) ListGuildScheduledEvents200ResponseInner {
	return ListGuildScheduledEvents200ResponseInner{
		VoiceScheduledEventResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGuildScheduledEvents200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExternalScheduledEventResponse
	err = newStrictDecoder(data).Decode(&dst.ExternalScheduledEventResponse)
	if err == nil {
		jsonExternalScheduledEventResponse, _ := json.Marshal(dst.ExternalScheduledEventResponse)
		if string(jsonExternalScheduledEventResponse) == "{}" { // empty struct
			dst.ExternalScheduledEventResponse = nil
		} else {
			if err = validator.Validate(dst.ExternalScheduledEventResponse); err != nil {
				dst.ExternalScheduledEventResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ExternalScheduledEventResponse = nil
	}

	// try to unmarshal data into StageScheduledEventResponse
	err = newStrictDecoder(data).Decode(&dst.StageScheduledEventResponse)
	if err == nil {
		jsonStageScheduledEventResponse, _ := json.Marshal(dst.StageScheduledEventResponse)
		if string(jsonStageScheduledEventResponse) == "{}" { // empty struct
			dst.StageScheduledEventResponse = nil
		} else {
			if err = validator.Validate(dst.StageScheduledEventResponse); err != nil {
				dst.StageScheduledEventResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StageScheduledEventResponse = nil
	}

	// try to unmarshal data into VoiceScheduledEventResponse
	err = newStrictDecoder(data).Decode(&dst.VoiceScheduledEventResponse)
	if err == nil {
		jsonVoiceScheduledEventResponse, _ := json.Marshal(dst.VoiceScheduledEventResponse)
		if string(jsonVoiceScheduledEventResponse) == "{}" { // empty struct
			dst.VoiceScheduledEventResponse = nil
		} else {
			if err = validator.Validate(dst.VoiceScheduledEventResponse); err != nil {
				dst.VoiceScheduledEventResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.VoiceScheduledEventResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ExternalScheduledEventResponse = nil
		dst.StageScheduledEventResponse = nil
		dst.VoiceScheduledEventResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGuildScheduledEvents200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGuildScheduledEvents200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGuildScheduledEvents200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.ExternalScheduledEventResponse != nil {
		return json.Marshal(&src.ExternalScheduledEventResponse)
	}

	if src.StageScheduledEventResponse != nil {
		return json.Marshal(&src.StageScheduledEventResponse)
	}

	if src.VoiceScheduledEventResponse != nil {
		return json.Marshal(&src.VoiceScheduledEventResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGuildScheduledEvents200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ExternalScheduledEventResponse != nil {
		return obj.ExternalScheduledEventResponse
	}

	if obj.StageScheduledEventResponse != nil {
		return obj.StageScheduledEventResponse
	}

	if obj.VoiceScheduledEventResponse != nil {
		return obj.VoiceScheduledEventResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListGuildScheduledEvents200ResponseInner) GetActualInstanceValue() (interface{}) {
	if obj.ExternalScheduledEventResponse != nil {
		return *obj.ExternalScheduledEventResponse
	}

	if obj.StageScheduledEventResponse != nil {
		return *obj.StageScheduledEventResponse
	}

	if obj.VoiceScheduledEventResponse != nil {
		return *obj.VoiceScheduledEventResponse
	}

	// all schemas are nil
	return nil
}

type NullableListGuildScheduledEvents200ResponseInner struct {
	value *ListGuildScheduledEvents200ResponseInner
	isSet bool
}

func (v NullableListGuildScheduledEvents200ResponseInner) Get() *ListGuildScheduledEvents200ResponseInner {
	return v.value
}

func (v *NullableListGuildScheduledEvents200ResponseInner) Set(val *ListGuildScheduledEvents200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListGuildScheduledEvents200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListGuildScheduledEvents200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGuildScheduledEvents200ResponseInner(val *ListGuildScheduledEvents200ResponseInner) *NullableListGuildScheduledEvents200ResponseInner {
	return &NullableListGuildScheduledEvents200ResponseInner{value: val, isSet: true}
}

func (v NullableListGuildScheduledEvents200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGuildScheduledEvents200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


