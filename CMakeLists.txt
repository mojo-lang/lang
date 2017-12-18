cmake_minimum_required(VERSION 3.7)
project(lang)

set(CMAKE_CXX_STANDARD 14)

set(SOURCE_FILES
        include/mojo/lang/array_adaptor.hpp
        include/mojo/lang/array_type.hpp
        include/mojo/lang/attribute.hpp
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
        include/mojo/lang/option_type.hpp
        include/mojo/lang/package.hpp
        include/mojo/lang/service_method.hpp
        include/mojo/lang/service_type.hpp
        include/mojo/lang/service_type_builder.hpp
        include/mojo/lang/service_type_of.hpp
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
        test/mojo/lang/data_type_test.cpp)

add_executable(lang ${SOURCE_FILES})