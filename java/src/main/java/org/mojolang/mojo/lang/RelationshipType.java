// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

/**
 * Protobuf enum {@code mojo.lang.RelationType}
 */
public enum RelationType
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <code>RELATIONSHIP_TYPE_UNSPECIFIED = 0;</code>
   */
  RELATIONSHIP_TYPE_UNSPECIFIED(0),
  /**
   * <code>RELATIONSHIP_TYPE_O2_O = 15;</code>
   */
  RELATIONSHIP_TYPE_O2_O(15),
  /**
   * <code>RELATIONSHIP_TYPE_O2_O_TWO_TYPES = 1;</code>
   */
  RELATIONSHIP_TYPE_O2_O_TWO_TYPES(1),
  /**
   * <code>RELATIONSHIP_TYPE_O2_O_SAME_TYPE = 2;</code>
   */
  RELATIONSHIP_TYPE_O2_O_SAME_TYPE(2),
  /**
   * <code>RELATIONSHIP_TYPE_O2_O_BIDIRECTIONAL = 3;</code>
   */
  RELATIONSHIP_TYPE_O2_O_BIDIRECTIONAL(3),
  /**
   * <code>RELATIONSHIP_TYPE_O2_M = 240;</code>
   */
  RELATIONSHIP_TYPE_O2_M(240),
  /**
   * <code>RELATIONSHIP_TYPE_O2_M_TWO_TYPES = 16;</code>
   */
  RELATIONSHIP_TYPE_O2_M_TWO_TYPES(16),
  /**
   * <code>RELATIONSHIP_TYPE_O2_M_SAME_TYPE = 32;</code>
   */
  RELATIONSHIP_TYPE_O2_M_SAME_TYPE(32),
  /**
   * <code>RELATIONSHIP_TYPE_M2_M = 3840;</code>
   */
  RELATIONSHIP_TYPE_M2_M(3840),
  /**
   * <code>RELATIONSHIP_TYPE_M2_M_TWO_TYPES = 256;</code>
   */
  RELATIONSHIP_TYPE_M2_M_TWO_TYPES(256),
  /**
   * <code>RELATIONSHIP_TYPE_M2_M_SAME_TYPE = 512;</code>
   */
  RELATIONSHIP_TYPE_M2_M_SAME_TYPE(512),
  /**
   * <code>RELATIONSHIP_TYPE_M2_M_BIDIRECTIONAL = 768;</code>
   */
  RELATIONSHIP_TYPE_M2_M_BIDIRECTIONAL(768),
  UNRECOGNIZED(-1),
  ;

  /**
   * <code>RELATIONSHIP_TYPE_UNSPECIFIED = 0;</code>
   */
  public static final int RELATIONSHIP_TYPE_UNSPECIFIED_VALUE = 0;
  /**
   * <code>RELATIONSHIP_TYPE_O2_O = 15;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_O_VALUE = 15;
  /**
   * <code>RELATIONSHIP_TYPE_O2_O_TWO_TYPES = 1;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_O_TWO_TYPES_VALUE = 1;
  /**
   * <code>RELATIONSHIP_TYPE_O2_O_SAME_TYPE = 2;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_O_SAME_TYPE_VALUE = 2;
  /**
   * <code>RELATIONSHIP_TYPE_O2_O_BIDIRECTIONAL = 3;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_O_BIDIRECTIONAL_VALUE = 3;
  /**
   * <code>RELATIONSHIP_TYPE_O2_M = 240;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_M_VALUE = 240;
  /**
   * <code>RELATIONSHIP_TYPE_O2_M_TWO_TYPES = 16;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_M_TWO_TYPES_VALUE = 16;
  /**
   * <code>RELATIONSHIP_TYPE_O2_M_SAME_TYPE = 32;</code>
   */
  public static final int RELATIONSHIP_TYPE_O2_M_SAME_TYPE_VALUE = 32;
  /**
   * <code>RELATIONSHIP_TYPE_M2_M = 3840;</code>
   */
  public static final int RELATIONSHIP_TYPE_M2_M_VALUE = 3840;
  /**
   * <code>RELATIONSHIP_TYPE_M2_M_TWO_TYPES = 256;</code>
   */
  public static final int RELATIONSHIP_TYPE_M2_M_TWO_TYPES_VALUE = 256;
  /**
   * <code>RELATIONSHIP_TYPE_M2_M_SAME_TYPE = 512;</code>
   */
  public static final int RELATIONSHIP_TYPE_M2_M_SAME_TYPE_VALUE = 512;
  /**
   * <code>RELATIONSHIP_TYPE_M2_M_BIDIRECTIONAL = 768;</code>
   */
  public static final int RELATIONSHIP_TYPE_M2_M_BIDIRECTIONAL_VALUE = 768;


  public final int getNumber() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalArgumentException(
          "Can't get the number of an unknown enum value.");
    }
    return value;
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   * @deprecated Use {@link #forNumber(int)} instead.
   */
  @java.lang.Deprecated
  public static RelationType valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static RelationType forNumber(int value) {
    switch (value) {
      case 0: return RELATIONSHIP_TYPE_UNSPECIFIED;
      case 15: return RELATIONSHIP_TYPE_O2_O;
      case 1: return RELATIONSHIP_TYPE_O2_O_TWO_TYPES;
      case 2: return RELATIONSHIP_TYPE_O2_O_SAME_TYPE;
      case 3: return RELATIONSHIP_TYPE_O2_O_BIDIRECTIONAL;
      case 240: return RELATIONSHIP_TYPE_O2_M;
      case 16: return RELATIONSHIP_TYPE_O2_M_TWO_TYPES;
      case 32: return RELATIONSHIP_TYPE_O2_M_SAME_TYPE;
      case 3840: return RELATIONSHIP_TYPE_M2_M;
      case 256: return RELATIONSHIP_TYPE_M2_M_TWO_TYPES;
      case 512: return RELATIONSHIP_TYPE_M2_M_SAME_TYPE;
      case 768: return RELATIONSHIP_TYPE_M2_M_BIDIRECTIONAL;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<RelationType>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      RelationType> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<RelationType>() {
          public RelationType findValueByNumber(int number) {
            return RelationType.forNumber(number);
          }
        };

  public final com.google.protobuf.Descriptors.EnumValueDescriptor
      getValueDescriptor() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalStateException(
          "Can't get the descriptor of an unrecognized enum value.");
    }
    return getDescriptor().getValues().get(ordinal());
  }
  public final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptorForType() {
    return getDescriptor();
  }
  public static final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptor() {
    return org.mojolang.mojo.lang.LangProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final RelationType[] VALUES = values();

  public static RelationType valueOf(
      com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
    if (desc.getType() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "EnumValueDescriptor is not for this type.");
    }
    if (desc.getIndex() == -1) {
      return UNRECOGNIZED;
    }
    return VALUES[desc.getIndex()];
  }

  private final int value;

  private RelationType(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:mojo.lang.RelationType)
}

