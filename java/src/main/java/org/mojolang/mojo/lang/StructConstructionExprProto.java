// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/struct_construction_expr.proto

package org.mojolang.mojo.lang;

public final class StructConstructionExprProto {
  private StructConstructionExprProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_StructConstructionExpr_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_StructConstructionExpr_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n(mojo/lang/struct_construction_expr.pro" +
      "to\022\tmojo.lang\032\024mojo/lang/expr.proto\032\030moj" +
      "o/lang/position.proto\"\220\001\n\026StructConstruc" +
      "tionExpr\022+\n\016start_position\030\001 \001(\0132\023.mojo." +
      "lang.Position\022)\n\014end_position\030\002 \001(\0132\023.mo" +
      "jo.lang.Position\022\014\n\004kind\030\003 \001(\005\022\020\n\010implic" +
      "it\030\004 \001(\010Bh\n\026org.mojolang.mojo.langB\033Stru" +
      "ctConstructionExprProtoP\001Z/github.com/mo" +
      "jo-lang/lang/go/pkg/mojo/lang;langb\006prot" +
      "o3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_StructConstructionExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_StructConstructionExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_StructConstructionExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "Kind", "Implicit", });
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}