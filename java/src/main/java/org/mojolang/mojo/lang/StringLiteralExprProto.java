// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/string_literal_expr.proto

package org.mojolang.mojo.lang;

public final class StringLiteralExprProto {
  private StringLiteralExprProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_StringLiteralExpr_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_StringLiteralExpr_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n#mojo/lang/string_literal_expr.proto\022\tm" +
      "ojo.lang\032\024mojo/lang/expr.proto\032\034mojo/lan" +
      "g/literal_expr.proto\032\030mojo/lang/position" +
      ".proto\"\232\001\n\021StringLiteralExpr\022+\n\016start_po" +
      "sition\030\001 \001(\0132\023.mojo.lang.Position\022)\n\014end" +
      "_position\030\002 \001(\0132\023.mojo.lang.Position\022\014\n\004" +
      "kind\030\003 \001(\005\022\020\n\010implicit\030\004 \001(\010\022\r\n\005value\030\017 " +
      "\001(\tBc\n\026org.mojolang.mojo.langB\026StringLit" +
      "eralExprProtoP\001Z/github.com/mojo-lang/la" +
      "ng/go/pkg/mojo/lang;langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.LiteralExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_StringLiteralExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_StringLiteralExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_StringLiteralExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "Kind", "Implicit", "Value", });
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.LiteralExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}