cmake_minimum_required(VERSION 3.7)
project(lang)

set(CMAKE_CXX_STANDARD 14)

INCLUDE_DIRECTORIES(
        /usr/local/include
        ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft-core/include
        ${PROJECT_SOURCE_DIR}/../../aimap/ncraft/include
        ${PROJECT_SOURCE_DIR}/../core/include
        ${PROJECT_SOURCE_DIR}/include
        ${PROJECT_SOURCE_DIR}/test
)

set(MOJO_HEADER_FILES
        include/mojo/lang/position.mojo.hpp
        include/mojo/lang/struct_decl.mojo.hpp
        include/mojo/lang/type_decl.mojo.hpp
        include/mojo/lang/numeric_literal_expr.mojo.hpp
        include/mojo/lang/decl.mojo.hpp
        include/mojo/lang/expression.mojo.hpp
        include/mojo/lang/literal_expr.mojo.hpp include/mojo/lang/integer_literal_expr.mojo.hpp include/mojo/lang/data_decl.mojo.hpp include/mojo/lang/value_decl.mojo.hpp include/mojo/lang/type_alias_decl.mojo.hpp include/mojo/lang/enum_decl.mojo.hpp)

set(SOURCE_FILES
        include/mojo/lang/array_adaptor.hpp
        include/mojo/lang/array_type.hpp
        include/mojo/lang/attribute.mojo.hpp
        include/mojo/lang/bool_adaptor.hpp
        include/mojo/lang/bool_type.hpp
        include/mojo/lang/callable.hpp
        include/mojo/lang/data_adaptor.hpp
        include/mojo/lang/data_adaptor_of.hpp
        include/mojo/lang/data_mutator.hpp
        include/mojo/lang/data_type.hpp
        include/mojo/lang/data_type_of.hpp
        include/mojo/lang/data_type_preprocessors.hpp
        include/mojo/lang/data_types.hpp
        include/mojo/lang/data_visitor.hpp
        include/mojo/lang/dictionary_adaptor.hpp
        include/mojo/lang/dictionary_type.hpp
        include/mojo/lang/enum_type.hpp
        include/mojo/lang/enum_type_builder.hpp
        include/mojo/lang/enum_type_of.hpp
        include/mojo/lang/enum_value.hpp
        include/mojo/lang/function.hpp
        include/mojo/lang/generic_type.hpp
        include/mojo/lang/source_file.hpp
        include/mojo/lang/numeric_adaptor.hpp
        include/mojo/lang/numeric_type.hpp
        include/mojo/lang/optional_type.hpp
        include/mojo/lang/package.mojo.hpp
        include/mojo/lang/interface_method.hpp
        include/mojo/lang/interface_type.hpp
        include/mojo/lang/interface_type_builder.hpp
        include/mojo/lang/interface_type_of.hpp
        include/mojo/lang/string_adaptor.hpp
        include/mojo/lang/string_type.hpp
        include/mojo/lang/struct_adaptor.hpp
        include/mojo/lang/struct_field.hpp
        include/mojo/lang/struct_field_adaptor.hpp
        include/mojo/lang/struct_field_of.hpp
        include/mojo/lang/struct_object_mutator.hpp
        include/mojo/lang/struct_object_spectator.hpp
        include/mojo/lang/struct_type.hpp
        include/mojo/lang/struct_type_builder.hpp
        include/mojo/lang/struct_type_of.hpp
        include/mojo/lang/struct_virtual_field.hpp
        include/mojo/lang/term.hpp
        src/mojo/lang/data_type.cpp
        src/mojo/lang/numeric_type.cpp
        src/mojo/lang/string_type.cpp
        test/mojo/lang/data_type_test.cpp include/mojo/lang/string_literal_expr.mojo.hpp include/mojo/lang/nominal_type.mojo.hpp include/mojo/lang/struct_decl_builder.hpp include/mojo/lang/decl.hpp)

add_executable(lang ${SOURCE_FILES} ${MOJO_HEADER_FILES})