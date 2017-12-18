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
class ServiceType : ncraft::NonCopyable {
public:
    /**
     *
     */
    struct Inherit {
        /**
         *
         */
        const ServiceType* serviceType = nullptr;

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

MOJO_BUILD_STRUCT_TYPE(mojo, lang, ServiceType) {
    type->attributes = {{},{}};
    type->document = "";

    field(&ServiceType::name, "name", 1);
    field(&ServiceType::full_name, "full_name", 2);
    field(&ServiceType::package, "", 3);
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, ServiceType, Inherit) {
    field(&ServiceType::Inherit::serviceType, "service_type", 1);
    field(&ServiceType::Inherit::attributes, "attributes", 2);
}

}
}

#endif //MOJO_LANG_SERVICE_TYPE_HPP
