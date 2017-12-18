#ifndef MOJO_LANG_SERVICE_TYPE_BUILDER_HPP
#define MOJO_LANG_SERVICE_TYPE_BUILDER_HPP

#include <mojo/lang/service_type.hpp>

namespace mojo {
namespace lang {

template <typename T>
struct ServiceTypeBuilder {
    ServiceTypeOf<T>* type;

    void name(std::string&& name, std::string&& full_name) {
        // msg->name_ = std::move(name);
    }

    template <typename Req, typename Resp>
    ServiceMethodOf<T, Req, Resp>& method(Resp T::*func(Req), std::string name, int number) {
        auto p = new MethodFor<T, Req, Resp>{func, std::move(name)};
        // msg->fields_.append(std::unique_ptr<FieldOf<T>>(p));
        return *p;
    }

    ServiceTypeOf<T>* operator()() {
        return type;
    }
};

}
}

#define MOJO_BUILD_SERVICE_TYPE(TYPE)                                                     \
    struct ServiceTypeBuilderFor##TYPE : ::ncraft::proto::ServiceTypeBuilder<TYPE> { \
        void build();                                                                \
        void build_with_defaults() {                                                 \
            name(#TYPE);                                                             \
            build();                                                                 \
        }                                                                            \
    };                                                                               \
    inline const ::ncraft::proto::ServiceType<TYPE>* buildType(                      \
        const ::ncraft::core::TypeIdentifier<TYPE>* dummy) {                         \
        ServiceTypeBuilderFor##TYPE builder;                                         \
        builder.type = new ::ncraft::proto::ServiceTypeOf<TYPE>;                    \
        builder.build_with_defaults();                                               \
        return builder();                                                            \
    }                                                                                \
    inline void ServiceTypeBuilderFor##TYPE::build()

#endif //MOJO_LANG_SERVICE_TYPE_BUILDER_HPP
