// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface AttributeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Attribute)
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
   * <code>bool implicit = 5;</code>
   * @return The implicit.
   */
  boolean getImplicit();

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
   * <code>repeated string fields = 11;</code>
   * @return A list containing the fields.
   */
  java.util.List<java.lang.String>
      getFieldsList();
  /**
   * <code>repeated string fields = 11;</code>
   * @return The count of fields.
   */
  int getFieldsCount();
  /**
   * <code>repeated string fields = 11;</code>
   * @param index The index of the element to return.
   * @return The fields at the given index.
   */
  java.lang.String getFields(int index);
  /**
   * <code>repeated string fields = 11;</code>
   * @param index The index of the value to return.
   * @return The bytes of the fields at the given index.
   */
  com.google.protobuf.ByteString
      getFieldsBytes(int index);

  /**
   * <code>.mojo.lang.AttributeDecl declaration = 13 [(.mojo.reference) = ""];</code>
   * @return Whether the declaration field is set.
   */
  boolean hasDeclaration();
  /**
   * <code>.mojo.lang.AttributeDecl declaration = 13 [(.mojo.reference) = ""];</code>
   * @return The declaration.
   */
  org.mojolang.mojo.lang.AttributeDecl getDeclaration();
  /**
   * <code>.mojo.lang.AttributeDecl declaration = 13 [(.mojo.reference) = ""];</code>
   */
  org.mojolang.mojo.lang.AttributeDeclOrBuilder getDeclarationOrBuilder();

  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 15;</code>
   */
  java.util.List<org.mojolang.mojo.lang.NominalType> 
      getGenericArgumentsList();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 15;</code>
   */
  org.mojolang.mojo.lang.NominalType getGenericArguments(int index);
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 15;</code>
   */
  int getGenericArgumentsCount();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 15;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.NominalTypeOrBuilder> 
      getGenericArgumentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 15;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getGenericArgumentsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Argument arguments = 16;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Argument> 
      getArgumentsList();
  /**
   * <code>repeated .mojo.lang.Argument arguments = 16;</code>
   */
  org.mojolang.mojo.lang.Argument getArguments(int index);
  /**
   * <code>repeated .mojo.lang.Argument arguments = 16;</code>
   */
  int getArgumentsCount();
  /**
   * <code>repeated .mojo.lang.Argument arguments = 16;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ArgumentOrBuilder> 
      getArgumentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Argument arguments = 16;</code>
   */
  org.mojolang.mojo.lang.ArgumentOrBuilder getArgumentsOrBuilder(
      int index);

  /**
   * <code>.mojo.lang.Position name_position = 19 [(.mojo.db_ignore) = true];</code>
   * @return Whether the namePosition field is set.
   */
  boolean hasNamePosition();
  /**
   * <code>.mojo.lang.Position name_position = 19 [(.mojo.db_ignore) = true];</code>
   * @return The namePosition.
   */
  org.mojolang.mojo.lang.Position getNamePosition();
  /**
   * <code>.mojo.lang.Position name_position = 19 [(.mojo.db_ignore) = true];</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getNamePositionOrBuilder();

  /**
   * <code>.mojo.lang.Expression value = 20;</code>
   * @return Whether the value field is set.
   */
  boolean hasValue();
  /**
   * <code>.mojo.lang.Expression value = 20;</code>
   * @return The value.
   */
  org.mojolang.mojo.lang.Expression getValue();
  /**
   * <code>.mojo.lang.Expression value = 20;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getValueOrBuilder();
}
