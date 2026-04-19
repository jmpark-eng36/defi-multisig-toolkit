# defi-multisig-toolkit

Go CLI for inspecting and managing Gnosis Safe multisigs across EVM chains.

Built for protocol operators who need quick visibility into signer status, thresholds, and activity without touching the Safe UI.

## Features

- **Inspect** — View owners, threshold, nonce, balances, and TX counts in one command
- **Activity Tracking** — Check signer activity and detect dormant keys
- **Multi-chain** — Works with any EVM chain via RPC flag

## Install

```bash
git clone https://github.com/jmpark-eng36/defi-multisig-toolkit.git
cd defi-multisig-toolkit
go build -o msig .
```

## Usage

### Inspect a Safe

```bash
./msig inspect 0xYourSafeAddress --rpc https://eth.llamarpc.com
```

Output:
```
  Safe: 0xYour...
  Threshold: 3/5
  Nonce: 89

    #   OWNER        BALANCE (ETH)   TX COUNT
    0   0xAA...1234  1.2300          456
    1   0xBB...5678  0.0500          12
    2   0xCC...9abc  0.0000          0
    3   0xDD...def0  5.4321          892
    4   0xEE...1234  0.1000          34
```

### Multi-chain

```bash
./msig inspect 0x... --rpc https://arb1.arbitrum.io/rpc    # Arbitrum
./msig inspect 0x... --rpc https://mainnet.base.org         # Base
```

## Roadmap

- [x] Basic Safe inspection
- [x] Signer activity tracking
- [ ] Transaction simulation on forked state
- [ ] Inactivity alerts (signer dormant > N days)
- [ ] Batch TX builder
- [ ] Support for non-Gnosis multisigs

## License

MIT
