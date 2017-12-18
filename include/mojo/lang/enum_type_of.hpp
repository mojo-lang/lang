#ifndef MOJO_LANG_ENUM_TYPE_OF_HPP
#define MOJO_LANG_ENUM_TYPE_OF_HPP

#include <mojo/lang/enum_type.hpp>
#include <mojo/lang/data_type_of.hpp>

namespace mojo {
namespace lang {

template <typename T>
class EnumTypeOf : public DataTypeOf<T, EnumType> {
public:

};

}
}

#endif //MOJO_LANG_ENUM_TYPE_OF_HPP
