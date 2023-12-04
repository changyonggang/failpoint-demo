/*
 * Copyright (c) 2020 KingSoft.com, Inc. All Rights Reserved
 *
 * @file main
 * @author changyonggang(changyonggang@kingsoft.com)
 * @date 2023/12/4 19:29
 * @brief
 *
 */

package main

import (
	"fmt"
	"github.com/pingcap/failpoint"
)

func test() {
	failpoint.Inject("testValue", func(v failpoint.Value) {
		fmt.Println(v)
	})
}

func main() {
	for i := 0; i < 100; i++ {
		test()
	}
}
