package chaincode
import ( 
	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"time"
	"strconv"
	"encoding/json"
	"fmt"
)
type SmartContract struct{
	contractapi.Contract
}

type Input struct{
	ReadyTime int64
	Money uint
	Amount uint
	Doctype string 
}

type Output struct{
	ReadyTime int64
	Money uint
	Amount uint
 	Purpose string 
	Doctype string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func(s *SmartContract) Inputmoney(ctx contractapi.TransactionContextInterface, money uint) (*Input, error) {

		nowtime :=time.Now()
		unixtime :=nowtime.Unix()
	
		var amount uint
	
		amountbyte, err := ctx.GetStub().GetState("balance")
		
		if err != nil {
			return nil, err
		}
		if amountbyte == nil {
			amount = 0
		} else {
			amountInt, _ := strconv.Atoi(string(amountbyte))
			amount = uint(amountInt)
		}
		
		amount = amount + money
	
		minputmoney := Input{
			ReadyTime : unixtime,
			Money : money,
			Amount : amount,
			Doctype : "inputmoney", //조회기능에 활용
			// Name : name,
		}
	
		inputmoneyBytes, _ := json.Marshal(minputmoney)
		// carAsBytes, _ := json.Marshal(car)
	
		ctx.GetStub().PutState("inputmoney"+strconv.Itoa(int(unixtime)),inputmoneyBytes)
		ctx.GetStub().PutState("balance",[]byte(strconv.Itoa(int(amount))))
	
		return &minputmoney, nil 
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func(s *SmartContract) Outputmoney(ctx contractapi.TransactionContextInterface, money uint, purpose string) (*Output, error) {

	nowtime :=time.Now()
	unixtime :=nowtime.Unix()

	var amount uint

	amountbyte, err := ctx.GetStub().GetState("balance")
		
	if err != nil {
		return nil, err
	} else {
		amountInt, _ := strconv.Atoi(string(amountbyte))
		amount = uint(amountInt)
	}

	amount = amount - money

	if amount < 0 {
		return nil, fmt.Errorf("Exceeded budget!!")
	}

	moutputmoney := Output{
		ReadyTime : unixtime,
		Money : money,
		Amount : amount,
		Purpose : purpose,
		Doctype : "outputmoney",
	}

	outputmoneyBytes, _ := json.Marshal(moutputmoney)

	ctx.GetStub().PutState("outputmoney"+strconv.Itoa(int(unixtime)),outputmoneyBytes)
	ctx.GetStub().PutState("balance",[]byte(strconv.Itoa(int(amount))))

	return &moutputmoney, nil 
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func(s *SmartContract) GetInputmoney(ctx contractapi.TransactionContextInterface) ([]*Input,error) {
	queryString := `{"selector":{"Doctype":"inputmoney"}}`

	resultsIterator,err := ctx.GetStub().GetQueryResult(queryString)
	if err !=nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var inputs []*Input
	for resultsIterator.HasNext() {
		queryResult,err := resultsIterator.Next()
		if err !=nil {
			return nil, err
		}
		var input Input
		err = json.Unmarshal(queryResult.Value, &input)
		if err !=nil {
			return nil, err
		}
		inputs = append(inputs, &input)
	}
	return inputs, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func(s *SmartContract) GetOutputmoney(ctx contractapi.TransactionContextInterface) ([]*Output,error) {
	queryString := `{"selector":{"Doctype":"outputmoney"}}`

	resultsIterator,err := ctx.GetStub().GetQueryResult(queryString)
	if err !=nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var outputs []*Output
	for resultsIterator.HasNext() {
		queryResult,err := resultsIterator.Next()
		if err !=nil {
			return nil, err
		}
		var output Output
		err = json.Unmarshal(queryResult.Value, &output)
		if err !=nil {
			return nil, err
		}
		outputs = append(outputs, &output)
	}
	return outputs, nil
}