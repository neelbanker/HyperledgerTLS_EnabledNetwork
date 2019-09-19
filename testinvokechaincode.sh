export CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
export ORG1_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/manufacturer.product.com/users/Admin@manufacturer.product.com/msp
export ORG1_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt
export retailer_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/retailer.product.com/users/Admin@retailer.product.com/msp
export retailer_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/retailer.product.com/peers/peer0.retailer.product.com/tls/ca.crt
export ORDERER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem

export CORE_PEER_LOCALMSPID=ManufacturerMSP
export CORE_PEER_ADDRESS=peer0.manufacturer.product.com:7051
export CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH}
export CORE_PEER_TLS_ROOTCERT_FILE=${ORG1_TLS_ROOTCERT_FILE}

export CHANNEL_NAME=mychannel
peer channel create -o orderer.product.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem
peer channel join -b mychannel.block
peer channel update -o orderer.product.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/ManufacturerMSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem

peer chaincode install -n fabcar -v 1.0 -p github.com/chaincode/fabcar/go
peer chaincode install -n newdistributor -v 1.0 -p github.com/chaincode/newdistributor

#fabcar chaincode commands
peer chaincode instantiate -o orderer.product.com:7050 -C mychannel -n fabcar -v 1.0 -c '{"Args":[]}' -P "OR('ManufacturerMSP.member')" --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem --peerAddresses peer0.manufacturer.product.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt
peer chaincode invoke -o orderer.product.com:7050 -C mychannel -n fabcar -c '{"function":"initLedger","Args":[]}' --waitForEvent --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem --peerAddresses peer0.manufacturer.product.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt
peer chaincode invoke -o orderer.product.com:7050 -C mychannel -n fabcar -c '{"function":"queryAllCars","Args":[]}' --waitForEvent --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem --peerAddresses peer0.manufacturer.product.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt

#cama chaincode commands
peer chaincode instantiate -o orderer.product.com:7050 -C mychannel -n newdistributor -v 1.0 -c '{"Args":[]}' -P "OR('ManufacturerMSP.member')" --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem --peerAddresses peer0.manufacturer.product.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt
peer chaincode invoke -o orderer.product.com:7050 -C mychannel -n newdistributor -c '{"Args":["newTransaction","1","5ce6336abdf8b650a4e6343c","Africa","Pearson","5499","theatres","2000","true","false","false","5","dollar","$19000.00","12","$2035.00","$560.00","12%","true","$50000.00","test","test","test","test"]}' --waitForEvent --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem --peerAddresses peer0.manufacturer.product.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt
peer chaincode invoke -o orderer.product.com:7050 -C mychannel -n newdistributor -c '{"Args":["showTransaction","1"]}' --waitForEvent --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/product.com/orderers/orderer.product.com/msp/tlscacerts/tlsca.product.com-cert.pem --peerAddresses peer0.manufacturer.product.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.product.com/peers/peer0.manufacturer.product.com/tls/ca.crt