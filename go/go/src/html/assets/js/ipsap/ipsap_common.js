const ipsap_common_js = true;

if (typeof ipsap_common_func_js == 'undefined') document.write("<script src='/assets/js/ipsap/ipsap_common_func.js'></script>");

const IPSAP = {
  DEMO_MODE: false,
  JUDGE_TYPE: {
    NONE: 0,
    IACUC: 1,
    IBC: 2,
    IRB: 3
  },
  APPLICATION_TYPE: {
    NEW: 1,
    CHANGE: 2,
    RENEW: 3,
    BRING: 4,
    CHECKLIST: 5,
    END: 6,
    MAINTANACE: 7,
    ADVERSE: 8,
    VIOLATION: 9,
    PROBLEM_OCCUR: 10 },
  APPLICATION_STEP: {
    WRITE: 0,
    CHECK: 1,
    PRO_JUDGE: 2,
    COMM_JUDGE: 3,
    FINAL: 4,
    PERFORMANCE: 5 },
  APPLICATION_RESULT: {
    TEMP: 0,
    SUPPLEMENT: 1,
    CHECKING: 2,
    CHECKING_2: 3,
    JUDGE_ING: 4,
    JUDGE_ING_2: 5,
    JUDGE_DELAY: 6,
    DECISION_ING: 8,
    REJECT: 9,
    REQUIRE_RETRY: 10,
    EXPERIMENT_ING_A: 11,
    EXPERIMENT_ING_AC: 12,
    EXPERIMENT_FINISH: 13,
    TASK_FINISH: 14,
    ACCEPT: 15,
    ACCEPT_AC: 16,
    DELETED: 17 },
  APP_LIST_TYPE: {
    ALL: 0,
    CHECK: 1,
    JUDGE_HISTROY: 3,
    JUDGE_FINISH: 4,
    FINAL_FINISH: 5,
    INSPECT_RECORD: 6 },
  APP_LIST_TYPE_PARAM: {
    ADMIN_ALL: "1",
    ADMIN_REVIEW_OFFICE: "2",
    ADMIN_REVIEW_CLOSE: "3",
    ADMIN_REVIEW_CONFIRM: "4",
    RESEARCHER_ALL: "5",
    CHAIRMAN_REVIEW_CONFIRM: "6",
    COMMITTEE_REVIEW: "7",
    CHAIRMAN_REVIEW_LIST: "8",
    COMMITTEE_REVIEW_LIST: "9",
    INSPECT_RECORD: "10",
    MY_INSPECTION: "11" },
  APPROVED_APP_LIST_TYPE_PARAM: {
    ADMIN_ALL: "1",
    ADMDIN_ING: "2",
    RESEARCHER_ALL: "3",
    CHAIRMAN_ING: "4" },
  APP_SUBMIT_TYPE_PARAM: {
    SUPPLEMENT: "1",
    CHECKING: "2",
    CHECKING_2: "3",
    JUDGE_PRO: "4",
    JUDGE_PRO_2: "5",
    JUDGE_NORMAL: "6",
    JUDGE_RESUME: "7",
    JUDGE_FINISH: "8",
    JUDGE_FINAL_A: "9",
    JUDGE_FINAL_AC: "10",
    JUDGE_FINAL_REJECT: "11",
    JUDGE_FINAL_REQUIRE_RETRY: "12",
    EXPER_FINISH: "13",
    TASK_FINISH: "14",
    RETRY_CHECKING: "15",

    CHILD_SUBMIT: "20",
    CHECKING1_FINISH_FAST: "21",

    JUMP_FINAL: "30" },
  APP_PROC_BTN_KIND: {
    CHECK: { url: ["./reviewOffice_list-review.html", "./reviewOffice_list-review_change.html", "./reviewOffice_list-review_renew.html", "./reviewOffice_list-review_bring.html", "./reviewOffice_list-review_inspection.html", "./reviewOffice_list-review_end.html"],
      text: "행정 검토" },
    JUDGE_SET: { url: ["./reviewOffice_list-setup.html", "./reviewOffice_list-setup.html", "./reviewOffice_list-setup.html"],
      text: "심사 설정" },
    JUDGE_FINISH: { url: ["./reviewClose_list-review.html"],
      text: "심사 종료" },
    JUDGE_FINISH2: { url: ["./reviewClose_list-review.html"],
      text: "심사 종료 설정" },
    FINAL: { url: ["./reviewConfirm_list-review.html", "./reviewConfirm_list-review_change.html", "./reviewConfirm_list-review_renew.html", "./reviewConfirm_list-review_bring.html", "./reviewConfirm_list-review_inspection.html", "./reviewConfirm_list-review_end.html"],
      text: "최종 심의" },
    PRO_JUDGE: { url: ["./review_list-review_IACUC_expert_1.html", "./review_list-review_IACUC_expert_1_change.html", "./review_list-review_IACUC_expert_1_renew.html"],
      text: "전문 심사" },
    NOR_JUDGE: { url: ["./review_list-review_IACUC_general_1.html", "./review_list-review_IACUC_general_1_change.html", "./review_list-review_IACUC_general_1_renew.html"],
      text: "일반 심사" } },
  IBC_APP_PROC_BTN_KIND: {
    CHECK: { url: ["./IBC/reviewOffice_list-IBC_review.html", "./IBC/reviewOffice_list-IBC_review_change.html"],
      text: "행정 검토" },
    JUDGE_SET: { url: ["./IBC/reviewOffice_list-IBC_setup.html", "./IBC/reviewOffice_list-IBC_setup.html"],
      text: "심사 설정" },
    JUDGE_FINISH: { url: ["./IBC/reviewClose_list-review.html"],
      text: "심사 종료" },
    JUDGE_FINISH2: { url: ["./IBC/reviewClose_list-review.html"],
      text: "심사 종료 설정" },
    FINAL: { url: ["./IBC/reviewConfirm_list-IBC_review.html", "./IBC/reviewConfirm_list-IBC_review_change.html"],
      text: "최종 심의" },
    PRO_JUDGE: { url: ["./review_list-review_IBC_expert_1.html", "./review_list-review_IBC_expert_1_change.html"],
      text: "전문 심사" },
    NOR_JUDGE: { url: ["./review_list-review_IBC_general_1.html", "./review_list-review_IBC_general_1_change.html"],
      text: "일반 심사" } },
  IRB_APP_PROC_BTN_KIND: {
    CHECK: {
      url: ["./IRB/reviewOffice_list-IRB_review.html", "./IRB/reviewOffice_list-IBC_review_change.html", "", "", "", "", ""],
      text: "행정 검토"
    },
    JUDGE_SET: {
      url: ["./IRB/reviewOffice_list-IRB_setup.html", "./IRB/reviewOffice_list-IRB_setup.html", "./IRB/reviewOffice_list-IRB_setup.html"],
      text: "심사 설정"
    },
    FINAL: {
      url: ["./IRB/reviewConfirm_list-IRB_review.html", "./IRB/reviewConfirm_list-IRB_review_change.html", "./IRB/reviewConfirm_list-IRB_review_renew.html"],
      text: "최종 심의"
    },
    PRO_JUDGE: {
      url: ["./review_list-review_IRB_expert_1.html", "./review_list-review_IRB_expert_1_change.html", "./review_list-review_IRB_expert_1_renew.html"],
      text: "전문 심사"
    },
    NOR_JUDGE: {
      url: ["./review_list-review_IRB_general_1.html", "./review_list-review_IRB_general_1_change.html", "./review_list-review_IRB_general_1_renew.html"],
      text: "일반 심사"
    } },
  COL: {
    ALL: 0,
    ODD: 1,
    EVEN: 2,
    HORI: 3 },
  APP_PARAM_NAME: "APP_PARAM",
  DEFAULT_URL: "/index.html",
  USER_TYPE: {
    ALL: 0,
    ADMIN_OFFICER: 1,
    CHAIRMAN: 2,
    RESEARCHER: 3,
    COMMITTEE: 4,
    OFFICER: 5 },
  USER_AUTH: {
    AUTH_NOMARL: 0,
    AUTH_INSTITUTION: 1,
    AUTH_PLATFORM: 9,
    AUTH_SYSTEM: 10 },
  REVIEW_TAG: {
    general_title: "1-1. 동물실험 과제명 - 보완 요청",
    general_date_cnt: "1-2. 연구기간 및 동물실험 횟수 - 보완 요청",
    general_fund_org: "1-3. 연구비 지원 기관 - 보완 요청",
    general_director: "1-4. 연구 책임자 - 보완 요청",
    general_expt: "1-5. 실험 수행자 - 보완 요청",
    general_ref: "1-6. 관련 자료 첨부 - 보완 요청",
    purpose_1: "2-1. 동물실험 목적에 따른 분류 Ⅰ - 보완 요청",
    purpose_2: "2-2. 동물실험 목적에 따른 분류 Ⅱ - 보완 요청",
    purpose_3: "2-3. 동물실험원칙(3Rs)에 다른 대안방법 모색 - 보완 요청",
    purpose_keyword: "2-4. 검색어 (Key Words) - 보완 요청",
    purpose_search_date: "2-5. 정보 검색일 - 보완 요청",
    purpose_alter_result: "2-6. 대안 방법 검토결과 - 보완 요청",
    animal_type: "3-1. 실험 동물의 종류 - 보완 요청",
    animal_species_reason: "3-2. 해당 동물 종(Species)과 계통(Strain)을 선택한 합리적인 이유 - 보완 요청",
    animal_cnt_reason: "3-3. 사용 동물 수에 대한 합리적인 근거 - 보완 요청",
    substance_dosage_flag: "4-1. 실험물질 투여 유무 - 보완 요청",
    animal_exp_summary: "5-1. 동물 실험의 개요 및 일정 - 보완 요청",
    animal_exp_form: "5-2. 동물 실험의 형태 - 보완 요청",
    animal_exp_restraint: "5-3. 보정법 - 보완 요청",
    animal_exp_identify: "5-4. 식별법 - 보완 요청",
    animal_exp_surgical: "5-5. 외과적 처치 - 보완 요청",
    animal_exp_export: "5-6. 동물 반출 - 보완 요청",
    animal_exp_sample: "5-7. 시료 채취 - 보완 요청",
    animal_exp_substance: "5-8. 실험 물질 - 보완 요청",
    pain_grage: "6-1. 고통 등급 분류 - 보완 요청",
    pain_d: "6-2. 고통등급 D (실험 동물의 고통 경감 방안) - 보완 요청",
    pain_e: "6-3. 고통등급 E (동물실험을 수행하는 사유 및 관리 방안) - 보완 요청",
    special_consideration: "6-4. 특별 고려사항(Special Consideration) – AAALAC PROGRAM : Animal Usage - 보완 요청",
    pain_relief_eval: "7-1. 고통 및 스트레스에 대한 평가 방법 - 보완 요청",
    pain_relief_end: "7-2. 인도적인 종료시점을 적용할 수 없는 사유 - 보완 요청",
    pain_relief_psych_medicine: "7-3. 마약, 향정신성 의약품 사용 유무 - 보완 요청",
    pain_relief_animal_medicine: "7-4. 동물 의약품 사용 유무 - 보완 요청",
    pain_relief_veterinary_mng: "7-5. 고통경감을 위한 수의학적 관리 - 보완 요청",
    breed_mng_condition: "8-1. 특별한 주거 및 사육 조건 - 보완 요청",
    breed_mng_environment: "8-2. 사육 환경 - 보완 요청",
    breed_mng_rich_tool: "8-3. 풍부화 도구 - 보완 요청",
    breed_mng_notrich_reason: "8-4. 풍부화 불가능 사유 - 보완 요청",
    euthanasia_method: "9-1. 안락사 방법 - 보완 요청",
    euthanasia_corpse: "9-2. 사체 처리 방법 - 보완 요청",
    euthanasia_share: "9-3. 실험동물 유래자원 공유 - 보완 요청",
    compliance_matters_agree: "10-1. 작업 환경 및 실험도구의 안전성 관리 - 보완 요청",
    compliance_matters_pledge: "10-2. 준수 사항 - 보완 요청"
  },
  IBC_REVIEW_TAG: {
    ibc_main_general: "1-1. 연구과제 기본 정보 - 보완 요청",
    ibc_general_researcher_fclty: "1-2. 연구원 및 연구시설 정보 - 보완 요청",
    ibc_general_experiment: "1-3. 실험 분류 - 보완 요청",
    ibc_general_animal_flag: "1-4. 실험동물 사용 유무 - 보완 요청",
    ibc_general_ref: "1-5. 기타 관련자료의 제출 - 보완 요청",
    ibc_plan_project: "2-1. 연구과제 기본 정보 - 보완 요청",
    ibc_plan_classification: "2-2. 실험 구분 - 보완 요청",
    ibc_plan_purpose_performance: "2-3. 연구목적 및 예상 성과 - 보완 요청",
    ibc_plan_content_range: "2-4. 연구내용 및 범위 - 보완 요청",
    ibc_plan_method: "2-5. 연구방법 - 보완 요청",
    ibc_risk_organisms_substance: "3-1. 취급 생물체 및 물질 정보 - 보완 요청",
    ibc_risk_bio: "3-2. 생물안전정보 - 보완 요청",
    ibc_director_pledge: "4-1. 연구 책임자 서약서 - 보완 요청"
  },
  EVALUATION_METHOD: {
    SCORE: 1,
    YN: 2,
    NONE: 3 },
  APP_MAIN_TAG_NAME: {
    "01-1": "1-1. 동물실험 과제명",
    "01-2": "1-2. 연구기간 및 동물실험 횟수",
    "01-3": "1-3. 연구비 지원 기관",
    "01-4": "1-4. 연구 책임자",
    "01-5": "1-5. 실험 수행자",
    "01-6": "1-6. 관련 자료 첨부",
    "02-1": "2-1. 동물실험 목적에 따른 분류 Ⅰ",
    "02-2": "2-2. 동물실험 목적에 따른 분류 Ⅱ",
    "02-3": "2-3. 동물실험원칙(3Rs)에 다른 대안방법 모색",
    "02-4": "2-4. 검색어 (Key Words)",
    "02-5": "2-5. 정보 검색일",
    "02-6": "2-6. 대안 방법 검토결과",
    "03-1": "3-1. 실험 동물의 종류",
    "03-2": "3-2. 해당 동물 종(Species)과 계통(Strain)을 선택한 합리적인 이유",
    "03-2": "3-3. 사용 동물 수에 대한 합리적인 근거",
    "04-1": "4-1. 실험물질 투여 유무",
    "05-1": "5-1. 동물 실험의 개요 및 일정",
    "05-2": "5-2. 동물 실험의 형태",
    "05-3": "5-3. 보정법",
    "05-4": "5-4. 식별법",
    "05-5": "5-5. 외과적 처치",
    "05-6": "5-6. 동물 반출",
    "05-7": "5-7. 시료 채취",
    "05-8": "5-8. 실험 물질",
    "06-1": "6-1. 고통 등급 분류",
    "06-2": "6-2. 고통등급 D (실험 동물의 고통 경감 방안)",
    "06-3": "6-3. 고통등급 E (동물실험을 수행하는 사유 및 관리 방안)",
    "07-1": "7-1. 고통 및 스트레스에 대한 평가 방법",
    "07-2": "7-2. 인도적인 종료시점을 적용할 수 없는 사유",
    "07-3": "7-3. 마약, 향정신성 의약품 사용 유무",
    "07-4": "7-4. 동물 의약품 사용 유무",
    "07-5": "7-5. 고통경감을 위한 수의학적 관리",
    "08-1": "8-1. 특별한 주거 및 사육 조건",
    "08-2": "8-2. 사육 환경",
    "08-3": "8-3. 풍부화 도구",
    "08-4": "8-4. 풍부화 불가능 사유",
    "09-1": "9-1. 안락사 방법",
    "09-2": "9-2. 사체 처리 방법",
    "09-3": "9-3. 실험동물 유래자원 공유",
    "10-1": "10-1. 작업 환경 및 실험도구의 안전성 관리",
    "10-2": "10-2. 준수 사항"
  },
  IBC_MAIN_TAGE_NAME: {
    "01-1": "1-1. 연구과제 기본 정보",
    "01-2": "1-2. 연구원 및 연구시설 정보",
    "01-3": "1-3. 실험 분류",
    "01-4": "1-4. 실험동물 사용 유무",
    "01-5": "1-5. 기타 관련자료의 제출",
    "02-1": "2-1. 연구과제 기본 정보",
    "02-2": "2-2. 실험 구분",
    "02-3": "2-3. 연구목적 및 예상 성과",
    "02-4": "2-4. 연구내용 및 범위",
    "02-5": "2-5. 연구방법",
    "03-1": "3-1. 취급 생물체 및 물질 정보",
    "03-2": "3-2. 생물안전정보"
  },
  IRB_MAIN_TAGE_NAME: {
    "01-1": "1-1. 연구과제 기본 정보",
    "01-2": "1-2. 연구원 정보",
    "01-3": "1-3. 연구 유형 및 심의 분류",
    "01-4": "1-4. 연구대상",
    "02-1": "2-1. 연구 목적",
    "02-2": "2-2. 배경 및 이론적 증거",
    "02-3": "2-3. 연구 방법",
    "02-4": "2-4. 관찰 및 검사 항목",
    "02-5": "2-5. 평가 기준 및 평가 방법",
    "03-1": "3-1. 동의취득 정보",
    "04-1": "4-1. 필수 제출 서류",
    "04-2": "4-2. 추가 제출 서류",
    "04-3": "4-3. 기타 제출 서류"
  },
  LIST_SIZE: {
    CASE1: 50,
    CASE2: 30
  }
};

(function (ipasp, $) {
  ipasp.setStor = function (name, val) {
    sessionStorage.setItem(name, JSON.stringify(val));
  };

  ipasp.removeStor = function (name) {
    sessionStorage.removeItem(name);
  };

  ipasp.clearStor = function (name, val) {
    sessionStorage.clear();
  };

  ipasp.getStor = function (name) {
    let item = sessionStorage.getItem(name);
    if (item != null) {
      return JSON.parse(item);
    }
    return item;
  };
})(IPSAP, $);