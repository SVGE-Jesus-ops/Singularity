# Singularity - Quantum-Enhanced Blockchain

## Overview

Singularity is a next-generation blockchain platform that integrates quantum random number generation (QRNG) for enhanced security and unpredictability. The project leverages the Australian National University's Quantum Random Number Generator to create truly random values for blockchain operations, ensuring cryptographic strength beyond traditional pseudo-random approaches.

## SINCERELY Token ($NC)

The SINCERELY token ($NC) is the native utility token of the Singularity blockchain ecosystem.

### Token Details
- **Symbol**: $NC
- **Standard**: ERC20 Compatible
- **Use Cases**: 
  - Transaction fees on the Singularity network
  - Staking and governance
  - Access to quantum-enhanced features
  - Smart contract execution

## Quantum RNG Integration

### ANU Quantum Random Number Generator

Singularity integrates with the ANU QRNG API to generate truly random numbers from quantum vacuum fluctuations.

**API Endpoint**: https://qrng.anu.edu.au/API/jsonI.php

**Features**:
- Generation rate: 5.7 Gbits/s (limited by network bandwidth)
- True quantum randomness from vacuum fluctuations
- Used for block nonces, cryptographic keys, and unpredictable blockchain operations

**Parameters**:
- `length`: 1-1024 (number of random values)
- `type`: uint8, uint16, or hex16
- `size`: 1-1024 (array size)

## Architecture Overview

### Core Components

#### Blockchain Layer (`src/blockchain`)
- **block.go**: Blockchain block structure with quantum RNG integration
- **quantumrng.go**: ANU Quantum Random Number Generator API client
- Quantum implementations: Aus, Quan, Quant, Quantum, Ram, Func

#### Smart Contracts
- **token-erc20.sol.txt**: ERC20 token contract for SINCERELY ($NC)
- Solidity contracts with security considerations

#### Firebase Integration
- **CoinForge - Firebase Studio**: Backend services integration
- Firestore for real-time databases
- Cloud Functions for serverless operations
- Authentication and hosting services

#### Workflow Automation (`.github/workflows`)
- IBM Watson integration
- Docker image builds
- CodeQL security scanning
- Defender for DevOps
- MetaMask integration
- Chromium EC crypto keys

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Node.js 16+ (for smart contract development)
- Git
- Docker (optional, for containerized deployment)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/SVGE-Jesus-ops/Singularity.git
   cd Singularity
   ```

2. **Install Go dependencies**:
   ```bash
   go mod download
   ```

3. **Build the blockchain**:
   ```bash
   go build -o singularity ./src/blockchain
   ```

4. **Run the node**:
   ```bash
   ./singularity start
   ```

### Configuration

Create a `config.yaml` file:

```yaml
network:
  port: 8545
  rpc_enabled: true

quantum:
  anu_api_url: https://qrng.anu.edu.au/API/jsonI.php
  default_length: 1024
  default_type: hex16

firebase:
  project_id: your-project-id
  credentials: path/to/credentials.json
```

## API Documentation

### Blockchain API

#### Get Block
```
GET /api/v1/block/:hash
```

Returns block information by hash.

#### Create Transaction
```
POST /api/v1/transaction
Content-Type: application/json

{
  "from": "0x...",
  "to": "0x...",
  "value": "1000000000000000000",
  "data": "0x"
}
```

#### Get Quantum Random Number
```
GET /api/v1/quantum/random?length=10&type=hex16
```

Returns quantum-generated random numbers.

### Smart Contract Interaction

#### Deploy Contract
```bash
npx hardhat run scripts/deploy.js --network singularity
```

#### Interact with SINCERELY Token
```javascript
const token = await ethers.getContractAt("SINCERELY", tokenAddress);
const balance = await token.balanceOf(address);
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Quality

The project uses:
- CodeQL for security analysis
- GitHub Actions for CI/CD
- Defender for DevOps for security monitoring

### Known Issues

[CHECKMARK] **Fixed Issues**:
- block.go typos corrected (commit 8590333a)

### Roadmap

- [ ] Consolidate quantum RNG implementations into quantumrng.go
- [ ] Review and fix integer overflow warnings in smart contracts
- [ ] Complete Firebase Studio environment setup for CoinForge
- [ ] Configure MetaMask SDK for Web3 interactions
- [ ] Implement cross-chain bridge functionality
- [ ] Launch mainnet with SINCERELY token

## Documentation Archives

The repository includes comprehensive documentation:
- `chainlist (1).json`: Blockchain network configurations (16KB)
- `pyetherscan-readthedocs-io-en-latest.zip`: Etherscan API documentation (3.2MB)
- `us-digital-registry-master.zip`: US Digital Registry (16MB)
- PGP keys and checksums for verification

## Security

### Quantum Security Features

- True random number generation for cryptographic operations
- Enhanced unpredictability in block mining
- Quantum-resistant algorithm considerations

### Reporting Security Issues

Please report security vulnerabilities to: security@singularity-blockchain.io

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

- **Project Lead**: SVGE-Jesus-ops
- **Repository**: https://github.com/SVGE-Jesus-ops/Singularity
- **Issues**: https://github.com/SVGE-Jesus-ops/Singularity/issues
- **Discussions**: https://github.com/SVGE-Jesus-ops/Singularity/discussions

## Acknowledgments

- Australian National University for the Quantum Random Number Generator API
- The Ethereum community for ERC20 standards
- Firebase team for cloud infrastructure support
- All contributors and supporters of the Singularity project

---

**Built with Quantum Precision | Powered by SINCERELY ($NC)**