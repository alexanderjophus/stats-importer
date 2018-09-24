package main

import "time"

type LiveFeedResponse struct {
	Copyright string `json:"copyright"`
	GamePk    int    `json:"gamePk"`
	Link      string `json:"link"`
	MetaData  struct {
		Wait      int    `json:"wait"`
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	GameData struct {
		Game struct {
			Pk     int    `json:"pk"`
			Season string `json:"season"`
			Type   string `json:"type"`
		} `json:"game"`
		Datetime struct {
			DateTime    time.Time `json:"dateTime"`
			EndDateTime time.Time `json:"endDateTime"`
		} `json:"datetime"`
		Status struct {
			AbstractGameState string `json:"abstractGameState"`
			CodedGameState    string `json:"codedGameState"`
			DetailedState     string `json:"detailedState"`
			StatusCode        string `json:"statusCode"`
			StartTimeTBD      bool   `json:"startTimeTBD"`
		} `json:"status"`
		Teams struct {
			Away struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Link  string `json:"link"`
				Venue struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Link     string `json:"link"`
					City     string `json:"city"`
					TimeZone struct {
						ID     string `json:"id"`
						Offset int    `json:"offset"`
						Tz     string `json:"tz"`
					} `json:"timeZone"`
				} `json:"venue"`
				Abbreviation    string `json:"abbreviation"`
				TriCode         string `json:"triCode"`
				TeamName        string `json:"teamName"`
				LocationName    string `json:"locationName"`
				FirstYearOfPlay string `json:"firstYearOfPlay"`
				Division        struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Link string `json:"link"`
				} `json:"division"`
				Conference struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Link string `json:"link"`
				} `json:"conference"`
				Franchise struct {
					FranchiseID int    `json:"franchiseId"`
					TeamName    string `json:"teamName"`
					Link        string `json:"link"`
				} `json:"franchise"`
				ShortName       string `json:"shortName"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				FranchiseID     int    `json:"franchiseId"`
				Active          bool   `json:"active"`
			} `json:"away"`
			Home struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Link  string `json:"link"`
				Venue struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Link     string `json:"link"`
					City     string `json:"city"`
					TimeZone struct {
						ID     string `json:"id"`
						Offset int    `json:"offset"`
						Tz     string `json:"tz"`
					} `json:"timeZone"`
				} `json:"venue"`
				Abbreviation    string `json:"abbreviation"`
				TriCode         string `json:"triCode"`
				TeamName        string `json:"teamName"`
				LocationName    string `json:"locationName"`
				FirstYearOfPlay string `json:"firstYearOfPlay"`
				Division        struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Link string `json:"link"`
				} `json:"division"`
				Conference struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Link string `json:"link"`
				} `json:"conference"`
				Franchise struct {
					FranchiseID int    `json:"franchiseId"`
					TeamName    string `json:"teamName"`
					Link        string `json:"link"`
				} `json:"franchise"`
				ShortName       string `json:"shortName"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				FranchiseID     int    `json:"franchiseId"`
				Active          bool   `json:"active"`
			} `json:"home"`
		} `json:"teams"`
		Venue struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"venue"`
	} `json:"gameData"`
	LiveData struct {
		Plays struct {
			AllPlays []struct {
				Result struct {
					Event       string `json:"event"`
					EventCode   string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx            int       `json:"eventIdx"`
					EventID             int       `json:"eventId"`
					Period              int       `json:"period"`
					PeriodType          string    `json:"periodType"`
					OrdinalNum          string    `json:"ordinalNum"`
					PeriodTime          string    `json:"periodTime"`
					PeriodTimeRemaining string    `json:"periodTimeRemaining"`
					DateTime            time.Time `json:"dateTime"`
					Goals               struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
				} `json:"coordinates"`
				Players []struct {
					Player struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"player"`
					PlayerType string `json:"playerType"`
				} `json:"players,omitempty"`
				Team struct {
					ID      int    `json:"id"`
					Name    string `json:"name"`
					Link    string `json:"link"`
					TriCode string `json:"triCode"`
				} `json:"team,omitempty"`
			} `json:"allPlays"`
			ScoringPlays  []int `json:"scoringPlays"`
			PenaltyPlays  []int `json:"penaltyPlays"`
			PlaysByPeriod []struct {
				StartIndex int   `json:"startIndex"`
				Plays      []int `json:"plays"`
				EndIndex   int   `json:"endIndex"`
			} `json:"playsByPeriod"`
			CurrentPlay struct {
				Result struct {
					Event       string `json:"event"`
					EventCode   string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx            int       `json:"eventIdx"`
					EventID             int       `json:"eventId"`
					Period              int       `json:"period"`
					PeriodType          string    `json:"periodType"`
					OrdinalNum          string    `json:"ordinalNum"`
					PeriodTime          string    `json:"periodTime"`
					PeriodTimeRemaining string    `json:"periodTimeRemaining"`
					DateTime            time.Time `json:"dateTime"`
					Goals               struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
				} `json:"coordinates"`
			} `json:"currentPlay"`
		} `json:"plays"`
		Linescore struct {
			CurrentPeriod              int    `json:"currentPeriod"`
			CurrentPeriodOrdinal       string `json:"currentPeriodOrdinal"`
			CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining"`
			Periods                    []struct {
				PeriodType string    `json:"periodType"`
				StartTime  time.Time `json:"startTime"`
				EndTime    time.Time `json:"endTime"`
				Num        int       `json:"num"`
				OrdinalNum string    `json:"ordinalNum"`
				Home       struct {
					Goals       int    `json:"goals"`
					ShotsOnGoal int    `json:"shotsOnGoal"`
					RinkSide    string `json:"rinkSide"`
				} `json:"home"`
				Away struct {
					Goals       int    `json:"goals"`
					ShotsOnGoal int    `json:"shotsOnGoal"`
					RinkSide    string `json:"rinkSide"`
				} `json:"away"`
			} `json:"periods"`
			ShootoutInfo struct {
				Away struct {
					Scores   int `json:"scores"`
					Attempts int `json:"attempts"`
				} `json:"away"`
				Home struct {
					Scores   int `json:"scores"`
					Attempts int `json:"attempts"`
				} `json:"home"`
			} `json:"shootoutInfo"`
			Teams struct {
				Home struct {
					Team struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode      string `json:"triCode"`
					} `json:"team"`
					Goals        int  `json:"goals"`
					ShotsOnGoal  int  `json:"shotsOnGoal"`
					GoaliePulled bool `json:"goaliePulled"`
					NumSkaters   int  `json:"numSkaters"`
					PowerPlay    bool `json:"powerPlay"`
				} `json:"home"`
				Away struct {
					Team struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode      string `json:"triCode"`
					} `json:"team"`
					Goals        int  `json:"goals"`
					ShotsOnGoal  int  `json:"shotsOnGoal"`
					GoaliePulled bool `json:"goaliePulled"`
					NumSkaters   int  `json:"numSkaters"`
					PowerPlay    bool `json:"powerPlay"`
				} `json:"away"`
			} `json:"teams"`
			PowerPlayStrength string `json:"powerPlayStrength"`
			HasShootout       bool   `json:"hasShootout"`
			IntermissionInfo  struct {
				IntermissionTimeRemaining int  `json:"intermissionTimeRemaining"`
				IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed"`
				InIntermission            bool `json:"inIntermission"`
			} `json:"intermissionInfo"`
			PowerPlayInfo struct {
				SituationTimeRemaining int  `json:"situationTimeRemaining"`
				SituationTimeElapsed   int  `json:"situationTimeElapsed"`
				InSituation            bool `json:"inSituation"`
			} `json:"powerPlayInfo"`
		} `json:"linescore"`
		Decisions struct {
			Winner struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"winner"`
			Loser struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"loser"`
			FirstStar struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"firstStar"`
			SecondStar struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"secondStar"`
			ThirdStar struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"thirdStar"`
		} `json:"decisions"`
	} `json:"liveData"`
}
