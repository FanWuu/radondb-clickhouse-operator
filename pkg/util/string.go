// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

// RandomString generates random string
func RandomString() string {
	b := make([]byte, 3)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)
	return hex.EncodeToString(b)
}

const (
	StringBool0        = "0"
	StringBool1        = "1"
	StringBoolFalse    = "false"
	StringBoolTrue     = "true"
	StringBoolNo       = "no"
	StringBoolYes      = "yes"
	StringBoolOff      = "off"
	StringBoolOn       = "on"
	StringBoolDisabled = "disabled"
	StringBoolEnabled  = "enabled"
)

// IsStringBool checks whether str is a string as bool value
func IsStringBool(str string) bool {
	switch strings.ToLower(str) {
	case
		StringBool0,
		StringBool1,

		StringBoolFalse,
		StringBoolTrue,

		StringBoolNo,
		StringBoolYes,

		StringBoolOff,
		StringBoolOn,

		StringBoolDisabled,
		StringBoolEnabled:
		return true

	default:
		return false
	}
}

// IsStringBool checks whether str is a string as bool "false" value
func IsStringBoolFalse(str string) bool {
	switch strings.ToLower(str) {
	case
		StringBool0,
		StringBoolFalse,
		StringBoolNo,
		StringBoolOff,
		StringBoolDisabled:
		return true

	default:
		return false
	}
}

// IsStringBool checks whether str is a string as bool "true" value
func IsStringBoolTrue(str string) bool {
	switch strings.ToLower(str) {
	case
		StringBool1,
		StringBoolTrue,
		StringBoolYes,
		StringBoolOn,
		StringBoolEnabled:
		return true

	default:
		return false
	}
}

// CastStringBoolTo01 casts string-bool into string "0/1"
func CastStringBoolTo01(str string, defaultValue bool) string {
	// True and False values
	t := StringBool1
	f := StringBool0

	if IsStringBoolTrue(str) {
		return t
	}
	if IsStringBoolFalse(str) {
		return f
	}

	// String value unrecognized, return default value

	if defaultValue {
		return t
	} else {
		return f
	}
}

// CastStringBoolToTrueFalse casts string-bool into string "true/false"
func CastStringBoolToTrueFalse(str string, defaultValue bool) string {
	// True and False values
	t := StringBoolTrue
	f := StringBoolFalse

	if IsStringBoolTrue(str) {
		return t
	}
	if IsStringBoolFalse(str) {
		return f
	}

	// String value unrecognized, return default value

	if defaultValue {
		return t
	} else {
		return f
	}
}

// CreateStringID creates HEX hash ID out of string.
// In case hashLen == 0 the whole hash is returned
func CreateStringID(str string, hashLen int) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	hash := hex.EncodeToString(hasher.Sum(nil))
	if hashLen == 0 {
		return hash
	} else {
		return hash[len(hash)-hashLen:]
	}
}

func StringHead(str string, headLen int) string {
	if len(str) <= headLen {
		return str
	} else {
		return str[:headLen]
	}

}
