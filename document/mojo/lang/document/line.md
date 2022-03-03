| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | 否 |  | position of first character belonging to the LineDocument |
| `endPosition` | `mojo.lang.Position` |  | 否 |  | position of first character immediately after the LineDocument |
| `following` | `boolean` |  | 否 |  | following after the decl, but not at the beginning of the line |
| `private` | `boolean` |  | 否 |  | private or inner document, which will not to publish to public |
| `text` | `string` |  | 否 |  | the content of the line comment, remove the prefix and the EOL |
