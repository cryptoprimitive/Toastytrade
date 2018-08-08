window.addEventListener('load', function() {

  window.accountInfoPanelVueapp = new Vue({
    delimiters: ['${', '}'],
    el: '#accountInfoPanelVueapp',
    data: {
      ethereumAddress: "",
      email: ""
    }
  });

  if (web3.eth.accounts.length == 0) {
    $("#no-eth-address-warning").show();
    $("#ethereum-address-output").hide();
  }
  else {
    accountInfoPanelVueapp.ethereumAddress = web3.eth.accounts[0];
  }
});
