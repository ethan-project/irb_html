const ipsap_application_common_js = true;

const APP_NAVIGATION = {
   IACUC: {
      STEP_CNT: 11,
      STEP_INFO: [
         { TITLE: "일반정보", DOT: 2 },
         { TITLE: "동물실험 목적", DOT: 3 },
         { TITLE: "실험동물 정보", DOT: 2 },
         { TITLE: "실험물질 정보", DOT: 1 },
         { TITLE: "동물실험 방법", DOT: 2 },
         { TITLE: "고통등급", DOT: 1 },
         { TITLE: "고통경감 방안", DOT: 2 },
         { TITLE: "사육관리", DOT: 1 },
         { TITLE: "안락사", DOT: 1 },
         { TITLE: "연구책임자 준수사항", DOT: 1 },
         { TITLE: "자가 점검", DOT: 1 },
      ],
      PAGE_INFO: [
         { PAGE_ID: "PAGE_1_1", URL: "/html/researcher/application_new-01_1.html" },
         { PAGE_ID: "PAGE_1_2", URL: "/html/researcher/application_new-01_2.html" },
         { PAGE_ID: "PAGE_2_1", URL: "/html/researcher/application_new-02_1.html" },
         { PAGE_ID: "PAGE_2_2", URL: "/html/researcher/application_new-02_2.html" },
         { PAGE_ID: "PAGE_2_3", URL: "/html/researcher/application_new-02_3.html" },
         { PAGE_ID: "PAGE_3_1", URL: "/html/researcher/application_new-03_1.html" },
         { PAGE_ID: "PAGE_3_2", URL: "/html/researcher/application_new-03_2.html" },
         { PAGE_ID: "PAGE_4_1", URL: "/html/researcher/application_new-04_1.html" },
         { PAGE_ID: "PAGE_5_1", URL: "/html/researcher/application_new-05_1.html" },
         { PAGE_ID: "PAGE_5_2", URL: "/html/researcher/application_new-05_2.html" },
         { PAGE_ID: "PAGE_6_1", URL: "/html/researcher/application_new-06_1.html" },
         { PAGE_ID: "PAGE_7_1", URL: "/html/researcher/application_new-07_1.html" },
         { PAGE_ID: "PAGE_7_2", URL: "/html/researcher/application_new-07_2.html" },
         { PAGE_ID: "PAGE_8_1", URL: "/html/researcher/application_new-08_1.html" },
         { PAGE_ID: "PAGE_9_1", URL: "/html/researcher/application_new-09_1.html" },
         { PAGE_ID: "PAGE_10_1", URL: "/html/researcher/application_new-10_1.html" },
         { PAGE_ID: "PAGE_11_1", URL: "/html/researcher/application_new-11_1.html" },
      ],
   },
   TMP: 0,
};

const APP_IBC_NAVIGATION = {
   IBC: {
      STEP_CNT: 4,
      STEP_INFO: [
         { TITLE: "심의 신청서", DOT: 4 },
         { TITLE: "연구 계획서", DOT: 2 },
         { TITLE: "위해성 평가서", DOT: 5 },
         { TITLE: "연구책임자 서약서", DOT: 1 },
      ],
      PAGE_INFO: [
         { PAGE_ID: "PAGE_1_1", URL: "/html/researcher/IBC/application_new-IBC_01_1.html" },
         { PAGE_ID: "PAGE_1_2", URL: "/html/researcher/IBC/application_new-IBC_01_2.html" },
         { PAGE_ID: "PAGE_1_3", URL: "/html/researcher/IBC/application_new-IBC_01_3.html" },
         { PAGE_ID: "PAGE_1_4", URL: "/html/researcher/IBC/application_new-IBC_01_4.html" },
         { PAGE_ID: "PAGE_2_1", URL: "/html/researcher/IBC/application_new-IBC_02_1.html" },
         { PAGE_ID: "PAGE_2_2", URL: "/html/researcher/IBC/application_new-IBC_02_2.html" },
         { PAGE_ID: "PAGE_3_1", URL: "/html/researcher/IBC/application_new-IBC_03_1.html" },
         { PAGE_ID: "PAGE_3_2", URL: "/html/researcher/IBC/application_new-IBC_03_2.html" },
         { PAGE_ID: "PAGE_3_3", URL: "/html/researcher/IBC/application_new-IBC_03_3.html" },
         { PAGE_ID: "PAGE_3_4", URL: "/html/researcher/IBC/application_new-IBC_03_4.html" },
         { PAGE_ID: "PAGE_3_5", URL: "/html/researcher/IBC/application_new-IBC_03_5.html" },
         { PAGE_ID: "PAGE_4_1", URL: "/html/researcher/IBC/application_new-IBC_04_1.html" },
      ],
   },
   TMP: 0,
};

const APP_IRB_NAVIGATION = {
   IRB: {
      STEP_CNT: 5,
      STEP_INFO: [
         { TITLE: "심의신청서", DOT: 4 },
         { TITLE: "연구 상세요약서", DOT: 2 },
         { TITLE: "동의 취득", DOT: 1 },
         { TITLE: "제출 서류", DOT: 2 },
         { TITLE: "연구책임자 서약서", DOT: 1 },
      ],
      PAGE_INFO: [
         { PAGE_ID: "PAGE_1_1", URL: "/html/researcher/IRB/application_new-IRB_01_1.html" },
         { PAGE_ID: "PAGE_1_2", URL: "/html/researcher/IRB/application_new-IRB_01_2.html" },
         { PAGE_ID: "PAGE_1_3", URL: "/html/researcher/IRB/application_new-IRB_01_3.html" },
         // { PAGE_ID: "PAGE_1_3_1", URL: "/html/researcher/IRB/application_new-IRB_01_3_1.html" },
         { PAGE_ID: "PAGE_1_4", URL: "/html/researcher/IRB/application_new-IRB_01_4.html" },
         // { PAGE_ID: "PAGE_1_5", URL: "/html/researcher/IRB/application_new-IRB_01_5.html" },
         { PAGE_ID: "PAGE_2_1", URL: "/html/researcher/IRB/application_new-IRB_02_1.html" },
         { PAGE_ID: "PAGE_2_2", URL: "/html/researcher/IRB/application_new-IRB_02_2.html" },
         { PAGE_ID: "PAGE_3_1", URL: "/html/researcher/IRB/application_new-IRB_03_1.html" },
         { PAGE_ID: "PAGE_4_1", URL: "/html/researcher/IRB/application_new-IRB_04_1.html" },
         { PAGE_ID: "PAGE_4_2", URL: "/html/researcher/IRB/application_new-IRB_04_2.html" },
         { PAGE_ID: "PAGE_5_1", URL: "/html/researcher/IRB/application_new-IRB_05_1.html" },
      ],
   },
   TMP: 0,
};

(function (app_navigation, $) {
   app_navigation.getUrlFromId = function (pageID) {
      for (var i = 0; i < this.IACUC.PAGE_INFO.length; ++i) {
         if (this.IACUC.PAGE_INFO[i].PAGE_ID == pageID) return this.IACUC.PAGE_INFO[i].URL;
      }
      return "";
   };

   app_navigation.navigate = function (pageID) {
      var url = this.getUrlFromId(pageID);
      if (url == "") return;

      g_AppInfo.PageID = pageID;
      g_AppInfo.saveParamsAndNavigate(url);
   };

   return app_navigation;
})(APP_NAVIGATION, $);

(function (app_ibc_navigation, $) {
   app_ibc_navigation.getUrlFromId = function (pageID) {
      for (var i = 0; i < this.IBC.PAGE_INFO.length; ++i) {
         if (this.IBC.PAGE_INFO[i].PAGE_ID == pageID) return this.IBC.PAGE_INFO[i].URL;
      }
      return "";
   };

   app_ibc_navigation.navigate = function (pageID) {
      var url = this.getUrlFromId(pageID);
      if (url == "") return;

      g_AppInfo.PageID = pageID;
      g_AppInfo.saveParamsAndNavigate(url);
   };

   return app_ibc_navigation;
})(APP_IBC_NAVIGATION, $);

(function (app_irb_navigation, $) {
   app_irb_navigation.getUrlFromId = function (pageID) {
      for (var i = 0; i < this.IRB.PAGE_INFO.length; ++i) {
         if (this.IRB.PAGE_INFO[i].PAGE_ID == pageID) return this.IRB.PAGE_INFO[i].URL;
      }
      return "";
   };

   app_irb_navigation.navigate = function (pageID) {
      var url = this.getUrlFromId(pageID);
      if (url == "") return;

      g_AppInfo.PageID = pageID;
      g_AppInfo.saveParamsAndNavigate(url);
   };

   return app_irb_navigation;
})(APP_IRB_NAVIGATION, $);

class ApplicationInfo {
   constructor() {
      this.appSeq = 0;
      this.childAppSeq = 0;
      this.appObj = {};
      this.PageID = "PAGE_1_1";
   }

   initNew(newObj) {
      this.appSeq = 0;
      this.childAppSeq = 0;
      this.appObj = newObj;
      this.PageID = "PAGE_1_1";
      return true;
   }

   initWithAppObj(appObj) {
      if (appObj.application_seq == undefined) return false;

      this.appSeq = appObj.application_seq;
      this.appObj = appObj;
      this.PageID = "PAGE_1_1";

      if (appObj.application_type != IPSAP.APPLICATION_TYPE.NEW) {
         this.appSeq = appObj.parent_app_seq;
         this.childAppSeq = appObj.application_seq;
      }

      return true;
   }

   saveParamsAndNavigate(url) {
      IPSAP.setStor(IPSAP.APP_PARAM_NAME, this);
      window.location.href = url;
   }

   loadParams() {
      var param = IPSAP.getStor(IPSAP.APP_PARAM_NAME);

      if (param == null || param.PageID == undefined || param.appObj == undefined || param.childAppSeq == undefined || param.appSeq == undefined) {
         alert("로그인 후 화면 조회가 가능합니다. 메인페이지로 이동합니다.");

         window.location.href = IPSAP.DEFAULT_URL;
         return;
      }

      this.appSeq = param.appSeq;
      this.childAppSeq = param.childAppSeq;
      this.appObj = param.appObj;
      this.PageID = param.PageID;
   }
}

var g_AppInfo = new ApplicationInfo();
