#！/bin/bash
# 使用org2身份
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.university.cn/peers/peer0.org2.university.cn/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.university.cn/users/Admin@org2.university.cn/msp
export CORE_PEER_ADDRESS=localhost:9051