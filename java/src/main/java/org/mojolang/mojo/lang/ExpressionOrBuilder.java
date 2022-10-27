// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/lang.proto

package org.mojolang.mojo.lang;

public interface ExpressionOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Expression)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.lang.NullLiteralExpr null_literal_expr = 1;</code>
   * @return Whether the nullLiteralExpr field is set.
   */
  boolean hasNullLiteralExpr();
  /**
   * <code>.mojo.lang.NullLiteralExpr null_literal_expr = 1;</code>
   * @return The nullLiteralExpr.
   */
  org.mojolang.mojo.lang.NullLiteralExpr getNullLiteralExpr();
  /**
   * <code>.mojo.lang.NullLiteralExpr null_literal_expr = 1;</code>
   */
  org.mojolang.mojo.lang.NullLiteralExprOrBuilder getNullLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.IntegerLiteralExpr integer_literal_expr = 2;</code>
   * @return Whether the integerLiteralExpr field is set.
   */
  boolean hasIntegerLiteralExpr();
  /**
   * <code>.mojo.lang.IntegerLiteralExpr integer_literal_expr = 2;</code>
   * @return The integerLiteralExpr.
   */
  org.mojolang.mojo.lang.IntegerLiteralExpr getIntegerLiteralExpr();
  /**
   * <code>.mojo.lang.IntegerLiteralExpr integer_literal_expr = 2;</code>
   */
  org.mojolang.mojo.lang.IntegerLiteralExprOrBuilder getIntegerLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.FloatLiteralExpr float_literal_expr = 3;</code>
   * @return Whether the floatLiteralExpr field is set.
   */
  boolean hasFloatLiteralExpr();
  /**
   * <code>.mojo.lang.FloatLiteralExpr float_literal_expr = 3;</code>
   * @return The floatLiteralExpr.
   */
  org.mojolang.mojo.lang.FloatLiteralExpr getFloatLiteralExpr();
  /**
   * <code>.mojo.lang.FloatLiteralExpr float_literal_expr = 3;</code>
   */
  org.mojolang.mojo.lang.FloatLiteralExprOrBuilder getFloatLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.BoolLiteralExpr bool_literal_expr = 4;</code>
   * @return Whether the boolLiteralExpr field is set.
   */
  boolean hasBoolLiteralExpr();
  /**
   * <code>.mojo.lang.BoolLiteralExpr bool_literal_expr = 4;</code>
   * @return The boolLiteralExpr.
   */
  org.mojolang.mojo.lang.BoolLiteralExpr getBoolLiteralExpr();
  /**
   * <code>.mojo.lang.BoolLiteralExpr bool_literal_expr = 4;</code>
   */
  org.mojolang.mojo.lang.BoolLiteralExprOrBuilder getBoolLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.StringLiteralExpr string_literal_expr = 5;</code>
   * @return Whether the stringLiteralExpr field is set.
   */
  boolean hasStringLiteralExpr();
  /**
   * <code>.mojo.lang.StringLiteralExpr string_literal_expr = 5;</code>
   * @return The stringLiteralExpr.
   */
  org.mojolang.mojo.lang.StringLiteralExpr getStringLiteralExpr();
  /**
   * <code>.mojo.lang.StringLiteralExpr string_literal_expr = 5;</code>
   */
  org.mojolang.mojo.lang.StringLiteralExprOrBuilder getStringLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.ObjectLiteralExpr object_literal_expr = 6;</code>
   * @return Whether the objectLiteralExpr field is set.
   */
  boolean hasObjectLiteralExpr();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr object_literal_expr = 6;</code>
   * @return The objectLiteralExpr.
   */
  org.mojolang.mojo.lang.ObjectLiteralExpr getObjectLiteralExpr();
  /**
   * <code>.mojo.lang.ObjectLiteralExpr object_literal_expr = 6;</code>
   */
  org.mojolang.mojo.lang.ObjectLiteralExprOrBuilder getObjectLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.ArrayLiteralExpr array_literal_expr = 7;</code>
   * @return Whether the arrayLiteralExpr field is set.
   */
  boolean hasArrayLiteralExpr();
  /**
   * <code>.mojo.lang.ArrayLiteralExpr array_literal_expr = 7;</code>
   * @return The arrayLiteralExpr.
   */
  org.mojolang.mojo.lang.ArrayLiteralExpr getArrayLiteralExpr();
  /**
   * <code>.mojo.lang.ArrayLiteralExpr array_literal_expr = 7;</code>
   */
  org.mojolang.mojo.lang.ArrayLiteralExprOrBuilder getArrayLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.MapLiteralExpr map_literal_expr = 8;</code>
   * @return Whether the mapLiteralExpr field is set.
   */
  boolean hasMapLiteralExpr();
  /**
   * <code>.mojo.lang.MapLiteralExpr map_literal_expr = 8;</code>
   * @return The mapLiteralExpr.
   */
  org.mojolang.mojo.lang.MapLiteralExpr getMapLiteralExpr();
  /**
   * <code>.mojo.lang.MapLiteralExpr map_literal_expr = 8;</code>
   */
  org.mojolang.mojo.lang.MapLiteralExprOrBuilder getMapLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.RangeLiteralExpr range_literal_expr = 9;</code>
   * @return Whether the rangeLiteralExpr field is set.
   */
  boolean hasRangeLiteralExpr();
  /**
   * <code>.mojo.lang.RangeLiteralExpr range_literal_expr = 9;</code>
   * @return The rangeLiteralExpr.
   */
  org.mojolang.mojo.lang.RangeLiteralExpr getRangeLiteralExpr();
  /**
   * <code>.mojo.lang.RangeLiteralExpr range_literal_expr = 9;</code>
   */
  org.mojolang.mojo.lang.RangeLiteralExprOrBuilder getRangeLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.IdentifierExpr identifier_expr = 10;</code>
   * @return Whether the identifierExpr field is set.
   */
  boolean hasIdentifierExpr();
  /**
   * <code>.mojo.lang.IdentifierExpr identifier_expr = 10;</code>
   * @return The identifierExpr.
   */
  org.mojolang.mojo.lang.IdentifierExpr getIdentifierExpr();
  /**
   * <code>.mojo.lang.IdentifierExpr identifier_expr = 10;</code>
   */
  org.mojolang.mojo.lang.IdentifierExprOrBuilder getIdentifierExprOrBuilder();

  /**
   * <code>.mojo.lang.NumericSuffixLiteralExpr numeric_suffix_literal_expr = 11;</code>
   * @return Whether the numericSuffixLiteralExpr field is set.
   */
  boolean hasNumericSuffixLiteralExpr();
  /**
   * <code>.mojo.lang.NumericSuffixLiteralExpr numeric_suffix_literal_expr = 11;</code>
   * @return The numericSuffixLiteralExpr.
   */
  org.mojolang.mojo.lang.NumericSuffixLiteralExpr getNumericSuffixLiteralExpr();
  /**
   * <code>.mojo.lang.NumericSuffixLiteralExpr numeric_suffix_literal_expr = 11;</code>
   */
  org.mojolang.mojo.lang.NumericSuffixLiteralExprOrBuilder getNumericSuffixLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.StringPrefixLiteralExpr string_prefix_literal_expr = 12;</code>
   * @return Whether the stringPrefixLiteralExpr field is set.
   */
  boolean hasStringPrefixLiteralExpr();
  /**
   * <code>.mojo.lang.StringPrefixLiteralExpr string_prefix_literal_expr = 12;</code>
   * @return The stringPrefixLiteralExpr.
   */
  org.mojolang.mojo.lang.StringPrefixLiteralExpr getStringPrefixLiteralExpr();
  /**
   * <code>.mojo.lang.StringPrefixLiteralExpr string_prefix_literal_expr = 12;</code>
   */
  org.mojolang.mojo.lang.StringPrefixLiteralExprOrBuilder getStringPrefixLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.StringSuffixLiteralExpr string_suffix_literal_expr = 13;</code>
   * @return Whether the stringSuffixLiteralExpr field is set.
   */
  boolean hasStringSuffixLiteralExpr();
  /**
   * <code>.mojo.lang.StringSuffixLiteralExpr string_suffix_literal_expr = 13;</code>
   * @return The stringSuffixLiteralExpr.
   */
  org.mojolang.mojo.lang.StringSuffixLiteralExpr getStringSuffixLiteralExpr();
  /**
   * <code>.mojo.lang.StringSuffixLiteralExpr string_suffix_literal_expr = 13;</code>
   */
  org.mojolang.mojo.lang.StringSuffixLiteralExprOrBuilder getStringSuffixLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.StructLiteralExpr struct_literal_expr = 14;</code>
   * @return Whether the structLiteralExpr field is set.
   */
  boolean hasStructLiteralExpr();
  /**
   * <code>.mojo.lang.StructLiteralExpr struct_literal_expr = 14;</code>
   * @return The structLiteralExpr.
   */
  org.mojolang.mojo.lang.StructLiteralExpr getStructLiteralExpr();
  /**
   * <code>.mojo.lang.StructLiteralExpr struct_literal_expr = 14;</code>
   */
  org.mojolang.mojo.lang.StructLiteralExprOrBuilder getStructLiteralExprOrBuilder();

  /**
   * <code>.mojo.lang.ClosureExpr closure_expr = 15;</code>
   * @return Whether the closureExpr field is set.
   */
  boolean hasClosureExpr();
  /**
   * <code>.mojo.lang.ClosureExpr closure_expr = 15;</code>
   * @return The closureExpr.
   */
  org.mojolang.mojo.lang.ClosureExpr getClosureExpr();
  /**
   * <code>.mojo.lang.ClosureExpr closure_expr = 15;</code>
   */
  org.mojolang.mojo.lang.ClosureExprOrBuilder getClosureExprOrBuilder();

  /**
   * <code>.mojo.lang.ParenthesizedExpr parenthesized_expr = 16;</code>
   * @return Whether the parenthesizedExpr field is set.
   */
  boolean hasParenthesizedExpr();
  /**
   * <code>.mojo.lang.ParenthesizedExpr parenthesized_expr = 16;</code>
   * @return The parenthesizedExpr.
   */
  org.mojolang.mojo.lang.ParenthesizedExpr getParenthesizedExpr();
  /**
   * <code>.mojo.lang.ParenthesizedExpr parenthesized_expr = 16;</code>
   */
  org.mojolang.mojo.lang.ParenthesizedExprOrBuilder getParenthesizedExprOrBuilder();

  /**
   * <code>.mojo.lang.ImplicitMemberExpr implicit_member_expr = 17;</code>
   * @return Whether the implicitMemberExpr field is set.
   */
  boolean hasImplicitMemberExpr();
  /**
   * <code>.mojo.lang.ImplicitMemberExpr implicit_member_expr = 17;</code>
   * @return The implicitMemberExpr.
   */
  org.mojolang.mojo.lang.ImplicitMemberExpr getImplicitMemberExpr();
  /**
   * <code>.mojo.lang.ImplicitMemberExpr implicit_member_expr = 17;</code>
   */
  org.mojolang.mojo.lang.ImplicitMemberExprOrBuilder getImplicitMemberExprOrBuilder();

  /**
   * <code>.mojo.lang.WildcardExpr wildcard_expr = 18;</code>
   * @return Whether the wildcardExpr field is set.
   */
  boolean hasWildcardExpr();
  /**
   * <code>.mojo.lang.WildcardExpr wildcard_expr = 18;</code>
   * @return The wildcardExpr.
   */
  org.mojolang.mojo.lang.WildcardExpr getWildcardExpr();
  /**
   * <code>.mojo.lang.WildcardExpr wildcard_expr = 18;</code>
   */
  org.mojolang.mojo.lang.WildcardExprOrBuilder getWildcardExprOrBuilder();

  /**
   * <code>.mojo.lang.StructConstructionExpr struct_construction_expr = 19;</code>
   * @return Whether the structConstructionExpr field is set.
   */
  boolean hasStructConstructionExpr();
  /**
   * <code>.mojo.lang.StructConstructionExpr struct_construction_expr = 19;</code>
   * @return The structConstructionExpr.
   */
  org.mojolang.mojo.lang.StructConstructionExpr getStructConstructionExpr();
  /**
   * <code>.mojo.lang.StructConstructionExpr struct_construction_expr = 19;</code>
   */
  org.mojolang.mojo.lang.StructConstructionExprOrBuilder getStructConstructionExprOrBuilder();

  /**
   * <code>.mojo.lang.TupleExpr tuple_expr = 20;</code>
   * @return Whether the tupleExpr field is set.
   */
  boolean hasTupleExpr();
  /**
   * <code>.mojo.lang.TupleExpr tuple_expr = 20;</code>
   * @return The tupleExpr.
   */
  org.mojolang.mojo.lang.TupleExpr getTupleExpr();
  /**
   * <code>.mojo.lang.TupleExpr tuple_expr = 20;</code>
   */
  org.mojolang.mojo.lang.TupleExprOrBuilder getTupleExprOrBuilder();

  /**
   * <code>.mojo.lang.PrefixUnaryExpr prefix_unary_expr = 30;</code>
   * @return Whether the prefixUnaryExpr field is set.
   */
  boolean hasPrefixUnaryExpr();
  /**
   * <code>.mojo.lang.PrefixUnaryExpr prefix_unary_expr = 30;</code>
   * @return The prefixUnaryExpr.
   */
  org.mojolang.mojo.lang.PrefixUnaryExpr getPrefixUnaryExpr();
  /**
   * <code>.mojo.lang.PrefixUnaryExpr prefix_unary_expr = 30;</code>
   */
  org.mojolang.mojo.lang.PrefixUnaryExprOrBuilder getPrefixUnaryExprOrBuilder();

  /**
   * <code>.mojo.lang.PostfixUnaryExpr postfix_unary_expr = 31;</code>
   * @return Whether the postfixUnaryExpr field is set.
   */
  boolean hasPostfixUnaryExpr();
  /**
   * <code>.mojo.lang.PostfixUnaryExpr postfix_unary_expr = 31;</code>
   * @return The postfixUnaryExpr.
   */
  org.mojolang.mojo.lang.PostfixUnaryExpr getPostfixUnaryExpr();
  /**
   * <code>.mojo.lang.PostfixUnaryExpr postfix_unary_expr = 31;</code>
   */
  org.mojolang.mojo.lang.PostfixUnaryExprOrBuilder getPostfixUnaryExprOrBuilder();

  /**
   * <code>.mojo.lang.FunctionCallExpr function_call_expr = 32;</code>
   * @return Whether the functionCallExpr field is set.
   */
  boolean hasFunctionCallExpr();
  /**
   * <code>.mojo.lang.FunctionCallExpr function_call_expr = 32;</code>
   * @return The functionCallExpr.
   */
  org.mojolang.mojo.lang.FunctionCallExpr getFunctionCallExpr();
  /**
   * <code>.mojo.lang.FunctionCallExpr function_call_expr = 32;</code>
   */
  org.mojolang.mojo.lang.FunctionCallExprOrBuilder getFunctionCallExprOrBuilder();

  /**
   * <code>.mojo.lang.ExplicitMemberExpr explicit_member_expr = 33;</code>
   * @return Whether the explicitMemberExpr field is set.
   */
  boolean hasExplicitMemberExpr();
  /**
   * <code>.mojo.lang.ExplicitMemberExpr explicit_member_expr = 33;</code>
   * @return The explicitMemberExpr.
   */
  org.mojolang.mojo.lang.ExplicitMemberExpr getExplicitMemberExpr();
  /**
   * <code>.mojo.lang.ExplicitMemberExpr explicit_member_expr = 33;</code>
   */
  org.mojolang.mojo.lang.ExplicitMemberExprOrBuilder getExplicitMemberExprOrBuilder();

  /**
   * <code>.mojo.lang.SubscriptExpr subscript_expr = 34;</code>
   * @return Whether the subscriptExpr field is set.
   */
  boolean hasSubscriptExpr();
  /**
   * <code>.mojo.lang.SubscriptExpr subscript_expr = 34;</code>
   * @return The subscriptExpr.
   */
  org.mojolang.mojo.lang.SubscriptExpr getSubscriptExpr();
  /**
   * <code>.mojo.lang.SubscriptExpr subscript_expr = 34;</code>
   */
  org.mojolang.mojo.lang.SubscriptExprOrBuilder getSubscriptExprOrBuilder();

  /**
   * <code>.mojo.lang.BinaryExpr binary_expr = 40;</code>
   * @return Whether the binaryExpr field is set.
   */
  boolean hasBinaryExpr();
  /**
   * <code>.mojo.lang.BinaryExpr binary_expr = 40;</code>
   * @return The binaryExpr.
   */
  org.mojolang.mojo.lang.BinaryExpr getBinaryExpr();
  /**
   * <code>.mojo.lang.BinaryExpr binary_expr = 40;</code>
   */
  org.mojolang.mojo.lang.BinaryExprOrBuilder getBinaryExprOrBuilder();

  /**
   * <code>.mojo.lang.ConditionalExpr conditional_expr = 41;</code>
   * @return Whether the conditionalExpr field is set.
   */
  boolean hasConditionalExpr();
  /**
   * <code>.mojo.lang.ConditionalExpr conditional_expr = 41;</code>
   * @return The conditionalExpr.
   */
  org.mojolang.mojo.lang.ConditionalExpr getConditionalExpr();
  /**
   * <code>.mojo.lang.ConditionalExpr conditional_expr = 41;</code>
   */
  org.mojolang.mojo.lang.ConditionalExprOrBuilder getConditionalExprOrBuilder();

  /**
   * <code>.mojo.lang.TypeCastingExpr type_casting_expr = 42;</code>
   * @return Whether the typeCastingExpr field is set.
   */
  boolean hasTypeCastingExpr();
  /**
   * <code>.mojo.lang.TypeCastingExpr type_casting_expr = 42;</code>
   * @return The typeCastingExpr.
   */
  org.mojolang.mojo.lang.TypeCastingExpr getTypeCastingExpr();
  /**
   * <code>.mojo.lang.TypeCastingExpr type_casting_expr = 42;</code>
   */
  org.mojolang.mojo.lang.TypeCastingExprOrBuilder getTypeCastingExprOrBuilder();

  /**
   * <code>.mojo.lang.AssignmentExpr assignment_expr = 43;</code>
   * @return Whether the assignmentExpr field is set.
   */
  boolean hasAssignmentExpr();
  /**
   * <code>.mojo.lang.AssignmentExpr assignment_expr = 43;</code>
   * @return The assignmentExpr.
   */
  org.mojolang.mojo.lang.AssignmentExpr getAssignmentExpr();
  /**
   * <code>.mojo.lang.AssignmentExpr assignment_expr = 43;</code>
   */
  org.mojolang.mojo.lang.AssignmentExprOrBuilder getAssignmentExprOrBuilder();

  /**
   * <code>.mojo.lang.ErrorExpr error_expr = 60;</code>
   * @return Whether the errorExpr field is set.
   */
  boolean hasErrorExpr();
  /**
   * <code>.mojo.lang.ErrorExpr error_expr = 60;</code>
   * @return The errorExpr.
   */
  org.mojolang.mojo.lang.ErrorExpr getErrorExpr();
  /**
   * <code>.mojo.lang.ErrorExpr error_expr = 60;</code>
   */
  org.mojolang.mojo.lang.ErrorExprOrBuilder getErrorExprOrBuilder();

  public org.mojolang.mojo.lang.Expression.ExpressionCase getExpressionCase();
}
