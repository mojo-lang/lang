#ifndef MOJO_LANG_CALLABLE_TYPE_HPP
#define MOJO_LANG_CALLABLE_TYPE_HPP

namespace mojo {
namespace lang {

class Callable {
public:
    virtual ~Callable() = default;

    virtual Future<Any> call(AnyRef req) const = 0;
};

}
}

#endif //MOJO_LANG_CALLABLE_TYPE_HPP
