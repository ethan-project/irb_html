

(function ($) {

  'use strict';

  function initDateRangrPicker() {
    if ($('#Dash_Date').length == 0) {
      return;
    }

    var picker = $('#Dash_Date');
    var start = moment();
    var end = moment();

    function cb(start, end, label) {
      var title = '';
      var range = '';

      if (end - start < 100 || label == 'Today') {
        title = 'Today:';
        range = start.format('MMM D');
      } else if (label == 'Yesterday') {
        title = 'Yesterday:';
        range = start.format('MMM D');
      } else {
        range = start.format('MMM D') + ' - ' + end.format('MMM D');
      }

      picker.find('#Select_date').html(range);
      picker.find('#Day_Name').html(title);
    }

    picker.daterangepicker({
      startDate: start,
      endDate: end,
      opens: 'left',
      applyClass: "btn btn-sm btn-primary",
      cancelClass: "btn btn-sm btn-secondary",
      ranges: {
        'Today': [moment(), moment()],
        'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
        'Last 7 Days': [moment().subtract(6, 'days'), moment()],
        'Last 30 Days': [moment().subtract(29, 'days'), moment()],
        'This Month': [moment().startOf('month'), moment().endOf('month')],
        'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
      }
    }, cb);

    cb(start, end, '');
  }

  function initMetisMenu() {
    $(".metismenu").metisMenu();
    $(window).resize(function () {
      initEnlarge();
    });
  }

  function initLeftMenuCollapse() {
    $('.button-menu-mobile').on('click', function (event) {
      event.preventDefault();
      $("body").toggleClass("enlarge-menu");
    });
  }

  function initTooltipPlugin() {
    $('[data-toggle="tooltip"]').tooltip();
  }

  function initMainIconTabMenu() {
    $('.main-icon-menu .nav-link').on('click', function (e) {
      $("body").removeClass("enlarge-menu");
      e.preventDefault();
      $(this).addClass('active');
      $(this).siblings().removeClass('active');
      $('.main-menu-inner').addClass('active');
      var targ = $(this).attr('href');
      $(targ).addClass('active');
      $(targ).siblings().removeClass('active');
    });
  }

  function initActiveMenu() {
    $(".leftbar-tab-menu a, .left-sidenav a").each(function () {
      var pageUrl = window.location.href.split(/[?#]/)[0];
      if (this.href == pageUrl) {
        $(this).addClass("active");
        $(this).parent().addClass("active");
        $(this).parent().parent().addClass("in");
        $(this).parent().parent().addClass("mm-show");
        $(this).parent().parent().parent().addClass("mm-active");
        $(this).parent().parent().prev().addClass("active");
        $(this).parent().parent().parent().addClass("active");
        $(this).parent().parent().parent().parent().addClass("mm-show");
        $(this).parent().parent().parent().parent().parent().addClass("mm-active");
        var menu = $(this).closest('.main-icon-menu-pane').attr('id');
        $("a[href='#" + menu + "']").addClass('active');
      }
    });
  }

  function initFeatherIcon() {
    feather.replace();
  }

  function initMainIconMenu() {
    $(".navigation-menu a").each(function () {
      var pageUrl = window.location.href.split(/[?#]/)[0];
      if (this.href == pageUrl) {
        $(this).parent().addClass("active");
        $(this).parent().parent().parent().addClass("active");
        $(this).parent().parent().parent().parent().parent().addClass("active");
      }
    });
  }

  function initTopbarMenu() {
    $('.navbar-toggle').on('click', function (event) {
      $(this).toggleClass('open');
      $('#navigation').slideToggle(400);
    });

    $('.navigation-menu>li').slice(-2).addClass('last-elements');

    $('.navigation-menu li.has-submenu a[href="#"]').on('click', function (e) {
      if ($(window).width() < 992) {
        e.preventDefault();
        $(this).parent('li').toggleClass('open').find('.submenu:first').toggleClass('open');
      }
    });
  }

  function init() {
    initDateRangrPicker();
    initMetisMenu();
    initLeftMenuCollapse();
    initEnlarge();
    initTooltipPlugin();
    initMainIconTabMenu();
    initActiveMenu2();
    initFeatherIcon();
    initMainIconMenu();
    initTopbarMenu();
    Waves.init();
  }

  init();
})(jQuery);