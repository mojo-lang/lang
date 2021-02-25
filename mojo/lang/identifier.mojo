
// An Object describes a named language entity such as a package,
// constant, type, variable, function (incl. methods), or label.
//
// The Data fields contains object-specific data:
//
//	Kind    Data type         Data value
//	Pkg     *Scope            package scope
//	Con     int               iota for the respective declaration
//
type Identifier {
    package: String @1
    
    source_file: String @2

    position: Position @3 // identifier position

    enclosing_type_names: [String] @4

    type: String @5 // type-alias, enum, struct, interface, attribute, function, constant, variable

    name: String @6

    full_name: String @7

    declaration: Declaration @10 @reference @serialization(false)
}
