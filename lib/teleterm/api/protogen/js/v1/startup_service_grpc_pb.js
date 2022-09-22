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
var v1_startup_service_pb = require('../v1/startup_service_pb.js');

function serialize_teleport_terminal_v1_ResolveTshdEventsServerAddressRequest(arg) {
  if (!(arg instanceof v1_startup_service_pb.ResolveTshdEventsServerAddressRequest)) {
    throw new Error('Expected argument of type teleport.terminal.v1.ResolveTshdEventsServerAddressRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_teleport_terminal_v1_ResolveTshdEventsServerAddressRequest(buffer_arg) {
  return v1_startup_service_pb.ResolveTshdEventsServerAddressRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_teleport_terminal_v1_ResolveTshdEventsServerAddressResponse(arg) {
  if (!(arg instanceof v1_startup_service_pb.ResolveTshdEventsServerAddressResponse)) {
    throw new Error('Expected argument of type teleport.terminal.v1.ResolveTshdEventsServerAddressResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_teleport_terminal_v1_ResolveTshdEventsServerAddressResponse(buffer_arg) {
  return v1_startup_service_pb.ResolveTshdEventsServerAddressResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// StartupService is used by the lib/teleterm gRPC server to collect data necessary to start the
// main TerminalService.
//
// This lets us start TerminalService handler with the tshd events client already available, without
// adding additional complexity to the handle of TerminalService.
var StartupServiceService = exports.StartupServiceService = {
  // ResolveTshdEventsServerAddress is called by the Electron app after the tshd events server has
// started.
resolveTshdEventsServerAddress: {
    path: '/teleport.terminal.v1.StartupService/ResolveTshdEventsServerAddress',
    requestStream: false,
    responseStream: false,
    requestType: v1_startup_service_pb.ResolveTshdEventsServerAddressRequest,
    responseType: v1_startup_service_pb.ResolveTshdEventsServerAddressResponse,
    requestSerialize: serialize_teleport_terminal_v1_ResolveTshdEventsServerAddressRequest,
    requestDeserialize: deserialize_teleport_terminal_v1_ResolveTshdEventsServerAddressRequest,
    responseSerialize: serialize_teleport_terminal_v1_ResolveTshdEventsServerAddressResponse,
    responseDeserialize: deserialize_teleport_terminal_v1_ResolveTshdEventsServerAddressResponse,
  },
};

exports.StartupServiceClient = grpc.makeGenericClientConstructor(StartupServiceService);
