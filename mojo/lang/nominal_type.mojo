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
type NominalType {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position: Position @2

    ///
    document: Document @4

    /// user writing package name
    /// var foo : package.name.Bar
    /// ==> mojo.package.name.Bar
    package_name: String @5

    ///
    name: String @10

    ///
    type_declaration: TypeDeclaration @11 @reference

    ///
    generic_arguments: [GenericArgument] @12

    ///
    attributes: [Attribute] @13

    ///
    enclosing_type: NominalType @14
}

//func enclosing_types(NominalType type) -> [NominalType] {
//    var enclosing_types
//    while enclosing = type.enclosing_type {
//        enclosing_types.append(enclosing)
//        type = enclosing
//    }
//    return enclosing_types.reverse
//}