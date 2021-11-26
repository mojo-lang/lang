// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface MapLiteralExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.MapLiteralExpr)
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
   * <code>repeated .mojo.lang.MapLiteralExpr.Entry entries = 15;</code>
   */
  java.util.List<org.mojolang.mojo.lang.MapLiteralExpr.Entry> 
      getEntriesList();
  /**
   * <code>repeated .mojo.lang.MapLiteralExpr.Entry entries = 15;</code>
   */
  org.mojolang.mojo.lang.MapLiteralExpr.Entry getEntries(int index);
  /**
   * <code>repeated .mojo.lang.MapLiteralExpr.Entry entries = 15;</code>
   */
  int getEntriesCount();
  /**
   * <code>repeated .mojo.lang.MapLiteralExpr.Entry entries = 15;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.MapLiteralExpr.EntryOrBuilder> 
      getEntriesOrBuilderList();
  /**
   * <code>repeated .mojo.lang.MapLiteralExpr.Entry entries = 15;</code>
   */
  org.mojolang.mojo.lang.MapLiteralExpr.EntryOrBuilder getEntriesOrBuilder(
      int index);
}