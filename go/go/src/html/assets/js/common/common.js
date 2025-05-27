if (typeof const_js == 'undefined') document.write("<script src='/assets/js/common/const.js'></script>");
if (typeof api_js == 'undefined') document.write("<script src='/assets/js/common/api.js'></script>");
if (typeof ipsap_common_js == 'undefined') document.write("<script src='/assets/js/ipsap/ipsap_common.js'></script>");

var popup_window;

const COMM = {};
(function (comm, $) {
  comm.getLoginKey = function (str) {
    var tmp = str + moment().format('HH:mm:ss');
    for (var i = 0; i < tmp.length; i++) {
      if (tmp.length < 32) {
        tmp += tmp;
      } else if (tmp.length > 32) {
        tmp = tmp.substring(0, 32);
      } else {
        break;
      }
    }
    return tmp;
  };

  comm.getEncrypt = function (key, value) {
    var iv = CryptoJS.lib.WordArray.random(128 / 8);

    var utf8Key = CryptoJS.enc.Utf8.parse(key);
    var encrypted = CryptoJS.AES.encrypt(value + "", utf8Key, { iv: iv });

    var merge = iv.concat(encrypted.ciphertext);
    return CryptoJS.enc.Base64.stringify(merge);
  };

  comm.getDecrypt = function (key, value) {
    var utf8Key = CryptoJS.enc.Utf8.parse(key);
    var base64data = CryptoJS.enc.Base64.parse(value);
    var encrypted = new CryptoJS.lib.WordArray.init(base64data.words.slice(4));
    var iv = new CryptoJS.lib.WordArray.init(base64data.words.slice(0, 4));
    var cipher = CryptoJS.lib.CipherParams.create({ ciphertext: encrypted });
    var decrypted = CryptoJS.AES.decrypt(cipher, utf8Key, { iv: iv });

    return $.trim(decrypted.toString(CryptoJS.enc.Utf8).replace(/\0/g, ''));
  };

  comm.setStor = function (name, val) {
    sessionStorage.setItem(name, JSON.stringify(val));
  };

  comm.removeStor = function (name) {
    sessionStorage.removeItem(name);
  };

  comm.clearStor = function (name, val) {
    sessionStorage.clear();
  };

  comm.getStor = function (name) {
    let item = sessionStorage.getItem(name);
    if (item != null) {
      return JSON.parse(item);
    }
    return item;
  };

  comm.setCookie = function (name, val, mins) {
    mins = mins ? mins : CONST.COOKIE_EXPIRES_MIN;
    name = $.trim(encodeURIComponent(name));
    let exdate = new Date();
    exdate.setMinutes(exdate.getMinutes() + mins);
    document.cookie = name + '=' + val + '; expires=' + exdate.toString() + '; path=/';
  };

  comm.updateCookie = function (name, mins) {
    if (COMM.getCookie(name)) {
      mins = mins ? mins : CONST.COOKIE_EXPIRES_MIN;
      name = $.trim(encodeURIComponent(name));
      let exdate = new Date();
      exdate.setMinutes(exdate.getMinutes() + mins);
      document.cookie = name + '=' + COMM.getCookie(name) + '; expires=' + exdate.toString() + '; path=/';
    }
  };

  comm.removeCookie = function (name) {
    COMM.setCookie(name, '', -1 * 60 * 24);
  };

  comm.clearCookie = function () {
    let cookies = document.cookie.split('; ');
    $.each(cookies, function (key, item) {
      let arr = item.split('=');
      COMM.setCookie($.trim(arr[0]), '', -1 * 60 * 24);
    });
  };

  comm.getCookie = function (name) {
    let value = document.cookie.match('(^|;) ?' + name + '=([^;]*)(;|$)');
    return value ? value[2] : null;
  };

  comm.paramStr2Object = function (paramStr) {
    return JSON.parse('{"' + paramStr.replace(/&/g, '","').replace(/=/g, '":"') + '"}', function (key, value) {
      return key === '' ? value : decodeURIComponent(value);
    });
  };

  return comm;
})(COMM, $);

function setTriggerSelectOnChange() {
  $('select').on({
    change: e => {
      let $select = $(e.currentTarget),
          visibility = $select.find('option').filter((idx, ele) => {
        if ($(ele).attr('value') == $select.val()) {
          return true;
        }
      }).data('custom-toggle');
      if (typeof visibility !== 'undefined' && visibility == 'visible') {
        $select.parent().find('.hidden').show().filter('input:first').focus();
      } else {
        $select.parent().find('.hidden').not($select).hide();
      }
    }
  });
}

function setTriggerCheckBoxOnChange() {
  $('input[type=checkbox]').on({
    change: e => {
      let $checkbox = $(e.currentTarget);

      if (typeof $checkbox.data('custom-toggle') !== 'undefined') {
        if ($checkbox.data('custom-toggle') == 'visible') {
          if ($checkbox.prop('checked')) {
            $checkbox.parent().find('.hidden').filter((idx, ele) => {
              if ($(ele).hasClass('row')) {
                $(ele).css({ 'display': 'flex' });
              } else {
                $(ele).show();
              }
            }).filter('input:first').focus();
          } else {
            $checkbox.parent().find('.hidden').hide();
          }
        } else if ($checkbox.data('custom-toggle') == 'focus') {
          let target_id = $checkbox.data('target');
          if ($checkbox.prop('checked')) {
            $(target_id).focus();
          } else {
            $(target_id).blur();
          }
        } else if ($checkbox.data('custom-toggle') == 'disabled') {
          let target_id = $checkbox.data('target');
          if ($checkbox.prop('checked')) {
            $(target_id).prop({
              'checked': false,
              'disabled': true
            });
          } else {
            $(target_id).prop({
              'disabled': false
            });
          }
        }
      }
    }
  });
  $('input[type=radio]').on({
    change: e => {
      let $radio = $(e.currentTarget);

      if (typeof $radio.data('custom-toggle') !== 'undefined') {
        if ($radio.data('custom-toggle') == 'focus') {
          let target_id = $radio.data('target');
          if ($radio.prop('checked')) {
            $(target_id).focus();
          } else {
            $(target_id).blur();
          }
        } else if ($radio.data('custom-toggle') == 'disabled') {
          let target_id = $radio.data('target');
          if ($radio.prop('checked')) {
            $(target_id).prop({
              'checked': false,
              'disabled': true
            });
          } else {
            $(target_id).prop({
              'disabled': false
            });
          }
        }
      }
    }
  });
}

function setTriggerTrDataUrlOnClick() {
  $('tr[data-url]').on({
    click: function (e) {
      let $TRs = $(this);

      $TRs.each(function (idx, ele) {
        let $TR = $(ele),
            url = $TR.data('url'),
            id = $TR.data('id'),
            reqsvc = $TR.data('reqsvc'),
            board = $TR.data('board'),
            inst = $TR.data('inst'),
            plan = $TR.data('plan'),
            order = $TR.data('order');

        if (typeof url !== 'undefined' && url != '') {
          window.location = url;
        }

        if (typeof id !== 'undefined' && id != '') {
          IPSAP.setStor('user_seq', id);
        }

        if (typeof reqsvc !== 'undefined' && reqsvc != '') {
          IPSAP.setStor('reqsvc_seq', reqsvc);
        }

        if (typeof board !== 'undefined' && board != '') {
          IPSAP.setStor('board_seq', board);
        }

        if (typeof inst !== 'undefined' && inst != '') {
          IPSAP.setStor('inst_seq', inst);
        }

        if (typeof plan !== 'undefined' && plan != '') {
          IPSAP.setStor('plan_seq', plan);
        }

        if (typeof order !== 'undefined' && order != '') {
          IPSAP.setStor('order_seq', order);
        }
      });
    }
  });
}

function setTriggerDataMoveTo() {
  $('[data-move-to]').on({
    click: e => {

      if ($(e.currentTarget).hasClass('focus_area')) {
        $(e.currentTarget).addClass('active');
      }
      let pos_id = $(e.currentTarget).data('move-to'),
          pos = $('nav.navbar-custom').length ? $(pos_id).offset().top : $(pos_id).offset().top + 50;
      $('html, body').animate({ scrollTop: pos }, 300);
    }
  });

  $('select.move-to-selected').on({
    change: e => {
      if (!$(e.currentTarget).val() || $(e.currentTarget).val() == '#') {} else {
        let pos_id = $(e.currentTarget).val(),
            pos = $('nav.navbar-custom').length ? $(pos_id).offset().top : $(pos_id).offset().top + 50;
        $('html, body').animate({ scrollTop: pos }, 300);
      }
    }
  });
}

function setTriggerCheckBoxOnClickKeydownChange() {
  $('input[type=checkbox]').on({
    'click keydown change': function (e) {
      let $this_check = $(this),
          disabled_status = false;

      if ($(this).is('[readonly]')) {
        return false;
      }

      if ($this_check.data('custom-toggle') == 'disable') {
        if ($this_check.prop('checked')) {
          disabled_status = false;
        } else {
          disabled_status = true;
        }
        $($this_check.data('custom-target')).prop('disabled', disabled_status);
      }
    }
  });
}

function setTriggerRadioOnChange() {
  $('input[type=radio]').on({
    click: function (e) {
      let $this_radio = $(this),
          radio_name = $this_radio.attr('name'),
          $radios = $('input[name=' + radio_name + ']');

      if ($this_radio.data('custom-toggle') == 'visible') {
        $this_radio.parent().find('.hidden').show().filter('input:first').focus();
        $radios.not($this_radio).each(function () {
          $(this).parent().find('.hidden').hide();
        });
      }
    }
  });
}

var ipsapLogOut = function () {
  $('a#ipsapLogout').click(function (e) {
    COMM.clearCookie();
    IPSAP.clearStor('institution_list_data');
    IPSAP.clearStor('user_seq');
    IPSAP.clearStor('reqsvc_seq');
    IPSAP.clearStor('board_seq');
    IPSAP.clearStor('inst_seq');
    IPSAP.clearStor('plan_seq');
    IPSAP.clearStor('order_seq');
  });
};

function initBreadcrumbs() {
  let this_path = [],
      full_path = '';

  this_path.push($('.metismenu>li.active>a').text());
  if ($('a.nav-link.active').not('[data-toggle="tab"]').text()) {
    this_path.push('<i class="ti-control-record mRight3"></i>' + $('a.nav-link.active').text());
  }
  this_path.push($('.page-title').text());

  $.each(this_path, (idx, ele) => {
    if (idx) {
      full_path += `<i class="dripicons-chevron-right mHside3"></i><span>${ele}</span>`;
    } else {
      full_path += `<span>${ele}</span>`;
    }
  });

  $('.page-title-box .page-title').not('.noBread').after(`<span id='breadcrumbs'>
    <i class='far fa-folder-open mRight5'></i>${full_path}
  </span>`);
}

function initMetisMenu() {
  $(".metismenu").metisMenu();

  let limit_folding_menu = $('body').data('fold-menu') * 1,
      limit_folding_guide = $('body').data('fold-guide') * 1,
      urlPathName = window.location.pathname;

  $(window).on({
    'load resize': function () {
      if (!limit_folding_menu) {
        if (urlPathName.indexOf('application_new-') > -1 || urlPathName.indexOf('application_list-') > -1 || urlPathName.indexOf('reviewConfirm_list-') > -1) {
          initEnlarge(1500);
        } else if (urlPathName.indexOf('reviewOffice_list-') > -1) {
          initEnlarge(1300);
        } else {
          initEnlarge();
        }
      } else {
        initEnlarge(limit_folding_menu);
      }

      if (!limit_folding_guide) {
        initEnlargeGuide();
      } else {
        initEnlargeGuide(limit_folding_guide);
      }

      if (checkAppOverflow('.app_content>.col-sm-9')) {
        $('body').addClass('hidden_guide');
      } else {
        $('body').removeClass('hidden_guide');
      }
    }
  });
}

function checkEleOverflow(ele_id) {
  $(ele_id).each((idx, ele) => {
    let curOverflow = ele.style.overflow;
    if (!curOverflow || curOverflow === "visible") {
      ele.style.overflow = "hidden";
    }
    let isOverflow = ele.clientWidth < ele.scrollWidth || ele.clientHeight + 2 < ele.scrollHeight;
    ele.style.overflow = curOverflow;

    return isOverflow;
  });
}
function checkAppOverflow(ele_id) {
  isOverflowEven = false;
  $(ele_id).each((idx, ele) => {
    let curOverflow = ele.style.overflow;
    if (!curOverflow || curOverflow === "visible") {
      ele.style.overflow = "hidden";
    }
    let isOverflowOne = ele.clientWidth < ele.scrollWidth || ele.clientHeight + 2 < ele.scrollHeight;
    ele.style.overflow = curOverflow;

    isOverflowEven = isOverflowEven || isOverflowOne;
  });

  return isOverflowEven;
}

function initLeftMenuCollapse() {
  $('.button-menu-mobile').on('click', e => {
    e.preventDefault();
    $("body").toggleClass("enlarge-menu").removeClass('enlarge-menu-all');
  });
}

function initEnlarge(max_size = 1025) {
  if ($(window).width() < max_size) {
    $('body').addClass('enlarge-menu enlarge-menu-all');
  } else {
    if ($('body').hasClass('enlarge-menu-all')) {
      $('body').removeClass('enlarge-menu enlarge-menu-all');
    }
  }
}
function initEnlargeGuide(max_size = 1400) {
  if ($(window).width() < max_size) {
    $('body').addClass('hidden_guide');
  } else {
    $('body').removeClass('hidden_guide');
  }
}

function initActiveMenu2() {
  var pageUrl = window.location.href.split(/[?#]/)[0].split('-')[0].replace('.html', '').split('/').slice(-1)[0];

  $(".leftbar-tab-menu a, .left-sidenav a").each((idx, ele) => {
    let menuUrl = ele.href.split('/').slice(-1)[0].split('?')[0].split('-')[0].replace('.html', '');
    let menu = '';

    if (menuUrl == pageUrl) {
      $(ele).addClass("active");
      $(ele).parent().addClass("active");
      $(ele).parent().parent().addClass("in");
      $(ele).parent().parent().addClass("mm-show");
      $(ele).parent().parent().parent().addClass("mm-active");
      $(ele).parent().parent().prev().addClass("active");
      $(ele).parent().parent().parent().addClass("active");
      $(ele).parent().parent().parent().parent().addClass("mm-show");
      $(ele).parent().parent().parent().parent().parent().addClass("mm-active");
      menu = $(ele).closest('.main-icon-menu-pane').attr('id');
      $(`a[href="#${menu}"]`).addClass('active');
    }
  });

  initBreadcrumbs();
}

function scrollSpy() {
  var current_id,
      $anchors = $('.app_anchor').not('#top_anchor, #bottom_anchor, #review_top');

  $anchors.each((idx, ele) => {
    if ($(ele).offset().top <= $(window).scrollTop() + 40) {
      current_id = $(ele).attr('id').split('-')[0];
    }
  });

  if ($('#footer_anchor').length) {
    var $footer_anchor = $('#footer_anchor');
    if ($footer_anchor.offset().top <= window.pageYOffset + window.innerHeight) {
      $footer_anchor.next('.card-footer').removeClass('side_type');
    } else {
      $footer_anchor.next('.card-footer').addClass('side_type');
    }
  }

  if ($('.process_content').length) {
    let $last_anchor = $anchors.last();
    if ($last_anchor.length) {
      if ($last_anchor.offset().top <= $(window).scrollTop() + window.innerHeight - 100) {
        current_id = $anchors.last().attr('id') ? $anchors.last().attr('id').split('-')[0] : '';
      }
    }
    $('.process_content').removeClass('active');
    $(`.process_content[data-move-to^="#${current_id}"]`).addClass('active');

    if (!$('.process_content.active').length) {
      $('.process_content:first').addClass('active');
    }
  }

  if ($('.app_frame').length) {
    if ($('.app_frame').offset().top <= $(window).scrollTop() + 60) {
      $('.app_frame').addClass('summarized');
    } else {
      $('.app_frame').removeClass('summarized');
    }
  }
}

function scrollTop(ele) {
  $('html, body').animate({ scrollTop: $(ele).offset().top }, 500);
}

function newWindowPopup(url, width = 1350, height = window.screen.height) {
  width = !width ? 1350 : width;
  let leftPosition = window.screen.width / 2 - (width / 2 + 5),
      topPosition = window.screen.height / 2 - (height / 2 + 25);

  popup_window = window.open(url, '_blank', 'width=' + width + ', height=' + height + ', top=' + topPosition + ', left=' + leftPosition + ',screenY=' + topPosition + ',toolbar=no, menubar=no, scrollbars=no, location=no, directories=no');
}

function printRegion(region_id, ignore_id) {
  let $cloned_region = $(region_id).clone();

  $cloned_region.find(ignore_id).remove();
  $('body').children().css('display', 'none');
  $('body').append($cloned_region);
  $cloned_region.find('#page_breaker').css('margin-top', '200px');

  window.print();

  setTimeout(() => {
    $cloned_region.remove();
    $('body').children().removeAttr('style');
  }, 500);
}

function setPrintableForm() {
  $('#print_region').addClass('printable');
  $('input[type=checkbox]').not('.custom-control-input').each((idx, ele) => {
    let $this_form = $(ele).filter((idx, ele) => {
      return !$(ele).closest('.modal').length && $(ele).prop('checked');
    }),
        $this_label = $this_form.next('label'),
        label_class = $this_label.attr('class'),
        $wrapper = $this_form.closest('.checkbox');

    $this_form.replaceWith('<i class="fas fa-check-square"></i>');
    $this_label.replaceWith(`<span class="data">${$this_label.text()}</span>`);
    $wrapper.removeClass('pLeft3');
    if ($wrapper.parent('th').length || $wrapper.parent('td').length) {
      if ($wrapper.hasClass('mTop5')) {
        $wrapper.removeClass('mTop5').addClass('mTop2');
      }
    } else if ($wrapper.parent('.flexMid').length) {} else if ($wrapper.closest('th').length || $wrapper.closest('td').length) {} else {
      $wrapper.addClass('pBot7');
    }
  });

  $('input[type=radio]').each((idx, ele) => {
    let $this_form = $(ele).filter((idx, ele) => {
      return $(ele).prop('checked') && !$(ele).closest('.modal').length;
    }),
        $this_label = $this_form.next('label'),
        label_class = $this_label.attr('class');

    if ($this_label.length) {
      $this_form.replaceWith('<i class="fas fa-dot-circle"></i>');
      $this_label.replaceWith(`<span class="data">${$this_label.text()}</span>`);
    } else {
      $this_form.replaceWith('<i class="fas fa-dot-circle mRight0"></i>');
    }
  });

  $('select.form-control').each((idx, ele) => {
    let $this_form = $(ele).filter((idx, ele) => {
      return !$(ele).closest('.modal').length;
    }),
        $this_val = $this_form.find('option[selected]:not([disabled])');

    if ($this_val.length) {
      $this_form.replaceWith(`<span class="data">${$this_val.text()}</span>`);
    } else {
      $this_form.replaceWith(`-`);
    }
  });
}

function printPdf(register_num) {
  html2canvas(document.body, {
    onrendered: canvas => {
      var imgData = canvas.toDataURL('image/png'),
          pageWidth = 210,
          pageHeight = 297,
          imgHeight = canvas.height * pageWidth / canvas.width,
          restHeight = imgHeight,
          doc = new jsPDF('p', 'mm', 'a4'),
          position = 0;

      doc.addImage(imgData, 'PNG', 0, position, pageWidth, imgHeight);
      restHeight -= pageHeight;

      while (restHeight >= 20) {
        position = restHeight - (imgHeight - 10);
        doc.addPage();
        doc.addImage(imgData, 'PNG', 0, position, pageWidth, imgHeight);
        restHeight -= pageHeight;
      }

      doc.save(register_num);
      $('body').removeClass('pdfFormat');

      if ($('#modal_spinner').length) {
        $('#modal_spinner').modal('hide');
      }
    }
  });
}

function closeWindow() {
  window.location = '/';
}

function appSubmitHandleError(msg, data) {
  if (data.em) {
    alert(`${msg}\n(${data.em})`);
  } else if (data.msg) {
    alert(`${msg}\n(${data.msg})`);
  } else {
    alert(`${msg}`);
  }
}

function typeAndAuthCheckAfterRedirect(user_type, user_auth) {

  let service_status = Number(JSON.parse(COMM.getCookie('institution_info'))['service_status']);

  switch (parseInt(user_auth)) {
    case IPSAP.USER_AUTH.AUTH_NOMARL:
    case IPSAP.USER_AUTH.AUTH_INSTITUTION:
      if (user_type[IPSAP.USER_TYPE.ADMIN_OFFICER]) {
        location.replace('/html/officer/dashboard.html');
      } else if (user_type[IPSAP.USER_TYPE.OFFICER]) {
        location.replace('/html/officer/dashboard.html');
      } else if (user_type[IPSAP.USER_TYPE.CHAIRMAN]) {
        location.replace('/html/chairman/dashboard.html');
      } else if (user_type[IPSAP.USER_TYPE.COMMITTEE]) {
        if (service_status == 1) {
          location.replace('/html/committee/review_list.html');
        } else {
          location.replace('/html/committee/info_me.html');
        }
      } else if (user_type[IPSAP.USER_TYPE.RESEARCHER]) {
        location.replace('/html/researcher/dashboard.html');
      } else {
        alert('권한이 없습니다.');
      }
      break;
    case IPSAP.USER_AUTH.AUTH_PLATFORM:
    case IPSAP.USER_AUTH.AUTH_SYSTEM:
      location.replace('/html/admin/dashboard.html');
      break;
  }
}

function openModal(target_id) {
  $(target_id).modal('show');
}

function stackCards($myStacks) {
  $myStacks.each((idx, ele) => {
    let card_cnt = $(ele).find('.card:not(.empty_apps)').length,
        card_height = 134;
    if ($(ele).hasClass('exp_area')) {
      card_height = 90;
    }
    if ($(ele).height() < card_cnt * card_height + 20) {
      $(ele).addClass('stacked');
    } else {
      $(ele).removeClass('stacked');
    }
  });
}

function setDashboardFilters() {
  $('.list_filters').on({
    click: e => {
      let $filter = $(e.currentTarget),
          typeClass = $filter.data('filter');
      if ($filter.hasClass('active')) {
        $('.list_filters').removeClass('active');
        $('.myApps').removeClass('application report');
        $('.myApps').find('.card').removeAttr('style');
        $('.myApps').removeClass('empty_result').find('.empty_result').remove();
      } else {
        $('.list_filters').removeClass('active');
        $filter.addClass('active');
        $('.myApps').removeClass('application report').addClass(typeClass);
        $('.myApps').each((idx, ele) => {
          $(ele).find(`.card.${typeClass}`).first().css({ 'margin-top': 'initial' });
          if (!$(ele).find('.card').filter((idx2, ele2) => {
            return $(ele2).is(':visible');
          }).length) {
            $(ele).addClass('empty_result').find('.mCSB_container').append(empty_result);
          } else {
            $(ele).removeClass('empty_result').find('.empty_result').remove();
          }
        });
      }
      updateStackCounter($('.myApps'), '.card', '.myApp_cnt');
    }
  });

  $('.committee_filters').on({
    click: e => {
      let $filter = $(e.currentTarget);
      if (!$filter.hasClass('inactive')) {
        let typeClass = $filter.data('filter').substring(5);
        if ($filter.hasClass('active')) {
          $('.committee_filters').removeClass('op_30 active inactive');
          $('.myApps, .exp_area').find('.card').removeAttr('style');
          $('.myApps, .exp_area').removeClass('only_IACUC only_IBC only_IRB');
          $('.myApps, .exp_area').removeClass('empty_result').find('.empty_result').remove();
        } else {
          $('.committee_filters').removeClass('active').addClass('op_30 inactive');
          $filter.addClass('active').removeClass('op_30 inactive');
          $('.myApps, .exp_area').removeClass('only_IACUC only_IBC only_IRB').addClass(`only_${typeClass}`);
          $('.myApps, .exp_area').each((idx, ele) => {
            $(ele).find(`.card.${typeClass}`).first().css({ 'margin-top': '0' });
            if (!$(ele).find('.card').filter((idx2, ele2) => {
              return $(ele2).is(':visible');
            }).length) {
              $(ele).addClass('empty_result').find('.mCSB_container').append(empty_result);
            }
          });
        }
        updateStackCounter($('.myApps'), '.card', '.myApp_cnt');
      }
    }
  });

  $('#statistics_trigger').on({
    click: e => {
      let $trigger = $(e.currentTarget);
      if ($trigger.hasClass('active')) {
        $trigger.removeClass('active');
        $('.area_wrapper').removeClass('statistics');
      } else {
        $trigger.addClass('active');
        $('.area_wrapper').addClass('statistics');
      }
    }
  });

  $('.exp_area').filter(() => {
    return !$('body').hasClass('chairman');
  }).on({
    mouseenter: e => {
      $('.dashboard_box').addClass('exp_active');
      $(e.currentTarget).removeClass('stacked');
    },
    mouseleave: e => {
      $('.dashboard_box').removeClass('exp_active');
    },
    transitionend: e => {
      if (e.target == e.currentTarget) {
        stackCards($(e.currentTarget));
      }
    },
    animationend: e => {}
  });

  if ($('.myApps, .myExps').length) {
    $(".myApp_body, .myExps").mCustomScrollbar({
      axis: 'y',
      scrollbarPosition: "inside",
      scrollInertia: 300,
      autoHideScrollbar: true,
      autoExpandScrollbar: false,
      theme: 'dark-2'
    });
  }
}
function updateStackCounter($wrappers, object_id, counter_id, isInit = false) {
  $wrappers.each((idx, ele) => {
    let $wrapper = $(ele),
        obj_cnt = $wrapper.find(`${object_id}:visible:not(.empty_apps):not(.empty_result)`).length,
        $counter = $wrapper.find(counter_id);

    let update_class = isInit ? '' : 'empty';
    if (obj_cnt > 0) {
      $wrapper.removeClass(update_class);
      $counter.text(obj_cnt);
    } else {
      $wrapper.addClass(update_class);
      $counter.text(obj_cnt);
    }
  });
}

function setComma(pureNumber) {
  return pureNumber.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

function maxLengthCheck(object) {
  if (object.value.length > object.maxLength) {
    object.value = object.value.slice(0, object.maxLength);
  }

  if (object.value > parseInt(object.max)) {
    object.value = parseInt(object.max);
  }

  object.value = object.value.replace(/(^0+)/, "");
}

var currentPath = window.location.pathname,
    pathArray = currentPath.split('/'),
    currentUser = pathArray[2],
    leftMenu_visibility = false,
    nav_btn_visibility = true;

function serviceStatusCheckLeftNav() {

  let service_status = Number(JSON.parse(COMM.getCookie('institution_info'))['service_status']);

  switch (service_status) {
    case 0:
    case 3:
    case 2:
      $('.li1, .li2, .li3, .li4').css('display', 'none');
      break;
  }
}

function serviceStatusCheckTopBar() {

  if (COMM.getCookie('institution_info')) {
    var service_status = Number(JSON.parse(COMM.getCookie('institution_info'))['service_status']);
  }

  switch (service_status) {
    case 0:
    case 3:
    case 2:
      var inputUrl = window.location.pathname;
      var nonAcceptableArr = ['application', 'experiment', 'experimenting', 'inspection', 'review'];
      for (var i = 0; i < nonAcceptableArr.length; i++) {
        if (inputUrl.indexOf(nonAcceptableArr[i]) > -1) {
          alert('현재 소속된 기관의 이용 상태로는 접근할 수 없는 경로입니다.');
          location.replace('/index.html');
        }
      }
      $('#gnb_menu').children().eq(1).css('display', 'none');
      break;
  }
}

if (COMM.getCookie('user_info')) var currentUserType = JSON.parse(COMM.getCookie('user_info'))["user_type"];

$(function () {
  window.alert = function (msg) {
    let iframe = document.createElement("IFRAME");
    iframe.style.display = "none";
    iframe.setAttribute("src", 'data:text/plain');
    document.documentElement.appendChild(iframe);
    window.frames[0].window.alert(msg);
    iframe.parentNode.removeChild(iframe);
  };

  var inst_logo_src = COMM.getCookie('logo_file_src'),
      user_info = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));

  switch (currentUser) {

    case 'researcher':
      if (!currentUserType || !currentUserType[IPSAP.USER_TYPE.RESEARCHER]) {
        location.replace('/index.html');
      }

      $('#leftsidenav').load("/html/common/inc/leftsidenav.html #researcher", function () {
        feather.replace();
        initMetisMenu();
        initLeftMenuCollapse();
        initEnlarge();
        initActiveMenu2();
        $('.logo').html(`<img height="60" src="${inst_logo_src}">`);
        serviceStatusCheckLeftNav();
      });
      leftMenu_visibility = true;
      break;

    case 'officer':
      if (!currentUserType || !currentUserType[IPSAP.USER_TYPE.ADMIN_OFFICER] && !currentUserType[IPSAP.USER_TYPE.OFFICER]) {
        location.replace('/index.html');
      }

      $('#leftsidenav').load("/html/common/inc/leftsidenav.html #officer", function () {
        feather.replace();
        initMetisMenu();
        initLeftMenuCollapse();
        initEnlarge();
        initActiveMenu2();
        $('.logo').html(`<img height="60" src="${inst_logo_src}">`);
        if (user_info.user_type[IPSAP.USER_TYPE.ADMIN_OFFICER]) {
          $('.auth_label').text('행정 간사');
        } else if (user_info.user_type[IPSAP.USER_TYPE.OFFICER]) {
          $('.auth_label').text('행정 담당');
          $('a[href="/html/officer/reviewConfirm_list.html"]').parent().css('display', 'none');
        }

        serviceStatusCheckLeftNav();
      });
      leftMenu_visibility = true;
      break;

    case 'committee':
      if (!currentUserType || !currentUserType[IPSAP.USER_TYPE.COMMITTEE]) {
        location.replace('/index.html');
      }

      $('#leftsidenav').load("/html/common/inc/leftsidenav.html #committee", function () {
        feather.replace();
        initMetisMenu();
        initLeftMenuCollapse();
        initEnlarge();
        initActiveMenu2();
        $('.logo').html(`<img height="60" src="${inst_logo_src}">`);
        serviceStatusCheckLeftNav();
      });
      leftMenu_visibility = true;
      break;

    case 'chairman':
      if (!currentUserType || !currentUserType[IPSAP.USER_TYPE.CHAIRMAN]) {
        location.replace('/index.html');
      }

      $('#leftsidenav').load("/html/common/inc/leftsidenav.html #chairman", function () {
        feather.replace();
        initMetisMenu();
        initLeftMenuCollapse();
        initEnlarge();
        initActiveMenu2();
        $('.logo').html(`<img height="60" src="${inst_logo_src}">`);
        serviceStatusCheckLeftNav();
      });
      leftMenu_visibility = true;
      break;

    case 'register':
      $('.page-content').addClass('pHside30');
      break;

    case 'member':
      leftMenu_visibility = true;
      nav_btn_visibility = false;
      break;

    default:
      leftMenu_visibility = false;
      break;

  }

  $('#topbar').load("/html/common/inc/topbar.html", function () {
    feather.replace();
    toastMsg();

    function initUserAuthMenu() {

      let service_status = Number(JSON.parse(COMM.getCookie('institution_info'))['service_status']);

      if (!nav_btn_visibility) {
        $('#nav_btn').remove();
        $('.auth_menu').before($('.logo_single'));
      }

      if (IPSAP.DEMO_MODE) return;
      var user_info = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));

      var cb_url = window.location.pathname;
      if (user_info == null) {
        location.replace('/index.html?' + cb_url);
        return;
      }
      var user_type = user_info.user_type;

      if (!user_type[IPSAP.USER_TYPE.ADMIN_OFFICER]) {
        $('.auth_menu').find('.officer').html(`<i class="fas fa-check-circle onlyActive mRight5"></i>행정 담당`);
      }
      if (!user_type[IPSAP.USER_TYPE.OFFICER]) {
        $('.auth_menu').find('.officer').html(`<i class="fas fa-check-circle onlyActive mRight5"></i>행정 간사`);
      }
      if (!user_type[IPSAP.USER_TYPE.OFFICER] && !user_type[IPSAP.USER_TYPE.ADMIN_OFFICER]) {
        $('.auth_menu').find('.officer').css('display', 'none');
      }
      if (!user_type[IPSAP.USER_TYPE.CHAIRMAN]) {
        $('.auth_menu').find('.chairman').css('display', 'none');
      }
      if (!user_type[IPSAP.USER_TYPE.COMMITTEE]) {
        $('.auth_menu').find('.committee').css('display', 'none');
      }
      if (!user_type[IPSAP.USER_TYPE.RESEARCHER]) {
        $('.auth_menu').find('.researcher').css('display', 'none');
      }

      if (user_type[IPSAP.USER_TYPE.COMMITTEE] && service_status != 1) {
        $('.auth_menu').find('.committee').attr('href', '/html/committee/info_me.html');
      }

      $('.user_name').text(user_info.name);
      $('.user_depart').text(user_info.institution_name_ko + " (" + user_info.dept_str + ")");
    }

    serviceStatusCheckTopBar();

    if (leftMenu_visibility) {
      if (nav_btn_visibility) {
        initUserAuthMenu();
        initMetisMenu();
        initLeftMenuCollapse();
        initEnlarge();
      } else {
        initUserAuthMenu();
      }
    } else {
      $('.top_left').remove();
      $('.topbar-nav').remove();
      $('#gnb_menu').css({ 'position': 'relative' });
      $('footer').removeClass('text-sm-left').addClass('text-sm-center');
      $('body').addClass('noLeftMenu');
    }

    if (IPSAP.DEMO_MODE) return;

    API.load({
      url: CONST.API.INSTITUTION.MY_OTHER_INST,
      type: CONST.API_TYPE.GET,
      success: function (data) {
        var table = $('.inst_list').find('table').eq(0);
        var tableObj = table.clone();
        $('.inst_list').children().not('div.text-center').remove();

        if (data.length) {
          $.each(data, function (key, item) {
            tableObj.find('tbody').children().eq(0).find('.name_ko').text(item.institution_name_ko);
            tableObj.find('tbody').children().eq(0).find('.name_en').text(item.institution_name_en);
            tableObj.find('tbody').children().eq(1).find('.data').text(item.institution_code);
            let label_html = '';
            let user_type_arr = [];
            user_type_arr.push(item.user_type_str.split(','));
            $.each(user_type_arr[0], function (i, v) {
              label_html += `<span class="label auth_label">`;
              label_html += v + `</span>`;
            });
            tableObj.find('tbody').children().eq(2).find('.data').html(label_html);
            tableObj.find('tbody').children().eq(3).find('#dept').text(item.dept_str);
            tableObj.find('tbody').children().eq(3).find('#position').text(item.position_str);
            tableObj.find('tbody').children().eq(3).find('#major').text(item.major_field_str);
            tableObj.find('tbody').attr('data-data', item.user_seq);

            var html = tableObj.wrapAll("<table/>").parent().html();
            $('.inst_list').find('.text-center').after(html);
          });
        } else {
          $('.joining_comment').html(`현재 접속한 기관 외에 내가 <span class="red">소속된 타기관이 없습니다.</span>`);
        }
      }
    });

    $(document).on('click', '.inst_move', function () {
      var tmp_key = COMM.getCookie(CONST.COOKIE.IPSAP_TMP_KEY);
      var user_seq = $('.inst_list').find('table').find('tbody').data('data');
      var param = {
        tmp_key: tmp_key,
        user_seq: user_seq
      };

      API.load({
        url: CONST.API.INSTITUTION.MOVE_INST,
        type: CONST.API_TYPE.POST,
        data: param,
        success: function (data) {
          let user_info = data[0].user_info;
          let institution_info = data[0].ins_info;
          COMM.setCookie(CONST.COOKIE.IPSAP_TMP_KEY, COMM.getDecrypt(param.tmp_key, data[0].tmp_key));
          COMM.setCookie('user_info', JSON.stringify(user_info));
          COMM.setCookie('institution_info', JSON.stringify(institution_info));
          COMM.setCookie('logo_file_src', institution_info.logo_file_src);
          COMM.setCookie(CONST.COOKIE.IPSAP_TOKEN, data[0].token);
          typeAndAuthCheckAfterRedirect(user_info.user_type, user_info.user_auth);
        },
        error: function (err) {}
      });
    });

    ipsapLogOut();
  });

  $('#footer').load("/html/common/inc/footer.html", function () {
    feather.replace();
  });

  $('.btn_shrink').on({
    click: function (e) {
      e.stopPropagation();

      if ($(this).find('i').hasClass('fa-arrow-down')) {
        //down
        $(this).find('i').switchClass('fa-arrow-down', 'fa-arrow-up');
        $(this).closest('.modal-dialog').css({
          'margin-right': '0',
          'left': '0',
          'top': window.innerHeight - 80
        });
      } else {
        // up
         // height
          var screenHeight = $(window).height();
          console.log("screenHeight", screenHeight);
          var modalHeight = $("#modal_changeApp .modal-content").outerHeight();
          console.log("modalHeight", modalHeight);
          //margin
          var marginTop = parseInt($(".modal-dialog").css("margin-top"), 10);
          var marginBottom = parseInt($(".modal-dialog").css("margin-bottom"), 10);
          console.log("Top Margin:", marginTop, "px");
          console.log("Bottom Margin:", marginBottom, "px");
          //calcu
          var topPosition = screenHeight - modalHeight - marginTop - marginBottom;
          console.log("topPosition", topPosition);
        $(this).find('i').switchClass('fa-arrow-up', 'fa-arrow-down');
        $(this).closest('.modal-dialog').css({
          'margin-right': '0',
          'left': '0',
          // 'top': '0'
          'top': topPosition +"px"
        });
      }
    }
  });

  setTriggerDataMoveTo();

  $('#move_top').on({
    click: function (e) {
      e.stopPropagation();

      $('.focus_area:first').addClass('active');
    }
  });

  $(document).on('change', 'input[type=file]', function (e) {
    if (window.FileReader) {
      var filename = $(this)[0].files[0].name;
    } else {
      var filename = $(this).val().split('/').pop().split('\\').pop();
    }
    $(this).siblings('.custom-file-label')[0].innerHTML = filename;
  });

  $(document).on('keyup', 'input[type=number]', function (e) {
    var inputVal = $(this).val();
    $(this).val(inputVal.replace(/[^0-9]/gi, ''));

    if (!(e.keyCode > 95 && e.keyCode < 106 || e.keyCode > 47 && e.keyCode < 58 || e.keyCode == 8)) {
      return false;
    }
  });

  setTriggerSelectOnChange();
  setTriggerCheckBoxOnChange();
  setTriggerCheckBoxOnClickKeydownChange();
  setTriggerRadioOnChange();

  $('button').on({
    click: e => {
      let $button = $(e.currentTarget);

      if ($button.data('custom-toggle') == 'visible') {
        $button.parent().find('.hidden').filter((idx, ele) => {
          if (!$(ele).attr('style')) {
            return true;
          } else {
            $(ele).removeAttr('style');
          }
        }).show();
      }
    }
  });

  setTriggerTrDataUrlOnClick();

  $('.action_buttons, .action_buttons a').on({
    click: function (e) {
      e.stopPropagation();
    }
  });
  $('a.disabled').on({
    click: function (e) {
      return false;
    }
  });

  ipsapLogOut();

  let url = window.location.href;
  if (url.indexOf('application_new') < 0) {
    if (!$('body').hasClass('noScroll')) {
      scrollSpy();
    }
  }
});

var isScrollProgress = false;
$(window).on({
  load: function (e) {
    let url = window.location.href,
        urlArry = url.split("#"),
        urlCore = urlArry[0],
        urlHash = urlArry[1],
        urlQuery = url.split("?")[1];

    $('.modal.fade').on({
      'shown.bs.modal': e => {
        $(e.currentTarget).find('[autofocus]:visible').focus();
      }
    });
    $('[data-toggle=tab]').on({
      'shown.bs.tab': e => {
        e.target;
        e.relatedTarget;
        $($(e.target).attr('href')).find('[autofocus]:visible').focus();
      }
    });

    $('[data-toggle=modal]').on({
      click: e => {
        var is_full_modal = false;
        if (typeof $(e.currentTarget).data('modal-type') !== 'undefined') {
          if ($(e.currentTarget).data('modal-type') == 'full') {
            $('body').addClass("full_modal");
            $($(e.currentTarget).data('target')).addClass('full_modal').find('[data-toggle=modal]').data('modal-type', 'full');
            is_full_modal = true;
          }
        }
        if (!is_full_modal) {
          $('body').removeClass("full_modal");
          $($(e.currentTarget).data('target')).removeClass('full_modal').find('[data-toggle=modal]').data('modal-type', 'pop');
        }

        if (typeof $(e.currentTarget).data('target-tab') !== 'undefined') {
          $($(e.currentTarget).data('target')).find($(e.currentTarget).data('target-tab')).trigger('click');
        }
      }
    });

    if ($('.page-wrapper.popup').length) {
      var $page_wrapper = $('.page-wrapper.popup'),
          pop_width = window.outerWidth,
          frame_height = window.outerHeight - window.innerHeight,
          pop_height = window.screen.height;

      if (currentPath.indexOf('_popup_application') > -1) {
        window.resizeTo(pop_width, pop_height);
      } else {
        setTimeout(() => {
          pop_height = $('.page-content').children().height() + 40 + frame_height;
          pop_height = pop_height > window.screen.height ? window.screen.height : pop_height;
          window.resizeTo(pop_width, pop_height);

          let leftPosition = window.screen.width / 2 - (pop_width / 2 + 5),
              topPosition = window.screen.height / 2 - (pop_height / 2 + 25);
          window.moveTo(leftPosition, topPosition);
        }, 1000);
      }
    }

    $('input[type=checkbox], input[type=radio]').filter((idx, ele) => {
      return $(ele).is('[readonly]');
    }).on('click', e => {
      return false;
    });

    $('[data-trigger=click]').on({
      click: e => {
        if (typeof $(e.currentTarget).data('target') && $(e.currentTarget).data('target') != '') {
          $($(e.currentTarget).data('target')).trigger($(e.currentTarget).data('trigger'));
        }
      }
    });

    $('.tippy-btn').each((idx, ele) => {
      let tippy_offset = !$(ele).data('offset') ? '0, 0' : $(ele).data('offset'),
          tippy_target = !$(ele).data('tippy-target') ? ele : $(ele).data('tippy-target').split(',');

      tippy(ele, {
        offset: tippy_offset,
        triggerTarget: tippy_target,
        arrow: true
      });
    });

    collapseFun();

    if (url.indexOf('review_list-review_IACUC_expert') > -1) {
      $('.trace_label').remove();
    }
    setTimeout(function () {
      if (!!urlHash && $('#' + urlHash).length) {
        if (urlCore.indexOf('html/ipsap/experiment_02.html') > -1) {
          $('a[href="#' + urlHash + '"]').trigger('click');
          return false;
        }
        $('html, body').animate({ scrollTop: $('#' + urlHash).offset().top + 50 }, 500);
      }

      if (!!urlQuery && urlQuery.length) {
        if (urlQuery.indexOf('checkto=') > -1) {
          let $checkbox = $('#' + urlQuery.split('=')[1]);
          $checkbox.attr('checked', '').prop('checked', true);
          return false;
          if ($checkbox.length) {
            $('html, body').animate({ scrollTop: $checkbox.offset().top }, 500);
          }
        }
        if (urlQuery.indexOf('clickto=') > -1) {
          let $clickObj = $('#' + urlQuery.split('=')[1]);
          $clickObj.trigger('click');
          return false;
          if ($clickObj.length) {
            $('html, body').animate({ scrollTop: $clickObj.offset().top }, 500);
          }
        }
      }

      if ($('a[href^="javascript:printRegion"]').length && !$('body').hasClass('noPrint')) {
        setPrintableForm();
      }

      $('#btn_pdf').click(function () {
        let register_num = $('.register_num').eq(0).text();
        $('body').addClass('pdfFormat');
        $(window).scrollTop(0);
        setTimeout(function () {
          printPdf(register_num);
          if ($('#modal_spinner').length) {
            $('#modal_spinner').modal('show');
          }
        }, 300);
      });
    }, 300);
    if (!$('body').hasClass('noScroll')) {
      if (url.indexOf('application_new') < 0) {
        isScrollProgress = true;
      }
    }
  },

  resize: e => {
    let $myApps = $('.myApps');
    if ($myApps.length) {
      stackCards($myApps);
    }
  },

  scroll: function (e) {
    if ($(window).scrollTop() > 0) {
      $('.page-wrapper').addClass('scrolled');
    } else {
      $('.page-wrapper').removeClass('scrolled');
    }

    if (isScrollProgress) {
      scrollSpy();
    }
  }
});

function collapseFun() {
  $('.collapse').on({
    'show.bs.collapse hide.bs.collapse': e => {
      let $this = $(e.currentTarget);
      if ($this.is('[id]')) {
        $(`[data-target="#${$this.attr('id')}"]`).prop({ 'disabled': true });
      } else {
        if ($this.attr('class').indexOf('ins_') > -1) {
          let class_arry = $this.attr('class').split(' ');
          $.each(class_arry, (idx, val) => {
            if (val.indexOf('ins_') > -1) {
              $(`[data-target=".${val}"]`).prop({ 'disabled': true });
            }
          });
        }
      }
    },
    'shown.bs.collapse hidden.bs.collapse': e => {
      let $this = $(e.currentTarget);
      if ($this.is('[id]')) {
        $(`[data-target="#${$this.attr('id')}"]`).prop({ 'disabled': false });
      } else {
        if ($this.attr('class').indexOf('ins_') > -1) {
          let class_arry = $this.attr('class').split(' ');
          $.each(class_arry, (idx, val) => {
            if (val.indexOf('ins_') > -1) {
              $(`[data-target=".${val}"]`).prop({ 'disabled': false });
            }
          });
        }
      }
    },
    'shown.bs.collapse': e => {
      let $this = $(e.currentTarget);
      if ($this[0].tagName == 'INPUT' || $this[0].tagName == 'TEXTAREA') {
        $this.focus();
      } else {}
    }
  });
}

function toastMsg() {
  let url = window.location.href,
      urlArry = url.split("#"),
      urlCore = urlArry[0],
      urlHash = urlArry[1],
      urlQuery = url.split("?")[1],
      current_time = new Date().getTime(),
      inspect_info = JSON.parse(COMM.getCookie('inspect_info'));

  if (inspect_info) {
    var inspect_cookie_expireMin = Math.floor((current_time - inspect_info.inspect_info_time) / 1000 / 60);
  }

  let $body = $('body'),
      $myApps = $('.myApps:not(.exp_area)'),
      $expArea = $('.exp_area'),
      msg_alert_HTML = `<!-- toast -->
        <div class="toast_full_top">
          <div id="msg_alert" class="toast hide" data-autohide="false" data-delay="10000">
            <div class="toast-header">
              <i class="fas fa-exclamation-triangle red mRight10"></i>
              <span class="comment">
                진행중인 실험의 <span class="red">승인 후 점검 위원</span>으로 지정되었습니다. "<span class="underline">승인 후 점검</span>" 메뉴에서 <span class="red">승인 후 점검표</span>를 작성해 주세요.
              </span>
              <button type="button" class="close fs_same mLeft10 mTop3" data-dismiss="toast"><i class="fas fa-times red"></i></button>
              <a href="./inspection_list.html" class="btn btn-danger btn_xxs mLeft30">승인 후 점검표 작성<i class="fas fa-arrow-right mLeft10"></i></a>
            </div>
          </div>
        </div><!--//toast -->`;

  if (urlCore.indexOf('-') < 0 && urlCore.indexOf('payment') < 0 && (urlCore.indexOf('_list') > -1 || urlCore.indexOf('dashboard') > -1)) {
    $('.topbar').append(msg_alert_HTML);
    if (inspect_info && inspect_cookie_expireMin <= CONST.COOKIE_EXPIRES_MIN) {
      if ($body.hasClass('researcher') || $body.hasClass('officer') || $body.hasClass('committee') || $body.hasClass('chairman')) {
        $('#msg_alert').toast('show');
      }
    }
  }

  if (inspect_cookie_expireMin > CONST.COOKIE_EXPIRES_MIN) {
    let inspect_flag = true;
    inspect_info.inspect_flag = inspect_flag;
    inspect_info.inspect_info_time = new Date().getTime();
    COMM.setCookie('inspect_info', JSON.stringify(inspect_info));
    inspect_cookie_expireMin = 0;
  }
}