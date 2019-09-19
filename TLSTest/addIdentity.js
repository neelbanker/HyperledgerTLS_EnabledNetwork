'use strict';

// Bring key classes into scope, most importantly Fabric SDK network class
const fs = require('fs');
const { FileSystemWallet, X509WalletMixin } = require('fabric-network');

// A wallet stores a collection of identities for use
const wallet = new FileSystemWallet('./idwallet');

async function main(){

    // Main try/catch block
    try {

        // define the identity to use
        const cert = fs.readFileSync('./ca.crt').toString();
        const key = fs.readFileSync('./server.key').toString();
        const identityLabel = 'User1@manufacturer.product.com';

        // prep wallet and test it at the same time
        await wallet.import(identityLabel, X509WalletMixin.createIdentity('ManufacturerMSP', cert, key));

    } catch (error) {
        console.log(`Error adding to wallet. ${error}`);
        console.log(error.stack);
    }
}

main().then(()=>{
    console.log('done');
}).catch((e)=>{
    console.log(e);
    console.log(e.stack);
    process.exit(-1);
});