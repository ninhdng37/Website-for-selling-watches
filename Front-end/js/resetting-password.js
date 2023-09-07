$(document).ready(function(){

    $("#reset-password").click(function (event) {
      event.preventDefault();
        var postData = {
            password: $("#password").val()
          };

        $.ajax({
          url:"http://localhost:8888/customer/reset-password",
          method: "PUT",
          data: JSON.stringify(postData), // Convert data to JSON format
          contentType: "application/json", // Set content type to JSON
          success: function() {
                $(".content").hide();
                $(".title").html('Reset password successfully. Please click <a href="http://localhost:8888/customer/login-form">here</a>" to login')
          },
          error: function(error) {
            if (error.responseJSON["InternalServerError"] != undefined) {
              alert(error.responseJSON["InternalServerError"])
              return
            }
    
            if (error.responseJSON["BadRequest"] != undefined) {
              // $(".title").html(response["BadRequest"])
              alert(error.responseJSON["BadRequest"])
              return
            }
          }
        });
    });

    $("#password").on("keyup", function() {
        var passwordRegex = /^(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$/;
        var password = $(this).val();
        var passwordError = $("#password-error");
    
        if (password.length < 8 || password.length > 20) {
          passwordError.text('Password have to contain 8 characters at least and maximum 20 characters');
          passwordError.show();
        } else {
          passwordError.hide();
        }
    
        if (!passwordRegex.test(password)) {
          passwordError.text("Password have to contain at least one uppercase letter \
          at least one digit and one special character");
          passwordError.show();
        } else {
          passwordError.hide();
        }
    });
    
    $("#confirming-password").on("keyup", function() {
        var confirmimgPass = $(this).val();
        var confirmimgPassError = $("#confirming-password-error");
        var password = $("#password").val();
        if (confirmimgPass != password) {
          confirmimgPassError.text('Confirming password is incorrect!')
          confirmimgPassError.show();
        } else {
          confirmimgPassError.hide();
        }
    });

})