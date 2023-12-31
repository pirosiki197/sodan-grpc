// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=ts"
// @generated from file api/v1/api.proto (package api.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CloseSodanRequest, CreateReplyRequest, CreateReplyResponse, CreateSodanRequest, CreateSodanResponse, GetRepliesRequest, GetRepliesResponse, GetReplyRequest, GetReplyResponse, GetSodanListResponse, GetSodanRequest, GetSodanResponse, GetSodansByTagRequest, GetSodansByTagResponse, SubscribeSodanRequest, SubscribeSodanResponse } from "./api_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.v1.APIService
 */
export const APIService = {
  typeName: "api.v1.APIService",
  methods: {
    /**
     * Sodan
     *
     * @generated from rpc api.v1.APIService.CreateSodan
     */
    createSodan: {
      name: "CreateSodan",
      I: CreateSodanRequest,
      O: CreateSodanResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.GetSodan
     */
    getSodan: {
      name: "GetSodan",
      I: GetSodanRequest,
      O: GetSodanResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.GetSodanList
     */
    getSodanList: {
      name: "GetSodanList",
      I: Empty,
      O: GetSodanListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.GetSodansByTag
     */
    getSodansByTag: {
      name: "GetSodansByTag",
      I: GetSodansByTagRequest,
      O: GetSodansByTagResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.CloseSodan
     */
    closeSodan: {
      name: "CloseSodan",
      I: CloseSodanRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * Chat
     *
     * @generated from rpc api.v1.APIService.CreateReply
     */
    createReply: {
      name: "CreateReply",
      I: CreateReplyRequest,
      O: CreateReplyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.GetReply
     */
    getReply: {
      name: "GetReply",
      I: GetReplyRequest,
      O: GetReplyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.GetReplies
     */
    getReplies: {
      name: "GetReplies",
      I: GetRepliesRequest,
      O: GetRepliesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.APIService.SubscribeSodan
     */
    subscribeSodan: {
      name: "SubscribeSodan",
      I: SubscribeSodanRequest,
      O: SubscribeSodanResponse,
      kind: MethodKind.ServerStreaming,
    },
  }
} as const;

