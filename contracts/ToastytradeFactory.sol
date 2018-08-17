pragma solidity ^ 0.4.24;

import "BurnablePayment.sol";

contract ToastytradeFactory {
    address public notifyServiceProvider = 0x0;
    uint public notifyServiceFee = 1;

    event ToastytradeCreated(address toastytradeAddress, string details);

    function newToastytrade(address seller, uint autoreleaseInterval, string _details)
    payable
    external {
        require(msg.value >= notifyServiceFee, "Not enough ether to cover the notify service fee");
        uint valueToForwardToToastytrade = msg.value - notifyServiceFee;

        uint commitThreshold = valueToForwardToToastytrade / 5;

        BurnablePayment toastytrade = (new BurnablePayment).value(valueToForwardToToastytrade)(true, seller, commitThreshold, autoreleaseInterval, "", "");

        emit ToastytradeCreated(address(toastytrade), _details);

        //send the fee to our server
        notifyServiceProvider.transfer(notifyServiceFee);

        //Make sure there's no leftover ether
        assert(address(this).balance == 0);
    }

}
