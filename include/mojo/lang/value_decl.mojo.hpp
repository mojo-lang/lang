#ifndef LANG_VALUE_DECL_MOJO_HPP
#define LANG_VALUE_DECL_MOJO_HPP

#include <mojo/core/string.hpp>
#include <mojo/core/array.hpp>
#include <mojo/lang/decl.mojo.hpp>
#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/lang/expression.mojo.hpp>

namespace mojo {
namespace lang {

struct ValueDecl : Decl {
    String name;

    //type: NominalType @12 @link
    Array<Attribute> attributes;
    Expression value;
};

}
}

#endif //LANG_VALUE_DECL_MOJO_HPP
