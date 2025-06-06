function cbFuncAfterMapping(e) {
   console.log("e",e)
   resetLeftStepNaviIcon(),
      $("input[type=checkbox]").on({
         change: function (e) {
            resetLeftStepNaviIcon(),
               $(this).is(":checked")
                  ? ($(this).parent().addClass("checkbox-danger"), $(this).siblings().removeClass("blue_deep"))
                  : ($(this).parent().removeClass("checkbox-danger"), $(this).siblings().addClass("blue_deep"), $(this).parent().parent().next().find(".input_review").val(""));
         },
      }),
      $("input[type=checkbox]").is(":checked")
         ? ($("input[type=checkbox]").parent().removeClass("checkbox-danger"),
           $("input[type=checkbox]").siblings().addClass("blue_deep"),
           $("input[type=checkbox]").parent().parent().next().find(".input_review").val(""))
         : ($("input[type=checkbox]").parent().addClass("checkbox-danger"), $("input[type=checkbox]").siblings().removeClass("blue_deep")),
      $(".input_review").on({
         keyup: function (e) {
            resetLeftStepNaviIcon();
         },
      }),
      $("#btn_request").click(function () {
         $("#modal_request").modal("show");
      }),
      $("#btn_finish_review").click(function () {
         $("#modal_finish_review").modal("show");
      });
}
function resetLeftStepNaviIcon() {
   let e = 0,
      t = !0,
      s = !0,
      n = !1,
      i = !1;
   for (const [a, p] of Object.entries(IPSAP.REVIEW_TAG)) {
      console.log("a,p ", a,p)
      let o = "app_supplement_" + a,
         c = "app_supplement_text_" + a,
         l = p.split("-")[0];
      if (e != l) {
         if (e > 0) {
            let s = $(`.process_content[data-process-num=${e}]`);
            s.removeClass("completed pending supplement"), t ? (n ? s.addClass("supplement") : s.addClass("completed")) : s.addClass("pending");
         }
         (e = l), (t = !0), (n = !1);
      }
      t &&
         $("#" + o).is(":checked") &&
         "" ==
            $("#" + c)
               .children("textarea")
               .val() &&
         ((t = !1), (s = !1)),
         $("#" + o).is(":checked") &&
            "" !=
               $("#" + c)
                  .children("textarea")
                  .val() &&
            ((n = !0), (i = !0));
   }
   if (e > 0) {
      let s = $(`.process_content[data-process-num=${e}]`);
      s.removeClass("completed pending supplement"), t ? (n ? s.addClass("supplement") : s.addClass("completed")) : s.addClass("pending");
   }
   $("#btn_request").removeClass("disabled"),
      $("#btn_finish_review").removeClass("disabled"),
      (s && i) || $("#btn_request").addClass("disabled"),
      (s && !i) || $("#btn_finish_review").addClass("disabled");
}
function onReviewRequest() {
   (g_AppItemParser.targetSavedItemNames = []), g_AppItemParser.addMoreSaveTag("supplement", makeSupplementData()), appItemSubmit(onCompleteReviewRequest, IPSAP.APP_SUBMIT_TYPE_PARAM.SUPPLEMENT);
}
function onCompleteReviewRequest(e, t, s) {
   e ? (window.location.href = "./reviewOffice_list.html") : appSubmitHandleError("행정 보완 요청을 실패했습니다.", t);
}
function onFinishReview() {
   (g_AppItemParser.targetSavedItemNames = []), g_AppItemParser.addMoreSaveTag("supplement", makeSupplementData()), appItemSubmit(onCompleteFinishReview, IPSAP.APP_SUBMIT_TYPE_PARAM.CHECKING);
}
function onCompleteFinishReview(e, t, s) {
   e ? (window.location.href = "./reviewOffice_list-setup.html") : appSubmitHandleError("행정 검토 완료를 실패했습니다.", t);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   document.write("<script src='../researcher/application_new-0.js'></script>"),
   $(function () {
      loadApplicationParams();
      $(".page-contents");
      $("#all_app_r").load("/html/common/inc_application/all_app_readonly.html", function () {
         remakeApplicationInfoReadOnly(["supplement"], cbFuncAfterMapping);
      });
   });
