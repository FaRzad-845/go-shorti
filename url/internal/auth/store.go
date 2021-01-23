package auth

import (
	"github.com/jmoiron/sqlx"
)

// Store defines an interface to persists hits and other data.
type Store interface {
	// NewTx creates a new transaction and panic on failure.
	NewTx() *sqlx.Tx

	// Commit commits given transaction and logs the error.
	Commit(*sqlx.Tx)

	// Rollback rolls back given transaction and logs the error.
	Rollback(*sqlx.Tx)

	// CreateUser create new user.
	CreateUser(User) error

	// FindOneByEmail find and return user by email.
	FindOneByEmail(string) (User, error)

	// VerifyEmail verifies user.
	VerifyEmail(*sqlx.Tx, string) error

	//// FindOneByEmail find user by email
	//FindOneByEmail(*sqlx.Tx, *VisitorTimeStats) error

	//// SaveLanguageStats saves LanguageStats.
	//SaveLanguageStats(*sqlx.Tx, *LanguageStats) error
	//
	//// SaveReferrerStats saves ReferrerStats.
	//SaveReferrerStats(*sqlx.Tx, *ReferrerStats) error
	//
	//// SaveOSStats saves OSStats.
	//SaveOSStats(*sqlx.Tx, *OSStats) error
	//
	//// SaveBrowserStats saves BrowserStats.
	//SaveBrowserStats(*sqlx.Tx, *BrowserStats) error
	//
	//// SaveScreenStats saves ScreenStats.
	//SaveScreenStats(*sqlx.Tx, *ScreenStats) error
	//
	//// SaveCountryStats saves CountryStats.
	//SaveCountryStats(*sqlx.Tx, *CountryStats) error
	//
	//// Session returns the hits session timestamp for given fingerprint and max age.
	//Session(sql.NullInt64, string, time.Time) time.Time
	//
	//// HitDays returns the distinct days with at least one hit.
	//HitDays(sql.NullInt64) ([]time.Time, error)
	//
	//// HitPaths returns the distinct paths for given day.
	//HitPaths(sql.NullInt64, time.Time) ([]string, error)
	//
	//// Paths returns the distinct paths for given time frame.
	//Paths(sql.NullInt64, time.Time, time.Time) ([]string, error)
	//
	//// CountVisitors returns the visitor count for given day.
	//CountVisitors(*sqlx.Tx, sql.NullInt64, time.Time) *Stats
	//
	//// CountVisitorsByPath returns the visitor count for given day, path, and if the platform should be included or not.
	//CountVisitorsByPath(*sqlx.Tx, sql.NullInt64, time.Time, string, bool) ([]VisitorStats, error)
	//
	//// CountVisitorsByPathAndHour returns the visitor count for given day and path grouped by hour of day.
	//CountVisitorsByHour(*sqlx.Tx, sql.NullInt64, time.Time) ([]VisitorTimeStats, error)
	//
	//// CountVisitorsByPathAndLanguage returns the visitor count for given day and path grouped by language.
	//CountVisitorsByPathAndLanguage(*sqlx.Tx, sql.NullInt64, time.Time, string) ([]LanguageStats, error)
	//
	//// CountVisitorsByPathAndReferrer returns the visitor count for given day and path grouped by referrer.
	//CountVisitorsByPathAndReferrer(*sqlx.Tx, sql.NullInt64, time.Time, string) ([]ReferrerStats, error)
	//
	//// CountVisitorsByPathAndOS returns the visitor count for given day and path grouped by operating system and operating system version.
	//CountVisitorsByPathAndOS(*sqlx.Tx, sql.NullInt64, time.Time, string) ([]OSStats, error)
	//
	//// CountVisitorsByPathAndBrowser returns the visitor count for given day and path grouped by browser and browser version.
	//CountVisitorsByPathAndBrowser(*sqlx.Tx, sql.NullInt64, time.Time, string) ([]BrowserStats, error)
	//
	//// CountVisitorsByLanguage returns the visitor count for given day grouped by language.
	//CountVisitorsByLanguage(*sqlx.Tx, sql.NullInt64, time.Time) ([]LanguageStats, error)
	//
	//// CountVisitorsByReferrer returns the visitor count for given day grouped by referrer.
	//CountVisitorsByReferrer(*sqlx.Tx, sql.NullInt64, time.Time) ([]ReferrerStats, error)
	//
	//// CountVisitorsByOS returns the visitor count for given day grouped by operating system.
	//CountVisitorsByOS(*sqlx.Tx, sql.NullInt64, time.Time) ([]OSStats, error)
	//
	//// CountVisitorsByBrowser returns the visitor count for given day grouped by browser.
	//CountVisitorsByBrowser(*sqlx.Tx, sql.NullInt64, time.Time) ([]BrowserStats, error)
	//
	//// CountVisitorsByScreenSize returns the visitor count for given day grouped by screen size (width and height).
	//CountVisitorsByScreenSize(*sqlx.Tx, sql.NullInt64, time.Time) ([]ScreenStats, error)
	//
	//// CountVisitorsByScreenClass returns the visitor count for given day grouped by screen class.
	//CountVisitorsByScreenClass(*sqlx.Tx, sql.NullInt64, time.Time) ([]ScreenStats, error)
	//
	//// CountVisitorsByCountryCode returns the visitor count for given day grouped by country code.
	//CountVisitorsByCountryCode(*sqlx.Tx, sql.NullInt64, time.Time) ([]CountryStats, error)
	//
	//// CountVisitorsByPlatform returns the visitor count for given day grouped by platform.
	//CountVisitorsByPlatform(*sqlx.Tx, sql.NullInt64, time.Time) *VisitorStats
	//
	//// CountVisitorsByPathAndMaxOneHit returns the visitor count for given day and optional path with a maximum of one hit.
	//// This returns the absolut number of hits without further page calls and is used to calculate the bounce rate.
	//CountVisitorsByPathAndMaxOneHit(*sqlx.Tx, sql.NullInt64, time.Time, string) int
	//
	//// ActiveVisitors returns the active visitor count for given duration.
	//ActiveVisitors(sql.NullInt64, time.Time) int
	//
	//// ActivePageVisitors returns the active visitors grouped by path for given duration.
	//ActivePageVisitors(sql.NullInt64, time.Time) ([]Stats, error)
	//
	//// Visitors returns the visitors for given time frame grouped by days.
	//Visitors(sql.NullInt64, time.Time, time.Time) ([]Stats, error)
	//
	//// VisitorHours returns the visitors for given time frame grouped by hour of day.
	//VisitorHours(sql.NullInt64, time.Time, time.Time) ([]VisitorTimeStats, error)
	//
	//// VisitorLanguages returns the visitors for given time frame grouped by language.
	//VisitorLanguages(sql.NullInt64, time.Time, time.Time) ([]LanguageStats, error)
	//
	//// VisitorReferrer returns the visitor count for given time frame grouped by referrer.
	//VisitorReferrer(sql.NullInt64, time.Time, time.Time) ([]ReferrerStats, error)
	//
	//// VisitorOS returns the visitor count for given time frame grouped by operating system.
	//VisitorOS(sql.NullInt64, time.Time, time.Time) ([]OSStats, error)
	//
	//// VisitorBrowser returns the visitor count for given time frame grouped by browser.
	//VisitorBrowser(sql.NullInt64, time.Time, time.Time) ([]BrowserStats, error)
	//
	//// VisitorPlatform returns the visitor count for given time frame grouped by platform.
	//VisitorPlatform(sql.NullInt64, time.Time, time.Time) *VisitorStats
	//
	//// VisitorScreenSize returns the visitor count for given time frame grouped by screen size (width and height).
	//VisitorScreenSize(sql.NullInt64, time.Time, time.Time) ([]ScreenStats, error)
	//
	//// VisitorScreenClass returns the visitor count for given time frame grouped by screen class.
	//VisitorScreenClass(sql.NullInt64, time.Time, time.Time) ([]ScreenStats, error)
	//
	//// VisitorCountry returns the visitor count for given time frame grouped by country code.
	//VisitorCountry(sql.NullInt64, time.Time, time.Time) ([]CountryStats, error)
	//
	//// PageVisitors returns the visitors for given path and time frame grouped by days.
	//PageVisitors(sql.NullInt64, string, time.Time, time.Time) ([]Stats, error)
	//
	//// PageReferrer returns the visitors for given path and time frame grouped by referrer.
	//PageReferrer(sql.NullInt64, string, time.Time, time.Time) ([]ReferrerStats, error)
	//
	//// PageLanguages returns the visitors for given path and time frame grouped by language.
	//PageLanguages(sql.NullInt64, string, time.Time, time.Time) ([]LanguageStats, error)
	//
	//// PageOS returns the visitors for given path and time frame grouped by operating system.
	//PageOS(sql.NullInt64, string, time.Time, time.Time) ([]OSStats, error)
	//
	//// PageBrowser returns the visitors for given path and time frame grouped by browser.
	//PageBrowser(sql.NullInt64, string, time.Time, time.Time) ([]BrowserStats, error)
	//
	//// PagePlatform returns the visitors for given path and time frame grouped by platform.
	//PagePlatform(sql.NullInt64, string, time.Time, time.Time) *VisitorStats
	//
	//// VisitorsSum returns the sum of the visitors, sessions, and bounces for given time frame and path.
	//// The path is optional.
	//VisitorsSum(sql.NullInt64, time.Time, time.Time, string) (*Stats, error)
}
