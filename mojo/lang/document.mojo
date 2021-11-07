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

type Document {
    /// position of first character belonging to the Document
    start_position: Position @1

    /// position of first character immediately after the Document
    end_position: Position @2

    lines: [LineDocument] @10

    structured: document.Document @15
}

type LineDocument {
    /// position of first character belonging to the LineDocument
    start_position: Position @1

    /// position of first character immediately after the LineDocument
    end_position: Position @2

    following: Bool @3

    line: String @10
}
