// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/decl.proto

package org.mojolang.mojo.lang;

public final class DeclProto {
  private DeclProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_Decl_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_Decl_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\024mojo/lang/decl.proto\022\tmojo.lang\032\027mojo/" +
      "lang/comment.proto\032\030mojo/lang/document.p" +
      "roto\032\030mojo/lang/position.proto\"\365\001\n\004Decl\022" +
      "+\n\016start_position\030\001 \001(\0132\023.mojo.lang.Posi" +
      "tion\022)\n\014end_position\030\002 \001(\0132\023.mojo.lang.P" +
      "osition\022,\n\020leading_comments\030\003 \003(\0132\022.mojo" +
      ".lang.Comment\022\020\n\010implicit\030\005 \001(\010\022%\n\010docum" +
      "ent\030\006 \001(\0132\023.mojo.lang.Document\022\024\n\014packag" +
      "e_name\030\007 \001(\t\022\030\n\020source_file_name\030\010 \001(\tBV" +
      "\n\026org.mojolang.mojo.langB\tDeclProtoP\001Z/g" +
      "ithub.com/mojo-lang/lang/go/pkg/mojo/lan" +
      "g;langb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.CommentProto.getDescriptor(),
          org.mojolang.mojo.lang.DocumentProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
        });
    internal_static_mojo_lang_Decl_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_Decl_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_Decl_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "LeadingComments", "Implicit", "Document", "PackageName", "SourceFileName", });
    org.mojolang.mojo.lang.CommentProto.getDescriptor();
    org.mojolang.mojo.lang.DocumentProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
