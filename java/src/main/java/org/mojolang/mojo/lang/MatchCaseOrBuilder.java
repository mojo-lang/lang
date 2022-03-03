// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface MatchCaseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.MatchCase)
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
   * <code>int64 kind = 4;</code>
   * @return The kind.
   */
  long getKind();

  /**
   * <code>bool implicit = 5;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>.mojo.lang.Pattern pattern = 20;</code>
   * @return Whether the pattern field is set.
   */
  boolean hasPattern();
  /**
   * <code>.mojo.lang.Pattern pattern = 20;</code>
   * @return The pattern.
   */
  org.mojolang.mojo.lang.Pattern getPattern();
  /**
   * <code>.mojo.lang.Pattern pattern = 20;</code>
   */
  org.mojolang.mojo.lang.PatternOrBuilder getPatternOrBuilder();

  /**
   * <code>.mojo.lang.BlockStmt block_stmt = 21;</code>
   * @return Whether the blockStmt field is set.
   */
  boolean hasBlockStmt();
  /**
   * <code>.mojo.lang.BlockStmt block_stmt = 21;</code>
   * @return The blockStmt.
   */
  org.mojolang.mojo.lang.BlockStmt getBlockStmt();
  /**
   * <code>.mojo.lang.BlockStmt block_stmt = 21;</code>
   */
  org.mojolang.mojo.lang.BlockStmtOrBuilder getBlockStmtOrBuilder();

  /**
   * <code>.mojo.lang.Expression expression = 22;</code>
   * @return Whether the expression field is set.
   */
  boolean hasExpression();
  /**
   * <code>.mojo.lang.Expression expression = 22;</code>
   * @return The expression.
   */
  org.mojolang.mojo.lang.Expression getExpression();
  /**
   * <code>.mojo.lang.Expression expression = 22;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getExpressionOrBuilder();

  public org.mojolang.mojo.lang.MatchCase.BodyCase getBodyCase();
}
