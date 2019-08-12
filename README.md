# gobtclib

A client library with golang for Bitcoin

![GitHub release](https://img.shields.io/github/release/chainlibs/gobtclib.svg)
![Github commits (since latest release)](https://img.shields.io/github/commits-since/wanyvic/gobtclib/latest.svg)
[![Build Status](https://travis-ci.org/chainlibs/gobtclib.svg?branch=master)](https://travis-ci.org/chainlibs/gobtclib)
[![GitHub issues](https://img.shields.io/github/issues/chainlibs/gobtclib.svg)](https://github.com/wanyvic/gobtclib/issues)
[![GitHub stars](https://img.shields.io/github/stars/chainlibs/gobtclib.svg)](https://github.com/wanyvic/gobtclib/stargazers)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/wanyvic/gobtclib/master/LICENSE)
[![Gitter chat](https://badges.gitter.im/owner/repo.png)](https://gitter.im/gobtclib/Lobby)

What is Bitcoin?
----------------

Bitcoin is an experimental digital currency that enables instant payments to
anyone, anywhere in the world. Bitcoin uses peer-to-peer technology to operate
with no central authority: managing transactions and issuing money are carried
out collectively by the network. Bitcoin Core is the name of open source
software which enables the use of this currency.

The current project is the client library for bitcoin, you can read the document from
https://bitcoincore.org/en/doc for more information.

What is gobtclib?
----------------

gobtclib is a golang-based library for working with the Bitcoin protocol. It can maintain a wallet, send/receive transactions without
needing a local copy of Bitcoin Core and has many other advanced features.

Support RPC list
----------------

- BLOCKCHAIN
    - getbestblockhash
    - getblock
    - getblockchaininfo
    - getblockcount
    - getblockhash
    - getblockheader
    - getblockstats
    - getchaintips
    - getchaintxstats
    - getdifficulty
    - getmempoolancestors
    - getmempooldescendants
    - getmempoolentry
    - getmempoolinfo
    - getrawmempool
    - gettxout
    - gettxoutproof
    - gettxoutsetinfo
    - preciousblock
    - pruneblockchain
    - savemempool
    - verifychain
    - verifytxoutproof
- CONTROL
    - getmemoryinfo
    - help
    - logging
    - stop
    - uptime
- GENERATING
- MINING
- NETWORK
- RAWTRANSACTIONS
    - CombineRawTransaction
    - CreateRawTransaction
    - DecodeRawTransaction
    - DecodeScript
    - FundRawTransaction
    - GetRawTransaction
    - GetRawTransactionVerbose
    - SendRawTransaction
    - SignRawTransactionWithKey
    - TestMempoolAccept
- UTIL
- WALLET
    - abandontransaction
    - abortrescan
    - addmultisigaddress
    - backupwallet
    - bumpfee
    - createwallet
    - dumpprivkey
    - dumpwallet
    - encryptwallet
    - getaccount
    - getaccountaddress
    - getaddressesbylabel
    - getaddressinfo
    - getbalance
    - getnewaddress
    - getrawchangeaddress
    - getreceivedbyaccount
    - getreceivedbyaddress
    - gettransaction
    - getunconfirmedbalance
    - getwalletinfo
    - importaddress
    - importmulti
    - importprivkey
    - -importprunedfunds-
    - importpubkey
    - importwallet
    - keypoolrefill
    - listaddressgroupings
    - listlabels
    - listlockunspent
    - listreceivedbyaddress
    - listsinceblock
    - listtransactions
    - listunspent
    - listwallets
    - loadwallet
    - lockunspent
    - move
    - removeprunedfunds
    - rescanblockchain
    - sendfrom
    - sendmany
    - sendtoaddress
    - sethdseed
    - settxfee
    - signmessage
    - -signrawtransactionwithwallet-
    - unloadwallet
    - -walletcreatefundedpsbt-
    - walletlock
    - walletpassphrase
    - walletpassphrasechange
    - -walletprocesspsbt-

Usage
-----

1. download dependency library: go get -u go.uber.org/zap;go get -u github.com/gobasis/log;
1. Take example from [examples/main.go](/examples/main.go)

Development Process
-------------------

The `master` branch is regularly built and tested, but is not guaranteed to be
completely stable. [Tags](https://github.com/wanyvic/gobtclib/tags) are created
regularly to indicate new official, stable release versions of Bitcoin Core.

The contribution workflow is described in [CONTRIBUTING.md](CONTRIBUTING.md).

Testing
-------

Testing and code review is the bottleneck for development; we get more pull requests than
we can review and test on short notice. Please be patient and help out by testing
other people's pull requests, and remember this is a security-critical project where
any mistake might cost people lots of money.

### Automated Testing

Developers are strongly encouraged to write unit tests for new code, and to
submit new unit tests for old code. Unit tests can be compiled and run.

The Travis CI system makes sure that every pull request is built for Windows, Linux, and macOS,
and that unit/sanity tests are run automatically.

### Manual Quality Assurance (QA) Testing

Changes should be tested by somebody other than the developer who wrote the
code. This is especially important for large or high-risk changes. It is useful
to add a test plan to the pull request description if testing the changes is
not straightforward.

License
-------

Chainlibs/gobtclib is released under the terms of the MIT license. See [COPYING](LICENSE) for more
information or see https://opensource.org/licenses/MIT.
