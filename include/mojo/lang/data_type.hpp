#ifndef MOJO_LANG_DATA_TYPE_HPP
#define MOJO_LANG_DATA_TYPE_HPP

#include <map>
#include <ncraft/core/noncopyable.hpp>
#include <ncraft/core/type_info.hpp>
#include <mojo/core/any.hpp>
#include <mojo/core/dictionary.hpp>
#include <mojo/lang/term.hpp>

namespace mojo {
namespace lang {

class DataAdaptor;
class DataVisitor;
class Package;

using Attributes = Dictionary<String, TermPtr>;

/**
 *
 */
class DataType : ::ncraft::NonCopyable {
public:
    enum Kind { kNull, kScalar, kArray, kDictionary, kEnum, kStruct };

public:
    /**
     *
     */
    String name;

public:
    virtual ~DataType() = default;

    /**
     *
     * @return
     */
    virtual Kind kind() const noexcept {
        return kNull;
    }

    /**
     *
     * @return
     */
    virtual bool nullable() const noexcept {
        return false;
    }

    /**
     *
     * @return
     */
    virtual const ::ncraft::TypeInfo& typeInfo() const noexcept {
        return ncraft::core::GetTypeInfo<int>::Value;
    }

    /**
     *
     * @param data
     * @param visitor
     */
    virtual void visit(ncraft::AnyConstRef data, DataVisitor& visitor) const noexcept {
        //NOOP
    }

    /**
     *
     * @param data
     * @return
     */
    virtual const DataAdaptor* adaptor() const noexcept {
        return nullptr;
    }
};

template <typename T>
struct TypeIdentifier {
    constexpr TypeIdentifier() {
    }
};

// Implement this:
// const DataType* buildType(const TypeIdentifier<T>* null);

template <typename T>
auto getType() -> decltype(
    buildType(::std::declval<const TypeIdentifier<ncraft::traits::RemoveCVRef<T>>*>())) {
    static const TypeIdentifier<ncraft::traits::RemoveCVRef<T>> typeIdentifier;
    static const auto t = buildType(&typeIdentifier);
    return t;
}
}
}

#endif  // MOJO_LANG_DATA_TYPE_HPP
