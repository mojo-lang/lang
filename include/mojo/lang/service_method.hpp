#ifndef MOJO_LANG_SERVICE_METHOD_HPP
#define MOJO_LANG_SERVICE_METHOD_HPP

#include <mojo/lang/function.hpp>
#include <mojo/lang/service_type.hpp>

namespace mojo {
namespace lang {

template <typename T>
class ServiceMethod : public Function {
public:
    virtual ~ServiceMethod() {
    }

    virtual const ServiceType* service() const noexcept = 0;

    virtual void call(T&, ObjectVisitor& visitor) const = 0;
};

template <typename T, typename M>
class ServiceMethodOf : public ServiceMethod<T> {
public:
    using MemberPtr = M T::*;
    MemberPtr ptr_;

    ServiceMethodOf(MemberPtr ptr, std::string name) : ptr_(ptr) {
    }

    const std::string& name() const final {
        return name_;
    }
    const std::string& lowerCamelName() const final {
        return name_;
    }

    //void visit(T& record, ObjectVisitor& visitor) const override {
        // visitor[this->column()](get_known(record));
    //}
};

}
}

#endif //MOJO_LANG_SERVICE_METHOD_HPP
