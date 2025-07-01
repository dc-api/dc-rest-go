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
	"gopkg.in/validator.v2"
)

// ThreadSearchTagParameter - struct for ThreadSearchTagParameter
type ThreadSearchTagParameter struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsThreadSearchTagParameter is a convenience function that returns []string wrapped in ThreadSearchTagParameter
func ArrayOfStringAsThreadSearchTagParameter(v *[]string) ThreadSearchTagParameter {
	return ThreadSearchTagParameter{
		ArrayOfString: v,
	}
}

// stringAsThreadSearchTagParameter is a convenience function that returns string wrapped in ThreadSearchTagParameter
func StringAsThreadSearchTagParameter(v *string) ThreadSearchTagParameter {
	return ThreadSearchTagParameter{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThreadSearchTagParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
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
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ThreadSearchTagParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ThreadSearchTagParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThreadSearchTagParameter) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThreadSearchTagParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ThreadSearchTagParameter) GetActualInstanceValue() (interface{}) {
	if obj.ArrayOfString != nil {
		return *obj.ArrayOfString
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableThreadSearchTagParameter struct {
	value *ThreadSearchTagParameter
	isSet bool
}

func (v NullableThreadSearchTagParameter) Get() *ThreadSearchTagParameter {
	return v.value
}

func (v *NullableThreadSearchTagParameter) Set(val *ThreadSearchTagParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadSearchTagParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadSearchTagParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadSearchTagParameter(val *ThreadSearchTagParameter) *NullableThreadSearchTagParameter {
	return &NullableThreadSearchTagParameter{value: val, isSet: true}
}

func (v NullableThreadSearchTagParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadSearchTagParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


