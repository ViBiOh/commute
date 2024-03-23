package strava

import "time"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Activity struct {
	StartDateLocal             time.Time `json:"start_date_local"`
	StartDate                  time.Time `json:"start_date"`
	WorkoutType                any       `json:"workout_type,omitempty"`
	LocationState              any       `json:"location_state"`
	LocationCity               any       `json:"location_city"`
	UploadIDStr                string    `json:"upload_id_str"`
	Type                       string    `json:"type"`
	SportType                  string    `json:"sport_type"`
	GearID                     string    `json:"gear_id"`
	ExternalID                 string    `json:"external_id"`
	Timezone                   string    `json:"timezone"`
	Visibility                 string    `json:"visibility"`
	Name                       string    `json:"name"`
	LocationCountry            string    `json:"location_country"`
	Map                        Map       `json:"map"`
	StartLatlng                []float64 `json:"start_latlng"`
	EndLatlng                  []float64 `json:"end_latlng"`
	Athlete                    Athlete   `json:"athlete"`
	ElevLow                    float64   `json:"elev_low"`
	TotalElevationGain         float64   `json:"total_elevation_gain"`
	CommentCount               int       `json:"comment_count"`
	AthleteCount               int       `json:"athlete_count"`
	PhotoCount                 int       `json:"photo_count"`
	AchievementCount           int       `json:"achievement_count"`
	WeightedAverageWatts       int       `json:"weighted_average_watts,omitempty"`
	MaxWatts                   int       `json:"max_watts,omitempty"`
	AverageCadence             float64   `json:"average_cadence,omitempty"`
	AverageTemp                int       `json:"average_temp,omitempty"`
	UtcOffset                  float64   `json:"utc_offset"`
	MaxHeartrate               float64   `json:"max_heartrate,omitempty"`
	ID                         int       `json:"id"`
	Distance                   float64   `json:"distance"`
	ElapsedTime                int       `json:"elapsed_time"`
	AverageSpeed               float64   `json:"average_speed"`
	MaxSpeed                   float64   `json:"max_speed"`
	AverageWatts               float64   `json:"average_watts,omitempty"`
	Kilojoules                 float64   `json:"kilojoules,omitempty"`
	KudosCount                 int       `json:"kudos_count"`
	AverageHeartrate           float64   `json:"average_heartrate,omitempty"`
	TotalPhotoCount            int       `json:"total_photo_count"`
	PrCount                    int       `json:"pr_count"`
	ElevHigh                   float64   `json:"elev_high"`
	ResourceState              int       `json:"resource_state"`
	UploadID                   int64     `json:"upload_id"`
	MovingTime                 int       `json:"moving_time"`
	DeviceWatts                bool      `json:"device_watts,omitempty"`
	FromAcceptedTag            bool      `json:"from_accepted_tag"`
	DisplayHideHeartrateOption bool      `json:"display_hide_heartrate_option"`
	HeartrateOptOut            bool      `json:"heartrate_opt_out"`
	HasKudoed                  bool      `json:"has_kudoed"`
	HasHeartrate               bool      `json:"has_heartrate"`
	Flagged                    bool      `json:"flagged"`
	Private                    bool      `json:"private"`
	Manual                     bool      `json:"manual"`
	Commute                    bool      `json:"commute"`
	Trainer                    bool      `json:"trainer"`
}

type Athlete struct {
	ID            int `json:"id"`
	ResourceState int `json:"resource_state"`
}

type Map struct {
	ID              string `json:"id"`
	SummaryPolyline string `json:"summary_polyline"`
	ResourceState   int    `json:"resource_state"`
}
