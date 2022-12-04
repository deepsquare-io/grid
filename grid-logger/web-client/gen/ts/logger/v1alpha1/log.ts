// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "logger/v1alpha1/log.proto" (package "logger.v1alpha1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message logger.v1alpha1.WriteRequest
 */
export interface WriteRequest {
    /**
     * / log_name is used as key for the logs
     *
     * @generated from protobuf field: string log_name = 1;
     */
    logName: string;
    /**
     * / user represents the owner of the logs
     *
     * @generated from protobuf field: string user = 2;
     */
    user: string;
    /**
     * / data is the actual log
     *
     * @generated from protobuf field: bytes data = 3;
     */
    data: Uint8Array;
}
/**
 * @generated from protobuf message logger.v1alpha1.WriteResponse
 */
export interface WriteResponse {
}
/**
 * @generated from protobuf message logger.v1alpha1.ReadRequest
 */
export interface ReadRequest {
    /**
     * @generated from protobuf field: string log_name = 1;
     */
    logName: string;
}
/**
 * @generated from protobuf message logger.v1alpha1.ReadResponse
 */
export interface ReadResponse {
    /**
     * @generated from protobuf field: bytes data = 1;
     */
    data: Uint8Array;
}
// @generated message type with reflection information, may provide speed optimized methods
class WriteRequest$Type extends MessageType<WriteRequest> {
    constructor() {
        super("logger.v1alpha1.WriteRequest", [
            { no: 1, name: "log_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "user", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<WriteRequest>): WriteRequest {
        const message = { logName: "", user: "", data: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<WriteRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: WriteRequest): WriteRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string log_name */ 1:
                    message.logName = reader.string();
                    break;
                case /* string user */ 2:
                    message.user = reader.string();
                    break;
                case /* bytes data */ 3:
                    message.data = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: WriteRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string log_name = 1; */
        if (message.logName !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.logName);
        /* string user = 2; */
        if (message.user !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.user);
        /* bytes data = 3; */
        if (message.data.length)
            writer.tag(3, WireType.LengthDelimited).bytes(message.data);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message logger.v1alpha1.WriteRequest
 */
export const WriteRequest = new WriteRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class WriteResponse$Type extends MessageType<WriteResponse> {
    constructor() {
        super("logger.v1alpha1.WriteResponse", []);
    }
    create(value?: PartialMessage<WriteResponse>): WriteResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<WriteResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: WriteResponse): WriteResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: WriteResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message logger.v1alpha1.WriteResponse
 */
export const WriteResponse = new WriteResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReadRequest$Type extends MessageType<ReadRequest> {
    constructor() {
        super("logger.v1alpha1.ReadRequest", [
            { no: 1, name: "log_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ReadRequest>): ReadRequest {
        const message = { logName: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ReadRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ReadRequest): ReadRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string log_name */ 1:
                    message.logName = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ReadRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string log_name = 1; */
        if (message.logName !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.logName);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message logger.v1alpha1.ReadRequest
 */
export const ReadRequest = new ReadRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReadResponse$Type extends MessageType<ReadResponse> {
    constructor() {
        super("logger.v1alpha1.ReadResponse", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<ReadResponse>): ReadResponse {
        const message = { data: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ReadResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ReadResponse): ReadResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes data */ 1:
                    message.data = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ReadResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes data = 1; */
        if (message.data.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.data);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message logger.v1alpha1.ReadResponse
 */
export const ReadResponse = new ReadResponse$Type();
/**
 * @generated ServiceType for protobuf service logger.v1alpha1.LoggerAPI
 */
export const LoggerAPI = new ServiceType("logger.v1alpha1.LoggerAPI", [
    { name: "Write", clientStreaming: true, options: {}, I: WriteRequest, O: WriteResponse },
    { name: "Read", serverStreaming: true, options: {}, I: ReadRequest, O: ReadResponse }
]);
