// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface ExplicitMemberExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.ExplicitMemberExpr)
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
   * <code>string member = 20;</code>
   * @return The member.
   */
  java.lang.String getMember();
  /**
   * <code>string member = 20;</code>
   * @return The bytes for member.
   */
  com.google.protobuf.ByteString
      getMemberBytes();
}