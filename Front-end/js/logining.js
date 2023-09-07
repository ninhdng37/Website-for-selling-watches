function sendRequestLogin() {
  var postData = {
    email: $("#email").val(),
    password: $("#password").val()
  };
  console.log('abc')
  $.ajax({
      url:"http://localhost:8888/customer/login",
      method: "POST",
      data: JSON.stringify(postData), // Convert data to JSON format
      contentType: "application/json", // Set content type to JSON
      success: function() {
        var randomQueryParam = Date.now();
        window.location.href = "http://localhost:8888/customer/home?nocache=" + randomQueryParam;
      },
      error: function(error) {
        if (error.responseJSON["InternalServerError"] != undefined) {
          $("form .error").text(error.responseJSON["InternalServerError"]);
          return
        }

        if (error.responseJSON["BadRequest"] != undefined) {
          $("form .error").text(error.responseJSON["BadRequest"]);
          $("form .error").show();
        } else {
          $("form .error").hide();
        }
      }
    });
}

$(document).ready(function() {

    $("#login").click(function() {
        sendRequestLogin()
    });

    $("#email").keypress(function(event) {
      if (event.which === 13) {
        event.preventDefault();
        sendRequestLogin();
      }
    });

    $("#password").click(function(event) {
      if (event.which === 13) {
        event.preventDefault();
        sendRequestLogin();
      }
    });

    $("#login-form").submit(function(event) {
        event.preventDefault();
    });

    $("#forgot-password").click(function() {
        $.ajax({
            url:"http://localhost:8888/customer/forgot-password",
            method: "GET",
            data: JSON.stringify(postData), // Convert data to JSON format
            contentType: "application/json", // Set content type to JSON
            success: function() {
            },
            error: function() {

            }
          });
    });

});