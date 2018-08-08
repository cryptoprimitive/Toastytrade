window.onload = function() {
  window.app = new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: {
      test: "hiii"
    }
  });
}
