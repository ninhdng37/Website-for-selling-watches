$(document).ready(function() {
  $("#registry").click(function(event) {

    if ($("#password").val() != $("#confirmimg-pass").val()){

      $("#confirmimg-pass-error").html("Confirming password is incorrect!");
      $("#confirmimg-pass-error").show();
      event.preventDefault();
      return;
    } else {
      $("#confirmimg-pass-error").hide();
    }

    if ($("#phone-number").val().length !== 10) {
      event.preventDefault();
      return;
    }   

    var postData = {
        customerName: $("#fullname").val(),
        email: $("#email").val(),
        phoneNumber: $("#phone-number").val(),
        password: $("#password").val()
      };
    $.ajax({
      url:"http://localhost:8888/customer/create",
      method: "POST",
      data: JSON.stringify(postData), // Convert data to JSON format
      contentType: "application/json", // Set content type to JSON
      success: function() {
            $(".input-box").hide();
            $("#registry").hide();
            $(".title").html("Please check your email to verify your account!")
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
        console.error("Error sending email:", error.responseJSON["BadRequest"]);
      }
    });

  });

  $("#fullname").on("keyup", function() {
    var fullname = $(this).val();
    var fullnameError = $("#fullname-error");

    // Regular expression to match only alphabets and spaces
    var alphabetSpaceRegex = /^[A-Za-z\s]+$/;

    if (!alphabetSpaceRegex .test(fullname)) {
        fullnameError.text("Fullname only contains alphabets")
        fullnameError.show();
    } else {
        // If no special characters are found, hide the "fullname-error" element
        fullnameError.hide();
    }    
  });

  $("#email").on("keyup", function() {
    // Regular expression to match email address
    var emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

    var email = $(this).val();
    var emailError = $("#email-error");

    if (!emailRegex.test(email)) {
      emailError.text("Email is invalid!")
      emailError.show();
    } else {
      // Hide the error message if the input is valid
      emailError.hide();
    }
  });

  $("#phone-number").on("keyup", function() {
    var phoneNumber = $(this).val();
    var phoneNumberError = $("#phone-number-error");

    if (phoneNumber.length === 10) {
      phoneNumberError.text("Phone number must contain 10 digits")
      phoneNumberError.hide();
    } else {
      phoneNumberError.show();
    }
  });

  $("#password").on("keyup", function() {
    var passwordRegex = /^(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$/;
    var password = $(this).val();
    var passwordError = $("#pass-error");

    if (password.length < 8 || password.length > 20) {
      passwordError.text('Password have to contain 8 characters at least and maximum 20 characters');
      passwordError.show();
    } else {
      passwordError.hide();
    }

    if (!passwordRegex.test(password)) {
      passwordError.text("Password have to contain at least 1 uppercase letter \
      at least 1 digit and 1 special character");
      passwordError.show();
    } else {
      passwordError.hide();
    }
  });

  $("#confirmimg-pass").on("keyup", function() {
    var confirmimgPass = $(this).val();
    var confirmimgPassError = $("#confirmimg-pass-error");
    var password = $("#password").val();
    if (confirmimgPass != password) {
      confirmimgPassError.text('Confirming password is incorrect!')
      confirmimgPassError.show();
    } else {
      confirmimgPassError.hide();
    }
  });

});
