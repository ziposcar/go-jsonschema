// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package test
import "encoding/json"
import "fmt"
import "reflect"

// A simple schema.
type 101Description struct {
	// MyDescriptionlessField corresponds to the JSON schema field
	// "myDescriptionlessField".
	MyDescriptionlessField *string `json:"myDescriptionlessField,omitempty"`

	// A string field.
	MyField *string `json:"myField,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *101Description) UnmarshalJSON(b []byte) error {
	var v struct {
		101Description
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.101Description
	return nil
}


type 612EnumMyBooleanTypedEnumEnum bool
var enumValues_612EnumMyBooleanTypedEnumEnum  = []interface {}{
  true,
  false,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *612EnumMyBooleanTypedEnumEnum) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil { return err }
	var ok bool
	for _, expected := range enumValues_612EnumMyBooleanTypedEnumEnum {
	if reflect.DeepEqual(v, expected) { ok = true; break }
	}
	if !ok {
	return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_612EnumMyBooleanTypedEnumEnum, v)
	}
	*j = 612EnumMyBooleanTypedEnumEnum(v)
	return nil
}


type 612EnumMyMixedUntypedEnumEnum struct {
	Value interface{}
}
var enumValues_612EnumMyMixedUntypedEnumEnum  = []interface {}{
  "red",
  1,
  true,
  nil,
}

// MarshalJSON implements json.Marshaler.
func (j *612EnumMyMixedUntypedEnumEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}



// UnmarshalJSON implements json.Unmarshaler.
func (j *612EnumMyMixedUntypedEnumEnum) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil { return err }
	var ok bool
	for _, expected := range enumValues_612EnumMyMixedUntypedEnumEnum {
	if reflect.DeepEqual(v.Value, expected) { ok = true; break }
	}
	if !ok {
	return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_612EnumMyMixedUntypedEnumEnum, v.Value)
	}
	*j = 612EnumMyMixedUntypedEnumEnum(v)
	return nil
}


type 612EnumMyNullTypedEnumEnum struct {
	Value interface{}
}
var enumValues_612EnumMyNullTypedEnumEnum  = []interface {}{
  nil,
}

// MarshalJSON implements json.Marshaler.
func (j *612EnumMyNullTypedEnumEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}



// UnmarshalJSON implements json.Unmarshaler.
func (j *612EnumMyNullTypedEnumEnum) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil { return err }
	var ok bool
	for _, expected := range enumValues_612EnumMyNullTypedEnumEnum {
	if reflect.DeepEqual(v.Value, expected) { ok = true; break }
	}
	if !ok {
	return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_612EnumMyNullTypedEnumEnum, v.Value)
	}
	*j = 612EnumMyNullTypedEnumEnum(v)
	return nil
}


type 612EnumMyNumberTypedEnumEnum float64
var enumValues_612EnumMyNumberTypedEnumEnum  = []interface {}{
  1,
  2,
  3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *612EnumMyNumberTypedEnumEnum) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil { return err }
	var ok bool
	for _, expected := range enumValues_612EnumMyNumberTypedEnumEnum {
	if reflect.DeepEqual(v, expected) { ok = true; break }
	}
	if !ok {
	return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_612EnumMyNumberTypedEnumEnum, v)
	}
	*j = 612EnumMyNumberTypedEnumEnum(v)
	return nil
}


type 612EnumMyStringTypedEnumEnum string
var enumValues_612EnumMyStringTypedEnumEnum  = []interface {}{
  "red",
  "blue",
  "green",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *612EnumMyStringTypedEnumEnum) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil { return err }
	var ok bool
	for _, expected := range enumValues_612EnumMyStringTypedEnumEnum {
	if reflect.DeepEqual(v, expected) { ok = true; break }
	}
	if !ok {
	return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_612EnumMyStringTypedEnumEnum, v)
	}
	*j = 612EnumMyStringTypedEnumEnum(v)
	return nil
}


const 612EnumMyStringTypedEnumEnumRed 612EnumMyStringTypedEnumEnum = "red"
const 612EnumMyStringTypedEnumEnumBlue 612EnumMyStringTypedEnumEnum = "blue"
const 612EnumMyStringTypedEnumEnumGreen 612EnumMyStringTypedEnumEnum = "green"
type 612Enum struct {
	// MyBooleanTypedEnum corresponds to the JSON schema field "myBooleanTypedEnum".
	MyBooleanTypedEnum *612EnumMyBooleanTypedEnumEnum `json:"myBooleanTypedEnum,omitempty"`

	// MyBooleanUntypedEnum corresponds to the JSON schema field
	// "myBooleanUntypedEnum".
	MyBooleanUntypedEnum *612EnumMyBooleanTypedEnumEnum `json:"myBooleanUntypedEnum,omitempty"`

	// MyMixedUntypedEnum corresponds to the JSON schema field "myMixedUntypedEnum".
	MyMixedUntypedEnum *612EnumMyMixedUntypedEnumEnum `json:"myMixedUntypedEnum,omitempty"`

	// MyNullTypedEnum corresponds to the JSON schema field "myNullTypedEnum".
	MyNullTypedEnum *612EnumMyNullTypedEnumEnum `json:"myNullTypedEnum,omitempty"`

	// MyNullUntypedEnum corresponds to the JSON schema field "myNullUntypedEnum".
	MyNullUntypedEnum *612EnumMyNullTypedEnumEnum `json:"myNullUntypedEnum,omitempty"`

	// MyNumberTypedEnum corresponds to the JSON schema field "myNumberTypedEnum".
	MyNumberTypedEnum *612EnumMyNumberTypedEnumEnum `json:"myNumberTypedEnum,omitempty"`

	// MyNumberUntypedEnum corresponds to the JSON schema field "myNumberUntypedEnum".
	MyNumberUntypedEnum *612EnumMyNumberTypedEnumEnum `json:"myNumberUntypedEnum,omitempty"`

	// MyStringTypedEnum corresponds to the JSON schema field "myStringTypedEnum".
	MyStringTypedEnum *612EnumMyStringTypedEnumEnum `json:"myStringTypedEnum,omitempty"`

	// MyStringUntypedEnum corresponds to the JSON schema field "myStringUntypedEnum".
	MyStringUntypedEnum *612EnumMyStringTypedEnumEnum `json:"myStringUntypedEnum,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *612Enum) UnmarshalJSON(b []byte) error {
	var v struct {
		612Enum
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*j = v.612Enum
	return nil
}
