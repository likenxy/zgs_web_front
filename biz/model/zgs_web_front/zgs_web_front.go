// Code generated by thriftgo (0.1.7). DO NOT EDIT.

package zgs_web_front

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type ZgsWebFront interface {
	HeartBeat(ctx context.Context) (r string, err error)
}

type ZgsWebFrontClient struct {
	c thrift.TClient
}

func NewZgsWebFrontClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ZgsWebFrontClient {
	return &ZgsWebFrontClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewZgsWebFrontClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ZgsWebFrontClient {
	return &ZgsWebFrontClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewZgsWebFrontClient(c thrift.TClient) *ZgsWebFrontClient {
	return &ZgsWebFrontClient{
		c: c,
	}
}

func (p *ZgsWebFrontClient) Client_() thrift.TClient {
	return p.c
}

func (p *ZgsWebFrontClient) HeartBeat(ctx context.Context) (r string, err error) {
	var _args ZgsWebFrontHeartBeatArgs
	var _result ZgsWebFrontHeartBeatResult
	if err = p.Client_().Call(ctx, "HeartBeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type ZgsWebFrontProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ZgsWebFront
}

func (p *ZgsWebFrontProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ZgsWebFrontProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ZgsWebFrontProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewZgsWebFrontProcessor(handler ZgsWebFront) *ZgsWebFrontProcessor {
	self := &ZgsWebFrontProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("HeartBeat", &zgsWebFrontProcessorHeartBeat{handler: handler})
	return self
}
func (p *ZgsWebFrontProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type zgsWebFrontProcessorHeartBeat struct {
	handler ZgsWebFront
}

func (p *zgsWebFrontProcessorHeartBeat) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ZgsWebFrontHeartBeatArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HeartBeat", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ZgsWebFrontHeartBeatResult{}
	var retval string
	if retval, err2 = p.handler.HeartBeat(ctx); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HeartBeat: "+err2.Error())
		oprot.WriteMessageBegin("HeartBeat", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("HeartBeat", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ZgsWebFrontHeartBeatArgs struct {
}

func NewZgsWebFrontHeartBeatArgs() *ZgsWebFrontHeartBeatArgs {
	return &ZgsWebFrontHeartBeatArgs{}
}

var fieldIDToName_ZgsWebFrontHeartBeatArgs = map[int16]string{}

func (p *ZgsWebFrontHeartBeatArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ZgsWebFrontHeartBeatArgs) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("HeartBeat_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ZgsWebFrontHeartBeatArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ZgsWebFrontHeartBeatArgs(%+v)", *p)
}

type ZgsWebFrontHeartBeatResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewZgsWebFrontHeartBeatResult() *ZgsWebFrontHeartBeatResult {
	return &ZgsWebFrontHeartBeatResult{}
}

var ZgsWebFrontHeartBeatResult_Success_DEFAULT string

func (p *ZgsWebFrontHeartBeatResult) GetSuccess() (v string) {
	if !p.IsSetSuccess() {
		return ZgsWebFrontHeartBeatResult_Success_DEFAULT
	}
	return *p.Success
}

var fieldIDToName_ZgsWebFrontHeartBeatResult = map[int16]string{
	0: "success",
}

func (p *ZgsWebFrontHeartBeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ZgsWebFrontHeartBeatResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ZgsWebFrontHeartBeatResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ZgsWebFrontHeartBeatResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Success = &v
	}
	return nil
}

func (p *ZgsWebFrontHeartBeatResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("HeartBeat_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ZgsWebFrontHeartBeatResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Success); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *ZgsWebFrontHeartBeatResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ZgsWebFrontHeartBeatResult(%+v)", *p)
}
