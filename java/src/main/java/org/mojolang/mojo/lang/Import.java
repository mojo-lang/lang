// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/import.proto

package org.mojolang.mojo.lang;

/**
 * Protobuf type {@code mojo.lang.Import}
 */
public final class Import extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.lang.Import)
    ImportOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Import.newBuilder() to construct.
  private Import(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Import() {
    packageName_ = "";
    identifiers_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Import();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Import(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
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
          case 82: {
            java.lang.String s = input.readStringRequireUtf8();

            packageName_ = s;
            break;
          }
          case 106: {
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              identifiers_ = new java.util.ArrayList<org.mojolang.mojo.lang.Identifier>();
              mutable_bitField0_ |= 0x00000001;
            }
            identifiers_.add(
                input.readMessage(org.mojolang.mojo.lang.Identifier.parser(), extensionRegistry));
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
      if (((mutable_bitField0_ & 0x00000001) != 0)) {
        identifiers_ = java.util.Collections.unmodifiableList(identifiers_);
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.lang.ImportProto.internal_static_mojo_lang_Import_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.lang.ImportProto.internal_static_mojo_lang_Import_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.lang.Import.class, org.mojolang.mojo.lang.Import.Builder.class);
  }

  public static final int PACKAGE_NAME_FIELD_NUMBER = 10;
  private volatile java.lang.Object packageName_;
  /**
   * <code>string package_name = 10;</code>
   * @return The packageName.
   */
  @java.lang.Override
  public java.lang.String getPackageName() {
    java.lang.Object ref = packageName_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      packageName_ = s;
      return s;
    }
  }
  /**
   * <code>string package_name = 10;</code>
   * @return The bytes for packageName.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getPackageNameBytes() {
    java.lang.Object ref = packageName_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      packageName_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int IDENTIFIERS_FIELD_NUMBER = 13;
  private java.util.List<org.mojolang.mojo.lang.Identifier> identifiers_;
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  @java.lang.Override
  public java.util.List<org.mojolang.mojo.lang.Identifier> getIdentifiersList() {
    return identifiers_;
  }
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  @java.lang.Override
  public java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
      getIdentifiersOrBuilderList() {
    return identifiers_;
  }
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  @java.lang.Override
  public int getIdentifiersCount() {
    return identifiers_.size();
  }
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.Identifier getIdentifiers(int index) {
    return identifiers_.get(index);
  }
  /**
   * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.lang.IdentifierOrBuilder getIdentifiersOrBuilder(
      int index) {
    return identifiers_.get(index);
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
    if (!getPackageNameBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 10, packageName_);
    }
    for (int i = 0; i < identifiers_.size(); i++) {
      output.writeMessage(13, identifiers_.get(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getPackageNameBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(10, packageName_);
    }
    for (int i = 0; i < identifiers_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(13, identifiers_.get(i));
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
    if (!(obj instanceof org.mojolang.mojo.lang.Import)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.lang.Import other = (org.mojolang.mojo.lang.Import) obj;

    if (!getPackageName()
        .equals(other.getPackageName())) return false;
    if (!getIdentifiersList()
        .equals(other.getIdentifiersList())) return false;
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
    hash = (37 * hash) + PACKAGE_NAME_FIELD_NUMBER;
    hash = (53 * hash) + getPackageName().hashCode();
    if (getIdentifiersCount() > 0) {
      hash = (37 * hash) + IDENTIFIERS_FIELD_NUMBER;
      hash = (53 * hash) + getIdentifiersList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.lang.Import parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.Import parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.Import parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.Import parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.lang.Import prototype) {
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
   * Protobuf type {@code mojo.lang.Import}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.lang.Import)
      org.mojolang.mojo.lang.ImportOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.lang.ImportProto.internal_static_mojo_lang_Import_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.lang.ImportProto.internal_static_mojo_lang_Import_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.lang.Import.class, org.mojolang.mojo.lang.Import.Builder.class);
    }

    // Construct using org.mojolang.mojo.lang.Import.newBuilder()
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
        getIdentifiersFieldBuilder();
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      packageName_ = "";

      if (identifiersBuilder_ == null) {
        identifiers_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
      } else {
        identifiersBuilder_.clear();
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.lang.ImportProto.internal_static_mojo_lang_Import_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.Import getDefaultInstanceForType() {
      return org.mojolang.mojo.lang.Import.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.Import build() {
      org.mojolang.mojo.lang.Import result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.Import buildPartial() {
      org.mojolang.mojo.lang.Import result = new org.mojolang.mojo.lang.Import(this);
      int from_bitField0_ = bitField0_;
      result.packageName_ = packageName_;
      if (identifiersBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          identifiers_ = java.util.Collections.unmodifiableList(identifiers_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.identifiers_ = identifiers_;
      } else {
        result.identifiers_ = identifiersBuilder_.build();
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
      if (other instanceof org.mojolang.mojo.lang.Import) {
        return mergeFrom((org.mojolang.mojo.lang.Import)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.lang.Import other) {
      if (other == org.mojolang.mojo.lang.Import.getDefaultInstance()) return this;
      if (!other.getPackageName().isEmpty()) {
        packageName_ = other.packageName_;
        onChanged();
      }
      if (identifiersBuilder_ == null) {
        if (!other.identifiers_.isEmpty()) {
          if (identifiers_.isEmpty()) {
            identifiers_ = other.identifiers_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureIdentifiersIsMutable();
            identifiers_.addAll(other.identifiers_);
          }
          onChanged();
        }
      } else {
        if (!other.identifiers_.isEmpty()) {
          if (identifiersBuilder_.isEmpty()) {
            identifiersBuilder_.dispose();
            identifiersBuilder_ = null;
            identifiers_ = other.identifiers_;
            bitField0_ = (bitField0_ & ~0x00000001);
            identifiersBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getIdentifiersFieldBuilder() : null;
          } else {
            identifiersBuilder_.addAllMessages(other.identifiers_);
          }
        }
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
      org.mojolang.mojo.lang.Import parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.lang.Import) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object packageName_ = "";
    /**
     * <code>string package_name = 10;</code>
     * @return The packageName.
     */
    public java.lang.String getPackageName() {
      java.lang.Object ref = packageName_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        packageName_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string package_name = 10;</code>
     * @return The bytes for packageName.
     */
    public com.google.protobuf.ByteString
        getPackageNameBytes() {
      java.lang.Object ref = packageName_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        packageName_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string package_name = 10;</code>
     * @param value The packageName to set.
     * @return This builder for chaining.
     */
    public Builder setPackageName(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      packageName_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string package_name = 10;</code>
     * @return This builder for chaining.
     */
    public Builder clearPackageName() {
      
      packageName_ = getDefaultInstance().getPackageName();
      onChanged();
      return this;
    }
    /**
     * <code>string package_name = 10;</code>
     * @param value The bytes for packageName to set.
     * @return This builder for chaining.
     */
    public Builder setPackageNameBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      packageName_ = value;
      onChanged();
      return this;
    }

    private java.util.List<org.mojolang.mojo.lang.Identifier> identifiers_ =
      java.util.Collections.emptyList();
    private void ensureIdentifiersIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        identifiers_ = new java.util.ArrayList<org.mojolang.mojo.lang.Identifier>(identifiers_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.lang.Identifier, org.mojolang.mojo.lang.Identifier.Builder, org.mojolang.mojo.lang.IdentifierOrBuilder> identifiersBuilder_;

    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public java.util.List<org.mojolang.mojo.lang.Identifier> getIdentifiersList() {
      if (identifiersBuilder_ == null) {
        return java.util.Collections.unmodifiableList(identifiers_);
      } else {
        return identifiersBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public int getIdentifiersCount() {
      if (identifiersBuilder_ == null) {
        return identifiers_.size();
      } else {
        return identifiersBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public org.mojolang.mojo.lang.Identifier getIdentifiers(int index) {
      if (identifiersBuilder_ == null) {
        return identifiers_.get(index);
      } else {
        return identifiersBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder setIdentifiers(
        int index, org.mojolang.mojo.lang.Identifier value) {
      if (identifiersBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureIdentifiersIsMutable();
        identifiers_.set(index, value);
        onChanged();
      } else {
        identifiersBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder setIdentifiers(
        int index, org.mojolang.mojo.lang.Identifier.Builder builderForValue) {
      if (identifiersBuilder_ == null) {
        ensureIdentifiersIsMutable();
        identifiers_.set(index, builderForValue.build());
        onChanged();
      } else {
        identifiersBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder addIdentifiers(org.mojolang.mojo.lang.Identifier value) {
      if (identifiersBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureIdentifiersIsMutable();
        identifiers_.add(value);
        onChanged();
      } else {
        identifiersBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder addIdentifiers(
        int index, org.mojolang.mojo.lang.Identifier value) {
      if (identifiersBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureIdentifiersIsMutable();
        identifiers_.add(index, value);
        onChanged();
      } else {
        identifiersBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder addIdentifiers(
        org.mojolang.mojo.lang.Identifier.Builder builderForValue) {
      if (identifiersBuilder_ == null) {
        ensureIdentifiersIsMutable();
        identifiers_.add(builderForValue.build());
        onChanged();
      } else {
        identifiersBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder addIdentifiers(
        int index, org.mojolang.mojo.lang.Identifier.Builder builderForValue) {
      if (identifiersBuilder_ == null) {
        ensureIdentifiersIsMutable();
        identifiers_.add(index, builderForValue.build());
        onChanged();
      } else {
        identifiersBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder addAllIdentifiers(
        java.lang.Iterable<? extends org.mojolang.mojo.lang.Identifier> values) {
      if (identifiersBuilder_ == null) {
        ensureIdentifiersIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, identifiers_);
        onChanged();
      } else {
        identifiersBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder clearIdentifiers() {
      if (identifiersBuilder_ == null) {
        identifiers_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        identifiersBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public Builder removeIdentifiers(int index) {
      if (identifiersBuilder_ == null) {
        ensureIdentifiersIsMutable();
        identifiers_.remove(index);
        onChanged();
      } else {
        identifiersBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public org.mojolang.mojo.lang.Identifier.Builder getIdentifiersBuilder(
        int index) {
      return getIdentifiersFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public org.mojolang.mojo.lang.IdentifierOrBuilder getIdentifiersOrBuilder(
        int index) {
      if (identifiersBuilder_ == null) {
        return identifiers_.get(index);  } else {
        return identifiersBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public java.util.List<? extends org.mojolang.mojo.lang.IdentifierOrBuilder> 
         getIdentifiersOrBuilderList() {
      if (identifiersBuilder_ != null) {
        return identifiersBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(identifiers_);
      }
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public org.mojolang.mojo.lang.Identifier.Builder addIdentifiersBuilder() {
      return getIdentifiersFieldBuilder().addBuilder(
          org.mojolang.mojo.lang.Identifier.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public org.mojolang.mojo.lang.Identifier.Builder addIdentifiersBuilder(
        int index) {
      return getIdentifiersFieldBuilder().addBuilder(
          index, org.mojolang.mojo.lang.Identifier.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.lang.Identifier identifiers = 13;</code>
     */
    public java.util.List<org.mojolang.mojo.lang.Identifier.Builder> 
         getIdentifiersBuilderList() {
      return getIdentifiersFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.lang.Identifier, org.mojolang.mojo.lang.Identifier.Builder, org.mojolang.mojo.lang.IdentifierOrBuilder> 
        getIdentifiersFieldBuilder() {
      if (identifiersBuilder_ == null) {
        identifiersBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            org.mojolang.mojo.lang.Identifier, org.mojolang.mojo.lang.Identifier.Builder, org.mojolang.mojo.lang.IdentifierOrBuilder>(
                identifiers_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        identifiers_ = null;
      }
      return identifiersBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.lang.Import)
  }

  // @@protoc_insertion_point(class_scope:mojo.lang.Import)
  private static final org.mojolang.mojo.lang.Import DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.lang.Import();
  }

  public static org.mojolang.mojo.lang.Import getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Import>
      PARSER = new com.google.protobuf.AbstractParser<Import>() {
    @java.lang.Override
    public Import parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Import(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Import> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Import> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.lang.Import getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
