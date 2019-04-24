/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */package kinds

import (
	"github.com/creativesoftwarefdn/weaviate/entities/models"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/crossref"
	"github.com/go-openapi/strfmt"
)

// BatchAction is a helper type that groups all the info about one action in a
// batch that belongs together, i.e. uuid, action body and error state.
//
// Consumers of a BatchAction (i.e. database connector) should always check
// whether an error is already present by the time they receive a batch action.
// Errors can be introduced at all levels, e.g. validation.
//
// However, error'd actions are not removed to make sure that the list in
// BatchActions matches the order and content of the incoming batch request
type BatchAction struct {
	OriginalIndex int
	Err           error
	Action        *models.Action
	UUID          strfmt.UUID
}

// BatchActions groups many BatchAction items together. The order matches the
// order from the original request. It can be turned into the expected response
// type using the .Response() method
type BatchActions []BatchAction

// BatchThing is a helper type that groups all the info about one thing in a
// batch that belongs together, i.e. uuid, thing body and error state.
//
// Consumers of a Thing (i.e. database connector) should always check
// whether an error is already present by the time they receive a batch thing.
// Errors can be introduced at all levels, e.g. validation.
//
// However, error'd things are not removed to make sure that the list in
// Things matches the order and content of the incoming batch request
type BatchThing struct {
	OriginalIndex int
	Err           error
	Thing         *models.Thing
	UUID          strfmt.UUID
}

// BatchThings groups many Thing items together. The order matches the
// order from the original request. It can be turned into the expected response
// type using the .Response() method
type BatchThings []BatchThing

// BatchReference is a helper type that groups all the info about one references in a
// batch that belongs together, i.e. from, to, original index and error state
//
// Consumers of a Thing (i.e. database connector) should always check
// whether an error is already present by the time they receive a batch thing.
// Errors can be introduced at all levels, e.g. validation.
//
// However, error'd things are not removed to make sure that the list in
// Things matches the order and content of the incoming batch request
type BatchReference struct {
	OriginalIndex int
	Err           error
	From          *crossref.RefSource
	To            *crossref.Ref
}

// BatchReferences groups many Reference items together. The order matches the
// order from the original request. It can be turned into the expected response
// type using the .Response() method
type BatchReferences []BatchReference

// // Response uses the information contained in every Reference (from, to, error)
// // to form the expected response for the Batching request.
// func (b BatchReferences) Response() []*models.BatchReferenceResponse {
// 	response := make([]*models.BatchReferenceResponse, len(b), len(b))
// 	for i, ref := range b {
// 		var errorResponse *models.ErrorResponse
// 		var reference models.BatchReference

// 		status := models.BatchReferenceResponseAO1ResultStatusSUCCESS
// 		if ref.Err != nil {
// 			errorResponse = errPayloadFromSingleErr(ref.Err)
// 			status = models.BatchReferenceResponseAO1ResultStatusFAILED
// 		} else {
// 			reference.From = strfmt.URI(ref.From.String())
// 			reference.To = strfmt.URI(ref.To.String())
// 		}

// 		response[i] = &models.BatchReferenceResponse{
// 			BatchReference: reference,
// 			Result: &models.BatchReferenceResponseAO1Result{
// 				Errors: errorResponse,
// 				Status: &status,
// 			},
// 		}
// 	}

// 	return response
// }