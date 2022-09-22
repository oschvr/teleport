// package: teleport.terminal.v1
// file: v1/startup_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as v1_startup_service_pb from "../v1/startup_service_pb";

interface IStartupServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    resolveTshdEventsServerAddress: IStartupServiceService_IResolveTshdEventsServerAddress;
}

interface IStartupServiceService_IResolveTshdEventsServerAddress extends grpc.MethodDefinition<v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, v1_startup_service_pb.ResolveTshdEventsServerAddressResponse> {
    path: "/teleport.terminal.v1.StartupService/ResolveTshdEventsServerAddress";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<v1_startup_service_pb.ResolveTshdEventsServerAddressRequest>;
    requestDeserialize: grpc.deserialize<v1_startup_service_pb.ResolveTshdEventsServerAddressRequest>;
    responseSerialize: grpc.serialize<v1_startup_service_pb.ResolveTshdEventsServerAddressResponse>;
    responseDeserialize: grpc.deserialize<v1_startup_service_pb.ResolveTshdEventsServerAddressResponse>;
}

export const StartupServiceService: IStartupServiceService;

export interface IStartupServiceServer {
    resolveTshdEventsServerAddress: grpc.handleUnaryCall<v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, v1_startup_service_pb.ResolveTshdEventsServerAddressResponse>;
}

export interface IStartupServiceClient {
    resolveTshdEventsServerAddress(request: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, callback: (error: grpc.ServiceError | null, response: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse) => void): grpc.ClientUnaryCall;
    resolveTshdEventsServerAddress(request: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse) => void): grpc.ClientUnaryCall;
    resolveTshdEventsServerAddress(request: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse) => void): grpc.ClientUnaryCall;
}

export class StartupServiceClient extends grpc.Client implements IStartupServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public resolveTshdEventsServerAddress(request: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, callback: (error: grpc.ServiceError | null, response: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse) => void): grpc.ClientUnaryCall;
    public resolveTshdEventsServerAddress(request: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse) => void): grpc.ClientUnaryCall;
    public resolveTshdEventsServerAddress(request: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse) => void): grpc.ClientUnaryCall;
}
