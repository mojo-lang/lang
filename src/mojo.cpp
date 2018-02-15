#include <mojo/lang/package.mojo.hpp>
#include <mojo/lang/source_file.mojo.hpp>
#include <mojo/lang/statement.mojo.hpp>
#include <mojo/lang/struct_decl.mojo.hpp>
#include <mojo/lang/declaration.mojo.hpp>

using namespace mojo;
using namespace mojo::lang;

int main(int argc, char* argv[]) {

    auto package = std::make_shared<Package>();
    auto sourceFile = std::make_shared<SourceFile>();

    auto structField = new StructTypeField;
    structField->name = "foo";
    structField->type.name = "Bar";

    auto structType = new StructType{};
    structType->name = "Foo";
    structType->fields.push_back(structField);

    auto structDecl = std::make_shared<StructDecl>();
    structDecl->name = "Foo";
    structDecl->document = "test";
    structDecl->type = structType;
    structDecl->package = package;
    structDecl->sourceFile = sourceFile;

    sourceFile->name = "test";
    sourceFile->statements.push_back(Statement{Declaration{structDecl}});
    sourceFile->structDecls.emplace(structDecl->name, structDecl);
    package->sourceFiles.push_back(std::move(sourceFile));



    return 0;
}


