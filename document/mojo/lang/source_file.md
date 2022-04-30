| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  |
| `fullName` | `string` |  | N |  |
| `url` | `string` | `Url` | N |  | URl type<br>Uniform Resource Identifier (URI) is a string of characters used to identify a resource.<br>A Uniform Resource Name (URN) may be compared to a person's name,while a Uniform Resource Locator (URL) may be compared to their street address.In other words, a URN identifies an item and a URL provides a method for finding it.<br> |
| `attributes` | `Array<mojo.lang.Attribute>` |  | N |  | 循环依赖文件名称 'circle_dependency' |
| `packageName` | `string` |  | N |  |
| `imports` | `Array<mojo.lang.ImportDecl>` |  | N |  | imports decl index for this file, will still exist in the statements |
| `statements` | `Array<mojo.lang.Statement>` |  | N |  | statements in this file |
| `tailingComments` | `Array<mojo.lang.Comment>` |  | N |  | comments after all statements but before the end of the source file |
| `resolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | N |  | resolved identifiers not in this source file |
| `unresolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | N |  | unresolved identifiers in this file |
| `dependencies` | `Array<mojo.lang.SourceFile>` |  | N |  |  |
| `scope` | `mojo.lang.Scope` |  | N |  | package scope (this file only) |
