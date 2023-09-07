var employeeIDForDelete = 0;
function fillEmployeeData(response) {
    $("#employee__table tbody tr").remove();
    $.each(response, function(_, employee) {        
        var newRow = $("<tr>");
        newRow.append(`<td>${employee.employeeId}</td>`);
        newRow.append(`<td>${employee.fullname}</td>`);
        newRow.append(`<td>${employee.identityNumber}</td>`);
        newRow.append(`<td>${employee.position === 'TP' ? 'Manager' : 'Employee'}</td>`);
        newRow.append(`<td>${employee.email}</td>`);
        newRow.append(`<td>${employee.phoneNumber}</td>`);
        
        var actions = `<td class="text-right">
                            <div class="dropdown dropdown-action">
                                <a href="#" class="action-icon dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="fas fa-ellipsis-v ellipse_color"></i></a>
                                <div class="dropdown-menu dropdown-menu-right">
                                    <a class="dropdown-item edit__employee__link" href="#">
                                        <i class="fas fa-pencil-alt m-r-5"></i> Edit
                                    </a>
                                    <a class="dropdown-item delete__employee__link" href="#" >
                                        <i class="fas fa-trash-alt m-r-5"></i> Delete
                                    </a>
                                </div>
                            </div>
                        </td>`;
        newRow.append(actions);
        
        $("#employee__table tbody").append(newRow);
    });
}

$(document).ready(function() {
    const managerString = localStorage.getItem('manager');
    if (managerString === "") {
        alert("Please login to access this page!");
        window.location.href = "http://localhost:8888/manager/login-form";
    }
    const manager = JSON.parse(managerString);

    $('.user-text h6').text(manager.fullname);

    $('#all__employee__link').click(function(event) {
        event.preventDefault();
        
        $.ajax({
            url: 'http://localhost:8888/manager/get-all-employees',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillEmployeeData(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });
        
        $('.page-wrapper').hide();
        $('#employee__view').show();
    });

    $('#add__employee__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__employee').show();
    });

    $('#add__employee__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__employee').show();
    });

    $('#update__employee').click(function(){
        if ($('#employee__fullname').val() == "" ){
            alert('Fullname is empty')
            return;
        }
        if ($('#identity__number').val() == "" ){
            alert('Fullname is empty')
            return;
        }
        if ($('#position').val() == "" ){
            alert('Fullname is empty')
            return;
        }
        if ($('#employee__email').val() == "" ){
            alert('Fullname is empty')
            return;
        }
        if ($('#employee__phonenumber').val() == "" ){
            alert('Fullname is empty')
            return;
        }

        var employeeId = $('#employee__id').val();
        var fullname = $('#employee__fullname').val();
        var identityNumber = $('#identity__number').val();
        var position = $('#position').val();
        var email = $('#employee__email').val();
        var phoneNumber = $('#employee__phone__number').val();
        var data = {
            employeeId: parseInt(employeeId),
            fullname: fullname,
            identityNumber: identityNumber,
            position: position,
            email: email,
            phoneNumber: phoneNumber
        };

        $.ajax({
            url: 'http://localhost:8888/manager/update-employee',
            method: 'PUT',
            data: JSON.stringify(data),
            contentType: 'application/json',
            success: function() {
                alert('Update employee successfully')
                $.ajax({
                    url: 'http://localhost:8888/manager/get-all-employees',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillEmployeeData(response);
                    },
                    error: function(xhr, status, error) {
                        console.log(xhr);
                        console.log(status);
                        console.log(error);
                    }
                });
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                alert(xhr.responseJSON);
            }
        });
    });

    $('#create__employee').click(function(){
        if ($('#employee__name__create').val() == "" ){
            alert('Employee name is empty')
            return;
        }
        if ($('#identity__number__create').val() == "" ){
            alert('Identity number is empty')
            return;
        }
        if ($('#position__create').val() == "" ){
            alert('Position is empty')
            return;
        }
        if ($('#employee__email__create').val() == "" ){
            alert('Employee email is empty')
            return;
        }
        if ($('#employee__phone__number__create').val() == "" ){
            alert('Employee phone number is empty')
            return;
        }

        
        var fullname = $('#employee__name__create').val();
        var identityNumber = $('#identity__number__create').val();
        var position = $('#position__create').val();
        var email = $('#employee__email__create').val();
        var phoneNumber = $('#employee__phone__number__create').val();
        var data = {
            fullname: fullname,
            identityNumber: identityNumber,
            position: position,
            email: email,
            phoneNumber: phoneNumber
        };

        $.ajax({
            url: 'http://localhost:8888/manager/create',
            method: 'POST',
            data: JSON.stringify(data),
            contentType: 'application/json',
            success: function(response) {
                alert('Create employee successfully')
                $('#employee__name__create').val("");
                $('#identity__number__create').val("");
                $('#position__create').val("");
                $('#employee__email__create').val("");
                $('#employee__phone__number__create').val("");
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                alert(xhr.responseJSON);
            }
        });
    });
    
    $('#employee__delete__btn').click(function(){
        $.ajax({
            url: "http://localhost:8888/manager/delete-employee/" + employeeIDForDelete, // Replace with your API endpoint
            type: "DELETE",
            success: function() {
                alert("Employee deleted successfully");
                $('#employee__id').val('');
                $('#employee__fullname').val('');
                $('#identity__number').val('');
                $('#position').val('');
                $('#employee__email').val('');
                $('#employee__phonenumber').val('');
                $('.modal-content').hide();
                // Handle success, update UI, etc.
                $.ajax({
                    url: 'http://localhost:8888/manager/get-all-employees',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillEmployeeData(response);
                    },
                    error: function(xhr, status, error) {
                        alert(xhr.responseJSON);
                        console.log(status);
                        console.log(error);
                    }
                });
            },
            error: function(xhr, status, error) {
                console.error("Error deleting employee:", error);
                $('.modal-content').hide();
                // Handle error, show error message, etc.
            }
        });
    });

    $('#input__employee__id').keyup(function(){
        var employeeId = $(this).val();
        var employeeTable = $('#employee__table tbody tr');
        employeeTable.each(function(){
            var value = $(this).find("td:first-child").text();
            if (value.includes(employeeId)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

});
    
$(document).on('click', '.edit__employee__link', function(event) {
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var employeeData = $("#employee__data").find('.form-control');    
    // Find td elements within the closest tr
    var tds = closestTr.find("td");
    
    // Loop through td elements except the last one
    for (var i = 0; i < tds.length - 1; i++) {
        if (tds.eq(i).text() === 'Manager') {
            employeeData.eq(i).val('TP');
        } else if (tds.eq(i).text() === 'Employee') {
            employeeData.eq(i).val('NV');
        } else {
            employeeData.eq(i).val(tds.eq(i).text());
        }        
    }
});

$(document).on('click', '.delete__employee__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var tds = closestTr.find("td");
    employeeIDForDelete = tds.eq(0).text();
    if (window.confirm('Are you sure you want to delete this employee?')){
        var URL = "http://localhost:8888/manager/delete-employee/" + employeeIDForDelete;
            $.ajax({
                url: URL,
                method: 'DELETE',
                dataType: 'json',
                success: function() {
                    alert('Delete employee successfully');
                    $.ajax({
                        url: 'http://localhost:8888/manager/get-all-employees',
                        method: 'GET',
                        dataType: 'json',
                        success: function(response) {
                            fillEmployeeData(response);
                        },
                        error: function(xhr, status, error) {
                            alert(xhr.responseJSON);
                            console.log(status);
                            console.log(error);
                        }
                    });
                },
                error: function(xhr, status, error) {
                    alert(xhr.responseJSON);
                    console.log(status);
                    console.log(error);
                }
            });
    }
});