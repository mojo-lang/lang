// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/entity_relation_set.proto

package org.mojolang.mojo.lang;

public final class EntityRelationSetProto {
  private EntityRelationSetProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_EntityRelationSet_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_EntityRelationSet_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_EntityRelationSet_NodesEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_EntityRelationSet_NodesEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_EntityRelationSet_EdgesEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_EntityRelationSet_EdgesEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_EntityRelationSet_RelationsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_EntityRelationSet_RelationsEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_EntityRelationSet_NodeEdgesEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_EntityRelationSet_NodeEdgesEntry_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_EntityRelationSet_NodeRelationsEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_EntityRelationSet_NodeRelationsEntry_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n#mojo/lang/entity_relation_set.proto\022\tm" +
      "ojo.lang\032\024mojo/lang/lang.proto\"\300\005\n\021Entit" +
      "yRelationSet\0226\n\005nodes\030\001 \003(\0132\'.mojo.lang." +
      "EntityRelationSet.NodesEntry\0226\n\005edges\030\002 " +
      "\003(\0132\'.mojo.lang.EntityRelationSet.EdgesE" +
      "ntry\022>\n\trelations\030\005 \003(\0132+.mojo.lang.Enti" +
      "tyRelationSet.RelationsEntry\022?\n\nnode_edg" +
      "es\030\016 \003(\0132+.mojo.lang.EntityRelationSet.N" +
      "odeEdgesEntry\022G\n\016node_relations\030\017 \003(\0132/." +
      "mojo.lang.EntityRelationSet.NodeRelation" +
      "sEntry\032C\n\nNodesEntry\022\013\n\003key\030\001 \001(\t\022$\n\005val" +
      "ue\030\002 \001(\0132\025.mojo.lang.EntityNode:\0028\001\032C\n\nE" +
      "dgesEntry\022\013\n\003key\030\001 \001(\t\022$\n\005value\030\002 \001(\0132\025." +
      "mojo.lang.EntityEdge:\0028\001\032K\n\016RelationsEnt" +
      "ry\022\013\n\003key\030\001 \001(\t\022(\n\005value\030\002 \001(\0132\031.mojo.la" +
      "ng.EntityRelation:\0028\001\032H\n\016NodeEdgesEntry\022" +
      "\013\n\003key\030\001 \001(\t\022%\n\005value\030\002 \001(\0132\026.mojo.lang." +
      "EntityEdges:\0028\001\032P\n\022NodeRelationsEntry\022\013\n" +
      "\003key\030\001 \001(\t\022)\n\005value\030\002 \001(\0132\032.mojo.lang.En" +
      "tityRelations:\0028\001Bc\n\026org.mojolang.mojo.l" +
      "angB\026EntityRelationSetProtoP\001Z/github.co" +
      "m/mojo-lang/lang/go/pkg/mojo/lang;langb\006" +
      "proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.LangProto.getDescriptor(),
        });
    internal_static_mojo_lang_EntityRelationSet_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_EntityRelationSet_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_EntityRelationSet_descriptor,
        new java.lang.String[] { "Nodes", "Edges", "Relations", "NodeEdges", "NodeRelations", });
    internal_static_mojo_lang_EntityRelationSet_NodesEntry_descriptor =
      internal_static_mojo_lang_EntityRelationSet_descriptor.getNestedTypes().get(0);
    internal_static_mojo_lang_EntityRelationSet_NodesEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_EntityRelationSet_NodesEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_mojo_lang_EntityRelationSet_EdgesEntry_descriptor =
      internal_static_mojo_lang_EntityRelationSet_descriptor.getNestedTypes().get(1);
    internal_static_mojo_lang_EntityRelationSet_EdgesEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_EntityRelationSet_EdgesEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_mojo_lang_EntityRelationSet_RelationsEntry_descriptor =
      internal_static_mojo_lang_EntityRelationSet_descriptor.getNestedTypes().get(2);
    internal_static_mojo_lang_EntityRelationSet_RelationsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_EntityRelationSet_RelationsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_mojo_lang_EntityRelationSet_NodeEdgesEntry_descriptor =
      internal_static_mojo_lang_EntityRelationSet_descriptor.getNestedTypes().get(3);
    internal_static_mojo_lang_EntityRelationSet_NodeEdgesEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_EntityRelationSet_NodeEdgesEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    internal_static_mojo_lang_EntityRelationSet_NodeRelationsEntry_descriptor =
      internal_static_mojo_lang_EntityRelationSet_descriptor.getNestedTypes().get(4);
    internal_static_mojo_lang_EntityRelationSet_NodeRelationsEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_EntityRelationSet_NodeRelationsEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    org.mojolang.mojo.lang.LangProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
