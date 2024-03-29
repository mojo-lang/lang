// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/cached_document.proto

package org.mojolang.mojo.lang;

/**
 * Protobuf type {@code mojo.lang.CachedDocument}
 */
public final class CachedDocument extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.lang.CachedDocument)
    CachedDocumentOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CachedDocument.newBuilder() to construct.
  private CachedDocument(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CachedDocument() {
    cache_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new CachedDocument();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.lang.CachedDocumentProto.internal_static_mojo_lang_CachedDocument_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.lang.CachedDocumentProto.internal_static_mojo_lang_CachedDocument_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.lang.CachedDocument.class, org.mojolang.mojo.lang.CachedDocument.Builder.class);
  }

  public static final int CACHE_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private volatile java.lang.Object cache_ = "";
  /**
   * <code>string cache = 1;</code>
   * @return The cache.
   */
  @java.lang.Override
  public java.lang.String getCache() {
    java.lang.Object ref = cache_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      cache_ = s;
      return s;
    }
  }
  /**
   * <code>string cache = 1;</code>
   * @return The bytes for cache.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getCacheBytes() {
    java.lang.Object ref = cache_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      cache_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int VAL_FIELD_NUMBER = 2;
  private org.mojolang.mojo.document.Document val_;
  /**
   * <code>.mojo.document.Document val = 2;</code>
   * @return Whether the val field is set.
   */
  @java.lang.Override
  public boolean hasVal() {
    return val_ != null;
  }
  /**
   * <code>.mojo.document.Document val = 2;</code>
   * @return The val.
   */
  @java.lang.Override
  public org.mojolang.mojo.document.Document getVal() {
    return val_ == null ? org.mojolang.mojo.document.Document.getDefaultInstance() : val_;
  }
  /**
   * <code>.mojo.document.Document val = 2;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.document.DocumentOrBuilder getValOrBuilder() {
    return val_ == null ? org.mojolang.mojo.document.Document.getDefaultInstance() : val_;
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
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(cache_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, cache_);
    }
    if (val_ != null) {
      output.writeMessage(2, getVal());
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(cache_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, cache_);
    }
    if (val_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getVal());
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
    if (!(obj instanceof org.mojolang.mojo.lang.CachedDocument)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.lang.CachedDocument other = (org.mojolang.mojo.lang.CachedDocument) obj;

    if (!getCache()
        .equals(other.getCache())) return false;
    if (hasVal() != other.hasVal()) return false;
    if (hasVal()) {
      if (!getVal()
          .equals(other.getVal())) return false;
    }
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
    hash = (37 * hash) + CACHE_FIELD_NUMBER;
    hash = (53 * hash) + getCache().hashCode();
    if (hasVal()) {
      hash = (37 * hash) + VAL_FIELD_NUMBER;
      hash = (53 * hash) + getVal().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.CachedDocument parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.lang.CachedDocument prototype) {
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
   * Protobuf type {@code mojo.lang.CachedDocument}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.lang.CachedDocument)
      org.mojolang.mojo.lang.CachedDocumentOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.lang.CachedDocumentProto.internal_static_mojo_lang_CachedDocument_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.lang.CachedDocumentProto.internal_static_mojo_lang_CachedDocument_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.lang.CachedDocument.class, org.mojolang.mojo.lang.CachedDocument.Builder.class);
    }

    // Construct using org.mojolang.mojo.lang.CachedDocument.newBuilder()
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
      cache_ = "";
      val_ = null;
      if (valBuilder_ != null) {
        valBuilder_.dispose();
        valBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.lang.CachedDocumentProto.internal_static_mojo_lang_CachedDocument_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.CachedDocument getDefaultInstanceForType() {
      return org.mojolang.mojo.lang.CachedDocument.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.CachedDocument build() {
      org.mojolang.mojo.lang.CachedDocument result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.CachedDocument buildPartial() {
      org.mojolang.mojo.lang.CachedDocument result = new org.mojolang.mojo.lang.CachedDocument(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(org.mojolang.mojo.lang.CachedDocument result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.cache_ = cache_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.val_ = valBuilder_ == null
            ? val_
            : valBuilder_.build();
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
      if (other instanceof org.mojolang.mojo.lang.CachedDocument) {
        return mergeFrom((org.mojolang.mojo.lang.CachedDocument)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.lang.CachedDocument other) {
      if (other == org.mojolang.mojo.lang.CachedDocument.getDefaultInstance()) return this;
      if (!other.getCache().isEmpty()) {
        cache_ = other.cache_;
        bitField0_ |= 0x00000001;
        onChanged();
      }
      if (other.hasVal()) {
        mergeVal(other.getVal());
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
              cache_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000001;
              break;
            } // case 10
            case 18: {
              input.readMessage(
                  getValFieldBuilder().getBuilder(),
                  extensionRegistry);
              bitField0_ |= 0x00000002;
              break;
            } // case 18
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

    private java.lang.Object cache_ = "";
    /**
     * <code>string cache = 1;</code>
     * @return The cache.
     */
    public java.lang.String getCache() {
      java.lang.Object ref = cache_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        cache_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string cache = 1;</code>
     * @return The bytes for cache.
     */
    public com.google.protobuf.ByteString
        getCacheBytes() {
      java.lang.Object ref = cache_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        cache_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string cache = 1;</code>
     * @param value The cache to set.
     * @return This builder for chaining.
     */
    public Builder setCache(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      cache_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>string cache = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearCache() {
      cache_ = getDefaultInstance().getCache();
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <code>string cache = 1;</code>
     * @param value The bytes for cache to set.
     * @return This builder for chaining.
     */
    public Builder setCacheBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      cache_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    private org.mojolang.mojo.document.Document val_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.document.Document, org.mojolang.mojo.document.Document.Builder, org.mojolang.mojo.document.DocumentOrBuilder> valBuilder_;
    /**
     * <code>.mojo.document.Document val = 2;</code>
     * @return Whether the val field is set.
     */
    public boolean hasVal() {
      return ((bitField0_ & 0x00000002) != 0);
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     * @return The val.
     */
    public org.mojolang.mojo.document.Document getVal() {
      if (valBuilder_ == null) {
        return val_ == null ? org.mojolang.mojo.document.Document.getDefaultInstance() : val_;
      } else {
        return valBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    public Builder setVal(org.mojolang.mojo.document.Document value) {
      if (valBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        val_ = value;
      } else {
        valBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    public Builder setVal(
        org.mojolang.mojo.document.Document.Builder builderForValue) {
      if (valBuilder_ == null) {
        val_ = builderForValue.build();
      } else {
        valBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    public Builder mergeVal(org.mojolang.mojo.document.Document value) {
      if (valBuilder_ == null) {
        if (((bitField0_ & 0x00000002) != 0) &&
          val_ != null &&
          val_ != org.mojolang.mojo.document.Document.getDefaultInstance()) {
          getValBuilder().mergeFrom(value);
        } else {
          val_ = value;
        }
      } else {
        valBuilder_.mergeFrom(value);
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    public Builder clearVal() {
      bitField0_ = (bitField0_ & ~0x00000002);
      val_ = null;
      if (valBuilder_ != null) {
        valBuilder_.dispose();
        valBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    public org.mojolang.mojo.document.Document.Builder getValBuilder() {
      bitField0_ |= 0x00000002;
      onChanged();
      return getValFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    public org.mojolang.mojo.document.DocumentOrBuilder getValOrBuilder() {
      if (valBuilder_ != null) {
        return valBuilder_.getMessageOrBuilder();
      } else {
        return val_ == null ?
            org.mojolang.mojo.document.Document.getDefaultInstance() : val_;
      }
    }
    /**
     * <code>.mojo.document.Document val = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.document.Document, org.mojolang.mojo.document.Document.Builder, org.mojolang.mojo.document.DocumentOrBuilder> 
        getValFieldBuilder() {
      if (valBuilder_ == null) {
        valBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.document.Document, org.mojolang.mojo.document.Document.Builder, org.mojolang.mojo.document.DocumentOrBuilder>(
                getVal(),
                getParentForChildren(),
                isClean());
        val_ = null;
      }
      return valBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.lang.CachedDocument)
  }

  // @@protoc_insertion_point(class_scope:mojo.lang.CachedDocument)
  private static final org.mojolang.mojo.lang.CachedDocument DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.lang.CachedDocument();
  }

  public static org.mojolang.mojo.lang.CachedDocument getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CachedDocument>
      PARSER = new com.google.protobuf.AbstractParser<CachedDocument>() {
    @java.lang.Override
    public CachedDocument parsePartialFrom(
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

  public static com.google.protobuf.Parser<CachedDocument> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CachedDocument> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.lang.CachedDocument getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

