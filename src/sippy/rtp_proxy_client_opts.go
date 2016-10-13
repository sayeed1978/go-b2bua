// Copyright (c) 2003-2005 Maxim Sobolev. All rights reserved.
// Copyright (c) 2006-2014 Sippy Software, Inc. All rights reserved.
// Copyright (c) 2016 Andriy Pylypenko. All rights reserved.
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation and/or
// other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
// ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package sippy

import (
    "time"

    "sippy/conf"
)

type rtpProxyClientOpts struct {
    no_version_check    bool
    spath               string
    nworkers            *int
    bind_address        *sippy_conf.HostPort
    hrtb_retr_ival      time.Duration
    hrtb_ival           time.Duration
}

func NewRtpProxyClientOpts() *rtpProxyClientOpts {
    return &rtpProxyClientOpts{
        hrtb_retr_ival      : 60 * time.Second,
        hrtb_ival           : 10 * time.Second,
        no_version_check    : false,
    }
}

func (self *rtpProxyClientOpts) SetSocketPath(spath string) {
    self.spath = spath
}

func (self *rtpProxyClientOpts) SetHeartbeatInterval(ival time.Duration) {
    self.hrtb_ival = ival
}

func (self *rtpProxyClientOpts) SetHeartbeatRetryInterval(ival time.Duration) {
    self.hrtb_retr_ival = ival
}

func (self *rtpProxyClientOpts) GetNWorkers() *int {
    return self.nworkers
}

func (self *rtpProxyClientOpts) GetBindAddress() *sippy_conf.HostPort {
    return self.bind_address
}
