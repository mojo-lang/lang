#ifndef MOJO_LANG_SERVICE_TYPE_HPP
#define MOJO_LANG_SERVICE_TYPE_HPP

#include <ncraft/core/noncopyable.hpp>
#include <mojo/lang/data_type.hpp>
#include <mojo/lang/function.hpp>
#include <mojo/lang/struct_type_builder.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
class InterfaceType : ncraft::NonCopyable {
public:
    /**
     *
     */
    struct Inherit {
        /**
         *
         */
        const InterfaceType* serviceType = nullptr;

        /**
         *
         */
        Attributes attributes;
    };

public:
    /**
     *
     */
    String name;

    /**
     * the type's fully-qualified name (eg. foo.bar.Object).
     */
    String fullName;

    /**
     *
     */
    String document;

    /**
     *
     */
    Attributes attributes;

    /**
     *
     */
    PackagePtr package;

    /**
     *
     */
    Array<std::unique_ptr<Inherit>> inherits;

    /**
     *
     */
    Array<FunctionPtr> methods;

public:

};

MOJO_BUILD_STRUCT_TYPE(mojo, lang, InterfaceType) {
    type->attributes = {{},{}};
    type->document = "";

    field(&InterfaceType::name, "name", 1);
    field(&InterfaceType::full_name, "full_name", 2);
    field(&InterfaceType::package, "", 3);
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, InterfaceType, Inherit) {
    field(&InterfaceType::Inherit::serviceType, "interface_type", 1);
    field(&InterfaceType::Inherit::attributes, "attributes", 2);
}

}
}

#endif //MOJO_LANG_SERVICE_TYPE_HPP
