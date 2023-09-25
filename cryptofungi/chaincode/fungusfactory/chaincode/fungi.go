package chaincode

import ( 
	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"time"
	"strconv"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"math"
)

type SmartContract struct{
	contractapi.Contract
}

var funguscount int = 0

type Fungus struct{
	FungusId uint
	Name string
	owner string
	Dna uint
	ReadyTime uint32
}

func(s *SmartContract) Initialize(){

}

func (s *SmartContract) CreateRandomFungus(ctx contractapi.TransactionContextInterface, name string) {
	dna :=s._generateRandomDna(name)
	_createFunugs(ctx, name, dna)
}

func _createFunugs(ctx contractapi.TransactionContextInterface, name string, dna uint) {

	funguscountJSON, err := ctx.Getstub().Getstate("funguscount")
	fungusIdINT,_ := strconv.AtoI(string(funguscountJSON))

	clientID,_ := ctx.GetClientIdentity().GetID()


	fungus :=Fungus{
		FungusId: uint(funguscount)+1,
		Name : name,
		Owner: clientID,
		Dna: dna,
		ReadyTime: uint32(unixtime)+60,

	}

	fungusJSON, _ := json.Marshal(fungus)
	ctx.Getstub().putstate(strconv.Itoa(int(fungus.FungusId)), fungusJSON)

	fungusIdINT += 1
	ctx.Getstub().putstate("funguscount",[]byte(strconv.Itoa(fungusIdINT)))


}

func(s *SmartContract) _generateRandomDna(name string) uint {
	nowtime :=time.Now()
	unixtime :=nowtime.Unix()
	data := strconv.Itoa(int(unixtime)) + name
	hash := sha256.New()
	hash.Write([]byte(data))
	dnaHash :=uint(binary.BigEndian.Uint64(hash.Sum(nil)))
	dna := dnaHash % uint(math.Pow(10, 14))
	dna = dna -(dna % 100)
	return dna
}