// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/comment.proto

package org.mojolang.mojo.lang;

public final class CommentProto {
  private CommentProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_Comment_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_Comment_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\027mojo/lang/comment.proto\022\tmojo.lang\"\031\n\007" +
      "Comment\022\016\n\006values\030\001 \003(\tBY\n\026org.mojolang." +
      "mojo.langB\014CommentProtoP\001Z/github.com/mo" +
      "jo-lang/lang/go/pkg/mojo/lang;langb\006prot" +
      "o3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_lang_Comment_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_Comment_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_Comment_descriptor,
        new java.lang.String[] { "Values", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}