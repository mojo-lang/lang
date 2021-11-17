// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface ConditionalExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.ConditionalExpr)
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
   * <code>.mojo.lang.Expression condition = 10;</code>
   * @return Whether the condition field is set.
   */
  boolean hasCondition();
  /**
   * <code>.mojo.lang.Expression condition = 10;</code>
   * @return The condition.
   */
  org.mojolang.mojo.lang.Expression getCondition();
  /**
   * <code>.mojo.lang.Expression condition = 10;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getConditionOrBuilder();

  /**
   * <code>.mojo.lang.Expression then_branch = 11;</code>
   * @return Whether the thenBranch field is set.
   */
  boolean hasThenBranch();
  /**
   * <code>.mojo.lang.Expression then_branch = 11;</code>
   * @return The thenBranch.
   */
  org.mojolang.mojo.lang.Expression getThenBranch();
  /**
   * <code>.mojo.lang.Expression then_branch = 11;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getThenBranchOrBuilder();

  /**
   * <code>.mojo.lang.Expression else_branch = 12;</code>
   * @return Whether the elseBranch field is set.
   */
  boolean hasElseBranch();
  /**
   * <code>.mojo.lang.Expression else_branch = 12;</code>
   * @return The elseBranch.
   */
  org.mojolang.mojo.lang.Expression getElseBranch();
  /**
   * <code>.mojo.lang.Expression else_branch = 12;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getElseBranchOrBuilder();
}
