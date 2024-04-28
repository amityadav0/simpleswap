# simpleswap
**simpleswap** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).
● Liquidity providers can provide liquidity in any of stablecoins whitelisted on module
params
● Liquidity providers get share token after providing liquidity
● Tokens provided by liquidity providers are put on module account
● Users can swap any coins as long as liquidity's available
● Fee percentage (e.g. 0.3%) is configured on params
● Swap fee's given to liquidity providers
● Liquidity providers can withdraw liquidity and collected fees at any time

## Get started
- ignite version: `v0.27.1`
```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Interacting with simpleswap module

- On another terminal use the commands given below

#### Create Pool (whitlelist)
    - simpleswapd  tx simpleswap create-pool "1000eth,1000weth" "6,6" 100 --from alice
    - Note: Can only be done by owner (cosmos14tpfntxwkv30d6re3hrk8ny72r50vpalkapy2k)

#### Add liquidity
    - simpleswapd  tx simpleswap add-liquidity poole546c64ea2770e4bf76808611963b12673ed7af8d2a1194ce584ae9a5c987255 "1000eth,1000weth" --from alice
    - Note: pool_id can be queried using simpleswapd query or a event is emitted with pool id when a new pool is created.
    - By default pool id for weth and eth will be `poole546c64ea2770e4bf76808611963b12673ed7af8d2a1194ce584ae9a5c987255`

#### Swap
    - simpleswapd tx simpleswap swap poole546c64ea2770e4bf76808611963b12673ed7af8d2a1194ce584ae9a5c987255 100weth 1eth --from alice

#### Withdraw
    - simpleswapd tx simpleswap withdraw cosmos14tpfntxwkv30d6re3hrk8ny72r50vpalkapy2k poole546c64ea2770e4bf76808611963b12673ed7af8d2a1194ce584ae9a5c987255 200poole546c64ea2770e4bf76808611963b12673ed7af8d2a1194ce584ae9a5c987255 --from alice

### Query

- `simpleswapd query simpleswap pools all 10`
- This will return pool data

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/simpleswap@latest! | sudo bash
```
`username/simpleswap` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
