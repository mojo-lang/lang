// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

/**
 * Protobuf type {@code mojo.lang.ReturnStmt}
 */
public final class ReturnStmt extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.lang.ReturnStmt)
    ReturnStmtOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ReturnStmt.newBuilder() to construct.
  private ReturnStmt(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ReturnStmt() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new ReturnStmt();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private ReturnStmt(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            org.mojolang.mojo.lang.Position.Builder subBuilder = null;
            if (startPosition_ != null) {
              subBuilder = startPosition_.toBuilder();
            }
            startPosition_ = input.readMessage(org.mojolang.mojo.lang.Position.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(startPosition_);
              startPosition_ = subBuilder.buildPartial();
            }

            break;
          }
          case 18: {
            org.mojolang.mojo.lang.Position.Builder subBuilder = null;
            if (endPosition_ != null) {
              subBuilder = endPosition_.toBuilder();
            }
            endPosition_ = input.readMessage(org.mojolang.mojo.lang.Position.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(endPosition_);
              endPosition_ = subBuilder.buildPartial();
            }

            break;
          }
          case 24: {

            kind_ = input.readInt64();
            break;
          }
          case 32: {

            implicit_ = input.readBool();
            break;
          }
          case 82: {
            org.mojolang.mojo.lang.Expression.Builder subBuilder = null;
            if (result_ != null) {
              subBuilder = result_.toBuilder();
            }
            result_ = input.readMessage(org.mojolang.mojo.lang.Expression.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(result_);
              result_ = subBuilder.buildPartial();
            }

            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_ReturnStmt_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_ReturnStmt_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.lang.ReturnStmt.class, org.mojolang.mojo.lang.ReturnStmt.Builder.class);
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
    return getStartPosition();
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
    return getEndPosition();
  }

  public static final int KIND_FIELD_NUMBER = 3;
  private long kind_;
  /**
   * <code>int64 kind = 3;</code>
   * @return The kind.
   */
  @java.lang.Override
  public long getKind() {
    return kind_;
  }

  public static final int IMPLICIT_FIELD_NUMBER = 4;
  private boolean implicit_;
  /**
   * <code>bool implicit = 4;</code>
   * @return The implicit.
   */
  @java.lang.Override
  public boolean getImplicit() {
    return implicit_;
  }

  public static final int RESULT_FIELD_NUMBER = 10;
  private org.mojolang.mojo.lang.Expression result_;
  /**
   * <code>.mojo.lang.Expression result = 10;</code>
   * @return Whether the result field is set.
   */
  @java.lang.Override
  public boolean hasResult() {
    return result_ != null;
  }
  /**
   * <code>.mojo.lang.Expression result = 10;</code>
   * @return The result.
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.Expression getResult() {
    return result_ == null ? org.mojolang.mojo.lang.Expression.getDefaultInstance() : result_;
  }
  /**
   * <code>.mojo.lang.Expression result = 10;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.ExpressionOrBuilder getResultOrBuilder() {
    return getResult();
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
    if (kind_ != 0L) {
      output.writeInt64(3, kind_);
    }
    if (implicit_ != false) {
      output.writeBool(4, implicit_);
    }
    if (result_ != null) {
      output.writeMessage(10, getResult());
    }
    unknownFields.writeTo(output);
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
    if (kind_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(3, kind_);
    }
    if (implicit_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(4, implicit_);
    }
    if (result_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(10, getResult());
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.lang.ReturnStmt)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.lang.ReturnStmt other = (org.mojolang.mojo.lang.ReturnStmt) obj;

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
    if (hasResult() != other.hasResult()) return false;
    if (hasResult()) {
      if (!getResult()
          .equals(other.getResult())) return false;
    }
    if (!unknownFields.equals(other.unknownFields)) return false;
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
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getKind());
    hash = (37 * hash) + IMPLICIT_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getImplicit());
    if (hasResult()) {
      hash = (37 * hash) + RESULT_FIELD_NUMBER;
      hash = (53 * hash) + getResult().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.ReturnStmt parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.lang.ReturnStmt prototype) {
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
   * Protobuf type {@code mojo.lang.ReturnStmt}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.lang.ReturnStmt)
      org.mojolang.mojo.lang.ReturnStmtOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_ReturnStmt_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_ReturnStmt_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.lang.ReturnStmt.class, org.mojolang.mojo.lang.ReturnStmt.Builder.class);
    }

    // Construct using org.mojolang.mojo.lang.ReturnStmt.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (startPositionBuilder_ == null) {
        startPosition_ = null;
      } else {
        startPosition_ = null;
        startPositionBuilder_ = null;
      }
      if (endPositionBuilder_ == null) {
        endPosition_ = null;
      } else {
        endPosition_ = null;
        endPositionBuilder_ = null;
      }
      kind_ = 0L;

      implicit_ = false;

      if (resultBuilder_ == null) {
        result_ = null;
      } else {
        result_ = null;
        resultBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.lang.LangProto.internal_static_mojo_lang_ReturnStmt_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.ReturnStmt getDefaultInstanceForType() {
      return org.mojolang.mojo.lang.ReturnStmt.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.ReturnStmt build() {
      org.mojolang.mojo.lang.ReturnStmt result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.ReturnStmt buildPartial() {
      org.mojolang.mojo.lang.ReturnStmt result = new org.mojolang.mojo.lang.ReturnStmt(this);
      if (startPositionBuilder_ == null) {
        result.startPosition_ = startPosition_;
      } else {
        result.startPosition_ = startPositionBuilder_.build();
      }
      if (endPositionBuilder_ == null) {
        result.endPosition_ = endPosition_;
      } else {
        result.endPosition_ = endPositionBuilder_.build();
      }
      result.kind_ = kind_;
      result.implicit_ = implicit_;
      if (resultBuilder_ == null) {
        result.result_ = result_;
      } else {
        result.result_ = resultBuilder_.build();
      }
      onBuilt();
      return result;
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
      if (other instanceof org.mojolang.mojo.lang.ReturnStmt) {
        return mergeFrom((org.mojolang.mojo.lang.ReturnStmt)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.lang.ReturnStmt other) {
      if (other == org.mojolang.mojo.lang.ReturnStmt.getDefaultInstance()) return this;
      if (other.hasStartPosition()) {
        mergeStartPosition(other.getStartPosition());
      }
      if (other.hasEndPosition()) {
        mergeEndPosition(other.getEndPosition());
      }
      if (other.getKind() != 0L) {
        setKind(other.getKind());
      }
      if (other.getImplicit() != false) {
        setImplicit(other.getImplicit());
      }
      if (other.hasResult()) {
        mergeResult(other.getResult());
      }
      this.mergeUnknownFields(other.unknownFields);
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
      org.mojolang.mojo.lang.ReturnStmt parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.lang.ReturnStmt) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private org.mojolang.mojo.lang.Position startPosition_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> startPositionBuilder_;
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     * @return Whether the startPosition field is set.
     */
    public boolean hasStartPosition() {
      return startPositionBuilder_ != null || startPosition_ != null;
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
        onChanged();
      } else {
        startPositionBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder setStartPosition(
        org.mojolang.mojo.lang.Position.Builder builderForValue) {
      if (startPositionBuilder_ == null) {
        startPosition_ = builderForValue.build();
        onChanged();
      } else {
        startPositionBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder mergeStartPosition(org.mojolang.mojo.lang.Position value) {
      if (startPositionBuilder_ == null) {
        if (startPosition_ != null) {
          startPosition_ =
            org.mojolang.mojo.lang.Position.newBuilder(startPosition_).mergeFrom(value).buildPartial();
        } else {
          startPosition_ = value;
        }
        onChanged();
      } else {
        startPositionBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public Builder clearStartPosition() {
      if (startPositionBuilder_ == null) {
        startPosition_ = null;
        onChanged();
      } else {
        startPosition_ = null;
        startPositionBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start_position = 1;</code>
     */
    public org.mojolang.mojo.lang.Position.Builder getStartPositionBuilder() {
      
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
      return endPositionBuilder_ != null || endPosition_ != null;
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
        onChanged();
      } else {
        endPositionBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder setEndPosition(
        org.mojolang.mojo.lang.Position.Builder builderForValue) {
      if (endPositionBuilder_ == null) {
        endPosition_ = builderForValue.build();
        onChanged();
      } else {
        endPositionBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder mergeEndPosition(org.mojolang.mojo.lang.Position value) {
      if (endPositionBuilder_ == null) {
        if (endPosition_ != null) {
          endPosition_ =
            org.mojolang.mojo.lang.Position.newBuilder(endPosition_).mergeFrom(value).buildPartial();
        } else {
          endPosition_ = value;
        }
        onChanged();
      } else {
        endPositionBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public Builder clearEndPosition() {
      if (endPositionBuilder_ == null) {
        endPosition_ = null;
        onChanged();
      } else {
        endPosition_ = null;
        endPositionBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end_position = 2;</code>
     */
    public org.mojolang.mojo.lang.Position.Builder getEndPositionBuilder() {
      
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

    private long kind_ ;
    /**
     * <code>int64 kind = 3;</code>
     * @return The kind.
     */
    @java.lang.Override
    public long getKind() {
      return kind_;
    }
    /**
     * <code>int64 kind = 3;</code>
     * @param value The kind to set.
     * @return This builder for chaining.
     */
    public Builder setKind(long value) {
      
      kind_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 kind = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearKind() {
      
      kind_ = 0L;
      onChanged();
      return this;
    }

    private boolean implicit_ ;
    /**
     * <code>bool implicit = 4;</code>
     * @return The implicit.
     */
    @java.lang.Override
    public boolean getImplicit() {
      return implicit_;
    }
    /**
     * <code>bool implicit = 4;</code>
     * @param value The implicit to set.
     * @return This builder for chaining.
     */
    public Builder setImplicit(boolean value) {
      
      implicit_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bool implicit = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearImplicit() {
      
      implicit_ = false;
      onChanged();
      return this;
    }

    private org.mojolang.mojo.lang.Expression result_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Expression, org.mojolang.mojo.lang.Expression.Builder, org.mojolang.mojo.lang.ExpressionOrBuilder> resultBuilder_;
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     * @return Whether the result field is set.
     */
    public boolean hasResult() {
      return resultBuilder_ != null || result_ != null;
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     * @return The result.
     */
    public org.mojolang.mojo.lang.Expression getResult() {
      if (resultBuilder_ == null) {
        return result_ == null ? org.mojolang.mojo.lang.Expression.getDefaultInstance() : result_;
      } else {
        return resultBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    public Builder setResult(org.mojolang.mojo.lang.Expression value) {
      if (resultBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        result_ = value;
        onChanged();
      } else {
        resultBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    public Builder setResult(
        org.mojolang.mojo.lang.Expression.Builder builderForValue) {
      if (resultBuilder_ == null) {
        result_ = builderForValue.build();
        onChanged();
      } else {
        resultBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    public Builder mergeResult(org.mojolang.mojo.lang.Expression value) {
      if (resultBuilder_ == null) {
        if (result_ != null) {
          result_ =
            org.mojolang.mojo.lang.Expression.newBuilder(result_).mergeFrom(value).buildPartial();
        } else {
          result_ = value;
        }
        onChanged();
      } else {
        resultBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    public Builder clearResult() {
      if (resultBuilder_ == null) {
        result_ = null;
        onChanged();
      } else {
        result_ = null;
        resultBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    public org.mojolang.mojo.lang.Expression.Builder getResultBuilder() {
      
      onChanged();
      return getResultFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    public org.mojolang.mojo.lang.ExpressionOrBuilder getResultOrBuilder() {
      if (resultBuilder_ != null) {
        return resultBuilder_.getMessageOrBuilder();
      } else {
        return result_ == null ?
            org.mojolang.mojo.lang.Expression.getDefaultInstance() : result_;
      }
    }
    /**
     * <code>.mojo.lang.Expression result = 10;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Expression, org.mojolang.mojo.lang.Expression.Builder, org.mojolang.mojo.lang.ExpressionOrBuilder> 
        getResultFieldBuilder() {
      if (resultBuilder_ == null) {
        resultBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.lang.Expression, org.mojolang.mojo.lang.Expression.Builder, org.mojolang.mojo.lang.ExpressionOrBuilder>(
                getResult(),
                getParentForChildren(),
                isClean());
        result_ = null;
      }
      return resultBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.lang.ReturnStmt)
  }

  // @@protoc_insertion_point(class_scope:mojo.lang.ReturnStmt)
  private static final org.mojolang.mojo.lang.ReturnStmt DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.lang.ReturnStmt();
  }

  public static org.mojolang.mojo.lang.ReturnStmt getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ReturnStmt>
      PARSER = new com.google.protobuf.AbstractParser<ReturnStmt>() {
    @java.lang.Override
    public ReturnStmt parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new ReturnStmt(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<ReturnStmt> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ReturnStmt> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.lang.ReturnStmt getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

