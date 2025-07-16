

$(document).ready(function(){

	$('.testimonial-slider').owlCarousel({
		loop:true,
		nav:false,
		dots: false,
		items:1,

	});

	$('.slider-section').owlCarousel({
	    loop:true,
	    margin:25,
	    nav:false,
	    dots: true,
	    dotsEach: 1,
	    responsive: {
             0: {
                 items: 1
             },
             767: {
                 items: 2
             },
             991: {
                 items: 3
             },
             1000: {
                 items: 3
             },
             1199: {
                 items: 5
             },
             1920: {
                 items: 5
            }
        }
	});

	$('.testimonial-wrapper').owlCarousel({
	    loop:true,
		nav:false,
		dots: false,
		items:1,
	});

// Mobile menu=================

		$('.menu-toggle').click(function(){
			$(this).toggleClass('active');
			$('.main-menu').slideToggle();
		})

// sticky menu===================
	    var wind = $(window);
	    var sticky = $('#sticky-header');
	    wind.on('scroll', function () {
	        var scroll = wind.scrollTop();
	        if (scroll <100) {
	            sticky.removeClass('sticky-nav');
	        } else {
	            sticky.addClass('sticky-nav');
	        }
	    });
	
});


// scroll strat============================

		 $(window).on('scroll', function () {
        var scrolled = $(window).scrollTop();
        if (scrolled > 300) $('.go-top').addClass('active');
        if (scrolled < 300) $('.go-top').removeClass('active');
    });

    $('.go-top').on('click', function () {
        $("html, body").animate({
            scrollTop: "0"
        }, 1200);
    });