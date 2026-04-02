import * as jspb from 'google-protobuf'



export class CalcRequest extends jspb.Message {
  getA(): number;
  setA(value: number): CalcRequest;

  getB(): number;
  setB(value: number): CalcRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CalcRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CalcRequest): CalcRequest.AsObject;
  static serializeBinaryToWriter(message: CalcRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CalcRequest;
  static deserializeBinaryFromReader(message: CalcRequest, reader: jspb.BinaryReader): CalcRequest;
}

export namespace CalcRequest {
  export type AsObject = {
    a: number;
    b: number;
  };
}

export class CalcResponse extends jspb.Message {
  getResult(): number;
  setResult(value: number): CalcResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CalcResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CalcResponse): CalcResponse.AsObject;
  static serializeBinaryToWriter(message: CalcResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CalcResponse;
  static deserializeBinaryFromReader(message: CalcResponse, reader: jspb.BinaryReader): CalcResponse;
}

export namespace CalcResponse {
  export type AsObject = {
    result: number;
  };
}

