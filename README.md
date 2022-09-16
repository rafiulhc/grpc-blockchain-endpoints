<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->




<!-- PROJECT LOGO -->
<br />
<div align="center">


  <h3 align="center">Front-running and malicious node/validator resistant smart contract, written in Solidity to register vanity names/domains on EVM Blockchains</h3>


</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#inspiration">Inspiration</a></li>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#testing">Testing</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>


### Smart Contract Deplyed to Ethereum Rinkeby Testnet
Contract Address


    0x6b3198d1a1b29156A2a99f01948e298A64bb86a3






![deployed_to_rinkeby_testnet](https://user-images.githubusercontent.com/68476971/169980483-53b8974d-aa37-4c77-bbfb-2ef5708cf5e9.png)


### Smart Contract Registered to Oracle testnet for checking expiration of name registration after every 24 hours automatically

Step: 1

![Chainlink_upkeep](https://user-images.githubusercontent.com/68476971/169981854-d403bdfa-a1be-4ade-8e2d-ab7b1184b0fd.png)

Step: 2

![Chainlink](https://user-images.githubusercontent.com/68476971/169982003-f7a77a4e-9352-4761-8a15-6fb421b8817f.png)

Step: 3

![Active Oracle](https://user-images.githubusercontent.com/68476971/169982149-bdfb5a03-f15a-4e47-8e5a-e1cc67f545e7.png)


<!-- ABOUT THE PROJECT -->
## Inspiration

Blockchain technology enables decentralized applications (DApps) or smart contracts. Function calls (or transactions) to the DApp are processed by a decentralized network. Transactions are finalized in stages: they (generally) first relay around the network, then are selected by a miner and put into a valid block, and finally, the block is well-enough incorporated in the blockchain that is unlikely to be changed. Front-running is an attack where a malicious node observes a transaction after it is broadcast but before it is finalized, and attempts to have its own transaction confirmed before or instead of the observed transaction.

In this application I followed Consensys provided guideline to prevent frontrun and malicious activities by validator. Suggestion is to store
the keccak256 hash of the data in the first transaction, then reveal the data and verify it against the hash in the second transaction.


## About The Project

Front-running and malicious node/validator resistant smart contract, written in Solidity to register vanity names/domains on EVM Blockchains.
An unregistered name can be registered for a certain amount of time by locking a certain balance of an account. After the registration expires,
 the account loses ownership of the name and his balance is unlocked. The registration can be renewed by making an on-chain call to keep the
 name registered and balance locked. The fee to register the name depends directly on the size of the name. Also, a malicious node/validator
 should not be able to front-run the process by censoring transactions of an honest user and registering its name in its own public address.



<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Solidity](https://docs.soliditylang.org/en/v0.8.14/)
* [Node.js](https://nodejs.org/en/docs/)
* [Openzeppelin](https://www.openzeppelin.com/)
* [Chainlink-Oracle](https://docs.chain.link/docs/chainlink-keepers/register-upkeep/)
* [Truffle Suite](https://trufflesuite.com/)
* [Hardhat](https://hardhat.org/)
* [Ehers.js](https://docs.ethers.io/v5/)
* [Chai](https://www.npmjs.com/package/chai) - as testing library.



<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation



1. Clone the repo
   ```sh
   git clone https://github.com/rafiul86/Name-Registry-Smart-Contract.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```
<p align="right">(<a href="#top">back to top</a>)</p>

Change the secret variables as process.env.SECRET_KEY and process.env.ALCHEMY provided by API provider if deploy to mainnet (Ethereum, Binance, Fantom, Avalanche etc.) or remote testnet(rinkeby, ropsten, goerli etc.)

```npm install
  npx hardhat compile
  npx hardhat run --network rinkeby scripts/deploy.js
  ```

![Secret Variables](https://user-images.githubusercontent.com/68476971/169951589-da24b489-0cb6-44f8-a1fb-f9f02afca154.png)


<!--Testing-->
## Testing

  ```npm install
  npx hardhat compile
  npx hardhat run --network localhost scripts/deploy.js
  npx hardhat test
  ```


After running the commands in terminal should show the tests results....


![Tests](https://user-images.githubusercontent.com/68476971/169960734-8bad34dc-93ca-40b8-843d-e8e48992d94a.png)


<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

None

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Rafiul Hasan - [Linkedin](https://www.linkedin.com/in/hrafiul/)
               [Twitter](https://twitter.com/r_hasan_c)
               - rafiulhasan86@gmail.com

Project Link: [https://github.com/rafiul86/Name-Registry-Smart-Contract](https://github.com/rafiul86/Name-Registry-Smart-Contract)

<p align="right">(<a href="#top">back to top</a>)</p>
