#ifndef MOJO_LANG_ATTRIBUTE_HPP
#define MOJO_LANG_ATTRIBUTE_HPP

#include <mojo/lang/package.hpp>

namespace mojo {
namespace lang {

struct Attribute {
    String name;
    String full_name;

    String document;
    TypePtr type;

    PackagePtr package;
};

}
}

#endif //MOJO_LANG_ATTRIBUTE_HPP
