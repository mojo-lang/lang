#ifndef MOJO_LANG_ATTRIBUTE_MOJO_HPP
#define MOJO_LANG_ATTRIBUTE_MOJO_HPP

#include <mojo/lang/expression.mojo.hpp>
#include <mojo/lang/package_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct Attribute {
    /**
     *
     */
    String name;

    /**
     *
     */
    Expression expression;

    /**
     *
     */
    PackagePtr package;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_ATTRIBUTE_MOJO_HPP
