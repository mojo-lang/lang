// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/stmt.proto

package org.mojolang.mojo.lang;

public final class StmtProto {
  private StmtProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_Stmt_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_Stmt_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\024mojo/lang/stmt.proto\022\tmojo.lang\032\027mojo/" +
      "lang/comment.proto\032\030mojo/lang/position.p" +
      "roto\"\254\001\n\004Stmt\022+\n\016start_position\030\001 \001(\0132\023." +
      "mojo.lang.Position\022)\n\014end_position\030\002 \001(\013" +
      "2\023.mojo.lang.Position\022,\n\020leading_comment" +
      "s\030\003 \003(\0132\022.mojo.lang.Comment\022\014\n\004kind\030\004 \001(" +
      "\003\022\020\n\010implicit\030\005 \001(\010BV\n\026org.mojolang.mojo" +
      ".langB\tStmtProtoP\001Z/github.com/mojo-lang" +
      "/lang/go/pkg/mojo/lang;langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.CommentProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_Stmt_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_Stmt_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_Stmt_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "LeadingComments", "Kind", "Implicit", });
    org.mojolang.mojo.lang.CommentProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
