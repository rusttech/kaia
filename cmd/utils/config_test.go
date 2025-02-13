// Copyright 2024 The Kaia Authors
// This file is part of the Kaia library.
//
// The Kaia library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Kaia library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Kaia library. If not, see <http://www.gnu.org/licenses/>.
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func newAppWithAction(action func(ctx *cli.Context) error) *cli.App {
	app := cli.NewApp()
	app.Flags = append(app.Flags, AllNodeFlags()...)
	app.Before = func(ctx *cli.Context) error {
		if err := altsrc.InitInputSourceWithContext(
			AllNodeFlags(),
			altsrc.NewYamlSourceFromFlagFunc("conf"),
		)(ctx); err != nil {
			return err
		}
		return nil
	}
	app.Action = action
	return app
}

func TestLoadYaml(t *testing.T) {
	testcases := map[string]bool{
		"conf":                                      true,
		"ntp.disable":                               true,
		"ntp.server":                                true,
		"docroot":                                   false,
		"bootnodes":                                 true,
		"identity":                                  false,
		"unlock":                                    true,
		"password":                                  true,
		"dbtype":                                    true,
		"datadir":                                   false,
		"overwrite-genesis":                         true,
		"start-block-num":                           true,
		"keystore":                                  false,
		"txpool.nolocals":                           true,
		"txpool.allow-local-anchortx":               true,
		"txpool.deny.remotetx":                      true,
		"txpool.journal":                            true,
		"txpool.journal-interval":                   true,
		"txpool.pricelimit":                         true,
		"txpool.pricebump":                          true,
		"txpool.exec-slots.account":                 true,
		"txpool.exec-slots.all":                     true,
		"txpool.nonexec-slots.account":              true,
		"txpool.nonexec-slots.all":                  true,
		"txpool.lifetime":                           true,
		"txpool.keeplocals":                         true,
		"syncmode":                                  false,
		"gcmode":                                    true,
		"lightkdf":                                  true,
		"db.single":                                 true,
		"db.num-statetrie-shards":                   true,
		"db.leveldb.compression":                    true,
		"db.leveldb.no-buffer-pool":                 true,
		"db.no-perf-metrics":                        true,
		"db.dynamo.tablename":                       true,
		"db.dynamo.region":                          true,
		"db.dynamo.is-provisioned":                  true,
		"db.dynamo.read-capacity":                   true, // TODO-check after bugfix
		"db.dynamo.write-capacity":                  true, // TODO-check after bugfix
		"db.dynamo.read-only":                       true,
		"db.leveldb.cache-size":                     true,
		"db.pebbledb.cache-size":                    true,
		"db.no-parallel-write":                      true,
		"db.rocksdb.secondary":                      true,
		"db.rocksdb.cache-size":                     true,
		"db.rocksdb.dump-memory-stat":               true,
		"db.rocksdb.compression-type":               true,
		"db.rocksdb.bottommost-compression-type":    true,
		"db.rocksdb.filter-policy":                  true,
		"db.rocksdb.disable-metrics":                true,
		"sendertxhashindexing":                      true,
		"state.cache-size":                          true,
		"state.block-interval":                      true,
		"state.tries-in-memory":                     true,
		"state.live-pruning":                        true,
		"state.live-pruning-retention":              true,
		"cache.type":                                true,
		"cache.scale":                               true,
		"cache.level":                               false,
		"cache.memory":                              true,
		"statedb.cache.type":                        true,
		"statedb.cache.num-fetcher-prefetch-worker": true,
		"statedb.cache.use-snapshot-for-prefetch":   true,
		"state.trie-cache-limit":                    true,
		"state.trie-cache-save-period":              true,
		"statedb.cache.redis.endpoints":             true,
		"statedb.cache.redis.cluster":               true,
		"statedb.cache.redis.publish":               true,
		"statedb.cache.redis.subscribe":             true,
		"port":                                      true,
		"subport":                                   true,
		"multichannel":                              true,
		"maxconnections":                            true,
		"maxRequestContentLength":                   true,
		"maxpendpeers":                              true,
		"targetgaslimit":                            true,
		"nat":                                       true,
		"nodiscover":                                true,
		"rwtimerwaittime":                           true,
		"rwtimerinterval":                           true,
		"netrestrict":                               false,
		"nodekey":                                   false,
		"nodekeyhex":                                false,
		"vmdebug":                                   true,
		"vmlog":                                     true,
		"vm.internaltx":                             true,
		"networkid":                                 true,
		"metrics":                                   true,
		"prometheus":                                true,
		"prometheusport":                            true,
		"extradata":                                 false,
		"srvtype":                                   true,
		"autorestart.enable":                        true,
		"autorestart.timeout":                       true,
		"autorestart.daemon.path":                   true,
		"config":                                    false,
		"api.filter.getLogs.maxitems":               true,
		"api.filter.getLogs.deadline":               true,
		"opcode-computation-cost-limit":             true,
		"snapshot":                                  true,
		"snapshot.cache-size":                       true,
		"snapshot.async-gen":                        true,
		"rpc":                                       true,
		"rpcaddr":                                   true,
		"rpcport":                                   true,
		"rpcapi":                                    true,
		"rpc.gascap":                                true,
		"rpc.ethtxfeecap":                           true,
		"rpccorsdomain":                             false,
		"rpcvhosts":                                 true,
		"rpc.eth.noncompatible":                     true,
		"ws":                                        true,
		"wsaddr":                                    true,
		"wsport":                                    true,
		"grpc":                                      true,
		"grpcaddr":                                  true,
		"grpcport":                                  true,
		"rpc.concurrencylimit":                      true,
		"wsapi":                                     true,
		"wsorigins":                                 true,
		"wsmaxsubscriptionperconn":                  true,
		"wsreaddeadline":                            true, // TODO-check after bugfix
		"wswritedeadline":                           true,
		"wsmaxconnections":                          true,
		"ipcdisable":                                true,
		"ipcpath":                                   false,
		"rpcreadtimeout":                            true,
		"rpcwritetimeout":                           true,
		"rpcidletimeout":                            true,
		"rpcexecutiontimeout":                       true,
		"jspath":                                    true,
		"exec":                                      false,
		"preload":                                   false,
		"verbosity":                                 true,
		"vmodule":                                   true,
		"backtrace":                                 true,
		"debug":                                     true,
		"pprof":                                     true,
		"pprofaddr":                                 true,
		"pprofport":                                 true,
		"memprofile":                                true,
		"memprofilerate":                            true,
		"blockprofilerate":                          true,
		"cpuprofile":                                true,
		"trace":                                     true,
		"chaindatafetcher":                          true,
		"chaindatafetcher.mode":                     true,
		"chaindatafetcher.no.default":               true,
		"chaindatafetcher.num.handlers":             true,
		"chaindatafetcher.job.channel.size":         true,
		"chaindatafetcher.block.channel.size":       true,
		"chaindatafetcher.max.processing.data.size": true,
		"chaindatafetcher.kas.db.host":              false,
		"chaindatafetcher.kas.db.port":              false,
		"chaindatafetcher.kas.db.name":              false,
		"chaindatafetcher.kas.db.user":              false,
		"chaindatafetcher.kas.db.password":          false,
		"chaindatafetcher.kas.cache.use":            false,
		"chaindatafetcher.kas.cache.url":            false,
		"chaindatafetcher.kas.xchainid":             false,
		"chaindatafetcher.kas.basic.auth.param":     false,
		"chaindatafetcher.kafka.replicas":           true, // TODO-check after bugfix
		"chaindatafetcher.kafka.brokers":            true,
		"chaindatafetcher.kafka.partitions":         true,
		"chaindatafetcher.kafka.topic.resource":     true,
		"chaindatafetcher.kafka.topic.environment":  true,
		"chaindatafetcher.kafka.max.message.bytes":  true,
		"chaindatafetcher.kafka.segment.size":       true,
		"chaindatafetcher.kafka.required.acks":      true,
		"chaindatafetcher.kafka.msg.version":        true,
		"chaindatafetcher.kafka.producer.id":        false,
		"dst.dbtype":                                false,
		"dst.datadir":                               false,
		"db.dst.single":                             false,
		"db.dst.leveldb.cache-size":                 false,
		"db.dst.leveldb.compression":                false,
		"db.dst.num-statetrie-shards":               false,
		"db.dst.dynamo.tablename":                   false,
		"db.dst.dynamo.region":                      false,
		"db.dst.dynamo.is-provisioned":              false,
		"db.dst.dynamo.read-capacity":               false,
		"db.dst.dynamo.write-capacity":              false,
		"genkey":                                    false,
		"writeaddress":                              true,
		"bnaddr":                                    true,
		"authorized-nodes":                          false,
		"rewardbase":                                false,
		"mainnet":                                   true,
		"kairos":                                    true,
		"block-generation-interval":                 true,
		"block-generation-time-limit":               true,
		"txresend.interval":                         true,
		"txresend.max-count":                        true,
		"txresend.use-legacy":                       true,
		"txpool.spamthrottler.disable":              true,
		"scsigner":                                  false,
		"childchainindexing":                        true,
		"mainbridge":                                true,
		"mainbridgeport":                            true,
		"kes.nodetype.service":                      true,
		"dbsyncer":                                  true,
		"dbsyncer.db.host":                          true,
		"dbsyncer.db.port":                          true,
		"dbsyncer.db.name":                          true,
		"dbsyncer.db.user":                          true,
		"dbsyncer.db.password":                      true,
		"dbsyncer.logmode":                          true,
		"dbsyncer.db.max.idle":                      true,
		"dbsyncer.db.max.open":                      true,
		"dbsyncer.db.max.lifetime":                  true,
		"dbsyncer.block.channel.size":               true,
		"dbsyncer.mode":                             true,
		"dbsyncer.genquery.th":                      true,
		"dbsyncer.insert.th":                        true,
		"dbsyncer.bulk.size":                        true,
		"dbsyncer.event.mode":                       true,
		"dbsyncer.max.block.diff":                   true,
		"chaintxperiod":                             true,
		"chaintxlimit":                              true,
		"subbridge":                                 true,
		"subbridgeport":                             true,
		"parentchainid":                             true,
		"vtrecovery":                                true,
		"vtrecoveryinterval":                        true,
		"scnewaccount":                              true,
		"anchoring":                                 true,
		"sc.parentoperator.gaslimit":                true,
		"sc.childoperator.gaslimit":                 true,
		"kas.sc.anchor":                             false,
		"kas.sc.anchor.period":                      false,
		"kas.sc.anchor.url":                         false,
		"kas.sc.anchor.operator":                    false,
		"kas.secretkey":                             false,
		"kas.accesskey":                             false,
		"kas.x-chain-id":                            false,
		"kas.sc.anchor.request.timeout":             false,
	}

	printFlags := func(ctx *cli.Context) error {
		for _, flag := range AllNodeFlags() {
			assert.Equal(t, testcases[flag.Names()[0]], ctx.IsSet(flag.Names()[0]), "IsSet returned unexpected result", flag.Names()[0])
		}
		return nil
	}
	app := newAppWithAction(printFlags)
	err := app.Run([]string{"testApp", "--conf", "nodecmd/testdata/test-config.yaml"})
	if err != nil {
		t.Error(err)
	}
}
