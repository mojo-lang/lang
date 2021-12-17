| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | package name |
| `fullName` | `string` |  | 否 |  |  |
| `version` | `mojo.core.Version` |  | 否 |  |  |
| `children` | `Array<mojo.lang.Package>` |  | 否 |  |  |
| `authors` | `Array<mojo.lang.Package.Author>` |  | 否 |  |
| `summary` | `string` |  | 否 |  |
| `description` | `string` |  | 否 |  | A description of the package. |
| `document` | `mojo.lang.CachedDocument` |  | 否 |  |  |
| `repository` | `string` | `url` | 否 |  | URL of the package source repository. |
| `license` | `string` |  | 否 |  | The package license. |
| `exports` | `Array<string>` |  | 否 |  |  |
| `sourceFiles` | `Array<mojo.lang.SourceFile>` |  | 否 |  |  |
| `scope` | `mojo.lang.Scope` |  | 否 |  | package scope across all files |
| `extraInfo` | `Map<string, mojo.core.Value>` |  | 否 |  |
| `entityRelationship` | `mojo.lang.EntityRelationship` |  | 否 |  | entity relationship analysis in the package |
| `dependencies` | `Map<string, mojo.lang.Package.Requirement>` |  | 否 |  |  |
| `resolvedDependencies` | `Map<string, mojo.lang.Package>` |  | 否 |  |  |
