| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `startPosition` | `mojo.lang.Position` |  | N |  | position of first character belonging to the Expr |
| `endPosition` | `mojo.lang.Position` |  | N |  | position of first character immediately after the Expr |
| `kind` | `integer` | `Int8` | N |  |  |
| `implicit` | `boolean` |  | N |  | Whether the Expr represents something directly written in source orit was implicitly generated by the type-checker. |
