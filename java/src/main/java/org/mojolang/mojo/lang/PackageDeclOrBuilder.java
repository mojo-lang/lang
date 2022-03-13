// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface PackageDeclOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.PackageDecl)
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
   * <code>bool implicit = 5;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>.mojo.lang.Document document = 6;</code>
   * @return Whether the document field is set.
   */
  boolean hasDocument();
  /**
   * <code>.mojo.lang.Document document = 6;</code>
   * @return The document.
   */
  org.mojolang.mojo.lang.Document getDocument();
  /**
   * <code>.mojo.lang.Document document = 6;</code>
   */
  org.mojolang.mojo.lang.DocumentOrBuilder getDocumentOrBuilder();

  /**
   * <code>string package_name = 7;</code>
   * @return The packageName.
   */
  java.lang.String getPackageName();
  /**
   * <code>string package_name = 7;</code>
   * @return The bytes for packageName.
   */
  com.google.protobuf.ByteString
      getPackageNameBytes();

  /**
   * <code>string source_file_name = 8;</code>
   * @return The sourceFileName.
   */
  java.lang.String getSourceFileName();
  /**
   * <code>string source_file_name = 8;</code>
   * @return The bytes for sourceFileName.
   */
  com.google.protobuf.ByteString
      getSourceFileNameBytes();

  /**
   * <code>.mojo.lang.Position keyword_position = 9;</code>
   * @return Whether the keywordPosition field is set.
   */
  boolean hasKeywordPosition();
  /**
   * <code>.mojo.lang.Position keyword_position = 9;</code>
   * @return The keywordPosition.
   */
  org.mojolang.mojo.lang.Position getKeywordPosition();
  /**
   * <code>.mojo.lang.Position keyword_position = 9;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getKeywordPositionOrBuilder();

  /**
   * <code>string name = 10;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 10;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>.mojo.lang.ObjectLiteralExpr package_literal_expr = 15;</code>
   * @return Whether the packageLiteralExpr field is set.
   */
  boolean hasPackageLiteralExpr();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr package_literal_expr = 15;</code>
   * @return The packageLiteralExpr.
   */
  org.mojolang.mojo.lang.ObjectLiteralExpr getPackageLiteralExpr();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr package_literal_expr = 15;</code>
   */
  org.mojolang.mojo.lang.ObjectLiteralExprOrBuilder getPackageLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.Package package = 16;</code>
   * @return Whether the package field is set.
   */
  boolean hasPackage();
  /**
   * <code>.mojo.lang.Package package = 16;</code>
   * @return The package.
   */
  org.mojolang.mojo.lang.Package getPackage();
  /**
   * <code>.mojo.lang.Package package = 16;</code>
   */
  org.mojolang.mojo.lang.PackageOrBuilder getPackageOrBuilder();

  /**
   * <code>.mojo.lang.Position name_position = 19;</code>
   * @return Whether the namePosition field is set.
   */
  boolean hasNamePosition();
  /**
   * <code>.mojo.lang.Position name_position = 19;</code>
   * @return The namePosition.
   */
  org.mojolang.mojo.lang.Position getNamePosition();
  /**
   * <code>.mojo.lang.Position name_position = 19;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getNamePositionOrBuilder();
}
