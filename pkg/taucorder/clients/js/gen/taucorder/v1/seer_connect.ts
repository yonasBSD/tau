// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file taucorder/v1/seer.proto (package taucorder.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { LocationRequest, NodesListRequest, NodesUsageRequest, PeerLocation, PeerUsage } from "./seer_pb.js";
import { Peer } from "./common_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Service
 *
 * @generated from service taucorder.v1.SeerService
 */
export const SeerService = {
  typeName: "taucorder.v1.SeerService",
  methods: {
    /**
     * @generated from rpc taucorder.v1.SeerService.List
     */
    list: {
      name: "List",
      I: NodesListRequest,
      O: Peer,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc taucorder.v1.SeerService.Usage
     */
    usage: {
      name: "Usage",
      I: NodesUsageRequest,
      O: PeerUsage,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc taucorder.v1.SeerService.Location
     */
    location: {
      name: "Location",
      I: LocationRequest,
      O: PeerLocation,
      kind: MethodKind.ServerStreaming,
    },
  }
} as const;
