// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/type_casting_expr.proto

package org.mojolang.mojo.lang;

public final class TypeCastingExprProto {
  private TypeCastingExprProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_TypeCastingExpr_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_TypeCastingExpr_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n!mojo/lang/type_casting_expr.proto\022\tmoj" +
      "o.lang\032\027mojo/lang/comment.proto\032\024mojo/la" +
      "ng/expr.proto\032\030mojo/lang/position.proto\"" +
      "\267\001\n\017TypeCastingExpr\022+\n\016start_position\030\001 " +
      "\001(\0132\023.mojo.lang.Position\022)\n\014end_position" +
      "\030\002 \001(\0132\023.mojo.lang.Position\022,\n\020leading_c" +
      "omments\030\003 \003(\0132\022.mojo.lang.Comment\022\014\n\004kin" +
      "d\030\004 \001(\005\022\020\n\010implicit\030\005 \001(\010Ba\n\026org.mojolan" +
      "g.mojo.langB\024TypeCastingExprProtoP\001Z/git" +
      "hub.com/mojo-lang/lang/go/pkg/mojo/lang;" +
      "langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.CommentProto.getDescriptor(),
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_TypeCastingExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_TypeCastingExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_TypeCastingExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "LeadingComments", "Kind", "Implicit", });
    org.mojolang.mojo.lang.CommentProto.getDescriptor();
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
