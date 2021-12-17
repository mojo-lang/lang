| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | 否 |  | position of first character belonging to the Expr |
| `endPosition` | `mojo.lang.Position` |  | 否 |  | position of first character immediately after the Expr |
| `document` | `mojo.lang.Document` |  | 否 |  |  |
| `packageName` | `string` |  | 否 |  | user writing package namevar foo : package.name.Bar==> mojo.package.name.Bar |
| `name` | `string` |  | 否 |  |  |
| `typeDeclaration` | `mojo.lang.TypeDeclaration` |  | 否 |  |  |
| `genericArguments` | `Array<mojo.lang.NominalType>` |  | 否 |  |  |
| `attributes` | `Array<mojo.lang.Attribute>` |  | 否 |  |  |
| `enclosingType` | `mojo.lang.NominalType` |  | 否 |  |  |
