package main

import (
	"github.com/gobasis/log"
	"github.com/gobasis/log/zapimpl"
	"github.com/wanyvic/gobtclib/client"
	"github.com/wanyvic/gobtclib/examples/demos"
)

/*
Description:
show demos of gobtclib/client
 * Author: architect.bian
 * Date: 2018/09/02 17:40
*/
func main() {
	log.UseLog(&zapimpl.Logger{}) // use zap log
	log.SetLevel(log.DevDebugLevel)
	log.Info("start up bitcoin rpc client")
	//cfg := &client.Config{
	//	Host:         "172.16.2.35:3333",
	//	User:         "et",
	//	Pass:         "www.et.com",
	//}
	cfg := &client.Config{
		Host: "localhost:19442",
		User: "user",
		Pass: "pwd",
	}
	demos.Initialize(cfg)
	defer demos.Shutdown()
	//common rpc method
	//demos.SendTest() cli.Send("command", args...)
	demos.SendManyEntires()
	//Blockchain rpc methods
	//demos.GetBestBlockHashTest()
	//demos.GetBlockBytesTest()
	//demos.GetBlockTest()
	//demos.GetBlockVerboseTXTest()
	// demos.GetBlockVerboseTXTest()
	//demos.GetBlockChainInfoTest()
	//demos.GetBlockCountTest()
	//demos.GetBlockHashTest()
	//demos.GetBlockHeaderTest()
	//demos.GetBlockStatsTest()
	//demos.GetChainTipsTest()
	//demos.GetChainTXStatsTest()
	//demos.GetChainTXStatsEntireTest()
	//demos.GetDifficultyTest()
	//demos.GetMempoolAncestorsTest()
	//demos.GetMempoolAncestorsVerboseTest()
	//demos.GetMempoolDescendantsTest()
	//demos.GetMempoolDescendantsVerboseTest()
	//demos.GetMempoolEntryTest()
	//demos.GetMempoolInfoTest()
	//demos.GetRawMempoolTest()
	//demos.GetRawMempoolVerboseTest()
	//demos.GetTXOutTest()
	//demos.GetTXOutProofTest()
	//demos.GetTXOutSetInfoTest()
	//demos.PreciousBlockTest()
	//demos.PruneBlockchainTest()
	//demos.SaveMempoolTest()
	//demos.VerifyChainTest()
	//demos.VerifyTXOutProofTest()
	//demos.VerifyChainLevelTest()
	//demos.VerifyChainBlocksTest()

	//Control rpc methods
	//demos.GetMemoryInfoTest()
	//demos.GetMemoryInfo4MallocInfoTest()
	//demos.HelpTest()
	//demos.StopTest()
	//demos.UptimeTest()

	//Generate rpc methods
	//demos.GenerateTest()
	//demos.GenerateToAddressTest()

	//Network rpc methods
	//demos.GetNetworkInfoTest()

	//RawTransactions rpc methods
	//demos.CombinePSBTTest() //ok
	//demos.CombineRawTransactionTest() //ok
	//demos.ConvertToPSBTTest() //ok
	//demos.CreatepSBTTest() //ok
	//demos.CreateRawTransactionTest() //ok
	//demos.DecodePSBTTest() //ok
	//demos.DecodeRawTransactionTest() //ok
	//demos.DecodeScriptTest() //ok
	//demos.FinalizePSBTTest() //ok
	//demos.FundRawTransactionTest() //ok
	//demos.GetRawTransactionTest() //ok
	//demos.GetRawTransactionVerboseTest() //ok
	//demos.SendRawTransactionTest() //ok
	//demos.SignRawTransactionWithKeyTest() //ok
	//demos.TestMempoolAcceptTest() //ok

	//Wallet rpc methods
	//demos.WalletTestInitialize() //ok
	//demos.AbandonTransactionTest() //ok
	//demos.AbortRescanTest() //ok
	//demos.AddMultiSigAddressTest() //ok
	//demos.BackupWalletTest() //ok
	//demos.BumpFeeTest() //ok
	//demos.CreateWalletTest() //ok
	//demos.DumpPrivkeyTest() //ok
	//demos.DumpWalletTest() //ok
	//demos.EncryptWalletTest() //ok
	//demos.GetAddressesByLabelTest() //ok
	//demos.GetAddressInfoTest() //ok
	//demos.GetBalanceTest() //ok
	//demos.GetBalanceEntireTest() //ok
	//demos.GetNewAddressTest() //ok
	//demos.GetNewAddressEntireTest() //ok
	//demos.GetRawChangeAddressTest() //ok
	//demos.GetReceivedByAddressTest() //ok
	//demos.GetTransactionTest() //ok
	//demos.GetUnconfirmedBalanceTest() //ok
	//demos.GetWalletInfoTest() //ok
	//demos.ImportaddressTest() //ok
	//demos.ImportMultiTest() //ok
	//demos.ImportPrivkeyTest() //ok
	//demos.ImportPrunedFundsTest() //TODO
	//demos.ImportPubkeyTest() //ok
	//demos.ImportWalletTest() //test exception
	//demos.KeypoolRefillTest() //ok
	//demos.ListAddressGroupingsTest() //ok
	//demos.ListLabelsTest() //ok
	//demos.ListLockUnspentTest() // todo https://bitcoincore.org/en/doc/0.17.0/rpc/wallet/listlockunspent/
	//demos.ListReceivedByAddressTest() //ok
	//demos.ListSinceBlockTest() //ok
	//demos.ListTransactionsTest() //ok
	//demos.ListUnspentTest() //ok
	//demos.ListUnspentEntireTest() //ok
	//demos.ListWalletsTest() //ok
	//demos.LoadWalletTest() //ok
	//demos.LockUnspentTest() //ok
	//demos.RemovePrunedFundsTest() //ok
	//demos.RescanBlockChainTest() //ok
	//demos.SendManyTest() //ok
	//demos.SendToAddressTest() //ok
	//demos.SendToAddressEntireTest() //ok
	//demos.SetHDSeedTest() //ok
	//demos.SetTXFeeTest() //ok
	//demos.SignMessageTest() //ok
	//demos.SignRawtransactionWithWalletTest() //TODO
	//demos.UnloadWalletTest() //ok
	//demos.WalletCreateFundedPSBTTest() //TODO no implement
	//demos.WalletLockTest() //ok
	//demos.WalletPassphraseTest() //ok
	//demos.WalletPassphraseChangeTest() //ok
	//demos.WalletProcessPSBTTest() //TODO
}
