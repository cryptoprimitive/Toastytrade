window.addEventListener('load', function() {

  window.registerVueapp = new Vue({
    delimiters: ['${', '}'],
    el: '#registerVueapp',
    data: {}
  });

});

function submitRegister() {
  var ethAddress = accountInfoPanelVueapp.ethereumAddress;
  var email = $('#emailInput').val();

  $.post("/doRegister", {ethAddress: ethAddress, email: email})
}
