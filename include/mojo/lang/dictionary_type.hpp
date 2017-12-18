#ifndef MOJO_DICTIONARY_TYPE_HPP
#define MOJO_DICTIONARY_TYPE_HPP

#include <mojo/lang/data_type_of.hpp>
#include <mojo/lang/generic_type.hpp>

namespace mojo {
namespace lang {

template<typename K, typename V>
class DictionaryType : public DataTypeOf<GenericType<>> {
public:

};

template <typename K, typename V> inline
const DictionaryType<K,V>* buildType(const TypeIdentifier<DictionaryType<K,V>>*) {
    static const DictionaryType<K,V>* p = new DictionaryType<K,V>{};
    return p;
}

}
}

#endif //MOJO_DICTIONARY_TYPE_HPP
