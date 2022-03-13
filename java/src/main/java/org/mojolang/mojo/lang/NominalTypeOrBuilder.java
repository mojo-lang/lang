// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface NominalTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.NominalType)
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
   * <code>bool implicit = 5;</code>
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
   * <code>.mojo.lang.TypeDeclaration type_declaration = 11;</code>
   * @return Whether the typeDeclaration field is set.
   */
  boolean hasTypeDeclaration();
  /**
   * <code>.mojo.lang.TypeDeclaration type_declaration = 11;</code>
   * @return The typeDeclaration.
   */
  org.mojolang.mojo.lang.TypeDeclaration getTypeDeclaration();
  /**
   * <code>.mojo.lang.TypeDeclaration type_declaration = 11;</code>
   */
  org.mojolang.mojo.lang.TypeDeclarationOrBuilder getTypeDeclarationOrBuilder();

  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 12;</code>
   */
  java.util.List<org.mojolang.mojo.lang.NominalType> 
      getGenericArgumentsList();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 12;</code>
   */
  org.mojolang.mojo.lang.NominalType getGenericArguments(int index);
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 12;</code>
   */
  int getGenericArgumentsCount();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 12;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.NominalTypeOrBuilder> 
      getGenericArgumentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 12;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getGenericArgumentsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Attribute attributes = 13;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Attribute> 
      getAttributesList();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 13;</code>
   */
  org.mojolang.mojo.lang.Attribute getAttributes(int index);
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 13;</code>
   */
  int getAttributesCount();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 13;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.AttributeOrBuilder> 
      getAttributesOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 13;</code>
   */
  org.mojolang.mojo.lang.AttributeOrBuilder getAttributesOrBuilder(
      int index);

  /**
   * <code>.mojo.lang.NominalType enclosing_type = 14;</code>
   * @return Whether the enclosingType field is set.
   */
  boolean hasEnclosingType();
  /**
   * <code>.mojo.lang.NominalType enclosing_type = 14;</code>
   * @return The enclosingType.
   */
  org.mojolang.mojo.lang.NominalType getEnclosingType();
  /**
   * <code>.mojo.lang.NominalType enclosing_type = 14;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getEnclosingTypeOrBuilder();
}
