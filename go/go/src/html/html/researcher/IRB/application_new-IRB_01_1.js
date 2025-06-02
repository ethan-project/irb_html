function dataMappingFunc(e) {
   0 == e.app_seq && e.addMoreSaveTag("application_info", { data: g_AppInfo.appObj });
   var a = "general_title";
   (t = e.getItemData(a)).applyTextMapValue(a + "_name_ko", "name_ko"), t.applyTextMapValue(a + "_name_en", "name_en");
   a = "general_end_date";
   (t = e.getItemData(a)).applyCalanderValue(a), setDate();
   a = "general_fund_org";
   (t = e.getItemData(a)).applyMainCheckValueInCheckBox("general_fund_org-main_check") && $("#input_01-3_group").addClass("collapse").removeClass("show");
   a = "general_fund_org_name";
   (t = e.getItemData(a)).applyTextValue();
   a = "general_fund_conflict";
   null != (t = e.getItemData(a)) && t.makeHtmlRadioType();
   a = "general_object";
   var t,
      n = (t = e.getItemData(a, "nosave")).getStringValue("0");
   irbResetLeftNavi(n), $(".card").removeClass("hidden");
}
function setDate() {
   let e = new Date();
   e.setDate(e.getDate() + 7);
   let a = moment(new Date()).format("YYYY-MM-DD");
   $("#general_end_date").val(a);
   let t = new Date();
   t.setDate(t.getDate() + 1095),
      $("#general_end_date").on("change", function () {
         let n = $("#general_end_date").val().split("-"),
            l = new Date(n[0], parseInt(n[1] - 1), n[2]);
         e.getTime() > l.getTime()
            ? (alert("오늘로부터 일주일 이후 날짜부터 선택이 가능합니다."), $("#general_end_date").val(a))
            : t.getTime() < l.getTime() && (alert("연구 기간은 신청일 기준 3년까지 가능합니다."), $("#general_end_date").val(a));
      });
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load({ "filter.query_items": "general_title, general_end_date, general_fund_org, general_fund_org_name, general_fund_conflict, general_object" }, dataMappingFunc);
   });
