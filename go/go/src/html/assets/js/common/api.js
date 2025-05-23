const api_js = true;

const API = {};
(function (api, $) {
  function ajax(options) {
    var defaultOption = {
      url: '',
      type: CONST.API_TYPE.GET,
      enctype: '',
      dataType: 'json',
      async: true,
      headers: {
        'token': '',
        'Content-Type': 'application/json'
      },
      success: null,
      error: null,
      complete: null };
    var option = $.extend({}, defaultOption, options);

    var token = COMM.getCookie(CONST.COOKIE.IPSAP_TOKEN);
    if (token != null) option.headers.token = token;

    if (option.dataType == false) {
      option.dataType = undefined;
    }

    if (option.enctype == false) {
      option.enctype = undefined;
    }

    if (option.contentType == false) {
      option.headers['Content-Type'] = undefined;
    }

    if (!option.url) {
      return false;
    }
    option.url = CONST.API_PATH + option.url;

    if (option.data && option.contentType != false) {
      switch (option.type.toUpperCase()) {
        case CONST.API_TYPE.POST:
        case CONST.API_TYPE.PATCH:
        case CONST.API_TYPE.DELETE:
          option.data = JSON.stringify(option.data);
          break;
      }
    }

    let fncSuccess = option.success;
    option.success = null;

    let fncError;
    if (option.error) {
      fncError = option.error;
      option.error = null;
    } else {
      fncError = function (err) {
        if (err.readyState !== 0) {
          $('#loading').remove();
          if ($('.modal').length > 0) {
            $('.modal').modal('hide');
          }
          let msgObj = {
            title: '작업을 진행할 수 없습니다.',
            titleIcon: 'alert-triangle',
            apprTxt: '확인'
          };
          if (typeof err.responseJSON !== 'undefined') {
            msgObj.desc = err.responseJSON.em;
            if (err.responseJSON.e === 4) {
              msgObj.apprCallBack = function () {
                location.replace('/');
              };
            } else if (err.responseJSON.e === 5) {
              msgObj.apprCallBack = function () {
                COMM.clearCookie();
                location.replace('/');
              };
            } else if (err.responseJSON.e === 6 || 14) {
              msgObj.apprCallBack = function () {
                $('.modal').modal('hide');
              };
            }
          }
        }
      };
    }

    return $.ajax(option).done(function (results) {
      fncSuccess(results);
    }).fail(function (xhr, status, errorThrown) {
      switch (xhr.status) {
        case 404:
          fncError(xhr);
          break;
        default:
          fncError(xhr);
          break;
      }
    });
  }

  api.load = function (options) {
    return ajax(options);
  };

  return api;
})(API, $);

$(function () {
  let ajaxCnt = 0;
  $(document).ajaxStart(function () {
    ajaxCnt++;
    if ($('#loading').length == 0) {
      $('body').append(`
      <div id="loading" style="display: none;">
        <div class="spinner-border" style="position: absolute;width: 3rem;height: 3rem;top: 50%;left: 50%;z-index: 9999;" role="status">
          <span class="sr-only">Loading...</span>
        </div>
        <div style="position: fixed;top: 0;left: 0;z-index: 9998;width: 100vw;height: 100vh;background-color: #000;opacity: 0.1;"></div>
        </div>
      </div>
      `);
    }
    $('#loading').show();
  }).ajaxStop(function () {
    ajaxCnt--;

    if (ajaxCnt <= 0) {
      $('#loading').remove();
    }
  });
});