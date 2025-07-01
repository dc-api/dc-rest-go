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

// ErrorDetails - struct for ErrorDetails
type ErrorDetails struct {
	InnerErrors *InnerErrors
	MapmapOfStringErrorDetails *map[string]ErrorDetails
}

// InnerErrorsAsErrorDetails is a convenience function that returns InnerErrors wrapped in ErrorDetails
func InnerErrorsAsErrorDetails(v *InnerErrors) ErrorDetails {
	return ErrorDetails{
		InnerErrors: v,
	}
}

// map[string]ErrorDetailsAsErrorDetails is a convenience function that returns map[string]ErrorDetails wrapped in ErrorDetails
func MapmapOfStringErrorDetailsAsErrorDetails(v *map[string]ErrorDetails) ErrorDetails {
	return ErrorDetails{
		MapmapOfStringErrorDetails: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ErrorDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InnerErrors
	err = newStrictDecoder(data).Decode(&dst.InnerErrors)
	if err == nil {
		jsonInnerErrors, _ := json.Marshal(dst.InnerErrors)
		if string(jsonInnerErrors) == "{}" { // empty struct
			dst.InnerErrors = nil
		} else {
			if err = validator.Validate(dst.InnerErrors); err != nil {
				dst.InnerErrors = nil
			} else {
				match++
			}
		}
	} else {
		dst.InnerErrors = nil
	}

	// try to unmarshal data into MapmapOfStringErrorDetails
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringErrorDetails)
	if err == nil {
		jsonMapmapOfStringErrorDetails, _ := json.Marshal(dst.MapmapOfStringErrorDetails)
		if string(jsonMapmapOfStringErrorDetails) == "{}" { // empty struct
			dst.MapmapOfStringErrorDetails = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringErrorDetails); err != nil {
				dst.MapmapOfStringErrorDetails = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringErrorDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InnerErrors = nil
		dst.MapmapOfStringErrorDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ErrorDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ErrorDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ErrorDetails) MarshalJSON() ([]byte, error) {
	if src.InnerErrors != nil {
		return json.Marshal(&src.InnerErrors)
	}

	if src.MapmapOfStringErrorDetails != nil {
		return json.Marshal(&src.MapmapOfStringErrorDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ErrorDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InnerErrors != nil {
		return obj.InnerErrors
	}

	if obj.MapmapOfStringErrorDetails != nil {
		return obj.MapmapOfStringErrorDetails
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ErrorDetails) GetActualInstanceValue() (interface{}) {
	if obj.InnerErrors != nil {
		return *obj.InnerErrors
	}

	if obj.MapmapOfStringErrorDetails != nil {
		return *obj.MapmapOfStringErrorDetails
	}

	// all schemas are nil
	return nil
}

type NullableErrorDetails struct {
	value *ErrorDetails
	isSet bool
}

func (v NullableErrorDetails) Get() *ErrorDetails {
	return v.value
}

func (v *NullableErrorDetails) Set(val *ErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetails(val *ErrorDetails) *NullableErrorDetails {
	return &NullableErrorDetails{value: val, isSet: true}
}

func (v NullableErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


