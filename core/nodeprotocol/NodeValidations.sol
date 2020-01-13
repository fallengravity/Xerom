pragma solidity 0.4.25;
pragma experimental ABIEncoderV2;

contract  NodeValidations {

    mapping(address => signedNodeValidation) public lastNodeActivity;
    
    struct signedNodeValidation {
        address nodeAddress;
        bytes[] signatures;
        uint signatureCount;
	bytes signatureBlockHash;
        bytes publicKey;
        uint blockHeight;
    }
    
    constructor() public {
    }
    
    function nodeCheckIn(bytes[] signatures, bytes publicKey, bytes signatureBlockHash) public {
        signedNodeValidation memory validation = signedNodeValidation({nodeAddress:tx.origin, signatures:signatures, signatureCount:signatures.length, signatureBlockHash:signatureBlockHash, publicKey:publicKey, blockHeight:block.number});
        lastNodeActivity[tx.origin] = validation;
    }
    
    function getSignatures(address nodeAddress) public view returns (bytes[]) {
        return lastNodeActivity[nodeAddress].signatures;
    }
}
