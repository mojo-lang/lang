| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  |  |
| `fullName` | `string` |  | 否 |  |  |
| `attributes` | `Array<mojo.lang.Attribute>` |  | 否 |  |  |
| `genericParameters` | `Array<mojo.lang.GenericParameter>` |  | 否 |  |  |
| `group` | `mojo.lang.GroupDecl` |  | 否 |  |  |
| `resolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | 否 |  |  |
| `unresolvedIdentifiers` | `Array<mojo.lang.Identifier>` |  | 否 |  | unresolved identifiers in this file |
| `namePosition` | `mojo.lang.Position` |  | 否 |  |  |
| `attribute` | `mojo.lang.Attribute` |  | 否 |  |  |
| `scope` | `mojo.lang.Scope` |  | 否 |  | the 'generic-parameter' type identifier will be in here |
