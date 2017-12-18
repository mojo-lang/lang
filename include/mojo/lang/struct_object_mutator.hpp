#ifndef MOJO_LANG_STRUCT_VALUE_SETTER_HPP
#define MOJO_LANG_STRUCT_VALUE_SETTER_HPP

#include <mojo/lang/struct_type.hpp>

namespace mojo {
namespace lang {

/// setter from object (like json) to struct type object
class StructObjectMutator {
public:
    StructObjectMutator(ncraft::AnyRef instance, const StructType* structType)
        : instance_(instance), structType_(structType) {
    }

    ncraft::Status objectStart() {
        if (!stack_.empty()) {
            auto& cursor = stack_.back();
            auto result = cursor.field->get(cursor.object);
            if (result) {
                const StructType* type = dynamic_cast<const StructType*>(cursor.field->type);
                if (type) {
                    stack_.push_back({*result, type, nullptr});
                }
                else { // dictionary type
                }
            }
            else {
            }
        }
        else { // first time into object
            stack_.push_back({instance_, structType_, nullptr});
        }
    }

    void objectEnd() {
        if (!stack_.back().inDictionary) {
            stack_.pop_back();
        }
    }

    void arrayStart() {
        //stack_.back().type_
    }

    void arrayEnd() {
    }

    void field(const nc::StringView& name) {
        const Field* fld = type->field(name);
        if (fld) {
            stack_.back().field = fld;
        }
        else {
        }
    }

    template<typename T>
    void value(T& v) {
        Frame& current = stack_.back();
        current.field->set(current.msg, v);

        mutator.set(v);

        if (stack_.back().inDictionary) {
            mutator.add();
        }
    }

    void error(Error& error) {
    }

private:
    struct Frame {
        bool inDictionary = false;
        AnyRef object;
        const StructType* type{nullptr};
        const StructType::Field* field{nullptr};

        //DataMutator mutator{object, type};
    };

    ncraft::AnyRef instance_;
    const StructType* structType_;

    std::deque<Frame> stack_;
};
}
}

#endif //MOJO_LANG_STRUCT_VALUE_SETTER_HPP
