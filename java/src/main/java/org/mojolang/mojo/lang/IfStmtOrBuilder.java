// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface IfStmtOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.IfStmt)
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
   * <code>int64 kind = 3;</code>
   * @return The kind.
   */
  long getKind();

  /**
   * <code>bool implicit = 4;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>.mojo.lang.Expression condition = 20;</code>
   * @return Whether the condition field is set.
   */
  boolean hasCondition();
  /**
   * <code>.mojo.lang.Expression condition = 20;</code>
   * @return The condition.
   */
  org.mojolang.mojo.lang.Expression getCondition();
  /**
   * <code>.mojo.lang.Expression condition = 20;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getConditionOrBuilder();

  /**
   * <code>.mojo.lang.BlockStmt then_branch = 21;</code>
   * @return Whether the thenBranch field is set.
   */
  boolean hasThenBranch();
  /**
   * <code>.mojo.lang.BlockStmt then_branch = 21;</code>
   * @return The thenBranch.
   */
  org.mojolang.mojo.lang.BlockStmt getThenBranch();
  /**
   * <code>.mojo.lang.BlockStmt then_branch = 21;</code>
   */
  org.mojolang.mojo.lang.BlockStmtOrBuilder getThenBranchOrBuilder();

  /**
   * <code>.mojo.lang.BlockStmt block_stmt = 22;</code>
   * @return Whether the blockStmt field is set.
   */
  boolean hasBlockStmt();
  /**
   * <code>.mojo.lang.BlockStmt block_stmt = 22;</code>
   * @return The blockStmt.
   */
  org.mojolang.mojo.lang.BlockStmt getBlockStmt();
  /**
   * <code>.mojo.lang.BlockStmt block_stmt = 22;</code>
   */
  org.mojolang.mojo.lang.BlockStmtOrBuilder getBlockStmtOrBuilder();

  /**
   * <code>.mojo.lang.IfStmt if_stmt = 23;</code>
   * @return Whether the ifStmt field is set.
   */
  boolean hasIfStmt();
  /**
   * <code>.mojo.lang.IfStmt if_stmt = 23;</code>
   * @return The ifStmt.
   */
  org.mojolang.mojo.lang.IfStmt getIfStmt();
  /**
   * <code>.mojo.lang.IfStmt if_stmt = 23;</code>
   */
  org.mojolang.mojo.lang.IfStmtOrBuilder getIfStmtOrBuilder();

  public org.mojolang.mojo.lang.IfStmt.ElseBranchCase getElseBranchCase();
}