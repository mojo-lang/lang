#ifndef MOJO_LANG_STRUCT_FIELD_HPP
#define MOJO_LANG_STRUCT_FIELD_HPP

#include <ncraft/core/result.hpp>
#include <mojo/lang/struct_type.hpp>
#include <mojo/lang/data_mutator.hpp>
#include <mojo/lang/data_type.hpp>

namespace mojo {
namespace lang {

template <typename T>
class StructField : public StructType::Field {
public:
    virtual ~StructField() = default;

    virtual bool hasValue(const T& msg) const = 0;

    virtual void visit(T&, DataVisitor& visitor) const = 0;

    virtual DataMutator mutator(T&) const = 0;
};

}
}

#endif //MOJO_LANG_STRUCT_FIELD_HPP
