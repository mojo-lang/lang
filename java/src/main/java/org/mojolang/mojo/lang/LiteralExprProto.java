// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/literal_expr.proto

package org.mojolang.mojo.lang;

public final class LiteralExprProto {
  private LiteralExprProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_LiteralExpr_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_LiteralExpr_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\034mojo/lang/literal_expr.proto\022\tmojo.lan" +
      "g\032\024mojo/lang/expr.proto\032\030mojo/lang/posit" +
      "ion.proto\"\205\001\n\013LiteralExpr\022+\n\016start_posit" +
      "ion\030\001 \001(\0132\023.mojo.lang.Position\022)\n\014end_po" +
      "sition\030\002 \001(\0132\023.mojo.lang.Position\022\014\n\004kin" +
      "d\030\003 \001(\005\022\020\n\010implicit\030\004 \001(\010B]\n\026org.mojolan" +
      "g.mojo.langB\020LiteralExprProtoP\001Z/github." +
      "com/mojo-lang/lang/go/pkg/mojo/lang;lang" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_LiteralExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_LiteralExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_LiteralExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "Kind", "Implicit", });
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
