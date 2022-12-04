// @generated by protobuf-ts 2.8.2
// @generated from protobuf file "logger/v1alpha1/log.proto" (package "logger.v1alpha1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { LoggerAPI } from "./log";
import type { ReadResponse } from "./log";
import type { ReadRequest } from "./log";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { WriteResponse } from "./log";
import type { WriteRequest } from "./log";
import type { ClientStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service logger.v1alpha1.LoggerAPI
 */
export interface ILoggerAPIClient {
    /**
     * @generated from protobuf rpc: Write(stream logger.v1alpha1.WriteRequest) returns (logger.v1alpha1.WriteResponse);
     */
    write(options?: RpcOptions): ClientStreamingCall<WriteRequest, WriteResponse>;
    /**
     * @generated from protobuf rpc: Read(logger.v1alpha1.ReadRequest) returns (stream logger.v1alpha1.ReadResponse);
     */
    read(input: ReadRequest, options?: RpcOptions): ServerStreamingCall<ReadRequest, ReadResponse>;
}
/**
 * @generated from protobuf service logger.v1alpha1.LoggerAPI
 */
export class LoggerAPIClient implements ILoggerAPIClient, ServiceInfo {
    typeName = LoggerAPI.typeName;
    methods = LoggerAPI.methods;
    options = LoggerAPI.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Write(stream logger.v1alpha1.WriteRequest) returns (logger.v1alpha1.WriteResponse);
     */
    write(options?: RpcOptions): ClientStreamingCall<WriteRequest, WriteResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<WriteRequest, WriteResponse>("clientStreaming", this._transport, method, opt);
    }
    /**
     * @generated from protobuf rpc: Read(logger.v1alpha1.ReadRequest) returns (stream logger.v1alpha1.ReadResponse);
     */
    read(input: ReadRequest, options?: RpcOptions): ServerStreamingCall<ReadRequest, ReadResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<ReadRequest, ReadResponse>("serverStreaming", this._transport, method, opt, input);
    }
}
