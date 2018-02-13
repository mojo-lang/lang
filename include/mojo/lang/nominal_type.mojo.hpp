#ifndef LANG_NOMINAL_TYPE_MOJO_HPP
#define LANG_NOMINAL_TYPE_MOJO_HPP

#include <mojo/core/string.hpp>

namespace mojo {
namespace lang {

struct NominalType {
    ///
    String name;

    ///
    //type_declaration: TypeDeclaration

    ///
    //generic_arguments: [GenericArgument]

    ///
    //attributes: [Attribute]
};

}
}

#endif //LANG_NOMINAL_TYPE_MOJO_HPP