// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface SourceFileOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.SourceFile)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string name = 1;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 1;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>string full_name = 2;</code>
   * @return The fullName.
   */
  java.lang.String getFullName();
  /**
   * <code>string full_name = 2;</code>
   * @return The bytes for fullName.
   */
  com.google.protobuf.ByteString
      getFullNameBytes();

  /**
   * <code>.mojo.core.Url url = 3;</code>
   * @return Whether the url field is set.
   */
  boolean hasUrl();
  /**
   * <code>.mojo.core.Url url = 3;</code>
   * @return The url.
   */
  org.mojolang.mojo.core.Url getUrl();
  /**
   * <code>.mojo.core.Url url = 3;</code>
   */
  org.mojolang.mojo.core.UrlOrBuilder getUrlOrBuilder();

  /**
   * <code>repeated .mojo.lang.Attribute attributes = 4;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Attribute> 
      getAttributesList();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 4;</code>
   */
  org.mojolang.mojo.lang.Attribute getAttributes(int index);
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 4;</code>
   */
  int getAttributesCount();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 4;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.AttributeOrBuilder> 
      getAttributesOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Attribute attributes = 4;</code>
   */
  org.mojolang.mojo.lang.AttributeOrBuilder getAttributesOrBuilder(
      int index);

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
   * <code>repeated .mojo.lang.ImportDecl imports = 11;</code>
   */
  java.util.List<org.mojolang.mojo.lang.ImportDecl> 
      getImportsList();
  /**
   * <code>repeated .mojo.lang.ImportDecl imports = 11;</code>
   */
  org.mojolang.mojo.lang.ImportDecl getImports(int index);
  /**
   * <code>repeated .mojo.lang.ImportDecl imports = 11;</code>
   */
  int getImportsCount();
  /**
   * <code>repeated .mojo.lang.ImportDecl imports = 11;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ImportDeclOrBuilder> 
      getImportsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.ImportDecl imports = 11;</code>
   */
  org.mojolang.mojo.lang.ImportDeclOrBuilder getImportsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Statement statements = 12;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Statement> 
      getStatementsList();
  /**
   * <code>repeated .mojo.lang.Statement statements = 12;</code>
   */
  org.mojolang.mojo.lang.Statement getStatements(int index);
  /**
   * <code>repeated .mojo.lang.Statement statements = 12;</code>
   */
  int getStatementsCount();
  /**
   * <code>repeated .mojo.lang.Statement statements = 12;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.StatementOrBuilder> 
      getStatementsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Statement statements = 12;</code>
   */
  org.mojolang.mojo.lang.StatementOrBuilder getStatementsOrBuilder(
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

  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 20;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Identifier> 
      getResolvedIdentifiersList();
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 20;</code>
   */
  org.mojolang.mojo.lang.Identifier getResolvedIdentifiers(int index);
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 20;</code>
   */
  int getResolvedIdentifiersCount();
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 20;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getResolvedIdentifiersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Identifier resolved_identifiers = 20;</code>
   */
  org.mojolang.mojo.lang.IdentifierOrBuilder getResolvedIdentifiersOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 21;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Identifier> 
      getUnresolvedIdentifiersList();
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 21;</code>
   */
  org.mojolang.mojo.lang.Identifier getUnresolvedIdentifiers(int index);
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 21;</code>
   */
  int getUnresolvedIdentifiersCount();
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 21;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getUnresolvedIdentifiersOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Identifier unresolved_identifiers = 21;</code>
   */
  org.mojolang.mojo.lang.IdentifierOrBuilder getUnresolvedIdentifiersOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.SourceFile dependencies = 22;</code>
   */
  java.util.List<org.mojolang.mojo.lang.SourceFile> 
      getDependenciesList();
  /**
   * <code>repeated .mojo.lang.SourceFile dependencies = 22;</code>
   */
  org.mojolang.mojo.lang.SourceFile getDependencies(int index);
  /**
   * <code>repeated .mojo.lang.SourceFile dependencies = 22;</code>
   */
  int getDependenciesCount();
  /**
   * <code>repeated .mojo.lang.SourceFile dependencies = 22;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.SourceFileOrBuilder> 
      getDependenciesOrBuilderList();
  /**
   * <code>repeated .mojo.lang.SourceFile dependencies = 22;</code>
   */
  org.mojolang.mojo.lang.SourceFileOrBuilder getDependenciesOrBuilder(
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
