// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface PrefixUnaryExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.PrefixUnaryExpr)
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
   * <code>int32 kind = 3;</code>
   * @return The kind.
   */
  int getKind();

  /**
   * <code>bool implicit = 4;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>.mojo.lang.Expression callee = 10;</code>
   * @return Whether the callee field is set.
   */
  boolean hasCallee();
  /**
   * <code>.mojo.lang.Expression callee = 10;</code>
   * @return The callee.
   */
  org.mojolang.mojo.lang.Expression getCallee();
  /**
   * <code>.mojo.lang.Expression callee = 10;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getCalleeOrBuilder();

  /**
   * <code>.mojo.lang.Operator operator = 15;</code>
   * @return Whether the operator field is set.
   */
  boolean hasOperator();
  /**
   * <code>.mojo.lang.Operator operator = 15;</code>
   * @return The operator.
   */
  org.mojolang.mojo.lang.Operator getOperator();
  /**
   * <code>.mojo.lang.Operator operator = 15;</code>
   */
  org.mojolang.mojo.lang.OperatorOrBuilder getOperatorOrBuilder();

  /**
   * <code>.mojo.lang.Expression argument = 16;</code>
   * @return Whether the argument field is set.
   */
  boolean hasArgument();
  /**
   * <code>.mojo.lang.Expression argument = 16;</code>
   * @return The argument.
   */
  org.mojolang.mojo.lang.Expression getArgument();
  /**
   * <code>.mojo.lang.Expression argument = 16;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getArgumentOrBuilder();
}
