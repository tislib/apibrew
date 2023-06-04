// @generated by protoc-gen-es v1.2.1 with parameter "target=ts"
// @generated from file model/query.proto (package model, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Value } from "@bufbuild/protobuf";

/**
 * @generated from message model.CompoundBooleanExpression
 */
export class CompoundBooleanExpression extends Message<CompoundBooleanExpression> {
  /**
   * @generated from field: repeated model.BooleanExpression expressions = 1;
   */
  expressions: BooleanExpression[] = [];

  constructor(data?: PartialMessage<CompoundBooleanExpression>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "model.CompoundBooleanExpression";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "expressions", kind: "message", T: BooleanExpression, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompoundBooleanExpression {
    return new CompoundBooleanExpression().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompoundBooleanExpression {
    return new CompoundBooleanExpression().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompoundBooleanExpression {
    return new CompoundBooleanExpression().fromJsonString(jsonString, options);
  }

  static equals(a: CompoundBooleanExpression | PlainMessage<CompoundBooleanExpression> | undefined, b: CompoundBooleanExpression | PlainMessage<CompoundBooleanExpression> | undefined): boolean {
    return proto3.util.equals(CompoundBooleanExpression, a, b);
  }
}

/**
 * @generated from message model.RefValue
 */
export class RefValue extends Message<RefValue> {
  /**
   * @generated from field: string namespace = 1;
   */
  namespace = "";

  /**
   * @generated from field: string resource = 2;
   */
  resource = "";

  /**
   * @generated from field: map<string, google.protobuf.Value> properties = 3;
   */
  properties: { [key: string]: Value } = {};

  constructor(data?: PartialMessage<RefValue>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "model.RefValue";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "resource", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "properties", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Value} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RefValue {
    return new RefValue().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RefValue {
    return new RefValue().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RefValue {
    return new RefValue().fromJsonString(jsonString, options);
  }

  static equals(a: RefValue | PlainMessage<RefValue> | undefined, b: RefValue | PlainMessage<RefValue> | undefined): boolean {
    return proto3.util.equals(RefValue, a, b);
  }
}

/**
 * @generated from message model.Expression
 */
export class Expression extends Message<Expression> {
  /**
   * @generated from oneof model.Expression.expression
   */
  expression: {
    /**
     * @generated from field: string property = 1;
     */
    value: string;
    case: "property";
  } | {
    /**
     * @generated from field: google.protobuf.Value value = 3;
     */
    value: Value;
    case: "value";
  } | {
    /**
     * @generated from field: model.RefValue refValue = 4;
     */
    value: RefValue;
    case: "refValue";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Expression>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "model.Expression";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "property", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "expression" },
    { no: 3, name: "value", kind: "message", T: Value, oneof: "expression" },
    { no: 4, name: "refValue", kind: "message", T: RefValue, oneof: "expression" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Expression {
    return new Expression().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Expression {
    return new Expression().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Expression {
    return new Expression().fromJsonString(jsonString, options);
  }

  static equals(a: Expression | PlainMessage<Expression> | undefined, b: Expression | PlainMessage<Expression> | undefined): boolean {
    return proto3.util.equals(Expression, a, b);
  }
}

/**
 * @generated from message model.PairExpression
 */
export class PairExpression extends Message<PairExpression> {
  /**
   * @generated from field: model.Expression left = 1;
   */
  left?: Expression;

  /**
   * @generated from field: model.Expression right = 2;
   */
  right?: Expression;

  constructor(data?: PartialMessage<PairExpression>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "model.PairExpression";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "left", kind: "message", T: Expression },
    { no: 2, name: "right", kind: "message", T: Expression },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PairExpression {
    return new PairExpression().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PairExpression {
    return new PairExpression().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PairExpression {
    return new PairExpression().fromJsonString(jsonString, options);
  }

  static equals(a: PairExpression | PlainMessage<PairExpression> | undefined, b: PairExpression | PlainMessage<PairExpression> | undefined): boolean {
    return proto3.util.equals(PairExpression, a, b);
  }
}

/**
 * @generated from message model.RegexMatchExpression
 */
export class RegexMatchExpression extends Message<RegexMatchExpression> {
  /**
   * @generated from field: string pattern = 1;
   */
  pattern = "";

  /**
   * @generated from field: model.Expression expression = 2;
   */
  expression?: Expression;

  constructor(data?: PartialMessage<RegexMatchExpression>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "model.RegexMatchExpression";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pattern", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expression", kind: "message", T: Expression },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegexMatchExpression {
    return new RegexMatchExpression().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegexMatchExpression {
    return new RegexMatchExpression().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegexMatchExpression {
    return new RegexMatchExpression().fromJsonString(jsonString, options);
  }

  static equals(a: RegexMatchExpression | PlainMessage<RegexMatchExpression> | undefined, b: RegexMatchExpression | PlainMessage<RegexMatchExpression> | undefined): boolean {
    return proto3.util.equals(RegexMatchExpression, a, b);
  }
}

/**
 * @generated from message model.BooleanExpression
 */
export class BooleanExpression extends Message<BooleanExpression> {
  /**
   * @generated from oneof model.BooleanExpression.expression
   */
  expression: {
    /**
     * logical expressions
     *
     * @generated from field: model.CompoundBooleanExpression and = 1;
     */
    value: CompoundBooleanExpression;
    case: "and";
  } | {
    /**
     * @generated from field: model.CompoundBooleanExpression or = 2;
     */
    value: CompoundBooleanExpression;
    case: "or";
  } | {
    /**
     * @generated from field: model.BooleanExpression not = 3;
     */
    value: BooleanExpression;
    case: "not";
  } | {
    /**
     * basic comparison
     *
     * @generated from field: model.PairExpression equal = 4;
     */
    value: PairExpression;
    case: "equal";
  } | {
    /**
     * @generated from field: model.PairExpression lessThan = 5;
     */
    value: PairExpression;
    case: "lessThan";
  } | {
    /**
     * @generated from field: model.PairExpression greaterThan = 6;
     */
    value: PairExpression;
    case: "greaterThan";
  } | {
    /**
     * @generated from field: model.PairExpression lessThanOrEqual = 7;
     */
    value: PairExpression;
    case: "lessThanOrEqual";
  } | {
    /**
     * @generated from field: model.PairExpression greaterThanOrEqual = 8;
     */
    value: PairExpression;
    case: "greaterThanOrEqual";
  } | {
    /**
     * @generated from field: model.PairExpression in = 9;
     */
    value: PairExpression;
    case: "in";
  } | {
    /**
     * @generated from field: model.Expression isNull = 10;
     */
    value: Expression;
    case: "isNull";
  } | {
    /**
     * other
     *
     * @generated from field: model.RegexMatchExpression regexMatch = 11;
     */
    value: RegexMatchExpression;
    case: "regexMatch";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<BooleanExpression>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "model.BooleanExpression";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "and", kind: "message", T: CompoundBooleanExpression, oneof: "expression" },
    { no: 2, name: "or", kind: "message", T: CompoundBooleanExpression, oneof: "expression" },
    { no: 3, name: "not", kind: "message", T: BooleanExpression, oneof: "expression" },
    { no: 4, name: "equal", kind: "message", T: PairExpression, oneof: "expression" },
    { no: 5, name: "lessThan", kind: "message", T: PairExpression, oneof: "expression" },
    { no: 6, name: "greaterThan", kind: "message", T: PairExpression, oneof: "expression" },
    { no: 7, name: "lessThanOrEqual", kind: "message", T: PairExpression, oneof: "expression" },
    { no: 8, name: "greaterThanOrEqual", kind: "message", T: PairExpression, oneof: "expression" },
    { no: 9, name: "in", kind: "message", T: PairExpression, oneof: "expression" },
    { no: 10, name: "isNull", kind: "message", T: Expression, oneof: "expression" },
    { no: 11, name: "regexMatch", kind: "message", T: RegexMatchExpression, oneof: "expression" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BooleanExpression {
    return new BooleanExpression().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BooleanExpression {
    return new BooleanExpression().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BooleanExpression {
    return new BooleanExpression().fromJsonString(jsonString, options);
  }

  static equals(a: BooleanExpression | PlainMessage<BooleanExpression> | undefined, b: BooleanExpression | PlainMessage<BooleanExpression> | undefined): boolean {
    return proto3.util.equals(BooleanExpression, a, b);
  }
}
