/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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
	"fmt"
	"gopkg.in/validator.v2"
)

// ExternalScheduledEventPatchRequestPartialStatus - struct for ExternalScheduledEventPatchRequestPartialStatus
type ExternalScheduledEventPatchRequestPartialStatus struct {
	GuildScheduledEventStatuses *GuildScheduledEventStatuses
}

// GuildScheduledEventStatusesAsExternalScheduledEventPatchRequestPartialStatus is a convenience function that returns GuildScheduledEventStatuses wrapped in ExternalScheduledEventPatchRequestPartialStatus
func GuildScheduledEventStatusesAsExternalScheduledEventPatchRequestPartialStatus(v *GuildScheduledEventStatuses) ExternalScheduledEventPatchRequestPartialStatus {
	return ExternalScheduledEventPatchRequestPartialStatus{
		GuildScheduledEventStatuses: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExternalScheduledEventPatchRequestPartialStatus) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into GuildScheduledEventStatuses
	err = newStrictDecoder(data).Decode(&dst.GuildScheduledEventStatuses)
	if err == nil {
		jsonGuildScheduledEventStatuses, _ := json.Marshal(dst.GuildScheduledEventStatuses)
		if string(jsonGuildScheduledEventStatuses) == "{}" { // empty struct
			dst.GuildScheduledEventStatuses = nil
		} else {
			if err = validator.Validate(dst.GuildScheduledEventStatuses); err != nil {
				dst.GuildScheduledEventStatuses = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildScheduledEventStatuses = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GuildScheduledEventStatuses = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ExternalScheduledEventPatchRequestPartialStatus)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExternalScheduledEventPatchRequestPartialStatus)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExternalScheduledEventPatchRequestPartialStatus) MarshalJSON() ([]byte, error) {
	if src.GuildScheduledEventStatuses != nil {
		return json.Marshal(&src.GuildScheduledEventStatuses)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExternalScheduledEventPatchRequestPartialStatus) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GuildScheduledEventStatuses != nil {
		return obj.GuildScheduledEventStatuses
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ExternalScheduledEventPatchRequestPartialStatus) GetActualInstanceValue() (interface{}) {
	if obj.GuildScheduledEventStatuses != nil {
		return *obj.GuildScheduledEventStatuses
	}

	// all schemas are nil
	return nil
}

type NullableExternalScheduledEventPatchRequestPartialStatus struct {
	value *ExternalScheduledEventPatchRequestPartialStatus
	isSet bool
}

func (v NullableExternalScheduledEventPatchRequestPartialStatus) Get() *ExternalScheduledEventPatchRequestPartialStatus {
	return v.value
}

func (v *NullableExternalScheduledEventPatchRequestPartialStatus) Set(val *ExternalScheduledEventPatchRequestPartialStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalScheduledEventPatchRequestPartialStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalScheduledEventPatchRequestPartialStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalScheduledEventPatchRequestPartialStatus(val *ExternalScheduledEventPatchRequestPartialStatus) *NullableExternalScheduledEventPatchRequestPartialStatus {
	return &NullableExternalScheduledEventPatchRequestPartialStatus{value: val, isSet: true}
}

func (v NullableExternalScheduledEventPatchRequestPartialStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalScheduledEventPatchRequestPartialStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

