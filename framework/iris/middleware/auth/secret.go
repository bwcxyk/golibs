/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package auth

import (
	"github.com/bwcxyk/golibs/framework/iris/services"
	"github.com/kataras/iris/v12"
)

func Secret(ctx iris.Context) {
	token := ctx.GetHeader("token")
	loginId, err := services.GetLoginId(token)
	if err != nil {
		ctx.StopWithError(iris.StatusUnauthorized, err)
		return
	}
	ctx.Values().Set("loginId", loginId)
	ctx.Next()
}
