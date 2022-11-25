package tnyuri

/*
 * URL
 * Structure of tiny url database
 *
 */
type URL struct {
	Id    int    `db:"id"`
	Url   string `db:"url"`
	Short string `db:"short"`
	User  string `db:"user"`
	Uid   string `db:"uid"`
	Time  string `db:"timestamp"`
}

/*
 * Stat
 * Structure of database table of statistics
 *
 * TODO:
 * Make more complex statistics:
 * 	- Add more statistics (ip ua...)
 * 	- split counter into unique | repeatable
 *
 */
type Stats struct {
	Id      int `db:"id"`
	Uid     int `db:"url_id"`
	Counter int `db:"counter"`
}

type Model interface {
	URL | Stats
}
