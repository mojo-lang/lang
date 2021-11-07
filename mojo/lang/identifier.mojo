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

// An Object describes a named language entity such as a package,
// constant, type, variable, function (incl. methods), or label.
//
// The Data fields contains object-specific data:
//
//	Kind    Data type         Data value
//	Pkg     *Scope            package scope
//	Con     int               iota for the respective declaration
//
type Identifier {
    enum Kind {
        unspecified @0
        enum @3
        struct @4
        type_alias @5 @alias("type-alias")
        interface @6
        constant @10
        variable @11
        attribute @12
        function @13
        generic_parameter @14 @alias("generic-parameter")

        type @20
        value @21
    }

    start_position: Position @1 // identifier position

    end_position: Position @2 // identifier position

    kind: Kind @3

    package: String @5
    
    source_file: String @6

    name: String @10

    declaration: Declaration @11 @reference @serialization(false)

    alias: String @12

    full_name: String @20

    enclosing_type_names: [String] @21

    //domain: Scope @11
}
