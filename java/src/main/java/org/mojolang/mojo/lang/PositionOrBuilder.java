// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface PositionOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Position)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string filename = 1;</code>
   * @return The filename.
   */
  java.lang.String getFilename();
  /**
   * <code>string filename = 1;</code>
   * @return The bytes for filename.
   */
  com.google.protobuf.ByteString
      getFilenameBytes();

  /**
   * <code>int64 offset = 2;</code>
   * @return The offset.
   */
  long getOffset();

  /**
   * <code>int64 line = 3;</code>
   * @return The line.
   */
  long getLine();

  /**
   * <code>int64 column = 4;</code>
   * @return The column.
   */
  long getColumn();

  /**
   * <code>repeated .mojo.lang.Comment leading_comments = 10;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Comment> 
      getLeadingCommentsList();
  /**
   * <code>repeated .mojo.lang.Comment leading_comments = 10;</code>
   */
  org.mojolang.mojo.lang.Comment getLeadingComments(int index);
  /**
   * <code>repeated .mojo.lang.Comment leading_comments = 10;</code>
   */
  int getLeadingCommentsCount();
  /**
   * <code>repeated .mojo.lang.Comment leading_comments = 10;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.CommentOrBuilder> 
      getLeadingCommentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Comment leading_comments = 10;</code>
   */
  org.mojolang.mojo.lang.CommentOrBuilder getLeadingCommentsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Comment tailing_comments = 13;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Comment> 
      getTailingCommentsList();
  /**
   * <code>repeated .mojo.lang.Comment tailing_comments = 13;</code>
   */
  org.mojolang.mojo.lang.Comment getTailingComments(int index);
  /**
   * <code>repeated .mojo.lang.Comment tailing_comments = 13;</code>
   */
  int getTailingCommentsCount();
  /**
   * <code>repeated .mojo.lang.Comment tailing_comments = 13;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.CommentOrBuilder> 
      getTailingCommentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Comment tailing_comments = 13;</code>
   */
  org.mojolang.mojo.lang.CommentOrBuilder getTailingCommentsOrBuilder(
      int index);
}
