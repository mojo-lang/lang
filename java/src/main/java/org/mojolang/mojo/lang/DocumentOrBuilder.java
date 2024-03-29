// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface DocumentOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Document)
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
   * <code>bool following = 3;</code>
   * @return The following.
   */
  boolean getFollowing();

  /**
   * <code>bool private = 4;</code>
   * @return The private.
   */
  boolean getPrivate();

  /**
   * <code>repeated .mojo.lang.Document.Line lines = 10;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Document.Line> 
      getLinesList();
  /**
   * <code>repeated .mojo.lang.Document.Line lines = 10;</code>
   */
  org.mojolang.mojo.lang.Document.Line getLines(int index);
  /**
   * <code>repeated .mojo.lang.Document.Line lines = 10;</code>
   */
  int getLinesCount();
  /**
   * <code>repeated .mojo.lang.Document.Line lines = 10;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.Document.LineOrBuilder> 
      getLinesOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Document.Line lines = 10;</code>
   */
  org.mojolang.mojo.lang.Document.LineOrBuilder getLinesOrBuilder(
      int index);

  /**
   * <code>.mojo.document.Document structured = 15;</code>
   * @return Whether the structured field is set.
   */
  boolean hasStructured();
  /**
   * <code>.mojo.document.Document structured = 15;</code>
   * @return The structured.
   */
  org.mojolang.mojo.document.Document getStructured();
  /**
   * <code>.mojo.document.Document structured = 15;</code>
   */
  org.mojolang.mojo.document.DocumentOrBuilder getStructuredOrBuilder();
}
