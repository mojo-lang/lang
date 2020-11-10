/// ForStmt - foreach statement that iterates over the elements in a
/// container.
///
/// Example:
/// ```
/// for i in 0..10 {
///   print(i.to<String>)
/// }
/// ```
type ForStmt : LoopStmt {
    iterator: Expression
    sequence: Expression
}
