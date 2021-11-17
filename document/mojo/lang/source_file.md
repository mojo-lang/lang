| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` | `` | 否 |  |
| `fullName` | `string` | `` | 否 |  |
| `url` | `string` | `url` | 否 |  |
| `attributes` | `Array<mojo.lang.Attribute>` | `` | 否 |  | 循环依赖文件名称 'circle_dependency' |
| `packageName` | `string` | `` | 否 |  |
| `imports` | `Array<unspecified>` | `` | 否 |  | imports in this file |
| `statements` | `Array<unspecified>` | `` | 否 |  |
| `resolvedIdentifiers` | `Array<mojo.lang.Identifier>` | `` | 否 |  |
| `unresolvedIdentifiers` | `Array<mojo.lang.Identifier>` | `` | 否 |  | unresolved identifiers in this file |
| `dependencies` | `Array<mojo.lang.SourceFile>` | `` | 否 |  |  |
| `scope` | `mojo.lang.Scope` | `` | 否 |  | package scope (this file only) |
