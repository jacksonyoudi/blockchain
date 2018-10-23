pragma solidity ^0.4.24;

//import "../../v1/this.sol";

contract Adoption {

    address[16] public adopters; //保存领养者的地址

    // 领养宠物
    function adpot(uint petId) public returns (uint) {
        require(petId >= 0 && petId <= 15);

        adopters[petId] = msg.sender;

        return petId;
    }

    // 返回领养的宠物
    function getAdopters() public view returns (address[16]) {
        return adopters;
    }

}


