// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface ClosureExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.ClosureExpr)
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
   * <code>string full_name = 11;</code>
   * @return The fullName.
   */
  java.lang.String getFullName();
  /**
   * <code>string full_name = 11;</code>
   * @return The bytes for fullName.
   */
  com.google.protobuf.ByteString
      getFullNameBytes();

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
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 13;</code>
   */
  java.util.List<org.mojolang.mojo.lang.GenericParameter> 
      getGenericParametersList();
  /**
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 13;</code>
   */
  org.mojolang.mojo.lang.GenericParameter getGenericParameters(int index);
  /**
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 13;</code>
   */
  int getGenericParametersCount();
  /**
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 13;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.GenericParameterOrBuilder> 
      getGenericParametersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 13;</code>
   */
  org.mojolang.mojo.lang.GenericParameterOrBuilder getGenericParametersOrBuilder(
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
   * <code>.mojo.lang.FunctionSignature signature = 20;</code>
   * @return Whether the signature field is set.
   */
  boolean hasSignature();
  /**
   * <code>.mojo.lang.FunctionSignature signature = 20;</code>
   * @return The signature.
   */
  org.mojolang.mojo.lang.FunctionSignature getSignature();
  /**
   * <code>.mojo.lang.FunctionSignature signature = 20;</code>
   */
  org.mojolang.mojo.lang.FunctionSignatureOrBuilder getSignatureOrBuilder();

  /**
   * <code>.mojo.lang.BlockStmt body = 21;</code>
   * @return Whether the body field is set.
   */
  boolean hasBody();
  /**
   * <code>.mojo.lang.BlockStmt body = 21;</code>
   * @return The body.
   */
  org.mojolang.mojo.lang.BlockStmt getBody();
  /**
   * <code>.mojo.lang.BlockStmt body = 21;</code>
   */
  org.mojolang.mojo.lang.BlockStmtOrBuilder getBodyOrBuilder();

  /**
   * <code>.mojo.lang.Scope scope = 30;</code>
   * @return Whether the scope field is set.
   */
  boolean hasScope();
  /**
   * <code>.mojo.lang.Scope scope = 30;</code>
   * @return The scope.
   */
  org.mojolang.mojo.lang.Scope getScope();
  /**
   * <code>.mojo.lang.Scope scope = 30;</code>
   */
  org.mojolang.mojo.lang.ScopeOrBuilder getScopeOrBuilder();

  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 31 [(.mojo.db_ignore) = true];</code>
   * @return Whether the interfaceDecl field is set.
   */
  boolean hasInterfaceDecl();
  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 31 [(.mojo.db_ignore) = true];</code>
   * @return The interfaceDecl.
   */
  org.mojolang.mojo.lang.InterfaceDecl getInterfaceDecl();
  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 31 [(.mojo.db_ignore) = true];</code>
   */
  org.mojolang.mojo.lang.InterfaceDeclOrBuilder getInterfaceDeclOrBuilder();

  /**
   * <code>repeated string captures = 40;</code>
   * @return A list containing the captures.
   */
  java.util.List<java.lang.String>
      getCapturesList();
  /**
   * <code>repeated string captures = 40;</code>
   * @return The count of captures.
   */
  int getCapturesCount();
  /**
   * <code>repeated string captures = 40;</code>
   * @param index The index of the element to return.
   * @return The captures at the given index.
   */
  java.lang.String getCaptures(int index);
  /**
   * <code>repeated string captures = 40;</code>
   * @param index The index of the value to return.
   * @return The bytes of the captures at the given index.
   */
  com.google.protobuf.ByteString
      getCapturesBytes(int index);
}
