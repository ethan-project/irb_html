function makeExptUserList() {
   var e = "general_expt",
      t = g_AppItemParser.getItemData(e),
      n = '<tbody id="general_expt">';
   let a = g_AppItemParser.managedItems[e].member;
   $.each(a, function (e, a) {
      (a.info.tmp_phoneno = a.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")), (a.info.tmp_edu_date = a.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3"));
      var r = `general_expt_${a.user_seq}_select_1`,
         o = `general_expt_${a.user_seq}_select_2`;
      (combo_html_1 = makeComboList(r, "<option selected disabled>구분 선택</option>", t.dataObj.codes, "exp_type_code")),
         (combo_html_2 = makeComboList(o, "<option selected disabled>경력</option>", t.dataObj.codes, "exp_year_code")),
         (n += `\n      <tr data-id="${a.user_seq}" item_name="general_expt">\n        <td>${combo_html_1}</td>\n        <td>${a.info.name}</td>\n        <td>${a.info.dept_str}</td>\n        <td>${a.info.position_str}</td>\n        <td>${a.info.tmp_phoneno}</td>\n        <td>${a.info.email}</td>\n        <td>${a.info.edu_course_num}</td>\n        <td>${a.info.edu_institution_str}</td>\n        <td>${combo_html_2}</td>\n        <td>\n          <a href="javascript:void(0);" class="btn btn-xs btn-outline-danger btn-staff-delete"><i class="far fa-trash-alt mRight5"></i>삭제</a>\n        </td>\n      </tr>`);
   }),
      (n +=
         '\n  <tr class="add_row">\n    <td colspan="11">\n      <a href="javascript:void(0);" onclick="openModelOtherStaff();" class="btn btn-xs btn-outline-primary" ><i class="fas fa-user-plus mRight5"></i>실험 수행자 선택 추가</a>\n    </td>\n  </tr>\n  </tbody>'),
      $("#general_expt").replaceWith(n),
      $.each(a, function (e, n) {
         var a = `general_expt_${n.user_seq}_select_1`,
            r = `general_expt_${n.user_seq}_select_2`;
         $("#" + a)
            .find(`option[value='${n.exp_type_code}'`)
            .attr("selected", !0),
            $("#" + r)
               .find(`option[value='${n.exp_year_code}'`)
               .attr("selected", !0),
            $("#" + a).on("change", function () {
               t.changeAttrForMember(n.user_seq, "exp_type_code", this.value);
            }),
            $("#" + r).on("change", function () {
               t.changeAttrForMember(n.user_seq, "exp_year_code", this.value);
            });
      });
}
function openModelOtherStaff() {
   makeOtherStaffList(), $("#modal_staff").modal("show");
}
function onStaffSelect(e) {
   $.each(g_other_userlist, function (t, n) {
      if (Number(n.user_seq) == e && !g_AppItemParser.hasMemberExists(n.user_seq)) return g_AppItemParser.insertMemberIRB("general_expt", n), makeExptUserList(), void $("#modal_staff").modal("hide");
   });
}
function makeOtherStaffList() {
   let e = '<tbody id="other-staff">';
   $.each(g_other_userlist, function (t, n) {
      var a = n.user_type.split(",");
      g_AppItemParser.hasMemberExists(n.user_seq) ||
         -1 == $.inArray("3", a) ||
         ((n.tmp_phoneno = n.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
         (e += `\n      <tr data-url="" class="staff-row" onclick="onStaffSelect(${n.user_seq})">\n        <td class="text-center">${n.name}</td>\n        <td>${n.dept_str}</td>\n        <td>${n.position_str}</td>\n        <td>${n.major_field_str}</td>\n        <td>${n.tmp_phoneno}</td>\n        <td>${n.email}</td>\n        <td>${n.edu_course_num}</td>\n        <td>${n.edu_institution_str}</td>\n      </tr>`));
   }),
      (e += "</tbody>"),
      $("#other-staff").replaceWith(e);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      function e(e) {
         !(function (e) {
            var t = "general_director",
               n = e.getItemData(t);
            if (null == n.dataObj.saved_data) return;
            (n.dataObj.saved_data.data[0].animal_mng_flag = 0), n.initMembers();
            let a = n.dataObj.saved_data.data[0];
            (a.info.tmp_phoneno = a.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
               (a.info.tmp_edu_date = a.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3")),
               $("#general_director_name").text(a.info.name),
               $("#general_director_email").text(a.info.email),
               $("#general_director_dept").text(a.info.dept_str),
               $("#general_director_edu_date").text(a.info.tmp_edu_date),
               $("#general_director_position").text(a.info.position_str),
               $("#general_director_edu_instition").text(a.info.edu_institution_str),
               $("#general_director_major_field").text(a.info.major_field_str),
               $("#general_director_edu_course_num").text(a.info.edu_course_num),
               $("#general_director_phoneno").text(a.info.tmp_phoneno);
            var r = "general_director_select";
            (combo_html = makeComboList(r, "<option selected disabled>경력 선택</option>", n.dataObj.codes, "exp_year_code")),
               $("#" + r).replaceWith(combo_html),
               $("#" + r)
                  .find(`option[value='${a.exp_year_code}'`)
                  .attr("selected", !0),
               $("#" + r).on("change", function () {
                  n.changeAttrForMember(a.user_seq, "exp_year_code", this.value);
               });
         })(e),
            (function (e) {
               var t = e.getItemData("general_expt");
               null != t.dataObj.saved_data && (t.initMembersIRB(), makeExptUserList());
            })(e);
         var t = e.getItemData("general_object", "nosave").getStringValue("0");
         irbResetLeftNavi(t), $(".card").removeClass("hidden");
      }
      loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load({ "filter.query_items": "general_director, general_expt, general_object" }, e),
         (function () {
            let e = CONST.API.INSTITUTION.USER;
            API.load({
               url: e.replace("${inst_seq}", g_AppInfo.appObj.institution_seq),
               data: { "filter.user_type": IPSAP.USER_TYPE.RESEARCHER },
               type: CONST.API_TYPE.GET,
               success: function (e) {
                  g_other_userlist = e;
               },
            });
         })(),
         $(document).on("click", ".btn-staff-delete", function () {
            if ($(this).closest("tr").find("input:radio").is(":checked")) return !1;
            var e = $(this).closest("tr").attr("data-id"),
               t = $(this).closest("tr").attr("item_name");
            g_AppItemParser.removeMembers(t, e), $(this).closest("tr").remove();
         });
   });
