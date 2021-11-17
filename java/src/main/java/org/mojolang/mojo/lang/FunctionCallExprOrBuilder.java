// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface FunctionCallExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.FunctionCallExpr)
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
   * <code>int32 kind = 3;</code>
   * @return The kind.
   */
  int getKind();

  /**
   * <code>bool implicit = 4;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>.mojo.lang.Expression callee = 10;</code>
   * @return Whether the callee field is set.
   */
  boolean hasCallee();
  /**
   * <code>.mojo.lang.Expression callee = 10;</code>
   * @return The callee.
   */
  org.mojolang.mojo.lang.Expression getCallee();
  /**
   * <code>.mojo.lang.Expression callee = 10;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getCalleeOrBuilder();

  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 11;</code>
   */
  java.util.List<org.mojolang.mojo.lang.NominalType> 
      getGenericArgumentsList();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 11;</code>
   */
  org.mojolang.mojo.lang.NominalType getGenericArguments(int index);
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 11;</code>
   */
  int getGenericArgumentsCount();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 11;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.NominalTypeOrBuilder> 
      getGenericArgumentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.NominalType generic_arguments = 11;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getGenericArgumentsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Argument arguments = 12;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Argument> 
      getArgumentsList();
  /**
   * <code>repeated .mojo.lang.Argument arguments = 12;</code>
   */
  org.mojolang.mojo.lang.Argument getArguments(int index);
  /**
   * <code>repeated .mojo.lang.Argument arguments = 12;</code>
   */
  int getArgumentsCount();
  /**
   * <code>repeated .mojo.lang.Argument arguments = 12;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ArgumentOrBuilder> 
      getArgumentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Argument arguments = 12;</code>
   */
  org.mojolang.mojo.lang.ArgumentOrBuilder getArgumentsOrBuilder(
      int index);
}
