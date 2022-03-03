| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  |
| `fullName` | `string` |  | 否 |  |
| `url` | `string` | `url` | 否 |  |
| `attributes` | `Array<mojo.lang.Attribute>` |  | 否 |  | 循环依赖文件名称 'circle_dependency' |
| `packageName` | `string` |  | 否 |  |
| `imports` | `Array<mojo.lang.ImportDecl>` |  | 否 |  | imports decl index for this file, will still exist in the statements |
| `statements` | `Array<mojo.lang.Statement>` |  | 否 |  | statements in this file |
| `tailingComments` | `Array<mojo.lang.Comment>` |  | 否 |  | comments after all statements but before the end of the source file |
| `resolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | 否 |  | resolved identifiers not in this source file |
| `unresolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | 否 |  | unresolved identifiers in this file |
| `dependencies` | `Array<mojo.lang.SourceFile>` |  | 否 |  |  |
| `scope` | `mojo.lang.Scope` |  | 否 |  | package scope (this file only) |
