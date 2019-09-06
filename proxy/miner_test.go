package proxy
import (
	"bytes"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"strconv"
	"strings"

	//"strconv"
	//"strings"
	"testing"
	"math/big"
	"github.com/NginProject/ngind/common"
)

func TestProxyServer_processShareGetwork(t *testing.T) {
	// Pool -> Miner
	task := []string{
	// 	"", // header 

	"0x887c6f943fdbb298ceb40d58c4e545264f148ceac7ae00c103c44f7d5f8aedd1",
	"0x851eaf8bf2dfc0087cf1ecb186d1a435fec582ffb1da52c3581e7d1a71558141",
	"0xf6b948a4cdc38615719a9052e23e675d1faf0b3429ec7bd800f6b948a4cd",
	
	// 	"", // seed
	// 	"", // diff
	}

	// Miner -> Pool
	params := []string{
		"0x398717c5156e21b2",
		"0x887c6f943fdbb298ceb40d58c4e545264f148ceac7ae00c103c44f7d5f8aedd1",
		"0x00006b600be445c653716084d9e3f386ac1f886dc0389c1aede6c4eb5266f90a",

		// "0xd74f80dc271fedfe", // nonce
		// "0xe3e4fca9c5e1d1b3095ea11bbc81849065930d5a432d8b1918954a8530b61857", // header
		// "0x00001394e7d2e54c2a4d96df1a21c9066151569bd1ba5d4d4717e73f36fa384c", // hash
	}
	nonceHex := params[0]
	nonce, _ := strconv.ParseUint(strings.Replace(nonceHex, "0x", "", -1), 16, 64)
	log.Println("Nonce is:", nonce)
	hashNoNonce := params[1]
	mixDigest := params[2]

	blob := common.HexToHash(hashNoNonce)

	maxUint256 := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
	var i int
	for {
		nonce++
		i++
		log.Println(i)
		digest, _ := hashcryptonight(blob.Bytes(), nonce)
		if !bytes.Equal(common.HexToHash(mixDigest).Bytes(), digest) {
			log.Println("digest incorrect!")
		}else{
			log.Println("digest correct!!!!!!!!!!!!!!!!!!!!!!!!!")
			return
		}

	diff, _ := hexutil.Decode(task[2])
	bigDiff := new(big.Int).SetBytes(diff)
		//:= strconv.ParseUint(, 16, 64)
		targetShare := new(big.Int).Div(maxUint256, new(big.Int).SetBytes(diff))
		log.Println("targetShare is :", targetShare)
		bDigest := common.HexToHash(mixDigest).Bytes()
		// bigHash := new(big.Int).SetBytes(digest)
		if new(big.Int).SetBytes(digest).Cmp(bigDiff) > 0 {
			log.Println("in their algo:", new(big.Int).SetBytes(bDigest))
			log.Println("in our algo:", new(big.Int).SetBytes(digest))
			log.Println("larger than target:", bigDiff)
		}else{
			log.Println("correct nonce!!!!!!!!!!!!!!!!!!!!!!")
			return
		}
	// }

	// shareDiff := s.config.Proxy.Difficulty

	// maxUint256 := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

	//2229777929363975550
	//4145308117253366194
	//13109809067366014541
	//7367306370600957031

	// h, ok := t.headers[hashNoNonce]
	// if !ok {
	// 	log.Printf("Stale share from %v@%v", login, ip)
	// 	return false, false
	// }

	// share := Block{
	// 	number:      h.height,
	// 	hashNoNonce: common.HexToHash(hashNoNonce),
	// 	difficulty:  big.NewInt(shareDiff),
	// 	nonce:       nonce,
	// 	mixDigest:   common.HexToHash(mixDigest),
	// }

	// block := Block{
	// 	number:      h.height,
	// 	hashNoNonce: common.HexToHash(hashNoNonce),
	// 	difficulty:  h.diff,
	// 	nonce:       nonce,
	// 	mixDigest:   common.HexToHash(mixDigest),
	// }
	// log.Printf("%+v", shareDiff)
	// log.Printf("%#v,", h.diff)

	// log.Printf("%+x", share)
	// log.Printf("%#v", share)
	// log.Printf("%+x", block)
	// log.Printf("%#v", block)

	// blob := common.HexToHash(hashNoNonce)
	// digest, _ := hashcryptonight(blob.Bytes(), nonce)
	// params[2] = common.ToHex(digest)

	// if bytes.Equal(digest, mix) {
	// }

	// if !bytes.Equal(block.mixDigest.Bytes(), digest) {
	// 	log.Printf("digest incorrect!")
	// 	return false, false
	// }

	// targetShare := new(big.Int).Div(maxUint256, share.difficulty)
	// if new(big.Int).SetBytes(digest).Cmp(targetShare) <= 0 {
	// 	log.Printf("share accepted")
	// } else {
	// 	return false, false
	// }
	//if !hasher.Verify(share) { return false, false }

	// targetBlock := new(big.Int).Div(maxUint256, block.difficulty)
	// if new(big.Int).SetBytes(digest).Cmp(targetBlock) <= 0 {
	// 	//if hasher.Verify(block) {
	// 	ok, err := s.rpc().SubmitBlock(params)
	// 	if err != nil {
	// 		log.Printf("Block submission failure at height %v for %v: %v", h.height, t.Header, err)
	// 	} else if !ok {
	// 		log.Printf("Block rejected at height %v for %v", h.height, t.Header)
	// 		return false, false
	// 	} else {
	// 		s.fetchBlockTemplate()
	// 		exist, err := s.backend.WriteBlock(login, id, params, shareDiff, h.diff.Int64(), h.height, s.hashrateExpiration)
	// 		if exist {
	// 			return true, false
	// 		}
	// 		if err != nil {
	// 			log.Println("Failed to insert block candidate into backend:", err)
	// 		} else {
	// 			log.Printf("Inserted block %v to backend", h.height)
	// 		}
	// 		log.Printf("Block found by miner %v@%v at height %d", login, ip, h.height)
	// 	}
	// } else {
	// 	exist, err := s.backend.WriteShare(login, id, params, shareDiff, h.height, s.hashrateExpiration)
	// 	if exist {
	// 		return true, false
	// 	}
	// 	if err != nil {
	// 		log.Println("Failed to insert share data into backend:", err)
	// 	}
	// }

	// ok, err := s.rpc().SubmitBlock(params)
	// if err != nil {
	// 	log.Printf("Block submission failure at height %v for %v: %v", h.height, t.Header, err)
	// } else if !ok {
	// 	log.Printf("Block rejected at height %v for %v", h.height, t.Header)
	// } else {
	// 	s.fetchBlockTemplate()
	// 	_, err := s.backend.WriteBlock(login, id, params, shareDiff, h.diff.Int64(), h.height, s.hashrateExpiration)
	// 	// if exist {
	// 	// 	return true, false
	// 	// }
	// 	if err != nil {
	// 		log.Println("Failed to insert block candidate into backend:", err)
	// 	} else {
	// 		log.Printf("Inserted block %v to backend", h.height)
	// 	}
	// 	log.Printf("Block found by miner %v@%v at height %d", login, ip, h.height)
	// 	// return false, true
	// }

	// _, err = s.backend.WriteShare(login, id, params, shareDiff, h.height, s.hashrateExpiration)
	// // if exist {
	// // 	return true, false
	// // }
	// if err != nil {
	// 	log.Println("Failed to insert share data into backend:", err)
	}

	// return false, true
}

