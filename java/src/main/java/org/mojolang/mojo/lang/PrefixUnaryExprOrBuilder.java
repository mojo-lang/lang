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
   * <code>string operator = 15;</code>
   * @return The operator.
   */
  java.lang.String getOperator();
  /**
   * <code>string operator = 15;</code>
   * @return The bytes for operator.
   */
  com.google.protobuf.ByteString
      getOperatorBytes();

  /**
   * <code>.mojo.lang.Expression expression = 16;</code>
   * @return Whether the expression field is set.
   */
  boolean hasExpression();
  /**
   * <code>.mojo.lang.Expression expression = 16;</code>
   * @return The expression.
   */
  org.mojolang.mojo.lang.Expression getExpression();
  /**
   * <code>.mojo.lang.Expression expression = 16;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getExpressionOrBuilder();
}
