// Copyright (c) 2003-2005 Maxim Sobolev. All rights reserved.
// Copyright (c) 2006-2015 Sippy Software, Inc. All rights reserved.
// Copyright (c) 2015 Andrii Pylypenko. All rights reserved.
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
package sippy_header

import (
    "strconv"

    "sippy/conf"
)

type SipExpires struct {
    normalName
    Number int
}

var _sip_expires_name normalName = newNormalName("Expires")

func NewSipExpires() *SipExpires {
    return &SipExpires{
        normalName  : _sip_expires_name,
        Number      : 300,
    }
}

func ParseSipExpires(body string) ([]SipHeader, error) {
    number, err := strconv.Atoi(body)
    if err != nil {
        return nil, err
    }
    return []SipHeader{ &SipExpires{
        normalName  : _sip_expires_name,
        Number      : number,
    } }, nil
}

func (self *SipExpires) Body() string {
    return strconv.Itoa(self.Number)
}

func (self *SipExpires) String() string {
    return self.Name() + ": " + self.Body()
}

func (self *SipExpires) LocalStr(hostport *sippy_conf.HostPort, compact bool) string {
    return self.String()
}

func (self *SipExpires) GetCopy() *SipExpires {
    tmp := *self
    return &tmp
}

func (self *SipExpires) GetCopyAsIface() SipHeader {
    return self.GetCopy()
}
