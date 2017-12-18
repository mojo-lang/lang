#ifndef MOJO_STRUCT_FIELD_VIRTUAL_OF_HPP
#define MOJO_STRUCT_FIELD_VIRTUAL_OF_HPP

#include <mojo/lang/struct_field.hpp>

namespace mojo {
namespace lang {

template <typename T, typename M>
class StructVirtualField : public StructField<T> {
public:
    using MemberGetPtr = M (T::*)();
    using MemberSetPtr = void (T::*)(const M&);

public:
    StructVirtualField(MemberGetPtr getter, MemberSetPtr setter, String&& name, int idx)
        : getter_(getter), setter_(setter) {
        StructType::Field::type = getType<M>();
    }

private:
    MemberGetPtr getter_;
    MemberSetPtr setter_;
};
}
}

#endif  // MOJO_STRUCT_FIELD_VIRTUAL_OF_HPP
