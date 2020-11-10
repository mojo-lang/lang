
type Int

type A : Array<Int>


type B<T> {
    type C<U> {
        foo : [String]
    }

    a: Int32
    b: String
    c: C<T>
    d: T
}

enum Foobar : Int {
    Foo :1
    Bar :2
}

Foobar.Foo
Foobar.Bar

type Test {
    a: Foobar
}

TypeAlias {
}

DataTypeDecl {
    name: 'Int'
}

DataTypeDecl {
    name: 'Array'
    generic_params: [{
        name: 'T'
    }]
}

DataTypeDecl {
    name: 'Dictionary'
    generic_params: [{
        name: 'Key'
    },{
        name: 'Value'
    }]
}

TypeDecl {
    name: 'A'
    type: TypeName {
        name: 'Array'
        generic_arguments: [{
            TypeName{
                name: 'Int'
            }
        }]
    }
}

StructDecl {
    name: 'B'
    generic_params: [{
        name: 'T'
    }]
    type: StructType {
        fields: [{
            name: 'a',
            type: ''
        }, {
        }]
    }
    struct_decls: [{
        name: 'C'
        generic_params:[]
        type: Struct {
            fields:[{
                
            }]
        }
    }],
    enum_decls: []
    type_alias_decls: []
}