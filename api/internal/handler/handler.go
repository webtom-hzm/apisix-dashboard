/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shiningrush/droplet/data"
)

type RegisterFactory func() (RouteRegister, error)

type RouteRegister interface {
	ApplyRoute(r *gin.Engine)
}

func SpecCodeResponse(err error) *data.SpecCodeResponse {
	errMsg := err.Error()
	if strings.Contains(errMsg, "required") ||
		strings.Contains(errMsg, "conflicted") ||
		strings.Contains(errMsg, "invalid") ||
		strings.Contains(errMsg, "missing") ||
		strings.Contains(errMsg, "scheme validate fail") {
		return &data.SpecCodeResponse{StatusCode: http.StatusBadRequest}
	}

	if strings.Contains(errMsg, "not found") {
		return &data.SpecCodeResponse{StatusCode: http.StatusNotFound}
	}

	return &data.SpecCodeResponse{StatusCode: http.StatusInternalServerError}
}
