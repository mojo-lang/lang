type FunctionCallExpr : Expr {
    // 函数名称调用
    // lambda 匿名调用
    callee: Expression @10

    ///
    generic_arguments: [GenericArgument] @11

    ///
    arguments: [Argument] @12
}
