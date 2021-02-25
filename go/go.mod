module github.com/mojo-lang/lang/go

replace github.com/mojo-lang/core/go => ../../core/go

replace github.com/mojo-lang/document/go => ../../document/go

go 1.16

require (
	github.com/gogo/protobuf v1.3.2
	github.com/iancoleman/strcase v0.1.3 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/mojo-lang/core/go v0.0.0-00010101000000-000000000000
	github.com/mojo-lang/document/go v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
)
