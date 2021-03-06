package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"chainspace.io/chainspace-go/config"
	"chainspace.io/chainspace-go/internal/log"
	"chainspace.io/chainspace-go/internal/log/fld"
)

func cmdInit(args []string, usage string) {
	opts := newOpts("init NETWORK_NAME [OPTIONS]", usage+`.

  If the --registry host is specified, then a 36-byte token is randomly
  generated and the hex-encoded form is set as a shared secret across
  all of the initialised nodes.
`)

	configRoot := opts.Flags("--config-root").Label("PATH").String("Path to the Chainspace root directory [~/.chainspace]", defaultRootDir())
	registry := opts.Flags("--registry").Label("HOST").String("Address of the network registry")
	shardCount := opts.Flags("--shard-count").Label("N").Int("Number of shards in the network [3]")
	shardSize := opts.Flags("--shard-size").Label("N").Int("Number of nodes in each shard [4]")
	httpPort := opts.Flags("--http-port").Label("PORT").Int("HTTP port to use with the shards")
	disableSBAC := opts.Flags("--disable-sbac").Label("BOOL").Bool("Disable sbac")
	manageContracts := opts.Flags("--manage-contracts").Label("BOOL").Bool("Manage docker contracts")
	enablePubsub := opts.Flags("--enable-pubsub").Label("BOOL").Bool("Enable pubsub")
	pubsubPort := opts.Flags("--pubsub-port").Label("PORT").Int("Port to use for the pubsub server")

	roundInterval := opts.Flags("--round-interval").Label("N").Int("Round interval in milliseconds")
	params := opts.Parse(args)

	if err := ensureDir(*configRoot); err != nil {
		log.Fatal("Could not ensure the existence of the root directory", fld.Err(err))
	}

	if len(params) < 1 {
		opts.PrintUsage()
		os.Exit(1)
	}

	networkName := params[0]
	netDir := filepath.Join(*configRoot, networkName)
	createUnlessExists(netDir)

	ri := 1 * time.Second
	if roundInterval != nil && *roundInterval != 0 {
		ri = time.Duration(*roundInterval) * time.Millisecond
	}
	announce := &config.Announce{}
	bootstrap := &config.Bootstrap{}
	consensus := &config.NetConsensus{
		BlockReferencesSizeLimit:   10 * config.MB,
		BlockTransactionsSizeLimit: 100 * config.MB,
		NonceExpiration:            30 * time.Second,
		RoundInterval:              ri,
		ViewTimeout:                15,
	}

	peers := map[uint64]*config.Peer{}
	registries := []config.Registry{}
	if *registry == "" {
		announce.MDNS = true
		bootstrap.MDNS = true
	} else {
		announce.Registry = true
		bootstrap.Registry = true
		token := make([]byte, 36)
		_, err := rand.Read(token)
		if err != nil {
			log.Fatal("Unable to generate random token for use with the network registry", fld.Err(err))
		}
		registries = append(registries, config.Registry{
			Host:  *registry,
			Token: hex.EncodeToString(token),
		})
	}

	shard := &config.Shard{
		Count: *shardCount,
		Size:  *shardSize,
	}

	network := &config.Network{
		Consensus:  consensus,
		MaxPayload: 128 * config.MB,
		Shard:      shard,
		SeedNodes:  peers,
	}

	broadcast := &config.Broadcast{
		InitialBackoff: 1 * time.Second,
		MaxBackoff:     2 * time.Second,
	}

	connections := &config.Connections{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logging := &config.Logging{
		ConsoleLevel: log.DebugLevel,
		FileLevel:    log.ErrorLevel,
		FilePath:     "log/chainspace.log",
	}

	rateLimit := &config.RateLimit{
		InitialRate:  10000,
		RateDecrease: 0.8,
		RateIncrease: 1000,
	}

	storage := &config.Storage{
		Type: "badger",
	}

	var epb bool
	if enablePubsub != nil {
		epb = *enablePubsub
	}

	var psPort *int
	if pubsubPort != nil && *pubsubPort != 0 {
		psPort = pubsubPort
	}
	pubsub := &config.Pubsub{
		Enabled: epb,
		Port:    psPort,
	}

	var mContracts bool
	if manageContracts != nil {
		mContracts = *manageContracts
	}
	nodecontracts := &config.NodeContracts{
		Manage: mContracts,
		Docker: true,
	}

	dummyimage := "gcr.io/acoustic-atom-211511/chainspace.io/contract-dummy:latest"
	cts := &config.Contracts{
		DockerMinimalVersion: "1.30",
		DockerContracts: []config.DockerContract{
			{
				Name:           "dummy",
				Procedures:     []string{"dummy_ok", "dummy_ko"},
				Image:          dummyimage,
				Addr:           "http://0.0.0.0",
				HostPort:       "1789",
				Port:           "8080",
				HealthCheckURL: "/healthcheck",
			},
			{
				Name: "cs-coin",
				Procedures: []string{
					"create-wallet", "add-funds", "transfer-funds"},
				Image:          "chainspace.io/cs-coin:latest",
				Addr:           "http://0.0.0.0",
				HostPort:       "1492",
				Port:           "1789",
				HealthCheckURL: "/healthcheck",
			},
		},
	}

	/*
		if ((3 * (*shardSize / 3)) + 1) != *shardSize {
			log.Fatal("The given --shard-size does not satisfy the 3f+1 requirement", fld.ShardSize(uint64(*shardSize)))
		}
	*/

	totalNodes := *shardCount * *shardSize
	for i := 1; i <= totalNodes; i++ {

		log.Info("Generating node", fld.NetworkName(networkName), fld.NodeID(uint64(i)))

		nodeID := uint64(i)
		dirName := fmt.Sprintf("node-%d", i)
		nodeDir := filepath.Join(netDir, dirName)
		createUnlessExists(nodeDir)

		// Create keys.yaml
		signingKey, cert, err := genKeys(filepath.Join(nodeDir, "keys.yaml"), networkName, nodeID)
		if err != nil {
			log.Fatal("Could not generate keys", fld.Err(err))
		}

		var httpcfg config.HTTP
		if httpPort != nil && *httpPort != 0 {
			httpcfg = config.HTTP{
				Enabled: true,
				Port:    *httpPort,
			}
		} else {
			httpcfg = config.HTTP{
				Enabled: true,
				Port:    8000 + i,
			}
		}

		consensus := &config.NodeConsensus{
			DriftTolerance:      10 * time.Millisecond,
			InitialWorkDuration: 100 * time.Millisecond,
			RateLimit:           rateLimit,
		}

		var disableTxtor bool
		if disableSBAC != nil {
			disableTxtor = *disableSBAC
		}

		// Create node.yaml
		cfg := &config.Node{
			Announce:    announce,
			Bootstrap:   bootstrap,
			Broadcast:   broadcast,
			Connections: connections,
			Consensus:   consensus,
			Contracts:   nodecontracts,
			DisableSBAC: disableTxtor,
			HTTP:        httpcfg,
			Logging:     logging,
			Pubsub:      pubsub,
			Registries:  registries,
			Storage:     storage,
		}

		if err := writeYAML(filepath.Join(nodeDir, "node.yaml"), cfg); err != nil {
			log.Fatal("Could not write to node.yaml", fld.Err(err))
		}

		peers[nodeID] = &config.Peer{
			SigningKey: &config.PeerKey{
				Type:  signingKey.Algorithm().String(),
				Value: b32.EncodeToString(signingKey.PublicKey().Value()),
			},
			TransportCert: &config.PeerKey{
				Type:  cert.Type.String(),
				Value: cert.Public,
			},
		}

	}

	networkID, err := network.Hash()
	if err != nil {
		log.Fatal("Could not generate the Network ID", fld.Err(err))
	}

	network.ID = b32.EncodeToString(networkID)
	if err := writeYAML(filepath.Join(netDir, "network.yaml"), network); err != nil {
		log.Fatal("Could not write to network.yaml", fld.Err(err))
	}

	if err := writeYAML(filepath.Join(netDir, "contracts.yaml"), cts); err != nil {
		log.Fatal("Could not write to contracts.yaml", fld.Err(err))
	}
}
