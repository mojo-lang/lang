// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/position.proto

package org.mojolang.mojo.lang;

public final class PositionProto {
  private PositionProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_Position_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_Position_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_PositionSpan_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_PositionSpan_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_PosSpan_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_PosSpan_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\030mojo/lang/position.proto\022\tmojo.lang\"J\n" +
      "\010Position\022\020\n\010filename\030\001 \001(\t\022\016\n\006offset\030\002 " +
      "\001(\003\022\014\n\004line\030\003 \001(\003\022\016\n\006column\030\004 \001(\003\"3\n\014Pos" +
      "itionSpan\022#\n\006values\030\001 \003(\0132\023.mojo.lang.Po" +
      "sition\"\031\n\007PosSpan\022\016\n\006values\030\001 \003(\003BZ\n\026org" +
      ".mojolang.mojo.langB\rPositionProtoP\001Z/gi" +
      "thub.com/mojo-lang/lang/go/pkg/mojo/lang" +
      ";langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_lang_Position_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_Position_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_Position_descriptor,
        new java.lang.String[] { "Filename", "Offset", "Line", "Column", });
    internal_static_mojo_lang_PositionSpan_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_mojo_lang_PositionSpan_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_PositionSpan_descriptor,
        new java.lang.String[] { "Values", });
    internal_static_mojo_lang_PosSpan_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_mojo_lang_PosSpan_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_PosSpan_descriptor,
        new java.lang.String[] { "Values", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}