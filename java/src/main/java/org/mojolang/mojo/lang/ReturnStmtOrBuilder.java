// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface ReturnStmtOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.ReturnStmt)
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
   * <code>.mojo.lang.Expression result = 10;</code>
   * @return Whether the result field is set.
   */
  boolean hasResult();
  /**
   * <code>.mojo.lang.Expression result = 10;</code>
   * @return The result.
   */
  org.mojolang.mojo.lang.Expression getResult();
  /**
   * <code>.mojo.lang.Expression result = 10;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getResultOrBuilder();
}
