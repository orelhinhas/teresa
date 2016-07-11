package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTeamDetailParams creates a new GetTeamDetailParams object
// with the default values initialized.
func NewGetTeamDetailParams() *GetTeamDetailParams {
	var ()
	return &GetTeamDetailParams{}
}

/*GetTeamDetailParams contains all the parameters to send to the API endpoint
for the get team detail operation typically these are written to a http.Request
*/
type GetTeamDetailParams struct {

	/*TeamID
	  Team ID

	*/
	TeamID int64
}

// WithTeamID adds the teamId to the get team detail params
func (o *GetTeamDetailParams) WithTeamID(TeamID int64) *GetTeamDetailParams {
	o.TeamID = TeamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}