package router

import (
	"strings"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
)

func ErrToDescription(err error) string{
	errStr := err.Error()
	eles := strings.Split(err.Error(), ":")
	return strings.Replace(errStr, eles[0], common.ErrToDescription[eles[0]], -1)
	
}