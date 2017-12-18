#ifndef MOJO_STRING_ADAPTOR_HPP
#define MOJO_STRING_ADAPTOR_HPP

#include <mojo/lang/data_adaptor_of.hpp>

namespace mojo {
namespace lang {

class StringAdaptor : public DataAdaptorOf<String> {
public:
    StringAdaptor() = default;
    virtual ~StringAdaptor() = default;

    virtual int size(const String& str) const noexcept {
        return str.size();
    }

    virtual bool has() const noexcept = 0;

    virtual AnyConstRef at(const String& str, const String&) const noexcept = 0;

    virtual AnyConstRef at(const String& str, int) const noexcept = 0;

    virtual ncraft::Status setNull(AnyRef) const noexcept {
    }

    virtual ncraft::Status setBool(AnyRef, bool) const noexcept {
    }

    virtual ncraft::Status set(String& str, AnyConstRef) const noexcept {
    }

    virtual MResult mutableAt(String& str, const String&) const noexcept {
    }

    virtual MResult mutableAt(String& str, int) const noexcept {
    }

    virtual MResult mutateField(String&,
                                const String&,
                                const StructTypeField* = nullptr) const noexcept {
    }

    virtual MResult add(String& str) const noexcept {
    }

    virtual void erase(String& str, const String&) const noexcept {
    }

    virtual void erase(String& str, int) const noexcept {
    }

    virtual void clear(String& str) const noexcept {
        str.clear();
    }
};

}
}

#endif //MOJO_STRING_ADAPTOR_HPP
