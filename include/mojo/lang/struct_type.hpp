#ifndef MOJO_LANG_STRUCT_TYPE_HPP
#define MOJO_LANG_STRUCT_TYPE_HPP

#include <mojo/core/string.hpp>
#include <mojo/core/numeric.hpp>
#include <mojo/lang/data_type.hpp>
#include <mojo/lang/enum_type.hpp>

namespace mojo {
namespace lang {

struct StructTypeField {
    UInt id = 0;
    String name;
    String camelName;

    String document;
    Attributes attributes;

    const DataType* type;
};

class StructType : public DataType {
public:
    using Field = StructTypeField;

public:
    const StructType* enclosingType;
    Array<const StructType*> nestedTypes;  // nested types
    Array<const EnumType*> nestedEnums;    // nested enums

    Array<Field*> fields;
/*
    class ConstFieldIterator
        : public std::iterator<std::forward_iterator_tag, const Field> {
    public:
        explicit ConstFieldIterator(const ObjectType& objectType)
            : objectType_(objectType) {
        }

        static ConstFieldIterator end(const ObjectType& objectType) {
            ConstFieldIterator itr(objectType);
            itr.index_ = -1;
            return itr;
        }

    public:
        void operator++() {
            ++index_;
        }

        const Field& operator*() const {
            return *objectType_.field(index_);
        }

        const Field& operator->() const {
            return *objectType_.field(index_);
        }

        bool operator!=(const ConstFieldIterator& other) const {
            return !((*this) == other);
        }
        bool operator==(const ConstFieldIterator& other) const {
            return &objectType_ == &(other.objectType_) && index_ == other.index_;
        }

    private:
        size_t index_{0};
        const ObjectType& objectType_;
    };

    class ConstFields : Copyable {
    public:
        ConstFields(const ObjectType& objectType) : objectType_(objectType) {
        }

        ConstFieldIterator begin() const {
            return objectType_.fieldBegin();
        }
        ConstFieldIterator end() const {
            return objectType_.fieldEnd();
        }

    private:
        const ObjectType& objectType_;
    };*/

    /**
     * The number of fields that belong to the ObjectType.
     * @return
     */
    virtual size_t fieldCount() const = 0;

    /**
     *
     * @param name
     * @return
     */
    virtual const Field* field(const String& name) const = 0;

    /**
     *
     * @param number
     * @return
     */
    virtual const Field* field(size_t number) const = 0;

    /**
     *
     * @return
     */
    //ConstFieldIterator fieldBegin() const {
    //    return ConstFieldIterator(*this);
    //}

    /**
     *
     * @return
     */
    //ConstFieldIterator fieldEnd() const {
    //    return ConstFieldIterator::end(*this);
    //}

    /**
     *
     * @return
     */
    //ConstFields fields() const {
    //    return ConstFields(*this);
    //}
};
}
}

#endif  // MOJO_LANG_STRUCT_TYPE_HPP
