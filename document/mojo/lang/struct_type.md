| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | 否 |  | position of the character "{" belonging to the StructType |
| `endPosition` | `mojo.lang.Position` |  | 否 |  | position of the character "}" |
| `fields` | `Array<mojo.lang.ValueDecl>` |  | 否 |  | A Field represents a Field declaration list in a struct type,a method list in an interface type, or a parameter/result declarationin a signature. |
| `inherits` | `Array<mojo.lang.NominalType>` |  | 否 |  |  |
| `groups` | `Array<mojo.lang.ValueGroupDecl>` |  | 否 |  |  |
| `inheritePosition` | `mojo.lang.Position` |  | 否 |  |  |
