#!/bin/bash

C_YELLOW='\033[1;33m'
C_BLUE='\033[0;34m'
C_RESET='\033[0m'

# subinfoln echos in blue color
function infoln() {
  echo -e "${C_YELLOW}${1}${C_RESET}"
}

function subinfoln() {
  echo -e "${C_BLUE}${1}${C_RESET}"
}

# add PATH to ensure we are picking up the correct binaries
export PATH=${HOME}/fabric-samples/bin:$PATH
export FABRIC_CFG_PATH=${PWD}/network/config

# Chaincode config variable

# CHANNEL_NAME="mychannel"
CC_NAME="downchaincode"
CHANNEL_NAME="mychannel"
ORDERER_CA=${PWD}/network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
PEER_CONN_PARMS="--peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"


export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

## TEST1 : Invoking the chaincode
infoln "TEST1 : Invoking the chaincode"
set -x
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CC_NAME} $PEER_CONN_PARMS -c '{"function":"Inputmoney","Args":["5000"]}' >&log.txt
{ set +x; } 2>/dev/null
cat log.txt
sleep 3

## TEST2 : Query the chaincode
infoln "TEST2 : Query the chaincode"
set -x
peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["GetInputmoney"]}' >&log.txt
{ set +x; } 2>/dev/null
cat log.txt

## TEST3 : Invoking the chaincode
infoln "TEST3 : Invoking the chaincode"
set -x
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CC_NAME} $PEER_CONN_PARMS -c '{"function":"Outputmoney","Args":["5000","Test"]}' >&log.txt
{ set +x; } 2>/dev/null
cat log.txt
sleep 3

## TEST4 : Query the chaincode
infoln "TEST4 : Query the chaincode"
set -x
peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["GetOutputmoney"]}' >&log.txt
{ set +x; } 2>/dev/null
cat log.txt


