// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface FloatLiteralExprOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.FloatLiteralExpr)
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
   * <code>int32 kind = 4;</code>
   * @return The kind.
   */
  int getKind();

  /**
   * <code>bool implicit = 5;</code>
   * @return The implicit.
   */
  boolean getImplicit();

  /**
   * <code>bool is_negative = 20;</code>
   * @return The isNegative.
   */
  boolean getIsNegative();

  /**
   * <code>double value = 30;</code>
   * @return The value.
   */
  double getValue();

  /**
   * <code>int64 integral_part = 35;</code>
   * @return The integralPart.
   */
  long getIntegralPart();

  /**
   * <code>int64 fractional_part = 36;</code>
   * @return The fractionalPart.
   */
  long getFractionalPart();

  /**
   * <code>int64 exponent_part = 37;</code>
   * @return The exponentPart.
   */
  long getExponentPart();
}
