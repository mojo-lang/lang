type Attribute = Attr | GroupAttr

type Attr {
    package: Package @1 | String @2

    name: String @2
    expression: Expression @3
}

type GroupAttr {
    attributes: ObjectExpr @3
}
