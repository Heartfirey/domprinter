// Code generated by hertz generator.

package domprinter_service

import (
	"context"
	"fmt"

	domprinter "github.com/Dup4/domprinter/hertz_client/domprinter"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	FetchPrintTask(context context.Context, req *domprinter.FetchPrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.FetchPrintTaskResp, rawResponse *protocol.Response, err error)

	SubmitPrintTask(context context.Context, req *domprinter.SubmitPrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.SubmitPrintTaskResp, rawResponse *protocol.Response, err error)

	UpdatePrintTask(context context.Context, req *domprinter.UpdatePrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.UpdatePrintTaskResp, rawResponse *protocol.Response, err error)
}

type DOMPrinterServiceClient struct {
	client *cli
}

func NewDOMPrinterServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &DOMPrinterServiceClient{
		client: cli,
	}, nil
}

func (s *DOMPrinterServiceClient) FetchPrintTask(context context.Context, req *domprinter.FetchPrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.FetchPrintTaskResp, rawResponse *protocol.Response, err error) {
	httpResp := &domprinter.FetchPrintTaskResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"TaskState":    req.GetTaskState(),
			"OffsetTaskID": req.GetOffsetTaskID(),
			"LimitTaskNum": req.GetLimitTaskNum(),
			"BaseReq":      req.GetBaseReq(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/print-task")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *DOMPrinterServiceClient) SubmitPrintTask(context context.Context, req *domprinter.SubmitPrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.SubmitPrintTaskResp, rawResponse *protocol.Response, err error) {
	httpResp := &domprinter.SubmitPrintTaskResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/print-task")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *DOMPrinterServiceClient) UpdatePrintTask(context context.Context, req *domprinter.UpdatePrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.UpdatePrintTaskResp, rawResponse *protocol.Response, err error) {
	httpResp := &domprinter.UpdatePrintTaskResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("PATCH", "/print-task")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewDOMPrinterServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewDOMPrinterServiceClient("", ops...)
	return
}

func FetchPrintTask(context context.Context, req *domprinter.FetchPrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.FetchPrintTaskResp, rawResponse *protocol.Response, err error) {
	return defaultClient.FetchPrintTask(context, req, reqOpt...)
}

func SubmitPrintTask(context context.Context, req *domprinter.SubmitPrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.SubmitPrintTaskResp, rawResponse *protocol.Response, err error) {
	return defaultClient.SubmitPrintTask(context, req, reqOpt...)
}

func UpdatePrintTask(context context.Context, req *domprinter.UpdatePrintTaskReq, reqOpt ...config.RequestOption) (resp *domprinter.UpdatePrintTaskResp, rawResponse *protocol.Response, err error) {
	return defaultClient.UpdatePrintTask(context, req, reqOpt...)
}