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
   * <code>string filter = 11;</code>
   * @return The filter.
   */
  java.lang.String getFilter();
  /**
   * <code>string filter = 11;</code>
   * @return The bytes for filter.
   */
  com.google.protobuf.ByteString
      getFilterBytes();

  /**
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Attribute> 
      getAttributesList();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  org.mojolang.mojo.lang.Attribute getAttributes(int index);
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  int getAttributesCount();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.AttributeOrBuilder> 
      getAttributesOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  org.mojolang.mojo.lang.AttributeOrBuilder getAttributesOrBuilder(
      int index);

  /**
   * <code>string import_package_name = 13;</code>
   * @return The importPackageName.
   */
  java.lang.String getImportPackageName();
  /**
   * <code>string import_package_name = 13;</code>
   * @return The bytes for importPackageName.
   */
  com.google.protobuf.ByteString
      getImportPackageNameBytes();

  /**
   * <code>string import_package_alias = 14;</code>
   * @return The importPackageAlias.
   */
  java.lang.String getImportPackageAlias();
  /**
   * <code>string import_package_alias = 14;</code>
   * @return The bytes for importPackageAlias.
   */
  com.google.protobuf.ByteString
      getImportPackageAliasBytes();

  /**
   * <code>string import_file_name = 15;</code>
   * @return The importFileName.
   */
  java.lang.String getImportFileName();
  /**
   * <code>string import_file_name = 15;</code>
   * @return The bytes for importFileName.
   */
  com.google.protobuf.ByteString
      getImportFileNameBytes();

  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 16;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Identifier> 
      getIdentifiersList();
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 16;</code>
   */
  org.mojolang.mojo.lang.Identifier getIdentifiers(int index);
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 16;</code>
   */
  int getIdentifiersCount();
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 16;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getIdentifiersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 16;</code>
   */
  org.mojolang.mojo.lang.IdentifierOrBuilder getIdentifiersOrBuilder(
      int index);
}
