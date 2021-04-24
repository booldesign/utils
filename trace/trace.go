package trace

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/4/24 08:41
 * @Desc:
 */

var LocalIP = net.ParseIP("127.0.0.1")

type Trace struct {
	TraceId     string
	SpanId      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

func NewTrace() *Trace {
	trace := &Trace{}
	trace.TraceId = GetTraceId()
	trace.SpanId = NewSpanId()
	return trace
}

func NewSpanId() string {
	return generateSpanId(LocalIP.To4())
}

func GetTraceId() (traceId string) {
	return generateTraceId(LocalIP.String())
}

func generateSpanId(ip net.IP) string {
	timestamp := uint32(time.Now().Unix())
	ipToLong := binary.BigEndian.Uint32(ip)
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%08x", ipToLong^timestamp))
	b.WriteString(fmt.Sprintf("%08x", rand.Int31()))
	return b.String()
}

func generateTraceId(ip string) string {
	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	pid := os.Getpid()

	b := bytes.Buffer{}
	netIP := net.ParseIP(ip)
	if netIP == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}
	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))
	b.WriteString(fmt.Sprintf("%04x", pid&0xffff))
	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))
	b.WriteString("b0") // 末两位标记来源,b0为go

	return b.String()
}
