// This file was auto-generated by the veyron vdl tool.
// Source: speaker.vdl

package sample

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_io "io"
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

// Speaker allows clients to control the music being played.
// Speaker is the interface the client binds and uses.
// Speaker_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Speaker_ExcludingUniversal interface {
	// Play starts or continues the current song.
	Play(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// PlaySong plays back the given song title, if possible.
	PlaySong(ctx _gen_context.T, songName string, opts ..._gen_ipc.CallOpt) (err error)
	// PlayStream plays the given stream of music data.
	PlayStream(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply SpeakerPlayStreamCall, err error)
	// GetSong retrieves the title of the Speaker's current song, if any.
	GetSong(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error)
	// Pause playback of the Speaker's current song.
	Pause(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Stop playback of the Speaker's current song.
	Stop(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Volume adjusts the Speaker's volume.
	Volume(ctx _gen_context.T, volumeLevel uint16, opts ..._gen_ipc.CallOpt) (err error)
	// GetVolume retrieves the Speaker's volume.
	GetVolume(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply uint16, err error)
}
type Speaker interface {
	_gen_ipc.UniversalServiceMethods
	Speaker_ExcludingUniversal
}

// SpeakerService is the interface the server implements.
type SpeakerService interface {

	// Play starts or continues the current song.
	Play(context _gen_ipc.ServerContext) (err error)
	// PlaySong plays back the given song title, if possible.
	PlaySong(context _gen_ipc.ServerContext, songName string) (err error)
	// PlayStream plays the given stream of music data.
	PlayStream(context _gen_ipc.ServerContext, stream SpeakerServicePlayStreamStream) (err error)
	// GetSong retrieves the title of the Speaker's current song, if any.
	GetSong(context _gen_ipc.ServerContext) (reply string, err error)
	// Pause playback of the Speaker's current song.
	Pause(context _gen_ipc.ServerContext) (err error)
	// Stop playback of the Speaker's current song.
	Stop(context _gen_ipc.ServerContext) (err error)
	// Volume adjusts the Speaker's volume.
	Volume(context _gen_ipc.ServerContext, volumeLevel uint16) (err error)
	// GetVolume retrieves the Speaker's volume.
	GetVolume(context _gen_ipc.ServerContext) (reply uint16, err error)
}

// SpeakerPlayStreamCall is the interface for call object of the method
// PlayStream in the service interface Speaker.
type SpeakerPlayStreamCall interface {

	// SendStream returns the send portion of the stream
	SendStream() interface {
		// Send places the item onto the output stream, blocking if there is no
		// buffer space available.  Calls to Send after having called Close
		// or Cancel will fail.  Any blocked Send calls will be unblocked upon
		// calling Cancel.
		Send(item []byte) error

		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.  This is
		// an optional call - it's used by streaming clients that need the
		// server to receive the io.EOF terminator before the client calls
		// Finish (for example, if the client needs to continue receiving items
		// from the server after having finished sending).
		// Calls to Close after having called Cancel will fail.
		// Like Send, Close blocks when there's no buffer space available.
		Close() error
	}

	// Finish performs the equivalent of SendStream().Close, then blocks until the server
	// is done, and returns the positional return values for call.
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.
	// Finish should be called at most once.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implSpeakerPlayStreamStreamSender struct {
	clientCall _gen_ipc.Call
}

func (c *implSpeakerPlayStreamStreamSender) Send(item []byte) error {
	return c.clientCall.Send(item)
}

func (c *implSpeakerPlayStreamStreamSender) Close() error {
	return c.clientCall.CloseSend()
}

// Implementation of the SpeakerPlayStreamCall interface that is not exported.
type implSpeakerPlayStreamCall struct {
	clientCall  _gen_ipc.Call
	writeStream implSpeakerPlayStreamStreamSender
}

func (c *implSpeakerPlayStreamCall) SendStream() interface {
	Send(item []byte) error
	Close() error
} {
	return &c.writeStream
}

func (c *implSpeakerPlayStreamCall) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implSpeakerPlayStreamCall) Cancel() {
	c.clientCall.Cancel()
}

type implSpeakerServicePlayStreamStreamIterator struct {
	serverCall _gen_ipc.ServerCall
	val        []byte
	err        error
}

func (s *implSpeakerServicePlayStreamStreamIterator) Advance() bool {
	s.err = s.serverCall.Recv(&s.val)
	return s.err == nil
}

func (s *implSpeakerServicePlayStreamStreamIterator) Value() []byte {
	return s.val
}

func (s *implSpeakerServicePlayStreamStreamIterator) Err() error {
	if s.err == _gen_io.EOF {
		return nil
	}
	return s.err
}

// SpeakerServicePlayStreamStream is the interface for streaming responses of the method
// PlayStream in the service interface Speaker.
type SpeakerServicePlayStreamStream interface {
	// RecvStream returns the recv portion of the stream
	RecvStream() interface {
		// Advance stages an element so the client can retrieve it
		// with Value.  Advance returns true iff there is an
		// element to retrieve.  The client must call Advance before
		// calling Value.  Advance may block if an element is not
		// immediately available.
		Advance() bool

		// Value returns the element that was staged by Advance.
		// Value may panic if Advance returned false or was not
		// called at all.  Value does not block.
		Value() []byte

		// Err returns a non-nil error iff the stream encountered
		// any errors.  Err does not block.
		Err() error
	}
}

// Implementation of the SpeakerServicePlayStreamStream interface that is not exported.
type implSpeakerServicePlayStreamStream struct {
	reader implSpeakerServicePlayStreamStreamIterator
}

func (s *implSpeakerServicePlayStreamStream) RecvStream() interface {
	// Advance stages an element so the client can retrieve it
	// with Value.  Advance returns true iff there is an
	// element to retrieve.  The client must call Advance before
	// calling Value.  The client must call Cancel if it does
	// not iterate through all elements (i.e. until Advance
	// returns false).  Advance may block if an element is not
	// immediately available.
	Advance() bool

	// Value returns the element that was staged by Advance.
	// Value may panic if Advance returned false or was not
	// called at all.  Value does not block.
	Value() []byte

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error
} {
	return &s.reader
}

// BindSpeaker returns the client stub implementing the Speaker
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindSpeaker(name string, opts ..._gen_ipc.BindOpt) (Speaker, error) {
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
	stub := &clientStubSpeaker{defaultClient: client, name: name}

	return stub, nil
}

// NewServerSpeaker creates a new server stub.
//
// It takes a regular server implementing the SpeakerService
// interface, and returns a new server stub.
func NewServerSpeaker(server SpeakerService) interface{} {
	return &ServerStubSpeaker{
		service: server,
	}
}

// clientStubSpeaker implements Speaker.
type clientStubSpeaker struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubSpeaker) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubSpeaker) Play(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Play", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) PlaySong(ctx _gen_context.T, songName string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "PlaySong", []interface{}{songName}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) PlayStream(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply SpeakerPlayStreamCall, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "PlayStream", nil, opts...); err != nil {
		return
	}
	reply = &implSpeakerPlayStreamCall{clientCall: call, writeStream: implSpeakerPlayStreamStreamSender{clientCall: call}}
	return
}

func (__gen_c *clientStubSpeaker) GetSong(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetSong", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) Pause(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Pause", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) Stop(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Stop", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) Volume(ctx _gen_context.T, volumeLevel uint16, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Volume", []interface{}{volumeLevel}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) GetVolume(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply uint16, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetVolume", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSpeaker) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubSpeaker wraps a server that implements
// SpeakerService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubSpeaker struct {
	service SpeakerService
}

func (__gen_s *ServerStubSpeaker) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Play":
		return []interface{}{}, nil
	case "PlaySong":
		return []interface{}{}, nil
	case "PlayStream":
		return []interface{}{}, nil
	case "GetSong":
		return []interface{}{}, nil
	case "Pause":
		return []interface{}{}, nil
	case "Stop":
		return []interface{}{}, nil
	case "Volume":
		return []interface{}{}, nil
	case "GetVolume":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubSpeaker) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["GetSong"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 3},
			{Name: "", Type: 65},
		},
	}
	result.Methods["GetVolume"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 51},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Pause"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Play"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["PlaySong"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "songName", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["PlayStream"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
		InStream: 67,
	}
	result.Methods["Stop"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Volume"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "volumeLevel", Type: 51},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubSpeaker) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubSpeaker) Play(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Play(call)
	return
}

func (__gen_s *ServerStubSpeaker) PlaySong(call _gen_ipc.ServerCall, songName string) (err error) {
	err = __gen_s.service.PlaySong(call, songName)
	return
}

func (__gen_s *ServerStubSpeaker) PlayStream(call _gen_ipc.ServerCall) (err error) {
	stream := &implSpeakerServicePlayStreamStream{reader: implSpeakerServicePlayStreamStreamIterator{serverCall: call}}
	err = __gen_s.service.PlayStream(call, stream)
	return
}

func (__gen_s *ServerStubSpeaker) GetSong(call _gen_ipc.ServerCall) (reply string, err error) {
	reply, err = __gen_s.service.GetSong(call)
	return
}

func (__gen_s *ServerStubSpeaker) Pause(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Pause(call)
	return
}

func (__gen_s *ServerStubSpeaker) Stop(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Stop(call)
	return
}

func (__gen_s *ServerStubSpeaker) Volume(call _gen_ipc.ServerCall, volumeLevel uint16) (err error) {
	err = __gen_s.service.Volume(call, volumeLevel)
	return
}

func (__gen_s *ServerStubSpeaker) GetVolume(call _gen_ipc.ServerCall) (reply uint16, err error) {
	reply, err = __gen_s.service.GetVolume(call)
	return
}
