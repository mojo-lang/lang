#ifndef MOJO_LANG_ENUM_TYPE_HPP
#define MOJO_LANG_ENUM_TYPE_HPP

#include <mojo/core/any.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/data_type.hpp>
#include <mojo/lang/string_type.hpp>

namespace mojo {
namespace lang {

class StructType;

class EnumType : public DataType {
public:
    struct Value {
        Int id;
        String name;
        String camelName;

        String document;
        Attributes attributes;

        Any value;
    };

    const StructType* enclosingType = nullptr;
    Dictionary<String, ::std::unique_ptr<Value>> values;
};
}
}

#endif  // MOJO_LANG_ENUM_TYPE_HPP
