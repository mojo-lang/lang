// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface DeclarationOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Declaration)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.lang.PackageDecl package_decl = 1;</code>
   * @return Whether the packageDecl field is set.
   */
  boolean hasPackageDecl();
  /**
   * <code>.mojo.lang.PackageDecl package_decl = 1;</code>
   * @return The packageDecl.
   */
  org.mojolang.mojo.lang.PackageDecl getPackageDecl();
  /**
   * <code>.mojo.lang.PackageDecl package_decl = 1;</code>
   */
  org.mojolang.mojo.lang.PackageDeclOrBuilder getPackageDeclOrBuilder();

  /**
   * <code>.mojo.lang.ImportDecl import_decl = 2;</code>
   * @return Whether the importDecl field is set.
   */
  boolean hasImportDecl();
  /**
   * <code>.mojo.lang.ImportDecl import_decl = 2;</code>
   * @return The importDecl.
   */
  org.mojolang.mojo.lang.ImportDecl getImportDecl();
  /**
   * <code>.mojo.lang.ImportDecl import_decl = 2;</code>
   */
  org.mojolang.mojo.lang.ImportDeclOrBuilder getImportDeclOrBuilder();

  /**
   * <code>.mojo.lang.EnumDecl enum_decl = 3;</code>
   * @return Whether the enumDecl field is set.
   */
  boolean hasEnumDecl();
  /**
   * <code>.mojo.lang.EnumDecl enum_decl = 3;</code>
   * @return The enumDecl.
   */
  org.mojolang.mojo.lang.EnumDecl getEnumDecl();
  /**
   * <code>.mojo.lang.EnumDecl enum_decl = 3;</code>
   */
  org.mojolang.mojo.lang.EnumDeclOrBuilder getEnumDeclOrBuilder();

  /**
   * <code>.mojo.lang.StructDecl struct_decl = 4;</code>
   * @return Whether the structDecl field is set.
   */
  boolean hasStructDecl();
  /**
   * <code>.mojo.lang.StructDecl struct_decl = 4;</code>
   * @return The structDecl.
   */
  org.mojolang.mojo.lang.StructDecl getStructDecl();
  /**
   * <code>.mojo.lang.StructDecl struct_decl = 4;</code>
   */
  org.mojolang.mojo.lang.StructDeclOrBuilder getStructDeclOrBuilder();

  /**
   * <code>.mojo.lang.TypeAliasDecl type_alias_decl = 5;</code>
   * @return Whether the typeAliasDecl field is set.
   */
  boolean hasTypeAliasDecl();
  /**
   * <code>.mojo.lang.TypeAliasDecl type_alias_decl = 5;</code>
   * @return The typeAliasDecl.
   */
  org.mojolang.mojo.lang.TypeAliasDecl getTypeAliasDecl();
  /**
   * <code>.mojo.lang.TypeAliasDecl type_alias_decl = 5;</code>
   */
  org.mojolang.mojo.lang.TypeAliasDeclOrBuilder getTypeAliasDeclOrBuilder();

  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 6;</code>
   * @return Whether the interfaceDecl field is set.
   */
  boolean hasInterfaceDecl();
  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 6;</code>
   * @return The interfaceDecl.
   */
  org.mojolang.mojo.lang.InterfaceDecl getInterfaceDecl();
  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 6;</code>
   */
  org.mojolang.mojo.lang.InterfaceDeclOrBuilder getInterfaceDeclOrBuilder();

  /**
   * <code>.mojo.lang.ConstantDecl constant_decl = 10;</code>
   * @return Whether the constantDecl field is set.
   */
  boolean hasConstantDecl();
  /**
   * <code>.mojo.lang.ConstantDecl constant_decl = 10;</code>
   * @return The constantDecl.
   */
  org.mojolang.mojo.lang.ConstantDecl getConstantDecl();
  /**
   * <code>.mojo.lang.ConstantDecl constant_decl = 10;</code>
   */
  org.mojolang.mojo.lang.ConstantDeclOrBuilder getConstantDeclOrBuilder();

  /**
   * <code>.mojo.lang.VariableDecl variable_decl = 11;</code>
   * @return Whether the variableDecl field is set.
   */
  boolean hasVariableDecl();
  /**
   * <code>.mojo.lang.VariableDecl variable_decl = 11;</code>
   * @return The variableDecl.
   */
  org.mojolang.mojo.lang.VariableDecl getVariableDecl();
  /**
   * <code>.mojo.lang.VariableDecl variable_decl = 11;</code>
   */
  org.mojolang.mojo.lang.VariableDeclOrBuilder getVariableDeclOrBuilder();

  /**
   * <code>.mojo.lang.AttributeDecl attribute_decl = 12;</code>
   * @return Whether the attributeDecl field is set.
   */
  boolean hasAttributeDecl();
  /**
   * <code>.mojo.lang.AttributeDecl attribute_decl = 12;</code>
   * @return The attributeDecl.
   */
  org.mojolang.mojo.lang.AttributeDecl getAttributeDecl();
  /**
   * <code>.mojo.lang.AttributeDecl attribute_decl = 12;</code>
   */
  org.mojolang.mojo.lang.AttributeDeclOrBuilder getAttributeDeclOrBuilder();

  /**
   * <code>.mojo.lang.FunctionDecl function_decl = 13;</code>
   * @return Whether the functionDecl field is set.
   */
  boolean hasFunctionDecl();
  /**
   * <code>.mojo.lang.FunctionDecl function_decl = 13;</code>
   * @return The functionDecl.
   */
  org.mojolang.mojo.lang.FunctionDecl getFunctionDecl();
  /**
   * <code>.mojo.lang.FunctionDecl function_decl = 13;</code>
   */
  org.mojolang.mojo.lang.FunctionDeclOrBuilder getFunctionDeclOrBuilder();

  /**
   * <code>.mojo.lang.ConstructorDecl constructor_decl = 14;</code>
   * @return Whether the constructorDecl field is set.
   */
  boolean hasConstructorDecl();
  /**
   * <code>.mojo.lang.ConstructorDecl constructor_decl = 14;</code>
   * @return The constructorDecl.
   */
  org.mojolang.mojo.lang.ConstructorDecl getConstructorDecl();
  /**
   * <code>.mojo.lang.ConstructorDecl constructor_decl = 14;</code>
   */
  org.mojolang.mojo.lang.ConstructorDeclOrBuilder getConstructorDeclOrBuilder();

  /**
   * <code>.mojo.lang.GenericParameter generic_parameter = 19;</code>
   * @return Whether the genericParameter field is set.
   */
  boolean hasGenericParameter();
  /**
   * <code>.mojo.lang.GenericParameter generic_parameter = 19;</code>
   * @return The genericParameter.
   */
  org.mojolang.mojo.lang.GenericParameter getGenericParameter();
  /**
   * <code>.mojo.lang.GenericParameter generic_parameter = 19;</code>
   */
  org.mojolang.mojo.lang.GenericParameterOrBuilder getGenericParameterOrBuilder();

  /**
   * <code>.mojo.lang.GroupDecl group_decl = 20;</code>
   * @return Whether the groupDecl field is set.
   */
  boolean hasGroupDecl();
  /**
   * <code>.mojo.lang.GroupDecl group_decl = 20;</code>
   * @return The groupDecl.
   */
  org.mojolang.mojo.lang.GroupDecl getGroupDecl();
  /**
   * <code>.mojo.lang.GroupDecl group_decl = 20;</code>
   */
  org.mojolang.mojo.lang.GroupDeclOrBuilder getGroupDeclOrBuilder();

  public org.mojolang.mojo.lang.Declaration.DeclarationCase getDeclarationCase();
}
