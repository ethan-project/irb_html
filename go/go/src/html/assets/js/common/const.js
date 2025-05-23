const const_js = true;

const CONST = {
  API_PATH: 'https://www.ipsap.co.kr:7375/api/v1.0',

  API_TYPE: {
    GET: 'GET',
    POST: 'POST',
    PUT: 'PUT',
    DELETE: 'DELETE',
    PATCH: 'PATCH' },
  COOKIE: {
    IPSAP_TOKEN: 'ipsap_token',
    IPSAP_TMP_KEY: 'ipsap_tmp_key',
    IPSAP_USER_INFO: 'user_info',
    IPSAP_INST_INFO: 'institution_info' },
  COOKIE_EXPIRES_MIN: 30,
  API: {
    ADMIN: {
      LIST_USER: '/admin/user',
      CREATE: '/admin/user',
      PATCH_USER: '/admin/user',
      CREATE_BATCH: '/admin/user-batch',
      DELETE_USER: '/admin/user/${user_seq}',
      RESEND_MSG: '/admin/user/${user_seq}/resend-msg',
      PATCH_PASS: '/admin/user/${user_seq}/reset-password',
      EXPELL: '/admin/user/${user_seq}/withdraw' },
    AUTH: {
      LOGIN: '/auth/login',
      LOGIN_EMAIL: '/auth/email/login',
      CANCEL_WITHDRAW: '/auth/email/cancel-withdraw',
      FIND_ID: '/auth/find-id',
      FIND_PWD: '/auth/find-pwd' },
    APPLICATION: {
      LIST: '/application',
      LIST_INFO: '/application/${app_seq}/info',
      LIST_APPROVED: '/application-approved',
      LIST_POSSIBLE: '/application-possible',
      GETINFO: '/application/${app_seq}',
      PATCH: '/application/${app_seq}',
      POST: '/application/${app_seq}',
      DELETE: '/application/${app_seq}',
      DELETE2: '/app/${app_seq}',
      DUPPLICATE: '/application/${app_seq}/copy',
      COPY_RETRIAL: '/application/${app_seq}/retrial-copy',
      COPY_IACUC: '/application/${app_seq}/iacuc-copy',
      INSPECTOR: '/application/${app_seq}/inspector',
      PATCH_INSPECTOR: '/application/${app_seq}/inspector/${user_seq}' },
    APPLICATION_CHANGE: {
      ALLINFO: '/application/${app_seq}/change',
      ANIMALINFO: '/application/${app_seq}/change/animal',
      BEFORECHANGE_ENDDATE: '/application/${app_seq}/change/end-date',
      BEFORECHANGE_MEMBER: '/application/${app_seq}/change/member' },
    BILLING: {
      CREATE_KEY: '/billing-key',
      REQ_ASSIGN: '/billing-key/assign' },
    BOARD: {
      LIST: '/board',
      NEW: '/board',
      INFO: '/board/${board_seq}',
      DELETE: '/board/${board_seq}',
      PATCH: '/board/${board_seq}',
      INSTITUTIONLIST: '/institution-board' },
    COMMON: {
      DUP_CHECK: {
        EMAIL: '/common/dup-check/email',
        INST_CODE: '/common/dup-check/institution-code' }
    },
    DASHBOARD: {
      LIST: '/dashboard',
      DOWNLOAD: '/file-animal' },
    INSTITUTION: {
      LIST: '/institution',
      DETAIL: '/institution/${inst_seq}',
      WITHDRAW: '/institution/${inst_seq}',
      PATCH: '/institution/${inst_seq}',
      PATCH_PAYMENT: '/institution/${inst_seq}/payment',
      LIST_PAYMENT: '/institution/${inst_seq}/using-membership/purchased',
      COUNT_OFFICER: '/institution/${inst_seq}/admin-count',
      USER: '/institution/${inst_seq}/user',
      MEMBERSHIP: '/institution/${inst_seq}/using-membership/purchased',
      MOVE_INST: '/move-institution',
      MY_INST: '/my/institution',
      MY_OTHER_INST: '/my/other-institution' },
    MEMBERSHIP: {
      DETAIL: '/membership',
      PAYMENT_SETTING: '/institution/${inst_seq}/payment-setting',
      CHANGE: '/institution/${inst_seq}/product/${product_seq}',
      CANCEL: '/membership',
      REQ_REFUND_AMT: '/membership/cancel',
      PLAN: {
        NEW: '/membership/plan',
        LIST: '/membership/plan',
        DETAIL: '/membership/plan/${plan_seq}',
        PATCH: '/membership/plan/${plan_seq}',
        DELETE: '/membership/plan/${plan_seq}' },
      FREE: {
        LIST: '/membership/free',
        AVAILABLE_LIST: '/membership-free/institution',
        APPLY: '/membership/free',
        DELETE: '/membership/free/${free_seq}' }
    },
    ORDERS: {
      LIST: '/orders',
      PAY: '/orders',
      REQ_ASSIGN: '/orders/assign',
      DETAIL: '/orders/${order_seq}',
      CANCEL: '/orders/${order_seq}' },
    PLATFORM: {
      ADMIN_USER_INFO: '/platform/admin-user',
      ADMIN_USER_PATCH: '/platform/admin-user',
      INSTITUTION_PATCH: '/platform/institution',
      REGI_USER_INST_LIST: '/platform/other-institution/${user_seq}',
      INSTITUTION_STATUS: '/platform/payment/cancel/institution',
      REGI_USER_LIST: '/platform/user',
      REGI_USER_CREATE: '/platform/user',
      REGI_USER_PATCH: '/platform/user',
      REGI_USER_INFO: '/platform/user/${user_seq}',
      REGI_USER_DELETE: '/platform/user/${user_seq}',
      REGI_USER_RESEND_MSG: '/platform/user/${user_seq}/resend-msg',
      REGI_USER_CREATE_BATCH: '/platform/user-batch',
      REGI_USER_RESET_PWD: '/platform/user/${user_seq}/reset-password' },
    PRODUCTS: {
      NEW: '/products',
      LIST: '/products',
      DETAIL: '/products/${product_seq}',
      DELETE: '/products/${product_seq}',
      PATCH: '/products/${product_seq}' },
    REQUEST: {
      SERVICE: '/request/service',
      REGISTER: '/request/service',
      DETAIL: '/request/service/${reqsvc_seq}',
      PATCH: '/request/service/${reqsvc_seq}',
      HANDLE: '/request/service/${reqsvc_seq}/handle' },
    USER: {
      DETAIL: '/user/${user_seq}',
      DELETE: '/user/${user_seq}',
      PATCH: '/user/${user_seq}',
      PATCH_PASS: '/user/${user_seq}/change-password',
      WITHDRAW: '/user/${user_seq}/institution',
      REGISTER: '/user/${user_seq}/register' }
  },

  PASS_POINT: 80,
  PASS_ACCOUNT: 8,
  VISIBLE_AVAILABLE_ITEMS: 5 };