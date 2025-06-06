function cbFuncAfterMapping(e) {
   resetLeftStepNaviIcon(),
      $("input[type=checkbox]").on({
         change: function (e) {
            resetLeftStepNaviIcon(),
               $(this).is(":checked")
                  ? ($(this).parent().addClass("checkbox-danger"), $(this).siblings().removeClass("blue_deep"))
                  : ($(this).parent().removeClass("checkbox-danger"), $(this).siblings().addClass("blue_deep"), $(this).parent().parent().next().find(".input_review").val(""));
         },
      }),
      $(".input_review").on({
         keyup: function (e) {
            resetLeftStepNaviIcon();
         },
      });
   var t = e.getItemData("ibc_general_experiment").getStringValue("0"),
      s = [];
   "" != t && (s = t.split(",")), Math.min.apply(null, s) <= 8 ? $("#btn_finish_review_fast").hide() : $("#btn_finish_review").hide();
}
function resetLeftStepNaviIcon() {
   let e = 0,
      t = !0,
      s = !0,
      i = !1,
      n = !1;
   for (const [a, p] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
      let r = "app_supplement_" + a,
         o = "app_supplement_text_" + a,
         l = p.split("-")[0];
      if (e != l) {
         if (e > 0) {
            let s = $(`.process_content[data-process-num=${e}]`);
            s.removeClass("completed pending supplement"), t ? (i ? s.addClass("supplement") : s.addClass("completed")) : s.addClass("pending");
         }
         (e = l), (t = !0), (i = !1);
      }
      t &&
         $("#" + r).is(":checked") &&
         "" ==
            $("#" + o)
               .children("textarea")
               .val() &&
         ((t = !1), (s = !1)),
         $("#" + r).is(":checked") &&
            "" !=
               $("#" + o)
                  .children("textarea")
                  .val() &&
            ((i = !0), (n = !0));
   }
   if (e > 0) {
      let s = $(`.process_content[data-process-num=${e}]`);
      s.removeClass("completed pending supplement"), t ? (i ? s.addClass("supplement") : s.addClass("completed")) : s.addClass("pending");
   }
   $("#btn_request").removeClass("disabled"),
      $("#btn_finish_review").removeClass("disabled"),
      $("#btn_finish_review_fast").removeClass("disabled"),
      (s && n) || $("#btn_request").addClass("disabled"),
      (s && !n) || ($("#btn_finish_review").addClass("disabled"), $("#btn_finish_review_fast").addClass("disabled"));
}
function onReviewRequest() {
   (g_AppItemParser.targetSavedItemNames = []), g_AppItemParser.addMoreSaveTag("supplement", makeSupplementData()), appItemSubmit(onCompleteReviewRequest, IPSAP.APP_SUBMIT_TYPE_PARAM.SUPPLEMENT);
}
function onCompleteReviewRequest(e, t, s) {
   e ? (window.location.href = "../reviewOffice_list.html") : appSubmitHandleError("행정 보완 요청을 실패했습니다.", t);
}
function onFinishReview() {
   (g_AppItemParser.targetSavedItemNames = []), g_AppItemParser.addMoreSaveTag("supplement", makeSupplementData()), appItemSubmit(onCompleteFinishReview, IPSAP.APP_SUBMIT_TYPE_PARAM.CHECKING);
}
function onClickFastFinish() {
   (g_AppItemParser.targetSavedItemNames = []),
      g_AppItemParser.addMoreSaveTag("supplement", makeSupplementData()),
      appItemSubmit(onCompleteFinishReview, IPSAP.APP_SUBMIT_TYPE_PARAM.CHECKING1_FINISH_FAST);
}
function onCompleteFinishReview(e, t, s) {
   e
      ? s == IPSAP.APP_SUBMIT_TYPE_PARAM.CHECKING
         ? (window.location.href = "./reviewOffice_list-IBC_setup.html")
         : (window.location.href = "../reviewConfirm_list.html")
      : appSubmitHandleError("행정 검토 완료에 실패했습니다.", t);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   document.write("<script src='/html/researcher/IBC/common_application_new-IBC.js'></script>"),
   document.write("<script src='/html/researcher/sub_inc/3_2_ibc_risk_emergency_network.js'></script>"),
   $(function () {
      loadApplicationParams();
      $(".page-contents");
      $("#all_app_r").load("/html/common/inc_application/all_app_readonly_ibc.html", function () {
         remakeApplicationInfoReadOnly(["supplement"], cbFuncAfterMapping);
      });
   });
