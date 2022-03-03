| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `filename` | `string` |  | 否 |  | filename, if any |
| `offset` | `integer` | `int64` | 否 |  | offset, starting at 0 (byte count) |
| `line` | `integer` | `int64` | 否 |  | line number, starting at 1 |
| `column` | `integer` | `int64` | 否 |  | column number, starting at 1 (byte count) |
| `leadingComments` | `Array<mojo.lang.Comment>` |  | 否 |  | the comment before the AST node |
| `tailingComments` | `Array<mojo.lang.Comment>` |  | 否 |  | usually exist in the end position of the AST nodeand only one of the following situations:<br>```mojo
   a = 5 // following line comment 

```
 |
