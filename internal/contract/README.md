## Run codegen 

```shell
abigen --abi internal/contract/erc20.abi --pkg contract --type ERC20 --out internal/contract/erc20.go
abigen --abi internal/contract/erc1271.abi --pkg contract --type ERC1271 --out internal/contract/erc1271.go
abigen --abi internal/contract/factory.abi --pkg contract --type Factory --out internal/contract/factory.go
abigen --abi internal/contract/pool.abi --pkg contract --type Pool --out internal/contract/pool.go
abigen --abi internal/contract/rfq.abi --pkg contract --type RFQ --out internal/contract/rfq.go
```
