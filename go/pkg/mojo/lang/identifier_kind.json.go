// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.
//
// Copyright 2021 Mojo-lang.org
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

package lang

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("lang.Identifier_Kind", &IdentifierKindCodec{})
	jsoniter.RegisterTypeEncoder("lang.Identifier_Kind", &IdentifierKindCodec{})
}

type IdentifierKindCodec struct {
}

func (codec *IdentifierKindCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	e := (*Identifier_Kind)(ptr)
	if any.ValueType() == jsoniter.StringValue {
		e.Parse(any.ToString())
	} else if any.ValueType() == jsoniter.NumberValue {
		value := any.ToInt32()
		if _, ok := IdentifierKindNames[value]; ok {
			*e = Identifier_Kind(value)
		}
	}
}

func (codec *IdentifierKindCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	e := (*Identifier_Kind)(ptr)
	stream.WriteString(e.Format())
}

func (codec *IdentifierKindCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Identifier_Kind)(ptr)
	return e == nil || *e == 0
}