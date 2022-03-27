#！/bin/bash
# 使用org3身份
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.university.cn/peers/peer0.org3.university.cn/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.university.cn/users/Admin@org3.university.cn/msp
export CORE_PEER_ADDRESS=localhost:11051