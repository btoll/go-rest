// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// Event sport response (default view)
//
// Identifier: application/nmgapi.evententity; view=default
type EventMedia struct {
	EndDtTm time.Time `datastore:"endDtTm,noindex" json:"endDtTm,omitempty"`
	// Not guaranteed to be unique
	EventID string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Location.defaultLoc.id
	LocationID string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// e.g., March Madness
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Sport ID
	SportID   string    `datastore:"sportId,noindex" json:"sportId,omitempty"`
	StartDtTm time.Time `datastore:"startDtTm,noindex" json:"startDtTm,omitempty"`
	SubTitle  string    `datastore:"subTitle,noindex" json:"subTitle,omitempty"`
	// EventAdvanceMethod.singleElimination || doubleElim || bestOf
	TeamAdvanceMethod string `datastore:"teamAdvanceMethod,noindex" json:"teamAdvanceMethod,omitempty"`
}

// Validate validates the EventMedia media type instance.
func (mt *EventMedia) Validate() (err error) {
	if mt.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sportId"))
	}
	if mt.EventID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eventId"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.SubTitle == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subTitle"))
	}

	if mt.LocationID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "locationId"))
	}
	if mt.TeamAdvanceMethod == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "teamAdvanceMethod"))
	}
	return
}

// Event sport response (tiny view)
//
// Identifier: application/nmgapi.evententity; view=tiny
type EventMediaTiny struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the EventMediaTiny media type instance.
func (mt *EventMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeEventMedia decodes the EventMedia instance encoded in resp body.
func (c *Client) DecodeEventMedia(resp *http.Response) (*EventMedia, error) {
	var decoded EventMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeEventMediaTiny decodes the EventMediaTiny instance encoded in resp body.
func (c *Client) DecodeEventMediaTiny(resp *http.Response) (*EventMediaTiny, error) {
	var decoded EventMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Game response (default view)
//
// Identifier: application/nmgapi.gameentity; view=default
type GameMedia struct {
	// Any public id for this game, not unique
	ExternalID string `datastore:"externalId,noindex" json:"externalId,omitempty"`
	// Favorite team id
	FavTeamID      string    `datastore:"favTeamId,noindex" json:"favTeamId,omitempty"`
	FinishedAtDtTm time.Time `datastore:"finishedAtDtTm,noindex" json:"finishedAtDtTm,omitempty"`
	// TeamGameStatus.preGame || tradeable || gameOn || ended
	GameStatus string `datastore:"gameStatus,noindex" json:"gameStatus,omitempty"`
	// Name of location
	Location string `datastore:"location,noindex" json:"location,omitempty"`
	// True GPS location
	LocationID string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// TeamGameStatus.eliminated
	LoserProgressStatus string `datastore:"loserProgressStatus,noindex" json:"loserProgressStatus,omitempty"`
	// 0
	OddsForFav float64   `datastore:"oddsForFav,noindex" json:"oddsForFav,omitempty"`
	PlayDtTm   time.Time `datastore:"playDtTm,noindex" json:"playDtTm,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`
	// Public free form name
	Title       string `datastore:"title,noindex" json:"title,omitempty"`
	UnderTeamID string `datastore:"underTeamId,noindex" json:"underTeamId,omitempty"`
	// Empty until game completed
	WinnerTeamID string `datastore:"winnerTeamId,noindex" json:"winnerTeamId,omitempty"`
}

// Validate validates the GameMedia media type instance.
func (mt *GameMedia) Validate() (err error) {
	if mt.FavTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "favTeamId"))
	}
	if mt.UnderTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "underTeamId"))
	}
	if mt.WinnerTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "winnerTeamId"))
	}
	if mt.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sportId"))
	}

	if mt.ExternalID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "externalId"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.LocationID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "locationId"))
	}
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if mt.GameStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "gameStatus"))
	}

	if mt.LoserProgressStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "loserProgressStatus"))
	}
	return
}

// Game response (tiny view)
//
// Identifier: application/nmgapi.gameentity; view=tiny
type GameMediaTiny struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the GameMediaTiny media type instance.
func (mt *GameMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeGameMedia decodes the GameMedia instance encoded in resp body.
func (c *Client) DecodeGameMedia(resp *http.Response) (*GameMedia, error) {
	var decoded GameMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeGameMediaTiny decodes the GameMediaTiny instance encoded in resp body.
func (c *Client) DecodeGameMediaTiny(resp *http.Response) (*GameMediaTiny, error) {
	var decoded GameMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Sport response (default view)
//
// Identifier: application/nmgapi.sportentity; view=default
type SportMedia struct {
	// Is in season?
	Active bool `datastore:"active,noindex" json:"active,omitempty"`
	// sport_icon
	BackgroundImageName string `datastore:"backgroundImageName,noindex" json:"backgroundImageName,omitempty"`
	// Tournament
	EventTerm string `datastore:"eventTerm,noindex" json:"eventTerm,omitempty"`
	// Game
	GameTerm string `datastore:"gameTerm,noindex" json:"gameTerm,omitempty"`
	// sport_icon
	IconName         string  `datastore:"iconName,noindex" json:"iconName,omitempty"`
	MaxPreSplitPrice float64 `datastore:"maxPreSplitPrice,noindex" json:"maxPreSplitPrice,omitempty"`
	// Sport name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
}

// Validate validates the SportMedia media type instance.
func (mt *SportMedia) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.GameTerm == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "gameTerm"))
	}
	if mt.EventTerm == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eventTerm"))
	}
	if mt.IconName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "iconName"))
	}
	if mt.BackgroundImageName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "backgroundImageName"))
	}
	return
}

// Sport response (tiny view)
//
// Identifier: application/nmgapi.sportentity; view=tiny
type SportMediaTiny struct {
	// Sport name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
}

// Validate validates the SportMediaTiny media type instance.
func (mt *SportMediaTiny) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeSportMedia decodes the SportMedia instance encoded in resp body.
func (c *Client) DecodeSportMedia(resp *http.Response) (*SportMedia, error) {
	var decoded SportMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeSportMediaTiny decodes the SportMediaTiny instance encoded in resp body.
func (c *Client) DecodeSportMediaTiny(resp *http.Response) (*SportMediaTiny, error) {
	var decoded SportMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Team sport response (default view)
//
// Identifier: application/nmgapi.teamentity; view=default
type TeamMedia struct {
	// Team Win-Loss Record
	CurrentWinRecord string `datastore:"currentWinRecord,noindex" json:"currentWinRecord,omitempty"`
	// Team Logo
	FullLogo string `datastore:"fullLogo,noindex" json:"fullLogo,omitempty"`
	// Sport HomeTown ID
	HomeTownID string `datastore:"homeTownId,noindex" json:"homeTownId,omitempty"`
	// Team Icon
	IconName string `datastore:"iconName,noindex" json:"iconName,omitempty"`
	// Team name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Team Nickname
	ShortName string `datastore:"shortName,noindex" json:"shortName,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`
}

// Validate validates the TeamMedia media type instance.
func (mt *TeamMedia) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sportId"))
	}
	if mt.HomeTownID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "homeTownId"))
	}
	if mt.ShortName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "shortName"))
	}
	if mt.CurrentWinRecord == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "currentWinRecord"))
	}
	if mt.IconName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "iconName"))
	}
	if mt.FullLogo == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fullLogo"))
	}
	return
}

// Team sport response (tiny view)
//
// Identifier: application/nmgapi.teamentity; view=tiny
type TeamMediaTiny struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the TeamMediaTiny media type instance.
func (mt *TeamMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeTeamMedia decodes the TeamMedia instance encoded in resp body.
func (c *Client) DecodeTeamMedia(resp *http.Response) (*TeamMedia, error) {
	var decoded TeamMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeTeamMediaTiny decodes the TeamMediaTiny instance encoded in resp body.
func (c *Client) DecodeTeamMediaTiny(resp *http.Response) (*TeamMediaTiny, error) {
	var decoded TeamMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Team sport response (default view)
//
// Identifier: application/nmgapi.teamopeningconfigentity; view=default
type TeamOpeningConfigMedia struct {
	BuyIncrementPrice  float64   `datastore:"buyIncrementPrice,noindex" json:"buyIncrementPrice,omitempty"`
	BuyIncrementQuan   int       `datastore:"buyIncrementQuan,noindex" json:"buyIncrementQuan,omitempty"`
	LiquidationFee     float64   `datastore:"liquidationFee,noindex" json:"liquidationFee,omitempty"`
	OpeningPrice       float64   `datastore:"openingPrice,noindex" json:"openingPrice,omitempty"`
	OpeningShares      int       `datastore:"openingShares,noindex" json:"openingShares,omitempty"`
	SellDecrementPrice float64   `datastore:"sellDecrementPrice,noindex" json:"sellDecrementPrice,omitempty"`
	SellDecrementQuan  int       `datastore:"sellDecrementQuan,noindex" json:"sellDecrementQuan,omitempty"`
	StartTradeDtTm     time.Time `datastore:"startTradeDtTm,noindex" json:"startTradeDtTm,omitempty"`
}

// Validate validates the TeamOpeningConfigMedia media type instance.
func (mt *TeamOpeningConfigMedia) Validate() (err error) {

	return
}

// Team sport response (tiny view)
//
// Identifier: application/nmgapi.teamopeningconfigentity; view=tiny
type TeamOpeningConfigMediaTiny struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the TeamOpeningConfigMediaTiny media type instance.
func (mt *TeamOpeningConfigMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeTeamOpeningConfigMedia decodes the TeamOpeningConfigMedia instance encoded in resp body.
func (c *Client) DecodeTeamOpeningConfigMedia(resp *http.Response) (*TeamOpeningConfigMedia, error) {
	var decoded TeamOpeningConfigMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeTeamOpeningConfigMediaTiny decodes the TeamOpeningConfigMediaTiny instance encoded in resp body.
func (c *Client) DecodeTeamOpeningConfigMediaTiny(resp *http.Response) (*TeamOpeningConfigMediaTiny, error) {
	var decoded TeamOpeningConfigMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
