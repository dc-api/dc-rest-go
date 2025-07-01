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
	"fmt"
	"gopkg.in/validator.v2"
)

// BasicMessageResponseNonce - struct for BasicMessageResponseNonce
type BasicMessageResponseNonce struct {
	Int64 *int64
	String *string
}

// int64AsBasicMessageResponseNonce is a convenience function that returns int64 wrapped in BasicMessageResponseNonce
func Int64AsBasicMessageResponseNonce(v *int64) BasicMessageResponseNonce {
	return BasicMessageResponseNonce{
		Int64: v,
	}
}

// stringAsBasicMessageResponseNonce is a convenience function that returns string wrapped in BasicMessageResponseNonce
func StringAsBasicMessageResponseNonce(v *string) BasicMessageResponseNonce {
	return BasicMessageResponseNonce{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BasicMessageResponseNonce) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			if err = validator.Validate(dst.Int64); err != nil {
				dst.Int64 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BasicMessageResponseNonce)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BasicMessageResponseNonce)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BasicMessageResponseNonce) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BasicMessageResponseNonce) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BasicMessageResponseNonce) GetActualInstanceValue() (interface{}) {
	if obj.Int64 != nil {
		return *obj.Int64
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableBasicMessageResponseNonce struct {
	value *BasicMessageResponseNonce
	isSet bool
}

func (v NullableBasicMessageResponseNonce) Get() *BasicMessageResponseNonce {
	return v.value
}

func (v *NullableBasicMessageResponseNonce) Set(val *BasicMessageResponseNonce) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicMessageResponseNonce) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicMessageResponseNonce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicMessageResponseNonce(val *BasicMessageResponseNonce) *NullableBasicMessageResponseNonce {
	return &NullableBasicMessageResponseNonce{value: val, isSet: true}
}

func (v NullableBasicMessageResponseNonce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicMessageResponseNonce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


