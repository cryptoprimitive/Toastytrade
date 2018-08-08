window.addEventListener('load', function() {
  window.accountInfoPanelVueapp = new Vue({
    delimiters: ['${', '}'],
    el: '#accountInfoPanelVueapp',
    data: {
      ethereumAddress: "",
      email: ""
    }
  });
});
