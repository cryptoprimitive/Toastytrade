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
    $("#noEthAddressWarning").show();
    $("#ethereumAddressOutput").hide();
  }
  else {
    accountInfoPanelVueapp.ethereumAddress = web3.eth.accounts[0];
  }
});
