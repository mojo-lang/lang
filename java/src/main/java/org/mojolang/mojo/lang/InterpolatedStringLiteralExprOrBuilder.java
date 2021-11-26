// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/interpolated_string_literal_expr.proto

package org.mojolang.mojo.lang;

public interface InterpolatedStringLiteralExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.InterpolatedStringLiteralExpr)
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
   * <code>repeated .mojo.lang.Expr segments = 20;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Expr> 
      getSegmentsList();
  /**
   * <code>repeated .mojo.lang.Expr segments = 20;</code>
   */
  org.mojolang.mojo.lang.Expr getSegments(int index);
  /**
   * <code>repeated .mojo.lang.Expr segments = 20;</code>
   */
  int getSegmentsCount();
  /**
   * <code>repeated .mojo.lang.Expr segments = 20;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ExprOrBuilder> 
      getSegmentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Expr segments = 20;</code>
   */
  org.mojolang.mojo.lang.ExprOrBuilder getSegmentsOrBuilder(
      int index);

  /**
   * <code>.mojo.lang.Expr semantic_expr = 21;</code>
   * @return Whether the semanticExpr field is set.
   */
  boolean hasSemanticExpr();
  /**
   * <code>.mojo.lang.Expr semantic_expr = 21;</code>
   * @return The semanticExpr.
   */
  org.mojolang.mojo.lang.Expr getSemanticExpr();
  /**
   * <code>.mojo.lang.Expr semantic_expr = 21;</code>
   */
  org.mojolang.mojo.lang.ExprOrBuilder getSemanticExprOrBuilder();
}