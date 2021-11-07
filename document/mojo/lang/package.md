| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `parent` | `0` |  | 否 |  |  |
| `document` | `mojo.lang.CachedDocument` |  | 否 |  |  |
| `exports` | `Array<string>` |  | 否 |  |  |
| `sourceFiles` | `Array<mojo.lang.SourceFile>` |  | 否 |  |  |
| `scope` | `mojo.lang.Scope` |  | 否 |  | package scope across all files |
| `fullName` | `string` |  | 否 |  |  |
| `description` | `string` |  | 否 |  | A description of the package. |
| `repository` | `string` | `url` | 否 |  | URL of the package source repository. |
| `dependencies` |  |  | 否 |  |  |
| `resolvedDependencies` |  |  | 否 |  |  |
| `name` | `string` |  | 否 |  | package name |
| `version` | `mojo.core.Version` |  | 否 |  |  |
| `authors` | `Array<mojo.lang.Package.Author>` |  | 否 |  |
| `summary` | `string` |  | 否 |  |
| `license` | `string` |  | 否 |  | The package license. |
| `children` | `Array<mojo.lang.Package>` |  | 否 |  |  |
| `extraInfo` |  |  | 否 |  |
