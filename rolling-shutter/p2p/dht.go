package p2p

import (
	"context"
	"math/rand"
	"time"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/rs/zerolog/log"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/encodeable/env"
)

var DefaultOptions []dht.Option

const (
	// resulting protocol string will be: '/<prefix>{/<extension>}/kad/1.0.0'.
	dhtProtocolPrefix           protocol.ID = "/shutter"
	dhtProtocolExtensionStaging protocol.ID = "/staging"
	dhtProtocolExtensionLocal   protocol.ID = "/local"

	findPeerInterval = 10 * time.Second
)

var (
	peerLow    = pubsub.GossipSubDlo * 2
	peerTarget = pubsub.GossipSubDhi * 3
	peerHigh   = pubsub.GossipSubDhi * 6
)

func dhtRoutingOptions(config *p2pNodeConfig) []dht.Option {
	// options with higher index in the array will overwrite existing ones
	opts := []dht.Option{
		dht.ProtocolPrefix(dhtProtocolPrefix),
		dht.BootstrapPeers(config.BootstrapPeers...),
	}

	switch config.Environment { //nolint: exhaustive
	case env.EnvironmentStaging:
		opts = append(opts,
			dht.ProtocolExtension(dhtProtocolExtensionStaging),
		)
	case env.EnvironmentLocal:
		opts = append(opts,
			dht.ProtocolExtension(dhtProtocolExtensionLocal),
			// auto mode will not work when the AutoNAT sets the
			// reachability to "private" when we are not reachable
			// over a public IP.
			dht.Mode(dht.ModeServer),
		)
	default:
	}

	if config.IsBootstrapNode {
		opts = append(opts, dht.Mode(dht.ModeServer))
	} else {
		opts = append(opts, dht.Mode(dht.ModeAutoServer))
	}

	return opts
}

func findPeers(ctx context.Context, h host.Host, d discovery.Discoverer, ns string) error {
	log.Info().Str("namespace", ns).Msg("starting peer discovery")

	ticker := time.NewTicker(findPeerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			peersBefore := len(h.Network().Peers())
			if peersBefore >= peerTarget {
				log.Debug().Int("peers-before", peersBefore).Int("peer-target", peerTarget).Msg("have enough peers")
				continue
			}

			peers, err := util.FindPeers(ctx, d, ns)
			if err != nil {
				log.Error().Err(err).Msg("error finding peers")
			}

			newConnections := 0
			failedDials := 0
			ourId := h.ID().String()

			randomisedPeers := randomisePeers(peers)

			for _, p := range randomisedPeers {
				collectPeerAddresses(p)
				if p.ID == h.ID() {
					continue
				}
				connectedness := h.Network().Connectedness(p.ID)
				metricsP2PPeerConnectedness.WithLabelValues(ourId, p.ID.String()).Set(float64(connectedness))
				peerPing := h.Peerstore().LatencyEWMA(p.ID)
				if peerPing != 0 {
					metricsP2PPeerPing.WithLabelValues(ourId, p.ID.String()).Set(peerPing.Seconds())
				}
				if connectedness != network.Connected {
					_, err = h.Network().DialPeer(ctx, p.ID)
					if err != nil {
						log.Debug().
							Err(err).
							Str("peer", p.ID.String()).
							Msg("error dialing peer")
						failedDials++
					}
					newConnections++
				}
			}

			log.Debug().
				Int("peers-before", peersBefore).
				Int("peer-target", peerTarget).
				Int("new-connections", newConnections).
				Int("failed-dials", failedDials).
				Msg("looking for peers")
		}
	}
}

func randomisePeers(peers []peer.AddrInfo) []peer.AddrInfo {
	randomIndexes := rand.Perm(len(peers))
	randomised := make([]peer.AddrInfo, len(peers))
	for _, index := range randomIndexes {
		randomised = append(randomised, peers[index])
	}
	return randomised
}
