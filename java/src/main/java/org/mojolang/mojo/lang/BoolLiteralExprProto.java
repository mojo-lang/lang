// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/bool_literal_expr.proto

package org.mojolang.mojo.lang;

public final class BoolLiteralExprProto {
  private BoolLiteralExprProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_BoolLiteralExpr_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_BoolLiteralExpr_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n!mojo/lang/bool_literal_expr.proto\022\tmoj" +
      "o.lang\032\024mojo/lang/expr.proto\032\034mojo/lang/" +
      "literal_expr.proto\032\030mojo/lang/position.p" +
      "roto\"\230\001\n\017BoolLiteralExpr\022+\n\016start_positi" +
      "on\030\001 \001(\0132\023.mojo.lang.Position\022)\n\014end_pos" +
      "ition\030\002 \001(\0132\023.mojo.lang.Position\022\014\n\004kind" +
      "\030\003 \001(\005\022\020\n\010implicit\030\004 \001(\010\022\r\n\005value\030\020 \001(\010B" +
      "a\n\026org.mojolang.mojo.langB\024BoolLiteralEx" +
      "prProtoP\001Z/github.com/mojo-lang/lang/go/" +
      "pkg/mojo/lang;langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.ExprProto.getDescriptor(),
          org.mojolang.mojo.lang.LiteralExprProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_BoolLiteralExpr_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_BoolLiteralExpr_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_BoolLiteralExpr_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "Kind", "Implicit", "Value", });
    org.mojolang.mojo.lang.ExprProto.getDescriptor();
    org.mojolang.mojo.lang.LiteralExprProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
