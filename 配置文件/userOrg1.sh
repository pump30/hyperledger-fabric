#！/bin/bash
# 使用org1身份
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.university.cn/peers/peer0.org1.university.cn/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.university.cn/users/Admin@org1.university.cn/msp
export CORE_PEER_ADDRESS=localhost:7051