npx hardhat export-abi
npx hardhat flatten contracts/$1.sol > resources/flattened/$1.sol
docker run -v $PWD:/sources ethereum/solc:0.8.9 --ir-optimized --optimize --optimize-runs=200 --bin /sources/contracts/$1.sol --include-path /sources/node_modules/ --base-path /sources -o /sources/$1.bin --overwrite
abigen --abi=abi/contracts/$1.sol/$1.json --pkg=$1 --out=go-contracts/$1.go --bin $1.bin/$1.bin
