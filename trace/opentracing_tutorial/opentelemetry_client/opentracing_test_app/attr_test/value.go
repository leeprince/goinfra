// Copyright The OpenTelemetry Authors
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

package attr_test // import "go.opentelemetry.io/otel/attribute"

/*
go get golang.org/x/tools/cmd/stringer
	$ stringer -help
*/
//go:generate stringer -type=Type
//go:generate stringer -type=Type -linecomment

// Type describes the type of the data Value holds.
type Type int

// const (
// 	// INVALID is used for a Value with no value set.
// 	INVALID Type = iota
// 	// BOOL is a boolean Type Value.
// 	BOOL
// 	// INT64 is a 64-bit signed integral Type Value.
// 	INT64
// 	// FLOAT64 is a 64-bit floating point Type Value.
// 	FLOAT64
// 	// STRING is a string Type Value.
// 	STRING
// 	// BOOLSLICE is a slice of booleans Type Value.
// 	BOOLSLICE
// 	// INT64SLICE is a slice of 64-bit signed integral numbers Type Value.
// 	INT64SLICE
// 	// FLOAT64SLICE is a slice of 64-bit floating point numbers Type Value.
// 	FLOAT64SLICE
// 	// STRINGSLICE is a slice of strings Type Value.
// 	STRINGSLICE
// )

const (
	// INVALID is used for a Value with no value set.
	INVALID  Type = iota // prince-INVALID
	// BOOL is a boolean Type Value.
	BOOL // prince-BOOL
	// INT64 is a 64-bit signed integral Type Value.
	INT64 // prince-INT64
	// FLOAT64 is a 64-bit floating point Type Value.
	FLOAT64 // prince-FLOAT64
	// STRING is a string Type Value.
	STRING // prince-STRING
	// BOOLSLICE is a slice of booleans Type Value.
	BOOLSLICE // prince-BOOLSLICE
	// INT64SLICE is a slice of 64-bit signed integral numbers Type Value.
	INT64SLICE // prince-INT64SLICE
	// FLOAT64SLICE is a slice of 64-bit floating point numbers Type Value.
	FLOAT64SLICE // prince-FLOAT64SLICE
	// STRINGSLICE is a slice of strings Type Value.
	STRINGSLICE // prince-STRINGSLICE
)
