// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface FunctionSignatureOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.FunctionSignature)
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
   * <code>.mojo.lang.FunctionSignature.Parameter parameter = 10;</code>
   * @return Whether the parameter field is set.
   */
  boolean hasParameter();
  /**
   * <code>.mojo.lang.FunctionSignature.Parameter parameter = 10;</code>
   * @return The parameter.
   */
  org.mojolang.mojo.lang.FunctionSignature.Parameter getParameter();
  /**
   * <code>.mojo.lang.FunctionSignature.Parameter parameter = 10;</code>
   */
  org.mojolang.mojo.lang.FunctionSignature.ParameterOrBuilder getParameterOrBuilder();

  /**
   * <code>.mojo.lang.FunctionSignature.Result result = 11;</code>
   * @return Whether the result field is set.
   */
  boolean hasResult();
  /**
   * <code>.mojo.lang.FunctionSignature.Result result = 11;</code>
   * @return The result.
   */
  org.mojolang.mojo.lang.FunctionSignature.Result getResult();
  /**
   * <code>.mojo.lang.FunctionSignature.Result result = 11;</code>
   */
  org.mojolang.mojo.lang.FunctionSignature.ResultOrBuilder getResultOrBuilder();
}