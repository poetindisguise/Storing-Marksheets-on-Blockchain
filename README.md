# Storing-Marksheets-on-Blockchain

## Overview
Our project is to create a web application to store educational certificates, using
a blockchain network. <br>
Educational certificates issued to students will be stored by authorities directly
on a blockchain network, which can be accessed by students from the website. The blockchain network will serve as an immutable ledger of certificates, which can be updated only by authorities, and hence will prevent fraud of tampering and using fake certificates.
The distributed and decentralized nature of the network ensures that it even more trusted, as no central organization can tamper with the marksheets.  <br>
Additionally, students can easily access this verified copy from any part of the world, in absence of the original hard copy. This is useful while applying for higher education or jobs, in order to be able to produce an authentic copy which can be trusted by the authorities.  <br>
## Hyperledger Fabric
Hyperledger fabric is an open-source, enterprise-graded permission distributed
ledger technology platform. It is used majorly in enterprise contexts and has
differentiating capabilities over popular distributed ledger or blockchain
platforms.
### Hyperledger CA:
The Hyperledger fabric CA is the default Certificate Authority component 
which issues PKI-based certificates to network member organizations and their
users. The CA issues one root certificate to each member and one enrollment
certificate to each authorized user.
## Key Terms

### 1.Peer
Peer nodes are the basic network entities, each of which can hold copies of
ledger and smart contracts. They maintain the ledger and depending on the type
of peer node, can perform operations to the ledger. Peers are owned and
maintained by members.
### 2.Member
A blockchain consists of one or more members, which are unique and legally
separate entities in the network. A member, such as an organization, will run
one or more separate peer nodes. Network components will be linked to a
member.
### 3.Ledger
The ledger is a channel’s chain and current state data which is maintained by
each peer on the channel.
### 4.Chaincode
<b>Smart contracts</b> are lines of code; domain specific programs, that define the
transaction logic of an object in the blockchain. A chaincode is a technical
container of a group of related smart contracts. The chaincode governs how
smart contracts are packaged for deployments. Multiple smart contracts can
be defined in the chaincode, and when the chaincode is deployed to the
blockchain, all the smart contracts within it are made available.
Chaincode execution results in a set of key-value writes(write set) that can be
submitted to the network and applied to the ledger on all peers.

### 5.Ordering service
The ordering service consists of ordered nodes, which arrange the transactions
in a chronological fashion. The ordering service exists independent of the peer
processes, and the order transactions on first-come-first-serve basis for all the
channels. The ordering service is common to the whole network, and contains
the cryptographic identity material linked to each member.
### 6.Channel
A channel is a component of a private blockchain which allows for data
isolation and confidentiality. A ledger can be made channel specific, and will be
only shared across the peers in that channel. To interact with a channel,
transacting parties must be properly authenticated to it. A channel is defined
by members(organization), anchor peers per member, the shared ledger,
chaincode applications and the ordering service nodes.
### 7.Consensus
A consensus is a broader term overarching the whole network. The peer nodes
come to an agreement about an order, which serves to maintain correctness of
the transactions. Common consensus algorithms include proof of work and
proof of stake. In Fabric, Ordering Service is responsible for consensus (via Kafka).
### 8.Types of Peers
1. Committing Peer: Maintains the ledger and state, commits transactions and
may hold smart contracts
2. Endorsing Peer: It is a specialized committing peer that receives a
transaction proposal for endorsement, responds granting or denying
endorsement. Must hold smart contract.
3. Ordering node: Approves the inclusion of transaction blocks into the ledger
and communicates with committing and endorsing peer nodes. Doesn’t hold
the smart contract and ledger. <br>
<img src = "https://user-images.githubusercontent.com/66271769/124391488-78267900-dd0e-11eb-9306-ca27cdb84904.jpeg" width= "550" height = "280">


## Blockchain Network
### Description 
Our blockchain network will store the marksheets as JSON objects. We have used <b>Hyperledger Fabric</b> for the blockchain architecture. The marksheets will form an immutable ledger which cannot be tampered with. Since Fabric is a permissioned network, only the admins will have the ability to add marksheets. <br>
Through the API, a transaction request is sent to the endorsing peers. The endorsing peers will validate the transaction and will reach a consensus whether the transaction is valid. In case of <b>viewing marksheets</b>, i.e., querying the ledger, the marksheets will be returned as response to the API in case of consensus. For <b>adding or modifying marksheets</b>, the application will check if the endorsement policy is fulfilled and submit the transaction to the Ordering Service, which will arrange the transactions in chronological order and send the block of marksheets to all the leader peers in the organisation. The leader peers will deliver the block to all the peers in their organisation through gossip protocol, using the channel of the respective organisation. The ledger will then be updated, and the API will be notified once the marksheet is successfully added, and this will be reflected on the website. 
The organisation will be all the centres responsible for updating the blockchain, such as the NIC centres, education boards etc. <br>
We have deployed the chaincode, which is a technical container of a group of related smart contracts, which will automate the transactions and ensure they are all following the same rules. Our smart contract stores the marksheet data of each student as a JSON object. Structure of our basic certificate: <br>
<img src = "https://user-images.githubusercontent.com/66271769/124385207-7353cc00-dcf2-11eb-8d48-fd5b5df21565.jpeg" width= "300" height = "150">

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

## Website
### Description
Our front-end website consists of a home page, from which students can view
their marksheets and admin organizations can login. The students have to enter their
credentials: Board, Name of Examination, Year of Examination and Roll number
in a form, which will send a request to the API and retrieve the marksheet from
the blockchain and display it to the student. The admins have to enter their login
credentials, after which they will be able to add, modify or view marksheets. On
entering the board name, year, school and exam, they will receive a list of students in
that school with the respective roll number, and once the admin uploads the
marksheet, the name will be flagged. The data of the boards and schools will be
stored in an SQLite relational database and the form is populated using the
same. The database is queried by the API to give the school wise data. <br>
<br>
<img src = "https://user-images.githubusercontent.com/66271769/124385712-93848a80-dcf4-11eb-8c2e-4b9e39d7f81c.jpeg" width= "400" height = "150"> <br>
<br>
<img src = "https://user-images.githubusercontent.com/66271769/124385714-94b5b780-dcf4-11eb-9e6b-6e306f612624.jpeg" width= "400" height = "200">

### Tools used
We have used HTML, CSS and Javascript to design the website. The API
endpoints are called using HTTP get and post requests, using AJAX and JQuery.
### Working
We have made the website with the necessary forms. We have linked the
forms to the database and to the API, so that students' and schools' data can be retrieved from
the database and the marksheets can be added to the blockchain and viewed
on request. A comfortable and easy-to-use User Interface has been prepared.
## API
### Description
The application programming interface consists of 3 end points: <b>the RDBMS</b>,
which will consist of the board, school and student data, <b>the Blockchain
network</b>, which will have the marksheets stored as JSON objects, and <b>the
website for users</b>. API Endpoints: <br>
<img src = "https://user-images.githubusercontent.com/66271769/124386053-fde9fa80-dcf5-11eb-80db-85e7fd43f82d.jpeg" width= "450" height = "150">

### Tools used
The API has been written in GoLang, using the Gin framework. GORM, which is
a very powerful Object Relational Mapping library for GoLang, has also been
used.
### Working
We firstly learnt the basics of GoLang, understood how the API will function,
and learnt the basics of the Gin framework. After deciding the endpoints etc.,
we implemented it. HTTP requests are being sent and received using ajax
requests.<br>
The data for the dropdown forms can be fetched from the database using the
API, the values entered will be sent back to the API on submission which will
accordingly fetch the marksheets from the blockchain, or add new marksheets
to the blockchain as JSON.
## Admin Access
Apart from all the above milestones, we have also prepared an application for
the Admin. It authorizes the Admin to add nodes, turning up the network and
adding channels to the blockchain. Upon deployment at a national level, more
functionalities can be added into it.
A snapshot of the Admin’s Frontend: <br>
<img src = "https://user-images.githubusercontent.com/66271769/124385623-2a9d1280-dcf4-11eb-9825-be6070a7c584.jpeg" width= "350" height = "200">
 ## How to Run                                                                                                                                             
- install Golang, verson 1.14 +
- install Gin
- set up the fabric samples network 
- open the teminal and cd into NIC_Project directory
- run: "go run main.go"
- the website should be hosted at localhost:8080 <br>
Chaincode of the blockchain network is being invoked using BASH scripts. For example, 
this will return all the marksheets stored in the blockchain: <br>
<img src = "https://user-images.githubusercontent.com/66271769/124385934-84520c80-dcf5-11eb-9ed2-11909f82aeed.jpeg" width= "650" height = "200">
