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
type Decl {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position:   Position @2

    /// Whether the Decl represents something directly written in source or
    /// it was implicitly generated by the type-checker.
    implicit: Bool @5

    ///
    document : Document @6

    /// package full name which the decl in
    package_name: String @7

    /// source file full name which the decl in
    source_file_name: String @8
}
