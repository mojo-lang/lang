| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | 否 |  | position of first character belonging to the Document |
| `endPosition` | `mojo.lang.Position` |  | 否 |  | position of first character immediately after the Document |
| `following` | `boolean` |  | 否 |  | all the lines are following after the decl, but not at the beginning of the line |
| `private` | `boolean` |  | 否 |  | private or inner document, which will not to publish to public |
| `lines` | `Array<mojo.lang.Document.Line>` |  | 否 |  | original line documents |
| `structured` | `mojo.document.Document` |  | 否 |  | structured document for all the lines |
