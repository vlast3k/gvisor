package sandbox

import (
	"net"
	"testing"

	"github.com/vishvananda/netlink"
)

func TestFindLinkAddressPreservesPeerAttributes(t *testing.T) {
	local := net.ParseIP("10.152.28.207")
	peer := net.ParseIP("10.152.28.1")
	existing := []netlink.Addr{
		{IPNet: &net.IPNet{IP: local, Mask: net.CIDRMask(24, 32)}},
		{IPNet: &net.IPNet{IP: local, Mask: net.CIDRMask(32, 32)}, Peer: &net.IPNet{IP: peer, Mask: net.CIDRMask(32, 32)}, Scope: 253, Flags: 128},
	}

	got := findLinkAddress(existing, &net.IPNet{IP: local, Mask: net.CIDRMask(32, 32)})
	if got == nil {
		t.Fatal("findLinkAddress() = nil, want point-to-point address")
	}
	if got.Peer == nil || !got.Peer.IP.Equal(peer) {
		t.Fatalf("findLinkAddress() peer = %v, want %v", got.Peer, peer)
	}
	if got.Scope != 253 || got.Flags != 128 {
		t.Fatalf("findLinkAddress() scope/flags = %d/%d, want 253/128", got.Scope, got.Flags)
	}
}

func TestFindLinkAddressMissingIsNil(t *testing.T) {
	existing := []netlink.Addr{{
		IPNet: &net.IPNet{IP: net.ParseIP("10.152.28.207"), Mask: net.CIDRMask(32, 32)},
	}}
	got := findLinkAddress(existing, &net.IPNet{IP: net.ParseIP("10.152.28.208"), Mask: net.CIDRMask(32, 32)})
	if got != nil {
		t.Fatalf("findLinkAddress() = %v, want nil", got)
	}
}
