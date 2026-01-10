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
	"fmt"
	"gopkg.in/validator.v2"
)

// PollCreateRequestLayoutType - struct for PollCreateRequestLayoutType
type PollCreateRequestLayoutType struct {
	Int32 *int32
}

// int32AsPollCreateRequestLayoutType is a convenience function that returns int32 wrapped in PollCreateRequestLayoutType
func Int32AsPollCreateRequestLayoutType(v *int32) PollCreateRequestLayoutType {
	return PollCreateRequestLayoutType{
		Int32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PollCreateRequestLayoutType) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PollCreateRequestLayoutType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PollCreateRequestLayoutType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PollCreateRequestLayoutType) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PollCreateRequestLayoutType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj PollCreateRequestLayoutType) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullablePollCreateRequestLayoutType struct {
	value *PollCreateRequestLayoutType
	isSet bool
}

func (v NullablePollCreateRequestLayoutType) Get() *PollCreateRequestLayoutType {
	return v.value
}

func (v *NullablePollCreateRequestLayoutType) Set(val *PollCreateRequestLayoutType) {
	v.value = val
	v.isSet = true
}

func (v NullablePollCreateRequestLayoutType) IsSet() bool {
	return v.isSet
}

func (v *NullablePollCreateRequestLayoutType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollCreateRequestLayoutType(val *PollCreateRequestLayoutType) *NullablePollCreateRequestLayoutType {
	return &NullablePollCreateRequestLayoutType{value: val, isSet: true}
}

func (v NullablePollCreateRequestLayoutType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollCreateRequestLayoutType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

