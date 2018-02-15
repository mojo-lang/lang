#ifndef LANG_FUNC_DECL_MOJO_HPP
#define LANG_FUNC_DECL_MOJO_HPP

#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/lang/decl.mojo.hpp>
#include <mojo/lang/func_decl_ptr.hpp>
#include <mojo/lang/func_type.mojo.hpp>
#include <mojo/lang/generic_parameter.mojo.hpp>

namespace mojo {
namespace lang {

struct FuncDecl : Decl {
public:
    /**
     *
     */
    String name;

    /**
     *
     */
    Array<GenericParameterPtr> genericParams;

    /**
     *
     */
    Array<Attribute> attributes;

    /**
     *
     */
    FuncTypePtr type;

    /**
     *
     */
    // body;
};

}  // namespace lang
}  // namespace mojo

#endif  // LANG_FUNC_DECL_MOJO_HPP
