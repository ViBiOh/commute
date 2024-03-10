package wahoo

import "time"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type WorkoutsResponse struct {
	Workouts []Workout `json:"workouts"`
}

type Workout struct {
	Starts        time.Time `json:"starts"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PlanID        any       `json:"plan_id"`
	Name          string    `json:"name"`
	WorkoutToken  string    `json:"workout_token"`
	ID            int       `json:"id"`
	Minutes       int       `json:"minutes"`
	WorkoutTypeID int       `json:"workout_type_id"`
}
