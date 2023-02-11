# testhellochain
**testhellochain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Install Ignite CLI

```
ignite version while scaffolding the chain in repo for first time was v0.24.0
install same ignite version by following 2 commands

curl https://get.ignite.com/cli@v0.24.0 | bash
sudo mv ignite /usr/local/bin/

```

## Install Python and FastApi 

```
move to api folder
install python >= 3.9
after that run commands

pip install "fastapi[all]"
uvicorn main:app --reload
```

## Get started

```
to start testhellochain:
ignite chain serve

to call venue api: (it calls external api and returns name and category of cryptoatm venuue at random index between 0 to 10)
testhellochaind q testhellochain callapi

to create venue: (can be called to add data to blockchain)
testhellochaind tx testhellochain create-venue foo(category) bar(name) --from alice

to show venues: 
testhellochaind q testhellochain showvenues

to get geolocations using Dids:
testhellochaind q testhellochain getgeolocations 1,2,3

note: here string 1,2,3 are driver ids separated with comas. this query intakes string and outputs array of strings (locations)


to call estimate transaction:
testhellochaind tx estimator estimate (start) (end) --from alice

to list all api data from blockchain there are 3 queries respectively:
testhellochaind q estimator list-api-data
testhellochaind q estimator list-api-count-map 
testhellochaind q estimator show-api-count-map address_of_dapp



```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

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
curl https://get.ignite.com/username/testhellochain@latest! | sudo bash
```
`username/testhellochain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
