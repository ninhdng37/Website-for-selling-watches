$(document).ready(function(){
    $("#email").on("keyup", function(){
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

    $("#send-email").click(function(event){
        event.preventDefault();
        var postData = {
            email: $("#email").val()
          };
        $.ajax({
          url:"http://localhost:8888/customer/send-email",
          method: "POST",
          data: JSON.stringify(postData), // Convert data to JSON format
          contentType: "application/json", // Set content type to JSON
          success: function() {
                $(".title").text("Please check your email verify your account!");
                $(".content").hide();
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

})