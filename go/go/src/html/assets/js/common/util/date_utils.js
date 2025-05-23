const date_utils_js = true;

moment.locale('ko', {
  weekdays: ["일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"],
  weekdaysShort: ["일", "월", "화", "수", "목", "금", "토"]
});

const MOMENT_BASE_FORMAT = 'YYYY-MM-DD HH:mm:ss';

const MOMENT_DATE_FORMAT = 'YYYY-MM-DD';

const MOMENT_TIME_FORMAT = 'HH:mm:ss';

const SERVER_DATE_FORMAT = 'YYYYMMDD';

function getByDateFormat(date1, date2) {

  function f(date) {
    return moment(date, SERVER_DATE_FORMAT).format(MOMENT_DATE_FORMAT);
  }

  if (typeof date2 === 'undefined' || date2 == null) {
    return f(date1);
  } else {
    return f(date1) + ' ~ ' + f(date2);
  }
}

function converPeriod(period) {
  if (period && period.indexOf(' ~ ') > -1) {
    let tmp = period.split(' ~ ');
    return {
      startDate: tmp[0].replace(/-/gi, ''),
      endDate: tmp[1].replace(/-/gi, '')
    };
  }
  return {
    startDate: undefined,
    endDate: undefined
  };
}

function convertServerDate(date) {
  return moment(date * 1000);
}

function getToDay() {
  let date = moment();

  return {
    dt: date.format(MOMENT_DATE_FORMAT),
    tm: date.format(MOMENT_TIME_FORMAT)
  };
}

function getDttm(dttm) {
  let rtn = {
    dt: '-',
    tm: ''
  };

  if (dttm == '0' || dttm == '') {
    return rtn;
  } else {
    let date = convertServerDate(dttm);
    rtn.dt = date.format(MOMENT_DATE_FORMAT), rtn.tm = date.format(MOMENT_TIME_FORMAT);
    return rtn;
  }
}

function getDateDiff(date1, date2) {
  if (date1.length == 8) {
    date1 = moment(date1, SERVER_DATE_FORMAT);
  } else {
    date1 = convertServerDate(date1).hours(0).minutes(0).seconds(0);
  }

  if (date2.length == 8) {
    date2 = moment(date2, SERVER_DATE_FORMAT);
  } else {
    date2 = convertServerDate(date2).hours(0).minutes(0).seconds(0);
  }

  return date1.diff(date2, 'days');
}

function getDateDiffWith10(date1, date2) {
  date1 = moment(date1, MOMENT_DATE_FORMAT);
  date2 = moment(date2, MOMENT_DATE_FORMAT);
  return date1.diff(date2, 'days');
}

function getDttmForHistoryList(date, current_svr_time) {

  if (date !== '0') {
    let dttm = getDttm(date);

    let diff = getDateDiff(current_svr_time, date);
    if (diff < 2) {
      let dateStr = null;
      if (diff === 0) {
        dateStr = '오늘';
      } else {
        dateStr = '어제';
      }
      dttm.dt = '<span class="numeric">' + dttm.dt + '</span><span class="human">' + dateStr + '</span>';
    }
    return dttm;
  }

  return {
    dt: '-',
    tm: ''
  };
}

function getTimeStrFromSec(secs) {
  let tail = "전";
  if (secs < 0) {
    secs *= -1;
    tail = "초과";
  }

  let day = parseInt(secs / (60 * 60 * 24));
  secs -= day * (60 * 60 * 24);
  let hour = parseInt(secs / (60 * 60));
  secs -= hour * (60 * 60);
  let min = parseInt(secs / 60);
  secs -= min * 60;

  let ret = ``;
  if (day > 0) ret += `${day}일 `;
  if (hour > 0) ret += `${hour}시간 `;
  if (min > 0) ret += `${min}분 `;
  if (secs > 0) ret += `${secs}초 `;

  ret += tail;
  return ret;
}

function getTimeStrFromDateType(date) {
  var year = date.getFullYear();
  var month = 1 + date.getMonth();
  month = month >= 10 ? month : '0' + month;
  var day = date.getDate();
  day = day >= 10 ? day : '0' + day;

  return year + '년 ' + month + '월 ' + day + '일';
}

function getYMDHMSFromDateType(date, hideSecs) {
  var s = leadingZeros(date.getFullYear(), 4) + '-' + leadingZeros(date.getMonth() + 1, 2) + '-' + leadingZeros(date.getDate(), 2) + ' ' + leadingZeros(date.getHours(), 2) + ':' + leadingZeros(date.getMinutes(), 2);

  if (hideSecs != true) s += ':' + leadingZeros(date.getSeconds(), 2);
  return s;
}

function getYMDHMSFromUnixtime(utime, hideSecs) {
  var date = new Date(utime * 1000);
  return getYMDHMSFromDateType(date, hideSecs);
}

function leadingZeros(n, digits) {
  var zero = '';
  n = n.toString();

  if (n.length < digits) {
    for (i = 0; i < digits - n.length; i++) zero += '0';
  }
  return zero + n;
}

function getDateByStr(dateStr) {
  let tmp = getByDateFormat(dateStr).split('-');
  return new Date(tmp[0], parseInt(tmp[1]) - 1, tmp[2]);
}

function getToDayStr() {
  let date = moment();

  return date.format('YYYY년 MM월 DD일');
}

function getDate(dttm) {
  let date = convertServerDate(dttm);
  return date.format('YYYY년 MM월 DD일');
}

function getYYYYMMDDHHMMSSfromDate(date) {
  var year = date.getFullYear().toString();

  var month = date.getMonth() + 1;
  month = month < 10 ? '0' + month.toString() : month.toString();

  var day = date.getDate();
  day = day < 10 ? '0' + day.toString() : day.toString();

  var hour = date.getHours();
  hour = hour < 10 ? '0' + hour.toString() : hour.toString();

  var minites = date.getMinutes();
  minites = minites < 10 ? '0' + minites.toString() : minites.toString();

  var seconds = date.getSeconds();
  seconds = seconds < 10 ? '0' + seconds.toString() : seconds.toString();

  return year + month + day + hour + minites + seconds;
}