// This file was auto-generated by the veyron vdl tool.
// Source: alarm.vdl

package sample

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Alarm allows clients to manipulate an alarm and query its status.
// Alarm is the interface the client binds and uses.
// Alarm_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Alarm_ExcludingUniversal interface {
	// Status returns the current status of the Alarm (i.e., armed, unarmed, panicking).
	Status(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error)
	// Arm sets the Alarm to the armed state.
	Arm(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// DelayArm sets the Alarm to the armed state after the given delay in seconds.
	DelayArm(ctx _gen_context.T, seconds uint16, opts ..._gen_ipc.CallOpt) (err error)
	// Unarm sets the Alarm to the unarmed state.
	Unarm(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Panic sets the Alarm to the panicking state.
	Panic(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
}
type Alarm interface {
	_gen_ipc.UniversalServiceMethods
	Alarm_ExcludingUniversal
}

// AlarmService is the interface the server implements.
type AlarmService interface {

	// Status returns the current status of the Alarm (i.e., armed, unarmed, panicking).
	Status(context _gen_ipc.ServerContext) (reply string, err error)
	// Arm sets the Alarm to the armed state.
	Arm(context _gen_ipc.ServerContext) (err error)
	// DelayArm sets the Alarm to the armed state after the given delay in seconds.
	DelayArm(context _gen_ipc.ServerContext, seconds uint16) (err error)
	// Unarm sets the Alarm to the unarmed state.
	Unarm(context _gen_ipc.ServerContext) (err error)
	// Panic sets the Alarm to the panicking state.
	Panic(context _gen_ipc.ServerContext) (err error)
}

// BindAlarm returns the client stub implementing the Alarm
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindAlarm(name string, opts ..._gen_ipc.BindOpt) (Alarm, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubAlarm{defaultClient: client, name: name}

	return stub, nil
}

// NewServerAlarm creates a new server stub.
//
// It takes a regular server implementing the AlarmService
// interface, and returns a new server stub.
func NewServerAlarm(server AlarmService) interface{} {
	return &ServerStubAlarm{
		service: server,
	}
}

// clientStubAlarm implements Alarm.
type clientStubAlarm struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubAlarm) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubAlarm) Status(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Status", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) Arm(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Arm", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) DelayArm(ctx _gen_context.T, seconds uint16, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "DelayArm", []interface{}{seconds}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) Unarm(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Unarm", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) Panic(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Panic", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAlarm) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubAlarm wraps a server that implements
// AlarmService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubAlarm struct {
	service AlarmService
}

func (__gen_s *ServerStubAlarm) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Status":
		return []interface{}{}, nil
	case "Arm":
		return []interface{}{}, nil
	case "DelayArm":
		return []interface{}{}, nil
	case "Unarm":
		return []interface{}{}, nil
	case "Panic":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubAlarm) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Arm"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["DelayArm"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "seconds", Type: 51},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Panic"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Status"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 3},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Unarm"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubAlarm) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubAlarm) Status(call _gen_ipc.ServerCall) (reply string, err error) {
	reply, err = __gen_s.service.Status(call)
	return
}

func (__gen_s *ServerStubAlarm) Arm(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Arm(call)
	return
}

func (__gen_s *ServerStubAlarm) DelayArm(call _gen_ipc.ServerCall, seconds uint16) (err error) {
	err = __gen_s.service.DelayArm(call, seconds)
	return
}

func (__gen_s *ServerStubAlarm) Unarm(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Unarm(call)
	return
}

func (__gen_s *ServerStubAlarm) Panic(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Panic(call)
	return
}
