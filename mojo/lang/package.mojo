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

/// A Package represents a set of source files
/// collectively building a Mojo package.
type Package {
    ///
    ///
    type Requirement {
        type Version {
            enum Type {
                caret @0
                tilde @1
                wildcard @2
                comparison @3
            }
            type: Type @1
            range: VersionRange @2
        }

        /// ^1.2.3
        /// ~1.2.3
        /// >= 1.2, < 1.5
        version: Version @1

        ///
        registry: String @3
        
        ///
        path: String @4
        
        ///
        repository: Url @5
        
        ///
        branch: String @6
    }

    type Author {
        /// the people that are considered the "author" of the package.
        author: String @1
        
        /// the organization that are considered the "author" of the package.
        organization: String @2
        
        ///
        email: EmailAddress @3
    }

    /// package name
    name : String @1

    ///
    full_name: String @2

    ///
    version: Version @3

    ///
    parent: Package @4 @reference('full_name')

    ///
    children: [Package] @5

    //
    authors: [Author] @6

    summary: String @7

    /// A description of the package.
    description: String @10

    /// 
    document: Cached<document.Document> @11

    /// URL of the package source repository.
    repository: Url @12

    /// The package license.
    license: String @13

    //
    //imports: [Import] @14

    ///
    exports: [String] @15

    ///
    source_files: [SourceFile] @16

    /// package scope across all files
    scope: Scope @20

    extra_info: Object @25

    /// entity relationship analysis in the package
    entity_relationship: EntityRelationship @26

    ///
    dependencies: {String: Requirement} @30

    /// 
    resolved_dependencies: {String: Package} @31
}

//type ReferencedPackage = Referenced<Package>

func to<T:String>(package: Package) -> T {
    return package.full_name
}
