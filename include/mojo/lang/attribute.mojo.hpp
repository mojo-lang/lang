#ifndef MOJO_LANG_ATTRIBUTE_MOJO_HPP
#define MOJO_LANG_ATTRIBUTE_MOJO_HPP

#include <mojo/lang/package.mojo.hpp>
#include <mojo/lang/expression.mojo.hpp>

namespace mojo {
namespace lang {

struct Attribute {
    String name;
    String fullName;

    Expression expression;
    PackagePtr package;
};

}
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, Attribute) {
    field(&mojo::lang::Attribute::name, "name", 1);
    field(&mojo::lang::Attribute::fullName, "full_name", "fullName", 2);
    field(&mojo::lang::Attribute::expression, "expression", 3);
    field(&mojo::lang::Attribute::package, "package", 4);
}

#endif //MOJO_LANG_ATTRIBUTE_MOJO_HPP
