/*
   Copyright (c) [2021] IT.SOS
   golibs is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package es

import (
	"sync"

	"github.com/bwcxyk/golibs/config"
	"github.com/elastic/go-elasticsearch/v7"
)

var esOnce sync.Once
var esNew *elasticsearch.Client

func NewEs() *elasticsearch.Client {
	esOnce.Do(func() {
		var err error
		cfg := elasticsearch.Config{
			Addresses: config.Config.GetEs(),
		}
		esNew, err = elasticsearch.NewClient(cfg)
		if err != nil {
			panic("es connect fail" + err.Error())
		}
	})
	return esNew
}
