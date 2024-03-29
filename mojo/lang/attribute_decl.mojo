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

///
///
type AttributeDecl : Decl {
    ///
    name: String @10

    /// 
    full_name: String @11

    ///
    attributes: [Attribute] @12

    ///
    generic_parameters: [GenericParameter] @13

    ///
    group: GroupDecl @15

    ///
    resolved_identifiers: [Identifier] @16

    ///
    unresolved_identifiers:  [Identifier] @17 //< unresolved identifiers in this file

    /// 
    name_position: Position @19 @ignore

    ///
    type: NominalType @20 | StructType @21

    ///
    default_value: Expression @22

    /// the 'generic-parameter' type identifier will be in here
    scope:       Scope @30
}
