// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

/**
 * Protobuf type {@code mojo.lang.StructConstructionExpr}
 */
public final class StructConstructionExpr extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.lang.StructConstructionExpr)
    StructConstructionExprOrBuilder {
private static final long serialVersionUID = 0L;
  // Use StructConstructionExpr.newBuilder() to construct.
  private StructConstructionExpr(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private StructConstructionExpr() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new StructConstructionExpr();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_StructConstructionExpr_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_StructConstructionExpr_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.lang.StructConstructionExpr.class, org.mojolang.mojo.lang.StructConstructionExpr.Builder.class);
  }

  public static final int START_POSITION_FIELD_NUMBER = 1;
  private org.mojolang.mojo.lang.Position startPosition_;
  /**
   * <code>.mojo.lang.Position start_position = 1;</code>
   * @return Whether the startPosition field is set.
   */
  @java.lang.Override
  public boolean hasStartPosition() {
    return startPosition_ != null;
  }
  /**
   * <code>.mojo.lang.Position start_position = 1;</code>
   * @return The startPosition.
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.Position getStartPosition() {
    return startPosition_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : startPosition_;
  }
  /**
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.PositionOrBuilder getStartPositionOrBuilder() {
    return startPosition_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : startPosition_;
  }

  public static final int END_POSITION_FIELD_NUMBER = 2;
  private org.mojolang.mojo.lang.Position endPosition_;
  /**
   * <code>.mojo.lang.Position end_position = 2;</code>
   * @return Whether the endPosition field is set.
   */
  @java.lang.Override
  public boolean hasEndPosition() {
    return endPosition_ != null;
  }
  /**
   * <code>.mojo.lang.Position end_position = 2;</code>
   * @return The endPosition.
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.Position getEndPosition() {
    return endPosition_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : endPosition_;
  }
  /**
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.PositionOrBuilder getEndPositionOrBuilder() {
    return endPosition_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : endPosition_;
  }

  public static final int KIND_FIELD_NUMBER = 4;
  private int kind_ = 0;
  /**
   * <code>int32 kind = 4;</code>
   * @return The kind.
   */
  @java.lang.Override
  public int getKind() {
    return kind_;
  }

  public static final int IMPLICIT_FIELD_NUMBER = 5;
  private boolean implicit_ = false;
  /**
   * <code>bool implicit = 5;</code>
   * @return The implicit.
   */
  @java.lang.Override
  public boolean getImplicit() {
    return implicit_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (startPosition_ != null) {
      output.writeMessage(1, getStartPosition());
    }
    if (endPosition_ != null) {
      output.writeMessage(2, getEndPosition());
    }
    if (kind_ != 0) {
      output.writeInt32(4, kind_);
    }
    if (implicit_ != false) {
      output.writeBool(5, implicit_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (startPosition_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getStartPosition());
    }
    if (endPosition_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getEndPosition());
    }
    if (kind_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(4, kind_);
    }
    if (implicit_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(5, implicit_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.lang.StructConstructionExpr)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.lang.StructConstructionExpr other = (org.mojolang.mojo.lang.StructConstructionExpr) obj;

    if (hasStartPosition() != other.hasStartPosition()) return false;
    if (hasStartPosition()) {
      if (!getStartPosition()
          .equals(other.getStartPosition())) return false;
    }
    if (hasEndPosition() != other.hasEndPosition()) return false;
    if (hasEndPosition()) {
      if (!getEndPosition()
          .equals(other.getEndPosition())) return false;
    }
    if (getKind()
        != other.getKind()) return false;
    if (getImplicit()
        != other.getImplicit()) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (hasStartPosition()) {
      hash = (37 * hash) + START_POSITION_FIELD_NUMBER;
      hash = (53 * hash) + getStartPosition().hashCode();
    }
    if (hasEndPosition()) {
      hash = (37 * hash) + END_POSITION_FIELD_NUMBER;
      hash = (53 * hash) + getEndPosition().hashCode();
    }
    hash = (37 * hash) + KIND_FIELD_NUMBER;
    hash = (53 * hash) + getKind();
    hash = (37 * hash) + IMPLICIT_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getImplicit());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.StructConstructionExpr parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojolang.mojo.lang.StructConstructionExpr prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code mojo.lang.StructConstructionExpr}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.lang.StructConstructionExpr)
      org.mojolang.mojo.lang.StructConstructionExprOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_StructConstructionExpr_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_StructConstructionExpr_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.lang.StructConstructionExpr.class, org.mojolang.mojo.lang.StructConstructionExpr.Builder.class);
    }

    // Construct using org.mojolang.mojo.lang.StructConstructionExpr.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      startPosition_ = null;
      if (startPositionBuilder_ != null) {
        startPositionBuilder_.dispose();
        startPositionBuilder_ = null;
      }
      endPosition_ = null;
      if (endPositionBuilder_ != null) {
        endPositionBuilder_.dispose();
        endPositionBuilder_ = null;
      }
      kind_ = 0;
      implicit_ = false;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_StructConstructionExpr_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.StructConstructionExpr getDefaultInstanceForType() {
      return org.mojolang.mojo.lang.StructConstructionExpr.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.StructConstructionExpr build() {
      org.mojolang.mojo.lang.StructConstructionExpr result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.StructConstructionExpr buildPartial() {
      org.mojolang.mojo.lang.StructConstructionExpr result = new org.mojolang.mojo.lang.StructConstructionExpr(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(org.mojolang.mojo.lang.StructConstructionExpr result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.startPosition_ = startPositionBuilder_ == null
            ? startPosition_
            : startPositionBuilder_.build();
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.endPosition_ = endPositionBuilder_ == null
            ? endPosition_
            : endPositionBuilder_.build();
      }
      if (((from_bitField0_ & 0x00000004) != 0)) {
        result.kind_ = kind_;
      }
      if (((from_bitField0_ & 0x00000008) != 0)) {
        result.implicit_ = implicit_;
      }
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojolang.mojo.lang.StructConstructionExpr) {
        return mergeFrom((org.mojolang.mojo.lang.StructConstructionExpr)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.lang.StructConstructionExpr other) {
      if (other == org.mojolang.mojo.lang.StructConstructionExpr.getDefaultInstance()) return this;
      if (other.hasStartPosition()) {
        mergeStartPosition(other.getStartPosition());
      }
      if (other.hasEndPosition()) {
        mergeEndPosition(other.getEndPosition());
      }
      if (other.getKind() != 0) {
        setKind(other.getKind());
      }
      if (other.getImplicit() != false) {
        setImplicit(other.getImplicit());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              input.readMessage(
                  getStartPositionFieldBuilder().getBuilder(),
                  extensionRegistry);
              bitField0_ |= 0x00000001;
              break;
            } // case 10
            case 18: {
              input.readMessage(
                  getEndPositionFieldBuilder().getBuilder(),
                  extensionRegistry);
              bitField0_ |= 0x00000002;
              break;
            } // case 18
            case 32: {
              kind_ = input.readInt32();
              bitField0_ |= 0x00000004;
              break;
            } // case 32
            case 40: {
              implicit_ = input.readBool();
              bitField0_ |= 0x00000008;
              break;
            } // case 40
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private org.mojolang.mojo.lang.Position startPosition_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> startPositionBuilder_;
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     * @return Whether the startPosition field is set.
     */
    public boolean hasStartPosition() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     * @return The startPosition.
     */
    public org.mojolang.mojo.lang.Position getStartPosition() {
      if (startPositionBuilder_ == null) {
        return startPosition_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : startPosition_;
      } else {
        return startPositionBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder setStartPosition(org.mojolang.mojo.lang.Position value) {
      if (startPositionBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        startPosition_ = value;
      } else {
        startPositionBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder setStartPosition(
        org.mojolang.mojo.lang.Position.Builder builderForValue) {
      if (startPositionBuilder_ == null) {
        startPosition_ = builderForValue.build();
      } else {
        startPositionBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder mergeStartPosition(org.mojolang.mojo.lang.Position value) {
      if (startPositionBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0) &&
          startPosition_ != null &&
          startPosition_ != org.mojolang.mojo.lang.Position.getDefaultInstance()) {
          getStartPositionBuilder().mergeFrom(value);
        } else {
          startPosition_ = value;
        }
      } else {
        startPositionBuilder_.mergeFrom(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder clearStartPosition() {
      bitField0_ = (bitField0_ & ~0x00000001);
      startPosition_ = null;
      if (startPositionBuilder_ != null) {
        startPositionBuilder_.dispose();
        startPositionBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public org.mojolang.mojo.lang.Position.Builder getStartPositionBuilder() {
      bitField0_ |= 0x00000001;
      onChanged();
      return getStartPositionFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public org.mojolang.mojo.lang.PositionOrBuilder getStartPositionOrBuilder() {
      if (startPositionBuilder_ != null) {
        return startPositionBuilder_.getMessageOrBuilder();
      } else {
        return startPosition_ == null ?
            org.mojolang.mojo.lang.Position.getDefaultInstance() : startPosition_;
      }
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> 
        getStartPositionFieldBuilder() {
      if (startPositionBuilder_ == null) {
        startPositionBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder>(
                getStartPosition(),
                getParentForChildren(),
                isClean());
        startPosition_ = null;
      }
      return startPositionBuilder_;
    }

    private org.mojolang.mojo.lang.Position endPosition_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> endPositionBuilder_;
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     * @return Whether the endPosition field is set.
     */
    public boolean hasEndPosition() {
      return ((bitField0_ & 0x00000002) != 0);
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     * @return The endPosition.
     */
    public org.mojolang.mojo.lang.Position getEndPosition() {
      if (endPositionBuilder_ == null) {
        return endPosition_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : endPosition_;
      } else {
        return endPositionBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder setEndPosition(org.mojolang.mojo.lang.Position value) {
      if (endPositionBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        endPosition_ = value;
      } else {
        endPositionBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder setEndPosition(
        org.mojolang.mojo.lang.Position.Builder builderForValue) {
      if (endPositionBuilder_ == null) {
        endPosition_ = builderForValue.build();
      } else {
        endPositionBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder mergeEndPosition(org.mojolang.mojo.lang.Position value) {
      if (endPositionBuilder_ == null) {
        if (((bitField0_ & 0x00000002) != 0) &&
          endPosition_ != null &&
          endPosition_ != org.mojolang.mojo.lang.Position.getDefaultInstance()) {
          getEndPositionBuilder().mergeFrom(value);
        } else {
          endPosition_ = value;
        }
      } else {
        endPositionBuilder_.mergeFrom(value);
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder clearEndPosition() {
      bitField0_ = (bitField0_ & ~0x00000002);
      endPosition_ = null;
      if (endPositionBuilder_ != null) {
        endPositionBuilder_.dispose();
        endPositionBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public org.mojolang.mojo.lang.Position.Builder getEndPositionBuilder() {
      bitField0_ |= 0x00000002;
      onChanged();
      return getEndPositionFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public org.mojolang.mojo.lang.PositionOrBuilder getEndPositionOrBuilder() {
      if (endPositionBuilder_ != null) {
        return endPositionBuilder_.getMessageOrBuilder();
      } else {
        return endPosition_ == null ?
            org.mojolang.mojo.lang.Position.getDefaultInstance() : endPosition_;
      }
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> 
        getEndPositionFieldBuilder() {
      if (endPositionBuilder_ == null) {
        endPositionBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder>(
                getEndPosition(),
                getParentForChildren(),
                isClean());
        endPosition_ = null;
      }
      return endPositionBuilder_;
    }

    private int kind_ ;
    /**
     * <code>int32 kind = 4;</code>
     * @return The kind.
     */
    @java.lang.Override
    public int getKind() {
      return kind_;
    }
    /**
     * <code>int32 kind = 4;</code>
     * @param value The kind to set.
     * @return This builder for chaining.
     */
    public Builder setKind(int value) {
      
      kind_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>int32 kind = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearKind() {
      bitField0_ = (bitField0_ & ~0x00000004);
      kind_ = 0;
      onChanged();
      return this;
    }

    private boolean implicit_ ;
    /**
     * <code>bool implicit = 5;</code>
     * @return The implicit.
     */
    @java.lang.Override
    public boolean getImplicit() {
      return implicit_;
    }
    /**
     * <code>bool implicit = 5;</code>
     * @param value The implicit to set.
     * @return This builder for chaining.
     */
    public Builder setImplicit(boolean value) {
      
      implicit_ = value;
      bitField0_ |= 0x00000008;
      onChanged();
      return this;
    }
    /**
     * <code>bool implicit = 5;</code>
     * @return This builder for chaining.
     */
    public Builder clearImplicit() {
      bitField0_ = (bitField0_ & ~0x00000008);
      implicit_ = false;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.lang.StructConstructionExpr)
  }

  // @@protoc_insertion_point(class_scope:mojo.lang.StructConstructionExpr)
  private static final org.mojolang.mojo.lang.StructConstructionExpr DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.lang.StructConstructionExpr();
  }

  public static org.mojolang.mojo.lang.StructConstructionExpr getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<StructConstructionExpr>
      PARSER = new com.google.protobuf.AbstractParser<StructConstructionExpr>() {
    @java.lang.Override
    public StructConstructionExpr parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<StructConstructionExpr> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<StructConstructionExpr> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.lang.StructConstructionExpr getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

