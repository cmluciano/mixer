// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package template

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	rpc "github.com/googleapis/googleapis/google/rpc"
	"github.com/hashicorp/go-multierror"

	"istio.io/api/mixer/v1/config/descriptor"
	"istio.io/mixer/pkg/adapter"
	adptConfig "istio.io/mixer/pkg/adapter/config"
	adptTmpl "istio.io/mixer/pkg/adapter/template"
	"istio.io/mixer/pkg/attribute"
	"istio.io/mixer/pkg/expr"
	"istio.io/mixer/pkg/status"

	"istio.io/mixer/template/sample/check"

	"istio.io/mixer/template/sample/quota"

	"istio.io/mixer/template/sample/report"
)

var (
	SupportedTmplInfo = map[string]Info{

		istio_mixer_adapter_sample_check.TemplateName: {
			CtrCfg:    &istio_mixer_adapter_sample_check.ConstructorParam{},
			Variety:   adptTmpl.TEMPLATE_VARIETY_CHECK,
			BldrName:  "istio.io/mixer/template/sample/check.SampleProcessorBuilder",
			HndlrName: "istio.io/mixer/template/sample/check.SampleProcessor",
			SupportsTemplate: func(hndlrBuilder adptConfig.HandlerBuilder) bool {
				_, ok := hndlrBuilder.(istio_mixer_adapter_sample_check.SampleProcessorBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adptConfig.Handler) bool {
				_, ok := hndlr.(istio_mixer_adapter_sample_check.SampleProcessor)
				return ok
			},
			InferType: func(cp proto.Message, tEvalFn TypeEvalFn) (proto.Message, error) {
				var err error = nil
				cpb := cp.(*istio_mixer_adapter_sample_check.ConstructorParam)
				infrdType := &istio_mixer_adapter_sample_check.Type{}

				infrdType.CheckExpression = istio_mixer_v1_config_descriptor.STRING

				_ = cpb
				return infrdType, err
			},
			ConfigureType: func(types map[string]proto.Message, builder *adptConfig.HandlerBuilder) error {
				// Mixer framework should have ensured the type safety.
				castedBuilder := (*builder).(istio_mixer_adapter_sample_check.SampleProcessorBuilder)
				castedTypes := make(map[string]*istio_mixer_adapter_sample_check.Type)
				for k, v := range types {
					// Mixer framework should have ensured the type safety.
					v1 := v.(*istio_mixer_adapter_sample_check.Type)
					castedTypes[k] = v1
				}
				return castedBuilder.ConfigureSample(castedTypes)
			},

			ProcessCheck: func(ctrs map[string]proto.Message, attrs attribute.Bag, mapper expr.Evaluator,
				handler adptConfig.Handler) (rpc.Status, adptConfig.CacheabilityInfo) {
				var found bool
				var err error

				var instances []*istio_mixer_adapter_sample_check.Instance
				castedCnstrs := make(map[string]*istio_mixer_adapter_sample_check.ConstructorParam)
				for k, v := range ctrs {
					v1 := v.(*istio_mixer_adapter_sample_check.ConstructorParam)
					castedCnstrs[k] = v1
				}
				for name, md := range castedCnstrs {

					CheckExpression, err := mapper.Eval(md.CheckExpression, attrs)

					if err != nil {
						return status.WithError(err), adptConfig.CacheabilityInfo{}
					}

					instances = append(instances, &istio_mixer_adapter_sample_check.Instance{
						Name: name,

						CheckExpression: CheckExpression.(string),
					})
				}
				var cacheInfo adptConfig.CacheabilityInfo
				if found, cacheInfo, err = handler.(istio_mixer_adapter_sample_check.SampleProcessor).CheckSample(instances); err != nil {
					return status.WithError(err), adptConfig.CacheabilityInfo{}
				}

				if found {
					return status.OK, cacheInfo
				}

				return status.WithPermissionDenied(fmt.Sprintf("%s rejected", instances)), adptConfig.CacheabilityInfo{}
			},
			ProcessReport: nil,
			ProcessQuota:  nil,
		},

		istio_mixer_adapter_sample_quota.TemplateName: {
			CtrCfg:    &istio_mixer_adapter_sample_quota.ConstructorParam{},
			Variety:   adptTmpl.TEMPLATE_VARIETY_QUOTA,
			BldrName:  "istio.io/mixer/template/sample/quota.QuotaProcessorBuilder",
			HndlrName: "istio.io/mixer/template/sample/quota.QuotaProcessor",
			SupportsTemplate: func(hndlrBuilder adptConfig.HandlerBuilder) bool {
				_, ok := hndlrBuilder.(istio_mixer_adapter_sample_quota.QuotaProcessorBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adptConfig.Handler) bool {
				_, ok := hndlr.(istio_mixer_adapter_sample_quota.QuotaProcessor)
				return ok
			},
			InferType: func(cp proto.Message, tEvalFn TypeEvalFn) (proto.Message, error) {
				var err error = nil
				cpb := cp.(*istio_mixer_adapter_sample_quota.ConstructorParam)
				infrdType := &istio_mixer_adapter_sample_quota.Type{}

				infrdType.Dimensions = make(map[string]istio_mixer_v1_config_descriptor.ValueType)
				for k, v := range cpb.Dimensions {
					if infrdType.Dimensions[k], err = tEvalFn(v); err != nil {
						return nil, err
					}
				}

				_ = cpb
				return infrdType, err
			},
			ConfigureType: func(types map[string]proto.Message, builder *adptConfig.HandlerBuilder) error {
				// Mixer framework should have ensured the type safety.
				castedBuilder := (*builder).(istio_mixer_adapter_sample_quota.QuotaProcessorBuilder)
				castedTypes := make(map[string]*istio_mixer_adapter_sample_quota.Type)
				for k, v := range types {
					// Mixer framework should have ensured the type safety.
					v1 := v.(*istio_mixer_adapter_sample_quota.Type)
					castedTypes[k] = v1
				}
				return castedBuilder.ConfigureQuota(castedTypes)
			},

			ProcessQuota: func(quotaName string, cnstr proto.Message, attrs attribute.Bag, mapper expr.Evaluator, handler adptConfig.Handler,
				qma adapter.QuotaRequestArgs) (rpc.Status, adptConfig.CacheabilityInfo, adapter.QuotaResult) {
				castedCnstr := cnstr.(*istio_mixer_adapter_sample_quota.ConstructorParam)

				Dimensions, err := evalAll(castedCnstr.Dimensions, attrs, mapper)

				if err != nil {
					msg := fmt.Sprintf("failed to eval Dimensions for constructor '%s': %v", quotaName, err)
					glog.Error(msg)
					return status.WithInvalidArgument(msg), adptConfig.CacheabilityInfo{}, adapter.QuotaResult{}
				}

				instance := &istio_mixer_adapter_sample_quota.Instance{
					Name: quotaName,

					Dimensions: Dimensions,
				}

				var qr adapter.QuotaResult
				var cacheInfo adptConfig.CacheabilityInfo
				if qr, cacheInfo, err = handler.(istio_mixer_adapter_sample_quota.QuotaProcessor).AllocQuota(instance, qma); err != nil {
					glog.Errorf("Quota allocation failed: %v", err)
					return status.WithError(err), adptConfig.CacheabilityInfo{}, adapter.QuotaResult{}
				}
				if qr.Amount == 0 {
					msg := fmt.Sprintf("Unable to allocate %v units from quota %s", qma.QuotaAmount, quotaName)
					glog.Warning(msg)
					return status.WithResourceExhausted(msg), adptConfig.CacheabilityInfo{}, adapter.QuotaResult{}
				}
				if glog.V(2) {
					glog.Infof("Allocated %v units from quota %s", qma.QuotaAmount, quotaName)
				}
				return status.OK, cacheInfo, qr
			},
			ProcessReport: nil,
			ProcessCheck:  nil,
		},

		istio_mixer_adapter_sample_report.TemplateName: {
			CtrCfg:    &istio_mixer_adapter_sample_report.ConstructorParam{},
			Variety:   adptTmpl.TEMPLATE_VARIETY_REPORT,
			BldrName:  "istio.io/mixer/template/sample/report.SampleProcessorBuilder",
			HndlrName: "istio.io/mixer/template/sample/report.SampleProcessor",
			SupportsTemplate: func(hndlrBuilder adptConfig.HandlerBuilder) bool {
				_, ok := hndlrBuilder.(istio_mixer_adapter_sample_report.SampleProcessorBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adptConfig.Handler) bool {
				_, ok := hndlr.(istio_mixer_adapter_sample_report.SampleProcessor)
				return ok
			},
			InferType: func(cp proto.Message, tEvalFn TypeEvalFn) (proto.Message, error) {
				var err error = nil
				cpb := cp.(*istio_mixer_adapter_sample_report.ConstructorParam)
				infrdType := &istio_mixer_adapter_sample_report.Type{}

				if infrdType.Value, err = tEvalFn(cpb.Value); err != nil {
					return nil, err
				}

				infrdType.Dimensions = make(map[string]istio_mixer_v1_config_descriptor.ValueType)
				for k, v := range cpb.Dimensions {
					if infrdType.Dimensions[k], err = tEvalFn(v); err != nil {
						return nil, err
					}
				}

				_ = cpb
				return infrdType, err
			},
			ConfigureType: func(types map[string]proto.Message, builder *adptConfig.HandlerBuilder) error {
				// Mixer framework should have ensured the type safety.
				castedBuilder := (*builder).(istio_mixer_adapter_sample_report.SampleProcessorBuilder)
				castedTypes := make(map[string]*istio_mixer_adapter_sample_report.Type)
				for k, v := range types {
					// Mixer framework should have ensured the type safety.
					v1 := v.(*istio_mixer_adapter_sample_report.Type)
					castedTypes[k] = v1
				}
				return castedBuilder.ConfigureSample(castedTypes)
			},

			ProcessReport: func(ctrs map[string]proto.Message, attrs attribute.Bag, mapper expr.Evaluator, handler adptConfig.Handler) rpc.Status {
				result := &multierror.Error{}
				var instances []*istio_mixer_adapter_sample_report.Instance

				castedCnstrs := make(map[string]*istio_mixer_adapter_sample_report.ConstructorParam)
				for k, v := range ctrs {
					v1 := v.(*istio_mixer_adapter_sample_report.ConstructorParam)
					castedCnstrs[k] = v1
				}
				for name, md := range castedCnstrs {

					Value, err := mapper.Eval(md.Value, attrs)

					if err != nil {
						result = multierror.Append(result, fmt.Errorf("failed to eval Value for constructor '%s': %v", name, err))
						continue
					}

					Dimensions, err := evalAll(md.Dimensions, attrs, mapper)

					if err != nil {
						result = multierror.Append(result, fmt.Errorf("failed to eval Dimensions for constructor '%s': %v", name, err))
						continue
					}

					instances = append(instances, &istio_mixer_adapter_sample_report.Instance{
						Name: name,

						Value: Value,

						Dimensions: Dimensions,
					})
				}

				if err := handler.(istio_mixer_adapter_sample_report.SampleProcessor).ReportSample(instances); err != nil {
					result = multierror.Append(result, fmt.Errorf("failed to report all values: %v", err))
				}

				err := result.ErrorOrNil()
				if err != nil {
					return status.WithError(err)
				}

				return status.OK
			},
			ProcessCheck: nil,
			ProcessQuota: nil,
		},
	}
)
