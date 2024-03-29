// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface StructTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.StructType)
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
   * <code>repeated .mojo.lang.ValueDecl fields = 10;</code>
   */
  java.util.List<org.mojolang.mojo.lang.ValueDecl> 
      getFieldsList();
  /**
   * <code>repeated .mojo.lang.ValueDecl fields = 10;</code>
   */
  org.mojolang.mojo.lang.ValueDecl getFields(int index);
  /**
   * <code>repeated .mojo.lang.ValueDecl fields = 10;</code>
   */
  int getFieldsCount();
  /**
   * <code>repeated .mojo.lang.ValueDecl fields = 10;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ValueDeclOrBuilder> 
      getFieldsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.ValueDecl fields = 10;</code>
   */
  org.mojolang.mojo.lang.ValueDeclOrBuilder getFieldsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.NominalType inherits = 11;</code>
   */
  java.util.List<org.mojolang.mojo.lang.NominalType> 
      getInheritsList();
  /**
   * <code>repeated .mojo.lang.NominalType inherits = 11;</code>
   */
  org.mojolang.mojo.lang.NominalType getInherits(int index);
  /**
   * <code>repeated .mojo.lang.NominalType inherits = 11;</code>
   */
  int getInheritsCount();
  /**
   * <code>repeated .mojo.lang.NominalType inherits = 11;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.NominalTypeOrBuilder> 
      getInheritsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.NominalType inherits = 11;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getInheritsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.GroupDecl groups = 12;</code>
   */
  java.util.List<org.mojolang.mojo.lang.GroupDecl> 
      getGroupsList();
  /**
   * <code>repeated .mojo.lang.GroupDecl groups = 12;</code>
   */
  org.mojolang.mojo.lang.GroupDecl getGroups(int index);
  /**
   * <code>repeated .mojo.lang.GroupDecl groups = 12;</code>
   */
  int getGroupsCount();
  /**
   * <code>repeated .mojo.lang.GroupDecl groups = 12;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.GroupDeclOrBuilder> 
      getGroupsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.GroupDecl groups = 12;</code>
   */
  org.mojolang.mojo.lang.GroupDeclOrBuilder getGroupsOrBuilder(
      int index);

  /**
   * <code>.mojo.lang.Position inherit_position = 19;</code>
   * @return Whether the inheritPosition field is set.
   */
  boolean hasInheritPosition();
  /**
   * <code>.mojo.lang.Position inherit_position = 19;</code>
   * @return The inheritPosition.
   */
  org.mojolang.mojo.lang.Position getInheritPosition();
  /**
   * <code>.mojo.lang.Position inherit_position = 19;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getInheritPositionOrBuilder();
}
