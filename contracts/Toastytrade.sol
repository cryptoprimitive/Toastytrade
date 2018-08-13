pragma solidity ^ 0.4.24;

import "BurnablePayment.sol";

contract ToastytradeFactory {
    address public notifyServiceProvider = 0x0;
    uint public notifyServiceFee = 1;

    address[] public toastytrades;

    event ToastytradeCreated(address toastytradeAddress, address BurnablePaymentAddress, string offerDetails);

    function newToastytrade(address seller, uint autoreleaseInterval, string _details)
    payable
    external {
        require(msg.value >= notifyServiceFee, "Not enough ether to cover the notify service fee");
        uint valueToForwardToToastytrade = msg.value - notifyServiceFee;

        Toastytrade toastytrade = (new Toastytrade).value(valueToForwardToToastytrade)(seller, autoreleaseInterval, _details);

        emit ToastytradeCreated(address(toastytrade), toastytrade.BP(), _details);

        toastytrades.push(address(toastytrade));

        //send the fee to our server
        notifyServiceProvider.transfer(notifyServiceFee);

        //Make sure there's no leftover ether
        assert(address(this).balance == 0);
    }

}

contract Toastytrade {
    address public BP;

    string public details;

    constructor(address seller, uint autoreleaseInterval, string _details)
    payable {
        //Prepare vars and create BP
        uint commitThreshold = msg.value / 3;

        BP = (new BurnablePayment).value(msg.value)(true, seller, commitThreshold, autoreleaseInterval, "", "");

        details = _details;
    }
}
