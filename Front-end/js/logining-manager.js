function sendRequestLogin() {
  var postData = {
    employeeId: parseInt($("#email").val()),
    password: $("#password").val()
  };
  
  $.ajax({
      url:"http://localhost:8888/manager/login",
      method: "POST",
      data: JSON.stringify(postData), // Convert data to JSON format
      contentType: "application/json", // Set content type to JSON
      success: function(response) {
        localStorage.setItem('manager', JSON.stringify(response));
        if (response.position == 'TP'){
          window.location.href = "http://localhost:8888/manager/home";
        }
        else {
          alert("You can not access this page!");
        }
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
    var manager = localStorage.getItem('manager');
    // var manager = JSON.parse(managerString);
    manager = "";
    localStorage.setItem('manager', manager);

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
            url:"http://localhost:8888/manager/forgot-password",
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