// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface InterfaceDeclOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.InterfaceDecl)
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
   * <code>.mojo.lang.GroupDecl group = 15;</code>
   * @return Whether the group field is set.
   */
  boolean hasGroup();
  /**
   * <code>.mojo.lang.GroupDecl group = 15;</code>
   * @return The group.
   */
  org.mojolang.mojo.lang.GroupDecl getGroup();
  /**
   * <code>.mojo.lang.GroupDecl group = 15;</code>
   */
  org.mojolang.mojo.lang.GroupDeclOrBuilder getGroupOrBuilder();

  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 16;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Identifier> 
      getResolvedIdentifiersList();
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 16;</code>
   */
  org.mojolang.mojo.lang.Identifier getResolvedIdentifiers(int index);
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 16;</code>
   */
  int getResolvedIdentifiersCount();
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 16;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getResolvedIdentifiersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 16;</code>
   */
  org.mojolang.mojo.lang.IdentifierOrBuilder getResolvedIdentifiersOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 17;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Identifier> 
      getUnresolvedIdentifiersList();
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 17;</code>
   */
  org.mojolang.mojo.lang.Identifier getUnresolvedIdentifiers(int index);
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 17;</code>
   */
  int getUnresolvedIdentifiersCount();
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 17;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getUnresolvedIdentifiersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 17;</code>
   */
  org.mojolang.mojo.lang.IdentifierOrBuilder getUnresolvedIdentifiersOrBuilder(
      int index);

  /**
   * <code>.mojo.lang.InterfaceType type = 20;</code>
   * @return Whether the type field is set.
   */
  boolean hasType();
  /**
   * <code>.mojo.lang.InterfaceType type = 20;</code>
   * @return The type.
   */
  org.mojolang.mojo.lang.InterfaceType getType();
  /**
   * <code>.mojo.lang.InterfaceType type = 20;</code>
   */
  org.mojolang.mojo.lang.InterfaceTypeOrBuilder getTypeOrBuilder();

  /**
   * <code>repeated .mojo.lang.TypeAliasDecl type_alias_decls = 21;</code>
   */
  java.util.List<org.mojolang.mojo.lang.TypeAliasDecl> 
      getTypeAliasDeclsList();
  /**
   * <code>repeated .mojo.lang.TypeAliasDecl type_alias_decls = 21;</code>
   */
  org.mojolang.mojo.lang.TypeAliasDecl getTypeAliasDecls(int index);
  /**
   * <code>repeated .mojo.lang.TypeAliasDecl type_alias_decls = 21;</code>
   */
  int getTypeAliasDeclsCount();
  /**
   * <code>repeated .mojo.lang.TypeAliasDecl type_alias_decls = 21;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.TypeAliasDeclOrBuilder> 
      getTypeAliasDeclsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.TypeAliasDecl type_alias_decls = 21;</code>
   */
  org.mojolang.mojo.lang.TypeAliasDeclOrBuilder getTypeAliasDeclsOrBuilder(
      int index);

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
}
