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

type Identifier {
    @case_style("kebab")
    enum Kind {
        unspecified @0
        package @1
        enum @3
        struct @4
        type_alias @5
        interface @6
        constant @10
        variable @11
        attribute @12
        attribute_alias @13
        function @14
        generic_parameter @19
    }

    start_position: Position @1 // identifier position

    end_position: Position @2 // identifier position

    kind: Kind @3

    /// Whether the Decl represents something directly written in source or
    /// it was implicitly generated by the type-checker.
    implicit: Bool @5

    source_file_name: String @8

    package_name: String @7
    
    name: String @10

    full_name: String @11

    alias: String @12

    declaration: Declaration @13 @reference @serialization(false)

    enclosing: Identifier @14

    scope: Scope @30
}

func is_type(ik: Identifier.Kind) -> Bool
func is_value(ik: Identifier.Kind) -> Bool