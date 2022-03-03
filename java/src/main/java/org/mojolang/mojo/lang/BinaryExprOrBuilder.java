// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface BinaryExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.BinaryExpr)
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
   * <code>.mojo.lang.Operator operator = 20;</code>
   * @return Whether the operator field is set.
   */
  boolean hasOperator();
  /**
   * <code>.mojo.lang.Operator operator = 20;</code>
   * @return The operator.
   */
  org.mojolang.mojo.lang.Operator getOperator();
  /**
   * <code>.mojo.lang.Operator operator = 20;</code>
   */
  org.mojolang.mojo.lang.OperatorOrBuilder getOperatorOrBuilder();

  /**
   * <code>.mojo.lang.Expression left_hand_argument = 21;</code>
   * @return Whether the leftHandArgument field is set.
   */
  boolean hasLeftHandArgument();
  /**
   * <code>.mojo.lang.Expression left_hand_argument = 21;</code>
   * @return The leftHandArgument.
   */
  org.mojolang.mojo.lang.Expression getLeftHandArgument();
  /**
   * <code>.mojo.lang.Expression left_hand_argument = 21;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getLeftHandArgumentOrBuilder();

  /**
   * <code>.mojo.lang.Expression right_hand_argument = 22;</code>
   * @return Whether the rightHandArgument field is set.
   */
  boolean hasRightHandArgument();
  /**
   * <code>.mojo.lang.Expression right_hand_argument = 22;</code>
   * @return The rightHandArgument.
   */
  org.mojolang.mojo.lang.Expression getRightHandArgument();
  /**
   * <code>.mojo.lang.Expression right_hand_argument = 22;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getRightHandArgumentOrBuilder();
}
