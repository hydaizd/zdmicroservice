package app

import (
	"fmt"
	"github.com/go-chassis/go-chassis/core/registry"
	rf "github.com/go-chassis/go-chassis/server/restful"
	"net/http"
	"strconv"
	"time"
)

/*
 * 具体业务逻辑
 */

// 定义服务结构体
type CalculateBmi struct {
}

// 指定cse对应的 URL 路由
func (m *CalculateBmi) URLPatterns() []rf.Route {
	return []rf.Route{
		{
			Method:           http.MethodGet,
			Path:             "/calculator/bmi",
			ResourceFuncName: "Calculate",
		},
	}
}

// 编写 handler 函数，其中restful.Context必须作为入参传入
func (m *CalculateBmi) Calculate(ctx *rf.Context) {
	var height, weight, bmi float64
	var err error
	result := struct {
		Result     float64 `json:"result"`
		InstanceId string  `json:"instanceId"`
		CallTime   string  `json:"callTime"`
	}{}
	errorResponse := struct {
		Error string `json:"error"`
	}{}

	heightStr := ctx.ReadQueryParameter("height")
	weightStr := ctx.ReadQueryParameter("weight")

	if height, err = strconv.ParseFloat(heightStr, 10); err != nil {
		errorResponse.Error = err.Error()
		ctx.WriteHeaderAndJSON(http.StatusBadRequest, errorResponse, "application/json")
		return
	}
	if weight, err = strconv.ParseFloat(weightStr, 10); err != nil {
		errorResponse.Error = err.Error()
		ctx.WriteHeaderAndJSON(http.StatusBadRequest, errorResponse, "application/json")
		return
	}

	if bmi, err = m.BMIIndex(height, weight); err != nil {
		errorResponse.Error = err.Error()
		ctx.WriteHeaderAndJSON(http.StatusBadRequest, errorResponse, "application/json")
		return
	}

	result.Result = bmi
	// Get InstanceID（为了便于区分不同的运行实例，在体质指数计算器（calculator）的实现中新增了返回实例 ID 的代码。）
	items := registry.SelfInstancesCache.Items()
	for microserviceID := range items {
		instanceID, exist := registry.SelfInstancesCache.Get(microserviceID)
		if exist {
			result.InstanceId = instanceID.([]string)[0]
		}
	}
	result.CallTime = time.Now().String()
	ctx.WriteJSON(result, "application/json")
}

// 编写计算体质指数（BMI）函数，该函数根据公式 \(BMI=\frac{weight}{height^2}\)进行实现
func (m *CalculateBmi) BMIIndex(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, fmt.Errorf("Arugments must be above 0")
	}
	heightInMeter := height / 100
	bmi := weight / (heightInMeter * heightInMeter)
	return bmi, nil
}
