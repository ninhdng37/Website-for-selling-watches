/*=============== SHOW MENU ===============*/
const navMenu = document.getElementById('nav-menu'),
      navToggle = document.getElementById('nav-toggle'),
      navClose = document.getElementById('nav-close')

/*===== MENU SHOW =====*/
/* Validate if constant exists */
if(navToggle){
    navToggle.addEventListener('click', () =>{
        navMenu.classList.add('show-menu')
    })
}

/*===== MENU HIDDEN =====*/
/* Validate if constant exists */
if(navClose){
    navClose.addEventListener('click', () =>{
        navMenu.classList.remove('show-menu')
    })
}

/*=============== REMOVE MENU MOBILE ===============*/
const navLink = document.querySelectorAll('.nav__link')

function linkAction(){
    const navMenu = document.getElementById('nav-menu')
    // When we click on each nav__link, we remove the show-menu class
    navMenu.classList.remove('show-menu')
}
navLink.forEach(n => n.addEventListener('click', linkAction))

/*=============== CHANGE BACKGROUND HEADER ===============*/
function scrollHeader(){
    const header = document.getElementById('header')
    // When the scroll is greater than 50 viewport height, add the scroll-header class to the header tag
    if(this.scrollY >= 50) header.classList.add('scroll-header'); else header.classList.remove('scroll-header')
}
window.addEventListener('scroll', scrollHeader)

/*=============== TESTIMONIAL SWIPER ===============*/
let testimonialSwiper = new Swiper(".testimonial-swiper", {
    spaceBetween: 30,
    loop: 'true',
    navigation: {
      nextEl: ".swiper-button-next",
      prevEl: ".swiper-button-prev",
    },
  });

/*=============== NEW SWIPER ===============*/
let newSwiper = new Swiper(".new-swiper", {
    spaceBetween: 24,
    loop: 'true',
    
    breakpoints: {
        576: {
          slidesPerView: 2,
        },
        768: {
          slidesPerView: 3,
        },
        1024: {
          slidesPerView: 4,
        },
      },
  });

/*=============== SCROLL SECTIONS ACTIVE LINK ===============*/
const sections = document.querySelectorAll('section[id]')

// function scrollActive(){
//     const scrollY = window.pageYOffset

//     sections.forEach(current =>{
//         const sectionHeight = current.offsetHeight,
//               sectionTop = current.offsetTop - 58,
//               sectionId = current.getAttribute('id')

//         if(scrollY > sectionTop && scrollY <= sectionTop + sectionHeight){
//             document.querySelector('.nav__menu a[href*=' + sectionId + ']').classList.add('active-link')
//         }else{
//             document.querySelector('.nav__menu a[href*=' + sectionId + ']').classList.remove('active-link')
//         }
//     })
// }
// window.addEventListener('scroll', scrollActive)

/*=============== SHOW SCROLL UP ===============*/ 
function scrollUp(){
  const scrollUp = document.getElementById('scroll-up');
  // When the scroll is higher than 400 viewport height, add the show-scroll class to the a tag with the scroll-top class
  if(this.scrollY >= 400) scrollUp.classList.add('show-scroll'); else scrollUp.classList.remove('show-scroll')
}
window.addEventListener('scroll', scrollUp)

/*=============== SHOW CART ===============*/
const cart = document.getElementById('cart'),
      cartShop = document.getElementById('cart-shop'),
      cartClose = document.getElementById('cart-close')

/*===== MENU SHOW =====*/
/* Validate if constant exists */
if(cartShop){
    cartShop.addEventListener('click', () =>{
        cart.classList.add('show-cart')
    })
}

/*===== MENU HIDDEN =====*/
/* Validate if constant exists */
if(cartClose){
    cartClose.addEventListener('click', () =>{
        cart.classList.remove('show-cart')
    })
}

/*=============== DARK LIGHT THEME ===============*/ 
const themeButton = document.getElementById('theme-button')
const darkTheme = 'dark-theme'
const iconTheme = 'bx-sun'

// Previously selected topic (if user selected)
const selectedTheme = localStorage.getItem('selected-theme')
const selectedIcon = localStorage.getItem('selected-icon')

// We obtain the current theme that the interface has by validating the dark-theme class
const getCurrentTheme = () => document.body.classList.contains(darkTheme) ? 'dark' : 'light'
const getCurrentIcon = () => themeButton.classList.contains(iconTheme) ? 'bx bx-moon' : 'bx bx-sun'

// We validate if the user previously chose a topic
if (selectedTheme) {
  // If the validation is fulfilled, we ask what the issue was to know if we activated or deactivated the dark
  document.body.classList[selectedTheme === 'dark' ? 'add' : 'remove'](darkTheme)
  themeButton.classList[selectedIcon === 'bx bx-moon' ? 'add' : 'remove'](iconTheme)
}

// Activate / deactivate the theme manually with the button
themeButton.addEventListener('click', () => {
    // Add or remove the dark / icon theme
    document.body.classList.toggle(darkTheme)
    themeButton.classList.toggle(iconTheme)
    // We save the theme and the current icon that the user chose
    localStorage.setItem('selected-theme', getCurrentTheme())
    localStorage.setItem('selected-icon', getCurrentIcon())
})
function generateProductHTML(product) {
    var quantityText = product.Quantity > 0 ? `Quantity: ${product.Quantity}` : 'Out of stock';
    var articleHtml = `
        <article class="products__card">
            <img src="${product.Image}" alt="" class="products__img">
            <h3 class="products__title">${product.WatchName}</h3>
            <span class="products__price">$${product.Price}</span> <br>
            <span class="products__price products__quantity">${quantityText}</span>
            <button class="products__button">
                <i class='bx bx-shopping-bag'></i>
            </button>
        </article>
    `;
    return articleHtml;
}

// Define a function to update the cart item quantity
function addCartItemQuantity(cartItem, quantity) {
    var cartAmountNumber = cartItem.find('.cart__amount-number');
    var currentAmount = parseInt(cartAmountNumber.text(), 10);
    if (currentAmount === 10) {
        alert("The quantity that you can buy for each item is only 10")
        return;
    }
    cartAmountNumber.text(currentAmount + quantity);
}

function addToContainer(container, article) {
    var articleHtml = `<article class="cart__card">
                    <div class="cart__box">
                        <img src="${article.find('img').attr('src')}" alt="" class="cart__img">
                    </div>
    
                    <div class="cart__details">
                        <h3 class="cart__title">${article.find('.products__title').text()}</h3>
                        <span class="cart__price">${article.find('.products__price').eq(0).text()}</span>
    
                        <div class="cart__amount">
                            <div class="cart__amount-content">
                                <span class="cart__amount-box">
                                    <i class='bx bx-minus' ></i>
                                </span>
    
                                <span class="cart__amount-number">1</span>
    
                                <span class="cart__amount-box">
                                    <i class='bx bx-plus' ></i>
                                </span>
                            </div>
    
                            <i class='bx bx-trash-alt cart__amount-trash' ></i>
                        </div>
                    </div>
                </article>`
                container.append(articleHtml)
}

function checkItemExistInCart(item, article){
    for (var i = 0; i < item.length; i++) {
        if ($(this).find('.cart__title').text() === article.find('.products__title').text()) {
            addCartItemQuantity($(this), 1);
            // itemFound = true;
            return false; // Exit the loop since item is found
        }
    }
    return true;
}

function getPrice(element, productName) {
    var price;
    element.each(function() {
        // Find the products__price element within the current product card
        if ($(this).find('.products__title').text() === productName) {
            var priceElement = $(this).find('.products__price');
            price = priceElement.text().slice(1);
            price = parseInt(price);
            return;
        }
    });
    return price;
}

function updateTotalPrice(price, action){
    var totalPrice = $('.cart__prices .cart__prices-total').text().slice(1);
    totalPrice = parseInt(totalPrice, 10);
    if (action === '+') {
        totalPrice += price;
    } else totalPrice -= price;
    
    $('.cart__prices .cart__prices-total').text('$' + totalPrice);
}

function updatePriceForItem(element, price, action) {
    itemPrice = element.text().slice(1);
    itemPrice = parseInt(itemPrice);
    if (action === '+')
        itemPrice += price;
    else itemPrice -= price;
    element.text('$' + itemPrice);
}

function updateItemNumberTotal(element, action){
    var total = element.text()[0];
    total = parseInt(total);
    if (action === '+') total++;
    else total--;
    element.text(total + ' items')
}

document.addEventListener("DOMContentLoaded", function() {
    var openModalBtn = document.getElementById("buying-btn");
    var modal = document.getElementById("myModal");
    var closeModal = document.getElementsByClassName("close")[0];
    var modalSubmitBtn = document.getElementById("modalSubmitBtn");

    openModalBtn.onclick = function() {
        modal.style.display = "block";
    };

    closeModal.onclick = function() {
        modal.style.display = "none";
    };
    
    modalSubmitBtn.onclick = function() {
        // console.log($('cart__card'));
        var cartData = [];
        $('.cart__card').each(function(){
            var $cartCard = $(this);
            var title = $cartCard.find('.cart__title').text();
            var unitPrice = parseInt(getPrice($('.products__card'), 
            $cartCard.find('.cart__title').text()));
            var quantity = parseInt($cartCard.find('.cart__amount-number').text(), 10);
            var cartItem = {
                "title": title,
                "unitPrice": unitPrice,
                "quantity": quantity
            };
        
            cartData.push(cartItem);
        });

        var postData = {
            provinceID: parseInt($('#province__combo').val()),
            districtId: parseInt($('#district__combo').val()),
            wardId: parseInt($('#ward__combo').val()),
            apartmentNumber: $('#apartment__number').val(),
            items: cartData
          };
        $.ajax({
            url:"http://localhost:8888/customer/add-order",
            method: "POST",
            data: JSON.stringify(postData), // Convert data to JSON format
            contentType: "application/json", // Set content type to JSON
            success: function() {
                $('.cart__card').each(function(){
                    $(this).remove();
                });
                $('#myModal').hide();
                $('#cart').removeClass('show-cart');
                cart.classList.remove('show-cart');
                $('.cart__item-count').text(0);
                $('.cart__prices-item').text('0 item');
                $('.cart__prices-total').text('$0');
                alert('You bought successfully!');
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
    };

    window.onclick = function(event) {
        if (event.target === modal) {
            modal.style.display = "none";
        }
    };
});

function updateItemCount(action){
    itemCount = $('.nav__shop .cart__item-count').text();
    itemCount = parseInt(itemCount);
    if (action === '+') itemCount++;
    else itemCount--;
    $('.nav__shop .cart__item-count').text(itemCount);
}

$(document).ready(function(){
    var provinceCombo = $('#province__combo');
    var districtCombo = $('#district__combo');
    var wardCombo = $('#ward__combo');
    $.ajax({
        url:"http://localhost:8888/customer/get-all",
        method: "GET",
        success: function(response) {
            // console.log(response.isCustomer);
            if (response.isCustomer) {
                $('.nav__shop .bx.bx-log-out').show();
                $('.nav__shop .bx.bx-log-in').hide();
            } else {
                $('.nav__shop .bx.bx-log-out').hide();
                $('.nav__shop .bx.bx-log-in').show();
            }

            $('.nav__shop .bx.bx-log-in').on('click', function(){
                window.location.href = 'http://localhost:8888/customer/login-form';
            });
            // provinceCombo.empty();
            var optionElement = $('<option>')
                    .attr('value','')
                    .text('Province');

                provinceCombo.append(optionElement);

            districtCombo.empty();
            districtCombo.append($("<option>")
            .val('')
            .text('District'));

            wardCombo.empty();
            wardCombo.append($("<option>")
            .val('')
            .text('Ward'));

            $.each(response.provinces, function(_, province){
                var optionElement = $('<option>')
                    .attr('value', province.provinceID)
                    .text(province.provinceName);
                provinceCombo.append(optionElement);
            });
            
            provinceCombo.on('change', function() {
                var selectedProvince = $(this).val();
                selectedProvince = parseInt(selectedProvince);

                districtCombo.empty();
                districtCombo.append($("<option>")
                .val('')
                .text('District'));

                wardCombo.empty();
                wardCombo.append($("<option>")
                .val('')
                .text('Ward'));

                $.each(response.districts, function(_, district){
                    var val = parseInt(district.provinceId);
                    // console.log(val);
                    // console.log(selectedProvince);
                    if (val === selectedProvince){
                        districtCombo.append($("<option>")
                        .val(district.districtId)
                        .text(district.districtName));
                    }                    
                });
            });

            districtCombo.on('change', function() {
                var selectedDistrict = $(this).val();
                selectedDistrict = parseInt(selectedDistrict);

                wardCombo.empty();
                wardCombo.append($("<option>")
                .val('')
                .text('Ward'));

                $.each(response.wards, function(_, ward){
                    var val = parseInt(ward.districtId)
                    if (val === selectedDistrict){
                        wardCombo.append($("<option>")
                        .val(ward.wardId)
                        .text(ward.wardName));
                    }                    
                });
            });

            // Select the container element using its class
            var container = $('.products__container.grid');
            
            $.each(response.watches, function(_, watch){
                var articleHtml = generateProductHTML(watch);
                // Append the generated HTML to the container
                container.append(articleHtml);
            });

            $('.products__card').on('click', '.products__button', function() {
                // Get the product details from the clicked article
                var $article = $(this).closest('.products__card');
                var container = $('.cart__container');                
                var productPrice = $article.find('.products__price').text().slice(1);
                productPrice = parseInt(productPrice, 10);
                // Loop through each cart item and update the quantity if found
                var itemFound = false;
                $.each(container.find('.cart__details'), function() {
                    if ($(this).find('.cart__title').text() === $article.find('.products__title').text()) {
                        itemFound = true;
                        var quantityText = $article.find('.products__quantity').text();
                        var quantity = parseInt(quantityText.match(/\d+/)[0]);

                        var quantityInCartText = $(this).find('.cart__amount-number').text();
                        var quantityInCart = parseInt(quantityInCartText);
                        if (quantityInCart == quantity) {
                            alert('The product amount that you want to buy cannot exceed the curren amount');
                            return;
                        }

                        addCartItemQuantity($(this), 1);
                        updatePriceForItem($(this).find('.cart__price'), productPrice, '+');
                        updateTotalPrice(productPrice, '+');
                        return; // Exit the loop since item is found
                    }
                });

                if (container.find('.cart__details').length === 4) {
                    alert("Cart only contains maximun 4 items");
                    return;
                }
                
                if (!itemFound) {
                    addToContainer(container, $article);
                    updateItemNumberTotal($('.cart__prices .cart__prices-item'), '+');
                    updateTotalPrice(productPrice, '+');
                    updateItemCount('+');
                }                
            });

            $('.cart__container').on('click', '.bx.bx-plus', function() {
                var $cartCard = $(this).closest('.cart__card');
                var cartAmountNumber = $cartCard.find('.cart__amount-number');
                
                var currentAmount = parseInt(cartAmountNumber.text(), 10);
                if (currentAmount === 10) {
                    alert('The quantity that you can buy for each item is only 10');
                    return;
                }                

                var productName = $cartCard.find('.cart__title').text();
                var price = getPrice($('.products__card'), productName);
                var cartPrice = $cartCard.find('.cart__price');
                updatePriceForItem(cartPrice, price, '+');
                updateTotalPrice(price, '+');

                var newAmount = currentAmount + 1;
                cartAmountNumber.text(newAmount);
            });

            $('.cart__container').on('click', '.bx.bx-minus', function() {
                var $cartCard = $(this).closest('.cart__card');
                var cartAmountNumber = $cartCard.find('.cart__amount-number');
                
                var currentAmount = parseInt(cartAmountNumber.text(), 10);
                if (currentAmount === 1) {
                    return;
                }

                var productName = $cartCard.find('.cart__title').text();
                var price = getPrice($('.products__card'), productName);
                var cartPrice = $cartCard.find('.cart__price');
                updatePriceForItem(cartPrice, price, '-');
                updateTotalPrice(price, '-');

                var newAmount = currentAmount - 1;
                
                cartAmountNumber.text(newAmount);
            });

            $('.cart__container').on('click', '.cart__amount-trash', function() {
                var $cartCard = $(this).closest('.cart__card');
                updateItemNumberTotal($('.cart__prices .cart__prices-item'), '-');
                var price = $cartCard.find('.cart__price').text().slice(1);
                price = parseInt(price);
                updateTotalPrice(price, '-')
                updateItemCount('-')
                $cartCard.remove();
            });

            // $('.nav__shop .bx.bx-log-in').on('click', function(){
            //     window.location.href = 'http://localhost:8888/customer/login-form';
            // });
            $('.nav__shop .bx.bx-log-out').on('click', function(){
                $.ajax({
                    url:"http://localhost:8888/customer/log-out",
                    method: "GET",
                    success: function() {
                        window.location.href = "http://localhost:8888/customer/login-form";
                    },
                    error: function(error) {
                        console.log(error)
                    }
                });    
            });

            $('.nav__search-input').keyup(function() {
                $('.warpper').hide();
                $('.products__container.grid').show();
                var watchName = $(".nav__search-input").val();
                watchName = watchName.trim();
                watchName = watchName.replace(/\s+/g, ' ');
                watchName = watchName.toUpperCase();
                $.each($('.products__card'), function(){
                    var title =  $(this).find('.products__title').text();
                    if (title.includes(watchName)){
                        $(this).show();
                    } else {
                        $(this).hide();
                    }
                });
            });

            $('#home__link').click(function(){
                $('.warpper').hide();
                $('.products__container.grid').show();
            });

            $('#order__link.nav__link').click(function(){
                $.ajax({
                    url: 'http://localhost:8888/customer/get-orders', // URL to your API or resource
                    method: 'GET',
                    dataType: 'json', // Specify the expected data type
                    success: function(response) {
                        // Handle the successful response
                        // console.log(response);
                        // $('#result').html(JSON.stringify(response));
                        $('.products__container').hide();

                        $('#one').removeAttr('checked');                        
                        $('#two').removeAttr('checked');
                        $('#three').removeAttr('checked');
                        $('#four').removeAttr('checked');
                        $('#five').removeAttr('checked');
                        $('.warpper').show();

                        

                        $('#one-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 3) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div class="order__id">Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity} </span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
                                    <button class="order__cancelation">Cancel order</button>
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#two-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 4) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>                                    
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#three-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 2) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
                                    
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#four-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 1) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#five-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                // console.log(order.Status);
                                if (order.Status == 0) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });

                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });
                    },
                    error: function(xhr, status, error) {
                        // Handle errors
                        console.error('Error:', error);
                    }
                });
            });            
            
        },
        error: function(error) {
          console.log(error)
        }
    });    
    
})

$(document).on('click', '.order__cancelation', function() {
    var text = "Do you want to cancel this order?";
    if (confirm(text)) {
        var orderCanceled = $(this).closest('.card');
        var orderID = orderCanceled.find('.order__id').text().split(':')[1].trim();
        var URL = 'http://localhost:8888/customer/cancel-orders?orderID=' + orderID
        $.ajax({
            url: URL, // URL to your API or resource
            method: 'PUT',
            dataType: 'json', // Specify the expected data type
            success: function() {
                orderCanceled.remove();
                alert("Cancel order successfully!");
                $.ajax({
                    url: 'http://localhost:8888/customer/get-orders', // URL to your API or resource
                    method: 'GET',
                    dataType: 'json', // Specify the expected data type
                    success: function(response) {

                        $('#one').removeAttr('checked');                        
                        $('#two').removeAttr('checked');
                        $('#three').removeAttr('checked');
                        $('#four').removeAttr('checked');
                        $('#five').removeAttr('checked');
                        $('#one-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 3) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div class="order__id">Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity} </span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
                                    <button class="order__cancelation">Cancel order</button>
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#two-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 4) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>                                    
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#three-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 2) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
                                    
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#four-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                if (order.Status == 1) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });
                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                }
                            });
                        });

                        $('#five-tab').click(function(){
                            var panel = $('.panel')
                            $('.card').remove();
                            $.each(response, function(_, order){
                                // console.log(order.Status);
                                if (order.Status == 0) {
                                    var html = `<div class="card" style="border-bottom: 2px solid #ccc;padding-bottom: 20px;">
                                    <div>Order ID:${order.OrderID} </div>
                                <div>Order Date: ${order.OrderDate}</div>`
                                    var totalMoney = 0;
                                    $.each(order.OrderDetails, function(_, item){
                                        totalMoney += parseInt(item.Quantity) * parseInt(item.UnitPrice)
                                        html += `<div class="item" style="display: flex;">
                                        <img src="${item.WatchImage}" class="picture"/>  
                                        <div class="info">
                                            <h4><bold>${item.WatchName}</bold></h4>
                                            
                                            <h5>Quantity: <span>${item.Quantity}</span></h5>
                                        </div>
                                        <div class="item-price">
                                            <h4 >  <span style="color: red">$${(parseInt(item.Quantity)
                                                 * parseInt(item.UnitPrice)).toLocaleString()}</span></h4>
                                        </div>
                                    </div>`;
                                    });

                                    html += `<div class="panel-footer">
                                    <div class="item-price" style="margin-right: 30px;">
                                        <svg width="16" height="17" viewBox="0 0 253 263" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M126.5 0.389801C126.5 0.389801 82.61 27.8998 5.75 26.8598C5.08763 26.8507 4.43006 26.9733 3.81548 27.2205C3.20091 27.4677 2.64159 27.8346 2.17 28.2998C1.69998 28.7657 1.32713 29.3203 1.07307 29.9314C0.819019 30.5425 0.688805 31.198 0.689995 31.8598V106.97C0.687073 131.07 6.77532 154.78 18.3892 175.898C30.003 197.015 46.7657 214.855 67.12 227.76L118.47 260.28C120.872 261.802 123.657 262.61 126.5 262.61C129.343 262.61 132.128 261.802 134.53 260.28L185.88 227.73C206.234 214.825 222.997 196.985 234.611 175.868C246.225 154.75 252.313 131.04 252.31 106.94V31.8598C252.31 31.1973 252.178 30.5414 251.922 29.9303C251.667 29.3191 251.292 28.7649 250.82 28.2998C250.35 27.8358 249.792 27.4696 249.179 27.2225C248.566 26.9753 247.911 26.852 247.25 26.8598C170.39 27.8998 126.5 0.389801 126.5 0.389801Z" fill="#ee4d2d"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M207.7 149.66L119.61 107.03C116.386 105.472 113.914 102.697 112.736 99.3154C111.558 95.9342 111.772 92.2235 113.33 88.9998C114.888 85.7761 117.663 83.3034 121.044 82.1257C124.426 80.948 128.136 81.1617 131.36 82.7198L215.43 123.38C215.7 120.38 215.85 117.38 215.85 114.31V61.0298C215.848 60.5592 215.753 60.0936 215.57 59.6598C215.393 59.2232 215.128 58.8281 214.79 58.4998C214.457 58.1705 214.063 57.909 213.63 57.7298C213.194 57.5576 212.729 57.4727 212.26 57.4798C157.69 58.2298 126.5 38.6798 126.5 38.6798C126.5 38.6798 95.31 58.2298 40.71 57.4798C40.2401 57.4732 39.7735 57.5602 39.3376 57.7357C38.9017 57.9113 38.5051 58.1719 38.1709 58.5023C37.8367 58.8328 37.5717 59.2264 37.3913 59.6604C37.2108 60.0943 37.1186 60.5599 37.12 61.0298V108.03L118.84 147.57C121.591 148.902 123.808 151.128 125.129 153.884C126.45 156.64 126.797 159.762 126.113 162.741C125.429 165.72 123.755 168.378 121.363 170.282C118.972 172.185 116.006 173.221 112.95 173.22C110.919 173.221 108.915 172.76 107.09 171.87L40.24 139.48C46.6407 164.573 62.3785 186.277 84.24 200.16L124.49 225.7C125.061 226.053 125.719 226.24 126.39 226.24C127.061 226.24 127.719 226.053 128.29 225.7L168.57 200.16C187.187 188.399 201.464 170.892 209.24 150.29C208.715 150.11 208.2 149.9 207.7 149.66Z" fill="#fff"></path></svg>
                                        <h4 >Total <span style="color: red">$${totalMoney.toLocaleString()}</span></h4>
                                    </div>
        
                                   </div>
                                   </div>`
                                   panel.append(html);
                                   console.log(10);
                                }
                            });
                        });
                    },
                    error: function(xhr, status, error) {
                        // Handle errors
                        console.error('Error:', error);
                        console.log(xhr);
                        console.log(status);
                    }
                });
            },
            error: function(xhr, status, error) {
                // Handle errors
                console.error('Error:', error);
                console.log(xhr);
                console.log(status);
            }
        });
    }
});

