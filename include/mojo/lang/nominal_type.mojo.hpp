#ifndef LANG_NOMINAL_TYPE_MOJO_HPP
#define LANG_NOMINAL_TYPE_MOJO_HPP

#include <mojo/core/string.hpp>
#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/lang/type_declaration.mojo.hpp>
#include <mojo/lang/nominal_type_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct NominalType {
    /**
     *
     */
    String name;

    /**
     *
     */
    TypeDeclaration typeDeclaration;

    /**
     *
     */
    Array<NominalTypePtr> genericArguments;

    /**
     *
     */
    Array<Attribute> attributes;

    NominalType() : typeDeclaration(DataDeclPtr{}) {}
};

}
}

#endif //LANG_NOMINAL_TYPE_MOJO_HPP