// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
'use strict';
var grpc = require('@grpc/grpc-js');
var v1_tshd_events_service_pb = require('../v1/tshd_events_service_pb.js');

function serialize_teleport_terminal_v1_NewGatewayConnectionAcceptedRequest(arg) {
  if (!(arg instanceof v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest)) {
    throw new Error('Expected argument of type teleport.terminal.v1.NewGatewayConnectionAcceptedRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_teleport_terminal_v1_NewGatewayConnectionAcceptedRequest(buffer_arg) {
  return v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_teleport_terminal_v1_NewGatewayConnectionAcceptedResponse(arg) {
  if (!(arg instanceof v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse)) {
    throw new Error('Expected argument of type teleport.terminal.v1.NewGatewayConnectionAcceptedResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_teleport_terminal_v1_NewGatewayConnectionAcceptedResponse(buffer_arg) {
  return v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// TODO: Add comment.
var TshdEventsServiceService = exports.TshdEventsServiceService = {
  // TODO: Add comment.
newGatewayConnectionAccepted: {
    path: '/teleport.terminal.v1.TshdEventsService/NewGatewayConnectionAccepted',
    requestStream: false,
    responseStream: false,
    requestType: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedRequest,
    responseType: v1_tshd_events_service_pb.NewGatewayConnectionAcceptedResponse,
    requestSerialize: serialize_teleport_terminal_v1_NewGatewayConnectionAcceptedRequest,
    requestDeserialize: deserialize_teleport_terminal_v1_NewGatewayConnectionAcceptedRequest,
    responseSerialize: serialize_teleport_terminal_v1_NewGatewayConnectionAcceptedResponse,
    responseDeserialize: deserialize_teleport_terminal_v1_NewGatewayConnectionAcceptedResponse,
  },
};

exports.TshdEventsServiceClient = grpc.makeGenericClientConstructor(TshdEventsServiceService);
