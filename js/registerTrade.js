function submitRegisterTrade() {
  var address = $('#addressInput').val();

  $.post("/doRegisterTrade", {toastytradeAddress: address})
}
