#ifndef MOJO_LANG_DECL_MOJO_HPP
#define MOJO_LANG_DECL_MOJO_HPP

namespace mojo {
namespace lang {

struct Decl {
    /// position of first character belonging to the Expr
    //start_pos: Pos @1

    /// position of first character immediately after the Expr
    //end_pos  : Pos @2

    ///
    //document : Document @3

    ///
    //package: Package @4 @link
};

}
}

#endif //MOJO_LANG_DECL_MOJO_HPP
