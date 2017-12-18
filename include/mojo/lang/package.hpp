#ifndef MOJO_LANG_PACKAGE_HPP
#define MOJO_LANG_PACKAGE_HPP

#include <vector>
#include <memory>
#include <mojo/core/string.hpp>
#include <mojo/core/array.hpp>
#include <mojo/lang/struct_type_builder.hpp>

namespace mojo {
namespace lang {

struct Package;
using PackagePtr = ::std::shared_ptr<Package>;

struct Package {
    String name;
    String fullName;

    PackagePtr parent;
    Array<PackagePtr> children;
};

}
}

#endif //MOJO_LANG_PACKAGE_HPP
