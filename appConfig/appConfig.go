package appConfig
import (
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	SugarAPI struct {
		CacheDir     string `json:"cacheDir"`
		UserName     string `json:"userName"`
		UserPassword string `json:"userPassword"`
		RestUrl      string `json:"restUrl"`
	} `json:"sugarAPI"`

	TwitterAPI struct {
		CacheDir          string `json:"cacheDir"`
		ConsumerKey       string `json:"consumerKey"`
		ConsumerSecret    string `json:"consumerSecret"`
		AccessToken       string `json:"accessToken"`
		AccessTokenSecret string `json:"accessTokenSecret"`
	} `json:"twitterAPI"`

	GoogleAPI struct {
		CertificateEmail              string `json:"certificateEmail"`
		CertificateEmailNew           string `json:"certificateEmailNew"`
		AnalyticsScope                string `json:"analyticsScope"`
		KeyFile                       string `json:"keyFile"`
		KeyFileNew                    string `json:"keyFileNew"`
		SpreadsheetId                 string `json:"spreadsheetId"`
		SpreadsheetId2                string `json:"spreadsheetId2"`
		WorksheetTitleAnalytic        string `json:"worksheetTitleAnalytic"`
		WorksheetTitleWP              string `json:"worksheetTitleWP"`
		WorksheetTitleBlog            string `json:"worksheetTitleBlog"`
		WorksheetTitleAltoroslabsBlog string `json:"worksheetTitleAltoroslabsBlog"`
		WorksheetTitleSMM             string `json:"worksheetTitleSMM"`
		WorksheetTitleRawData         string `json:"worksheetTitleRawData"`
		WorksheetTitleCFLive          string `json:"worksheetTitleCFLive"`
		Blog_url_count                string `json:"blog_url_count"`
	} `json:"googleApi"`

	BingAPI struct {
		CacheDir        string `json:"cacheDir"`
		CacheExtractDir string `json:"cacheExtractDir"`
		UserName        string `json:"userName"`
		Password        string `json:"password"`
		DeveloperToken  string `json:"developerToken"`
		CustomerId      string `json:"customerId"`
		AccountId       string `json:"accountId"`
		ClientId        string `json:"clientId"`
		ClientSecret    string `json:"clientSecret"`
		RedirectUri     string `json:"redirectUri"`
	} `json:"bingAPI"`

	BingAPINorway struct {
		CacheDir        string `json:"cacheDir"`
		CacheExtractDir string `json:"cacheExtractDir"`
		UserName        string `json:"userName"`
		Password        string `json:"password"`
		DeveloperToken  string `json:"developerToken"`
		CustomerId      string `json:"customerId"`
		AccountId       string `json:"accountId"`
		ClientId        string `json:"clientId"`
		ClientSecret    string `json:"clientSecret"`
		RedirectUri     string `json:"redirectUri"`
	} `json:"bingAPINorway"`

	OktopostAPI struct {
		CacheDir string `json:"cacheDir"`
		User     string `json:"user"`
		Key      string `json:"key"`
	} `json:"oktopostAPI"`

	FacebookAPI struct {
		App_id     string `json:"app_id"`
		App_secret string `json:"pp_secret"`
	} `json:"facebookAPI"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

