window.addEventListener('load', function() {

  window.accountInfoPanelVueapp = new Vue({
    delimiters: ['${', '}'],
    el: '#accountInfoPanelVueapp',
    data: {
      ethereumAddress: "",
      email: "loading..."
    }
  });

  //Try to load Ethereum account from web3; display warning on fail.
  if (web3.eth.accounts.length == 0) {
    $("#noEthAddressWarning").show();
    $("#ethereumAddressOutput").hide();
  }
  else {
    accountInfoPanelVueapp.ethereumAddress = web3.eth.accounts[0];
    $.post("/getUserEmail", {ethAddress: web3.eth.accounts[0]}, function(data) {
      if (data != "") {
        window.accountInfoPanelVueapp.email = data;
      }
      else {
        //empty string return indicates no email was found.
        window.accountInfoPanelVueapp.email = "No email found!";
      }
    });
  }
});
