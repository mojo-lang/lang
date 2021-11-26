// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface GroupDeclOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.GroupDecl)
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
   * <code>.mojo.lang.Identifier.Kind type = 20;</code>
   * @return The enum numeric value on the wire for type.
   */
  int getTypeValue();
  /**
   * <code>.mojo.lang.Identifier.Kind type = 20;</code>
   * @return The type.
   */
  org.mojolang.mojo.lang.Identifier.Kind getType();

  /**
   * <code>repeated .mojo.lang.Declaration declarations = 21;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Declaration> 
      getDeclarationsList();
  /**
   * <code>repeated .mojo.lang.Declaration declarations = 21;</code>
   */
  org.mojolang.mojo.lang.Declaration getDeclarations(int index);
  /**
   * <code>repeated .mojo.lang.Declaration declarations = 21;</code>
   */
  int getDeclarationsCount();
  /**
   * <code>repeated .mojo.lang.Declaration declarations = 21;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.DeclarationOrBuilder> 
      getDeclarationsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Declaration declarations = 21;</code>
   */
  org.mojolang.mojo.lang.DeclarationOrBuilder getDeclarationsOrBuilder(
      int index);
}