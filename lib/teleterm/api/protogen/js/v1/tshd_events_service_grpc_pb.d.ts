// package: teleport.terminal.v1
// file: v1/tshd_events_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as v1_tshd_events_service_pb from "../v1/tshd_events_service_pb";

interface ITshdEventsServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    newGatewayConnectionAccepted: ITshdEventsServiceService_INewGatewayConnectionAccepted;
}

interface ITshdEventsServiceService_INewGatewayConnectionAccepted extends grpc.MethodDefinition<v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse> {
    path: "/teleport.terminal.v1.TshdEventsService/NewGatewayConnectionAccepted";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest>;
    requestDeserialize: grpc.deserialize<v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest>;
    responseSerialize: grpc.serialize<v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse>;
    responseDeserialize: grpc.deserialize<v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse>;
}

export const TshdEventsServiceService: ITshdEventsServiceService;

export interface ITshdEventsServiceServer {
    newGatewayConnectionAccepted: grpc.handleUnaryCall<v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse>;
}

export interface ITshdEventsServiceClient {
    newGatewayConnectionAccepted(request: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, callback: (error: grpc.ServiceError | null, response: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse) => void): grpc.ClientUnaryCall;
    newGatewayConnectionAccepted(request: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse) => void): grpc.ClientUnaryCall;
    newGatewayConnectionAccepted(request: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse) => void): grpc.ClientUnaryCall;
}

export class TshdEventsServiceClient extends grpc.Client implements ITshdEventsServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public newGatewayConnectionAccepted(request: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, callback: (error: grpc.ServiceError | null, response: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse) => void): grpc.ClientUnaryCall;
    public newGatewayConnectionAccepted(request: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse) => void): grpc.ClientUnaryCall;
    public newGatewayConnectionAccepted(request: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse) => void): grpc.ClientUnaryCall;
}
