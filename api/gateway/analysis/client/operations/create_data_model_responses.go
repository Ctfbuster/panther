// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// CreateDataModelReader is a Reader for the CreateDataModel structure.
type CreateDataModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDataModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDataModelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDataModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDataModelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDataModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDataModelCreated creates a CreateDataModelCreated with default headers values
func NewCreateDataModelCreated() *CreateDataModelCreated {
	return &CreateDataModelCreated{}
}

/*CreateDataModelCreated handles this case with default header values.

DataModel created successfully
*/
type CreateDataModelCreated struct {
	Payload *models.DataModel
}

func (o *CreateDataModelCreated) Error() string {
	return fmt.Sprintf("[POST /datamodel][%d] createDataModelCreated  %+v", 201, o.Payload)
}

func (o *CreateDataModelCreated) GetPayload() *models.DataModel {
	return o.Payload
}

func (o *CreateDataModelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataModelBadRequest creates a CreateDataModelBadRequest with default headers values
func NewCreateDataModelBadRequest() *CreateDataModelBadRequest {
	return &CreateDataModelBadRequest{}
}

/*CreateDataModelBadRequest handles this case with default header values.

Bad request
*/
type CreateDataModelBadRequest struct {
	Payload *models.Error
}

func (o *CreateDataModelBadRequest) Error() string {
	return fmt.Sprintf("[POST /datamodel][%d] createDataModelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDataModelBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDataModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataModelConflict creates a CreateDataModelConflict with default headers values
func NewCreateDataModelConflict() *CreateDataModelConflict {
	return &CreateDataModelConflict{}
}

/*CreateDataModelConflict handles this case with default header values.

DataModel with the given ID already exists
*/
type CreateDataModelConflict struct {
}

func (o *CreateDataModelConflict) Error() string {
	return fmt.Sprintf("[POST /datamodel][%d] createDataModelConflict ", 409)
}

func (o *CreateDataModelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDataModelInternalServerError creates a CreateDataModelInternalServerError with default headers values
func NewCreateDataModelInternalServerError() *CreateDataModelInternalServerError {
	return &CreateDataModelInternalServerError{}
}

/*CreateDataModelInternalServerError handles this case with default header values.

Internal server error
*/
type CreateDataModelInternalServerError struct {
}

func (o *CreateDataModelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datamodel][%d] createDataModelInternalServerError ", 500)
}

func (o *CreateDataModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
