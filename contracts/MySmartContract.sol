pragma solidity ^0.8.0;

contract MySmartContract {
    uint256 balance = 100;
    address public admin;

    constructor() {
        admin = msg.sender;
        updateBalance();
    }

    receive() external payable {
        updateBalance();
    }

    function updateBalance() internal {
        balance += msg.value;
    }

    function Withdrawl(uint256 _amt) public{
        balance = balance - _amt;
    }

    function Deposite(uint256 amt) public returns (uint256) {
        return balance = balance + amt;
    }

    function Balance() public view returns (uint256) {
        return balance;
    }

}