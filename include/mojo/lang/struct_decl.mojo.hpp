#ifndef MOJO_LANG_STRUCT_DECL_MOJO_HPP
#define MOJO_LANG_STRUCT_DECL_MOJO_HPP

#include <mojo/lang/struct_type_builder.hpp>

namespace mojo {
namespace lang {

struct StructDecl : Decl {
    String name;
    Array<GenericParam> genericParams;

    Type* type;
    Array<Attribute> attributes;
};

MOJO_BUILD_STRUCT_TYPE(mojo, lang, StructDecl) {
    field(&mojo::lang::StructDecl::name, "name", 1);
    field(&mojo::lang::StructDecl::fullName, "full_name", "fullName", 2);
    field(&mojo::lang::StructDecl::package, "package", 3);
}

}
}

#endif //MOJO_LANG_STRUCT_DECL_MOJO_HPP
