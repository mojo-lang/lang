// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface ImportDeclOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.ImportDecl)
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
   * <code>.mojo.lang.Document document = 4;</code>
   * @return Whether the document field is set.
   */
  boolean hasDocument();
  /**
   * <code>.mojo.lang.Document document = 4;</code>
   * @return The document.
   */
  org.mojolang.mojo.lang.Document getDocument();
  /**
   * <code>.mojo.lang.Document document = 4;</code>
   */
  org.mojolang.mojo.lang.DocumentOrBuilder getDocumentOrBuilder();

  /**
   * <code>string package_name = 5;</code>
   * @return The packageName.
   */
  java.lang.String getPackageName();
  /**
   * <code>string package_name = 5;</code>
   * @return The bytes for packageName.
   */
  com.google.protobuf.ByteString
      getPackageNameBytes();

  /**
   * <code>string source_file_name = 6;</code>
   * @return The sourceFileName.
   */
  java.lang.String getSourceFileName();
  /**
   * <code>string source_file_name = 6;</code>
   * @return The bytes for sourceFileName.
   */
  com.google.protobuf.ByteString
      getSourceFileNameBytes();

  /**
   * <code>bool implicit = 7;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>string import_package_name = 10;</code>
   * @return The importPackageName.
   */
  java.lang.String getImportPackageName();
  /**
   * <code>string import_package_name = 10;</code>
   * @return The bytes for importPackageName.
   */
  com.google.protobuf.ByteString
      getImportPackageNameBytes();

  /**
   * <code>string import_package_alias = 11;</code>
   * @return The importPackageAlias.
   */
  java.lang.String getImportPackageAlias();
  /**
   * <code>string import_package_alias = 11;</code>
   * @return The bytes for importPackageAlias.
   */
  com.google.protobuf.ByteString
      getImportPackageAliasBytes();

  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Identifier> 
      getIdentifiersList();
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  org.mojolang.mojo.lang.Identifier getIdentifiers(int index);
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  int getIdentifiersCount();
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getIdentifiersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  org.mojolang.mojo.lang.IdentifierOrBuilder getIdentifiersOrBuilder(
      int index);
}
