package main

import (
	"bytes"
	"testing"
)

func TestParseCompact(t *testing.T) {
	procNetSnmp := `Ip: Forwarding DefaultTTL InReceives InHdrErrors InAddrErrors ForwDatagrams InUnknownProtos InDiscards InDelivers OutRequests OutDiscards OutNoRoutes ReasmTimeout ReasmReqds ReasmOKs ReasmFails FragOKs FragFails FragCreates
Ip: 1 64 3000945 0 0 5563 0 0 1587382 1341624 0 90 0 0 0 0 0 0 0
Icmp: InMsgs InErrors InCsumErrors InDestUnreachs InTimeExcds InParmProbs InSrcQuenchs InRedirects InEchos InEchoReps InTimestamps InTimestampReps InAddrMasks InAddrMaskReps OutMsgs OutErrors OutDestUnreachs OutTimeExcds OutParmProbs OutSrcQuenchs OutRedirects OutEchos OutEchoReps OutTimestamps OutTimestampReps OutAddrMasks OutAddrMaskReps
Icmp: 1010 754 3 755 81 0 1 2 159 9 0 0 0 0 8128 0 7969 0 0 0 0 0 159 0 0 0 0
IcmpMsg: InType0 InType3 InType4 InType5 InType8 InType11 OutType0 OutType3
IcmpMsg: 9 755 1 2 159 81 159 7969
Tcp: RtoAlgorithm RtoMin RtoMax MaxConn ActiveOpens PassiveOpens AttemptFails EstabResets CurrEstab InSegs OutSegs RetransSegs InErrs OutRsts InCsumErrors
Tcp: 1 200 120000 -1 43031 23536 390 101 1 1577612 3154151 11794 0 936 0
Udp: InDatagrams NoPorts InErrors OutDatagrams RcvbufErrors SndbufErrors InCsumErrors IgnoredMulti
Udp: 117068 30 0 68905 0 0 0 58
UdpLite: InDatagrams NoPorts InErrors OutDatagrams RcvbufErrors SndbufErrors InCsumErrors IgnoredMulti
UdpLite: 0 0 0 0 0 0 0 0
`
	parsed := parseCompact(bytes.NewBufferString(procNetSnmp))

	if len(parsed) != 85 {
		t.Errorf("Parsed value should have the correct number of entries. %d exptected / %d actual.", 85, len(parsed))
	}

	if parsed["IpForwarding"] != 1 {
		t.Errorf("Value for IpForwarding should be 1. %d actual.", parsed["IpForwarding"])
	}

	if parsed["IcmpInMsgs"] != 1010 {
		t.Errorf("Value for IcmpInMsgs should be 1010. %d actual.", parsed["IcmpInErrors"])
	}
}

func TestParseTable(t *testing.T) {
	procNetSnmp6 := `Ip6InReceives                           1416837
Ip6InHdrErrors                          9
Ip6InTooBigErrors                       0
Ip6InNoRoutes                           31
Ip6InAddrErrors                         0
Ip6InUnknownProtos                      0
Ip6InTruncatedPkts                      0
Ip6InDiscards                           0
Ip6InDelivers                           261635
Ip6OutForwDatagrams                     0
Ip6OutRequests                          230111
`

	parsed := parseTable(bytes.NewBufferString(procNetSnmp6))

	if len(parsed) != 11 {
		t.Errorf("Parsed value should have the correct number of entries. %d expected /%d actual.", 11, len(parsed))
	}

	if parsed["Ip6InDelivers"] != 261635 {
		t.Errorf("Parsed value should have the correct value.")
	}

}
