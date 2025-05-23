package api

import (
	"ipsap/common"
	"ipsap/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"

	// "strings"
	"fmt"
	// "log"
)

// @Tags Board
// @Summary 게시판 리스트
// @Description board_type = 1 공지사항
// @Description board_type = 2 자료실
// @Description board_type = 3 FAQ
// @Description qurey param board_type 이 없으면 전체
// @Description qurey param institution_seq가 없으면 공통게시판
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.board_type query string false "1"
// @Param filter.institution_seq query string false "1"
// @Router /board [get]
// @Success 200
func BoardList(c *gin.Context) {

	log.Println("c : ", c)
	tokenMap := common.Check_token(c)

	log.Println("tokenMap : ", tokenMap)

	if nil == tokenMap {
		return
	}

	seq_str := common.ToStr(c.Request.URL.Query().Get("filter.institution_seq"))
	institution_seq := uint(0)
	var err interface{}
	if "" != seq_str {
		institution_seq, err = cast.ToUintE(seq_str)
		if nil != err {
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
		}
		if 0 != institution_seq {
			tokenMap := common.Check_token(c)
			if nil == tokenMap {
				return
			}
			if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
				if common.ToUint(tokenMap["institution_seq"]) != institution_seq {
					common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_unauthorized)
					return
				}
			}
		}
	}

	boardType := common.ToStr(c.Request.URL.Query().Get("filter.board_type"))
	moreCondition := fmt.Sprintf(` AND board.institution_seq = %d`, institution_seq)
	if "" != boardType {
		cd := model.Code{}
		cd.Type = "board_type"
		cd.Id = common.ToUint(boardType)
		if !cd.CheckCodeError(c) {
			return
		}
		moreCondition += fmt.Sprintf(` AND board.board_type  = %v`, boardType)
	}

	board := model.Board{}
	board.Board_type = common.ToUint(boardType)
	rows := board.LoadList(moreCondition)
	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Board
// @Summary 등록 기관 공지사항 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /institution-board [get]
// @Success 200
func InstitutionBoardList(c *gin.Context) {
	board := model.Board{}
	board.Board_type = model.DEF_BOARD_TYPE_NOTICE
	moreCondition := fmt.Sprintf(`AND board.institution_seq > 0
																AND board.board_type  = %v`, model.DEF_BOARD_TYPE_NOTICE)
	rows := board.LoadList(moreCondition)
	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Board
// @Summary 게시판 상세 정보
// @Description 게시판 상세 정보
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param board_seq path uint true "board_seq"
// @Router /board/{board_seq} [get]
// @Success 200
func BoardInfo(c *gin.Context) {

	board_seq, succ, _, _ := getBoardInfoFromPath(c)
	if !succ {
		return
	}

	board := model.Board{}
	board.Board_seq = common.ToUint(board_seq)
	if !board.Load(true) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	common.FinishApi(c, common.Api_status_ok, board.Data)
}

// @Tags Board
// @Summary	게시판 등록
// @Description 게시판 등록
// @Description institution_seq = 0 인거는 공통!
// @Description board_type = 1 공지사항
// @Description board_type = 2 자료실
// @Description board_type = 3 FAQ
// @Description view_order 값은 FAQ 등록시 필요함
// @Accept  mpfd
// @Produce  mpfd
// @Security ApiKeyAuth
// @Param board_file formData file false "관련자료첨부"
// @Param param formData string true "json format"
// @Param test body model.Board false "test용 Json Data 실제사용 X"
// @Router /board [post]
// @Success 200
func BoardCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	data, eMsg := common.UnmarshalFormData(c, "param")
	if data == nil {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	board := model.Board{}
	if err := mapstructure.Decode(data, &board); nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		if 0 == board.Institution_seq {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_unauthorized)
			return
		} else {
			if common.ToUint(tokenMap["institution_seq"]) != board.Institution_seq {
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_unauthorized)
				return
			}
		}
	}

	board.User_seq = common.ToUint(tokenMap["user_seq"])
	board.File_idx = "1"
	if !board.InsertBoard(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Board
// @Summary	게시판 수정
// @Description 게시판 수정
// @Description institution_seq = 0 인거는 공통!
// @Description board_type = 1 공지사항
// @Description board_type = 2 자료실
// @Description board_type = 3 FAQ
// @Description file_delete 삭제시 true 아니면 flase
// @Accept  mpfd
// @Produce  mpfd
// @Security ApiKeyAuth
// @Param board_seq path uint true "board_seq"
// @Param board_file formData file false "관련자료첨부"
// @Param param formData string true "json format"
// @Param test body model.Board false "test용 Json Data 실제사용 X"
// @Router /board/{board_seq} [patch]
// @Success 200
func BoardPatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	data, eMsg := common.UnmarshalFormData(c, "param")
	if data == nil {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	board := model.Board{}
	if err := mapstructure.Decode(data, &board); nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	board_seq, succ, user_seq, _ := getBoardInfoFromPath(c)
	if !succ {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		if user_seq != tokenMap["user_seq"] {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_unauthorized)
			return
		}
	}

	board.Board_seq = common.ToUint(board_seq)
	board.Chg_user_seq = common.ToUint(tokenMap["user_seq"])
	if !board.UpdateBoard(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Board
// @Summary 게시판 삭제
// @Description 게시판 삭제
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param board_seq path uint true "board_seq"
// @Router /board/{board_seq} [delete]
// @Success 200
func BoardDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	board_seq, succ, user_seq, _ := getBoardInfoFromPath(c)
	if !succ {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_INSTITUTION {
		if user_seq != tokenMap["user_seq"] {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_unauthorized)
			return
		}
	}

	board := model.Board{}
	board.Board_seq = common.ToUint(board_seq)
	if !board.DeleteBoard(c) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}
