# Storing-Marksheets-on-Blockchain

## Overview
Our project is to create a web application to store educational certificates, using
a blockchain network. <br>
Educational certificates issued to students will be stored by authorities directly
on a blockchain network, which can be accessed by students from the website. The blockchain network will serve as an immutable ledger of certificates, which can be updated only by authorities, and hence will prevent fraud of tampering and using fake certificates.
The distributed and decentralized nature of the network ensures that it even more trusted, as no central organization can tamper with the marksheets.  <br>
Additionally, students can easily access this verified copy from any part of the world, in absence of the original hard copy. This is useful while applying for higher education or jobs, in order to be able to produce an authentic copy which can be trusted by the authorities.  <br>
## Blockchain Network
### Description
Our blockchain network will store the marksheets as JSON objects. We have used <b>Hyperledger Fabric</b> for the blockchain architecture. The marksheets will form an immutable ledger which cannot be tampered with. Since Fabric is a permissioned network, only the admins will have the ability to add marksheets. <br>
Through the API, a transaction request is sent to the endorsing peers. The endorsing peers will validate the transaction and will reach a consensus whether the transaction is valid. In case of <b>viewing marksheets</b>, i.e., querying the ledger, the marksheets will be returned as response to the API in case of consensus. For <b>adding or modifying marksheets</b>, the application will check if the endorsement policy is fulfilled and submit the transaction to the Ordering Service, which will arrange the transactions in chronological order and send the block of marksheets to all the leader peers in the organisation. The leader peers will deliver the block to all the peers in their organisation through gossip protocol, using the channel of the respective organisation. The ledger will then be updated, and the API will be notified once the marksheet is successfully added, and this will be reflected on the website. 
The organisation will be all the centres responsible for updating the blockchain, such as the NIC centres, education boards etc. <br>
We have deployed the chaincode, which is a technical container of a group of related smart contracts, which will automate the transactions and ensure they are all following the same rules. Our smart contract stores the marksheet data of each student as a JSON object.

### Tools used
1. Hyperledger Fabric
2. Docker
3. GoLang to write the chaincode 
### Working
We changed the chaincode at Hyperledger Fabric according to the requirements of our certificate and
then deployed it into the network. Marksheet data will be stored as a JSON
object. Our network currently consists of 2 organisations with 1 peer each and
1 ordering service with 1 peer. We linked the blockchain to the API which will be
able to request transactions. <br>
- clone this directory
- open terminal in this directory 
- make sure .sh file in this directory has excutable permission given if not "chmod +x ./createchannel.sh"
- run : "go run main.go"
- website should be hosted on localhost:8080/form
- make sure you have the fabric-samples cloned
- cd into that directory
- run the command: "./network.sh up"
- website ready to use.
