function fillBrandData(response) {
    $("#brand__table tbody tr").remove();
    $.each(response, function(_, brand) {        
        var newRow = $("<tr>");
        newRow.append(`<td>${brand.brandID}</td>`);
        newRow.append(`<td>${brand.brandName}</td>`);
        
        var actions = `<td class="text-right">
                            <div class="dropdown dropdown-action">
                                <a href="#" class="action-icon dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="fas fa-ellipsis-v ellipse_color"></i></a>
                                <div class="dropdown-menu dropdown-menu-right">
                                    <a class="dropdown-item edit__brand__link" href="#">
                                        <i class="fas fa-pencil-alt m-r-5"></i> Edit
                                    </a>
                                    <a class="dropdown-item delete__brand__link" href="#" data-toggle="modal" data-target="#delete_asset">
                                        <i class="fas fa-trash-alt m-r-5"></i> Delete
                                    </a>
                                </div>
                            </div>
                        </td>`;
        newRow.append(actions);
        
        $("#brand__table tbody").append(newRow);
    });
}

function fillProvinceData(response) {
    $("#province__table tbody tr").remove();
    $.each(response, function(_, province) {        
        var newRow = $("<tr>");
        newRow.append(`<td>${province.provinceID}</td>`);
        newRow.append(`<td>${province.provinceName}</td>`);
        
        var actions = `<td class="text-right">
                            <div class="dropdown dropdown-action">
                                <a href="#" class="action-icon dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="fas fa-ellipsis-v ellipse_color"></i></a>
                                <div class="dropdown-menu dropdown-menu-right">
                                    <a class="dropdown-item edit__province__link" href="#">
                                        <i class="fas fa-pencil-alt m-r-5"></i> Edit
                                    </a>
                                    <a class="dropdown-item delete__province__link" href="#" >
                                        <i class="fas fa-trash-alt m-r-5"></i> Delete
                                    </a>

                                </div>
                            </div>
                        </td>`;
        newRow.append(actions);
        
        $("#province__table tbody").append(newRow);
    });
}

function fillTypeOfWatchData(response) {
    $("#type__of__watch__table tbody tr").remove();
    $.each(response, function(_, typeOfWatch) {        
        var newRow = $("<tr>");
        newRow.append(`<td>${typeOfWatch.typeOfWatchID}</td>`);
        newRow.append(`<td>${typeOfWatch.typeOfWatchName}</td>`);
        
        var actions = `<td class="text-right">
                            <div class="dropdown dropdown-action">
                                <a href="#" class="action-icon dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="fas fa-ellipsis-v ellipse_color"></i></a>
                                <div class="dropdown-menu dropdown-menu-right">
                                    <a class="dropdown-item edit__type__of__watch__link" href="#">
                                        <i class="fas fa-pencil-alt m-r-5"></i> Edit
                                    </a>
                                    <a class="dropdown-item delete__type__of__watch__link" href="#" >
                                        <i class="fas fa-trash-alt m-r-5"></i> Delete
                                    </a>

                                </div>
                            </div>
                        </td>`;
        newRow.append(actions);
        
        $("#type__of__watch__table tbody").append(newRow);
    });
}

function fillOrderData(response) {
    $("#order__table tbody tr").remove();
    function formatOrderStatus(status) {
        switch (status) {
            case 0:
                return "Canceled";
            case 1:
                return "Finished";
            case 2:
                return "Being transported";
            case 3:
                return "Pending";
            case 4:
                return "Approved";
        }
    }
    $.each(response, function(_, order) {        
        var newRow = $("<tr>");
        newRow.append(`<td>${order.orderID}</td>`);
        newRow.append(`<td>${order.orderDateString}</td>`);
        newRow.append(`<td>${order.customerName}</td>`);
        newRow.append(`<td>${order.employeeName = order.employeeName == null ? "" : order.employeeName}</td>`);
        newRow.append(`<td>${order.invoiceID = order.invoiceID == null ? "" : order.invoiceID}</td>`);
        newRow.append(`<td>${formatOrderStatus(order.status)}</td>`);
        newRow.append(`<td>${order.provinceName}</td>`);
        newRow.append(`<td>${order.districtName}</td>`);
        newRow.append(`<td>${order.wardName}</td>`);
        newRow.append(`<td>${order.apartmentNumber}</td>`);
        
        var actions = `<td class="text-right">
                            <div class="dropdown dropdown-action">
                                <a href="#" class="action-icon dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="fas fa-ellipsis-v ellipse_color"></i></a>
                                <div class="dropdown-menu dropdown-menu-right">
                                    <a class="dropdown-item edit__order__link" href="#">
                                        <i class="fas fa-pencil-alt m-r-5"></i> Edit
                                    </a>
                                </div>
                            </div>
                        </td>`;
        newRow.append(actions);
        
        $("#order__table tbody").append(newRow);
    });
}

function fillWatchData(response) {
    $("#watch__table tbody tr").remove();
    function formatWatchStatus(status) {
        switch (status) {
            case 0:
                return "Out of stock";
            case 1:
                return "Available";
            case 2:
                return "No longer available for sale";
        }
    }
    $.each(response.watches, function(_, watch) {
           
        var newRow = $("<tr>");
        newRow.append(`<td>${watch.watchId}</td>`);
        newRow.append(`<td>${watch.watchName}</td>`);
        newRow.append(`<td>${watch.price}</td>`);

        newRow.append(`<td>
                            <h2 class="table-avatar">
                                <a href="#" class="avatar avatar-sm mr-2">
                                    <img class="avatar-img rounded-circle" src="${watch.image}" alt="User Image">
                                </a>
                            </h2>
                        </td>`);
        newRow.append(`<td>${formatWatchStatus(watch.status)}</td>`);
        newRow.append(`<td>${watch.quantity}</td>`);
        newRow.append(`<td>${watch.brand.brandName}</td>`);
        newRow.append(`<td>${watch.typeOfWatch.typeOfWatchName}</td>`);
        
        var actions = `<td class="text-right">
                            <div class="dropdown dropdown-action">
                                <a href="#" class="action-icon dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="fas fa-ellipsis-v ellipse_color"></i></a>
                                <div class="dropdown-menu dropdown-menu-right">
                                    <a class="dropdown-item edit__watch__link" href="#">
                                        <i class="fas fa-pencil-alt m-r-5"></i> Edit
                                    </a>
                                    <a class="dropdown-item delete__watch__link" href="#" >
                                        <i class="fas fa-trash-alt m-r-5"></i> Delete
                                    </a>
                                </div>
                            </div>
                        </td>`;
        newRow.append(actions);
        
        $("#watch__table tbody").append(newRow);
    });

    $("#brand__id__for__watch").empty();
    $("#brand__id__for__watch").append('<option value="">Select</option>');
    $.each(response.brands, function(_, brand) {
        $("#brand__id__for__watch").append(`<option value="${brand.brandID}">
        ${brand.brandName}</option>`);
    });

    $("#watch__type__id").empty();
    $("#watch__type__id").append('<option value="">Select</option>');
    $.each(response.typesOfWatch, function(_, typeOfWatch) {
    $("#watch__type__id").append(`<option value="${typeOfWatch.typeOfWatchID}">
        ${typeOfWatch.typeOfWatchName}</option>`);
    });


}

function fillOrderDetailData(response) {
    $('#order__detail__data .col-md-4').remove();
    var orderDetails =  $('#order__detail__data .formtype')
    // orderDetails.empty();        
    $.each(response, function(_, orderDetail){
        console.log(orderDetail.WatchName);
        var newWatch = `<div class="col-md-4">
                        <div class="form-group">
                            <label>Watch Name</label>
                            <input class="form-control" value="${orderDetail.WatchName}" type="text" disabled>
                        </div>
                        </div>

                        <div class="col-md-4">
                            <div class="form-group">
                                <label>Quantity</label>
                                <input class="form-control" value="${orderDetail.Quantity}" type="text" disabled>
                            </div>
                        </div>

                        <div class="col-md-4">
                            <div class="form-group">
                                <label>Unit Price</label>
                                <input class="form-control" value="${orderDetail.UnitPrice}" type="text" disabled>
                            </div>
                        </div>`
        orderDetails.append(newWatch)
    });    
}

function fillWatchDataWhenAdding(response){
    var brands = $('#watch__brand__create');
    var typesOfWatch = $('#watch__type__create');

    brands.append('<option value="">Select</option>');
    $.each(response.brands, function(_, brand) {
        brands.append(`<option value="${brand.brandID}">
        ${brand.brandName}</option>`);
    });

    typesOfWatch.append('<option value="">Select</option>');
    $.each(response.typesOfWatch, function(_, typeOfWatch) {
        typesOfWatch.append(`<option value="${typeOfWatch.typeOfWatchID}">
        ${typeOfWatch.typeOfWatchName}</option>`);
    });
}

var brandIDForDelete = 0;
var provinceIDForDelete = 0;
var typeOfWatchIDForDelete = 0;

$(document).ready(function() {
    var employeeString = localStorage.getItem('employee');
    if (employeeString === "") {
        alert("Please login to access this page!");
        window.location.href = "http://localhost:8888/employee/login-form";
    }

    var employee = JSON.parse(employeeString);

    $('.user-text h6').text(employee.fullname);

    $('#logout').click(function(){
        employee = "";
        localStorage.setItem('employee',employee);
    });

    //BRAND
    $('#all__brand__link').click(function(event) {
        event.preventDefault();

        $.ajax({
            url: 'http://localhost:8888/employee/get-all-brands',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillBrandData(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });

        $('.page-wrapper').hide();

        $('#brand__view').show();
    });

    $('#add__brand__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__brand').show();
    });

    $('#add__brand__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__brand').show();
    });

    $('#input__brand__name').keyup(function(){
        var brandName = $(this).val().toLowerCase();
        var brandTable = $('#brand__table tbody tr');
        brandTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(brandName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    $('#add__brand__btn').click(function(event){
        event.preventDefault();

        var brandName = $("#brand__name__create").val();

        // Create an object to hold the data
        var postData = {
            brandName: brandName
        };

        $.ajax({
            url: 'http://localhost:8888/employee/create-brand',
            method: 'POST',
            data: JSON.stringify(postData),
            contentType: "application/json",
            success: function() {
                alert('Create new brand successfully!');
                $("#brand__name__create").val('');
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });
        
    });

    $('#update__brand__btn').click(function(){
        var brandID = parseInt($('#brand__id').val());
        var brandName = $("#brand__name").val();

        // Create an object to hold the data
        var postData = {
            brandID: brandID,
            brandName: brandName
        };

        $.ajax({
            url: 'http://localhost:8888/employee/update-brand',
            method: 'PUT',
            data: JSON.stringify(postData),
            contentType: "application/json",
            success: function() {
                alert('Update brand successfully!');
                $('#brand__id').val('');
                $("#brand__name").val('');
                $.ajax({
                    url: 'http://localhost:8888/employee/get-all-brands',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillBrandData(response);
                    },
                    error: function(xhr, status, error) {
                        console.log(xhr);
                        console.log(status);
                        console.log(error);
                    }
                });
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });
    });

    $('#delete__brand__btn').click(function(){
        var URL = 'http://localhost:8888/employee/delete-brand/' + brandIDForDelete;
        $.ajax({
            url: URL,
            method: 'DELETE',
            dataType: 'json',
            success: function() {
                alert('Delete brand successfully');
                $.ajax({
                    url: 'http://localhost:8888/employee/get-all-brands',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillBrandData(response);
                    },
                    error: function(xhr, status, error) {
                        console.log(xhr);
                        console.log(status);
                        console.log(error);
                    }
                });
                $('.modal-content').hide();
            },
            error: function(xhr, status, error) {
                alert(xhr.responseJSON);
                $('.modal-content').hide();
                console.log(status);
                console.log(error);
            }
        });
    });

    //DISTRICT
    $('#all__district__link').click(function(event) {
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#district__view').show();
    });

    $('#add__district__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__district').show();
    });

    $('#add__district__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__district').show();
    });

    $('.edit__district__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var districtData = $("#district__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            districtData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#input__district__name').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#district__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    //INVOICES
    $('#all__invoice__link').click(function(event) {
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#invoice__view').show();
    });

    $('.edit__invoice__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var invoiceData = $("#invoice__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            invoiceData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#input__invoice__id').keyup(function(){
        var invoiceName = $(this).val().toLowerCase();
        var invoiceTable = $('#invoice__table tbody tr');
        invoiceTable.each(function(){
            var value = $(this).find("td:nth-child(1)").text().toLowerCase();
            if (value.includes(invoiceName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    //ORDER DETAIL
    $('#all__order__detail__link').click(function(event) {
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#order__detail__view').show();
    });

    $('.edit__order__detail__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var districtData = $("#order__detail__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            districtData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#input__order__detail__id').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#order__detail__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    //ORDER
    $('#all__order__link').click(function(event) {
        event.preventDefault();
        $.ajax({
            url: 'http://localhost:8888/employee/get-all-orders',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillOrderData(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });

        var orderDetails = $("#order__data").find('.form-control'); 
        $.each(orderDetails, function(_, orderDetail){
            $(orderDetail).val('');
        });

        $('#order__detail__data .col-md-4').remove();
        $('.page-wrapper').hide();
        $('#order__view').show();
    });
    
    $('#update__order__btn').click(function(){
        var orderID = $('#order__id').val();
        var status = parseInt($('#order__status').val());
        var employeeID = employee.employeeId;
        var data = {
            orderID: parseInt(orderID),
            status: status,
            employeeID: parseInt(employeeID)
        }
        // console.log(employee);
        $.ajax({
            url: 'http://localhost:8888/employee/update-order-status',
            method: 'PUT',
            data: JSON.stringify(data),
            contentType: "application/json",
            success: function() {
                alert('Update order status successfully!');
                $.ajax({
                    url: 'http://localhost:8888/employee/get-all-orders',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillOrderData(response);
                    },
                    error: function(xhr, status, error) {
                        console.log(xhr);
                        console.log(status);
                        console.log(error);
                    }
                });
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });

    });

    $('#input__order__id').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#order__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(1)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    //PROMOTION DETAIL
    $('#all__promotion__detail__link').click(function(event) {
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#promotion__detail__view').show();
    });

    $('.edit__promotion__detail__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var districtData = $("#promotion__detail__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            districtData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#input__promotion__detail__id').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#promotion__detail__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(1)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    //PROMOTION 
    $('#all__promotion__link').click(function(event) {
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#promotion__view').show();
    });

    $('.edit__promotion__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var districtData = $("#promotion__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            districtData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#input__promotion__id').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#promotion__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(1)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    $('#add__promotion__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__promotion').show();
    });

    $('#add__promotion__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__promotion').show();
    });
    
    $("#add__watch__name").click(function() {
        var newWatchNameField = `
            <div class="col-md-4">
                <div class="form-group">
                    <label>Watch name</label>
                    <input class="form-control" type="text">
                </div>
            </div>
        `;
        $("#watch__name__list .formtype").append(newWatchNameField);
    });

    $("#remove__watch__name").click(function() {
        var watchNameFields = $("#watch__name__list .formtype .col-md-4");
        if (watchNameFields.length > 4) {
            watchNameFields.last().remove();
            // updateRemoveButtonState();
        }
    });

    //PROVINCE
    $('#all__province__link').click(function(event) {
        event.preventDefault();
        $.ajax({
            url: 'http://localhost:8888/employee/get-all-provinces',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillProvinceData(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });
        $('.page-wrapper').hide();
        $('#province__view').show();
    });

    $('#add__province__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__province').show();
    });

    $('#add__province__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__province').show();
    });

    $('#input__province__name').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#province__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    $('#add__province__btn').click(function(event){
        event.preventDefault();

        var provinceName = $("#province__name__create").val();

        // Create an object to hold the data
        var postData = {
            provinceName: provinceName
        };

        $.ajax({
            url: 'http://localhost:8888/employee/create-province',
            method: 'POST',
            data: JSON.stringify(postData),
            contentType: "application/json",
            success: function() {
                alert('Create new province successfully!');
                $("#province__name__create").val('');
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });
        
    });
    
    $('#update__province__btn').click(function(){
        var provinceID = parseInt($('#province__id').val());
        var provinceName = $("#province__name").val();
        console.log(provinceName);
        // Create an object to hold the data
        var postData = {
            provinceID: provinceID,
            provinceName: provinceName
        };

        $.ajax({
            url: 'http://localhost:8888/employee/update-province',
            method: 'PUT',
            data: JSON.stringify(postData),
            contentType: "application/json",
            success: function() {
                alert('Update province successfully!');
                $('#province__id').val('');
                $("#province__name").val('');
                $.ajax({
                    url: 'http://localhost:8888/employee/get-all-provinces',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillProvinceData(response);
                    },
                    error: function(xhr, status, error) {
                        console.log(xhr);
                        console.log(status);
                        console.log(error);
                    }
                });
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });
    });

    

    //TYPE OF WATCH
    $('#all__type__of__watch__link').click(function(event) {
        event.preventDefault();
        $.ajax({
            url: 'http://localhost:8888/employee/get-all-types-of-watch',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillTypeOfWatchData(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });
        $('.page-wrapper').hide();
        $('#type__of__watch__view').show();
    });

    $('.edit__type__of__watch__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var districtData = $("#type__of__watch__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            districtData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#add__type__of__watch__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__type__of__watch').show();
    });

    $('#add__type__of__watch__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__type__of__watch').show();
    });

    $('#input__type__of__watch__name').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#type__of__watch__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    $('#add__type__of__watch__btn').click(function(event){
        event.preventDefault();

        var typeOfWatchName = $("#type__of__watch__name__create").val();

        // Create an object to hold the data
        var postData = {
            typeOfWatchName: typeOfWatchName
        };

        $.ajax({
            url: 'http://localhost:8888/employee/create-type-of-watch',
            method: 'POST',
            data: JSON.stringify(postData),
            contentType: "application/json",
            success: function() {
                alert('Create new type of watch successfully!');
                $("#type__of__watch__name__create").val('');
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });
        
    });
    
    $('#update__type__of__watch__btn').click(function(){
        var typeOfWatchID = parseInt($('#type__of__watch__id').val());
        var typeOfWatchName = $("#type__of__watch__name").val();
        // console.log(provinceName);
        // Create an object to hold the data
        var postData = {
            typeOfWatchID: typeOfWatchID,
            typeOfWatchName: typeOfWatchName
        };

        $.ajax({
            url: 'http://localhost:8888/employee/update-type-of-watch',
            method: 'PUT',
            data: JSON.stringify(postData),
            contentType: "application/json",
            success: function() {
                alert('Update this type of watch successfully!');
                $('#type__of__watch__id').val('');
                $("#type__of__watch__name").val('');
                $.ajax({
                    url: 'http://localhost:8888/employee/get-all-types-of-watch',
                    method: 'GET',
                    dataType: 'json',
                    success: function(response) {
                        fillTypeOfWatchData(response);
                    },
                    error: function(xhr, status, error) {
                        console.log(xhr);
                        console.log(status);
                        console.log(error);
                    }
                });
            },
            error: function(xhr, _, _) {
                alert(xhr.responseJSON);
                // console.log(status);
                // alert(error.responseJSON);
                // console.log(error)
            }
        });
    });

    //WARDS
    $('#all__ward__link').click(function(event) {
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#ward__view').show();
    });

    $('.edit__ward__link').click(function(event){
        event.preventDefault();
        // Find the closest parent tr element
        var closestTr = $(this).closest("tr");
        var districtData = $("#ward__data").find('.form-control');    
        // Find td elements within the closest tr
        var tds = closestTr.find("td");
        
        // Loop through td elements except the last one
        for (var i = 0; i < tds.length - 1; i++) {
            districtData.eq(i).val(tds.eq(i).text());
        }
    });

    $('#input__ward_name').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#ward__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    $('#add__ward__link').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__ward').show();
    });

    $('#add__ward__link__view').click(function(event){
        event.preventDefault();
        $('.page-wrapper').hide();
        $('#add__ward').show();
    });

    //WATCH
    $('#all__watch__link').click(function(event) {
        event.preventDefault();
        $.ajax({
            url: 'http://localhost:8888/employee/get-all-watches',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillWatchData(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });
        $('.page-wrapper').hide();
        $('#watch__view').show();
    });

    $('#input__watch_name').keyup(function(){
        var districtName = $(this).val().toLowerCase();
        var districtTable = $('#watch__table tbody tr');
        districtTable.each(function(){
            var value = $(this).find("td:nth-child(2)").text().toLowerCase();
            if (value.includes(districtName)) {
                $(this).show();
            } else {
                $(this).hide();
            }
        });
    });

    $('#add__watch__link').click(function(event){
        event.preventDefault();
        $.ajax({
            url: 'http://localhost:8888/employee/get-brans-and-types-of-watch',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillWatchDataWhenAdding(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });
        $('.page-wrapper').hide();
        $('#add__watch').show();
    });

    $('#add__watch__link__view').click(function(event){
        event.preventDefault();
        $.ajax({
            url: 'http://localhost:8888/employee/get-brans-and-types-of-watch',
            method: 'GET',
            dataType: 'json',
            success: function(response) {
                fillWatchDataWhenAdding(response);
            },
            error: function(xhr, status, error) {
                console.log(xhr);
                console.log(status);
                console.log(error);
            }
        });
        $('.page-wrapper').hide();
        $('#add__watch').show();
    });

    $('#watch__create__btn').click(function(){
        var fileInput = $("#watch__img__create")[0];
        var base64Image;
        if (fileInput.files && fileInput.files[0]){
            var reader = new FileReader();
            reader.onload = function (e) {
                base64Image = e.target.result.split(',')[1]; // Extract base64 data
                
                var formData = {
                    watchName: $("#watch__name__create").val(),
                    price: parseInt($("#watch__price__create").val()),
                    image: base64Image,
                    quantity: parseInt($('#watch__quantity__create').val()),
                    brandId: parseInt($('#watch__brand__create').val()),
                    typeOfWatchId: parseInt($('#watch__type__create').val()),
                };
                
                $.ajax({
                    url: "http://localhost:8888/employee/create-watch",  // Update with your server endpoint
                    type: "POST",
                    contentType: "application/json",
                    data: JSON.stringify(formData),
                    success: function () {
                        alert("Create watch successfully!");
                        $("#watch__name__create").val('');
                        $("#watch__price__create").val('');
                        $('#imagePreview').attr('src', "");
                        $('#watch__quantity__create').val('');
                        $('#watch__brand__create').val('');
                        $('#watch__type__create').val('');
                    },
                    error: function (error) {
                        alert("Have error!");
                        console.error("Error submitting data:", error);
                    }
                });
            };
        }
        reader.readAsDataURL(fileInput.files[0]);
        
    });

    $('#update__watch__btn').click(function(){
        var fileInput = $("#watch__img")[0];
        if (fileInput.files && fileInput.files[0]){
            var reader = new FileReader();
            reader.onload = function (e) {
                var imageData = e.target.result.split(',')[1];
                
                var formData = {
                    watchId: parseInt($('#watch__id').val()),
                    watchName: $("#watch__name").val(),
                    price: parseInt($("#watch__price").val()),
                    image: imageData,
                    status: parseInt($('#watch__status').val()),
                    quantity: parseInt($('#watch__quantity').val()),
                    brandId: parseInt($('#brand__id__for__watch').val()),
                    typeOfWatchId: parseInt($('#watch__type__id').val()),
                };
                
                $.ajax({
                    url: "http://localhost:8888/employee/update-watch",  // Update with your server endpoint
                    type: "PUT",
                    contentType: "application/json",
                    data: JSON.stringify(formData),
                    success: function () {
                        alert("Update watch successfully!");
                    },
                    error: function (error) {
                        alert("Have error!");
                        console.error("Error submitting data:", error);
                    }
                });
            };
            reader.readAsDataURL(fileInput.files[0]);
        } else {
            console.log(parseInt($('#watch__id').val()));
            var formData = {
                watchId: parseInt($('#watch__id').val()),
                watchName: $("#watch__name").val(),
                price: parseInt($("#watch__price").val()),
                status: parseInt($('#watch__status').val()),
                quantity: parseInt($('#watch__quantity').val()),
                brandId: parseInt($('#brand__id__for__watch').val()),
                typeOfWatchId: parseInt($('#watch__type__id').val()),
            };
            
            $.ajax({
                url: "http://localhost:8888/employee/update-watch",  // Update with your server endpoint
                type: "PUT",
                contentType: "application/json",
                data: JSON.stringify(formData),
                success: function () {
                    alert("Update watch successfully!");
                },
                error: function (error) {
                    alert("Have error!");
                    console.error("Error submitting data:", error);
                }
            });
        }        
        
    });

});

//BRAND
$(document).on('click', '.edit__brand__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var brandData = $("#brand__data").find('.form-control');    
    // Find td elements within the closest tr
    var tds = closestTr.find("td");
    
    // Loop through td elements except the last one
    for (var i = 0; i < tds.length - 1; i++) {
        brandData.eq(i).val(tds.eq(i).text());
    }
});

$(document).on('click', '.delete__brand__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var tds = closestTr.find("td");
    brandIDForDelete = parseInt(tds.eq(0).text());
    $('.modal-content').show();
});

//PROVINCE

$(document).on('click', '.edit__province__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var brandData = $("#province__data").find('.form-control');    
    // Find td elements within the closest tr
    var tds = closestTr.find("td");
    
    // Loop through td elements except the last one
    for (var i = 0; i < tds.length - 1; i++) {
        brandData.eq(i).val(tds.eq(i).text());
    }
});

$(document).on('click', '.delete__province__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var tds = closestTr.find("td");
    provinceIDForDelete = parseInt(tds.eq(0).text());
    if (window.confirm('Are you sure you want to delete this province?')){
        var URL = 'http://localhost:8888/employee/delete-province/' + provinceIDForDelete;
            $.ajax({
                url: URL,
                method: 'DELETE',
                dataType: 'json',
                success: function() {
                    alert('Delete province successfully');
                    $.ajax({
                        url: 'http://localhost:8888/employee/get-all-provinces',
                        method: 'GET',
                        dataType: 'json',
                        success: function(response) {
                            fillProvinceData(response);
                        },
                        error: function(xhr, status, error) {
                            console.log(xhr);
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

//TYPE OF WATCH

$(document).on('click', '.edit__type__of__watch__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var brandData = $("#type__of__watch__data").find('.form-control');    
    // Find td elements within the closest tr
    var tds = closestTr.find("td");
    
    // Loop through td elements except the last one
    for (var i = 0; i < tds.length - 1; i++) {
        brandData.eq(i).val(tds.eq(i).text());
    }
});

$(document).on('click', '.delete__type__of__watch__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var tds = closestTr.find("td");
    typeOfWatchIDForDelete = tds.eq(0).text();
    if (window.confirm('Are you sure you want to delete this type of watch?')){
        var URL = 'http://localhost:8888/employee/delete-type-of-watch/' + typeOfWatchIDForDelete;
            $.ajax({
                url: URL,
                method: 'DELETE',
                dataType: 'json',
                success: function() {
                    alert('Delete type of watch successfully');
                    $.ajax({
                        url: 'http://localhost:8888/employee/get-all-types-of-watch',
                        method: 'GET',
                        dataType: 'json',
                        success: function(response) {
                            fillTypeOfWatchData(response);
                        },
                        error: function(xhr, status, error) {
                            console.log(xhr);
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

//ORDER
$(document).on('click', '.edit__order__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var districtData = $("#order__data").find('.form-control');    
    // Find td elements within the closest tr
    var tds = closestTr.find("td");
    
    // Loop through td elements except the last one
    for (var i = 0; i < tds.length - 1; i++) {
        if (tds.eq(i).text() == 'Pending'){
            districtData.eq(i).val('3');
        } else if (tds.eq(i).text() == 'Approved'){
            districtData.eq(i).val('4');
        } else if (tds.eq(i).text() == 'Being transported'){
            districtData.eq(i).val('2');
        } else if (tds.eq(i).text() == 'Finished'){
            districtData.eq(i).val('1');
        } else if (tds.eq(i).text() == 'Canceled'){
            districtData.eq(i).val('0');
        } else {
            districtData.eq(i).val(tds.eq(i).text());
        }
    }

    $.ajax({
        url: 'http://localhost:8888/employee/get-order-details-by-order-id/' + tds.eq(0).text(),
        method: 'GET',
        dataType: 'json',
        success: function(response) {
            fillOrderDetailData(response);
        },
        error: function(xhr, status, error) {
            console.log(xhr);
            console.log(status);
            console.log(error);
        }
    });

    
});

//WATCH
$(document).on('click', '.edit__watch__link', function(event){
    event.preventDefault();
    // Find the closest parent tr element
    var closestTr = $(this).closest("tr");
    var districtData = $("#watch__data").find('.form-control');    
    // Find td elements within the closest tr
    var tds = closestTr.find("td");

    for (var i = 0; i < tds.length - 1; i++) {

        if (i === 3){
            var imgSrc = tds.find('img.avatar-img').attr('src');
            $('#image__view').attr('src', imgSrc);
            continue;
        }

        var foundValue = false;
        if (i === 4 || i === 6 || i === 7) {
            var targetSelectId = '';
            if (i === 4) {
                targetSelectId = '#watch__status';
            } else if (i === 6) {
                targetSelectId = '#brand__id__for__watch';
            } else if (i === 7) {
                targetSelectId = '#watch__type__id';
            }
            foundValue = true;
            var selectElement = $(targetSelectId);
            var tdText = tds.eq(i).text().trim();
            selectElement.find('option').each(function() {
                var optionText = $(this).text().trim();
                if (optionText === tdText) {
                    districtData.eq(i).val($(this).val());
                    return;
                }
            });

        }        

        if (foundValue){
            continue;
        }

        districtData.eq(i).val(tds.eq(i).text());
    }
       
});

$(document).on('click', '.delete__watch__link', function(event){
    event.preventDefault();
    var closestTr = $(this).closest("tr");
    var tds = closestTr.find("td");
    var watchID = tds.eq(0).text();
    if (window.confirm('Are you sure you want to delete this watch?')){
        var URL = 'http://localhost:8888/employee/delete-watch/' + watchID;
            $.ajax({
                url: URL,
                method: 'DELETE',
                dataType: 'json',
                success: function() {
                    alert('Delete watch successfully');
                    console.log("Before removing closestTr");
                    closestTr.remove();
                    console.log("After removing closestTr");
                },
                error: function(xhr, status, error) {
                    alert(xhr.responseJSON);
                    console.log(status);
                    console.log(error);
                }
            });
    }
});