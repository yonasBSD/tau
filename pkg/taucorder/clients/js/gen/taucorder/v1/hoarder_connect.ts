// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file taucorder/v1/hoarder.proto (package taucorder.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, Node } from "./common_pb.js";
import { StashedItem, StashRequest } from "./hoarder_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Service
 *
 * @generated from service taucorder.v1.HoarderService
 */
export const HoarderService = {
  typeName: "taucorder.v1.HoarderService",
  methods: {
    /**
     * @generated from rpc taucorder.v1.HoarderService.List
     */
    list: {
      name: "List",
      I: Node,
      O: StashedItem,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc taucorder.v1.HoarderService.Stash
     */
    stash: {
      name: "Stash",
      I: StashRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;
