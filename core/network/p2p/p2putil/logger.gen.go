// Code generated by berty.tech/core/.scripts/generate-logger.sh

package p2putil

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("core.network.p2p.p2putil")
}