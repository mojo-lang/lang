// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/struct_literal_expr.proto

package org.mojolang.mojo.lang;

public final class StructLiteralExprProto {
  private StructLiteralExprProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_StructLiteralExpr_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_StructLiteralExpr_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n#mojo/lang/struct_literal_expr.proto\022\tm" +
      "ojo.lang\032\032mojo/lang/apply_expr.proto\032\024mo" +
      "jo/lang/expr.proto\032\030mojo/lang/position.p" +
      "roto\"\213\001\n\021StructLiteralExpr\022+\n\016start_posi" +
      "tion\030\001 \001(\0132\023.mojo.lang.Position\022)\n\014end_p" +
      "osition\030\002 \001(\0132\023.mojo.lang.Position\022\014\n\004ki" +
      "nd\030\003 \001(\005\022\020\n\010implicit\030\004 \001(\010Bc\n\026org.mojola" +
      "ng.mojo.langB\026StructLiteralExprProtoP\001Z/" +
      "github.com/mojo-lang/lang/go/pkg/mojo/la" +
      "ng;langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.ApplyExprProto.getDescriptor(),
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_StructLiteralExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_StructLiteralExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_StructLiteralExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "Kind", "Implicit", });
    org.mojolang.mojo.lang.ApplyExprProto.getDescriptor();
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}