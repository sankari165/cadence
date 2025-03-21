// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package grpc

// Code generated by gowrap. DO NOT EDIT.
// template: ../../../templates/grpc.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	matchingv1 "github.com/uber/cadence/.gen/proto/matching/v1"
	"github.com/uber/cadence/common/types/mapper/proto"
	"github.com/uber/cadence/service/matching/handler"
)

type GRPCHandler struct {
	h handler.Handler
}

func NewGRPCHandler(h handler.Handler) GRPCHandler {
	return GRPCHandler{h}
}

func (g GRPCHandler) AddActivityTask(ctx context.Context, request *matchingv1.AddActivityTaskRequest) (*matchingv1.AddActivityTaskResponse, error) {
	response, err := g.h.AddActivityTask(ctx, proto.ToMatchingAddActivityTaskRequest(request))
	return proto.FromMatchingAddActivityTaskResponse(response), proto.FromError(err)
}

func (g GRPCHandler) AddDecisionTask(ctx context.Context, request *matchingv1.AddDecisionTaskRequest) (*matchingv1.AddDecisionTaskResponse, error) {
	response, err := g.h.AddDecisionTask(ctx, proto.ToMatchingAddDecisionTaskRequest(request))
	return proto.FromMatchingAddDecisionTaskResponse(response), proto.FromError(err)
}

func (g GRPCHandler) CancelOutstandingPoll(ctx context.Context, request *matchingv1.CancelOutstandingPollRequest) (*matchingv1.CancelOutstandingPollResponse, error) {
	err := g.h.CancelOutstandingPoll(ctx, proto.ToMatchingCancelOutstandingPollRequest(request))
	return &matchingv1.CancelOutstandingPollResponse{}, proto.FromError(err)
}

func (g GRPCHandler) DescribeTaskList(ctx context.Context, request *matchingv1.DescribeTaskListRequest) (*matchingv1.DescribeTaskListResponse, error) {
	response, err := g.h.DescribeTaskList(ctx, proto.ToMatchingDescribeTaskListRequest(request))
	return proto.FromMatchingDescribeTaskListResponse(response), proto.FromError(err)
}

func (g GRPCHandler) GetTaskListsByDomain(ctx context.Context, request *matchingv1.GetTaskListsByDomainRequest) (*matchingv1.GetTaskListsByDomainResponse, error) {
	response, err := g.h.GetTaskListsByDomain(ctx, proto.ToMatchingGetTaskListsByDomainRequest(request))
	return proto.FromMatchingGetTaskListsByDomainResponse(response), proto.FromError(err)
}

func (g GRPCHandler) ListTaskListPartitions(ctx context.Context, request *matchingv1.ListTaskListPartitionsRequest) (*matchingv1.ListTaskListPartitionsResponse, error) {
	response, err := g.h.ListTaskListPartitions(ctx, proto.ToMatchingListTaskListPartitionsRequest(request))
	return proto.FromMatchingListTaskListPartitionsResponse(response), proto.FromError(err)
}

func (g GRPCHandler) PollForActivityTask(ctx context.Context, request *matchingv1.PollForActivityTaskRequest) (*matchingv1.PollForActivityTaskResponse, error) {
	response, err := g.h.PollForActivityTask(ctx, proto.ToMatchingPollForActivityTaskRequest(request))
	return proto.FromMatchingPollForActivityTaskResponse(response), proto.FromError(err)
}

func (g GRPCHandler) PollForDecisionTask(ctx context.Context, request *matchingv1.PollForDecisionTaskRequest) (*matchingv1.PollForDecisionTaskResponse, error) {
	response, err := g.h.PollForDecisionTask(ctx, proto.ToMatchingPollForDecisionTaskRequest(request))
	return proto.FromMatchingPollForDecisionTaskResponse(response), proto.FromError(err)
}

func (g GRPCHandler) QueryWorkflow(ctx context.Context, request *matchingv1.QueryWorkflowRequest) (*matchingv1.QueryWorkflowResponse, error) {
	response, err := g.h.QueryWorkflow(ctx, proto.ToMatchingQueryWorkflowRequest(request))
	return proto.FromMatchingQueryWorkflowResponse(response), proto.FromError(err)
}

func (g GRPCHandler) RefreshTaskListPartitionConfig(ctx context.Context, request *matchingv1.RefreshTaskListPartitionConfigRequest) (*matchingv1.RefreshTaskListPartitionConfigResponse, error) {
	response, err := g.h.RefreshTaskListPartitionConfig(ctx, proto.ToMatchingRefreshTaskListPartitionConfigRequest(request))
	return proto.FromMatchingRefreshTaskListPartitionConfigResponse(response), proto.FromError(err)
}

func (g GRPCHandler) RespondQueryTaskCompleted(ctx context.Context, request *matchingv1.RespondQueryTaskCompletedRequest) (*matchingv1.RespondQueryTaskCompletedResponse, error) {
	err := g.h.RespondQueryTaskCompleted(ctx, proto.ToMatchingRespondQueryTaskCompletedRequest(request))
	return &matchingv1.RespondQueryTaskCompletedResponse{}, proto.FromError(err)
}

func (g GRPCHandler) UpdateTaskListPartitionConfig(ctx context.Context, request *matchingv1.UpdateTaskListPartitionConfigRequest) (*matchingv1.UpdateTaskListPartitionConfigResponse, error) {
	response, err := g.h.UpdateTaskListPartitionConfig(ctx, proto.ToMatchingUpdateTaskListPartitionConfigRequest(request))
	return proto.FromMatchingUpdateTaskListPartitionConfigResponse(response), proto.FromError(err)
}
