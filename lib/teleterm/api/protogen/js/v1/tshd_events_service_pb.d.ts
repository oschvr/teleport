// package: teleport.terminal.v1
// file: v1/tshd_events_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class NewGatewayConnectionAcceptedRequest extends jspb.Message { 
    getGatewayUri(): string;
    setGatewayUri(value: string): NewGatewayConnectionAcceptedRequest;

    getTargetUri(): string;
    setTargetUri(value: string): NewGatewayConnectionAcceptedRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewGatewayConnectionAcceptedRequest.AsObject;
    static toObject(includeInstance: boolean, msg: NewGatewayConnectionAcceptedRequest): NewGatewayConnectionAcceptedRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewGatewayConnectionAcceptedRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewGatewayConnectionAcceptedRequest;
    static deserializeBinaryFromReader(message: NewGatewayConnectionAcceptedRequest, reader: jspb.BinaryReader): NewGatewayConnectionAcceptedRequest;
}

export namespace NewGatewayConnectionAcceptedRequest {
    export type AsObject = {
        gatewayUri: string,
        targetUri: string,
    }
}

export class NewGatewayConnectionAcceptedResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NewGatewayConnectionAcceptedResponse.AsObject;
    static toObject(includeInstance: boolean, msg: NewGatewayConnectionAcceptedResponse): NewGatewayConnectionAcceptedResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NewGatewayConnectionAcceptedResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NewGatewayConnectionAcceptedResponse;
    static deserializeBinaryFromReader(message: NewGatewayConnectionAcceptedResponse, reader: jspb.BinaryReader): NewGatewayConnectionAcceptedResponse;
}

export namespace NewGatewayConnectionAcceptedResponse {
    export type AsObject = {
    }
}
