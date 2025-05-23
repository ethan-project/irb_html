
package model

import (
  "github.com/nleeper/goment"
	"github.com/gin-gonic/gin"
  "github.com/xuri/excelize/v2"
  "encoding/json"
  "database/sql"
  "ipsap/common"
  "strings"
  "fmt"
  "log"
  "os"
)

type AppAnimal struct {
  Application         *Application
  Datas               []map[string]interface{}   //  []string
}

func (ins *AppAnimal)Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

  sql := ""
  if (item_name == "animal_sum") {
    // 파이널이 있는지 체크
    sql2 := fmt.Sprintf(`SELECT COUNT(application_seq) AS cnt FROM t_application_animal WHERE application_seq = %d  AND item_name = 'animal_type_final'`, ins.Application.Application_seq)
    row2 := common.DB_fetch_one(sql2, nil)
    item_name2 := "animal_type"
    if 0 != common.ToUint(row2["cnt"])  {
      item_name2 = "animal_type_final"
    }
    sql = fmt.Sprintf(`
            SELECT animal_code, SUM(male_cnt) AS male_cnt, SUM(female_cnt) AS female_cnt,
                  mb_grade, breeding_place, strain,
                  week_age, age_unit, weight_unit, size, size_unit,
                  weight_gram, supplier_type, supplier_name,
                  lmo_flag, ibc_num, genetic_type, lmo_type,
                  animal_code_str, mb_grade_str, breeding_place_str, IF(mb_grade_str <> '-', group_concat(mb_grade_str), '-') AS group_mb_grade_str
             FROM t_application_animal
            WHERE application_seq = %d
              AND item_name = '%s'
            GROUP BY animal_code, animal_code_str`, ins.Application.Application_seq, item_name2)
  } else {
    sql = fmt.Sprintf(`
            SELECT animal_code, male_cnt, female_cnt,
                   mb_grade, breeding_place, strain,
                   week_age, age_unit, weight_unit, size, size_unit,
                   weight_gram, supplier_type, supplier_name,
                   lmo_flag, ibc_num, genetic_type, lmo_type,
                   animal_code_str, mb_grade_str, breeding_place_str
              FROM t_application_animal
             WHERE application_seq = %d
               AND item_name = '%s'
             ORDER BY view_order`, ins.Application.Application_seq, item_name)
  }

  filter := func(row map[string]interface{}) {
    code := Code {}

    code.Type = "animal_code"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId(row[code.Type + "_str"])

    code.Type = "mb_grade"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId(row[code.Type + "_str"])

    code.Type = "breeding_place"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId(row[code.Type + "_str"])

    code.Type = "supplier_type"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId()

    code.Type = "lmo_type"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId()

    code.Type = "age_unit"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId()

    code.Type = "weight_unit"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId()

    code.Type = "size_unit"
    code.Id = common.ToUint(row[code.Type])
    row[code.Type + "_str"] = code.GetCodeStrFromTypeAndId()
  }

  ins.Datas = common.DB_fetch_all(sql, filter)

  return true
}

func (ins *AppAnimal)GetJsonData() (ret interface{}) {
  ret = ins.Datas
  return
}

func (ins *AppAnimal)UpdateItem(tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  dataArr := data1.([]interface{})
  sql := `DELETE FROM t_application_animal
           WHERE application_seq = ?
             AND item_name = ?`
  _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
  if err != nil {
    log.Println(err)
    return
  }

  for idx, data2 := range dataArr {
    data := data2.(map[string]interface{})

    code := Code {}

    code.Type = "animal_code"
    code.Id = common.ToUint(data[code.Type]);
    data[code.Type + "_str"] = code.GetCodeStrFromTypeAndId(data[code.Type + "_str"])

    code.Type = "mb_grade"
    code.Id = common.ToUint(data[code.Type]);
    data[code.Type + "_str"] = code.GetCodeStrFromTypeAndId(data[code.Type + "_str"])

    code.Type = "breeding_place"
    code.Id = common.ToUint(data[code.Type]);
    data[code.Type + "_str"] = code.GetCodeStrFromTypeAndId(data[code.Type + "_str"])

    sql := `INSERT INTO t_application_animal(application_seq, item_name,          animal_code,  animal_code_str,
                                             male_cnt,        female_cnt,         mb_grade,     mb_grade_str,
                                             breeding_place,  breeding_place_str, strain,       week_age,
                                             age_unit,        weight_unit,        size,         size_unit,
                                             weight_gram,     supplier_type,      supplier_name,
                                             lmo_flag,        ibc_num,            genetic_type,
                                             lmo_type,        view_order)
            VALUES(?,?,?,?,
                   ?,?,?,?,
                   ?,?,?,?,
                   ?,?,?,?,
                   ?,?,?,
                   ?,?,?,
                   ?,?)`
    _, err := tx.Exec(sql,
                      ins.Application.Application_seq,  item_name,                  data["animal_code"],  data["animal_code_str"],
                      data["male_cnt"],                 data["female_cnt"],         data["mb_grade"],     data["mb_grade_str"],
                      data["breeding_place"],           data["breeding_place_str"], data["strain"],       data["week_age"],
                      data["age_unit"],                 data["weight_unit"],        data["size"],         data["size_unit"],
                      data["weight_gram"],              data["supplier_type"],      data["supplier_name"],
                      data["lmo_flag"],                 data["ibc_num"],            data["genetic_type"],
                      data["lmo_type"],                 idx)
    if err != nil {
      log.Println(err)
      err_msg = "data 값이 잘못 되었습니다."
      log.Println(data)
      return
    }
  }

  ret = true
  return
}

func (ins *AppAnimal)UpdateFinalAnimal(tx *sql.Tx) (ret bool, err_msg string) {
  app_etc := AppEtc{ Application : ins.Application }
  if !app_etc.Load("ca_regular_item")  {   //  변경 사항 로드
    err_msg = "변경신청 사항을 로드할 수 없습니다."
    return
  }

  data, exists := app_etc.Data["104"]   //  동물 종/수량 변경  아이템 번호 : 하드코딩 Fix Number
  if !exists {
    ret = true
    return
  }

  var jsonData []string

  if err := json.Unmarshal([]byte(data), &jsonData); err != nil {
    err_msg = fmt.Sprintf("데이터 오류입니다.(%v)", data);
    return
	}

  var jsonAnimal []interface{}
  if err := json.Unmarshal([]byte(jsonData[1]), &jsonAnimal); err != nil {
    err_msg = fmt.Sprintf("데이터 오류입니다.(%v)", jsonData[1]);
    return
	}

  sql := `DELETE FROM t_application_animal
          WHERE application_seq = ?
            AND item_name = ?`
  _, err := tx.Exec(sql, ins.Application.Data["parent_app_seq"],  "animal_type_final")
  if err != nil {
    log.Println(err)
    err_msg = "data 값이 잘못 되었습니다."
    log.Println(data)
    return
  }

  for idx, ani := range jsonAnimal {
    aniData := ani.([]interface{})
    sql := `INSERT INTO t_application_animal(application_seq, item_name,          animal_code,  animal_code_str,
                                             male_cnt,        female_cnt,         mb_grade,     mb_grade_str,
                                             breeding_place,  breeding_place_str, strain,       week_age,
                                             weight_gram,     supplier_type,      supplier_name,
                                             lmo_flag,        ibc_num,            genetic_type,
                                             lmo_type,        view_order)
            VALUES(?,?,?,?,
                   ?,?,?,?,
                   ?,?,?,?,
                   ?,?,?,
                   ?,?,?,
                   ?,?)`

     _, err := tx.Exec(sql,
                       ins.Application.Data["parent_app_seq"],  "animal_type_final", 0,  aniData[3],
                       aniData[4], aniData[5],  0, "-",
                       0, "-", "-", 0,
                       0, 0, "-",
                       1, "", "", 9, idx)
    if err != nil {
      log.Println(err)
      err_msg = "data 값이 잘못 되었습니다."
      log.Println(data)
      return
    }
  }

  ret = true
  return
}

// parent_app_seq
func (ins *AppAnimal)InsertBringAnimal(tx *sql.Tx) (ret bool, err_msg string) {
  ret = false;
  err_msg = "data 값이 잘못 되었습니다."
  sql := `SELECT *
            FROM t_application_animal
           WHERE application_seq = ?
             AND item_name = 'animal_type_bring'`
  rows := common.DB_Tx_fetch_all(tx, sql, nil, ins.Application.Data["application_seq"])

  for _, row := range rows {
    sql = `SELECT SUM(male_cnt) AS sum_male_cnt, SUM(female_cnt) AS sum_female_cnt
             FROM (SELECT male_cnt, female_cnt
                     FROM t_application_animal
                    WHERE application_seq = ?
                      AND item_name = 'animal_type_bring'
                      AND animal_code = ?
                      AND animal_code_str = ?
                      AND view_order = ?
                    UNION ALL
                   SELECT male_cnt, female_cnt
                     FROM t_application_animal
                    WHERE application_seq = ?
                      AND item_name = 'animal_type_bring'
                      AND animal_code = ?
                      AND animal_code_str = ?
                      AND view_order = ?) AS a`
    sum_row := common.DB_Tx_fetch_one(tx, sql, nil,
                                      ins.Application.Data["application_seq"],
                                      row["animal_code"],
                                      row["animal_code_str"],
                                      row["view_order"],
                                      ins.Application.Data["parent_app_seq"],
                                      row["animal_code"],
                                      row["animal_code_str"],
                                      row["view_order"])

    sql = `INSERT INTO t_application_animal(
          	application_seq, item_name, animal_code,
          	animal_code_str, male_cnt, female_cnt,
          	mb_grade, mb_grade_str, breeding_place,
          	breeding_place_str, strain, week_age,
            age_unit, weight_unit, size, size_unit,
            weight_gram, supplier_type, supplier_name,
            lmo_flag, ibc_num, genetic_type,
            lmo_type, view_order
          )
          VALUES(?,?,?,
              	 ?,?,?,
              	 ?,?,?,
              	 ?,?,?,
                 ?,?,?,?,
              	 ?,?,?,
              	 ?,?,?,
              	 ?,?)
          ON DUPLICATE KEY UPDATE animal_code = ?, animal_code_str = ?,
                                  male_cnt = ?, female_cnt = ?, supplier_type = ?,
                                  supplier_name = ?, genetic_type = ?`
    _, err := tx.Exec(sql,
                      ins.Application.Data["parent_app_seq"], row["item_name"], row["animal_code"],
                      row["animal_code_str"], sum_row["sum_male_cnt"], sum_row["sum_female_cnt"],
                      row["mb_grade"], row["mb_grade_str"], row["breeding_place"],
                      row["breeding_place_str"], row["strain"], row["week_age"],
                      row["age_unit"], row["weight_unit"], row["size"], row["size_unit"],
                      row["weight_gram"], row["supplier_type"], row["supplier_name"],
                      row["lmo_flag"], row["ibc_num"], row["genetic_type"],
                      row["lmo_type"], row["view_order"],
                      row["animal_code"], row["animal_code_str"],
                      sum_row["sum_male_cnt"], sum_row["sum_female_cnt"], row["supplier_type"],
                      row["supplier_name"], row["genetic_type"])
    if err != nil {
      log.Println(err)
      err_msg = "data 값이 잘못 되었습니다."
      return
    }
  }

  ret = true
  return
}

func AnimalDownload(c *gin.Context, institution_seq uint, reg_user_seq interface{}) {
  var columns []string
  for i := 'A'; i <= 'S'; i++ {
    columns = append(columns, fmt.Sprintf("%c", i))
  }

  categories := map[string]string{"A1"	:	"기관코드",
                                  "B1"	: "과제명",
                                  "C1"	: "고통등급",
                                  "D1"	: "Special Consideration",
                                  "E1"	: "책임연구자",
                                  "F1"	: "IACUC 접수번호",
                                  "G1"	: "실험게시일",
                                  "H1"	: "실험종료일",
                                  "I1"	: "동물실험 목적에 따른 분류Ⅰ(식약처 보고사항)",
                                  "J1"	: "동물실험 목적에 따른 분류Ⅱ(검역본부 보고사항)",
                                  "K1"	: "동물실험 계획수량",
                                  "L1" 	: "반입보고서 접수번호",
                                  "M1"	: "반입일자",
                                  "N1"	: "동물종류",
                                  "O1"	: "M",
                                  "P1"	: "F",
                                  "Q1"	: "계",
                                  "R1"	: "공급처구분",
                                  "S1"  : "공급처명"}

  moreCondition := ""
  if (common.ToStr(reg_user_seq) != "") {
    moreCondition = fmt.Sprintf(`AND reg_user_seq = %v`, reg_user_seq)
  }

  // (
  //  IF((SELECT COUNT(*)
  //        FROM t_application_animal
  //       WHERE application_seq = app.parent_app_seq
  //         AND animal_code_str = ani.animal_code_str
  //         AND item_name = 'animal_type_final') = 0,
  //      (SELECT (SUM(male_cnt) + SUM(female_cnt))
  //         FROM t_application_animal
  //        WHERE application_seq = app.parent_app_seq
  //          AND animal_code_str = ani.animal_code_str
  //          AND item_name = 'animal_type'),
  //      (SELECT (male_cnt + female_cnt)
  //         FROM t_application_animal
  //        WHERE application_seq = app.parent_app_seq
  //          AND animal_code_str = ani.animal_code_str
  //          AND item_name = 'animal_type_final'
  //          ORDER BY view_order Desc
  //          LIMIT 1))
  // )

  sql := fmt.Sprintf(`
            SELECT	istt.institution_code AS A,
                    app.name_ko AS B,
                    (
                      SELECT CASE
                             WHEN select_ids = 1
                             THEN 'A'
                             WHEN select_ids = 2
                             THEN 'B'
                             WHEN select_ids = 3
                             THEN 'C'
                             WHEN select_ids = 4
                             THEN 'D'
                             WHEN select_ids = 5
                             THEN 'E'
                             END
                        FROM t_application_select
                       WHERE item_name = 'pain_grage'
                         AND application_seq = app.parent_app_seq
                    ) AS C,
                    user.name AS E,
                    (
                      SELECT
                        app2.application_no
                      FROM
                        t_application AS app2
                      WHERE
                        app2.application_seq = app.parent_app_seq
                    ) AS F,
                    (
                      SELECT
                        FROM_UNIXTIME(app2.approved_dttm, '%%Y-%%m-%%d')
                      FROM
                        t_application AS app2
                      WHERE
                        app2.application_seq = app.parent_app_seq
                    ) AS G,
                    (
                      SELECT
                        etc.contents
                      FROM
                        t_application_etc etc
                      WHERE
                        etc.item_name = 'general_end_date'
                        AND etc.application_seq = app.parent_app_seq
                    ) AS H,
                    (
                      SELECT
                        group_concat(value)
                      FROM
                        t_item_select
                      WHERE
                        item_name = 'purpose_1'
                        AND id REGEXP (
                          SELECT
                            REPLACE(
                              GROUP_CONCAT(select_ids),
                              ',',
                              '|'
                            )
                          FROM  t_application_select
                          WHERE item_name = 'purpose_1'
                            AND application_seq = app.parent_app_seq
                        )
                    ) AS I,
                    app.parent_app_seq AS seq,
                    '' AS K,
                    application_no AS L,
                    ani.genetic_type AS M,
                    ani.animal_code_str AS N,
                    ani.male_cnt AS O,
                    ani.female_cnt AS P,
                    (ani.male_cnt + ani.female_cnt) AS Q,
                    (
                      SELECT
                        value
                      FROM
                        t_code
                      WHERE
                        type = 'supplier_type'
                        AND id = ani.supplier_type
                    ) AS R,
                    ani.supplier_name AS S
                  FROM
                    t_application AS app
                    LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
                    LEFT OUTER JOIN t_institution istt ON (
                      app.institution_seq = istt.institution_seq
                    )
                    LEFT OUTER JOIN t_application_animal ani ON (
                      app.application_seq = ani.application_seq
                    )
                  WHERE
                    1 = 1
                    AND application_type = 4
                    AND application_step = 4
                    AND application_result IN (15, 16)
                    AND istt.institution_seq = %v
                    %v #moreCondition
                    `, institution_seq, moreCondition)

  filter := func(row map[string]interface{}) {
    sql2 := `SELECT group_concat(A) AS J
             FROM (SELECT CASE
                    WHEN item_name = 'purpose_2_basic'
                    THEN '기초 연구'
                    WHEN item_name = 'purpose_2_applied'
                    THEN '중개 및 응용 연구'
                    WHEN item_name = 'purpose_2_test'
                    THEN '법적인 요구사항을 만족하기 위한 규제시험(Regulatory Test)'
                    WHEN item_name = 'purpose_2_nature'
                    THEN '사람이나 동물의 건강이나 복지를 위한 자연환경 연구'
                    WHEN item_name = 'purpose_2_species'
                    THEN '종 보존을 위한 연구'
                    WHEN item_name = 'purpose_2_training'
                    THEN '교육이나 훈련'
                    WHEN item_name = 'purpose_2_forensic'
                    THEN '법의학 관련 연구'
                    WHEN item_name = 'purpose_2_gene'
                    THEN '유전자 변형 형질 동물 생산'
                    END AS A
               FROM t_application_main_check
               WHERE application_seq = ?
               AND checked = 1
               AND item_name LIKE 'purpose_2%'
               AND item_name NOT IN ('purpose_2_etc_check', 'purpose_2_test2')
               UNION
               SELECT value AS A
               FROM		t_item_select as item
               WHERE  item_name IN (SELECT item_name
                                      FROM t_application_main_check
                                     WHERE application_seq = ?
                                       AND checked = 1
                                       AND item_name LIKE 'purpose_2%'
                                       AND item_name NOT IN ('purpose_2_etc_check', 'purpose_2_test2'))
               AND		id REGEXP (SELECT REPLACE(GROUP_CONCAT(select_ids),',','|')
                                   FROM t_application_select
                                  WHERE item_name = item.item_name
                                    AND application_seq = ?
                                    AND select_ids <> '')
               AND value <> '기타'
               UNION
               SELECT	input AS A
               FROM		t_application_select_input
               WHERE item_name LIKE 'purpose_2%'
                 AND application_seq = ?
               UNION
               SELECT contents AS A
               FROM		t_application_etc
               WHERE	application_seq = ?
               AND		item_name = 'purpose_2_etc_input'
               AND 		contents <> '') AS T`
    row2 := common.DB_fetch_one(sql2, nil, row["seq"], row["seq"], row["seq"], row["seq"], row["seq"])
    row["J"] = row2["J"]

    sql3 := `SELECT select_ids
               FROM t_application_select
              WHERE application_seq = ?
                AND item_name = 'special_consideration'`
    row3 := common.DB_fetch_one(sql3, nil, row["seq"])
    ids := common.ToStr(row3["select_ids"]);
    if ids != "" {
      r := strings.NewReplacer("1", "Survival Surgery(SS)",
                               "2", "Multiple Survival Surgery(MSS)",
                               "3", "Food or Fluid Regulation(FFR)",
                               "4", "Prolonged Restraint(PR)",
                               "5", "Hazardous Agent Use(HAU)",
                               "6", "Non-Centralized Use(NCU)")
      ids = r.Replace(ids)
    }
    row["D"] = ids
  }

  rows := common.DB_fetch_all(sql, filter)

  // 신규 신청서 기준
  sql = fmt.Sprintf(`SELECT istt.institution_code AS A,
                app.name_ko AS B,
                (
                  SELECT CASE
                         WHEN select_ids = 1
                         THEN 'A'
                         WHEN select_ids = 2
                         THEN 'B'
                         WHEN select_ids = 3
                         THEN 'C'
                         WHEN select_ids = 4
                         THEN 'D'
                         WHEN select_ids = 5
                         THEN 'E'
                         END
                    FROM t_application_select
                   WHERE item_name = 'pain_grage'
                     AND application_seq = app.application_seq
                ) AS C,
                user.name AS E,
                app.application_no AS F,
                FROM_UNIXTIME(app.approved_dttm, '%%Y-%%m-%%d') AS G,
                (
                  SELECT
                    etc.contents
                  FROM
                    t_application_etc etc
                  WHERE
                    etc.item_name = 'general_end_date'
                    AND etc.application_seq = app.application_seq
                ) AS H,
                (
                  SELECT
                    group_concat(value)
                  FROM
                    t_item_select
                  WHERE
                    item_name = 'purpose_1'
                    AND id REGEXP (
                      SELECT
                        REPLACE(
                          GROUP_CONCAT(select_ids),
                          ',',
                          '|'
                        )
                      FROM  t_application_select
                      WHERE item_name = 'purpose_1'
                        AND application_seq = app.application_seq
                    )
                ) AS I,
                SUM(ani.male_cnt) + SUM(ani.female_cnt) AS K,
                '' AS L,
                '' AS M,
                ani.animal_code_str AS N,
                SUM(ani.male_cnt) AS O,
                SUM(ani.female_cnt) AS P,
                SUM(ani.male_cnt) + SUM(ani.female_cnt) AS Q,
                (
                  SELECT value
                    FROM t_code
                   WHERE type = 'supplier_type'
                     AND id = ani.supplier_type
                ) AS R,
                ani.supplier_name AS S,
                app.application_seq
              FROM
                t_application AS app
                LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
                LEFT OUTER JOIN t_institution istt ON (
                  app.institution_seq = istt.institution_seq
                )
                LEFT OUTER JOIN t_application_animal ani ON (
                  app.application_seq = ani.application_seq
                )
               WHERE 1 = 1
                 AND app.application_type = 1
                 AND app.application_step = 5
                 AND app.application_result <> 17
                 AND app.judge_type = 1
                 AND IF((SELECT COUNT(application_seq) AS cnt FROM t_application_animal WHERE application_seq = app.application_seq AND item_name = 'animal_type_final') > 0, ani.item_name = 'animal_type_final', ani.item_name = 'animal_type')
                 AND istt.institution_seq = %v
                 %v #moreCondition
                 GROUP BY ani.animal_code, ani.animal_code_str, ani.application_seq
                 ORDER BY app.approved_dttm`, institution_seq, moreCondition)
  filter2 := func(row map[string]interface{}) {
   sql2 := `SELECT group_concat(A) AS J
            FROM (SELECT CASE
                   WHEN item_name = 'purpose_2_basic'
                   THEN '기초 연구'
                   WHEN item_name = 'purpose_2_applied'
                   THEN '중개 및 응용 연구'
                   WHEN item_name = 'purpose_2_test'
                   THEN '법적인 요구사항을 만족하기 위한 규제시험(Regulatory Test)'
                   WHEN item_name = 'purpose_2_nature'
                   THEN '사람이나 동물의 건강이나 복지를 위한 자연환경 연구'
                   WHEN item_name = 'purpose_2_species'
                   THEN '종 보존을 위한 연구'
                   WHEN item_name = 'purpose_2_training'
                   THEN '교육이나 훈련'
                   WHEN item_name = 'purpose_2_forensic'
                   THEN '법의학 관련 연구'
                   WHEN item_name = 'purpose_2_gene'
                   THEN '유전자 변형 형질 동물 생산'
                   END AS A
              FROM t_application_main_check
              WHERE application_seq = ?
              AND checked = 1
              AND item_name LIKE 'purpose_2%'
              AND item_name NOT IN ('purpose_2_etc_check', 'purpose_2_test2')
              UNION
              SELECT value AS A
              FROM		t_item_select as item
              WHERE  item_name IN (SELECT item_name
                                     FROM t_application_main_check
                                    WHERE application_seq = ?
                                      AND checked = 1
                                      AND item_name LIKE 'purpose_2%'
                                      AND item_name NOT IN ('purpose_2_etc_check', 'purpose_2_test2'))
              AND		id REGEXP (SELECT REPLACE(GROUP_CONCAT(select_ids),',','|')
                                  FROM t_application_select
                                 WHERE item_name = item.item_name
                                   AND application_seq = ?
                                   AND select_ids <> '')
              AND value <> '기타'
              UNION
              SELECT	input AS A
              FROM		t_application_select_input
              WHERE item_name LIKE 'purpose_2%'
                AND application_seq = ?
              UNION
              SELECT contents AS A
              FROM		t_application_etc
              WHERE	application_seq = ?
              AND		item_name = 'purpose_2_etc_input'
              AND 		contents <> '') AS T`
   row2 := common.DB_fetch_one(sql2, nil, row["application_seq"], row["application_seq"], row["application_seq"], row["application_seq"], row["application_seq"])
   row["J"] = row2["J"]
   sql3 := `SELECT select_ids
              FROM t_application_select
             WHERE application_seq = ?
               AND item_name = 'special_consideration'`
   row3 := common.DB_fetch_one(sql3, nil, row["application_seq"])
   ids := common.ToStr(row3["select_ids"]);
   if ids != "" {
     r := strings.NewReplacer("1", "Survival Surgery(SS)",
                              "2", "Multiple Survival Surgery(MSS)",
                              "3", "Food or Fluid Regulation(FFR)",
                              "4", "Prolonged Restraint(PR)",
                              "5", "Hazardous Agent Use(HAU)",
                              "6", "Non-Centralized Use(NCU)")
     ids = r.Replace(ids)
   }
   row["D"] = ids
   row = nil
 }

  rows2 := common.DB_fetch_all(sql, filter2)
  rows3 := make([]map[string]interface{}, 0, 0)
  for _, v1 := range rows2 {
    rows3 = append(rows3, v1)
    for _, v2 := range rows {
      if common.ToUint(v2["seq"]) == common.ToUint(v1["application_seq"]) {
        rows3 = append(rows3, v2)
      }
    }
  }

  values := make(map[string]interface{})
  for i, row := range rows3 {
    for _ , column := range columns {
      data := column + common.ToStr(i+2);
      values[data] = row[column];
    }
  }

  f := excelize.NewFile()
  for k, v := range categories {
      f.SetCellValue("Sheet1", k, v)
  }

  for k, v := range values {
      f.SetCellValue("Sheet1", k, v)
  }

  // Save spreadsheet by the given path.
  now, _ := goment.New()
  dateTime := now.Format("YYYYMMDDHHmmss")

  fileName := fmt.Sprintf("data_%s.xlsx", dateTime)
  if err := f.SaveAs(fileName); err != nil {
     fmt.Println(err)
     return
  }

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.ms-excel")
	c.File(fileName)
  os.Remove(fileName)
}
