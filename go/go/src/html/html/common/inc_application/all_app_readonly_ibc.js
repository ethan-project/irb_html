"undefined" == typeof ipsap_common_func_js && document.write("<script src='/assets/js/ipsap/ipsap_common_func.js'></script>");
const A = "ibc_risk_pathogen",
   B = "ibc_risk_resistant_gene",
   C = "ibc_risk_toxin_gene",
   D = "ibc_risk_toxin_handling",
   targetA = ".ibc-3-1-A",
   targetB = ".ibc-3-1-B",
   targetC = ".ibc-3-1-C",
   targetD = ".ibc-3-1-D";
var g_templateA,
   g_templateB,
   g_templateC,
   g_templateD,
   g_min_value,
   g_inserted_data_cnt_A = 0,
   g_inserted_data_cnt_B = 0,
   g_inserted_data_cnt_C = 0,
   g_inserted_data_cnt_D = 0;
function insertDataA(e) {
   g_inserted_data_cnt_A++;
   var t = g_templateA.html().replaceAll("03-1-A", `03-1-A-${g_inserted_data_cnt_A}`);
   t.replaceAll("input_03-1-A", `input_03-1-A-${g_inserted_data_cnt_A}`);
   var a,
      i = $(t);
   if ("" != e && 1 == (a = JSON.parse(e)).v) {
      let e = i.children("div").eq(1);
      "" != a.d[1] && e.find(".data").eq(0).text(a.d[1]), "" != a.d[2] && e.find(".data").eq(1).text(a.d[2]);
      for (let t = 0; t < 4; t++)
         try {
            a.d[3][t] > 0 && e.find("input:checkbox").eq(t).prop("checked", !0);
         } catch (e) {}
      for (let t = 0; t < 4; t++)
         try {
            a.d[4][t] > 0 &&
               e
                  .find("input:checkbox")
                  .eq(4 + t)
                  .prop("checked", !0);
         } catch (e) {}
      e.find("input:radio[value=" + a.d[5] + "]").prop("checked", !0),
         "" != a.d[6] && e.find(".data").eq(2).text(a.d[6]),
         "" != a.d[7] && e.find(".data").eq(3).text(a.d[7]),
         "" != a.d[8] && e.find(".data").eq(4).text(a.d[8]),
         e.find(`input:radio[name='radio_03-1-A-${g_inserted_data_cnt_A}_5']:radio[value='${a.d[5]}']`).prop("checked", !0),
         e.find(`input:radio[name='radio_03-1-A-${g_inserted_data_cnt_A}_9']:radio[value='${a.d[9].R}']`).prop("checked", !0),
         e.find(`input:radio[name='radio_03-1-A-${g_inserted_data_cnt_A}_10']:radio[value='${a.d[10].R}']`).prop("checked", !0),
         "" != a.d[9].T1 && 1 == a.d[9].R && e.find(".data").eq(5).text(a.d[9].T1),
         "" != a.d[9].T2 && 4 == a.d[9].R && e.find(".data").eq(6).text(a.d[9].T1),
         "" != a.d[10].T1 && 3 == a.d[10].R && e.find(".data").eq(7).text(a.d[10].T1),
         e.find("pre").eq(0).text(a.d[11]),
         e.find("pre").eq(1).text(a.d[12]),
         a.d[13] > 0 && e.find("input:checkbox").eq(8).prop("checked", !0);
      for (let t = 0; t < 4; t++)
         try {
            a.d[14][t] > 0 &&
               e
                  .find("input:checkbox")
                  .eq(9 + t)
                  .prop("checked", !0);
         } catch (e) {}
      a.d[15] > 0 && e.find("input:checkbox").eq(13).prop("checked", !0);
      for (let t = 0; t < 3; t++)
         try {
            a.d[16][t] > 0 &&
               e
                  .find("input:checkbox")
                  .eq(14 + t)
                  .prop("checked", !0);
         } catch (e) {}
   }
   ($(`.${A}`).append(i), "" != e) &&
      1 == (a = JSON.parse(e)).v &&
      (1 == a.d[9].R && $(`#radio_03-1-A-${g_inserted_data_cnt_A}_9-1`).siblings().addClass("show"),
      4 == a.d[9].R && $(`#radio_03-1-A-${g_inserted_data_cnt_A}_9-4`).siblings().addClass("show"),
      1 != a.d[9].R && $(`#radio_03-1-A-${g_inserted_data_cnt_A}_9-1`).siblings().removeClass("show"),
      3 == a.d[10].R && $(`#radio_03-1-A-${g_inserted_data_cnt_A}_10-3`).siblings().addClass("show"),
      a.d[13] > 0 && ($(`#input_03-1-A-${g_inserted_data_cnt_A}_13_group`).removeClass("show"), $(`#input_03-1-A-${g_inserted_data_cnt_A}_13_group`).addClass("collapse")),
      a.d[15] > 0 && ($(`#input_03-1-A-${g_inserted_data_cnt_A}_14_group`).removeClass("show"), $(`#input_03-1-A-${g_inserted_data_cnt_A}_14_group`).addClass("collapse")));
}
function insertDataB(e) {
   g_inserted_data_cnt_B++;
   var t = g_templateB.html().replaceAll("03-1-B", `03-1-B-${g_inserted_data_cnt_B}`);
   t.replaceAll("input_03-1-B", `input_03-1-B-${g_inserted_data_cnt_B}`);
   var a,
      i = $(t);
   if ("" != e)
      if (1 == (a = JSON.parse(e)).v) {
         let e = i.children("div").eq(1);
         "" != a.d[1] && e.find(".data").eq(0).text(a.d[1]),
            a.d[2].C > 0 && e.find("input:checkbox").eq(0).prop("checked", !0),
            e.find(".data").eq(1).text(a.d[2].T1),
            a.d[3].C > 0 && e.find("input:checkbox").eq(1).prop("checked", !0),
            e.find(".data").eq(2).text(a.d[3].T1),
            a.d[4].C > 0 && e.find("input:checkbox").eq(2).prop("checked", !0),
            e.find(".data").eq(3).text(a.d[4].T1),
            a.d[5].R > 0 && (e.find(`input:radio[name='radio_03-1-B-${g_inserted_data_cnt_B}_3']:radio[value='${a.d[5].R}']`).prop("checked", !0), e.find(".data").eq(4).text(a.d[5].T1)),
            a.d[6].R > 0 &&
               (e.find(`input:radio[name='radio_03-1-B-${g_inserted_data_cnt_B}_4']:radio[value='${a.d[6].R}']`).prop("checked", !0),
               e.find(".data").eq(5).text(a.d[6].T1),
               e.find(".data").eq(6).text(a.d[6].T2)),
            a.d[7].R > 0 && (e.find(`input:radio[name='radio_03-1-B-${g_inserted_data_cnt_B}_5']:radio[value='${a.d[7].R}']`).prop("checked", !0), e.find(".data").eq(7).text(a.d[7].T1)),
            e.find("pre").eq(0).text(a.d[8]),
            e.find("pre").eq(1).text(a.d[9]);
      } else alert("Data version Error!!");
   ($(`.${B}`).append(i), "" != e) &&
      1 == (a = JSON.parse(e)).v &&
      (a.d[2].C > 0 && $(`#check_03-1-B-${g_inserted_data_cnt_B}_2-1`).siblings().eq(1).css("display", "block"),
      a.d[3].C > 0 && $(`#check_03-1-B-${g_inserted_data_cnt_B}_2-2`).siblings().eq(1).css("display", "block"),
      a.d[4].C > 0 && $(`#check_03-1-B-${g_inserted_data_cnt_B}_2-3`).siblings().eq(1).css("display", "block"),
      4 == a.d[5].R && $(`#radio_03-1-B-${g_inserted_data_cnt_B}_3-4`).siblings().eq(1).css("display", "block"),
      4 == a.d[6].R && $(`#radio_03-1-B-${g_inserted_data_cnt_B}_4-4`).siblings().addClass("show"),
      1 == a.d[6].R && $(`#radio_03-1-B-${g_inserted_data_cnt_B}_4-1`).siblings().removeClass("hidden"),
      1 != a.d[6].R && $(`#radio_03-1-B-${g_inserted_data_cnt_B}_4-1`).siblings().removeClass("show"),
      3 == a.d[7].R && $(`#radio_03-1-B-${g_inserted_data_cnt_B}_5-3`).siblings().addClass("show"));
}
function insertDataC(e) {
   g_inserted_data_cnt_C++;
   var t = g_templateC.html().replaceAll("03-1-C", `03-1-C-${g_inserted_data_cnt_C}`);
   t.replaceAll("input_03-1-C", `input_03-1-C-${g_inserted_data_cnt_C}`);
   var a,
      i = $(t);
   if ("" != e && 1 == (a = JSON.parse(e)).v) {
      let e = i.children("div").eq(1);
      "" != a.d[1] && e.find(".data").eq(0).text(a.d[1]),
         a.d[2].C1 > 0 && e.find("input:checkbox").eq(0).prop("checked", !0),
         e.find(".data").eq(1).text(a.d[2].T1),
         a.d[2].C2 > 0 && e.find("input:checkbox").eq(1).prop("checked", !0),
         e.find(".data").eq(2).text(a.d[2].T2),
         a.d[2].C3 > 0 && e.find("input:checkbox").eq(2).prop("checked", !0),
         e.find(".data").eq(3).text(a.d[2].T3),
         e.find(`input:radio[name='radio_03-1-C-${g_inserted_data_cnt_C}_3']:radio[value='${a.d[3].R}']`).prop("checked", !0),
         e.find(".data").eq(4).text(a.d[3].T1),
         "" != a.d[4] && e.find(".data").eq(5).text(a.d[4]),
         a.d[5].R > 0 &&
            (e.find(`input:radio[name='radio_03-1-C-${g_inserted_data_cnt_C}_5']:radio[value='${a.d[5].R}']`).prop("checked", !0),
            e.find(".data").eq(6).text(a.d[5].T1),
            e.find(".data").eq(7).text(a.d[5].T2)),
         a.d[6].R > 0 && (e.find(`input:radio[name='radio_03-1-C-${g_inserted_data_cnt_C}_6']:radio[value='${a.d[6].R}']`).prop("checked", !0), e.find(".data").eq(8).text(a.d[6].T1)),
         e.find("pre").eq(0).text(a.d[7]),
         e.find("pre").eq(1).text(a.d[8]);
   }
   ($(`.${C}`).append(i), "" != e) &&
      1 == (a = JSON.parse(e)).v &&
      (a.d[2].C1 > 0 && $(`#check_03-1-C-${g_inserted_data_cnt_C}_2-1`).siblings().eq(1).css("display", "block"),
      a.d[2].C2 > 0 && $(`#check_03-1-C-${g_inserted_data_cnt_C}_2-2`).siblings().eq(1).css("display", "block"),
      a.d[2].C3 > 0 && $(`#check_03-1-C-${g_inserted_data_cnt_C}_2-3`).siblings().eq(1).css("display", "block"),
      3 == a.d[3].R && $(`#radio_03-1-C-${g_inserted_data_cnt_C}_3-3`).siblings().eq(1).css("display", "block"),
      4 == a.d[5].R && $(`#radio_03-1-C-${g_inserted_data_cnt_C}_5-4`).siblings().addClass("show"),
      1 != a.d[5].R && $(`#radio_03-1-C-${g_inserted_data_cnt_C}_5-1`).siblings().removeClass("show"),
      1 == a.d[5].R && $(`#radio_03-1-C-${g_inserted_data_cnt_C}_5-1`).siblings().removeClass("hidden"),
      3 == a.d[6].R && $(`#radio_03-1-C-${g_inserted_data_cnt_C}_6-3`).siblings().addClass("show"));
}
function insertDataD(e) {
   g_inserted_data_cnt_D++;
   var t = g_templateD.html().replaceAll("03-1-D", `03-1-D-${g_inserted_data_cnt_D}`);
   t.replaceAll("input_03-1-D", `input_03-1-D-${g_inserted_data_cnt_D}`);
   var a,
      i = $(t);
   if ("" != e && 1 == (a = JSON.parse(e)).v) {
      let e = i.children("div").eq(1);
      "" != a.d[1] && e.find(".data").eq(0).text(a.d[1]),
         "" != a.d[2] && e.find(".data").eq(1).text(a.d[2]),
         "" != a.d[3] && e.find(".data").eq(2).text(a.d[3]),
         "" != a.d[4] && e.find(".data").eq(3).text(a.d[4]),
         a.d[5].R > 0 &&
            (e.find(`input:radio[name='radio_03-1-D-${g_inserted_data_cnt_D}_5']:radio[value='${a.d[5].R}']`).prop("checked", !0),
            e.find(".data").eq(4).text(a.d[5].T1),
            e.find(".data").eq(5).text(a.d[5].T2)),
         a.d[6].R > 0 && (e.find(`input:radio[name='radio_03-1-D-${g_inserted_data_cnt_D}_6']:radio[value='${a.d[6].R}']`).prop("checked", !0), e.find(".data").eq(6).text(a.d[6].T1)),
         e.find("pre").eq(0).text(a.d[7]),
         e.find("pre").eq(1).text(a.d[8]),
         a.d[9] > 0 && e.find("input:checkbox").eq(0).prop("checked", !0);
      for (let t = 0; t < 3; t++)
         try {
            a.d[10][t] > 0 &&
               e
                  .find("input:checkbox")
                  .eq(1 + t)
                  .prop("checked", !0);
         } catch (e) {}
   }
   ($(`.${D}`).append(i), "" != e) &&
      1 == (a = JSON.parse(e)).v &&
      (4 == a.d[5].R && $(`#radio_03-1-D-${g_inserted_data_cnt_D}_5-4`).siblings().addClass("show"),
      1 != a.d[5].R && $(`#radio_03-1-D-${g_inserted_data_cnt_D}_5-1`).siblings().removeClass("show"),
      3 == a.d[6].R && $(`#radio_03-1-D-${g_inserted_data_cnt_D}_6-3`).siblings().addClass("show"),
      a.d[9] > 0 && ($(`#input_03-1-D-${g_inserted_data_cnt_D}_9_group`).removeClass("show"), $(`#input_03-1-D-${g_inserted_data_cnt_D}_9_group`).addClass("collapse")));
}
function remakeApplicationInfoReadOnly(e, t, a) {
        console.log("e, a, t", e, a, t)

   var _ = "ibc_all";
   for (let t = 0; t < e.length; ++t) _ = _ + "," + e[t];
   let n = { "filter.query_items": _ };
   if (null != a && g_AppInfo.childAppSeq > 0 && a.length > 0) {
      var d = a[0];
      for (let e = 1; e < a.length; ++e) d = d + "," + a[e];
      (n["filter.child_app_seq"] = g_AppInfo.childAppSeq), (n["filter.child_items"] = d);
   }
   if (isInArray(e, "supplement")) {
      for (const [e, t] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
        console.log("e, t", e, t)
         let a = "app_supplement_" + e,
            i = "app_supplement_text_" + e;
         (html = `\n      <div class="col-sm-9 mTop20">\n      <div class="checkbox">\n        <input type="checkbox" id="${a}" value="1" data-toggle="collapse" data-target='#${i}'>\n        <label for="${a}" class="blue_deep">${t}</label>\n      </div>\n    </div>\n    <div class="col-sm-9 collapse" id="${i}">\n      <textarea class="form-control red_border_shallow input_review" rows="2" placeholder="${"구체적인 보완 요청 사항을 입력해 주세요"}">${""}</textarea>\n    </div>`),
            $("#" + a).replaceWith(html);
      }
   } else for (const [e, t] of Object.entries(IPSAP.IBC_REVIEW_TAG)) $("#app_supplement_" + e).remove();
   (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
      g_AppItemParser.load(n, function (a) {
         !(function (t) {
            if (!isInArray(e, "supplement")) return;
            var a = t.getItemData("supplement");
            for (const [e, t] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
               let t = a.dataObj.saved_data.data[e];
               if (null == t || "" == t) continue;
               let i = "app_supplement_" + e,
                  _ = "app_supplement_text_" + e;
               0 == t.substring(0, 1)
                  ? ($("#" + i).prop("checked", !1),
                    $("#" + i)
                       .parent()
                       .removeClass("checkbox-danger"),
                    $("#" + i)
                       .siblings()
                       .addClass("blue_deep"),
                    $("#" + _).addClass("collapse"),
                    $("#" + _).removeClass("show"))
                  : $("#" + _)
                       .children("textarea")
                       .val(t.substring(2));
            }
         })(a);
         var _ = "general_title";
         (n = a.getItemData(_)).applyTextMapValueReadOnly(_ + "_name_ko", "name_ko"), n.applyTextMapValueReadOnly(_ + "_name_en", "name_en");
         {
            var n = a.getItemData("application_info");
            let e = Number(n.getStringValue("approved_dttm"));
            0 == e ? $("#approved_date").text("IBC 위원회 승인일") : ($("#approved_date").text(getDttm(e).dt), $(".task_approved_date").text(getDttm(e).dt));
            let t = n.getStringValue("application_no");
            "" != t && ($(".register_num").text(t), $(".register_num2").text(t));
         }
         {
            _ = "general_end_date";
            let e = (n = a.getItemData(_, !1)).applyTextValueReadOnly();
            $(".task_period").text(e), 0 == g_AppInfo.appObj.approved_dttm ? $(".task_days").text("00") : $(".task_days").text(getDateDiffWith10(e, getDttm(g_AppInfo.appObj.approved_dttm).dt));
         }
         _ = "ibc_general_experiment_cnt";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_general_experiment_degree";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         (_ = "ibc_general_fund_org"), (n = a.getItemData(_));
         (l = n.applyMainCheckValueInCheckBoxReadOnly("ibc_general_fund_org-main_check")) &&
            ($("#ibc_general_fund_org-main_check").attr("onclick", "return false;"), $("#input_01-3_group").addClass("collapse").removeClass("show"));
         _ = "ibc_general_fund_org_name";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_general_fund_conflict";
         null != (n = a.getItemData(_)) && n.makeHtmlRadioTypeReadOnly();
         (_ = "ibc_general_fund_start_date"), (n = a.getItemData(_));
         $("#ibc_general_fund_start_date").text(n.applyTextValueReadOnly());
         (_ = "ibc_general_fund_end_date"), (n = a.getItemData(_));
         $("#ibc_general_fund_end_date").text(n.applyTextValueReadOnly()),
            (function (e) {
               var t = "general_director",
                  a = e.getItemData(t);
               if (null == a.dataObj.saved_data) return;
               a.initMembersIBC();
               let i = a.dataObj.saved_data.data[0];
               (i.info.tmp_phoneno = i.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
                  (i.info.tmp_edu_date = i.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3"));
               let _ = "";
               if ("" != i.edu_course) {
                  let e = JSON.parse(i.edu_course);
                  for (key in e) _ += `\n        <div class="flexMid">\n          <span class="min_w120">${key}</span>\n          <span class="mLeft10">${e[key]}</span>\n        </div>`;
               }
               $("#general_director_name").text(i.info.name),
                  $("#general_director_email").text(i.info.email),
                  $("#general_director_dept").text(i.info.dept_str),
                  $("#general_director_edu_date").text(i.info.tmp_edu_date),
                  $("#general_director_position").text(i.info.position_str),
                  $("#general_director_edu_instition").text(i.info.edu_institution_str),
                  $("#general_director_major_field").text(i.info.major_field_str),
                  $("#general_director_edu_course_num").text(i.info.edu_course_num),
                  $("#general_director_phoneno").text(i.info.tmp_phoneno),
                  $("#general_director_edu_course").append(_);
            })(a),
            (function (e) {
               var t = "general_expt",
                  a = e.getItemData(t);
               if (null == a.dataObj.saved_data) return;
               a.initMembersIBC(),
                  (function () {
                     var e = "general_expt",
                        t = (g_AppItemParser.getItemData(e), `<tbody class="data" id="${e}">`);
                     let a = g_AppItemParser.managedItems[e].member;
                     $.each(a, function (a, i) {
                        (i.info.tmp_phoneno = i.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
                           (i.info.tmp_edu_date = i.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3"));
                        var _ = `${e}_${i.user_seq}_radio`;
                        let n = "";
                        if ("" != i.edu_course) {
                           let e = JSON.parse(i.edu_course);
                           for (a in e) n += `\n          <div class="flexMid">\n            <span class="min_w120">${a}</span>\n            <span class="mLeft10">${e[a]}</span>\n          </div>`;
                        }
                        t += `\n        <tr data-id="${i.user_seq}" item_name="${e}">\n          <td>\n            <div class="custom-radio">\n              <input type="radio" id="${_}" name="exp_manager" value="${i.user_seq}">\n            </div>\n          </td>\n          \x3c!-- <td>${i.info.user_type_str}</td> --\x3e\n          <td>${i.info.name}</td>\n          <td>${i.info.dept_str}</td>\n          <td>${i.info.position_str}</td>\n          <td>${i.info.major_field_str}</td>\n          <td>${i.info.tmp_phoneno}</td>\n          <td class="ellipsis max_w150" title="${i.info.email}">${i.info.email}</td>\n          <td>${n}</td>\n        </tr>\n        `;
                     }),
                        $("#general_expt").replaceWith(t),
                        $.each(a, function (t, a) {
                           var i = `${e}_${a.user_seq}_radio`;
                           a.animal_mng_flag > 0 ? ($("#" + i).prop("checked", !0), $("#" + i).prop("readonly", !0)) : $("#" + i).prop("disabled", !0);
                        }),
                        (t += "</tbody>");
                  })();
            })(a),
            mappingIbcFacility(a, !0);
         var d = "ibc_general_experiment",
            c = (n = a.getItemData(d)).getStringValue("0"),
            s = [];
         "" != c && (s = c.split(","));
         for (
            $.each(s, function (e, t) {
               let a = "check_01-3-" + (Number(t) + 1);
               $("#" + a)
                  .prop("checked", !0)
                  .prop("readonly", !0);
            }),
               i = 0;
            i < $("input:checkbox[id*=check_01-3]").length;
            i++
         )
            0 == $("#check_01-3-" + (i + 1)).is(":checked") && $("#check_01-3-" + (i + 1)).prop("disabled", !0);
         (_ = "ibc_general_animal_flag"), (n = a.getItemData(_));
         (l = n.applyMainCheckValueInCheckBoxReadOnly("ibc_general_animal_flag-main_check")) && $("#input_01-4_group").addClass("collapse").removeClass("show");
         _ = "ibc_general_animal_name";
         (n = a.getItemData(_)).applyTextValueReadOnly(_);
         _ = "ibc_general_animal_iacuc_num";
         (n = a.getItemData(_)).applyTextValueReadOnly(_);
         _ = "ibc_general_animal_iacuc_flag";
         null != (n = a.getItemData(_)) && (n.makeHtmlRadioTypeReadOnly(IPSAP.COL.EVEN), n.makeHtmlRadioTypeReadOnly(IPSAP.COL.ODD));
         _ = "ibc_general_animal_exp_flag";
         null != (n = a.getItemData(_)) && (n.makeHtmlRadioTypeReadOnly(IPSAP.COL.EVEN), n.makeHtmlRadioTypeReadOnly(IPSAP.COL.ODD));
         _ = "ibc_general_animal_purpose";
         (n = a.getItemData(_)).applyTextValueReadOnly(_);
         _ = "ibc_general_ref";
         (n = a.getItemData(_)).applyMultiFileValueReadOnly(_, "관련 자료", "mRight10 mBot0 w110"),
            $("#general_title_name_ko_2").text($("#general_title_name_ko").text()),
            $("#general_title_name_en_2").text($("#general_title_name_en").text());
         {
            let e = $("#general_director").clone();
            $("#general_director_2").empty(), $("#general_director_2").replaceWith(e);
         }
         (_ = "ibc_general_experiment"), (c = (n = a.getItemData(_)).getStringValue("0")), (s = []);
         "" != c && (s = c.split(","));
         if (s.length > 0) {
            ibcResetLeftNavi(s);
            let e = Math.min.apply(null, s);
            (g_min_value = e),
               e >= 0 && e < 4
                  ? $("#classification").text("국가 승인 실험")
                  : e >= 4 && e < 9
                  ? $("#classification").text("기관 승인 실험")
                  : e >= 9 && e < 11
                  ? ($("#classification").text("기관 신고 실험"),
                    $("#form_03").addClass("hidden"),
                    $("#app_divider_2").addClass("hidden"),
                    $("#form_03 input:checkbox[id*=app_supplement]").prop("checked", !1))
                  : ($("#form_02").addClass("hidden"),
                    $("#form_02 input:checkbox[id*=app_supplement]").prop("checked", !1),
                    $("#form_03").addClass("hidden"),
                    $("#form_03 input:checkbox[id*=app_supplement]").prop("checked", !1),
                    $("#app_divider_1").addClass("hidden"),
                    $("#app_divider_2").addClass("hidden"),
                    $("#classification").text("면제실험"));
         }
         a.deleteSaveTag(["ibc_general_experiment"]);
         _ = "ibc_plan_classification";
         (n = a.getItemData(_)).makeHtmlCheckListReadOnly(_);
         _ = "ibc_plan_purpose_perf_text";
         (n = a.getItemData(_)).applyTextValueReadOnly(_);
         _ = "ibc_plan_purpose_perf_file";
         (n = a.getItemData(_)).applyMultiFileValueReadOnly(_, "관련 자료", "mRight10 mBot0 w80");
         _ = "ibc_plan_content_range_text";
         (n = a.getItemData(_)).applyTextValueReadOnly(_);
         _ = "ibc_plan_content_range_file";
         (n = a.getItemData(_)).applyMultiFileValueReadOnly(_, "관련 자료", "mRight10 mBot0 w80");
         _ = "ibc_plan_method_text";
         (n = a.getItemData(_)).applyTextValueReadOnly(_);
         _ = "ibc_plan_method_file";
         (n = a.getItemData(_)).applyMultiFileValueReadOnly(_, "관련 자료", "mRight10 mBot0 w80");
         (_ = A + "_check"), (n = a.getItemData(_));
         if ((l = n.applyMainCheckValueInCheckBoxReadOnly(_ + "-main_check")))
            $("#" + _ + "-main_check")
               .parent()
               .parent()
               .removeClass("hidden");
         else {
            var r = (n = a.getItemData(A)).dataObj.saved_data.data;
            for (const [e, t] of Object.entries(r)) insertDataA(t);
            $(".ibc-3-1-A").removeClass("hidden");
         }
         (_ = B + "_check"), (n = a.getItemData(_));
         if ((l = n.applyMainCheckValueInCheckBoxReadOnly(_ + "-main_check")))
            $("#" + _ + "-main_check")
               .parent()
               .parent()
               .removeClass("hidden");
         else {
            r = (n = a.getItemData(B)).dataObj.saved_data.data;
            for (const [e, t] of Object.entries(r)) insertDataB(t);
            $(".ibc-3-1-B").removeClass("hidden");
         }
         (_ = C + "_check"), (n = a.getItemData(_));
         if ((l = n.applyMainCheckValueInCheckBoxReadOnly(_ + "-main_check")))
            $("#" + _ + "-main_check")
               .parent()
               .parent()
               .removeClass("hidden");
         else {
            r = (n = a.getItemData(C)).dataObj.saved_data.data;
            for (const [e, t] of Object.entries(r)) insertDataC(t);
            $(".ibc-3-1-C").removeClass("hidden");
         }
         (_ = D + "_check"), (n = a.getItemData(_));
         if ((l = n.applyMainCheckValueInCheckBoxReadOnly(_ + "-main_check")))
            $("#" + _ + "-main_check")
               .parent()
               .parent()
               .removeClass("hidden");
         else {
            r = (n = a.getItemData(D)).dataObj.saved_data.data;
            for (const [e, t] of Object.entries(r)) insertDataD(t);
            $(".ibc-3-1-D").removeClass("hidden");
         }
         (d = "ibc_risk_bio_matters_check"), (c = (n = a.getItemData(d)).getStringValue("0")), (s = new Object());
         "" != c && (s = c.split(","));
         $("input:checkbox[name=ibc_risk_bio_matters_check]").each(function (e) {
            isInArray(s, this.value) && (this.checked = !0);
         });
         _ = "ibc_risk_bio_mask_check";
         var l = (n = a.getItemData(_)).applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_mask_check-main_check"),
            p = "ibc_risk_bio_mask_input";
         a.getItemData(p).applyTextValueReadOnly(), l ? ($("#input_03-2_7-6_group").addClass("show"), $("#ibc_risk_bio_mask_input").css("display", "block")) : $(`#${p}`).val("");
         (_ = "ibc_risk_bio_etc_protection_check"),
            (l = (n = a.getItemData(_)).applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_etc_protection_check-main_check")),
            (p = "ibc_risk_bio_etc_protection_input");
         a.getItemData(p).applyTextValueReadOnly(), l ? ($("#input_03-2_7-7_group").addClass("show"), $("#ibc_risk_bio_etc_protection_input").css("display", "block")) : $(`#${p}`).val("");
         (_ = "ibc_risk_bio_safety_workstation_check"),
            (l = (n = a.getItemData(_)).applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_safety_workstation_check-main_check")),
            (p = "ibc_risk_bio_safety_workstation_input");
         a.getItemData(p).applyTextValueReadOnly(), l ? ($("#input_03-2_8-5_group").addClass("show"), $("#ibc_risk_bio_safety_workstation_input").css("display", "block")) : $(`#${p}`).val("");
         (_ = "ibc_risk_bio_etc_safety_check"), (l = (n = a.getItemData(_)).applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_etc_safety_check-main_check")), (p = "ibc_risk_bio_etc_safety_input");
         a.getItemData(p).applyTextValueReadOnly(), l ? ($("#input_03-2_8-7_group").addClass("show"), $("#ibc_risk_bio_etc_safety_input").css("display", "block")) : $(`#${p}`).val("");
         (d = "ibc_risk_bio_grade"), (c = (n = a.getItemData(d)).getStringValue("0")), (s = new Object());
         "" != c && (s = c.split(","));
         $("input:checkbox[name=ibc_risk_bio_grade]").each(function (e) {
            isInArray(s, this.value) && (this.checked = !0);
         });
         _ = "ibc_risk_bio_reason";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bios_disinf_sterlztn";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bio_inf_treatment";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bio_countermeasure";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bio_disposal";
         (n = a.getItemData(_)).applyTextValueReadOnly(), mappingIbcInfectionChance(a, !0), mappingIbcEmergencyNetwork(a, !0);
         (_ = "ibc_risk_bios_prvnt_check"), (l = (n = a.getItemData(_)).applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bios_prvnt_check-main_check")), (p = "ibc_risk_bios_prvnt_input");
         a.getItemData(p).applyTextValueReadOnly(), l ? ($("#input_03-2_5-1_group").addClass("show"), $("#ibc_risk_bios_prvnt_input").css("display", "block")) : $(`#${p}`).val("");
         (_ = "ibc_risk_bios_treatment_check"), (n = a.getItemData(_));
         (l = n.applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bios_treatment_check-main_check")) &&
            ($("#input_03-2_5-2_group").addClass("show"), $("#ibc_risk_bios_treatment_input").css("display", "block"));
         _ = "ibc_risk_bios_treatment_input";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bio_packing";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         (_ = "ibc_risk_bio_transport_check"), (n = a.getItemData(_));
         (l = n.applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_transport_check-main_check")) &&
            ($("#input_03-2_10-1_group").addClass("show"), $("#ibc_risk_bio_transport_input").css("display", "block"));
         _ = "ibc_risk_bio_transport_input";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bio_keeping";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bio_lmo";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         _ = "ibc_risk_bios_prvnt_list";
         (n = a.getItemData(_)).applyMultiFileValueReadOnly(_, "예방백신 명단 제출", "mRight10 mBot0 mTop10 w120");
         (_ = "ibc_risk_bio_transport_check"), (n = a.getItemData(_));
         (l = n.applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_transport_check-main_check")) &&
            ($("#input_03-2_10-1_group").addClass("show"), $("#ibc_risk_bio_transport_input").css("display", "block"));
         _ = "ibc_risk_bio_transport_input";
         (n = a.getItemData(_)).applyTextValueReadOnly();
         (_ = "ibc_risk_bio_storage_check"), (n = a.getItemData(_));
         (l = n.applyMainCheckValueInCheckBoxReadOnly("ibc_risk_bio_storage_check-main_check")) &&
            ($("#input_03-2_10-2_group").addClass("show"), $("#ibc_risk_bio_storage_input").css("display", "block"));
         _ = "ibc_risk_bio_storage_input";
         (n = a.getItemData(_)).applyTextValueReadOnly(),
            $("input[id^='check_03']").each(function () {
               this.checked ? ((this.disabled = !1), (this.readOnly = !0), $(this).attr("onclick", "return false")) : (this.disabled = !0);
            }),
            $("input[id^='radio_03']").each(function () {
               this.checked && ((this.disabled = !1), (this.readOnly = !0));
            }),
            $(".task_director").text(g_AppInfo.appObj.user_name);
         var o = getDttm(g_AppInfo.appObj.submit_dttm).dt.split("-"),
            g = o[0] + "년 " + o[1] + "월 " + o[2] + "일";
         0 == g_AppInfo.appObj.submit_dttm && (g = getTimeStrFromDateType(new Date()));
         $(".agree_date").text(g), null != t && t(a);
         $(".card").removeClass("hidden");
      });
}
function makeSupplementData() {
   var e = {};
   for (const [t, a] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
      let a = "app_supplement_text_" + t,
         i = $("#" + ("app_supplement_" + t)).is(":checked"),
         _ = $("#" + a)
            .children("textarea")
            .val();
      if ((i || (_ = ""), !i || "" != _)) {
         let a = "0|";
         i && (a = "1|"), (a += _), (e[t] = a);
      }
   }
   return { data: e };
}
$(document).ready(function () {
   $(".ibc-3-1-A").wrap(`<span class="${A}"></span>`),
      (g_templateA = $(`.${A}`).clone()),
      $(".ibc-3-1-A").empty(),
      $(".ibc-3-1-B").wrap(`<span class="${B}"></span>`),
      (g_templateB = $(`.${B}`).clone()),
      $(".ibc-3-1-B").empty(),
      $(".ibc-3-1-C").wrap(`<span class="${C}"></span>`),
      (g_templateC = $(`.${C}`).clone()),
      $(".ibc-3-1-C").empty(),
      $(".ibc-3-1-D").wrap(`<span class="${D}"></span>`),
      (g_templateD = $(`.${D}`).clone()),
      $(".ibc-3-1-D").empty();
});
