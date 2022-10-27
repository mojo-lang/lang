// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface StructLiteralExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.StructLiteralExpr)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.lang.Position start_position = 1;</code>
   * @return Whether the startPosition field is set.
   */
  boolean hasStartPosition();
  /**
   * <code>.mojo.lang.Position start_position = 1;</code>
   * @return The startPosition.
   */
  org.mojolang.mojo.lang.Position getStartPosition();
  /**
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getStartPositionOrBuilder();

  /**
   * <code>.mojo.lang.Position end_position = 2;</code>
   * @return Whether the endPosition field is set.
   */
  boolean hasEndPosition();
  /**
   * <code>.mojo.lang.Position end_position = 2;</code>
   * @return The endPosition.
   */
  org.mojolang.mojo.lang.Position getEndPosition();
  /**
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getEndPositionOrBuilder();

  /**
   * <code>int32 kind = 4;</code>
   * @return The kind.
   */
  int getKind();

  /**
   * <code>bool implicit = 5;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>.mojo.lang.Expression expression = 10;</code>
   * @return Whether the expression field is set.
   */
  boolean hasExpression();
  /**
   * <code>.mojo.lang.Expression expression = 10;</code>
   * @return The expression.
   */
  org.mojolang.mojo.lang.Expression getExpression();
  /**
   * <code>.mojo.lang.Expression expression = 10;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getExpressionOrBuilder();

  /**
   * <code>.mojo.lang.Operator operator = 11;</code>
   * @return Whether the operator field is set.
   */
  boolean hasOperator();
  /**
   * <code>.mojo.lang.Operator operator = 11;</code>
   * @return The operator.
   */
  org.mojolang.mojo.lang.Operator getOperator();
  /**
   * <code>.mojo.lang.Operator operator = 11;</code>
   */
  org.mojolang.mojo.lang.OperatorOrBuilder getOperatorOrBuilder();

  /**
   * <code>.mojo.lang.ObjectLiteralExpr argument = 21;</code>
   * @return Whether the argument field is set.
   */
  boolean hasArgument();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr argument = 21;</code>
   * @return The argument.
   */
  org.mojolang.mojo.lang.ObjectLiteralExpr getArgument();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr argument = 21;</code>
   */
  org.mojolang.mojo.lang.ObjectLiteralExprOrBuilder getArgumentOrBuilder();

  public org.mojolang.mojo.lang.StructLiteralExpr.CalleeCase getCalleeCase();
}
