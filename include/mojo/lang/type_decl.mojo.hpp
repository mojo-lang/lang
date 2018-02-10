#ifndef MOJO_LANG_TYPE_DECL_MOJO_HPP
#define MOJO_LANG_TYPE_DECL_MOJO_HPP

#include <mojo/lang/struct_type_builder.hpp>

namespace mojo {
namespace lang {

struct TypeDecl : Decl {
    String name;
    Array<GenericParam> genericParams;

    DataType* type;
    Array<Attribute> attributes;
};

MOJO_BUILD_STRUCT_TYPE(mojo, lang, TypeDecl) {
    field(&mojo::lang::DataType::name, "name", 1);
    field(&mojo::lang::DataType::fullName, "full_name", "fullName", 2);
    field(&mojo::lang::DataType::package, "package", 3);
}

MOJO_BUILD_STRUCT_DECL() {
    name();
    document();
    genericParam();

    type<>();
    attribute();
}

MOJO_BUILD_SOURCE_FILE() {
}

}
}

#endif //MOJO_LANG_TYPE_DECL_MOJO_HPP
