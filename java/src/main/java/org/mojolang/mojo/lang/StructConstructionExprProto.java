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
      "to\022\tmojo.lang\032\027mojo/lang/comment.proto\032\024" +
      "mojo/lang/expr.proto\032\030mojo/lang/position" +
      ".proto\"\276\001\n\026StructConstructionExpr\022+\n\016sta" +
      "rt_position\030\001 \001(\0132\023.mojo.lang.Position\022)" +
      "\n\014end_position\030\002 \001(\0132\023.mojo.lang.Positio" +
      "n\022,\n\020leading_comments\030\003 \003(\0132\022.mojo.lang." +
      "Comment\022\014\n\004kind\030\004 \001(\005\022\020\n\010implicit\030\005 \001(\010B" +
      "h\n\026org.mojolang.mojo.langB\033StructConstru" +
      "ctionExprProtoP\001Z/github.com/mojo-lang/l" +
      "ang/go/pkg/mojo/lang;langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.CommentProto.getDescriptor(),
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_StructConstructionExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_StructConstructionExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_StructConstructionExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "LeadingComments", "Kind", "Implicit", });
    org.mojolang.mojo.lang.CommentProto.getDescriptor();
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
