# kvs
**kvs** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## TX API 

DataProposal - Propose data to store in kv storage
data-proposal [key] [value]

AddressRegistration - Registration of three addresses to confirm proposal
address-registration [addresses]

DataConfirmation - Confirmation of the proposed data at the trusted module address
data-confirmation [key]


## QUERY API 

ShowData - shows successfully confirmed data
show-data [index]

ListData - list of all successfully confirmed data
list-data

ShowProposal - shows a Proposal
show-proposal [index]

ListProposal - list all Proposal
list-proposal

ShowAcl - shows acl
show-acl