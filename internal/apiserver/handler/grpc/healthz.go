// Copyright 2025 码龙 <maloong2022@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/maloong2022/miniblog. The professional
// version of this repository is https://github.com/maloong2022/onex.

package grpc

import (
	"context"
	"github.com/maloong2022/miniblog/internal/pkg/log"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	apiv1 "github.com/maloong2022/miniblog/pkg/api/apiserver/v1"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(ctx context.Context, rq *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	log.W(ctx).Infow("Healthz handler is called", "method", "Healthz", "status", "healthy")
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
