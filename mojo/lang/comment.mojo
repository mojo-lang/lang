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

@discriminator('@type')
@label_format('{}')
type Comment = BlockComment @1 //< block comment using /**/
             | MultiLineComment @2 //< continuous multi-line comment
             | LineBreakComment @3 //< user line break comment
             | Document @4 //< free floating document

type BlockComment {
    /// position of first character belonging to the Document
    start_position: Position @1

    /// position of first character immediately after the Document
    end_position: Position @2

    /// the content of the block comment
    text: String @10

    /// the head /* is after some code in the same line
    head_embeded: Bool @13

    /// the tail */ is before some code in the same line
    tail_embeded: Bool @14
}

type LineComment {
    /// position of first character belonging to the LineComment
    start_position: Position @1

    /// position of first character immediately after the LineComment
    end_position: Position @2

    following: Bool @3

    text: String @10
}

type MultiLineComment {
    /// position of first character belonging to the MultiLineComment
    start_position: Position @1

    /// position of first character immediately after the MultiLineComment
    end_position: Position @2

    lines: [LineComment] @10
}

/// user line break before a ValueDecl in struct or enum scope, or statement in function/lamda scope
type LineBreakComment {
    /// position of first character belonging to the LineBreakComment
    start_position: Position @1

    /// position of first character immediately after the LineBreakComment
    end_position: Position @2
}
