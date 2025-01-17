//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package test

import (
	"testing"

	"github.com/semi-technologies/weaviate/test/acceptance/helper"
)

// Helper function to get all the names of Object classes.
func GetObjectClassNames(t *testing.T) []string {
	resp, err := helper.Client(t).Schema.SchemaDump(nil, nil)
	var names []string

	// Extract all names
	helper.AssertRequestOk(t, resp, err, func() {
		for _, class := range resp.Payload.Classes {
			names = append(names, class.Class)
		}
	})

	return names
}

// Helper function to get all the names of Action classes.
// func GetActionClassNames(t *testing.T) []string {
// 	resp, err := helper.Client(t).Schema.SchemaDump(nil, nil)
// 	var names []string

// 	// Extract all names
// 	helper.AssertRequestOk(t, resp, err, func() {
// 		for _, class := range resp.Payload.Actions.Classes {
// 			names = append(names, class.Class)
// 		}
// 	})

// 	return names
// }
