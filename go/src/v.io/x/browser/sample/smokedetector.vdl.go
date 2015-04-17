// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: smokedetector.vdl

package sample

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

// SmokeDetectorClientMethods is the client interface
// containing SmokeDetector methods.
//
// SmokeDetector allows clients to monitor and adjust a smoke detector.
type SmokeDetectorClientMethods interface {
	// Status retrieves the current status and sensitivity of the SmokeDetector.
	Status(*context.T, ...rpc.CallOpt) (status string, sensitivity int16, err error)
	// Test the SmokeDetector to check if it is working.
	Test(*context.T, ...rpc.CallOpt) (bool, error)
	// Sensitivity adjusts the SmokeDetector's sensitivity to smoke.
	Sensitivity(ctx *context.T, sens int16, opts ...rpc.CallOpt) error
}

// SmokeDetectorClientStub adds universal methods to SmokeDetectorClientMethods.
type SmokeDetectorClientStub interface {
	SmokeDetectorClientMethods
	rpc.UniversalServiceMethods
}

// SmokeDetectorClient returns a client stub for SmokeDetector.
func SmokeDetectorClient(name string) SmokeDetectorClientStub {
	return implSmokeDetectorClientStub{name}
}

type implSmokeDetectorClientStub struct {
	name string
}

func (c implSmokeDetectorClientStub) Status(ctx *context.T, opts ...rpc.CallOpt) (o0 string, o1 int16, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Status", nil, []interface{}{&o0, &o1}, opts...)
	return
}

func (c implSmokeDetectorClientStub) Test(ctx *context.T, opts ...rpc.CallOpt) (o0 bool, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Test", nil, []interface{}{&o0}, opts...)
	return
}

func (c implSmokeDetectorClientStub) Sensitivity(ctx *context.T, i0 int16, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Sensitivity", []interface{}{i0}, nil, opts...)
	return
}

// SmokeDetectorServerMethods is the interface a server writer
// implements for SmokeDetector.
//
// SmokeDetector allows clients to monitor and adjust a smoke detector.
type SmokeDetectorServerMethods interface {
	// Status retrieves the current status and sensitivity of the SmokeDetector.
	Status(*context.T, rpc.ServerCall) (status string, sensitivity int16, err error)
	// Test the SmokeDetector to check if it is working.
	Test(*context.T, rpc.ServerCall) (bool, error)
	// Sensitivity adjusts the SmokeDetector's sensitivity to smoke.
	Sensitivity(ctx *context.T, call rpc.ServerCall, sens int16) error
}

// SmokeDetectorServerStubMethods is the server interface containing
// SmokeDetector methods, as expected by rpc.Server.
// There is no difference between this interface and SmokeDetectorServerMethods
// since there are no streaming methods.
type SmokeDetectorServerStubMethods SmokeDetectorServerMethods

// SmokeDetectorServerStub adds universal methods to SmokeDetectorServerStubMethods.
type SmokeDetectorServerStub interface {
	SmokeDetectorServerStubMethods
	// Describe the SmokeDetector interfaces.
	Describe__() []rpc.InterfaceDesc
}

// SmokeDetectorServer returns a server stub for SmokeDetector.
// It converts an implementation of SmokeDetectorServerMethods into
// an object that may be used by rpc.Server.
func SmokeDetectorServer(impl SmokeDetectorServerMethods) SmokeDetectorServerStub {
	stub := implSmokeDetectorServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implSmokeDetectorServerStub struct {
	impl SmokeDetectorServerMethods
	gs   *rpc.GlobState
}

func (s implSmokeDetectorServerStub) Status(ctx *context.T, call rpc.ServerCall) (string, int16, error) {
	return s.impl.Status(ctx, call)
}

func (s implSmokeDetectorServerStub) Test(ctx *context.T, call rpc.ServerCall) (bool, error) {
	return s.impl.Test(ctx, call)
}

func (s implSmokeDetectorServerStub) Sensitivity(ctx *context.T, call rpc.ServerCall, i0 int16) error {
	return s.impl.Sensitivity(ctx, call, i0)
}

func (s implSmokeDetectorServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implSmokeDetectorServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{SmokeDetectorDesc}
}

// SmokeDetectorDesc describes the SmokeDetector interface.
var SmokeDetectorDesc rpc.InterfaceDesc = descSmokeDetector

// descSmokeDetector hides the desc to keep godoc clean.
var descSmokeDetector = rpc.InterfaceDesc{
	Name:    "SmokeDetector",
	PkgPath: "v.io/x/browser/sample",
	Doc:     "// SmokeDetector allows clients to monitor and adjust a smoke detector.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Status",
			Doc:  "// Status retrieves the current status and sensitivity of the SmokeDetector. ",
			OutArgs: []rpc.ArgDesc{
				{"status", ``},      // string
				{"sensitivity", ``}, // int16
			},
		},
		{
			Name: "Test",
			Doc:  "// Test the SmokeDetector to check if it is working.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // bool
			},
		},
		{
			Name: "Sensitivity",
			Doc:  "// Sensitivity adjusts the SmokeDetector's sensitivity to smoke.",
			InArgs: []rpc.ArgDesc{
				{"sens", ``}, // int16
			},
		},
	},
}